package main

import (
	"log"
	"melvad/internal/app"
	"melvad/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	err = app.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
