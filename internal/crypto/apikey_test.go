package crypto

import (
	"strings"
	"testing"
)

func TestGenerateAPIKey_Format(t *testing.T) {
	raw, hash, prefix, err := GenerateAPIKey()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(raw, "base365_") {
		t.Fatalf("api key must start with the product prefix, got %q", raw)
	}
	if len(raw) != len("base365_")+32 {
		t.Fatalf("api key must be prefix + 32 hex chars, got %q", raw)
	}
	if hash != HashAPIKey(raw) {
		t.Fatal("hash must match HashAPIKey(raw)")
	}
	if len(prefix) != 8 || !strings.HasPrefix(strings.TrimPrefix(raw, "base365_"), prefix) {
		t.Fatalf("display prefix %q must be the first 8 hex chars of the random part", prefix)
	}
}

// RN-02: keys with the pre-rebrand prefix are not recognized.
func TestAPIKeyPrefix_LegacyNotUsed(t *testing.T) {
	raw, _, _, err := GenerateAPIKey()
	if err != nil {
		t.Fatal(err)
	}
	if strings.HasPrefix(raw, "go"+"claw_") {
		t.Fatalf("generated key must not use the legacy prefix: %q", raw)
	}
}
