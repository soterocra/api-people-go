Executar banco de dados no Docker:

```shell
docker compose up -d
```

Para iniciar o módulo antes de instalar as dependências eu usei o seguinte comando:

```shell
go mod init api-people-go
```

Fazer o download do driver mysql:

```shell
go get github.com/go-sql-driver/mysql
```

Listar todas as dependências:

```shell
go list -m all
```

Entender o porque uma dependência está no projeto:

```shell
go mod why filippo.io/edwards25519
```


Comando usado para limpar e organizar o `go.mod`:

```shell
go mod tidy
```

`Create` da tabela:

```sql
CREATE TABLE pessoas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```