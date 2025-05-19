package models

func (stager *Stager) initMaps() {
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

	stager.Map_id_attributeDefinitionENUM = make(map[string]*ATTRIBUTE_DEFINITION_ENUMERATION)
	{
		attributeDefinitionENUMs := *GetGongstructInstancesSet[ATTRIBUTE_DEFINITION_ENUMERATION](stager.GetStage())
		for attributeDefinition := range attributeDefinitionENUMs {
			stager.Map_id_attributeDefinitionENUM[attributeDefinition.IDENTIFIER] = attributeDefinition
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
