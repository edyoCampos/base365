# Testes após a troca (T025 a T029)

| Verificação | Resultado |
|-------------|-----------|
| `gofmt -w` | 419 arquivos Go reformatados (imports reordenados e comentários realinhados). |
| `go build ./...` | OK |
| `go build -tags sqliteonly ./...` | OK |
| `go vet ./...` | OK |
| `go test ./...` | OK, 88 pacotes, **0 falhas** (igual à baseline) |

**Fixtures quebrados:** nenhum. As ações T026, T027 e T028 (correção de fixtures) não tiveram o que corrigir. A tabela T029 repete o resultado acima.

**`go fix`:** `go fix -diff ./...` propõe 63 arquivos de modernização sem relação com o rebrand (`for range`, `maps`, entre outros). Não foram aplicados, para não misturar mudança de comportamento no diff do rebrand. Fica como sugestão separada.
