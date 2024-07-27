package runtime

import (
	"errors"
	"log"
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
	switch i.fetcher {
	case "curl":
		err = i.runCommand("curl -o nvm.sh https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh")
		if err != nil {
			return errors.New("Could not access curl!")
		}
	case "wget":
		err = i.runCommand("wget -o https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh")
		if err != nil {
			log.Println(err)
			return errors.New("Could not access wget!")
		}
	}

	commands := []Installstep{
		{cmd: []string{"chmod", "+x", "nvm.sh"}, err: "Could not change permission to execute"},
		{cmd: []string{"/bin/bash", "nvm.sh"}, err: "Could not run the instal script"},
		{cmd: []string{"sh", "-c", `[ -s "$HOME/.nvm/nvm.sh" ] && \. "$HOME/.nvm/nvm.sh" && nvm install ` + i.Version}, err: "Could not change permission to execute"},
		{cmd: []string{"sh", "-c", `[ -s "$HOME/.nvm/nvm.sh" ] && \. "$HOME/.nvm/nvm.sh" && nvm use ` + i.Version + ` && node --version`}, err: "Could not change permission to execute "},
	}
	for _, cmd := range commands {
		cmdR := exec.Command(cmd.cmd[0], cmd.cmd[1:]...)
		buf, err := cmdR.CombinedOutput()
		if err != nil {
			log.Println(cmd.err + err.Error())
		}
		log.Println("this is a output", string(buf))
	}
	return nil
}
