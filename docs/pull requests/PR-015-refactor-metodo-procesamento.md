# PR-015 - Refatoração do Fluxo de Processamento de Ordens

**Status:** Merged

**Tipo:** Refactor

**Commit:** `05f1647`

**Data:** 06/08/2026

---

# Contexto

Após a conclusão do MVP do motor de processamento de ordens, foi identificado um débito técnico na implementação do caso de uso `ProcessPendingOrders`.

Embora o fluxo estivesse funcional, toda a lógica de processamento encontrava-se concentrada em um único método, acumulando responsabilidades como busca da posição, criação de novas posições, decisão do tipo de operação (compra ou venda), atualização da carteira e alteração do status da ordem.

Esta entrega tem como objetivo reduzir essa complexidade, tornando o processamento mais modular, extensível e aderente aos princípios de responsabilidade única.

---

# Objetivo

Refatorar o fluxo de processamento de ordens, separando responsabilidades e removendo decisões de negócio concentradas em um único método, sem alterar o comportamento funcional da aplicação.

---

# Decisões Técnicas

## Introdução do Strategy Pattern

A decisão de como processar cada ordem deixa de ser realizada através de um `switch` dentro do caso de uso.

Foi introduzida uma fábrica de estratégias responsável por selecionar a implementação adequada de acordo com o tipo da ordem.

```
BUY  -> BuyStrategy
SELL -> SellStrategy
```

O processamento passa a delegar a responsabilidade da execução para cada estratégia especializada.

Essa abordagem reduz o acoplamento do caso de uso e facilita a inclusão de novos tipos de operações futuramente.

---

## Extração da resolução da posição

A responsabilidade de localizar ou criar a posição do usuário foi isolada no método:

```
getPositions(...)
```

Agora o fluxo principal apenas solicita uma posição válida, sem conhecer os detalhes necessários para obtê-la.

Isso elimina duplicação de responsabilidade e melhora significativamente a legibilidade do processamento.

---

## Padronização da atualização do status

O método anteriormente chamado de:

```
UpdateStatusOrders(...)
```

foi renomeado para:

```
markOrderProcessed(...)
```

A nova nomenclatura representa melhor a intenção da operação.

Em vez de expor uma atualização genérica de status, o caso de uso passa a expressar claramente a transição de uma ordem para o estado **PROCESSADO**. Além de ser um metodo privado apenas para a funcionalidade de processamento das ordens.

---

## Simplificação do fluxo principal

Após a refatoração, o processamento passou a seguir um fluxo de alto nível:

1. Buscar ordens pendentes.
2. Resolver a posição correspondente.
3. Selecionar a estratégia adequada.
4. Executar o processamento da ordem.
5. Marcar a ordem como processada.

Cada responsabilidade permanece isolada em seu próprio componente.

---

# Ações Realizadas

- Removido o `switch` de processamento do caso de uso principal.
- Integrado o Strategy Pattern ao fluxo de processamento.
- Extraída a lógica de busca/criação da posição para `getPositions`.
- Renomeado `UpdateStatusOrders` para `MarkOrderProcessed`.
- Simplificado o método `ProcessPendingOrders`.
- Centralizada a responsabilidade de atualização do status após o processamento.

---

# Resultado

O fluxo de processamento permanece com o mesmo comportamento funcional, porém apresenta uma estrutura significativamente mais modular e legível.

A responsabilidade de cada etapa do processamento passa a estar distribuída entre componentes especializados, reduzindo o acoplamento do caso de uso principal e facilitando futuras evoluções, como inclusão de novos tipos de ordens ou novas estratégias de processamento, sem necessidade de modificar o fluxo central.

---

# Evolução Arquitetural

Esta entrega representa a quitação de um débito técnico identificado durante a implementação do MVP.

Ao priorizar inicialmente a funcionalidade e realizar a refatoração posteriormente, foi possível validar toda a regra de negócio antes de reorganizar sua estrutura interna.

O resultado é um motor de processamento mais aderente aos princípios de Clean Architecture e SOLID, no qual o caso de uso atua como um orquestrador das etapas do processamento, enquanto as regras específicas permanecem encapsuladas em componentes especializados.