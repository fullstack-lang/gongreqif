package models

func GetAttributeDefinitionIsDisplayed[Attr Attribute](stager *Stager, attribute Attr) (res bool) {

	ref := attribute.GetAttributeDefinitionRef()

	switch any(attribute).(type) {
	case *ATTRIBUTE_VALUE_STRING:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_STRING[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitle[attributeDefinition]

	}

	return
}
