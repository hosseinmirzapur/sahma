package main

import (
	"fmt"
	"os"
	"sahma/internal/config"
	"sahma/internal/database/adapters"
	"sahma/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Register database
	err := adapters.RegisterMysql()
	if err != nil {
		config.Logger().Fatalln(err)
	}

	// Migrate database models
	err = adapters.Migrate()
	if err != nil {
		config.Logger().Fatalln(err)
	}

	// Run server
	s := server.NewServer()
	port := os.Getenv("PORT")
	err = s.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		config.Logger().Panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
