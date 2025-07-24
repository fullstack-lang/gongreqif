package specifications

import (
	// Corrected path
	m "github.com/fullstack-lang/gongreqif/go/models"
)

// fillUpTitleWithAttributes
func fillUpTitleWithAttributes(stager *m.Stager, specObjectType *m.SPEC_OBJECT_TYPE, specObject *m.SPEC_OBJECT) (titleComplement string) {

	titleComplement += parseAttributes(stager, specObject.GetValues().ATTRIBUTE_VALUE_STRING)

	return
}

func parseAttributes[Attr m.Attribute](stager *m.Stager, attributes []Attr) (titleComplement string) {

	for _, attr := range attributes {

		// fetch the definition of the attribute, check if it is to be displayed in the title
		// if yes, append to titleComplement

		isInTitle := m.GetAttributeDefinitionIsDisplayed(stager, attr)

		if isInTitle {
			titleComplement += attr.GetValue()
		}
	}

	return
}
