package models

func GetAttributeDefinitionIsDisplayed[Attr Attribute](stager *Stager, attribute Attr) (res bool) {

	ref := attribute.GetAttributeDefinitionRef()

	switch any(attribute).(type) {
	case *ATTRIBUTE_VALUE_STRING:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_STRING[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitle[attributeDefinition]
	case *ATTRIBUTE_VALUE_XHTML:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitle[attributeDefinition]
	case *ATTRIBUTE_VALUE_BOOLEAN:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitle[attributeDefinition]
	case *ATTRIBUTE_VALUE_DATE:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_DATE[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitle[attributeDefinition]
	case *ATTRIBUTE_VALUE_REAL:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_REAL[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitle[attributeDefinition]
	case *ATTRIBUTE_VALUE_INTEGER:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitle[attributeDefinition]
	case *ATTRIBUTE_VALUE_ENUMERATION:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitle[attributeDefinition]

	}

	return
}
