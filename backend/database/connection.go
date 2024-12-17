package database

import (
	"context"
	"fmt"
	"log"

	"imobiliaria_crm/backend/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

// Connect initializes the database connection using the given configuration.
func Connect(cfg config.Config) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	var err error
	db, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}

	log.Println("Connected to the database successfully")
}

// GetDB returns the database connection pool.
func GetDB() *pgxpool.Pool {
	return db
}

// CloseDB closes the database connection pool.
func CloseDB() {
	db.Close()
}
