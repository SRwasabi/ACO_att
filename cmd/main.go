package main

import (
	"github.com/SRwasabi/ACO_att/pkg/config"
)

func main() {
	cfg := config.Default()
	
	config.OverrideWithCLI(&cfg)

	cfg.Print()
}