package repository

import (
	"api-people-go/domain"
	"database/sql"
	"fmt"
)

// Interface, define o "contrato" que deve ser feito, não "como".
// Se fosse em java seria como 'public interface PessoaRepository'
type PessoaRepository interface {
	Create(pessoa domain.Pessoa) (domain.Pessoa, error)
	// Demandas métodos aqui (FindByID, FindAll, etc...)
}

// Essa struct implementa a interface, mas não diretamente. Veja a descrição dos demais itens.
type mysqlPessoaRepository struct {
	db *sql.DB // O "pool" de conexão
}

// Em GO não há construtores, esse método faz esse trabalho, atuando como uma Factory e retornando a interface
// É aqui que a injeção de dependência acontece, recebe o *sql.DB e o armazena na struct
func NewPessoaRepository(db *sql.DB) PessoaRepository {
	return &mysqlPessoaRepository{
		db: db,
	}
}

// Aqui fica a função acoplada na struct que implementa os métodos da interface, ficam separados em GO.
// Se atentar a ligação feita na primeira parte da func, onde foi dado um apelido que deve ser curto, ele representa o "this" do java.
func (r *mysqlPessoaRepository) Create(pessoa domain.Pessoa) (domain.Pessoa, error) {
	res, err := r.db.Exec("INSERT INTO pessoas (nome, email) VALUES (?, ?)", pessoa.Nome, pessoa.Email)
	if err != nil {
		return domain.Pessoa{}, fmt.Errorf("erro ao inserir no banco: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Pessoa{}, fmt.Errorf("erro ao buscar o ID gerado: %w", err)
	}

	pessoa.ID = int(id)
	return pessoa, nil
}
