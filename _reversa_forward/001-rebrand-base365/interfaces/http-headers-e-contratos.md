# Interface: cabeçalhos HTTP, webhook e campo `model`

> Identificador: `001-rebrand-base365`
> Data: `2026-10-02`
> Base: `_reversa_sdd/http/`, `_reversa_sdd/flowcharts/webhooks.md`, `_reversa_sdd/openapi/http-api.yaml`
> Tipo: HTTP | Confidência: 🟢

## 1. Cabeçalhos próprios do produto

Medição em 2026-10-02 (ocorrências no repositório fora de `_reversa_*`):

| Antes | Depois | Ocorrências | Uso |
|-------|--------|-------------|-----|
| `X-GoClaw-User-Id` | `X-Base365-User-Id` | 270 | Identidade do usuário (`internal/http/auth.go:46`). |
| `X-GoClaw-Tenant-Id` | `X-Base365-Tenant-Id` | 36 | Tenant (`auth.go:189`, `:211`, `:229`). |
| `X-GoClaw-Signature` | `X-Base365-Signature` | 14 | Assinatura HMAC do webhook (`internal/http/webhooks_auth.go:238`, `:284`). |
| `X-GoClaw-Sender-Id` | `X-Base365-Sender-Id` | 11 | Usuário pareado (`auth.go:225`). |
| `X-GoClaw-Agent-Id` | `X-Base365-Agent-Id` | 9 | Agente alvo (`auth.go:69`). |
| `X-GoClaw-Upgrade-Token` | `X-Base365-Upgrade-Token` | 3 | Upgrade protegido (`internal/http/gateway_upgrade.go:24`). |
| `X-GoClaw-Agent` | `X-Base365-Agent` | 3 | Alias legado do agente alvo (`auth.go:72`). |

**Consumidores que precisam mudar junto (RF-07):** interface web e desktop (cliente HTTP e WebSocket), `scripts/zuey/goclaw-upgrade-release.sh`, `scripts/zuey/goclaw-deploy.sh`, documentação (`docs/webhooks.md`, `api-reference.md`), `openapi/` e testes.

## 2. Webhook de entrada (HMAC)

Regras de autenticação **não mudam**: HMAC, tolerância de 300 s e cache de nonce (`internal/http/webhooks_auth.go`).

| Cenário | Resultado esperado |
|---------|--------------------|
| Assinado em `X-Base365-Signature` | aceito |
| Assinado só em `X-GoClaw-Signature` | rejeitado (sem cabeçalho de assinatura reconhecido) |
| Ambos presentes | usa `X-Base365-Signature`; o antigo é ignorado |

Idempotência e timeouts: inalterados.

## 3. Campo `model` em `/v1/chat/completions`

`internal/http/auth.go:61` aceita `model` no formato `goclaw:<agentId>` (ou `agent:<agentId>`). O prefixo passa a `base365:<agentId>`; `agent:<agentId>` continua. Clientes que enviam `goclaw:<id>` deixam de resolver o agente (RN-02).

## 4. Erros

Nenhum código de erro ou corpo de resposta muda, exceto o texto que cita o produto (catálogos de i18n, RF-12).

## 5. Verificação

Teste de contrato em `tests/` com os casos acima. Teste negativo para o cabeçalho antigo.
