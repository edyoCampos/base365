# Requirements: Rebrand GoClaw → Base365

> Identificador: `001-rebrand-base365`
> Data: `2026-10-01`
> Pasta da extração reversa: `_reversa_sdd/`
> Insumo: `_reversa_sdd/brainstorms/001-rebrand-goclaw-base365/pre-spec.md` (sessão de ideação completa na mesma pasta)
> Confidência: 🟢 CONFIRMADO, 🟡 INFERIDO, 🔴 LACUNA / DÚVIDA

## 1. Resumo executivo

O produto é um fork adquirido que será lançado como produto próprio com o nome **Base365**. Todas as superfícies ainda exibem a identidade GoClaw:
- as mensagens dos agentes nos canais;
- o painel web e o app desktop;
- o comando de instalação, as variáveis de ambiente e as imagens de contêiner;
- o próprio código.

Esta feature entrega uma troca **total** de identidade. A palavra "goclaw", em qualquer grafia, deixa de existir no repositório. A paleta laranja-terracota é substituída pela paleta violeta do manual da marca, e logo, ícones e favicons passam a ser os ativos oficiais. Ninguém usa o produto em produção, então nenhuma compatibilidade retroativa é exigida.

## 2. Contexto a partir do legado

| Fonte | Trecho relevante | Confidência |
|-------|------------------|-------------|
| `_reversa_sdd/architecture.md#1. Visão geral` | Gateway multi-tenant com interface web, app desktop (edição Lite) e canais de mensageria. Todas as superfícies carregam o nome. | 🟢 |
| `_reversa_sdd/inventory.md#5. CI/CD, Docker e deploy` | As imagens são publicadas em `ghcr.io/nextlevelbuilder/goclaw` e `digitop/goclaw`. A pasta `.github/` não existe no checkout. Há scripts de instalação (`install-lite.sh`, `install-lite.ps1`) e `compose.d/00-goclaw.yml`. | 🟡 |
| `_reversa_sdd/deployment.md#6. Publicação de imagens` | Quatro variantes de imagem (latest, base, full, web) mais beta, todas com o nome do produto. | 🟢 |
| `_reversa_sdd/design-system/color-palette.md#1. Identidade da marca` | A paleta inteira gira em torno de um laranja de matiz 38 (`#e04f1a` no claro, `#f1683b` no escuro) sobre neutros quentes. Web: 37 tokens por tema. Desktop: 15 tokens copiados à mão da web. | 🟢 |
| `_reversa_sdd/design-system/design-system.md#2. Tokens resumidos` | O desktop declara fontes que não carrega e usa tokens próprios. A web usa Inter. | 🟢 |
| `_reversa_sdd/code-analysis.md#Módulo: crypto` (Arquivos principais) | As chaves de API seguem o formato `goclaw_<32hex>`. `env_denylist.go` rejeita o prefixo `GOCLAW_` nas variáveis de ambiente de credenciais de CLI. | 🟢 |
| `_reversa_sdd/code-analysis.md#Módulo: webhooks` (Regras) | Os webhooks de entrada são autenticados por HMAC no cabeçalho `X-GoClaw-Signature` (HMAC, código de autenticação de mensagem com hash). | 🟢 |
| `_reversa_sdd/flowcharts/webhooks.md` | O fluxo de autenticação do webhook decide pelo cabeçalho `X-GoClaw-Signature`. | 🟢 |
| `_reversa_sdd/code-analysis.md#Módulo: updater` | O atualizador do desktop consulta releases de outro repositório. | 🟢 |
| `_reversa_sdd/code-analysis.md#Módulo: tools` | A ferramenta `web_fetch` tem uma cadeia de extratores. O primeiro é um serviço externo hospedado pelos autores originais. | 🟡 |
| `_reversa_sdd/code-analysis.md#Módulo: bootstrap` | Os templates de contexto dos agentes (enviados ao modelo de linguagem) citam o produto pelo nome. | 🟡 |
| `_reversa_sdd/code-analysis.md#Módulo: i18n` | Há catálogos de mensagens no backend. O `CLAUDE.md` cita en, vi e zh, e o código também tem ko e ru. | 🟡 |
| `_reversa_sdd/domain.md#2. Glossário` (Lane) | Parâmetros de negócio são configurados por variáveis `GOCLAW_*`, por exemplo `GOCLAW_LANE_*`. | 🟢 |
| `_reversa_sdd/brainstorms/001-rebrand-goclaw-base365/decision.md#Pendências externas a providenciar` | Lista de 12 recursos externos (domínio, documentação, serviço de extração, registro de imagens, releases, integração contínua, ou CI) que o dono do produto precisa providenciar. | 🟡 |

