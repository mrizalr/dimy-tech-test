package cmd

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/mrizalr/dimy-tech-test/internal/server"
	"github.com/mrizalr/dimy-tech-test/pkg/db/mysql"
)

func StartApplication() {
	// Loading .env files
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error when loading .env files | err = %v", err)
	}

	// Create db connection & auto migrating tables
	db, err := mysql.New()
	if err != nil {
		log.Fatalf("Error when connecting to database | err = %v", err)
	}
	err = db.AutoMigrate()
	if err != nil {
		log.Fatalf("Error wheh auto migrating database | err = %v", err)
	}

	// Create a new server
	s := server.New(db)

	// Running the server
	err = s.Run()
	if err != nil {
		log.Fatalf("Error when running the server | err = %v", err)
	}

}
