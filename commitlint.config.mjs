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
        "clean",    // Limpeza de workspace
        "config",   // Configurações
        "engine",   // Motor de execução
        "infra",    // Infraestrutura/Git
        "models",   // IA e modelos
        "risk",     // Gestão de risco
        "state",    // Estado/Persistência
        "strategy", // Estratégia e indicadores
        "sync",     // Sincronia de dados
        "terminal", // Interface/Logs
        "domain",   // Core Domain/Business
        "logic",    // Trading Logic
        "physics",  // Physics Engines
      ],
    ],
    "scope-empty": [0],
    "subject-empty": [2, "never"],
    "type-empty": [2, "never"],
    "header-max-length": [2, "always", 100],
    "subject-case": [0],
  },
};
