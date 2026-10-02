# Baseline de build e testes (antes da troca) — 2026-10-02

Executada depois dos testes de caracterização da Fase 2 (já incluídos).

| Verificação | Resultado |
|-------------|-----------|
| `go build ./...` (PostgreSQL) | OK |
| `go build -tags sqliteonly ./...` | OK |
| `go vet ./...` | OK, sem apontamentos |
| `go test ./...` | OK, 88 pacotes com testes, 0 falhas |
| `ui/web`: `pnpm test` (vitest) | OK, 56 arquivos, 357 testes |
| `ui/web`: `pnpm build` | OK (aviso de chunk > 500 kB, pré-existente) |
| `ui/desktop/frontend`: `pnpm build` | **FALHA pré-existente**: 4 erros TS2307, todos por falta de `wailsjs/` (bindings gerados pelo Wails, que não está instalado neste ambiente). Nenhum erro além desses. |

**Consequência para os portões:** o build do desktop frontend só passa onde o Wails gera `wailsjs/`. Neste ambiente, o critério vira "mesmos 4 erros de `wailsjs`, nenhum erro novo".
