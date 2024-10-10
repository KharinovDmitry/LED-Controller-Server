package main

import (
	"DynamicLED/config"
	"DynamicLED/internal/app"
	"log"
)

func main() {
	cfg, err := config.ReadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	application := app.New(cfg)
	if err := application.Build(); err != nil {
		log.Fatal(err.Error())
	}
	if err := application.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
