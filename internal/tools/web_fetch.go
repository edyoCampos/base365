package tools

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

// Matching TS src/agents/tools/web-fetch.ts constants.
const (
	defaultFetchMaxChars    = 60000
	defaultFetchMaxRedirect = 3
	defaultErrorMaxChars    = 4000
	fetchTimeoutSeconds     = 30
	fetchUserAgent          = "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_7_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
)

// WebFetchTool implements the web_fetch tool matching TS src/agents/tools/web-fetch.ts.
type WebFetchTool struct {
	maxChars       int
	cache          *webCache
	policy         string   // "allow_all" (default), "allowlist"
	allowedDomains []string // domains when policy="allowlist" (supports "*.example.com")
	blockedDomains []string // always checked regardless of policy (supports "*.example.com")
	mu             sync.RWMutex

	// Test seams, nil in production. Unexported on purpose: no constructor or config sets them.
	transport http.RoundTripper         // nil = default transport
	checkSSRF func(rawURL string) error // nil = CheckSSRF
}

// ssrf applies the SSRF check (CheckSSRF unless a test replaced it).
func (t *WebFetchTool) ssrf(rawURL string) error {
	if t.checkSSRF != nil {
		return t.checkSSRF(rawURL)
	}
	return CheckSSRF(rawURL)
}

// WebFetchConfig holds configuration for the web fetch tool.
type WebFetchConfig struct {
	MaxChars       int
	CacheTTL       time.Duration
	Policy         string   // "allow_all" (default), "allowlist"
	AllowedDomains []string // domains when policy="allowlist"
	BlockedDomains []string // always blocked regardless of policy
}

func NewWebFetchTool(cfg WebFetchConfig) *WebFetchTool {
	maxChars := cfg.MaxChars
	if maxChars <= 0 {
		maxChars = defaultFetchMaxChars
	}
	ttl := cfg.CacheTTL
	if ttl <= 0 {
		ttl = defaultCacheTTL
	}
	policy := cfg.Policy
	if policy == "" {
		policy = "allow_all"
	}
	return &WebFetchTool{
		maxChars:       maxChars,
		cache:          newWebCache(defaultCacheMaxEntries, ttl),
		policy:         policy,
		allowedDomains: cfg.AllowedDomains,
		blockedDomains: cfg.BlockedDomains,
	}
}

// UpdatePolicy replaces the domain policy at runtime (called via pub/sub on config change).
func (t *WebFetchTool) UpdatePolicy(policy string, allowed, blocked []string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if policy == "" {
		policy = "allow_all"
	}
	t.policy = policy
	t.allowedDomains = allowed
	t.blockedDomains = blocked
	slog.Info("web_fetch policy updated", "policy", policy, "allowed", len(allowed), "blocked", len(blocked))
}

// webFetchPolicy holds the resolved domain policy for a single request.
type webFetchPolicy struct {
	mode           string // "allow_all" | "allowlist"
	allowedDomains []string
	blockedDomains []string
}

// webFetchPolicyOverride is the tenant settings shape for web_fetch
// (stored in builtin_tool_tenant_configs.settings).
type webFetchPolicyOverride struct {
	Policy         string   `json:"policy,omitempty"`
	AllowedDomains []string `json:"allowed_domains,omitempty"`
	BlockedDomains []string `json:"blocked_domains,omitempty"`
}

// resolvePolicy returns the effective domain policy for this request.
// Checks tenant override via BuiltinToolSettingsFromCtx first; falls back
// to the tool's default policy when no override is present.
func (t *WebFetchTool) resolvePolicy(ctx context.Context) webFetchPolicy {
	if settings := BuiltinToolSettingsFromCtx(ctx); settings != nil {
		if raw, ok := settings["web_fetch"]; ok && len(raw) > 0 {
			var override webFetchPolicyOverride
			if err := json.Unmarshal(raw, &override); err != nil {
				slog.Warn("web_fetch: failed to parse tenant override, using defaults", "error", err)
			} else if override.Policy != "" {
				return webFetchPolicy{
					mode:           override.Policy,
					allowedDomains: override.AllowedDomains,
					blockedDomains: override.BlockedDomains,
				}
			}
		}
	}
	// Fall back to tool defaults
	t.mu.RLock()
	defer t.mu.RUnlock()
	return webFetchPolicy{
		mode:           t.policy,
		allowedDomains: t.allowedDomains,
		blockedDomains: t.blockedDomains,
	}
}

