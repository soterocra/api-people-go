package database

import (
	"api-people-go/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // O driver é necessário aqui
	"github.com/joho/godotenv"
)

// Connect será nossa função "public" (dado que começa com maiúscula)
// Ela retorna o "pool" de conexões (*sql.DB) ou um erro.
func Connect(cfg config.Config) (*sql.DB, error) {

	// 1. Carregar o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Erro ao carregar arquivo .env")
	}

	// 2. Construir DSN
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)

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
