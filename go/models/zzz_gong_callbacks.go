// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		if stage.OnAfterALTERNATIVE_IDCreateCallback != nil {
			stage.OnAfterALTERNATIVE_IDCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		if stage.OnAfterATTRIBUTE_DEFINITION_DATECreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_DATECreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		if stage.OnAfterATTRIBUTE_DEFINITION_REALCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_REALCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterATTRIBUTE_VALUE_DATECreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_DATECreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterATTRIBUTE_VALUE_REALCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_REALCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterATTRIBUTE_VALUE_STRINGCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_STRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ALTERNATIVE_ID:
		if stage.OnAfterA_ALTERNATIVE_IDCreateCallback != nil {
			stage.OnAfterA_ALTERNATIVE_IDCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterA_ATTRIBUTE_VALUE_DATECreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_DATECreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterA_ATTRIBUTE_VALUE_REALCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_REALCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_STRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1CreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1CreateCallback.OnAfterCreate(stage, target)
		}
	case *A_CHILDREN:
		if stage.OnAfterA_CHILDRENCreateCallback != nil {
			stage.OnAfterA_CHILDRENCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_CORE_CONTENT:
		if stage.OnAfterA_CORE_CONTENTCreateCallback != nil {
			stage.OnAfterA_CORE_CONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPES:
		if stage.OnAfterA_DATATYPESCreateCallback != nil {
			stage.OnAfterA_DATATYPESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_EDITABLE_ATTS:
		if stage.OnAfterA_EDITABLE_ATTSCreateCallback != nil {
			stage.OnAfterA_EDITABLE_ATTSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ENUM_VALUE_REF:
		if stage.OnAfterA_ENUM_VALUE_REFCreateCallback != nil {
			stage.OnAfterA_ENUM_VALUE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_OBJECT:
		if stage.OnAfterA_OBJECTCreateCallback != nil {
			stage.OnAfterA_OBJECTCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_PROPERTIES:
		if stage.OnAfterA_PROPERTIESCreateCallback != nil {
			stage.OnAfterA_PROPERTIESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_RELATION_GROUP_TYPE_REF:
		if stage.OnAfterA_RELATION_GROUP_TYPE_REFCreateCallback != nil {
			stage.OnAfterA_RELATION_GROUP_TYPE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SOURCE_1:
		if stage.OnAfterA_SOURCE_1CreateCallback != nil {
			stage.OnAfterA_SOURCE_1CreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SOURCE_SPECIFICATION_1:
		if stage.OnAfterA_SOURCE_SPECIFICATION_1CreateCallback != nil {
			stage.OnAfterA_SOURCE_SPECIFICATION_1CreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPECIFICATIONS:
		if stage.OnAfterA_SPECIFICATIONSCreateCallback != nil {
			stage.OnAfterA_SPECIFICATIONSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPECIFICATION_TYPE_REF:
		if stage.OnAfterA_SPECIFICATION_TYPE_REFCreateCallback != nil {
			stage.OnAfterA_SPECIFICATION_TYPE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPECIFIED_VALUES:
		if stage.OnAfterA_SPECIFIED_VALUESCreateCallback != nil {
			stage.OnAfterA_SPECIFIED_VALUESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_ATTRIBUTES:
		if stage.OnAfterA_SPEC_ATTRIBUTESCreateCallback != nil {
			stage.OnAfterA_SPEC_ATTRIBUTESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_OBJECTS:
		if stage.OnAfterA_SPEC_OBJECTSCreateCallback != nil {
			stage.OnAfterA_SPEC_OBJECTSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_OBJECT_TYPE_REF:
		if stage.OnAfterA_SPEC_OBJECT_TYPE_REFCreateCallback != nil {
			stage.OnAfterA_SPEC_OBJECT_TYPE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_RELATIONS:
		if stage.OnAfterA_SPEC_RELATIONSCreateCallback != nil {
			stage.OnAfterA_SPEC_RELATIONSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_RELATION_GROUPS:
		if stage.OnAfterA_SPEC_RELATION_GROUPSCreateCallback != nil {
			stage.OnAfterA_SPEC_RELATION_GROUPSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_RELATION_REF:
		if stage.OnAfterA_SPEC_RELATION_REFCreateCallback != nil {
			stage.OnAfterA_SPEC_RELATION_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_RELATION_TYPE_REF:
		if stage.OnAfterA_SPEC_RELATION_TYPE_REFCreateCallback != nil {
			stage.OnAfterA_SPEC_RELATION_TYPE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_TYPES:
		if stage.OnAfterA_SPEC_TYPESCreateCallback != nil {
			stage.OnAfterA_SPEC_TYPESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_THE_HEADER:
		if stage.OnAfterA_THE_HEADERCreateCallback != nil {
			stage.OnAfterA_THE_HEADERCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_TOOL_EXTENSIONS:
		if stage.OnAfterA_TOOL_EXTENSIONSCreateCallback != nil {
			stage.OnAfterA_TOOL_EXTENSIONSCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_DATE:
		if stage.OnAfterDATATYPE_DEFINITION_DATECreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_DATECreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_INTEGER:
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_REAL:
		if stage.OnAfterDATATYPE_DEFINITION_REALCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_REALCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_STRING:
		if stage.OnAfterDATATYPE_DEFINITION_STRINGCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_STRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_XHTML:
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *EMBEDDED_VALUE:
		if stage.OnAfterEMBEDDED_VALUECreateCallback != nil {
			stage.OnAfterEMBEDDED_VALUECreateCallback.OnAfterCreate(stage, target)
		}
	case *ENUM_VALUE:
		if stage.OnAfterENUM_VALUECreateCallback != nil {
			stage.OnAfterENUM_VALUECreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryCreateCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_SPECIFICATION_Nodes_expandedEntry:
		if stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryCreateCallback != nil {
			stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryCreateCallback != nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryCreateCallback != nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryCreateCallback != nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryCreateCallback.OnAfterCreate(stage, target)
		}
	case *RELATION_GROUP:
		if stage.OnAfterRELATION_GROUPCreateCallback != nil {
			stage.OnAfterRELATION_GROUPCreateCallback.OnAfterCreate(stage, target)
		}
	case *RELATION_GROUP_TYPE:
		if stage.OnAfterRELATION_GROUP_TYPECreateCallback != nil {
			stage.OnAfterRELATION_GROUP_TYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF:
		if stage.OnAfterREQ_IFCreateCallback != nil {
			stage.OnAfterREQ_IFCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF_CONTENT:
		if stage.OnAfterREQ_IF_CONTENTCreateCallback != nil {
			stage.OnAfterREQ_IF_CONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF_HEADER:
		if stage.OnAfterREQ_IF_HEADERCreateCallback != nil {
			stage.OnAfterREQ_IF_HEADERCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF_TOOL_EXTENSION:
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback != nil {
			stage.OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *RenderingConfiguration:
		if stage.OnAfterRenderingConfigurationCreateCallback != nil {
			stage.OnAfterRenderingConfigurationCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONCreateCallback != nil {
			stage.OnAfterSPECIFICATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECIFICATION_TYPE:
		if stage.OnAfterSPECIFICATION_TYPECreateCallback != nil {
			stage.OnAfterSPECIFICATION_TYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_HIERARCHY:
		if stage.OnAfterSPEC_HIERARCHYCreateCallback != nil {
			stage.OnAfterSPEC_HIERARCHYCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_OBJECT:
		if stage.OnAfterSPEC_OBJECTCreateCallback != nil {
			stage.OnAfterSPEC_OBJECTCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_OBJECT_TYPE:
		if stage.OnAfterSPEC_OBJECT_TYPECreateCallback != nil {
			stage.OnAfterSPEC_OBJECT_TYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_RELATION:
		if stage.OnAfterSPEC_RELATIONCreateCallback != nil {
			stage.OnAfterSPEC_RELATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_RELATION_TYPE:
		if stage.OnAfterSPEC_RELATION_TYPECreateCallback != nil {
			stage.OnAfterSPEC_RELATION_TYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *StaticWebSite:
		if stage.OnAfterStaticWebSiteCreateCallback != nil {
			stage.OnAfterStaticWebSiteCreateCallback.OnAfterCreate(stage, target)
		}
	case *StaticWebSiteChapter:
		if stage.OnAfterStaticWebSiteChapterCreateCallback != nil {
			stage.OnAfterStaticWebSiteChapterCreateCallback.OnAfterCreate(stage, target)
		}
	case *StaticWebSiteGeneratedImage:
		if stage.OnAfterStaticWebSiteGeneratedImageCreateCallback != nil {
			stage.OnAfterStaticWebSiteGeneratedImageCreateCallback.OnAfterCreate(stage, target)
		}
	case *StaticWebSiteImage:
		if stage.OnAfterStaticWebSiteImageCreateCallback != nil {
			stage.OnAfterStaticWebSiteImageCreateCallback.OnAfterCreate(stage, target)
		}
	case *StaticWebSiteParagraph:
		if stage.OnAfterStaticWebSiteParagraphCreateCallback != nil {
			stage.OnAfterStaticWebSiteParagraphCreateCallback.OnAfterCreate(stage, target)
		}
	case *XHTML_CONTENT:
		if stage.OnAfterXHTML_CONTENTCreateCallback != nil {
			stage.OnAfterXHTML_CONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type, mouseEvent *Gong__MouseEvent) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		newTarget := any(new).(*ALTERNATIVE_ID)
		if stage.OnAfterALTERNATIVE_IDUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterALTERNATIVE_IDUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterALTERNATIVE_IDUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterALTERNATIVE_IDUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_BOOLEAN)
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_DATE)
		if stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_ENUMERATION)
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_INTEGER)
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_REAL)
		if stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_STRING)
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_XHTML)
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		newTarget := any(new).(*ATTRIBUTE_VALUE_BOOLEAN)
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_VALUE_DATE:
		newTarget := any(new).(*ATTRIBUTE_VALUE_DATE)
		if stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_VALUE_DATEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_VALUE_DATEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		newTarget := any(new).(*ATTRIBUTE_VALUE_ENUMERATION)
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		newTarget := any(new).(*ATTRIBUTE_VALUE_INTEGER)
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_VALUE_REAL:
		newTarget := any(new).(*ATTRIBUTE_VALUE_REAL)
		if stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_VALUE_REALUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_VALUE_REALUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_VALUE_STRING:
		newTarget := any(new).(*ATTRIBUTE_VALUE_STRING)
		if stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ATTRIBUTE_VALUE_XHTML:
		newTarget := any(new).(*ATTRIBUTE_VALUE_XHTML)
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ALTERNATIVE_ID:
		newTarget := any(new).(*A_ALTERNATIVE_ID)
		if stage.OnAfterA_ALTERNATIVE_IDUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ALTERNATIVE_IDUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ALTERNATIVE_IDUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ALTERNATIVE_IDUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_DATE_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_REAL_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_STRING_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_BOOLEAN)
		if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_VALUE_DATE:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_DATE)
		if stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_ENUMERATION)
		if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_INTEGER)
		if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_VALUE_REAL:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_REAL)
		if stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_VALUE_STRING:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_STRING)
		if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_VALUE_XHTML:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_XHTML)
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_XHTML_1)
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_CHILDREN:
		newTarget := any(new).(*A_CHILDREN)
		if stage.OnAfterA_CHILDRENUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_CHILDRENUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_CHILDRENUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_CHILDRENUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_CORE_CONTENT:
		newTarget := any(new).(*A_CORE_CONTENT)
		if stage.OnAfterA_CORE_CONTENTUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_CORE_CONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_CORE_CONTENTUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_CORE_CONTENTUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_DATATYPES:
		newTarget := any(new).(*A_DATATYPES)
		if stage.OnAfterA_DATATYPESUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_DATATYPESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_DATATYPESUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_DATATYPESUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_DATE_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_INTEGER_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_REAL_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_STRING_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_XHTML_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_EDITABLE_ATTS:
		newTarget := any(new).(*A_EDITABLE_ATTS)
		if stage.OnAfterA_EDITABLE_ATTSUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_EDITABLE_ATTSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_EDITABLE_ATTSUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_EDITABLE_ATTSUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_ENUM_VALUE_REF:
		newTarget := any(new).(*A_ENUM_VALUE_REF)
		if stage.OnAfterA_ENUM_VALUE_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_ENUM_VALUE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_ENUM_VALUE_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_ENUM_VALUE_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_OBJECT:
		newTarget := any(new).(*A_OBJECT)
		if stage.OnAfterA_OBJECTUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_OBJECTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_OBJECTUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_OBJECTUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_PROPERTIES:
		newTarget := any(new).(*A_PROPERTIES)
		if stage.OnAfterA_PROPERTIESUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_PROPERTIESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_PROPERTIESUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_PROPERTIESUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_RELATION_GROUP_TYPE_REF:
		newTarget := any(new).(*A_RELATION_GROUP_TYPE_REF)
		if stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SOURCE_1:
		newTarget := any(new).(*A_SOURCE_1)
		if stage.OnAfterA_SOURCE_1UpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SOURCE_1UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SOURCE_1UpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SOURCE_1UpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SOURCE_SPECIFICATION_1:
		newTarget := any(new).(*A_SOURCE_SPECIFICATION_1)
		if stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPECIFICATIONS:
		newTarget := any(new).(*A_SPECIFICATIONS)
		if stage.OnAfterA_SPECIFICATIONSUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPECIFICATIONSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPECIFICATIONSUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPECIFICATIONSUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPECIFICATION_TYPE_REF:
		newTarget := any(new).(*A_SPECIFICATION_TYPE_REF)
		if stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPECIFIED_VALUES:
		newTarget := any(new).(*A_SPECIFIED_VALUES)
		if stage.OnAfterA_SPECIFIED_VALUESUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPECIFIED_VALUESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPECIFIED_VALUESUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPECIFIED_VALUESUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPEC_ATTRIBUTES:
		newTarget := any(new).(*A_SPEC_ATTRIBUTES)
		if stage.OnAfterA_SPEC_ATTRIBUTESUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPEC_ATTRIBUTESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPEC_ATTRIBUTESUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPEC_ATTRIBUTESUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPEC_OBJECTS:
		newTarget := any(new).(*A_SPEC_OBJECTS)
		if stage.OnAfterA_SPEC_OBJECTSUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPEC_OBJECTSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPEC_OBJECTSUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPEC_OBJECTSUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPEC_OBJECT_TYPE_REF:
		newTarget := any(new).(*A_SPEC_OBJECT_TYPE_REF)
		if stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPEC_RELATIONS:
		newTarget := any(new).(*A_SPEC_RELATIONS)
		if stage.OnAfterA_SPEC_RELATIONSUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPEC_RELATIONSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPEC_RELATIONSUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPEC_RELATIONSUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPEC_RELATION_GROUPS:
		newTarget := any(new).(*A_SPEC_RELATION_GROUPS)
		if stage.OnAfterA_SPEC_RELATION_GROUPSUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPEC_RELATION_GROUPSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPEC_RELATION_GROUPSUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPEC_RELATION_GROUPSUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPEC_RELATION_REF:
		newTarget := any(new).(*A_SPEC_RELATION_REF)
		if stage.OnAfterA_SPEC_RELATION_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPEC_RELATION_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPEC_RELATION_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPEC_RELATION_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPEC_RELATION_TYPE_REF:
		newTarget := any(new).(*A_SPEC_RELATION_TYPE_REF)
		if stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_SPEC_TYPES:
		newTarget := any(new).(*A_SPEC_TYPES)
		if stage.OnAfterA_SPEC_TYPESUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_SPEC_TYPESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_SPEC_TYPESUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_SPEC_TYPESUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_THE_HEADER:
		newTarget := any(new).(*A_THE_HEADER)
		if stage.OnAfterA_THE_HEADERUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_THE_HEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_THE_HEADERUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_THE_HEADERUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *A_TOOL_EXTENSIONS:
		newTarget := any(new).(*A_TOOL_EXTENSIONS)
		if stage.OnAfterA_TOOL_EXTENSIONSUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterA_TOOL_EXTENSIONSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterA_TOOL_EXTENSIONSUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterA_TOOL_EXTENSIONSUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		newTarget := any(new).(*DATATYPE_DEFINITION_BOOLEAN)
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *DATATYPE_DEFINITION_DATE:
		newTarget := any(new).(*DATATYPE_DEFINITION_DATE)
		if stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterDATATYPE_DEFINITION_DATEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterDATATYPE_DEFINITION_DATEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		newTarget := any(new).(*DATATYPE_DEFINITION_ENUMERATION)
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *DATATYPE_DEFINITION_INTEGER:
		newTarget := any(new).(*DATATYPE_DEFINITION_INTEGER)
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *DATATYPE_DEFINITION_REAL:
		newTarget := any(new).(*DATATYPE_DEFINITION_REAL)
		if stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterDATATYPE_DEFINITION_REALUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterDATATYPE_DEFINITION_REALUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *DATATYPE_DEFINITION_STRING:
		newTarget := any(new).(*DATATYPE_DEFINITION_STRING)
		if stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *DATATYPE_DEFINITION_XHTML:
		newTarget := any(new).(*DATATYPE_DEFINITION_XHTML)
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *EMBEDDED_VALUE:
		newTarget := any(new).(*EMBEDDED_VALUE)
		if stage.OnAfterEMBEDDED_VALUEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterEMBEDDED_VALUEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterEMBEDDED_VALUEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterEMBEDDED_VALUEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *ENUM_VALUE:
		newTarget := any(new).(*ENUM_VALUE)
		if stage.OnAfterENUM_VALUEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterENUM_VALUEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterENUM_VALUEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterENUM_VALUEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		newTarget := any(new).(*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry)
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_SPECIFICATION_Nodes_expandedEntry:
		newTarget := any(new).(*Map_SPECIFICATION_Nodes_expandedEntry)
		if stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		newTarget := any(new).(*Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry)
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		newTarget := any(new).(*Map_SPEC_OBJECT_TYPE_showIdentifierEntry)
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		newTarget := any(new).(*Map_SPEC_OBJECT_TYPE_showNameEntry)
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *RELATION_GROUP:
		newTarget := any(new).(*RELATION_GROUP)
		if stage.OnAfterRELATION_GROUPUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterRELATION_GROUPUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterRELATION_GROUPUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterRELATION_GROUPUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *RELATION_GROUP_TYPE:
		newTarget := any(new).(*RELATION_GROUP_TYPE)
		if stage.OnAfterRELATION_GROUP_TYPEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterRELATION_GROUP_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterRELATION_GROUP_TYPEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterRELATION_GROUP_TYPEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *REQ_IF:
		newTarget := any(new).(*REQ_IF)
		if stage.OnAfterREQ_IFUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterREQ_IFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterREQ_IFUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterREQ_IFUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *REQ_IF_CONTENT:
		newTarget := any(new).(*REQ_IF_CONTENT)
		if stage.OnAfterREQ_IF_CONTENTUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterREQ_IF_CONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterREQ_IF_CONTENTUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterREQ_IF_CONTENTUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *REQ_IF_HEADER:
		newTarget := any(new).(*REQ_IF_HEADER)
		if stage.OnAfterREQ_IF_HEADERUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterREQ_IF_HEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterREQ_IF_HEADERUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterREQ_IF_HEADERUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *REQ_IF_TOOL_EXTENSION:
		newTarget := any(new).(*REQ_IF_TOOL_EXTENSION)
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *RenderingConfiguration:
		newTarget := any(new).(*RenderingConfiguration)
		if stage.OnAfterRenderingConfigurationUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterRenderingConfigurationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterRenderingConfigurationUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterRenderingConfigurationUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *SPECIFICATION:
		newTarget := any(new).(*SPECIFICATION)
		if stage.OnAfterSPECIFICATIONUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSPECIFICATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSPECIFICATIONUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSPECIFICATIONUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *SPECIFICATION_TYPE:
		newTarget := any(new).(*SPECIFICATION_TYPE)
		if stage.OnAfterSPECIFICATION_TYPEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSPECIFICATION_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSPECIFICATION_TYPEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSPECIFICATION_TYPEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *SPEC_HIERARCHY:
		newTarget := any(new).(*SPEC_HIERARCHY)
		if stage.OnAfterSPEC_HIERARCHYUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSPEC_HIERARCHYUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSPEC_HIERARCHYUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSPEC_HIERARCHYUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *SPEC_OBJECT:
		newTarget := any(new).(*SPEC_OBJECT)
		if stage.OnAfterSPEC_OBJECTUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSPEC_OBJECTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSPEC_OBJECTUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSPEC_OBJECTUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *SPEC_OBJECT_TYPE:
		newTarget := any(new).(*SPEC_OBJECT_TYPE)
		if stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSPEC_OBJECT_TYPEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSPEC_OBJECT_TYPEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *SPEC_RELATION:
		newTarget := any(new).(*SPEC_RELATION)
		if stage.OnAfterSPEC_RELATIONUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSPEC_RELATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSPEC_RELATIONUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSPEC_RELATIONUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *SPEC_RELATION_TYPE:
		newTarget := any(new).(*SPEC_RELATION_TYPE)
		if stage.OnAfterSPEC_RELATION_TYPEUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterSPEC_RELATION_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterSPEC_RELATION_TYPEUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterSPEC_RELATION_TYPEUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *StaticWebSite:
		newTarget := any(new).(*StaticWebSite)
		if stage.OnAfterStaticWebSiteUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterStaticWebSiteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterStaticWebSiteUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterStaticWebSiteUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *StaticWebSiteChapter:
		newTarget := any(new).(*StaticWebSiteChapter)
		if stage.OnAfterStaticWebSiteChapterUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterStaticWebSiteChapterUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterStaticWebSiteChapterUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterStaticWebSiteChapterUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *StaticWebSiteGeneratedImage:
		newTarget := any(new).(*StaticWebSiteGeneratedImage)
		if stage.OnAfterStaticWebSiteGeneratedImageUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterStaticWebSiteGeneratedImageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterStaticWebSiteGeneratedImageUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterStaticWebSiteGeneratedImageUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *StaticWebSiteImage:
		newTarget := any(new).(*StaticWebSiteImage)
		if stage.OnAfterStaticWebSiteImageUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterStaticWebSiteImageUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterStaticWebSiteImageUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterStaticWebSiteImageUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *StaticWebSiteParagraph:
		newTarget := any(new).(*StaticWebSiteParagraph)
		if stage.OnAfterStaticWebSiteParagraphUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterStaticWebSiteParagraphUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterStaticWebSiteParagraphUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterStaticWebSiteParagraphUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	case *XHTML_CONTENT:
		newTarget := any(new).(*XHTML_CONTENT)
		if stage.OnAfterXHTML_CONTENTUpdateCallback != nil && mouseEvent == nil {
			stage.OnAfterXHTML_CONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
		if stage.OnAfterXHTML_CONTENTUpdateWithMouseEventCallback != nil && mouseEvent != nil {
			stage.OnAfterXHTML_CONTENTUpdateWithMouseEventCallback.OnAfterUpdateWithMouseEvent(stage, oldTarget, newTarget, mouseEvent)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		if stage.OnAfterALTERNATIVE_IDDeleteCallback != nil {
			staged := any(staged).(*ALTERNATIVE_ID)
			stage.OnAfterALTERNATIVE_IDDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_BOOLEAN)
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		if stage.OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_DATE)
			stage.OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_ENUMERATION)
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_INTEGER)
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		if stage.OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_REAL)
			stage.OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_STRING)
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_XHTML)
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_BOOLEAN)
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterATTRIBUTE_VALUE_DATEDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_DATE)
			stage.OnAfterATTRIBUTE_VALUE_DATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_ENUMERATION)
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_INTEGER)
			stage.OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterATTRIBUTE_VALUE_REALDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_REAL)
			stage.OnAfterATTRIBUTE_VALUE_REALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_STRING)
			stage.OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_XHTML)
			stage.OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ALTERNATIVE_ID:
		if stage.OnAfterA_ALTERNATIVE_IDDeleteCallback != nil {
			staged := any(staged).(*A_ALTERNATIVE_ID)
			stage.OnAfterA_ALTERNATIVE_IDDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_DATE_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_REAL_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_STRING_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_BOOLEAN)
			stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterA_ATTRIBUTE_VALUE_DATEDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_DATE)
			stage.OnAfterA_ATTRIBUTE_VALUE_DATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_ENUMERATION)
			stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_INTEGER)
			stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterA_ATTRIBUTE_VALUE_REALDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_REAL)
			stage.OnAfterA_ATTRIBUTE_VALUE_REALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_STRING)
			stage.OnAfterA_ATTRIBUTE_VALUE_STRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_XHTML)
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1DeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_XHTML_1)
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_CHILDREN:
		if stage.OnAfterA_CHILDRENDeleteCallback != nil {
			staged := any(staged).(*A_CHILDREN)
			stage.OnAfterA_CHILDRENDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_CORE_CONTENT:
		if stage.OnAfterA_CORE_CONTENTDeleteCallback != nil {
			staged := any(staged).(*A_CORE_CONTENT)
			stage.OnAfterA_CORE_CONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPES:
		if stage.OnAfterA_DATATYPESDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPES)
			stage.OnAfterA_DATATYPESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_DATE_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_INTEGER_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_REAL_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_STRING_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_XHTML_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_EDITABLE_ATTS:
		if stage.OnAfterA_EDITABLE_ATTSDeleteCallback != nil {
			staged := any(staged).(*A_EDITABLE_ATTS)
			stage.OnAfterA_EDITABLE_ATTSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ENUM_VALUE_REF:
		if stage.OnAfterA_ENUM_VALUE_REFDeleteCallback != nil {
			staged := any(staged).(*A_ENUM_VALUE_REF)
			stage.OnAfterA_ENUM_VALUE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_OBJECT:
		if stage.OnAfterA_OBJECTDeleteCallback != nil {
			staged := any(staged).(*A_OBJECT)
			stage.OnAfterA_OBJECTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_PROPERTIES:
		if stage.OnAfterA_PROPERTIESDeleteCallback != nil {
			staged := any(staged).(*A_PROPERTIES)
			stage.OnAfterA_PROPERTIESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_RELATION_GROUP_TYPE_REF:
		if stage.OnAfterA_RELATION_GROUP_TYPE_REFDeleteCallback != nil {
			staged := any(staged).(*A_RELATION_GROUP_TYPE_REF)
			stage.OnAfterA_RELATION_GROUP_TYPE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SOURCE_1:
		if stage.OnAfterA_SOURCE_1DeleteCallback != nil {
			staged := any(staged).(*A_SOURCE_1)
			stage.OnAfterA_SOURCE_1DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SOURCE_SPECIFICATION_1:
		if stage.OnAfterA_SOURCE_SPECIFICATION_1DeleteCallback != nil {
			staged := any(staged).(*A_SOURCE_SPECIFICATION_1)
			stage.OnAfterA_SOURCE_SPECIFICATION_1DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPECIFICATIONS:
		if stage.OnAfterA_SPECIFICATIONSDeleteCallback != nil {
			staged := any(staged).(*A_SPECIFICATIONS)
			stage.OnAfterA_SPECIFICATIONSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPECIFICATION_TYPE_REF:
		if stage.OnAfterA_SPECIFICATION_TYPE_REFDeleteCallback != nil {
			staged := any(staged).(*A_SPECIFICATION_TYPE_REF)
			stage.OnAfterA_SPECIFICATION_TYPE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPECIFIED_VALUES:
		if stage.OnAfterA_SPECIFIED_VALUESDeleteCallback != nil {
			staged := any(staged).(*A_SPECIFIED_VALUES)
			stage.OnAfterA_SPECIFIED_VALUESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_ATTRIBUTES:
		if stage.OnAfterA_SPEC_ATTRIBUTESDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_ATTRIBUTES)
			stage.OnAfterA_SPEC_ATTRIBUTESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_OBJECTS:
		if stage.OnAfterA_SPEC_OBJECTSDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_OBJECTS)
			stage.OnAfterA_SPEC_OBJECTSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_OBJECT_TYPE_REF:
		if stage.OnAfterA_SPEC_OBJECT_TYPE_REFDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_OBJECT_TYPE_REF)
			stage.OnAfterA_SPEC_OBJECT_TYPE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_RELATIONS:
		if stage.OnAfterA_SPEC_RELATIONSDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_RELATIONS)
			stage.OnAfterA_SPEC_RELATIONSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_RELATION_GROUPS:
		if stage.OnAfterA_SPEC_RELATION_GROUPSDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_RELATION_GROUPS)
			stage.OnAfterA_SPEC_RELATION_GROUPSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_RELATION_REF:
		if stage.OnAfterA_SPEC_RELATION_REFDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_RELATION_REF)
			stage.OnAfterA_SPEC_RELATION_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_RELATION_TYPE_REF:
		if stage.OnAfterA_SPEC_RELATION_TYPE_REFDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_RELATION_TYPE_REF)
			stage.OnAfterA_SPEC_RELATION_TYPE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_TYPES:
		if stage.OnAfterA_SPEC_TYPESDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_TYPES)
			stage.OnAfterA_SPEC_TYPESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_THE_HEADER:
		if stage.OnAfterA_THE_HEADERDeleteCallback != nil {
			staged := any(staged).(*A_THE_HEADER)
			stage.OnAfterA_THE_HEADERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_TOOL_EXTENSIONS:
		if stage.OnAfterA_TOOL_EXTENSIONSDeleteCallback != nil {
			staged := any(staged).(*A_TOOL_EXTENSIONS)
			stage.OnAfterA_TOOL_EXTENSIONSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_BOOLEAN)
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_DATE:
		if stage.OnAfterDATATYPE_DEFINITION_DATEDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_DATE)
			stage.OnAfterDATATYPE_DEFINITION_DATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_ENUMERATION)
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_INTEGER:
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_INTEGER)
			stage.OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_REAL:
		if stage.OnAfterDATATYPE_DEFINITION_REALDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_REAL)
			stage.OnAfterDATATYPE_DEFINITION_REALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_STRING:
		if stage.OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_STRING)
			stage.OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_XHTML:
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_XHTML)
			stage.OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EMBEDDED_VALUE:
		if stage.OnAfterEMBEDDED_VALUEDeleteCallback != nil {
			staged := any(staged).(*EMBEDDED_VALUE)
			stage.OnAfterEMBEDDED_VALUEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ENUM_VALUE:
		if stage.OnAfterENUM_VALUEDeleteCallback != nil {
			staged := any(staged).(*ENUM_VALUE)
			stage.OnAfterENUM_VALUEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDeleteCallback != nil {
			staged := any(staged).(*Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry)
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_SPECIFICATION_Nodes_expandedEntry:
		if stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryDeleteCallback != nil {
			staged := any(staged).(*Map_SPECIFICATION_Nodes_expandedEntry)
			stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryDeleteCallback != nil {
			staged := any(staged).(*Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry)
			stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryDeleteCallback != nil {
			staged := any(staged).(*Map_SPEC_OBJECT_TYPE_showIdentifierEntry)
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryDeleteCallback != nil {
			staged := any(staged).(*Map_SPEC_OBJECT_TYPE_showNameEntry)
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RELATION_GROUP:
		if stage.OnAfterRELATION_GROUPDeleteCallback != nil {
			staged := any(staged).(*RELATION_GROUP)
			stage.OnAfterRELATION_GROUPDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RELATION_GROUP_TYPE:
		if stage.OnAfterRELATION_GROUP_TYPEDeleteCallback != nil {
			staged := any(staged).(*RELATION_GROUP_TYPE)
			stage.OnAfterRELATION_GROUP_TYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF:
		if stage.OnAfterREQ_IFDeleteCallback != nil {
			staged := any(staged).(*REQ_IF)
			stage.OnAfterREQ_IFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF_CONTENT:
		if stage.OnAfterREQ_IF_CONTENTDeleteCallback != nil {
			staged := any(staged).(*REQ_IF_CONTENT)
			stage.OnAfterREQ_IF_CONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF_HEADER:
		if stage.OnAfterREQ_IF_HEADERDeleteCallback != nil {
			staged := any(staged).(*REQ_IF_HEADER)
			stage.OnAfterREQ_IF_HEADERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF_TOOL_EXTENSION:
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback != nil {
			staged := any(staged).(*REQ_IF_TOOL_EXTENSION)
			stage.OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RenderingConfiguration:
		if stage.OnAfterRenderingConfigurationDeleteCallback != nil {
			staged := any(staged).(*RenderingConfiguration)
			stage.OnAfterRenderingConfigurationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONDeleteCallback != nil {
			staged := any(staged).(*SPECIFICATION)
			stage.OnAfterSPECIFICATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECIFICATION_TYPE:
		if stage.OnAfterSPECIFICATION_TYPEDeleteCallback != nil {
			staged := any(staged).(*SPECIFICATION_TYPE)
			stage.OnAfterSPECIFICATION_TYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_HIERARCHY:
		if stage.OnAfterSPEC_HIERARCHYDeleteCallback != nil {
			staged := any(staged).(*SPEC_HIERARCHY)
			stage.OnAfterSPEC_HIERARCHYDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_OBJECT:
		if stage.OnAfterSPEC_OBJECTDeleteCallback != nil {
			staged := any(staged).(*SPEC_OBJECT)
			stage.OnAfterSPEC_OBJECTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_OBJECT_TYPE:
		if stage.OnAfterSPEC_OBJECT_TYPEDeleteCallback != nil {
			staged := any(staged).(*SPEC_OBJECT_TYPE)
			stage.OnAfterSPEC_OBJECT_TYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_RELATION:
		if stage.OnAfterSPEC_RELATIONDeleteCallback != nil {
			staged := any(staged).(*SPEC_RELATION)
			stage.OnAfterSPEC_RELATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_RELATION_TYPE:
		if stage.OnAfterSPEC_RELATION_TYPEDeleteCallback != nil {
			staged := any(staged).(*SPEC_RELATION_TYPE)
			stage.OnAfterSPEC_RELATION_TYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StaticWebSite:
		if stage.OnAfterStaticWebSiteDeleteCallback != nil {
			staged := any(staged).(*StaticWebSite)
			stage.OnAfterStaticWebSiteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StaticWebSiteChapter:
		if stage.OnAfterStaticWebSiteChapterDeleteCallback != nil {
			staged := any(staged).(*StaticWebSiteChapter)
			stage.OnAfterStaticWebSiteChapterDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StaticWebSiteGeneratedImage:
		if stage.OnAfterStaticWebSiteGeneratedImageDeleteCallback != nil {
			staged := any(staged).(*StaticWebSiteGeneratedImage)
			stage.OnAfterStaticWebSiteGeneratedImageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StaticWebSiteImage:
		if stage.OnAfterStaticWebSiteImageDeleteCallback != nil {
			staged := any(staged).(*StaticWebSiteImage)
			stage.OnAfterStaticWebSiteImageDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StaticWebSiteParagraph:
		if stage.OnAfterStaticWebSiteParagraphDeleteCallback != nil {
			staged := any(staged).(*StaticWebSiteParagraph)
			stage.OnAfterStaticWebSiteParagraphDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *XHTML_CONTENT:
		if stage.OnAfterXHTML_CONTENTDeleteCallback != nil {
			staged := any(staged).(*XHTML_CONTENT)
			stage.OnAfterXHTML_CONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		if stage.OnAfterALTERNATIVE_IDReadCallback != nil {
			stage.OnAfterALTERNATIVE_IDReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		if stage.OnAfterATTRIBUTE_DEFINITION_DATEReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_DATEReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		if stage.OnAfterATTRIBUTE_DEFINITION_REALReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_REALReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterATTRIBUTE_VALUE_DATEReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_DATEReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_INTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterATTRIBUTE_VALUE_REALReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_REALReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterATTRIBUTE_VALUE_STRINGReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_STRINGReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_XHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *A_ALTERNATIVE_ID:
		if stage.OnAfterA_ALTERNATIVE_IDReadCallback != nil {
			stage.OnAfterA_ALTERNATIVE_IDReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterA_ATTRIBUTE_VALUE_DATEReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_DATEReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterA_ATTRIBUTE_VALUE_REALReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_REALReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_STRINGReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1ReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1ReadCallback.OnAfterRead(stage, target)
		}
	case *A_CHILDREN:
		if stage.OnAfterA_CHILDRENReadCallback != nil {
			stage.OnAfterA_CHILDRENReadCallback.OnAfterRead(stage, target)
		}
	case *A_CORE_CONTENT:
		if stage.OnAfterA_CORE_CONTENTReadCallback != nil {
			stage.OnAfterA_CORE_CONTENTReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPES:
		if stage.OnAfterA_DATATYPESReadCallback != nil {
			stage.OnAfterA_DATATYPESReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_EDITABLE_ATTS:
		if stage.OnAfterA_EDITABLE_ATTSReadCallback != nil {
			stage.OnAfterA_EDITABLE_ATTSReadCallback.OnAfterRead(stage, target)
		}
	case *A_ENUM_VALUE_REF:
		if stage.OnAfterA_ENUM_VALUE_REFReadCallback != nil {
			stage.OnAfterA_ENUM_VALUE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_OBJECT:
		if stage.OnAfterA_OBJECTReadCallback != nil {
			stage.OnAfterA_OBJECTReadCallback.OnAfterRead(stage, target)
		}
	case *A_PROPERTIES:
		if stage.OnAfterA_PROPERTIESReadCallback != nil {
			stage.OnAfterA_PROPERTIESReadCallback.OnAfterRead(stage, target)
		}
	case *A_RELATION_GROUP_TYPE_REF:
		if stage.OnAfterA_RELATION_GROUP_TYPE_REFReadCallback != nil {
			stage.OnAfterA_RELATION_GROUP_TYPE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SOURCE_1:
		if stage.OnAfterA_SOURCE_1ReadCallback != nil {
			stage.OnAfterA_SOURCE_1ReadCallback.OnAfterRead(stage, target)
		}
	case *A_SOURCE_SPECIFICATION_1:
		if stage.OnAfterA_SOURCE_SPECIFICATION_1ReadCallback != nil {
			stage.OnAfterA_SOURCE_SPECIFICATION_1ReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPECIFICATIONS:
		if stage.OnAfterA_SPECIFICATIONSReadCallback != nil {
			stage.OnAfterA_SPECIFICATIONSReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPECIFICATION_TYPE_REF:
		if stage.OnAfterA_SPECIFICATION_TYPE_REFReadCallback != nil {
			stage.OnAfterA_SPECIFICATION_TYPE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPECIFIED_VALUES:
		if stage.OnAfterA_SPECIFIED_VALUESReadCallback != nil {
			stage.OnAfterA_SPECIFIED_VALUESReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_ATTRIBUTES:
		if stage.OnAfterA_SPEC_ATTRIBUTESReadCallback != nil {
			stage.OnAfterA_SPEC_ATTRIBUTESReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_OBJECTS:
		if stage.OnAfterA_SPEC_OBJECTSReadCallback != nil {
			stage.OnAfterA_SPEC_OBJECTSReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_OBJECT_TYPE_REF:
		if stage.OnAfterA_SPEC_OBJECT_TYPE_REFReadCallback != nil {
			stage.OnAfterA_SPEC_OBJECT_TYPE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_RELATIONS:
		if stage.OnAfterA_SPEC_RELATIONSReadCallback != nil {
			stage.OnAfterA_SPEC_RELATIONSReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_RELATION_GROUPS:
		if stage.OnAfterA_SPEC_RELATION_GROUPSReadCallback != nil {
			stage.OnAfterA_SPEC_RELATION_GROUPSReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_RELATION_REF:
		if stage.OnAfterA_SPEC_RELATION_REFReadCallback != nil {
			stage.OnAfterA_SPEC_RELATION_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_RELATION_TYPE_REF:
		if stage.OnAfterA_SPEC_RELATION_TYPE_REFReadCallback != nil {
			stage.OnAfterA_SPEC_RELATION_TYPE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_TYPES:
		if stage.OnAfterA_SPEC_TYPESReadCallback != nil {
			stage.OnAfterA_SPEC_TYPESReadCallback.OnAfterRead(stage, target)
		}
	case *A_THE_HEADER:
		if stage.OnAfterA_THE_HEADERReadCallback != nil {
			stage.OnAfterA_THE_HEADERReadCallback.OnAfterRead(stage, target)
		}
	case *A_TOOL_EXTENSIONS:
		if stage.OnAfterA_TOOL_EXTENSIONSReadCallback != nil {
			stage.OnAfterA_TOOL_EXTENSIONSReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_DATE:
		if stage.OnAfterDATATYPE_DEFINITION_DATEReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_DATEReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_INTEGER:
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_INTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_REAL:
		if stage.OnAfterDATATYPE_DEFINITION_REALReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_REALReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_STRING:
		if stage.OnAfterDATATYPE_DEFINITION_STRINGReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_STRINGReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_XHTML:
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_XHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *EMBEDDED_VALUE:
		if stage.OnAfterEMBEDDED_VALUEReadCallback != nil {
			stage.OnAfterEMBEDDED_VALUEReadCallback.OnAfterRead(stage, target)
		}
	case *ENUM_VALUE:
		if stage.OnAfterENUM_VALUEReadCallback != nil {
			stage.OnAfterENUM_VALUEReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		if stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryReadCallback != nil {
			stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_SPECIFICATION_Nodes_expandedEntry:
		if stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryReadCallback != nil {
			stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryReadCallback != nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryReadCallback != nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryReadCallback.OnAfterRead(stage, target)
		}
	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		if stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryReadCallback != nil {
			stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryReadCallback.OnAfterRead(stage, target)
		}
	case *RELATION_GROUP:
		if stage.OnAfterRELATION_GROUPReadCallback != nil {
			stage.OnAfterRELATION_GROUPReadCallback.OnAfterRead(stage, target)
		}
	case *RELATION_GROUP_TYPE:
		if stage.OnAfterRELATION_GROUP_TYPEReadCallback != nil {
			stage.OnAfterRELATION_GROUP_TYPEReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF:
		if stage.OnAfterREQ_IFReadCallback != nil {
			stage.OnAfterREQ_IFReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF_CONTENT:
		if stage.OnAfterREQ_IF_CONTENTReadCallback != nil {
			stage.OnAfterREQ_IF_CONTENTReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF_HEADER:
		if stage.OnAfterREQ_IF_HEADERReadCallback != nil {
			stage.OnAfterREQ_IF_HEADERReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF_TOOL_EXTENSION:
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONReadCallback != nil {
			stage.OnAfterREQ_IF_TOOL_EXTENSIONReadCallback.OnAfterRead(stage, target)
		}
	case *RenderingConfiguration:
		if stage.OnAfterRenderingConfigurationReadCallback != nil {
			stage.OnAfterRenderingConfigurationReadCallback.OnAfterRead(stage, target)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONReadCallback != nil {
			stage.OnAfterSPECIFICATIONReadCallback.OnAfterRead(stage, target)
		}
	case *SPECIFICATION_TYPE:
		if stage.OnAfterSPECIFICATION_TYPEReadCallback != nil {
			stage.OnAfterSPECIFICATION_TYPEReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_HIERARCHY:
		if stage.OnAfterSPEC_HIERARCHYReadCallback != nil {
			stage.OnAfterSPEC_HIERARCHYReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_OBJECT:
		if stage.OnAfterSPEC_OBJECTReadCallback != nil {
			stage.OnAfterSPEC_OBJECTReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_OBJECT_TYPE:
		if stage.OnAfterSPEC_OBJECT_TYPEReadCallback != nil {
			stage.OnAfterSPEC_OBJECT_TYPEReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_RELATION:
		if stage.OnAfterSPEC_RELATIONReadCallback != nil {
			stage.OnAfterSPEC_RELATIONReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_RELATION_TYPE:
		if stage.OnAfterSPEC_RELATION_TYPEReadCallback != nil {
			stage.OnAfterSPEC_RELATION_TYPEReadCallback.OnAfterRead(stage, target)
		}
	case *StaticWebSite:
		if stage.OnAfterStaticWebSiteReadCallback != nil {
			stage.OnAfterStaticWebSiteReadCallback.OnAfterRead(stage, target)
		}
	case *StaticWebSiteChapter:
		if stage.OnAfterStaticWebSiteChapterReadCallback != nil {
			stage.OnAfterStaticWebSiteChapterReadCallback.OnAfterRead(stage, target)
		}
	case *StaticWebSiteGeneratedImage:
		if stage.OnAfterStaticWebSiteGeneratedImageReadCallback != nil {
			stage.OnAfterStaticWebSiteGeneratedImageReadCallback.OnAfterRead(stage, target)
		}
	case *StaticWebSiteImage:
		if stage.OnAfterStaticWebSiteImageReadCallback != nil {
			stage.OnAfterStaticWebSiteImageReadCallback.OnAfterRead(stage, target)
		}
	case *StaticWebSiteParagraph:
		if stage.OnAfterStaticWebSiteParagraphReadCallback != nil {
			stage.OnAfterStaticWebSiteParagraphReadCallback.OnAfterRead(stage, target)
		}
	case *XHTML_CONTENT:
		if stage.OnAfterXHTML_CONTENTReadCallback != nil {
			stage.OnAfterXHTML_CONTENTReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVE_ID:
		stage.OnAfterALTERNATIVE_IDUpdateCallback = any(callback).(OnAfterUpdateInterface[ALTERNATIVE_ID])
	
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_BOOLEAN])
	
	case *ATTRIBUTE_DEFINITION_DATE:
		stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_DATE])
	
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_ENUMERATION])
	
	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_INTEGER])
	
	case *ATTRIBUTE_DEFINITION_REAL:
		stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_REAL])
	
	case *ATTRIBUTE_DEFINITION_STRING:
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_STRING])
	
	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_XHTML])
	
	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_BOOLEAN])
	
	case *ATTRIBUTE_VALUE_DATE:
		stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_DATE])
	
	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_ENUMERATION])
	
	case *ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_INTEGER])
	
	case *ATTRIBUTE_VALUE_REAL:
		stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_REAL])
	
	case *ATTRIBUTE_VALUE_STRING:
		stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_STRING])
	
	case *ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_XHTML])
	
	case *A_ALTERNATIVE_ID:
		stage.OnAfterA_ALTERNATIVE_IDUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ALTERNATIVE_ID])
	
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF])
	
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_DATE_REF])
	
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF])
	
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF])
	
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_REAL_REF])
	
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_STRING_REF])
	
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF])
	
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_BOOLEAN])
	
	case *A_ATTRIBUTE_VALUE_DATE:
		stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_DATE])
	
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_ENUMERATION])
	
	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_INTEGER])
	
	case *A_ATTRIBUTE_VALUE_REAL:
		stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_REAL])
	
	case *A_ATTRIBUTE_VALUE_STRING:
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_STRING])
	
	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_XHTML])
	
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_XHTML_1])
	
	case *A_CHILDREN:
		stage.OnAfterA_CHILDRENUpdateCallback = any(callback).(OnAfterUpdateInterface[A_CHILDREN])
	
	case *A_CORE_CONTENT:
		stage.OnAfterA_CORE_CONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[A_CORE_CONTENT])
	
	case *A_DATATYPES:
		stage.OnAfterA_DATATYPESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPES])
	
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF])
	
	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_DATE_REF])
	
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF])
	
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_INTEGER_REF])
	
	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_REAL_REF])
	
	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_STRING_REF])
	
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_XHTML_REF])
	
	case *A_EDITABLE_ATTS:
		stage.OnAfterA_EDITABLE_ATTSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_EDITABLE_ATTS])
	
	case *A_ENUM_VALUE_REF:
		stage.OnAfterA_ENUM_VALUE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ENUM_VALUE_REF])
	
	case *A_OBJECT:
		stage.OnAfterA_OBJECTUpdateCallback = any(callback).(OnAfterUpdateInterface[A_OBJECT])
	
	case *A_PROPERTIES:
		stage.OnAfterA_PROPERTIESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_PROPERTIES])
	
	case *A_RELATION_GROUP_TYPE_REF:
		stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_RELATION_GROUP_TYPE_REF])
	
	case *A_SOURCE_1:
		stage.OnAfterA_SOURCE_1UpdateCallback = any(callback).(OnAfterUpdateInterface[A_SOURCE_1])
	
	case *A_SOURCE_SPECIFICATION_1:
		stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateCallback = any(callback).(OnAfterUpdateInterface[A_SOURCE_SPECIFICATION_1])
	
	case *A_SPECIFICATIONS:
		stage.OnAfterA_SPECIFICATIONSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPECIFICATIONS])
	
	case *A_SPECIFICATION_TYPE_REF:
		stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPECIFICATION_TYPE_REF])
	
	case *A_SPECIFIED_VALUES:
		stage.OnAfterA_SPECIFIED_VALUESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPECIFIED_VALUES])
	
	case *A_SPEC_ATTRIBUTES:
		stage.OnAfterA_SPEC_ATTRIBUTESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_ATTRIBUTES])
	
	case *A_SPEC_OBJECTS:
		stage.OnAfterA_SPEC_OBJECTSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_OBJECTS])
	
	case *A_SPEC_OBJECT_TYPE_REF:
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_OBJECT_TYPE_REF])
	
	case *A_SPEC_RELATIONS:
		stage.OnAfterA_SPEC_RELATIONSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_RELATIONS])
	
	case *A_SPEC_RELATION_GROUPS:
		stage.OnAfterA_SPEC_RELATION_GROUPSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_RELATION_GROUPS])
	
	case *A_SPEC_RELATION_REF:
		stage.OnAfterA_SPEC_RELATION_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_RELATION_REF])
	
	case *A_SPEC_RELATION_TYPE_REF:
		stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_RELATION_TYPE_REF])
	
	case *A_SPEC_TYPES:
		stage.OnAfterA_SPEC_TYPESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_TYPES])
	
	case *A_THE_HEADER:
		stage.OnAfterA_THE_HEADERUpdateCallback = any(callback).(OnAfterUpdateInterface[A_THE_HEADER])
	
	case *A_TOOL_EXTENSIONS:
		stage.OnAfterA_TOOL_EXTENSIONSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_TOOL_EXTENSIONS])
	
	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_BOOLEAN])
	
	case *DATATYPE_DEFINITION_DATE:
		stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_DATE])
	
	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_ENUMERATION])
	
	case *DATATYPE_DEFINITION_INTEGER:
		stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_INTEGER])
	
	case *DATATYPE_DEFINITION_REAL:
		stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_REAL])
	
	case *DATATYPE_DEFINITION_STRING:
		stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_STRING])
	
	case *DATATYPE_DEFINITION_XHTML:
		stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_XHTML])
	
	case *EMBEDDED_VALUE:
		stage.OnAfterEMBEDDED_VALUEUpdateCallback = any(callback).(OnAfterUpdateInterface[EMBEDDED_VALUE])
	
	case *ENUM_VALUE:
		stage.OnAfterENUM_VALUEUpdateCallback = any(callback).(OnAfterUpdateInterface[ENUM_VALUE])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry])
	
	case *Map_SPECIFICATION_Nodes_expandedEntry:
		stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_SPECIFICATION_Nodes_expandedEntry])
	
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry])
	
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_SPEC_OBJECT_TYPE_showIdentifierEntry])
	
	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryUpdateCallback = any(callback).(OnAfterUpdateInterface[Map_SPEC_OBJECT_TYPE_showNameEntry])
	
	case *RELATION_GROUP:
		stage.OnAfterRELATION_GROUPUpdateCallback = any(callback).(OnAfterUpdateInterface[RELATION_GROUP])
	
	case *RELATION_GROUP_TYPE:
		stage.OnAfterRELATION_GROUP_TYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[RELATION_GROUP_TYPE])
	
	case *REQ_IF:
		stage.OnAfterREQ_IFUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF])
	
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF_CONTENT])
	
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF_HEADER])
	
	case *REQ_IF_TOOL_EXTENSION:
		stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF_TOOL_EXTENSION])
	
	case *RenderingConfiguration:
		stage.OnAfterRenderingConfigurationUpdateCallback = any(callback).(OnAfterUpdateInterface[RenderingConfiguration])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECIFICATION])
	
	case *SPECIFICATION_TYPE:
		stage.OnAfterSPECIFICATION_TYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECIFICATION_TYPE])
	
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_HIERARCHY])
	
	case *SPEC_OBJECT:
		stage.OnAfterSPEC_OBJECTUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_OBJECT])
	
	case *SPEC_OBJECT_TYPE:
		stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_OBJECT_TYPE])
	
	case *SPEC_RELATION:
		stage.OnAfterSPEC_RELATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_RELATION])
	
	case *SPEC_RELATION_TYPE:
		stage.OnAfterSPEC_RELATION_TYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_RELATION_TYPE])
	
	case *StaticWebSite:
		stage.OnAfterStaticWebSiteUpdateCallback = any(callback).(OnAfterUpdateInterface[StaticWebSite])
	
	case *StaticWebSiteChapter:
		stage.OnAfterStaticWebSiteChapterUpdateCallback = any(callback).(OnAfterUpdateInterface[StaticWebSiteChapter])
	
	case *StaticWebSiteGeneratedImage:
		stage.OnAfterStaticWebSiteGeneratedImageUpdateCallback = any(callback).(OnAfterUpdateInterface[StaticWebSiteGeneratedImage])
	
	case *StaticWebSiteImage:
		stage.OnAfterStaticWebSiteImageUpdateCallback = any(callback).(OnAfterUpdateInterface[StaticWebSiteImage])
	
	case *StaticWebSiteParagraph:
		stage.OnAfterStaticWebSiteParagraphUpdateCallback = any(callback).(OnAfterUpdateInterface[StaticWebSiteParagraph])
	
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[XHTML_CONTENT])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVE_ID:
		stage.OnAfterALTERNATIVE_IDCreateCallback = any(callback).(OnAfterCreateInterface[ALTERNATIVE_ID])
	
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_BOOLEAN])
	
	case *ATTRIBUTE_DEFINITION_DATE:
		stage.OnAfterATTRIBUTE_DEFINITION_DATECreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_DATE])
	
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_ENUMERATION])
	
	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_INTEGER])
	
	case *ATTRIBUTE_DEFINITION_REAL:
		stage.OnAfterATTRIBUTE_DEFINITION_REALCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_REAL])
	
	case *ATTRIBUTE_DEFINITION_STRING:
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_STRING])
	
	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_XHTML])
	
	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_BOOLEAN])
	
	case *ATTRIBUTE_VALUE_DATE:
		stage.OnAfterATTRIBUTE_VALUE_DATECreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_DATE])
	
	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_ENUMERATION])
	
	case *ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_INTEGER])
	
	case *ATTRIBUTE_VALUE_REAL:
		stage.OnAfterATTRIBUTE_VALUE_REALCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_REAL])
	
	case *ATTRIBUTE_VALUE_STRING:
		stage.OnAfterATTRIBUTE_VALUE_STRINGCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_STRING])
	
	case *ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_XHTML])
	
	case *A_ALTERNATIVE_ID:
		stage.OnAfterA_ALTERNATIVE_IDCreateCallback = any(callback).(OnAfterCreateInterface[A_ALTERNATIVE_ID])
	
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF])
	
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_DATE_REF])
	
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF])
	
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF])
	
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_REAL_REF])
	
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_STRING_REF])
	
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF])
	
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_BOOLEAN])
	
	case *A_ATTRIBUTE_VALUE_DATE:
		stage.OnAfterA_ATTRIBUTE_VALUE_DATECreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_DATE])
	
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_ENUMERATION])
	
	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_INTEGER])
	
	case *A_ATTRIBUTE_VALUE_REAL:
		stage.OnAfterA_ATTRIBUTE_VALUE_REALCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_REAL])
	
	case *A_ATTRIBUTE_VALUE_STRING:
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_STRING])
	
	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_XHTML])
	
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1CreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_XHTML_1])
	
	case *A_CHILDREN:
		stage.OnAfterA_CHILDRENCreateCallback = any(callback).(OnAfterCreateInterface[A_CHILDREN])
	
	case *A_CORE_CONTENT:
		stage.OnAfterA_CORE_CONTENTCreateCallback = any(callback).(OnAfterCreateInterface[A_CORE_CONTENT])
	
	case *A_DATATYPES:
		stage.OnAfterA_DATATYPESCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPES])
	
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF])
	
	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_DATE_REF])
	
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF])
	
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_INTEGER_REF])
	
	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_REAL_REF])
	
	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_STRING_REF])
	
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_XHTML_REF])
	
	case *A_EDITABLE_ATTS:
		stage.OnAfterA_EDITABLE_ATTSCreateCallback = any(callback).(OnAfterCreateInterface[A_EDITABLE_ATTS])
	
	case *A_ENUM_VALUE_REF:
		stage.OnAfterA_ENUM_VALUE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ENUM_VALUE_REF])
	
	case *A_OBJECT:
		stage.OnAfterA_OBJECTCreateCallback = any(callback).(OnAfterCreateInterface[A_OBJECT])
	
	case *A_PROPERTIES:
		stage.OnAfterA_PROPERTIESCreateCallback = any(callback).(OnAfterCreateInterface[A_PROPERTIES])
	
	case *A_RELATION_GROUP_TYPE_REF:
		stage.OnAfterA_RELATION_GROUP_TYPE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_RELATION_GROUP_TYPE_REF])
	
	case *A_SOURCE_1:
		stage.OnAfterA_SOURCE_1CreateCallback = any(callback).(OnAfterCreateInterface[A_SOURCE_1])
	
	case *A_SOURCE_SPECIFICATION_1:
		stage.OnAfterA_SOURCE_SPECIFICATION_1CreateCallback = any(callback).(OnAfterCreateInterface[A_SOURCE_SPECIFICATION_1])
	
	case *A_SPECIFICATIONS:
		stage.OnAfterA_SPECIFICATIONSCreateCallback = any(callback).(OnAfterCreateInterface[A_SPECIFICATIONS])
	
	case *A_SPECIFICATION_TYPE_REF:
		stage.OnAfterA_SPECIFICATION_TYPE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_SPECIFICATION_TYPE_REF])
	
	case *A_SPECIFIED_VALUES:
		stage.OnAfterA_SPECIFIED_VALUESCreateCallback = any(callback).(OnAfterCreateInterface[A_SPECIFIED_VALUES])
	
	case *A_SPEC_ATTRIBUTES:
		stage.OnAfterA_SPEC_ATTRIBUTESCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_ATTRIBUTES])
	
	case *A_SPEC_OBJECTS:
		stage.OnAfterA_SPEC_OBJECTSCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_OBJECTS])
	
	case *A_SPEC_OBJECT_TYPE_REF:
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_OBJECT_TYPE_REF])
	
	case *A_SPEC_RELATIONS:
		stage.OnAfterA_SPEC_RELATIONSCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_RELATIONS])
	
	case *A_SPEC_RELATION_GROUPS:
		stage.OnAfterA_SPEC_RELATION_GROUPSCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_RELATION_GROUPS])
	
	case *A_SPEC_RELATION_REF:
		stage.OnAfterA_SPEC_RELATION_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_RELATION_REF])
	
	case *A_SPEC_RELATION_TYPE_REF:
		stage.OnAfterA_SPEC_RELATION_TYPE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_RELATION_TYPE_REF])
	
	case *A_SPEC_TYPES:
		stage.OnAfterA_SPEC_TYPESCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_TYPES])
	
	case *A_THE_HEADER:
		stage.OnAfterA_THE_HEADERCreateCallback = any(callback).(OnAfterCreateInterface[A_THE_HEADER])
	
	case *A_TOOL_EXTENSIONS:
		stage.OnAfterA_TOOL_EXTENSIONSCreateCallback = any(callback).(OnAfterCreateInterface[A_TOOL_EXTENSIONS])
	
	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_BOOLEAN])
	
	case *DATATYPE_DEFINITION_DATE:
		stage.OnAfterDATATYPE_DEFINITION_DATECreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_DATE])
	
	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_ENUMERATION])
	
	case *DATATYPE_DEFINITION_INTEGER:
		stage.OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_INTEGER])
	
	case *DATATYPE_DEFINITION_REAL:
		stage.OnAfterDATATYPE_DEFINITION_REALCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_REAL])
	
	case *DATATYPE_DEFINITION_STRING:
		stage.OnAfterDATATYPE_DEFINITION_STRINGCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_STRING])
	
	case *DATATYPE_DEFINITION_XHTML:
		stage.OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_XHTML])
	
	case *EMBEDDED_VALUE:
		stage.OnAfterEMBEDDED_VALUECreateCallback = any(callback).(OnAfterCreateInterface[EMBEDDED_VALUE])
	
	case *ENUM_VALUE:
		stage.OnAfterENUM_VALUECreateCallback = any(callback).(OnAfterCreateInterface[ENUM_VALUE])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry])
	
	case *Map_SPECIFICATION_Nodes_expandedEntry:
		stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_SPECIFICATION_Nodes_expandedEntry])
	
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry])
	
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_SPEC_OBJECT_TYPE_showIdentifierEntry])
	
	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryCreateCallback = any(callback).(OnAfterCreateInterface[Map_SPEC_OBJECT_TYPE_showNameEntry])
	
	case *RELATION_GROUP:
		stage.OnAfterRELATION_GROUPCreateCallback = any(callback).(OnAfterCreateInterface[RELATION_GROUP])
	
	case *RELATION_GROUP_TYPE:
		stage.OnAfterRELATION_GROUP_TYPECreateCallback = any(callback).(OnAfterCreateInterface[RELATION_GROUP_TYPE])
	
	case *REQ_IF:
		stage.OnAfterREQ_IFCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF])
	
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF_CONTENT])
	
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF_HEADER])
	
	case *REQ_IF_TOOL_EXTENSION:
		stage.OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF_TOOL_EXTENSION])
	
	case *RenderingConfiguration:
		stage.OnAfterRenderingConfigurationCreateCallback = any(callback).(OnAfterCreateInterface[RenderingConfiguration])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONCreateCallback = any(callback).(OnAfterCreateInterface[SPECIFICATION])
	
	case *SPECIFICATION_TYPE:
		stage.OnAfterSPECIFICATION_TYPECreateCallback = any(callback).(OnAfterCreateInterface[SPECIFICATION_TYPE])
	
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYCreateCallback = any(callback).(OnAfterCreateInterface[SPEC_HIERARCHY])
	
	case *SPEC_OBJECT:
		stage.OnAfterSPEC_OBJECTCreateCallback = any(callback).(OnAfterCreateInterface[SPEC_OBJECT])
	
	case *SPEC_OBJECT_TYPE:
		stage.OnAfterSPEC_OBJECT_TYPECreateCallback = any(callback).(OnAfterCreateInterface[SPEC_OBJECT_TYPE])
	
	case *SPEC_RELATION:
		stage.OnAfterSPEC_RELATIONCreateCallback = any(callback).(OnAfterCreateInterface[SPEC_RELATION])
	
	case *SPEC_RELATION_TYPE:
		stage.OnAfterSPEC_RELATION_TYPECreateCallback = any(callback).(OnAfterCreateInterface[SPEC_RELATION_TYPE])
	
	case *StaticWebSite:
		stage.OnAfterStaticWebSiteCreateCallback = any(callback).(OnAfterCreateInterface[StaticWebSite])
	
	case *StaticWebSiteChapter:
		stage.OnAfterStaticWebSiteChapterCreateCallback = any(callback).(OnAfterCreateInterface[StaticWebSiteChapter])
	
	case *StaticWebSiteGeneratedImage:
		stage.OnAfterStaticWebSiteGeneratedImageCreateCallback = any(callback).(OnAfterCreateInterface[StaticWebSiteGeneratedImage])
	
	case *StaticWebSiteImage:
		stage.OnAfterStaticWebSiteImageCreateCallback = any(callback).(OnAfterCreateInterface[StaticWebSiteImage])
	
	case *StaticWebSiteParagraph:
		stage.OnAfterStaticWebSiteParagraphCreateCallback = any(callback).(OnAfterCreateInterface[StaticWebSiteParagraph])
	
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTCreateCallback = any(callback).(OnAfterCreateInterface[XHTML_CONTENT])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVE_ID:
		stage.OnAfterALTERNATIVE_IDDeleteCallback = any(callback).(OnAfterDeleteInterface[ALTERNATIVE_ID])
	
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_BOOLEAN])
	
	case *ATTRIBUTE_DEFINITION_DATE:
		stage.OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_DATE])
	
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_ENUMERATION])
	
	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_INTEGER])
	
	case *ATTRIBUTE_DEFINITION_REAL:
		stage.OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_REAL])
	
	case *ATTRIBUTE_DEFINITION_STRING:
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_STRING])
	
	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_XHTML])
	
	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_BOOLEAN])
	
	case *ATTRIBUTE_VALUE_DATE:
		stage.OnAfterATTRIBUTE_VALUE_DATEDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_DATE])
	
	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_ENUMERATION])
	
	case *ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_INTEGER])
	
	case *ATTRIBUTE_VALUE_REAL:
		stage.OnAfterATTRIBUTE_VALUE_REALDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_REAL])
	
	case *ATTRIBUTE_VALUE_STRING:
		stage.OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_STRING])
	
	case *ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_XHTML])
	
	case *A_ALTERNATIVE_ID:
		stage.OnAfterA_ALTERNATIVE_IDDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ALTERNATIVE_ID])
	
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF])
	
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_DATE_REF])
	
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF])
	
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF])
	
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_REAL_REF])
	
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_STRING_REF])
	
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF])
	
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_BOOLEAN])
	
	case *A_ATTRIBUTE_VALUE_DATE:
		stage.OnAfterA_ATTRIBUTE_VALUE_DATEDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_DATE])
	
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_ENUMERATION])
	
	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_INTEGER])
	
	case *A_ATTRIBUTE_VALUE_REAL:
		stage.OnAfterA_ATTRIBUTE_VALUE_REALDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_REAL])
	
	case *A_ATTRIBUTE_VALUE_STRING:
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_STRING])
	
	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_XHTML])
	
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1DeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_XHTML_1])
	
	case *A_CHILDREN:
		stage.OnAfterA_CHILDRENDeleteCallback = any(callback).(OnAfterDeleteInterface[A_CHILDREN])
	
	case *A_CORE_CONTENT:
		stage.OnAfterA_CORE_CONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[A_CORE_CONTENT])
	
	case *A_DATATYPES:
		stage.OnAfterA_DATATYPESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPES])
	
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF])
	
	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_DATE_REF])
	
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF])
	
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_INTEGER_REF])
	
	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_REAL_REF])
	
	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_STRING_REF])
	
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_XHTML_REF])
	
	case *A_EDITABLE_ATTS:
		stage.OnAfterA_EDITABLE_ATTSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_EDITABLE_ATTS])
	
	case *A_ENUM_VALUE_REF:
		stage.OnAfterA_ENUM_VALUE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ENUM_VALUE_REF])
	
	case *A_OBJECT:
		stage.OnAfterA_OBJECTDeleteCallback = any(callback).(OnAfterDeleteInterface[A_OBJECT])
	
	case *A_PROPERTIES:
		stage.OnAfterA_PROPERTIESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_PROPERTIES])
	
	case *A_RELATION_GROUP_TYPE_REF:
		stage.OnAfterA_RELATION_GROUP_TYPE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_RELATION_GROUP_TYPE_REF])
	
	case *A_SOURCE_1:
		stage.OnAfterA_SOURCE_1DeleteCallback = any(callback).(OnAfterDeleteInterface[A_SOURCE_1])
	
	case *A_SOURCE_SPECIFICATION_1:
		stage.OnAfterA_SOURCE_SPECIFICATION_1DeleteCallback = any(callback).(OnAfterDeleteInterface[A_SOURCE_SPECIFICATION_1])
	
	case *A_SPECIFICATIONS:
		stage.OnAfterA_SPECIFICATIONSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPECIFICATIONS])
	
	case *A_SPECIFICATION_TYPE_REF:
		stage.OnAfterA_SPECIFICATION_TYPE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPECIFICATION_TYPE_REF])
	
	case *A_SPECIFIED_VALUES:
		stage.OnAfterA_SPECIFIED_VALUESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPECIFIED_VALUES])
	
	case *A_SPEC_ATTRIBUTES:
		stage.OnAfterA_SPEC_ATTRIBUTESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_ATTRIBUTES])
	
	case *A_SPEC_OBJECTS:
		stage.OnAfterA_SPEC_OBJECTSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_OBJECTS])
	
	case *A_SPEC_OBJECT_TYPE_REF:
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_OBJECT_TYPE_REF])
	
	case *A_SPEC_RELATIONS:
		stage.OnAfterA_SPEC_RELATIONSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_RELATIONS])
	
	case *A_SPEC_RELATION_GROUPS:
		stage.OnAfterA_SPEC_RELATION_GROUPSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_RELATION_GROUPS])
	
	case *A_SPEC_RELATION_REF:
		stage.OnAfterA_SPEC_RELATION_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_RELATION_REF])
	
	case *A_SPEC_RELATION_TYPE_REF:
		stage.OnAfterA_SPEC_RELATION_TYPE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_RELATION_TYPE_REF])
	
	case *A_SPEC_TYPES:
		stage.OnAfterA_SPEC_TYPESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_TYPES])
	
	case *A_THE_HEADER:
		stage.OnAfterA_THE_HEADERDeleteCallback = any(callback).(OnAfterDeleteInterface[A_THE_HEADER])
	
	case *A_TOOL_EXTENSIONS:
		stage.OnAfterA_TOOL_EXTENSIONSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_TOOL_EXTENSIONS])
	
	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_BOOLEAN])
	
	case *DATATYPE_DEFINITION_DATE:
		stage.OnAfterDATATYPE_DEFINITION_DATEDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_DATE])
	
	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_ENUMERATION])
	
	case *DATATYPE_DEFINITION_INTEGER:
		stage.OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_INTEGER])
	
	case *DATATYPE_DEFINITION_REAL:
		stage.OnAfterDATATYPE_DEFINITION_REALDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_REAL])
	
	case *DATATYPE_DEFINITION_STRING:
		stage.OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_STRING])
	
	case *DATATYPE_DEFINITION_XHTML:
		stage.OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_XHTML])
	
	case *EMBEDDED_VALUE:
		stage.OnAfterEMBEDDED_VALUEDeleteCallback = any(callback).(OnAfterDeleteInterface[EMBEDDED_VALUE])
	
	case *ENUM_VALUE:
		stage.OnAfterENUM_VALUEDeleteCallback = any(callback).(OnAfterDeleteInterface[ENUM_VALUE])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry])
	
	case *Map_SPECIFICATION_Nodes_expandedEntry:
		stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_SPECIFICATION_Nodes_expandedEntry])
	
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry])
	
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_SPEC_OBJECT_TYPE_showIdentifierEntry])
	
	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryDeleteCallback = any(callback).(OnAfterDeleteInterface[Map_SPEC_OBJECT_TYPE_showNameEntry])
	
	case *RELATION_GROUP:
		stage.OnAfterRELATION_GROUPDeleteCallback = any(callback).(OnAfterDeleteInterface[RELATION_GROUP])
	
	case *RELATION_GROUP_TYPE:
		stage.OnAfterRELATION_GROUP_TYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[RELATION_GROUP_TYPE])
	
	case *REQ_IF:
		stage.OnAfterREQ_IFDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF])
	
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF_CONTENT])
	
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF_HEADER])
	
	case *REQ_IF_TOOL_EXTENSION:
		stage.OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF_TOOL_EXTENSION])
	
	case *RenderingConfiguration:
		stage.OnAfterRenderingConfigurationDeleteCallback = any(callback).(OnAfterDeleteInterface[RenderingConfiguration])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECIFICATION])
	
	case *SPECIFICATION_TYPE:
		stage.OnAfterSPECIFICATION_TYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECIFICATION_TYPE])
	
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_HIERARCHY])
	
	case *SPEC_OBJECT:
		stage.OnAfterSPEC_OBJECTDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_OBJECT])
	
	case *SPEC_OBJECT_TYPE:
		stage.OnAfterSPEC_OBJECT_TYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_OBJECT_TYPE])
	
	case *SPEC_RELATION:
		stage.OnAfterSPEC_RELATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_RELATION])
	
	case *SPEC_RELATION_TYPE:
		stage.OnAfterSPEC_RELATION_TYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_RELATION_TYPE])
	
	case *StaticWebSite:
		stage.OnAfterStaticWebSiteDeleteCallback = any(callback).(OnAfterDeleteInterface[StaticWebSite])
	
	case *StaticWebSiteChapter:
		stage.OnAfterStaticWebSiteChapterDeleteCallback = any(callback).(OnAfterDeleteInterface[StaticWebSiteChapter])
	
	case *StaticWebSiteGeneratedImage:
		stage.OnAfterStaticWebSiteGeneratedImageDeleteCallback = any(callback).(OnAfterDeleteInterface[StaticWebSiteGeneratedImage])
	
	case *StaticWebSiteImage:
		stage.OnAfterStaticWebSiteImageDeleteCallback = any(callback).(OnAfterDeleteInterface[StaticWebSiteImage])
	
	case *StaticWebSiteParagraph:
		stage.OnAfterStaticWebSiteParagraphDeleteCallback = any(callback).(OnAfterDeleteInterface[StaticWebSiteParagraph])
	
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[XHTML_CONTENT])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVE_ID:
		stage.OnAfterALTERNATIVE_IDReadCallback = any(callback).(OnAfterReadInterface[ALTERNATIVE_ID])
	
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_BOOLEAN])
	
	case *ATTRIBUTE_DEFINITION_DATE:
		stage.OnAfterATTRIBUTE_DEFINITION_DATEReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_DATE])
	
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_ENUMERATION])
	
	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_INTEGER])
	
	case *ATTRIBUTE_DEFINITION_REAL:
		stage.OnAfterATTRIBUTE_DEFINITION_REALReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_REAL])
	
	case *ATTRIBUTE_DEFINITION_STRING:
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_STRING])
	
	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_XHTML])
	
	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_BOOLEAN])
	
	case *ATTRIBUTE_VALUE_DATE:
		stage.OnAfterATTRIBUTE_VALUE_DATEReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_DATE])
	
	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_ENUMERATION])
	
	case *ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterATTRIBUTE_VALUE_INTEGERReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_INTEGER])
	
	case *ATTRIBUTE_VALUE_REAL:
		stage.OnAfterATTRIBUTE_VALUE_REALReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_REAL])
	
	case *ATTRIBUTE_VALUE_STRING:
		stage.OnAfterATTRIBUTE_VALUE_STRINGReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_STRING])
	
	case *ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterATTRIBUTE_VALUE_XHTMLReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_XHTML])
	
	case *A_ALTERNATIVE_ID:
		stage.OnAfterA_ALTERNATIVE_IDReadCallback = any(callback).(OnAfterReadInterface[A_ALTERNATIVE_ID])
	
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF])
	
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_DATE_REF])
	
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF])
	
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF])
	
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_REAL_REF])
	
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_STRING_REF])
	
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF])
	
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_BOOLEAN])
	
	case *A_ATTRIBUTE_VALUE_DATE:
		stage.OnAfterA_ATTRIBUTE_VALUE_DATEReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_DATE])
	
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_ENUMERATION])
	
	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_INTEGER])
	
	case *A_ATTRIBUTE_VALUE_REAL:
		stage.OnAfterA_ATTRIBUTE_VALUE_REALReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_REAL])
	
	case *A_ATTRIBUTE_VALUE_STRING:
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_STRING])
	
	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_XHTML])
	
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1ReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_XHTML_1])
	
	case *A_CHILDREN:
		stage.OnAfterA_CHILDRENReadCallback = any(callback).(OnAfterReadInterface[A_CHILDREN])
	
	case *A_CORE_CONTENT:
		stage.OnAfterA_CORE_CONTENTReadCallback = any(callback).(OnAfterReadInterface[A_CORE_CONTENT])
	
	case *A_DATATYPES:
		stage.OnAfterA_DATATYPESReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPES])
	
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF])
	
	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_DATE_REF])
	
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF])
	
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_INTEGER_REF])
	
	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_REAL_REF])
	
	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_STRING_REF])
	
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_XHTML_REF])
	
	case *A_EDITABLE_ATTS:
		stage.OnAfterA_EDITABLE_ATTSReadCallback = any(callback).(OnAfterReadInterface[A_EDITABLE_ATTS])
	
	case *A_ENUM_VALUE_REF:
		stage.OnAfterA_ENUM_VALUE_REFReadCallback = any(callback).(OnAfterReadInterface[A_ENUM_VALUE_REF])
	
	case *A_OBJECT:
		stage.OnAfterA_OBJECTReadCallback = any(callback).(OnAfterReadInterface[A_OBJECT])
	
	case *A_PROPERTIES:
		stage.OnAfterA_PROPERTIESReadCallback = any(callback).(OnAfterReadInterface[A_PROPERTIES])
	
	case *A_RELATION_GROUP_TYPE_REF:
		stage.OnAfterA_RELATION_GROUP_TYPE_REFReadCallback = any(callback).(OnAfterReadInterface[A_RELATION_GROUP_TYPE_REF])
	
	case *A_SOURCE_1:
		stage.OnAfterA_SOURCE_1ReadCallback = any(callback).(OnAfterReadInterface[A_SOURCE_1])
	
	case *A_SOURCE_SPECIFICATION_1:
		stage.OnAfterA_SOURCE_SPECIFICATION_1ReadCallback = any(callback).(OnAfterReadInterface[A_SOURCE_SPECIFICATION_1])
	
	case *A_SPECIFICATIONS:
		stage.OnAfterA_SPECIFICATIONSReadCallback = any(callback).(OnAfterReadInterface[A_SPECIFICATIONS])
	
	case *A_SPECIFICATION_TYPE_REF:
		stage.OnAfterA_SPECIFICATION_TYPE_REFReadCallback = any(callback).(OnAfterReadInterface[A_SPECIFICATION_TYPE_REF])
	
	case *A_SPECIFIED_VALUES:
		stage.OnAfterA_SPECIFIED_VALUESReadCallback = any(callback).(OnAfterReadInterface[A_SPECIFIED_VALUES])
	
	case *A_SPEC_ATTRIBUTES:
		stage.OnAfterA_SPEC_ATTRIBUTESReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_ATTRIBUTES])
	
	case *A_SPEC_OBJECTS:
		stage.OnAfterA_SPEC_OBJECTSReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_OBJECTS])
	
	case *A_SPEC_OBJECT_TYPE_REF:
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_OBJECT_TYPE_REF])
	
	case *A_SPEC_RELATIONS:
		stage.OnAfterA_SPEC_RELATIONSReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_RELATIONS])
	
	case *A_SPEC_RELATION_GROUPS:
		stage.OnAfterA_SPEC_RELATION_GROUPSReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_RELATION_GROUPS])
	
	case *A_SPEC_RELATION_REF:
		stage.OnAfterA_SPEC_RELATION_REFReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_RELATION_REF])
	
	case *A_SPEC_RELATION_TYPE_REF:
		stage.OnAfterA_SPEC_RELATION_TYPE_REFReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_RELATION_TYPE_REF])
	
	case *A_SPEC_TYPES:
		stage.OnAfterA_SPEC_TYPESReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_TYPES])
	
	case *A_THE_HEADER:
		stage.OnAfterA_THE_HEADERReadCallback = any(callback).(OnAfterReadInterface[A_THE_HEADER])
	
	case *A_TOOL_EXTENSIONS:
		stage.OnAfterA_TOOL_EXTENSIONSReadCallback = any(callback).(OnAfterReadInterface[A_TOOL_EXTENSIONS])
	
	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_BOOLEAN])
	
	case *DATATYPE_DEFINITION_DATE:
		stage.OnAfterDATATYPE_DEFINITION_DATEReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_DATE])
	
	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_ENUMERATION])
	
	case *DATATYPE_DEFINITION_INTEGER:
		stage.OnAfterDATATYPE_DEFINITION_INTEGERReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_INTEGER])
	
	case *DATATYPE_DEFINITION_REAL:
		stage.OnAfterDATATYPE_DEFINITION_REALReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_REAL])
	
	case *DATATYPE_DEFINITION_STRING:
		stage.OnAfterDATATYPE_DEFINITION_STRINGReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_STRING])
	
	case *DATATYPE_DEFINITION_XHTML:
		stage.OnAfterDATATYPE_DEFINITION_XHTMLReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_XHTML])
	
	case *EMBEDDED_VALUE:
		stage.OnAfterEMBEDDED_VALUEReadCallback = any(callback).(OnAfterReadInterface[EMBEDDED_VALUE])
	
	case *ENUM_VALUE:
		stage.OnAfterENUM_VALUEReadCallback = any(callback).(OnAfterReadInterface[ENUM_VALUE])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry])
	
	case *Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		stage.OnAfterMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryReadCallback = any(callback).(OnAfterReadInterface[Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry])
	
	case *Map_SPECIFICATION_Nodes_expandedEntry:
		stage.OnAfterMap_SPECIFICATION_Nodes_expandedEntryReadCallback = any(callback).(OnAfterReadInterface[Map_SPECIFICATION_Nodes_expandedEntry])
	
	case *Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryReadCallback = any(callback).(OnAfterReadInterface[Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry])
	
	case *Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_showIdentifierEntryReadCallback = any(callback).(OnAfterReadInterface[Map_SPEC_OBJECT_TYPE_showIdentifierEntry])
	
	case *Map_SPEC_OBJECT_TYPE_showNameEntry:
		stage.OnAfterMap_SPEC_OBJECT_TYPE_showNameEntryReadCallback = any(callback).(OnAfterReadInterface[Map_SPEC_OBJECT_TYPE_showNameEntry])
	
	case *RELATION_GROUP:
		stage.OnAfterRELATION_GROUPReadCallback = any(callback).(OnAfterReadInterface[RELATION_GROUP])
	
	case *RELATION_GROUP_TYPE:
		stage.OnAfterRELATION_GROUP_TYPEReadCallback = any(callback).(OnAfterReadInterface[RELATION_GROUP_TYPE])
	
	case *REQ_IF:
		stage.OnAfterREQ_IFReadCallback = any(callback).(OnAfterReadInterface[REQ_IF])
	
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTReadCallback = any(callback).(OnAfterReadInterface[REQ_IF_CONTENT])
	
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERReadCallback = any(callback).(OnAfterReadInterface[REQ_IF_HEADER])
	
	case *REQ_IF_TOOL_EXTENSION:
		stage.OnAfterREQ_IF_TOOL_EXTENSIONReadCallback = any(callback).(OnAfterReadInterface[REQ_IF_TOOL_EXTENSION])
	
	case *RenderingConfiguration:
		stage.OnAfterRenderingConfigurationReadCallback = any(callback).(OnAfterReadInterface[RenderingConfiguration])
	
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONReadCallback = any(callback).(OnAfterReadInterface[SPECIFICATION])
	
	case *SPECIFICATION_TYPE:
		stage.OnAfterSPECIFICATION_TYPEReadCallback = any(callback).(OnAfterReadInterface[SPECIFICATION_TYPE])
	
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYReadCallback = any(callback).(OnAfterReadInterface[SPEC_HIERARCHY])
	
	case *SPEC_OBJECT:
		stage.OnAfterSPEC_OBJECTReadCallback = any(callback).(OnAfterReadInterface[SPEC_OBJECT])
	
	case *SPEC_OBJECT_TYPE:
		stage.OnAfterSPEC_OBJECT_TYPEReadCallback = any(callback).(OnAfterReadInterface[SPEC_OBJECT_TYPE])
	
	case *SPEC_RELATION:
		stage.OnAfterSPEC_RELATIONReadCallback = any(callback).(OnAfterReadInterface[SPEC_RELATION])
	
	case *SPEC_RELATION_TYPE:
		stage.OnAfterSPEC_RELATION_TYPEReadCallback = any(callback).(OnAfterReadInterface[SPEC_RELATION_TYPE])
	
	case *StaticWebSite:
		stage.OnAfterStaticWebSiteReadCallback = any(callback).(OnAfterReadInterface[StaticWebSite])
	
	case *StaticWebSiteChapter:
		stage.OnAfterStaticWebSiteChapterReadCallback = any(callback).(OnAfterReadInterface[StaticWebSiteChapter])
	
	case *StaticWebSiteGeneratedImage:
		stage.OnAfterStaticWebSiteGeneratedImageReadCallback = any(callback).(OnAfterReadInterface[StaticWebSiteGeneratedImage])
	
	case *StaticWebSiteImage:
		stage.OnAfterStaticWebSiteImageReadCallback = any(callback).(OnAfterReadInterface[StaticWebSiteImage])
	
	case *StaticWebSiteParagraph:
		stage.OnAfterStaticWebSiteParagraphReadCallback = any(callback).(OnAfterReadInterface[StaticWebSiteParagraph])
	
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTReadCallback = any(callback).(OnAfterReadInterface[XHTML_CONTENT])
	
	}
}
