<p align="center">
  <img src="../_statics/base365.png" alt="Base365" />
</p>

<h1 align="center">Base365</h1>

<p align="center"><strong>Enterprise AI Agent Platform</strong></p>

<p align="center">
Multi-agent AI gateway built in Go. 20+ LLM providers. 7 channels. Multi-tenant PostgreSQL.<br/>
Single binary. Production-tested. Agents that orchestrate for you.
</p>

<p align="center">
  <a href="https://edyocampos.github.io/base365">Dokumentation</a> •
  <a href="https://edyocampos.github.io/base365/#quick-start">Snabbstart</a> •
  <a href="https://x.com/nlb_io">Twitter / X</a>
</p>

<p align="center">
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go_1.26-00ADD8?style=flat-square&logo=go&logoColor=white" alt="Go" /></a>
  <a href="https://www.postgresql.org/"><img src="https://img.shields.io/badge/PostgreSQL_18-316192?style=flat-square&logo=postgresql&logoColor=white" alt="PostgreSQL" /></a>
  <a href="https://www.docker.com/"><img src="https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=docker&logoColor=white" alt="Docker" /></a>
  <a href="https://developer.mozilla.org/en-US/docs/Web/API/WebSocket"><img src="https://img.shields.io/badge/WebSocket-010101?style=flat-square&logo=socket.io&logoColor=white" alt="WebSocket" /></a>
  <a href="https://opentelemetry.io/"><img src="https://img.shields.io/badge/OpenTelemetry-000000?style=flat-square&logo=opentelemetry&logoColor=white" alt="OpenTelemetry" /></a>
  <a href="https://www.anthropic.com/"><img src="https://img.shields.io/badge/Anthropic-191919?style=flat-square&logo=anthropic&logoColor=white" alt="Anthropic" /></a>
  <a href="https://openai.com/"><img src="https://img.shields.io/badge/OpenAI_Compatible-412991?style=flat-square&logo=openai&logoColor=white" alt="OpenAI" /></a>
  <img src="https://img.shields.io/badge/License-MIT-yellow?style=flat-square" alt="License: MIT" />
</p>

**Base365** är en multi-agent AI-gateway som kopplar ihop LLM:er med dina verktyg, kanaler och data — driftsatt som en enda Go-binärfil utan externa körtidsberoenden. Den orkestrerar agentteam och inter-agent-delegering mellan 20+ LLM-leverantörer med full flertenantsisolering.

🌐 **Språk:**
[🇺🇸 English](../README.md) ·
[🇨🇳 简体中文](README.zh-CN.md) ·
[🇯🇵 日本語](README.ja.md) ·
[🇰🇷 한국어](README.ko.md) ·
[🇻🇳 Tiếng Việt](README.vi.md) ·
[🇵🇭 Tagalog](README.tl.md) ·
[🇪🇸 Español](README.es.md) ·
[🇧🇷 Português](README.pt.md) ·
[🇮🇹 Italiano](README.it.md) ·
[🇩🇪 Deutsch](README.de.md) ·
[🇫🇷 Français](README.fr.md) ·
[🇸🇦 العربية](README.ar.md) ·
[🇮🇳 हिन्दी](README.hi.md) ·
[🇷🇺 Русский](README.ru.md) ·
[🇧🇩 বাংলা](README.bn.md) ·
[🇮🇱 עברית](README.he.md) ·
[🇵🇱 Polski](README.pl.md) ·
[🇨🇿 Čeština](README.cs.md) ·
[🇳🇱 Nederlands](README.nl.md) ·
[🇹🇷 Türkçe](README.tr.md) ·
[🇺🇦 Українська](README.uk.md) ·
[🇮🇩 Bahasa Indonesia](README.id.md) ·
[🇹🇭 ไทย](README.th.md) ·
[🇵🇰 اردو](README.ur.md) ·
[🇷🇴 Română](README.ro.md) ·
[🇸🇪 Svenska](README.sv.md) ·
[🇬🇷 Ελληνικά](README.el.md) ·
[🇭🇺 Magyar](README.hu.md) ·
[🇫🇮 Suomi](README.fi.md) ·
[🇩🇰 Dansk](README.da.md) ·
[🇳🇴 Norsk](README.nb.md)

