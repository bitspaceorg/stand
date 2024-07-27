package parser

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// BuildFileParser decodes the config file for the
// project build
type BuildFileParser struct {

	//the data which is contained in the yml file
	Data map[interface{}]interface{}

	//path to the build file
	filePath string
}


func NewBuildFileParser(filePath string) *BuildFileParser {
	return &BuildFileParser{
		Data:     make(map[interface{}]interface{}),
		filePath: filePath,
	}
}

// It tries to prase the given yml|yaml  file
// into the interface which is passed in
func (p *BuildFileParser) Parse(data interface{}) error {
	file, err := os.ReadFile(p.filePath)
	if err != nil {
		return fmt.Errorf("Error opening the file: %v", err.Error())
	}

	err = yaml.Unmarshal(file, data)

	if err != nil {
		log.Fatalf("Error parsing file into interface : %v", err.Error())
	}

	return nil
}
