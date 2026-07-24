# PR 004 - Worker assíncrono para atualização de posições

**Status:** Merged

**Tipo:** Feature

**Commit:** `9639984`

**Data:** 11/03/2026

---

## Objetivo

Após disponibilizar a API e os casos de uso principais, o próximo passo foi desacoplar responsabilidades que não precisam fazer parte do fluxo síncrono da requisição.

Até este momento, uma ordem era persistida, porém nenhuma atualização da carteira (positions) acontecia automaticamente. Este PR introduz um fluxo responsável por processar esse efeito colateral através de um worker dedicado.

Além disso, foram realizados ajustes na modelagem da camada de Position para permitir tanto consultas quanto atualização incremental das posições do usuário.

---

## O que foi implementado

### Criação da camada de Worker

Foi criada uma camada responsável por executar operações decorrentes da criação de uma ordem sem concentrar toda essa responsabilidade dentro do OrdersUsecase.

O fluxo passou a ser:

```
HTTP Request
      │
      ▼
OrdersUsecase
      │
      ├── Persiste Order
      │
      └── Worker
              │
              ▼
      PositionUsecase
              │
              ▼
      PositionsRepository
```

Essa separação permite evoluir futuramente para workers executados por filas, mantendo a regra de negócio isolada.

---

### Atualização automática das posições

Após salvar uma nova ordem, o sistema passa a disparar automaticamente um processamento responsável por refletir essa operação na carteira do usuário.

Foi criada a operação:

- SavePositionByNewOrder()

responsável por:

- construir uma Position baseada na Order criada;
- persistir uma nova posição quando inexistente;
- atualizar posições existentes utilizando UPSERT.

---

### Modelagem da entidade Position

A entidade passou a encapsular sua própria construção através do método:

```
BuildPositionByOrder()
```

Esse método centraliza a transformação entre Order → Position, evitando que essa lógica fique espalhada entre usecases ou repositories.

---

### UPSERT de posições

O repository passou a utilizar:

```
INSERT ...
ON CONFLICT (...)
DO UPDATE
```

para permitir atualização incremental da carteira.

Com isso, novas compras do mesmo ativo deixam de gerar registros duplicados e passam a atualizar:

- quantidade total;
- preço médio;
- data de atualização.

---

### Evolução do PositionUsecase

O caso de uso deixou de ser apenas uma camada de consulta e passou também a expor operações de escrita.

Foram adicionadas responsabilidades como:

- salvar posição;
- construir Position a partir de Order;
- orquestrar atualização da carteira.

---

### Ampliação dos contratos

As interfaces do domínio evoluíram para suportar o novo fluxo.

Foram adicionados métodos como:

- SavePositionByNewOrder()
- SaveNewPosition()
- GetPositionByUserId()
- GetOrdersInPendingByUserId()

Essa evolução mantém o domínio desacoplado das implementações concretas.

---

### Injeção de dependências

O bootstrap da aplicação também evoluiu.

Agora o OrdersUsecase recebe:

- PositionUsecase
- WorkerSaveNewPosition

permitindo orquestrar o processamento sem conhecer detalhes de persistência.

---

## Motivação arquitetural

Este PR introduz a primeira separação explícita entre:

- processamento síncrono da API;
- processamento interno decorrente de eventos da aplicação.

Embora o worker ainda execute dentro do mesmo processo, sua criação prepara a arquitetura para uma futura migração para processamento assíncrono baseado em filas (Kafka, RabbitMQ, SQS etc.), sem necessidade de alterar as regras de negócio.

Ao mesmo tempo, a responsabilidade de atualização da carteira deixa de estar acoplada ao endpoint de criação de ordens, tornando cada componente responsável apenas por sua própria etapa do fluxo.

---

## Próximos passos

Com essa base estabelecida, as próximas evoluções poderão focar em:

- substituição do worker interno por mensageria;
- processamento resiliente com retry;
- consolidação de posições mais complexas (compra, venda e preço médio);
- cálculo de lucro/prejuízo;
- geração de alertas tributários;
- tratamento transacional entre Order e Position.

---

## Resultado

Este PR marca a primeira evolução da aplicação em direção a um processamento orientado a eventos.

Além de introduzir um worker dedicado, a mudança reorganiza responsabilidades entre usecases, entidades e repositories, permitindo que a atualização das posições ocorra de forma desacoplada da criação das ordens e preparando a arquitetura para futuras evoluções distribuídas.