**Medição direta do repositório em 2026-10-01** (fora do `_reversa_sdd/`, confidência 🟢):
- 8.272 ocorrências de "goclaw", sem diferenciar maiúsculas, em 2.001 arquivos;
- 1.552 arquivos Go importam o módulo;
- 152 variáveis `GOCLAW_*` distintas;
- 97 arquivos citam o diretório `.goclaw`;
- 12 arquivos ou pastas têm "goclaw" no nome;
- 647 arquivos de teste citam o nome;
- 162 classes `orange-*` em 38 arquivos da web e 36 em 8 do desktop;
- 414 classes `amber-*` na web e 60 no desktop.

## 3. Personas e cenários de uso

| Persona | Objetivo | Cenário-chave |
|---------|----------|---------------|
| Usuário final | Conversar com um agente pelo Telegram, WhatsApp ou chat web | Pergunta ao agente "quem é você?" e a resposta não menciona GoClaw. |
| Cliente ou administrador (admin) | Configurar agentes, canais e ferramentas no painel web ou no app desktop | Abre o painel e vê nome, logo, favicon e cores Base365 em todas as telas, nos temas claro e escuro. |
| Operador ou técnico | Instalar e operar o produto | Instala pela imagem de contêiner ou pelo script, configura variáveis `BASE365_*` e encontra os dados em `~/.base365/`. |
| Desenvolvedor | Manter e evoluir o código | Clona `github.com/edyoCampos/base365`, compila e roda os testes sem encontrar "goclaw". Se reintroduzir o nome, a checagem automática falha. |

## 4. Regras de negócio novas ou alteradas

1. **RN-01:** O produto tem uma identidade única: Base365. Nenhum artefato do repositório (conteúdo ou nome de arquivo) contém "goclaw" em qualquer grafia, exceto o que constar numa lista de exceções versionada e justificada item a item. 🟢
   - Origem: `_reversa_sdd/brainstorms/001-rebrand-goclaw-base365/framing.md#Restrições e premissas declaradas`
   - Tipo: nova
2. **RN-02:** Nenhuma compatibilidade retroativa com a identidade antiga é exigida. Variáveis, diretórios, cabeçalhos e prefixos antigos deixam de ser aceitos e não ganham aliases. 🟢
   - Origem: `_reversa_sdd/brainstorms/001-rebrand-goclaw-base365/framing.md#Fora de escopo declarado`
   - Tipo: nova
3. **RN-03:** Os controles de segurança baseados no prefixo de variáveis de ambiente do produto passam a valer para `BASE365_`, com cobertura igual à de hoje: 🟢
   - a lista de bloqueio das variáveis de credenciais de CLI, no backend e no espelho da interface web;
   - a lista de bloqueio de variáveis da workstation remota;
   - a remoção de credenciais do ambiente de processos filhos;
   - os padrões de comando bloqueados que tentam ler segredos do produto.
   - Origem no legado: `_reversa_sdd/code-analysis.md#Módulo: crypto` (`env_denylist.go`) e os arquivos `internal/workstation/security/allowlist.go:189-190`, `internal/tools/env_scrub.go:23` e `internal/tools/shell_deny_groups.go:233-236`
   - Tipo: alterada
4. **RN-04:** As chaves de API geradas pelo gateway seguem o formato `base365_<32 caracteres hexadecimais>`. O resto continua igual (16 bytes aleatórios, hash SHA-256, prefixo de exibição). 🟢
   - Origem no legado: `_reversa_sdd/code-analysis.md#Módulo: crypto` (`goclaw_<32hex>`)
   - Tipo: alterada
