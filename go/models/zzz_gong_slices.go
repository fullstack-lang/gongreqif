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

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML
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
	clear(stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_BOOLEAN)
	for a_attribute_value_boolean := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		_ = a_attribute_value_boolean
		for _, _attribute_value_boolean := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
			stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = a_attribute_value_boolean
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_DATE
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_DATE)
	for a_attribute_value_date := range stage.A_ATTRIBUTE_VALUE_DATEs {
		_ = a_attribute_value_date
		for _, _attribute_value_date := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
			stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = a_attribute_value_date
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_ENUMERATION)
	for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		_ = a_attribute_value_enumeration
		for _, _attribute_value_enumeration := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
			stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = a_attribute_value_enumeration
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_INTEGER
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_INTEGER)
	for a_attribute_value_integer := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		_ = a_attribute_value_integer
		for _, _attribute_value_integer := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
			stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = a_attribute_value_integer
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_REAL
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_REAL)
	for a_attribute_value_real := range stage.A_ATTRIBUTE_VALUE_REALs {
		_ = a_attribute_value_real
		for _, _attribute_value_real := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
			stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = a_attribute_value_real
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_STRING
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_STRING)
	for a_attribute_value_string := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		_ = a_attribute_value_string
		for _, _attribute_value_string := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
			stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = a_attribute_value_string
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML)
	for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		_ = a_attribute_value_xhtml
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
			stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = a_attribute_value_xhtml
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML_1
	// insertion point per field
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_boolean := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_date := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_integer := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_real := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_string := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = a_attribute_value_xhtml_1
		}
	}
	clear(stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap)
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = a_attribute_value_xhtml_1
		}
	}

	// Compute reverse map for named struct A_CHILDREN
	// insertion point per field
	clear(stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap)
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
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap = make(map[*DATATYPE_DEFINITION_BOOLEAN]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
			stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[_datatype_definition_boolean] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap = make(map[*DATATYPE_DEFINITION_DATE]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
			stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[_datatype_definition_date] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap = make(map[*DATATYPE_DEFINITION_ENUMERATION]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
			stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[_datatype_definition_enumeration] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap = make(map[*DATATYPE_DEFINITION_INTEGER]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
			stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[_datatype_definition_integer] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap = make(map[*DATATYPE_DEFINITION_REAL]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
			stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[_datatype_definition_real] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap)
	stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap = make(map[*DATATYPE_DEFINITION_STRING]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
			stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[_datatype_definition_string] = a_datatypes
		}
	}
	clear(stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap)
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
	clear(stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap)
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
	clear(stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap)
	stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap = make(map[*ENUM_VALUE]*A_SPECIFIED_VALUES)
	for a_specified_values := range stage.A_SPECIFIED_VALUESs {
		_ = a_specified_values
		for _, _enum_value := range a_specified_values.ENUM_VALUE {
			stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[_enum_value] = a_specified_values
		}
	}

	// Compute reverse map for named struct A_SPEC_ATTRIBUTES
	// insertion point per field
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = a_spec_attributes
		}
	}
	clear(stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap)
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = a_spec_attributes
		}
	}

	// Compute reverse map for named struct A_SPEC_OBJECTS
	// insertion point per field
	clear(stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap)
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
	clear(stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap)
	stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap = make(map[*SPEC_RELATION]*A_SPEC_RELATIONS)
	for a_spec_relations := range stage.A_SPEC_RELATIONSs {
		_ = a_spec_relations
		for _, _spec_relation := range a_spec_relations.SPEC_RELATION {
			stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap[_spec_relation] = a_spec_relations
		}
	}

	// Compute reverse map for named struct A_SPEC_RELATION_GROUPS
	// insertion point per field
	clear(stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap)
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
	clear(stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap)
	stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap = make(map[*RELATION_GROUP_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
			stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[_relation_group_type] = a_spec_types
		}
	}
	clear(stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap)
	stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap = make(map[*SPEC_OBJECT_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
			stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[_spec_object_type] = a_spec_types
		}
	}
	clear(stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap)
	stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap = make(map[*SPEC_RELATION_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
			stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[_spec_relation_type] = a_spec_types
		}
	}
	clear(stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap)
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
	clear(stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap)
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

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
	// insertion point per field

	// Compute reverse map for named struct Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
	// insertion point per field

	// Compute reverse map for named struct Map_SPECIFICATION_Nodes_expandedEntry
	// insertion point per field

	// Compute reverse map for named struct Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry
	// insertion point per field

	// Compute reverse map for named struct Map_SPEC_OBJECT_TYPE_showIdentifierEntry
	// insertion point per field

	// Compute reverse map for named struct Map_SPEC_OBJECT_TYPE_showNameEntry
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
	clear(stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries_reverseMap)
	stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries_reverseMap = make(map[*Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_spec_object_type_isnodeexpandedentry := range renderingconfiguration.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries {
			stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries_reverseMap[_map_spec_object_type_isnodeexpandedentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_xhtml_showintitleentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries_reverseMap[_map_attribute_definition_xhtml_showintitleentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_string_showintitleentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries_reverseMap[_map_attribute_definition_string_showintitleentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_boolean_showintitleentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries_reverseMap[_map_attribute_definition_boolean_showintitleentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_integer_showintitleentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries_reverseMap[_map_attribute_definition_integer_showintitleentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_date_showintitleentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries_reverseMap[_map_attribute_definition_date_showintitleentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_real_showintitleentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries_reverseMap[_map_attribute_definition_real_showintitleentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_enumeration_showintitleentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries_reverseMap[_map_attribute_definition_enumeration_showintitleentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_xhtml_showintableentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries_reverseMap[_map_attribute_definition_xhtml_showintableentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_string_showintableentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries_reverseMap[_map_attribute_definition_string_showintableentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_boolean_showintableentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries_reverseMap[_map_attribute_definition_boolean_showintableentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_integer_showintableentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries_reverseMap[_map_attribute_definition_integer_showintableentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_date_showintableentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries_reverseMap[_map_attribute_definition_date_showintableentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_real_showintableentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries_reverseMap[_map_attribute_definition_real_showintableentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_enumeration_showintableentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries_reverseMap[_map_attribute_definition_enumeration_showintableentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_xhtml_showinsubjectentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries_reverseMap[_map_attribute_definition_xhtml_showinsubjectentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_string_showinsubjectentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries_reverseMap[_map_attribute_definition_string_showinsubjectentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_boolean_showinsubjectentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries_reverseMap[_map_attribute_definition_boolean_showinsubjectentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_integer_showinsubjectentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries_reverseMap[_map_attribute_definition_integer_showinsubjectentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_date_showinsubjectentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries_reverseMap[_map_attribute_definition_date_showinsubjectentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_real_showinsubjectentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries_reverseMap[_map_attribute_definition_real_showinsubjectentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries_reverseMap)
	stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries_reverseMap = make(map[*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_attribute_definition_enumeration_showinsubjectentry := range renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries {
			stage.RenderingConfiguration_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries_reverseMap[_map_attribute_definition_enumeration_showinsubjectentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_SPECIFICATION_Nodes_expandedEntries_reverseMap)
	stage.RenderingConfiguration_Map_SPECIFICATION_Nodes_expandedEntries_reverseMap = make(map[*Map_SPECIFICATION_Nodes_expandedEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_specification_nodes_expandedentry := range renderingconfiguration.Map_SPECIFICATION_Nodes_expandedEntries {
			stage.RenderingConfiguration_Map_SPECIFICATION_Nodes_expandedEntries_reverseMap[_map_specification_nodes_expandedentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showIdentifierEntries_reverseMap)
	stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showIdentifierEntries_reverseMap = make(map[*Map_SPEC_OBJECT_TYPE_showIdentifierEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_spec_object_type_showidentifierentry := range renderingconfiguration.Map_SPEC_OBJECT_TYPE_showIdentifierEntries {
			stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showIdentifierEntries_reverseMap[_map_spec_object_type_showidentifierentry] = renderingconfiguration
		}
	}
	clear(stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showNameEntries_reverseMap)
	stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showNameEntries_reverseMap = make(map[*Map_SPEC_OBJECT_TYPE_showNameEntry]*RenderingConfiguration)
	for renderingconfiguration := range stage.RenderingConfigurations {
		_ = renderingconfiguration
		for _, _map_spec_object_type_shownameentry := range renderingconfiguration.Map_SPEC_OBJECT_TYPE_showNameEntries {
			stage.RenderingConfiguration_Map_SPEC_OBJECT_TYPE_showNameEntries_reverseMap[_map_spec_object_type_shownameentry] = renderingconfiguration
		}
	}

	// Compute reverse map for named struct SPECIFICATION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_HIERARCHY
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct StaticWebSite
	// insertion point per field
	clear(stage.StaticWebSite_Chapters_reverseMap)
	stage.StaticWebSite_Chapters_reverseMap = make(map[*StaticWebSiteChapter]*StaticWebSite)
	for staticwebsite := range stage.StaticWebSites {
		_ = staticwebsite
		for _, _staticwebsitechapter := range staticwebsite.Chapters {
			stage.StaticWebSite_Chapters_reverseMap[_staticwebsitechapter] = staticwebsite
		}
	}

	// Compute reverse map for named struct StaticWebSiteChapter
	// insertion point per field
	clear(stage.StaticWebSiteChapter_Paragraphs_reverseMap)
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
