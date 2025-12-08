package specobjects

import (
	// Corrected path
	"log"

	m "github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/xhtml"
)

type TargetEnum string

const (
	Title   TargetEnum = "Title"
	Subject TargetEnum = "Subject"
)

// FillUpStringWithAttributes
func FillUpStringWithAttributes(stager *m.Stager, specObject *m.SPEC_OBJECT, target TargetEnum) (titleComplement string) {

	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_STRING, target)
	titleComplement += parseXHTMLAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_XHTML, target)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_INTEGER, target)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_REAL, target)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_DATE, target)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_ENUMERATION, target)

	return
}

func parseXHTMLAttributes(stager *m.Stager, attributes []*m.ATTRIBUTE_VALUE_XHTML, target TargetEnum) (attributesString string) {

	for _, attr := range attributes {
		var isIn bool
		switch target {
		case Title:
			isIn = *GetAttributeDefinition(stager, attr).GetShowInTitlePtr()
		case Subject:
			isIn = *GetAttributeDefinition(stager, attr).GetShowInSubjectPtr()
		}

		if isIn {
			if attributesString != "" {
				attributesString += " - "
			}

			// Get the raw XHTML value
			rawXHTML := attr.GetValue()

			// Convert the XHTML string to its Markdown equivalent
			markdown, err := xhtml.ToMarkdown(rawXHTML)
			if err != nil {
				// Log the error and fall back to the raw value to avoid crashing
				log.Printf("WARN: Could not convert XHTML to Markdown for attribute. Error: %v", err)
				attributesString += rawXHTML // Fallback
			} else {
				attributesString += markdown // Use the converted Markdown
			}
		}
	}
	return
}

func parseAttributes[Attr m.Attribute](stager *m.Stager, attributes []Attr, target TargetEnum) (attributesString string) {

	for _, attr := range attributes {

		var isIn bool

		// fetch the definition of the attribute, check if it is to be retained
		// if yes, append to attributesString
		switch target {
		case Title:
			isIn = *GetAttributeDefinition(stager, attr).GetShowInTitlePtr()
		case Subject:
			isIn = *GetAttributeDefinition(stager, attr).GetShowInSubjectPtr()
		}

		if isIn {
			if attributesString != "" {
				attributesString += " - "
			}

			attributesString += attr.GetValue()
		}

	}

	return
}
