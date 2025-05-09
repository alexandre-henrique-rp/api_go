package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// inicializa a conexão com o banco de dados

func SetupDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PATH := os.Getenv("DATABASE_URL")
	connection, errDb := sql.Open("sqlite3", PATH)
	if errDb != nil {
		log.Fatal("Error connecting to database")
		log.Fatal(errDb)
	}

	teste := connection.Ping()

	if teste != nil {
		log.Fatal("Error pinging database")
		log.Fatal(teste)
	}

	log.Println("Database connected successfully")

	return connection
}