5. **RN-05:** Os cabeçalhos HTTP próprios do produto usam o prefixo `X-Base365-`, por exemplo `X-Base365-Signature`, `X-Base365-User-Id`, `X-Base365-Tenant-Id` e `X-Base365-Upgrade-Token`. As regras de autenticação continuam iguais: HMAC, tolerância de 300 s e cache de nonce. 🟢
   - Origem no legado: `_reversa_sdd/code-analysis.md#Módulo: webhooks` (Regras), `_reversa_sdd/flowcharts/webhooks.md`
   - Tipo: alterada
6. **RN-06:** As cores de marca da interface vêm exclusivamente do manual da marca (`base365-marca.pptx`, slide 7, e SVGs exportados): 🟢

   | Nome | Cor |
   |---|---|
   | Violeta 365 | `#7c4fe0` |
   | Berinjela noturna | `#1f1a2e` |
   | Violeta claro | `#a487f2` |
   | Névoa | `#f4f2f7` |
   | Cinza-violeta (apoio) | `#6b6880` |
   | Grafite violeta (apoio) | `#4a4560` |
   | Lavanda (apoio) | `#c9c3dc` |

   - Origem no legado: `_reversa_sdd/design-system/color-palette.md#1. Identidade da marca` (laranja de matiz 38)
   - Tipo: alterada
7. **RN-07:** A ferramenta de leitura de páginas usa por padrão a **extração interna** do gateway. O extrator externo (`defuddle`) vem desativado de fábrica, não aponta mais para o serviço dos autores originais e pode ser ligado por configuração quando existir um serviço próprio do Base365. O plano para criar esse serviço está em `_reversa_forward/backlog/extrator-proprio-base365.md`. 🟢
   - Origem no legado: `_reversa_sdd/code-analysis.md#Módulo: tools`; decisão em `_reversa_sdd/brainstorms/001-rebrand-goclaw-base365/decision.md#Investigação do fetch.goclaw.sh`
   - Tipo: alterada

## 5. Requisitos Funcionais

