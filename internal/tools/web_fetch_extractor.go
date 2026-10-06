package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

// ContentExtractor extracts readable content from a URL.
type ContentExtractor interface {
	Extract(ctx context.Context, req ExtractRequest) (ExtractOutput, error)
	Name() string
}

// ExtractRequest describes one extraction.
type ExtractRequest struct {
	URL      string
	MaxChars int  // 0 = defaultFetchMaxChars
	FullPage bool // skip main-content extraction
}

// ExtractOutput is the content an extractor produced, with what it knows about the response.
type ExtractOutput struct {
	Content        string
	Extractor      string // content extractor actually used (e.g. "json" for a JSON response); empty = Name()
	FinalURL       string // URL after redirects
	Status         int    // origin HTTP status; 0 = unknown (external service)
	Method         string // HTML: contentMain | contentFullPage; "external" for an extraction service
	FallbackReason string // why the full page was used instead of the main content
	Meta           pageMeta
}

// ExtractResult holds the output from a successful extraction.
type ExtractResult struct {
	Content        string
	Extractor      string // name of the extractor that produced the content
	FinalURL       string
	Status         int
	Method         string
	FallbackReason string
	Meta           pageMeta
}

// ExtractorChain tries extractors in order until one returns quality content.
type ExtractorChain struct {
	extractors []ContentExtractor
	maxRetries []int           // per-extractor max attempts (default 1)
	timeouts   []time.Duration // per-extractor chain-level timeout (0 = no chain timeout)
	// inProcess runs a full-page request when every configured extractor is skipped
	// (only external extractors enabled). nil when the chain was built without a tool.
	inProcess *InProcessExtractor
}

// NewExtractorChain creates a chain from ordered extractors with default settings (1 attempt, no chain timeout).
func NewExtractorChain(extractors ...ContentExtractor) *ExtractorChain {
	maxRetries := make([]int, len(extractors))
	timeouts := make([]time.Duration, len(extractors))
	for i := range extractors {
		maxRetries[i] = 1
	}
	return &ExtractorChain{extractors: extractors, maxRetries: maxRetries, timeouts: timeouts}
}

// Extract runs each extractor in order with per-entry retry and optional timeout.
// Returns the first quality result or cascades to the next extractor.
func (c *ExtractorChain) Extract(ctx context.Context, req ExtractRequest) (ExtractResult, error) {
	rawURL := req.URL
	var lastErr error
	ran := false
	for i, ext := range c.extractors {
		// A full-page request must not go to an external service: it returns its own
		// extraction, not the page.
		if _, inProcess := ext.(*InProcessExtractor); req.FullPage && !inProcess {
			continue
		}
		ran = true
		maxRetries := c.maxRetries[i]

		for attempt := 1; attempt <= maxRetries; attempt++ {
			// Apply chain-level timeout if configured.
			callCtx, cancel := ctx, context.CancelFunc(nil)
			if timeout := c.timeouts[i]; timeout > 0 {
				callCtx, cancel = context.WithTimeout(ctx, timeout)
			}

			out, err := ext.Extract(callCtx, req)
			content := out.Content
			if cancel != nil {
				cancel()
			}
			if err != nil {
				// An HTTP error from the origin or a redirect refused by policy is final: every
				// extractor reads the same URL, and another one could hide the error behind a
				// 200 or follow the refused redirect itself.
				if isTerminalFetchError(err) {
					return ExtractResult{}, err
				}
				lastErr = err
				if ctx.Err() != nil {
					return ExtractResult{}, fmt.Errorf("context cancelled: %w", lastErr)
				}
				if attempt < maxRetries {
					slog.Warn("extractor_chain: attempt failed, retrying",
						"extractor", ext.Name(), "url", rawURL,
						"attempt", attempt, "max_retries", maxRetries,
						"error", err)
				} else {
					slog.Debug("extractor failed", "extractor", ext.Name(), "url", rawURL, "error", err)
				}
				continue
			}
			if !isQualityContent(content) {
				slog.Debug("extractor returned low quality content", "extractor", ext.Name(), "url", rawURL, "chars", len(content))
				lastErr = fmt.Errorf("%s: content below quality threshold (%d chars)", ext.Name(), len(content))
				break // low quality is not transient — don't retry, cascade to next
			}
			return newExtractResult(out, ext.Name()), nil
		}

		slog.Debug("extractor_chain: extractor exhausted, moving to next",
			"extractor", ext.Name(), "max_retries", maxRetries)
	}
	if !ran && c.inProcess != nil {
		out, err := c.inProcess.Extract(ctx, req)
		if err != nil {
			return ExtractResult{}, err
		}
		return newExtractResult(out, c.inProcess.Name()), nil
	}
	if lastErr != nil {
		return ExtractResult{}, fmt.Errorf("all extractors failed for %s: %w", rawURL, lastErr)
	}
	return ExtractResult{}, fmt.Errorf("no extractors configured")
}

