#!/usr/bin/env python3
"""Motor de rebrand GoClaw -> Base365.

Uso: apply.py [--dry-run] [--content-only] [--rename-only] [--report ARQUIVO]
Aplica scripts/rebrand/rules.txt em ordem ao conteúdo de arquivos de texto e depois
renomeia caminhos cujo nome contém o termo antigo. Idempotente. Respeita
scripts/brand-exceptions.txt. Em --dry-run não grava nada e gera relatório por regra.
"""
import os, re, sys, subprocess, collections

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), "..", ".."))
os.chdir(ROOT)
args = sys.argv[1:]
DRY = "--dry-run" in args
CONTENT = "--rename-only" not in args
RENAME = "--content-only" not in args
REPORT = args[args.index("--report") + 1] if "--report" in args else None

def load_rules():
    rules = []
    for line in open("scripts/rebrand/rules.txt", encoding="utf-8"):
        if not line.strip() or line.startswith("#"):
            continue
        parts = line.rstrip("\n").split("\t")
        rules.append((re.compile(parts[0]), parts[1], parts[2] if len(parts) > 2 else ""))
    return rules

def load_exceptions():
    ex = []
    for line in open("scripts/brand-exceptions.txt", encoding="utf-8"):
        p = line.split("#")[0].strip()
        if p:
            ex.append(p)
    return ex

def excluded(path, ex):
    if "node_modules" in path.split("/"):
        return True
    for e in ex:
        if e.endswith("/"):
            if path == e[:-1] or path.startswith(e):
                return True
        elif path == e:
            return True
    return False

def all_files(ex):
    out = subprocess.run(["git", "ls-files", "-z", "--cached", "--others", "--exclude-standard"],
                         capture_output=True, check=True).stdout.decode("utf-8", "surrogateescape")
    for p in out.split("\0"):
        if p and os.path.isfile(p) and not os.path.islink(p) and not excluded(p, ex):
            yield p

def is_text(path):
    try:
        with open(path, "rb") as f:
            return b"\0" not in f.read(8192)
    except OSError:
        return False

TRIGGER = re.compile(r"goclaw|nextlevelbuilder|digitop", re.I)
rules = load_rules(); ex = load_exceptions()
counts = collections.Counter(); samples = collections.defaultdict(list)
cat_files = collections.Counter(); touched = 0

if CONTENT:
    for p in all_files(ex):
        if not is_text(p):
            continue
        try:
            s = open(p, encoding="utf-8", newline="").read()
        except UnicodeDecodeError:
            print("AVISO: não é UTF-8, ignorado:", p, file=sys.stderr); continue
        if not TRIGGER.search(s):
            continue
        orig = s
        for i, (rx, rep, why) in enumerate(rules):
            def sub(m, i=i, rep=rep):
                counts[i] += 1
                if len(samples[i]) < 3:
                    samples[i].append((p, m.group(0)))
                return m.expand(rep) if "\\" in rep else rep
            s = rx.sub(sub, s)
        if s != orig:
            touched += 1
            if not DRY:
                open(p, "w", encoding="utf-8", newline="").write(s)

renames = []
if RENAME:
    # do mais profundo para o mais raso; renomeia só o último componente
    paths = []
    for dp, dns, fns in os.walk("."):
        rel = os.path.relpath(dp, ".")
        if rel != "." and excluded(rel + "/", ex):
            dns[:] = []; continue
        dns[:] = [d for d in dns if not excluded(os.path.normpath(os.path.join(rel, d)) + "/", ex)]
        for n in dns + fns:
            if re.search("goclaw", n, re.I):
                paths.append(os.path.normpath(os.path.join(rel, n)))
    for p in sorted(paths, key=lambda x: -x.count(os.sep)):
        d, n = os.path.split(p)
        new = n
        for rx, rep, why in rules:
            if "/" not in rx.pattern:  # regras de URL/caminho não se aplicam a nomes de arquivo
                new = rx.sub(rep, new)
        if new != n:
            renames.append((p, os.path.join(d, new)))
            if not DRY:
                os.rename(p, os.path.join(d, new))

lines = ["# Relatório do motor de rebrand" + (" (dry-run)" if DRY else ""), "",
         f"Arquivos de conteúdo alterados: **{touched}**", f"Caminhos renomeados: **{len(renames)}**", "",
         "| # | Regra | Substituições | Exemplo |", "|---|-------|---------------|---------|"]
for i, (rx, rep, why) in enumerate(rules):
    ex_s = "; ".join(f"`{a}`: `{b}`" for a, b in samples[i][:2])
    lines.append(f"| {i+1} | `{rx.pattern}` → `{rep}` ({why}) | {counts[i]} | {ex_s} |")
lines += ["", "## Renomeios", ""] + [f"- `{a}` → `{b}`" for a, b in renames]
text = "\n".join(lines) + "\n"
if REPORT:
    open(REPORT, "w", encoding="utf-8").write(text)
print(text)
