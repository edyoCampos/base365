# Actions: Rebrand GoClaw → Base365

> Identificador: `001-rebrand-base365`
> Data: `2026-10-02`
> Roadmap: `_reversa_forward/001-rebrand-base365/roadmap.md`

## Resumo

| Métrica | Valor |
|---------|-------|
| Total de ações | 78 |
| Paralelizáveis (`[//]`) | 44 |
| Maior cadeia de dependência | 26 |

**Ordem global:** as fases seguem a numeração, e as ações só começam com as dependências concluídas. Os relatórios citados (`reports/`) ficam em `_reversa_forward/001-rebrand-base365/reports/`. Os portões finais correspondem ao critério de pronto do roadmap (seção 10).

**Bloqueio humano:** a ação T063 exige resposta do usuário (licença). As de remoção de menções (T064 a T067) não rodam sem ela.

**Ações estruturais fora do `actions.md`:** o `.github/` não existe neste checkout, então fluxo de release e CI ficam como pendência (ação T070), não como ação de código.


## Fase 1, Preparação

Rede de segurança (baseline, checagem de marca) e motor de substituição.

| ID | Descrição | Dependências | Paralelismo | Arquivo alvo | Confidência | Status |
|----|-----------|--------------|-------------|--------------|-------------|--------|
| T001 | Criar baseline versionado do projeto (`git init` e commit inicial; se o usuário preferir, `tar` fora do projeto). Pedir confirmação antes de executar. | - | - | `/home/edyo/Projects/base365/` (raiz) | 🟢 | `[X]` |
| T002 | Registrar as contagens iniciais (conteúdo por grafia, nomes, `GOCLAW_*` distintas, `nextlevelbuilder`, `digitop`, `tamgiac`, `zuey`, `orange-`, `amber-`) em `reports/baseline-counts.md`. | T001 | - | `_reversa_forward/001-rebrand-base365/reports/baseline-counts.md` | 🟢 | `[X]` |
| T003 | Arquivar a lista completa de nomes `goclaw_*` das ferramentas MCP (`grep -rhoE '"goclaw_[a-z_]+"'`) e a contagem, antes da troca. | T001 | `[//]` | `_reversa_forward/001-rebrand-base365/reports/mcp-tools-before.txt` | 🟡 | `[X]` |
| T004 | Criar a lista de exceções da checagem com as justificativas: `_reversa_sdd/`, `_reversa_forward/`, `_reversa_docs/`, `.reversa/`, `.claude/`, `.agents/`, o próprio arquivo e o script. | T001 | `[//]` | `scripts/brand-exceptions.txt` | 🟢 | `[X]` |
| T005 | Criar o arquivo de regras ordenadas do motor: (1) módulo, (2) URLs externas e imagens (mapeamento T-08), (3) prefixos de segurança, (4) cabeçalhos, (5) prefixos de contrato (`goclaw_`, `goclaw:`), (6) quatro grafias genéricas. | T001 | `[//]` | `scripts/rebrand/rules.txt` | 🟢 | `[X]` |
| T006 | Criar `scripts/check-brand.sh`: varre conteúdo e nomes sem diferenciar maiúsculas, respeita a lista de exceções, imprime arquivo e linha e sai com código diferente de zero se achar algo. | T004 | - | `scripts/check-brand.sh` | 🟢 | `[X]` |
| T007 | Adicionar o alvo `make check-brand` ao Makefile. | T006 | - | `Makefile` | 🟢 | `[X]` |
| T008 | Rodar `make check-brand` e registrar o estado vermelho inicial, com a contagem, em `reports/baseline-counts.md`. | T007, T002 | - | `_reversa_forward/001-rebrand-base365/reports/baseline-counts.md` | 🟢 | `[X]` |
| T009 | Criar `scripts/rebrand/apply.sh`: aplica as regras em ordem, renomeia caminhos, respeita exceções, tem `--dry-run` com relatório de substituições por regra e é idempotente. | T004, T005 | - | `scripts/rebrand/apply.sh` | 🟢 | `[X]` |

## Fase 2, Testes

Testes de caracterização escritos **antes** da troca, com o prefixo antigo. O motor os converte, e eles precisam continuar verdes.

