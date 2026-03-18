# Go Wallet Core API Service (Go)

Este projeto segue os princípios de **Domain-Driven Design (DDD)** e **Clean Architecture** para garantir uma base de código modular, testável e manutenível.

## Arquitetura do Projeto

A estrutura de diretórios foi organizada dentro da pasta `src/` para separar as responsabilidades em camadas:

- **src/cmd/app**: Ponto de entrada da aplicação. Responsável por inicializar as dependências e subir o servidor.
- **src/internal/domain**: O coração da aplicação. Contém as regras de negócio, entidades e interfaces (contratos) de repositórios.
- **src/internal/usecase**: Camada de aplicação que orquestra a execução das regras de negócio usando as entidades do domínio.
- **src/internal/infrastructure**: Implementações detalhadas como persistência no banco de dados (mocked), clientes de APIs externas, etc.
- **src/internal/interface**: Adaptadores de interface para expor a lógica interna (HTTP Handlers, gRPC, CLI).


## Como Executar

Se você tiver o Go instalado:

1.  Aponte para o diretório raiz do projeto.
2.  Assure-se de que o `go.mod` está configurado corretamente.
3.  Execute o comando:

```bash
go run src/cmd/app/main.go
```

## 📖 Documentação da API

A API segue padrões REST e possui documentação técnica disponível na pasta `api/`. 

- [OpenAPI (Swagger)](api/swagger.yaml)
- [REST Client (VSCode)](api/requests.http)

### Endpoints Disponíveis

#### **1. Health Check**
`GET /health`
Verifica se o servidor está online.
```bash
curl.exe -X GET "http://localhost:8080/health"
```

#### **2. Criar Usuário**
`POST /users`
Cria um novo usuário no sistema.
- **Body**: `{ "id": string, "name": string, "email": string }`
```bash
curl.exe -X POST "http://localhost:8080/users" \
  -H "Content-Type: application/json" \
  -d '{ \"id\": \"USER-001\", \"name\": \"Victor\", \"email\": \"victor@teste.com\" }'
```

#### **3. Processar Transação (Ledger)**
`POST /wallet/transaction`
Registra movimentos de entrada (CREDIT) ou saída (DEBIT) na carteira.
- **Body**: `{ "account_id": string, "type": "DEBIT"|"CREDIT", "amount": float, "description": string }`
```bash
# Exemplo de Crédito
curl.exe -X POST "http://localhost:8080/wallet/transaction" \
  -H "Content-Type: application/json" \
  -d '{ \"account_id\": \"ACC-001\", \"type\": \"CREDIT\", \"amount\": 250.50, \"description\": \"Recebimento PIX\" }'
```

## Próximos Passos

1.  Configurar persistência real (PostgreSQL, MongoDB, etc.) no `internal/infrastructure`.
2.  Adicionar logs e monitoramento.
3.  Implementar mais casos de uso.
4.  Adicionar Testes Unitários e de Integração.
