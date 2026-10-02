#!/usr/bin/env bash
# Checagem de marca: falha se o nome antigo aparecer no conteúdo ou nos nomes de arquivo.
# Uso: scripts/check-brand.sh [raiz]     (padrão: raiz do repositório)
# Exceções: scripts/brand-exceptions.txt
set -u
ROOT="${1:-$(cd "$(dirname "$0")/.." && pwd)}"
cd "$ROOT" || exit 2
PATTERN='goclaw'
EXC="scripts/brand-exceptions.txt"

dir_args=(); file_args=(); find_args=()
if [[ -f "$EXC" ]]; then
  while IFS= read -r line; do
    p="${line%%#*}"; p="$(echo "$p" | sed 's/[[:space:]]*$//; s/^[[:space:]]*//')"
    [[ -z "$p" ]] && continue
    if [[ "$p" == */ ]]; then
      p="${p%/}"; dir_args+=(--exclude-dir="$(basename "$p")"); find_args+=(-not -path "./$p/*" -not -path "./$p")
    else
      file_args+=(--exclude="$(basename "$p")"); find_args+=(-not -path "./$p")
    fi
  done < "$EXC"
fi

fail=0
echo "== conteúdo"
if grep -rIniE "${dir_args[@]}" "${file_args[@]}" "$PATTERN" . 2>/dev/null | head -n "${BRAND_MAX_LINES:-200}"; then :; fi
content_hits=$(grep -rIiE "${dir_args[@]}" "${file_args[@]}" "$PATTERN" . 2>/dev/null | wc -l)
echo "ocorrências no conteúdo: $content_hits"
[[ "$content_hits" -gt 0 ]] && fail=1

echo "== nomes"
name_hits=$(find . "${find_args[@]}" -iname "*${PATTERN}*" 2>/dev/null | tee /dev/stderr | wc -l)
echo "caminhos com o nome antigo: $name_hits"
[[ "$name_hits" -gt 0 ]] && fail=1

# Lista de aviso (não falha): marcas da infraestrutura herdada.
WARN='digitop|tamgiac|zuey|nextlevelbuilder'
warn_hits=$(grep -rIiE "${dir_args[@]}" "${file_args[@]}" "$WARN" . 2>/dev/null | wc -l)
echo "aviso (não falha): $warn_hits linhas com marcas herdadas ($WARN)"

if [[ $fail -ne 0 ]]; then echo "FALHOU: marca antiga encontrada"; exit 1; fi
echo "OK: nenhuma ocorrência fora das exceções"