| ID | Descrição | Dependências | Paralelismo | Arquivo alvo | Confidência | Status |
|----|-----------|--------------|-------------|--------------|-------------|--------|
| T010 | Escrever teste de caracterização do denylist de credenciais de CLI com o prefixo **antigo** (`GOCLAW_X` e `GOCLAW_GATEWAY_TOKEN` rejeitados). O motor o converte para `BASE365_` e ele deve continuar verde. | T001 | `[//]` | `internal/crypto/env_denylist_test.go` | 🟢 | `[X]` |
| T011 | Escrever teste de caracterização da allowlist da workstation (variável `GOCLAW_X` rejeitada, `internal/workstation/security/allowlist.go:189-190`). | T001 | `[//]` | `internal/workstation/security/allowlist_test.go` | 🟢 | `[X]` |
| T012 | Escrever teste de caracterização da remoção de credenciais do ambiente de processo filho (`GOCLAW_GATEWAY_TOKEN` não chega ao filho, `internal/tools/env_scrub.go:23`). | T001 | `[//]` | `internal/tools/env_scrub_test.go` | 🟢 | `[X]` |
| T013 | Estender o teste dos padrões de comando bloqueados com os 4 regex de `GOCLAW_` (`echo`, `printf`, `python os.environ`, `node process.env`, `internal/tools/shell_deny_groups.go:233-236`). | T001 | `[//]` | `internal/tools/shell_deny_test.go` | 🟢 | `[X]` |
| T014 | Escrever teste do espelho web `ENV_DENYLIST_PREFIXES` (`GOCLAW_` rejeitado), usando o framework de teste já presente em `ui/web`. | T001 | `[//]` | `ui/web/src/pages/cli-credentials/cli-credential-grant-env-section.test.ts` | 🟡 | `[X]` |
| T015 | Escrever teste do formato da chave de API (`goclaw_<32hex>`, `internal/crypto/apikey.go:10`, `:12`). O motor o converte para `base365_`. | T001 | `[//]` | `internal/crypto/apikey_test.go` | 🟢 | `[X]` |
| T016 | Rodar a baseline completa **antes** da troca (`go build` PG e `-tags sqliteonly`, `go vet`, `go test ./...`, `pnpm build` de web e desktop) e registrar falhas pré-existentes em `reports/baseline-tests.md`. | T010, T011, T012, T013, T014, T015 | - | `_reversa_forward/001-rebrand-base365/reports/baseline-tests.md` | 🟢 | `[X]` |

## Fase 3, Núcleo

Aplicação das regras, renomeio de caminhos, `gofmt`, builds e correção de fixtures.

| ID | Descrição | Dependências | Paralelismo | Arquivo alvo | Confidência | Status |
|----|-----------|--------------|-------------|--------------|-------------|--------|
| T017 | Rodar o motor em `--dry-run` e gerar o relatório de substituições por regra e por categoria. | T009, T016, T003, T008 | - | `_reversa_forward/001-rebrand-base365/reports/dry-run.md` | 🟢 | `[X]` |
| T018 | Revisar amostras do relatório por categoria (módulo, URLs, segurança, cabeçalhos, contratos, genérica) e ajustar as regras onde a troca estiver errada. | T017 | - | `scripts/rebrand/rules.txt` | 🟢 | `[X]` |
| T019 | Aplicar as regras de conteúdo em todo o repositório (sem renomear caminhos ainda). | T018 | - | `scripts/rebrand/apply.sh` (repositório inteiro) | 🟡 | `[X]` |
| T020 | Renomear os 10 caminhos de fonte com "goclaw" no nome (`compose.d/00-goclaw.yml`, `_statics/goclaw*`, `plan/goclaw-mcp-integration.md`, `skills/goclaw`, `scripts/zuey/goclaw-*.sh`, `ui/web/public/goclaw-icon.svg`, `ui/desktop/frontend/public/goclaw-icon.svg`) e corrigir as referências a eles. O `dist/` é tratado em a ação T076 (build). | T019 | - | `scripts/rebrand/apply.sh` (renomeios) | 🟡 | `[X]` |
| T021 | Rodar `gofmt -w .` e `go fix ./...` (reordena imports e realinha comentários). | T020 | - | `**/*.go` | 🟢 | `[X]` |
| T022 | Compilar com `go build ./...` e corrigir erros de compilação decorrentes da troca. | T021 | `[//]` | `**/*.go` | 🟢 | `[X]` |
| T023 | Compilar com `go build -tags sqliteonly ./...` e corrigir erros de compilação decorrentes da troca. | T022 | - | `**/*.go` | 🟢 | `[X]` |
| T024 | Rodar `go vet ./...` e corrigir os apontamentos. | T022, T023 | - | `**/*.go` | 🟢 | `[X]` |
| T025 | Rodar `go test ./...` e listar falhas novas em comparação com `reports/baseline-tests.md`. | T024 | - | `_reversa_forward/001-rebrand-base365/reports/post-rename-tests.md` | 🟢 | `[X]` |
| T026 | Corrigir fixtures de teste quebrados em `internal/http`, `internal/gateway` e `internal/channels` (assinatura HMAC, snapshots, URLs). Corrige o fixture, nunca a regra. | T025 | `[//]` | `internal/http/*_test.go`, `internal/gateway/*_test.go`, `internal/channels/**/*_test.go` | 🟡 | `[X]` |
| T027 | Corrigir fixtures quebrados em `internal/agent`, `internal/bootstrap` e `internal/tokencount` (limites de tamanho de prompt, contagem de tokens). | T025 | `[//]` | `internal/agent/*_test.go`, `internal/bootstrap/*_test.go`, `internal/tokencount/*_test.go` | 🟡 | `[X]` |
| T028 | Corrigir fixtures quebrados nos demais pacotes listados em `reports/post-rename-tests.md`. | T026, T027 | - | `internal/**/*_test.go`, `cmd/*_test.go`, `tests/**` | 🟡 | `[X]` |
| T029 | Rodar de novo `go test ./...` e confirmar zero falhas novas em relação à baseline. | T026, T027, T028 | - | `_reversa_forward/001-rebrand-base365/reports/post-rename-tests.md` | 🟢 | `[X]` |

