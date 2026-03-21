**Valores monetarios (centavos)**

Todo valor financeiro na API e no dominio e representado em **centavos** como `int64`.

| Real (BRL) | Centavos (API) |
|------------|----------------|
| R$ 1,00    | `100`          |
| R$ 250,50  | `25050`        |
| R$ 500,00  | `50000`        |

Motivos:

- Evita erros de arredondamento de ponto flutuante.
- Alinha-se a praticas comuns em sistemas financeiros (minor units).

Campos afetados:

- `Account.Balance`, `ProcessTransactionRequest.amount`, `ProcessTransactionResponse.new_balance`.
