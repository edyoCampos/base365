# Pendências externas do rebrand Base365

Lista de recursos fora do código que o rebrand ainda aponta de forma **provisória**. Cada item diz o que o código usa hoje, o que providenciar e o impacto se faltar. Mantenha esta lista: quando um item for resolvido, troque o endereço no código e remova a linha.

A checagem de marca (`make check-brand`) imprime um aviso com a contagem de termos herdados da infraestrutura original (`digitop`, `tamgiac`, `zuey`, `nextlevelbuilder`).

| # | Recurso | O que o código usa hoje | O que providenciar | Impacto se faltar |
|---|---------|-------------------------|--------------------|-------------------|
| 1 | Documentação e site | `https://edyocampos.github.io/base365/` (GitHub Pages do repositório). Os links com âncoras de páginas, como `#quick-start`, ainda não existem. | Publicar o GitHub Pages com as páginas, ou um domínio próprio. | Links de documentação quebrados na interface, nos READMEs e nas mensagens. |
| 2 | Domínio e e-mail de contato | `base365.example.com` (exemplos), `contact@base365.example.com` (`ui/desktop/wails.json`). | Domínio próprio e caixa de e-mail. | E-mail de autor errado nos metadados do app desktop. |
| 4 | Registro de imagens | `ghcr.io/edyocampos/base365` (e `-web`). | Habilitar o GHCR no repositório e publicar as imagens. | `docker pull` e os compose falham até a primeira publicação. |
| 5 | Docker Hub | `edyocampos/base365`. | Criar o repositório no Docker Hub ou remover as referências. | Instruções de instalação por Docker Hub não funcionam. |
| 6 | Releases do app desktop | `edyoCampos/base365` (tags `lite-v*`), artefatos `base365-lite.app` e `base365-lite.exe`. | Fluxo de release (`release-desktop.yaml`) publicando esses nomes. O `.github/` **não existe** neste checkout. | O atualizador do desktop não encontra versões. |
| 7 | CI e releases do gateway | Documentado em `CLAUDE.md` e `docs/deployment-guide.md`, mas sem workflows no repositório. | Criar os workflows (`ci`, `dev-beta-release`, `release`, `release-beta`) e adaptar. Incluir `make check-brand` como etapa. | Sem CI, sem releases automáticas e sem a checagem de marca no PR. |
| 8 | Infraestrutura de deploy herdada ("zuey") | Pasta `scripts/zuey/`, variáveis `ZUEY_*` e jobs `*_zuey_*` nos testes de CI (`scripts/ci/`). Host, porta e IP foram trocados por `<VPS_HOST>` e `<SSH_PORT>` em `docs/deployment-guide.md`. | Definir o servidor de beta próprio e renomear `scripts/zuey/`, `ZUEY_*` e os jobs. | O deploy beta automático aponta para infraestrutura que não é sua. |
| 9 | CLI separada | `edyoCampos/base365-cli` (antes um repositório da organização original). | Criar o repositório ou remover as referências. Verificar o canal de release antes de declarar paridade da CLI (regra do `CLAUDE.md`). | Referências a um pacote CLI inexistente em `docs/10-tracing-observability.md` e `docs/project-changelog.md`. |
| 10 | Identificador de parceiro AIMLAPI | `X-AIMLAPI-Partner-ID: base365` (`internal/providers/aimlapi.go`). Antes era o identificador da organização original. | Registrar o seu identificador de parceiro no AIMLAPI, ou remover o cabeçalho. | Atribuição de uso do provedor AIMLAPI incorreta. |
| 11 | Scripts de instalação | `scripts/install.sh`, `scripts/install-lite.sh`, `scripts/install-lite.ps1` apontam para `edyoCampos/base365` e para o GitHub Pages. | Publicar as releases e, se quiser, uma URL curta. | Instalação por `curl \| bash` falha até haver releases. |
| 12 | Bots e contas nos provedores de canal | Nomes de bots em Telegram, Zalo, Feishu e Discord ficam fora do código. | Renomear ou criar os bots com a marca nova. | O nome visível nos canais continua o antigo. |
| 13 | Licença | `LICENSE` é CC BY-NC 4.0 (atribuição, uso não comercial). A confirmação de que a aquisição cobre uso comercial e a remoção de créditos está em `_reversa_forward/001-rebrand-base365/reports/licenca-confirmacao.md`. Os READMEs traduzidos ainda mostram o selo "MIT" e o texto de licença MIT. | Confirmar o texto de licença correto e alinhar os READMEs traduzidos ao `LICENSE`. | Informação de licença contraditória. |

## Resolvidos

- **3. Serviço de extração de páginas** (2026-10-06): deixou de ser necessário. A ferramenta `web_fetch` passou a extrair o conteúdo principal dentro do próprio gateway (título, autor, data e texto sem menus nem banners), com detecção de charset e sem enviar URLs a terceiros. O extrator externo `defuddle` continua opcional e desativado. Os números dos demais itens foram mantidos.
