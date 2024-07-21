// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *ALTERNATIVE_ID:
		ok = stage.IsStagedALTERNATIVE_ID(target)

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_BOOLEAN(target)

	case *ATTRIBUTE_DEFINITION_DATE:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_DATE(target)

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_ENUMERATION(target)

	case *ATTRIBUTE_DEFINITION_INTEGER:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_INTEGER(target)

	case *ATTRIBUTE_DEFINITION_REAL:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_REAL(target)

	case *ATTRIBUTE_DEFINITION_STRING:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_STRING(target)

	case *ATTRIBUTE_DEFINITION_XHTML:
		ok = stage.IsStagedATTRIBUTE_DEFINITION_XHTML(target)

	case *ATTRIBUTE_VALUE_BOOLEAN:
		ok = stage.IsStagedATTRIBUTE_VALUE_BOOLEAN(target)

	case *ATTRIBUTE_VALUE_DATE:
		ok = stage.IsStagedATTRIBUTE_VALUE_DATE(target)

	case *ATTRIBUTE_VALUE_ENUMERATION:
		ok = stage.IsStagedATTRIBUTE_VALUE_ENUMERATION(target)

	case *ATTRIBUTE_VALUE_INTEGER:
		ok = stage.IsStagedATTRIBUTE_VALUE_INTEGER(target)

	case *ATTRIBUTE_VALUE_REAL:
		ok = stage.IsStagedATTRIBUTE_VALUE_REAL(target)

	case *ATTRIBUTE_VALUE_STRING:
		ok = stage.IsStagedATTRIBUTE_VALUE_STRING(target)

	case *ATTRIBUTE_VALUE_XHTML:
		ok = stage.IsStagedATTRIBUTE_VALUE_XHTML(target)

	case *AnyType:
		ok = stage.IsStagedAnyType(target)

	case *DATATYPE_DEFINITION_BOOLEAN:
		ok = stage.IsStagedDATATYPE_DEFINITION_BOOLEAN(target)

	case *DATATYPE_DEFINITION_DATE:
		ok = stage.IsStagedDATATYPE_DEFINITION_DATE(target)

	case *DATATYPE_DEFINITION_ENUMERATION:
		ok = stage.IsStagedDATATYPE_DEFINITION_ENUMERATION(target)

	case *DATATYPE_DEFINITION_INTEGER:
		ok = stage.IsStagedDATATYPE_DEFINITION_INTEGER(target)

	case *DATATYPE_DEFINITION_REAL:
		ok = stage.IsStagedDATATYPE_DEFINITION_REAL(target)

	case *DATATYPE_DEFINITION_STRING:
		ok = stage.IsStagedDATATYPE_DEFINITION_STRING(target)

	case *DATATYPE_DEFINITION_XHTML:
		ok = stage.IsStagedDATATYPE_DEFINITION_XHTML(target)

	case *EMBEDDED_VALUE:
		ok = stage.IsStagedEMBEDDED_VALUE(target)

	case *ENUM_VALUE:
		ok = stage.IsStagedENUM_VALUE(target)

	case *RELATION_GROUP:
		ok = stage.IsStagedRELATION_GROUP(target)

	case *RELATION_GROUP_TYPE:
		ok = stage.IsStagedRELATION_GROUP_TYPE(target)

	case *REQ_IF:
		ok = stage.IsStagedREQ_IF(target)

	case *REQ_IF_CONTENT:
		ok = stage.IsStagedREQ_IF_CONTENT(target)

	case *REQ_IF_HEADER:
		ok = stage.IsStagedREQ_IF_HEADER(target)

	case *REQ_IF_TOOL_EXTENSION:
		ok = stage.IsStagedREQ_IF_TOOL_EXTENSION(target)

	case *SPECIFICATION:
		ok = stage.IsStagedSPECIFICATION(target)

	case *SPECIFICATION_TYPE:
		ok = stage.IsStagedSPECIFICATION_TYPE(target)

	case *SPEC_HIERARCHY:
		ok = stage.IsStagedSPEC_HIERARCHY(target)

	case *SPEC_OBJECT:
		ok = stage.IsStagedSPEC_OBJECT(target)

	case *SPEC_OBJECT_TYPE:
		ok = stage.IsStagedSPEC_OBJECT_TYPE(target)

	case *SPEC_RELATION:
		ok = stage.IsStagedSPEC_RELATION(target)

	case *SPEC_RELATION_TYPE:
		ok = stage.IsStagedSPEC_RELATION_TYPE(target)

	case *XHTML_CONTENT:
		ok = stage.IsStagedXHTML_CONTENT(target)

	case *Xhtml_InlPres_type:
		ok = stage.IsStagedXhtml_InlPres_type(target)

	case *Xhtml_a_type:
		ok = stage.IsStagedXhtml_a_type(target)

	case *Xhtml_abbr_type:
		ok = stage.IsStagedXhtml_abbr_type(target)

	case *Xhtml_acronym_type:
		ok = stage.IsStagedXhtml_acronym_type(target)

	case *Xhtml_address_type:
		ok = stage.IsStagedXhtml_address_type(target)

	case *Xhtml_blockquote_type:
		ok = stage.IsStagedXhtml_blockquote_type(target)

	case *Xhtml_br_type:
		ok = stage.IsStagedXhtml_br_type(target)

	case *Xhtml_caption_type:
		ok = stage.IsStagedXhtml_caption_type(target)

	case *Xhtml_cite_type:
		ok = stage.IsStagedXhtml_cite_type(target)

	case *Xhtml_code_type:
		ok = stage.IsStagedXhtml_code_type(target)

	case *Xhtml_col_type:
		ok = stage.IsStagedXhtml_col_type(target)

	case *Xhtml_colgroup_type:
		ok = stage.IsStagedXhtml_colgroup_type(target)

	case *Xhtml_dd_type:
		ok = stage.IsStagedXhtml_dd_type(target)

	case *Xhtml_dfn_type:
		ok = stage.IsStagedXhtml_dfn_type(target)

	case *Xhtml_div_type:
		ok = stage.IsStagedXhtml_div_type(target)

	case *Xhtml_dl_type:
		ok = stage.IsStagedXhtml_dl_type(target)

	case *Xhtml_dt_type:
		ok = stage.IsStagedXhtml_dt_type(target)

	case *Xhtml_edit_type:
		ok = stage.IsStagedXhtml_edit_type(target)

	case *Xhtml_em_type:
		ok = stage.IsStagedXhtml_em_type(target)

	case *Xhtml_h1_type:
		ok = stage.IsStagedXhtml_h1_type(target)

	case *Xhtml_h2_type:
		ok = stage.IsStagedXhtml_h2_type(target)

	case *Xhtml_h3_type:
		ok = stage.IsStagedXhtml_h3_type(target)

	case *Xhtml_h4_type:
		ok = stage.IsStagedXhtml_h4_type(target)

	case *Xhtml_h5_type:
		ok = stage.IsStagedXhtml_h5_type(target)

	case *Xhtml_h6_type:
		ok = stage.IsStagedXhtml_h6_type(target)

	case *Xhtml_heading_type:
		ok = stage.IsStagedXhtml_heading_type(target)

	case *Xhtml_hr_type:
		ok = stage.IsStagedXhtml_hr_type(target)

	case *Xhtml_kbd_type:
		ok = stage.IsStagedXhtml_kbd_type(target)

	case *Xhtml_li_type:
		ok = stage.IsStagedXhtml_li_type(target)

	case *Xhtml_object_type:
		ok = stage.IsStagedXhtml_object_type(target)

	case *Xhtml_ol_type:
		ok = stage.IsStagedXhtml_ol_type(target)

	case *Xhtml_p_type:
		ok = stage.IsStagedXhtml_p_type(target)

	case *Xhtml_param_type:
		ok = stage.IsStagedXhtml_param_type(target)

	case *Xhtml_pre_type:
		ok = stage.IsStagedXhtml_pre_type(target)

	case *Xhtml_q_type:
		ok = stage.IsStagedXhtml_q_type(target)

	case *Xhtml_samp_type:
		ok = stage.IsStagedXhtml_samp_type(target)

	case *Xhtml_span_type:
		ok = stage.IsStagedXhtml_span_type(target)

	case *Xhtml_strong_type:
		ok = stage.IsStagedXhtml_strong_type(target)

	case *Xhtml_table_type:
		ok = stage.IsStagedXhtml_table_type(target)

	case *Xhtml_tbody_type:
		ok = stage.IsStagedXhtml_tbody_type(target)

	case *Xhtml_td_type:
		ok = stage.IsStagedXhtml_td_type(target)

	case *Xhtml_tfoot_type:
		ok = stage.IsStagedXhtml_tfoot_type(target)

	case *Xhtml_th_type:
		ok = stage.IsStagedXhtml_th_type(target)

	case *Xhtml_thead_type:
		ok = stage.IsStagedXhtml_thead_type(target)

	case *Xhtml_tr_type:
		ok = stage.IsStagedXhtml_tr_type(target)

	case *Xhtml_ul_type:
		ok = stage.IsStagedXhtml_ul_type(target)

	case *Xhtml_var_type:
		ok = stage.IsStagedXhtml_var_type(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID) (ok bool) {

	_, ok = stage.ALTERNATIVE_IDs[alternative_id]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_REALs[attribute_value_real]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string]

	return
}

func (stage *StageStruct) IsStagedATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml]

	return
}

func (stage *StageStruct) IsStagedAnyType(anytype *AnyType) (ok bool) {

	_, ok = stage.AnyTypes[anytype]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_REALs[datatype_definition_real]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string]

	return
}

func (stage *StageStruct) IsStagedDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml]

	return
}

func (stage *StageStruct) IsStagedEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE) (ok bool) {

	_, ok = stage.EMBEDDED_VALUEs[embedded_value]

	return
}

func (stage *StageStruct) IsStagedENUM_VALUE(enum_value *ENUM_VALUE) (ok bool) {

	_, ok = stage.ENUM_VALUEs[enum_value]

	return
}

func (stage *StageStruct) IsStagedRELATION_GROUP(relation_group *RELATION_GROUP) (ok bool) {

	_, ok = stage.RELATION_GROUPs[relation_group]

	return
}

func (stage *StageStruct) IsStagedRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) (ok bool) {

	_, ok = stage.RELATION_GROUP_TYPEs[relation_group_type]

	return
}

func (stage *StageStruct) IsStagedREQ_IF(req_if *REQ_IF) (ok bool) {

	_, ok = stage.REQ_IFs[req_if]

	return
}

func (stage *StageStruct) IsStagedREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) (ok bool) {

	_, ok = stage.REQ_IF_CONTENTs[req_if_content]

	return
}

func (stage *StageStruct) IsStagedREQ_IF_HEADER(req_if_header *REQ_IF_HEADER) (ok bool) {

	_, ok = stage.REQ_IF_HEADERs[req_if_header]

	return
}

func (stage *StageStruct) IsStagedREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION) (ok bool) {

	_, ok = stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension]

	return
}

func (stage *StageStruct) IsStagedSPECIFICATION(specification *SPECIFICATION) (ok bool) {

	_, ok = stage.SPECIFICATIONs[specification]

	return
}

func (stage *StageStruct) IsStagedSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) (ok bool) {

	_, ok = stage.SPECIFICATION_TYPEs[specification_type]

	return
}

func (stage *StageStruct) IsStagedSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) (ok bool) {

	_, ok = stage.SPEC_HIERARCHYs[spec_hierarchy]

	return
}

func (stage *StageStruct) IsStagedSPEC_OBJECT(spec_object *SPEC_OBJECT) (ok bool) {

	_, ok = stage.SPEC_OBJECTs[spec_object]

	return
}

func (stage *StageStruct) IsStagedSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) (ok bool) {

	_, ok = stage.SPEC_OBJECT_TYPEs[spec_object_type]

	return
}

func (stage *StageStruct) IsStagedSPEC_RELATION(spec_relation *SPEC_RELATION) (ok bool) {

	_, ok = stage.SPEC_RELATIONs[spec_relation]

	return
}

func (stage *StageStruct) IsStagedSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) (ok bool) {

	_, ok = stage.SPEC_RELATION_TYPEs[spec_relation_type]

	return
}

func (stage *StageStruct) IsStagedXHTML_CONTENT(xhtml_content *XHTML_CONTENT) (ok bool) {

	_, ok = stage.XHTML_CONTENTs[xhtml_content]

	return
}

