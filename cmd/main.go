package main

import (
	"log"

	parser "github.com/bitspaceorg/STAND/internal/build-parser"
)

// import (
// 	"context"
// 	"log"
//
// 	"github.com/bitspaceorg/STAND/internal/runnable"
// )

// func main() {
//
// 	cfg := runnable.
// 		NewStandConfig("testProj", "env", "/tmp/stand", "/tmp/stand").
// 		SetLogCompression(false)
//
// 	runner, err := runnable.NewStandRunner(context.Background(), cfg)
//
// 	if err != nil {
// 		log.Fatalf("Error: %v", err)
// 	}
//
// 	runner.SetEnv("HELLO=WORLD", "TEST=test")
//
// 	if err := runner.Run(); err != nil {
// 		log.Fatalf("Error: %v", err)
// 	}
//
// }

func main() {
	p := parser.NewBuildFileParser("./internal/build-parser/test.yml")
	cfg := new(parser.PythonBuildConfig)
	err := p.Parse(cfg)

	if err != nil {
		log.Fatalf("Error while parsing %v", err.Error())
	}

	log.Printf("Data: %+v",cfg)
}
