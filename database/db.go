package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // O driver é necessário aqui
	"github.com/joho/godotenv"
)

// Connect será nossa função "public" (dado que começa com maiúscula)
// Ela retorna o "pool" de conexões (*sql.DB) ou um erro.
func Connect() (*sql.DB, error) {

	// 1. Carregar o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Erro ao carregar arquivo .env")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	if dbName == "" {
		dbName = "empty"
	}
	if dbUser == "" {
		dbUser = "empty"
	}
	if dbPass == "" {
		dbPass = "empty"
	}

	// 2. Construir DSN
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPass, dbName)

	// 3. Abrir conexão
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// ao invés de apenas logar, retornamos o erro
		return nil, fmt.Errorf("Erro ao preparar a conexão: %w", err)
	}

	// 4. Testar a conexão
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Erro ao conectar com o banco: %w", err)
	}

	log.Println("Conexão com o MySQL estabelecida com sucesso!")

	// Retorna o 'db' (pool) pronto para ser usado
	return db, nil

}
