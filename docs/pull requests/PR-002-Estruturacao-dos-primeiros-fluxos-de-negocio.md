# PR 002 - Estruturação dos primeiros fluxos de negócio da aplicação

**Status:** Merged

**Tipo:** Feature

**Commit:** `4412083`

**Data:** 16/02/2026

---

## Objetivo

Após estabelecer o esqueleto da aplicação no PR anterior, esta etapa tem como objetivo implementar os primeiros fluxos de negócio ponta a ponta, validando que a arquitetura proposta suporta funcionalidades reais.

Além da criação dos endpoints iniciais, este PR consolida a separação entre as camadas da aplicação, permitindo que cada responsabilidade possua seu próprio contrato e implementação.

---

## Principais entregas

### Estruturação da aplicação

- Centralização da inicialização da aplicação através da `Application`.
- Implementação da composição dos componentes (Composition Root).
- Injeção manual de dependências entre serviços, repositórios e casos de uso.
- Inicialização da conexão com PostgreSQL durante o bootstrap da aplicação.
- Organização dos casos de uso disponíveis pela aplicação.

---

### Camada de Domínio

Criação das entidades responsáveis por representar os principais conceitos do sistema:

- User
- Order
- Position
- TaxAlert

Também foram definidos os contratos (interfaces) que representam:

- Casos de uso
- Repositórios

Essa separação permite que a regra de negócio permaneça desacoplada das implementações de infraestrutura.

---

### Casos de Uso

Implementação dos primeiros casos de uso da aplicação:

- Criação de usuário
- Login
- Criação de ordens
- Consulta de posições

Durante essa etapa também foram adicionadas regras de negócio como:

- geração de UUIDs;
- definição do status inicial das ordens;
- hash de senha utilizando bcrypt;
- autenticação baseada em JWT.

---

### Persistência

Implementação da camada de repositórios para comunicação com PostgreSQL.

Os primeiros repositórios implementados foram responsáveis por:

- criação de usuários;
- autenticação de usuários;
- criação de ordens;
- consulta de posições.

Com isso, os casos de uso passaram a possuir uma implementação concreta para persistência dos dados.

---

### API HTTP

Criação dos primeiros endpoints da aplicação:

| Método | Endpoint | Responsabilidade |
|---------|----------|------------------|
| GET | `/health/health` | Verificação de saúde da aplicação |
| POST | `/user/create` | Cadastro de usuário |
| POST | `/user/login` | Autenticação |
| POST | `/v1/orders` | Criação de ordens |
| GET | `/v1/portfolio/:id` | Consulta de posições |

---

## Resultado

Ao final desta etapa, a aplicação deixa de possuir apenas uma estrutura inicial e passa a executar fluxos completos de negócio.

Cada requisição agora percorre toda a arquitetura definida:

```
HTTP
    ↓
Handler
    ↓
Use Case
    ↓
Repository
    ↓
PostgreSQL
```

Essa entrega valida a arquitetura proposta e estabelece uma base sólida para a evolução das próximas funcionalidades, permitindo que novos casos de uso sejam implementados reutilizando a mesma estrutura de composição e separação de responsabilidades.