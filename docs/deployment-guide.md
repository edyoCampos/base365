# Deployment Guide

## Overview

Production target uses a hybrid deployment:

- Base365 gateway runs as a bare-metal `systemd` service.
- PostgreSQL 18 + pgvector runs in Docker.
- Nginx reverse proxies public HTTP/HTTPS traffic to Base365 on localhost.
- Codex CLI is installed on the host for future agent-controlled CLI work.

Current VPS shape:

| Item | Value |
|---|---|
| Host | Private deployment value |
| SSH | Private deployment value |
| Domain | Private deployment value |
| Gateway listen | `127.0.0.1:18790` |
| Public proxy | Nginx on `80/443` |
| Postgres | Docker container `base365-postgres` |
| Postgres bind | `127.0.0.1:5432` |

Keep concrete hostnames, IPs, SSH users, SSH ports, and credentials in private operator notes or environment variables outside this public repository.

For command examples below:

```bash
export BASE365_HOST=<server-ip-or-hostname>
export BASE365_SSH_USER=<ssh-user>
export BASE365_SSH_PORT=<ssh-port>
export BASE365_DOMAIN=<public-domain>
```

## Server Layout

| Path | Purpose |
|---|---|
| `/opt/base365/releases/<release>` | Immutable Base365 release directory |
| `/opt/base365/current` | Symlink to active release |
| `/opt/base365/shared/docker-compose.postgres.yml` | Postgres compose file |
| `/opt/base365/shared/postgres.env` | Postgres env file |
| `/opt/base365/backups/` | Uploaded DB backup files |
| `/etc/base365/config.json` | Gateway config |
| `/etc/base365/base365.env` | Gateway env and secrets |
| `/etc/base365/r2-backup.env` | R2 backup env and secrets |
| `/var/lib/base365/data` | Base365 persistent data |
| `/var/lib/base365/workspace` | Agent workspace |
| `/var/lib/base365/postgres` | Postgres Docker data |
| `/usr/local/bin/base365-deploy` | Release switch, upgrade, health-check, rollback (source: [`scripts/zuey/base365-deploy.sh`](../scripts/zuey/base365-deploy.sh)) |
| `/usr/local/bin/base365-issue-ssl` | Certbot wrapper for the deployment domain |
| `/usr/local/bin/base365-backup-r2` | Postgres dump, R2 upload, retention cleanup |
| `/usr/local/bin/base365-upgrade-release` | Download and deploy a GitHub Release tarball (source: [`scripts/zuey/base365-upgrade-release.sh`](../scripts/zuey/base365-upgrade-release.sh)) |

Secrets are stored only in server env files. Do not copy tokens or database passwords into repo docs.

## Runtime Services

Check status:

```bash
ssh -p "$BASE365_SSH_PORT" "$BASE365_SSH_USER@$BASE365_HOST"
systemctl status base365 --no-pager
systemctl status nginx --no-pager
systemctl status base365-backup-r2.timer --no-pager
docker ps --filter name=base365-postgres
curl -fsS http://127.0.0.1:18790/health
```

Expected health:

```json
{"status":"ok","protocol":3}
```

## Initial Deployment Record

Completed on 2026-05-17:

- Installed Docker, Docker Compose v2, Nginx, Certbot, Node.js 22, Codex CLI.
- Added the operator workstation SSH public key to the deployment user.
- Installed Codex CLI. For agent-controlled `codex exec`, authenticate the Linux service user that runs Base365 (`base365`), not only the SSH operator user.
- Restored the latest private PostgreSQL backup into Docker Postgres.
- Upgraded restored schema from `57` to `65`.
- Started the initial Base365 release by `systemd`.
- Added automated database backup to a private Cloudflare R2 bucket.
- Verified local and Nginx health endpoints.

Current verification snapshot:

```text
base365_active=active
nginx_active=active
health_local={"status":"ok","protocol":3}
health_nginx={"status":"ok","protocol":3}
docker=base365-postgres pgvector/pgvector:pg18 healthy
schema=65
codex=codex-cli 0.130.0
```

## Codex CLI Auth For Agents

The gateway runs under the `base365` Linux user. Codex auth is stored under the
effective home directory, so an SSH-session login such as `codex login
--device-auth` under an operator user only writes that user's `~/.codex/auth.json`.
Agents invoking `codex` through the exec tool use the service user's home:

```text
/var/lib/base365/.codex/auth.json
```

