package main

import "github.com/bitspaceorg/STAND-FOSSHACK/internal/deploy"

func main() {
	go deploy.DeployGo("./example/test.yml")
	select {}
}