func (stage *StageStruct) IsStagedXhtml_InlPres_type(xhtml_inlpres_type *Xhtml_InlPres_type) (ok bool) {

	_, ok = stage.Xhtml_InlPres_types[xhtml_inlpres_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_a_type(xhtml_a_type *Xhtml_a_type) (ok bool) {

	_, ok = stage.Xhtml_a_types[xhtml_a_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_abbr_type(xhtml_abbr_type *Xhtml_abbr_type) (ok bool) {

	_, ok = stage.Xhtml_abbr_types[xhtml_abbr_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_acronym_type(xhtml_acronym_type *Xhtml_acronym_type) (ok bool) {

	_, ok = stage.Xhtml_acronym_types[xhtml_acronym_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_address_type(xhtml_address_type *Xhtml_address_type) (ok bool) {

	_, ok = stage.Xhtml_address_types[xhtml_address_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_blockquote_type(xhtml_blockquote_type *Xhtml_blockquote_type) (ok bool) {

	_, ok = stage.Xhtml_blockquote_types[xhtml_blockquote_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_br_type(xhtml_br_type *Xhtml_br_type) (ok bool) {

	_, ok = stage.Xhtml_br_types[xhtml_br_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_caption_type(xhtml_caption_type *Xhtml_caption_type) (ok bool) {

	_, ok = stage.Xhtml_caption_types[xhtml_caption_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_cite_type(xhtml_cite_type *Xhtml_cite_type) (ok bool) {

	_, ok = stage.Xhtml_cite_types[xhtml_cite_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_code_type(xhtml_code_type *Xhtml_code_type) (ok bool) {

	_, ok = stage.Xhtml_code_types[xhtml_code_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_col_type(xhtml_col_type *Xhtml_col_type) (ok bool) {

	_, ok = stage.Xhtml_col_types[xhtml_col_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_colgroup_type(xhtml_colgroup_type *Xhtml_colgroup_type) (ok bool) {

	_, ok = stage.Xhtml_colgroup_types[xhtml_colgroup_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_dd_type(xhtml_dd_type *Xhtml_dd_type) (ok bool) {

	_, ok = stage.Xhtml_dd_types[xhtml_dd_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_dfn_type(xhtml_dfn_type *Xhtml_dfn_type) (ok bool) {

	_, ok = stage.Xhtml_dfn_types[xhtml_dfn_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_div_type(xhtml_div_type *Xhtml_div_type) (ok bool) {

	_, ok = stage.Xhtml_div_types[xhtml_div_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_dl_type(xhtml_dl_type *Xhtml_dl_type) (ok bool) {

	_, ok = stage.Xhtml_dl_types[xhtml_dl_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_dt_type(xhtml_dt_type *Xhtml_dt_type) (ok bool) {

	_, ok = stage.Xhtml_dt_types[xhtml_dt_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_edit_type(xhtml_edit_type *Xhtml_edit_type) (ok bool) {

	_, ok = stage.Xhtml_edit_types[xhtml_edit_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_em_type(xhtml_em_type *Xhtml_em_type) (ok bool) {

	_, ok = stage.Xhtml_em_types[xhtml_em_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_h1_type(xhtml_h1_type *Xhtml_h1_type) (ok bool) {

	_, ok = stage.Xhtml_h1_types[xhtml_h1_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_h2_type(xhtml_h2_type *Xhtml_h2_type) (ok bool) {

	_, ok = stage.Xhtml_h2_types[xhtml_h2_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_h3_type(xhtml_h3_type *Xhtml_h3_type) (ok bool) {

	_, ok = stage.Xhtml_h3_types[xhtml_h3_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_h4_type(xhtml_h4_type *Xhtml_h4_type) (ok bool) {

	_, ok = stage.Xhtml_h4_types[xhtml_h4_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_h5_type(xhtml_h5_type *Xhtml_h5_type) (ok bool) {

	_, ok = stage.Xhtml_h5_types[xhtml_h5_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_h6_type(xhtml_h6_type *Xhtml_h6_type) (ok bool) {

	_, ok = stage.Xhtml_h6_types[xhtml_h6_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_heading_type(xhtml_heading_type *Xhtml_heading_type) (ok bool) {

	_, ok = stage.Xhtml_heading_types[xhtml_heading_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_hr_type(xhtml_hr_type *Xhtml_hr_type) (ok bool) {

	_, ok = stage.Xhtml_hr_types[xhtml_hr_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_kbd_type(xhtml_kbd_type *Xhtml_kbd_type) (ok bool) {

	_, ok = stage.Xhtml_kbd_types[xhtml_kbd_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_li_type(xhtml_li_type *Xhtml_li_type) (ok bool) {

	_, ok = stage.Xhtml_li_types[xhtml_li_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_object_type(xhtml_object_type *Xhtml_object_type) (ok bool) {

	_, ok = stage.Xhtml_object_types[xhtml_object_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_ol_type(xhtml_ol_type *Xhtml_ol_type) (ok bool) {

	_, ok = stage.Xhtml_ol_types[xhtml_ol_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_p_type(xhtml_p_type *Xhtml_p_type) (ok bool) {

	_, ok = stage.Xhtml_p_types[xhtml_p_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_param_type(xhtml_param_type *Xhtml_param_type) (ok bool) {

	_, ok = stage.Xhtml_param_types[xhtml_param_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_pre_type(xhtml_pre_type *Xhtml_pre_type) (ok bool) {

	_, ok = stage.Xhtml_pre_types[xhtml_pre_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_q_type(xhtml_q_type *Xhtml_q_type) (ok bool) {

	_, ok = stage.Xhtml_q_types[xhtml_q_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_samp_type(xhtml_samp_type *Xhtml_samp_type) (ok bool) {

	_, ok = stage.Xhtml_samp_types[xhtml_samp_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_span_type(xhtml_span_type *Xhtml_span_type) (ok bool) {

	_, ok = stage.Xhtml_span_types[xhtml_span_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_strong_type(xhtml_strong_type *Xhtml_strong_type) (ok bool) {

	_, ok = stage.Xhtml_strong_types[xhtml_strong_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_table_type(xhtml_table_type *Xhtml_table_type) (ok bool) {

	_, ok = stage.Xhtml_table_types[xhtml_table_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_tbody_type(xhtml_tbody_type *Xhtml_tbody_type) (ok bool) {

	_, ok = stage.Xhtml_tbody_types[xhtml_tbody_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_td_type(xhtml_td_type *Xhtml_td_type) (ok bool) {

	_, ok = stage.Xhtml_td_types[xhtml_td_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_tfoot_type(xhtml_tfoot_type *Xhtml_tfoot_type) (ok bool) {

	_, ok = stage.Xhtml_tfoot_types[xhtml_tfoot_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_th_type(xhtml_th_type *Xhtml_th_type) (ok bool) {

	_, ok = stage.Xhtml_th_types[xhtml_th_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_thead_type(xhtml_thead_type *Xhtml_thead_type) (ok bool) {

	_, ok = stage.Xhtml_thead_types[xhtml_thead_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_tr_type(xhtml_tr_type *Xhtml_tr_type) (ok bool) {

	_, ok = stage.Xhtml_tr_types[xhtml_tr_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_ul_type(xhtml_ul_type *Xhtml_ul_type) (ok bool) {

	_, ok = stage.Xhtml_ul_types[xhtml_ul_type]

	return
}

func (stage *StageStruct) IsStagedXhtml_var_type(xhtml_var_type *Xhtml_var_type) (ok bool) {

	_, ok = stage.Xhtml_var_types[xhtml_var_type]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *ALTERNATIVE_ID:
		stage.StageBranchALTERNATIVE_ID(target)

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.StageBranchATTRIBUTE_DEFINITION_BOOLEAN(target)

	case *ATTRIBUTE_DEFINITION_DATE:
		stage.StageBranchATTRIBUTE_DEFINITION_DATE(target)

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.StageBranchATTRIBUTE_DEFINITION_ENUMERATION(target)

	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.StageBranchATTRIBUTE_DEFINITION_INTEGER(target)

	case *ATTRIBUTE_DEFINITION_REAL:
		stage.StageBranchATTRIBUTE_DEFINITION_REAL(target)

	case *ATTRIBUTE_DEFINITION_STRING:
		stage.StageBranchATTRIBUTE_DEFINITION_STRING(target)

	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.StageBranchATTRIBUTE_DEFINITION_XHTML(target)

	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.StageBranchATTRIBUTE_VALUE_BOOLEAN(target)

	case *ATTRIBUTE_VALUE_DATE:
		stage.StageBranchATTRIBUTE_VALUE_DATE(target)

	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.StageBranchATTRIBUTE_VALUE_ENUMERATION(target)

	case *ATTRIBUTE_VALUE_INTEGER:
		stage.StageBranchATTRIBUTE_VALUE_INTEGER(target)

	case *ATTRIBUTE_VALUE_REAL:
		stage.StageBranchATTRIBUTE_VALUE_REAL(target)

	case *ATTRIBUTE_VALUE_STRING:
		stage.StageBranchATTRIBUTE_VALUE_STRING(target)

	case *ATTRIBUTE_VALUE_XHTML:
		stage.StageBranchATTRIBUTE_VALUE_XHTML(target)

	case *AnyType:
		stage.StageBranchAnyType(target)

	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.StageBranchDATATYPE_DEFINITION_BOOLEAN(target)

	case *DATATYPE_DEFINITION_DATE:
		stage.StageBranchDATATYPE_DEFINITION_DATE(target)

	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.StageBranchDATATYPE_DEFINITION_ENUMERATION(target)

	case *DATATYPE_DEFINITION_INTEGER:
		stage.StageBranchDATATYPE_DEFINITION_INTEGER(target)

	case *DATATYPE_DEFINITION_REAL:
		stage.StageBranchDATATYPE_DEFINITION_REAL(target)

	case *DATATYPE_DEFINITION_STRING:
		stage.StageBranchDATATYPE_DEFINITION_STRING(target)

	case *DATATYPE_DEFINITION_XHTML:
		stage.StageBranchDATATYPE_DEFINITION_XHTML(target)

	case *EMBEDDED_VALUE:
		stage.StageBranchEMBEDDED_VALUE(target)

	case *ENUM_VALUE:
		stage.StageBranchENUM_VALUE(target)

	case *RELATION_GROUP:
		stage.StageBranchRELATION_GROUP(target)

	case *RELATION_GROUP_TYPE:
		stage.StageBranchRELATION_GROUP_TYPE(target)

	case *REQ_IF:
		stage.StageBranchREQ_IF(target)

	case *REQ_IF_CONTENT:
		stage.StageBranchREQ_IF_CONTENT(target)

	case *REQ_IF_HEADER:
		stage.StageBranchREQ_IF_HEADER(target)

	case *REQ_IF_TOOL_EXTENSION:
		stage.StageBranchREQ_IF_TOOL_EXTENSION(target)

	case *SPECIFICATION:
		stage.StageBranchSPECIFICATION(target)

	case *SPECIFICATION_TYPE:
		stage.StageBranchSPECIFICATION_TYPE(target)

	case *SPEC_HIERARCHY:
		stage.StageBranchSPEC_HIERARCHY(target)

	case *SPEC_OBJECT:
		stage.StageBranchSPEC_OBJECT(target)

	case *SPEC_OBJECT_TYPE:
		stage.StageBranchSPEC_OBJECT_TYPE(target)

	case *SPEC_RELATION:
		stage.StageBranchSPEC_RELATION(target)

	case *SPEC_RELATION_TYPE:
		stage.StageBranchSPEC_RELATION_TYPE(target)

	case *XHTML_CONTENT:
		stage.StageBranchXHTML_CONTENT(target)

	case *Xhtml_InlPres_type:
		stage.StageBranchXhtml_InlPres_type(target)

	case *Xhtml_a_type:
		stage.StageBranchXhtml_a_type(target)

	case *Xhtml_abbr_type:
		stage.StageBranchXhtml_abbr_type(target)

	case *Xhtml_acronym_type:
		stage.StageBranchXhtml_acronym_type(target)

	case *Xhtml_address_type:
		stage.StageBranchXhtml_address_type(target)

	case *Xhtml_blockquote_type:
		stage.StageBranchXhtml_blockquote_type(target)

	case *Xhtml_br_type:
		stage.StageBranchXhtml_br_type(target)

	case *Xhtml_caption_type:
		stage.StageBranchXhtml_caption_type(target)

	case *Xhtml_cite_type:
		stage.StageBranchXhtml_cite_type(target)

	case *Xhtml_code_type:
		stage.StageBranchXhtml_code_type(target)

	case *Xhtml_col_type:
		stage.StageBranchXhtml_col_type(target)

	case *Xhtml_colgroup_type:
		stage.StageBranchXhtml_colgroup_type(target)

	case *Xhtml_dd_type:
		stage.StageBranchXhtml_dd_type(target)

	case *Xhtml_dfn_type:
		stage.StageBranchXhtml_dfn_type(target)

	case *Xhtml_div_type:
		stage.StageBranchXhtml_div_type(target)

	case *Xhtml_dl_type:
		stage.StageBranchXhtml_dl_type(target)

	case *Xhtml_dt_type:
		stage.StageBranchXhtml_dt_type(target)

	case *Xhtml_edit_type:
		stage.StageBranchXhtml_edit_type(target)

	case *Xhtml_em_type:
		stage.StageBranchXhtml_em_type(target)

	case *Xhtml_h1_type:
		stage.StageBranchXhtml_h1_type(target)

	case *Xhtml_h2_type:
		stage.StageBranchXhtml_h2_type(target)

	case *Xhtml_h3_type:
		stage.StageBranchXhtml_h3_type(target)

	case *Xhtml_h4_type:
		stage.StageBranchXhtml_h4_type(target)

	case *Xhtml_h5_type:
		stage.StageBranchXhtml_h5_type(target)

	case *Xhtml_h6_type:
		stage.StageBranchXhtml_h6_type(target)

	case *Xhtml_heading_type:
		stage.StageBranchXhtml_heading_type(target)

	case *Xhtml_hr_type:
		stage.StageBranchXhtml_hr_type(target)

	case *Xhtml_kbd_type:
		stage.StageBranchXhtml_kbd_type(target)

	case *Xhtml_li_type:
		stage.StageBranchXhtml_li_type(target)

	case *Xhtml_object_type:
		stage.StageBranchXhtml_object_type(target)

	case *Xhtml_ol_type:
		stage.StageBranchXhtml_ol_type(target)

	case *Xhtml_p_type:
		stage.StageBranchXhtml_p_type(target)

	case *Xhtml_param_type:
		stage.StageBranchXhtml_param_type(target)

	case *Xhtml_pre_type:
		stage.StageBranchXhtml_pre_type(target)

	case *Xhtml_q_type:
		stage.StageBranchXhtml_q_type(target)

	case *Xhtml_samp_type:
		stage.StageBranchXhtml_samp_type(target)

	case *Xhtml_span_type:
		stage.StageBranchXhtml_span_type(target)

	case *Xhtml_strong_type:
		stage.StageBranchXhtml_strong_type(target)

	case *Xhtml_table_type:
		stage.StageBranchXhtml_table_type(target)

	case *Xhtml_tbody_type:
		stage.StageBranchXhtml_tbody_type(target)

	case *Xhtml_td_type:
		stage.StageBranchXhtml_td_type(target)

	case *Xhtml_tfoot_type:
		stage.StageBranchXhtml_tfoot_type(target)

	case *Xhtml_th_type:
		stage.StageBranchXhtml_th_type(target)

	case *Xhtml_thead_type:
		stage.StageBranchXhtml_thead_type(target)

	case *Xhtml_tr_type:
		stage.StageBranchXhtml_tr_type(target)

	case *Xhtml_ul_type:
		stage.StageBranchXhtml_ul_type(target)

	case *Xhtml_var_type:
		stage.StageBranchXhtml_var_type(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID) {

	// check if instance is already staged
	if IsStaged(stage, alternative_id) {
		return
	}

	alternative_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_boolean) {
		return
	}

	attribute_definition_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_boolean := range attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
		StageBranch(stage, _attribute_value_boolean)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_date) {
		return
	}

	attribute_definition_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_date := range attribute_definition_date.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE {
		StageBranch(stage, _attribute_value_date)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_enumeration) {
		return
	}

	attribute_definition_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_enumeration := range attribute_definition_enumeration.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION {
		StageBranch(stage, _attribute_value_enumeration)
	}
	for _, _alternative_id := range attribute_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_integer) {
		return
	}

	attribute_definition_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_integer := range attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
		StageBranch(stage, _attribute_value_integer)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_real) {
		return
	}

	attribute_definition_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_real := range attribute_definition_real.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL {
		StageBranch(stage, _attribute_value_real)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_string) {
		return
	}

	attribute_definition_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_string := range attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
		StageBranch(stage, _attribute_value_string)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, attribute_definition_xhtml) {
		return
	}

	attribute_definition_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_xhtml := range attribute_definition_xhtml.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML {
		StageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_boolean) {
		return
	}

	attribute_value_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_date) {
		return
	}

	attribute_value_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_enumeration) {
		return
	}

	attribute_value_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_integer) {
		return
	}

	attribute_value_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_real) {
		return
	}

	attribute_value_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_string) {
		return
	}

	attribute_value_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, attribute_value_xhtml) {
		return
	}

	attribute_value_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xhtml_content := range attribute_value_xhtml.THE_VALUE {
		StageBranch(stage, _xhtml_content)
	}
	for _, _xhtml_content := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
		StageBranch(stage, _xhtml_content)
	}

}

func (stage *StageStruct) StageBranchAnyType(anytype *AnyType) {

	// check if instance is already staged
	if IsStaged(stage, anytype) {
		return
	}

	anytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_boolean) {
		return
	}

	datatype_definition_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_date) {
		return
	}

	datatype_definition_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_enumeration) {
		return
	}

	datatype_definition_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _enum_value := range datatype_definition_enumeration.SPECIFIED_VALUES.ENUM_VALUE {
		StageBranch(stage, _enum_value)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_integer) {
		return
	}

	datatype_definition_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_real) {
		return
	}

	datatype_definition_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_string) {
		return
	}

	datatype_definition_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) StageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) {

	// check if instance is already staged
	if IsStaged(stage, datatype_definition_xhtml) {
		return
	}

	datatype_definition_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) StageBranchEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE) {

	// check if instance is already staged
	if IsStaged(stage, embedded_value) {
		return
	}

	embedded_value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchENUM_VALUE(enum_value *ENUM_VALUE) {

	// check if instance is already staged
	if IsStaged(stage, enum_value) {
		return
	}

	enum_value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range enum_value.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _embedded_value := range enum_value.PROPERTIES.EMBEDDED_VALUE {
		StageBranch(stage, _embedded_value)
	}

}

func (stage *StageStruct) StageBranchRELATION_GROUP(relation_group *RELATION_GROUP) {

	// check if instance is already staged
	if IsStaged(stage, relation_group) {
		return
	}

	relation_group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) StageBranchRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, relation_group_type) {
		return
	}

	relation_group_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range relation_group_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_definition_boolean := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		StageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		StageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		StageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		StageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		StageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		StageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		StageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) StageBranchREQ_IF(req_if *REQ_IF) {

	// check if instance is already staged
	if IsStaged(stage, req_if) {
		return
	}

	req_if.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if.THE_HEADER.REQ_IF_HEADER != nil {
		StageBranch(stage, req_if.THE_HEADER.REQ_IF_HEADER)
	}
	if req_if.CORE_CONTENT.REQ_IF_CONTENT != nil {
		StageBranch(stage, req_if.CORE_CONTENT.REQ_IF_CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_tool_extension := range req_if.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION {
		StageBranch(stage, _req_if_tool_extension)
	}

}

func (stage *StageStruct) StageBranchREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) {

	// check if instance is already staged
	if IsStaged(stage, req_if_content) {
		return
	}

	req_if_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatype_definition_boolean := range req_if_content.DATATYPES.DATATYPE_DEFINITION_BOOLEAN {
		StageBranch(stage, _datatype_definition_boolean)
	}
	for _, _datatype_definition_date := range req_if_content.DATATYPES.DATATYPE_DEFINITION_DATE {
		StageBranch(stage, _datatype_definition_date)
	}
	for _, _datatype_definition_enumeration := range req_if_content.DATATYPES.DATATYPE_DEFINITION_ENUMERATION {
		StageBranch(stage, _datatype_definition_enumeration)
	}
	for _, _datatype_definition_integer := range req_if_content.DATATYPES.DATATYPE_DEFINITION_INTEGER {
		StageBranch(stage, _datatype_definition_integer)
	}
	for _, _datatype_definition_real := range req_if_content.DATATYPES.DATATYPE_DEFINITION_REAL {
		StageBranch(stage, _datatype_definition_real)
	}
	for _, _datatype_definition_string := range req_if_content.DATATYPES.DATATYPE_DEFINITION_STRING {
		StageBranch(stage, _datatype_definition_string)
	}
	for _, _datatype_definition_xhtml := range req_if_content.DATATYPES.DATATYPE_DEFINITION_XHTML {
		StageBranch(stage, _datatype_definition_xhtml)
	}
	for _, _relation_group_type := range req_if_content.SPEC_TYPES.RELATION_GROUP_TYPE {
		StageBranch(stage, _relation_group_type)
	}
	for _, _spec_object_type := range req_if_content.SPEC_TYPES.SPEC_OBJECT_TYPE {
		StageBranch(stage, _spec_object_type)
	}
	for _, _spec_relation_type := range req_if_content.SPEC_TYPES.SPEC_RELATION_TYPE {
		StageBranch(stage, _spec_relation_type)
	}
	for _, _specification_type := range req_if_content.SPEC_TYPES.SPECIFICATION_TYPE {
		StageBranch(stage, _specification_type)
	}
	for _, _spec_object := range req_if_content.SPEC_OBJECTS.SPEC_OBJECT {
		StageBranch(stage, _spec_object)
	}
	for _, _spec_relation := range req_if_content.SPEC_RELATIONS.SPEC_RELATION {
		StageBranch(stage, _spec_relation)
	}
	for _, _specification := range req_if_content.SPECIFICATIONS.SPECIFICATION {
		StageBranch(stage, _specification)
	}
	for _, _relation_group := range req_if_content.SPEC_RELATION_GROUPS.RELATION_GROUP {
		StageBranch(stage, _relation_group)
	}

}

func (stage *StageStruct) StageBranchREQ_IF_HEADER(req_if_header *REQ_IF_HEADER) {

	// check if instance is already staged
	if IsStaged(stage, req_if_header) {
		return
	}

	req_if_header.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION) {

	// check if instance is already staged
	if IsStaged(stage, req_if_tool_extension) {
		return
	}

	req_if_tool_extension.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchSPECIFICATION(specification *SPECIFICATION) {

	// check if instance is already staged
	if IsStaged(stage, specification) {
		return
	}

	specification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range specification.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_boolean := range specification.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		StageBranch(stage, _attribute_value_boolean)
	}
	for _, _attribute_value_date := range specification.VALUES.ATTRIBUTE_VALUE_DATE {
		StageBranch(stage, _attribute_value_date)
	}
	for _, _attribute_value_enumeration := range specification.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		StageBranch(stage, _attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range specification.VALUES.ATTRIBUTE_VALUE_INTEGER {
		StageBranch(stage, _attribute_value_integer)
	}
	for _, _attribute_value_real := range specification.VALUES.ATTRIBUTE_VALUE_REAL {
		StageBranch(stage, _attribute_value_real)
	}
	for _, _attribute_value_string := range specification.VALUES.ATTRIBUTE_VALUE_STRING {
		StageBranch(stage, _attribute_value_string)
	}
	for _, _attribute_value_xhtml := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
		StageBranch(stage, _attribute_value_xhtml)
	}
	for _, _spec_hierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
		StageBranch(stage, _spec_hierarchy)
	}

}

func (stage *StageStruct) StageBranchSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, specification_type) {
		return
	}

	specification_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range specification_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_definition_boolean := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		StageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		StageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		StageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		StageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		StageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		StageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		StageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) StageBranchSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) {

	// check if instance is already staged
	if IsStaged(stage, spec_hierarchy) {
		return
	}

	spec_hierarchy.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_hierarchy.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _spec_hierarchy := range spec_hierarchy.CHILDREN.SPEC_HIERARCHY {
		StageBranch(stage, _spec_hierarchy)
	}

}

func (stage *StageStruct) StageBranchSPEC_OBJECT(spec_object *SPEC_OBJECT) {

	// check if instance is already staged
	if IsStaged(stage, spec_object) {
		return
	}

	spec_object.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_object.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_boolean := range spec_object.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		StageBranch(stage, _attribute_value_boolean)
	}
	for _, _attribute_value_date := range spec_object.VALUES.ATTRIBUTE_VALUE_DATE {
		StageBranch(stage, _attribute_value_date)
	}
	for _, _attribute_value_enumeration := range spec_object.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		StageBranch(stage, _attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range spec_object.VALUES.ATTRIBUTE_VALUE_INTEGER {
		StageBranch(stage, _attribute_value_integer)
	}
	for _, _attribute_value_real := range spec_object.VALUES.ATTRIBUTE_VALUE_REAL {
		StageBranch(stage, _attribute_value_real)
	}
	for _, _attribute_value_string := range spec_object.VALUES.ATTRIBUTE_VALUE_STRING {
		StageBranch(stage, _attribute_value_string)
	}
	for _, _attribute_value_xhtml := range spec_object.VALUES.ATTRIBUTE_VALUE_XHTML {
		StageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) StageBranchSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, spec_object_type) {
		return
	}

	spec_object_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_object_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_definition_boolean := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		StageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		StageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		StageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		StageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		StageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		StageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		StageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) StageBranchSPEC_RELATION(spec_relation *SPEC_RELATION) {

	// check if instance is already staged
	if IsStaged(stage, spec_relation) {
		return
	}

	spec_relation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_relation.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_boolean := range spec_relation.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		StageBranch(stage, _attribute_value_boolean)
	}
	for _, _attribute_value_date := range spec_relation.VALUES.ATTRIBUTE_VALUE_DATE {
		StageBranch(stage, _attribute_value_date)
	}
	for _, _attribute_value_enumeration := range spec_relation.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		StageBranch(stage, _attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range spec_relation.VALUES.ATTRIBUTE_VALUE_INTEGER {
		StageBranch(stage, _attribute_value_integer)
	}
	for _, _attribute_value_real := range spec_relation.VALUES.ATTRIBUTE_VALUE_REAL {
		StageBranch(stage, _attribute_value_real)
	}
	for _, _attribute_value_string := range spec_relation.VALUES.ATTRIBUTE_VALUE_STRING {
		StageBranch(stage, _attribute_value_string)
	}
	for _, _attribute_value_xhtml := range spec_relation.VALUES.ATTRIBUTE_VALUE_XHTML {
		StageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) StageBranchSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) {

	// check if instance is already staged
	if IsStaged(stage, spec_relation_type) {
		return
	}

	spec_relation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_relation_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
		StageBranch(stage, _alternative_id)
	}
	for _, _attribute_definition_boolean := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		StageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		StageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		StageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		StageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		StageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		StageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		StageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) StageBranchXHTML_CONTENT(xhtml_content *XHTML_CONTENT) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_content) {
		return
	}

	xhtml_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_InlPres_type(xhtml_inlpres_type *Xhtml_InlPres_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_inlpres_type) {
		return
	}

	xhtml_inlpres_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_a_type(xhtml_a_type *Xhtml_a_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_a_type) {
		return
	}

	xhtml_a_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_abbr_type(xhtml_abbr_type *Xhtml_abbr_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_abbr_type) {
		return
	}

	xhtml_abbr_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_acronym_type(xhtml_acronym_type *Xhtml_acronym_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_acronym_type) {
		return
	}

	xhtml_acronym_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_address_type(xhtml_address_type *Xhtml_address_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_address_type) {
		return
	}

	xhtml_address_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_blockquote_type(xhtml_blockquote_type *Xhtml_blockquote_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_blockquote_type) {
		return
	}

	xhtml_blockquote_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_br_type(xhtml_br_type *Xhtml_br_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_br_type) {
		return
	}

	xhtml_br_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_caption_type(xhtml_caption_type *Xhtml_caption_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_caption_type) {
		return
	}

	xhtml_caption_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_cite_type(xhtml_cite_type *Xhtml_cite_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_cite_type) {
		return
	}

	xhtml_cite_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_code_type(xhtml_code_type *Xhtml_code_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_code_type) {
		return
	}

	xhtml_code_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_col_type(xhtml_col_type *Xhtml_col_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_col_type) {
		return
	}

	xhtml_col_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_colgroup_type(xhtml_colgroup_type *Xhtml_colgroup_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_colgroup_type) {
		return
	}

	xhtml_colgroup_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_dd_type(xhtml_dd_type *Xhtml_dd_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_dd_type) {
		return
	}

	xhtml_dd_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_dfn_type(xhtml_dfn_type *Xhtml_dfn_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_dfn_type) {
		return
	}

	xhtml_dfn_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_div_type(xhtml_div_type *Xhtml_div_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_div_type) {
		return
	}

	xhtml_div_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_dl_type(xhtml_dl_type *Xhtml_dl_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_dl_type) {
		return
	}

	xhtml_dl_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_dt_type(xhtml_dt_type *Xhtml_dt_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_dt_type) {
		return
	}

	xhtml_dt_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_edit_type(xhtml_edit_type *Xhtml_edit_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_edit_type) {
		return
	}

	xhtml_edit_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_em_type(xhtml_em_type *Xhtml_em_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_em_type) {
		return
	}

	xhtml_em_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_h1_type(xhtml_h1_type *Xhtml_h1_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_h1_type) {
		return
	}

	xhtml_h1_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_h2_type(xhtml_h2_type *Xhtml_h2_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_h2_type) {
		return
	}

	xhtml_h2_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_h3_type(xhtml_h3_type *Xhtml_h3_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_h3_type) {
		return
	}

	xhtml_h3_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_h4_type(xhtml_h4_type *Xhtml_h4_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_h4_type) {
		return
	}

	xhtml_h4_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_h5_type(xhtml_h5_type *Xhtml_h5_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_h5_type) {
		return
	}

	xhtml_h5_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_h6_type(xhtml_h6_type *Xhtml_h6_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_h6_type) {
		return
	}

	xhtml_h6_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_heading_type(xhtml_heading_type *Xhtml_heading_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_heading_type) {
		return
	}

	xhtml_heading_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_hr_type(xhtml_hr_type *Xhtml_hr_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_hr_type) {
		return
	}

	xhtml_hr_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_kbd_type(xhtml_kbd_type *Xhtml_kbd_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_kbd_type) {
		return
	}

	xhtml_kbd_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_li_type(xhtml_li_type *Xhtml_li_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_li_type) {
		return
	}

	xhtml_li_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_object_type(xhtml_object_type *Xhtml_object_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_object_type) {
		return
	}

	xhtml_object_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_ol_type(xhtml_ol_type *Xhtml_ol_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_ol_type) {
		return
	}

	xhtml_ol_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_p_type(xhtml_p_type *Xhtml_p_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_p_type) {
		return
	}

	xhtml_p_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_param_type(xhtml_param_type *Xhtml_param_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_param_type) {
		return
	}

	xhtml_param_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_pre_type(xhtml_pre_type *Xhtml_pre_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_pre_type) {
		return
	}

	xhtml_pre_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_q_type(xhtml_q_type *Xhtml_q_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_q_type) {
		return
	}

	xhtml_q_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_samp_type(xhtml_samp_type *Xhtml_samp_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_samp_type) {
		return
	}

	xhtml_samp_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_span_type(xhtml_span_type *Xhtml_span_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_span_type) {
		return
	}

	xhtml_span_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_strong_type(xhtml_strong_type *Xhtml_strong_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_strong_type) {
		return
	}

	xhtml_strong_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_table_type(xhtml_table_type *Xhtml_table_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_table_type) {
		return
	}

	xhtml_table_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_tbody_type(xhtml_tbody_type *Xhtml_tbody_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_tbody_type) {
		return
	}

	xhtml_tbody_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_td_type(xhtml_td_type *Xhtml_td_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_td_type) {
		return
	}

	xhtml_td_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_tfoot_type(xhtml_tfoot_type *Xhtml_tfoot_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_tfoot_type) {
		return
	}

	xhtml_tfoot_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_th_type(xhtml_th_type *Xhtml_th_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_th_type) {
		return
	}

	xhtml_th_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_thead_type(xhtml_thead_type *Xhtml_thead_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_thead_type) {
		return
	}

	xhtml_thead_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_tr_type(xhtml_tr_type *Xhtml_tr_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_tr_type) {
		return
	}

	xhtml_tr_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_ul_type(xhtml_ul_type *Xhtml_ul_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_ul_type) {
		return
	}

	xhtml_ul_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchXhtml_var_type(xhtml_var_type *Xhtml_var_type) {

	// check if instance is already staged
	if IsStaged(stage, xhtml_var_type) {
		return
	}

	xhtml_var_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *ALTERNATIVE_ID:
		toT := CopyBranchALTERNATIVE_ID(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		toT := CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_DATE:
		toT := CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		toT := CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_INTEGER:
		toT := CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_REAL:
		toT := CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_STRING:
		toT := CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_XHTML:
		toT := CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_BOOLEAN:
		toT := CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_DATE:
		toT := CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_ENUMERATION:
		toT := CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_INTEGER:
		toT := CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_REAL:
		toT := CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_STRING:
		toT := CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_XHTML:
		toT := CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AnyType:
		toT := CopyBranchAnyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_BOOLEAN:
		toT := CopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_DATE:
		toT := CopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_ENUMERATION:
		toT := CopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_INTEGER:
		toT := CopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_REAL:
		toT := CopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_STRING:
		toT := CopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_XHTML:
		toT := CopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EMBEDDED_VALUE:
		toT := CopyBranchEMBEDDED_VALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ENUM_VALUE:
		toT := CopyBranchENUM_VALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RELATION_GROUP:
		toT := CopyBranchRELATION_GROUP(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RELATION_GROUP_TYPE:
		toT := CopyBranchRELATION_GROUP_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF:
		toT := CopyBranchREQ_IF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF_CONTENT:
		toT := CopyBranchREQ_IF_CONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF_HEADER:
		toT := CopyBranchREQ_IF_HEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF_TOOL_EXTENSION:
		toT := CopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFICATION:
		toT := CopyBranchSPECIFICATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFICATION_TYPE:
		toT := CopyBranchSPECIFICATION_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_HIERARCHY:
		toT := CopyBranchSPEC_HIERARCHY(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_OBJECT:
		toT := CopyBranchSPEC_OBJECT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_OBJECT_TYPE:
		toT := CopyBranchSPEC_OBJECT_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_RELATION:
		toT := CopyBranchSPEC_RELATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_RELATION_TYPE:
		toT := CopyBranchSPEC_RELATION_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XHTML_CONTENT:
		toT := CopyBranchXHTML_CONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_InlPres_type:
		toT := CopyBranchXhtml_InlPres_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_a_type:
		toT := CopyBranchXhtml_a_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_abbr_type:
		toT := CopyBranchXhtml_abbr_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_acronym_type:
		toT := CopyBranchXhtml_acronym_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_address_type:
		toT := CopyBranchXhtml_address_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_blockquote_type:
		toT := CopyBranchXhtml_blockquote_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_br_type:
		toT := CopyBranchXhtml_br_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_caption_type:
		toT := CopyBranchXhtml_caption_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_cite_type:
		toT := CopyBranchXhtml_cite_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_code_type:
		toT := CopyBranchXhtml_code_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_col_type:
		toT := CopyBranchXhtml_col_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_colgroup_type:
		toT := CopyBranchXhtml_colgroup_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_dd_type:
		toT := CopyBranchXhtml_dd_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_dfn_type:
		toT := CopyBranchXhtml_dfn_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_div_type:
		toT := CopyBranchXhtml_div_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_dl_type:
		toT := CopyBranchXhtml_dl_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_dt_type:
		toT := CopyBranchXhtml_dt_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_edit_type:
		toT := CopyBranchXhtml_edit_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_em_type:
		toT := CopyBranchXhtml_em_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_h1_type:
		toT := CopyBranchXhtml_h1_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_h2_type:
		toT := CopyBranchXhtml_h2_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_h3_type:
		toT := CopyBranchXhtml_h3_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_h4_type:
		toT := CopyBranchXhtml_h4_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_h5_type:
		toT := CopyBranchXhtml_h5_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_h6_type:
		toT := CopyBranchXhtml_h6_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_heading_type:
		toT := CopyBranchXhtml_heading_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_hr_type:
		toT := CopyBranchXhtml_hr_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_kbd_type:
		toT := CopyBranchXhtml_kbd_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_li_type:
		toT := CopyBranchXhtml_li_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_object_type:
		toT := CopyBranchXhtml_object_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_ol_type:
		toT := CopyBranchXhtml_ol_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_p_type:
		toT := CopyBranchXhtml_p_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_param_type:
		toT := CopyBranchXhtml_param_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_pre_type:
		toT := CopyBranchXhtml_pre_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_q_type:
		toT := CopyBranchXhtml_q_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_samp_type:
		toT := CopyBranchXhtml_samp_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_span_type:
		toT := CopyBranchXhtml_span_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_strong_type:
		toT := CopyBranchXhtml_strong_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_table_type:
		toT := CopyBranchXhtml_table_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_tbody_type:
		toT := CopyBranchXhtml_tbody_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_td_type:
		toT := CopyBranchXhtml_td_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_tfoot_type:
		toT := CopyBranchXhtml_tfoot_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_th_type:
		toT := CopyBranchXhtml_th_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_thead_type:
		toT := CopyBranchXhtml_thead_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_tr_type:
		toT := CopyBranchXhtml_tr_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_ul_type:
		toT := CopyBranchXhtml_ul_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Xhtml_var_type:
		toT := CopyBranchXhtml_var_type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchALTERNATIVE_ID(mapOrigCopy map[any]any, alternative_idFrom *ALTERNATIVE_ID) (alternative_idTo *ALTERNATIVE_ID) {

	// alternative_idFrom has already been copied
	if _alternative_idTo, ok := mapOrigCopy[alternative_idFrom]; ok {
		alternative_idTo = _alternative_idTo.(*ALTERNATIVE_ID)
		return
	}

	alternative_idTo = new(ALTERNATIVE_ID)
	mapOrigCopy[alternative_idFrom] = alternative_idTo
	alternative_idFrom.CopyBasicFields(alternative_idTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy map[any]any, attribute_definition_booleanFrom *ATTRIBUTE_DEFINITION_BOOLEAN) (attribute_definition_booleanTo *ATTRIBUTE_DEFINITION_BOOLEAN) {

	// attribute_definition_booleanFrom has already been copied
	if _attribute_definition_booleanTo, ok := mapOrigCopy[attribute_definition_booleanFrom]; ok {
		attribute_definition_booleanTo = _attribute_definition_booleanTo.(*ATTRIBUTE_DEFINITION_BOOLEAN)
		return
	}

	attribute_definition_booleanTo = new(ATTRIBUTE_DEFINITION_BOOLEAN)
	mapOrigCopy[attribute_definition_booleanFrom] = attribute_definition_booleanTo
	attribute_definition_booleanFrom.CopyBasicFields(attribute_definition_booleanTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_booleanFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_booleanTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_booleanTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_value_boolean := range attribute_definition_booleanFrom.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
		attribute_definition_booleanTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN = append(attribute_definition_booleanTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN, CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, _attribute_value_boolean))
	}

	return
}

func CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy map[any]any, attribute_definition_dateFrom *ATTRIBUTE_DEFINITION_DATE) (attribute_definition_dateTo *ATTRIBUTE_DEFINITION_DATE) {

	// attribute_definition_dateFrom has already been copied
	if _attribute_definition_dateTo, ok := mapOrigCopy[attribute_definition_dateFrom]; ok {
		attribute_definition_dateTo = _attribute_definition_dateTo.(*ATTRIBUTE_DEFINITION_DATE)
		return
	}

	attribute_definition_dateTo = new(ATTRIBUTE_DEFINITION_DATE)
	mapOrigCopy[attribute_definition_dateFrom] = attribute_definition_dateTo
	attribute_definition_dateFrom.CopyBasicFields(attribute_definition_dateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_dateFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_dateTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_dateTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_value_date := range attribute_definition_dateFrom.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE {
		attribute_definition_dateTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE = append(attribute_definition_dateTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE, CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, _attribute_value_date))
	}

	return
}

func CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy map[any]any, attribute_definition_enumerationFrom *ATTRIBUTE_DEFINITION_ENUMERATION) (attribute_definition_enumerationTo *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// attribute_definition_enumerationFrom has already been copied
	if _attribute_definition_enumerationTo, ok := mapOrigCopy[attribute_definition_enumerationFrom]; ok {
		attribute_definition_enumerationTo = _attribute_definition_enumerationTo.(*ATTRIBUTE_DEFINITION_ENUMERATION)
		return
	}

	attribute_definition_enumerationTo = new(ATTRIBUTE_DEFINITION_ENUMERATION)
	mapOrigCopy[attribute_definition_enumerationFrom] = attribute_definition_enumerationTo
	attribute_definition_enumerationFrom.CopyBasicFields(attribute_definition_enumerationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_enumeration := range attribute_definition_enumerationFrom.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION {
		attribute_definition_enumerationTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION = append(attribute_definition_enumerationTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION, CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, _attribute_value_enumeration))
	}
	for _, _alternative_id := range attribute_definition_enumerationFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_enumerationTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_enumerationTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}

	return
}

func CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy map[any]any, attribute_definition_integerFrom *ATTRIBUTE_DEFINITION_INTEGER) (attribute_definition_integerTo *ATTRIBUTE_DEFINITION_INTEGER) {

	// attribute_definition_integerFrom has already been copied
	if _attribute_definition_integerTo, ok := mapOrigCopy[attribute_definition_integerFrom]; ok {
		attribute_definition_integerTo = _attribute_definition_integerTo.(*ATTRIBUTE_DEFINITION_INTEGER)
		return
	}

	attribute_definition_integerTo = new(ATTRIBUTE_DEFINITION_INTEGER)
	mapOrigCopy[attribute_definition_integerFrom] = attribute_definition_integerTo
	attribute_definition_integerFrom.CopyBasicFields(attribute_definition_integerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_integerFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_integerTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_integerTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_value_integer := range attribute_definition_integerFrom.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
		attribute_definition_integerTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER = append(attribute_definition_integerTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER, CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, _attribute_value_integer))
	}

	return
}

func CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy map[any]any, attribute_definition_realFrom *ATTRIBUTE_DEFINITION_REAL) (attribute_definition_realTo *ATTRIBUTE_DEFINITION_REAL) {

	// attribute_definition_realFrom has already been copied
	if _attribute_definition_realTo, ok := mapOrigCopy[attribute_definition_realFrom]; ok {
		attribute_definition_realTo = _attribute_definition_realTo.(*ATTRIBUTE_DEFINITION_REAL)
		return
	}

	attribute_definition_realTo = new(ATTRIBUTE_DEFINITION_REAL)
	mapOrigCopy[attribute_definition_realFrom] = attribute_definition_realTo
	attribute_definition_realFrom.CopyBasicFields(attribute_definition_realTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_realFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_realTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_realTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_value_real := range attribute_definition_realFrom.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL {
		attribute_definition_realTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL = append(attribute_definition_realTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL, CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, _attribute_value_real))
	}

	return
}

func CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy map[any]any, attribute_definition_stringFrom *ATTRIBUTE_DEFINITION_STRING) (attribute_definition_stringTo *ATTRIBUTE_DEFINITION_STRING) {

	// attribute_definition_stringFrom has already been copied
	if _attribute_definition_stringTo, ok := mapOrigCopy[attribute_definition_stringFrom]; ok {
		attribute_definition_stringTo = _attribute_definition_stringTo.(*ATTRIBUTE_DEFINITION_STRING)
		return
	}

	attribute_definition_stringTo = new(ATTRIBUTE_DEFINITION_STRING)
	mapOrigCopy[attribute_definition_stringFrom] = attribute_definition_stringTo
	attribute_definition_stringFrom.CopyBasicFields(attribute_definition_stringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_stringFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_stringTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_stringTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_value_string := range attribute_definition_stringFrom.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
		attribute_definition_stringTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING = append(attribute_definition_stringTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING, CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, _attribute_value_string))
	}

	return
}

func CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy map[any]any, attribute_definition_xhtmlFrom *ATTRIBUTE_DEFINITION_XHTML) (attribute_definition_xhtmlTo *ATTRIBUTE_DEFINITION_XHTML) {

	// attribute_definition_xhtmlFrom has already been copied
	if _attribute_definition_xhtmlTo, ok := mapOrigCopy[attribute_definition_xhtmlFrom]; ok {
		attribute_definition_xhtmlTo = _attribute_definition_xhtmlTo.(*ATTRIBUTE_DEFINITION_XHTML)
		return
	}

	attribute_definition_xhtmlTo = new(ATTRIBUTE_DEFINITION_XHTML)
	mapOrigCopy[attribute_definition_xhtmlFrom] = attribute_definition_xhtmlTo
	attribute_definition_xhtmlFrom.CopyBasicFields(attribute_definition_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_xhtmlFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		attribute_definition_xhtmlTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(attribute_definition_xhtmlTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_value_xhtml := range attribute_definition_xhtmlFrom.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML {
		attribute_definition_xhtmlTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML = append(attribute_definition_xhtmlTo.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML, CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, _attribute_value_xhtml))
	}

	return
}

func CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy map[any]any, attribute_value_booleanFrom *ATTRIBUTE_VALUE_BOOLEAN) (attribute_value_booleanTo *ATTRIBUTE_VALUE_BOOLEAN) {

	// attribute_value_booleanFrom has already been copied
	if _attribute_value_booleanTo, ok := mapOrigCopy[attribute_value_booleanFrom]; ok {
		attribute_value_booleanTo = _attribute_value_booleanTo.(*ATTRIBUTE_VALUE_BOOLEAN)
		return
	}

	attribute_value_booleanTo = new(ATTRIBUTE_VALUE_BOOLEAN)
	mapOrigCopy[attribute_value_booleanFrom] = attribute_value_booleanTo
	attribute_value_booleanFrom.CopyBasicFields(attribute_value_booleanTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy map[any]any, attribute_value_dateFrom *ATTRIBUTE_VALUE_DATE) (attribute_value_dateTo *ATTRIBUTE_VALUE_DATE) {

	// attribute_value_dateFrom has already been copied
	if _attribute_value_dateTo, ok := mapOrigCopy[attribute_value_dateFrom]; ok {
		attribute_value_dateTo = _attribute_value_dateTo.(*ATTRIBUTE_VALUE_DATE)
		return
	}

	attribute_value_dateTo = new(ATTRIBUTE_VALUE_DATE)
	mapOrigCopy[attribute_value_dateFrom] = attribute_value_dateTo
	attribute_value_dateFrom.CopyBasicFields(attribute_value_dateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy map[any]any, attribute_value_enumerationFrom *ATTRIBUTE_VALUE_ENUMERATION) (attribute_value_enumerationTo *ATTRIBUTE_VALUE_ENUMERATION) {

	// attribute_value_enumerationFrom has already been copied
	if _attribute_value_enumerationTo, ok := mapOrigCopy[attribute_value_enumerationFrom]; ok {
		attribute_value_enumerationTo = _attribute_value_enumerationTo.(*ATTRIBUTE_VALUE_ENUMERATION)
		return
	}

	attribute_value_enumerationTo = new(ATTRIBUTE_VALUE_ENUMERATION)
	mapOrigCopy[attribute_value_enumerationFrom] = attribute_value_enumerationTo
	attribute_value_enumerationFrom.CopyBasicFields(attribute_value_enumerationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy map[any]any, attribute_value_integerFrom *ATTRIBUTE_VALUE_INTEGER) (attribute_value_integerTo *ATTRIBUTE_VALUE_INTEGER) {

	// attribute_value_integerFrom has already been copied
	if _attribute_value_integerTo, ok := mapOrigCopy[attribute_value_integerFrom]; ok {
		attribute_value_integerTo = _attribute_value_integerTo.(*ATTRIBUTE_VALUE_INTEGER)
		return
	}

	attribute_value_integerTo = new(ATTRIBUTE_VALUE_INTEGER)
	mapOrigCopy[attribute_value_integerFrom] = attribute_value_integerTo
	attribute_value_integerFrom.CopyBasicFields(attribute_value_integerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy map[any]any, attribute_value_realFrom *ATTRIBUTE_VALUE_REAL) (attribute_value_realTo *ATTRIBUTE_VALUE_REAL) {

	// attribute_value_realFrom has already been copied
	if _attribute_value_realTo, ok := mapOrigCopy[attribute_value_realFrom]; ok {
		attribute_value_realTo = _attribute_value_realTo.(*ATTRIBUTE_VALUE_REAL)
		return
	}

	attribute_value_realTo = new(ATTRIBUTE_VALUE_REAL)
	mapOrigCopy[attribute_value_realFrom] = attribute_value_realTo
	attribute_value_realFrom.CopyBasicFields(attribute_value_realTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy map[any]any, attribute_value_stringFrom *ATTRIBUTE_VALUE_STRING) (attribute_value_stringTo *ATTRIBUTE_VALUE_STRING) {

	// attribute_value_stringFrom has already been copied
	if _attribute_value_stringTo, ok := mapOrigCopy[attribute_value_stringFrom]; ok {
		attribute_value_stringTo = _attribute_value_stringTo.(*ATTRIBUTE_VALUE_STRING)
		return
	}

	attribute_value_stringTo = new(ATTRIBUTE_VALUE_STRING)
	mapOrigCopy[attribute_value_stringFrom] = attribute_value_stringTo
	attribute_value_stringFrom.CopyBasicFields(attribute_value_stringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy map[any]any, attribute_value_xhtmlFrom *ATTRIBUTE_VALUE_XHTML) (attribute_value_xhtmlTo *ATTRIBUTE_VALUE_XHTML) {

	// attribute_value_xhtmlFrom has already been copied
	if _attribute_value_xhtmlTo, ok := mapOrigCopy[attribute_value_xhtmlFrom]; ok {
		attribute_value_xhtmlTo = _attribute_value_xhtmlTo.(*ATTRIBUTE_VALUE_XHTML)
		return
	}

	attribute_value_xhtmlTo = new(ATTRIBUTE_VALUE_XHTML)
	mapOrigCopy[attribute_value_xhtmlFrom] = attribute_value_xhtmlTo
	attribute_value_xhtmlFrom.CopyBasicFields(attribute_value_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xhtml_content := range attribute_value_xhtmlFrom.THE_VALUE {
		attribute_value_xhtmlTo.THE_VALUE = append(attribute_value_xhtmlTo.THE_VALUE, CopyBranchXHTML_CONTENT(mapOrigCopy, _xhtml_content))
	}
	for _, _xhtml_content := range attribute_value_xhtmlFrom.THE_ORIGINAL_VALUE {
		attribute_value_xhtmlTo.THE_ORIGINAL_VALUE = append(attribute_value_xhtmlTo.THE_ORIGINAL_VALUE, CopyBranchXHTML_CONTENT(mapOrigCopy, _xhtml_content))
	}

	return
}

func CopyBranchAnyType(mapOrigCopy map[any]any, anytypeFrom *AnyType) (anytypeTo *AnyType) {

	// anytypeFrom has already been copied
	if _anytypeTo, ok := mapOrigCopy[anytypeFrom]; ok {
		anytypeTo = _anytypeTo.(*AnyType)
		return
	}

	anytypeTo = new(AnyType)
	mapOrigCopy[anytypeFrom] = anytypeTo
	anytypeFrom.CopyBasicFields(anytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy map[any]any, datatype_definition_booleanFrom *DATATYPE_DEFINITION_BOOLEAN) (datatype_definition_booleanTo *DATATYPE_DEFINITION_BOOLEAN) {

	// datatype_definition_booleanFrom has already been copied
	if _datatype_definition_booleanTo, ok := mapOrigCopy[datatype_definition_booleanFrom]; ok {
		datatype_definition_booleanTo = _datatype_definition_booleanTo.(*DATATYPE_DEFINITION_BOOLEAN)
		return
	}

	datatype_definition_booleanTo = new(DATATYPE_DEFINITION_BOOLEAN)
	mapOrigCopy[datatype_definition_booleanFrom] = datatype_definition_booleanTo
	datatype_definition_booleanFrom.CopyBasicFields(datatype_definition_booleanTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_booleanFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		datatype_definition_booleanTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(datatype_definition_booleanTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}

	return
}

func CopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy map[any]any, datatype_definition_dateFrom *DATATYPE_DEFINITION_DATE) (datatype_definition_dateTo *DATATYPE_DEFINITION_DATE) {

	// datatype_definition_dateFrom has already been copied
	if _datatype_definition_dateTo, ok := mapOrigCopy[datatype_definition_dateFrom]; ok {
		datatype_definition_dateTo = _datatype_definition_dateTo.(*DATATYPE_DEFINITION_DATE)
		return
	}

	datatype_definition_dateTo = new(DATATYPE_DEFINITION_DATE)
	mapOrigCopy[datatype_definition_dateFrom] = datatype_definition_dateTo
	datatype_definition_dateFrom.CopyBasicFields(datatype_definition_dateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_dateFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		datatype_definition_dateTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(datatype_definition_dateTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}

	return
}

func CopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy map[any]any, datatype_definition_enumerationFrom *DATATYPE_DEFINITION_ENUMERATION) (datatype_definition_enumerationTo *DATATYPE_DEFINITION_ENUMERATION) {

	// datatype_definition_enumerationFrom has already been copied
	if _datatype_definition_enumerationTo, ok := mapOrigCopy[datatype_definition_enumerationFrom]; ok {
		datatype_definition_enumerationTo = _datatype_definition_enumerationTo.(*DATATYPE_DEFINITION_ENUMERATION)
		return
	}

	datatype_definition_enumerationTo = new(DATATYPE_DEFINITION_ENUMERATION)
	mapOrigCopy[datatype_definition_enumerationFrom] = datatype_definition_enumerationTo
	datatype_definition_enumerationFrom.CopyBasicFields(datatype_definition_enumerationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_enumerationFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		datatype_definition_enumerationTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(datatype_definition_enumerationTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _enum_value := range datatype_definition_enumerationFrom.SPECIFIED_VALUES.ENUM_VALUE {
		datatype_definition_enumerationTo.SPECIFIED_VALUES.ENUM_VALUE = append(datatype_definition_enumerationTo.SPECIFIED_VALUES.ENUM_VALUE, CopyBranchENUM_VALUE(mapOrigCopy, _enum_value))
	}

	return
}

func CopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy map[any]any, datatype_definition_integerFrom *DATATYPE_DEFINITION_INTEGER) (datatype_definition_integerTo *DATATYPE_DEFINITION_INTEGER) {

	// datatype_definition_integerFrom has already been copied
	if _datatype_definition_integerTo, ok := mapOrigCopy[datatype_definition_integerFrom]; ok {
		datatype_definition_integerTo = _datatype_definition_integerTo.(*DATATYPE_DEFINITION_INTEGER)
		return
	}

	datatype_definition_integerTo = new(DATATYPE_DEFINITION_INTEGER)
	mapOrigCopy[datatype_definition_integerFrom] = datatype_definition_integerTo
	datatype_definition_integerFrom.CopyBasicFields(datatype_definition_integerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_integerFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		datatype_definition_integerTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(datatype_definition_integerTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}

	return
}

func CopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy map[any]any, datatype_definition_realFrom *DATATYPE_DEFINITION_REAL) (datatype_definition_realTo *DATATYPE_DEFINITION_REAL) {

	// datatype_definition_realFrom has already been copied
	if _datatype_definition_realTo, ok := mapOrigCopy[datatype_definition_realFrom]; ok {
		datatype_definition_realTo = _datatype_definition_realTo.(*DATATYPE_DEFINITION_REAL)
		return
	}

	datatype_definition_realTo = new(DATATYPE_DEFINITION_REAL)
	mapOrigCopy[datatype_definition_realFrom] = datatype_definition_realTo
	datatype_definition_realFrom.CopyBasicFields(datatype_definition_realTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_realFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		datatype_definition_realTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(datatype_definition_realTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}

	return
}

func CopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy map[any]any, datatype_definition_stringFrom *DATATYPE_DEFINITION_STRING) (datatype_definition_stringTo *DATATYPE_DEFINITION_STRING) {

	// datatype_definition_stringFrom has already been copied
	if _datatype_definition_stringTo, ok := mapOrigCopy[datatype_definition_stringFrom]; ok {
		datatype_definition_stringTo = _datatype_definition_stringTo.(*DATATYPE_DEFINITION_STRING)
		return
	}

	datatype_definition_stringTo = new(DATATYPE_DEFINITION_STRING)
	mapOrigCopy[datatype_definition_stringFrom] = datatype_definition_stringTo
	datatype_definition_stringFrom.CopyBasicFields(datatype_definition_stringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_stringFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		datatype_definition_stringTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(datatype_definition_stringTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}

	return
}

func CopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy map[any]any, datatype_definition_xhtmlFrom *DATATYPE_DEFINITION_XHTML) (datatype_definition_xhtmlTo *DATATYPE_DEFINITION_XHTML) {

	// datatype_definition_xhtmlFrom has already been copied
	if _datatype_definition_xhtmlTo, ok := mapOrigCopy[datatype_definition_xhtmlFrom]; ok {
		datatype_definition_xhtmlTo = _datatype_definition_xhtmlTo.(*DATATYPE_DEFINITION_XHTML)
		return
	}

	datatype_definition_xhtmlTo = new(DATATYPE_DEFINITION_XHTML)
	mapOrigCopy[datatype_definition_xhtmlFrom] = datatype_definition_xhtmlTo
	datatype_definition_xhtmlFrom.CopyBasicFields(datatype_definition_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_xhtmlFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		datatype_definition_xhtmlTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(datatype_definition_xhtmlTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}

	return
}

func CopyBranchEMBEDDED_VALUE(mapOrigCopy map[any]any, embedded_valueFrom *EMBEDDED_VALUE) (embedded_valueTo *EMBEDDED_VALUE) {

	// embedded_valueFrom has already been copied
	if _embedded_valueTo, ok := mapOrigCopy[embedded_valueFrom]; ok {
		embedded_valueTo = _embedded_valueTo.(*EMBEDDED_VALUE)
		return
	}

	embedded_valueTo = new(EMBEDDED_VALUE)
	mapOrigCopy[embedded_valueFrom] = embedded_valueTo
	embedded_valueFrom.CopyBasicFields(embedded_valueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchENUM_VALUE(mapOrigCopy map[any]any, enum_valueFrom *ENUM_VALUE) (enum_valueTo *ENUM_VALUE) {

	// enum_valueFrom has already been copied
	if _enum_valueTo, ok := mapOrigCopy[enum_valueFrom]; ok {
		enum_valueTo = _enum_valueTo.(*ENUM_VALUE)
		return
	}

	enum_valueTo = new(ENUM_VALUE)
	mapOrigCopy[enum_valueFrom] = enum_valueTo
	enum_valueFrom.CopyBasicFields(enum_valueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range enum_valueFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		enum_valueTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(enum_valueTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _embedded_value := range enum_valueFrom.PROPERTIES.EMBEDDED_VALUE {
		enum_valueTo.PROPERTIES.EMBEDDED_VALUE = append(enum_valueTo.PROPERTIES.EMBEDDED_VALUE, CopyBranchEMBEDDED_VALUE(mapOrigCopy, _embedded_value))
	}

	return
}

func CopyBranchRELATION_GROUP(mapOrigCopy map[any]any, relation_groupFrom *RELATION_GROUP) (relation_groupTo *RELATION_GROUP) {

	// relation_groupFrom has already been copied
	if _relation_groupTo, ok := mapOrigCopy[relation_groupFrom]; ok {
		relation_groupTo = _relation_groupTo.(*RELATION_GROUP)
		return
	}

	relation_groupTo = new(RELATION_GROUP)
	mapOrigCopy[relation_groupFrom] = relation_groupTo
	relation_groupFrom.CopyBasicFields(relation_groupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range relation_groupFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		relation_groupTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(relation_groupTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}

	return
}

func CopyBranchRELATION_GROUP_TYPE(mapOrigCopy map[any]any, relation_group_typeFrom *RELATION_GROUP_TYPE) (relation_group_typeTo *RELATION_GROUP_TYPE) {

	// relation_group_typeFrom has already been copied
	if _relation_group_typeTo, ok := mapOrigCopy[relation_group_typeFrom]; ok {
		relation_group_typeTo = _relation_group_typeTo.(*RELATION_GROUP_TYPE)
		return
	}

	relation_group_typeTo = new(RELATION_GROUP_TYPE)
	mapOrigCopy[relation_group_typeFrom] = relation_group_typeTo
	relation_group_typeFrom.CopyBasicFields(relation_group_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range relation_group_typeFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		relation_group_typeTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(relation_group_typeTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_definition_boolean := range relation_group_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, _attribute_definition_boolean))
	}
	for _, _attribute_definition_date := range relation_group_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, _attribute_definition_date))
	}
	for _, _attribute_definition_enumeration := range relation_group_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, _attribute_definition_enumeration))
	}
	for _, _attribute_definition_integer := range relation_group_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, _attribute_definition_integer))
	}
	for _, _attribute_definition_real := range relation_group_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, _attribute_definition_real))
	}
	for _, _attribute_definition_string := range relation_group_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, _attribute_definition_string))
	}
	for _, _attribute_definition_xhtml := range relation_group_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(relation_group_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, _attribute_definition_xhtml))
	}

	return
}

func CopyBranchREQ_IF(mapOrigCopy map[any]any, req_ifFrom *REQ_IF) (req_ifTo *REQ_IF) {

	// req_ifFrom has already been copied
	if _req_ifTo, ok := mapOrigCopy[req_ifFrom]; ok {
		req_ifTo = _req_ifTo.(*REQ_IF)
		return
	}

	req_ifTo = new(REQ_IF)
	mapOrigCopy[req_ifFrom] = req_ifTo
	req_ifFrom.CopyBasicFields(req_ifTo)

	//insertion point for the staging of instances referenced by pointers
	if req_ifFrom.THE_HEADER.REQ_IF_HEADER != nil {
		req_ifTo.THE_HEADER.REQ_IF_HEADER = CopyBranchREQ_IF_HEADER(mapOrigCopy, req_ifFrom.THE_HEADER.REQ_IF_HEADER)
	}
	if req_ifFrom.CORE_CONTENT.REQ_IF_CONTENT != nil {
		req_ifTo.CORE_CONTENT.REQ_IF_CONTENT = CopyBranchREQ_IF_CONTENT(mapOrigCopy, req_ifFrom.CORE_CONTENT.REQ_IF_CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_tool_extension := range req_ifFrom.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION {
		req_ifTo.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION = append(req_ifTo.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION, CopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy, _req_if_tool_extension))
	}

	return
}

func CopyBranchREQ_IF_CONTENT(mapOrigCopy map[any]any, req_if_contentFrom *REQ_IF_CONTENT) (req_if_contentTo *REQ_IF_CONTENT) {

	// req_if_contentFrom has already been copied
	if _req_if_contentTo, ok := mapOrigCopy[req_if_contentFrom]; ok {
		req_if_contentTo = _req_if_contentTo.(*REQ_IF_CONTENT)
		return
	}

	req_if_contentTo = new(REQ_IF_CONTENT)
	mapOrigCopy[req_if_contentFrom] = req_if_contentTo
	req_if_contentFrom.CopyBasicFields(req_if_contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatype_definition_boolean := range req_if_contentFrom.DATATYPES.DATATYPE_DEFINITION_BOOLEAN {
		req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_BOOLEAN = append(req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_BOOLEAN, CopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy, _datatype_definition_boolean))
	}
	for _, _datatype_definition_date := range req_if_contentFrom.DATATYPES.DATATYPE_DEFINITION_DATE {
		req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_DATE = append(req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_DATE, CopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy, _datatype_definition_date))
	}
	for _, _datatype_definition_enumeration := range req_if_contentFrom.DATATYPES.DATATYPE_DEFINITION_ENUMERATION {
		req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_ENUMERATION = append(req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_ENUMERATION, CopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy, _datatype_definition_enumeration))
	}
	for _, _datatype_definition_integer := range req_if_contentFrom.DATATYPES.DATATYPE_DEFINITION_INTEGER {
		req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_INTEGER = append(req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_INTEGER, CopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy, _datatype_definition_integer))
	}
	for _, _datatype_definition_real := range req_if_contentFrom.DATATYPES.DATATYPE_DEFINITION_REAL {
		req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_REAL = append(req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_REAL, CopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy, _datatype_definition_real))
	}
	for _, _datatype_definition_string := range req_if_contentFrom.DATATYPES.DATATYPE_DEFINITION_STRING {
		req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_STRING = append(req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_STRING, CopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy, _datatype_definition_string))
	}
	for _, _datatype_definition_xhtml := range req_if_contentFrom.DATATYPES.DATATYPE_DEFINITION_XHTML {
		req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_XHTML = append(req_if_contentTo.DATATYPES.DATATYPE_DEFINITION_XHTML, CopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy, _datatype_definition_xhtml))
	}
	for _, _relation_group_type := range req_if_contentFrom.SPEC_TYPES.RELATION_GROUP_TYPE {
		req_if_contentTo.SPEC_TYPES.RELATION_GROUP_TYPE = append(req_if_contentTo.SPEC_TYPES.RELATION_GROUP_TYPE, CopyBranchRELATION_GROUP_TYPE(mapOrigCopy, _relation_group_type))
	}
	for _, _spec_object_type := range req_if_contentFrom.SPEC_TYPES.SPEC_OBJECT_TYPE {
		req_if_contentTo.SPEC_TYPES.SPEC_OBJECT_TYPE = append(req_if_contentTo.SPEC_TYPES.SPEC_OBJECT_TYPE, CopyBranchSPEC_OBJECT_TYPE(mapOrigCopy, _spec_object_type))
	}
	for _, _spec_relation_type := range req_if_contentFrom.SPEC_TYPES.SPEC_RELATION_TYPE {
		req_if_contentTo.SPEC_TYPES.SPEC_RELATION_TYPE = append(req_if_contentTo.SPEC_TYPES.SPEC_RELATION_TYPE, CopyBranchSPEC_RELATION_TYPE(mapOrigCopy, _spec_relation_type))
	}
	for _, _specification_type := range req_if_contentFrom.SPEC_TYPES.SPECIFICATION_TYPE {
		req_if_contentTo.SPEC_TYPES.SPECIFICATION_TYPE = append(req_if_contentTo.SPEC_TYPES.SPECIFICATION_TYPE, CopyBranchSPECIFICATION_TYPE(mapOrigCopy, _specification_type))
	}
	for _, _spec_object := range req_if_contentFrom.SPEC_OBJECTS.SPEC_OBJECT {
		req_if_contentTo.SPEC_OBJECTS.SPEC_OBJECT = append(req_if_contentTo.SPEC_OBJECTS.SPEC_OBJECT, CopyBranchSPEC_OBJECT(mapOrigCopy, _spec_object))
	}
	for _, _spec_relation := range req_if_contentFrom.SPEC_RELATIONS.SPEC_RELATION {
		req_if_contentTo.SPEC_RELATIONS.SPEC_RELATION = append(req_if_contentTo.SPEC_RELATIONS.SPEC_RELATION, CopyBranchSPEC_RELATION(mapOrigCopy, _spec_relation))
	}
	for _, _specification := range req_if_contentFrom.SPECIFICATIONS.SPECIFICATION {
		req_if_contentTo.SPECIFICATIONS.SPECIFICATION = append(req_if_contentTo.SPECIFICATIONS.SPECIFICATION, CopyBranchSPECIFICATION(mapOrigCopy, _specification))
	}
	for _, _relation_group := range req_if_contentFrom.SPEC_RELATION_GROUPS.RELATION_GROUP {
		req_if_contentTo.SPEC_RELATION_GROUPS.RELATION_GROUP = append(req_if_contentTo.SPEC_RELATION_GROUPS.RELATION_GROUP, CopyBranchRELATION_GROUP(mapOrigCopy, _relation_group))
	}

	return
}

func CopyBranchREQ_IF_HEADER(mapOrigCopy map[any]any, req_if_headerFrom *REQ_IF_HEADER) (req_if_headerTo *REQ_IF_HEADER) {

	// req_if_headerFrom has already been copied
	if _req_if_headerTo, ok := mapOrigCopy[req_if_headerFrom]; ok {
		req_if_headerTo = _req_if_headerTo.(*REQ_IF_HEADER)
		return
	}

	req_if_headerTo = new(REQ_IF_HEADER)
	mapOrigCopy[req_if_headerFrom] = req_if_headerTo
	req_if_headerFrom.CopyBasicFields(req_if_headerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy map[any]any, req_if_tool_extensionFrom *REQ_IF_TOOL_EXTENSION) (req_if_tool_extensionTo *REQ_IF_TOOL_EXTENSION) {

	// req_if_tool_extensionFrom has already been copied
	if _req_if_tool_extensionTo, ok := mapOrigCopy[req_if_tool_extensionFrom]; ok {
		req_if_tool_extensionTo = _req_if_tool_extensionTo.(*REQ_IF_TOOL_EXTENSION)
		return
	}

	req_if_tool_extensionTo = new(REQ_IF_TOOL_EXTENSION)
	mapOrigCopy[req_if_tool_extensionFrom] = req_if_tool_extensionTo
	req_if_tool_extensionFrom.CopyBasicFields(req_if_tool_extensionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSPECIFICATION(mapOrigCopy map[any]any, specificationFrom *SPECIFICATION) (specificationTo *SPECIFICATION) {

	// specificationFrom has already been copied
	if _specificationTo, ok := mapOrigCopy[specificationFrom]; ok {
		specificationTo = _specificationTo.(*SPECIFICATION)
		return
	}

	specificationTo = new(SPECIFICATION)
	mapOrigCopy[specificationFrom] = specificationTo
	specificationFrom.CopyBasicFields(specificationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range specificationFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		specificationTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(specificationTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_value_boolean := range specificationFrom.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		specificationTo.VALUES.ATTRIBUTE_VALUE_BOOLEAN = append(specificationTo.VALUES.ATTRIBUTE_VALUE_BOOLEAN, CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, _attribute_value_boolean))
	}
	for _, _attribute_value_date := range specificationFrom.VALUES.ATTRIBUTE_VALUE_DATE {
		specificationTo.VALUES.ATTRIBUTE_VALUE_DATE = append(specificationTo.VALUES.ATTRIBUTE_VALUE_DATE, CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, _attribute_value_date))
	}
	for _, _attribute_value_enumeration := range specificationFrom.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		specificationTo.VALUES.ATTRIBUTE_VALUE_ENUMERATION = append(specificationTo.VALUES.ATTRIBUTE_VALUE_ENUMERATION, CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, _attribute_value_enumeration))
	}
	for _, _attribute_value_integer := range specificationFrom.VALUES.ATTRIBUTE_VALUE_INTEGER {
		specificationTo.VALUES.ATTRIBUTE_VALUE_INTEGER = append(specificationTo.VALUES.ATTRIBUTE_VALUE_INTEGER, CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, _attribute_value_integer))
	}
	for _, _attribute_value_real := range specificationFrom.VALUES.ATTRIBUTE_VALUE_REAL {
		specificationTo.VALUES.ATTRIBUTE_VALUE_REAL = append(specificationTo.VALUES.ATTRIBUTE_VALUE_REAL, CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, _attribute_value_real))
	}
	for _, _attribute_value_string := range specificationFrom.VALUES.ATTRIBUTE_VALUE_STRING {
		specificationTo.VALUES.ATTRIBUTE_VALUE_STRING = append(specificationTo.VALUES.ATTRIBUTE_VALUE_STRING, CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, _attribute_value_string))
	}
	for _, _attribute_value_xhtml := range specificationFrom.VALUES.ATTRIBUTE_VALUE_XHTML {
		specificationTo.VALUES.ATTRIBUTE_VALUE_XHTML = append(specificationTo.VALUES.ATTRIBUTE_VALUE_XHTML, CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, _attribute_value_xhtml))
	}
	for _, _spec_hierarchy := range specificationFrom.CHILDREN.SPEC_HIERARCHY {
		specificationTo.CHILDREN.SPEC_HIERARCHY = append(specificationTo.CHILDREN.SPEC_HIERARCHY, CopyBranchSPEC_HIERARCHY(mapOrigCopy, _spec_hierarchy))
	}

	return
}

func CopyBranchSPECIFICATION_TYPE(mapOrigCopy map[any]any, specification_typeFrom *SPECIFICATION_TYPE) (specification_typeTo *SPECIFICATION_TYPE) {

	// specification_typeFrom has already been copied
	if _specification_typeTo, ok := mapOrigCopy[specification_typeFrom]; ok {
		specification_typeTo = _specification_typeTo.(*SPECIFICATION_TYPE)
		return
	}

	specification_typeTo = new(SPECIFICATION_TYPE)
	mapOrigCopy[specification_typeFrom] = specification_typeTo
	specification_typeFrom.CopyBasicFields(specification_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range specification_typeFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		specification_typeTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(specification_typeTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_definition_boolean := range specification_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, _attribute_definition_boolean))
	}
	for _, _attribute_definition_date := range specification_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, _attribute_definition_date))
	}
	for _, _attribute_definition_enumeration := range specification_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, _attribute_definition_enumeration))
	}
	for _, _attribute_definition_integer := range specification_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, _attribute_definition_integer))
	}
	for _, _attribute_definition_real := range specification_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, _attribute_definition_real))
	}
	for _, _attribute_definition_string := range specification_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, _attribute_definition_string))
	}
	for _, _attribute_definition_xhtml := range specification_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(specification_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, _attribute_definition_xhtml))
	}

	return
}

func CopyBranchSPEC_HIERARCHY(mapOrigCopy map[any]any, spec_hierarchyFrom *SPEC_HIERARCHY) (spec_hierarchyTo *SPEC_HIERARCHY) {

	// spec_hierarchyFrom has already been copied
	if _spec_hierarchyTo, ok := mapOrigCopy[spec_hierarchyFrom]; ok {
		spec_hierarchyTo = _spec_hierarchyTo.(*SPEC_HIERARCHY)
		return
	}

	spec_hierarchyTo = new(SPEC_HIERARCHY)
	mapOrigCopy[spec_hierarchyFrom] = spec_hierarchyTo
	spec_hierarchyFrom.CopyBasicFields(spec_hierarchyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_hierarchyFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		spec_hierarchyTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(spec_hierarchyTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _spec_hierarchy := range spec_hierarchyFrom.CHILDREN.SPEC_HIERARCHY {
		spec_hierarchyTo.CHILDREN.SPEC_HIERARCHY = append(spec_hierarchyTo.CHILDREN.SPEC_HIERARCHY, CopyBranchSPEC_HIERARCHY(mapOrigCopy, _spec_hierarchy))
	}

	return
}

func CopyBranchSPEC_OBJECT(mapOrigCopy map[any]any, spec_objectFrom *SPEC_OBJECT) (spec_objectTo *SPEC_OBJECT) {

	// spec_objectFrom has already been copied
	if _spec_objectTo, ok := mapOrigCopy[spec_objectFrom]; ok {
		spec_objectTo = _spec_objectTo.(*SPEC_OBJECT)
		return
	}

	spec_objectTo = new(SPEC_OBJECT)
	mapOrigCopy[spec_objectFrom] = spec_objectTo
	spec_objectFrom.CopyBasicFields(spec_objectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_objectFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		spec_objectTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(spec_objectTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_value_boolean := range spec_objectFrom.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		spec_objectTo.VALUES.ATTRIBUTE_VALUE_BOOLEAN = append(spec_objectTo.VALUES.ATTRIBUTE_VALUE_BOOLEAN, CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, _attribute_value_boolean))
	}
	for _, _attribute_value_date := range spec_objectFrom.VALUES.ATTRIBUTE_VALUE_DATE {
		spec_objectTo.VALUES.ATTRIBUTE_VALUE_DATE = append(spec_objectTo.VALUES.ATTRIBUTE_VALUE_DATE, CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, _attribute_value_date))
	}
	for _, _attribute_value_enumeration := range spec_objectFrom.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		spec_objectTo.VALUES.ATTRIBUTE_VALUE_ENUMERATION = append(spec_objectTo.VALUES.ATTRIBUTE_VALUE_ENUMERATION, CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, _attribute_value_enumeration))
	}
	for _, _attribute_value_integer := range spec_objectFrom.VALUES.ATTRIBUTE_VALUE_INTEGER {
		spec_objectTo.VALUES.ATTRIBUTE_VALUE_INTEGER = append(spec_objectTo.VALUES.ATTRIBUTE_VALUE_INTEGER, CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, _attribute_value_integer))
	}
	for _, _attribute_value_real := range spec_objectFrom.VALUES.ATTRIBUTE_VALUE_REAL {
		spec_objectTo.VALUES.ATTRIBUTE_VALUE_REAL = append(spec_objectTo.VALUES.ATTRIBUTE_VALUE_REAL, CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, _attribute_value_real))
	}
	for _, _attribute_value_string := range spec_objectFrom.VALUES.ATTRIBUTE_VALUE_STRING {
		spec_objectTo.VALUES.ATTRIBUTE_VALUE_STRING = append(spec_objectTo.VALUES.ATTRIBUTE_VALUE_STRING, CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, _attribute_value_string))
	}
	for _, _attribute_value_xhtml := range spec_objectFrom.VALUES.ATTRIBUTE_VALUE_XHTML {
		spec_objectTo.VALUES.ATTRIBUTE_VALUE_XHTML = append(spec_objectTo.VALUES.ATTRIBUTE_VALUE_XHTML, CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, _attribute_value_xhtml))
	}

	return
}

func CopyBranchSPEC_OBJECT_TYPE(mapOrigCopy map[any]any, spec_object_typeFrom *SPEC_OBJECT_TYPE) (spec_object_typeTo *SPEC_OBJECT_TYPE) {

	// spec_object_typeFrom has already been copied
	if _spec_object_typeTo, ok := mapOrigCopy[spec_object_typeFrom]; ok {
		spec_object_typeTo = _spec_object_typeTo.(*SPEC_OBJECT_TYPE)
		return
	}

	spec_object_typeTo = new(SPEC_OBJECT_TYPE)
	mapOrigCopy[spec_object_typeFrom] = spec_object_typeTo
	spec_object_typeFrom.CopyBasicFields(spec_object_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_object_typeFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		spec_object_typeTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(spec_object_typeTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_definition_boolean := range spec_object_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, _attribute_definition_boolean))
	}
	for _, _attribute_definition_date := range spec_object_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, _attribute_definition_date))
	}
	for _, _attribute_definition_enumeration := range spec_object_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, _attribute_definition_enumeration))
	}
	for _, _attribute_definition_integer := range spec_object_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, _attribute_definition_integer))
	}
	for _, _attribute_definition_real := range spec_object_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, _attribute_definition_real))
	}
	for _, _attribute_definition_string := range spec_object_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, _attribute_definition_string))
	}
	for _, _attribute_definition_xhtml := range spec_object_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(spec_object_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, _attribute_definition_xhtml))
	}

	return
}

func CopyBranchSPEC_RELATION(mapOrigCopy map[any]any, spec_relationFrom *SPEC_RELATION) (spec_relationTo *SPEC_RELATION) {

	// spec_relationFrom has already been copied
	if _spec_relationTo, ok := mapOrigCopy[spec_relationFrom]; ok {
		spec_relationTo = _spec_relationTo.(*SPEC_RELATION)
		return
	}

	spec_relationTo = new(SPEC_RELATION)
	mapOrigCopy[spec_relationFrom] = spec_relationTo
	spec_relationFrom.CopyBasicFields(spec_relationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_relationFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		spec_relationTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(spec_relationTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_value_boolean := range spec_relationFrom.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		spec_relationTo.VALUES.ATTRIBUTE_VALUE_BOOLEAN = append(spec_relationTo.VALUES.ATTRIBUTE_VALUE_BOOLEAN, CopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, _attribute_value_boolean))
	}
	for _, _attribute_value_date := range spec_relationFrom.VALUES.ATTRIBUTE_VALUE_DATE {
		spec_relationTo.VALUES.ATTRIBUTE_VALUE_DATE = append(spec_relationTo.VALUES.ATTRIBUTE_VALUE_DATE, CopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, _attribute_value_date))
	}
	for _, _attribute_value_enumeration := range spec_relationFrom.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		spec_relationTo.VALUES.ATTRIBUTE_VALUE_ENUMERATION = append(spec_relationTo.VALUES.ATTRIBUTE_VALUE_ENUMERATION, CopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, _attribute_value_enumeration))
	}
	for _, _attribute_value_integer := range spec_relationFrom.VALUES.ATTRIBUTE_VALUE_INTEGER {
		spec_relationTo.VALUES.ATTRIBUTE_VALUE_INTEGER = append(spec_relationTo.VALUES.ATTRIBUTE_VALUE_INTEGER, CopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, _attribute_value_integer))
	}
	for _, _attribute_value_real := range spec_relationFrom.VALUES.ATTRIBUTE_VALUE_REAL {
		spec_relationTo.VALUES.ATTRIBUTE_VALUE_REAL = append(spec_relationTo.VALUES.ATTRIBUTE_VALUE_REAL, CopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, _attribute_value_real))
	}
	for _, _attribute_value_string := range spec_relationFrom.VALUES.ATTRIBUTE_VALUE_STRING {
		spec_relationTo.VALUES.ATTRIBUTE_VALUE_STRING = append(spec_relationTo.VALUES.ATTRIBUTE_VALUE_STRING, CopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, _attribute_value_string))
	}
	for _, _attribute_value_xhtml := range spec_relationFrom.VALUES.ATTRIBUTE_VALUE_XHTML {
		spec_relationTo.VALUES.ATTRIBUTE_VALUE_XHTML = append(spec_relationTo.VALUES.ATTRIBUTE_VALUE_XHTML, CopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, _attribute_value_xhtml))
	}

	return
}