func newExtractResult(out ExtractOutput, name string) ExtractResult {
	extractor := out.Extractor
	if extractor == "" {
		extractor = name
	}
	return ExtractResult{
		Content:        out.Content,
		Extractor:      extractor,
		FinalURL:       out.FinalURL,
		Status:         out.Status,
		Method:         out.Method,
		FallbackReason: out.FallbackReason,
		Meta:           out.Meta,
	}
}

// isQualityContent checks if extracted content meets minimum quality thresholds.
// Returns false for empty, very short (<100 chars), or low word count (<10 words) content.
func isQualityContent(content string) bool {
	trimmed := strings.TrimSpace(content)
	if len(trimmed) < 100 {
		return false
	}
	return len(strings.Fields(trimmed)) >= 10
}

// ---------------------------------------------------------------------------
// Extractor chain settings — stored in builtin_tools.settings for web_fetch
// ---------------------------------------------------------------------------

// ExtractorEntry represents a single extractor in the chain settings JSON.
type ExtractorEntry struct {
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	Timeout    int    `json:"timeout,omitempty"`     // seconds, 0 = use extractor default
	MaxRetries int    `json:"max_retries,omitempty"` // default 1 (no retry)
	BaseURL    string `json:"base_url,omitempty"`    // for defuddle: extraction service URL (required, no default)
}

// extractorChainSettings is the JSON schema for web_fetch builtin_tools.settings.
type extractorChainSettings struct {
	Extractors []ExtractorEntry `json:"extractors,omitempty"`
}

// ResolveExtractorChain parses builtin_tools.settings from context and builds
// an ordered ExtractorChain. Returns nil if no extractors are enabled.
func ResolveExtractorChain(ctx context.Context, tool *WebFetchTool) *ExtractorChain {
	if settings := BuiltinToolSettingsFromCtx(ctx); settings != nil {
		if raw, ok := settings["web_fetch"]; ok && len(raw) > 0 {
			chain := parseExtractorChainSettings(raw, tool)
			if chain != nil {
				return chain
			}
		}
	}
	// Default fallback: InProcess only (no external extractors).
	chain := NewExtractorChain(&InProcessExtractor{tool: tool})
	chain.inProcess = &InProcessExtractor{tool: tool}
	return chain
}

// parseExtractorChainSettings parses the settings JSON and builds a chain.
func parseExtractorChainSettings(raw []byte, tool *WebFetchTool) *ExtractorChain {
	var settings extractorChainSettings
	if err := json.Unmarshal(raw, &settings); err != nil {
		slog.Warn("web_fetch: failed to parse extractor chain settings", "error", err)
		return nil
	}

	var extractors []ContentExtractor
	var maxRetries []int
	var timeouts []time.Duration
	for _, entry := range settings.Extractors {
		if !entry.Enabled || entry.Name == "" {
			continue
		}
		switch entry.Name {
		case "defuddle":
			// No built-in endpoint: an enabled defuddle entry without base_url is skipped so
			// no URL is ever sent to a service the operator did not configure.
			if strings.TrimSpace(entry.BaseURL) == "" {
				slog.Warn("web_fetch: defuddle extractor enabled without base_url, skipping")
				continue
			}
			extractors = append(extractors, NewDefuddleExtractorFromEntry(entry))
		case "html-to-markdown":
			extractors = append(extractors, &InProcessExtractor{tool: tool})
		default:
			slog.Warn("web_fetch: unknown extractor in chain, skipping", "name", entry.Name)
			continue
		}
		retries := entry.MaxRetries
		if retries <= 0 {
			retries = 1
		}
		maxRetries = append(maxRetries, retries)
		timeouts = append(timeouts, time.Duration(entry.Timeout)*time.Second)
	}
	if len(extractors) == 0 {
		return nil
	}
	return &ExtractorChain{extractors: extractors, maxRetries: maxRetries, timeouts: timeouts, inProcess: &InProcessExtractor{tool: tool}}
}
