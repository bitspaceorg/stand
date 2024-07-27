package parser

import "fmt"

type Command struct {
	Name string `yaml:"name" json:"name"`
	Cmd  string `yaml:"cmd" json:"cmd"`
<<<<<<< HEAD
}

type Env struct {
	Name  string `yaml:"name" json:"name"`
	Value string `yaml:"value" json:"value"`
=======
>>>>>>> 3fab022 ([Add] rest handler new project 🚧)
}

type Env struct {
	Name  string `yaml:"name" json:"name"`
	Value string `yaml:"value" json:"value"`
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
		Name   string `yaml:"name" json:"name"`
		Home   string `yaml:"home" json:"home"`
		LogDir string `yaml:"log" json:"log"`
	}
	Requirements struct {
		Language string `yaml:"language" json:"language"`
		Version  string `yaml:"version" json:"version"`
	}
	Build []Command `yaml:"build" json:"build"`

	Run []Command `yaml:"run" json:"run"`

	Envs []Env `yaml:"env" json:"env"`
}

func (c *NodeBuildConfig) GetEnv() []string {
	envs := []string{}
	for _, v := range c.Envs {
		envs = append(envs, fmt.Sprintf("%v=%v", v.Name, v.Value))
	}

	return envs
}