## Vad som skiljer den åt

- **Agentteam och orkestrering** — Team med delade uppgiftstavlor, inter-agent-delegering (synkron/asynkron) och hybrid agentupptäckt
- **Flertenants PostgreSQL** — Per-användararbetsytor, per-användarkontextfiler, krypterade API-nycklar (AES-256-GCM), isolerade sessioner
- **Enstaka binärfil** — ~25 MB statisk Go-binärfil, ingen Node.js-körningstid, <1s uppstartstid, körs på en $5 VPS
- **Produktionssäkerhet** — 5-lagers behörighetssystem (gateway-autentisering → global verktygspolicy → per-agent → per-kanal → ägarexklusivt) plus hastighetsbegränsning, detektering av prompt injection, SSRF-skydd, nekade skalkommandon och AES-256-GCM-kryptering
- **20+ LLM-leverantörer** — Anthropic (native HTTP+SSE med prompt caching), OpenAI, OpenRouter, Groq, DeepSeek, Gemini, Mistral, xAI, MiniMax, Cohere, Perplexity, DashScope, Bailian, Zai, Ollama, Ollama Cloud, Claude CLI, Codex, ACP och valfri OpenAI-kompatibel slutpunkt
- **7 meddelandekanaler** — Telegram, Discord, Slack, Zalo OA, Zalo Personal, Feishu/Lark, WhatsApp
- **Extended Thinking** — Per-leverantörs tänkande-läge (Anthropic budgettokens, OpenAI resonemangsnivå, DashScope thinking budget) med stöd för strömning
- **Heartbeat** — Periodiska agentincheckning via HEARTBEAT.md-checklistor med undertryckning vid OK-status, aktiva timmar, återförsökslogik och kanalleverans
- **Schemaläggning och Cron** — `at`, `every` och cron-uttryck för automatiserade agentuppgifter med filbaserad samtidighetskontroll
- **Observerbarhet** — Inbyggd LLM-anropsspårning med spann och prompt cache-mätvärden, valfri OpenTelemetry OTLP-export

## Claw-ekosystemet

|                 | **Base365**                              |
| --------------- | --------------------------------------- |
| Språk           | **Go**                                  |
| Binärstorlek    | **~25 MB** (bas) / **~36 MB** (+ OTel) |
| Docker-avbild   | **~50 MB** (Alpine)                     |
| RAM (vilande)   | **~35 MB**                              |
| Uppstart        | **< 1 s**                               |
| Målhårdvara     | **$5 VPS+**                             |

| Funktion                   | **Base365**                     |
| -------------------------- | ------------------------------ |
| Flertenants (PostgreSQL)   | ✅                             |
| MCP-integration            | ✅ (stdio/SSE/streamable-http) |
| Agentteam                  | ✅ Uppgiftstavla + brevlåda    |
| Säkerhetshärdning          | ✅ 5-lagers försvar            |
| OTel-observerbarhet        | ✅ OTLP (valfri build-tag)     |
| Prompt caching             | ✅ Anthropic + OpenAI-compat   |
| Kunskapsgraf               | ✅ LLM-extraktion + traversering |
| Färdighetssystem           | ✅ BM25 + pgvector hybrid      |
| Filbaserad schemaläggare   | ✅ (main/subagent/team/cron)   |
| Meddelandekanaler          | 7+                             |
| Medföljande appar          | Webbinstrumentpanel            |
| Live Canvas / Röst         | TTS (4 leverantörer)           |
| LLM-leverantörer           | **20+**                        |
| Per-användararbetsytor     | ✅ (PostgreSQL)                |
| Krypterade hemligheter     | ✅ AES-256-GCM i DB            |

## Arkitektur

