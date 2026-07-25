# PR 007 - Refatoração da Arquitetura da Aplicação e Início do Fluxo de Processamento de Ordens

**Status:** Merged

**Tipo:** Chore

**Commit:** `6e7475`

**Data:** 25/07/2026

---

## Contexto

Após a implementação inicial dos casos de uso, a camada de aplicação ainda mantinha todos os módulos concentrados em um único pacote (`application/usecases`). Embora funcional, essa organização aumentava o acoplamento entre funcionalidades e dificultava a evolução independente dos domínios.

Além disso, iniciou-se a implementação do fluxo responsável pelo processamento das ordens pendentes, preparando a aplicação para consolidar automaticamente a posição patrimonial do usuário após a execução das ordens.

---

## Objetivo

Reestruturar a camada de aplicação para refletir a separação por módulos de negócio e iniciar a implementação do pipeline responsável pelo processamento de ordens pendentes.

---

## Decisões Técnicas

- Reorganizar a camada `application`, separando os módulos de negócio em:
  - `application/users`
  - `application/orders`
  - `application/positions`
- Reestruturar a camada de repositórios seguindo a mesma divisão por domínio.
- Introduzir a camada `positions/service`, responsável por disponibilizar regras de negócio compartilhadas relacionadas às posições, permitindo que outros módulos consumam essas funcionalidades sem depender diretamente de contratos de persistência.
- Alterar o fluxo de processamento para trabalhar a partir do `userId`, buscando todas as ordens pendentes antes da consolidação das posições.
- Manter a atualização do status das ordens fora deste PR, priorizando primeiro a validação do fluxo completo de processamento.

---

## Ações Realizadas

- Refatorada toda a organização da camada `application`, separando os módulos por contexto de negócio.
- Reorganizados os repositórios da infraestrutura seguindo a nova estrutura modular.
- Criada a estrutura inicial da camada `positions/service`.
- Iniciado o pipeline de processamento de ordens pendentes.
- Implementada a consulta de ordens com status `PENDING`.
- Implementada a consolidação das ordens por ativo, preparando os dados para atualização da posição do usuário.
- Adicionado o campo `TotalCost` na entidade de posições para suportar futuros cálculos de preço médio.
- Mantida a criação de ordens com status `PENDING`, deixando a atualização para `COMPLETED` para a próxima etapa da implementação.

---

## Resultado

A arquitetura da aplicação passou a refletir melhor os limites entre os domínios de negócio, reduzindo o acoplamento entre módulos e tornando a evolução do sistema mais previsível.

Também foi construída a base do pipeline de processamento de ordens, deixando preparada a implementação das próximas etapas, como atualização da posição consolidada do usuário, cálculo do preço médio e alteração do status das ordens após o processamento completo.