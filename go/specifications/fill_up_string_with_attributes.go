package specifications

import (
	// Corrected path
	m "github.com/fullstack-lang/gongreqif/go/models"
)

type targetEnum string

const (
	Title   targetEnum = "Title"
	Subject targetEnum = "Subject"
)

// fillUpStringWithAttributes
func fillUpStringWithAttributes(stager *m.Stager, specObject *m.SPEC_OBJECT, target targetEnum) (titleComplement string) {

	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_STRING, target)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_XHTML, target)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_INTEGER, target)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_REAL, target)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_DATE, target)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_ENUMERATION, target)

	return
}

func parseAttributes[Attr m.Attribute](stager *m.Stager, attributes []Attr, target targetEnum) (attributesString string) {

	for _, attr := range attributes {

		var isIn bool

		// fetch the definition of the attribute, check if it is to be retained
		// if yes, append to attributesString
		switch target {
		case Title:
			isIn = m.GetAttributeDefinitionIsDisplayedInTitle(stager, attr)
		case Subject:
			isIn = m.GetAttributeDefinitionIsDisplayedInSubject(stager, attr)
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