// matchDomainList checks if a hostname matches any pattern in the list.
// Supports exact match ("github.com") and wildcard prefix ("*.example.com").
func matchDomainList(hostname string, patterns []string) bool {
	hostname = strings.ToLower(hostname)
	for _, pattern := range patterns {
		pattern = strings.ToLower(strings.TrimSpace(pattern))
		if pattern == hostname {
			return true
		}
		// Wildcard: *.example.com matches sub.example.com, a.b.example.com
		if strings.HasPrefix(pattern, "*.") {
			suffix := pattern[1:] // ".example.com"
			if strings.HasSuffix(hostname, suffix) && hostname != suffix[1:] {
				return true
			}
		}
	}
	return false
}

func (t *WebFetchTool) Name() string { return "web_fetch" }

func (t *WebFetchTool) Description() string {
	return "Fetch a URL and extract its content. For HTML pages returns the main content (article text without menus, sidebars or cookie banners) with title, author, date and site when available, falling back to the whole page for listings and front pages; set fullPage to get the whole page. Also supports JSON and plain text. If content exceeds the character limit, full content is saved to a temp file — use shell or read_file to access it. Includes SSRF protection."
}

func (t *WebFetchTool) Parameters() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"url": map[string]any{
				"type":        "string",
				"description": "HTTP or HTTPS URL to fetch.",
			},
			"extractMode": map[string]any{
				"type":        "string",
				"description": `Extraction mode ("markdown" or "text"). Default: "markdown".`,
				"enum":        []string{"markdown", "text"},
			},
			"maxChars": map[string]any{
				"type":        "number",
				"description": "Maximum characters to return (truncates when exceeded). Default: 60000. Omit to use the default.",
				"minimum":     100.0,
			},
			"fullPage": map[string]any{
				"type":        "boolean",
				"description": "Return the whole page instead of only the main content. Use for front pages, listings, search results, or when the main-content result is missing something you need (e.g. a price). Default: false.",
			},
		},
		"required": []string{"url"},
	}
}

