package config

import (
	"database/sql"
	"log"
	"os"
	"path/filepath" // Import necessário para manipular caminhos de arquivos

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// inicializa a conexão com o banco de dados

func SetupDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtém o caminho do banco de dados a partir da variável de ambiente
	PATH := os.Getenv("DATABASE_URL")

	dir := filepath.Dir(PATH)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errDir := os.MkdirAll(dir, 0755)
		if errDir != nil {
			log.Fatalf("Erro ao criar diretório do banco de dados: %v", errDir)
		}
	}

	// Abre a conexão com o banco de dados SQLite
	connection, errDb := sql.Open("sqlite3", PATH)
	if errDb != nil {
		log.Fatal("Error connecting to database")
		log.Fatal(errDb)
	}

	// Testa a conexão
	teste := connection.Ping()
	if teste != nil {
		log.Fatal("Error pinging database")
		log.Fatal(teste)
	}

	log.Println("Database connected successfully")
	return connection
}


