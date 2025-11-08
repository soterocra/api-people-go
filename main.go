package main

import (
	"api-people-go/database"
	"api-people-go/handler"
	"api-people-go/repository"
	"api-people-go/router"
	"api-people-go/server"
	"log"
)

func main() {

	// 1. Conexão de dados
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	// Encerra em ordem LIFO os defer quando para o main. Tem alguma semelhança ao finally do Java.
	defer db.Close()

	// 2. Injeção de dependência "O LEGO"
	pessoaRepo := repository.NewPessoaRepository(db)
	pessoaHandler := handler.NewPessoaHandler(pessoaRepo)

	// 3. Camanda de Roteamento HTTP. Fábrica de rotas
	mux := router.NewRouter(pessoaHandler)

	// 4. Iniciar o servidor com o 'mux' personalizado
	srv := server.NewServer(":8080", mux)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