func (t *WebFetchTool) Execute(ctx context.Context, args map[string]any) *Result {
	rawURL, _ := args["url"].(string)
	if rawURL == "" {
		return ErrorResult("url is required")
	}

	// Validate URL scheme
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return ErrorResult(fmt.Sprintf("invalid URL: %v", err))
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return ErrorResult("only http and https URLs are supported")
	}
	if parsed.Host == "" {
		return ErrorResult("missing hostname in URL")
	}

	// SSRF protection
	if err := t.ssrf(rawURL); err != nil {
		return ErrorResult(fmt.Sprintf("SSRF protection: %v", err))
	}

	// Resolve domain policy (tenant override via ctx, or tool defaults)
	pol := t.resolvePolicy(ctx)
	hostname := parsed.Hostname()

	// Domain blocklist check (always enforced regardless of policy)
	if matchDomainList(hostname, pol.blockedDomains) {
		return ErrorResult(fmt.Sprintf("domain %q is blocked by policy", hostname))
	}

	// Domain allowlist check
	if pol.mode == "allowlist" && !matchDomainList(hostname, pol.allowedDomains) {
		return ErrorResult(fmt.Sprintf("domain %q is not in the allowed domains list", hostname))
	}

	extractMode := "markdown"
	if em, ok := args["extractMode"].(string); ok && (em == "markdown" || em == "text") {
		extractMode = em
	}

	maxChars := t.maxChars
	if mc, ok := args["maxChars"].(float64); ok && int(mc) >= 100 {
		maxChars = int(mc)
	}

	// Adaptive maxChars: reduce as iterations progress to prevent context bloat.
	if prog, ok := IterationProgressFromCtx(ctx); ok && prog.Max > 0 {
		ratio := float64(prog.Current) / float64(prog.Max)
		switch {
		case ratio >= 0.75:
			maxChars = min(maxChars, 10000)
		case ratio >= 0.50:
			maxChars = min(maxChars, 20000)
		}
	}

	fullPage, _ := args["fullPage"].(bool)

	// Check cache (scoped per channel to prevent cross-channel cache poisoning)
	channel := ToolChannelFromCtx(ctx)
	cacheKey := fmt.Sprintf("fetch:%s:%s:%s:%d:%t", channel, rawURL, extractMode, maxChars, fullPage)
	if cached, ok := t.cache.get(cacheKey); ok {
		slog.Debug("web_fetch cache hit", "url", rawURL)
		return NewResult(cached)
	}

	// Fetch
	result, err := t.doFetch(ctx, rawURL, extractMode, maxChars, fullPage, pol)
	if err != nil {
		return ErrorResult(fetchErrorMessage(err))
	}

	wrapped := wrapExternalContent(result, "Web Fetch", true)
	t.cache.set(cacheKey, wrapped)
	return NewResult(wrapped)
}

// fetchErrorMessage builds the error shown to the model. The diagnostic part is truncated; an
// origin error page excerpt is appended afterwards, whole and wrapped as untrusted content, so
// truncation can never cut the wrapper.
func fetchErrorMessage(err error) string {
	msg := "fetch failed: " + truncateStr(err.Error(), defaultErrorMaxChars)
	var statusErr *httpStatusError
	if errors.As(err, &statusErr) && statusErr.Excerpt != "" {
		msg += "\n\nResponse body excerpt:\n" + wrapExternalContent(statusErr.Excerpt, "Web Fetch error response", true)
	}
	return msg
}

func (t *WebFetchTool) doFetch(ctx context.Context, rawURL, extractMode string, maxChars int, fullPage bool, pol webFetchPolicy) (string, error) {
	// For markdown mode, use the extractor chain (Defuddle → InProcess waterfall)
	// resolved from builtin_tools settings stored in context.
	// InProcessExtractor delegates to fetchRaw (same path as doDirectFetch),
	// so no fallthrough is needed — it would just retry the same request.
	if extractMode == "markdown" {
		chain := ResolveExtractorChain(ctx, t)
		if chain != nil {
			result, err := chain.Extract(ctx, ExtractRequest{URL: rawURL, MaxChars: maxChars, FullPage: fullPage})
			if err != nil {
				if isTerminalFetchError(err) {
					return "", err
				}
				return "", fmt.Errorf("all extractors failed: %w", err)
			}
			return formatFetchResult(fetchHeader{
				requestedURL:   rawURL,
				finalURL:       result.FinalURL,
				status:         result.Status,
				extractor:      result.Extractor,
				method:         result.Method,
				fallbackReason: result.FallbackReason,
				meta:           result.Meta,
			}, result.Content, maxChars, ctx), nil
		}
	}

	// Text mode or no chain available — use direct HTTP fetch.
	return t.doDirectFetch(ctx, rawURL, extractMode, maxChars, fullPage, pol)
}

// fetchRawResult holds the output from fetchRaw.
type fetchRawResult struct {
	content        string
	extractor      string
	finalURL       string
	statusCode     int
	method         string // HTML only: contentMain | contentFullPage
	fallbackReason string // HTML only, when method is contentFullPage
	meta           pageMeta
}

// fetchOptions controls a single fetch.
type fetchOptions struct {
	extractMode string // "markdown" | "text"
	maxChars    int
	fullPage    bool // HTML: skip main-content extraction
}

