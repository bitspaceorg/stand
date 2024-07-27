package runnable

import (
	"os"
)

// Config to initialize a new StandRunner
type StandConfig struct {
	ProjectName string
	CmdString   string
	HomeDir     string
	LogDir      string
	Env         []string

	//log management config
	LogMaxAge      int64
	LogMaxSize     int64
	LogMaxBackups  int64
	LogCompression bool
}

// Creates the log and home directories of the project folder
// and the log folder
func (cfg *StandConfig) CreateDirectories() error {
	err := os.MkdirAll(cfg.HomeDir, os.ModeAppend|os.ModePerm)
	if err != nil {
		return err
	}
	err = os.MkdirAll(cfg.LogDir, os.ModeAppend|os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// gives the default config
func NewStandConfig(ProjectName, CmdString, HomeDir, LogDir string) *StandConfig {
	return &StandConfig{
		ProjectName: ProjectName,
		CmdString:   CmdString,
		HomeDir:     HomeDir,
		LogDir:      LogDir,
		Env:         []string{},
		LogMaxAge:   14,
		//max sise of the file in bytes
		LogMaxSize: 10,
		//number of backup files
		LogMaxBackups:  2,
		LogCompression: true,
	}
}

func (cfg *StandConfig) SetLogMaxAge(age int64) *StandConfig {
	cfg.LogMaxAge = age
	return cfg
}

func (cfg *StandConfig) SetLogMaxSize(size int64) *StandConfig {
	cfg.LogMaxSize = size
	return cfg
}

func (cfg *StandConfig) SetLogMaxBackups(x int64) *StandConfig {
	cfg.LogMaxBackups = x
	return cfg
}

func (cfg *StandConfig) SetLogCompression(compress bool) *StandConfig {
	cfg.LogCompression = compress
	return cfg
}
