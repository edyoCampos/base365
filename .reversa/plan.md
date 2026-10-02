# Plano de Exploração — base365

> Criado pelo Reversa em 2026-09-22
> Marque cada tarefa com ✅ quando concluída.
> Você pode editar este plano antes de iniciar: adicione, remova ou reordene tarefas conforme necessário.

---

## Fase 1: Reconhecimento 🔍

- [x] **Scout** — Mapeamento de estrutura de pastas e tecnologias
- [x] **Scout** — Análise de dependências e gerenciadores de pacotes
- [x] **Scout** — Identificação de entry points, CI/CD e configurações

## Decisão de organização das specs 🗂️

> Entre o Scout e o Arqueólogo, o Reversa pergunta como você quer organizar as specs (por módulo, caso de uso, endpoint, híbrida, por features ou customizada). A escolha fica persistida em `.reversa/config.toml` na seção `[specs]` e não será reperguntada em execuções futuras. Para reapresentar o menu, remova manualmente a seção.

## Fase 2: Escavação 🏗️

- [x] **Arqueólogo** — Análise do módulo `agent`
- [x] **Arqueólogo** — Análise do módulo `bootstrap`
- [x] **Arqueólogo** — Análise do módulo `channels` (núcleo + `channels/whatsapp`)
- [x] **Arqueólogo** — Análise do módulo `channels/bitrix24` (~21,4K linhas — maior subpacote de canal)
- [x] **Arqueólogo** — Análise do módulo `channels/telegram`
- [x] **Arqueólogo** — Análise do módulo `channels/feishu`
- [x] **Arqueólogo** — Análise do módulo `channels/zalo` (+ `zalo/personal` + `zalo/personal/protocol`)
- [x] **Arqueólogo** — Análise do módulo `channels/pancake`
- [x] **Arqueólogo** — Análise do módulo `channels/discord`
- [x] **Arqueólogo** — Análise do módulo `channels/slack`
- [x] **Arqueólogo** — Análise do módulo `channels/facebook` (+ `channels/media`, `channels/replycontext`, `channels/typing`)
- [x] **Arqueólogo** — Análise do módulo `config`
- [x] **Arqueólogo** — Análise do módulo `consolidation`
- [x] **Arqueólogo** — Análise do módulo `crypto`
- [x] **Arqueólogo** — Análise do módulo `cron` (+ `cronexec`)
- [x] **Arqueólogo** — Análise do módulo `edition`
- [x] **Arqueólogo** — Análise do módulo `eventbus` (+ `bus`)
- [x] **Arqueólogo** — Análise do módulo `gateway` (+ `gateway/methods`)
- [x] **Arqueólogo** — Análise do módulo `hooks`
- [x] **Arqueólogo** — Análise do módulo `http`
- [x] **Arqueólogo** — Análise do módulo `i18n`
- [x] **Arqueólogo** — Análise do módulo `knowledgegraph`
- [x] **Arqueólogo** — Análise do módulo `mcp`
- [x] **Arqueólogo** — Análise do módulo `memory`
- [x] **Arqueólogo** — Análise do módulo `oauth`
- [x] **Arqueólogo** — Análise do módulo `orchestration`
- [x] **Arqueólogo** — Análise do módulo `permissions`
- [x] **Arqueólogo** — Análise do módulo `pipeline`
- [x] **Arqueólogo** — Análise do módulo `providerresolve` (+ `providers`)
- [x] **Arqueólogo** — Análise do módulo `sandbox`
- [x] **Arqueólogo** — Análise do módulo `scheduler`
- [x] **Arqueólogo** — Análise do módulo `security`
- [x] **Arqueólogo** — Análise do módulo `sessions`
- [x] **Arqueólogo** — Análise do módulo `skills`
- [x] **Arqueólogo** — Análise do módulo `store` (+ `store/base`, `store/pg`, `store/sqlitestore`)
- [x] **Arqueólogo** — Análise do módulo `tasks` (+ `teamworkclassify`, `childrun`)
- [x] **Arqueólogo** — Análise do módulo `tokencount`
- [x] **Arqueólogo** — Análise do módulo `tools`
- [x] **Arqueólogo** — Análise do módulo `tracing`
- [x] **Arqueólogo** — Análise do módulo `tts`
- [x] **Arqueólogo** — Análise do módulo `updater`
- [x] **Arqueólogo** — Análise do módulo `upgrade`
- [x] **Arqueólogo** — Análise do módulo `usage`
- [x] **Arqueólogo** — Análise do módulo `vault`
- [x] **Arqueólogo** — Análise do módulo `webhooks`
- [x] **Arqueólogo** — Análise do módulo `webui`
- [x] **Arqueólogo** — Análise do módulo `workspace` (+ `workstation`)
- [x] **Arqueólogo** — Análise dos módulos de suporte (`audio`, `backup`, `bgalert`, `cache`, `channelmemory`, `heartbeat`, `logs`, `media`, `mediabudget`, `safego`, `systemmessages`, `testutil`, `version`)

