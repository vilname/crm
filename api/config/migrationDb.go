package config

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func InitMigrationDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbSql, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPass, dbName, dbPort),
	)

	driver, _ := postgres.WithInstance(dbSql, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"postgres", driver)

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		_ = fmt.Errorf("error when migration up: %v", err)
	}

	dbSql.Close()
}
