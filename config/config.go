package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Database config
type Database struct {
	Host string
	Name string
}

var (
	DB      Database
	GinMode = "debug"
)

func init() {
	CurrentEnvironment, ok := os.LookupEnv("GO_ENV")

	var err error

	if !ok {
		CurrentEnvironment = "development"
		err = godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Environment: " + CurrentEnvironment)

	DB = setDatabaseConfig()
}

func setDatabaseConfig() Database {
	db := Database{
		Host: os.Getenv("MONGODB_HOST"),
		Name: os.Getenv("MONGODB_DATABASE"),
	}

	return db
}
