package models

func GetAttributeDefinitionIsDisplayedInSubject[Attr Attribute](stager *Stager, attribute Attr) (res bool) {

	ref := attribute.GetAttributeDefinitionRef()

	switch any(attribute).(type) {
	case *ATTRIBUTE_VALUE_STRING:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_STRING[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubject[attributeDefinition]
	case *ATTRIBUTE_VALUE_XHTML:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubject[attributeDefinition]
	case *ATTRIBUTE_VALUE_BOOLEAN:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubject[attributeDefinition]
	case *ATTRIBUTE_VALUE_DATE:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_DATE[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubject[attributeDefinition]
	case *ATTRIBUTE_VALUE_REAL:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_REAL[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubject[attributeDefinition]
	case *ATTRIBUTE_VALUE_INTEGER:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubject[attributeDefinition]
	case *ATTRIBUTE_VALUE_ENUMERATION:
		attributeDefinition := stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[ref]
		res = stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubject[attributeDefinition]

	}

	return
}
