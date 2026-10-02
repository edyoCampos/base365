# Interface: ferramentas MCP expostas

> Identificador: `001-rebrand-base365`
> Data: `2026-10-02`
> Base: `_reversa_sdd/mcp/`
> Tipo: MCP (stdio e HTTP) | Confidência: 🟡 (nomes verificados, lista completa a conferir na F2)

## 1. Mudança

Os nomes das ferramentas expostas pelo servidor MCP começam com `goclaw_` e passam a começar com `base365_`. O nome do servidor na chave `name` também muda (`internal/gateway/router.go:321`, `:383`, `internal/mcp/manager_connect.go:39`, valor `"goclaw"` → `"base365"`).

Exemplos medidos em `internal/` e `cmd/`: `goclaw_skills_update`, `goclaw_session`, `goclaw_llm_complete`, `goclaw_agent_get`, `goclaw_voices_list`, `goclaw_teams_get`, `goclaw_teams_create`, `goclaw_sessions_list`, `goclaw_pairing_device_request`, `goclaw_chat_send`, e outros. A lista completa é gerada na F2 com `grep -rhoE '"goclaw_[a-z_]+"'` antes da troca e arquivada no relatório da fase.

## 2. Impacto em clientes

Qualquer cliente MCP configurado com os nomes antigos (por exemplo, instruções de agente que citam `goclaw_chat_send`) deixa de encontrar a ferramenta. Sem compatibilidade retroativa (RN-02). Os templates de bootstrap e as skills (`skills/goclaw/SKILL.md`, a pasta vira `skills/base365/`) citam esses nomes e mudam junto.

## 3. Verificação

Teste que lista as ferramentas do servidor e confirma que nenhuma começa com `goclaw_`, e que a lista tem o mesmo tamanho de antes.
