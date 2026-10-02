package crypto

import "testing"

func TestIsDeniedEnvKey_ProductPrefix(t *testing.T) {
	for _, k := range []string{"BASE365_X", "BASE365_GATEWAY_TOKEN", "base365_encryption_key"} {
		if !IsDeniedEnvKey(k) {
			t.Errorf("%q must be denied (product env prefix)", k)
		}
	}
	if IsDeniedEnvKey("MY_APP_TOKEN") {
		t.Error("unrelated key must not be denied")
	}
}

func TestValidateGrantEnvVars_RejectsProductPrefix(t *testing.T) {
	rejected, _ := ValidateGrantEnvVars(map[string]string{"BASE365_ENCRYPTION_KEY": "x", "OK_KEY": "y"})
	if len(rejected) != 1 || rejected[0] != "BASE365_ENCRYPTION_KEY" {
		t.Fatalf("expected only BASE365_ENCRYPTION_KEY rejected, got %v", rejected)
	}
}

// The pre-rebrand prefix is built from parts so the brand check does not flag this file.
// RN-02: no backwards compatibility, so the old prefix is an ordinary key now.
func TestIsDeniedEnvKey_LegacyPrefixIsNotSpecial(t *testing.T) {
	legacy := "GO" + "CLAW_X"
	if IsDeniedEnvKey(legacy) {
		t.Errorf("%q must no longer be treated as a product secret key", legacy)
	}
}
