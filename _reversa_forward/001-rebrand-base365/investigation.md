# Investigação: Rebrand GoClaw → Base365

> Identificador: `001-rebrand-base365`
> Data: `2026-10-02`
> Confidência: 🟢 CONFIRMADO, 🟡 INFERIDO, 🔴 LACUNA

## 1. Medições de base (2026-10-02)

Fora de `_reversa_sdd/`, `_reversa_forward/`, `.reversa/`, `.claude/` e `node_modules/`:

| Medida | Valor |
|--------|-------|
| Arquivos de texto com "goclaw" (qualquer grafia) | 1.997 |
| Linhas com "goclaw" | 7.756 |
| Ocorrências por grafia | `goclaw` 6.432, `GoClaw` 1.111, `GOCLAW` 687, `Goclaw` 30 |
| Variantes com separador (`go-claw`, `go_claw`, `Go Claw`) | 0 |
| Arquivos `.go` com o caminho do módulo | 1.552 (3.707 linhas) |
| Variáveis `GOCLAW_*` distintas | 152 |
| Arquivos ou pastas com "goclaw" no nome | 11 |
| Arquivos `_test.go` citando o nome | 638 |
| Arquivos binários citando o nome | 0 |
| `nextlevelbuilder` (todas as formas) | 3.815 linhas, ou seja, cerca de 108 além do caminho do módulo |
| `docs.goclaw.sh` | 33 arquivos, 437 linhas |
| `goclaw.sh` | 52 arquivos, 469 linhas |
| `digitop` / `tamgiac` / `zuey` | 49 / 20 / 12 arquivos |
| `openclaw` / `zeroclaw` / `picoclaw` | 44 / 32 / 32 arquivos |
| Arquivos com `orange-` / com `amber-` (web e desktop) | 49 / 124 |

**Diferença em relação ao `requirements.md`:** ele cita 8.272 ocorrências em 2.001 arquivos, medidas sobre uma área maior (incluía diretórios de ferramentas). Os números acima têm fronteira explícita. O que vale para a conclusão é a checagem de marca (T-15), não os números.

**Não é repositório git** (`git status` falha em `/home/edyo/Projects/base365`).

## 2. Pontos onde a troca não é mecânica

| Ponto | Por que não é mecânico | Tratamento |
|-------|------------------------|------------|
| Prefixo de segurança `GOCLAW_` | Aparece em lista de bloqueio, remoção de ambiente, regex de comando e espelho na web. Perder um ponto abre brecha de segredo. | T-06, testes de caracterização antes e depois. |
| Domínios (`docs.goclaw.sh`, `goclaw.sh`) | A troca genérica criaria `docs.base365.sh`, que não existe. | T-03, T-08: regras de URL antes da genérica. |
| Org `nextlevelbuilder` fora do módulo | A regra "goclaw" não pega (ex.: `github.com/nextlevelbuilder/goclaw-docs`). | T-08 e lista de pendências. |
| `digitop`, `tamgiac`, `zuey` | Infraestrutura dos autores originais, sem "goclaw" no nome. | R-08, F6. |
| Extrator `fetch.goclaw.sh` | É serviço de terceiros e vaza URLs lidas pelos agentes. | T-09. |
| Ordem dos imports e alinhamento Go | 7 letras contra 6. | `gofmt -w`. |
| `ui/desktop/frontend/dist/` | Saída de build embutida no binário, contém `goclaw-icon.svg`. | Regenerar (R-06). |
| Ícones binários (`.ico`, `.icns`, `.png`) | Não são texto, a regra não os altera. | T-13. |
| Licença do repositório | `LICENSE` é CC BY-NC 4.0, com exigência de atribuição e proibição de uso comercial. | R-02. |

## 3. Alternativas avaliadas

| Tema | Alternativa | Resultado |
|------|-------------|-----------|
| Estratégia | Rename incremental por módulo, com compatibilidade temporária | Descartada: RN-02 dispensa compatibilidade, e o estado intermediário misturaria marcas. |
| Estratégia | Big-bang scriptado com portões | **Escolhida.** |
| Ferramenta | `sed`/`perl` por regra em arquivo de dados | **Escolhida.** Auditável e reexecutável. |
| Ferramenta | `gopls rename` por símbolo | Descartada: o problema é texto (strings, comentários, docs), não só símbolos. |
| Migrations | Migration nova de renomeação | Descartada: sem produção, só ruído (T-07). |
| Ícones | ImageMagick | Indisponível no ambiente. Pillow 12.1.1 disponível. |
| Extrator | Apontar para um Worker próprio agora | Descartada: o serviço ainda não existe. Ver `_reversa_forward/backlog/extrator-proprio-base365.md`. |

## 4. Fontes

- Manual de cores: `/home/edyo/Downloads/Base365 logo e ícone.pdf` (V1): tokens web e desktop, badges, contraste (87 pares), regra do âmbar. 🟢
- Ativos: `/home/edyo/Downloads/Base365 logo e ícone/export/` (SVG: `simbolo*.svg`, `icone-app*.svg`, `favicon-16/32.svg`; PNG: `favicon-16/32/64`, `icone-app-180/512/1024`, `icone-app-escuro-1024`, `simbolo-1024`, `simbolo-branco-1024`, logos horizontais e empilhado). 🟢
- Legado: `_reversa_sdd/brainstorms/001-rebrand-goclaw-base365/` (framing, decision, pre-spec). 🟢
- Padrões do projeto: `CLAUDE.md` (migrations dual-DB, i18n, checklist pós-implementação). 🟢

## 5. Questões que ficam para a execução

- Lista completa dos nomes `goclaw_*` das ferramentas MCP (arquivar antes da troca).
- Mapeamento do e-mail de contato do desktop (pendência externa).
- Tratamento caso a caso de `digitop`, `tamgiac` e `zuey` (R-08).
