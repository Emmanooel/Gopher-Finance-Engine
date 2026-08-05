# PR-014 - Autenticação das Operações da Carteira

**Status:** Merged

**Tipo:** Feature

**Commit:** `be40e70`

**Data:** 05/08/2026

---

# Contexto

Com a interface da aplicação integrada ao backend, foi necessário adequar os endpoints responsáveis pelas operações da carteira para utilizar exclusivamente o contexto do usuário autenticado.

Anteriormente, o endpoint de criação de ordens não exigia autenticação, permitindo que o identificador do usuário pudesse ser informado externamente. Essa abordagem não representa o comportamento esperado em uma aplicação real.

Esta entrega centraliza a identificação do usuário através do JWT, garantindo que todas as operações ocorram em nome do usuário autenticado.

---

# Objetivo

Garantir que as operações relacionadas à carteira utilizem exclusivamente o usuário autenticado, eliminando a dependência de informações enviadas pelo cliente e tornando o fluxo compatível com uma aplicação de produção.

---

# Decisões Técnicas

## Proteção do endpoint de criação de ordens

O endpoint:

```
POST /v1/orders
```

passa a utilizar o middleware de autenticação.

A partir deste momento, apenas usuários autenticados podem registrar novas ordens.

---

## Utilização do contexto de autenticação

Após a validação do JWT, o `user_id` passa a ser recuperado diretamente do contexto da requisição.

```go
body.UserId = c.GetString("user_id")
```

Dessa forma:

- o cliente não informa o usuário da operação;
- o backend torna-se responsável pela associação da ordem ao usuário autenticado;
- elimina-se a possibilidade de criação de ordens para outro usuário através da API.

---

## Organização da estrutura interna

O pacote de utilidades foi movido de:

```
internal/application/utils
```

para

```
internal/utils
```

A alteração reflete melhor sua responsabilidade, permitindo que funções utilitárias sejam compartilhadas entre diferentes módulos da aplicação sem depender da camada de Application.

---

## Ajustes de inicialização

Foi realizado um pequeno refinamento na criação da aplicação para garantir que o serviço de autenticação seja inicializado antes da construção das dependências que fazem uso dele.

Essa alteração melhora a organização do bootstrap da aplicação sem impactar o comportamento existente.

---

## Observabilidade

Foram adicionados logs auxiliares durante o fluxo de autenticação e criação de ordens para facilitar a validação da integração entre frontend, JWT e processamento das operações.

---

# Ações Realizadas

- Protegido o endpoint `POST /v1/orders` com autenticação JWT.
- Associação automática do `user_id` à ordem criada.
- Removida a dependência de identificação enviada pelo cliente.
- Reorganizado o pacote de utilidades para `internal/utils`.
- Refinado o processo de inicialização da aplicação.
- Adicionados logs para apoio ao processo de integração e depuração.

---

# Resultado

A API passa a operar utilizando exclusivamente o contexto do usuário autenticado, tornando o fluxo de criação de ordens consistente com os demais endpoints protegidos da aplicação.

Com esta entrega, o backend fica completamente preparado para consumo por uma interface web, mantendo a autenticação como única fonte de identidade do usuário durante todas as operações da carteira.

---

# Marco

Este foi o primeiro commit validado através da interface da aplicação integrada ao backend.

Pela primeira vez foi possível visualizar, em uma experiência completa de ponta a ponta, todo o fluxo implementado ao longo do projeto:

- autenticação do usuário;
- criação de ordens;
- processamento assíncrono pelo worker;
- atualização automática da carteira;
- consulta da posição consolidada;
- visualização do histórico de ordens.

Mais do que uma alteração técnica, esta entrega representou a materialização do sistema funcionando como um produto, permitindo validar visualmente toda a arquitetura construída durante o projeto.