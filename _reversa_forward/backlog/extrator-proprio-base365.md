# Backlog: serviço de extração próprio do Base365

> Origem: `001-rebrand-base365` (RN-07, RF-17, RF-18), conversa de 2026-10-02
> Estado: **pendente, a cargo do dono do produto (Edyo)**. Não bloqueia o rebrand.
> Quando retomar: depois do merge da feature `001-rebrand-base365`, ou antes do primeiro lançamento público, o que vier primeiro.

## Contexto

O `web_fetch` tem uma cadeia de extratores. O primeiro, `defuddle`, chama um serviço externo:
`GET <base_url>/<domínio>/<caminho>` e recebe markdown limpo (Defuddle, licença MIT).
Hoje o endereço é `https://fetch.goclaw.sh/` (infraestrutura dos autores originais, que vê toda URL lida pelos agentes).
Código: `internal/tools/web_fetch_extractor_defuddle.go` (cliente), `cmd/gateway_builtin_tools.go:42` e `:253` (padrão semeado).
O segundo extrator, `html-to-markdown`, roda dentro do gateway e é o fallback.

O GitHub Pages é estático e não executa código, então não hospeda esse serviço.

## Decisão para o rebrand (feature 001)

Opção 3: o padrão semeado deixa o `defuddle` **desativado** e a leitura de páginas usa a extração interna. Nenhuma URL vai para terceiros. O endereço do serviço próprio entra na lista de pendências externas (RF-17).

## Plano para depois (opção 1: Cloudflare Worker próprio)

| Passo | Ação | Pronto quando |
|-------|------|---------------|
| 1 | Criar conta/projeto no Cloudflare e um Worker (ex.: `base365-fetch`), endereço `base365-fetch.<conta>.workers.dev`. | Worker responde 200 em uma URL de teste. |
| 2 | Implementar o Worker: busca a página de `<domínio>/<caminho>`, extrai com Defuddle (`kepano/defuddle`, MIT) e devolve markdown com frontmatter. Mesmo contrato do serviço original. | `curl <worker>/example.com` devolve markdown. |
| 3 | Proteger contra abuso: limite de taxa, bloqueio de IPs privados e loopback (SSRF), limite de tamanho (1 MB, como o cliente), timeout. | Requisição a `localhost` ou IP privado é recusada. |
| 4 | Guardar o código do Worker num repositório próprio (ou pasta `deploy/` do repositório, se preferir). | Código versionado. |
| 5 | No Base365: trocar `base_url` e ligar o extrator `defuddle` nas configurações (painel ou semente). Sem mudança de código. | Com o Worker no ar, `web_fetch` usa o `defuddle`. Com ele fora, cai no `html-to-markdown`. |
| 6 | Atualizar a lista de pendências externas do repositório (RF-17), marcando o item como resolvido. | Item removido ou marcado. |

## Alternativas descartadas por agora

- **Extrator em Go no gateway (ex.: `go-readability`):** zera dependência externa, mas é desenvolvimento de funcionalidade, fora do escopo de um rebrand. Reavaliar se o Worker não valer a manutenção.
- **Container Node + Defuddle no compose:** boa opção para quem hospeda por conta própria. Adiciona um processo a operar.
- **Vercel, Deno Deploy, Fly.io:** equivalentes ao Worker, com mais um fornecedor.

## Lembrete

Quando o dono do produto mencionar release, deploy, domínio ou `web_fetch`, lembrar deste item.
