# Interface: CLI, variáveis de ambiente e diretórios

> Identificador: `001-rebrand-base365`
> Data: `2026-10-02`
> Base: `_reversa_sdd/config/`, `_reversa_sdd/deployment.md`
> Tipo: arquivo / CLI | Confidência: 🟢

## 1. Executável e CLI (RF-04)

| Antes | Depois |
|-------|--------|
| binário `goclaw` (`Makefile:3`, `BINARY = goclaw`) | `base365` |
| desktop `goclaw-lite` (`ui/desktop/wails.json`) | `base365-lite` |
| comando `./goclaw onboard`, `./goclaw migrate up` | `./base365 onboard`, `./base365 migrate up` |
| descrição (`cmd/root.go:27`) | "Base365: multi-agent AI platform…" (sem "A Go port of OpenClaw") |

## 2. Variáveis de ambiente (RF-05)

152 variáveis distintas com prefixo `GOCLAW_` (medição em 2026-10-02) passam a `BASE365_`. Exemplos mais usados: `GOCLAW_ENCRYPTION_KEY`, `GOCLAW_GATEWAY_TOKEN`, `GOCLAW_POSTGRES_DSN`, `GOCLAW_CONFIG`, `GOCLAW_SERVER`, `GOCLAW_HOST`, `GOCLAW_PORT`, `GOCLAW_DATA_DIR`, `GOCLAW_WORKSPACE`, `GOCLAW_SSH_PORT`, `GOCLAW_TRACE_VERBOSE`.

Variáveis com prefixo antigo deixam de ser lidas e não geram aviso de migração (RN-02). Locais: `internal/config/config.go`, `cmd/`, exemplos de `.env`, `prepare-env.sh`, `docker-compose*.yml`, `compose.d/`, `scripts/`, documentação, valores `secretEnv` da web (`ui/web/src/pages/config/sections/channels-section.tsx`).

## 3. Diretórios e arquivos (RF-06)

| Antes | Depois |
|-------|--------|
| `~/.goclaw/` (dados, `workspace`, `secrets`, `skills-store`) | `~/.base365/` |
| `/opt/goclaw/current`, `/srv/goclaw/workspace`, `/usr/local/bin/goclaw-*`, `/app/goclaw` | caminhos equivalentes `base365` |
| `goclaw.db` | `base365.db` |
| keyring `goclaw-desktop` (`ui/desktop/keyring.go:17`) | `base365-desktop` |
| `compose.d/00-goclaw.yml` | `compose.d/00-base365.yml` |
| `scripts/zuey/goclaw-*.sh` | `scripts/zuey/base365-*.sh` |

A pasta `.goclaw` aparece também na proteção de caminhos do `filesystem` (`internal/tools/filesystem.go:29`, `:454`) e nos padrões de bloqueio do shell. Esses pontos também mudam e têm teste (`internal/tools/shell_deny_test.go:192`).

## 4. Imagens de contêiner

| Antes | Depois (provisório) |
|-------|---------------------|
| `ghcr.io/nextlevelbuilder/goclaw` | `ghcr.io/edyocampos/base365` |
| `digitop/goclaw` (Docker Hub) | `edyocampos/base365` (pendência) |
| `goclaw-sandbox:bookworm-slim` | `base365-sandbox:bookworm-slim` |
| `goclaw-postgres`, serviço `goclaw` no compose | `base365-postgres`, `base365` |

## 5. Verificação

Teste de configuração com apenas `GOCLAW_CONFIG` definida: o arquivo não é carregado. Teste análogo com `BASE365_CONFIG`: é carregado.
