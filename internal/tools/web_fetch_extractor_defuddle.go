package tools

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	defuddleTimeout = 10 * time.Second
	defuddleMaxBody = 1 << 20 // 1MB
)

// DefuddleExtractor calls an operator-provided extraction service (for example a Cloudflare
// Worker running Defuddle) to extract clean markdown. There is no built-in endpoint: the
// extractor only runs when base_url is configured (see parseExtractorChainSettings).
type DefuddleExtractor struct {
	baseURL string
	client  *http.Client
}

// NewDefuddleExtractorFromEntry creates a DefuddleExtractor from chain settings.
func NewDefuddleExtractorFromEntry(entry ExtractorEntry) *DefuddleExtractor {
	baseURL := entry.BaseURL
	// Ensure trailing slash for URL construction.
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	timeout := time.Duration(entry.Timeout) * time.Second
	if timeout <= 0 {
		timeout = defuddleTimeout
	}
	return newDefuddleExtractor(baseURL, timeout)
}

// newDefuddleExtractor creates a DefuddleExtractor with custom base URL and timeout (for testing).
func newDefuddleExtractor(baseURL string, timeout time.Duration) *DefuddleExtractor {
	return &DefuddleExtractor{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				ForceAttemptHTTP2:   true,
				MaxIdleConns:        5,
				IdleConnTimeout:     30 * time.Second,
				TLSHandshakeTimeout: 10 * time.Second,
			},
		},
	}
}

func (d *DefuddleExtractor) Name() string { return "defuddle" }

// Extract sends a GET request to <base_url>/<domain>/<path> (no scheme)
// and returns the plain markdown response. The service does not report the origin's
// status or final URL, so Status is 0 (unknown) and FinalURL is the requested URL.
func (d *DefuddleExtractor) Extract(ctx context.Context, req ExtractRequest) (ExtractOutput, error) {
	rawURL := req.URL
	// Strip scheme: https://example.com/path → example.com/path
	target := strings.TrimPrefix(strings.TrimPrefix(rawURL, "https://"), "http://")
	fetchURL := d.baseURL + target

	// Context timeout ensures cancellation propagates even if transport ignores client.Timeout.
	ctx, cancel := context.WithTimeout(ctx, d.client.Timeout)
	defer cancel()

	httpReq, err := http.NewRequestWithContext(ctx, "GET", fetchURL, nil)
	if err != nil {
		return ExtractOutput{}, fmt.Errorf("create defuddle request: %w", err)
	}
	httpReq.Header.Set("User-Agent", fetchUserAgent)

	resp, err := d.client.Do(httpReq)
	if err != nil {
		return ExtractOutput{}, fmt.Errorf("defuddle fetch: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ExtractOutput{}, fmt.Errorf("defuddle returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, defuddleMaxBody))
	if err != nil {
		return ExtractOutput{}, fmt.Errorf("read defuddle response: %w", err)
	}

	return ExtractOutput{Content: string(body), FinalURL: rawURL, Method: "external"}, nil
}
