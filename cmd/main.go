package main

import "github.com/bitspaceorg/STAND/internal/deploy"

func main() {
	go deploy.DeployGo("./internal/build-parser/test.yml")
	go deploy.DeployGo("./internal/build-parser/test2.yml")
	select {}
}
