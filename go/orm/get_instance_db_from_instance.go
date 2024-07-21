// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongreqif/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | ALTERNATIVE_IDDB | ATTRIBUTE_DEFINITION_BOOLEANDB | ATTRIBUTE_DEFINITION_DATEDB | ATTRIBUTE_DEFINITION_ENUMERATIONDB | ATTRIBUTE_DEFINITION_INTEGERDB | ATTRIBUTE_DEFINITION_REALDB | ATTRIBUTE_DEFINITION_STRINGDB | ATTRIBUTE_DEFINITION_XHTMLDB | ATTRIBUTE_VALUE_BOOLEANDB | ATTRIBUTE_VALUE_DATEDB | ATTRIBUTE_VALUE_ENUMERATIONDB | ATTRIBUTE_VALUE_INTEGERDB | ATTRIBUTE_VALUE_REALDB | ATTRIBUTE_VALUE_STRINGDB | ATTRIBUTE_VALUE_XHTMLDB | AnyTypeDB | DATATYPE_DEFINITION_BOOLEANDB | DATATYPE_DEFINITION_DATEDB | DATATYPE_DEFINITION_ENUMERATIONDB | DATATYPE_DEFINITION_INTEGERDB | DATATYPE_DEFINITION_REALDB | DATATYPE_DEFINITION_STRINGDB | DATATYPE_DEFINITION_XHTMLDB | EMBEDDED_VALUEDB | ENUM_VALUEDB | RELATION_GROUPDB | RELATION_GROUP_TYPEDB | REQ_IFDB | REQ_IF_CONTENTDB | REQ_IF_HEADERDB | REQ_IF_TOOL_EXTENSIONDB | SPECIFICATIONDB | SPECIFICATION_TYPEDB | SPEC_HIERARCHYDB | SPEC_OBJECTDB | SPEC_OBJECT_TYPEDB | SPEC_RELATIONDB | SPEC_RELATION_TYPEDB | XHTML_CONTENTDB | Xhtml_InlPres_typeDB | Xhtml_a_typeDB | Xhtml_abbr_typeDB | Xhtml_acronym_typeDB | Xhtml_address_typeDB | Xhtml_blockquote_typeDB | Xhtml_br_typeDB | Xhtml_caption_typeDB | Xhtml_cite_typeDB | Xhtml_code_typeDB | Xhtml_col_typeDB | Xhtml_colgroup_typeDB | Xhtml_dd_typeDB | Xhtml_dfn_typeDB | Xhtml_div_typeDB | Xhtml_dl_typeDB | Xhtml_dt_typeDB | Xhtml_edit_typeDB | Xhtml_em_typeDB | Xhtml_h1_typeDB | Xhtml_h2_typeDB | Xhtml_h3_typeDB | Xhtml_h4_typeDB | Xhtml_h5_typeDB | Xhtml_h6_typeDB | Xhtml_heading_typeDB | Xhtml_hr_typeDB | Xhtml_kbd_typeDB | Xhtml_li_typeDB | Xhtml_object_typeDB | Xhtml_ol_typeDB | Xhtml_p_typeDB | Xhtml_param_typeDB | Xhtml_pre_typeDB | Xhtml_q_typeDB | Xhtml_samp_typeDB | Xhtml_span_typeDB | Xhtml_strong_typeDB | Xhtml_table_typeDB | Xhtml_tbody_typeDB | Xhtml_td_typeDB | Xhtml_tfoot_typeDB | Xhtml_th_typeDB | Xhtml_thead_typeDB | Xhtml_tr_typeDB | Xhtml_ul_typeDB | Xhtml_var_typeDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
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
	case *models.AnyType:
		anytypeInstance := any(concreteInstance).(*models.AnyType)
		ret2 := backRepo.BackRepoAnyType.GetAnyTypeDBFromAnyTypePtr(anytypeInstance)
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
	case *models.XHTML_CONTENT:
		xhtml_contentInstance := any(concreteInstance).(*models.XHTML_CONTENT)
		ret2 := backRepo.BackRepoXHTML_CONTENT.GetXHTML_CONTENTDBFromXHTML_CONTENTPtr(xhtml_contentInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_InlPres_type:
		xhtml_inlpres_typeInstance := any(concreteInstance).(*models.Xhtml_InlPres_type)
		ret2 := backRepo.BackRepoXhtml_InlPres_type.GetXhtml_InlPres_typeDBFromXhtml_InlPres_typePtr(xhtml_inlpres_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_a_type:
		xhtml_a_typeInstance := any(concreteInstance).(*models.Xhtml_a_type)
		ret2 := backRepo.BackRepoXhtml_a_type.GetXhtml_a_typeDBFromXhtml_a_typePtr(xhtml_a_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_abbr_type:
		xhtml_abbr_typeInstance := any(concreteInstance).(*models.Xhtml_abbr_type)
		ret2 := backRepo.BackRepoXhtml_abbr_type.GetXhtml_abbr_typeDBFromXhtml_abbr_typePtr(xhtml_abbr_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_acronym_type:
		xhtml_acronym_typeInstance := any(concreteInstance).(*models.Xhtml_acronym_type)
		ret2 := backRepo.BackRepoXhtml_acronym_type.GetXhtml_acronym_typeDBFromXhtml_acronym_typePtr(xhtml_acronym_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_address_type:
		xhtml_address_typeInstance := any(concreteInstance).(*models.Xhtml_address_type)
		ret2 := backRepo.BackRepoXhtml_address_type.GetXhtml_address_typeDBFromXhtml_address_typePtr(xhtml_address_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_blockquote_type:
		xhtml_blockquote_typeInstance := any(concreteInstance).(*models.Xhtml_blockquote_type)
		ret2 := backRepo.BackRepoXhtml_blockquote_type.GetXhtml_blockquote_typeDBFromXhtml_blockquote_typePtr(xhtml_blockquote_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_br_type:
		xhtml_br_typeInstance := any(concreteInstance).(*models.Xhtml_br_type)
		ret2 := backRepo.BackRepoXhtml_br_type.GetXhtml_br_typeDBFromXhtml_br_typePtr(xhtml_br_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_caption_type:
		xhtml_caption_typeInstance := any(concreteInstance).(*models.Xhtml_caption_type)
		ret2 := backRepo.BackRepoXhtml_caption_type.GetXhtml_caption_typeDBFromXhtml_caption_typePtr(xhtml_caption_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_cite_type:
		xhtml_cite_typeInstance := any(concreteInstance).(*models.Xhtml_cite_type)
		ret2 := backRepo.BackRepoXhtml_cite_type.GetXhtml_cite_typeDBFromXhtml_cite_typePtr(xhtml_cite_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_code_type:
		xhtml_code_typeInstance := any(concreteInstance).(*models.Xhtml_code_type)
		ret2 := backRepo.BackRepoXhtml_code_type.GetXhtml_code_typeDBFromXhtml_code_typePtr(xhtml_code_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_col_type:
		xhtml_col_typeInstance := any(concreteInstance).(*models.Xhtml_col_type)
		ret2 := backRepo.BackRepoXhtml_col_type.GetXhtml_col_typeDBFromXhtml_col_typePtr(xhtml_col_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_colgroup_type:
		xhtml_colgroup_typeInstance := any(concreteInstance).(*models.Xhtml_colgroup_type)
		ret2 := backRepo.BackRepoXhtml_colgroup_type.GetXhtml_colgroup_typeDBFromXhtml_colgroup_typePtr(xhtml_colgroup_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_dd_type:
		xhtml_dd_typeInstance := any(concreteInstance).(*models.Xhtml_dd_type)
		ret2 := backRepo.BackRepoXhtml_dd_type.GetXhtml_dd_typeDBFromXhtml_dd_typePtr(xhtml_dd_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_dfn_type:
		xhtml_dfn_typeInstance := any(concreteInstance).(*models.Xhtml_dfn_type)
		ret2 := backRepo.BackRepoXhtml_dfn_type.GetXhtml_dfn_typeDBFromXhtml_dfn_typePtr(xhtml_dfn_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_div_type:
		xhtml_div_typeInstance := any(concreteInstance).(*models.Xhtml_div_type)
		ret2 := backRepo.BackRepoXhtml_div_type.GetXhtml_div_typeDBFromXhtml_div_typePtr(xhtml_div_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_dl_type:
		xhtml_dl_typeInstance := any(concreteInstance).(*models.Xhtml_dl_type)
		ret2 := backRepo.BackRepoXhtml_dl_type.GetXhtml_dl_typeDBFromXhtml_dl_typePtr(xhtml_dl_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_dt_type:
		xhtml_dt_typeInstance := any(concreteInstance).(*models.Xhtml_dt_type)
		ret2 := backRepo.BackRepoXhtml_dt_type.GetXhtml_dt_typeDBFromXhtml_dt_typePtr(xhtml_dt_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_edit_type:
		xhtml_edit_typeInstance := any(concreteInstance).(*models.Xhtml_edit_type)
		ret2 := backRepo.BackRepoXhtml_edit_type.GetXhtml_edit_typeDBFromXhtml_edit_typePtr(xhtml_edit_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_em_type:
		xhtml_em_typeInstance := any(concreteInstance).(*models.Xhtml_em_type)
		ret2 := backRepo.BackRepoXhtml_em_type.GetXhtml_em_typeDBFromXhtml_em_typePtr(xhtml_em_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_h1_type:
		xhtml_h1_typeInstance := any(concreteInstance).(*models.Xhtml_h1_type)
		ret2 := backRepo.BackRepoXhtml_h1_type.GetXhtml_h1_typeDBFromXhtml_h1_typePtr(xhtml_h1_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_h2_type:
		xhtml_h2_typeInstance := any(concreteInstance).(*models.Xhtml_h2_type)
		ret2 := backRepo.BackRepoXhtml_h2_type.GetXhtml_h2_typeDBFromXhtml_h2_typePtr(xhtml_h2_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_h3_type:
		xhtml_h3_typeInstance := any(concreteInstance).(*models.Xhtml_h3_type)
		ret2 := backRepo.BackRepoXhtml_h3_type.GetXhtml_h3_typeDBFromXhtml_h3_typePtr(xhtml_h3_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_h4_type:
		xhtml_h4_typeInstance := any(concreteInstance).(*models.Xhtml_h4_type)
		ret2 := backRepo.BackRepoXhtml_h4_type.GetXhtml_h4_typeDBFromXhtml_h4_typePtr(xhtml_h4_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_h5_type:
		xhtml_h5_typeInstance := any(concreteInstance).(*models.Xhtml_h5_type)
		ret2 := backRepo.BackRepoXhtml_h5_type.GetXhtml_h5_typeDBFromXhtml_h5_typePtr(xhtml_h5_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_h6_type:
		xhtml_h6_typeInstance := any(concreteInstance).(*models.Xhtml_h6_type)
		ret2 := backRepo.BackRepoXhtml_h6_type.GetXhtml_h6_typeDBFromXhtml_h6_typePtr(xhtml_h6_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_heading_type:
		xhtml_heading_typeInstance := any(concreteInstance).(*models.Xhtml_heading_type)
		ret2 := backRepo.BackRepoXhtml_heading_type.GetXhtml_heading_typeDBFromXhtml_heading_typePtr(xhtml_heading_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_hr_type:
		xhtml_hr_typeInstance := any(concreteInstance).(*models.Xhtml_hr_type)
		ret2 := backRepo.BackRepoXhtml_hr_type.GetXhtml_hr_typeDBFromXhtml_hr_typePtr(xhtml_hr_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_kbd_type:
		xhtml_kbd_typeInstance := any(concreteInstance).(*models.Xhtml_kbd_type)
		ret2 := backRepo.BackRepoXhtml_kbd_type.GetXhtml_kbd_typeDBFromXhtml_kbd_typePtr(xhtml_kbd_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_li_type:
		xhtml_li_typeInstance := any(concreteInstance).(*models.Xhtml_li_type)
		ret2 := backRepo.BackRepoXhtml_li_type.GetXhtml_li_typeDBFromXhtml_li_typePtr(xhtml_li_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_object_type:
		xhtml_object_typeInstance := any(concreteInstance).(*models.Xhtml_object_type)
		ret2 := backRepo.BackRepoXhtml_object_type.GetXhtml_object_typeDBFromXhtml_object_typePtr(xhtml_object_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_ol_type:
		xhtml_ol_typeInstance := any(concreteInstance).(*models.Xhtml_ol_type)
		ret2 := backRepo.BackRepoXhtml_ol_type.GetXhtml_ol_typeDBFromXhtml_ol_typePtr(xhtml_ol_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_p_type:
		xhtml_p_typeInstance := any(concreteInstance).(*models.Xhtml_p_type)
		ret2 := backRepo.BackRepoXhtml_p_type.GetXhtml_p_typeDBFromXhtml_p_typePtr(xhtml_p_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_param_type:
		xhtml_param_typeInstance := any(concreteInstance).(*models.Xhtml_param_type)
		ret2 := backRepo.BackRepoXhtml_param_type.GetXhtml_param_typeDBFromXhtml_param_typePtr(xhtml_param_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_pre_type:
		xhtml_pre_typeInstance := any(concreteInstance).(*models.Xhtml_pre_type)
		ret2 := backRepo.BackRepoXhtml_pre_type.GetXhtml_pre_typeDBFromXhtml_pre_typePtr(xhtml_pre_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_q_type:
		xhtml_q_typeInstance := any(concreteInstance).(*models.Xhtml_q_type)
		ret2 := backRepo.BackRepoXhtml_q_type.GetXhtml_q_typeDBFromXhtml_q_typePtr(xhtml_q_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_samp_type:
		xhtml_samp_typeInstance := any(concreteInstance).(*models.Xhtml_samp_type)
		ret2 := backRepo.BackRepoXhtml_samp_type.GetXhtml_samp_typeDBFromXhtml_samp_typePtr(xhtml_samp_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_span_type:
		xhtml_span_typeInstance := any(concreteInstance).(*models.Xhtml_span_type)
		ret2 := backRepo.BackRepoXhtml_span_type.GetXhtml_span_typeDBFromXhtml_span_typePtr(xhtml_span_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_strong_type:
		xhtml_strong_typeInstance := any(concreteInstance).(*models.Xhtml_strong_type)
		ret2 := backRepo.BackRepoXhtml_strong_type.GetXhtml_strong_typeDBFromXhtml_strong_typePtr(xhtml_strong_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_table_type:
		xhtml_table_typeInstance := any(concreteInstance).(*models.Xhtml_table_type)
		ret2 := backRepo.BackRepoXhtml_table_type.GetXhtml_table_typeDBFromXhtml_table_typePtr(xhtml_table_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_tbody_type:
		xhtml_tbody_typeInstance := any(concreteInstance).(*models.Xhtml_tbody_type)
		ret2 := backRepo.BackRepoXhtml_tbody_type.GetXhtml_tbody_typeDBFromXhtml_tbody_typePtr(xhtml_tbody_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_td_type:
		xhtml_td_typeInstance := any(concreteInstance).(*models.Xhtml_td_type)
		ret2 := backRepo.BackRepoXhtml_td_type.GetXhtml_td_typeDBFromXhtml_td_typePtr(xhtml_td_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_tfoot_type:
		xhtml_tfoot_typeInstance := any(concreteInstance).(*models.Xhtml_tfoot_type)
		ret2 := backRepo.BackRepoXhtml_tfoot_type.GetXhtml_tfoot_typeDBFromXhtml_tfoot_typePtr(xhtml_tfoot_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_th_type:
		xhtml_th_typeInstance := any(concreteInstance).(*models.Xhtml_th_type)
		ret2 := backRepo.BackRepoXhtml_th_type.GetXhtml_th_typeDBFromXhtml_th_typePtr(xhtml_th_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_thead_type:
		xhtml_thead_typeInstance := any(concreteInstance).(*models.Xhtml_thead_type)
		ret2 := backRepo.BackRepoXhtml_thead_type.GetXhtml_thead_typeDBFromXhtml_thead_typePtr(xhtml_thead_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_tr_type:
		xhtml_tr_typeInstance := any(concreteInstance).(*models.Xhtml_tr_type)
		ret2 := backRepo.BackRepoXhtml_tr_type.GetXhtml_tr_typeDBFromXhtml_tr_typePtr(xhtml_tr_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_ul_type:
		xhtml_ul_typeInstance := any(concreteInstance).(*models.Xhtml_ul_type)
		ret2 := backRepo.BackRepoXhtml_ul_type.GetXhtml_ul_typeDBFromXhtml_ul_typePtr(xhtml_ul_typeInstance)
		ret = any(ret2).(*T2)
	case *models.Xhtml_var_type:
		xhtml_var_typeInstance := any(concreteInstance).(*models.Xhtml_var_type)
		ret2 := backRepo.BackRepoXhtml_var_type.GetXhtml_var_typeDBFromXhtml_var_typePtr(xhtml_var_typeInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
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
	case *models.AnyType:
		tmp := GetInstanceDBFromInstance[models.AnyType, AnyTypeDB](
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
	case *models.XHTML_CONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTML_CONTENT, XHTML_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_InlPres_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_InlPres_type, Xhtml_InlPres_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_a_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_a_type, Xhtml_a_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_abbr_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_abbr_type, Xhtml_abbr_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_acronym_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_acronym_type, Xhtml_acronym_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_address_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_address_type, Xhtml_address_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_blockquote_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_blockquote_type, Xhtml_blockquote_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_br_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_br_type, Xhtml_br_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_caption_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_caption_type, Xhtml_caption_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_cite_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_cite_type, Xhtml_cite_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_code_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_code_type, Xhtml_code_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_col_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_col_type, Xhtml_col_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_colgroup_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_colgroup_type, Xhtml_colgroup_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_dd_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_dd_type, Xhtml_dd_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_dfn_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_dfn_type, Xhtml_dfn_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_div_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_div_type, Xhtml_div_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_dl_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_dl_type, Xhtml_dl_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_dt_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_dt_type, Xhtml_dt_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_edit_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_edit_type, Xhtml_edit_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_em_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_em_type, Xhtml_em_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h1_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h1_type, Xhtml_h1_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h2_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h2_type, Xhtml_h2_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h3_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h3_type, Xhtml_h3_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h4_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h4_type, Xhtml_h4_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h5_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h5_type, Xhtml_h5_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h6_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h6_type, Xhtml_h6_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_heading_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_heading_type, Xhtml_heading_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_hr_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_hr_type, Xhtml_hr_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_kbd_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_kbd_type, Xhtml_kbd_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_li_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_li_type, Xhtml_li_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_object_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_object_type, Xhtml_object_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_ol_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_ol_type, Xhtml_ol_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_p_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_p_type, Xhtml_p_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_param_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_param_type, Xhtml_param_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_pre_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_pre_type, Xhtml_pre_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_q_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_q_type, Xhtml_q_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_samp_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_samp_type, Xhtml_samp_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_span_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_span_type, Xhtml_span_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_strong_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_strong_type, Xhtml_strong_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_table_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_table_type, Xhtml_table_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_tbody_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_tbody_type, Xhtml_tbody_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_td_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_td_type, Xhtml_td_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_tfoot_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_tfoot_type, Xhtml_tfoot_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_th_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_th_type, Xhtml_th_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_thead_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_thead_type, Xhtml_thead_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_tr_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_tr_type, Xhtml_tr_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_ul_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_ul_type, Xhtml_ul_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_var_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_var_type, Xhtml_var_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
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
	case *models.AnyType:
		tmp := GetInstanceDBFromInstance[models.AnyType, AnyTypeDB](
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
	case *models.XHTML_CONTENT:
		tmp := GetInstanceDBFromInstance[models.XHTML_CONTENT, XHTML_CONTENTDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_InlPres_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_InlPres_type, Xhtml_InlPres_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_a_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_a_type, Xhtml_a_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_abbr_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_abbr_type, Xhtml_abbr_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_acronym_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_acronym_type, Xhtml_acronym_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_address_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_address_type, Xhtml_address_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_blockquote_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_blockquote_type, Xhtml_blockquote_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_br_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_br_type, Xhtml_br_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_caption_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_caption_type, Xhtml_caption_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_cite_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_cite_type, Xhtml_cite_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_code_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_code_type, Xhtml_code_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_col_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_col_type, Xhtml_col_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_colgroup_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_colgroup_type, Xhtml_colgroup_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_dd_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_dd_type, Xhtml_dd_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_dfn_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_dfn_type, Xhtml_dfn_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_div_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_div_type, Xhtml_div_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_dl_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_dl_type, Xhtml_dl_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_dt_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_dt_type, Xhtml_dt_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_edit_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_edit_type, Xhtml_edit_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_em_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_em_type, Xhtml_em_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h1_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h1_type, Xhtml_h1_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h2_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h2_type, Xhtml_h2_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h3_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h3_type, Xhtml_h3_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h4_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h4_type, Xhtml_h4_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h5_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h5_type, Xhtml_h5_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_h6_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_h6_type, Xhtml_h6_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_heading_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_heading_type, Xhtml_heading_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_hr_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_hr_type, Xhtml_hr_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_kbd_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_kbd_type, Xhtml_kbd_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_li_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_li_type, Xhtml_li_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_object_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_object_type, Xhtml_object_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_ol_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_ol_type, Xhtml_ol_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_p_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_p_type, Xhtml_p_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_param_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_param_type, Xhtml_param_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_pre_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_pre_type, Xhtml_pre_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_q_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_q_type, Xhtml_q_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_samp_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_samp_type, Xhtml_samp_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_span_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_span_type, Xhtml_span_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_strong_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_strong_type, Xhtml_strong_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_table_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_table_type, Xhtml_table_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_tbody_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_tbody_type, Xhtml_tbody_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_td_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_td_type, Xhtml_td_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_tfoot_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_tfoot_type, Xhtml_tfoot_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_th_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_th_type, Xhtml_th_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_thead_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_thead_type, Xhtml_thead_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_tr_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_tr_type, Xhtml_tr_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_ul_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_ul_type, Xhtml_ul_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Xhtml_var_type:
		tmp := GetInstanceDBFromInstance[models.Xhtml_var_type, Xhtml_var_typeDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
