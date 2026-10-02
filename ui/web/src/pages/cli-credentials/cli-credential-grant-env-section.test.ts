import { describe, it, expect } from "vitest";
import { readFileSync } from "node:fs";
import { join } from "node:path";

// ENV_DENYLIST_PREFIXES is a module-private const mirrored from the backend denylist.
// Read the source so the mirror cannot drift silently from internal/crypto/env_denylist.go.
describe("grant env denylist mirror", () => {
  const src = readFileSync(join(__dirname, "cli-credential-grant-env-section.tsx"), "utf8");
  const match = src.match(/const ENV_DENYLIST_PREFIXES = \[([^\]]*)\]/);

  it("declares the product env prefix", () => {
    expect(match).not.toBeNull();
    expect(match![1]).toContain('"BASE365_"');
  });

  it("keeps the other denied prefixes", () => {
    for (const p of ["DYLD_", "LD_", "NPM_CONFIG_"]) expect(match![1]).toContain(`"${p}"`);
  });
});
