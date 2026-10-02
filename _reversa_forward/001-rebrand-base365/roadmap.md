# Roadmap: Rebrand GoClaw → Base365

> Identificador: `001-rebrand-base365`
> Data: `2026-10-02`
> Requirements: `_reversa_forward/001-rebrand-base365/requirements.md`
> Confidência: 🟢 CONFIRMADO, 🟡 INFERIDO, 🔴 LACUNA

## 1. Resumo da abordagem

O rebrand é uma troca mecânica em larga escala, e o risco está nos pontos onde a troca não é mecânica. A abordagem tem quatro eixos.

1. **Rede de segurança primeiro.** O diretório não é um repositório git (verificado: `git status` falha). Antes de qualquer alteração, cria-se um baseline versionado. Em seguida entra a checagem de marca (RF-23), que já nasce falhando e vira o critério de progresso.
2. **Motor de substituição ordenado.** Um script único aplica regras específicas antes das genéricas (caminho do módulo, URLs externas, prefixos de segurança, cabeçalhos e, por fim, as quatro grafias). Depois renomeia os 11 caminhos e roda `gofmt`, porque "base365" tem 7 letras e "goclaw" tem 6, o que desalinha comentários e a ordem dos imports.
3. **Intervenções manuais só onde há semântica.** São quatro: os controles de segurança (RN-03), o extrator (RN-07), os metadados do desktop e as sementes de banco. Cada uma tem teste próprio.
4. **Visual por último.** Primeiro os tokens de cor e os ativos, depois a varredura de classes `orange-*` e `amber-*` com a regra do âmbar. O visual vem depois porque não altera comportamento e é verificável por build.

O resultado é validado por seis portões: checagem de marca, build PostgreSQL, build SQLite, `go vet`, testes, e build das duas interfaces.

## 2. Princípios aplicados

`.reversa/principles.md` não existe neste projeto. Por isso a feature foi avaliada contra as regras vigentes do `CLAUDE.md` (não é reescrita nem atenuação de princípio).

| Regra do `CLAUDE.md` | Como a feature se relaciona | Status |
|-----------------------|------------------------------|--------|
| Migrations dual-DB (PG `migrations/` + SQLite `schema.sql`/`schema.go`) | O texto `goclaw` em comentários e sementes é trocado nos dois motores. Sem coluna nem tabela nova, então a versão não é incrementada (T-07). | respeita |
| i18n: chave nova em todos os catálogos | Não há chave nova. Os valores trocam nos catálogos existentes (backend: en, vi, zh, ko, ru; web: en, vi, zh, ko, ru; desktop: en, vi, zh, ru). O teste de paridade `internal/i18n/git_keys_parity_test.go` continua valendo. | respeita |
| Paridade entre superfícies (gateway, API, web, CLI) | Todas as quatro são afetadas e cobertas (seção 5). | respeita |
| Tenant-scope guards em escrita admin | Nenhum handler de escrita muda de escopo. | n/a |
| Sem testes de carga | Nenhum teste de carga é criado. | respeita |
| Checklist pós-implementação (`go fix`, `go build` PG e SQLite, `go vet`, testes) | Vira o portão P3 do critério de pronto. | respeita |
| Regras mobile da web | A troca de cores não altera layout, viewport nem alvos de toque. | respeita |

## 3. Decisões técnicas

