# PROCESS-001 - Processar ordens

## Objetivo

Processar ordens pendentes e atualizar automaticamente a posição consolidada do usuário.

---

## Critérios de aceite

- [x] Buscar ordens pendentes.
- [x] Processar cada ordem.
- [x] Calcular o preço médio.
- [x] Atualizar a posição do usuário.
- [x] Alterar o status da ordem após o processamento.

---

## Regras de negócio

- Cada ordem deve ser processada apenas uma vez.
- O processamento deve refletir corretamente a posição consolidada do usuário.

---

## Status

Done