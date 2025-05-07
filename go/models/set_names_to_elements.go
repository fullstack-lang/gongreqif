package models

// SetNamesToREQIF_Elements fill up names of reqif objects
func SetNamesToREQIF_Elements(stage *Stage) {

	for x := range *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_ENUMERATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_XHTML](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_ENUMERATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_XHTML](stage) {
		x.Name = x.LONG_NAME
	}

	// Attribute Value

	// for x := range *GetGongstructInstancesSet[ATTRIBUTE_VALUE_ENUMERATION](stage) {
	// 	x.Name = x.DEFINITION
	// }

	// anonymous without ATTRIBUTE DEF

	for x := range *GetGongstructInstancesSet[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](stage) {
		x.Name = x.ATTRIBUTE_DEFINITION_ENUMERATION_REF
	}

	for x := range *GetGongstructInstancesSet[A_ATTRIBUTE_DEFINITION_XHTML_REF](stage) {
		x.Name = x.ATTRIBUTE_DEFINITION_XHTML_REF
	}

	// for x := range *GetGongstructInstancesSet[A_ATTRIBUTE_VALUE_XHTML_1](stage) {
	// 	if x.
	// 	x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	// }

	for x := range *GetGongstructInstancesSet[A_DATATYPE_DEFINITION_ENUMERATION_REF](stage) {
		x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	}

	for x := range *GetGongstructInstancesSet[A_DATATYPE_DEFINITION_XHTML_REF](stage) {
		x.Name = x.DATATYPE_DEFINITION_XHTML_REF
	}

	for x := range *GetGongstructInstancesSet[A_ENUM_VALUE_REF](stage) {
		x.Name = x.ENUM_VALUE_REF
	}

	for x := range *GetGongstructInstancesSet[A_OBJECT](stage) {
		x.Name = x.SPEC_OBJECT_REF
	}

	for x := range *GetGongstructInstancesSet[A_SPECIFICATION_TYPE_REF](stage) {
		x.Name = x.SPECIFICATION_TYPE_REF
	}

	for x := range *GetGongstructInstancesSet[A_SPEC_OBJECT_TYPE_REF](stage) {
		x.Name = x.SPEC_OBJECT_TYPE_REF
	}

	for x := range *GetGongstructInstancesSet[DATATYPE_DEFINITION_ENUMERATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[DATATYPE_DEFINITION_XHTML](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[EMBEDDED_VALUE](stage) {
		x.Name = x.OTHER_CONTENT
	}

	for x := range *GetGongstructInstancesSet[ENUM_VALUE](stage) {
		x.Name = x.LONG_NAME
	}

	// DATATYPE_DEFINITION_XHTML EMBEDDED_VALUE ENUM_VALUE

	for x := range *GetGongstructInstancesSet[SPECIFICATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[SPECIFICATION_TYPE](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[SPEC_HIERARCHY](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[SPEC_OBJECT](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[SPEC_OBJECT_TYPE](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[SPEC_RELATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *GetGongstructInstancesSet[SPEC_RELATION_TYPE](stage) {
		x.Name = x.LONG_NAME
	}
}