| ID | Decisão | Justificativa | Alternativas descartadas | Confidência |
|----|---------|---------------|--------------------------|-------------|
| T-01 | Criar baseline versionado (`git init` + commit inicial, ou cópia `tar` fora do projeto) **antes** de qualquer alteração. | O diretório não é repositório git, e o RF-23 e o critério de aceite usam commit. Sem baseline, uma regra errada de substituição em 2.000 arquivos não é reversível. | Trabalhar sem baseline. | 🟢 |
| T-02 | Um script de rebrand idempotente, versionado em `scripts/rebrand/`, com regras em arquivo de dados ordenado. Substitui conteúdo e renomeia caminhos. É removido ou arquivado ao final. | Reexecutável depois de correções. Auditável por regra. Evita 8.000 edições manuais. | `sed -i` solto. IDE refactor. | 🟢 |
| T-03 | Ordem das regras: (1) caminho do módulo; (2) URLs e endereços externos; (3) prefixos de segurança; (4) cabeçalhos; (5) formato de chave e prefixos de contrato; (6) as 4 grafias genéricas. | Regra genérica antes das específicas destrói o mapeamento (ex.: `docs.goclaw.sh` virando `docs.base365.sh`, domínio inexistente). | Substituição genérica única. | 🟢 |
| T-04 | Mapeamento de grafias: `goclaw`→`base365`, `GoClaw`→`Base365`, `GOCLAW`→`BASE365`, `Goclaw`→`Base365`. Medição em 2026-10-02: 6.432, 1.111, 687 e 30 ocorrências. | Cobre todas as grafias presentes. Não há variantes com separador (`go-claw` etc.: 0 ocorrências). | — | 🟢 |
| T-05 | Caminho do módulo `github.com/nextlevelbuilder/goclaw` → `github.com/edyoCampos/base365` (3.707 linhas, 1.552 arquivos `.go`), seguido de `gofmt -w` e `go build`. Outras referências a `nextlevelbuilder` (cerca de 108 linhas) são tratadas na regra de URLs (T-08). | A ordem dos imports muda e `gofmt` reordena. O build prova a consistência. | `go mod edit` isolado (não cobre imports). | 🟢 |
| T-06 | O prefixo de segurança passa a `BASE365_` nos quatro controles. Cada ponto ganha teste novo com o prefixo novo, e um teste negativo garante que `GOCLAW_` deixa de ser tratado como segredo. Pontos: `deniedPrefixes` e `isBlockedEnvKey` em `internal/crypto/env_denylist.go` (prefixo, nome `GOCLAW_GATEWAY_TOKEN` na lista de remoção); `internal/workstation/security/allowlist.go`; `internal/tools/env_scrub.go`; os 4 regex em `internal/tools/shell_deny_groups.go`; `ENV_DENYLIST_PREFIXES` em `ui/web/src/pages/cli-credentials/cli-credential-grant-env-section.tsx:35`. | A substituição genérica altera esses pontos sem que ninguém perceba. O teste explícito garante que não houve perda de cobertura. | Confiar só na substituição genérica. | 🟢 |
| T-07 | As migrations existentes são editadas no próprio arquivo, sem migration nova e sem incrementar `RequiredSchemaVersion` (97) nem `SchemaVersion` do SQLite (60). Arquivos PG afetados: `000001`, `000020`, `000061`, `000095` (a chave `_goclaw_recovery` vira `_base365_recovery`). | Ninguém usa em produção (RN-02). O texto muda só em comentário e em um valor semeado. Uma migration nova só pelo nome seria ruído. | Migration nova de renomeação. | 🟡 |
| T-08 | Mapeamento de endereços externos: `docs.goclaw.sh`, `goclaw.sh` e `github.com/nextlevelbuilder/goclaw-docs` → `https://edyocampos.github.io/base365/`; `github.com/nextlevelbuilder/goclaw` → `github.com/edyoCampos/base365`; `ghcr.io/nextlevelbuilder/goclaw` → `ghcr.io/edyocampos/base365`; `digitopvn/goclaw` (Docker Hub) → `edyocampos/base365`; `goclaw.example.com` → `base365.example.com`. Cada endereço provisório entra na lista de pendências (T-14). | O GitHub Pages foi a decisão do usuário (clarify, 2026-10-02). Domínios de exemplo seguem o padrão já usado nos testes. | Domínio `*.example` inventado. | 🟡 |
| T-09 | O extrator `defuddle` passa a vir **desativado** no padrão semeado, sem endereço dos autores originais: `cmd/gateway_builtin_tools.go:42` e `:253` (função `backfillWebFetchSettings`). A constante `defuddleBaseURL` em `internal/tools/web_fetch_extractor_defuddle.go:13` é removida. Uma entrada `defuddle` ativa **sem** `base_url` é ignorada com aviso, nunca chama endereço padrão. O fallback sem configuração já é só extração interna (`internal/tools/web_fetch_extractor.go`, `ResolveExtractorChain`). | Cumpre RN-07 e RF-18 sem depender de serviço inexistente e sem enviar URLs a terceiros. O plano do serviço próprio está em `_reversa_forward/backlog/extrator-proprio-base365.md`. | Apontar para um Worker ainda não criado. | 🟢 |
| T-10 | O formato da chave de API é `base365_<32hex>` em `internal/crypto/apikey.go:10` (constante `apiKeyPrefix`), com comentário e testes atualizados. Chaves antigas deixam de validar (RN-02). | Cumpre RN-04 e RF-08. | Aceitar os dois prefixos (descartado por RN-02). | 🟢 |
| T-11 | Os valores de metadado do desktop mudam em `ui/desktop/wails.json` (nome, `outputfilename`, empresa, produto, e-mail), `ui/desktop/build/windows/info.json`, `ui/desktop/build/windows/wails.exe.manifest` (identidade `com.goclaw.lite`), `ui/desktop/main.go:21` (título), `ui/desktop/keyring.go:17` (serviço do cofre) e `:73` (diretório `.goclaw`). O nome do executável vira `base365-lite`, e o produto `Base365 Lite`. O e-mail de contato recebe valor provisório registrado nas pendências. | Cumpre RF-06 e RF-10. O nome "Lite" já existia e a edição não muda. | Renomear a edição. | 🟡 |
| T-12 | O atualizador consulta `edyoCampos/base365` em `internal/updater/updater.go:25` (constante `githubRepo`). Os nomes temporários `goclaw-lite.app` e `goclaw-lite.exe` (`:228`, `:287`) seguem o novo nome do artefato (`base365-lite.*`). **Pendência cruzada:** o fluxo de release que publica esses artefatos (`release-desktop.yaml`) não existe neste checkout (`.github/` ausente) e precisa usar os mesmos nomes. | Sem isso o atualizador procura um arquivo que a release não publica. | — | 🟡 |
| T-13 | Ativos visuais: a origem é `/home/edyo/Downloads/Base365 logo e ícone/export/`. Substituem `_statics/goclaw*` (renomeados), `ui/web/public/goclaw-icon.svg`, `ui/web/public/favicon.svg`, `ui/desktop/frontend/public/goclaw-icon.svg` e `ui/desktop/build/appicon.png`. Os formatos `icon.ico` e `iconfile.icns` são gerados do PNG de 1024 px com Pillow (disponível, 12.1.1). `ui/desktop/frontend/dist/` é saída de build embutida por `go:embed all:frontend/dist` (`ui/desktop/main.go:14`) e é **regenerada**, nunca editada à mão. | Os ativos oficiais estão no `export/` (SVG e PNG nas variantes clara e escura). Não há ImageMagick nem `icotool` no ambiente. | Editar o `dist/` à mão. | 🟡 |
| T-14 | Criar `docs/rebrand-pendencias.md` (lista RF-17) e `scripts/brand-exceptions.txt` (lista RF-23), cada item com justificativa. Pendências iniciais: domínio e site, e-mail de contato, extrator próprio, registro Docker Hub, release desktop e fluxo de CI, bots nos provedores, URLs de instalação. | Cumpre RF-17, RF-23 e o RNF de manutenibilidade. Origem: `_reversa_sdd/brainstorms/001-rebrand-goclaw-base365/decision.md#Pendências externas a providenciar`. | Só memória do assistente. | 🟢 |
| T-15 | A checagem de marca é um script (`scripts/check-brand.sh`) com alvo `make check-brand`. Varre conteúdo e nomes (sem diferenciar maiúsculas), respeita `scripts/brand-exceptions.txt` e imprime arquivo e linha. É escrita **antes** do motor e começa falhando. Exceções iniciais: `_reversa_sdd/`, `_reversa_forward/`, `_reversa_docs/`, `.reversa/`, `.claude/`, `.agents/` (saída e ferramentas do Reversa, fora do produto; hoje só `_reversa_*` e `.reversa/` contêm o termo), mais o próprio arquivo de exceções e o script. O gancho de CI fica documentado, porque `.github/` não existe aqui. | O estado vermelho→verde mede o progresso. | Busca manual por `grep`. | 🟢 |
| T-16 | Tokens de cor: aplicar os valores do PDF (`Base365 logo e ícone.pdf`, V1) em `ui/web/src/index.css` e `ui/desktop/frontend/src/index.css`, tema claro e escuro, mais os badges de status. Em seguida varrer `orange-*` (49 arquivos) e `amber-*` (124 arquivos) com a regra do âmbar (seção 06 do PDF): aviso fica no token `warning`, decoração vira `primary`/`accent`. A classificação vira uma tabela revisável (`docs/rebrand-amber-review.md`) para não perder decisões. | A regra de decisão veio do clarify. Medição em 2026-10-02: os arquivos citados incluem os de `ui/web/src` e `ui/desktop/frontend/src`. | Trocar todo âmbar por violeta. | 🟡 |
| T-17 | Referências a OpenClaw, ZeroClaw e PicoClaw são removidas (comentários, README, `cmd/root.go:27`), **depois** da verificação de licença descrita no risco R-02. A descrição da CLI vira "Base365: multi-agent AI platform with WebSocket RPC, tool execution, and channel integration." | Decisão do clarify. A medição indica 44, 32 e 32 arquivos. | — | 🟡 |
| T-18 | Nenhuma lista de aliases ou compatibilidade retroativa é criada. Variáveis `GOCLAW_*`, diretório `~/.goclaw`, cabeçalhos `X-GoClaw-*`, prefixo `goclaw:` do campo `model` e as chaves `goclaw_` deixam de existir. | RN-02. | Shims de transição. | 🟢 |

