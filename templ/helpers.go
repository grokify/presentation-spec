package templ

import (
	"regexp"
	"strings"
)

// markdownToHTML provides a simple markdown to HTML conversion.
// For full markdown support, use a library like goldmark.
func markdownToHTML(md string) string {
	if md == "" {
		return ""
	}

	html := md

	// Escape HTML first
	html = strings.ReplaceAll(html, "&", "&amp;")
	html = strings.ReplaceAll(html, "<", "&lt;")
	html = strings.ReplaceAll(html, ">", "&gt;")

	// Headers
	html = regexp.MustCompile(`(?m)^### (.+)$`).ReplaceAllString(html, "<h4>$1</h4>")
	html = regexp.MustCompile(`(?m)^## (.+)$`).ReplaceAllString(html, "<h3>$1</h3>")
	html = regexp.MustCompile(`(?m)^# (.+)$`).ReplaceAllString(html, "<h2>$1</h2>")

	// Bold and italic
	html = regexp.MustCompile(`\*\*(.+?)\*\*`).ReplaceAllString(html, "<strong>$1</strong>")
	html = regexp.MustCompile(`\*(.+?)\*`).ReplaceAllString(html, "<em>$1</em>")

	// Code
	html = regexp.MustCompile("`([^`]+)`").ReplaceAllString(html, "<code>$1</code>")

	// Links
	html = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`).ReplaceAllString(html, `<a href="$2">$1</a>`)

	// Unordered lists
	lines := strings.Split(html, "\n")
	var result []string
	inList := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- ") {
			if !inList {
				result = append(result, "<ul>")
				inList = true
			}
			result = append(result, "<li>"+strings.TrimPrefix(trimmed, "- ")+"</li>")
		} else {
			if inList {
				result = append(result, "</ul>")
				inList = false
			}
			result = append(result, line)
		}
	}
	if inList {
		result = append(result, "</ul>")
	}
	html = strings.Join(result, "\n")

	// Ordered lists
	lines = strings.Split(html, "\n")
	result = nil
	inList = false
	numPattern := regexp.MustCompile(`^(\d+)\. (.+)$`)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if matches := numPattern.FindStringSubmatch(trimmed); matches != nil {
			if !inList {
				result = append(result, "<ol>")
				inList = true
			}
			result = append(result, "<li>"+matches[2]+"</li>")
		} else {
			if inList {
				result = append(result, "</ol>")
				inList = false
			}
			result = append(result, line)
		}
	}
	if inList {
		result = append(result, "</ol>")
	}
	html = strings.Join(result, "\n")

	// Paragraphs - wrap non-tagged lines
	lines = strings.Split(html, "\n")
	result = nil
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			result = append(result, "")
		} else if !strings.HasPrefix(trimmed, "<") {
			result = append(result, "<p>"+trimmed+"</p>")
		} else {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}
