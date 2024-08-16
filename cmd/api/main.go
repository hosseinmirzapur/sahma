package main

import (
	"fmt"
	"log"
	"os"
	"sahma/internal/database"
	"sahma/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Register database
	err := database.Register()
	if err != nil {
		log.Fatalln(err)
	}

	// Migrate database models
	err = database.Migrate()
	if err != nil {
		log.Fatalln(err)
	}

	// Run server
	s := server.NewServer()
	port := os.Getenv("PORT")
	err = s.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
