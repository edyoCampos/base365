#!/usr/bin/env node
// Compara os tokens de cor aplicados na web e no desktop com os valores do documento de marca
// (scripts/brand-tokens.json, extraído de "Base365 logo e ícone.pdf", V1). Sai com 1 se algum divergir (RF-19).
import { readFileSync } from "node:fs";

const spec = JSON.parse(readFileSync("scripts/brand-tokens.json", "utf8"));
const web = readFileSync("ui/web/src/index.css", "utf8");
const desk = readFileSync("ui/desktop/frontend/src/index.css", "utf8");

const block = (css, start) => {
  const i = css.indexOf(start);
  if (i < 0) throw new Error(`bloco não encontrado: ${start}`);
  return css.slice(i, css.indexOf("\n}", i));
};
const problems = [];
const check = (where, name, got, want) => {
  if (got?.toLowerCase() !== want) problems.push(`${where} ${name}: obtido ${got ?? "ausente"}, esperado ${want}`);
};

for (const [theme, start] of [["light", ":root {"], ["dark", ".dark {"]]) {
  const b = block(web, start);
  for (const [name, hex] of Object.entries(spec[theme])) {
    const m = b.match(new RegExp(`--${name}:\\s*[^;]+;\\s*/\\*\\s*(#[0-9a-fA-F]{6})\\s*\\*/`));
    check(`web ${theme}`, name, m?.[1], hex);
  }
}
for (const [theme, start, i] of [["light", ":root:not(.dark) {", 0], ["dark", ".dark {", 1]]) {
  const b = block(desk, start);
  for (const [name, pair] of Object.entries(spec.desktop)) {
    const m = b.match(new RegExp(`--color-${name}:\\s*(#[0-9a-fA-F]{6});`));
    check(`desktop ${theme}`, name, m?.[1], pair[i]);
  }
}
const total = Object.keys(spec.light).length * 2 + Object.keys(spec.desktop).length * 2;
if (problems.length) {
  console.error(problems.join("\n"));
  console.error(`FALHA: ${problems.length} de ${total} tokens divergem do documento de marca`);
  process.exit(1);
}
console.log(`OK: ${total} tokens (web e desktop, claro e escuro) conferem com o documento de marca`);
