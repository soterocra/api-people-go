package service

import (
	"api-people-go/domain"
	"api-people-go/repository"
	"fmt"
)

// A interface, define o que o serviço faz (contrato)
type PessoaService interface {
	Create(pessoa domain.Pessoa) (domain.Pessoa, error)
}

// A struct, a implementação. Letra minúscula por ser "privada".
// Aqui ele vai guardar as dependências.
type pessoaService struct {
	// O serviço depende da interface do repository
	repo repository.PessoaRepository
}

// A fábrica/construtor. Ela retorna a interface e esconde a implementação.
func NewPessoaService(repo repository.PessoaRepository) PessoaService {
	return &pessoaService{
		repo: repo,
	}
}

// O metodo, ele deve refletir o que é pedido na interface. O "s *pessoaService" é como se fosse o "this" do java.
func (s *pessoaService) Create(pessoa domain.Pessoa) (domain.Pessoa, error) {
	if pessoa.Nome == "" || pessoa.Email == "" {
		return domain.Pessoa{}, fmt.Errorf("nome e email são obrigatórios")
	}

	return s.repo.Create(pessoa)
}
