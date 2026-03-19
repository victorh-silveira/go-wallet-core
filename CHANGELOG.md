## [1.0.3](https://github.com/victorh-silveira/go-wallet-core/compare/v1.0.2...v1.0.3) (2026-03-19)

### Documentacao

* **config:** adicionar badges e remover emojis de release e changelog ([1ab4b8f](https://github.com/victorh-silveira/go-wallet-core/commit/1ab4b8fa9a94fa99c4a598efc3c7dd7a618b5bb0))

## [1.0.2](https://github.com/victorh-silveira/go-wallet-core/compare/v1.0.1...v1.0.2) (2026-03-19)

### 🐛 Bug Fix

* **infra:** reforcar validacoes de transacao e qualidade local ([c531ff5](https://github.com/victorh-silveira/go-wallet-core/commit/c531ff5e3c3f7d87857ebdfbfb4d8942913ec0b1))

## [1.0.1](https://github.com/victorh-silveira/go-wallet-core/compare/v1.0.0...v1.0.1) (2026-03-19)

### Chore

* **config:** alinhar ciclo de commit, pre-commit e release ([452cdcc](https://github.com/victorh-silveira/go-wallet-core/commit/452cdcc62998fa0bc192c31d1aee02504c56a944))

### Documentacao

* **readme:** remover emoji do README ([7301666](https://github.com/victorh-silveira/go-wallet-core/commit/7301666b2fe9cd07730e9644e454591aad11bb73))

### Bug Fix

* **config:** alinhar validacao local com commitlint do CI ([1f2413b](https://github.com/victorh-silveira/go-wallet-core/commit/1f2413b74055b09263a7afb22a8abb9f1ca9f651))

## 1.0.0 (2026-03-18)

### Nova Funcionalidade

* estrutura inicial com DDD e Clean Architecture aplicada ao Go Wallet Core ([01055c5](https://github.com/victorh-silveira/go-wallet-core/commit/01055c5cad883000342ef17ec13c3b0f94d0671d))
* sincroniza scripts locais de limpeza com as auditorias do CI/CD (lint e sec) ([d7ecafd](https://github.com/victorh-silveira/go-wallet-core/commit/d7ecafdcdcc893ea0f3ef2b52c7f50b66abb62f7))

### Correcao de Erros

* compila golangci-lint manualmente no CI para compatibilidade com Go 1.26 ([b188048](https://github.com/victorh-silveira/go-wallet-core/commit/b188048e303bc6ccba5727c8864301a483249245))
* forçar reconstrução do golangci-lint com Go 1.26 para compatibilidade ([1f9c1d0](https://github.com/victorh-silveira/go-wallet-core/commit/1f9c1d0d5c1674e906dfe5da20892c60d56ff38e))
* resolve vulnerabilidades de infra com Go 1.23 e protege tags [#nosec](https://github.com/victorh-silveira/go-wallet-core/issues/nosec) no script de limpeza ([ec8a749](https://github.com/victorh-silveira/go-wallet-core/commit/ec8a7495bbd1bd669614dde9f4c4ac8fb548eee5))
* resolve vulnerabilidades de segurança (gosec) no servidor e scripts ([61b5685](https://github.com/victorh-silveira/go-wallet-core/commit/61b56858a077f482bbf8c025cc971acfd227cfe3))
* restaura anotações [#nosec](https://github.com/victorh-silveira/go-wallet-core/issues/nosec) e protege contra remoção automática ([f2a4027](https://github.com/victorh-silveira/go-wallet-core/commit/f2a40271215b814f92c9abaa3334c92046de7f9b))
* trata retornos de erro ignorados apontados pelo golangci-lint ([dba2413](https://github.com/victorh-silveira/go-wallet-core/commit/dba241335a87ce293259baf4a924b8d3b4d27f44))
* upgrade geral para Go 1.26 em todo ecossistema para eliminar vulnerabilidades de stdlib ([ed9a568](https://github.com/victorh-silveira/go-wallet-core/commit/ed9a568f721acb7909987bd7fe965f2d4c31f8ae))

### Documentacao

* adiciona documentação interativa Swagger UI e especificações da API ([c81d4dc](https://github.com/victorh-silveira/go-wallet-core/commit/c81d4dcb432e442bd1f2cbe30ee757d9e5678daa))

### Integracao Continua (CI/CD)

* adiciona suporte à branch master no pipeline de CI/CD ([7c41eab](https://github.com/victorh-silveira/go-wallet-core/commit/7c41eabbbd1ca053267c6912804c1c747b562b6d))

### Manutencao e Configuracoes

* purificação final do script de limpeza e ecossistema sincronizado ([c9f0978](https://github.com/victorh-silveira/go-wallet-core/commit/c9f0978bf2215e6fb9967ef0043dd9241b8c5629))
* remove menção à branch master das configurações de CI e release ([2878324](https://github.com/victorh-silveira/go-wallet-core/commit/28783245142cc0a3678b7a0b1e1a156552e2ea9e))