## 4. Premissas

Não restam `[DÚVIDA]` no `requirements.md`. As premissas abaixo vêm da seção "Premissas assumidas" e das lacunas 🟡.

| Premissa | Origem (`requirements.md`) | Risco se errada |
|----------|----------------------------|-----------------|
| O registro de imagens é `ghcr.io/edyocampos/base365`. | Seção 10, premissas | Imagens apontam para um registro que não existe até o primeiro push. |
| O endereço do GitHub Pages é `https://edyocampos.github.io/base365/`. | Seção 10, premissas | Links de documentação quebram até o Pages ser publicado. |
| As migrations podem ser editadas no próprio arquivo. | Seção 10, premissas | Um banco já criado fica com os textos antigos. Como não há produção, o impacto é nulo. |
| Os 30 READMEs traduzidos são mantidos, só com o nome trocado. | Seção 10, premissas | Nenhum. |
| O e-mail de contato fica como pendência, sem valor definitivo. | Seção 10, lacunas | O campo do desktop (`wails.json`) mostra um valor provisório. |

## 5. Delta arquitetural

Nenhum componente ganha ou perde responsabilidade. Mudam identificadores, constantes e valores padrão. A coluna "Superfície" cobre a regra de paridade do `CLAUDE.md`.

| Componente | Arquivo de origem no legado | Superfície | Tipo de mudança | Resumo |
|------------|-----------------------------|------------|-----------------|--------|
| Todo o código Go (módulo) | `_reversa_sdd/inventory.md` | Gateway | regra-alterada | Caminho do módulo e imports nos 1.552 arquivos `.go` (T-05). |
| `crypto` (chaves e denylist) | `_reversa_sdd/code-analysis.md#Módulo: crypto` | Gateway, web | contrato-alterado | Prefixo `base365_` nas chaves, `BASE365_` no denylist (T-06, T-10). |
| `tools` (env scrub, shell deny, `web_fetch`) | `_reversa_sdd/code-analysis.md#Módulo: tools` | Gateway | regra-alterada | Prefixo de segurança e extrator desativado por padrão (T-06, T-09). |
| `workstation/security` | `_reversa_sdd/code-analysis.md` | Gateway | regra-alterada | Allowlist de variáveis com o prefixo novo (T-06). |
| `webhooks` | `_reversa_sdd/code-analysis.md#Módulo: webhooks` | Gateway, API | contrato-alterado | Cabeçalho `X-Base365-Signature` (interface `http-headers-e-contratos.md`). |
| `http` (auth, model prefix) | `_reversa_sdd/code-analysis.md#Módulo: http` | API, web | contrato-alterado | Cabeçalhos `X-Base365-*`, prefixo `base365:<agentId>` no campo `model` (idem). |
| `mcp` (ferramentas expostas) | `_reversa_sdd/code-analysis.md#Módulo: mcp` | API, CLI | contrato-alterado | Nomes `base365_*` nas ferramentas MCP (interface `mcp-tools.md`). |
| `config` (variáveis de ambiente) | `_reversa_sdd/code-analysis.md#Módulo: config` | Gateway, CLI | contrato-alterado | 152 variáveis `BASE365_*`, diretório `~/.base365` (interface `cli-env-e-diretorios.md`). |
| `bootstrap` (templates de contexto) | `_reversa_sdd/code-analysis.md#Módulo: bootstrap` | Gateway | regra-alterada | Textos semeados dizem Base365 (RF-13). |
| `tracing/otelexport` | `_reversa_sdd/code-analysis.md#Módulo: tracing` | Gateway | regra-alterada | Serviço padrão `base365-gateway` (`internal/tracing/otelexport/exporter.go:26`, `:45`; `internal/config/config.go:543`). |
| `updater` | `_reversa_sdd/code-analysis.md#Módulo: updater` | Desktop | regra-alterada | Repositório `edyoCampos/base365` (T-12). |
| `i18n` (backend) | `_reversa_sdd/code-analysis.md#Módulo: i18n` | Gateway | regra-alterada | Valores nos 5 catálogos Go. |
| App desktop (Wails) | `_reversa_sdd/desktop-core/` | Desktop | regra-alterada | Metadados, keyring, título, ícones (T-11, T-13). |
| Web UI e desktop frontend | `_reversa_sdd/design-system/` | Web, desktop | regra-alterada | Tokens, classes, logo, favicon, i18n, nome dos pacotes (T-13, T-16). |
| Scripts, compose e Docker | `_reversa_sdd/deployment.md` | CLI/runtime | regra-alterada | `compose.d/00-goclaw.yml` e `scripts/zuey/goclaw-*.sh` renomeados, `install*.sh/.ps1`, nomes de imagem, Makefile (`BINARY = goclaw`). |
| Documentação | `_reversa_sdd/inventory.md` | Docs | regra-alterada | `README.md`, 30 READMEs em `_readmes/`, `docs/`, `AGENTS.md`, `CLAUDE.md`, `CONTRIBUTING.md`. |
| Checagem de marca | — | CLI/CI | componente-novo | `scripts/check-brand.sh` e `scripts/brand-exceptions.txt` (T-15). |

