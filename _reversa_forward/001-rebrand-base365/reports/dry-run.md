# Relatório do motor de rebrand (dry-run)

Arquivos de conteúdo alterados: **2001**
Caminhos renomeados: **11**

| # | Regra | Substituições | Exemplo |
|---|-------|---------------|---------|
| 1 | `https://github\.com/nextlevelbuilder/goclaw-docs` → `https://edyocampos.github.io/base365` (repositório de docs dos autores originais) | 31 | `README.md`: `https://github.com/nextlevelbuilder/goclaw-docs`; `_readmes/README.ar.md`: `https://github.com/nextlevelbuilder/goclaw-docs` |
| 2 | `https?://docs\.goclaw\.sh` → `https://edyocampos.github.io/base365` (site de documentação) | 437 | `CHANGELOG.md`: `https://docs.goclaw.sh`; `README.md`: `https://docs.goclaw.sh` |
| 3 | `\bdocs\.goclaw\.sh` → `edyocampos.github.io/base365` (idem, sem esquema) | 33 | `CHANGELOG.md`: `docs.goclaw.sh`; `README.md`: `docs.goclaw.sh` |
| 4 | `https?://fetch\.goclaw\.sh` → `https://fetch.base365.example.com` (serviço de extração dos autores originais (padrão removido à mão na T-extrator)) | 5 | `cmd/gateway_builtin_tools.go`: `https://fetch.goclaw.sh`; `cmd/gateway_builtin_tools.go`: `https://fetch.goclaw.sh` |
| 5 | `\bfetch\.goclaw\.sh` → `fetch.base365.example.com` (idem, sem esquema) | 11 | `internal/tools/web_fetch_extractor_defuddle.go`: `fetch.goclaw.sh`; `internal/tools/web_fetch_extractor_defuddle.go`: `fetch.goclaw.sh` |
| 6 | `https?://goclaw\.sh` → `https://edyocampos.github.io/base365` (site principal) | 16 | `cmd/gateway_providers.go`: `https://goclaw.sh`; `cmd/gateway_providers.go`: `https://goclaw.sh` |
| 7 | `\bgoclaw\.sh\b` → `edyocampos.github.io/base365` (idem, sem esquema) | 0 |  |
| 8 | `goclaw\.tamgiac\.com` → `base365.example.com` (implantação do autor original) | 45 | `cmd/bitrix_portal.go`: `goclaw.tamgiac.com`; `internal/channels/bitrix24/portal_test.go`: `goclaw.tamgiac.com` |
| 9 | `goclaw\.zuey\.me` → `base365.example.com` (implantação do autor original) | 2 | `docs/deployment-guide.md`: `goclaw.zuey.me`; `plans/260609-1713-agent-behavior-ux-overrides/phase-06-validation-and-zuey-beta-handoff.md`: `goclaw.zuey.me` |
| 10 | `ghcr\.io/nextlevelbuilder/goclaw` → `ghcr.io/edyocampos/base365` (registro GHCR (minúsculas obrigatórias)) | 8 | `AGENTS.md`: `ghcr.io/nextlevelbuilder/goclaw`; `CLAUDE.md`: `ghcr.io/nextlevelbuilder/goclaw` |
| 11 | `digitop/goclaw` → `edyocampos/base365` (imagem do Docker Hub (pendência)) | 5 | `AGENTS.md`: `digitop/goclaw`; `CLAUDE.md`: `digitop/goclaw` |
| 12 | `digitopvn/goclaw` → `edyoCampos/base365` (repositório GitHub antigo) | 64 | `docs/deployment-guide.md`: `digitopvn/goclaw`; `docs/journals/260518-1741-beta-skill-grants-tenant-scope.md`: `digitopvn/goclaw` |
| 13 | `nextlevelbuilder/goclaw` → `edyoCampos/base365` (caminho do módulo Go e URLs do repositório) | 3776 | `CHANGELOG.md`: `nextlevelbuilder/goclaw`; `Dockerfile`: `nextlevelbuilder/goclaw` |
| 14 | `GOCLAW` → `BASE365` (maiúsculas) | 701 | `AGENTS.md`: `GOCLAW`; `AGENTS.md`: `GOCLAW` |
| 15 | `GoClaw` → `Base365` (capitalizada) | 1111 | `AGENTS.md`: `GoClaw`; `CHANGELOG.md`: `GoClaw` |
| 16 | `Goclaw` → `Base365` (capitalização mista) | 30 | `cmd/backup.go`: `Goclaw`; `docs/browser-backends.md`: `Goclaw` |
| 17 | `goclaw` → `base365` (minúsculas) | 2003 | `AGENTS.md`: `goclaw`; `AGENTS.md`: `goclaw` |

## Renomeios

- `ui/desktop/frontend/public/goclaw-icon.svg` → `ui/desktop/frontend/public/base365-icon.svg`
- `ui/desktop/frontend/dist/goclaw-icon.svg` → `ui/desktop/frontend/dist/base365-icon.svg`
- `ui/web/public/goclaw-icon.svg` → `ui/web/public/base365-icon.svg`
- `scripts/zuey/goclaw-deploy.sh` → `scripts/zuey/base365-deploy.sh`
- `scripts/zuey/goclaw-upgrade-release.sh` → `scripts/zuey/base365-upgrade-release.sh`
- `plan/goclaw-mcp-integration.md` → `plan/base365-mcp-integration.md`
- `skills/goclaw` → `skills/base365`
- `_statics/goclaw-logo.svg` → `_statics/base365-logo.svg`
- `_statics/goclaw-icon.svg` → `_statics/base365-icon.svg`
- `_statics/goclaw.png` → `_statics/base365.png`
- `compose.d/00-goclaw.yml` → `compose.d/00-base365.yml`

## Revisão por categoria (T018)

| Categoria | Conclusão |
|-----------|-----------|
| Módulo Go (regra 13) | 3.776 trocas, sufixos conferidos: `goclaw`, `goclaw.git`, `goclaw-docs`, `goclaw-cli`, `goclaw-web`. Nenhum repositório de terceiros da mesma organização é afetado (só `/goclaw*`). |
| URLs externas (regras 1 a 9) | Mapeamento correto. Sobra a regra 7 sem uso (0), mantida por segurança. |
| Imagens (regras 10 a 12) | `ghcr.io/nextlevelbuilder/goclaw-web` vira `ghcr.io/edyocampos/base365-web` (variante web mantida). |
| Genérica (regras 14 a 17) | Identificadores colados (`goclawprotocol`, `goclawKey`, `goclawGID`, `dotgoclaw_*`) viram `base365protocol` etc., válidos em Go e TS. |
| Fora da regra "goclaw" | `X-AIMLAPI-Partner-ID: nextlevelbuilder` (`internal/providers/aimlapi.go:28`) é um identificador de parceiro dos autores originais e **não** é alterado pelo motor. `nextlevelbuilder/goclaw-cli` (CLI separada) vira `edyoCampos/base365-cli`, repositório que ainda não existe. Ambos vão para `docs/rebrand-pendencias.md`. |

Nenhuma regra precisou de ajuste.