<p align="center">
  <img src="../_statics/architecture.jpg" alt="Base365 Architecture" width="800" />
</p>

## Snabbstart

**Förutsättningar:** Go 1.26+, PostgreSQL 18 med pgvector, Docker (valfritt)

### Från källkod

```bash
git clone https://github.com/edyoCampos/base365.git && cd base365
make build
./base365 onboard        # Interaktiv installationsguide
source .env.local && ./base365
```

### Med Docker

```bash
# Generera .env med automatiskt skapade hemligheter
chmod +x prepare-env.sh && ./prepare-env.sh

# Lägg till minst en BASE365_*_API_KEY i .env, sedan:
make up

# Webbinstrumentpanel på http://localhost:18790
# Hälsokontroll: curl http://localhost:18790/health
```

När miljövariabler av typen `BASE365_*_API_KEY` är inställda registreras gatewayen automatiskt utan interaktiva promptar — identifierar leverantör, kör migrationer och sår standarddata.

> För byggvarianter (OTel, Tailscale, Redis), Docker-avbildstaggar och compose-överlägg, se [Driftsättningsguiden](https://edyocampos.github.io/base365/#deploy-docker-compose).

## Multi-agent-orkestrering

Base365 stöder agentteam och inter-agent-delegering — varje agent körs med sin egen identitet, verktyg, LLM-leverantör och kontextfiler.

### Agentdelegering

<p align="center">
  <img src="../_statics/agent-delegation.jpg" alt="Agent Delegation" width="700" />
</p>

| Läge | Hur det fungerar | Passar bäst för |
|------|-----------------|-----------------|
| **Synkron** | Agent A frågar Agent B och **väntar** på svaret | Snabba slagningar, faktakontroller |
| **Asynkron** | Agent A frågar Agent B och **fortsätter**. B meddelar senare | Långa uppgifter, rapporter, djupanalys |

Agenter kommunicerar via explicita **behörighetslänkar** med riktningskontroll (`outbound`, `inbound`, `bidirectional`) och samtidighetsbegränsningar på både per-länk- och per-agentnivå.

### Agentteam

<p align="center">
  <img src="../_statics/agent-teams.jpg" alt="Agent Teams Workflow" width="800" />
</p>

- **Delad uppgiftstavla** — Skapa, ta, slutföra och söka uppgifter med `blocked_by`-beroenden
- **Teamets brevlåda** — Direkt peer-to-peer-meddelanden och utsändningar
- **Verktyg**: `team_tasks` för uppgiftshantering, `team_message` för brevlådan

> För detaljer om delegering, behörighetslänkar och samtidighetskontroll, se [dokumentationen för agentteam](https://edyocampos.github.io/base365/#teams-what-are-teams).

## Inbyggda verktyg

| Verktyg            | Grupp         | Beskrivning                                                  |
| ------------------ | ------------- | ------------------------------------------------------------ |
| `read_file`        | fs            | Läs filinnehåll (med virtuell FS-routing)                    |
| `write_file`       | fs            | Skriv/skapa filer                                            |
| `edit_file`        | fs            | Applicera riktade redigeringar på befintliga filer           |
| `list_files`       | fs            | Lista kataloginnehåll                                        |
| `search`           | fs            | Sök filinnehåll med mönster                                  |
| `glob`             | fs            | Hitta filer med glob-mönster                                 |
| `exec`             | runtime       | Kör skalkommandon (med godkännandeflöde)                     |
| `web_search`       | web           | Sök på webben (Brave, DuckDuckGo)                            |
| `web_fetch`        | web           | Hämta och tolka webbinnehåll                                 |
| `memory_search`    | memory        | Sök i långtidsminnet (FTS + vektor)                          |
| `memory_get`       | memory        | Hämta minnesposter                                           |
| `skill_search`     | —             | Sök färdigheter (BM25 + inbäddningshybrid)                   |
| `knowledge_graph_search` | memory  | Sök entiteter och traversera kunskapsgrafrelationer         |
| `create_image`     | media         | Bildgenerering (DashScope, MiniMax)                          |
| `create_audio`     | media         | Ljudgenerering (OpenAI, ElevenLabs, MiniMax, Suno)           |
| `create_video`     | media         | Videogenerering (MiniMax, Veo)                               |
| `read_document`    | media         | Dokumentläsning (Gemini File API, leverantörskedja)          |
| `read_image`       | media         | Bildanalys                                                   |
| `read_audio`       | media         | Ljudtransskription och analys                                |
| `read_video`       | media         | Videoanalys                                                  |
| `message`          | messaging     | Skicka meddelanden till kanaler                              |
| `tts`              | —             | Text-till-tal-syntes                                         |
| `spawn`            | —             | Starta en underagent                                         |
| `subagents`        | sessions      | Styr körande underagenter                                    |
| `team_tasks`       | teams         | Delad uppgiftstavla (lista, skapa, ta, slutföra, söka)       |
| `team_message`     | teams         | Teamets brevlåda (skicka, sända ut, läsa)                    |
| `sessions_list`    | sessions      | Lista aktiva sessioner                                       |
| `sessions_history` | sessions      | Visa sessionshistorik                                        |
| `sessions_send`    | sessions      | Skicka meddelande till en session                            |
| `sessions_spawn`   | sessions      | Starta en ny session                                         |
| `session_status`   | sessions      | Kontrollera sessionsstatus                                   |
| `cron`             | automation    | Schemalägg och hantera cron-jobb                             |
| `gateway`          | automation    | Gatewayadministration                                        |
| `browser`          | ui            | Webbläsarautomation (navigera, klicka, skriva, skärmdump)    |
| `announce_queue`   | automation    | Asynkron resultatannonsering (för asynkrona delegeringar)    |

## Dokumentation

Fullständig dokumentation på **[edyocampos.github.io/base365](https://edyocampos.github.io/base365)** — eller bläddra i källkoden i [`base365-docs/`](https://edyocampos.github.io/base365)

| Avsnitt | Ämnen |
|---------|-------|
| [Kom igång](https://edyocampos.github.io/base365/#what-is-base365) | Installation, Snabbstart, Konfiguration, Rundtur i webbinstrumentpanelen |
| [Grundläggande begrepp](https://edyocampos.github.io/base365/#how-base365-works) | Agentloop, Sessioner, Verktyg, Minne, Flertenancy |
| [Agenter](https://edyocampos.github.io/base365/#creating-agents) | Skapa agenter, Kontextfiler, Personlighet, Delning och åtkomst |
| [Leverantörer](https://edyocampos.github.io/base365/#providers-overview) | Anthropic, OpenAI, OpenRouter, Gemini, DeepSeek, +15 fler |
| [Kanaler](https://edyocampos.github.io/base365/#channels-overview) | Telegram, Discord, Slack, Feishu, Zalo, WhatsApp, WebSocket |
| [Agentteam](https://edyocampos.github.io/base365/#teams-what-are-teams) | Team, Uppgiftstavla, Meddelanden, Delegering och överlämning |
| [Avancerat](https://edyocampos.github.io/base365/#custom-tools) | Anpassade verktyg, MCP, Färdigheter, Cron, Sandbox, Krokar, RBAC |
| [Driftsättning](https://edyocampos.github.io/base365/#deploy-docker-compose) | Docker Compose, Databas, Säkerhet, Observerbarhet, Tailscale |
| [Referens](https://edyocampos.github.io/base365/#cli-commands) | CLI-kommandon, REST API, WebSocket-protokoll, Miljövariabler |

## Testning

```bash
go test ./...                                    # Enhetstester
go test -v ./tests/integration/ -timeout 120s    # Integrationstester (kräver körande gateway)
```

## Projektstatus

Se [CHANGELOG.md](CHANGELOG.md) för detaljerad funktionsstatus, inklusive vad som testats i produktion och vad som fortfarande pågår.

## Licens

MIT
