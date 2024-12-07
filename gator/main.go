package main

import (
	"fmt"
	"log"

	"github.com/Osirous/Gator/internal/config"
)

const configFileName = ".gatorconfig.json"

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	fmt.Fprintln(cfg)
}
