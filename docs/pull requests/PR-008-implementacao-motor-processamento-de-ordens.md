# PR 008 - Implementação do processamento de posições

**Status:** Merged

**Tipo:** Feature

**Commit:** `4d283ee`

**Data:** 16/07/2026

---

# Contexto

Até este momento o projeto era capaz de cadastrar ordens e persistir posições, porém ainda não existia um fluxo responsável por transformar uma ordem em uma alteração real da carteira do usuário.

A arquitetura já possuía as camadas necessárias, mas a regra central do domínio financeiro ainda não havia sido implementada.

Este PR representa a implementação dessa primeira versão do motor de processamento.

---

# Objetivo

Implementar a regra de negócio responsável por processar uma ordem e refletir seu impacto na posição do usuário, separando responsabilidades entre os módulos da aplicação e iniciando a construção do motor financeiro do projeto.

---

# Decisões Técnicas

## Modularização da camada Application

A organização da camada de Application foi revisada.

Os casos de uso deixaram de ficar concentrados em um único pacote e passaram a ser organizados por contexto de negócio.

Exemplo:

```
application/
    users/
    orders/
    positions/
```

Essa organização reduz o acoplamento entre funcionalidades e facilita a evolução independente de cada domínio.

---

## Criação do serviço de processamento de Position

Foi criada uma camada responsável exclusivamente pelo processamento da posição do usuário baseada na ordem criada.

Esse componente concentra as regras de atualização da carteira, permitindo que outros módulos utilizem essa funcionalidade sem conhecer detalhes de persistência.

A responsabilidade da Position deixa de ser apenas consultar dados e passa a representar efetivamente o comportamento do agregado.

---

## Separação entre persistência e regra de negócio

O processamento passou a ocorrer na camada de aplicação.

Os repositórios permanecem responsáveis apenas pelas operações de acesso aos dados, enquanto toda decisão sobre criação, atualização e processamento das posições passa a ocorrer na camada de domínio.

Essa mudança reduz o acoplamento com o banco de dados e facilita futuras evoluções das regras financeiras.

---

## Implementação do fluxo de processamento

Foi implementado o fluxo responsável por:

- receber uma nova ordem;
- localizar a posição existente;
- decidir se a posição deve ser criada ou atualizada;
- executar a lógica correspondente.

Esse processamento foi executado diretamente através do UseCase durante os testes de desenvolvimento, permitindo validar o comportamento do domínio antes da introdução de mecanismos assíncronos.

---

## Evolução das entidades

A entidade Position passou a possuir responsabilidades relacionadas ao próprio domínio.

Foi adicionada a construção da posição inicial a partir de uma ordem, preparando o modelo para futuras evoluções, como cálculo de preço médio, vendas e atualização de quantidade.

---

# Fluxo implementado

Após este PR, o processamento passa a seguir o seguinte fluxo:

```
                          Nova Ordem
                              │
                              ▼
                        Orders UseCase
                              │
                              ▼
                           Position
                              │
                 ┌────────────┴────────────┐
                 │                         │
            Não encontrada             Encontrada
                 │                         │
                 ▼                         ▼
            cria Position              calcula position
                 │                         │
                 ▼                         ▼
            SavePosition()             UpdatePosition()
```

---

# Benefícios

- Implementação da primeira versão do processamento de posições.
- Separação entre CRUD e regras de negócio.
- Modularização da camada Application.
- Centralização da lógica de processamento dentro do domínio.
- Base preparada para evolução das regras financeiras.
- Redução do acoplamento entre módulos.

---

# Próximos passos

Os próximos PRs darão continuidade ao motor de processamento, incluindo:

- processamento de vendas;
- execução assíncrona através de Workers;
- processamento baseado em filas e eventos.

---

# Marco do Projeto

Este PR representa um dos marcos mais importantes do desenvolvimento do **Gopher Finance Engine**.

Até aqui o projeto possuía a infraestrutura necessária para receber e armazenar informações. A partir desta implementação, ele passa a executar uma regra real do domínio financeiro: transformar uma ordem em uma atualização da carteira do usuário.

Foi também o primeiro momento em que a arquitetura modular começou a demonstrar seu propósito, permitindo que diferentes módulos colaborassem através de responsabilidades bem definidas, sem depender diretamente da camada de persistência.

Em outras palavras, este foi o ponto em que o projeto deixou de ser apenas uma API para cadastro de dados e começou a se comportar como um motor financeiro.

Esse processamento foi validado executando diretamente os casos de uso durante o desenvolvimento, estabelecendo a base sobre a qual as próximas evoluções — processamento assíncrono, workers, filas e eventos — serão construídas.