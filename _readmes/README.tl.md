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
  <a href="https://edyocampos.github.io/base365">Dokumentasyon</a> •
  <a href="https://edyocampos.github.io/base365/#quick-start">Mabilis na Simula</a> •
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

Ang **Base365** ay isang multi-agent AI gateway na nagkokonekta ng mga LLM sa iyong mga kasangkapan, channel, at datos — na inilunsad bilang isang Go binary na walang runtime dependencies. Inooorkestra nito ang mga agent team at inter-agent delegation sa 20+ LLM provider na may ganap na multi-tenant isolation.

🌐 **Mga Wika:**
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

## Ano ang Nagpapaiba Rito

- **Mga Agent Team at Orkestrasyon** — Mga team na may pinagsamang task board, inter-agent delegation (sync/async), at hybrid agent discovery
- **Multi-Tenant PostgreSQL** — Bawat user ay may sariling workspace, context file, naka-encrypt na API key (AES-256-GCM), at nakahiwalay na session
- **Iisang Binary** — ~25 MB static Go binary, walang Node.js runtime, <1s na pagsisimula, tumatakbo sa $5 VPS
- **Seguridad sa Produksyon** — 5-layer na sistema ng pahintulot (gateway auth → global tool policy → per-agent → per-channel → owner-only) kasama ang rate limiting, pagtuklas ng prompt injection, proteksyon sa SSRF, shell deny patterns, at AES-256-GCM encryption
- **20+ LLM Provider** — Anthropic (native HTTP+SSE na may prompt caching), OpenAI, OpenRouter, Groq, DeepSeek, Gemini, Mistral, xAI, MiniMax, Cohere, Perplexity, DashScope, Bailian, Zai, Ollama, Ollama Cloud, Claude CLI, Codex, ACP, at anumang OpenAI-compatible na endpoint
- **7 Messaging Channel** — Telegram, Discord, Slack, Zalo OA, Zalo Personal, Feishu/Lark, WhatsApp
- **Extended Thinking** — Bawat provider ay may thinking mode (Anthropic budget tokens, OpenAI reasoning effort, DashScope thinking budget) na may suporta sa streaming
- **Heartbeat** — Pana-panahong pag-check-in ng agent sa pamamagitan ng HEARTBEAT.md checklist na may suppress-on-OK, aktibong oras, retry logic, at paghahatid sa channel
- **Pag-iskedyul at Cron** — `at`, `every`, at cron expressions para sa mga automated na gawain ng agent na may lane-based na concurrency
- **Observability** — Built-in na LLM call tracing na may mga span at prompt cache metrics, opsyonal na OpenTelemetry OTLP export

## Claw Ecosystem

|                 | **Base365**                              |
| --------------- | --------------------------------------- |
| Wika            | **Go**                                  |
| Laki ng binary  | **~25 MB** (base) / **~36 MB** (+ OTel) |
| Docker image    | **~50 MB** (Alpine)                     |
| RAM (walang gawa) | **~35 MB**                              |
| Pagsisimula     | **< 1 s**                               |
| Target hardware | **$5 VPS+**                             |

| Tampok                     | **Base365**                     |
| -------------------------- | ------------------------------ |
| Multi-tenant (PostgreSQL)  | ✅                             |
| MCP integration            | ✅ (stdio/SSE/streamable-http) |
| Mga agent team             | ✅ Task board + mailbox        |
| Pagpapatibay ng seguridad  | ✅ 5-layer defense             |
| OTel observability         | ✅ OTLP (opt-in build tag)     |
| Prompt caching             | ✅ Anthropic + OpenAI-compat   |
| Knowledge graph            | ✅ LLM extraction + traversal  |
| Sistema ng skill           | ✅ BM25 + pgvector hybrid      |
| Lane-based scheduler       | ✅ (main/subagent/team/cron)   |
| Mga messaging channel      | 7+                             |
| Mga kasama na app          | Web dashboard                  |
| Live Canvas / Boses        | TTS (4 provider)               |
| Mga LLM provider           | **20+**                        |
| Bawat user ay may workspace | ✅ (PostgreSQL)                |
| Naka-encrypt na lihim      | ✅ AES-256-GCM sa DB           |

## Arkitektura

<p align="center">
  <img src="../_statics/architecture.jpg" alt="Base365 Architecture" width="800" />
</p>

## Mabilis na Simula

**Mga Kinakailangan:** Go 1.26+, PostgreSQL 18 na may pgvector, Docker (opsyonal)