**Caminhos com "goclaw" no nome (11 itens, medição em 2026-10-02, fora de `_reversa_*`, `.reversa/`, `.claude/`):** `compose.d/00-goclaw.yml`, `_statics/goclaw-logo.svg`, `_statics/goclaw-icon.svg`, `_statics/goclaw.png`, `plan/goclaw-mcp-integration.md`, `skills/goclaw` (pasta), `scripts/zuey/goclaw-deploy.sh`, `scripts/zuey/goclaw-upgrade-release.sh`, `ui/web/public/goclaw-icon.svg`, `ui/desktop/frontend/public/goclaw-icon.svg` e `ui/desktop/frontend/dist/goclaw-icon.svg` (saída de build). O requirements citava 12, e a remedição de 2026-10-02 achou 11. O número oficial é o que `scripts/check-brand.sh` medir.

## 6. Delta no modelo de dados

- Resumo das mudanças: nenhuma tabela, coluna ou índice muda. Mudam comentários SQL, um valor semeado (`_goclaw_recovery`), um valor padrão de configuração semeada do `web_fetch` e o nome do arquivo de banco SQLite padrão (`goclaw.db` → `base365.db`, `internal/config/config.go:196`).
- Detalhe completo em: `_reversa_forward/001-rebrand-base365/data-delta.md`