const (
	fetchReadFloor        = 512 * 1024
	fetchHTMLReadFloor    = 2 << 20 // extraction shrinks HTML, so read more of it
	fetchReadCeiling      = 8 << 20 // hard cap on bytes read, whatever maxChars asks for
	httpErrorExcerptRunes = 500
)

// httpStatusError is an HTTP 4xx/5xx response from the origin. It is permanent: retrying or
// asking another extractor for the same URL cannot help. Excerpt is untrusted page text and
// must be wrapped with wrapExternalContent before it reaches the model.
type httpStatusError struct {
	Code    int
	Text    string
	Excerpt string
}

func (e *httpStatusError) Error() string {
	return fmt.Sprintf("HTTP %d %s", e.Code, e.Text)
}

// fetchPolicyError is a redirect refused by SSRF protection, the domain policy or the hop limit. It is
// permanent: another extractor (an external service following redirects itself) must not be
// asked for the same URL, or it would fetch the blocked target.
type fetchPolicyError struct{ err error }

func (e *fetchPolicyError) Error() string { return e.err.Error() }
func (e *fetchPolicyError) Unwrap() error { return e.err }

// isTerminalFetchError reports whether err must end the extractor chain without retry or cascade.
func isTerminalFetchError(err error) bool {
	var statusErr *httpStatusError
	var policyErr *fetchPolicyError
	return errors.As(err, &statusErr) || errors.As(err, &policyErr)
}

