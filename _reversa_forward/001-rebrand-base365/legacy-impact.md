# Legacy impact: Rebrand GoClaw → Base365

> Identificador: `001-rebrand-base365`
> Data: `2026-10-02`
> Cenário: legado (âncora: `_reversa_sdd/architecture.md` e `_reversa_sdd/domain.md`)
> Política de edição do legado na execução: `allowLegacyEdits: true`, `allowedPaths: []` (projeto inteiro liberado, liberação irrestrita).
> Escopo medido: `git diff --shortstat 90bd368 HEAD` → **2.322 arquivos alterados** (11.880 inserções, 10.904 remoções), mais as alterações ainda não commitadas desta rodada.
> Execução parcial: 76 de 78 ações concluídas. A T077 (integração PostgreSQL) está bloqueada por ambiente, e a T078 (arquivar o motor) depende dela.

## 1. Arquivos afetados por componente

| Arquivo afetado | Componente | Tipo | Severidade | Justificativa |
|-----------------|------------|------|------------|---------------|
| `go.mod` e 1.552 arquivos `.go` | todo o código Go | regra-alterada | MEDIUM | Caminho do módulo `github.com/edyoCampos/base365`. Build e testes provam a consistência. |
| `internal/crypto/apikey.go`, `internal/crypto/env_denylist.go` | `crypto` | delta-de-contrato-externo | HIGH | Chave de API `base365_<32hex>`. Prefixo `BASE365_` no denylist de credenciais de CLI. Chaves antigas deixam de validar. |
| `internal/workstation/security/allowlist.go`, `internal/tools/env_scrub.go`, `internal/tools/shell_deny_groups.go`, `ui/web/.../cli-credential-grant-env-section.tsx` | `workstation`, `tools` (controles de segurança) | regra-alterada | HIGH | Os 4 controles de segurança por prefixo passaram a `BASE365_`. Testes de caracterização antes e negativos depois comprovam cobertura igual. |
| `internal/http/auth.go`, `internal/http/webhooks_auth.go`, `internal/http/gateway_upgrade.go` | `http`, `webhooks` | delta-de-contrato-externo | HIGH | Cabeçalhos `X-Base365-*` (User-Id, Tenant-Id, Signature, Sender-Id, Agent-Id, Agent, Upgrade-Token) e prefixo `base365:<agentId>` no campo `model`. Cabeçalhos antigos deixam de ser aceitos. |
| `internal/mcp/*` | `mcp` | delta-de-contrato-externo | HIGH | 178 ferramentas CRUD com prefixo `base365_`, nome do servidor `base365`. |
| `internal/config/*`, `cmd/*`, `.env` de exemplo, `docker-compose*.yml`, `compose.d/00-base365.yml`, `scripts/*` | `config`, CLI, deploy | delta-de-contrato-externo | HIGH | 152 variáveis `BASE365_*`, diretório `~/.base365/`, banco `base365.db`, binário `base365`. Variáveis antigas deixam de ser lidas. |
| `cmd/gateway_builtin_tools.go`, `internal/tools/web_fetch_extractor*.go` | `tools` (`web_fetch`) | regra-alterada | MEDIUM | O extrator externo `defuddle` vem desativado e sem endereço. Uma entrada ativa sem `base_url` é ignorada. |
| `internal/bootstrap/templates/*`, `internal/bootstrap/seed_store.go`, `internal/knowledgegraph/extractor_prompt.go` | `bootstrap`, `knowledgegraph` | regra-alterada | MEDIUM | Textos enviados ao modelo dizem Base365. |
| `migrations/000001`, `000020`, `000061`, `000095`, `internal/store/sqlitestore/schema.sql`, `internal/store/pg/pool.go` | `store` | delta-de-dados | LOW | Só comentários, um valor semeado (`_base365_recovery`) e o nome do pool. Versões de schema inalteradas (97 e 60). |
| `internal/i18n/catalog_*.go`, `ui/web/src/i18n/**`, `ui/desktop/frontend/src/i18n/**` | `i18n` | regra-alterada | LOW | Valores trocados. Chaves intactas, teste de paridade verde. |
| `internal/tracing/otelexport/exporter.go` | `tracing` | regra-alterada | LOW | Serviço padrão `base365-gateway`, constante `DefaultServiceName`. |
| `internal/updater/updater.go`, `ui/desktop/wails.json`, `ui/desktop/build/**`, `ui/desktop/keyring.go`, `ui/desktop/main.go` | `updater`, app desktop | regra-alterada | MEDIUM | Repositório `edyoCampos/base365`, metadados, ícones, keyring `base365-desktop`. O fluxo de release que publica os artefatos não existe neste checkout. |
| `ui/web/src/index.css`, `ui/desktop/frontend/src/index.css`, 144 arquivos de componentes | design system | regra-alterada | MEDIUM | Tokens de cor do PDF V1, orange/amber convertidos pela regra do âmbar (tabela em `docs/rebrand-amber-review.md`). |
| `_statics/*`, `ui/web/public/*`, `ui/desktop/frontend/public/*`, favicons | ativos visuais | regra-alterada | LOW | Ativos oficiais da marca. |
| `README.md`, `_readmes/*`, `docs/**`, `CLAUDE.md`, `AGENTS.md`, `CONTRIBUTING.md` | documentação | regra-alterada | LOW | Nome, endereços provisórios, créditos a projetos de origem removidos. |
| `scripts/check-brand.sh`, `scripts/brand-exceptions.txt`, `scripts/check-brand.test.sh`, `scripts/rebrand/*`, `docs/rebrand-pendencias.md` | checagem de marca | componente-novo | LOW | Rede de segurança contra o nome antigo. |

