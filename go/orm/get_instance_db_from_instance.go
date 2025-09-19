// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongreqif/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.ALTERNATIVE_ID:
		alternative_idInstance := any(concreteInstance).(*models.ALTERNATIVE_ID)
		ret2 := backRepo.BackRepoALTERNATIVE_ID.GetALTERNATIVE_IDDBFromALTERNATIVE_IDPtr(alternative_idInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		attribute_definition_booleanInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_BOOLEAN)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.GetATTRIBUTE_DEFINITION_BOOLEANDBFromATTRIBUTE_DEFINITION_BOOLEANPtr(attribute_definition_booleanInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		attribute_definition_dateInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_DATE)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.GetATTRIBUTE_DEFINITION_DATEDBFromATTRIBUTE_DEFINITION_DATEPtr(attribute_definition_dateInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		attribute_definition_enumerationInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_ENUMERATION)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.GetATTRIBUTE_DEFINITION_ENUMERATIONDBFromATTRIBUTE_DEFINITION_ENUMERATIONPtr(attribute_definition_enumerationInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		attribute_definition_integerInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_INTEGER)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.GetATTRIBUTE_DEFINITION_INTEGERDBFromATTRIBUTE_DEFINITION_INTEGERPtr(attribute_definition_integerInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		attribute_definition_realInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_REAL)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.GetATTRIBUTE_DEFINITION_REALDBFromATTRIBUTE_DEFINITION_REALPtr(attribute_definition_realInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		attribute_definition_stringInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_STRING)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.GetATTRIBUTE_DEFINITION_STRINGDBFromATTRIBUTE_DEFINITION_STRINGPtr(attribute_definition_stringInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		attribute_definition_xhtmlInstance := any(concreteInstance).(*models.ATTRIBUTE_DEFINITION_XHTML)
		ret2 := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.GetATTRIBUTE_DEFINITION_XHTMLDBFromATTRIBUTE_DEFINITION_XHTMLPtr(attribute_definition_xhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		attribute_value_booleanInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_BOOLEAN)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.GetATTRIBUTE_VALUE_BOOLEANDBFromATTRIBUTE_VALUE_BOOLEANPtr(attribute_value_booleanInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_DATE:
		attribute_value_dateInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_DATE)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_DATE.GetATTRIBUTE_VALUE_DATEDBFromATTRIBUTE_VALUE_DATEPtr(attribute_value_dateInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		attribute_value_enumerationInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_ENUMERATION)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.GetATTRIBUTE_VALUE_ENUMERATIONDBFromATTRIBUTE_VALUE_ENUMERATIONPtr(attribute_value_enumerationInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		attribute_value_integerInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_INTEGER)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.GetATTRIBUTE_VALUE_INTEGERDBFromATTRIBUTE_VALUE_INTEGERPtr(attribute_value_integerInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_REAL:
		attribute_value_realInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_REAL)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_REAL.GetATTRIBUTE_VALUE_REALDBFromATTRIBUTE_VALUE_REALPtr(attribute_value_realInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_STRING:
		attribute_value_stringInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_STRING)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_STRING.GetATTRIBUTE_VALUE_STRINGDBFromATTRIBUTE_VALUE_STRINGPtr(attribute_value_stringInstance)
		ret = any(ret2).(*T2)
	case *models.ATTRIBUTE_VALUE_XHTML:
		attribute_value_xhtmlInstance := any(concreteInstance).(*models.ATTRIBUTE_VALUE_XHTML)
		ret2 := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.GetATTRIBUTE_VALUE_XHTMLDBFromATTRIBUTE_VALUE_XHTMLPtr(attribute_value_xhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.A_ALTERNATIVE_ID:
		a_alternative_idInstance := any(concreteInstance).(*models.A_ALTERNATIVE_ID)
		ret2 := backRepo.BackRepoA_ALTERNATIVE_ID.GetA_ALTERNATIVE_IDDBFromA_ALTERNATIVE_IDPtr(a_alternative_idInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		a_attribute_definition_boolean_refInstance := any(concreteInstance).(*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_BOOLEAN_REF.GetA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDBFromA_ATTRIBUTE_DEFINITION_BOOLEAN_REFPtr(a_attribute_definition_boolean_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		a_attribute_definition_date_refInstance := any(concreteInstance).(*models.A_ATTRIBUTE_DEFINITION_DATE_REF)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_DATE_REF.GetA_ATTRIBUTE_DEFINITION_DATE_REFDBFromA_ATTRIBUTE_DEFINITION_DATE_REFPtr(a_attribute_definition_date_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		a_attribute_definition_enumeration_refInstance := any(concreteInstance).(*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_ENUMERATION_REF.GetA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDBFromA_ATTRIBUTE_DEFINITION_ENUMERATION_REFPtr(a_attribute_definition_enumeration_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		a_attribute_definition_integer_refInstance := any(concreteInstance).(*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.GetA_ATTRIBUTE_DEFINITION_INTEGER_REFDBFromA_ATTRIBUTE_DEFINITION_INTEGER_REFPtr(a_attribute_definition_integer_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		a_attribute_definition_real_refInstance := any(concreteInstance).(*models.A_ATTRIBUTE_DEFINITION_REAL_REF)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_REAL_REF.GetA_ATTRIBUTE_DEFINITION_REAL_REFDBFromA_ATTRIBUTE_DEFINITION_REAL_REFPtr(a_attribute_definition_real_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		a_attribute_definition_string_refInstance := any(concreteInstance).(*models.A_ATTRIBUTE_DEFINITION_STRING_REF)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.GetA_ATTRIBUTE_DEFINITION_STRING_REFDBFromA_ATTRIBUTE_DEFINITION_STRING_REFPtr(a_attribute_definition_string_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		a_attribute_definition_xhtml_refInstance := any(concreteInstance).(*models.A_ATTRIBUTE_DEFINITION_XHTML_REF)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_XHTML_REF.GetA_ATTRIBUTE_DEFINITION_XHTML_REFDBFromA_ATTRIBUTE_DEFINITION_XHTML_REFPtr(a_attribute_definition_xhtml_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		a_attribute_value_booleanInstance := any(concreteInstance).(*models.A_ATTRIBUTE_VALUE_BOOLEAN)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.GetA_ATTRIBUTE_VALUE_BOOLEANDBFromA_ATTRIBUTE_VALUE_BOOLEANPtr(a_attribute_value_booleanInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_VALUE_DATE:
		a_attribute_value_dateInstance := any(concreteInstance).(*models.A_ATTRIBUTE_VALUE_DATE)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.GetA_ATTRIBUTE_VALUE_DATEDBFromA_ATTRIBUTE_VALUE_DATEPtr(a_attribute_value_dateInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		a_attribute_value_enumerationInstance := any(concreteInstance).(*models.A_ATTRIBUTE_VALUE_ENUMERATION)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.GetA_ATTRIBUTE_VALUE_ENUMERATIONDBFromA_ATTRIBUTE_VALUE_ENUMERATIONPtr(a_attribute_value_enumerationInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		a_attribute_value_integerInstance := any(concreteInstance).(*models.A_ATTRIBUTE_VALUE_INTEGER)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.GetA_ATTRIBUTE_VALUE_INTEGERDBFromA_ATTRIBUTE_VALUE_INTEGERPtr(a_attribute_value_integerInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_VALUE_REAL:
		a_attribute_value_realInstance := any(concreteInstance).(*models.A_ATTRIBUTE_VALUE_REAL)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_VALUE_REAL.GetA_ATTRIBUTE_VALUE_REALDBFromA_ATTRIBUTE_VALUE_REALPtr(a_attribute_value_realInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_VALUE_STRING:
		a_attribute_value_stringInstance := any(concreteInstance).(*models.A_ATTRIBUTE_VALUE_STRING)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.GetA_ATTRIBUTE_VALUE_STRINGDBFromA_ATTRIBUTE_VALUE_STRINGPtr(a_attribute_value_stringInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_VALUE_XHTML:
		a_attribute_value_xhtmlInstance := any(concreteInstance).(*models.A_ATTRIBUTE_VALUE_XHTML)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.GetA_ATTRIBUTE_VALUE_XHTMLDBFromA_ATTRIBUTE_VALUE_XHTMLPtr(a_attribute_value_xhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		a_attribute_value_xhtml_1Instance := any(concreteInstance).(*models.A_ATTRIBUTE_VALUE_XHTML_1)
		ret2 := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML_1.GetA_ATTRIBUTE_VALUE_XHTML_1DBFromA_ATTRIBUTE_VALUE_XHTML_1Ptr(a_attribute_value_xhtml_1Instance)
		ret = any(ret2).(*T2)
	case *models.A_CHILDREN:
		a_childrenInstance := any(concreteInstance).(*models.A_CHILDREN)
		ret2 := backRepo.BackRepoA_CHILDREN.GetA_CHILDRENDBFromA_CHILDRENPtr(a_childrenInstance)
		ret = any(ret2).(*T2)
	case *models.A_CORE_CONTENT:
		a_core_contentInstance := any(concreteInstance).(*models.A_CORE_CONTENT)
		ret2 := backRepo.BackRepoA_CORE_CONTENT.GetA_CORE_CONTENTDBFromA_CORE_CONTENTPtr(a_core_contentInstance)
		ret = any(ret2).(*T2)
	case *models.A_DATATYPES:
		a_datatypesInstance := any(concreteInstance).(*models.A_DATATYPES)
		ret2 := backRepo.BackRepoA_DATATYPES.GetA_DATATYPESDBFromA_DATATYPESPtr(a_datatypesInstance)
		ret = any(ret2).(*T2)
	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		a_datatype_definition_boolean_refInstance := any(concreteInstance).(*models.A_DATATYPE_DEFINITION_BOOLEAN_REF)
		ret2 := backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.GetA_DATATYPE_DEFINITION_BOOLEAN_REFDBFromA_DATATYPE_DEFINITION_BOOLEAN_REFPtr(a_datatype_definition_boolean_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		a_datatype_definition_date_refInstance := any(concreteInstance).(*models.A_DATATYPE_DEFINITION_DATE_REF)
		ret2 := backRepo.BackRepoA_DATATYPE_DEFINITION_DATE_REF.GetA_DATATYPE_DEFINITION_DATE_REFDBFromA_DATATYPE_DEFINITION_DATE_REFPtr(a_datatype_definition_date_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		a_datatype_definition_enumeration_refInstance := any(concreteInstance).(*models.A_DATATYPE_DEFINITION_ENUMERATION_REF)
		ret2 := backRepo.BackRepoA_DATATYPE_DEFINITION_ENUMERATION_REF.GetA_DATATYPE_DEFINITION_ENUMERATION_REFDBFromA_DATATYPE_DEFINITION_ENUMERATION_REFPtr(a_datatype_definition_enumeration_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		a_datatype_definition_integer_refInstance := any(concreteInstance).(*models.A_DATATYPE_DEFINITION_INTEGER_REF)
		ret2 := backRepo.BackRepoA_DATATYPE_DEFINITION_INTEGER_REF.GetA_DATATYPE_DEFINITION_INTEGER_REFDBFromA_DATATYPE_DEFINITION_INTEGER_REFPtr(a_datatype_definition_integer_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		a_datatype_definition_real_refInstance := any(concreteInstance).(*models.A_DATATYPE_DEFINITION_REAL_REF)
		ret2 := backRepo.BackRepoA_DATATYPE_DEFINITION_REAL_REF.GetA_DATATYPE_DEFINITION_REAL_REFDBFromA_DATATYPE_DEFINITION_REAL_REFPtr(a_datatype_definition_real_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		a_datatype_definition_string_refInstance := any(concreteInstance).(*models.A_DATATYPE_DEFINITION_STRING_REF)
		ret2 := backRepo.BackRepoA_DATATYPE_DEFINITION_STRING_REF.GetA_DATATYPE_DEFINITION_STRING_REFDBFromA_DATATYPE_DEFINITION_STRING_REFPtr(a_datatype_definition_string_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		a_datatype_definition_xhtml_refInstance := any(concreteInstance).(*models.A_DATATYPE_DEFINITION_XHTML_REF)
		ret2 := backRepo.BackRepoA_DATATYPE_DEFINITION_XHTML_REF.GetA_DATATYPE_DEFINITION_XHTML_REFDBFromA_DATATYPE_DEFINITION_XHTML_REFPtr(a_datatype_definition_xhtml_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_EDITABLE_ATTS:
		a_editable_attsInstance := any(concreteInstance).(*models.A_EDITABLE_ATTS)
		ret2 := backRepo.BackRepoA_EDITABLE_ATTS.GetA_EDITABLE_ATTSDBFromA_EDITABLE_ATTSPtr(a_editable_attsInstance)
		ret = any(ret2).(*T2)
	case *models.A_ENUM_VALUE_REF:
		a_enum_value_refInstance := any(concreteInstance).(*models.A_ENUM_VALUE_REF)
		ret2 := backRepo.BackRepoA_ENUM_VALUE_REF.GetA_ENUM_VALUE_REFDBFromA_ENUM_VALUE_REFPtr(a_enum_value_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_OBJECT:
		a_objectInstance := any(concreteInstance).(*models.A_OBJECT)
		ret2 := backRepo.BackRepoA_OBJECT.GetA_OBJECTDBFromA_OBJECTPtr(a_objectInstance)
		ret = any(ret2).(*T2)
	case *models.A_PROPERTIES:
		a_propertiesInstance := any(concreteInstance).(*models.A_PROPERTIES)
		ret2 := backRepo.BackRepoA_PROPERTIES.GetA_PROPERTIESDBFromA_PROPERTIESPtr(a_propertiesInstance)
		ret = any(ret2).(*T2)
	case *models.A_RELATION_GROUP_TYPE_REF:
		a_relation_group_type_refInstance := any(concreteInstance).(*models.A_RELATION_GROUP_TYPE_REF)
		ret2 := backRepo.BackRepoA_RELATION_GROUP_TYPE_REF.GetA_RELATION_GROUP_TYPE_REFDBFromA_RELATION_GROUP_TYPE_REFPtr(a_relation_group_type_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_SOURCE_1:
		a_source_1Instance := any(concreteInstance).(*models.A_SOURCE_1)
		ret2 := backRepo.BackRepoA_SOURCE_1.GetA_SOURCE_1DBFromA_SOURCE_1Ptr(a_source_1Instance)
		ret = any(ret2).(*T2)
	case *models.A_SOURCE_SPECIFICATION_1:
		a_source_specification_1Instance := any(concreteInstance).(*models.A_SOURCE_SPECIFICATION_1)
		ret2 := backRepo.BackRepoA_SOURCE_SPECIFICATION_1.GetA_SOURCE_SPECIFICATION_1DBFromA_SOURCE_SPECIFICATION_1Ptr(a_source_specification_1Instance)
		ret = any(ret2).(*T2)
	case *models.A_SPECIFICATIONS:
		a_specificationsInstance := any(concreteInstance).(*models.A_SPECIFICATIONS)
		ret2 := backRepo.BackRepoA_SPECIFICATIONS.GetA_SPECIFICATIONSDBFromA_SPECIFICATIONSPtr(a_specificationsInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPECIFICATION_TYPE_REF:
		a_specification_type_refInstance := any(concreteInstance).(*models.A_SPECIFICATION_TYPE_REF)
		ret2 := backRepo.BackRepoA_SPECIFICATION_TYPE_REF.GetA_SPECIFICATION_TYPE_REFDBFromA_SPECIFICATION_TYPE_REFPtr(a_specification_type_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPECIFIED_VALUES:
		a_specified_valuesInstance := any(concreteInstance).(*models.A_SPECIFIED_VALUES)
		ret2 := backRepo.BackRepoA_SPECIFIED_VALUES.GetA_SPECIFIED_VALUESDBFromA_SPECIFIED_VALUESPtr(a_specified_valuesInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPEC_ATTRIBUTES:
		a_spec_attributesInstance := any(concreteInstance).(*models.A_SPEC_ATTRIBUTES)
		ret2 := backRepo.BackRepoA_SPEC_ATTRIBUTES.GetA_SPEC_ATTRIBUTESDBFromA_SPEC_ATTRIBUTESPtr(a_spec_attributesInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPEC_OBJECTS:
		a_spec_objectsInstance := any(concreteInstance).(*models.A_SPEC_OBJECTS)
		ret2 := backRepo.BackRepoA_SPEC_OBJECTS.GetA_SPEC_OBJECTSDBFromA_SPEC_OBJECTSPtr(a_spec_objectsInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPEC_OBJECT_TYPE_REF:
		a_spec_object_type_refInstance := any(concreteInstance).(*models.A_SPEC_OBJECT_TYPE_REF)
		ret2 := backRepo.BackRepoA_SPEC_OBJECT_TYPE_REF.GetA_SPEC_OBJECT_TYPE_REFDBFromA_SPEC_OBJECT_TYPE_REFPtr(a_spec_object_type_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPEC_RELATIONS:
		a_spec_relationsInstance := any(concreteInstance).(*models.A_SPEC_RELATIONS)
		ret2 := backRepo.BackRepoA_SPEC_RELATIONS.GetA_SPEC_RELATIONSDBFromA_SPEC_RELATIONSPtr(a_spec_relationsInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPEC_RELATION_GROUPS:
		a_spec_relation_groupsInstance := any(concreteInstance).(*models.A_SPEC_RELATION_GROUPS)
		ret2 := backRepo.BackRepoA_SPEC_RELATION_GROUPS.GetA_SPEC_RELATION_GROUPSDBFromA_SPEC_RELATION_GROUPSPtr(a_spec_relation_groupsInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPEC_RELATION_REF:
		a_spec_relation_refInstance := any(concreteInstance).(*models.A_SPEC_RELATION_REF)
		ret2 := backRepo.BackRepoA_SPEC_RELATION_REF.GetA_SPEC_RELATION_REFDBFromA_SPEC_RELATION_REFPtr(a_spec_relation_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPEC_RELATION_TYPE_REF:
		a_spec_relation_type_refInstance := any(concreteInstance).(*models.A_SPEC_RELATION_TYPE_REF)
		ret2 := backRepo.BackRepoA_SPEC_RELATION_TYPE_REF.GetA_SPEC_RELATION_TYPE_REFDBFromA_SPEC_RELATION_TYPE_REFPtr(a_spec_relation_type_refInstance)
		ret = any(ret2).(*T2)
	case *models.A_SPEC_TYPES:
		a_spec_typesInstance := any(concreteInstance).(*models.A_SPEC_TYPES)
		ret2 := backRepo.BackRepoA_SPEC_TYPES.GetA_SPEC_TYPESDBFromA_SPEC_TYPESPtr(a_spec_typesInstance)
		ret = any(ret2).(*T2)
	case *models.A_THE_HEADER:
		a_the_headerInstance := any(concreteInstance).(*models.A_THE_HEADER)
		ret2 := backRepo.BackRepoA_THE_HEADER.GetA_THE_HEADERDBFromA_THE_HEADERPtr(a_the_headerInstance)
		ret = any(ret2).(*T2)
	case *models.A_TOOL_EXTENSIONS:
		a_tool_extensionsInstance := any(concreteInstance).(*models.A_TOOL_EXTENSIONS)
		ret2 := backRepo.BackRepoA_TOOL_EXTENSIONS.GetA_TOOL_EXTENSIONSDBFromA_TOOL_EXTENSIONSPtr(a_tool_extensionsInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		datatype_definition_booleanInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_BOOLEAN)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.GetDATATYPE_DEFINITION_BOOLEANDBFromDATATYPE_DEFINITION_BOOLEANPtr(datatype_definition_booleanInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_DATE:
		datatype_definition_dateInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_DATE)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_DATE.GetDATATYPE_DEFINITION_DATEDBFromDATATYPE_DEFINITION_DATEPtr(datatype_definition_dateInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		datatype_definition_enumerationInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_ENUMERATION)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.GetDATATYPE_DEFINITION_ENUMERATIONDBFromDATATYPE_DEFINITION_ENUMERATIONPtr(datatype_definition_enumerationInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_INTEGER:
		datatype_definition_integerInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_INTEGER)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.GetDATATYPE_DEFINITION_INTEGERDBFromDATATYPE_DEFINITION_INTEGERPtr(datatype_definition_integerInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_REAL:
		datatype_definition_realInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_REAL)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_REAL.GetDATATYPE_DEFINITION_REALDBFromDATATYPE_DEFINITION_REALPtr(datatype_definition_realInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_STRING:
		datatype_definition_stringInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_STRING)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_STRING.GetDATATYPE_DEFINITION_STRINGDBFromDATATYPE_DEFINITION_STRINGPtr(datatype_definition_stringInstance)
		ret = any(ret2).(*T2)
	case *models.DATATYPE_DEFINITION_XHTML:
		datatype_definition_xhtmlInstance := any(concreteInstance).(*models.DATATYPE_DEFINITION_XHTML)
		ret2 := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.GetDATATYPE_DEFINITION_XHTMLDBFromDATATYPE_DEFINITION_XHTMLPtr(datatype_definition_xhtmlInstance)
		ret = any(ret2).(*T2)
	case *models.EMBEDDED_VALUE:
		embedded_valueInstance := any(concreteInstance).(*models.EMBEDDED_VALUE)
		ret2 := backRepo.BackRepoEMBEDDED_VALUE.GetEMBEDDED_VALUEDBFromEMBEDDED_VALUEPtr(embedded_valueInstance)
		ret = any(ret2).(*T2)
	case *models.ENUM_VALUE:
		enum_valueInstance := any(concreteInstance).(*models.ENUM_VALUE)
		ret2 := backRepo.BackRepoENUM_VALUE.GetENUM_VALUEDBFromENUM_VALUEPtr(enum_valueInstance)
		ret = any(ret2).(*T2)
	case *models.JpgImage:
		jpgimageInstance := any(concreteInstance).(*models.JpgImage)
		ret2 := backRepo.BackRepoJpgImage.GetJpgImageDBFromJpgImagePtr(jpgimageInstance)
		ret = any(ret2).(*T2)
	case *models.Kill:
		killInstance := any(concreteInstance).(*models.Kill)
		ret2 := backRepo.BackRepoKill.GetKillDBFromKillPtr(killInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		map_attribute_definition_boolean_showinsubjectentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry.GetMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDBFromMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryPtr(map_attribute_definition_boolean_showinsubjectentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		map_attribute_definition_boolean_showintableentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry.GetMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDBFromMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryPtr(map_attribute_definition_boolean_showintableentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		map_attribute_definition_boolean_showintitleentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry.GetMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDBFromMap_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryPtr(map_attribute_definition_boolean_showintitleentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		map_attribute_definition_date_showinsubjectentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry.GetMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDBFromMap_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryPtr(map_attribute_definition_date_showinsubjectentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		map_attribute_definition_date_showintableentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry.GetMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDBFromMap_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryPtr(map_attribute_definition_date_showintableentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		map_attribute_definition_date_showintitleentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry.GetMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDBFromMap_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryPtr(map_attribute_definition_date_showintitleentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		map_attribute_definition_enumeration_showinsubjectentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDBFromMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPtr(map_attribute_definition_enumeration_showinsubjectentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		map_attribute_definition_enumeration_showintableentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry.GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDBFromMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryPtr(map_attribute_definition_enumeration_showintableentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		map_attribute_definition_enumeration_showintitleentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDBFromMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPtr(map_attribute_definition_enumeration_showintitleentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		map_attribute_definition_integer_showinsubjectentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry.GetMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDBFromMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryPtr(map_attribute_definition_integer_showinsubjectentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		map_attribute_definition_integer_showintableentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry.GetMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDBFromMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryPtr(map_attribute_definition_integer_showintableentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		map_attribute_definition_integer_showintitleentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry.GetMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDBFromMap_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryPtr(map_attribute_definition_integer_showintitleentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		map_attribute_definition_real_showinsubjectentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry.GetMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDBFromMap_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryPtr(map_attribute_definition_real_showinsubjectentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		map_attribute_definition_real_showintableentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry.GetMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDBFromMap_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryPtr(map_attribute_definition_real_showintableentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		map_attribute_definition_real_showintitleentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry.GetMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDBFromMap_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryPtr(map_attribute_definition_real_showintitleentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		map_attribute_definition_string_showinsubjectentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry.GetMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDBFromMap_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryPtr(map_attribute_definition_string_showinsubjectentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		map_attribute_definition_string_showintableentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry.GetMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDBFromMap_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryPtr(map_attribute_definition_string_showintableentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		map_attribute_definition_string_showintitleentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry.GetMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDBFromMap_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryPtr(map_attribute_definition_string_showintitleentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		map_attribute_definition_xhtml_showinsubjectentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry.GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDBFromMap_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryPtr(map_attribute_definition_xhtml_showinsubjectentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		map_attribute_definition_xhtml_showintableentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDBFromMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPtr(map_attribute_definition_xhtml_showintableentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		map_attribute_definition_xhtml_showintitleentryInstance := any(concreteInstance).(*models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry)
		ret2 := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDBFromMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPtr(map_attribute_definition_xhtml_showintitleentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_SPECIFICATION_Nodes_expandedEntry:
		map_specification_nodes_expandedentryInstance := any(concreteInstance).(*models.Map_SPECIFICATION_Nodes_expandedEntry)
		ret2 := backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.GetMap_SPECIFICATION_Nodes_expandedEntryDBFromMap_SPECIFICATION_Nodes_expandedEntryPtr(map_specification_nodes_expandedentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		map_spec_object_type_isnodeexpandedentryInstance := any(concreteInstance).(*models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry)
		ret2 := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.GetMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryDBFromMap_SPEC_OBJECT_TYPE_isNodeExpandedEntryPtr(map_spec_object_type_isnodeexpandedentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		map_spec_object_type_showidentifierentryInstance := any(concreteInstance).(*models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry)
		ret2 := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.GetMap_SPEC_OBJECT_TYPE_showIdentifierEntryDBFromMap_SPEC_OBJECT_TYPE_showIdentifierEntryPtr(map_spec_object_type_showidentifierentryInstance)
		ret = any(ret2).(*T2)
	case *models.Map_SPEC_OBJECT_TYPE_showNameEntry:
		map_spec_object_type_shownameentryInstance := any(concreteInstance).(*models.Map_SPEC_OBJECT_TYPE_showNameEntry)
		ret2 := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.GetMap_SPEC_OBJECT_TYPE_showNameEntryDBFromMap_SPEC_OBJECT_TYPE_showNameEntryPtr(map_spec_object_type_shownameentryInstance)
		ret = any(ret2).(*T2)
	case *models.PngImage:
		pngimageInstance := any(concreteInstance).(*models.PngImage)
		ret2 := backRepo.BackRepoPngImage.GetPngImageDBFromPngImagePtr(pngimageInstance)
		ret = any(ret2).(*T2)
	case *models.RELATION_GROUP:
		relation_groupInstance := any(concreteInstance).(*models.RELATION_GROUP)
		ret2 := backRepo.BackRepoRELATION_GROUP.GetRELATION_GROUPDBFromRELATION_GROUPPtr(relation_groupInstance)
		ret = any(ret2).(*T2)
	case *models.RELATION_GROUP_TYPE:
		relation_group_typeInstance := any(concreteInstance).(*models.RELATION_GROUP_TYPE)
		ret2 := backRepo.BackRepoRELATION_GROUP_TYPE.GetRELATION_GROUP_TYPEDBFromRELATION_GROUP_TYPEPtr(relation_group_typeInstance)
		ret = any(ret2).(*T2)
	case *models.REQ_IF:
		req_ifInstance := any(concreteInstance).(*models.REQ_IF)
		ret2 := backRepo.BackRepoREQ_IF.GetREQ_IFDBFromREQ_IFPtr(req_ifInstance)
		ret = any(ret2).(*T2)
	case *models.REQ_IF_CONTENT:
		req_if_contentInstance := any(concreteInstance).(*models.REQ_IF_CONTENT)
		ret2 := backRepo.BackRepoREQ_IF_CONTENT.GetREQ_IF_CONTENTDBFromREQ_IF_CONTENTPtr(req_if_contentInstance)
		ret = any(ret2).(*T2)
	case *models.REQ_IF_HEADER:
		req_if_headerInstance := any(concreteInstance).(*models.REQ_IF_HEADER)
		ret2 := backRepo.BackRepoREQ_IF_HEADER.GetREQ_IF_HEADERDBFromREQ_IF_HEADERPtr(req_if_headerInstance)
		ret = any(ret2).(*T2)
	case *models.REQ_IF_TOOL_EXTENSION:
		req_if_tool_extensionInstance := any(concreteInstance).(*models.REQ_IF_TOOL_EXTENSION)
		ret2 := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.GetREQ_IF_TOOL_EXTENSIONDBFromREQ_IF_TOOL_EXTENSIONPtr(req_if_tool_extensionInstance)
		ret = any(ret2).(*T2)
	case *models.RenderingConfiguration:
		renderingconfigurationInstance := any(concreteInstance).(*models.RenderingConfiguration)
		ret2 := backRepo.BackRepoRenderingConfiguration.GetRenderingConfigurationDBFromRenderingConfigurationPtr(renderingconfigurationInstance)
		ret = any(ret2).(*T2)
	case *models.SPECIFICATION:
		specificationInstance := any(concreteInstance).(*models.SPECIFICATION)
		ret2 := backRepo.BackRepoSPECIFICATION.GetSPECIFICATIONDBFromSPECIFICATIONPtr(specificationInstance)
		ret = any(ret2).(*T2)
	case *models.SPECIFICATION_TYPE:
		specification_typeInstance := any(concreteInstance).(*models.SPECIFICATION_TYPE)
		ret2 := backRepo.BackRepoSPECIFICATION_TYPE.GetSPECIFICATION_TYPEDBFromSPECIFICATION_TYPEPtr(specification_typeInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_HIERARCHY:
		spec_hierarchyInstance := any(concreteInstance).(*models.SPEC_HIERARCHY)
		ret2 := backRepo.BackRepoSPEC_HIERARCHY.GetSPEC_HIERARCHYDBFromSPEC_HIERARCHYPtr(spec_hierarchyInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_OBJECT:
		spec_objectInstance := any(concreteInstance).(*models.SPEC_OBJECT)
		ret2 := backRepo.BackRepoSPEC_OBJECT.GetSPEC_OBJECTDBFromSPEC_OBJECTPtr(spec_objectInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_OBJECT_TYPE:
		spec_object_typeInstance := any(concreteInstance).(*models.SPEC_OBJECT_TYPE)
		ret2 := backRepo.BackRepoSPEC_OBJECT_TYPE.GetSPEC_OBJECT_TYPEDBFromSPEC_OBJECT_TYPEPtr(spec_object_typeInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_RELATION:
		spec_relationInstance := any(concreteInstance).(*models.SPEC_RELATION)
		ret2 := backRepo.BackRepoSPEC_RELATION.GetSPEC_RELATIONDBFromSPEC_RELATIONPtr(spec_relationInstance)
		ret = any(ret2).(*T2)
	case *models.SPEC_RELATION_TYPE:
		spec_relation_typeInstance := any(concreteInstance).(*models.SPEC_RELATION_TYPE)
		ret2 := backRepo.BackRepoSPEC_RELATION_TYPE.GetSPEC_RELATION_TYPEDBFromSPEC_RELATION_TYPEPtr(spec_relation_typeInstance)
		ret = any(ret2).(*T2)
	case *models.StaticWebSite:
		staticwebsiteInstance := any(concreteInstance).(*models.StaticWebSite)
		ret2 := backRepo.BackRepoStaticWebSite.GetStaticWebSiteDBFromStaticWebSitePtr(staticwebsiteInstance)
		ret = any(ret2).(*T2)
	case *models.StaticWebSiteChapter:
		staticwebsitechapterInstance := any(concreteInstance).(*models.StaticWebSiteChapter)
		ret2 := backRepo.BackRepoStaticWebSiteChapter.GetStaticWebSiteChapterDBFromStaticWebSiteChapterPtr(staticwebsitechapterInstance)
		ret = any(ret2).(*T2)
	case *models.StaticWebSiteGeneratedImage:
		staticwebsitegeneratedimageInstance := any(concreteInstance).(*models.StaticWebSiteGeneratedImage)
		ret2 := backRepo.BackRepoStaticWebSiteGeneratedImage.GetStaticWebSiteGeneratedImageDBFromStaticWebSiteGeneratedImagePtr(staticwebsitegeneratedimageInstance)
		ret = any(ret2).(*T2)
	case *models.StaticWebSiteImage:
		staticwebsiteimageInstance := any(concreteInstance).(*models.StaticWebSiteImage)
		ret2 := backRepo.BackRepoStaticWebSiteImage.GetStaticWebSiteImageDBFromStaticWebSiteImagePtr(staticwebsiteimageInstance)
		ret = any(ret2).(*T2)
	case *models.StaticWebSiteParagraph:
		staticwebsiteparagraphInstance := any(concreteInstance).(*models.StaticWebSiteParagraph)
		ret2 := backRepo.BackRepoStaticWebSiteParagraph.GetStaticWebSiteParagraphDBFromStaticWebSiteParagraphPtr(staticwebsiteparagraphInstance)
		ret = any(ret2).(*T2)
	case *models.SvgImage:
		svgimageInstance := any(concreteInstance).(*models.SvgImage)
		ret2 := backRepo.BackRepoSvgImage.GetSvgImageDBFromSvgImagePtr(svgimageInstance)
		ret = any(ret2).(*T2)
	case *models.XHTML_CONTENT:
		xhtml_contentInstance := any(concreteInstance).(*models.XHTML_CONTENT)
		ret2 := backRepo.BackRepoXHTML_CONTENT.GetXHTML_CONTENTDBFromXHTML_CONTENTPtr(xhtml_contentInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.ALTERNATIVE_ID:
		tmp := GetInstanceDBFromInstance[models.ALTERNATIVE_ID, ALTERNATIVE_IDDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_BOOLEAN, ATTRIBUTE_DEFINITION_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_DATE, ATTRIBUTE_DEFINITION_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_ENUMERATION, ATTRIBUTE_DEFINITION_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_INTEGER, ATTRIBUTE_DEFINITION_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_REAL, ATTRIBUTE_DEFINITION_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_STRING, ATTRIBUTE_DEFINITION_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_XHTML, ATTRIBUTE_DEFINITION_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_BOOLEAN, ATTRIBUTE_VALUE_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_DATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_DATE, ATTRIBUTE_VALUE_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_ENUMERATION, ATTRIBUTE_VALUE_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_INTEGER, ATTRIBUTE_VALUE_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_REAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_REAL, ATTRIBUTE_VALUE_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_STRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_STRING, ATTRIBUTE_VALUE_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_XHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_XHTML, ATTRIBUTE_VALUE_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ALTERNATIVE_ID:
		tmp := GetInstanceDBFromInstance[models.A_ALTERNATIVE_ID, A_ALTERNATIVE_IDDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF, A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_DATE_REF, A_ATTRIBUTE_DEFINITION_DATE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF, A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF, A_ATTRIBUTE_DEFINITION_INTEGER_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_REAL_REF, A_ATTRIBUTE_DEFINITION_REAL_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_STRING_REF, A_ATTRIBUTE_DEFINITION_STRING_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_XHTML_REF, A_ATTRIBUTE_DEFINITION_XHTML_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_BOOLEAN, A_ATTRIBUTE_VALUE_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_DATE:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_DATE, A_ATTRIBUTE_VALUE_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_ENUMERATION, A_ATTRIBUTE_VALUE_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_INTEGER, A_ATTRIBUTE_VALUE_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_REAL:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_REAL, A_ATTRIBUTE_VALUE_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_STRING:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_STRING, A_ATTRIBUTE_VALUE_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_XHTML:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_XHTML, A_ATTRIBUTE_VALUE_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_XHTML_1, A_ATTRIBUTE_VALUE_XHTML_1DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_CHILDREN:
		tmp := GetInstanceDBFromInstance[models.A_CHILDREN, A_CHILDRENDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_CORE_CONTENT:
		tmp := GetInstanceDBFromInstance[models.A_CORE_CONTENT, A_CORE_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPES:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPES, A_DATATYPESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_BOOLEAN_REF, A_DATATYPE_DEFINITION_BOOLEAN_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_DATE_REF, A_DATATYPE_DEFINITION_DATE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_ENUMERATION_REF, A_DATATYPE_DEFINITION_ENUMERATION_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_INTEGER_REF, A_DATATYPE_DEFINITION_INTEGER_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_REAL_REF, A_DATATYPE_DEFINITION_REAL_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_STRING_REF, A_DATATYPE_DEFINITION_STRING_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_XHTML_REF, A_DATATYPE_DEFINITION_XHTML_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_EDITABLE_ATTS:
		tmp := GetInstanceDBFromInstance[models.A_EDITABLE_ATTS, A_EDITABLE_ATTSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ENUM_VALUE_REF:
		tmp := GetInstanceDBFromInstance[models.A_ENUM_VALUE_REF, A_ENUM_VALUE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_OBJECT:
		tmp := GetInstanceDBFromInstance[models.A_OBJECT, A_OBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_PROPERTIES:
		tmp := GetInstanceDBFromInstance[models.A_PROPERTIES, A_PROPERTIESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_RELATION_GROUP_TYPE_REF:
		tmp := GetInstanceDBFromInstance[models.A_RELATION_GROUP_TYPE_REF, A_RELATION_GROUP_TYPE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SOURCE_1:
		tmp := GetInstanceDBFromInstance[models.A_SOURCE_1, A_SOURCE_1DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SOURCE_SPECIFICATION_1:
		tmp := GetInstanceDBFromInstance[models.A_SOURCE_SPECIFICATION_1, A_SOURCE_SPECIFICATION_1DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPECIFICATIONS:
		tmp := GetInstanceDBFromInstance[models.A_SPECIFICATIONS, A_SPECIFICATIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPECIFICATION_TYPE_REF:
		tmp := GetInstanceDBFromInstance[models.A_SPECIFICATION_TYPE_REF, A_SPECIFICATION_TYPE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPECIFIED_VALUES:
		tmp := GetInstanceDBFromInstance[models.A_SPECIFIED_VALUES, A_SPECIFIED_VALUESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_ATTRIBUTES:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_ATTRIBUTES, A_SPEC_ATTRIBUTESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_OBJECTS:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_OBJECTS, A_SPEC_OBJECTSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_OBJECT_TYPE_REF:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_OBJECT_TYPE_REF, A_SPEC_OBJECT_TYPE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_RELATIONS:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_RELATIONS, A_SPEC_RELATIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_RELATION_GROUPS:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_RELATION_GROUPS, A_SPEC_RELATION_GROUPSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_RELATION_REF:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_RELATION_REF, A_SPEC_RELATION_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_RELATION_TYPE_REF:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_RELATION_TYPE_REF, A_SPEC_RELATION_TYPE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_TYPES:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_TYPES, A_SPEC_TYPESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_THE_HEADER:
		tmp := GetInstanceDBFromInstance[models.A_THE_HEADER, A_THE_HEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_TOOL_EXTENSIONS:
		tmp := GetInstanceDBFromInstance[models.A_TOOL_EXTENSIONS, A_TOOL_EXTENSIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_BOOLEAN, DATATYPE_DEFINITION_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_DATE:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_DATE, DATATYPE_DEFINITION_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_ENUMERATION, DATATYPE_DEFINITION_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_INTEGER:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_INTEGER, DATATYPE_DEFINITION_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_REAL:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_REAL, DATATYPE_DEFINITION_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_STRING:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_STRING, DATATYPE_DEFINITION_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_XHTML:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_XHTML, DATATYPE_DEFINITION_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.EMBEDDED_VALUE:
		tmp := GetInstanceDBFromInstance[models.EMBEDDED_VALUE, EMBEDDED_VALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ENUM_VALUE:
		tmp := GetInstanceDBFromInstance[models.ENUM_VALUE, ENUM_VALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.JpgImage:
		tmp := GetInstanceDBFromInstance[models.JpgImage, JpgImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Kill:
		tmp := GetInstanceDBFromInstance[models.Kill, KillDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_SPECIFICATION_Nodes_expandedEntry:
		tmp := GetInstanceDBFromInstance[models.Map_SPECIFICATION_Nodes_expandedEntry, Map_SPECIFICATION_Nodes_expandedEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		tmp := GetInstanceDBFromInstance[models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry, Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		tmp := GetInstanceDBFromInstance[models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry, Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_SPEC_OBJECT_TYPE_showNameEntry:
		tmp := GetInstanceDBFromInstance[models.Map_SPEC_OBJECT_TYPE_showNameEntry, Map_SPEC_OBJECT_TYPE_showNameEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.PngImage:
		tmp := GetInstanceDBFromInstance[models.PngImage, PngImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATION_GROUP:
		tmp := GetInstanceDBFromInstance[models.RELATION_GROUP, RELATION_GROUPDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATION_GROUP_TYPE:
		tmp := GetInstanceDBFromInstance[models.RELATION_GROUP_TYPE, RELATION_GROUP_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF:
		tmp := GetInstanceDBFromInstance[models.REQ_IF, REQ_IFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_CONTENT:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_CONTENT, REQ_IF_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_HEADER:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_HEADER, REQ_IF_HEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_TOOL_EXTENSION:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_TOOL_EXTENSION, REQ_IF_TOOL_EXTENSIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RenderingConfiguration:
		tmp := GetInstanceDBFromInstance[models.RenderingConfiguration, RenderingConfigurationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION, SPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION_TYPE, SPECIFICATION_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_HIERARCHY:
		tmp := GetInstanceDBFromInstance[models.SPEC_HIERARCHY, SPEC_HIERARCHYDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_OBJECT:
		tmp := GetInstanceDBFromInstance[models.SPEC_OBJECT, SPEC_OBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_OBJECT_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPEC_OBJECT_TYPE, SPEC_OBJECT_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_RELATION:
		tmp := GetInstanceDBFromInstance[models.SPEC_RELATION, SPEC_RELATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_RELATION_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPEC_RELATION_TYPE, SPEC_RELATION_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSite:
		tmp := GetInstanceDBFromInstance[models.StaticWebSite, StaticWebSiteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSiteChapter:
		tmp := GetInstanceDBFromInstance[models.StaticWebSiteChapter, StaticWebSiteChapterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSiteGeneratedImage:
		tmp := GetInstanceDBFromInstance[models.StaticWebSiteGeneratedImage, StaticWebSiteGeneratedImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSiteImage:
		tmp := GetInstanceDBFromInstance[models.StaticWebSiteImage, StaticWebSiteImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSiteParagraph:
		tmp := GetInstanceDBFromInstance[models.StaticWebSiteParagraph, StaticWebSiteParagraphDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SvgImage:
		tmp := GetInstanceDBFromInstance[models.SvgImage, SvgImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XHTML_CONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTML_CONTENT, XHTML_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.ALTERNATIVE_ID:
		tmp := GetInstanceDBFromInstance[models.ALTERNATIVE_ID, ALTERNATIVE_IDDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_BOOLEAN, ATTRIBUTE_DEFINITION_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_DATE, ATTRIBUTE_DEFINITION_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_ENUMERATION, ATTRIBUTE_DEFINITION_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_INTEGER, ATTRIBUTE_DEFINITION_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_REAL, ATTRIBUTE_DEFINITION_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_STRING, ATTRIBUTE_DEFINITION_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_DEFINITION_XHTML, ATTRIBUTE_DEFINITION_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_BOOLEAN, ATTRIBUTE_VALUE_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_DATE:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_DATE, ATTRIBUTE_VALUE_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_ENUMERATION, ATTRIBUTE_VALUE_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_INTEGER, ATTRIBUTE_VALUE_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_REAL:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_REAL, ATTRIBUTE_VALUE_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_STRING:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_STRING, ATTRIBUTE_VALUE_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ATTRIBUTE_VALUE_XHTML:
		tmp := GetInstanceDBFromInstance[models.ATTRIBUTE_VALUE_XHTML, ATTRIBUTE_VALUE_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ALTERNATIVE_ID:
		tmp := GetInstanceDBFromInstance[models.A_ALTERNATIVE_ID, A_ALTERNATIVE_IDDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF, A_ATTRIBUTE_DEFINITION_BOOLEAN_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_DATE_REF, A_ATTRIBUTE_DEFINITION_DATE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF, A_ATTRIBUTE_DEFINITION_ENUMERATION_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF, A_ATTRIBUTE_DEFINITION_INTEGER_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_REAL_REF, A_ATTRIBUTE_DEFINITION_REAL_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_STRING_REF, A_ATTRIBUTE_DEFINITION_STRING_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_DEFINITION_XHTML_REF, A_ATTRIBUTE_DEFINITION_XHTML_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_BOOLEAN, A_ATTRIBUTE_VALUE_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_DATE:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_DATE, A_ATTRIBUTE_VALUE_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_ENUMERATION, A_ATTRIBUTE_VALUE_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_INTEGER, A_ATTRIBUTE_VALUE_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_REAL:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_REAL, A_ATTRIBUTE_VALUE_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_STRING:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_STRING, A_ATTRIBUTE_VALUE_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_XHTML:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_XHTML, A_ATTRIBUTE_VALUE_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		tmp := GetInstanceDBFromInstance[models.A_ATTRIBUTE_VALUE_XHTML_1, A_ATTRIBUTE_VALUE_XHTML_1DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_CHILDREN:
		tmp := GetInstanceDBFromInstance[models.A_CHILDREN, A_CHILDRENDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_CORE_CONTENT:
		tmp := GetInstanceDBFromInstance[models.A_CORE_CONTENT, A_CORE_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPES:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPES, A_DATATYPESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_BOOLEAN_REF, A_DATATYPE_DEFINITION_BOOLEAN_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_DATE_REF, A_DATATYPE_DEFINITION_DATE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_ENUMERATION_REF, A_DATATYPE_DEFINITION_ENUMERATION_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_INTEGER_REF, A_DATATYPE_DEFINITION_INTEGER_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_REAL_REF, A_DATATYPE_DEFINITION_REAL_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_STRING_REF, A_DATATYPE_DEFINITION_STRING_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		tmp := GetInstanceDBFromInstance[models.A_DATATYPE_DEFINITION_XHTML_REF, A_DATATYPE_DEFINITION_XHTML_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_EDITABLE_ATTS:
		tmp := GetInstanceDBFromInstance[models.A_EDITABLE_ATTS, A_EDITABLE_ATTSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_ENUM_VALUE_REF:
		tmp := GetInstanceDBFromInstance[models.A_ENUM_VALUE_REF, A_ENUM_VALUE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_OBJECT:
		tmp := GetInstanceDBFromInstance[models.A_OBJECT, A_OBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_PROPERTIES:
		tmp := GetInstanceDBFromInstance[models.A_PROPERTIES, A_PROPERTIESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_RELATION_GROUP_TYPE_REF:
		tmp := GetInstanceDBFromInstance[models.A_RELATION_GROUP_TYPE_REF, A_RELATION_GROUP_TYPE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SOURCE_1:
		tmp := GetInstanceDBFromInstance[models.A_SOURCE_1, A_SOURCE_1DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SOURCE_SPECIFICATION_1:
		tmp := GetInstanceDBFromInstance[models.A_SOURCE_SPECIFICATION_1, A_SOURCE_SPECIFICATION_1DB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPECIFICATIONS:
		tmp := GetInstanceDBFromInstance[models.A_SPECIFICATIONS, A_SPECIFICATIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPECIFICATION_TYPE_REF:
		tmp := GetInstanceDBFromInstance[models.A_SPECIFICATION_TYPE_REF, A_SPECIFICATION_TYPE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPECIFIED_VALUES:
		tmp := GetInstanceDBFromInstance[models.A_SPECIFIED_VALUES, A_SPECIFIED_VALUESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_ATTRIBUTES:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_ATTRIBUTES, A_SPEC_ATTRIBUTESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_OBJECTS:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_OBJECTS, A_SPEC_OBJECTSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_OBJECT_TYPE_REF:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_OBJECT_TYPE_REF, A_SPEC_OBJECT_TYPE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_RELATIONS:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_RELATIONS, A_SPEC_RELATIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_RELATION_GROUPS:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_RELATION_GROUPS, A_SPEC_RELATION_GROUPSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_RELATION_REF:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_RELATION_REF, A_SPEC_RELATION_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_RELATION_TYPE_REF:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_RELATION_TYPE_REF, A_SPEC_RELATION_TYPE_REFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_SPEC_TYPES:
		tmp := GetInstanceDBFromInstance[models.A_SPEC_TYPES, A_SPEC_TYPESDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_THE_HEADER:
		tmp := GetInstanceDBFromInstance[models.A_THE_HEADER, A_THE_HEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.A_TOOL_EXTENSIONS:
		tmp := GetInstanceDBFromInstance[models.A_TOOL_EXTENSIONS, A_TOOL_EXTENSIONSDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_BOOLEAN, DATATYPE_DEFINITION_BOOLEANDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_DATE:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_DATE, DATATYPE_DEFINITION_DATEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_ENUMERATION, DATATYPE_DEFINITION_ENUMERATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_INTEGER:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_INTEGER, DATATYPE_DEFINITION_INTEGERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_REAL:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_REAL, DATATYPE_DEFINITION_REALDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_STRING:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_STRING, DATATYPE_DEFINITION_STRINGDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.DATATYPE_DEFINITION_XHTML:
		tmp := GetInstanceDBFromInstance[models.DATATYPE_DEFINITION_XHTML, DATATYPE_DEFINITION_XHTMLDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.EMBEDDED_VALUE:
		tmp := GetInstanceDBFromInstance[models.EMBEDDED_VALUE, EMBEDDED_VALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.ENUM_VALUE:
		tmp := GetInstanceDBFromInstance[models.ENUM_VALUE, ENUM_VALUEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.JpgImage:
		tmp := GetInstanceDBFromInstance[models.JpgImage, JpgImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Kill:
		tmp := GetInstanceDBFromInstance[models.Kill, KillDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_BOOLEAN_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_DATE_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_DATE_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_DATE_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_INTEGER_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_REAL_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_REAL_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_REAL_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_STRING_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_STRING_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_STRING_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntry, Map_ATTRIBUTE_DEFINITION_XHTML_ShowInSubjectEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry, Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry:
		tmp := GetInstanceDBFromInstance[models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry, Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_SPECIFICATION_Nodes_expandedEntry:
		tmp := GetInstanceDBFromInstance[models.Map_SPECIFICATION_Nodes_expandedEntry, Map_SPECIFICATION_Nodes_expandedEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry:
		tmp := GetInstanceDBFromInstance[models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry, Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry:
		tmp := GetInstanceDBFromInstance[models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry, Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Map_SPEC_OBJECT_TYPE_showNameEntry:
		tmp := GetInstanceDBFromInstance[models.Map_SPEC_OBJECT_TYPE_showNameEntry, Map_SPEC_OBJECT_TYPE_showNameEntryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.PngImage:
		tmp := GetInstanceDBFromInstance[models.PngImage, PngImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATION_GROUP:
		tmp := GetInstanceDBFromInstance[models.RELATION_GROUP, RELATION_GROUPDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RELATION_GROUP_TYPE:
		tmp := GetInstanceDBFromInstance[models.RELATION_GROUP_TYPE, RELATION_GROUP_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF:
		tmp := GetInstanceDBFromInstance[models.REQ_IF, REQ_IFDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_CONTENT:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_CONTENT, REQ_IF_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_HEADER:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_HEADER, REQ_IF_HEADERDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.REQ_IF_TOOL_EXTENSION:
		tmp := GetInstanceDBFromInstance[models.REQ_IF_TOOL_EXTENSION, REQ_IF_TOOL_EXTENSIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.RenderingConfiguration:
		tmp := GetInstanceDBFromInstance[models.RenderingConfiguration, RenderingConfigurationDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION, SPECIFICATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPECIFICATION_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPECIFICATION_TYPE, SPECIFICATION_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_HIERARCHY:
		tmp := GetInstanceDBFromInstance[models.SPEC_HIERARCHY, SPEC_HIERARCHYDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_OBJECT:
		tmp := GetInstanceDBFromInstance[models.SPEC_OBJECT, SPEC_OBJECTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_OBJECT_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPEC_OBJECT_TYPE, SPEC_OBJECT_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_RELATION:
		tmp := GetInstanceDBFromInstance[models.SPEC_RELATION, SPEC_RELATIONDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SPEC_RELATION_TYPE:
		tmp := GetInstanceDBFromInstance[models.SPEC_RELATION_TYPE, SPEC_RELATION_TYPEDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSite:
		tmp := GetInstanceDBFromInstance[models.StaticWebSite, StaticWebSiteDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSiteChapter:
		tmp := GetInstanceDBFromInstance[models.StaticWebSiteChapter, StaticWebSiteChapterDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSiteGeneratedImage:
		tmp := GetInstanceDBFromInstance[models.StaticWebSiteGeneratedImage, StaticWebSiteGeneratedImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSiteImage:
		tmp := GetInstanceDBFromInstance[models.StaticWebSiteImage, StaticWebSiteImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.StaticWebSiteParagraph:
		tmp := GetInstanceDBFromInstance[models.StaticWebSiteParagraph, StaticWebSiteParagraphDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.SvgImage:
		tmp := GetInstanceDBFromInstance[models.SvgImage, SvgImageDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XHTML_CONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTML_CONTENT, XHTML_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
