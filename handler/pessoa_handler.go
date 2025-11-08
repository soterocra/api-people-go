package handler

import (
	"api-people-go/domain"
	"api-people-go/service"
	"encoding/json"
	"net/http"
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
