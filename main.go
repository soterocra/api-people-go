package main

import (
	"database/sql"
	"encoding/json" // import para lidar com json
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // usamos o '_' pois é uma dependência impplícita, usada pela 'database/sql'
	"github.com/joho/godotenv"
)

// Modelo de dados
// O detalhe abaixo: 'json:"..."' é o equivalente ao @JsonProperty("nome") no Java
// São 'Struct Tags', são metadados que dizem ao GO como esse campo se comporta em diferentes contextos.
// json:"nome" diz ao pacote `json` do Go: "Quando você for transformar esta struct em JSON, ou ler um JSON para esta struct, o campo `Nome` deve ser chamado de `nome`."
type Pessoa struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func main() {

	// --- Carrega do .env chaves e valores ---
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Erro ao carregar arquivo .env. Usando variáveis de ambiente (se existirem).")
	}

	// Ler as variáveis do ambiente (que o godotenv acabou de carregar)
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// --- Conexão com o Banco ---

	// 1. DSN (Data Source Name): Connection String do Go.
	// Formato: usuario:senha@tcp(host:portal)/nome_do_banco
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPass, dbName)

	// 2. Abrir "pool" de conexão
	// Ele ainda não abre, mas prepara o pool. Verifica se a 'dsn' é válida
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao preparar a conexão com o banco: %v", err)
	}

	// 'defer' é o 'finally' do Go. Garante que o pool será fechado quando a função 'main' terminar.
	defer db.Close()

	// 3. Testar a conexão (aqui é realmente onde conecta)
	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	log.Println("Conexão com o MySQL estabelecida com sucesso!")

	// --- Handlers HTTP ---

	// 1. Registras um "handler" para a rota raiz.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 2. Escreve a resposta para quem pediu
		fmt.Fprintf(w, "Olá, API!")
	})

	// 2. Rota /pessoas
	http.HandleFunc("/pessoas", func(w http.ResponseWriter, r *http.Request) {

		// Etapa 1: Só permitir método POST
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}

		// Etapa 2: Decodificar o JSON do corpo da requisição
		var p Pessoa
		// Em Java: ObjectMapper.readValue(r.getBody(), Pessoa.class)
		// Em Go, criamos um "Decoder" e mandamos ele "preencher" a nossa struct 'p'.
		// O '&p' passa a "referência" (ponteiro) da struct.

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Corpo do JSON inválido "+err.Error(), http.StatusBadRequest)
			return
		}

		// Etapa 3: Inserir pessoa no banco
		// Usamos '?' para evitar SQL Injection
		res, err := db.Exec("INSERT INTO pessoas (nome, email) VALUES (?, ?)", p.Nome, p.Email)
		if err != nil {
			// Se o e-mail for duplicado (UNIQUE), o erro vai cair aqui
			http.Error(w, "Erro ao inserir no banco: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Etapa 4: Pegar o ID que o MySQL acabou de gerar (AUTO_INCREMENT)
		id, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Errro ao buscar o ID gerado: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Etapa 5: Responder com sucesso
		p.ID = int(id) // Coloca o ID do banco na struct

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(p)

	})

	// 4. Inicia o servidor web
	// O  'nil' é o multiplexador padrão (onde registramos a rota)
	log.Println("Servidor escutando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
