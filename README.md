**Go Wallet Core API Service (Go)**

[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](go.mod)
[![Lint](https://img.shields.io/badge/Lint-golangci--lint-00C7B7?logo=go&logoColor=white)](.github/actions/lint/action.yml)
[![Tests](https://img.shields.io/badge/Tests-go%20test-0F9D58?logo=go&logoColor=white)](#testes)
[![Pre-commit](https://img.shields.io/badge/Hooks-pre--commit-FAB040?logo=pre-commit&logoColor=white)](.pre-commit-config.yaml)
[![CI](https://img.shields.io/github/actions/workflow/status/victor-silveira/go-wallet-core/ci.yml?branch=main&logo=github&label=CI)](https://github.com/victor-silveira/go-wallet-core/actions/workflows/ci.yml)
[![Release](https://img.shields.io/badge/Release-semantic--release-494949?logo=semantic-release&logoColor=white)](tools/releaserc.json)
[![API](https://img.shields.io/badge/API-REST-0A1E3F?logo=fastapi&logoColor=white)](api/swagger.yaml)
[![OpenAPI](https://img.shields.io/badge/Spec-OpenAPI-6BA539?logo=swagger&logoColor=white)](api/swagger.yaml)

Este projeto segue os principios de **Domain-Driven Design (DDD)** e **Clean Architecture** para garantir uma base de codigo modular, testavel e manutenivel.

**Arquitetura do Projeto**

A estrutura de diretorios foi organizada dentro da pasta `src/` para separar as responsabilidades em camadas:

- **src/main.go**: Ponto de entrada, logs estruturados (`log/slog`), `http.ServeMux` com rotas por metodo (Go 1.22+), encerramento gracioso (`Shutdown` + SIGINT/SIGTERM).
- **src/domain**: Entidades, erros de dominio e contratos de repositorio.
- **src/application**: Casos de uso da aplicacao.
- **src/infrastructure**: Implementacoes concretas (repositorio **em memoria** em `repository/memory`).
- **src/interfaces**: Handlers HTTP e serializacao de respostas.

**Valores monetarios**

Todos os valores sao **centavos** (`int64` na API JSON). Ex.: R$ 500,00 = `50000`. Ver [docs/money.md](docs/money.md).

**Como Executar**

Se voce tiver o Go instalado:

1.  Aponte para o diretorio raiz do projeto.
2.  Assure-se de que o `go.mod` esta configurado corretamente.
3.  Execute o comando (inclui todo o pacote `main` em `src/`, ex. `version.go`):

```bash
go run ./src
```

Por padrao, a aplicacao inicia com uma conta seed:

- `account_id`: `ACC-001`
- `balance`: `50000` centavos (R$ 500,00)

Os logs sao emitidos em **JSON** no stdout (nivel `INFO`). A primeira entrada inclui **`version`**, **`commit`** e **`go`** (runtime). Valores padrao sem build de release: `version` e `commit` = `dev`. Para encerrar com seguranca, use **Ctrl+C** (SIGINT) ou SIGTERM; o servidor conclui requisicoes ativas ate o timeout de shutdown.

Para iniciar sem seed:

```bash
SEED_DEFAULT_ACCOUNT=false go run ./src
```

**Versao (Git) e build**

Consultar versao do repositorio local (tags e estado de working tree):

```bash
git describe --tags --always --dirty
git rev-parse --short HEAD
```

Arranque com versao e commit injetados via `ldflags`:

```bash
go run -ldflags "-X main.Version=$(git describe --tags --always --dirty) -X main.Commit=$(git rev-parse --short HEAD)" ./src
```

Binario de producao (exemplo):

```bash
go build -o bin/go-wallet-core -ldflags "-X main.Version=$(git describe --tags --always --dirty) -X main.Commit=$(git rev-parse --short HEAD)" ./src
```

**Qualidade e Ciclo de Commits**

O fluxo de qualidade deste projeto e obrigatorio em todo o ciclo:

- **Local (`pre-commit`)**: roda `clean-workspace` com `go run ./scripts/core/clean_workspace.go` e `go test ./...`.
- **Local (`commit-msg`)**: roda `commitlint` com Conventional Commits.
- **CI**: lint em paralelo com testes; em seguida seguranca; na `main`, release apos seguranca.
- **Release**: `semantic-release` gera versao e atualiza `CHANGELOG.md`.

**Padrao de commit**

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

<a id="testes"></a>

**Testes**

Executar testes localmente:

```bash
go test ./...
```

Organizacao dos testes (ver [`tests/readme.md`](tests/readme.md)):

- `tests/unit/domain/entity` — entidades
- `tests/unit/application` — casos de uso
- `tests/unit/handler` — HTTP (`httptest`)
- `tests/unit/infrastructure` — repositorio em memoria
- `tests/integration` — fluxo com repositorios reais em memoria

**Release e Changelog**

- A liberacao e automatica na branch `main` via `semantic-release`.
- O changelog oficial e versionado no arquivo `CHANGELOG.md`.
- O commit automatico de release segue o formato:
  - `chore(release): <versao> [skip ci]`

**Documentacao da API**

A API segue padroes REST e possui documentacao tecnica disponivel na pasta `api/`. 

- [OpenAPI (Swagger)](api/swagger.yaml)
- [REST Client (VSCode)](api/requests.http)
- [Centavos e contrato monetario](docs/money.md)

**Endpoints Disponiveis**

**1. Health Check**
`GET /health`
Verifica se o servidor esta online.
```bash
curl.exe -X GET "http://localhost:8080/health"
```

**2. Criar Usuario**
`POST /users`
Cria um novo usuario no sistema.
- **Body**: `{ "id": string, "name": string, "email": string }`
```bash
curl.exe -X POST "http://localhost:8080/users" \
  -H "Content-Type: application/json" \
  -d '{ \"id\": \"USER-001\", \"name\": \"Victor\", \"email\": \"victor@teste.com\" }'
```

**3. Processar Transacao (Ledger)**
`POST /wallet/transaction`
Registra movimentos de entrada (CREDIT) ou saida (DEBIT) na carteira.
- **Body**: `{ "account_id": string, "type": "DEBIT"|"CREDIT", "amount": int (centavos), "description": string }`
```bash
curl.exe -X POST "http://localhost:8080/wallet/transaction" \
  -H "Content-Type: application/json" \
  -d '{ \"account_id\": \"ACC-001\", \"type\": \"CREDIT\", \"amount\": 25050, \"description\": \"Recebimento PIX\" }'
```

**Proximos Passos**

1.  Persistencia real (PostgreSQL ou outro) mantendo `repository` interfaces.
2.  Middleware de observabilidade (metricas, tracing).
3.  Contrato de erros HTTP padronizado (codigos de erro de dominio).
4.  Contrato OpenAPI automatizado (schema vs respostas) e E2E opcional.
