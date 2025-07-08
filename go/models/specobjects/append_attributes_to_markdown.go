package specobjects

import (
	"fmt"
	"log"
	"strings"

	m "github.com/fullstack-lang/gongreqif/go/models"
)

// AppendAttributesToMarkdown populates a markdown string with all attributes from a SPEC_OBJECT.
// It iterates through each attribute type and appends its formatted content.
func AppendAttributesToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	appendAttributeXHTMLToMarkdown(stager, specObject, markdownContent)
	appendAttributeStringToMarkdown(stager, specObject, markdownContent)
	appendAttributeBooleanToMarkdown(stager, specObject, markdownContent)
	appendAttributeIntegerToMarkdown(stager, specObject, markdownContent)
	appendAttributeRealToMarkdown(stager, specObject, markdownContent)
	appendAttributeDateToMarkdown(stager, specObject, markdownContent)
	appendAttributeEnumToMarkdown(stager, specObject, markdownContent)
}

// appendAttributeXHTMLToMarkdown formats and appends XHTML attributes to the markdown string.
func appendAttributeXHTMLToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_XHTML) > 0 {
		*markdownContent += "\n### XHTML Attributes\n"
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_XHTML_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF, "unknown ref")
			}

			// Clean up XHTML tags for markdown readability
			enclosedText := attribute.THE_VALUE.EnclosedText
			enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:div>", "")
			enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:div>", "")
			enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:br />", "\n")
			enclosedText = strings.TrimSpace(enclosedText)

			*markdownContent += fmt.Sprintf("* **%s:** %s\n", attributeDefinition, enclosedText)
		}
	}
}

// appendAttributeStringToMarkdown formats and appends String attributes to the markdown string.
func appendAttributeStringToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_STRING) > 0 {
		*markdownContent += "\n### String Attributes\n"
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_STRING {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_STRING[attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_STRING_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF, "unknown ref")
			}
			*markdownContent += fmt.Sprintf("* **%s:** %s\n", attributeDefinition, attribute.THE_VALUE)
		}
	}
}

// appendAttributeBooleanToMarkdown formats and appends Boolean attributes to the markdown string.
func appendAttributeBooleanToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN) > 0 {
		*markdownContent += "\n### Boolean Attributes\n"
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN[attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_BOOLEAN_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF, "unknown ref")
			}
			*markdownContent += fmt.Sprintf("* **%s:** %t\n", attributeDefinition, attribute.THE_VALUE)
		}
	}
}

// appendAttributeIntegerToMarkdown formats and appends Integer attributes to the markdown string.
func appendAttributeIntegerToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_INTEGER) > 0 {
		*markdownContent += "\n### Integer Attributes\n"
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_INTEGER {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[attribute.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_INTEGER_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF, "unknown ref")
			}
			*markdownContent += fmt.Sprintf("* **%s:** %d\n", attributeDefinition, attribute.THE_VALUE)
		}
	}
}

// appendAttributeDateToMarkdown formats and appends Date attributes to the markdown string.
func appendAttributeDateToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_DATE) > 0 {
		*markdownContent += "\n### Date Attributes\n"
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_DATE {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_DATE[attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_DATE_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF, "unknown ref")
			}
			*markdownContent += fmt.Sprintf("* **%s:** %s\n", attributeDefinition, attribute.THE_VALUE)
		}
	}
}

// appendAttributeRealToMarkdown formats and appends Real attributes to the markdown string.
func appendAttributeRealToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_REAL) > 0 {
		*markdownContent += "\n### Real Attributes\n"
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_REAL {
			var attributeDefinition string
			if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_REAL[attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF]; ok {
				attributeDefinition = datatype.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_REAL_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF, "unknown ref")
			}
			*markdownContent += fmt.Sprintf("* **%s:** %f\n", attributeDefinition, attribute.THE_VALUE)
		}
	}
}

// appendAttributeEnumToMarkdown formats and appends Enumeration attributes to the markdown string.
func appendAttributeEnumToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	if len(specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION) > 0 {
		*markdownContent += "\n### Enumeration Attributes\n"
		for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			var enumTypeString string
			if enumType, ok := stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF]; ok {
				enumTypeString = enumType.LONG_NAME
			} else {
				log.Panic("ATTRIBUTE_DEFINITION_ENUMERATION_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF, "unknown ref")
			}

			var enumValueString string
			// Check if the enumeration reference string is not empty
			if len(attribute.VALUES.ENUM_VALUE_REF) > 0 {
				// FIX: Use the entire string as the identifier, not just the first byte.
				// The error occurred because ENUM_VALUE_REF is a string, and indexing it
				// (e.g., ENUM_VALUE_REF[0]) returns a byte, which cannot be used as a string key in a map.
				valueIdentifier := attribute.VALUES.ENUM_VALUE_REF
				if enumValue, ok := stager.Map_id_ENUM_VALUE[valueIdentifier]; ok {
					enumValueString = enumValue.LONG_NAME
				} else {
					log.Panic("ENUM_VALUE_REF", valueIdentifier, "unknown ref in Map_id_ENUM_VALUE")
				}
			}

			*markdownContent += fmt.Sprintf("* **%s:** %s\n", enumTypeString, enumValueString)
		}
	}
}
