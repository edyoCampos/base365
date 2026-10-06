package tools

import (
	"fmt"
	"log/slog"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"

	readability "codeberg.org/readeck/go-readability/v2"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	// mainContentMinRatio is the smallest main-content/full-page size ratio accepted. Below it
	// the page is a listing or front page whose "article" is a fragment, so the full page wins.
	mainContentMinRatio = 0.15
	// mainContentMaxElems skips main-content extraction on documents larger than this.
	mainContentMaxElems = 100_000
	// pageMetaMaxRunes caps each metadata value shown in the result header.
	pageMetaMaxRunes = 300
	// headingWrapperMaxRunes is the most text a heading wrapper may hold besides the heading
	// (edit links, anchors) for the wrapper to be unwrapped.
	headingWrapperMaxRunes = 32
)

// Content methods and fallback reasons reported in the web_fetch result header.
const (
	contentMain     = "main"
	contentFullPage = "full-page"

	fallbackRequested       = "requested"
	fallbackExtractionError = "extraction-error"
	fallbackEmpty           = "empty"
	fallbackLowQuality      = "low-quality"
	fallbackBelowRatio      = "below-ratio"
	fallbackTooManyElements = "too-many-elements"
)

// pageMeta holds page metadata. Values are untrusted page content, already single-line and capped.
type pageMeta struct {
	Title       string
	Author      string
	Published   string // YYYY-MM-DD
	Description string
	Site        string
	Language    string
}

// pageResult is the outcome of HTML extraction.
type pageResult struct {
	content        string
	method         string // contentMain | contentFullPage
	fallbackReason string // empty when method is contentMain
	meta           pageMeta
}

// extractPage renders an HTML page as markdown or text. Unless fullPage is set it extracts the
// main content (readability) and falls back to the whole page when that fails or is a fragment.
// When nothing reaches the quality bar, a second pass renders <noscript> content from <body>.
// charsetName is only logged.
func extractPage(htmlText string, pageURL *url.URL, mode convertMode, fullPage bool, charsetName string) pageResult {
	start := time.Now()
	res, elems := safeExtractPageOnce(htmlText, pageURL, mode, fullPage, false)
	secondPass := false
	if !isQualityContent(res.content) {
		if retry, _ := safeExtractPageOnce(htmlText, pageURL, mode, fullPage, true); isQualityContent(retry.content) {
			res, secondPass = retry, true
		}
	}
	slog.Debug("web_fetch.extract",
		"method", res.method, "fallback_reason", res.fallbackReason, "noscript_pass", secondPass,
		"charset", charsetName, "elements", elems, "ms", time.Since(start).Milliseconds())
	return res
}

// safeExtractPageOnce runs extractPageOnce, turning a panic anywhere in parsing, readability or
// metadata parsing (page-controlled input) into an extraction error with no content.
func safeExtractPageOnce(htmlText string, pageURL *url.URL, mode convertMode, fullPage, unwrapNoscript bool) (res pageResult, elems int) {
	defer func() {
		if r := recover(); r != nil {
			// Log only the panic's type: its value may carry page text.
			slog.Warn("web_fetch.extract_panic", "panic_type", fmt.Sprintf("%T", r))
			res, elems = pageResult{content: "", method: contentFullPage, fallbackReason: fallbackExtractionError}, 0
		}
	}()
	return extractPageOnce(htmlText, pageURL, mode, fullPage, unwrapNoscript)
}

