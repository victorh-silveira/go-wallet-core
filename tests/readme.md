**Testes**

Suite organizada por tipo:

| Pasta | Conteudo |
|-------|----------|
| `unit/domain/entity` | Entidades: saldo, usuario, transacao |
| `unit/application` | Casos de uso: criar usuario, processar transacao |
| `unit/handler` | HTTP: `httptest` + repositorios em memoria |
| `unit/infrastructure` | Repositorio `memory`: copias, erros, transacoes |
| `integration` | Fluxo com `memory` + caso de uso real |
| `e2e` | Placeholder — ver `e2e/readme.md` |
| `contract` | Placeholder — ver `contract/readme.md` |

Regras de negocio e valores monetarios (**centavos** `int64`) sao cobertos no dominio e na aplicacao.

**Executar**

```bash
go test ./...
```

Com cobertura:

```bash
go test -cover ./...
```
