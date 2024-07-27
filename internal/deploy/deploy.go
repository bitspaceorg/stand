package deploy

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strings"

	parser "github.com/bitspaceorg/STAND/internal/build-parser"
	"github.com/bitspaceorg/STAND/internal/runnable"
	"github.com/bitspaceorg/STAND/internal/runtime"
)

func DeployGo(builPath string) {
	var BuildConfig parser.NodeBuildConfig
	parser := parser.NewBuildFileParser(builPath)
	parser.Parse(&BuildConfig)
	if BuildConfig.Requirements.Language != "node" {
		log.Fatalf("Its not node!")
	}
	r := runtime.NodeRuntimeInstaller{
		Home: BuildConfig.Project.Home, Version: BuildConfig.Requirements.Version,
	}
	err := r.Install()
	if err != nil {
		if !runtime.IsExitCode(3, err) {
			log.Fatalf("[Error] :%v", err.Error())
		}
	}

	cmd := exec.Command("n", BuildConfig.Requirements.Version)
	if err := cmd.Run(); err != nil {
		log.Fatalf("Could not switch version")
	}

	for _, rawCmd := range BuildConfig.Build {
		cmds := strings.Split(rawCmd.Cmd, " ")
		buildCmd := exec.Command(cmds[0], cmds[1:]...)
		buildCmd.Dir = BuildConfig.Project.Home
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		if err := buildCmd.Run(); err != nil {
			log.Fatalf("Error in %v", rawCmd.Name)
		}
	}

	cfg := runnable.NewStandConfig(BuildConfig.Project.Name, BuildConfig.Run[0].Cmd, BuildConfig.Project.Home, BuildConfig.Project.LogDir)

	runner, err := runnable.NewStandRunner(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Error:%v", err)
	}
	runner.SetEnv(BuildConfig.GetEnv())
	runner.Run()
}