| ID | Requisito | Prioridade | Critério de aceite | Confidência |
|----|-----------|------------|--------------------|-------------|
| RF-01 | Substituir todas as grafias de "goclaw" (`goclaw`, `GoClaw`, `GOCLAW`, `Goclaw`) pela grafia equivalente de "base365" (`base365`, `Base365`, `BASE365`) no conteúdo de todos os arquivos de texto do repositório. | Must | Uma busca sem diferenciar maiúsculas por "goclaw" no conteúdo retorna 0 ocorrências fora da lista de exceções. | 🟢 |
| RF-02 | Renomear todo arquivo ou pasta cujo nome contenha "goclaw". | Must | Uma listagem de caminhos com "goclaw" no nome retorna 0 itens (hoje são 12). | 🟢 |
| RF-03 | O caminho do módulo Go passa a ser `github.com/edyoCampos/base365`. | Must | O arquivo de módulo declara o caminho novo, e nenhum import aponta para o caminho antigo. | 🟢 |
| RF-04 | O executável e a interface de linha de comando (CLI) se chamam `base365`, inclusive nos textos de ajuda, exemplos e documentação. | Must | A compilação gera o executável `base365`, e `base365 --help` não menciona GoClaw. | 🟢 |
| RF-05 | Todas as variáveis de ambiente do produto usam o prefixo `BASE365_`, incluindo os exemplos de `.env`, os arquivos de compose, os scripts e a documentação. | Must | Nenhuma das 152 variáveis atuais é lida com o prefixo antigo, e os exemplos usam `BASE365_*`. | 🟢 |
| RF-06 | O diretório de dados do usuário passa a ser `~/.base365/` (dados, área de trabalho, segredos). | Must | Numa máquina limpa, o app desktop cria `~/.base365/` e nada em `~/.goclaw/`. | 🟢 |
| RF-07 | Os cabeçalhos HTTP próprios do produto seguem a RN-05, no servidor e em todos os consumidores (interface web, app desktop, scripts de deploy). | Must | Um webhook assinado com `X-Base365-Signature` é aceito. Um assinado só com o cabeçalho antigo é rejeitado. | 🟢 |
| RF-08 | As chaves de API novas seguem a RN-04. | Must | Uma chave criada pelo painel começa com `base365_`. | 🟢 |
| RF-09 | Os controles de segurança da RN-03 bloqueiam e removem variáveis `BASE365_*` com a mesma cobertura de hoje. | Must | Para cada um dos 4 controles: uma credencial de CLI com variável `BASE365_X` é rejeitada; uma variável `BASE365_GATEWAY_TOKEN` não chega ao processo filho; um comando que imprime `$BASE365_ENCRYPTION_KEY` é bloqueado; a workstation rejeita a variável `BASE365_X`. | 🟢 |
| RF-10 | O app desktop exibe a identidade Base365 em todos os metadados: nome do produto, nome do executável, empresa, e-mail de contato, nome do serviço no cofre de senhas do sistema operacional (keyring) e título da janela. | Must | Os metadados do executável e o keyring mostram "Base365", sem "GoClaw". | 🟢 |
| RF-11 | O atualizador automático do desktop consulta releases no repositório `edyoCampos/base365`. | Must | O atualizador consulta o repositório novo, e nenhuma URL do repositório antigo aparece no código. | 🟢 |
| RF-12 | Os textos exibidos ao usuário final e ao admin usam Base365 em todos os idiomas existentes: catálogos de mensagens do backend e da interface web e desktop (en, vi, zh, ko, ru, quando existirem). | Must | Uma busca por "goclaw" nos catálogos retorna 0, e as chaves continuam com paridade entre idiomas. | 🟢 |
| RF-13 | Os textos de contexto enviados aos agentes (templates de bootstrap e conteúdo semeado pelo banco) se referem ao produto como Base365. | Must | Um agente recém-criado, perguntado sobre a plataforma em que roda, não menciona GoClaw. | 🟡 |
| RF-14 | O esquema do banco (migrations do PostgreSQL e esquema completo do SQLite) não contém "goclaw" em comentários, conteúdo semeado nem chaves internas, por exemplo `_goclaw_recovery`. | Must | Um banco criado do zero em cada motor não tem "goclaw" em nenhum valor semeado. | 🟢 |
| RF-15 | O nome padrão do serviço de telemetria passa a ser `base365-gateway`. | Should | Sem configuração explícita, os traces são emitidos com o nome `base365-gateway`. | 🟢 |
| RF-16 | Os pacotes da interface web e do desktop têm nomes Base365. | Should | Os manifestos dos dois pacotes declaram nomes com "base365". | 🟢 |
| RF-17 | Os endereços externos do produto (documentação, site, serviço de extração, scripts de instalação, e-mail de contato) apontam provisoriamente para o GitHub Pages do repositório (`https://edyocampos.github.io/base365/`), enquanto não houver domínio próprio. O registro de imagens segue a premissa `ghcr.io/edyocampos/base365`. Uma lista versionada no repositório registra cada endereço provisório e o recurso definitivo a providenciar. | Must | Nenhum endereço externo contém "goclaw". Nenhum endereço usa domínio inexistente ou de terceiros. A lista de pendências tem um item para cada endereço provisório. | 🟢 |
| RF-18 | A configuração padrão da ferramenta de leitura de páginas segue a RN-07, inclusive o valor semeado no banco, o valor aplicado em bancos sem configuração e o texto de ajuda da interface web e desktop. | Must | Uma instalação nova grava o extrator `defuddle` desativado, sem o endereço dos autores originais, e a leitura de uma página pública devolve conteúdo pela extração interna. Ligar o `defuddle` com um endereço configurado volta a usá-lo. O endereço do serviço próprio consta na lista de pendências (RF-17). | 🟢 |
| RF-19 | Os tokens de cor da interface web (shadcn/ui + Tailwind 4) e do desktop, nos temas claro e escuro, usam os valores do documento `Base365 logo e ícone.pdf` ("Tokens de cor para produto", V1): primária `#7c4fe0` no claro e `#a487f2` no escuro, fundo Névoa e card branco no claro, fundo `#15111f` no escuro, semânticas (success, warning, destructive, info) e escala `chart-1…5` conforme o documento. O desktop abre no escuro por padrão. | Must | A cor primária renderizada é `#7c4fe0` no claro e `#a487f2` no escuro, e cada token web e desktop corresponde ao valor do documento. O texto sobre a primária escura é Berinjela, nunca branco. | 🟢 |
| RF-20 | Nenhuma classe utilitária de cor laranja (`orange-*`) permanece nos componentes da web e do desktop. Classes de cor âmbar que representam a marca são convertidas para a paleta da marca. | Must | Uma busca por `orange-` nos componentes retorna 0. Cada ocorrência de `amber-*` fica classificada como aviso ou decoração numa lista revisável. A conversão do âmbar segue a regra do documento: âmbar que comunica atenção (limite, expira, pendente, alerta, validação fraca, reconexão) continua âmbar e usa o token `warning`. Âmbar decorativo (ícone de destaque, favorito, tag "novo", gradiente, hover) vira violeta (`primary`/`accent`). Âmbar nunca é cor de série em gráficos de status, que usam só `chart-1…5`. Os badges de status success, warning e info substituem emerald, amber e sky. | 🟢 |
| RF-21 | Logo, ícone da aplicação, favicons (16, 32 e 64 px) e ícones do executável desktop são os ativos oficiais em `Base365 logo e ícone/export/`, nas variantes clara e escura. | Must | Nenhum ativo visual antigo permanece. O favicon da web e o ícone do desktop correspondem aos arquivos de marca. | 🟢 |
| RF-22 | A documentação do repositório (README principal, os 30 READMEs traduzidos, `docs/` e arquivos de instrução para agentes de código) usa Base365. | Should | Uma busca por "goclaw" nesses arquivos retorna 0. | 🟢 |
| RF-23 | Uma checagem automática, que roda localmente e pode ser plugada no CI, falha quando "goclaw" aparece no conteúdo ou nos nomes de arquivo fora da lista de exceções. | Must | A checagem passa no estado final e falha num commit de teste que reintroduz "GoClaw" num arquivo qualquer. | 🟢 |
| RF-24 | Remover as menções a OpenClaw, ZeroClaw e PicoClaw: comentários de procedência no código, agradecimentos no README e a frase "A Go port of OpenClaw" na descrição da CLI. | Should | Uma busca por "openclaw", "zeroclaw" e "picoclaw" (sem diferenciar maiúsculas) retorna 0. Antes da remoção, o plano verifica se as licenças dos projetos exigem manter avisos de autoria e registra o resultado. | 🟢 |