## 7. Delta de contratos externos

| Contrato | Tipo | Arquivo de detalhe |
|----------|------|--------------------|
| Cabeçalhos HTTP, assinatura de webhook, prefixo `model` | HTTP | `_reversa_forward/001-rebrand-base365/interfaces/http-headers-e-contratos.md` |
| Ferramentas MCP `goclaw_*` | MCP | `_reversa_forward/001-rebrand-base365/interfaces/mcp-tools.md` |
| Variáveis de ambiente, diretórios, nome do binário | arquivo / CLI | `_reversa_forward/001-rebrand-base365/interfaces/cli-env-e-diretorios.md` |

## 8. Plano de migração

Não há migração de dados (RN-02). Há uma **sequência de execução** com portões entre as fases. Cada fase só começa com a anterior concluída e verde nos testes.

| Fase | Conteúdo | Portão para a próxima |
|------|----------|-----------------------|
| F0 | Baseline versionado (T-01). Registrar contagens iniciais. | Baseline restaurável. |
| F1 | Escrever `scripts/check-brand.sh` e a lista de exceções (T-15). Caracterizar com testes os 4 controles de segurança **antes** de mexer neles. | Checagem roda e falha com a contagem inicial. Testes de caracterização verdes. |
| F2 | Motor de substituição (T-02 a T-05, T-08) com rodada em seco. Revisar uma amostra do relatório de substituições por categoria. | Amostra revisada. |
| F3 | Aplicar substituição e renomear caminhos. `gofmt -w`, `go fix`, `go build` PG e `-tags sqliteonly`. | Os dois builds passam. |
| F4 | Intervenções manuais: T-06 (testes novos), T-07, T-09, T-10, T-11, T-12. | Testes novos e `go test` do que muda passam. |
| F5 | Ativos (T-13) e tokens de cor (T-16). Regenerar `ui/desktop/frontend/dist/`. Revisão do âmbar. | Builds de web e desktop passam. |
| F6 | Remoção das menções a projetos de origem (T-17) depois da verificação de licença. Documentação, pendências (T-14). | `make check-brand` verde. |
| F7 | Portões finais da seção 10. | Todos verdes. |

