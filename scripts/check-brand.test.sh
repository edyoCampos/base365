#!/usr/bin/env bash
# Teste da checagem de marca (RF-23): falha ao reintroduzir o nome antigo e passa depois de remover.
# O nome antigo é montado em partes para este arquivo não ser acusado pela própria checagem.
set -u
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OLD="Go""Claw"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

mkdir -p "$TMP/scripts" "$TMP/src"
cp "$ROOT/scripts/check-brand.sh" "$TMP/scripts/"
cat > "$TMP/scripts/brand-exceptions.txt" <<'EX'
scripts/brand-exceptions.txt
scripts/check-brand.sh
EX
echo "clean file" > "$TMP/src/ok.txt"

fail() { echo "FALHOU: $*"; exit 1; }

# 1. árvore limpa passa
"$TMP/scripts/check-brand.sh" "$TMP" >/dev/null 2>&1 || fail "árvore limpa deveria passar"

# 2. nome no conteúdo falha e indica arquivo e linha
printf 'line one\nsee %s docs\n' "$OLD" > "$TMP/src/bad.txt"
out="$("$TMP/scripts/check-brand.sh" "$TMP" 2>&1)" && fail "conteúdo com o nome antigo deveria falhar"
echo "$out" | grep -q "src/bad.txt:2" || fail "a saída deveria apontar src/bad.txt:2, veio: $out"
rm "$TMP/src/bad.txt"

# 3. nome no nome do arquivo falha
touch "$TMP/src/${OLD}-notes.txt"
"$TMP/scripts/check-brand.sh" "$TMP" >/dev/null 2>&1 && fail "nome de arquivo com o nome antigo deveria falhar"
rm "$TMP/src/${OLD}-notes.txt"

# 4. exceção declarada é respeitada
mkdir -p "$TMP/legacy"; echo "$OLD" > "$TMP/legacy/a.txt"
"$TMP/scripts/check-brand.sh" "$TMP" >/dev/null 2>&1 && fail "sem exceção deveria falhar"
echo "legacy/    # teste" >> "$TMP/scripts/brand-exceptions.txt"
"$TMP/scripts/check-brand.sh" "$TMP" >/dev/null 2>&1 || fail "com exceção deveria passar"

# 5. depois de remover tudo, passa
rm -rf "$TMP/legacy"
"$TMP/scripts/check-brand.sh" "$TMP" >/dev/null 2>&1 || fail "estado final deveria passar"

echo "OK: check-brand.test.sh"
