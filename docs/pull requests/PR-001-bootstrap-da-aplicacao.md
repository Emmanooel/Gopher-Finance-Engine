# PR-001 - Bootstrap da aplicação

**Status:** Merged

**Tipo:** Feature

**Commit:** `8eff3fc`

**Data:** 11/02/2026

---

## Contexto

O desenvolvimento deste projeto foi pensado desde o início em pequenas entregas incrementais.

Antes de implementar qualquer regra de negócio, era necessário garantir que a aplicação possuísse uma base estável sobre a qual as demais funcionalidades pudessem ser construídas.

Por esse motivo, a primeira entrega não foi uma funcionalidade do produto, mas sim o bootstrap da aplicação: uma API capaz de iniciar corretamente, responder a verificações de saúde e fornecer toda a infraestrutura necessária para a evolução do sistema.

A intenção era estabelecer um primeiro marco de desenvolvimento, permitindo que as próximas implementações fossem construídas de forma incremental e sempre sobre uma aplicação executável.

---

## Decisões desta entrega

Nesta etapa foi priorizada a construção da infraestrutura mínima da aplicação.

Nenhuma regra de negócio foi implementada intencionalmente.

O foco desta entrega foi garantir que:

- a aplicação pudesse ser executada;
- a estrutura do projeto estivesse organizada;
- os componentes básicos de infraestrutura estivessem disponíveis;
- as próximas funcionalidades pudessem ser implementadas incrementalmente sobre uma base estável.

---

## Objetivo

Estabelecer a primeira entrega técnica do projeto, criando uma aplicação executável, organizada e preparada para receber as funcionalidades previstas no backlog.

---

## Alterações

### Estrutura da aplicação

- Criação da estrutura inicial do projeto em Go.
- Organização dos diretórios principais.
- Separação dos pontos de entrada da aplicação (`cmd/api` e `cmd/worker`).

### Configuração

- Criação do carregamento de variáveis de ambiente.
- Estrutura inicial de configuração da aplicação.
- Estrutura para configuração de conexão com PostgreSQL.

### Infraestrutura

- Configuração inicial do servidor HTTP utilizando Gin.
- Implementação do endpoint `/health`.
- Estrutura inicial para inicialização da aplicação.
- Configuração do logger utilizando Zap.

### Ambiente de execução

- Adição de Dockerfile utilizando multi-stage build.
- Inclusão do `.gitignore`.
- Definição das dependências do projeto (`go.mod`).

---

## Impactos

A aplicação passa a possuir uma estrutura mínima para desenvolvimento.

Com esta entrega torna-se possível:

- iniciar a API;
- validar se a aplicação está saudável através do Health Check;
- evoluir novas funcionalidades sobre uma base organizada.

---

## Próximos passos

- Implementar autenticação.
- Criar os primeiros casos de uso.
- Implementar persistência dos dados.
- Disponibilizar os primeiros endpoints de negócio.

---

### Observações

Durante o bootstrap já foi prevista a existência de um Worker independente da API.

Embora ainda sem implementação, a separação foi criada desde o início para permitir o processamento assíncrono de ordens em futuras entregas.

---

## Histórico

Esta Pull Request foi reconstruída retroativamente a partir do histórico de commits do projeto.

Seu objetivo é preservar o contexto técnico e registrar a evolução da aplicação ao longo do desenvolvimento.