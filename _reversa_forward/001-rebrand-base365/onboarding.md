# Onboarding: testar o rebrand pela primeira vez

> Identificador: `001-rebrand-base365`
> Data: `2026-10-02`
> Público: quem vai validar a feature depois do `/reversa-coding`.
> Pré-requisitos: Go 1.26, pnpm, Docker (para o teste de integração em PostgreSQL), Python 3 com Pillow (só para regenerar ícones).

## 1. Antes de qualquer coisa

1. Confirme que existe o baseline da fase F0 (commit inicial ou cópia `tar`). Sem ele, não aplique nada.
2. Leia `docs/rebrand-pendencias.md` (gerado pela feature). Ele lista o que você ainda precisa providenciar fora do código.

## 2. Checagem de marca

```bash
make check-brand
```

Esperado: termina com sucesso e informa 0 ocorrências fora da lista de exceções. Para ver a checagem falhar, adicione "GoClaw" em qualquer arquivo temporário e rode de novo: ela deve apontar arquivo e linha. Remova o arquivo.

## 3. Builds e testes

```bash
go build ./...
go build -tags sqliteonly ./...
go vet ./...
go test -race ./...
cd ui/web && pnpm install && pnpm build
cd ../desktop/frontend && pnpm install && pnpm build
```

Esperado: todos terminam sem erro. O binário gerado se chama `base365`.

## 4. Gateway com variáveis novas

```bash
./base365 onboard
source .env.local        # variáveis BASE365_*
./base365
```

Conferir:
1. `GOCLAW_CONFIG=/algum/arquivo ./base365` não carrega o arquivo (variável antiga ignorada).
2. `BASE365_CONFIG=/algum/arquivo ./base365` carrega.
3. No painel, crie uma chave de API: ela começa com `base365_`.
4. Envie um webhook com HMAC no cabeçalho `X-Base365-Signature`: aceito. Só com `X-GoClaw-Signature`: rejeitado.
5. Com `BASE365_GATEWAY_TOKEN` definida, um agente executando `env` não vê essa variável.
6. Um agente executando `echo $BASE365_ENCRYPTION_KEY` é bloqueado.
7. Ao criar uma credencial de CLI, incluir `BASE365_X` no ambiente é rejeitado (interface e servidor).

## 5. Agente e extrator

1. Crie um agente novo e pergunte "em que plataforma você roda?". A resposta não menciona GoClaw.
2. Em Ferramentas → web_fetch, o extrator `defuddle` aparece desativado, sem o endereço `fetch.goclaw.sh`.
3. Peça ao agente para ler uma página pública. O conteúdo vem pela extração interna.

## 6. Interface

1. Abra o painel web nos temas claro e escuro. Primária `#7c4fe0` no claro e `#a487f2` no escuro, fundo Névoa no claro e `#15111f` no escuro.
2. Logo, favicon e título da aba mostram Base365.
3. Nenhum laranja de marca aparece. Avisos de limite, pendência e expiração continuam âmbar.
4. Texto sobre a primária no escuro é Berinjela (não branco).

## 7. Desktop

```bash
make desktop-build VERSION=0.1.0
```

Conferir: nome do produto, título da janela e metadados do executável dizem Base365; o diretório criado é `~/.base365/`; nada é criado em `~/.goclaw/`; o ícone é o oficial.

## 8. Banco novo (PostgreSQL e SQLite)

Crie um banco do zero em cada motor e rode a consulta por conteúdo `ILIKE '%goclaw%'` nas tabelas semeadas (`agent_context_files`, `builtin_tools`, entre outras). Esperado: 0 linhas.

## 9. O que NÃO testar nesta feature

Serviço de extração próprio, CI, domínio, e-mail, bots nos provedores. Estão na lista de pendências.
