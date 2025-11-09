package handler

import (
	"api-people-go/domain"
	"api-people-go/service"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

// É na struct que se "guarda" as dependências
type PessoaHandler struct {
	// Esse é o "D" do SOLID (Inversão de Dependência), em Java seria o "@Autowired private PessoaRepository repo;"
	service service.PessoaService
}

// Factory (Construtor)
func NewPessoaHandler(service service.PessoaService) *PessoaHandler {
	return &PessoaHandler{
		service: service,
	}
}

// O método (implementação). Peternce à struct 'PessoaHandler'. (h *PessoaHandler) é o "this".
func (h *PessoaHandler) CreatePessoa(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var p domain.Pessoa
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Corpo do JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	// O Handler não sabe SQL, ele apenas invoca o repositório. Em java seria "Pessoa novaPessoa = this.repo.save(p)"
	novaPessoa, err := h.service.Create(p)
	if err != nil {
		http.Error(w, "Erro ao criar pessoa: "+err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(novaPessoa)

}

func (h *PessoaHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("ERRO: %s %s -> %v", r.Method, r.URL.Path, err)

	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Recurso não encontrado", http.StatusNotFound)
		return
	}

	// Aqui pode-se checar outros erros

	// Erro padrão (catch-all)
	http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)

}

func (h *PessoaHandler) GetPessoaByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	pessoa, err := h.service.FindByID(id)
	if err != nil {
		h.handleError(w, r, err) // Usando o helper de erros.
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pessoa)

}
