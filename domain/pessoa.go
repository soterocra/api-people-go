package domain // O nome do pacote é o nome da pasta

// Em Go, para algo ser "public" (visivel fora do pacote) o nome deve começar com letra maiúscula.
// (Pense em 'Pessoa' como 'public class Pessoa')
type Pessoa struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