## 6. Requisitos Não Funcionais

| Tipo | Requisito | Evidência ou justificativa | Confidência |
|------|-----------|----------------------------|-------------|
| Segurança | Nenhum controle de segurança perde cobertura com a troca de prefixo (RN-03). Cada controle ganha um teste automatizado com o prefixo novo. | `internal/crypto/env_denylist.go:57`, `internal/workstation/security/allowlist.go:189-190`, `internal/tools/env_scrub.go:23`, `internal/tools/shell_deny_groups.go:233-236`, `ui/web/src/pages/cli-credentials/cli-credential-grant-env-section.tsx:35` | 🟢 |
| Privacidade | Com o padrão novo (RN-07), nenhuma URL lida pelos agentes é enviada a terceiros. | `_reversa_sdd/brainstorms/001-rebrand-goclaw-base365/decision.md#Investigação do fetch.goclaw.sh` | 🟢 |
| Confiabilidade | Com o extrator externo desligado ou indisponível, a leitura de páginas continua funcionando pela extração interna. | `cmd/gateway_builtin_tools.go:42` | 🟢 |
| Compatibilidade de build | O projeto compila nas duas edições (PostgreSQL e SQLite/desktop), passa na análise estática e nos testes unitários, e as interfaces web e desktop geram build de produção sem erro. | `CLAUDE.md#Post-Implementation Checklist` | 🟢 |
| Paridade de bancos | Toda alteração de esquema ou conteúdo semeado vale igualmente para PostgreSQL e SQLite. | `CLAUDE.md` (Migrations dual-DB), `_reversa_sdd/architecture.md#6.4 Paridade PostgreSQL × SQLite` | 🟢 |
| Acessibilidade | Texto e componentes interativos nas cores novas atingem contraste mínimo de 4,5:1 para texto normal e 3:1 para texto grande e ícones, nos dois temas (WCAG 2.1, nível AA). | Recomendação de mercado. O legado não declara meta de contraste. | 🟡 |
| Manutenibilidade | A lista de exceções da checagem (RF-23) e a lista de pendências externas (RF-17) ficam versionadas no repositório, com uma justificativa por item. | Pedido do usuário por "algo que me lembre disso depois" | 🟢 |
| Paridade de i18n | Toda chave de texto alterada continua presente em todos os idiomas existentes. | `CLAUDE.md#Key Patterns` (i18n) | 🟢 |

