package database

import (
	"context"
	"fmt"
	"log"

	"imobiliaria_crm/backend/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBconn struct {
	postgresDB *pgxpool.Pool
}

var db DBconn

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
	db.postgresDB, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}

	// Check if the "users" table already exists
	checkTableQuery := `SELECT to_regclass('public.users')`
	var result string

	err = db.postgresDB.QueryRow(context.Background(), checkTableQuery).Scan(&result)

	if err != nil && err.Error() != "no rows in result set" {
		fmt.Printf("failed to check table existence: %v", err)
		return
	}

	// If table doesn't exist, create it
	if result == "" {
		err = db.CreateTables(cfg)
		if err != nil {
			log.Fatalf("Unable to create tables in the database: %v\n", err)
		}
	}

	log.Println("Connected to the database successfully")

}

// GetDB returns the database connection pool.
func GetDB() DBconn {
	return db
}

// CloseDB closes the database connection pool.
func CloseDB() {
	db.postgresDB.Close()
}
