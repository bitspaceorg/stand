package parser

import "fmt"

type Command struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
}

type Env struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// default grammar for python build file
type PythonBuildConfig struct {
	Project struct {
		Name string `yaml:"name"`
		Home string `yaml:"home"`
	}
	Requirements struct {
		Language string `yaml:"language"`
		Version  string `yaml:"version"`
	}
	Build []Command `yaml:"build"`

	Run []Command `yaml:"run"`

	Envs []Env `yaml:"env"`
}

// default grammar for python build file
type NodeBuildConfig struct {
	Project struct {
		Name   string `yaml:"name"`
		Home   string `yaml:"home"`
		LogDir string `yaml:"log"`
	}
	Requirements struct {
		Language string `yaml:"language"`
		Version  string `yaml:"version"`
	}
	Build []Command `yaml:"build"`

	Run []Command `yaml:"run"`

	Envs []Env `yaml:"env"`
}

func (c *NodeBuildConfig) GetEnv() []string {
	envs := []string{}
	for _, v := range c.Envs {
		envs = append(envs, fmt.Sprintf("%v=%v", v.Name, v.Value))
	}

	return envs
}