## Fase 4, Integração

Segurança, contratos externos, extrator, banco, desktop, ativos e cores.

| ID | Descrição | Dependências | Paralelismo | Arquivo alvo | Confidência | Status |
|----|-----------|--------------|-------------|--------------|-------------|--------|
| T030 | Adicionar testes negativos de segurança (`GOCLAW_X` **não** é tratado como segredo) no denylist de CLI e na allowlist da workstation. | T029 | `[//]` | `internal/crypto/env_denylist_test.go`, `internal/workstation/security/allowlist_test.go` | 🟢 | `[X]` |
| T031 | Adicionar testes negativos de segurança no env scrub e nos padrões de comando bloqueados (`$GOCLAW_ENCRYPTION_KEY` deixa de ser bloqueado, `$BASE365_ENCRYPTION_KEY` é). | T029 | `[//]` | `internal/tools/env_scrub_test.go`, `internal/tools/shell_deny_test.go` | 🟢 | `[X]` |
| T032 | Adicionar teste negativo no espelho web do denylist (prefixo antigo aceito como chave comum, novo rejeitado). | T029 | `[//]` | `ui/web/src/pages/cli-credentials/cli-credential-grant-env-section.test.ts` | 🟡 | `[X]` |
| T033 | Conferir `apiKeyPrefix = "base365_"` e adicionar teste: chave com prefixo `goclaw_` não valida. | T029 | `[//]` | `internal/crypto/apikey.go`, `internal/crypto/apikey_test.go` | 🟢 | `[X]` |
| T034 | Adicionar testes de webhook: assinatura em `X-Base365-Signature` aceita, só `X-GoClaw-Signature` rejeitada, ambos presentes usa o novo. | T029 | `[//]` | `internal/http/webhooks_auth_test.go` | 🟢 | `[X]` |
| T035 | Adicionar teste do campo `model`: `base365:<agentId>` resolve o agente, `goclaw:<agentId>` não (`internal/http/auth.go:61`). | T029 | `[//]` | `internal/http/auth_test.go` | 🟢 | `[X]` |
| T036 | Adicionar teste que lista as ferramentas MCP e confirma que nenhuma começa com `goclaw_` e que a contagem bate com `reports/mcp-tools-before.txt`. | T029, T003 | `[//]` | `internal/mcp/tools_brand_test.go` | 🟡 | `[X]` |
| T037 | Mudar o padrão semeado do `web_fetch` para `defuddle` desativado, sem endereço de terceiros (`cmd/gateway_builtin_tools.go:42` e `:253`). | T029 | - | `cmd/gateway_builtin_tools.go` | 🟢 | `[X]` |
| T038 | Remover a constante `defuddleBaseURL` e ignorar com aviso uma entrada `defuddle` ativa sem `base_url` (nunca chamar endereço padrão). | T037 | - | `internal/tools/web_fetch_extractor_defuddle.go`, `internal/tools/web_fetch_extractor.go` | 🟢 | `[X]` |
| T039 | Testar o extrator: sem configuração usa só extração interna; entrada ativa sem `base_url` é ignorada; com `base_url` configurado o `defuddle` é usado. | T038 | - | `internal/tools/web_fetch_extractor_test.go` | 🟢 | `[X]` |
| T040 | Atualizar o texto de ajuda do extrator na interface web (extrator externo opcional, desativado por padrão) e os 5 locales `tools.json`. | T037 | `[//]` | `ui/web/src/pages/builtin-tools/web-fetch-extractor-chain-form.tsx`, `ui/web/src/i18n/locales/*/tools.json` | 🟢 | `[X]` |
| T041 | Atualizar o texto de ajuda do extrator no desktop e os locales `tools.json` de en, vi, zh e ru. | T037 | `[//]` | `ui/desktop/frontend/src/components/tools/extractor-chain-form.tsx`, `ui/desktop/frontend/src/i18n/locales/*/tools.json` | 🟢 | `[X]` |
| T042 | Conferir que as 4 migrations PG (`000001`, `000020`, `000061`, `000095`), `schema.sql`, `pool.go:17`, `docker-compose.postgres.yml` e `config.go:196` não têm "goclaw", e que `RequiredSchemaVersion` (97) e `SchemaVersion` (60) não mudaram. | T029 | - | `migrations/`, `internal/store/sqlitestore/schema.sql`, `internal/upgrade/version.go` | 🟡 | `[X]` |
| T043 | Teste de integração PG: banco novo sem "goclaw" em nenhuma linha semeada e com `web_fetch` com `defuddle` desativado. | T042, T039 | `[//]` | `tests/integration/rebrand_seed_test.go` | 🟡 | `[X]` |
| T044 | Teste de integração SQLite: banco novo sem "goclaw" em linha semeada e com `web_fetch` com `defuddle` desativado. | T042, T039 | `[//]` | `internal/store/sqlitestore/rebrand_seed_test.go` | 🟡 | `[X]` |
| T045 | Atualizar os metadados do desktop: nome, `outputfilename` (`base365-lite`), empresa, produto, e-mail provisório, `info.json`, identidade `com.base365.lite` do manifesto. | T029 | `[//]` | `ui/desktop/wails.json`, `ui/desktop/build/windows/info.json`, `ui/desktop/build/windows/wails.exe.manifest` | 🟡 | `[X]` |
| T046 | Conferir título da janela (`ui/desktop/main.go:21`), serviço do keyring (`keyring.go:17`) e diretório `~/.base365/secrets` (`keyring.go:73`), com teste do caminho. | T029 | `[//]` | `ui/desktop/main.go`, `ui/desktop/keyring.go` | 🟢 | `[X]` |
| T047 | Conferir `githubRepo = "edyoCampos/base365"` (`internal/updater/updater.go:25`) e os nomes temporários `base365-lite.app` e `.exe`, com teste. Registrar na pendência que o fluxo de release precisa publicar os mesmos nomes. | T029 | `[//]` | `internal/updater/updater.go` | 🟡 | `[X]` |
| T048 | Conferir o nome padrão `base365-gateway` (`internal/tracing/otelexport/exporter.go:26`, `:45`, `internal/config/config.go:543`) e o teste `exporter_test.go`. | T029 | `[//]` | `internal/tracing/otelexport/exporter.go` | 🟢 | `[X]` |
| T049 | Validar scripts e compose após a troca: `bash -n` em `install.sh`, `install-lite.sh`, `setup-docker.sh`, `prepare-env.sh`, `scripts/zuey/*.sh`; `docker compose config` nos compose; `make build` gera `base365`. | T029 | - | `scripts/`, `docker-compose*.yml`, `compose.d/`, `Makefile` | 🟡 | `[X]` |
| T050 | Copiar para `ui/web/public/` o favicon e o símbolo oficiais (`/home/edyo/Downloads/Base365 logo e ícone/export/`) e atualizar as referências no `index.html`. | T029 | `[//]` | `ui/web/public/`, `ui/web/index.html` | 🟢 | `[X]` |
| T051 | Copiar o símbolo oficial para `ui/desktop/frontend/public/` e atualizar as referências. | T029 | `[//]` | `ui/desktop/frontend/public/` | 🟢 | `[X]` |
| T052 | Substituir `_statics/base365-logo.svg`, `_statics/base365-icon.svg` e `_statics/base365.png` pelos ativos oficiais (logo horizontal e símbolo). | T020 | `[//]` | `_statics/` | 🟢 | `[X]` |
| T053 | Gerar `appicon.png`, `icon.ico` e `iconfile.icns` do desktop a partir de `icone-app-1024.png` com Pillow. | T029 | `[//]` | `ui/desktop/build/appicon.png`, `ui/desktop/build/windows/icon.ico`, `ui/desktop/build/darwin/iconfile.icns` | 🟡 | `[X]` |
| T054 | Aplicar os tokens de cor do PDF (V1) no tema claro e escuro da web, mais os badges de status. | T029 | `[//]` | `ui/web/src/index.css` | 🟢 | `[ ]` |
| T055 | Aplicar os tokens de cor do PDF (V1) no tema claro e escuro do desktop (abre no escuro por padrão). | T029 | `[//]` | `ui/desktop/frontend/src/index.css` | 🟢 | `[ ]` |
| T056 | Listar todas as ocorrências de `orange-*` e `amber-*` e classificar cada uma como aviso ou decoração pela regra do PDF, em `docs/rebrand-amber-review.md`. | T054, T055 | - | `docs/rebrand-amber-review.md` | 🟡 | `[ ]` |
| T057 | Converter as classes `orange-*` da web para os tokens da paleta. | T056 | `[//]` | `ui/web/src/**` | 🟡 | `[ ]` |
| T058 | Converter as classes `orange-*` do desktop para os tokens da paleta. | T056 | `[//]` | `ui/desktop/frontend/src/**` | 🟡 | `[ ]` |
| T059 | Converter as classes `amber-*` decorativas de `ui/web/src/components/**` conforme a tabela (aviso fica no token `warning`). | T057 | - | `ui/web/src/components/**` | 🟡 | `[ ]` |
| T060 | Converter as classes `amber-*` decorativas de `ui/web/src/pages/**` conforme a tabela. | T059 | - | `ui/web/src/pages/**` | 🟡 | `[ ]` |
| T061 | Converter as classes `amber-*` decorativas do desktop conforme a tabela. | T058 | - | `ui/desktop/frontend/src/**` | 🟡 | `[ ]` |
| T062 | Criar e rodar um verificador de contraste dos pares texto/fundo aplicados (mínimo 4,5:1 e 3:1 para foco, borda de campo e gráficos), nos dois temas. | T054, T055 | `[//]` | `scripts/rebrand/contrast-check.mjs` | 🟢 | `[ ]` |

