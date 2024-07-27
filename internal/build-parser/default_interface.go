package parser

type Command struct {
	Name string `yaml:"name"`
	Cmd  string `yaml:"cmd"`
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

	Ports struct {
		Expose string `yaml:"expose"`
	} `yaml:"ports"`
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

	Ports struct {
		Expose string `yaml:"expose"`
	} `yaml:"ports"`
}
