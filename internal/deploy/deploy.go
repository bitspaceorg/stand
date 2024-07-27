package deploy

import (
    "context"
    "fmt"
    "os"
    "os/exec"
    "strings"

    parser "github.com/bitspaceorg/STAND-FOSSHACK/internal/build-parser"
    "github.com/bitspaceorg/STAND-FOSSHACK/internal/runnable"
    "github.com/bitspaceorg/STAND-FOSSHACK/internal/runtime"
)

type DeployCallback func(message string, status bool)

func DeployGo(builPath string, cb DeployCallback) {
    var BuildConfig parser.NodeBuildConfig
    parser := parser.NewBuildFileParser(builPath)
    parser.Parse(&BuildConfig)
    if BuildConfig.Requirements.Language != "node" {
        cb("Only Node is supported", false)
    }
    r := runtime.NodeRuntimeInstaller{
        Home: BuildConfig.Project.Home, Version: BuildConfig.Requirements.Version,
    }
    err := r.Install()
    if err != nil {
        if !runtime.IsExitCode(3, err) {
            cb(fmt.Sprintf("[Error] :%v", err.Error()), false)
        }
    }

    cmd := exec.Command("n", BuildConfig.Requirements.Version)
    if err := cmd.Run(); err != nil {
        cb("Could not change node version", false)
    }

    for _, rawCmd := range BuildConfig.Build {
        cmds := strings.Split(rawCmd.Cmd, " ")
        buildCmd := exec.Command(cmds[0], cmds[1:]...)
        buildCmd.Dir = BuildConfig.Project.Home
        buildCmd.Stdout = os.Stdout
        buildCmd.Stderr = os.Stderr
        if err := buildCmd.Run(); err != nil {
            cb(fmt.Sprintf("[Error] :%v", rawCmd.Name), false)
        }
    }

    cfg := runnable.NewStandConfig(BuildConfig.Project.Name, BuildConfig.Run[0].Cmd, BuildConfig.Project.Home, BuildConfig.Project.LogDir)

    runner, err := runnable.NewStandRunner(context.Background(), cfg)
    if err != nil {
        cb(fmt.Sprintf("[Error] :%v", err.Error()), false)
    }
    runner.SetEnv(BuildConfig.GetEnv())
    cb("Build Successful", true)
    runner.Run()
}
