package main

import (
	"fmt"

	"imobiliaria_crm/backend/config"
)

func main() {

	// Get DB Config
	Config := config.LoadConfig()
	fmt.Println("Database Connection String:", config.GetEnv("DB_PASSWORD", Config.DBUser))

	// Here you would normally initialize your database connection
}
