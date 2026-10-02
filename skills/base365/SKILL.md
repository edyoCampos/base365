---
name: base365
description: Use this skill when administering, operating, or debugging a Base365 gateway through the Base365 CLI/runtime package. It covers CLI discovery, safe command inspection, gateway health/config diagnostics, agents, skills, MCP/tools, runtime packages, credentials, traces, sessions, channels, providers, cron/jobs, and troubleshooting. Always inspect the live `base365 --help` output first because command availability is version-dependent.
license: Proprietary. Part of Base365 bundled skills.
---

# Base365 Gateway CLI Administration

Use this skill when the user asks you to operate, inspect, administer, or debug a
Base365 gateway through the `base365` CLI/runtime package.

## Operating Rules

1. Identify the exact binary before running commands.
2. Prefer read-only inspection before mutating actions.
3. Never print, paste, or store API keys, bearer tokens, OAuth tokens, database
   credentials, private keys, cookies, or `.env` contents in chat, issues,
   comments, or logs.
4. Scope every action to the requested gateway, tenant, org, agent, team,
   session, channel, or provider.
5. Ask for explicit confirmation before destructive or security-sensitive
   actions.
6. If command help disagrees with this skill, trust the live `--help` output.

Do not run bare `base365` unless the user explicitly wants to start the gateway
server. Use `base365 --help`, `base365 version`, or a subcommand for inspection.

## CLI Discovery

Start with read-only discovery:

```bash
command -v base365
type -a base365
base365 version
base365 --help
base365 <command> --help
```

If `command -v base365` is empty on a managed server, also check deployment
runtime paths before concluding the CLI is absent:

```bash
ls -l /var/lib/base365/data/.runtime/bin/base365
ls -l /app/data/.runtime/bin/base365
```

Distinguish the command surface you found:

- Gateway server binary: running `base365` with no subcommand starts the gateway.
- Operator/admin CLI: subcommands such as `agent`, `skills`, `traces`,
  `sessions`, `providers`, `channels`, `cron`, `config`, `doctor`, `auth`,
  `backup`, `restore`, `migrate`, and `upgrade`.
- Remote operator mode: many admin commands can read `BASE365_SERVER`,
  `BASE365_GATEWAY_URL`, and `BASE365_GATEWAY_TOKEN` from the environment.
  Configure tokens through the shell/session secret manager, not CLI argv or
  pasted text.

Use placeholders in examples:

```bash
base365 traces list
```

## Read-Only Diagnostics

Run safe checks first:

```bash
base365 doctor
base365 config path
base365 config validate
base365 agent list
base365 skills list
base365 sessions list
base365 providers list
base365 channels list
base365 cron list
base365 traces list
```

If a command requires a running gateway, retry with the correct remote target:

```bash
base365 agent list
```

Use `base365 config show` only in a private local terminal when necessary. Treat
the output as sensitive even though the CLI redacts known secret fields; do not
paste it into chats, issues, PRs, or shared logs.

For HTTP-only surfaces, inspect the current API/docs before guessing endpoint
shape. Useful areas include `/health`, `/v1/skills`, `/v1/traces`,
`/v1/packages/runtimes`, `/v1/mcp`, and `/v1/providers` when available.

## Common Workflows

### Agents

Read first:

```bash
base365 agent --help
base365 agent list
```

Before create, update, chat, or delete operations, verify target tenant/org and
agent ID/key. Treat `agent delete` as destructive.

### Skills

Read first:

```bash
base365 skills --help
base365 skills list
base365 skills show <skill>
base365 skills deps status <skill-id-or-path>
base365 skills access get <skill-id>
```

Use dependency scan/check before install:

```bash
base365 skills deps scan <skill-id-or-path>
base365 skills deps check <skill-id-or-path>
```

Mutating skill commands need clear intent and scope:

```bash
base365 skills deps install <skill-id>
base365 skills access set <skill-id> --help
base365 skills grant agent <skill-id> <agent-id>
base365 skills revoke agent <skill-id> <agent-id>
```

