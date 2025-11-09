package service

import (
	"api-people-go/domain"
	"database/sql"
	"errors"
	"testing"
)

// Mock Repository
// Implementação "falsa" de "PessoaRepository"
type mockPessoaRepository struct {

	// Guardar funções para o teste definir como simular o comportamento
	onCreate   func(pessoa domain.Pessoa) (domain.Pessoa, error)
	onFindById func(id int) (domain.Pessoa, error)
}

// Implementando a interface
// Faz o mínimo para satisfazer a interface PessoaRepository, apenas chamando os métodos guardados.
func (m *mockPessoaRepository) Create(pessoa domain.Pessoa) (domain.Pessoa, error) {
	// Se o teste definiu uma função 'onCreate', chame-a
	if m.onCreate != nil {
		return m.onCreate(pessoa)
	}
	// Se não retorne um erro padrão
	return domain.Pessoa{}, errors.New("onCreate não foi implementado pelo mock")
}

func (m *mockPessoaRepository) FindByID(id int) (domain.Pessoa, error) {
	// Se o testes definiu uma função 'onFindByID', chame-a.
	if m.onFindById != nil {
		return m.onFindById(id)
	}
	// Se não, retorne um erro padrão
	return domain.Pessoa{}, errors.New("onFindByID não foi implementado pelo mock")
}

// Os testes

// Teste 1: Caminho feliz do Create
func TestPessoaService_Create_Success(t *testing.T) {
	// ARRANGE
	// 1. A pessoa que esperamos receber do banco
	pessoaEsperada := domain.Pessoa{ID: 1, Nome: "Ada", Email: "ada@email.com"}

	// 2. Cria o mock
	mockRepo := &mockPessoaRepository{
		// 3. Setup do mock para quando 'onCreate' for chamado, retorna o 'pessoaEsperada'
		onCreate: func(pessoa domain.Pessoa) (domain.Pessoa, error) {
			return pessoaEsperada, nil
		},
	}

	// 4. Cria o service injetando o mock
	service := NewPessoaService(mockRepo)

	// ACT
	pessoaInput := domain.Pessoa{Nome: "Ada", Email: "ada@email.com"}
	pessoaCriada, err := service.Create(pessoaInput)

	// ASSERT
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if pessoaCriada.ID != pessoaEsperada.ID {
		t.Errorf("ID esperado era %d, mas recebi %d", pessoaEsperada.ID, pessoaCriada.ID)
	}

}

// Teste 2: Testando a validação de nome vazio
func TestPessoaService_Create_Error_Validation(t *testing.T) {
	// ARRANGE
	// Como a regra de negócio vai falhar antes de chamar o mock, não precisa criar ele completamente.
	mockRepo := &mockPessoaRepository{}
	service := NewPessoaService(mockRepo)

	// ACT
	pessoaInput := domain.Pessoa{Nome: "", Email: "ada@email.com"} // nome vazio
	_, err := service.Create(pessoaInput)

	// ASSERT
	if err == nil {
		t.Errorf("Esperava um erro de validação, mas não recebi nenhum")
	}
	if err.Error() != "nome e email são obrigatórios" {
		t.Errorf("Mensagem de erro inesperada: %v", err)
	}
}

// Teste 3: Caminho feliz do FindByID
func TestPessoaService_FindByID_Success(t *testing.T) {
	// ARRANGE
	pessoaEsperada := domain.Pessoa{ID: 1, Nome: "Ada", Email: "ada@email.com"}
	mockRepo := &mockPessoaRepository{
		onFindById: func(id int) (domain.Pessoa, error) {
			if id == 1 {
				return pessoaEsperada, nil
			}
			return domain.Pessoa{}, sql.ErrNoRows // simular não encontrado
		},
	}
	// não precisa do prefixo service.NewPessoaService porque o teste está no mesmo package.
	service := NewPessoaService(mockRepo)

	// ACT
	pessoaEncontrada, err := service.FindByID(1)

	// ASSERT
	if err != nil {
		t.Errorf("Erro inesperado :%v", err)
	}
	if pessoaEncontrada.ID != pessoaEsperada.ID {
		t.Errorf("Pessoa errada encontrada")
	}

}

// Teste 4: Testando o FindByID não encontrado
func TestPessoaService_FindByID_NotFound(t *testing.T) {
	// ARRANGE
	mockRepo := &mockPessoaRepository{
		onFindById: func(id int) (domain.Pessoa, error) {
			return domain.Pessoa{}, sql.ErrNoRows
		},
	}
	service := NewPessoaService(mockRepo)

	// ACT
	_, err := service.FindByID(99) // Busca id que não existe

	// ASSERT
	if err == nil {
		t.Errorf("Esperava um erro de 'não encontrado', mas não recebi")
	}

	// Verifica se recebeu o erro correto
	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("Tipo de erro inesperado: %v", err)
	}
}
