# Data delta: Rebrand GoClaw → Base365

> Identificador: `001-rebrand-base365`
> Data: `2026-10-02`
> Base: `_reversa_sdd/data-dictionary.md`, `_reversa_sdd/erd-complete.md`
> Confidência: 🟢 CONFIRMADO, 🟡 INFERIDO, 🔴 LACUNA

## 1. Resumo

Nenhuma tabela, coluna, índice, constraint ou tipo muda. O delta é só de **texto** e de **valores padrão**. Por isso `RequiredSchemaVersion` (97, `internal/upgrade/version.go`) e `SchemaVersion` do SQLite (60, `internal/store/sqlitestore/schema.go:19`) **não são incrementados** (roadmap T-07, 🟡).

## 2. Mudanças

| # | Onde | Antes | Depois | Motor | Tipo |
|---|------|-------|--------|-------|------|
| 1 | `migrations/000001_init_schema.up.sql:1` | comentário `-- GoClaw Multi-Tenant Schema` | `-- Base365 Multi-Tenant Schema` | PG | comentário |
| 2 | `migrations/000020_secure_cli_and_api_keys.up.sql:2` | comentário "GoClaw auto-injects…" | "Base365 auto-injects…" | PG | comentário |
| 3 | `migrations/000061_webhooks_encrypted_secret.up.sql:1` | comentário `GOCLAW_ENCRYPTION_KEY` | `BASE365_ENCRYPTION_KEY` | PG | comentário |
| 4 | `migrations/000095_restore_misclassified_custom_skills.up.sql:13` | valor JSON `{"_goclaw_recovery":"bundled_slug_collision"}` | `{"_base365_recovery":"bundled_slug_collision"}` | PG | **valor semeado** |
| 5 | `internal/store/sqlitestore/schema.sql:1` e `:2094` | comentários | idem | SQLite | comentário |
| 6 | `internal/store/sqlitestore/*.go` (agents_access, evolution_metrics, cron_crud e outros que citam o nome) | comentários e textos | idem | SQLite | texto |
| 7 | Configuração semeada de `builtin_tools.settings` do `web_fetch` (`cmd/gateway_builtin_tools.go:42`) e a função de preenchimento (`:253`, `backfillWebFetchSettings`) | `defuddle` ativo apontando para `https://fetch.goclaw.sh/` | `defuddle` **desativado**, sem endereço dos autores originais | PG e SQLite | **valor padrão** |
| 8 | Nome do arquivo SQLite padrão (comentário em `internal/config/config.go:196`, e o ponto do código que monta o caminho) | `{dataDir}/goclaw.db` | `{dataDir}/base365.db` | SQLite | valor padrão |
| 9 | `internal/store/pg/pool.go:17` | `PoolApplicationName = "goclaw"` | `"base365"` | PG | valor (aparece em `pg_stat_activity`) |
| 10 | `docker-compose.postgres.yml:20,22,26,37` | usuário, banco e senha padrão `goclaw` | `base365` | PG (compose) | valor padrão |
| 11 | Textos semeados dos templates de bootstrap (`internal/bootstrap/templates/IDENTITY.md`, `AGENTS.md`, `AGENTS_CORE.md`, `AGENTS_TASK.md`, `seed_store.go`, `load_store.go`) | "GoClaw" | "Base365" | PG e SQLite | **conteúdo semeado** (RF-13, RF-14) |
| 12 | Exemplos no prompt de extração do grafo de conhecimento (`internal/knowledgegraph/extractor_prompt.go:29,80,88,89`) | entidade `goclaw` | `base365` | PG e SQLite | texto enviado ao modelo |

## 3. Comportamento em bancos já existentes

Sem produção (RN-02), nenhum banco precisa ser migrado. Um banco criado **antes** da mudança continua funcionando, mas mantém:
- o valor `_goclaw_recovery` já gravado (linha 4);
- as configurações já semeadas do `web_fetch` (`backfillWebFetchSettings` só age quando as configurações estão vazias).

Isso é aceitável porque os critérios de aceite do RF-14 e do RF-18 testam **banco criado do zero**.

## 4. Itens que NÃO mudam

- Nomes de tabela e coluna: nenhum contém "goclaw" (medição em `migrations/` e `schema.sql`, só comentários e o valor da linha 4).
- Chaves de criptografia e derivação: `internal/crypto` só cita o nome no prefixo da chave de API (`apikey.go:10`). Nenhum sal, contexto ou dado adicional de autenticação usa o nome, então os dados criptografados não são afetados.
- Chaves de contexto Go (`internal/store/context.go`, ex.: `"goclaw_user_id"`): valor de uso interno em memória, renomeado como parte da regra genérica, sem persistência.

## 5. Verificação

- Banco novo em PG e em SQLite (teste de integração `tests/integration/`): nenhuma linha semeada contém "goclaw" (consulta por conteúdo).
- `web_fetch` em banco novo: `settings.extractors[defuddle].enabled = false`.
