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

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_BOOLEAN_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML_Rendering
	// insertion point per field

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

	// Compute reverse map for named struct A_ALTERNATIVE_ID
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_DATE_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_INTEGER_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_REAL_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_STRING_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_XHTML_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_BOOLEAN
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_BOOLEAN)
	for a_attribute_value_boolean := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		_ = a_attribute_value_boolean
		for _, _attribute_value_boolean := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
			stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = a_attribute_value_boolean
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_DATE
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_DATE)
	for a_attribute_value_date := range stage.A_ATTRIBUTE_VALUE_DATEs {
		_ = a_attribute_value_date
		for _, _attribute_value_date := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
			stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = a_attribute_value_date
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_ENUMERATION)
	for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		_ = a_attribute_value_enumeration
		for _, _attribute_value_enumeration := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
			stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = a_attribute_value_enumeration
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_INTEGER
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_INTEGER)
	for a_attribute_value_integer := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		_ = a_attribute_value_integer
		for _, _attribute_value_integer := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
			stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = a_attribute_value_integer
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_REAL
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_REAL)
	for a_attribute_value_real := range stage.A_ATTRIBUTE_VALUE_REALs {
		_ = a_attribute_value_real
		for _, _attribute_value_real := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
			stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = a_attribute_value_real
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_STRING
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_STRING)
	for a_attribute_value_string := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		_ = a_attribute_value_string
		for _, _attribute_value_string := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
			stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = a_attribute_value_string
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML)
	for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		_ = a_attribute_value_xhtml
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
			stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = a_attribute_value_xhtml
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML_1
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_boolean := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_date := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_integer := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_real := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_string := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = a_attribute_value_xhtml_1
		}
	}

	// Compute reverse map for named struct A_CHILDREN
	// insertion point per field
	stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap = make(map[*SPEC_HIERARCHY]*A_CHILDREN)
	for a_children := range stage.A_CHILDRENs {
		_ = a_children
		for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
			stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap[_spec_hierarchy] = a_children
		}
	}

	// Compute reverse map for named struct A_CORE_CONTENT
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPES
	// insertion point per field
	stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap = make(map[*DATATYPE_DEFINITION_BOOLEAN]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
			stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[_datatype_definition_boolean] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap = make(map[*DATATYPE_DEFINITION_DATE]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
			stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[_datatype_definition_date] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap = make(map[*DATATYPE_DEFINITION_ENUMERATION]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
			stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[_datatype_definition_enumeration] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap = make(map[*DATATYPE_DEFINITION_INTEGER]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
			stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[_datatype_definition_integer] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap = make(map[*DATATYPE_DEFINITION_REAL]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
			stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[_datatype_definition_real] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap = make(map[*DATATYPE_DEFINITION_STRING]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
			stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[_datatype_definition_string] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap = make(map[*DATATYPE_DEFINITION_XHTML]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
			stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[_datatype_definition_xhtml] = a_datatypes
		}
	}

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_BOOLEAN_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_DATE_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_ENUMERATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_INTEGER_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_REAL_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_STRING_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_XHTML_REF
	// insertion point per field

	// Compute reverse map for named struct A_EDITABLE_ATTS
	// insertion point per field

	// Compute reverse map for named struct A_ENUM_VALUE_REF
	// insertion point per field

	// Compute reverse map for named struct A_OBJECT
	// insertion point per field

	// Compute reverse map for named struct A_PROPERTIES
	// insertion point per field

	// Compute reverse map for named struct A_RELATION_GROUP_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SOURCE_1
	// insertion point per field

	// Compute reverse map for named struct A_SOURCE_SPECIFICATION_1
	// insertion point per field

	// Compute reverse map for named struct A_SPECIFICATIONS
	// insertion point per field
	stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap = make(map[*SPECIFICATION]*A_SPECIFICATIONS)
	for a_specifications := range stage.A_SPECIFICATIONSs {
		_ = a_specifications
		for _, _specification := range a_specifications.SPECIFICATION {
			stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap[_specification] = a_specifications
		}
	}

	// Compute reverse map for named struct A_SPECIFICATION_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPECIFIED_VALUES
	// insertion point per field
	stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap = make(map[*ENUM_VALUE]*A_SPECIFIED_VALUES)
	for a_specified_values := range stage.A_SPECIFIED_VALUESs {
		_ = a_specified_values
		for _, _enum_value := range a_specified_values.ENUM_VALUE {
			stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[_enum_value] = a_specified_values
		}
	}

	// Compute reverse map for named struct A_SPEC_ATTRIBUTES
	// insertion point per field
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = a_spec_attributes
		}
	}

	// Compute reverse map for named struct A_SPEC_OBJECTS
	// insertion point per field
	stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap = make(map[*SPEC_OBJECT]*A_SPEC_OBJECTS)
	for a_spec_objects := range stage.A_SPEC_OBJECTSs {
		_ = a_spec_objects
		for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
			stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[_spec_object] = a_spec_objects
		}
	}

	// Compute reverse map for named struct A_SPEC_OBJECT_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_RELATIONS
	// insertion point per field
	stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap = make(map[*SPEC_RELATION]*A_SPEC_RELATIONS)
	for a_spec_relations := range stage.A_SPEC_RELATIONSs {
		_ = a_spec_relations
		for _, _spec_relation := range a_spec_relations.SPEC_RELATION {
			stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap[_spec_relation] = a_spec_relations
		}
	}

	// Compute reverse map for named struct A_SPEC_RELATION_GROUPS
	// insertion point per field
	stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap = make(map[*RELATION_GROUP]*A_SPEC_RELATION_GROUPS)
	for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
		_ = a_spec_relation_groups
		for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
			stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[_relation_group] = a_spec_relation_groups
		}
	}

	// Compute reverse map for named struct A_SPEC_RELATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_RELATION_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_TYPES
	// insertion point per field
	stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap = make(map[*RELATION_GROUP_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
			stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[_relation_group_type] = a_spec_types
		}
	}
	stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap = make(map[*SPEC_OBJECT_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
			stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[_spec_object_type] = a_spec_types
		}
	}
	stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap = make(map[*SPEC_RELATION_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
			stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[_spec_relation_type] = a_spec_types
		}
	}
	stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap = make(map[*SPECIFICATION_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
			stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[_specification_type] = a_spec_types
		}
	}

	// Compute reverse map for named struct A_THE_HEADER
	// insertion point per field

	// Compute reverse map for named struct A_TOOL_EXTENSIONS
	// insertion point per field
	stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap = make(map[*REQ_IF_TOOL_EXTENSION]*A_TOOL_EXTENSIONS)
	for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
		_ = a_tool_extensions
		for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
			stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[_req_if_tool_extension] = a_tool_extensions
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_BOOLEAN
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_DATE
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_INTEGER
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_REAL
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_STRING
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_XHTML
	// insertion point per field

	// Compute reverse map for named struct EMBEDDED_VALUE
	// insertion point per field

	// Compute reverse map for named struct ENUM_VALUE
	// insertion point per field

	// Compute reverse map for named struct EmbeddedJpgImage
	// insertion point per field

	// Compute reverse map for named struct EmbeddedPngImage
	// insertion point per field

	// Compute reverse map for named struct EmbeddedSvgImage
	// insertion point per field

	// Compute reverse map for named struct Kill
	// insertion point per field

	// Compute reverse map for named struct Map_identifier_bool
	// insertion point per field

	// Compute reverse map for named struct RELATION_GROUP
	// insertion point per field

	// Compute reverse map for named struct RELATION_GROUP_TYPE
	// insertion point per field

	// Compute reverse map for named struct REQ_IF
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_CONTENT
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_HEADER
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_TOOL_EXTENSION
	// insertion point per field

	// Compute reverse map for named struct RenderingConfiguration
	// insertion point per field
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries_reverseMap = make(map[*Map_identifier_bool]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_identifier_bool := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries_reverseMap[_map_identifier_bool] = renderingconfiguration
		}
	}

	// Compute reverse map for named struct SPECIFICATION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION_Rendering
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_HIERARCHY
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT_TYPE_Rendering
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct StaticWebSite
	// insertion point per field
	stage.StaticWebSite_Chapters_reverseMap = make(map[*StaticWebSiteChapter]*StaticWebSite)
	for staticwebsite := range stage.StaticWebSites {
		_ = staticwebsite
		for _, _staticwebsitechapter := range staticwebsite.Chapters {
			stage.StaticWebSite_Chapters_reverseMap[_staticwebsitechapter] = staticwebsite
		}
	}

	// Compute reverse map for named struct StaticWebSiteChapter
	// insertion point per field
	stage.StaticWebSiteChapter_Paragraphs_reverseMap = make(map[*StaticWebSiteParagraph]*StaticWebSiteChapter)
	for staticwebsitechapter := range stage.StaticWebSiteChapters {
		_ = staticwebsitechapter
		for _, _staticwebsiteparagraph := range staticwebsitechapter.Paragraphs {
			stage.StaticWebSiteChapter_Paragraphs_reverseMap[_staticwebsiteparagraph] = staticwebsitechapter
		}
	}

	// Compute reverse map for named struct StaticWebSiteGeneratedImage
	// insertion point per field

	// Compute reverse map for named struct StaticWebSiteImage
	// insertion point per field

	// Compute reverse map for named struct StaticWebSiteParagraph
	// insertion point per field

	// Compute reverse map for named struct XHTML_CONTENT
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.ALTERNATIVE_IDs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_DATEs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_DATE_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_REALs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_REAL_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_STRING_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_XHTML_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_DATEs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_INTEGERs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_REALs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_STRINGs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_XHTMLs {
		res = append(res, instance)
	}

	for instance := range stage.A_ALTERNATIVE_IDs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_DATEs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_REALs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		res = append(res, instance)
	}

	for instance := range stage.A_CHILDRENs {
		res = append(res, instance)
	}

	for instance := range stage.A_CORE_CONTENTs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPESs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_EDITABLE_ATTSs {
		res = append(res, instance)
	}

	for instance := range stage.A_ENUM_VALUE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_OBJECTs {
		res = append(res, instance)
	}

	for instance := range stage.A_PROPERTIESs {
		res = append(res, instance)
	}

	for instance := range stage.A_RELATION_GROUP_TYPE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SOURCE_1s {
		res = append(res, instance)
	}

	for instance := range stage.A_SOURCE_SPECIFICATION_1s {
		res = append(res, instance)
	}

	for instance := range stage.A_SPECIFICATIONSs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPECIFICATION_TYPE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPECIFIED_VALUESs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_ATTRIBUTESs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_OBJECTSs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_OBJECT_TYPE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_RELATIONSs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_RELATION_GROUPSs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_RELATION_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_RELATION_TYPE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_TYPESs {
		res = append(res, instance)
	}

	for instance := range stage.A_THE_HEADERs {
		res = append(res, instance)
	}

	for instance := range stage.A_TOOL_EXTENSIONSs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_BOOLEANs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_DATEs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_INTEGERs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_REALs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_STRINGs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_XHTMLs {
		res = append(res, instance)
	}

	for instance := range stage.EMBEDDED_VALUEs {
		res = append(res, instance)
	}

	for instance := range stage.ENUM_VALUEs {
		res = append(res, instance)
	}

	for instance := range stage.EmbeddedJpgImages {
		res = append(res, instance)
	}

	for instance := range stage.EmbeddedPngImages {
		res = append(res, instance)
	}

	for instance := range stage.EmbeddedSvgImages {
		res = append(res, instance)
	}

	for instance := range stage.Kills {
		res = append(res, instance)
	}

	for instance := range stage.Map_identifier_bools {
		res = append(res, instance)
	}

	for instance := range stage.RELATION_GROUPs {
		res = append(res, instance)
	}

	for instance := range stage.RELATION_GROUP_TYPEs {
		res = append(res, instance)
	}

	for instance := range stage.REQ_IFs {
		res = append(res, instance)
	}

	for instance := range stage.REQ_IF_CONTENTs {
		res = append(res, instance)
	}

	for instance := range stage.REQ_IF_HEADERs {
		res = append(res, instance)
	}

	for instance := range stage.REQ_IF_TOOL_EXTENSIONs {
		res = append(res, instance)
	}

	for instance := range stage.RenderingConfigurations {
		res = append(res, instance)
	}

	for instance := range stage.SPECIFICATIONs {
		res = append(res, instance)
	}

	for instance := range stage.SPECIFICATION_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.SPECIFICATION_TYPEs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_HIERARCHYs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_OBJECTs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_OBJECT_TYPEs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_OBJECT_TYPE_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_RELATIONs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_RELATION_TYPEs {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSites {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSiteChapters {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSiteGeneratedImages {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSiteImages {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSiteParagraphs {
		res = append(res, instance)
	}

	for instance := range stage.XHTML_CONTENTs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongCopy() GongstructIF {
	newInstance := *alternative_id
	return &newInstance
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongCopy() GongstructIF {
	newInstance := *attribute_definition_boolean
	return &newInstance
}

func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongCopy() GongstructIF {
	newInstance := *attribute_definition_boolean_rendering
	return &newInstance
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongCopy() GongstructIF {
	newInstance := *attribute_definition_date
	return &newInstance
}

func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongCopy() GongstructIF {
	newInstance := *attribute_definition_date_rendering
	return &newInstance
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongCopy() GongstructIF {
	newInstance := *attribute_definition_enumeration
	return &newInstance
}

func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongCopy() GongstructIF {
	newInstance := *attribute_definition_enumeration_rendering
	return &newInstance
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongCopy() GongstructIF {
	newInstance := *attribute_definition_integer
	return &newInstance
}

func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongCopy() GongstructIF {
	newInstance := *attribute_definition_integer_rendering
	return &newInstance
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongCopy() GongstructIF {
	newInstance := *attribute_definition_real
	return &newInstance
}

func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongCopy() GongstructIF {
	newInstance := *attribute_definition_real_rendering
	return &newInstance
}

func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongCopy() GongstructIF {
	newInstance := *attribute_definition_rendering
	return &newInstance
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongCopy() GongstructIF {
	newInstance := *attribute_definition_string
	return &newInstance
}

func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongCopy() GongstructIF {
	newInstance := *attribute_definition_string_rendering
	return &newInstance
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongCopy() GongstructIF {
	newInstance := *attribute_definition_xhtml
	return &newInstance
}

func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongCopy() GongstructIF {
	newInstance := *attribute_definition_xhtml_rendering
	return &newInstance
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongCopy() GongstructIF {
	newInstance := *attribute_value_boolean
	return &newInstance
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongCopy() GongstructIF {
	newInstance := *attribute_value_date
	return &newInstance
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongCopy() GongstructIF {
	newInstance := *attribute_value_enumeration
	return &newInstance
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongCopy() GongstructIF {
	newInstance := *attribute_value_integer
	return &newInstance
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongCopy() GongstructIF {
	newInstance := *attribute_value_real
	return &newInstance
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongCopy() GongstructIF {
	newInstance := *attribute_value_string
	return &newInstance
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongCopy() GongstructIF {
	newInstance := *attribute_value_xhtml
	return &newInstance
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongCopy() GongstructIF {
	newInstance := *a_alternative_id
	return &newInstance
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_boolean_ref
	return &newInstance
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_date_ref
	return &newInstance
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_enumeration_ref
	return &newInstance
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_integer_ref
	return &newInstance
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_real_ref
	return &newInstance
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_string_ref
	return &newInstance
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_xhtml_ref
	return &newInstance
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_boolean
	return &newInstance
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_date
	return &newInstance
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_enumeration
	return &newInstance
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_integer
	return &newInstance
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_real
	return &newInstance
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_string
	return &newInstance
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_xhtml
	return &newInstance
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_xhtml_1
	return &newInstance
}

func (a_children *A_CHILDREN) GongCopy() GongstructIF {
	newInstance := *a_children
	return &newInstance
}

func (a_core_content *A_CORE_CONTENT) GongCopy() GongstructIF {
	newInstance := *a_core_content
	return &newInstance
}

func (a_datatypes *A_DATATYPES) GongCopy() GongstructIF {
	newInstance := *a_datatypes
	return &newInstance
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_boolean_ref
	return &newInstance
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_date_ref
	return &newInstance
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_enumeration_ref
	return &newInstance
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_integer_ref
	return &newInstance
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_real_ref
	return &newInstance
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_string_ref
	return &newInstance
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_xhtml_ref
	return &newInstance
}

func (a_editable_atts *A_EDITABLE_ATTS) GongCopy() GongstructIF {
	newInstance := *a_editable_atts
	return &newInstance
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongCopy() GongstructIF {
	newInstance := *a_enum_value_ref
	return &newInstance
}

func (a_object *A_OBJECT) GongCopy() GongstructIF {
	newInstance := *a_object
	return &newInstance
}

func (a_properties *A_PROPERTIES) GongCopy() GongstructIF {
	newInstance := *a_properties
	return &newInstance
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongCopy() GongstructIF {
	newInstance := *a_relation_group_type_ref
	return &newInstance
}

func (a_source_1 *A_SOURCE_1) GongCopy() GongstructIF {
	newInstance := *a_source_1
	return &newInstance
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongCopy() GongstructIF {
	newInstance := *a_source_specification_1
	return &newInstance
}

func (a_specifications *A_SPECIFICATIONS) GongCopy() GongstructIF {
	newInstance := *a_specifications
	return &newInstance
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongCopy() GongstructIF {
	newInstance := *a_specification_type_ref
	return &newInstance
}

func (a_specified_values *A_SPECIFIED_VALUES) GongCopy() GongstructIF {
	newInstance := *a_specified_values
	return &newInstance
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongCopy() GongstructIF {
	newInstance := *a_spec_attributes
	return &newInstance
}

func (a_spec_objects *A_SPEC_OBJECTS) GongCopy() GongstructIF {
	newInstance := *a_spec_objects
	return &newInstance
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongCopy() GongstructIF {
	newInstance := *a_spec_object_type_ref
	return &newInstance
}

func (a_spec_relations *A_SPEC_RELATIONS) GongCopy() GongstructIF {
	newInstance := *a_spec_relations
	return &newInstance
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongCopy() GongstructIF {
	newInstance := *a_spec_relation_groups
	return &newInstance
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongCopy() GongstructIF {
	newInstance := *a_spec_relation_ref
	return &newInstance
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongCopy() GongstructIF {
	newInstance := *a_spec_relation_type_ref
	return &newInstance
}

func (a_spec_types *A_SPEC_TYPES) GongCopy() GongstructIF {
	newInstance := *a_spec_types
	return &newInstance
}

func (a_the_header *A_THE_HEADER) GongCopy() GongstructIF {
	newInstance := *a_the_header
	return &newInstance
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongCopy() GongstructIF {
	newInstance := *a_tool_extensions
	return &newInstance
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongCopy() GongstructIF {
	newInstance := *datatype_definition_boolean
	return &newInstance
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongCopy() GongstructIF {
	newInstance := *datatype_definition_date
	return &newInstance
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongCopy() GongstructIF {
	newInstance := *datatype_definition_enumeration
	return &newInstance
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongCopy() GongstructIF {
	newInstance := *datatype_definition_integer
	return &newInstance
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongCopy() GongstructIF {
	newInstance := *datatype_definition_real
	return &newInstance
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongCopy() GongstructIF {
	newInstance := *datatype_definition_string
	return &newInstance
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongCopy() GongstructIF {
	newInstance := *datatype_definition_xhtml
	return &newInstance
}

func (embedded_value *EMBEDDED_VALUE) GongCopy() GongstructIF {
	newInstance := *embedded_value
	return &newInstance
}

func (enum_value *ENUM_VALUE) GongCopy() GongstructIF {
	newInstance := *enum_value
	return &newInstance
}

func (embeddedjpgimage *EmbeddedJpgImage) GongCopy() GongstructIF {
	newInstance := *embeddedjpgimage
	return &newInstance
}

func (embeddedpngimage *EmbeddedPngImage) GongCopy() GongstructIF {
	newInstance := *embeddedpngimage
	return &newInstance
}

func (embeddedsvgimage *EmbeddedSvgImage) GongCopy() GongstructIF {
	newInstance := *embeddedsvgimage
	return &newInstance
}

func (kill *Kill) GongCopy() GongstructIF {
	newInstance := *kill
	return &newInstance
}

func (map_identifier_bool *Map_identifier_bool) GongCopy() GongstructIF {
	newInstance := *map_identifier_bool
	return &newInstance
}

func (relation_group *RELATION_GROUP) GongCopy() GongstructIF {
	newInstance := *relation_group
	return &newInstance
}

func (relation_group_type *RELATION_GROUP_TYPE) GongCopy() GongstructIF {
	newInstance := *relation_group_type
	return &newInstance
}

func (req_if *REQ_IF) GongCopy() GongstructIF {
	newInstance := *req_if
	return &newInstance
}

func (req_if_content *REQ_IF_CONTENT) GongCopy() GongstructIF {
	newInstance := *req_if_content
	return &newInstance
}

func (req_if_header *REQ_IF_HEADER) GongCopy() GongstructIF {
	newInstance := *req_if_header
	return &newInstance
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongCopy() GongstructIF {
	newInstance := *req_if_tool_extension
	return &newInstance
}

func (renderingconfiguration *RenderingConfiguration) GongCopy() GongstructIF {
	newInstance := *renderingconfiguration
	return &newInstance
}

func (specification *SPECIFICATION) GongCopy() GongstructIF {
	newInstance := *specification
	return &newInstance
}

func (specification_rendering *SPECIFICATION_Rendering) GongCopy() GongstructIF {
	newInstance := *specification_rendering
	return &newInstance
}

func (specification_type *SPECIFICATION_TYPE) GongCopy() GongstructIF {
	newInstance := *specification_type
	return &newInstance
}

func (spec_hierarchy *SPEC_HIERARCHY) GongCopy() GongstructIF {
	newInstance := *spec_hierarchy
	return &newInstance
}

func (spec_object *SPEC_OBJECT) GongCopy() GongstructIF {
	newInstance := *spec_object
	return &newInstance
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongCopy() GongstructIF {
	newInstance := *spec_object_type
	return &newInstance
}

func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongCopy() GongstructIF {
	newInstance := *spec_object_type_rendering
	return &newInstance
}

func (spec_relation *SPEC_RELATION) GongCopy() GongstructIF {
	newInstance := *spec_relation
	return &newInstance
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongCopy() GongstructIF {
	newInstance := *spec_relation_type
	return &newInstance
}

func (staticwebsite *StaticWebSite) GongCopy() GongstructIF {
	newInstance := *staticwebsite
	return &newInstance
}

func (staticwebsitechapter *StaticWebSiteChapter) GongCopy() GongstructIF {
	newInstance := *staticwebsitechapter
	return &newInstance
}

func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongCopy() GongstructIF {
	newInstance := *staticwebsitegeneratedimage
	return &newInstance
}

func (staticwebsiteimage *StaticWebSiteImage) GongCopy() GongstructIF {
	newInstance := *staticwebsiteimage
	return &newInstance
}

func (staticwebsiteparagraph *StaticWebSiteParagraph) GongCopy() GongstructIF {
	newInstance := *staticwebsiteparagraph
	return &newInstance
}

func (xhtml_content *XHTML_CONTENT) GongCopy() GongstructIF {
	newInstance := *xhtml_content
	return &newInstance
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {
	stage.reference = make(map[GongstructIF]GongstructIF)
	for _, instance := range stage.GetInstances() {
		stage.reference[instance] = instance.GongCopy()
	}
	stage.new = make(map[GongstructIF]struct{})
	stage.modified = make(map[GongstructIF]struct{})
	stage.deleted = make(map[GongstructIF]struct{})
}