## 9. Riscos e mitigações

| ID | Risco | Impacto | Probabilidade | Mitigação |
|----|-------|---------|---------------|-----------|
| R-01 | Substituição genérica quebra um ponto de segurança ou contrato sem aviso (ex.: `GOCLAW_` virando `BASE365_` na lista de remoção de ambiente). | alto | médio | Testes de caracterização antes da troca e testes novos depois (T-06). Revisão de amostra por categoria. |
| R-02 | **Licença.** O `LICENSE` atual é CC BY-NC 4.0 (`LICENSE`, "Copyright (c) 2025-2026 base365 Contributors"). A licença exige atribuição e proíbe uso comercial. Remover créditos aos projetos de origem (RF-24) pode violar a exigência de atribuição, e a cláusula NonCommercial pode conflitar com o lançamento comercial. | alto | médio | Antes da F6, o usuário confirma que a aquisição cobre o uso comercial e a remoção de créditos. A remoção de RF-24 só roda depois dessa confirmação, registrada no plano. Não é questão técnica, é jurídica. |
| R-03 | Testes com dados fixos de texto quebram: assinatura HMAC com cabeçalho no fixture, contagem de tokens de templates de bootstrap, snapshots (ex.: `internal/gateway/public_url_snapshot_test.go`, `internal/agent/systemprompt_target_test.go`). A remedição de 2026-10-02 encontrou 638 arquivos `_test.go` citando o nome (o requirements dizia 647, que incluía outros tipos de teste). | médio | alto | Rodar a suíte inteira por fase. Corrigir o fixture, nunca a regra. Pode haver teste que exige limite de caracteres no prompt: o texto ficou 1 caractere maior por ocorrência. |
| R-04 | Reordenação de imports e desalinhamento de comentários após a troca. | baixo | alto | `gofmt -w` e `go vet` no portão F3. |
| R-05 | Links de documentação apontam para páginas inexistentes do GitHub Pages (T-08). | baixo | alto | Aceito. Cada link está na lista de pendências. |
| R-06 | O diretório de saída `ui/desktop/frontend/dist/` contém marca antiga e é embutido no binário. | médio | alto | Regenerar com `pnpm build` e conferir com a checagem de marca antes do build do desktop. |
| R-07 | Ícones `.ico`/`.icns` gerados por Pillow ficam fora do padrão esperado pelo Wails ou pela App Store. | baixo | médio | Validar com `wails build` num ambiente que tenha a ferramenta. Alternativa: deixar o Wails gerar a partir de `appicon.png`. |
| R-08 | Referências à infraestrutura herdada que a regra "goclaw" não pega: `digitop` (49 arquivos), `tamgiac` (20) e `zuey` (12), como os scripts `scripts/zuey/*`, testes de portal Bitrix e `docs/deployment-guide.md`. | médio | alto | Tratar na F6, caso a caso. Os de teste ganham domínio de exemplo e os operacionais entram na lista de pendências. A checagem de marca ganha uma lista de aviso (não falha) para esses termos. |
| R-09 | Ambiente sem git impede o teste "commit que reintroduz GoClaw" (RF-23). | baixo | alto | T-01 resolve. O teste do RF-23 pode rodar também sem commit, sobre um arquivo temporário. |

