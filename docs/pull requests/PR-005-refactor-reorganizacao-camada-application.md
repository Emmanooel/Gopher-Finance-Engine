# PR-00 - Refactor: Reorganização da camada de Application

**Status:** Merged

**Tipo:** architetural refactor

**Commit:** `9d7188e`

**Data:** 24/07/2026

---

# Contexto

Durante o desenvolvimento do motor de processamento de ordens, tornou-se evidente que a estrutura da camada de Application não estava escalando bem para novas funcionalidades.

Os casos de uso estavam concentrados em arquivos únicos (`users.go`, `orders.go` e `positions.go`), fazendo com que diferentes responsabilidades fossem acumuladas no mesmo arquivo conforme novas funcionalidades eram adicionadas.

Além disso, as interfaces dos casos de uso estavam localizadas dentro do domínio (`internal/domain/application/usecases`), criando uma dependência conceitual incorreta entre Domain e Application.

Essa refatoração foi realizada antes da continuidade do desenvolvimento do motor de processamento para evitar o crescimento de uma estrutura que já apresentava sinais de baixa coesão.

---

# Decisões

## Organizar a camada de Application por módulos

A estrutura passou a ser organizada por contexto de negócio.

Antes

```
application/
    usecases/
        users.go
        orders.go
        positions.go
```

Depois

```
application/
    usecases/

        users/
            interface.go
            login.go
            create_user.go

        orders/
            interface.go
            create_orders.go
            process_orders.go

        positions/
            interface.go
            positions.go
```

Essa organização permite que cada ação evolua independentemente sem aumentar continuamente um único arquivo.

---

## Aproximar interfaces de seus respectivos módulos

As interfaces dos casos de uso deixaram de ficar dentro da camada de Domain e passaram a ficar junto de suas implementações.

Isso reduz a distância entre contrato e implementação e elimina uma dependência conceitual incorreta entre Domain e Application.

---

## Remover o acionamento direto do Worker

O caso de uso de criação de ordens deixou de acionar diretamente um Worker responsável pelo processamento das posições.

Essa mudança prepara a arquitetura para um modelo baseado em eventos, onde o caso de uso apenas conclui sua responsabilidade e um consumidor independente realiza o processamento posterior.

---

# Impactos

- Organização por contexto de negócio.
- Redução da responsabilidade de arquivos individuais.
- Melhor coesão da camada de Application.
- Estrutura preparada para crescimento de novos casos de uso.
- Preparação para processamento assíncrono baseado em eventos.

---

# Próximos passos

- Implementar o processamento das ordens.
- Definir estratégia de publicação de eventos.
- Evoluir os testes unitários por ação.