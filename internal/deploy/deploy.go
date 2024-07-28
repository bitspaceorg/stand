package deploy

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	parser "github.com/bitspaceorg/STAND-FOSSHACK/internal/build-parser"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/runnable"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/runtime"
	"github.com/bitspaceorg/STAND-FOSSHACK/utils"
)

type DeployCallback func(message string, status bool)

func DeployGo(builPath string, cb DeployCallback) {
	var BuildConfig parser.NodeBuildConfig
	parser := parser.NewBuildFileParser(builPath)
	parser.Parse(&BuildConfig)
	projectFolder := fmt.Sprintf("%s/%s/", utils.ShadowFolder, BuildConfig.Project.Home)
	log.Printf("%+v", BuildConfig)
	if BuildConfig.Requirements.Language != "node" {
		log.Println("Only node is supported")
		cb("Only Node is supported", false)
		return
	}
	r := runtime.NodeRuntimeInstaller{
		Home: projectFolder, Version: BuildConfig.Requirements.Version,
	}
	err := r.Install()
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
		buildCmd.Dir = projectFolder + "/" + BuildConfig.Project.Name
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
	cb("Build Successful", true)
	err = runner.Run()
	if err != nil {
		log.Println("Error in running :", err)
	}
}