func CopyBranchSPEC_RELATION_TYPE(mapOrigCopy map[any]any, spec_relation_typeFrom *SPEC_RELATION_TYPE) (spec_relation_typeTo *SPEC_RELATION_TYPE) {

	// spec_relation_typeFrom has already been copied
	if _spec_relation_typeTo, ok := mapOrigCopy[spec_relation_typeFrom]; ok {
		spec_relation_typeTo = _spec_relation_typeTo.(*SPEC_RELATION_TYPE)
		return
	}

	spec_relation_typeTo = new(SPEC_RELATION_TYPE)
	mapOrigCopy[spec_relation_typeFrom] = spec_relation_typeTo
	spec_relation_typeFrom.CopyBasicFields(spec_relation_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_relation_typeFrom.ALTERNATIVE_ID.ALTERNATIVE_ID {
		spec_relation_typeTo.ALTERNATIVE_ID.ALTERNATIVE_ID = append(spec_relation_typeTo.ALTERNATIVE_ID.ALTERNATIVE_ID, CopyBranchALTERNATIVE_ID(mapOrigCopy, _alternative_id))
	}
	for _, _attribute_definition_boolean := range spec_relation_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN = append(spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN, CopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, _attribute_definition_boolean))
	}
	for _, _attribute_definition_date := range spec_relation_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE = append(spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE, CopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, _attribute_definition_date))
	}
	for _, _attribute_definition_enumeration := range spec_relation_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION = append(spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION, CopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, _attribute_definition_enumeration))
	}
	for _, _attribute_definition_integer := range spec_relation_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER = append(spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER, CopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, _attribute_definition_integer))
	}
	for _, _attribute_definition_real := range spec_relation_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL = append(spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL, CopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, _attribute_definition_real))
	}
	for _, _attribute_definition_string := range spec_relation_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING = append(spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING, CopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, _attribute_definition_string))
	}
	for _, _attribute_definition_xhtml := range spec_relation_typeFrom.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML = append(spec_relation_typeTo.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, CopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, _attribute_definition_xhtml))
	}

	return
}