Verify both contexts when debugging auth:

```bash
codex login status
sudo -u base365 -H codex login status
```

If the operator user is logged in but `base365` is not, either run device auth as
the service user or copy the operator auth file with strict ownership:

```bash
sudo install -d -o base365 -g base365 -m 700 /var/lib/base365/.codex
sudo install -o base365 -g base365 -m 600 ~/.codex/auth.json /var/lib/base365/.codex/auth.json
sudo -u base365 -H codex login status
sudo -u base365 -H sh -lc 'mkdir -p /var/lib/base365/codex-smoke && cd /var/lib/base365/codex-smoke && codex exec --skip-git-repo-check --sandbox read-only "Reply with exactly: CODEX_AUTH_OK"'
```

Do not store Codex auth material in repository docs. Treat `auth.json` as a
credential.

## DNS And SSL

Cloudflare DNS record:

| Type | Name | Value | Proxy |
|---|---|---|---|
| `A` | Deployment subdomain | Deployment host IP | Proxied |

SSL was issued with Certbot for the deployment domain on 2026-05-17 and Certbot installed automatic renewal.

Verify HTTPS:

```bash
ssh -p "$BASE365_SSH_PORT" "$BASE365_SSH_USER@$BASE365_HOST"
curl -fsS "https://$BASE365_DOMAIN/health"
```

Re-issue manually if needed:

```bash
sudo /usr/local/bin/base365-issue-ssl
```

## Deploy A New Release

Preferred server-side upgrade flow:

```bash
sudo /usr/local/bin/base365-upgrade-release --dry-run latest
sudo /usr/local/bin/base365-upgrade-release latest
sudo /usr/local/bin/base365-upgrade-release v3.12.0
```

The script downloads the Linux amd64 GitHub Release tarball from `edyoCampos/base365`, follows GitHub release redirects, verifies `CHECKSUMS.sha256` when present, falls back to the GitHub release asset SHA256 digest for beta assets without checksum files, extracts to `/opt/base365/releases/<tag>`, and calls `base365-deploy`. When invoked from the running gateway service, it first re-launches itself as a transient `systemd-run` unit so `base365-deploy` can stop/restart `base365` without killing the upgrade job.

The HTTP API still accepts only `tag`; it does not accept repo names or custom download URLs.

Remote API trigger is available in builds that include the gateway upgrade endpoint:

```bash
curl -fsS -X POST "https://$BASE365_DOMAIN/v1/system/gateway/upgrade" \
  -H "Authorization: Bearer <gateway-token>" \
  -H "X-Base365-Upgrade-Token: <upgrade-token>" \
  -H "Content-Type: application/json" \
  --data '{"tag":"latest"}'
```

Check status:

```bash
curl -fsS "https://$BASE365_DOMAIN/v1/system/gateway/upgrade/status" \
  -H "Authorization: Bearer <gateway-token>" \
  -H "X-Base365-Upgrade-Token: <upgrade-token>"
```

Keep upgrade tokens in server env files or secret managers. Do not put real tokens in docs.

The remote trigger endpoint fails closed unless `BASE365_UPGRADE_TRIGGER_TOKEN` is configured in the gateway environment.

### Automatic Beta Deploy From `dev`

Pushing or merging into `dev` runs `.github/workflows/dev-beta-release.yaml`. After Go/Web checks pass, the workflow creates the next semantic beta tag, publishes the linux amd64 prerelease asset and checksum, then deploys that exact beta tag to the zuey VPS through the gateway upgrade endpoint.

Linux arm64 release assets, full checksums, multi-arch Docker images, and beta Docker aliases continue after the fast zuey deploy path. Release asset completion waits until zuey deploy finishes before using `--clobber`, so the VPS upgrade script cannot race with a release asset refresh while downloading the linux amd64 tarball. Those completion jobs are still required to pass; the workflow remains failed if later artifact completion breaks. Zuey deploy and Docker beta alias promotion both skip stale beta tags so older runs cannot roll back a newer beta.

Required GitHub Actions configuration:

| Name | Type | Value |
|---|---|---|
| `ZUEY_BASE365_URL` | Secret | Public gateway URL, for example `https://base365.example.com` |
| `ZUEY_BASE365_GATEWAY_TOKEN` | Secret | Gateway bearer token from the server env |
| `ZUEY_BASE365_UPGRADE_TOKEN` | Secret | Upgrade trigger token from the server env |
| `ZUEY_BASE365_USER_ID` | Variable | Optional owner identity, defaults to `system` |

