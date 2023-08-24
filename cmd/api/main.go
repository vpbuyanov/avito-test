package main

import (
	"avito-test/config"
	"avito-test/internal/app"
	"log"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	application := app.NewApp(cfg)
	err = application.Run()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
