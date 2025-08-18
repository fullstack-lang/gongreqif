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
	// The parser works better without the namespace prefixes.
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
		mdBuilder.WriteString(renderNode(c))
	}

	// Final cleanup: trim leading/trailing space and ensure no more than two consecutive newlines.
	finalOutput := strings.TrimSpace(mdBuilder.String())
	re := regexp.MustCompile(`\n{3,}`)
	finalOutput = re.ReplaceAllString(finalOutput, "\n\n")

	return finalOutput, nil
}

// renderNode is a recursive function that processes a node and its children into a Markdown string.
func renderNode(n *html.Node) string {
	switch n.Type {
	case html.TextNode:
		// If a text node is only whitespace, it's for formatting.
		// Return a single space to be safe; the parent's normalization will handle it.
		if strings.TrimSpace(n.Data) == "" {
			return " "
		}
		return n.Data

	case html.ElementNode:
		// Recursively process all children and concatenate their results into raw content.
		var contentBuilder strings.Builder
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			contentBuilder.WriteString(renderNode(c))
		}
		content := contentBuilder.String()

		// Apply Markdown formatting based on the element type.
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
			// Indent each line of the blockquote content.
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
					result.WriteString("- " + renderNode(c) + "\n")
				}
			}
			return result.String()
		case "ol":
			var result strings.Builder
			result.WriteString("\n")
			itemCounter := 1
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "li" {
					result.WriteString(fmt.Sprintf("%d. %s\n", itemCounter, renderNode(c)))
					itemCounter++
				}
			}
			return result.String()
		case "li":
			// For list items, we just return the cleaned content.
			// The parent <ol> or <ul> is responsible for the bullet/number.
			return normalizeWhitespace(content)

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
			return "  \n" // Markdown for a hard line break

		default:
			// For unknown tags, just return their processed content.
			return content
		}
	}
	return ""
}

// normalizeWhitespace collapses all forms of whitespace (spaces, tabs, newlines)
// into a single space and trims the leading/trailing space.
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
