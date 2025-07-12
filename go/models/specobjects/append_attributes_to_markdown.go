package specobjects

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

// Pre-compile regexes for efficiency.
var (
	// tagRegex matches any XML/XHTML-style tag.
	tagRegex = regexp.MustCompile(`<[^>]*>`)
	// spaceRegex matches one or more whitespace characters.
	spaceRegex = regexp.MustCompile(`\s+`)
)

// sanitizeForMarkdownTable escapes characters that can break a markdown table's structure.
func sanitizeForMarkdownTable(s string) string {
	// Replace pipe characters, which define columns
	s = strings.ReplaceAll(s, "|", "&#124;")
	// Replace newlines with a space to keep the row on a single line
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}

// AppendAttributesToMarkdown populates a markdown string with a table of attributes from a SPEC_OBJECT.
func AppendAttributesToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	var tableRows strings.Builder

	// Gather all attribute rows into the tableRows builder
	appendAttributeXHTMLRows(stager, specObject, &tableRows)
	appendAttributeStringRows(stager, specObject, &tableRows)
	appendAttributeBooleanRows(stager, specObject, &tableRows)
	appendAttributeIntegerRows(stager, specObject, &tableRows)
	appendAttributeRealRows(stager, specObject, &tableRows)
	appendAttributeDateRows(stager, specObject, &tableRows)
	appendAttributeEnumRows(stager, specObject, &tableRows)

	// If any attributes were found, create the table.
	if tableRows.Len() > 0 {
		*markdownContent += "\n|  |  |\n"
		*markdownContent += "|---|---|\n"
		*markdownContent += tableRows.String()
	}
}

// appendAttributeXHTMLRows formats and appends XHTML attributes as markdown table rows.
func appendAttributeXHTMLRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
		var attributeDefinition string
		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
			attributeDefinition = datatype.LONG_NAME
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_XHTML_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF, "unknown ref")
		}

		// Clean up XHTML content for markdown readability
		enclosedText := attribute.THE_VALUE.EnclosedText

		// 1. Remove all tags using a regular expression.
		textWithoutTags := tagRegex.ReplaceAllString(enclosedText, "")

		// 2. Normalize whitespace: replace newlines, tabs, and multiple spaces with a single space.
		normalizedText := spaceRegex.ReplaceAllString(textWithoutTags, " ")

		// 3. Trim leading/trailing space from the final string.
		finalText := strings.TrimSpace(normalizedText)

		tableRows.WriteString(fmt.Sprintf("| %s: | %s |\n",
			sanitizeForMarkdownTable(attributeDefinition),
			sanitizeForMarkdownTable(finalText)))
	}
}

// appendAttributeStringRows formats and appends String attributes as markdown table rows.
func appendAttributeStringRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_STRING {
		var attributeDefinition string
		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_STRING[attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF]; ok {
			attributeDefinition = datatype.LONG_NAME
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_STRING_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF, "unknown ref")
		}
		tableRows.WriteString(fmt.Sprintf("| %s: | %s |\n",
			sanitizeForMarkdownTable(attributeDefinition),
			sanitizeForMarkdownTable(attribute.THE_VALUE)))
	}
}

// appendAttributeBooleanRows formats and appends Boolean attributes as markdown table rows.
func appendAttributeBooleanRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		var attributeDefinition string
		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN[attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF]; ok {
			attributeDefinition = datatype.LONG_NAME
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_BOOLEAN_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF, "unknown ref")
		}
		tableRows.WriteString(fmt.Sprintf("| %s | %t |\n",
			sanitizeForMarkdownTable(attributeDefinition),
			attribute.THE_VALUE))
	}
}

// appendAttributeIntegerRows formats and appends Integer attributes as markdown table rows.
func appendAttributeIntegerRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_INTEGER {
		var attributeDefinition string
		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[attribute.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF]; ok {
			attributeDefinition = datatype.LONG_NAME
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_INTEGER_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF, "unknown ref")
		}
		tableRows.WriteString(fmt.Sprintf("| %s | %d |\n",
			sanitizeForMarkdownTable(attributeDefinition),
			attribute.THE_VALUE))
	}
}

// appendAttributeDateRows formats and appends Date attributes as markdown table rows.
func appendAttributeDateRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_DATE {
		var attributeDefinition string
		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_DATE[attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF]; ok {
			attributeDefinition = datatype.LONG_NAME
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_DATE_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF, "unknown ref")
		}
		tableRows.WriteString(fmt.Sprintf("| %s: | %s |\n",
			sanitizeForMarkdownTable(attributeDefinition),
			sanitizeForMarkdownTable(attribute.THE_VALUE)))
	}
}

// appendAttributeRealRows formats and appends Real attributes as markdown table rows.
func appendAttributeRealRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_REAL {
		var attributeDefinition string
		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_REAL[attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF]; ok {
			attributeDefinition = datatype.LONG_NAME
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_REAL_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF, "unknown ref")
		}
		tableRows.WriteString(fmt.Sprintf("| %s | %f |\n",
			sanitizeForMarkdownTable(attributeDefinition),
			attribute.THE_VALUE))
	}
}

// appendAttributeEnumRows formats and appends Enumeration attributes as markdown table rows.
func appendAttributeEnumRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		var enumTypeString string
		if enumType, ok := stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF]; ok {
			enumTypeString = enumType.LONG_NAME
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_ENUMERATION_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF, "unknown ref")
		}

		var enumValueString string
		if len(attribute.VALUES.ENUM_VALUE_REF) > 0 {
			valueIdentifier := attribute.VALUES.ENUM_VALUE_REF
			if enumValue, ok := stager.Map_id_ENUM_VALUE[valueIdentifier]; ok {
				enumValueString = enumValue.LONG_NAME
			} else {
				log.Panic("ENUM_VALUE_REF", valueIdentifier, "unknown ref in Map_id_ENUM_VALUE")
			}
		}

		tableRows.WriteString(fmt.Sprintf("| %s: | %s |\n",
			sanitizeForMarkdownTable(enumTypeString),
			sanitizeForMarkdownTable(enumValueString)))
	}
}