The deploy job sends:

```bash
POST /v1/system/gateway/upgrade {"tag":"vX.Y.Z-beta.N"}
GET /v1/system/gateway/upgrade/status
GET /health
```

The workflow fails if the upgrade status becomes `failed`, times out, or public health does not return `{"status":"ok"}`.

Manual local-build fallback:

Build locally with embedded web UI:

```bash
cd ui/web
pnpm install --frozen-lockfile
pnpm build
cd ../..
rm -rf internal/webui/dist
mkdir -p internal/webui/dist
cp -r ui/web/dist/* internal/webui/dist/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -tags embedui \
  -ldflags="-s -w -X github.com/edyoCampos/base365/cmd.Version=<version>" \
  -o dist/base365-linux-amd64 .
```

Upload:

```bash
release=<version-or-date>
ssh -p "$BASE365_SSH_PORT" "$BASE365_SSH_USER@$BASE365_HOST" "mkdir -p /opt/base365/releases/$release/migrations"
scp -P "$BASE365_SSH_PORT" dist/base365-linux-amd64 "$BASE365_SSH_USER@$BASE365_HOST:/opt/base365/releases/$release/base365"
scp -P "$BASE365_SSH_PORT" migrations/* "$BASE365_SSH_USER@$BASE365_HOST:/opt/base365/releases/$release/migrations/"
ssh -p "$BASE365_SSH_PORT" "$BASE365_SSH_USER@$BASE365_HOST" "chmod +x /opt/base365/releases/$release/base365 && sudo /usr/local/bin/base365-deploy /opt/base365/releases/$release"
```

`base365-deploy` does:

1. Validate binary and migrations.
2. Switch `/opt/base365/current`.
3. Run `base365 upgrade`.
4. Restart `base365`.
5. Poll `/health`.
6. Roll back symlink and restart if health fails.

### Self-loop symlink guard

`base365-deploy` captures the previous release for rollback via `readlink -f /opt/base365/current` before swinging the symlink. If `/opt/base365/current` is ever a self-loop (e.g. `current -> current`) or otherwise unresolvable, `readlink -f` exits non-zero. Combined with `set -euo pipefail`, an unguarded call aborts the script before `ln -sfn` runs, so deploys fail silently and zero rollback target is recorded — observed in production on 2026-05-27 where every `deploy_zuey_beta` CI run failed at this step.

The script swallows the readlink failure (`readlink -f ... 2>/dev/null || true`) and logs a warning when the symlink exists but resolves to empty, then proceeds to overwrite. Rollback is skipped in that case because no valid `previous` is available.

If you ever edit `/usr/local/bin/base365-deploy` on zuey, preserve this guard:

```bash
previous=""
if [ -L /opt/base365/current ]; then previous="$(readlink -f /opt/base365/current 2>/dev/null || true)"; fi
if [ -z "$previous" ] && [ -L /opt/base365/current ]; then
  echo "warn: /opt/base365/current is a symlink but readlink -f failed (likely self-loop or broken target); overwriting without rollback target" >&2
fi
```

The canonical source of this script lives in the repo at [`scripts/zuey/base365-deploy.sh`](../scripts/zuey/base365-deploy.sh), alongside [`scripts/zuey/base365-upgrade-release.sh`](../scripts/zuey/base365-upgrade-release.sh). The VPS copies at `/usr/local/bin/base365-deploy` and `/usr/local/bin/base365-upgrade-release` are downstream replicas. CI auto-syncs both on every beta release via the `Sync zuey ops scripts to VPS` step in `.github/workflows/dev-beta-release.yaml`. For manual sync (off-CI):

```bash
scp -P <SSH_PORT> scripts/zuey/base365-deploy.sh scripts/zuey/base365-upgrade-release.sh \
  zuey@<VPS_HOST>:/tmp/
ssh -p <SSH_PORT> zuey@<VPS_HOST> 'bash -s' <<'EOF'
set -euo pipefail
ts=$(date +%Y%m%d-%H%M%S)
for name in base365-deploy base365-upgrade-release; do
  if [ -f "/usr/local/bin/$name" ]; then
    sudo cp -p "/usr/local/bin/$name" "/usr/local/bin/${name}.bak-${ts}"
  fi
  sudo install -o root -g root -m 0755 "/tmp/${name}.sh" "/usr/local/bin/$name"
  sudo bash -n "/usr/local/bin/$name" && echo "OK: $name"
done
EOF
```

