package tools

import (
	"context"
)

// InProcessExtractor delegates to WebFetchTool.fetchRaw for HTML→markdown
// extraction with full security checks (SSRF, domain policy on redirects).
// This is the fallback when external extractors (Defuddle) are unavailable.
type InProcessExtractor struct {
	tool *WebFetchTool
}

func (e *InProcessExtractor) Name() string { return "html-to-markdown" }

// Extract fetches the URL via the tool's fetchRaw (full security checks)
// and returns the extracted markdown with the response details.
func (e *InProcessExtractor) Extract(ctx context.Context, req ExtractRequest) (ExtractOutput, error) {
	maxChars := req.MaxChars
	if maxChars <= 0 {
		maxChars = defaultFetchMaxChars
	}
	pol := e.tool.resolvePolicy(ctx)
	raw, err := e.tool.fetchRaw(ctx, req.URL, fetchOptions{extractMode: "markdown", maxChars: maxChars, fullPage: req.FullPage}, pol)
	if err != nil {
		return ExtractOutput{}, err
	}
	return ExtractOutput{
		Content:        raw.content,
		Extractor:      raw.extractor,
		FinalURL:       raw.finalURL,
		Status:         raw.statusCode,
		Method:         raw.method,
		FallbackReason: raw.fallbackReason,
		Meta:           raw.meta,
	}, nil
}
