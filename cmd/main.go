package main

import (
	"fmt"
	"os"

	"github.com/bitspaceorg/STAND-FOSSHACK/gui"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api"
	"github.com/bitspaceorg/STAND-FOSSHACK/user"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli [command]")
		fmt.Println("Commands: init, start")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		User.CreateUser()
	case "start":
		if User.ValidateUser() {
			// TODO
			// 1. Check for dependencies
			// 2. Initialize GUI
			gui.Init()
			// 3. Invoke metrics API
			api.Init()
			// 4. Start process?
		}
	default:
		fmt.Println("Invalid command:", command)
	}
}