## 7. Critérios de Aceitação

```gherkin
Cenário: Repositório sem a marca antiga (RF-01, RF-02, RF-22, RF-23)
  Dado o repositório no estado final da feature
  Quando a checagem automática de marca é executada
  Então ela termina com sucesso
  E nenhuma ocorrência de "goclaw", em qualquer grafia, é encontrada no conteúdo ou nos nomes de arquivo fora da lista de exceções

Cenário: Reintrodução da marca antiga é barrada (RF-23)
  Dado o repositório no estado final da feature
  Quando um commit adiciona o texto "GoClaw" a qualquer arquivo fora da lista de exceções
  Então a checagem automática falha e indica o arquivo e a linha

Cenário: Build nas duas edições (RF-03, RF-04, RF-16)
  Dado o código no estado final
  Quando o projeto é compilado nas edições PostgreSQL e SQLite e as interfaces web e desktop são construídas
  Então todos os builds terminam sem erro
  E o executável gerado se chama "base365"

Cenário: Operador configura o produto (RF-05, RF-06)
  Dado uma máquina sem instalação anterior
  Quando o operador define variáveis BASE365_* e inicia o produto
  Então o produto lê essas variáveis
  E grava os dados do desktop em ~/.base365/

Cenário: Variável com prefixo antigo é ignorada (RF-05, RN-02)
  Dado uma máquina com apenas GOCLAW_CONFIG definida
  Quando o produto é iniciado
  Então a configuração indicada por GOCLAW_CONFIG não é carregada

Cenário: Webhook assinado com cabeçalho novo (RF-07)
  Dado um webhook configurado com segredo HMAC
  Quando um sistema externo envia uma chamada assinada no cabeçalho X-Base365-Signature
  Então a chamada é aceita

Cenário: Webhook com cabeçalho antigo é rejeitado (RF-07, RN-02)
  Dado um webhook configurado com segredo HMAC
  Quando a chamada chega assinada só no cabeçalho X-GoClaw-Signature
  Então a chamada é rejeitada por falta de autenticação

Cenário: Nova chave de API (RF-08)
  Dado um admin no painel web
  Quando ele cria uma chave de API
  Então a chave exibida começa com "base365_"

Cenário: Segredo do produto não chega ao processo filho (RF-09)
  Dado o gateway rodando com BASE365_GATEWAY_TOKEN definida
  Quando um agente executa um comando sem credencial associada
  Então o ambiente do processo filho não contém BASE365_GATEWAY_TOKEN

Cenário: Credencial de CLI tenta sobrescrever variável do produto (RF-09)
  Dado um admin criando uma credencial de CLI
  Quando ele inclui a variável BASE365_ENCRYPTION_KEY no ambiente da credencial
  Então a interface e o servidor rejeitam a variável

Cenário: Comando de exfiltração com prefixo novo é bloqueado (RF-09)
  Dado um agente com a ferramenta de execução de comandos habilitada
  Quando o agente tenta executar "echo $BASE365_ENCRYPTION_KEY"
  Então o comando é bloqueado pelos padrões de negação

Cenário: Workstation remota recusa variável do produto (RF-09)
  Dado uma workstation remota vinculada
  Quando uma execução pede a variável BASE365_X no ambiente
  Então a workstation rejeita a variável

Cenário: Identidade do app desktop (RF-10, RF-11, RF-21)
  Dado o executável desktop gerado
  Quando o usuário o instala e abre
  Então nome, ícone, janela e metadados exibem Base365
  E o atualizador consulta releases do repositório edyoCampos/base365

Cenário: Agente se apresenta como Base365 (RF-12, RF-13, RF-14)
  Dado um banco criado do zero e um agente novo
  Quando o usuário final pergunta ao agente em que plataforma ele roda
  Então a resposta não menciona GoClaw

Cenário: Telemetria com nome novo (RF-15)
  Dado a telemetria habilitada sem nome de serviço configurado
  Quando o gateway emite um trace
  Então o nome do serviço é "base365-gateway"

Cenário: Endereços externos provisórios registrados (RF-17)
  Dado o repositório no estado final
  Quando o desenvolvedor abre a lista de pendências externas
  Então cada endereço Base365 que depende de recurso não providenciado aparece com o recurso a providenciar e o impacto se faltar

Cenário: Extração interna por padrão (RF-18)
  Dado uma instalação nova com a configuração padrão
  Quando um agente lê uma página pública
  Então o conteúdo é devolvido pela extração interna
  E o painel mostra o extrator externo desativado, sem o endereço dos autores originais

Cenário: Painel nas cores da marca (RF-19, RF-20, RF-21)
  Dado o painel web e o app desktop nos temas claro e escuro
  Quando o admin navega pelas telas principais
  Então a cor primária e o logo seguem o manual da marca
  E nenhuma classe laranja permanece nos componentes

Cenário: Contraste insuficiente é detectado (RNF de acessibilidade)
  Dado um par de cores de texto e fundo derivado da paleta nova
  Quando o contraste medido fica abaixo de 4,5:1 para texto normal
  Então o par é rejeitado na revisão e ajustado antes da entrega

Cenário: Referências a projetos de origem removidas (RF-24)
  Dado o repositório no estado final
  Quando é feita uma busca por "openclaw", "zeroclaw" e "picoclaw"
  Então nenhuma ocorrência é encontrada
```

