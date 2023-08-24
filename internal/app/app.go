package app

import (
	"avito-test/config"
	"avito-test/internal/databases/repos"
	"avito-test/internal/server"
	"avito-test/internal/usecase"
	"log"
	"os"
)

type App struct {
	config *config.Config
}

func (app *App) Run() error {
	db, err := app.connPostgres()
	if err != nil {
		return err
	}

	repositories := repos.New(db)
	myUseCase := usecase.NewUseCase(repositories)

	newServer := server.NewServer(*app.config, myUseCase)

	err = newServer.Run()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
	return nil
}

func NewApp(config *config.Config) *App {
	return &App{config: config}
}
