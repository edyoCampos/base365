#!/usr/bin/env bash
# Checagens de identidade das interfaces (RF-07, RF-16, RF-17, RF-20, RF-21). Sai com 1 se alguma falhar.
set -u
fail=0
bad() { echo "FALHA: $*"; fail=1; }

# RF-20: nenhuma classe orange-* nos componentes.
n=$(grep -rIl 'orange-' ui/web/src ui/desktop/frontend/src | wc -l)
[ "$n" -eq 0 ] || bad "RF-20: $n arquivo(s) ainda usam orange-*"

# RF-16: nomes dos pacotes.
for p in ui/web/package.json ui/desktop/frontend/package.json; do
  grep -q '"name": "base365-' "$p" || bad "RF-16: $p sem nome base365-*"
done

# RF-07: consumidores web e desktop usam o prefixo novo nos cabeçalhos.
grep -rq "X-Base365-" ui/web/src || bad "RF-07: ui/web/src sem cabeçalhos X-Base365-*"
grep -rq "X-Base365-" ui/desktop/frontend/src || bad "RF-07: ui/desktop/frontend/src sem cabeçalhos X-Base365-*"

# RF-21: favicons 16, 32 e 64 px e símbolo nas duas interfaces.
for s in 16 32 64; do
  f="ui/web/public/favicon-$s.png"
  [ -f "$f" ] || { bad "RF-21: falta $f"; continue; }
  file "$f" | grep -q "$s x $s" || bad "RF-21: $f não tem ${s}x${s}"
done
for f in ui/web/public/base365-icon.svg ui/desktop/frontend/public/base365-icon.svg ui/desktop/build/appicon.png; do
  [ -f "$f" ] || bad "RF-21: falta $f"
done

# RF-17: cada endereço provisório usado no repositório consta em docs/rebrand-pendencias.md.
for addr in 'edyocampos.github.io/base365' 'ghcr.io/edyocampos/base365' 'base365.example.com'; do
  grep -rIq --exclude-dir=node_modules --exclude-dir=.git --exclude-dir=dist --exclude-dir='_reversa_*' \
    --exclude=rebrand-pendencias.md "$addr" . || continue
  grep -q "$addr" docs/rebrand-pendencias.md || bad "RF-17: $addr em uso, mas ausente de docs/rebrand-pendencias.md"
done
grep -q 'edyocampos/base365' docs/rebrand-pendencias.md || bad "RF-17: Docker Hub edyocampos/base365 ausente da lista"

[ "$fail" -eq 0 ] && echo "OK: identidade das interfaces conferida (RF-07, 16, 17, 20, 21)"
exit "$fail"
