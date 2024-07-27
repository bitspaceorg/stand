package gui

import (
	"fmt"
	"github.com/bitspaceorg/STAND-FOSSHACK/utils"
	"os"
	"os/exec"
	"strings"
)

// Sequentially initializes the GUI
func Init() {
	fmt.Println("Starting GUI...")

	// Checks if GUI already exists - if not fetches from the remote URL
	if _, err := os.Stat(utils.GetShadowGUIFilePath()); os.IsNotExist(err) {
		fmt.Println("GUI not found. Fetching from remote...")
		fetchCmds := strings.Split(fmt.Sprintf("git clone %v %v", utils.GUIFetchLink, utils.GetShadowGUIFilePath()), " ")
		fetchExec := exec.Command(fetchCmds[0], fetchCmds[1:]...)
		if err := fetchExec.Run(); err != nil {
			fmt.Println("Error fetching GUI:", err)
			return
		}
	}

	// Checks if the GUI is already built - if not yarn build it
	if _, yarnbuilderr := os.Stat(fmt.Sprintf("%v/node_modules", utils.GetShadowGUIFilePath())); os.IsNotExist(yarnbuilderr) {
		fmt.Println("Building GUI...")
		yarnBuildCmds := strings.Split(fmt.Sprintf("cd %v && yarn", utils.GetShadowGUIFilePath()), " ")
		yarnBuildExec := exec.Command("sh", "-c", strings.Join(yarnBuildCmds, " "))
		if err := yarnBuildExec.Run(); err != nil {
			fmt.Println("Error building GUI:", err)
			return
		}
	}
	fmt.Println("[Running GUI on port 6789]")

	// Start the GUI on port [6789]
	yarnCmds := strings.Split(fmt.Sprintf("cd %v && yarn dev", utils.GetShadowGUIFilePath()), " ")
	yarnExec := exec.Command("sh", "-c", strings.Join(yarnCmds, " "))
	if err := yarnExec.Run(); err != nil {
		fmt.Println("Error running GUI:", err)
		return
	}
}
