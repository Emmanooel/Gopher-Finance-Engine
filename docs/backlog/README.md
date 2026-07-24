# Product Backlog

Este diretório contém o backlog funcional do projeto.

O objetivo deste backlog é documentar **o que o sistema precisa entregar**, sem definir detalhes de implementação.

Cada Epic representa um conjunto de funcionalidades relacionadas ao mesmo contexto de negócio.

Dentro de cada Epic existem Issues que descrevem uma funcionalidade específica e seus critérios de aceite.

---

# Estrutura

```
backlog/

    epics/

        authentication/

        users/

        orders/

        positions/

        processing/
```

---

# Fluxo

Cada funcionalidade segue o seguinte ciclo de desenvolvimento.

```
Epic
    ↓
Issue
    ↓
Refinamento
    ↓
Implementação
    ↓
Pull Request
    ↓
ADR (quando necessário)
    ↓
Merge
```

---

# Convenções

As Issues devem responder apenas:

> O que precisa ser desenvolvido?

Detalhes de implementação não pertencem ao backlog.

Questões arquiteturais devem ser registradas em `/docs/adr`.

Alterações realizadas devem ser registradas em `/docs/prs`.