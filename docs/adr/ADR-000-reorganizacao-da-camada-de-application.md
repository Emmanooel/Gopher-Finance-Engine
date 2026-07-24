# ADR-000 - Reorganização da camada de Application

**Status:** Accepted

**Data:** 24/07/2026

---

# Contexto

Durante a implementação do motor de processamento de ordens surgiu uma dúvida recorrente sobre a separação das responsabilidades dentro da arquitetura.

Até então, as interfaces dos casos de uso estavam localizadas dentro da camada de Domain, enquanto as implementações permaneciam na camada de Application.

Embora funcional, essa organização começou a gerar desconforto durante a evolução do projeto, principalmente quando novos casos de uso passaram a surgir.

Ao invés de continuar implementando funcionalidades sobre uma estrutura que já gerava dúvidas, foi decidido interromper momentaneamente o desenvolvimento para investigar a organização da arquitetura.

---

# Achados

## O domínio não deve conhecer os casos de uso da aplicação

Durante a análise foi identificado que o domínio representa apenas as regras de negócio centrais do sistema.

Casos de uso pertencem à camada de Application e representam a orquestração dessas regras.

Dessa forma, manter interfaces de Application dentro de Domain cria uma dependência conceitual desnecessária.

---

## Os casos de uso estavam acumulando responsabilidades

Inicialmente existia um único arquivo por contexto (`users.go`, `orders.go`, `positions.go`).

Essa organização parecia suficiente no início do projeto, porém conforme novas funcionalidades surgiam, os arquivos passaram a concentrar múltiplas responsabilidades.

Esse comportamento contrariava o princípio de responsabilidade única (SRP), dificultando a evolução do código e reduzindo sua coesão.

---

## O Worker estava acoplado incorretamente ao caso de uso

A implementação inicial fazia com que o caso de uso de criação de ordens acionasse diretamente um Worker.

Durante a investigação foi observado que essa abordagem descaracteriza o papel do Worker, transformando-o apenas em uma chamada indireta.

Em uma arquitetura orientada a eventos, o caso de uso deve concluir sua responsabilidade de negócio e publicar um evento de domínio. A responsabilidade pelo processamento posterior pertence ao consumidor desse evento, e não ao próprio caso de uso.

Embora o mecanismo de eventos ainda não esteja implementado neste projeto, a arquitetura passou a ser preparada para essa evolução.

---

# Considerações

Como resultado da investigação, a camada de Application foi reorganizada por módulos de negócio.

Cada módulo passou a possuir seu próprio contrato, implementações e ações específicas, permitindo que novas funcionalidades sejam adicionadas sem aumentar continuamente um único arquivo.

Essa reorganização também aproxima contratos e implementações, melhora a coesão da camada de Application e prepara a arquitetura para um modelo de processamento assíncrono baseado em eventos.

Mais do que uma mudança estrutural, esta decisão representa uma evolução no entendimento sobre as responsabilidades entre Domain e Application dentro da arquitetura adotada neste projeto.

---
