package main

import (
	"github.com/joho/godotenv"
	"github.com/stivenviedman/go-mongo-starter/src/config"
	"github.com/stivenviedman/go-mongo-starter/src/database"
	"github.com/stivenviedman/go-mongo-starter/src/router"
)

func main() {
	// Load environment variable from .env
	godotenv.Load()

	// Register env variables in global config.Env object
	config.RegisterEnv()

	// Configure DB
	database.Init()
	defer database.Close()

	// Configure router
	router.Init()
}