// fetchRaw performs HTTP GET with full security checks (SSRF, domain policy on redirects),
// decodes the charset and routes content by type. Returns extracted content without formatting.
// The caller (Execute) has already SSRF-checked rawURL.
func (t *WebFetchTool) fetchRaw(ctx context.Context, rawURL string, opts fetchOptions, pol webFetchPolicy) (fetchRawResult, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", rawURL, nil)
	if err != nil {
		return fetchRawResult{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("User-Agent", fetchUserAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	transport := t.transport
	if transport == nil {
		transport = &http.Transport{
			ForceAttemptHTTP2:   true,
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			TLSHandshakeTimeout: 15 * time.Second,
		}
	}
	redirectCount := 0
	client := &http.Client{
		Timeout:   time.Duration(fetchTimeoutSeconds) * time.Second,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Policy first, so every hop is checked; then the hop limit. Both are terminal: an
			// external extractor that follows redirects itself must not get the same URL.
			if err := t.ssrf(req.URL.String()); err != nil {
				return &fetchPolicyError{fmt.Errorf("redirect SSRF protection: %w", err)}
			}
			redirectHost := req.URL.Hostname()
			if matchDomainList(redirectHost, pol.blockedDomains) {
				return &fetchPolicyError{fmt.Errorf("redirect to %q blocked: domain is in blocklist", redirectHost)}
			}
			if pol.mode == "allowlist" && !matchDomainList(redirectHost, pol.allowedDomains) {
				return &fetchPolicyError{fmt.Errorf("redirect to %q blocked: domain not in allowlist", redirectHost)}
			}
			redirectCount++
			if redirectCount > defaultFetchMaxRedirect {
				return &fetchPolicyError{fmt.Errorf("stopped after %d redirects", defaultFetchMaxRedirect)}
			}
			return nil
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return fetchRawResult{}, err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	isHTML := strings.Contains(contentType, "text/html") || strings.Contains(contentType, "application/xhtml")
	readFloor := fetchReadFloor
	if isHTML {
		readFloor = fetchHTMLReadFloor
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, fetchReadLimit(opts.maxChars, readFloor)))
	if err != nil {
		return fetchRawResult{}, fmt.Errorf("read body: %w", err)
	}
	finalURL := resp.Request.URL

	if resp.StatusCode >= 400 {
		return fetchRawResult{}, &httpStatusError{
			Code:    resp.StatusCode,
			Text:    http.StatusText(resp.StatusCode),
			Excerpt: errorExcerpt(body, contentType, isHTML),
		}
	}

	res := fetchRawResult{finalURL: finalURL.String(), statusCode: resp.StatusCode}
	switch {
	case strings.Contains(contentType, "application/json"):
		res.content, res.extractor = extractJSON(body)

	case strings.Contains(contentType, "text/markdown"):
		res.content, _ = decodeBody(body, contentType)
		res.extractor = "cf-markdown"
		if opts.extractMode == "text" {
			res.content = markdownToText(res.content)
		}

	case isHTML:
		text, charsetName := decodeBody(body, contentType)
		mode, extractor := modeMarkdown, "html-to-markdown"
		if opts.extractMode == "text" {
			mode, extractor = modeText, "html-to-text"
		}
		page := extractPage(text, finalURL, mode, opts.fullPage, charsetName)
		res.content, res.extractor = page.content, extractor
		res.method, res.fallbackReason, res.meta = page.method, page.fallbackReason, page.meta
		if res.content == "" && len(body) > 0 {
			res.content = "[No content extracted. The page may require JavaScript to render, " +
				"or returned a bot-protection challenge. Try using browser automation instead.]"
		}

	case strings.HasPrefix(contentType, "text/"):
		res.content, _ = decodeBody(body, contentType)
		res.extractor = "raw"

	default:
		res.content = string(body)
		res.extractor = "raw"
	}
	return res, nil
}

// fetchReadLimit is max(maxChars×10, floor), capped at fetchReadCeiling. maxChars comes from
// the model and is otherwise unbounded.
func fetchReadLimit(maxChars, floor int) int64 {
	if maxChars <= 0 || maxChars > fetchReadCeiling/10 {
		return fetchReadCeiling
	}
	return int64(max(maxChars*10, floor))
}

// errorExcerpt returns the start of an error response body as plain text.
func errorExcerpt(body []byte, contentType string, isHTML bool) string {
	text, _ := decodeBody(body, contentType)
	if isHTML {
		text = htmlToText(text)
	}
	return truncateRunes(strings.TrimSpace(text), httpErrorExcerptRunes)
}

// doDirectFetch wraps fetchRaw with full HTTP metadata formatting.
// Used for text mode extraction and as ultimate fallback.
func (t *WebFetchTool) doDirectFetch(ctx context.Context, rawURL, extractMode string, maxChars int, fullPage bool, pol webFetchPolicy) (string, error) {
	raw, err := t.fetchRaw(ctx, rawURL, fetchOptions{extractMode: extractMode, maxChars: maxChars, fullPage: fullPage}, pol)
	if err != nil {
		return "", err
	}
	return formatFetchResult(fetchHeader{
		requestedURL:   rawURL,
		finalURL:       raw.finalURL,
		status:         raw.statusCode,
		extractor:      raw.extractor,
		method:         raw.method,
		fallbackReason: raw.fallbackReason,
		meta:           raw.meta,
	}, raw.content, maxChars, ctx), nil
}

// fetchHeader is what the result header reports about a fetch.
type fetchHeader struct {
	requestedURL   string
	finalURL       string // empty = requestedURL
	status         int    // 0 = unknown (external extraction service)
	extractor      string
	method         string // contentMain | contentFullPage | "external" | "" (non-HTML)
	fallbackReason string
	meta           pageMeta
}

// formatFetchResult builds the metadata-prefixed response. Empty fields are omitted.
func formatFetchResult(h fetchHeader, content string, maxChars int, ctx context.Context) string {
	finalURL := h.finalURL
	if finalURL == "" {
		finalURL = h.requestedURL
	}
	var sb strings.Builder
	line := func(key, value string) {
		if value != "" {
			fmt.Fprintf(&sb, "%s: %s\n", key, value)
		}
	}
	line("URL", finalURL)
	if finalURL != h.requestedURL {
		line("Redirected from", h.requestedURL)
	}
	if h.status > 0 {
		line("Status", strconv.Itoa(h.status))
	}
	line("Extractor", h.extractor)
	switch {
	case h.method == contentFullPage && h.fallbackReason != "":
		line("Content", fmt.Sprintf("%s (fallback: %s)", contentFullPage, h.fallbackReason))
	case h.method == contentMain || h.method == contentFullPage:
		line("Content", h.method)
	}
	line("Title", h.meta.Title)
	line("Author", h.meta.Author)
	line("Published", h.meta.Published)
	line("Description", h.meta.Description)
	line("Site", h.meta.Site)
	line("Language", h.meta.Language)
	appendContent(&sb, content, maxChars, finalURL, ctx)
	return sb.String()
}

// cutAtRuneBoundary returns the longest prefix of s that is at most n bytes and does not
// split a UTF-8 rune.
func cutAtRuneBoundary(s string, n int) string {
	if len(s) <= n {
		return s
	}
	for n > 0 && !utf8.RuneStart(s[n]) {
		n--
	}
	return s[:n]
}

// appendContent writes content to the builder, handling truncation and temp file overflow.
func appendContent(sb *strings.Builder, text string, maxChars int, sourceURL string, ctx context.Context) {
	if len(text) > maxChars {
		workspace := ToolWorkspaceFromCtx(ctx)
		tmpPath, writeErr := writeWebFetchTempFile(workspace, text, sourceURL)
		if writeErr != nil {
			slog.Warn("web_fetch: failed to write temp file, falling back to truncation", "error", writeErr)
			text = cutAtRuneBoundary(text, maxChars)
			sb.WriteString(fmt.Sprintf("Truncated: true (limit: %d chars)\n", maxChars))
			sb.WriteString(fmt.Sprintf("Length: %d\n", len(text)))
			sb.WriteString("\n")
			sb.WriteString(text)
		} else {
			sb.WriteString(fmt.Sprintf("Content-Length: %d chars (exceeds %d char limit)\n", len(text), maxChars))
			sb.WriteString(fmt.Sprintf("Full-Content-File: %s\n", tmpPath))
			sb.WriteString(fmt.Sprintf("Length: %d\n", maxChars))
			sb.WriteString("\n")
			sb.WriteString(cutAtRuneBoundary(text, maxChars))
			sb.WriteString(fmt.Sprintf("\n\n[Content truncated at %d chars. Full content (%d chars) saved to: %s — use shell/read_file to access the rest.]",
				maxChars, len(text), tmpPath))
		}
	} else {
		sb.WriteString(fmt.Sprintf("Length: %d\n", len(text)))
		sb.WriteString("\n")
		sb.WriteString(text)
	}
}

// writeWebFetchTempFile saves fetched content to a file with security sanitization.
// When workspace is non-empty, writes to {workspace}/web-fetch/; otherwise falls back to os.TempDir().
func writeWebFetchTempFile(workspace, content, sourceURL string) (string, error) {
	// Generate cryptographically random filename to prevent path prediction
	var randBytes [8]byte
	if _, err := rand.Read(randBytes[:]); err != nil {
		return "", fmt.Errorf("generate random name: %w", err)
	}
	filename := fmt.Sprintf("web-fetch-%s.txt", hex.EncodeToString(randBytes[:]))

	dir := os.TempDir()
	if workspace != "" {
		dir = filepath.Join(workspace, "web-fetch")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("create web-fetch dir: %w", err)
	}
	outPath := filepath.Join(dir, filename)

	// Sanitize content: strip any potential prompt injection markers
	sanitized := sanitizeMarkers(content)

	// Write with restrictive permissions (owner read/write only)
	if err := os.WriteFile(outPath, []byte(sanitized), 0600); err != nil {
		return "", fmt.Errorf("write file: %w", err)
	}

	slog.Info("web_fetch: content saved to file",
		"path", outPath,
		"chars", len(sanitized),
		"source_url", sourceURL,
	)
	return outPath, nil
}
