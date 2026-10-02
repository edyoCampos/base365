# Relatório do motor de rebrand

Arquivos de conteúdo alterados: **2001**
Caminhos renomeados: **0**

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

