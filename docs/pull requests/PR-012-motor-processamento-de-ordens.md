# PR 013 - Implementação do motor de processamento assíncrono de ordens

**Status:** Merged

**Tipo:** Feature

**Commit:** `4e6bcda`

**Data:** 03/08/2026

---

# Contexto

Até este momento o projeto já possuía toda a lógica necessária para processar ordens pendentes, atualizar posições, recalcular preço médio e concluir o processamento.

Entretanto, toda essa lógica ainda dependia de execução manual durante os testes.

Este PR introduz o primeiro processo assíncrono da aplicação, criando um Worker dedicado responsável por executar continuamente o processamento das ordens pendentes.

Com isso, o sistema deixa de depender de intervenção manual e passa a possuir um fluxo automático de processamento.

---

# Objetivo

Implementar uma aplicação dedicada ao processamento de ordens pendentes, executando periodicamente toda a regra de negócio já construída anteriormente.

Além disso, separar completamente o ciclo de vida do Worker da API HTTP, permitindo que ambos evoluam de forma independente.

---

# Decisões Técnicas

## Novo entrypoint da aplicação

Foi criado um novo executável em:

```
cmd/worker
```

Assim como a API possui seu próprio processo (`cmd/api`), o Worker passa a possuir seu próprio ponto de entrada.

Isso permite executar ambos os serviços de maneira independente em ambientes produtivos.

---

## Criação do Engine de processamento

Foi criada a camada:

```
internal/application/engine
```

Essa camada passa a ser responsável exclusivamente pelo ciclo de vida do Worker.

Responsabilidades:

- inicialização;
- execução periódica;
- controle do loop de processamento;
- encerramento da aplicação.

A regra de negócio continua pertencendo ao módulo de Orders.

---

## Processamento periódico

O Worker passa a executar automaticamente o processamento utilizando um `Ticker`.

Fluxo:

```
Inicialização
      │
      ▼
Executa processamento
      │
      ▼
Aguarda 15 segundos
      │
      ▼
Executa novamente
      │
      ▼
(repetição contínua)
```

Dessa forma, novas ordens pendentes são processadas continuamente sem necessidade de chamadas externas.

---

## Graceful Shutdown

O processo do Worker passa a responder corretamente aos sinais do sistema operacional.

Foram adicionados:

- SIGTERM
- CTRL+C (Interrupt)

permitindo finalizar o processamento de forma controlada antes do encerramento da aplicação.

---

## Evolução do processamento

O processamento deixou de operar sobre um usuário específico.

Antes:

```
ProcessOrders(userId)
```

Agora:

```
ProcessPendingOrders()
```

A busca pelas ordens pendentes passa a ocorrer diretamente no repositório, permitindo que o Worker processe múltiplos usuários em uma única execução.

---

## Responsabilidades preservadas

Mesmo com a criação do Worker, as responsabilidades permanecem bem definidas.

```
Worker
    │
    ▼
Orders Usecase
    │
    ▼
Position Service
    │
    ▼
Position Usecase
    │
    ▼
Repositories
```

O Worker apenas agenda e dispara o processamento.

Toda regra de negócio continua centralizada na camada de aplicação.

---

# Ações realizadas

- Criado novo executável `cmd/worker`.
- Implementada a camada `engine`.
- Criado ciclo de execução periódico utilizando `time.Ticker`.
- Implementado suporte a Graceful Shutdown.
- Adaptado o processamento para buscar ordens pendentes globalmente.
- Removida a dependência de processamento por usuário.
- Estruturada uma aplicação dedicada exclusivamente ao processamento assíncrono.

---

# Resultado

A arquitetura passa a ser composta por dois processos independentes:

```
                ┌────────────────────┐
                │      API HTTP      │
                └─────────┬──────────┘
                          │
                  Cria novas ordens
                          │
                          ▼
                     Banco de Dados
                          ▲
                          │
                Busca ordens pendentes
                          │
                ┌─────────┴──────────┐
                │      Worker        │
                └─────────┬──────────┘
                          │
                  Processa ordens
                          │
                          ▼
               Atualiza posições
               Atualiza preço médio
               Conclui processamento
```

---

# Marco

Este PR representa um dos maiores marcos arquiteturais do projeto.

Até este ponto, o Gopher Finance Engine era capaz de receber requisições e executar regras de negócio sob demanda.

A partir desta implementação, o sistema passa a possuir um processo assíncrono dedicado, responsável por executar continuamente o processamento das ordens, aproximando a arquitetura do modelo utilizado por plataformas financeiras reais.

Além de concluir o fluxo iniciado nos PRs anteriores, este commit estabelece a base para futuras evoluções, como filas de processamento, múltiplos Workers, escalabilidade horizontal, métricas de execução e mecanismos de retry, mantendo a separação entre entrada HTTP e processamento de domínio.