func extractPageOnce(htmlText string, pageURL *url.URL, mode convertMode, fullPage, unwrapNoscript bool) (pageResult, int) {
	doc, elems, err := prepareDocument(htmlText, unwrapNoscript)
	if err != nil {
		return pageResult{content: unparseableHTMLNotice, method: contentFullPage, fallbackReason: fallbackExtractionError}, 0
	}

	// Render the full page before readability mutates the tree.
	full := renderNode(findBody(doc), mode)
	fallback := func(reason string, meta pageMeta) pageResult {
		return pageResult{content: full, method: contentFullPage, fallbackReason: reason, meta: meta}
	}
	docMeta := pageMeta{Title: cleanMeta(documentTitle(doc))}

	if fullPage {
		return fallback(fallbackRequested, docMeta), elems
	}
	if elems > mainContentMaxElems {
		return fallback(fallbackTooManyElements, docMeta), elems
	}

	if pageURL == nil {
		pageURL = &url.URL{}
	}
	article, err := parseArticle(doc, pageURL)
	if err != nil {
		return fallback(fallbackExtractionError, docMeta), elems
	}
	meta := articleMeta(article, docMeta.Title)
	if article.Node == nil {
		return fallback(fallbackEmpty, meta), elems
	}

	main := renderNode(article.Node, mode)
	switch {
	case !isQualityContent(main):
		return fallback(fallbackLowQuality, meta), elems
	case float64(len(main)) < mainContentMinRatio*float64(len(full)):
		return fallback(fallbackBelowRatio, meta), elems
	}
	meta.Description = dropRedundantDescription(meta.Description, main)
	return pageResult{content: main, method: contentMain, meta: meta}, elems
}

// parseArticle runs readability on doc (mutating it), converting a panic on hostile input
// into an error. A new Parser per call: it carries per-parse state. Never use FromURL here:
// fetching happens only through the SSRF-protected client.
func parseArticle(doc *html.Node, pageURL *url.URL) (article readability.Article, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("readability panic: %v", r)
		}
	}()
	parser := readability.NewParser()
	parser.MaxElemsToParse = mainContentMaxElems
	return parser.ParseAndMutate(doc, pageURL)
}

// prepareDocument parses the page and cleans it for extraction. It drops hidden elements
// before extraction because readability strips the class attributes the hidden-element check
// relies on. With unwrapNoscript, each <noscript> inside <body> is replaced by its content,
// parsed as an isolated fragment (so an unclosed tag cannot swallow the rest of the page) and
// filtered like the rest. <noscript> in <head> is never unwrapped (tracking pixels).
// Returns the element count of the cleaned document.
func prepareDocument(htmlText string, unwrapNoscript bool) (*html.Node, int, error) {
	doc, err := html.Parse(strings.NewReader(htmlText))
	if err != nil {
		return nil, 0, err
	}
	if unwrapNoscript {
		if body := findBody(doc); body.DataAtom == atom.Body {
			unwrapNoscriptElements(body)
		}
	}
	// Drop every <noscript> left (all of <head>, and <body> on the first pass): the converter
	// skips them anyway, but readability would re-parse their raw text and reinsert images
	// whose hidden-marking attributes it then strips.
	removeElements(doc, atom.Noscript)
	removeHiddenElements(doc)
	unwrapHeadingWrappers(doc)
	return doc, countElements(doc), nil
}

func unwrapNoscriptElements(body *html.Node) {
	var found []*html.Node
	var collect func(*html.Node)
	collect = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && c.DataAtom == atom.Noscript {
				found = append(found, c)
				continue
			}
			collect(c)
		}
	}
	collect(body)

	context := &html.Node{Type: html.ElementNode, Data: "body", DataAtom: atom.Body}
	for _, ns := range found {
		var raw strings.Builder
		for c := ns.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				raw.WriteString(c.Data)
			}
		}
		nodes, err := html.ParseFragment(strings.NewReader(raw.String()), context)
		if err != nil {
			continue
		}
		for _, n := range nodes {
			ns.Parent.InsertBefore(n, ns)
		}
		ns.Parent.RemoveChild(ns)
	}
}

func removeElements(n *html.Node, a atom.Atom) {
	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		if c.Type == html.ElementNode && c.DataAtom == a {
			n.RemoveChild(c)
		} else {
			removeElements(c, a)
		}
		c = next
	}
}

func removeHiddenElements(n *html.Node) {
	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		if c.Type == html.ElementNode && isHiddenElement(c) {
			n.RemoveChild(c)
		} else {
			removeHiddenElements(c)
		}
		c = next
	}
}

// unwrapHeadingWrappers replaces wrappers such as <div class="mw-heading"><h2>…</h2><span>edit</span></div>
// with the heading, because readability drops headings nested in a <div>. A wrapper qualifies
// only when everything besides the heading is link decoration (edit links, anchors), so
// dropping that decoration loses no content.
func unwrapHeadingWrappers(n *html.Node) {
	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		if heading := wrappedHeading(c); heading != nil {
			c.RemoveChild(heading)
			n.InsertBefore(heading, c)
			n.RemoveChild(c)
		} else {
			unwrapHeadingWrappers(c)
		}
		c = next
	}
}

