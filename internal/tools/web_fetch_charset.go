package tools

import (
	"mime"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

// decodeBody converts a response body to UTF-8 text and returns the name of the charset used.
//
// Order: byte order mark, charset parameter of the Content-Type header, <meta>/sniff of the
// first 1 KB (trusted only when it names something other than the windows-1252 default; a
// utf-8 answer is trusted only when the whole body is valid UTF-8), UTF-8 validity of the whole
// body, windows-1252. A declared Latin-1 <meta> resolves through the last two steps, which give
// the same text. Never fails: undecodable bytes become U+FFFD.
func decodeBody(body []byte, contentType string) (string, string) {
	if len(body) == 0 {
		return "", "utf-8"
	}

	// 1. BOM. With an empty content type, DetermineEncoding is certain only for a BOM.
	if enc, name, certain := charset.DetermineEncoding(body, ""); certain {
		return decodeWith(enc, name, body)
	}

	// 2. Content-Type charset.
	if _, params, err := mime.ParseMediaType(contentType); err == nil {
		if label := params["charset"]; label != "" {
			if enc, name := charset.Lookup(label); enc != nil {
				return decodeWith(enc, name, body)
			}
		}
	}

	// 3. <meta> declaration or UTF-8 sniff of the first 1 KB.
	if enc, name, _ := charset.DetermineEncoding(body, ""); name != "windows-1252" {
		if name != "utf-8" {
			return decodeWith(enc, name, body)
		}
		if s, ok := validUTF8Prefix(body); ok {
			return s, "utf-8"
		}
		return decodeWith(charmap.Windows1252, "windows-1252", body)
	}

	// 4. Whole body is UTF-8.
	if s, ok := validUTF8Prefix(body); ok {
		return s, "utf-8"
	}

	// 5. Legacy default.
	return decodeWith(charmap.Windows1252, "windows-1252", body)
}

// decodeWith decodes body with enc, replacing invalid sequences instead of failing.
func decodeWith(enc encoding.Encoding, name string, body []byte) (string, string) {
	if name == "utf-8" {
		if s, ok := validUTF8Prefix(body); ok {
			return s, name
		}
		return strings.ToValidUTF8(string(body), "�"), name
	}
	out, err := enc.NewDecoder().Bytes(body)
	if err != nil {
		return strings.ToValidUTF8(string(body), "�"), name
	}
	return string(out), name
}

// validUTF8Prefix reports whether b is valid UTF-8, tolerating one incomplete rune at the end
// (left by a read limit cutting the body), which is dropped from the returned text.
func validUTF8Prefix(b []byte) (string, bool) {
	if utf8.Valid(b) {
		return string(b), true
	}
	for k := 1; k <= 3 && k < len(b); k++ {
		head, tail := b[:len(b)-k], b[len(b)-k:]
		if utf8.RuneStart(tail[0]) && !utf8.FullRune(tail) && utf8.Valid(head) {
			return string(head), true
		}
	}
	return "", false
}
