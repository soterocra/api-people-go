package main

import (
	"api-people-go/config"
	"api-people-go/database"
	"api-people-go/handler"
	"api-people-go/repository"
	"api-people-go/router"
	"api-people-go/server"
	"api-people-go/service"
	"log"
)

func main() {

	// 0. Configuração
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 1. Conexão de dados
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// Encerra em ordem LIFO os defer quando para o main. Tem alguma semelhança ao finally do Java.
	defer db.Close()

	// 2. Injeção de dependência "O LEGO"
	pessoaRepo := repository.NewPessoaRepository(db)
	pessoaService := service.NewPessoaService(pessoaRepo)
	pessoaHandler := handler.NewPessoaHandler(pessoaService)

	// 3. Camanda de Roteamento HTTP. Fábrica de rotas
	mux := router.NewRouter(pessoaHandler)

	// 4. Iniciar o servidor com o 'mux' personalizado
	srv := server.NewServer(cfg, mux)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
