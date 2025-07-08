package models

func (stager *Stager) initMaps() {

	datatypes_xhtml := *GetGongstructInstancesSet[DATATYPE_DEFINITION_XHTML](stager.GetStage())
	stager.Map_id_DATATYPE_DEFINITION_XHTML = make(map[string]*DATATYPE_DEFINITION_XHTML)
	for datatype_xhtml := range datatypes_xhtml {
		stager.Map_id_DATATYPE_DEFINITION_XHTML[datatype_xhtml.IDENTIFIER] = datatype_xhtml
	}

	datatypes_string := *GetGongstructInstancesSet[DATATYPE_DEFINITION_STRING](stager.GetStage())
	stager.Map_id_DATATYPE_DEFINITION_STRING = make(map[string]*DATATYPE_DEFINITION_STRING)
	for datatype_string := range datatypes_string {
		stager.Map_id_DATATYPE_DEFINITION_STRING[datatype_string.IDENTIFIER] = datatype_string
	}

	datatypes_boolean := *GetGongstructInstancesSet[DATATYPE_DEFINITION_BOOLEAN](stager.GetStage())
	stager.Map_id_DATATYPE_DEFINITION_BOOLEAN = make(map[string]*DATATYPE_DEFINITION_BOOLEAN)
	for datatype_boolean := range datatypes_boolean {
		stager.Map_id_DATATYPE_DEFINITION_BOOLEAN[datatype_boolean.IDENTIFIER] = datatype_boolean
	}

	datatypes_integer := *GetGongstructInstancesSet[DATATYPE_DEFINITION_INTEGER](stager.GetStage())
	stager.Map_id_DATATYPE_DEFINITION_INTEGER = make(map[string]*DATATYPE_DEFINITION_INTEGER)
	for datatype_integer := range datatypes_integer {
		stager.Map_id_DATATYPE_DEFINITION_INTEGER[datatype_integer.IDENTIFIER] = datatype_integer
	}

	datatypes_real := *GetGongstructInstancesSet[DATATYPE_DEFINITION_REAL](stager.GetStage())
	stager.Map_id_DATATYPE_DEFINITION_REAL = make(map[string]*DATATYPE_DEFINITION_REAL)
	for datatype_real := range datatypes_real {
		stager.Map_id_DATATYPE_DEFINITION_REAL[datatype_real.IDENTIFIER] = datatype_real
	}

	datatypes_date := *GetGongstructInstancesSet[DATATYPE_DEFINITION_DATE](stager.GetStage())
	stager.Map_id_DATATYPE_DEFINITION_DATE = make(map[string]*DATATYPE_DEFINITION_DATE)
	for datatype_date := range datatypes_date {
		stager.Map_id_DATATYPE_DEFINITION_DATE[datatype_date.IDENTIFIER] = datatype_date
	}

	datatypes_enumeration := *GetGongstructInstancesSet[DATATYPE_DEFINITION_ENUMERATION](stager.GetStage())
	stager.Map_id_DATATYPE_DEFINITION_ENUMERATION = make(map[string]*DATATYPE_DEFINITION_ENUMERATION)
	for datatype_enumeration := range datatypes_enumeration {
		stager.Map_id_DATATYPE_DEFINITION_ENUMERATION[datatype_enumeration.IDENTIFIER] = datatype_enumeration
	}

	stager.Map_id_SPEC_OBJECT_TYPE = make(map[string]*SPEC_OBJECT_TYPE)
	{
		specObjectTypes := *GetGongstructInstancesSet[SPEC_OBJECT_TYPE](stager.GetStage())
		for specObjectType := range specObjectTypes {
			stager.Map_id_SPEC_OBJECT_TYPE[specObjectType.IDENTIFIER] = specObjectType
		}
	}

	stager.Map_id_SPECIFICATION_TYPE = make(map[string]*SPECIFICATION_TYPE)
	{
		types := *GetGongstructInstancesSet[SPECIFICATION_TYPE](stager.GetStage())
		for type_ := range types {
			stager.Map_id_SPECIFICATION_TYPE[type_.IDENTIFIER] = type_
		}
	}

	stager.Map_id_SPEC_OBJECT = make(map[string]*SPEC_OBJECT)
	{
		specObjects := *GetGongstructInstancesSet[SPEC_OBJECT](stager.GetStage())
		for specObject := range specObjects {
			stager.Map_id_SPEC_OBJECT[specObject.IDENTIFIER] = specObject
		}
	}

	stager.Map_id_ATTRIBUTE_DEFINITION_XHTML = make(map[string]*ATTRIBUTE_DEFINITION_XHTML)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_XHTML](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_ATTRIBUTE_DEFINITION_STRING = make(map[string]*ATTRIBUTE_DEFINITION_STRING)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_STRING](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_ATTRIBUTE_DEFINITION_STRING[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN = make(map[string]*ATTRIBUTE_DEFINITION_BOOLEAN)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_BOOLEAN](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER = make(map[string]*ATTRIBUTE_DEFINITION_INTEGER)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_INTEGER](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_ATTRIBUTE_DEFINITION_DATE = make(map[string]*ATTRIBUTE_DEFINITION_DATE)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_DATE](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_ATTRIBUTE_DEFINITION_DATE[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_ATTRIBUTE_DEFINITION_REAL = make(map[string]*ATTRIBUTE_DEFINITION_REAL)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_REAL](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_ATTRIBUTE_DEFINITION_REAL[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION = make(map[string]*ATTRIBUTE_DEFINITION_ENUMERATION)
	{
		attributeDefinitionENUMs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_ENUMERATION](stager.GetStage())
		for attributeDefinition := range attributeDefinitionENUMs {
			stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_ENUM_VALUE = make(map[string]*ENUM_VALUE)
	{
		enumValuess := *GetGongstructInstancesSet[ENUM_VALUE](stager.GetStage())
		for enumValue := range enumValuess {
			stager.Map_id_ENUM_VALUE[enumValue.IDENTIFIER] = enumValue
		}
	}

	stager.Map_id_SPEC_RELATION_TYPE = make(map[string]*SPEC_RELATION_TYPE)
	{
		specRelationTypes := *GetGongstructInstancesSet[SPEC_RELATION_TYPE](stager.GetStage())
		for specRelationType := range specRelationTypes {
			stager.Map_id_SPEC_RELATION_TYPE[specRelationType.IDENTIFIER] = specRelationType
		}
	}

	stager.Map_SpecificationNodes_exapanded = make(map[*SPECIFICATION]bool)
}
