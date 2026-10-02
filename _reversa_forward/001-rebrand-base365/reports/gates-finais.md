# Portões finais (T074 a T076)

| Portão | Resultado |
|--------|-----------|
| `make check-brand` | OK: 0 ocorrências de conteúdo e 0 caminhos; 93 linhas de aviso de marcas herdadas (não falha). |
| `scripts/check-brand.test.sh` | OK |
| `go build ./...` e `go build -tags sqliteonly ./...` | OK |
| `go vet ./...` e `go vet -tags sqliteonly ./...` | OK |
| `go test ./internal/i18n/` (paridade de chaves) | OK |
| `go test -tags sqliteonly ./...` | OK, 92 pacotes, 0 falhas |
| `go test -race ./...` | 88 de 89 pacotes OK. **1 falha:** `internal/providers` `TestChatStream_MultipleToolCalls_OneTruncated`. |
| `ui/web`: `pnpm build` e `pnpm test` | OK (358 testes). `dist/` sem o nome antigo. |
| `ui/desktop/frontend`: `pnpm build` | Mesmos 4 erros TS2307 de `wailsjs/` da baseline, 0 erro novo. |
| Contraste (`scripts/rebrand/contrast-check.mjs`) | 88 pares verificados, 0 falhas. |

## Falha em `-race`: pré-existente

`TestChatStream_MultipleToolCalls_OneTruncated` é instável com `-race` (2 falhas em 15 execuções) e passa sem `-race`. Reproduzida **também na baseline** (commit `90bd368`, em worktree separado): mesmas 2 falhas em 15. Não é efeito do rebrand. Corrigir é trabalho separado.

## Não executado neste ambiente

- **Integração com PostgreSQL** (`go test -tags integration ./tests/integration/`, `make test-invariants`): o daemon do Docker não está rodando, então não há PostgreSQL pgvector na porta 5433. O teste `tests/integration/rebrand_seed_test.go` compila (`go vet -tags integration`) mas **não foi executado**.
- **Build do desktop com Wails:** o Wails não está instalado, então `wailsjs/` não é gerado.
