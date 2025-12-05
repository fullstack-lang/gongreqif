// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}

	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by ALTERNATIVE_ID
func (alternative_id *ALTERNATIVE_ID) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_BOOLEAN
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_definition_boolean.ALTERNATIVE_ID = GongCleanPointer(stage, attribute_definition_boolean.ALTERNATIVE_ID)
	attribute_definition_boolean.DEFAULT_VALUE = GongCleanPointer(stage, attribute_definition_boolean.DEFAULT_VALUE)
	attribute_definition_boolean.TYPE = GongCleanPointer(stage, attribute_definition_boolean.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_DATE
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_definition_date.ALTERNATIVE_ID = GongCleanPointer(stage, attribute_definition_date.ALTERNATIVE_ID)
	attribute_definition_date.DEFAULT_VALUE = GongCleanPointer(stage, attribute_definition_date.DEFAULT_VALUE)
	attribute_definition_date.TYPE = GongCleanPointer(stage, attribute_definition_date.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_ENUMERATION
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_definition_enumeration.ALTERNATIVE_ID = GongCleanPointer(stage, attribute_definition_enumeration.ALTERNATIVE_ID)
	attribute_definition_enumeration.DEFAULT_VALUE = GongCleanPointer(stage, attribute_definition_enumeration.DEFAULT_VALUE)
	attribute_definition_enumeration.TYPE = GongCleanPointer(stage, attribute_definition_enumeration.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_INTEGER
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_definition_integer.ALTERNATIVE_ID = GongCleanPointer(stage, attribute_definition_integer.ALTERNATIVE_ID)
	attribute_definition_integer.DEFAULT_VALUE = GongCleanPointer(stage, attribute_definition_integer.DEFAULT_VALUE)
	attribute_definition_integer.TYPE = GongCleanPointer(stage, attribute_definition_integer.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_REAL
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_definition_real.ALTERNATIVE_ID = GongCleanPointer(stage, attribute_definition_real.ALTERNATIVE_ID)
	attribute_definition_real.DEFAULT_VALUE = GongCleanPointer(stage, attribute_definition_real.DEFAULT_VALUE)
	attribute_definition_real.TYPE = GongCleanPointer(stage, attribute_definition_real.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_STRING
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_definition_string.ALTERNATIVE_ID = GongCleanPointer(stage, attribute_definition_string.ALTERNATIVE_ID)
	attribute_definition_string.DEFAULT_VALUE = GongCleanPointer(stage, attribute_definition_string.DEFAULT_VALUE)
	attribute_definition_string.TYPE = GongCleanPointer(stage, attribute_definition_string.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_XHTML
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_definition_xhtml.ALTERNATIVE_ID = GongCleanPointer(stage, attribute_definition_xhtml.ALTERNATIVE_ID)
	attribute_definition_xhtml.DEFAULT_VALUE = GongCleanPointer(stage, attribute_definition_xhtml.DEFAULT_VALUE)
	attribute_definition_xhtml.TYPE = GongCleanPointer(stage, attribute_definition_xhtml.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_BOOLEAN
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_value_boolean.DEFINITION = GongCleanPointer(stage, attribute_value_boolean.DEFINITION)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_DATE
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_value_date.DEFINITION = GongCleanPointer(stage, attribute_value_date.DEFINITION)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_ENUMERATION
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_value_enumeration.DEFINITION = GongCleanPointer(stage, attribute_value_enumeration.DEFINITION)
	attribute_value_enumeration.VALUES = GongCleanPointer(stage, attribute_value_enumeration.VALUES)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_INTEGER
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_value_integer.DEFINITION = GongCleanPointer(stage, attribute_value_integer.DEFINITION)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_REAL
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_value_real.DEFINITION = GongCleanPointer(stage, attribute_value_real.DEFINITION)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_STRING
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_value_string.DEFINITION = GongCleanPointer(stage, attribute_value_string.DEFINITION)
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_XHTML
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	attribute_value_xhtml.DEFINITION = GongCleanPointer(stage, attribute_value_xhtml.DEFINITION)
	attribute_value_xhtml.THE_VALUE = GongCleanPointer(stage, attribute_value_xhtml.THE_VALUE)
	attribute_value_xhtml.THE_ORIGINAL_VALUE = GongCleanPointer(stage, attribute_value_xhtml.THE_ORIGINAL_VALUE)
}

// Clean garbage collect unstaged instances that are referenced by A_ALTERNATIVE_ID
func (a_alternative_id *A_ALTERNATIVE_ID) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	a_alternative_id.ALTERNATIVE_ID = GongCleanPointer(stage, a_alternative_id.ALTERNATIVE_ID)
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_DATE_REF
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_INTEGER_REF
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_REAL_REF
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_STRING_REF
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_XHTML_REF
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_BOOLEAN
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongClean(stage *Stage) {
	// insertion point per field
	a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN = GongCleanSlice(stage, a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_DATE
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongClean(stage *Stage) {
	// insertion point per field
	a_attribute_value_date.ATTRIBUTE_VALUE_DATE = GongCleanSlice(stage, a_attribute_value_date.ATTRIBUTE_VALUE_DATE)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_ENUMERATION
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongClean(stage *Stage) {
	// insertion point per field
	a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION = GongCleanSlice(stage, a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_INTEGER
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongClean(stage *Stage) {
	// insertion point per field
	a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER = GongCleanSlice(stage, a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_REAL
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongClean(stage *Stage) {
	// insertion point per field
	a_attribute_value_real.ATTRIBUTE_VALUE_REAL = GongCleanSlice(stage, a_attribute_value_real.ATTRIBUTE_VALUE_REAL)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_STRING
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongClean(stage *Stage) {
	// insertion point per field
	a_attribute_value_string.ATTRIBUTE_VALUE_STRING = GongCleanSlice(stage, a_attribute_value_string.ATTRIBUTE_VALUE_STRING)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_XHTML
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongClean(stage *Stage) {
	// insertion point per field
	a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML = GongCleanSlice(stage, a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_XHTML_1
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongClean(stage *Stage) {
	// insertion point per field
	a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN = GongCleanSlice(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN)
	a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE = GongCleanSlice(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE)
	a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION = GongCleanSlice(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION)
	a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER = GongCleanSlice(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER)
	a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL = GongCleanSlice(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL)
	a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING = GongCleanSlice(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING)
	a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML = GongCleanSlice(stage, a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_CHILDREN
func (a_children *A_CHILDREN) GongClean(stage *Stage) {
	// insertion point per field
	a_children.SPEC_HIERARCHY = GongCleanSlice(stage, a_children.SPEC_HIERARCHY)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_CORE_CONTENT
func (a_core_content *A_CORE_CONTENT) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	a_core_content.REQ_IF_CONTENT = GongCleanPointer(stage, a_core_content.REQ_IF_CONTENT)
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPES
func (a_datatypes *A_DATATYPES) GongClean(stage *Stage) {
	// insertion point per field
	a_datatypes.DATATYPE_DEFINITION_BOOLEAN = GongCleanSlice(stage, a_datatypes.DATATYPE_DEFINITION_BOOLEAN)
	a_datatypes.DATATYPE_DEFINITION_DATE = GongCleanSlice(stage, a_datatypes.DATATYPE_DEFINITION_DATE)
	a_datatypes.DATATYPE_DEFINITION_ENUMERATION = GongCleanSlice(stage, a_datatypes.DATATYPE_DEFINITION_ENUMERATION)
	a_datatypes.DATATYPE_DEFINITION_INTEGER = GongCleanSlice(stage, a_datatypes.DATATYPE_DEFINITION_INTEGER)
	a_datatypes.DATATYPE_DEFINITION_REAL = GongCleanSlice(stage, a_datatypes.DATATYPE_DEFINITION_REAL)
	a_datatypes.DATATYPE_DEFINITION_STRING = GongCleanSlice(stage, a_datatypes.DATATYPE_DEFINITION_STRING)
	a_datatypes.DATATYPE_DEFINITION_XHTML = GongCleanSlice(stage, a_datatypes.DATATYPE_DEFINITION_XHTML)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_BOOLEAN_REF
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_DATE_REF
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_ENUMERATION_REF
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_INTEGER_REF
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_REAL_REF
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_STRING_REF
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_XHTML_REF
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_EDITABLE_ATTS
func (a_editable_atts *A_EDITABLE_ATTS) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_ENUM_VALUE_REF
func (a_enum_value_ref *A_ENUM_VALUE_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_OBJECT
func (a_object *A_OBJECT) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_PROPERTIES
func (a_properties *A_PROPERTIES) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	a_properties.EMBEDDED_VALUE = GongCleanPointer(stage, a_properties.EMBEDDED_VALUE)
}

// Clean garbage collect unstaged instances that are referenced by A_RELATION_GROUP_TYPE_REF
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SOURCE_1
func (a_source_1 *A_SOURCE_1) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SOURCE_SPECIFICATION_1
func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPECIFICATIONS
func (a_specifications *A_SPECIFICATIONS) GongClean(stage *Stage) {
	// insertion point per field
	a_specifications.SPECIFICATION = GongCleanSlice(stage, a_specifications.SPECIFICATION)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPECIFICATION_TYPE_REF
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPECIFIED_VALUES
func (a_specified_values *A_SPECIFIED_VALUES) GongClean(stage *Stage) {
	// insertion point per field
	a_specified_values.ENUM_VALUE = GongCleanSlice(stage, a_specified_values.ENUM_VALUE)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_ATTRIBUTES
func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongClean(stage *Stage) {
	// insertion point per field
	a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN = GongCleanSlice(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN)
	a_spec_attributes.ATTRIBUTE_DEFINITION_DATE = GongCleanSlice(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_DATE)
	a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION = GongCleanSlice(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION)
	a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER = GongCleanSlice(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER)
	a_spec_attributes.ATTRIBUTE_DEFINITION_REAL = GongCleanSlice(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_REAL)
	a_spec_attributes.ATTRIBUTE_DEFINITION_STRING = GongCleanSlice(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_STRING)
	a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML = GongCleanSlice(stage, a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_OBJECTS
func (a_spec_objects *A_SPEC_OBJECTS) GongClean(stage *Stage) {
	// insertion point per field
	a_spec_objects.SPEC_OBJECT = GongCleanSlice(stage, a_spec_objects.SPEC_OBJECT)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_OBJECT_TYPE_REF
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_RELATIONS
func (a_spec_relations *A_SPEC_RELATIONS) GongClean(stage *Stage) {
	// insertion point per field
	a_spec_relations.SPEC_RELATION = GongCleanSlice(stage, a_spec_relations.SPEC_RELATION)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_RELATION_GROUPS
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongClean(stage *Stage) {
	// insertion point per field
	a_spec_relation_groups.RELATION_GROUP = GongCleanSlice(stage, a_spec_relation_groups.RELATION_GROUP)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_RELATION_REF
func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_RELATION_TYPE_REF
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_TYPES
func (a_spec_types *A_SPEC_TYPES) GongClean(stage *Stage) {
	// insertion point per field
	a_spec_types.RELATION_GROUP_TYPE = GongCleanSlice(stage, a_spec_types.RELATION_GROUP_TYPE)
	a_spec_types.SPEC_OBJECT_TYPE = GongCleanSlice(stage, a_spec_types.SPEC_OBJECT_TYPE)
	a_spec_types.SPEC_RELATION_TYPE = GongCleanSlice(stage, a_spec_types.SPEC_RELATION_TYPE)
	a_spec_types.SPECIFICATION_TYPE = GongCleanSlice(stage, a_spec_types.SPECIFICATION_TYPE)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by A_THE_HEADER
func (a_the_header *A_THE_HEADER) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	a_the_header.REQ_IF_HEADER = GongCleanPointer(stage, a_the_header.REQ_IF_HEADER)
}

// Clean garbage collect unstaged instances that are referenced by A_TOOL_EXTENSIONS
func (a_tool_extensions *A_TOOL_EXTENSIONS) GongClean(stage *Stage) {
	// insertion point per field
	a_tool_extensions.REQ_IF_TOOL_EXTENSION = GongCleanSlice(stage, a_tool_extensions.REQ_IF_TOOL_EXTENSION)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_BOOLEAN
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	datatype_definition_boolean.ALTERNATIVE_ID = GongCleanPointer(stage, datatype_definition_boolean.ALTERNATIVE_ID)
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_DATE
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	datatype_definition_date.ALTERNATIVE_ID = GongCleanPointer(stage, datatype_definition_date.ALTERNATIVE_ID)
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_ENUMERATION
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	datatype_definition_enumeration.ALTERNATIVE_ID = GongCleanPointer(stage, datatype_definition_enumeration.ALTERNATIVE_ID)
	datatype_definition_enumeration.SPECIFIED_VALUES = GongCleanPointer(stage, datatype_definition_enumeration.SPECIFIED_VALUES)
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_INTEGER
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	datatype_definition_integer.ALTERNATIVE_ID = GongCleanPointer(stage, datatype_definition_integer.ALTERNATIVE_ID)
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_REAL
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	datatype_definition_real.ALTERNATIVE_ID = GongCleanPointer(stage, datatype_definition_real.ALTERNATIVE_ID)
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_STRING
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	datatype_definition_string.ALTERNATIVE_ID = GongCleanPointer(stage, datatype_definition_string.ALTERNATIVE_ID)
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_XHTML
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	datatype_definition_xhtml.ALTERNATIVE_ID = GongCleanPointer(stage, datatype_definition_xhtml.ALTERNATIVE_ID)
}

// Clean garbage collect unstaged instances that are referenced by EMBEDDED_VALUE
func (embedded_value *EMBEDDED_VALUE) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by ENUM_VALUE
func (enum_value *ENUM_VALUE) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	enum_value.ALTERNATIVE_ID = GongCleanPointer(stage, enum_value.ALTERNATIVE_ID)
	enum_value.PROPERTIES = GongCleanPointer(stage, enum_value.PROPERTIES)
}

// Clean garbage collect unstaged instances that are referenced by EmbeddedJpgImage
func (embeddedjpgimage *EmbeddedJpgImage) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by EmbeddedPngImage
func (embeddedpngimage *EmbeddedPngImage) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by EmbeddedSvgImage
func (embeddedsvgimage *EmbeddedSvgImage) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Kill
func (kill *Kill) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry
func (map_attribute_definition_boolean_showinsubjectentry *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry
func (map_attribute_definition_boolean_showintableentry *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry
func (map_attribute_definition_boolean_showintitleentry *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry
func (map_attribute_definition_date_showinsubjectentry *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry
func (map_attribute_definition_date_showintableentry *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry
func (map_attribute_definition_date_showintitleentry *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
func (map_attribute_definition_enumeration_showinsubjectentry *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry
func (map_attribute_definition_enumeration_showintableentry *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
func (map_attribute_definition_enumeration_showintitleentry *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry
func (map_attribute_definition_integer_showinsubjectentry *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry
func (map_attribute_definition_integer_showintableentry *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry
func (map_attribute_definition_integer_showintitleentry *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry
func (map_attribute_definition_real_showinsubjectentry *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry
func (map_attribute_definition_real_showintableentry *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry
func (map_attribute_definition_real_showintitleentry *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry
func (map_attribute_definition_string_showinsubjectentry *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry
func (map_attribute_definition_string_showintableentry *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry
func (map_attribute_definition_string_showintitleentry *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry
func (map_attribute_definition_xhtml_showinsubjectentry *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
func (map_attribute_definition_xhtml_showintableentry *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
func (map_attribute_definition_xhtml_showintitleentry *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_SPECIFICATION_Nodes_expandedEntry
func (map_specification_nodes_expandedentry *Map_SPECIFICATION_Nodes_expandedEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry
func (map_spec_object_type_isnodeexpandedentry *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_SPEC_OBJECT_TYPE_showIdentifierEntry
func (map_spec_object_type_showidentifierentry *Map_SPEC_OBJECT_TYPE_showIdentifierEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_SPEC_OBJECT_TYPE_showNameEntry
func (map_spec_object_type_shownameentry *Map_SPEC_OBJECT_TYPE_showNameEntry) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Map_SPEC_OBJECT_TYPE_showRelations
func (map_spec_object_type_showrelations *Map_SPEC_OBJECT_TYPE_showRelations) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by RELATION_GROUP
func (relation_group *RELATION_GROUP) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	relation_group.ALTERNATIVE_ID = GongCleanPointer(stage, relation_group.ALTERNATIVE_ID)
	relation_group.SOURCE_SPECIFICATION = GongCleanPointer(stage, relation_group.SOURCE_SPECIFICATION)
	relation_group.SPEC_RELATIONS = GongCleanPointer(stage, relation_group.SPEC_RELATIONS)
	relation_group.TARGET_SPECIFICATION = GongCleanPointer(stage, relation_group.TARGET_SPECIFICATION)
	relation_group.TYPE = GongCleanPointer(stage, relation_group.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by RELATION_GROUP_TYPE
func (relation_group_type *RELATION_GROUP_TYPE) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	relation_group_type.ALTERNATIVE_ID = GongCleanPointer(stage, relation_group_type.ALTERNATIVE_ID)
	relation_group_type.SPEC_ATTRIBUTES = GongCleanPointer(stage, relation_group_type.SPEC_ATTRIBUTES)
}

// Clean garbage collect unstaged instances that are referenced by REQ_IF
func (req_if *REQ_IF) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	req_if.THE_HEADER = GongCleanPointer(stage, req_if.THE_HEADER)
	req_if.CORE_CONTENT = GongCleanPointer(stage, req_if.CORE_CONTENT)
	req_if.TOOL_EXTENSIONS = GongCleanPointer(stage, req_if.TOOL_EXTENSIONS)
}

// Clean garbage collect unstaged instances that are referenced by REQ_IF_CONTENT
func (req_if_content *REQ_IF_CONTENT) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	req_if_content.DATATYPES = GongCleanPointer(stage, req_if_content.DATATYPES)
	req_if_content.SPEC_TYPES = GongCleanPointer(stage, req_if_content.SPEC_TYPES)
	req_if_content.SPEC_OBJECTS = GongCleanPointer(stage, req_if_content.SPEC_OBJECTS)
	req_if_content.SPEC_RELATIONS = GongCleanPointer(stage, req_if_content.SPEC_RELATIONS)
	req_if_content.SPECIFICATIONS = GongCleanPointer(stage, req_if_content.SPECIFICATIONS)
	req_if_content.SPEC_RELATION_GROUPS = GongCleanPointer(stage, req_if_content.SPEC_RELATION_GROUPS)
}

// Clean garbage collect unstaged instances that are referenced by REQ_IF_HEADER
func (req_if_header *REQ_IF_HEADER) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by REQ_IF_TOOL_EXTENSION
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by RenderingConfiguration
func (renderingconfiguration *RenderingConfiguration) GongClean(stage *Stage) {
	// insertion point per field
	renderingconfiguration.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries = GongCleanSlice(stage, renderingconfiguration.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntries)
	renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries = GongCleanSlice(stage, renderingconfiguration.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntries)
	renderingconfiguration.Map_SPECIFICATION_Nodes_expandedEntries = GongCleanSlice(stage, renderingconfiguration.Map_SPECIFICATION_Nodes_expandedEntries)
	renderingconfiguration.Map_SPEC_OBJECT_TYPE_showIdentifierEntries = GongCleanSlice(stage, renderingconfiguration.Map_SPEC_OBJECT_TYPE_showIdentifierEntries)
	renderingconfiguration.Map_SPEC_OBJECT_TYPE_showNameEntries = GongCleanSlice(stage, renderingconfiguration.Map_SPEC_OBJECT_TYPE_showNameEntries)
	renderingconfiguration.Map_SPEC_OBJECT_TYPE_showRelations = GongCleanSlice(stage, renderingconfiguration.Map_SPEC_OBJECT_TYPE_showRelations)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by SPECIFICATION
func (specification *SPECIFICATION) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	specification.ALTERNATIVE_ID = GongCleanPointer(stage, specification.ALTERNATIVE_ID)
	specification.TYPE = GongCleanPointer(stage, specification.TYPE)
	specification.CHILDREN = GongCleanPointer(stage, specification.CHILDREN)
	specification.VALUES = GongCleanPointer(stage, specification.VALUES)
}

// Clean garbage collect unstaged instances that are referenced by SPECIFICATION_TYPE
func (specification_type *SPECIFICATION_TYPE) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	specification_type.ALTERNATIVE_ID = GongCleanPointer(stage, specification_type.ALTERNATIVE_ID)
	specification_type.SPEC_ATTRIBUTES = GongCleanPointer(stage, specification_type.SPEC_ATTRIBUTES)
}

// Clean garbage collect unstaged instances that are referenced by SPEC_HIERARCHY
func (spec_hierarchy *SPEC_HIERARCHY) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spec_hierarchy.ALTERNATIVE_ID = GongCleanPointer(stage, spec_hierarchy.ALTERNATIVE_ID)
	spec_hierarchy.OBJECT = GongCleanPointer(stage, spec_hierarchy.OBJECT)
	spec_hierarchy.CHILDREN = GongCleanPointer(stage, spec_hierarchy.CHILDREN)
	spec_hierarchy.EDITABLE_ATTS = GongCleanPointer(stage, spec_hierarchy.EDITABLE_ATTS)
}

// Clean garbage collect unstaged instances that are referenced by SPEC_OBJECT
func (spec_object *SPEC_OBJECT) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spec_object.ALTERNATIVE_ID = GongCleanPointer(stage, spec_object.ALTERNATIVE_ID)
	spec_object.VALUES = GongCleanPointer(stage, spec_object.VALUES)
	spec_object.TYPE = GongCleanPointer(stage, spec_object.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by SPEC_OBJECT_TYPE
func (spec_object_type *SPEC_OBJECT_TYPE) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spec_object_type.ALTERNATIVE_ID = GongCleanPointer(stage, spec_object_type.ALTERNATIVE_ID)
	spec_object_type.SPEC_ATTRIBUTES = GongCleanPointer(stage, spec_object_type.SPEC_ATTRIBUTES)
}

// Clean garbage collect unstaged instances that are referenced by SPEC_RELATION
func (spec_relation *SPEC_RELATION) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spec_relation.ALTERNATIVE_ID = GongCleanPointer(stage, spec_relation.ALTERNATIVE_ID)
	spec_relation.VALUES = GongCleanPointer(stage, spec_relation.VALUES)
	spec_relation.SOURCE = GongCleanPointer(stage, spec_relation.SOURCE)
	spec_relation.TARGET = GongCleanPointer(stage, spec_relation.TARGET)
	spec_relation.TYPE = GongCleanPointer(stage, spec_relation.TYPE)
}

// Clean garbage collect unstaged instances that are referenced by SPEC_RELATION_TYPE
func (spec_relation_type *SPEC_RELATION_TYPE) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	spec_relation_type.ALTERNATIVE_ID = GongCleanPointer(stage, spec_relation_type.ALTERNATIVE_ID)
	spec_relation_type.SPEC_ATTRIBUTES = GongCleanPointer(stage, spec_relation_type.SPEC_ATTRIBUTES)
}

// Clean garbage collect unstaged instances that are referenced by StaticWebSite
func (staticwebsite *StaticWebSite) GongClean(stage *Stage) {
	// insertion point per field
	staticwebsite.Chapters = GongCleanSlice(stage, staticwebsite.Chapters)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by StaticWebSiteChapter
func (staticwebsitechapter *StaticWebSiteChapter) GongClean(stage *Stage) {
	// insertion point per field
	staticwebsitechapter.Paragraphs = GongCleanSlice(stage, staticwebsitechapter.Paragraphs)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by StaticWebSiteGeneratedImage
func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by StaticWebSiteImage
func (staticwebsiteimage *StaticWebSiteImage) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by StaticWebSiteParagraph
func (staticwebsiteparagraph *StaticWebSiteParagraph) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	staticwebsiteparagraph.Image = GongCleanPointer(stage, staticwebsiteparagraph.Image)
}

// Clean garbage collect unstaged instances that are referenced by XHTML_CONTENT
func (xhtml_content *XHTML_CONTENT) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