System skills are bundled and should not be edited in place. Prefer uploading or
publishing a tenant/custom skill override when customization is required.

### MCP Servers And Tools

First inspect whether this CLI version exposes MCP commands:

```bash
base365 mcp --help
```

If no MCP CLI exists, use the current gateway API/UI/docs for MCP discovery and
permission changes. Do not invent command names. For tool access issues, verify:

- the MCP server is configured and enabled;
- the tool is visible to the requested agent/team;
- permissions or approval rules allow the action;
- gateway logs/traces show the actual failure.

### Runtime Packages

Inspect current runtime/package support before acting:

```bash
base365 packages --help
base365 skills deps status <skill-id-or-path>
```

If there is no packages CLI, use the Packages UI or `/v1/packages/*` API when
available. Install/update/remove package operations are admin-level and may need
approval. Prefer the smallest dependency needed by the selected skill.

### Credentials And Auth

Use read-only status commands first:

```bash
base365 auth status
base365 auth status <provider>
```

For gateway bearer tokens, provider keys, OAuth refresh tokens, and CLI
credentials:

- never echo values;
- do not add them to issues, PRs, chat, or logs;
- prefer environment variables, OS keychain, or gateway credential UI;
- rotate credentials if they were exposed;
- confirm before logout, delete, or replace operations.

### Traces And Failed Runs

Use traces to debug provider errors, tool failures, channel delivery, and stuck
sessions:

```bash
base365 traces --help
base365 traces list --status error
base365 traces get <trace-id> -o json
base365 traces timeline <trace-id>
base365 traces follow --session <session-key>
base365 traces export <trace-id>
```

Redact prompts, tokens, URLs with secrets, headers, and customer data before
posting trace excerpts anywhere public.

### Sessions, Channels, Providers, Cron

Inspect help and list commands:

```bash
base365 sessions --help
base365 channels --help
base365 providers --help
base365 cron --help
```

Treat these as sensitive:

- `sessions delete` and `sessions reset`;
- channel add/delete or credential changes;
- provider add/update/delete and model verification with live credentials;
- cron delete/toggle/run.

## Troubleshooting Playbooks

### Agent Cannot Access A Tool Or MCP

1. Confirm the request is in the intended tenant/org and agent/team.
2. List available skills/tools for that agent context.
3. Check MCP server status/config and permission grants.
4. Inspect the relevant trace for permission, approval, or transport errors.
5. Apply the smallest grant or config change after confirmation.

### Package Install Or Update Failed

1. Check skill dependency status.
2. Inspect runtime availability for Python, Node, system packages, or GitHub
   release installers.
3. Look for pending approval, network, permission, disk, or checksum errors.
4. Retry only after the root cause is clear.

### CLI Credentials Missing Or Expired

1. Run `base365 auth status` or the relevant credential status command.
2. Verify required env vars exist without printing values.
3. Re-authenticate through the approved UI/CLI flow.
4. Re-run the original read-only command before mutating anything.

### Provider Or Model Error

1. Use traces to find the provider, model, status code, and error class.
2. Check `base365 providers list` and provider verification help.
3. Confirm model availability from the current provider config.
4. Avoid changing provider priority or credentials without user approval.

### Skill Not Visible Or Not Granted

1. Run `base365 skills list` and `base365 skills show <skill>`.
2. Inspect access mode and effective access for the target agent/user.
3. Check whether the skill is archived due to missing dependencies.
4. Grant or enable only the requested skill and scope.

### Channel Delivery Failure

1. Identify channel, session, sender, and trace ID.
2. Inspect `base365 channels list` and the failed trace.
3. Check provider/tool errors before blaming the channel.
4. Redact external message IDs and user data when reporting.

## Mutating Command Confirmation

Ask for confirmation before running commands that:

- delete, reset, revoke, restore, force, drop, or migrate down;
- change provider/channel credentials;
- install/update/remove runtime packages;
- grant broad skill/tool access;
- modify tenant/global config;
- trigger cron jobs with external side effects.

Confirmation should include the target, command category, expected effect, and
rollback path if one exists.
