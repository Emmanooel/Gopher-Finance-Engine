# PR-013 - Histórico de Ordens do Usuário

**Status:** Merged

**Tipo:** Feature

**Commit:** `a26cbc3`

**Data:** 03/08/2026

---

# Contexto

Após concluir todo o fluxo de processamento de ordens e atualização automática da carteira do usuário, ainda faltava disponibilizar uma forma do próprio usuário consultar todas as operações realizadas.

Esta entrega implementa o endpoint responsável por retornar o histórico completo de ordens do usuário autenticado, finalizando a camada de consulta da aplicação.

Além da funcionalidade em si, também foram concluídas as histórias de backlog relacionadas à autenticação e consulta de histórico.

---

# Objetivo

Disponibilizar uma API para consulta do histórico de ordens do usuário autenticado, permitindo que o frontend possa apresentar todas as operações realizadas utilizando o próprio contexto de autenticação (JWT), sem necessidade de informar manualmente o identificador do usuário.

---

# Decisões Técnicas

## Endpoint autenticado

O histórico passa a ser acessado através de:

```
GET /v1/historico
```

A identificação do usuário é realizada através do middleware de autenticação, utilizando o `user_id` presente no JWT.

Essa abordagem evita exposição de identificadores na URL e garante que cada usuário tenha acesso apenas ao seu próprio histórico.

---

## Novo caso de uso

Foi criado um novo fluxo na camada de aplicação responsável por recuperar todas as ordens pertencentes ao usuário.

Responsabilidades:

- solicitar os dados ao repositório;
- encapsular a resposta em um objeto próprio (`OrderResponse`);
- manter a regra de negócio isolada da camada HTTP.

---

## Padronização da resposta

Foi criada uma estrutura específica de resposta:

```go
OrderResponse
```

permitindo manter um contrato consistente entre os endpoints da aplicação e facilitando futuras implementações de metadados, paginação e filtros.

---

## Persistência

Foi implementado um novo método no repositório:

```
GetAllOrdersByUserId(...)
```

Responsável exclusivamente por recuperar todas as ordens pertencentes ao usuário autenticado.

A responsabilidade do repositório permanece limitada ao acesso aos dados, sem aplicação de regras de negócio.

---

## Camada HTTP

Foi criado o handler:

```
ShowOrdersHistory
```

Responsável por:

- recuperar o `user_id` do contexto;
- executar o caso de uso;
- retornar o histórico em formato JSON.

Toda a lógica permanece centralizada na camada de aplicação.

---

## Atualização do Backlog

Foram concluídas as histórias:

- ✅ AUTH-001 — Login
- ✅ ORDER-002 — Histórico de Ordens

Todos os critérios de aceite dessas histórias passaram para o estado **Done**.

---

# Ações Realizadas

- Implementado caso de uso para consulta de histórico de ordens.
- Criado `OrderResponse` como contrato de resposta.
- Adicionado método `GetAllOrdersByUserId` no repositório.
- Implementado handler `ShowOrdersHistory`.
- Criado endpoint autenticado `GET /v1/historico`.
- Integração com middleware JWT para identificação automática do usuário.
- Atualizadas as histórias do backlog para **Done**.

---

# Resultado

O backend passa a disponibilizar todas as funcionalidades previstas para o MVP:

- ✅ Cadastro de usuários
- ✅ Login com JWT
- ✅ Cadastro de ordens
- ✅ Processamento automático de ordens
- ✅ Atualização da carteira do usuário
- ✅ Consulta da carteira consolidada
- ✅ Consulta do histórico completo de ordens

Com esta entrega, todas as histórias funcionais inicialmente planejadas para o projeto foram concluídas, restando apenas atividades de documentação arquitetural e diagramas de apoio para encerramento do projeto.