// wrappedHeading returns the heading of a <div> whose first element child is h1–h6 and whose
// other children are whitespace or link decoration (<a>, or <span> holding only links/markup)
// with at most headingWrapperMaxRunes of text in total; otherwise nil.
func wrappedHeading(n *html.Node) *html.Node {
	if n.Type != html.ElementNode || n.DataAtom != atom.Div {
		return nil
	}
	var heading *html.Node
	extra := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch {
		case c.Type == html.TextNode:
			if strings.TrimSpace(c.Data) != "" {
				return nil
			}
		case c.Type != html.ElementNode:
		case heading == nil:
			if !isHeadingAtom(c.DataAtom) {
				return nil
			}
			heading = c
		case isLinkDecoration(c):
			extra += utf8.RuneCountInString(strings.TrimSpace(nodeText(c)))
		default:
			return nil
		}
	}
	if heading == nil || extra > headingWrapperMaxRunes {
		return nil
	}
	return heading
}

// isLinkDecoration reports whether n is an <a>, or a <span> whose text sits only inside links
// (e.g. Wikipedia's <span class="mw-editsection">[<a>edit</a>]</span>).
func isLinkDecoration(n *html.Node) bool {
	switch n.DataAtom {
	case atom.A:
		return true
	case atom.Span:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			switch {
			case c.Type == html.TextNode:
				if t := strings.Trim(strings.TrimSpace(c.Data), "[]|"); t != "" {
					return false
				}
			case c.Type == html.ElementNode && !isLinkDecoration(c):
				return false
			}
		}
		return true
	}
	return false
}

func isHeadingAtom(a atom.Atom) bool {
	switch a {
	case atom.H1, atom.H2, atom.H3, atom.H4, atom.H5, atom.H6:
		return true
	}
	return false
}

func nodeText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(nodeText(c))
	}
	return sb.String()
}

func countElements(n *html.Node) int {
	count := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			count++
		}
		count += countElements(c)
	}
	return count
}

// documentTitle returns the text of the first <title> element.
func documentTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.DataAtom == atom.Title {
		return nodeText(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if t := documentTitle(c); t != "" {
			return t
		}
	}
	return ""
}

func articleMeta(a readability.Article, fallbackTitle string) pageMeta {
	meta := pageMeta{
		Title:       cleanMeta(a.Title()),
		Author:      cleanMeta(a.Byline()),
		Description: cleanMeta(a.Excerpt()),
		Site:        cleanMeta(a.SiteName()),
		Language:    cleanMeta(a.Language()),
	}
	if meta.Title == "" {
		meta.Title = fallbackTitle
	}
	meta.Published = publishedDate(a)
	return meta
}

// publishedDate formats the article's published time as YYYY-MM-DD, or "" when absent or
// unparseable. The date parser runs on page-controlled text and can panic on hostile input.
func publishedDate(a readability.Article) (date string) {
	defer func() {
		if recover() != nil {
			date = ""
		}
	}()
	if t, err := a.PublishedTime(); err == nil && !t.IsZero() {
		return t.Format("2006-01-02")
	}
	return ""
}

// cleanMeta collapses whitespace (so a value cannot forge header lines) and caps the length.
func cleanMeta(s string) string {
	return truncateRunes(strings.Join(strings.Fields(s), " "), pageMetaMaxRunes)
}

// dropRedundantDescription drops a description that only repeats the start of the content
// (readability falls back to the first paragraph when the page has no description).
func dropRedundantDescription(desc, content string) string {
	if desc == "" {
		return ""
	}
	probe := truncateRunes(desc, 80)
	head := strings.Join(strings.Fields(truncateRunes(content, 1500)), " ")
	if strings.Contains(head, probe) {
		return ""
	}
	return desc
}

// truncateRunes cuts s to at most n runes.
func truncateRunes(s string, n int) string {
	if utf8.RuneCountInString(s) <= n {
		return s
	}
	i := 0
	for pos := range s {
		if i == n {
			return s[:pos]
		}
		i++
	}
	return s
}
