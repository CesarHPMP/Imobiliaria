package database

import (
	"context"
	"errors"
	"fmt"
	"log"

	"imobiliaria_crm/backend/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	stdpgx "github.com/jackc/pgx/v4/stdlib"
)

// RunMigrations runs database migrations from the migrations folder.
func (db DBconn) RunMigrations(cfg config.Config) error {

	if db.PostgresDB == nil {
		return errors.New("database connection pool is not initialized")
	}

	sqlDB := stdpgx.OpenDB(*db.PostgresDB.Config().ConnConfig, nil)

	// Create the Postgres migration driver
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}

	// Initialize the migrator
	m, err := migrate.NewWithDatabaseInstance(
		"file://backend/migrations", // Path to migration files
		cfg.DBName,                  // Database name
		driver,
	)
	if err != nil {
		return err
	}

	// Apply all migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migrations applied successfully")

	return nil
}

// CreateTables creates tables in the database
func (db *DBconn) CreateTables(cfg config.Config) error {
	if db.PostgresDB == nil {
		return errors.New("database connection pool is not initialized")
	}

	// Start a transaction
	transaction, err := db.PostgresDB.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	// Ensure rollback if something goes wrong
	defer transaction.Rollback(context.Background())

	// Example SQL query to create a table
	createTableQuery := `
        CREATE TABLE users (
        id BIGSERIAL PRIMARY KEY,
        nome VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        endereco TEXT,
        numero VARCHAR(20),
        senha_hash TEXT NOT NULL
    );
    `

	// Execute the query
	_, err = transaction.Exec(context.Background(), createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	// Commit the transaction
	if err := transaction.Commit(context.Background()); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	log.Println("Tables created successfully")
	err = db.RunMigrations(cfg)
	if err != nil {
		return err
	}
	return nil
}
