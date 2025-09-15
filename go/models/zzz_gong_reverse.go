// generated code - do not edit
package models

func GetReverseFieldOwnerName(
	stage *Stage,
	instance any,
	reverseField *ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_BOOLEAN":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_DATE":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_ENUMERATION":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_INTEGER":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_REAL":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_STRING":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *ATTRIBUTE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_XHTML":
				if _a_spec_attributes, ok := stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]; ok {
					res = _a_spec_attributes.Name
				}
			}
		}

	case *ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_BOOLEAN":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				if _a_attribute_value_boolean, ok := stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _a_attribute_value_boolean.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_DATE":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				if _a_attribute_value_date, ok := stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _a_attribute_value_date.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_ENUMERATION":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				if _a_attribute_value_enumeration, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _a_attribute_value_enumeration.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_INTEGER":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				if _a_attribute_value_integer, ok := stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _a_attribute_value_integer.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_REAL":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				if _a_attribute_value_real, ok := stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _a_attribute_value_real.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_STRING":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				if _a_attribute_value_string, ok := stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _a_attribute_value_string.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				if _a_attribute_value_xhtml, ok := stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml.Name
				}
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				if _a_attribute_value_xhtml_1, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]; ok {
					res = _a_attribute_value_xhtml_1.Name
				}
			}
		}

	case *A_ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_XHTML_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_CHILDREN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_CORE_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_DATE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_REAL_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_STRING_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_XHTML_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_EDITABLE_ATTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ENUM_VALUE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_PROPERTIES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_RELATION_GROUP_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SOURCE_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SOURCE_SPECIFICATION_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPECIFICATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPECIFICATION_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPECIFIED_VALUES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_ATTRIBUTES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_OBJECTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_OBJECT_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_RELATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_RELATION_GROUPS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_RELATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_RELATION_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_TYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_THE_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_TOOL_EXTENSIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *DATATYPE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_BOOLEAN":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *DATATYPE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_DATE":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *DATATYPE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_ENUMERATION":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *DATATYPE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_INTEGER":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *DATATYPE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_REAL":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *DATATYPE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_STRING":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *DATATYPE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_XHTML":
				if _a_datatypes, ok := stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[inst]; ok {
					res = _a_datatypes.Name
				}
			}
		}

	case *EMBEDDED_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ENUM_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPECIFIED_VALUES":
			switch reverseField.Fieldname {
			case "ENUM_VALUE":
				if _a_specified_values, ok := stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[inst]; ok {
					res = _a_specified_values.Name
				}
			}
		}

	case *Kill:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_SPECIFICATION_Nodes_expandedEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_SPECIFICATION_Nodes_expandedEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_SPECIFICATION_Nodes_expandedEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_SPEC_OBJECT_TYPE_showIdentifierEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showIdentifierEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_SPEC_OBJECT_TYPE_showNameEntries":
				if _renderingconfiguration, ok := stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showNameEntries_reverseMap[inst]; ok {
					res = _renderingconfiguration.Name
				}
			}
		}

	case *RELATION_GROUP:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_RELATION_GROUPS":
			switch reverseField.Fieldname {
			case "RELATION_GROUP":
				if _a_spec_relation_groups, ok := stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[inst]; ok {
					res = _a_spec_relation_groups.Name
				}
			}
		}

	case *RELATION_GROUP_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "RELATION_GROUP_TYPE":
				if _a_spec_types, ok := stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[inst]; ok {
					res = _a_spec_types.Name
				}
			}
		}

	case *REQ_IF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *REQ_IF_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *REQ_IF_TOOL_EXTENSION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_TOOL_EXTENSIONS":
			switch reverseField.Fieldname {
			case "REQ_IF_TOOL_EXTENSION":
				if _a_tool_extensions, ok := stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[inst]; ok {
					res = _a_tool_extensions.Name
				}
			}
		}

	case *RenderingConfiguration:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPECIFICATIONS":
			switch reverseField.Fieldname {
			case "SPECIFICATION":
				if _a_specifications, ok := stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap[inst]; ok {
					res = _a_specifications.Name
				}
			}
		}

	case *SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPECIFICATION_TYPE":
				if _a_spec_types, ok := stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[inst]; ok {
					res = _a_spec_types.Name
				}
			}
		}

	case *SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		case "A_CHILDREN":
			switch reverseField.Fieldname {
			case "SPEC_HIERARCHY":
				if _a_children, ok := stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap[inst]; ok {
					res = _a_children.Name
				}
			}
		}

	case *SPEC_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_OBJECTS":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT":
				if _a_spec_objects, ok := stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[inst]; ok {
					res = _a_spec_objects.Name
				}
			}
		}

	case *SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT_TYPE":
				if _a_spec_types, ok := stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[inst]; ok {
					res = _a_spec_types.Name
				}
			}
		}

	case *SPEC_RELATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_RELATIONS":
			switch reverseField.Fieldname {
			case "SPEC_RELATION":
				if _a_spec_relations, ok := stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap[inst]; ok {
					res = _a_spec_relations.Name
				}
			}
		}

	case *SPEC_RELATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPEC_RELATION_TYPE":
				if _a_spec_types, ok := stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[inst]; ok {
					res = _a_spec_types.Name
				}
			}
		}

	case *StaticWebSite:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *StaticWebSiteChapter:
		switch reverseField.GongstructName {
		// insertion point
		case "StaticWebSite":
			switch reverseField.Fieldname {
			case "Chapters":
				if _staticwebsite, ok := stage.StaticWebSite_Chapters_reverseMap[inst]; ok {
					res = _staticwebsite.Name
				}
			}
		}

	case *StaticWebSiteGeneratedImage:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *StaticWebSiteImage:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *StaticWebSiteParagraph:
		switch reverseField.GongstructName {
		// insertion point
		case "StaticWebSiteChapter":
			switch reverseField.Fieldname {
			case "Paragraphs":
				if _staticwebsitechapter, ok := stage.StaticWebSiteChapter_Paragraphs_reverseMap[inst]; ok {
					res = _staticwebsitechapter.Name
				}
			}
		}

	case *XHTML_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T Gongstruct](
	stage *Stage,
	instance *T,
	reverseField *ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_BOOLEAN":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_DATE":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_ENUMERATION":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_INTEGER":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_REAL":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_STRING":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_ATTRIBUTES":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_DEFINITION_XHTML":
				res = stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_BOOLEAN":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_BOOLEAN":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_DATE":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				res = stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_DATE":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_ENUMERATION":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_ENUMERATION":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_INTEGER":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				res = stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_INTEGER":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_REAL":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				res = stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_REAL":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_STRING":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				res = stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_STRING":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap[inst]
			}
		}

	case *ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_ATTRIBUTE_VALUE_XHTML":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
		case "A_ATTRIBUTE_VALUE_XHTML_1":
			switch reverseField.Fieldname {
			case "ATTRIBUTE_VALUE_XHTML":
				res = stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[inst]
			}
		}

	case *A_ALTERNATIVE_ID:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_DATE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_REAL:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_STRING:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ATTRIBUTE_VALUE_XHTML_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_CHILDREN:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_CORE_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_DATE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_REAL_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_STRING_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_DATATYPE_DEFINITION_XHTML_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_EDITABLE_ATTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_ENUM_VALUE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_PROPERTIES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_RELATION_GROUP_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SOURCE_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SOURCE_SPECIFICATION_1:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPECIFICATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPECIFICATION_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPECIFIED_VALUES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_ATTRIBUTES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_OBJECTS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_OBJECT_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_RELATIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_RELATION_GROUPS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_RELATION_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_RELATION_TYPE_REF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_SPEC_TYPES:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_THE_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *A_TOOL_EXTENSIONS:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *DATATYPE_DEFINITION_BOOLEAN:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_BOOLEAN":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[inst]
			}
		}

	case *DATATYPE_DEFINITION_DATE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_DATE":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[inst]
			}
		}

	case *DATATYPE_DEFINITION_ENUMERATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_ENUMERATION":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[inst]
			}
		}

	case *DATATYPE_DEFINITION_INTEGER:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_INTEGER":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[inst]
			}
		}

	case *DATATYPE_DEFINITION_REAL:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_REAL":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[inst]
			}
		}

	case *DATATYPE_DEFINITION_STRING:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_STRING":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[inst]
			}
		}

	case *DATATYPE_DEFINITION_XHTML:
		switch reverseField.GongstructName {
		// insertion point
		case "A_DATATYPES":
			switch reverseField.Fieldname {
			case "DATATYPE_DEFINITION_XHTML":
				res = stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[inst]
			}
		}

	case *EMBEDDED_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *ENUM_VALUE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPECIFIED_VALUES":
			switch reverseField.Fieldname {
			case "ENUM_VALUE":
				res = stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[inst]
			}
		}

	case *Kill:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries_reverseMap[inst]
			}
		}

	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries":
				res = stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries_reverseMap[inst]
			}
		}

	case *Map_SPECIFICATION_Nodes_expandedEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_SPECIFICATION_Nodes_expandedEntries":
				res = stage.RenderingConfiguration_Map_SPECIFICATION_Nodes_expandedEntries_reverseMap[inst]
			}
		}

	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries":
				res = stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries_reverseMap[inst]
			}
		}

	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_SPEC_OBJECT_TYPE_showIdentifierEntries":
				res = stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showIdentifierEntries_reverseMap[inst]
			}
		}

	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		switch reverseField.GongstructName {
		// insertion point
		case "RenderingConfiguration":
			switch reverseField.Fieldname {
			case "Map_SPEC_OBJECT_TYPE_showNameEntries":
				res = stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showNameEntries_reverseMap[inst]
			}
		}

	case *RELATION_GROUP:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_RELATION_GROUPS":
			switch reverseField.Fieldname {
			case "RELATION_GROUP":
				res = stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[inst]
			}
		}

	case *RELATION_GROUP_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "RELATION_GROUP_TYPE":
				res = stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[inst]
			}
		}

	case *REQ_IF:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *REQ_IF_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *REQ_IF_HEADER:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *REQ_IF_TOOL_EXTENSION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_TOOL_EXTENSIONS":
			switch reverseField.Fieldname {
			case "REQ_IF_TOOL_EXTENSION":
				res = stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[inst]
			}
		}

	case *RenderingConfiguration:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *SPECIFICATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPECIFICATIONS":
			switch reverseField.Fieldname {
			case "SPECIFICATION":
				res = stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap[inst]
			}
		}

	case *SPECIFICATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPECIFICATION_TYPE":
				res = stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[inst]
			}
		}

	case *SPEC_HIERARCHY:
		switch reverseField.GongstructName {
		// insertion point
		case "A_CHILDREN":
			switch reverseField.Fieldname {
			case "SPEC_HIERARCHY":
				res = stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap[inst]
			}
		}

	case *SPEC_OBJECT:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_OBJECTS":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT":
				res = stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[inst]
			}
		}

	case *SPEC_OBJECT_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPEC_OBJECT_TYPE":
				res = stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[inst]
			}
		}

	case *SPEC_RELATION:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_RELATIONS":
			switch reverseField.Fieldname {
			case "SPEC_RELATION":
				res = stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap[inst]
			}
		}

	case *SPEC_RELATION_TYPE:
		switch reverseField.GongstructName {
		// insertion point
		case "A_SPEC_TYPES":
			switch reverseField.Fieldname {
			case "SPEC_RELATION_TYPE":
				res = stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[inst]
			}
		}

	case *StaticWebSite:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *StaticWebSiteChapter:
		switch reverseField.GongstructName {
		// insertion point
		case "StaticWebSite":
			switch reverseField.Fieldname {
			case "Chapters":
				res = stage.StaticWebSite_Chapters_reverseMap[inst]
			}
		}

	case *StaticWebSiteGeneratedImage:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *StaticWebSiteImage:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *StaticWebSiteParagraph:
		switch reverseField.GongstructName {
		// insertion point
		case "StaticWebSiteChapter":
			switch reverseField.Fieldname {
			case "Paragraphs":
				res = stage.StaticWebSiteChapter_Paragraphs_reverseMap[inst]
			}
		}

	case *XHTML_CONTENT:
		switch reverseField.GongstructName {
		// insertion point
		}

	default:
		_ = inst
	}
	return res
}
