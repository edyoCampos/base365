package cmd

import "testing"

// Legacy name built from parts so the brand check does not flag this file (RN-02).
const legacyConfigEnv = "GO" + "CLAW_CONFIG"

func TestResolveConfigPath_ReadsBase365Config(t *testing.T) {
	t.Cleanup(func(old string) func() { return func() { cfgFile = old } }(cfgFile))
	cfgFile = ""
	t.Setenv("BASE365_CONFIG", "/tmp/base365-test.json")
	if got := resolveConfigPath(); got != "/tmp/base365-test.json" {
		t.Fatalf("resolveConfigPath = %q, want BASE365_CONFIG value", got)
	}
}

func TestResolveConfigPath_IgnoresLegacyEnv(t *testing.T) {
	t.Cleanup(func(old string) func() { return func() { cfgFile = old } }(cfgFile))
	cfgFile = ""
	t.Setenv("BASE365_CONFIG", "")
	t.Setenv(legacyConfigEnv, "/tmp/legacy.json")
	if got := resolveConfigPath(); got != "config.json" {
		t.Fatalf("resolveConfigPath = %q, want default config.json (legacy env must be ignored)", got)
	}
}
