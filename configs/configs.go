package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	App    *AppConfigs
	DbConn *PostgresConfigs
)

type AppConfigs struct {
	Port string `env:"PORT"`
	Env  string `env:"ENV"`
}

type PostgresConfigs struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DbName   string `env:"DB_NAME"`
}

func LoadConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	log.Println(".env inicializate")
	initializeApp()
	initializeDb()
}

func initializeApp() {
	App = &AppConfigs{
		Port: os.Getenv("PORT"),
		Env:  os.Getenv("ENV"),
	}
}

func initializeDb() {
	DbConn = &PostgresConfigs{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
	}
}
