export default {
  extends: ["@commitlint/config-conventional"],
  rules: {
    "type-enum": [
      2,
      "always",
      [
        "build",    // Build/Dependências
        "chore",    // Manutenção/Config
        "ci",       // CI/CD
        "docs",     // Documentação
        "feat",     // Nova funcionalidade
        "fix",      // Correção de erro
        "perf",     // Performance
        "qa",       // Auditoria de Qualidade
        "refactor", // Refatoração
        "revert",   // Reversão
        "style",    // Estilo/Formatação
        "test",     // Testes
      ],
    ],
    "scope-enum": [
      2,
      "always",
      [
        "api",      // Comunicação Deriv
        "build",    // Build e empacotamento
        "clean",    // Limpeza de workspace
        "cli",      // Entrypoints/linha de comando
        "config",   // Configurações
        "deps",     // Dependências
        "domain",   // Core Domain/Business
        "docs",     // Artefatos de documentação
        "infra",    // Infraestrutura
        "release",  // Versionamento/release
        "repo",     // Repositório de dados
        "security", // Segurança
        "tests",    // Testes
        "usecase",  // Casos de uso
        "wallet",   // Funcionalidades de carteira
      ],
    ],
    "scope-empty": [0],
    "subject-empty": [2, "never"],
    "type-empty": [2, "never"],
    "header-max-length": [2, "always", 100],
    "subject-case": [0],
  },
};
