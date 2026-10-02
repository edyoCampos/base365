#!/usr/bin/env node
// Verifica o contraste WCAG 2.1 dos pares texto/fundo aplicados nos tokens de cor (RNF de acessibilidade).
// Uso: node scripts/rebrand/contrast-check.mjs   (sai com código 1 se algum par falhar)
import { readFileSync } from "node:fs";

const lum = (hex) => {
  const [r, g, b] = [1, 3, 5].map((i) => parseInt(hex.slice(i, i + 2), 16) / 255)
    .map((c) => (c <= 0.03928 ? c / 12.92 : ((c + 0.055) / 1.055) ** 2.4));
  return 0.2126 * r + 0.7152 * g + 0.0722 * b;
};
const ratio = (a, b) => {
  const [x, y] = [lum(a), lum(b)].sort((p, q) => q - p);
  return (x + 0.05) / (y + 0.05);
};

// Lê um bloco CSS e devolve {token: "#hex"}; o hex vem do comentário (web) ou do valor (desktop).
function tokens(css, startMarker, prefix) {
  const i = css.indexOf(startMarker);
  const j = css.indexOf("\n}", i);
  const out = {};
  for (const m of css.slice(i, j).matchAll(new RegExp(`--${prefix}([a-z0-9-]+):\\s*([^;]+);(?:\\s*/\\*\\s*(#[0-9a-fA-F]{6})\\s*\\*/)?`, "g"))) {
    const hex = m[3] ?? (/^#[0-9a-fA-F]{6}$/.test(m[2].trim()) ? m[2].trim() : null);
    if (hex) out[m[1]] = hex.toLowerCase();
  }
  return out;
}

const web = readFileSync("ui/web/src/index.css", "utf8");
const desk = readFileSync("ui/desktop/frontend/src/index.css", "utf8");

const webText = [["foreground", "background"], ["card-foreground", "card"], ["popover-foreground", "popover"],
  ["primary-foreground", "primary"], ["secondary-foreground", "secondary"], ["muted-foreground", "muted"],
  ["muted-foreground", "background"], ["muted-foreground", "card"], ["accent-foreground", "accent"],
  ["destructive-foreground", "destructive"], ["success-foreground", "success"], ["warning-foreground", "warning"],
  ["sidebar-foreground", "sidebar"], ["sidebar-primary-foreground", "sidebar-primary"],
  ["sidebar-accent-foreground", "sidebar-accent"], ["foreground", "chat-bubble-user"],
  ["primary", "card"], ["primary", "background"], ["badge-success-foreground", "badge-success"],
  ["badge-warning-foreground", "badge-warning"], ["badge-info-foreground", "badge-info"]];
const webNonText = [["ring", "background"], ["ring", "card"], ["input", "card"], ["chart-1", "card"], ["chart-2", "card"],
  ["chart-3", "card"], ["chart-4", "card"], ["chart-5", "card"]];
const deskText = [["text-primary", "surface-primary"], ["text-primary", "surface-secondary"], ["text-primary", "surface-tertiary"],
  ["text-secondary", "surface-primary"], ["text-secondary", "surface-secondary"], ["text-muted", "surface-primary"],
  ["text-muted", "surface-secondary"], ["accent", "surface-primary"], ["accent", "surface-secondary"],
  ["accent-foreground", "accent"], ["text-primary", "user-bubble"], ["text-primary", "agent-bubble"],
  ["success", "surface-secondary"], ["warning", "surface-secondary"], ["error", "surface-secondary"]];

let fail = 0, total = 0;
function run(label, toks, pairs, min) {
  for (const [fg, bg] of pairs) {
    if (!toks[fg] || !toks[bg]) { console.log(`?  ${label} ${fg} × ${bg}: token ausente`); fail++; continue; }
    const r = ratio(toks[fg], toks[bg]); total++;
    const ok = r >= min;
    if (!ok) fail++;
    console.log(`${ok ? "ok" : "FALHA"}  ${label.padEnd(14)} ${fg} × ${bg}`.padEnd(64) + `${r.toFixed(2)}:1 (mín ${min})`);
  }
}
const wl = tokens(web, ":root {", ""), wd = tokens(web, ".dark {", "");
run("web claro", wl, webText, 4.5); run("web claro", wl, webNonText, 3);
run("web escuro", wd, webText, 4.5); run("web escuro", wd, webNonText, 3);
const dd = tokens(desk, "@theme {", "color-"), dl = tokens(desk, ":root:not(.dark) {", "color-");
run("desktop claro", dl, deskText, 4.5); run("desktop escuro", dd, deskText, 4.5);
console.log(`\n${total} pares verificados, ${fail} falha(s)`);
process.exit(fail ? 1 : 0);