## 10. Critério de pronto

- [ ] Todas as ações do `actions.md` marcadas `[X]`
- [ ] `make check-brand`: 0 ocorrências de "goclaw" fora da lista de exceções, conteúdo e nomes (RF-01, RF-02, RF-22, RF-23)
- [ ] P1 `go build ./...` e `go build -tags sqliteonly ./...` passam (RF-03, RF-04)
- [ ] P2 `go vet ./...` passa
- [ ] P3 `go test -race ./...` passa (inclui os testes novos da T-06)
- [ ] P4 `pnpm build` de `ui/web` e de `ui/desktop/frontend` passam (RF-16, RF-19, RF-20)
- [ ] Busca por `orange-` nos componentes retorna 0, e a tabela do âmbar está preenchida (RF-20)
- [ ] Busca por `openclaw`, `zeroclaw` e `picoclaw` retorna 0, com a verificação de licença registrada (RF-24, R-02)
- [ ] `docs/rebrand-pendencias.md` existe com um item por endereço provisório (RF-17)
- [ ] `cross-check.md` (se executado) sem CRITICAL nem HIGH
- [ ] `regression-watch.md` gerado
- [ ] Re-extração reversa executada e sem regressão vermelha (recomendado, não obrigatório)

## 11. Histórico de alterações

| Data | Alteração | Autor |
|------|-----------|-------|
| 2026-10-02 | Versão inicial gerada por `/reversa-plan` | reversa |
