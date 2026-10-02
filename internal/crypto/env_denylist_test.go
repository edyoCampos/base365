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
