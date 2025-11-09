package router

import (
	"api-people-go/handler"
	"fmt"
	"net/http"
)

// NewRouter usa o padrão de "Fábrica" (igual ao repository)
// Ele recebe as dependÊncias (o handler) e devolve o roteador
func NewRouter(pessoaHandler *handler.PessoaHandler) *http.ServeMux {

	// 1. Criamos nosso próprio roteador (não o global)
	mux := http.NewServeMux()

	// 2. Registramos as rotas nele (mux.HandleFunc, não http.HandleFunc)
	mux.HandleFunc("POST /pessoas", pessoaHandler.CreatePessoa)

	// {id} é o parâmetro
	mux.HandleFunc("GET /pessoas/{id}", pessoaHandler.GetPessoaByID)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API de Pessoas v2 (Com Roteador separado!)")
	})

	// 3. Retorna o roteador pronto
	return mux

}