Always back up the live copy before overwriting (the `cp -p` line above does this).

### CI auto-sync — required GitHub secrets

The `Sync zuey ops scripts to VPS` step in `dev-beta-release.yaml` runs `scp + sudo install` before triggering the gateway upgrade endpoint. It needs the following repository secrets configured under **Settings → Secrets and variables → Actions**:

| Secret | Purpose |
|---|---|
| `ZUEY_SSH_PRIVATE_KEY_B64` | CI-only ed25519/rsa key, **base64-encoded as a single line** (`base64 -w0 < /path/to/key`). Its public key must be appended to `zuey@<VPS_HOST>:~/.ssh/authorized_keys`. **Do not reuse the operator's personal key.** Base64 avoids the `error in libcrypto` failure caused by GitHub Secrets normalizing newlines inside multi-line PEM blocks. |
| `ZUEY_SUDO_PASS` | Same value as `ZUEY_BASE365_SUDO_PASS` in the operator's local `.env`; used by `sudo -S` over the SSH session to install scripts. |

Optional repository **variables** (override defaults if the VPS endpoint changes):

| Variable | Default |
|---|---|
| `ZUEY_SSH_HOST` | `<VPS_HOST>` |
| `ZUEY_SSH_PORT` | `<SSH_PORT>` |
| `ZUEY_SSH_USER` | `zuey` |

To rotate the CI SSH key:

```bash
ssh-keygen -t ed25519 -f /tmp/ci-zuey-key -N '' -C 'gh-actions-deploy-zuey-beta'
# add /tmp/ci-zuey-key.pub to zuey:~/.ssh/authorized_keys (consider restricting to scp+install via `command="..."` forced-command)
# base64-encode the private key on a single line, then paste into the
# ZUEY_SSH_PRIVATE_KEY_B64 secret:
base64 -w0 < /tmp/ci-zuey-key  # macOS: `base64 -i /tmp/ci-zuey-key | tr -d '\n'`
# (copy the single-line output and paste it into the secret value)
# then `shred -u /tmp/ci-zuey-key*` locally
```

## Backup And Restore

Automated backups:

| Item | Value |
|---|---|
| Timer | `base365-backup-r2.timer` |
| Schedule | Every 6 hours: `00:00`, `06:00`, `12:00`, `18:00` server time |
| Source | Docker Postgres container `base365-postgres` |
| Format | PostgreSQL custom dump, `pg_dump -Fc` |
| Local directory | `/opt/base365/backups/` |
| R2 bucket | Private deployment value |
| R2 prefix | Private deployment value |
| Retention | Keep latest 20 backups locally and in R2 |

Check timer and latest logs:

```bash
systemctl list-timers base365-backup-r2.timer --no-pager
journalctl -u base365-backup-r2.service -n 80 --no-pager
```

Run a manual backup:

```bash
sudo systemctl start base365-backup-r2.service
```

Create a database dump on server:

```bash
ts=$(date +%Y%m%d-%H%M%S)
docker exec base365-postgres pg_dump -U base365 -Fc -d base365 > /opt/base365/backups/base365-$ts.dump
```

Restore a dump:

```bash
systemctl stop base365
docker exec -i base365-postgres pg_restore -U base365 -d base365 --clean --if-exists --no-owner < /opt/base365/backups/<file>.dump
sudo /usr/local/bin/base365-deploy /opt/base365/current
```

## Operational Notes

- Gateway runs as Linux user `base365`.
- Host-control exceptions are deployment-specific and must be documented privately, not in this public runbook.
- Workspace restriction settings are deployment-specific and must be reviewed before enabling agent-controlled host operations.
- Postgres is bound only to localhost.
- UFW allows the private SSH port, `80/tcp`, and `443/tcp`.
- Reboot is recommended later because the VPS reports a pending kernel upgrade.

## Troubleshooting

Logs:

```bash
journalctl -u base365 -n 200 --no-pager
docker logs base365-postgres --tail 100
tail -n 100 /var/log/nginx/error.log
```

Restart:

```bash
systemctl restart base365
docker compose --env-file /opt/base365/shared/postgres.env -f /opt/base365/shared/docker-compose.postgres.yml restart postgres
systemctl reload nginx
```

Unresolved questions: none.
