package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (app *App) connPostgres() (*sql.DB, error) {
	urlDB := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		app.config.Postgres.Host, app.config.Postgres.Port,
		app.config.Postgres.User, app.config.Postgres.Password,
		app.config.Postgres.DBName)

	db, err := sql.Open("postgres", urlDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connect to the db")

	return db, nil
}