## 2. Diff conceitual por componente

- **crypto e segurança:** a regra "variáveis do produto não vazam para processos filhos, para credenciais de CLI nem para a workstation" continua igual, só com o prefixo novo. Os dados criptografados não são afetados, porque nenhum sal ou contexto de derivação usa o nome.
- **http e webhooks:** as regras de autenticação (HMAC, tolerância de 300 s, cache de nonce, resolução de tenant) não mudam. Muda o vocabulário do contrato: cabeçalhos e prefixo do `model`.
- **mcp:** mesmas ferramentas, nomes novos.
- **config e deploy:** mesmas variáveis, prefixo novo; o diretório de dados muda.
- **web_fetch:** o comportamento padrão passa a ser extração interna. Antes, o primeiro extrator chamava um serviço dos autores originais.
- **design system:** a identidade de cor passa de laranja (matiz 38) para violeta. A semântica de aviso permanece âmbar.

## 3. Preservadas (regras 🟢 do `_reversa_sdd/domain.md` que continuam intactas)

- Isolamento multi-tenant: agentes, sessões, memória, times e provedores por tenant.
- Resolução de tenant: um tenant não resolvido cai no master tenant sem erro (RN-T08), agora lido de `X-Base365-Tenant-Id`.
- Lanes de concorrência (`main` 30, `subagent` 50, `team` 100, `cron` 30). Só o prefixo da variável muda, de `GOCLAW_LANE_*` para `BASE365_LANE_*`.
- Autenticação do webhook por HMAC com tolerância de 300 s e cache de nonce.
- Pipeline de 8 estágios, memória em 3 camadas, Knowledge Vault, DomainEventBus.
- Esquemas PostgreSQL e SQLite: nenhuma tabela, coluna ou índice mudou.

## 4. Modificadas (regras 🟢 alteradas)

| Regra | Origem | Alteração |
|-------|--------|-----------|
| Chaves de API seguem `goclaw_<32hex>` | `_reversa_sdd/code-analysis.md#Módulo: crypto` | Passa a `base365_<32hex>`. |
| Variáveis de ambiente `GOCLAW_*` e bloqueio do prefixo em credenciais de CLI | `_reversa_sdd/code-analysis.md#Módulo: crypto`, `domain.md` (Lane) | Passa a `BASE365_*`. |
| Webhooks autenticados por `X-GoClaw-Signature` | `_reversa_sdd/code-analysis.md#Módulo: webhooks`, `flowcharts/webhooks.md` | Passa a `X-Base365-Signature`. |
| `X-GoClaw-Tenant-Id` (RN-T08) e demais cabeçalhos | `_reversa_sdd/domain.md#RN-T08`, `code-analysis.md#Módulo: http` | Passam a `X-Base365-*`. |
| O primeiro extrator do `web_fetch` é um serviço externo dos autores originais | `_reversa_sdd/code-analysis.md#Módulo: tools` | Extração interna por padrão, extrator externo opcional e sem endereço. |
| Atualizador do desktop consulta releases de outro repositório | `_reversa_sdd/code-analysis.md#Módulo: updater` | Consulta `edyoCampos/base365`. |
| Paleta laranja de matiz 38 | `_reversa_sdd/design-system/color-palette.md#1. Identidade da marca` | Paleta violeta do PDF V1. |
| Imagens publicadas em `ghcr.io/nextlevelbuilder/goclaw` e `digitop/goclaw` | `_reversa_sdd/deployment.md#6` | `ghcr.io/edyocampos/base365` e `edyocampos/base365` (provisório). |

A extração em `_reversa_sdd/` ainda descreve a identidade antiga. `/reversa-sync` converge o adendo.

## 5. Fora do escopo declarado e tratado como pendência

Itens de `docs/rebrand-pendencias.md`: domínio, e-mail, Docker Hub, release do desktop e CI (`.github/` ausente), infraestrutura de deploy herdada, CLI separada, identificador de parceiro AIMLAPI, bots nos provedores.