## Fase 5, Polimento

Licença, remoção de menções, infraestrutura herdada, pendências e portões finais.

| ID | Descrição | Dependências | Paralelismo | Arquivo alvo | Confidência | Status |
|----|-----------|--------------|-------------|--------------|-------------|--------|
| T063 | **Bloqueio humano (R-02):** perguntar ao usuário se a aquisição cobre uso comercial e a remoção dos créditos a projetos de origem, dado o `LICENSE` CC BY-NC 4.0. Registrar a resposta em `reports/licenca-confirmacao.md`. Sem resposta, não executar as ações de remoção. | - | - | `_reversa_forward/001-rebrand-base365/reports/licenca-confirmacao.md` | 🔴 | `[ ]` |
| T064 | Remover "A Go port of OpenClaw" da descrição da CLI (`cmd/root.go:27`). | T063, T029 | `[//]` | `cmd/root.go` | 🟡 | `[ ]` |
| T065 | Remover comentários de procedência de OpenClaw, ZeroClaw e PicoClaw em `internal/` e `pkg/`. | T063, T029 | `[//]` | `internal/**`, `pkg/**` | 🟡 | `[ ]` |
| T066 | Remover as menções nos agradecimentos do `README.md` e dos READMEs traduzidos em `_readmes/`. | T063, T029 | `[//]` | `README.md`, `_readmes/*.md` | 🟡 | `[ ]` |
| T067 | Remover as menções restantes em `docs/`, `plan/` e `plans/`, e conferir que a busca por `openclaw|zeroclaw|picoclaw` retorna 0. | T063, T029 | `[//]` | `docs/**`, `plan/**`, `plans/**` | 🟡 | `[ ]` |
| T068 | Trocar `digitop`, `tamgiac` e `zuey` por domínios de exemplo nos testes (portais Bitrix, snapshot de URL pública, prompt do sistema). | T065 | - | `internal/channels/bitrix24/*_test.go`, `internal/gateway/*_test.go`, `internal/agent/systemprompt_target_test.go` | 🟡 | `[ ]` |
| T069 | Tratar `digitop`, `tamgiac` e `zuey` em `docs/deployment-guide.md` e `scripts/zuey/*.sh`: trocar por valores genéricos e listar o que é infraestrutura real do usuário em `docs/rebrand-pendencias.md`. | T029 | `[//]` | `docs/deployment-guide.md`, `scripts/zuey/*.sh` | 🟡 | `[ ]` |
| T070 | Criar `docs/rebrand-pendencias.md` (RF-17): domínio e site, e-mail de contato, extrator próprio (apontando para `_reversa_forward/backlog/extrator-proprio-base365.md`), Docker Hub, release desktop e fluxo de CI, bots nos provedores, URLs de instalação, GitHub Pages. Cada item com recurso e impacto. | T069, T047 | - | `docs/rebrand-pendencias.md` | 🟢 | `[ ]` |
| T071 | Adicionar à checagem uma lista de aviso (sem falhar) para `digitop`, `tamgiac`, `zuey` e `nextlevelbuilder`. | T006, T068, T069 | - | `scripts/check-brand.sh` | 🟡 | `[ ]` |
| T072 | Documentar `make check-brand` e o gancho sugerido para CI no `CONTRIBUTING.md`, já que `.github/` não existe aqui. | T071 | `[//]` | `CONTRIBUTING.md` | 🟢 | `[ ]` |
| T073 | Escrever o teste da própria checagem: arquivo temporário com "GoClaw" faz a checagem falhar e apontar arquivo e linha; removido o arquivo, passa. | T071 | `[//]` | `scripts/check-brand.test.sh` | 🟢 | `[ ]` |
| T074 | Portão: `make check-brand` verde, com conteúdo e nomes (RF-01, RF-02, RF-22, RF-23). | T073, T070, T064, T065, T066, T067, T050, T051, T052, T057, T058, T059, T060, T061 | - | `scripts/check-brand.sh` | 🟢 | `[ ]` |
| T075 | Portão Go: `go fix`, `go build` PG e `-tags sqliteonly`, `go vet`, `go test -race ./...` e `go test ./internal/i18n/` (paridade de chaves). | T074, T030, T031, T033, T034, T035, T036, T039, T046, T048, T047, T049, T043, T044 | - | `**/*.go` | 🟢 | `[ ]` |
| T076 | Portão interface: `pnpm build` de `ui/web` e de `ui/desktop/frontend` (regenera `dist/` e remove o `goclaw-icon.svg` antigo), `make check-brand` de novo, e `go build -tags sqliteonly` do desktop com o `dist/` novo. | T075, T032, T062, T040, T041, T053, T045 | - | `ui/web/dist`, `ui/desktop/frontend/dist` | 🟡 | `[ ]` |
| T077 | Portão integração: `go test -tags integration ./tests/integration/` com PostgreSQL pgvector na porta 5433, e `make test-invariants`. | T075 | `[//]` | `tests/integration/` | 🟡 | `[ ]` |
| T078 | Mover `scripts/rebrand/` para `_reversa_forward/001-rebrand-base365/artifacts/` (fica fora do produto) e confirmar que `make check-brand` continua verde. | T076, T077 | - | `scripts/rebrand/` | 🟢 | `[ ]` |

## Notas de execução

- As ações `[//]` só são paralelas dentro da mesma fase e depois que as dependências estão `[X]`. Quando duas ações `[//]` pedem alteração nos mesmos arquivos de teste (ex.: as de segurança da Fase 4), execute em sequência.
- As ações de conversão de cores (Fase 4) trabalham sobre a tabela de `docs/rebrand-amber-review.md`. Cada ação é grande em quantidade de arquivos, mas é uma única regra aplicada mecanicamente.
- A ação de portão final de interface regenera `dist/`. O `dist/` nunca é editado à mão.

## Histórico de alterações

| Data | Alteração | Autor |
|------|-----------|-------|
| 2026-10-02 | Versão inicial gerada por `/reversa-to-do` | reversa |
