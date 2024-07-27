package runnable

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Implements the Runnable interface
type StandRunner struct {
	cmd     *exec.Cmd
	pty     *os.File
	logFile *lumberjack.Logger
}

func NewStandRunner(ctx context.Context, cfg *StandConfig) (*StandRunner, error) {
	args := strings.Split(cfg.CmdString, " ")
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	err := cfg.CreateDirectories()
	if err != nil {
		return nil, err
	}

	logger := &lumberjack.Logger{
		Filename:   cfg.LogDir + "/" + cfg.ProjectName + ".log",
		MaxSize:    10,
		MaxBackups: 2,
		MaxAge:     14,
		Compress:   true,
	}

	cmd.Stdout = logger
	cmd.Stderr = logger
	cmd.Dir = cfg.HomeDir

	return &StandRunner{
		cmd:     cmd,
		logFile: logger,
	}, nil
}

func (s *StandRunner) Env() []string {
	return s.cmd.Env
}

func (s *StandRunner) SetEnv(vars ...string) {
	s.cmd.Env = append(s.cmd.Env, vars...)
}

// Clean up Function
func (s *StandRunner) Flush() error {
	s.logFile.Close()
	return nil
}

func (s *StandRunner) Run() error {

	//clean up logic when the run ends
	defer func() {
		s.Flush()
	}()

	err := s.cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
