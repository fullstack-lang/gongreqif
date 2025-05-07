// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct ALTERNATIVE_ID
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_BOOLEAN
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_BOOLEAN)
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		_ = attribute_definition_boolean
		for _, _alternative_id := range attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_boolean
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*ATTRIBUTE_DEFINITION_BOOLEAN)
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		_ = attribute_definition_boolean
		for _, _attribute_value_boolean := range attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
			stage.ATTRIBUTE_DEFINITION_BOOLEAN_DEFAULT_VALUE_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = attribute_definition_boolean
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_DATE)
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		_ = attribute_definition_date
		for _, _alternative_id := range attribute_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_date
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*ATTRIBUTE_DEFINITION_DATE)
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		_ = attribute_definition_date
		for _, _attribute_value_date := range attribute_definition_date.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE {
			stage.ATTRIBUTE_DEFINITION_DATE_DEFAULT_VALUE_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = attribute_definition_date
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*ATTRIBUTE_DEFINITION_ENUMERATION)
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		_ = attribute_definition_enumeration
		for _, _attribute_value_enumeration := range attribute_definition_enumeration.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION {
			stage.ATTRIBUTE_DEFINITION_ENUMERATION_DEFAULT_VALUE_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = attribute_definition_enumeration
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_ENUMERATION)
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		_ = attribute_definition_enumeration
		for _, _alternative_id := range attribute_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_enumeration
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_INTEGER)
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		_ = attribute_definition_integer
		for _, _alternative_id := range attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_integer
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*ATTRIBUTE_DEFINITION_INTEGER)
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		_ = attribute_definition_integer
		for _, _attribute_value_integer := range attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
			stage.ATTRIBUTE_DEFINITION_INTEGER_DEFAULT_VALUE_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = attribute_definition_integer
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_REAL)
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		_ = attribute_definition_real
		for _, _alternative_id := range attribute_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_real
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*ATTRIBUTE_DEFINITION_REAL)
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		_ = attribute_definition_real
		for _, _attribute_value_real := range attribute_definition_real.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL {
			stage.ATTRIBUTE_DEFINITION_REAL_DEFAULT_VALUE_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = attribute_definition_real
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_STRING)
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		_ = attribute_definition_string
		for _, _alternative_id := range attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_string
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*ATTRIBUTE_DEFINITION_STRING)
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		_ = attribute_definition_string
		for _, _attribute_value_string := range attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
			stage.ATTRIBUTE_DEFINITION_STRING_DEFAULT_VALUE_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = attribute_definition_string
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML
	// insertion point per field
	clear(stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ATTRIBUTE_DEFINITION_XHTML)
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		_ = attribute_definition_xhtml
		for _, _alternative_id := range attribute_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ATTRIBUTE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = attribute_definition_xhtml
		}
	}
	clear(stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*ATTRIBUTE_DEFINITION_XHTML)
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		_ = attribute_definition_xhtml
		for _, _attribute_value_xhtml := range attribute_definition_xhtml.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML {
			stage.ATTRIBUTE_DEFINITION_XHTML_DEFAULT_VALUE_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = attribute_definition_xhtml
		}
	}

	// Compute reverse map for named struct ATTRIBUTE_VALUE_BOOLEAN
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_DATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_INTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_REAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_STRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_XHTML
	// insertion point per field
	clear(stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap)
	stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap = make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		_ = attribute_value_xhtml
		for _, _xhtml_content := range attribute_value_xhtml.THE_VALUE {
			stage.ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap[_xhtml_content] = attribute_value_xhtml
		}
	}
	clear(stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap)
	stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap = make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		_ = attribute_value_xhtml
		for _, _xhtml_content := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
			stage.ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap[_xhtml_content] = attribute_value_xhtml
		}
	}

	// Compute reverse map for named struct AnyType
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_BOOLEAN
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_BOOLEAN)
	for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
		_ = datatype_definition_boolean
		for _, _alternative_id := range datatype_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_BOOLEAN_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_boolean
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_DATE
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_DATE)
	for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
		_ = datatype_definition_date
		for _, _alternative_id := range datatype_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_DATE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_date
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_ENUMERATION
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_ENUMERATION)
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		_ = datatype_definition_enumeration
		for _, _alternative_id := range datatype_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_ENUMERATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_enumeration
		}
	}
	clear(stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_ENUM_VALUE_reverseMap)
	stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_ENUM_VALUE_reverseMap = make(map[*ENUM_VALUE]*DATATYPE_DEFINITION_ENUMERATION)
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		_ = datatype_definition_enumeration
		for _, _enum_value := range datatype_definition_enumeration.SPECIFIED_VALUES.ENUM_VALUE {
			stage.DATATYPE_DEFINITION_ENUMERATION_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[_enum_value] = datatype_definition_enumeration
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_INTEGER
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_INTEGER)
	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		_ = datatype_definition_integer
		for _, _alternative_id := range datatype_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_INTEGER_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_integer
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_REAL
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_REAL)
	for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
		_ = datatype_definition_real
		for _, _alternative_id := range datatype_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_REAL_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_real
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_STRING
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_STRING)
	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		_ = datatype_definition_string
		for _, _alternative_id := range datatype_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_STRING_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_string
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_XHTML
	// insertion point per field
	clear(stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*DATATYPE_DEFINITION_XHTML)
	for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
		_ = datatype_definition_xhtml
		for _, _alternative_id := range datatype_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.DATATYPE_DEFINITION_XHTML_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = datatype_definition_xhtml
		}
	}

	// Compute reverse map for named struct EMBEDDED_VALUE
	// insertion point per field

	// Compute reverse map for named struct ENUM_VALUE
	// insertion point per field
	clear(stage.ENUM_VALUE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.ENUM_VALUE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*ENUM_VALUE)
	for enum_value := range stage.ENUM_VALUEs {
		_ = enum_value
		for _, _alternative_id := range enum_value.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.ENUM_VALUE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = enum_value
		}
	}
	clear(stage.ENUM_VALUE_PROPERTIES_EMBEDDED_VALUE_reverseMap)
	stage.ENUM_VALUE_PROPERTIES_EMBEDDED_VALUE_reverseMap = make(map[*EMBEDDED_VALUE]*ENUM_VALUE)
	for enum_value := range stage.ENUM_VALUEs {
		_ = enum_value
		for _, _embedded_value := range enum_value.PROPERTIES.EMBEDDED_VALUE {
			stage.ENUM_VALUE_PROPERTIES_EMBEDDED_VALUE_reverseMap[_embedded_value] = enum_value
		}
	}

	// Compute reverse map for named struct RELATION_GROUP
	// insertion point per field
	clear(stage.RELATION_GROUP_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.RELATION_GROUP_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*RELATION_GROUP)
	for relation_group := range stage.RELATION_GROUPs {
		_ = relation_group
		for _, _alternative_id := range relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.RELATION_GROUP_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = relation_group
		}
	}

	// Compute reverse map for named struct RELATION_GROUP_TYPE
	// insertion point per field
	clear(stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _alternative_id := range relation_group_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.RELATION_GROUP_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_boolean := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_date := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_enumeration := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_integer := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_real := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_string := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = relation_group_type
		}
	}
	clear(stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*RELATION_GROUP_TYPE)
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		_ = relation_group_type
		for _, _attribute_definition_xhtml := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			stage.RELATION_GROUP_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = relation_group_type
		}
	}

	// Compute reverse map for named struct REQ_IF
	// insertion point per field
	clear(stage.REQ_IF_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap)
	stage.REQ_IF_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap = make(map[*REQ_IF_TOOL_EXTENSION]*REQ_IF)
	for req_if := range stage.REQ_IFs {
		_ = req_if
		for _, _req_if_tool_extension := range req_if.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION {
			stage.REQ_IF_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[_req_if_tool_extension] = req_if
		}
	}

	// Compute reverse map for named struct REQ_IF_CONTENT
	// insertion point per field
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap = make(map[*DATATYPE_DEFINITION_BOOLEAN]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_boolean := range req_if_content.DATATYPES.DATATYPE_DEFINITION_BOOLEAN {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[_datatype_definition_boolean] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap = make(map[*DATATYPE_DEFINITION_DATE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_date := range req_if_content.DATATYPES.DATATYPE_DEFINITION_DATE {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[_datatype_definition_date] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap = make(map[*DATATYPE_DEFINITION_ENUMERATION]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_enumeration := range req_if_content.DATATYPES.DATATYPE_DEFINITION_ENUMERATION {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[_datatype_definition_enumeration] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap = make(map[*DATATYPE_DEFINITION_INTEGER]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_integer := range req_if_content.DATATYPES.DATATYPE_DEFINITION_INTEGER {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[_datatype_definition_integer] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap = make(map[*DATATYPE_DEFINITION_REAL]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_real := range req_if_content.DATATYPES.DATATYPE_DEFINITION_REAL {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[_datatype_definition_real] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap = make(map[*DATATYPE_DEFINITION_STRING]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_string := range req_if_content.DATATYPES.DATATYPE_DEFINITION_STRING {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[_datatype_definition_string] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap)
	stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap = make(map[*DATATYPE_DEFINITION_XHTML]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _datatype_definition_xhtml := range req_if_content.DATATYPES.DATATYPE_DEFINITION_XHTML {
			stage.REQ_IF_CONTENT_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[_datatype_definition_xhtml] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap = make(map[*RELATION_GROUP_TYPE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _relation_group_type := range req_if_content.SPEC_TYPES.RELATION_GROUP_TYPE {
			stage.REQ_IF_CONTENT_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[_relation_group_type] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap = make(map[*SPEC_OBJECT_TYPE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _spec_object_type := range req_if_content.SPEC_TYPES.SPEC_OBJECT_TYPE {
			stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[_spec_object_type] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap = make(map[*SPEC_RELATION_TYPE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _spec_relation_type := range req_if_content.SPEC_TYPES.SPEC_RELATION_TYPE {
			stage.REQ_IF_CONTENT_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[_spec_relation_type] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap = make(map[*SPECIFICATION_TYPE]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _specification_type := range req_if_content.SPEC_TYPES.SPECIFICATION_TYPE {
			stage.REQ_IF_CONTENT_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[_specification_type] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_OBJECTS_SPEC_OBJECT_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_OBJECTS_SPEC_OBJECT_reverseMap = make(map[*SPEC_OBJECT]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _spec_object := range req_if_content.SPEC_OBJECTS.SPEC_OBJECT {
			stage.REQ_IF_CONTENT_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[_spec_object] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_RELATIONS_SPEC_RELATION_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_RELATIONS_SPEC_RELATION_reverseMap = make(map[*SPEC_RELATION]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _spec_relation := range req_if_content.SPEC_RELATIONS.SPEC_RELATION {
			stage.REQ_IF_CONTENT_SPEC_RELATIONS_SPEC_RELATION_reverseMap[_spec_relation] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPECIFICATIONS_SPECIFICATION_reverseMap)
	stage.REQ_IF_CONTENT_SPECIFICATIONS_SPECIFICATION_reverseMap = make(map[*SPECIFICATION]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _specification := range req_if_content.SPECIFICATIONS.SPECIFICATION {
			stage.REQ_IF_CONTENT_SPECIFICATIONS_SPECIFICATION_reverseMap[_specification] = req_if_content
		}
	}
	clear(stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap)
	stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap = make(map[*RELATION_GROUP]*REQ_IF_CONTENT)
	for req_if_content := range stage.REQ_IF_CONTENTs {
		_ = req_if_content
		for _, _relation_group := range req_if_content.SPEC_RELATION_GROUPS.RELATION_GROUP {
			stage.REQ_IF_CONTENT_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[_relation_group] = req_if_content
		}
	}

	// Compute reverse map for named struct REQ_IF_HEADER
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_TOOL_EXTENSION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION
	// insertion point per field
	clear(stage.SPECIFICATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPECIFICATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _alternative_id := range specification.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPECIFICATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_boolean := range specification.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_date := range specification.VALUES.ATTRIBUTE_VALUE_DATE {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_enumeration := range specification.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_integer := range specification.VALUES.ATTRIBUTE_VALUE_INTEGER {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_real := range specification.VALUES.ATTRIBUTE_VALUE_REAL {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_string := range specification.VALUES.ATTRIBUTE_VALUE_STRING {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = specification
		}
	}
	clear(stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _attribute_value_xhtml := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
			stage.SPECIFICATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = specification
		}
	}
	clear(stage.SPECIFICATION_CHILDREN_SPEC_HIERARCHY_reverseMap)
	stage.SPECIFICATION_CHILDREN_SPEC_HIERARCHY_reverseMap = make(map[*SPEC_HIERARCHY]*SPECIFICATION)
	for specification := range stage.SPECIFICATIONs {
		_ = specification
		for _, _spec_hierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
			stage.SPECIFICATION_CHILDREN_SPEC_HIERARCHY_reverseMap[_spec_hierarchy] = specification
		}
	}

	// Compute reverse map for named struct SPECIFICATION_TYPE
	// insertion point per field
	clear(stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _alternative_id := range specification_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPECIFICATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_boolean := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_date := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_enumeration := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_integer := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_real := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_string := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = specification_type
		}
	}
	clear(stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*SPECIFICATION_TYPE)
	for specification_type := range stage.SPECIFICATION_TYPEs {
		_ = specification_type
		for _, _attribute_definition_xhtml := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			stage.SPECIFICATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = specification_type
		}
	}

	// Compute reverse map for named struct SPEC_HIERARCHY
	// insertion point per field
	clear(stage.SPEC_HIERARCHY_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_HIERARCHY_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_HIERARCHY)
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		_ = spec_hierarchy
		for _, _alternative_id := range spec_hierarchy.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_HIERARCHY_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_hierarchy
		}
	}
	clear(stage.SPEC_HIERARCHY_CHILDREN_SPEC_HIERARCHY_reverseMap)
	stage.SPEC_HIERARCHY_CHILDREN_SPEC_HIERARCHY_reverseMap = make(map[*SPEC_HIERARCHY]*SPEC_HIERARCHY)
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		_ = spec_hierarchy
		for _, _spec_hierarchy := range spec_hierarchy.CHILDREN.SPEC_HIERARCHY {
			stage.SPEC_HIERARCHY_CHILDREN_SPEC_HIERARCHY_reverseMap[_spec_hierarchy] = spec_hierarchy
		}
	}

	// Compute reverse map for named struct SPEC_OBJECT
	// insertion point per field
	clear(stage.SPEC_OBJECT_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_OBJECT_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _alternative_id := range spec_object.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_OBJECT_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_boolean := range spec_object.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_date := range spec_object.VALUES.ATTRIBUTE_VALUE_DATE {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_enumeration := range spec_object.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_integer := range spec_object.VALUES.ATTRIBUTE_VALUE_INTEGER {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_real := range spec_object.VALUES.ATTRIBUTE_VALUE_REAL {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_string := range spec_object.VALUES.ATTRIBUTE_VALUE_STRING {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = spec_object
		}
	}
	clear(stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*SPEC_OBJECT)
	for spec_object := range stage.SPEC_OBJECTs {
		_ = spec_object
		for _, _attribute_value_xhtml := range spec_object.VALUES.ATTRIBUTE_VALUE_XHTML {
			stage.SPEC_OBJECT_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = spec_object
		}
	}

	// Compute reverse map for named struct SPEC_OBJECT_TYPE
	// insertion point per field
	clear(stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _alternative_id := range spec_object_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_OBJECT_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_boolean := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_date := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_enumeration := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_integer := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_real := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_string := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = spec_object_type
		}
	}
	clear(stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*SPEC_OBJECT_TYPE)
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		_ = spec_object_type
		for _, _attribute_definition_xhtml := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			stage.SPEC_OBJECT_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = spec_object_type
		}
	}

	// Compute reverse map for named struct SPEC_RELATION
	// insertion point per field
	clear(stage.SPEC_RELATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_RELATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _alternative_id := range spec_relation.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_RELATION_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_boolean := range spec_relation.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_date := range spec_relation.VALUES.ATTRIBUTE_VALUE_DATE {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_enumeration := range spec_relation.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_integer := range spec_relation.VALUES.ATTRIBUTE_VALUE_INTEGER {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_real := range spec_relation.VALUES.ATTRIBUTE_VALUE_REAL {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_string := range spec_relation.VALUES.ATTRIBUTE_VALUE_STRING {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = spec_relation
		}
	}
	clear(stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*SPEC_RELATION)
	for spec_relation := range stage.SPEC_RELATIONs {
		_ = spec_relation
		for _, _attribute_value_xhtml := range spec_relation.VALUES.ATTRIBUTE_VALUE_XHTML {
			stage.SPEC_RELATION_VALUES_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = spec_relation
		}
	}

	// Compute reverse map for named struct SPEC_RELATION_TYPE
	// insertion point per field
	clear(stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap)
	stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap = make(map[*ALTERNATIVE_ID]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _alternative_id := range spec_relation_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
			stage.SPEC_RELATION_TYPE_ALTERNATIVE_ID_ALTERNATIVE_ID_reverseMap[_alternative_id] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_boolean := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_date := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_enumeration := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_integer := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_real := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_string := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = spec_relation_type
		}
	}
	clear(stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*SPEC_RELATION_TYPE)
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		_ = spec_relation_type
		for _, _attribute_definition_xhtml := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			stage.SPEC_RELATION_TYPE_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = spec_relation_type
		}
	}

	// Compute reverse map for named struct XHTML_CONTENT
	// insertion point per field

	// Compute reverse map for named struct Xhtml_InlPres_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_a_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_abbr_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_acronym_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_address_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_blockquote_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_br_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_caption_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_cite_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_code_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_col_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_colgroup_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_dd_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_dfn_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_div_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_dl_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_dt_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_edit_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_em_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h1_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h2_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h3_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h4_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h5_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_h6_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_heading_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_hr_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_kbd_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_li_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_object_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_ol_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_p_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_param_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_pre_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_q_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_samp_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_span_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_strong_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_table_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_tbody_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_td_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_tfoot_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_th_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_thead_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_tr_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_ul_type
	// insertion point per field

	// Compute reverse map for named struct Xhtml_var_type
	// insertion point per field

}