### Mula sa Source

```bash
git clone https://github.com/edyoCampos/base365.git && cd base365
make build
./base365 onboard        # Interactive setup wizard
source .env.local && ./base365
```

### Gamit ang Docker

```bash
# Gumawa ng .env na may auto-generated na mga lihim
chmod +x prepare-env.sh && ./prepare-env.sh

# Magdagdag ng kahit isang BASE365_*_API_KEY sa .env, pagkatapos:
make up

# Web Dashboard sa http://localhost:18790
# Health check: curl http://localhost:18790/health
```

Kapag nakatakda ang mga environment variable na `BASE365_*_API_KEY`, ang gateway ay awtomatikong mag-onboard nang walang interactive na mga tanong — nakikita ang provider, nagpapatakbo ng mga migration, at naglalagay ng default na datos.

> Para sa mga variant ng build (OTel, Tailscale, Redis), mga Docker image tag, at compose overlay, tingnan ang [Gabay sa Deployment](https://edyocampos.github.io/base365/#deploy-docker-compose).

## Multi-Agent Orkestrasyon

Sinusuportahan ng Base365 ang mga agent team at inter-agent delegation — ang bawat agent ay tumatakbo gamit ang sariling pagkakakilanlan, kasangkapan, LLM provider, at context file.

### Delegasyon ng Agent

<p align="center">
  <img src="../_statics/agent-delegation.jpg" alt="Agent Delegation" width="700" />
</p>

| Mode | Paano gumagana | Pinakamainam para sa |
|------|----------------|----------------------|
| **Sync** | Ang Agent A ay nagtatanong sa Agent B at **naghihintay** ng sagot | Mabilis na paghahanap, pagtitiyak ng katotohanan |
| **Async** | Ang Agent A ay nagtatanong sa Agent B at **nagpapatuloy**. Ang B ay mag-aanunsyo sa ibang pagkakataon | Matagal na gawain, ulat, malalim na pagsusuri |

Ang mga agent ay nakikipag-ugnayan sa pamamagitan ng malinaw na **mga link ng pahintulot** na may kontrol sa direksyon (`outbound`, `inbound`, `bidirectional`) at mga limitasyon sa concurrency sa antas ng bawat link at bawat agent.

### Mga Agent Team

<p align="center">
  <img src="../_statics/agent-teams.jpg" alt="Agent Teams Workflow" width="800" />
</p>

- **Pinagsamang task board** — Lumikha, mag-claim, kumpletuhin, at maghanap ng mga gawain na may mga dependency na `blocked_by`
- **Mailbox ng team** — Direktang peer-to-peer na pagmemensahe at mga broadcast
- **Mga Kasangkapan**: `team_tasks` para sa pamamahala ng gawain, `team_message` para sa mailbox

> Para sa mga detalye ng delegasyon, mga link ng pahintulot, at kontrol ng concurrency, tingnan ang [dokumentasyon ng Mga Agent Team](https://edyocampos.github.io/base365/#teams-what-are-teams).

## Mga Built-in na Kasangkapan

| Kasangkapan        | Grupo         | Paglalarawan                                                  |
| ------------------ | ------------- | ------------------------------------------------------------- |
| `read_file`        | fs            | Basahin ang nilalaman ng file (na may virtual FS routing)     |
| `write_file`       | fs            | Sumulat/lumikha ng mga file                                   |
| `edit_file`        | fs            | Mag-apply ng mga targeted na pagbabago sa mga kasalukuyang file |
| `list_files`       | fs            | Ilista ang mga nilalaman ng direktoryo                        |
| `search`           | fs            | Maghanap ng nilalaman ng file ayon sa pattern                 |
| `glob`             | fs            | Hanapin ang mga file ayon sa glob pattern                     |
| `exec`             | runtime       | Mag-execute ng mga shell command (na may approval workflow)   |
| `web_search`       | web           | Maghanap sa web (Brave, DuckDuckGo)                           |
| `web_fetch`        | web           | I-fetch at i-parse ang nilalaman ng web                       |
| `memory_search`    | memory        | Maghanap sa pangmatagalang memorya (FTS + vector)             |
| `memory_get`       | memory        | Kunin ang mga entry sa memorya                                |
| `skill_search`     | —             | Maghanap ng mga skill (BM25 + embedding hybrid)               |
| `knowledge_graph_search` | memory  | Maghanap ng mga entity at dumaan sa mga relasyon ng knowledge graph |
| `create_image`     | media         | Paglikha ng imahe (DashScope, MiniMax)                        |
| `create_audio`     | media         | Paglikha ng audio (OpenAI, ElevenLabs, MiniMax, Suno)         |
| `create_video`     | media         | Paglikha ng video (MiniMax, Veo)                              |
| `read_document`    | media         | Pagbabasa ng dokumento (Gemini File API, provider chain)      |
| `read_image`       | media         | Pagsusuri ng imahe                                            |
| `read_audio`       | media         | Pag-transcribe at pagsusuri ng audio                          |
| `read_video`       | media         | Pagsusuri ng video                                            |
| `message`          | messaging     | Magpadala ng mga mensahe sa mga channel                       |
| `tts`              | —             | Text-to-Speech synthesis                                      |
| `spawn`            | —             | Mag-spawn ng subagent                                         |
| `subagents`        | sessions      | Kontrolin ang mga tumatakbong subagent                        |
| `team_tasks`       | teams         | Pinagsamang task board (ilista, lumikha, mag-claim, kumpletuhin, maghanap) |
| `team_message`     | teams         | Mailbox ng team (magpadala, mag-broadcast, magbasa)           |
| `sessions_list`    | sessions      | Ilista ang mga aktibong session                               |
| `sessions_history` | sessions      | Tingnan ang kasaysayan ng session                             |
| `sessions_send`    | sessions      | Magpadala ng mensahe sa isang session                         |
| `sessions_spawn`   | sessions      | Mag-spawn ng bagong session                                   |
| `session_status`   | sessions      | Suriin ang katayuan ng session                                |
| `cron`             | automation    | Mag-iskedyul at pamahalaan ang mga cron job                   |
| `gateway`          | automation    | Pangangasiwa ng gateway                                       |
| `browser`          | ui            | Awtomatikong paggamit ng browser (mag-navigate, mag-click, mag-type, screenshot) |
| `announce_queue`   | automation    | Pag-aanunsyo ng async na resulta (para sa mga async na delegasyon) |

## Dokumentasyon

Buong dokumentasyon sa **[edyocampos.github.io/base365](https://edyocampos.github.io/base365)** — o i-browse ang source sa [`base365-docs/`](https://edyocampos.github.io/base365)

| Seksyon | Mga Paksa |
|---------|-----------|
| [Pagsisimula](https://edyocampos.github.io/base365/#what-is-base365) | Pag-install, Mabilis na Simula, Konfiguraksyon, Pagsasaliksik sa Web Dashboard |
| [Mga Pangunahing Konsepto](https://edyocampos.github.io/base365/#how-base365-works) | Agent Loop, Mga Session, Mga Kasangkapan, Memorya, Multi-Tenancy |
| [Mga Agent](https://edyocampos.github.io/base365/#creating-agents) | Paglikha ng Mga Agent, Mga Context File, Personalidad, Pagbabahagi at Access |
| [Mga Provider](https://edyocampos.github.io/base365/#providers-overview) | Anthropic, OpenAI, OpenRouter, Gemini, DeepSeek, +15 pa |
| [Mga Channel](https://edyocampos.github.io/base365/#channels-overview) | Telegram, Discord, Slack, Feishu, Zalo, WhatsApp, WebSocket |
| [Mga Agent Team](https://edyocampos.github.io/base365/#teams-what-are-teams) | Mga Team, Task Board, Pagmemensahe, Delegasyon at Handoff |
| [Advanced](https://edyocampos.github.io/base365/#custom-tools) | Mga Custom na Kasangkapan, MCP, Mga Skill, Cron, Sandbox, Hooks, RBAC |
| [Deployment](https://edyocampos.github.io/base365/#deploy-docker-compose) | Docker Compose, Database, Seguridad, Observability, Tailscale |
| [Sanggunian](https://edyocampos.github.io/base365/#cli-commands) | Mga CLI Command, REST API, WebSocket Protocol, Mga Environment Variable |

## Pagsusuri

```bash
go test ./...                                    # Mga unit test
go test -v ./tests/integration/ -timeout 120s    # Mga integration test (nangangailangan ng tumatakbong gateway)
```

## Katayuan ng Proyekto

Tingnan ang [CHANGELOG.md](CHANGELOG.md) para sa detalyadong katayuan ng mga tampok kasama ang kung ano na ang nasubok sa produksyon at kung ano pa ang isinasagawa.

## Lisensya

MIT
