# PR 009- Implementação do endpoint `/me` e infraestrutura para autenticação

**Status:** Merged

**Tipo:** Feature

**Commit:** `feat: create endpoint /me for listing user infos`

**Data:** 03/08/2026

---

# Contexto

Até este momento a aplicação possuía apenas endpoints públicos, sem uma forma padronizada de identificar o usuário autenticado durante uma requisição.

Além disso, a User Story **USER-002 - Visualizar usuário** ainda não havia sido implementada, impedindo que o cliente consultasse seus próprios dados utilizando o token JWT obtido no login.

Esta entrega introduz a primeira rota autenticada do projeto e estabelece a infraestrutura necessária para reutilização da autenticação em novos endpoints.

---

# Objetivo

Disponibilizar um endpoint para consulta das informações do usuário autenticado, utilizando o JWT para identificação do usuário, além de consolidar a infraestrutura necessária para futuras rotas protegidas.

---

# Decisões Técnicas

## Middleware de autenticação

Foi criada uma camada de middleware responsável por:

- Validar o token JWT enviado no header `Authorization`.
- Rejeitar requisições sem token ou com token inválido.
- Disponibilizar o `user_id` no contexto da requisição para consumo pelos handlers.

Com isso, a autenticação deixa de ser responsabilidade dos handlers, permitindo reutilização em qualquer rota protegida.

---

## Endpoint `/user/me`

Foi criado o endpoint responsável por retornar os dados do usuário autenticado.

Fluxo:

1. Middleware valida o JWT.
2. Extrai o `user_id`.
3. Handler recupera o `user_id` do contexto.
4. Usecase busca o usuário.
5. Repository realiza a consulta ao banco.
6. A resposta retorna apenas informações públicas do usuário.

---

## Separação entre entidade de persistência e resposta da API

Foi criada uma estrutura específica (`User`) para representar os dados públicos retornados pela API.

A entidade de persistência (`Users`) continua representando o modelo completo armazenado no banco, incluindo senha e demais campos internos.

A conversão é realizada através do método:

- `BuildUser()`

Essa separação evita exposição acidental de informações sensíveis e reduz o acoplamento entre persistência e contratos HTTP.

---

## Compartilhamento do serviço de autenticação

O `AuthService` passou a ser inicializado na aplicação e compartilhado entre os componentes que necessitam realizar operações com JWT.

Com isso:

- Login continua responsável pela geração do token.
- Middleware passa a reutilizar a mesma implementação para validação.

Essa abordagem elimina duplicidade de instâncias e centraliza a responsabilidade da autenticação.

---

## Evolução do domínio de usuários

Também foram realizadas pequenas evoluções na modelagem do domínio:

- Inclusão do método `GetUserById`.
- Adequação da entidade `Users`.
- Ajustes na representação de senha.
- Inclusão do campo `DeletedAt`.
- Padronização do atributo `Rule`.

---

## Backlog

Foram atualizadas as histórias já concluídas e criada a User Story:

- **USER-002 - Visualizar usuário**

representando formalmente a funcionalidade implementada.

---

# Ações realizadas

- Implementado endpoint `GET /user/me`.
- Criado middleware de autenticação JWT.
- Disponibilização do `user_id` no contexto da requisição.
- Implementado fluxo de consulta do usuário por ID.
- Criada representação pública da entidade `User`.
- Implementado método `BuildUser()`.
- Adicionado `GetUserById` no Usecase e Repository.
- Centralizada a criação do `AuthService`.
- Atualizadas histórias do backlog concluídas.
- Criada a história **USER-002 - Visualizar usuário**.

---

# Resultado

A aplicação passa a possuir sua primeira rota autenticada baseada em JWT.

Além da funcionalidade entregue ao usuário final, esta implementação estabelece a infraestrutura que será reutilizada pelos próximos endpoints protegidos, tornando a autenticação um comportamento transversal da aplicação e reduzindo o acoplamento entre regras de autenticação e regras de negócio.