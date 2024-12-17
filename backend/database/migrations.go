package database

import (
	"log"

	"imobiliaria_crm/backend/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations runs database migrations from the migrations folder.
func RunMigrations(cfg config.Config) {
	db := GetDB()
	driver, err := postgres.WithInstance(db.Conn(), &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://backend/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	err = m.Up() // Apply all migrations
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
