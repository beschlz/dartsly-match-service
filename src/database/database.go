package database

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigrateDB() {

	db, err := sql.Open("postgres", "postgres://postgres:example@localhost:5432/dartsly_match_service?sslmode=disable")

	fmt.Println(err)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	fmt.Println(err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)

	fmt.Println(err)
	m.Up()
}
