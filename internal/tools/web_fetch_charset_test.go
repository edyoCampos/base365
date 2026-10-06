package tools

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const charsetFixtureText = "Informação técnica: a ação não é só para o coração. Pão, maçã e açúcar."

func readFixture(t *testing.T, parts ...string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(append([]string{"testdata", "web_fetch"}, parts...)...))
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}
	return data
}

func TestDecodeBody_Fixtures(t *testing.T) {
	tests := []struct {
		file        string
		contentType string
		wantCharset string
		wantExtra   string
	}{
		{"latin1_header.html", "text/html; charset=ISO-8859-1", "windows-1252", ""},
		{"latin1_meta.html", "text/html", "windows-1252", ""},
		{"cp1252_undeclared.html", "text/html", "windows-1252", "Ele disse “ola” – e saiu."},
		{"utf8_undeclared_ascii1k.html", "text/html", "utf-8", ""},
		// BOM wins over a wrong header.
		{"utf8_bom.html", "text/html; charset=iso-8859-1", "utf-8", ""},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			text, name := decodeBody(readFixture(t, "charset", tt.file), tt.contentType)
			if name != tt.wantCharset {
				t.Errorf("charset = %q, want %q", name, tt.wantCharset)
			}
			if !strings.Contains(text, charsetFixtureText) {
				t.Errorf("decoded text missing accented sentence: %q", text)
			}
			if tt.wantExtra != "" && !strings.Contains(text, tt.wantExtra) {
				t.Errorf("decoded text missing %q: %q", tt.wantExtra, text)
			}
			if strings.ContainsRune(text, '�') || strings.Contains(text, "Ã") {
				t.Errorf("decoded text has mojibake: %q", text)
			}
		})
	}
}

func TestDecodeBody_UnknownHeaderCharsetFallsThrough(t *testing.T) {
	text, name := decodeBody(readFixture(t, "charset", "latin1_meta.html"), "text/html; charset=x-unknown-thing")
	if name != "windows-1252" || !strings.Contains(text, charsetFixtureText) {
		t.Errorf("got charset %q, text %q", name, text)
	}
}

func TestDecodeBody_Empty(t *testing.T) {
	if text, _ := decodeBody(nil, "text/html"); text != "" {
		t.Errorf("expected empty text, got %q", text)
	}
}

func TestDecodeBody_UTF8CutMidRune(t *testing.T) {
	body := []byte("<p>" + charsetFixtureText + "</p> ação")
	// End on the first byte of the last "ã", as a read limit would.
	cut := body[:strings.LastIndex(string(body), "ã")+1]
	text, name := decodeBody(cut, "text/html")
	if name != "utf-8" {
		t.Fatalf("charset = %q, want utf-8 (a rune split by the read limit must not demote UTF-8)", name)
	}
	if !strings.Contains(text, charsetFixtureText) || strings.ContainsRune(text, '�') {
		t.Errorf("unexpected text: %q", text)
	}
}

func TestCutAtRuneBoundary(t *testing.T) {
	s := "aã" // 'a' + 2-byte rune
	if got := cutAtRuneBoundary(s, 2); got != "a" {
		t.Errorf("cutAtRuneBoundary = %q, want %q", got, "a")
	}
	if got := truncateStr("ããã", 3); got != "ã..." {
		t.Errorf("truncateStr = %q, want %q", got, "ã...")
	}
}
