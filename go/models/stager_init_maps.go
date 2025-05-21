package models

func (stager *Stager) initMaps() {

	datatypes_xhtml := *GetGongstructInstancesSet[DATATYPE_DEFINITION_XHTML](stager.GetStage())
	stager.Map_id_datatypes_xhtml = make(map[string]*DATATYPE_DEFINITION_XHTML)
	for datatype_xhtml := range datatypes_xhtml {
		stager.Map_id_datatypes_xhtml[datatype_xhtml.IDENTIFIER] = datatype_xhtml
	}

	datatypes_string := *GetGongstructInstancesSet[DATATYPE_DEFINITION_STRING](stager.GetStage())
	stager.Map_id_datatypes_string = make(map[string]*DATATYPE_DEFINITION_STRING)
	for datatype_string := range datatypes_string {
		stager.Map_id_datatypes_string[datatype_string.IDENTIFIER] = datatype_string
	}

	datatypes_boolean := *GetGongstructInstancesSet[DATATYPE_DEFINITION_BOOLEAN](stager.GetStage())
	stager.Map_id_datatypes_boolean = make(map[string]*DATATYPE_DEFINITION_BOOLEAN)
	for datatype_boolean := range datatypes_boolean {
		stager.Map_id_datatypes_boolean[datatype_boolean.IDENTIFIER] = datatype_boolean
	}

	datatypes_integer := *GetGongstructInstancesSet[DATATYPE_DEFINITION_INTEGER](stager.GetStage())
	stager.Map_id_datatypes_integer = make(map[string]*DATATYPE_DEFINITION_INTEGER)
	for datatype_integer := range datatypes_integer {
		stager.Map_id_datatypes_integer[datatype_integer.IDENTIFIER] = datatype_integer
	}

	datatypes_real := *GetGongstructInstancesSet[DATATYPE_DEFINITION_REAL](stager.GetStage())
	stager.Map_id_datatypes_real = make(map[string]*DATATYPE_DEFINITION_REAL)
	for datatype_real := range datatypes_real {
		stager.Map_id_datatypes_real[datatype_real.IDENTIFIER] = datatype_real
	}

	datatypes_date := *GetGongstructInstancesSet[DATATYPE_DEFINITION_DATE](stager.GetStage())
	stager.Map_id_datatypes_date = make(map[string]*DATATYPE_DEFINITION_DATE)
	for datatype_date := range datatypes_date {
		stager.Map_id_datatypes_date[datatype_date.IDENTIFIER] = datatype_date
	}

	datatypes_enumeration := *GetGongstructInstancesSet[DATATYPE_DEFINITION_ENUMERATION](stager.GetStage())
	stager.Map_id_datatypes_enumeration = make(map[string]*DATATYPE_DEFINITION_ENUMERATION)
	for datatype_enumeration := range datatypes_enumeration {
		stager.Map_id_datatypes_enumeration[datatype_enumeration.IDENTIFIER] = datatype_enumeration
	}

	stager.Map_id_specobjectTypes = make(map[string]*SPEC_OBJECT_TYPE)
	{
		specObjectTypes := *GetGongstructInstancesSet[SPEC_OBJECT_TYPE](stager.GetStage())
		for specObjectType := range specObjectTypes {
			stager.Map_id_specobjectTypes[specObjectType.IDENTIFIER] = specObjectType
		}
	}

	stager.Map_id_specificationType = make(map[string]*SPECIFICATION_TYPE)
	{
		types := *GetGongstructInstancesSet[SPECIFICATION_TYPE](stager.GetStage())
		for type_ := range types {
			stager.Map_id_specificationType[type_.IDENTIFIER] = type_
		}
	}

	stager.Map_id_specObject = make(map[string]*SPEC_OBJECT)
	{
		specObjects := *GetGongstructInstancesSet[SPEC_OBJECT](stager.GetStage())
		for specObject := range specObjects {
			stager.Map_id_specObject[specObject.IDENTIFIER] = specObject
		}
	}

	stager.Map_id_attributeDefinitionXHTML = make(map[string]*ATTRIBUTE_DEFINITION_XHTML)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_XHTML](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_attributeDefinitionXHTML[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_attributeDefinitionString = make(map[string]*ATTRIBUTE_DEFINITION_STRING)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_STRING](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_attributeDefinitionString[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_attributeDefinitionBoolean = make(map[string]*ATTRIBUTE_DEFINITION_BOOLEAN)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_BOOLEAN](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_attributeDefinitionBoolean[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_attributeDefinitionInteger = make(map[string]*ATTRIBUTE_DEFINITION_INTEGER)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_INTEGER](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_attributeDefinitionInteger[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_attributeDefinitionDate = make(map[string]*ATTRIBUTE_DEFINITION_DATE)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_DATE](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_attributeDefinitionDate[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_attributeDefinitionReal = make(map[string]*ATTRIBUTE_DEFINITION_REAL)
	{
		attributeDefinitionXHTMLs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_REAL](stager.GetStage())
		for attributeDefinition := range attributeDefinitionXHTMLs {
			stager.Map_id_attributeDefinitionReal[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_attributeDefinitionEnum = make(map[string]*ATTRIBUTE_DEFINITION_ENUMERATION)
	{
		attributeDefinitionENUMs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_ENUMERATION](stager.GetStage())
		for attributeDefinition := range attributeDefinitionENUMs {
			stager.Map_id_attributeDefinitionEnum[attributeDefinition.IDENTIFIER] = attributeDefinition
		}
	}

	stager.Map_id_enumValues = make(map[string]*ENUM_VALUE)
	{
		enumValuess := *GetGongstructInstancesSet[ENUM_VALUE](stager.GetStage())
		for enumValue := range enumValuess {
			stager.Map_id_enumValues[enumValue.IDENTIFIER] = enumValue
		}
	}

	stager.Map_id_specRelationType = make(map[string]*SPEC_RELATION_TYPE)
	{
		specRelationTypes := *GetGongstructInstancesSet[SPEC_RELATION_TYPE](stager.GetStage())
		for specRelationType := range specRelationTypes {
			stager.Map_id_specRelationType[specRelationType.IDENTIFIER] = specRelationType
		}
	}
}
