# Go Wallet Core API Service (Go)

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](go.mod)
[![Tests](https://img.shields.io/badge/Tests-go%20test-0F9D58?logo=go&logoColor=white)](#testes)
[![Lint](https://img.shields.io/badge/Lint-golangci--lint-00C7B7?logo=go&logoColor=white)](.github/actions/lint/action.yml)
[![Pre-commit](https://img.shields.io/badge/Hooks-pre--commit-FAB040?logo=pre-commit&logoColor=white)](.pre-commit-config.yaml)
[![Release](https://img.shields.io/badge/Release-semantic--release-494949?logo=semantic-release&logoColor=white)](.releaserc.json)

Este projeto segue os princípios de **Domain-Driven Design (DDD)** e **Clean Architecture** para garantir uma base de código modular, testável e manutenível.

## Arquitetura do Projeto

A estrutura de diretórios foi organizada dentro da pasta `src/` para separar as responsabilidades em camadas:

- **src/cmd/app**: Ponto de entrada da aplicação. Responsável por inicializar as dependências e subir o servidor.
- **src/internal/domain**: O coração da aplicação. Contém as regras de negócio, entidades e interfaces (contratos) de repositórios.
- **src/internal/usecase**: Camada de aplicação que orquestra a execução das regras de negócio usando as entidades do domínio.
- **src/internal/infrastructure**: Implementações detalhadas de infraestrutura. Neste projeto, a persistência atual e in-memory (sem banco externo).
- **src/internal/interface**: Adaptadores de interface para expor a lógica interna (HTTP Handlers, gRPC, CLI).


## Como Executar

Se você tiver o Go instalado:

1.  Aponte para o diretório raiz do projeto.
2.  Assure-se de que o `go.mod` está configurado corretamente.
3.  Execute o comando:

```bash
go run src/cmd/app/main.go
```

Por padrao, a aplicacao inicia com uma conta seed:

- `account_id`: `ACC-001`
- `balance`: `500.0`

Para iniciar sem seed:

```bash
SEED_DEFAULT_ACCOUNT=false go run src/cmd/app/main.go
```

## Qualidade e Ciclo de Commits

O fluxo de qualidade deste projeto e obrigatorio em todo o ciclo:

- **Local (`pre-commit`)**: roda `clean-workspace` com `go run ./scripts/core/clean_workspace.go` e `go test ./...`.
- **Local (`commit-msg`)**: roda `commitlint` com Conventional Commits.
- **CI**: valida novamente commitlint e executa qualidade/seguranca.
- **Release**: `semantic-release` gera versao e atualiza `CHANGELOG.md`.

### Padrao de commit

Use o formato:

```text
type(scope): descricao curta
```

Exemplos validos:

```text
feat(wallet): adicionar processamento de transacao
fix(api): corrigir serializacao de erro
chore(config): ajustar pipeline de release
```

Regras importantes:

- `type` e `scope` devem respeitar `commitlint.config.mjs`.

## Testes

Executar testes localmente:

```bash
go test ./...
```

Testes adicionados no projeto:

- `src/internal/usecase/wallet/process_transaction_test.go`
- `src/internal/infrastructure/repository/postgres/user_repository_test.go`
- `src/internal/infrastructure/repository/postgres/wallet_repository_test.go`

## Release e Changelog

- A liberacao e automatica na branch `main` via `semantic-release`.
- O changelog oficial e versionado no arquivo `CHANGELOG.md`.
- O commit automatico de release segue o formato:
  - `chore(release): <versao> [skip ci]`

## Documentação da API

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

1.  Migrar valores monetarios de `float64` para centavos (`int64`) para evitar problemas de precisao.
2.  Adicionar logs estruturados e middleware de observabilidade.
3.  Expandir casos de uso e cobrir novos fluxos com testes.
4.  Melhorar contrato de erros HTTP com padronizacao por tipo de erro.
