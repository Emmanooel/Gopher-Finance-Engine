# ADR 003 - Centralizar a Regra de Atualização de Posições no Domínio

**Status:** Accepted

**Tipo:** Refatoração Arquitetural

**Data:** 26/07/2026

---

# Contexto

Durante a implementação do fluxo de processamento das ordens foi identificado que a atualização das posições estava sendo realizada diretamente pela camada de persistência através de uma instrução SQL utilizando `INSERT ... ON CONFLICT DO UPDATE`.

Além de persistir os dados, a própria consulta era responsável por decidir quando uma posição deveria ser criada, quando deveria ser atualizada e como recalcular o preço médio do ativo.

Embora funcional, essa abordagem transferia parte das regras de negócio para o banco de dados, tornando a evolução do domínio dependente da implementação da camada de infraestrutura.

---

# Problema

O repositório passou a assumir responsabilidades que pertencem ao domínio da aplicação.

A decisão entre criar uma nova posição ou atualizar uma posição existente faz parte do fluxo de processamento das ordens e depende exclusivamente das regras de negócio da aplicação.

Da mesma forma, o cálculo do preço médio é uma regra do domínio financeiro e não deve estar embutido em uma instrução SQL.

Essa implementação viola o princípio de responsabilidade única (SRP), dificulta testes unitários e aumenta o acoplamento entre domínio e banco de dados.

---

# Objetivo

Garantir que toda decisão relacionada à evolução da carteira do usuário seja executada pela camada de aplicação.

A infraestrutura deve ser responsável apenas pela persistência dos dados, enquanto o domínio decide como esses dados evoluem.

---

# Decisão

Remover a lógica de atualização automática da query de persistência e dividir as responsabilidades entre consulta, criação e atualização.

O fluxo de processamento passará a ser responsável por:

1. Buscar a posição do usuário para determinado ativo.
2. Verificar se a posição já existe.
3. Caso não exista, criar uma nova posição.
4. Caso exista, recalcular quantidade, custo total e preço médio.
5. Persistir a nova versão da posição.

---

# Arquitetura Proposta

```text
                    ProcessOrdersUsecase
                             │
                             ▼
             GetPositionByUserAndSymbol()
                             │
                  ┌──────────┴──────────┐
                  │                     │
            Não encontrada         Encontrada
                  │                     │
                  ▼                     ▼
         BuildPosition()        ApplyOrder()
                  │                     │
                  ▼                     ▼
           SavePosition()      UpdatePosition()
```

Toda a decisão permanece na camada de aplicação.

O repositório apenas executa operações de persistência.

---

# Alterações na Camada de Persistência

A operação atual de persistência deixa de utilizar:

```sql
INSERT ... ON CONFLICT DO UPDATE
```

e passa a possuir operações específicas:

- `GetPositionByUserAndSymbol`
- `SavePosition`
- `UpdatePosition`

### SavePosition

Responsável exclusivamente por criar um novo registro.

Caso já exista uma posição para o mesmo `(user_id, symbol)`, a operação retorna erro de violação de chave única.

Nenhuma regra de atualização será executada automaticamente pelo banco.

### UpdatePosition

Responsável apenas por persistir uma posição previamente calculada pelo domínio.

Nenhum cálculo de preço médio, quantidade ou custo será realizado na camada de infraestrutura.

---

# Responsabilidades

## Repository

Responsável apenas por operações CRUD.

- Buscar posição.
- Inserir posição.
- Atualizar posição.

Nenhuma decisão de negócio deve existir nesta camada.

---

## Use Case

Responsável pela orquestração do processamento.

Fluxo esperado:

1. Buscar posição existente.
2. Verificar existência.
3. Criar nova posição ou atualizar a posição atual.
4. Aplicar todas as regras de negócio necessárias.
5. Persistir o resultado.

Toda a inteligência permanece centralizada no domínio da aplicação.

---

# Benefícios

- Repositórios passam a possuir responsabilidade única.
- O domínio deixa de depender de comportamentos específicos do banco de dados.
- O cálculo do preço médio torna-se testável através de testes unitários.
- Novas regras podem ser adicionadas sem alteração das consultas SQL.
- Facilita futuras implementações de:
  - venda parcial;
  - venda total;
  - split;
  - grupamento;
  - bonificação;
  - dividendos em ações;
  - eventos corporativos.

---

# Trade-offs

- O fluxo passa a realizar uma consulta antes da atualização da posição.
- O caso de uso assume maior responsabilidade de orquestração.
- Há aumento de uma operação de leitura em troca de maior clareza arquitetural e desacoplamento.

---

# Consequências

A evolução da carteira deixa de depender da tecnologia de persistência.

O banco de dados passa a armazenar apenas o estado da aplicação, enquanto toda decisão sobre como esse estado evolui permanece centralizada no domínio.

Essa abordagem fortalece os princípios da Clean Architecture, facilita a manutenção do sistema e permite a evolução das regras financeiras sem necessidade de alterar a camada de infraestrutura.