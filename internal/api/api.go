package api

import (
	"github.com/bitspaceorg/STAND-FOSSHACK/config"
	"log"
)

func Init() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("Error setting up env : %v", err)
	}
	log.Printf("[Running GUI on port %v]", cfg.ServerPort)
	StartServer(cfg)
}
