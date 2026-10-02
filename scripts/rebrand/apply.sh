#!/usr/bin/env bash
# Wrapper do motor de rebrand. Ver apply.py.
exec python3 "$(dirname "$0")/apply.py" "$@"
