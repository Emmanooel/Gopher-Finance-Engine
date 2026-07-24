# PR 003 - Extração da camada de Web Handlers

**Status:** Merged

**Tipo:** architetural refactor

**Commit:** `9639984`

**Data:** 24/02/2026

---

## Objetivo

Com os primeiros fluxos de negócio implementados no PR anterior, a camada HTTP passou a concentrar responsabilidades relacionadas tanto ao servidor quanto ao tratamento das requisições.

O objetivo desta etapa é separar essas responsabilidades, introduzindo uma camada dedicada de **Web Handlers**, responsável exclusivamente por receber requisições HTTP, realizar o binding dos dados e delegar a execução para os respectivos casos de uso.

Essa alteração melhora a organização da infraestrutura web e prepara a aplicação para crescer sem aumentar o acoplamento da camada de roteamento.

---

## Principais entregas

### Criação da camada de Handlers

Foi criada uma nova camada responsável exclusivamente pelo tratamento das requisições HTTP.

Cada Handler passa a possuir acesso apenas aos casos de uso necessários para executar sua responsabilidade, removendo essa responsabilidade da estrutura do servidor HTTP.

---

### Separação entre Servidor e Handlers

A estrutura `Server` deixa de conhecer diretamente todos os casos de uso da aplicação.

Sua responsabilidade passa a ser apenas:

- inicializar o servidor HTTP;
- criar a camada de Handlers;
- registrar as rotas disponíveis.

Essa alteração reduz o acoplamento entre o servidor e a lógica de processamento das requisições.

---

### Organização da infraestrutura Web

Todos os handlers foram movidos para um pacote dedicado:

- Health
- Users
- Orders
- Positions

Essa reorganização torna mais clara a separação entre:

- configuração do servidor;
- definição das rotas;
- implementação dos endpoints.

---

### Ajuste na composição das rotas

A configuração das rotas passa a utilizar a estrutura de `Handlers`, deixando explícito que as rotas dependem apenas da camada responsável pelo processamento HTTP.

Com isso, a definição das rotas permanece desacoplada da estrutura do servidor.

---

## Resultado

Ao final desta etapa, a camada Web passa a possuir uma divisão clara de responsabilidades:

```
Server
    │
    ├── inicializa o servidor HTTP
    ├── cria os Handlers
    └── registra as rotas

Handlers
    │
    ├── recebe a requisição
    ├── realiza o binding dos dados
    ├── chama o Use Case
    └── monta a resposta HTTP

Use Cases
    │
    └── executam a regra de negócio
```

Essa reorganização reduz o acoplamento da infraestrutura HTTP, melhora a legibilidade da aplicação e estabelece uma estrutura mais preparada para suportar o crescimento do número de endpoints sem concentrar responsabilidades em uma única camada.