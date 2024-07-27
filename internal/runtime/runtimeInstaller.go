package runtime

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Each step in the install of the runtime
type Installstep struct {
	cmd []string
	err string
}

type RuntimeInstaller interface {
	// installs the runtime of the specific language to the local repository
	Install()
}

type NodeRuntimeInstaller struct {
	fetcher string
	Home    string
	Version string
}

// runs the command and throws the error
// it will block until the command is completed with a exit code
func (i *NodeRuntimeInstaller) runCommand(cmdString string) error {
	cmds := strings.Split(cmdString, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Dir = i.Home

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// check for the installation of fetch tool
func (i *NodeRuntimeInstaller) GetFetcher() error {
	toRun := []string{"which curl", "which wget"}
	for _, v := range toRun {
		err := i.runCommand(v)
		if err == nil {
			i.fetcher = strings.Split(v, " ")[1]
			return nil
		}
	}
	return errors.New("Your majesty, wget or curl is required for this peasant!")
}

func (i *NodeRuntimeInstaller) Install() error {
	var err error
	err = i.GetFetcher()
	if err != nil {
		return err
	}
	err = i.runCommand(i.fetcher + " -L -o node_install.sh https://bit.ly/n-install")
	if err != nil {
		log.Printf("Fetcher error: %v", err.Error())
		return errors.New("Could not access " + i.fetcher)
	}

	err = i.runCommand("chmod +x node_install.sh")
	if err != nil {
		return errors.New("could not give permission")
	}
	err = i.runCommand("bash node_install.sh -y")
	if err != nil {
		log.Println(err)
		return errors.New("Error could not install !")
	}
	return nil
}
