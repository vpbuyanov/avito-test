package main

import (
	"github.com/golang-migrate/migrate/v4"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@0.0.0.0:5432/avitoTest?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil {
		log.Fatal(err)
	}
}
