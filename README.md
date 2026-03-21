# Go Wallet Core API Service (Go)

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](go.mod)
[![Lint](https://img.shields.io/badge/Lint-golangci--lint-00C7B7?logo=go&logoColor=white)](.github/actions/lint/action.yml)
[![Tests](https://img.shields.io/badge/Tests-go%20test-0F9D58?logo=go&logoColor=white)](#testes)
[![Pre-commit](https://img.shields.io/badge/Hooks-pre--commit-FAB040?logo=pre-commit&logoColor=white)](.pre-commit-config.yaml)
[![CI](https://github.com/victorh-silveira/go-wallet-core/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/victorh-silveira/go-wallet-core/actions/workflows/ci.yml)
[![Release](https://img.shields.io/badge/Release-semantic--release-494949?logo=semantic-release&logoColor=white)](tools/releaserc.json)
[![API](https://img.shields.io/badge/API-REST-0A1E3F?logo=fastapi&logoColor=white)](api/swagger.yaml)
[![OpenAPI](https://img.shields.io/badge/Spec-OpenAPI-6BA539?logo=swagger&logoColor=white)](api/swagger.yaml)

Este projeto segue os princípios de **Domain-Driven Design (DDD)** e **Clean Architecture** para garantir uma base de código modular, testável e manutenível.

## Arquitetura do Projeto

A estrutura de diretórios foi organizada dentro da pasta `src/` para separar as responsabilidades em camadas:

- **src/main.go**: Ponto de entrada da aplicação.
- **src/domain**: Entidades e contratos de repositório.
- **src/application**: Casos de uso da aplicação.
- **src/infrastructure**: Implementações concretas (repositório in-memory).
- **src/interfaces**: Handlers HTTP e serialização de respostas.


## Como Executar

Se você tiver o Go instalado:

1.  Aponte para o diretório raiz do projeto.
2.  Assure-se de que o `go.mod` está configurado corretamente.
3.  Execute o comando:

```bash
go run src/main.go
```

Por padrao, a aplicacao inicia com uma conta seed:

- `account_id`: `ACC-001`
- `balance`: `500.0`

Para iniciar sem seed:

```bash
SEED_DEFAULT_ACCOUNT=false go run src/main.go
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

- `tests/unit/application/process_transaction_test.go`
- `tests/unit/infrastructure/user_repository_test.go`
- `tests/unit/infrastructure/wallet_repository_test.go`

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