## Fase 2b: Escavação Frontend 🖥️ (adicionada em 2026-09-28 — gap identificado pelo usuário)

> O Scout mapeou tecnologias/entry points do frontend em `inventory.md`/`surface.json`, mas a heurística de identificação de módulo só reconheceu pacotes Go (`internal/<domínio>/`). `ui/web/src` e `ui/desktop/frontend/src` nunca viraram units de arqueologia. Agrupamento por domínio funcional aprovado pelo usuário em 2026-09-28.

- [x] **Arqueólogo** — `web-core` (`ui/web/src/adapters,api,lib,stores,schemas,types,constants,data,i18n`)
- [x] **Arqueólogo** — `web-components` (`ui/web/src/components/*`)
- [x] **Arqueólogo** — `web-pages-agents-teams` (`ui/web/src/pages/agents,teams,workstations,nodes`)
- [x] **Arqueólogo** — `web-pages-messaging` (`ui/web/src/pages/chat,channels,contacts,webhooks,pending-messages`)
- [x] **Arqueólogo** — `web-pages-admin` (`ui/web/src/pages/config,setup,tenants-admin,packages,builtin-tools,api-keys,cli-credentials,import-export`)
- [x] **Arqueólogo** — `web-pages-knowledge` (`ui/web/src/pages/knowledge-graph,memory,vault`)
- [x] **Arqueólogo** — `web-pages-observability` (`ui/web/src/pages/activity,events,logs,traces,usage`)
- [x] **Arqueólogo** — `web-pages-automation` (`ui/web/src/pages/cron,hooks,mcp,skills`)
- [x] **Arqueólogo** — `web-pages-outros` (`ui/web/src/pages/login,approvals,sessions,storage,backup-restore,providers,tts,overview`)
- [x] **Arqueólogo** — `desktop-core` (`ui/desktop/frontend/src/api,lib,stores,schemas,types,constants,data,services,i18n`)
- [x] **Arqueólogo** — `desktop-components-core` (`ui/desktop/frontend/src/components/agents,teams,channels,chat,common,layout,ui,onboarding`)
- [x] **Arqueólogo** — `desktop-components-tools` (`ui/desktop/frontend/src/components/builtin-tools,cron,mcp,skills,tools,traces`)
- [x] **Arqueólogo** — `desktop-settings-storage` (`ui/desktop/frontend/src/components/settings,storage,providers`)
- [x] **Redator** — Specs SDD para as 13 units acima (`_reversa_sdd/<unit>/requirements.md,design.md,tasks.md`)
- [x] **Revisor** — Revisão cruzada das 13 units novas + atualização de `confidence-report.md`

## Fase 3: Interpretação 🧠

- [x] **Detetive** — Arqueologia Git e ADRs retroativos (sem git: changelogs/journals/plans)
- [x] **Detetive** — Regras de negócio implícitas e máquinas de estado
- [x] **Detetive** — Matriz de permissões (RBAC/ACL)
- [x] **Arquiteto** — Diagramas C4 (Contexto, Containers, Componentes)
- [x] **Arquiteto** — ERD completo e integrações externas
- [x] **Arquiteto** — Spec Impact Matrix

## Fase 4: Geração 📝

- [x] **Redator** — Specs SDD por componente (40 units, 192 arquivos)
- [x] **Redator** — OpenAPI (`openapi/gateway-api.yaml`, `openapi/http-api.yaml`)
- [x] **Redator** — User Stories (4 fluxos cross-unit)
- [x] **Redator** — Code/Spec Matrix (`traceability/code-spec-matrix.md`)

## Fase 5: Revisão ✅

- [x] **Revisor** — Revisão cruzada de specs (verificação estrutural 40/40 units + consolidação de 91 lacunas)
- [x] **Revisor** — Resolução de lacunas com o usuário (adiada — aquisição sem equipe original, ver `_reversa_sdd/questions.md` e `gaps.md`)
- [x] **Revisor** — Relatório de confiança final (`_reversa_sdd/confidence-report.md`, 65,7%)

---

## Agentes Independentes

> Execute estes agentes quando os recursos estiverem disponíveis — podem rodar em qualquer fase.

- [ ] **Visor** — Análise de interface via screenshots
- [x] **Data Master** — Análise completa do banco de dados (`_reversa_sdd/database/`, 2026-10-01)
- [x] **Design System** — Extração de tokens de design (`_reversa_sdd/design-system/`, 2026-10-01)
- [ ] **Tracer** — Análise dinâmica (requer sistema acessível)

---

## Próximo passo

Após o Time de Descoberta concluir e o `_reversa_sdd/` estar populado, você pode disparar um dos fluxos seguintes:

- `/reversa-migrate`: orquestrador do **Time de Migração** (Paradigm Advisor → Curator → Strategist → Designer → Screen Translator → Inspector). Gera as specs do sistema novo. Saída em `_reversa_sdd/migration/` e `_reversa_sdd/screens/`.
- `/reversa-reconstructor`: gera plano bottom-up para reimplementar o software a partir das specs do legado (uma tarefa por sessão).