## 8. Prioridade MoSCoW

| Item | MoSCoW | Justificativa |
|------|--------|---------------|
| RF-01 a RF-14, RF-17 a RF-21, RF-23 | Must | Formam a restrição inegociável "goclaw não deve mais existir" e a paleta compatível com o manual. Sem eles o lançamento não acontece. |
| RF-09 e RNF de segurança | Must | Uma troca de prefixo sem cobertura equivalente abre brecha de segredo. |
| RF-15, RF-16, RF-22 | Should | Não são vistos pelo usuário final, mas deixam rastro da marca antiga para operador e desenvolvedor. |
| RF-24 | Should | Decisão tomada (remover). Não bloqueia a identidade visível. |
| RNF de acessibilidade | Should | O legado não tinha meta de contraste. A paleta nova é a oportunidade de fixar uma. |
| Redesenho do design system (fonte Nunito na interface, tokens unificados entre web e desktop) | Won't (nesta feature) | Não-objetivo declarado em `pre-spec.md#Não-objetivos`. |
| Criar o serviço de extração próprio, os workflows de CI e renomear bots nos provedores | Won't (nesta feature) | Pendências externas do usuário (`decision.md#Pendências externas a providenciar`). |

## 9. Esclarecimentos

### Sessão 2026-10-02

- **Q:** (D-01) Como as cores do manual viram tokens de interface (primária no escuro, fundos, semânticas, escala de gráficos)?
  **R:** Usar o documento `Base365 logo e ícone.pdf` ("Tokens de cor para produto", V1), que define todos os tokens web e desktop nos dois temas, os badges de status e a escala `chart-1…5`, com contraste WCAG 2.1 AA calculado. A primária escura é `#a487f2` com texto Berinjela. O muted-foreground claro é `#5d5873` (e não `#6b6880`, que falha o contraste).