func CopyBranchXHTML_CONTENT(mapOrigCopy map[any]any, xhtml_contentFrom *XHTML_CONTENT) (xhtml_contentTo *XHTML_CONTENT) {

	// xhtml_contentFrom has already been copied
	if _xhtml_contentTo, ok := mapOrigCopy[xhtml_contentFrom]; ok {
		xhtml_contentTo = _xhtml_contentTo.(*XHTML_CONTENT)
		return
	}

	xhtml_contentTo = new(XHTML_CONTENT)
	mapOrigCopy[xhtml_contentFrom] = xhtml_contentTo
	xhtml_contentFrom.CopyBasicFields(xhtml_contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_InlPres_type(mapOrigCopy map[any]any, xhtml_inlpres_typeFrom *Xhtml_InlPres_type) (xhtml_inlpres_typeTo *Xhtml_InlPres_type) {

	// xhtml_inlpres_typeFrom has already been copied
	if _xhtml_inlpres_typeTo, ok := mapOrigCopy[xhtml_inlpres_typeFrom]; ok {
		xhtml_inlpres_typeTo = _xhtml_inlpres_typeTo.(*Xhtml_InlPres_type)
		return
	}

	xhtml_inlpres_typeTo = new(Xhtml_InlPres_type)
	mapOrigCopy[xhtml_inlpres_typeFrom] = xhtml_inlpres_typeTo
	xhtml_inlpres_typeFrom.CopyBasicFields(xhtml_inlpres_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_a_type(mapOrigCopy map[any]any, xhtml_a_typeFrom *Xhtml_a_type) (xhtml_a_typeTo *Xhtml_a_type) {

	// xhtml_a_typeFrom has already been copied
	if _xhtml_a_typeTo, ok := mapOrigCopy[xhtml_a_typeFrom]; ok {
		xhtml_a_typeTo = _xhtml_a_typeTo.(*Xhtml_a_type)
		return
	}

	xhtml_a_typeTo = new(Xhtml_a_type)
	mapOrigCopy[xhtml_a_typeFrom] = xhtml_a_typeTo
	xhtml_a_typeFrom.CopyBasicFields(xhtml_a_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_abbr_type(mapOrigCopy map[any]any, xhtml_abbr_typeFrom *Xhtml_abbr_type) (xhtml_abbr_typeTo *Xhtml_abbr_type) {

	// xhtml_abbr_typeFrom has already been copied
	if _xhtml_abbr_typeTo, ok := mapOrigCopy[xhtml_abbr_typeFrom]; ok {
		xhtml_abbr_typeTo = _xhtml_abbr_typeTo.(*Xhtml_abbr_type)
		return
	}

	xhtml_abbr_typeTo = new(Xhtml_abbr_type)
	mapOrigCopy[xhtml_abbr_typeFrom] = xhtml_abbr_typeTo
	xhtml_abbr_typeFrom.CopyBasicFields(xhtml_abbr_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_acronym_type(mapOrigCopy map[any]any, xhtml_acronym_typeFrom *Xhtml_acronym_type) (xhtml_acronym_typeTo *Xhtml_acronym_type) {

	// xhtml_acronym_typeFrom has already been copied
	if _xhtml_acronym_typeTo, ok := mapOrigCopy[xhtml_acronym_typeFrom]; ok {
		xhtml_acronym_typeTo = _xhtml_acronym_typeTo.(*Xhtml_acronym_type)
		return
	}

	xhtml_acronym_typeTo = new(Xhtml_acronym_type)
	mapOrigCopy[xhtml_acronym_typeFrom] = xhtml_acronym_typeTo
	xhtml_acronym_typeFrom.CopyBasicFields(xhtml_acronym_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_address_type(mapOrigCopy map[any]any, xhtml_address_typeFrom *Xhtml_address_type) (xhtml_address_typeTo *Xhtml_address_type) {

	// xhtml_address_typeFrom has already been copied
	if _xhtml_address_typeTo, ok := mapOrigCopy[xhtml_address_typeFrom]; ok {
		xhtml_address_typeTo = _xhtml_address_typeTo.(*Xhtml_address_type)
		return
	}

	xhtml_address_typeTo = new(Xhtml_address_type)
	mapOrigCopy[xhtml_address_typeFrom] = xhtml_address_typeTo
	xhtml_address_typeFrom.CopyBasicFields(xhtml_address_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_blockquote_type(mapOrigCopy map[any]any, xhtml_blockquote_typeFrom *Xhtml_blockquote_type) (xhtml_blockquote_typeTo *Xhtml_blockquote_type) {

	// xhtml_blockquote_typeFrom has already been copied
	if _xhtml_blockquote_typeTo, ok := mapOrigCopy[xhtml_blockquote_typeFrom]; ok {
		xhtml_blockquote_typeTo = _xhtml_blockquote_typeTo.(*Xhtml_blockquote_type)
		return
	}

	xhtml_blockquote_typeTo = new(Xhtml_blockquote_type)
	mapOrigCopy[xhtml_blockquote_typeFrom] = xhtml_blockquote_typeTo
	xhtml_blockquote_typeFrom.CopyBasicFields(xhtml_blockquote_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_br_type(mapOrigCopy map[any]any, xhtml_br_typeFrom *Xhtml_br_type) (xhtml_br_typeTo *Xhtml_br_type) {

	// xhtml_br_typeFrom has already been copied
	if _xhtml_br_typeTo, ok := mapOrigCopy[xhtml_br_typeFrom]; ok {
		xhtml_br_typeTo = _xhtml_br_typeTo.(*Xhtml_br_type)
		return
	}

	xhtml_br_typeTo = new(Xhtml_br_type)
	mapOrigCopy[xhtml_br_typeFrom] = xhtml_br_typeTo
	xhtml_br_typeFrom.CopyBasicFields(xhtml_br_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_caption_type(mapOrigCopy map[any]any, xhtml_caption_typeFrom *Xhtml_caption_type) (xhtml_caption_typeTo *Xhtml_caption_type) {

	// xhtml_caption_typeFrom has already been copied
	if _xhtml_caption_typeTo, ok := mapOrigCopy[xhtml_caption_typeFrom]; ok {
		xhtml_caption_typeTo = _xhtml_caption_typeTo.(*Xhtml_caption_type)
		return
	}

	xhtml_caption_typeTo = new(Xhtml_caption_type)
	mapOrigCopy[xhtml_caption_typeFrom] = xhtml_caption_typeTo
	xhtml_caption_typeFrom.CopyBasicFields(xhtml_caption_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_cite_type(mapOrigCopy map[any]any, xhtml_cite_typeFrom *Xhtml_cite_type) (xhtml_cite_typeTo *Xhtml_cite_type) {

	// xhtml_cite_typeFrom has already been copied
	if _xhtml_cite_typeTo, ok := mapOrigCopy[xhtml_cite_typeFrom]; ok {
		xhtml_cite_typeTo = _xhtml_cite_typeTo.(*Xhtml_cite_type)
		return
	}

	xhtml_cite_typeTo = new(Xhtml_cite_type)
	mapOrigCopy[xhtml_cite_typeFrom] = xhtml_cite_typeTo
	xhtml_cite_typeFrom.CopyBasicFields(xhtml_cite_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_code_type(mapOrigCopy map[any]any, xhtml_code_typeFrom *Xhtml_code_type) (xhtml_code_typeTo *Xhtml_code_type) {

	// xhtml_code_typeFrom has already been copied
	if _xhtml_code_typeTo, ok := mapOrigCopy[xhtml_code_typeFrom]; ok {
		xhtml_code_typeTo = _xhtml_code_typeTo.(*Xhtml_code_type)
		return
	}

	xhtml_code_typeTo = new(Xhtml_code_type)
	mapOrigCopy[xhtml_code_typeFrom] = xhtml_code_typeTo
	xhtml_code_typeFrom.CopyBasicFields(xhtml_code_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_col_type(mapOrigCopy map[any]any, xhtml_col_typeFrom *Xhtml_col_type) (xhtml_col_typeTo *Xhtml_col_type) {

	// xhtml_col_typeFrom has already been copied
	if _xhtml_col_typeTo, ok := mapOrigCopy[xhtml_col_typeFrom]; ok {
		xhtml_col_typeTo = _xhtml_col_typeTo.(*Xhtml_col_type)
		return
	}

	xhtml_col_typeTo = new(Xhtml_col_type)
	mapOrigCopy[xhtml_col_typeFrom] = xhtml_col_typeTo
	xhtml_col_typeFrom.CopyBasicFields(xhtml_col_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_colgroup_type(mapOrigCopy map[any]any, xhtml_colgroup_typeFrom *Xhtml_colgroup_type) (xhtml_colgroup_typeTo *Xhtml_colgroup_type) {

	// xhtml_colgroup_typeFrom has already been copied
	if _xhtml_colgroup_typeTo, ok := mapOrigCopy[xhtml_colgroup_typeFrom]; ok {
		xhtml_colgroup_typeTo = _xhtml_colgroup_typeTo.(*Xhtml_colgroup_type)
		return
	}

	xhtml_colgroup_typeTo = new(Xhtml_colgroup_type)
	mapOrigCopy[xhtml_colgroup_typeFrom] = xhtml_colgroup_typeTo
	xhtml_colgroup_typeFrom.CopyBasicFields(xhtml_colgroup_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_dd_type(mapOrigCopy map[any]any, xhtml_dd_typeFrom *Xhtml_dd_type) (xhtml_dd_typeTo *Xhtml_dd_type) {

	// xhtml_dd_typeFrom has already been copied
	if _xhtml_dd_typeTo, ok := mapOrigCopy[xhtml_dd_typeFrom]; ok {
		xhtml_dd_typeTo = _xhtml_dd_typeTo.(*Xhtml_dd_type)
		return
	}

	xhtml_dd_typeTo = new(Xhtml_dd_type)
	mapOrigCopy[xhtml_dd_typeFrom] = xhtml_dd_typeTo
	xhtml_dd_typeFrom.CopyBasicFields(xhtml_dd_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_dfn_type(mapOrigCopy map[any]any, xhtml_dfn_typeFrom *Xhtml_dfn_type) (xhtml_dfn_typeTo *Xhtml_dfn_type) {

	// xhtml_dfn_typeFrom has already been copied
	if _xhtml_dfn_typeTo, ok := mapOrigCopy[xhtml_dfn_typeFrom]; ok {
		xhtml_dfn_typeTo = _xhtml_dfn_typeTo.(*Xhtml_dfn_type)
		return
	}

	xhtml_dfn_typeTo = new(Xhtml_dfn_type)
	mapOrigCopy[xhtml_dfn_typeFrom] = xhtml_dfn_typeTo
	xhtml_dfn_typeFrom.CopyBasicFields(xhtml_dfn_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_div_type(mapOrigCopy map[any]any, xhtml_div_typeFrom *Xhtml_div_type) (xhtml_div_typeTo *Xhtml_div_type) {

	// xhtml_div_typeFrom has already been copied
	if _xhtml_div_typeTo, ok := mapOrigCopy[xhtml_div_typeFrom]; ok {
		xhtml_div_typeTo = _xhtml_div_typeTo.(*Xhtml_div_type)
		return
	}

	xhtml_div_typeTo = new(Xhtml_div_type)
	mapOrigCopy[xhtml_div_typeFrom] = xhtml_div_typeTo
	xhtml_div_typeFrom.CopyBasicFields(xhtml_div_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_dl_type(mapOrigCopy map[any]any, xhtml_dl_typeFrom *Xhtml_dl_type) (xhtml_dl_typeTo *Xhtml_dl_type) {

	// xhtml_dl_typeFrom has already been copied
	if _xhtml_dl_typeTo, ok := mapOrigCopy[xhtml_dl_typeFrom]; ok {
		xhtml_dl_typeTo = _xhtml_dl_typeTo.(*Xhtml_dl_type)
		return
	}

	xhtml_dl_typeTo = new(Xhtml_dl_type)
	mapOrigCopy[xhtml_dl_typeFrom] = xhtml_dl_typeTo
	xhtml_dl_typeFrom.CopyBasicFields(xhtml_dl_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_dt_type(mapOrigCopy map[any]any, xhtml_dt_typeFrom *Xhtml_dt_type) (xhtml_dt_typeTo *Xhtml_dt_type) {

	// xhtml_dt_typeFrom has already been copied
	if _xhtml_dt_typeTo, ok := mapOrigCopy[xhtml_dt_typeFrom]; ok {
		xhtml_dt_typeTo = _xhtml_dt_typeTo.(*Xhtml_dt_type)
		return
	}

	xhtml_dt_typeTo = new(Xhtml_dt_type)
	mapOrigCopy[xhtml_dt_typeFrom] = xhtml_dt_typeTo
	xhtml_dt_typeFrom.CopyBasicFields(xhtml_dt_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_edit_type(mapOrigCopy map[any]any, xhtml_edit_typeFrom *Xhtml_edit_type) (xhtml_edit_typeTo *Xhtml_edit_type) {

	// xhtml_edit_typeFrom has already been copied
	if _xhtml_edit_typeTo, ok := mapOrigCopy[xhtml_edit_typeFrom]; ok {
		xhtml_edit_typeTo = _xhtml_edit_typeTo.(*Xhtml_edit_type)
		return
	}

	xhtml_edit_typeTo = new(Xhtml_edit_type)
	mapOrigCopy[xhtml_edit_typeFrom] = xhtml_edit_typeTo
	xhtml_edit_typeFrom.CopyBasicFields(xhtml_edit_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_em_type(mapOrigCopy map[any]any, xhtml_em_typeFrom *Xhtml_em_type) (xhtml_em_typeTo *Xhtml_em_type) {

	// xhtml_em_typeFrom has already been copied
	if _xhtml_em_typeTo, ok := mapOrigCopy[xhtml_em_typeFrom]; ok {
		xhtml_em_typeTo = _xhtml_em_typeTo.(*Xhtml_em_type)
		return
	}

	xhtml_em_typeTo = new(Xhtml_em_type)
	mapOrigCopy[xhtml_em_typeFrom] = xhtml_em_typeTo
	xhtml_em_typeFrom.CopyBasicFields(xhtml_em_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_h1_type(mapOrigCopy map[any]any, xhtml_h1_typeFrom *Xhtml_h1_type) (xhtml_h1_typeTo *Xhtml_h1_type) {

	// xhtml_h1_typeFrom has already been copied
	if _xhtml_h1_typeTo, ok := mapOrigCopy[xhtml_h1_typeFrom]; ok {
		xhtml_h1_typeTo = _xhtml_h1_typeTo.(*Xhtml_h1_type)
		return
	}

	xhtml_h1_typeTo = new(Xhtml_h1_type)
	mapOrigCopy[xhtml_h1_typeFrom] = xhtml_h1_typeTo
	xhtml_h1_typeFrom.CopyBasicFields(xhtml_h1_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_h2_type(mapOrigCopy map[any]any, xhtml_h2_typeFrom *Xhtml_h2_type) (xhtml_h2_typeTo *Xhtml_h2_type) {

	// xhtml_h2_typeFrom has already been copied
	if _xhtml_h2_typeTo, ok := mapOrigCopy[xhtml_h2_typeFrom]; ok {
		xhtml_h2_typeTo = _xhtml_h2_typeTo.(*Xhtml_h2_type)
		return
	}

	xhtml_h2_typeTo = new(Xhtml_h2_type)
	mapOrigCopy[xhtml_h2_typeFrom] = xhtml_h2_typeTo
	xhtml_h2_typeFrom.CopyBasicFields(xhtml_h2_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_h3_type(mapOrigCopy map[any]any, xhtml_h3_typeFrom *Xhtml_h3_type) (xhtml_h3_typeTo *Xhtml_h3_type) {

	// xhtml_h3_typeFrom has already been copied
	if _xhtml_h3_typeTo, ok := mapOrigCopy[xhtml_h3_typeFrom]; ok {
		xhtml_h3_typeTo = _xhtml_h3_typeTo.(*Xhtml_h3_type)
		return
	}

	xhtml_h3_typeTo = new(Xhtml_h3_type)
	mapOrigCopy[xhtml_h3_typeFrom] = xhtml_h3_typeTo
	xhtml_h3_typeFrom.CopyBasicFields(xhtml_h3_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_h4_type(mapOrigCopy map[any]any, xhtml_h4_typeFrom *Xhtml_h4_type) (xhtml_h4_typeTo *Xhtml_h4_type) {

	// xhtml_h4_typeFrom has already been copied
	if _xhtml_h4_typeTo, ok := mapOrigCopy[xhtml_h4_typeFrom]; ok {
		xhtml_h4_typeTo = _xhtml_h4_typeTo.(*Xhtml_h4_type)
		return
	}

	xhtml_h4_typeTo = new(Xhtml_h4_type)
	mapOrigCopy[xhtml_h4_typeFrom] = xhtml_h4_typeTo
	xhtml_h4_typeFrom.CopyBasicFields(xhtml_h4_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_h5_type(mapOrigCopy map[any]any, xhtml_h5_typeFrom *Xhtml_h5_type) (xhtml_h5_typeTo *Xhtml_h5_type) {

	// xhtml_h5_typeFrom has already been copied
	if _xhtml_h5_typeTo, ok := mapOrigCopy[xhtml_h5_typeFrom]; ok {
		xhtml_h5_typeTo = _xhtml_h5_typeTo.(*Xhtml_h5_type)
		return
	}

	xhtml_h5_typeTo = new(Xhtml_h5_type)
	mapOrigCopy[xhtml_h5_typeFrom] = xhtml_h5_typeTo
	xhtml_h5_typeFrom.CopyBasicFields(xhtml_h5_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_h6_type(mapOrigCopy map[any]any, xhtml_h6_typeFrom *Xhtml_h6_type) (xhtml_h6_typeTo *Xhtml_h6_type) {

	// xhtml_h6_typeFrom has already been copied
	if _xhtml_h6_typeTo, ok := mapOrigCopy[xhtml_h6_typeFrom]; ok {
		xhtml_h6_typeTo = _xhtml_h6_typeTo.(*Xhtml_h6_type)
		return
	}

	xhtml_h6_typeTo = new(Xhtml_h6_type)
	mapOrigCopy[xhtml_h6_typeFrom] = xhtml_h6_typeTo
	xhtml_h6_typeFrom.CopyBasicFields(xhtml_h6_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_heading_type(mapOrigCopy map[any]any, xhtml_heading_typeFrom *Xhtml_heading_type) (xhtml_heading_typeTo *Xhtml_heading_type) {

	// xhtml_heading_typeFrom has already been copied
	if _xhtml_heading_typeTo, ok := mapOrigCopy[xhtml_heading_typeFrom]; ok {
		xhtml_heading_typeTo = _xhtml_heading_typeTo.(*Xhtml_heading_type)
		return
	}

	xhtml_heading_typeTo = new(Xhtml_heading_type)
	mapOrigCopy[xhtml_heading_typeFrom] = xhtml_heading_typeTo
	xhtml_heading_typeFrom.CopyBasicFields(xhtml_heading_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_hr_type(mapOrigCopy map[any]any, xhtml_hr_typeFrom *Xhtml_hr_type) (xhtml_hr_typeTo *Xhtml_hr_type) {

	// xhtml_hr_typeFrom has already been copied
	if _xhtml_hr_typeTo, ok := mapOrigCopy[xhtml_hr_typeFrom]; ok {
		xhtml_hr_typeTo = _xhtml_hr_typeTo.(*Xhtml_hr_type)
		return
	}

	xhtml_hr_typeTo = new(Xhtml_hr_type)
	mapOrigCopy[xhtml_hr_typeFrom] = xhtml_hr_typeTo
	xhtml_hr_typeFrom.CopyBasicFields(xhtml_hr_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_kbd_type(mapOrigCopy map[any]any, xhtml_kbd_typeFrom *Xhtml_kbd_type) (xhtml_kbd_typeTo *Xhtml_kbd_type) {

	// xhtml_kbd_typeFrom has already been copied
	if _xhtml_kbd_typeTo, ok := mapOrigCopy[xhtml_kbd_typeFrom]; ok {
		xhtml_kbd_typeTo = _xhtml_kbd_typeTo.(*Xhtml_kbd_type)
		return
	}

	xhtml_kbd_typeTo = new(Xhtml_kbd_type)
	mapOrigCopy[xhtml_kbd_typeFrom] = xhtml_kbd_typeTo
	xhtml_kbd_typeFrom.CopyBasicFields(xhtml_kbd_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_li_type(mapOrigCopy map[any]any, xhtml_li_typeFrom *Xhtml_li_type) (xhtml_li_typeTo *Xhtml_li_type) {

	// xhtml_li_typeFrom has already been copied
	if _xhtml_li_typeTo, ok := mapOrigCopy[xhtml_li_typeFrom]; ok {
		xhtml_li_typeTo = _xhtml_li_typeTo.(*Xhtml_li_type)
		return
	}

	xhtml_li_typeTo = new(Xhtml_li_type)
	mapOrigCopy[xhtml_li_typeFrom] = xhtml_li_typeTo
	xhtml_li_typeFrom.CopyBasicFields(xhtml_li_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_object_type(mapOrigCopy map[any]any, xhtml_object_typeFrom *Xhtml_object_type) (xhtml_object_typeTo *Xhtml_object_type) {

	// xhtml_object_typeFrom has already been copied
	if _xhtml_object_typeTo, ok := mapOrigCopy[xhtml_object_typeFrom]; ok {
		xhtml_object_typeTo = _xhtml_object_typeTo.(*Xhtml_object_type)
		return
	}

	xhtml_object_typeTo = new(Xhtml_object_type)
	mapOrigCopy[xhtml_object_typeFrom] = xhtml_object_typeTo
	xhtml_object_typeFrom.CopyBasicFields(xhtml_object_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_ol_type(mapOrigCopy map[any]any, xhtml_ol_typeFrom *Xhtml_ol_type) (xhtml_ol_typeTo *Xhtml_ol_type) {

	// xhtml_ol_typeFrom has already been copied
	if _xhtml_ol_typeTo, ok := mapOrigCopy[xhtml_ol_typeFrom]; ok {
		xhtml_ol_typeTo = _xhtml_ol_typeTo.(*Xhtml_ol_type)
		return
	}

	xhtml_ol_typeTo = new(Xhtml_ol_type)
	mapOrigCopy[xhtml_ol_typeFrom] = xhtml_ol_typeTo
	xhtml_ol_typeFrom.CopyBasicFields(xhtml_ol_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_p_type(mapOrigCopy map[any]any, xhtml_p_typeFrom *Xhtml_p_type) (xhtml_p_typeTo *Xhtml_p_type) {

	// xhtml_p_typeFrom has already been copied
	if _xhtml_p_typeTo, ok := mapOrigCopy[xhtml_p_typeFrom]; ok {
		xhtml_p_typeTo = _xhtml_p_typeTo.(*Xhtml_p_type)
		return
	}

	xhtml_p_typeTo = new(Xhtml_p_type)
	mapOrigCopy[xhtml_p_typeFrom] = xhtml_p_typeTo
	xhtml_p_typeFrom.CopyBasicFields(xhtml_p_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_param_type(mapOrigCopy map[any]any, xhtml_param_typeFrom *Xhtml_param_type) (xhtml_param_typeTo *Xhtml_param_type) {

	// xhtml_param_typeFrom has already been copied
	if _xhtml_param_typeTo, ok := mapOrigCopy[xhtml_param_typeFrom]; ok {
		xhtml_param_typeTo = _xhtml_param_typeTo.(*Xhtml_param_type)
		return
	}

	xhtml_param_typeTo = new(Xhtml_param_type)
	mapOrigCopy[xhtml_param_typeFrom] = xhtml_param_typeTo
	xhtml_param_typeFrom.CopyBasicFields(xhtml_param_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_pre_type(mapOrigCopy map[any]any, xhtml_pre_typeFrom *Xhtml_pre_type) (xhtml_pre_typeTo *Xhtml_pre_type) {

	// xhtml_pre_typeFrom has already been copied
	if _xhtml_pre_typeTo, ok := mapOrigCopy[xhtml_pre_typeFrom]; ok {
		xhtml_pre_typeTo = _xhtml_pre_typeTo.(*Xhtml_pre_type)
		return
	}

	xhtml_pre_typeTo = new(Xhtml_pre_type)
	mapOrigCopy[xhtml_pre_typeFrom] = xhtml_pre_typeTo
	xhtml_pre_typeFrom.CopyBasicFields(xhtml_pre_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_q_type(mapOrigCopy map[any]any, xhtml_q_typeFrom *Xhtml_q_type) (xhtml_q_typeTo *Xhtml_q_type) {

	// xhtml_q_typeFrom has already been copied
	if _xhtml_q_typeTo, ok := mapOrigCopy[xhtml_q_typeFrom]; ok {
		xhtml_q_typeTo = _xhtml_q_typeTo.(*Xhtml_q_type)
		return
	}

	xhtml_q_typeTo = new(Xhtml_q_type)
	mapOrigCopy[xhtml_q_typeFrom] = xhtml_q_typeTo
	xhtml_q_typeFrom.CopyBasicFields(xhtml_q_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_samp_type(mapOrigCopy map[any]any, xhtml_samp_typeFrom *Xhtml_samp_type) (xhtml_samp_typeTo *Xhtml_samp_type) {

	// xhtml_samp_typeFrom has already been copied
	if _xhtml_samp_typeTo, ok := mapOrigCopy[xhtml_samp_typeFrom]; ok {
		xhtml_samp_typeTo = _xhtml_samp_typeTo.(*Xhtml_samp_type)
		return
	}

	xhtml_samp_typeTo = new(Xhtml_samp_type)
	mapOrigCopy[xhtml_samp_typeFrom] = xhtml_samp_typeTo
	xhtml_samp_typeFrom.CopyBasicFields(xhtml_samp_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_span_type(mapOrigCopy map[any]any, xhtml_span_typeFrom *Xhtml_span_type) (xhtml_span_typeTo *Xhtml_span_type) {

	// xhtml_span_typeFrom has already been copied
	if _xhtml_span_typeTo, ok := mapOrigCopy[xhtml_span_typeFrom]; ok {
		xhtml_span_typeTo = _xhtml_span_typeTo.(*Xhtml_span_type)
		return
	}

	xhtml_span_typeTo = new(Xhtml_span_type)
	mapOrigCopy[xhtml_span_typeFrom] = xhtml_span_typeTo
	xhtml_span_typeFrom.CopyBasicFields(xhtml_span_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_strong_type(mapOrigCopy map[any]any, xhtml_strong_typeFrom *Xhtml_strong_type) (xhtml_strong_typeTo *Xhtml_strong_type) {

	// xhtml_strong_typeFrom has already been copied
	if _xhtml_strong_typeTo, ok := mapOrigCopy[xhtml_strong_typeFrom]; ok {
		xhtml_strong_typeTo = _xhtml_strong_typeTo.(*Xhtml_strong_type)
		return
	}

	xhtml_strong_typeTo = new(Xhtml_strong_type)
	mapOrigCopy[xhtml_strong_typeFrom] = xhtml_strong_typeTo
	xhtml_strong_typeFrom.CopyBasicFields(xhtml_strong_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_table_type(mapOrigCopy map[any]any, xhtml_table_typeFrom *Xhtml_table_type) (xhtml_table_typeTo *Xhtml_table_type) {

	// xhtml_table_typeFrom has already been copied
	if _xhtml_table_typeTo, ok := mapOrigCopy[xhtml_table_typeFrom]; ok {
		xhtml_table_typeTo = _xhtml_table_typeTo.(*Xhtml_table_type)
		return
	}

	xhtml_table_typeTo = new(Xhtml_table_type)
	mapOrigCopy[xhtml_table_typeFrom] = xhtml_table_typeTo
	xhtml_table_typeFrom.CopyBasicFields(xhtml_table_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_tbody_type(mapOrigCopy map[any]any, xhtml_tbody_typeFrom *Xhtml_tbody_type) (xhtml_tbody_typeTo *Xhtml_tbody_type) {

	// xhtml_tbody_typeFrom has already been copied
	if _xhtml_tbody_typeTo, ok := mapOrigCopy[xhtml_tbody_typeFrom]; ok {
		xhtml_tbody_typeTo = _xhtml_tbody_typeTo.(*Xhtml_tbody_type)
		return
	}

	xhtml_tbody_typeTo = new(Xhtml_tbody_type)
	mapOrigCopy[xhtml_tbody_typeFrom] = xhtml_tbody_typeTo
	xhtml_tbody_typeFrom.CopyBasicFields(xhtml_tbody_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_td_type(mapOrigCopy map[any]any, xhtml_td_typeFrom *Xhtml_td_type) (xhtml_td_typeTo *Xhtml_td_type) {

	// xhtml_td_typeFrom has already been copied
	if _xhtml_td_typeTo, ok := mapOrigCopy[xhtml_td_typeFrom]; ok {
		xhtml_td_typeTo = _xhtml_td_typeTo.(*Xhtml_td_type)
		return
	}

	xhtml_td_typeTo = new(Xhtml_td_type)
	mapOrigCopy[xhtml_td_typeFrom] = xhtml_td_typeTo
	xhtml_td_typeFrom.CopyBasicFields(xhtml_td_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_tfoot_type(mapOrigCopy map[any]any, xhtml_tfoot_typeFrom *Xhtml_tfoot_type) (xhtml_tfoot_typeTo *Xhtml_tfoot_type) {

	// xhtml_tfoot_typeFrom has already been copied
	if _xhtml_tfoot_typeTo, ok := mapOrigCopy[xhtml_tfoot_typeFrom]; ok {
		xhtml_tfoot_typeTo = _xhtml_tfoot_typeTo.(*Xhtml_tfoot_type)
		return
	}

	xhtml_tfoot_typeTo = new(Xhtml_tfoot_type)
	mapOrigCopy[xhtml_tfoot_typeFrom] = xhtml_tfoot_typeTo
	xhtml_tfoot_typeFrom.CopyBasicFields(xhtml_tfoot_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_th_type(mapOrigCopy map[any]any, xhtml_th_typeFrom *Xhtml_th_type) (xhtml_th_typeTo *Xhtml_th_type) {

	// xhtml_th_typeFrom has already been copied
	if _xhtml_th_typeTo, ok := mapOrigCopy[xhtml_th_typeFrom]; ok {
		xhtml_th_typeTo = _xhtml_th_typeTo.(*Xhtml_th_type)
		return
	}

	xhtml_th_typeTo = new(Xhtml_th_type)
	mapOrigCopy[xhtml_th_typeFrom] = xhtml_th_typeTo
	xhtml_th_typeFrom.CopyBasicFields(xhtml_th_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_thead_type(mapOrigCopy map[any]any, xhtml_thead_typeFrom *Xhtml_thead_type) (xhtml_thead_typeTo *Xhtml_thead_type) {

	// xhtml_thead_typeFrom has already been copied
	if _xhtml_thead_typeTo, ok := mapOrigCopy[xhtml_thead_typeFrom]; ok {
		xhtml_thead_typeTo = _xhtml_thead_typeTo.(*Xhtml_thead_type)
		return
	}

	xhtml_thead_typeTo = new(Xhtml_thead_type)
	mapOrigCopy[xhtml_thead_typeFrom] = xhtml_thead_typeTo
	xhtml_thead_typeFrom.CopyBasicFields(xhtml_thead_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_tr_type(mapOrigCopy map[any]any, xhtml_tr_typeFrom *Xhtml_tr_type) (xhtml_tr_typeTo *Xhtml_tr_type) {

	// xhtml_tr_typeFrom has already been copied
	if _xhtml_tr_typeTo, ok := mapOrigCopy[xhtml_tr_typeFrom]; ok {
		xhtml_tr_typeTo = _xhtml_tr_typeTo.(*Xhtml_tr_type)
		return
	}

	xhtml_tr_typeTo = new(Xhtml_tr_type)
	mapOrigCopy[xhtml_tr_typeFrom] = xhtml_tr_typeTo
	xhtml_tr_typeFrom.CopyBasicFields(xhtml_tr_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_ul_type(mapOrigCopy map[any]any, xhtml_ul_typeFrom *Xhtml_ul_type) (xhtml_ul_typeTo *Xhtml_ul_type) {

	// xhtml_ul_typeFrom has already been copied
	if _xhtml_ul_typeTo, ok := mapOrigCopy[xhtml_ul_typeFrom]; ok {
		xhtml_ul_typeTo = _xhtml_ul_typeTo.(*Xhtml_ul_type)
		return
	}

	xhtml_ul_typeTo = new(Xhtml_ul_type)
	mapOrigCopy[xhtml_ul_typeFrom] = xhtml_ul_typeTo
	xhtml_ul_typeFrom.CopyBasicFields(xhtml_ul_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchXhtml_var_type(mapOrigCopy map[any]any, xhtml_var_typeFrom *Xhtml_var_type) (xhtml_var_typeTo *Xhtml_var_type) {

	// xhtml_var_typeFrom has already been copied
	if _xhtml_var_typeTo, ok := mapOrigCopy[xhtml_var_typeFrom]; ok {
		xhtml_var_typeTo = _xhtml_var_typeTo.(*Xhtml_var_type)
		return
	}

	xhtml_var_typeTo = new(Xhtml_var_type)
	mapOrigCopy[xhtml_var_typeFrom] = xhtml_var_typeTo
	xhtml_var_typeFrom.CopyBasicFields(xhtml_var_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *ALTERNATIVE_ID:
		stage.UnstageBranchALTERNATIVE_ID(target)

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.UnstageBranchATTRIBUTE_DEFINITION_BOOLEAN(target)

	case *ATTRIBUTE_DEFINITION_DATE:
		stage.UnstageBranchATTRIBUTE_DEFINITION_DATE(target)

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.UnstageBranchATTRIBUTE_DEFINITION_ENUMERATION(target)

	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.UnstageBranchATTRIBUTE_DEFINITION_INTEGER(target)

	case *ATTRIBUTE_DEFINITION_REAL:
		stage.UnstageBranchATTRIBUTE_DEFINITION_REAL(target)

	case *ATTRIBUTE_DEFINITION_STRING:
		stage.UnstageBranchATTRIBUTE_DEFINITION_STRING(target)

	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.UnstageBranchATTRIBUTE_DEFINITION_XHTML(target)

	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.UnstageBranchATTRIBUTE_VALUE_BOOLEAN(target)

	case *ATTRIBUTE_VALUE_DATE:
		stage.UnstageBranchATTRIBUTE_VALUE_DATE(target)

	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.UnstageBranchATTRIBUTE_VALUE_ENUMERATION(target)

	case *ATTRIBUTE_VALUE_INTEGER:
		stage.UnstageBranchATTRIBUTE_VALUE_INTEGER(target)

	case *ATTRIBUTE_VALUE_REAL:
		stage.UnstageBranchATTRIBUTE_VALUE_REAL(target)

	case *ATTRIBUTE_VALUE_STRING:
		stage.UnstageBranchATTRIBUTE_VALUE_STRING(target)

	case *ATTRIBUTE_VALUE_XHTML:
		stage.UnstageBranchATTRIBUTE_VALUE_XHTML(target)

	case *AnyType:
		stage.UnstageBranchAnyType(target)

	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.UnstageBranchDATATYPE_DEFINITION_BOOLEAN(target)

	case *DATATYPE_DEFINITION_DATE:
		stage.UnstageBranchDATATYPE_DEFINITION_DATE(target)

	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.UnstageBranchDATATYPE_DEFINITION_ENUMERATION(target)

	case *DATATYPE_DEFINITION_INTEGER:
		stage.UnstageBranchDATATYPE_DEFINITION_INTEGER(target)

	case *DATATYPE_DEFINITION_REAL:
		stage.UnstageBranchDATATYPE_DEFINITION_REAL(target)

	case *DATATYPE_DEFINITION_STRING:
		stage.UnstageBranchDATATYPE_DEFINITION_STRING(target)

	case *DATATYPE_DEFINITION_XHTML:
		stage.UnstageBranchDATATYPE_DEFINITION_XHTML(target)

	case *EMBEDDED_VALUE:
		stage.UnstageBranchEMBEDDED_VALUE(target)

	case *ENUM_VALUE:
		stage.UnstageBranchENUM_VALUE(target)

	case *RELATION_GROUP:
		stage.UnstageBranchRELATION_GROUP(target)

	case *RELATION_GROUP_TYPE:
		stage.UnstageBranchRELATION_GROUP_TYPE(target)

	case *REQ_IF:
		stage.UnstageBranchREQ_IF(target)

	case *REQ_IF_CONTENT:
		stage.UnstageBranchREQ_IF_CONTENT(target)

	case *REQ_IF_HEADER:
		stage.UnstageBranchREQ_IF_HEADER(target)

	case *REQ_IF_TOOL_EXTENSION:
		stage.UnstageBranchREQ_IF_TOOL_EXTENSION(target)

	case *SPECIFICATION:
		stage.UnstageBranchSPECIFICATION(target)

	case *SPECIFICATION_TYPE:
		stage.UnstageBranchSPECIFICATION_TYPE(target)

	case *SPEC_HIERARCHY:
		stage.UnstageBranchSPEC_HIERARCHY(target)

	case *SPEC_OBJECT:
		stage.UnstageBranchSPEC_OBJECT(target)

	case *SPEC_OBJECT_TYPE:
		stage.UnstageBranchSPEC_OBJECT_TYPE(target)

	case *SPEC_RELATION:
		stage.UnstageBranchSPEC_RELATION(target)

	case *SPEC_RELATION_TYPE:
		stage.UnstageBranchSPEC_RELATION_TYPE(target)

	case *XHTML_CONTENT:
		stage.UnstageBranchXHTML_CONTENT(target)

	case *Xhtml_InlPres_type:
		stage.UnstageBranchXhtml_InlPres_type(target)

	case *Xhtml_a_type:
		stage.UnstageBranchXhtml_a_type(target)

	case *Xhtml_abbr_type:
		stage.UnstageBranchXhtml_abbr_type(target)

	case *Xhtml_acronym_type:
		stage.UnstageBranchXhtml_acronym_type(target)

	case *Xhtml_address_type:
		stage.UnstageBranchXhtml_address_type(target)

	case *Xhtml_blockquote_type:
		stage.UnstageBranchXhtml_blockquote_type(target)

	case *Xhtml_br_type:
		stage.UnstageBranchXhtml_br_type(target)

	case *Xhtml_caption_type:
		stage.UnstageBranchXhtml_caption_type(target)

	case *Xhtml_cite_type:
		stage.UnstageBranchXhtml_cite_type(target)

	case *Xhtml_code_type:
		stage.UnstageBranchXhtml_code_type(target)

	case *Xhtml_col_type:
		stage.UnstageBranchXhtml_col_type(target)

	case *Xhtml_colgroup_type:
		stage.UnstageBranchXhtml_colgroup_type(target)

	case *Xhtml_dd_type:
		stage.UnstageBranchXhtml_dd_type(target)

	case *Xhtml_dfn_type:
		stage.UnstageBranchXhtml_dfn_type(target)

	case *Xhtml_div_type:
		stage.UnstageBranchXhtml_div_type(target)

	case *Xhtml_dl_type:
		stage.UnstageBranchXhtml_dl_type(target)

	case *Xhtml_dt_type:
		stage.UnstageBranchXhtml_dt_type(target)

	case *Xhtml_edit_type:
		stage.UnstageBranchXhtml_edit_type(target)

	case *Xhtml_em_type:
		stage.UnstageBranchXhtml_em_type(target)

	case *Xhtml_h1_type:
		stage.UnstageBranchXhtml_h1_type(target)

	case *Xhtml_h2_type:
		stage.UnstageBranchXhtml_h2_type(target)

	case *Xhtml_h3_type:
		stage.UnstageBranchXhtml_h3_type(target)

	case *Xhtml_h4_type:
		stage.UnstageBranchXhtml_h4_type(target)

	case *Xhtml_h5_type:
		stage.UnstageBranchXhtml_h5_type(target)

	case *Xhtml_h6_type:
		stage.UnstageBranchXhtml_h6_type(target)

	case *Xhtml_heading_type:
		stage.UnstageBranchXhtml_heading_type(target)

	case *Xhtml_hr_type:
		stage.UnstageBranchXhtml_hr_type(target)

	case *Xhtml_kbd_type:
		stage.UnstageBranchXhtml_kbd_type(target)

	case *Xhtml_li_type:
		stage.UnstageBranchXhtml_li_type(target)

	case *Xhtml_object_type:
		stage.UnstageBranchXhtml_object_type(target)

	case *Xhtml_ol_type:
		stage.UnstageBranchXhtml_ol_type(target)

	case *Xhtml_p_type:
		stage.UnstageBranchXhtml_p_type(target)

	case *Xhtml_param_type:
		stage.UnstageBranchXhtml_param_type(target)

	case *Xhtml_pre_type:
		stage.UnstageBranchXhtml_pre_type(target)

	case *Xhtml_q_type:
		stage.UnstageBranchXhtml_q_type(target)

	case *Xhtml_samp_type:
		stage.UnstageBranchXhtml_samp_type(target)

	case *Xhtml_span_type:
		stage.UnstageBranchXhtml_span_type(target)

	case *Xhtml_strong_type:
		stage.UnstageBranchXhtml_strong_type(target)

	case *Xhtml_table_type:
		stage.UnstageBranchXhtml_table_type(target)

	case *Xhtml_tbody_type:
		stage.UnstageBranchXhtml_tbody_type(target)

	case *Xhtml_td_type:
		stage.UnstageBranchXhtml_td_type(target)

	case *Xhtml_tfoot_type:
		stage.UnstageBranchXhtml_tfoot_type(target)

	case *Xhtml_th_type:
		stage.UnstageBranchXhtml_th_type(target)

	case *Xhtml_thead_type:
		stage.UnstageBranchXhtml_thead_type(target)

	case *Xhtml_tr_type:
		stage.UnstageBranchXhtml_tr_type(target)

	case *Xhtml_ul_type:
		stage.UnstageBranchXhtml_ul_type(target)

	case *Xhtml_var_type:
		stage.UnstageBranchXhtml_var_type(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID) {

	// check if instance is already staged
	if !IsStaged(stage, alternative_id) {
		return
	}

	alternative_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_boolean) {
		return
	}

	attribute_definition_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_boolean := range attribute_definition_boolean.DEFAULT_VALUE.ATTRIBUTE_VALUE_BOOLEAN {
		UnstageBranch(stage, _attribute_value_boolean)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_date) {
		return
	}

	attribute_definition_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_date := range attribute_definition_date.DEFAULT_VALUE.ATTRIBUTE_VALUE_DATE {
		UnstageBranch(stage, _attribute_value_date)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_enumeration) {
		return
	}

	attribute_definition_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_enumeration := range attribute_definition_enumeration.DEFAULT_VALUE.ATTRIBUTE_VALUE_ENUMERATION {
		UnstageBranch(stage, _attribute_value_enumeration)
	}
	for _, _alternative_id := range attribute_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_integer) {
		return
	}

	attribute_definition_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_integer := range attribute_definition_integer.DEFAULT_VALUE.ATTRIBUTE_VALUE_INTEGER {
		UnstageBranch(stage, _attribute_value_integer)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_real) {
		return
	}

	attribute_definition_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_real := range attribute_definition_real.DEFAULT_VALUE.ATTRIBUTE_VALUE_REAL {
		UnstageBranch(stage, _attribute_value_real)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_string) {
		return
	}

	attribute_definition_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_string := range attribute_definition_string.DEFAULT_VALUE.ATTRIBUTE_VALUE_STRING {
		UnstageBranch(stage, _attribute_value_string)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_definition_xhtml) {
		return
	}

	attribute_definition_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range attribute_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_xhtml := range attribute_definition_xhtml.DEFAULT_VALUE.ATTRIBUTE_VALUE_XHTML {
		UnstageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_boolean) {
		return
	}

	attribute_value_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_date) {
		return
	}

	attribute_value_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_enumeration) {
		return
	}

	attribute_value_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_integer) {
		return
	}

	attribute_value_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_real) {
		return
	}

	attribute_value_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_string) {
		return
	}

	attribute_value_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, attribute_value_xhtml) {
		return
	}

	attribute_value_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _xhtml_content := range attribute_value_xhtml.THE_VALUE {
		UnstageBranch(stage, _xhtml_content)
	}
	for _, _xhtml_content := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
		UnstageBranch(stage, _xhtml_content)
	}

}

func (stage *StageStruct) UnstageBranchAnyType(anytype *AnyType) {

	// check if instance is already staged
	if !IsStaged(stage, anytype) {
		return
	}

	anytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_boolean) {
		return
	}

	datatype_definition_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_boolean.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_date) {
		return
	}

	datatype_definition_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_date.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_enumeration) {
		return
	}

	datatype_definition_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_enumeration.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _enum_value := range datatype_definition_enumeration.SPECIFIED_VALUES.ENUM_VALUE {
		UnstageBranch(stage, _enum_value)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_integer) {
		return
	}

	datatype_definition_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_integer.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_real) {
		return
	}

	datatype_definition_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_real.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_string) {
		return
	}

	datatype_definition_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_string.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) {

	// check if instance is already staged
	if !IsStaged(stage, datatype_definition_xhtml) {
		return
	}

	datatype_definition_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range datatype_definition_xhtml.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE) {

	// check if instance is already staged
	if !IsStaged(stage, embedded_value) {
		return
	}

	embedded_value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchENUM_VALUE(enum_value *ENUM_VALUE) {

	// check if instance is already staged
	if !IsStaged(stage, enum_value) {
		return
	}

	enum_value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range enum_value.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _embedded_value := range enum_value.PROPERTIES.EMBEDDED_VALUE {
		UnstageBranch(stage, _embedded_value)
	}

}

func (stage *StageStruct) UnstageBranchRELATION_GROUP(relation_group *RELATION_GROUP) {

	// check if instance is already staged
	if !IsStaged(stage, relation_group) {
		return
	}

	relation_group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range relation_group.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}

}

func (stage *StageStruct) UnstageBranchRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, relation_group_type) {
		return
	}

	relation_group_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range relation_group_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_definition_boolean := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		UnstageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		UnstageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		UnstageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		UnstageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		UnstageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		UnstageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range relation_group_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		UnstageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchREQ_IF(req_if *REQ_IF) {

	// check if instance is already staged
	if !IsStaged(stage, req_if) {
		return
	}

	req_if.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if.THE_HEADER.REQ_IF_HEADER != nil {
		UnstageBranch(stage, req_if.THE_HEADER.REQ_IF_HEADER)
	}
	if req_if.CORE_CONTENT.REQ_IF_CONTENT != nil {
		UnstageBranch(stage, req_if.CORE_CONTENT.REQ_IF_CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_tool_extension := range req_if.TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION {
		UnstageBranch(stage, _req_if_tool_extension)
	}

}

func (stage *StageStruct) UnstageBranchREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, req_if_content) {
		return
	}

	req_if_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatype_definition_boolean := range req_if_content.DATATYPES.DATATYPE_DEFINITION_BOOLEAN {
		UnstageBranch(stage, _datatype_definition_boolean)
	}
	for _, _datatype_definition_date := range req_if_content.DATATYPES.DATATYPE_DEFINITION_DATE {
		UnstageBranch(stage, _datatype_definition_date)
	}
	for _, _datatype_definition_enumeration := range req_if_content.DATATYPES.DATATYPE_DEFINITION_ENUMERATION {
		UnstageBranch(stage, _datatype_definition_enumeration)
	}
	for _, _datatype_definition_integer := range req_if_content.DATATYPES.DATATYPE_DEFINITION_INTEGER {
		UnstageBranch(stage, _datatype_definition_integer)
	}
	for _, _datatype_definition_real := range req_if_content.DATATYPES.DATATYPE_DEFINITION_REAL {
		UnstageBranch(stage, _datatype_definition_real)
	}
	for _, _datatype_definition_string := range req_if_content.DATATYPES.DATATYPE_DEFINITION_STRING {
		UnstageBranch(stage, _datatype_definition_string)
	}
	for _, _datatype_definition_xhtml := range req_if_content.DATATYPES.DATATYPE_DEFINITION_XHTML {
		UnstageBranch(stage, _datatype_definition_xhtml)
	}
	for _, _relation_group_type := range req_if_content.SPEC_TYPES.RELATION_GROUP_TYPE {
		UnstageBranch(stage, _relation_group_type)
	}
	for _, _spec_object_type := range req_if_content.SPEC_TYPES.SPEC_OBJECT_TYPE {
		UnstageBranch(stage, _spec_object_type)
	}
	for _, _spec_relation_type := range req_if_content.SPEC_TYPES.SPEC_RELATION_TYPE {
		UnstageBranch(stage, _spec_relation_type)
	}
	for _, _specification_type := range req_if_content.SPEC_TYPES.SPECIFICATION_TYPE {
		UnstageBranch(stage, _specification_type)
	}
	for _, _spec_object := range req_if_content.SPEC_OBJECTS.SPEC_OBJECT {
		UnstageBranch(stage, _spec_object)
	}
	for _, _spec_relation := range req_if_content.SPEC_RELATIONS.SPEC_RELATION {
		UnstageBranch(stage, _spec_relation)
	}
	for _, _specification := range req_if_content.SPECIFICATIONS.SPECIFICATION {
		UnstageBranch(stage, _specification)
	}
	for _, _relation_group := range req_if_content.SPEC_RELATION_GROUPS.RELATION_GROUP {
		UnstageBranch(stage, _relation_group)
	}

}

func (stage *StageStruct) UnstageBranchREQ_IF_HEADER(req_if_header *REQ_IF_HEADER) {

	// check if instance is already staged
	if !IsStaged(stage, req_if_header) {
		return
	}

	req_if_header.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION) {

	// check if instance is already staged
	if !IsStaged(stage, req_if_tool_extension) {
		return
	}

	req_if_tool_extension.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchSPECIFICATION(specification *SPECIFICATION) {

	// check if instance is already staged
	if !IsStaged(stage, specification) {
		return
	}

	specification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range specification.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_boolean := range specification.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		UnstageBranch(stage, _attribute_value_boolean)
	}
	for _, _attribute_value_date := range specification.VALUES.ATTRIBUTE_VALUE_DATE {
		UnstageBranch(stage, _attribute_value_date)
	}
	for _, _attribute_value_enumeration := range specification.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		UnstageBranch(stage, _attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range specification.VALUES.ATTRIBUTE_VALUE_INTEGER {
		UnstageBranch(stage, _attribute_value_integer)
	}
	for _, _attribute_value_real := range specification.VALUES.ATTRIBUTE_VALUE_REAL {
		UnstageBranch(stage, _attribute_value_real)
	}
	for _, _attribute_value_string := range specification.VALUES.ATTRIBUTE_VALUE_STRING {
		UnstageBranch(stage, _attribute_value_string)
	}
	for _, _attribute_value_xhtml := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
		UnstageBranch(stage, _attribute_value_xhtml)
	}
	for _, _spec_hierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
		UnstageBranch(stage, _spec_hierarchy)
	}

}

func (stage *StageStruct) UnstageBranchSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, specification_type) {
		return
	}

	specification_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range specification_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_definition_boolean := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		UnstageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		UnstageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		UnstageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		UnstageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		UnstageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		UnstageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range specification_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		UnstageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) {

	// check if instance is already staged
	if !IsStaged(stage, spec_hierarchy) {
		return
	}

	spec_hierarchy.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_hierarchy.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _spec_hierarchy := range spec_hierarchy.CHILDREN.SPEC_HIERARCHY {
		UnstageBranch(stage, _spec_hierarchy)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_OBJECT(spec_object *SPEC_OBJECT) {

	// check if instance is already staged
	if !IsStaged(stage, spec_object) {
		return
	}

	spec_object.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_object.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_boolean := range spec_object.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		UnstageBranch(stage, _attribute_value_boolean)
	}
	for _, _attribute_value_date := range spec_object.VALUES.ATTRIBUTE_VALUE_DATE {
		UnstageBranch(stage, _attribute_value_date)
	}
	for _, _attribute_value_enumeration := range spec_object.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		UnstageBranch(stage, _attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range spec_object.VALUES.ATTRIBUTE_VALUE_INTEGER {
		UnstageBranch(stage, _attribute_value_integer)
	}
	for _, _attribute_value_real := range spec_object.VALUES.ATTRIBUTE_VALUE_REAL {
		UnstageBranch(stage, _attribute_value_real)
	}
	for _, _attribute_value_string := range spec_object.VALUES.ATTRIBUTE_VALUE_STRING {
		UnstageBranch(stage, _attribute_value_string)
	}
	for _, _attribute_value_xhtml := range spec_object.VALUES.ATTRIBUTE_VALUE_XHTML {
		UnstageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, spec_object_type) {
		return
	}

	spec_object_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_object_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_definition_boolean := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		UnstageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		UnstageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		UnstageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		UnstageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		UnstageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		UnstageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range spec_object_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		UnstageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_RELATION(spec_relation *SPEC_RELATION) {

	// check if instance is already staged
	if !IsStaged(stage, spec_relation) {
		return
	}

	spec_relation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_relation.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_value_boolean := range spec_relation.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		UnstageBranch(stage, _attribute_value_boolean)
	}
	for _, _attribute_value_date := range spec_relation.VALUES.ATTRIBUTE_VALUE_DATE {
		UnstageBranch(stage, _attribute_value_date)
	}
	for _, _attribute_value_enumeration := range spec_relation.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		UnstageBranch(stage, _attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range spec_relation.VALUES.ATTRIBUTE_VALUE_INTEGER {
		UnstageBranch(stage, _attribute_value_integer)
	}
	for _, _attribute_value_real := range spec_relation.VALUES.ATTRIBUTE_VALUE_REAL {
		UnstageBranch(stage, _attribute_value_real)
	}
	for _, _attribute_value_string := range spec_relation.VALUES.ATTRIBUTE_VALUE_STRING {
		UnstageBranch(stage, _attribute_value_string)
	}
	for _, _attribute_value_xhtml := range spec_relation.VALUES.ATTRIBUTE_VALUE_XHTML {
		UnstageBranch(stage, _attribute_value_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) {

	// check if instance is already staged
	if !IsStaged(stage, spec_relation_type) {
		return
	}

	spec_relation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _alternative_id := range spec_relation_type.ALTERNATIVE_ID.ALTERNATIVE_ID {
		UnstageBranch(stage, _alternative_id)
	}
	for _, _attribute_definition_boolean := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
		UnstageBranch(stage, _attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
		UnstageBranch(stage, _attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
		UnstageBranch(stage, _attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
		UnstageBranch(stage, _attribute_definition_integer)
	}
	for _, _attribute_definition_real := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
		UnstageBranch(stage, _attribute_definition_real)
	}
	for _, _attribute_definition_string := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
		UnstageBranch(stage, _attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range spec_relation_type.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		UnstageBranch(stage, _attribute_definition_xhtml)
	}

}

func (stage *StageStruct) UnstageBranchXHTML_CONTENT(xhtml_content *XHTML_CONTENT) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_content) {
		return
	}

	xhtml_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_InlPres_type(xhtml_inlpres_type *Xhtml_InlPres_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_inlpres_type) {
		return
	}

	xhtml_inlpres_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_a_type(xhtml_a_type *Xhtml_a_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_a_type) {
		return
	}

	xhtml_a_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_abbr_type(xhtml_abbr_type *Xhtml_abbr_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_abbr_type) {
		return
	}

	xhtml_abbr_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_acronym_type(xhtml_acronym_type *Xhtml_acronym_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_acronym_type) {
		return
	}

	xhtml_acronym_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_address_type(xhtml_address_type *Xhtml_address_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_address_type) {
		return
	}

	xhtml_address_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_blockquote_type(xhtml_blockquote_type *Xhtml_blockquote_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_blockquote_type) {
		return
	}

	xhtml_blockquote_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_br_type(xhtml_br_type *Xhtml_br_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_br_type) {
		return
	}

	xhtml_br_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_caption_type(xhtml_caption_type *Xhtml_caption_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_caption_type) {
		return
	}

	xhtml_caption_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_cite_type(xhtml_cite_type *Xhtml_cite_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_cite_type) {
		return
	}

	xhtml_cite_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_code_type(xhtml_code_type *Xhtml_code_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_code_type) {
		return
	}

	xhtml_code_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_col_type(xhtml_col_type *Xhtml_col_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_col_type) {
		return
	}

	xhtml_col_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_colgroup_type(xhtml_colgroup_type *Xhtml_colgroup_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_colgroup_type) {
		return
	}

	xhtml_colgroup_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_dd_type(xhtml_dd_type *Xhtml_dd_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_dd_type) {
		return
	}

	xhtml_dd_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_dfn_type(xhtml_dfn_type *Xhtml_dfn_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_dfn_type) {
		return
	}

	xhtml_dfn_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_div_type(xhtml_div_type *Xhtml_div_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_div_type) {
		return
	}

	xhtml_div_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_dl_type(xhtml_dl_type *Xhtml_dl_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_dl_type) {
		return
	}

	xhtml_dl_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_dt_type(xhtml_dt_type *Xhtml_dt_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_dt_type) {
		return
	}

	xhtml_dt_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_edit_type(xhtml_edit_type *Xhtml_edit_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_edit_type) {
		return
	}

	xhtml_edit_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_em_type(xhtml_em_type *Xhtml_em_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_em_type) {
		return
	}

	xhtml_em_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_h1_type(xhtml_h1_type *Xhtml_h1_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_h1_type) {
		return
	}

	xhtml_h1_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_h2_type(xhtml_h2_type *Xhtml_h2_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_h2_type) {
		return
	}

	xhtml_h2_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_h3_type(xhtml_h3_type *Xhtml_h3_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_h3_type) {
		return
	}

	xhtml_h3_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_h4_type(xhtml_h4_type *Xhtml_h4_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_h4_type) {
		return
	}

	xhtml_h4_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_h5_type(xhtml_h5_type *Xhtml_h5_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_h5_type) {
		return
	}

	xhtml_h5_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_h6_type(xhtml_h6_type *Xhtml_h6_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_h6_type) {
		return
	}

	xhtml_h6_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_heading_type(xhtml_heading_type *Xhtml_heading_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_heading_type) {
		return
	}

	xhtml_heading_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_hr_type(xhtml_hr_type *Xhtml_hr_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_hr_type) {
		return
	}

	xhtml_hr_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_kbd_type(xhtml_kbd_type *Xhtml_kbd_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_kbd_type) {
		return
	}

	xhtml_kbd_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_li_type(xhtml_li_type *Xhtml_li_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_li_type) {
		return
	}

	xhtml_li_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_object_type(xhtml_object_type *Xhtml_object_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_object_type) {
		return
	}

	xhtml_object_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_ol_type(xhtml_ol_type *Xhtml_ol_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_ol_type) {
		return
	}

	xhtml_ol_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_p_type(xhtml_p_type *Xhtml_p_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_p_type) {
		return
	}

	xhtml_p_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_param_type(xhtml_param_type *Xhtml_param_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_param_type) {
		return
	}

	xhtml_param_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_pre_type(xhtml_pre_type *Xhtml_pre_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_pre_type) {
		return
	}

	xhtml_pre_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_q_type(xhtml_q_type *Xhtml_q_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_q_type) {
		return
	}

	xhtml_q_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_samp_type(xhtml_samp_type *Xhtml_samp_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_samp_type) {
		return
	}

	xhtml_samp_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_span_type(xhtml_span_type *Xhtml_span_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_span_type) {
		return
	}

	xhtml_span_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_strong_type(xhtml_strong_type *Xhtml_strong_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_strong_type) {
		return
	}

	xhtml_strong_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_table_type(xhtml_table_type *Xhtml_table_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_table_type) {
		return
	}

	xhtml_table_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_tbody_type(xhtml_tbody_type *Xhtml_tbody_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_tbody_type) {
		return
	}

	xhtml_tbody_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_td_type(xhtml_td_type *Xhtml_td_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_td_type) {
		return
	}

	xhtml_td_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_tfoot_type(xhtml_tfoot_type *Xhtml_tfoot_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_tfoot_type) {
		return
	}

	xhtml_tfoot_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_th_type(xhtml_th_type *Xhtml_th_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_th_type) {
		return
	}

	xhtml_th_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_thead_type(xhtml_thead_type *Xhtml_thead_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_thead_type) {
		return
	}

	xhtml_thead_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_tr_type(xhtml_tr_type *Xhtml_tr_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_tr_type) {
		return
	}

	xhtml_tr_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_ul_type(xhtml_ul_type *Xhtml_ul_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_ul_type) {
		return
	}

	xhtml_ul_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchXhtml_var_type(xhtml_var_type *Xhtml_var_type) {

	// check if instance is already staged
	if !IsStaged(stage, xhtml_var_type) {
		return
	}

	xhtml_var_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
