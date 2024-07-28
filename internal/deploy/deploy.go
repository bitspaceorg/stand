package deploy

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"

	parser "github.com/bitspaceorg/STAND-FOSSHACK/internal/build-parser"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/runnable"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/runtime"
	"github.com/bitspaceorg/STAND-FOSSHACK/utils"
)

type DeployCallback func(message string, status bool)

type Deployer struct {
	//interface to install runtime
	RuntimeInstaller runtime.RuntimeInstaller

	//map to hold the list of all deployments
	Deploymnets map[string]runnable.Runnable
}

var instance *Deployer
var once sync.Once

// GetInstance returns the singleton instance
func GetInstance(installer runtime.RuntimeInstaller) *Deployer {
	once.Do(func() {
		instance = &Deployer{
			RuntimeInstaller: installer,
			Deploymnets:      make(map[string]runnable.Runnable),
		}
	})
	return instance
}

func NewDeployer(installer runtime.RuntimeInstaller) Deployer {
	return Deployer{
		RuntimeInstaller: installer,
		Deploymnets:      make(map[string]runnable.Runnable),
	}
}

func (d *Deployer) Kill(name string) error {
	runner, ok := d.Deploymnets[name]
	if !ok {
		log.Println(d.Deploymnets)
		return errors.New("Process does not exist!")
	}
	if err := runner.Kill(); err != nil {
		log.Println("Here =====>",err)
		return err
	}
	delete(d.Deploymnets, name)
	return nil
}

func (d *Deployer) Deploy(builPath string, cb DeployCallback) {
	var BuildConfig parser.NodeBuildConfig
	parser := parser.NewBuildFileParser(builPath)
	parser.Parse(&BuildConfig)
	projectFolder := fmt.Sprintf("%s/%s/", utils.ShadowFolder, BuildConfig.Project.Name)
	log.Printf("%+v", BuildConfig)
	if BuildConfig.Requirements.Language != "node" {
		log.Println("Only node is supported")
		cb("Only Node is supported", false)
		return
	}

	d.RuntimeInstaller.SetHome(projectFolder)
	d.RuntimeInstaller.SetVersion(BuildConfig.Requirements.Version)

	err := d.RuntimeInstaller.Install()
	if err != nil {
		if !runtime.IsExitCode(3, err) {
			log.Println(fmt.Sprintf("[Error] :%v", err.Error()))
			cb(fmt.Sprintf("[Error] :%v", err.Error()), false)
			return
		}
	}

	cmdi := exec.Command("n", "i", BuildConfig.Requirements.Version)
	if err := cmdi.Run(); err != nil {
		log.Println("could not install node version")
		cb("Could not install node version", false)
		return
	}

	cmd := exec.Command("n", BuildConfig.Requirements.Version)
	if err := cmd.Run(); err != nil {
		log.Println("could not change node version")
		cb("Could not change node version", false)
		return
	}

	for _, rawCmd := range BuildConfig.Build {
		cmds := strings.Split(rawCmd.Cmd, " ")
		buildCmd := exec.Command(cmds[0], cmds[1:]...)
		buildCmd.Dir = projectFolder+BuildConfig.Project.Name
		// buildCmd.Stdout = os.Stdout
		// buildCmd.Stderr = os.Stderr
		if buf, err := buildCmd.CombinedOutput(); err != nil {
			log.Println(fmt.Sprintf("[Error]:%v", rawCmd.Cmd))
			log.Println("OUTPUT:", string(buf))
			cb(fmt.Sprintf("[Error] :%v", rawCmd.Name), false)
			return
		}
	}

	cfg := runnable.NewStandConfig(BuildConfig.Project.Name, BuildConfig.Run[0].Cmd, projectFolder+BuildConfig.Project.Name, BuildConfig.Project.LogDir)

	runner, err := runnable.NewStandRunner(context.Background(), cfg)
	if err != nil {
		log.Println(fmt.Sprintf("[Error] :%v", err.Error()))
		cb(fmt.Sprintf("[Error] :%v", err.Error()), false)
		return
	}
	runner.SetEnv(BuildConfig.GetEnv())
	log.Println("Build Successful!")

	d.Deploymnets[BuildConfig.Project.Name] = runner
	cb("Build Successful", true)
	err = runner.Run()
	if err != nil {
		log.Println("Error in running :", err)
	}
}
