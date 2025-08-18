// File: xhtml/converter.go

package xhtml

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// ToMarkdown converts a given XHTML string to its Markdown equivalent.
func ToMarkdown(xhtmlString string) (string, error) {
	sanitizedXHTML := strings.ReplaceAll(xhtmlString, "xhtml:", "")

	doc, err := html.Parse(strings.NewReader(sanitizedXHTML))
	if err != nil {
		return "", fmt.Errorf("failed to parse xhtml: %w", err)
	}

	bodyNode := findBody(doc)
	if bodyNode == nil {
		return "", fmt.Errorf("could not find <body> tag in parsed document")
	}

	var mdBuilder strings.Builder
	for c := bodyNode.FirstChild; c != nil; c = c.NextSibling {
		// Pass the parent node to renderNode for context (important for lists)
		mdBuilder.WriteString(renderNode(c, c.Parent))
	}

	finalOutput := strings.TrimSpace(mdBuilder.String())
	re := regexp.MustCompile(`\n{3,}`)
	finalOutput = re.ReplaceAllString(finalOutput, "\n\n")

	return finalOutput, nil
}

// renderNode is the main recursive function for processing nodes.
// It now takes a parent node for context.
func renderNode(n *html.Node, parent *html.Node) string {
	switch n.Type {
	case html.TextNode:
		// Discard text nodes that are just formatting whitespace between block elements.
		if parent != nil && (parent.Data == "body" || parent.Data == "blockquote" || parent.Data == "div") {
			if strings.TrimSpace(n.Data) == "" {
				return ""
			}
		}
		return n.Data

	case html.ElementNode:
		var contentBuilder strings.Builder
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			contentBuilder.WriteString(renderNode(c, n))
		}
		content := contentBuilder.String()

		switch n.Data {
		// BLOCK-LEVEL ELEMENTS
		case "p":
			return "\n\n" + normalizeWhitespace(content)
		case "h1":
			return "\n\n# " + normalizeWhitespace(content)
		case "h2":
			return "\n\n## " + normalizeWhitespace(content)
		case "h3":
			return "\n\n### " + normalizeWhitespace(content)
		case "h4":
			return "\n\n#### " + normalizeWhitespace(content)
		case "blockquote":
			lines := strings.Split(normalizeWhitespace(content), "\n")
			for i, line := range lines {
				lines[i] = "> " + line
			}
			return "\n\n" + strings.Join(lines, "\n")
		case "ul":
			var result strings.Builder
			result.WriteString("\n")
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "li" {
					result.WriteString("- " + renderNode(c, n) + "\n")
				}
			}
			return result.String()
		case "ol":
			var result strings.Builder
			result.WriteString("\n")
			itemCounter := 1
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "li" {
					result.WriteString(fmt.Sprintf("%d. %s\n", itemCounter, renderNode(c, n)))
					itemCounter++
				}
			}
			return result.String()
		case "li":
			return normalizeWhitespace(content)

		// TABLE-SPECIFIC LOGIC
		case "table":
			return "\n\n" + renderTable(n) + "\n\n"

		// INLINE-LEVEL ELEMENTS
		case "strong", "b":
			return "**" + normalizeWhitespace(content) + "**"
		case "em", "i":
			return "*" + normalizeWhitespace(content) + "*"
		case "a":
			href := getAttr(n, "href")
			text := normalizeWhitespace(content)
			return fmt.Sprintf("[%s](%s)", text, href)
		case "br":
			return "  \n"
		default:
			return content
		}
	}
	return ""
}

// renderTable specifically handles the logic for converting a <table> node.
func renderTable(tableNode *html.Node) string {
	var b strings.Builder
	var headerRows [][]string
	var bodyRows [][]string

	// Find thead and tbody
	for child := tableNode.FirstChild; child != nil; child = child.NextSibling {
		if child.Type != html.ElementNode {
			continue
		}
		switch child.Data {
		case "thead":
			headerRows = parseTableRows(child, "th")
		case "tbody":
			bodyRows = parseTableRows(child, "td")
		}
	}

	// If no thead/tbody, parse directly from table
	if len(headerRows) == 0 && len(bodyRows) == 0 {
		bodyRows = parseTableRows(tableNode, "td")
	}

	// Write header
	if len(headerRows) > 0 {
		b.WriteString("| " + strings.Join(headerRows[0], " | ") + " |\n")
		// Write separator
		var separator []string
		for range headerRows[0] {
			separator = append(separator, "---")
		}
		b.WriteString("| " + strings.Join(separator, " | ") + " |\n")
	}

	// Write body
	for _, row := range bodyRows {
		b.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}

	return b.String()
}

// parseTableRows is a helper to extract cell content from a <tr>.
func parseTableRows(node *html.Node, cellTag string) [][]string {
	var rows [][]string
	for tr := node.FirstChild; tr != nil; tr = tr.NextSibling {
		if tr.Type == html.ElementNode && tr.Data == "tr" {
			var row []string
			for cell := tr.FirstChild; cell != nil; cell = cell.NextSibling {
				if cell.Type == html.ElementNode && (cell.Data == cellTag || cell.Data == "td" || cell.Data == "th") {
					row = append(row, normalizeWhitespace(renderNode(cell, tr)))
				}
			}
			rows = append(rows, row)
		}
	}
	return rows
}

// normalizeWhitespace collapses all forms of whitespace into a single space.
func normalizeWhitespace(s string) string {
	re := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(re.ReplaceAllString(s, " "))
}

// --- Helper functions ---

func getAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func findBody(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "body" {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findBody(c); result != nil {
			return result
		}
	}
	return nil
}