- **Q:** (D-01) Quais cores semânticas e escala de gráficos?
  **R:** As do mesmo documento: success `#157a4a`/`#3ecf8e`, warning `#f0a92a`/`#f5b740`, destructive `#c62a3a`/`#ef6b73`, info nos badges, e `chart-1…5` (violeta, verde-azulado, laranja, azul-céu, neutro violeta).
- **Q:** (D-01) Quais classes `amber-*` continuam âmbar e quais viram violeta?
  **R:** Regra do âmbar do documento (seção 06): se remover a cor faz o usuário perder a informação de que algo precisa de atenção, é aviso e usa `warning`; senão é decoração e vira violeta. Textos com "atenção", "limite", "expira" e "pendente" ficam âmbar. Âmbar nunca é série de gráfico de status.
- **Q:** (D-02) As menções a OpenClaw, ZeroClaw e PicoClaw continuam?
  **R:** Remover todas (comentários de procedência, README e descrição da CLI).
- **Q:** (D-03) Qual valor usar nos endereços externos provisórios?
  **R:** GitHub Pages do repositório, por enquanto, para não quebrar nada.

## 10. Lacunas

- 🟡 **Licenças (RF-24):** a remoção das menções a OpenClaw, ZeroClaw e PicoClaw foi decidida, mas o plano precisa verificar se as licenças exigem manter avisos de autoria no código portado, e registrar o resultado.
- 🟡 **Serviço de extração (RN-07, RF-18):** decidido em 2026-10-02 usar a extração interna por padrão e deixar o serviço próprio (Cloudflare Worker com Defuddle) para depois. O plano está em `_reversa_forward/backlog/extrator-proprio-base365.md` e o item entra na lista de pendências externas (RF-17).
- 🟡 **E-mail de contato (RF-10, RF-17):** o GitHub Pages não resolve e-mail. O endereço fica como pendência registrada na lista, sem valor definitivo.

**Premissas assumidas** (sem marcador de dúvida, reversíveis no `/reversa-clarify`):
- 🟡 O registro de imagens é `ghcr.io/edyocampos/base365`, inferido do repositório, com as mesmas variantes de hoje.
- 🟡 Os 30 READMEs traduzidos são mantidos, só com o nome trocado.
- 🟡 O endereço do GitHub Pages é `https://edyocampos.github.io/base365/`, derivado do repositório `edyoCampos/base365`.
- 🟡 Como não há produção, as migrations existentes podem ser alteradas no próprio arquivo, sem migration nova.

## Pendências de Qualidade

- **Q-018 (sem nome de produto ou biblioteca), reprovado de propósito:** um rebrand precisa citar identificadores concretos para ser verificável. São eles: o caminho do módulo, os prefixos de variável, os cabeçalhos, o formato de chave e os nomes de cor. As citações técnicas (SHA-256, HMAC, WCAG 2.1) só preservam comportamento que já existe ou fixam uma meta mensurável. Nenhuma prescreve a forma de implementar.
- **Q-017 (o quê, não o como), atenção:** a decisão de alterar as migrations no próprio arquivo aparece só como premissa reversível na seção 10. Quem decide é o `/reversa-plan`.

## 11. Histórico de alterações

| Data | Alteração | Autor |
|------|-----------|-------|
| 2026-10-01 | Versão inicial gerada por `/reversa-requirements` a partir da sessão de ideação `001-rebrand-goclaw-base365` | reversa |
| 2026-10-02 | `/reversa-clarify`: resolvidas D-01 (tokens de cor, escala de gráficos e regra do âmbar via PDF de tokens), D-02 (remover menções) e D-03 (GitHub Pages). RF-17, RF-19, RF-20 e RF-24 reescritos. | reversa |
| 2026-10-02 | RN-07 e RF-18 ajustados: extração interna por padrão, serviço próprio no backlog (`backlog/extrator-proprio-base365.md`) | reversa |
