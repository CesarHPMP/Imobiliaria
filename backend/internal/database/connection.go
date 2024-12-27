package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"imobiliaria_crm/backend/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBconn struct {
	PostgresDB *pgxpool.Pool
}

var db DBconn

// Connect initializes the database connection using the given configuration.
func Connect(cfg config.Config) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	var err error
	db.PostgresDB, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}

	// Check if the "users" table already exists
	checkTableQuery := `SELECT to_regclass('public.users')`
	var result *sql.NullString // Use *sql.NullString to handle potential NULL values

	err = db.PostgresDB.QueryRow(context.Background(), checkTableQuery).Scan(&result)
	if err != nil {
		fmt.Printf("Failed to check table existence: %v\n", err)
		return err
	}

	// If result is nil, the table doesn't exist
	if result == nil {
		err = db.CreateTables(cfg)
		if err != nil {
			log.Fatalf("Unable to create tables in the database: %v\n", err)
		}
	}

	log.Println("Connected to the database successfully")

	return nil
}

// GetDB returns the database connection pool.
func GetDB() DBconn {
	return db
}

// CloseDB closes the database connection pool.
func CloseDB() {
	db.PostgresDB.Close()
}
