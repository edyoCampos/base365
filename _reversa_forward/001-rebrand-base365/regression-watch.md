# Regression watch: Rebrand GoClaw → Base365

> Identificador: `001-rebrand-base365`
> Criado em: `2026-10-02`
> Usado por: re-extrações do `/reversa` e por `/reversa-sync`. Cada item diz o que precisa continuar verdadeiro depois do rebrand.

## Watch principal

| ID | Origem (arquivo, seção) | Regra esperada após a mudança | Tipo de verificação | Sinal de violação |
|----|-------------------------|-------------------------------|---------------------|-------------------|
| W001 | `_reversa_sdd/code-analysis.md#Módulo: crypto` | A chave de API gerada segue `base365_<32hex>`. | redação | A extração volta a documentar `goclaw_<32hex>`, ou `apiKeyPrefix` deixa de ser `base365_`. |
| W002 | `_reversa_sdd/code-analysis.md#Módulo: crypto` (`env_denylist.go`) | Credenciais de CLI com variável `BASE365_*` são rejeitadas no servidor e na interface web. | presença | `IsDeniedEnvKey("BASE365_X")` retorna falso, ou o espelho web perde o prefixo. |
| W003 | `internal/tools/env_scrub.go`, `internal/workstation/security/allowlist.go` | Variáveis `BASE365_GATEWAY_TOKEN` e `BASE365_*` não chegam a processos filhos nem à workstation. | presença | O teste `TestScrubCredentialEnv_StripsProductGatewayToken` ou o da allowlist falha. |
| W004 | `internal/tools/shell_deny_groups.go` | Comandos que imprimem `$BASE365_ENCRYPTION_KEY`, `$BASE365_GATEWAY_TOKEN` ou `$BASE365_POSTGRES_DSN` (echo, printf, python, node) são bloqueados. | presença | `TestProductSecretExfiltrationDeny` falha. |
| W005 | `_reversa_sdd/code-analysis.md#Módulo: webhooks`, `flowcharts/webhooks.md` | A autenticação HMAC do webhook lê `X-Base365-Signature`, tolerância de 300 s, cache de nonce. O cabeçalho antigo é ignorado. | redação | A extração cita `X-GoClaw-Signature`, ou `TestWebhookAuth_LegacySignatureHeaderRejected` falha. |
| W006 | `_reversa_sdd/domain.md#RN-T08` | Um `X-Base365-Tenant-Id` que não resolve cai no master tenant sem erro. | redação | A regra continua escrita com `X-GoClaw-Tenant-Id`. |
| W007 | `internal/http/auth.go` | O campo `model` aceita `base365:<agentId>` e `agent:<agentId>`. O prefixo antigo não resolve o agente. | presença | `TestExtractAgentID_LegacyNamesIgnored` falha. |
| W008 | `_reversa_sdd/code-analysis.md#Módulo: mcp` | As 178 ferramentas CRUD MCP começam com `base365_`. Nenhuma começa com o prefixo antigo. | presença | `TestCRUDToolNames_Prefix` falha, ou a contagem cai abaixo de 178. |
| W009 | `_reversa_sdd/code-analysis.md#Módulo: tools` (`web_fetch`) | Em instalação nova, o extrator `defuddle` vem desativado e sem endereço. Uma entrada ativa sem `base_url` é ignorada. A leitura usa extração interna. | redação | `TestWebFetchSeed_*` ou `TestResolveExtractorChain_DefuddleWithoutBaseURLSkipped` falha, ou a extração documenta `fetch.goclaw.sh` como padrão. |
| W010 | `_reversa_sdd/code-analysis.md#Módulo: config` | Variáveis do produto usam o prefixo `BASE365_`. Variáveis `GOCLAW_*` não são lidas. | ausência | Qualquer leitura de variável com o prefixo antigo reaparece. |
| W011 | `_reversa_sdd/domain.md` (Lane) | Parâmetros de lane são configurados por `BASE365_LANE_*`. | redação | A extração continua citando `GOCLAW_LANE_*`. |
| W012 | `_reversa_sdd/code-analysis.md#Módulo: updater` | O atualizador do desktop consulta `edyoCampos/base365` (tags `lite-v*`). | presença | `TestUpdaterTargetsProductRepository` falha. |
| W013 | `ui/desktop/keyring.go` | O desktop guarda segredos em `~/.base365/secrets` e usa o serviço de keyring `base365-desktop`. O diretório antigo não é criado. | ausência | `TestSecretsDir_UsesProductDirectory` falha. |
| W014 | `internal/tracing/otelexport/exporter.go` | O nome de serviço padrão é `base365-gateway`. | redação | `DefaultServiceName` muda. |
| W015 | `_reversa_sdd/design-system/color-palette.md#1. Identidade da marca` | A cor primária é `#7c4fe0` (claro) e `#a487f2` (escuro). Texto sobre a primária escura é Berinjela. Nenhuma classe `orange-*` nos componentes. | redação | A extração volta a descrever laranja de matiz 38, ou `orange-` reaparece. |
| W016 | `docs/rebrand-amber-review.md` | Âmbar de aviso (atenção, limite, pendente, expira) usa o token `warning`. Âmbar decorativo não existe. | redação | Classes `amber-*` reaparecem nos componentes. |
| W017 | `scripts/check-brand.sh` | `make check-brand` passa. O nome antigo não existe em conteúdo nem em nomes de arquivo, fora de `scripts/brand-exceptions.txt`. | ausência | O alvo falha, ou a lista de exceções cresce sem justificativa. |
| W018 | `_reversa_sdd/data-dictionary.md`, `internal/upgrade/version.go` | O esquema de dados não mudou: `RequiredSchemaVersion` = 97 e `SchemaVersion` do SQLite = 60. Nenhuma tabela contém o nome antigo em valor semeado. | confidência | A versão muda sem migration, ou `TestFreshDB_HasNoLegacyBrand` falha. |
| W019 | `_reversa_sdd/deployment.md#6` | As imagens são publicadas como `ghcr.io/edyocampos/base365` (provisório, ver pendências). | redação | A extração ou os compose voltam a citar `ghcr.io/nextlevelbuilder/goclaw`. |

## Observações (sem peso de regressão, regras 🟡 ou 🔴 na origem)

- O formato de `README` traduzido e a tabela comparativa com projetos de origem foram simplificados. Sem verificação automática.
- O texto de licença nos READMEs traduzidos (selo MIT) ainda diverge do `LICENSE` (CC BY-NC 4.0). Está em `docs/rebrand-pendencias.md`, item 13.
- O identificador de parceiro AIMLAPI (`base365`) é provisório (pendências, item 10).

## Histórico de re-extrações

_(vazio, será preenchido quando o `/reversa` rodar de novo)_

## Arquivadas

_(vazio)_
