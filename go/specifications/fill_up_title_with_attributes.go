package specifications

import (
	// Corrected path
	m "github.com/fullstack-lang/gongreqif/go/models"
)

// fillUpTitleWithAttributes
func fillUpTitleWithAttributes(stager *m.Stager, specObject *m.SPEC_OBJECT) (titleComplement string) {

	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_STRING)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_XHTML)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_INTEGER)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_REAL)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_DATE)
	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_ENUMERATION)

	return
}

func parseAttributes[Attr m.Attribute](stager *m.Stager, attributes []Attr) (titleComplement string) {

	for _, attr := range attributes {

		// fetch the definition of the attribute, check if it is to be displayed in the title
		// if yes, append to titleComplement

		isInTitle := m.GetAttributeDefinitionIsDisplayed(stager, attr)

		if isInTitle {
			if titleComplement != "" {
				titleComplement += " - "
			}

			titleComplement += attr.GetValue()
		}
	}

	return
}
