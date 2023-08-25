package main

import (
	"avito-test/config"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	url := fmt.Sprintf("postgres://%s:%s@0.0.0.0:%v/%s?sslmode=disable",
		conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Port, conf.Postgres.DBName)
	m, err := migrate.New(
		"file://migrations",
		url)

	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil {
		log.Fatal(err)
	}
}
