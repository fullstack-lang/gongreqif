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
	case *AnyType:
		if stage.OnAfterAnyTypeCreateCallback != nil {
			stage.OnAfterAnyTypeCreateCallback.OnAfterCreate(stage, target)
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
	case *XHTML_CONTENT:
		if stage.OnAfterXHTML_CONTENTCreateCallback != nil {
			stage.OnAfterXHTML_CONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_InlPres_type:
		if stage.OnAfterXhtml_InlPres_typeCreateCallback != nil {
			stage.OnAfterXhtml_InlPres_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_a_type:
		if stage.OnAfterXhtml_a_typeCreateCallback != nil {
			stage.OnAfterXhtml_a_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_abbr_type:
		if stage.OnAfterXhtml_abbr_typeCreateCallback != nil {
			stage.OnAfterXhtml_abbr_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_acronym_type:
		if stage.OnAfterXhtml_acronym_typeCreateCallback != nil {
			stage.OnAfterXhtml_acronym_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_address_type:
		if stage.OnAfterXhtml_address_typeCreateCallback != nil {
			stage.OnAfterXhtml_address_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_blockquote_type:
		if stage.OnAfterXhtml_blockquote_typeCreateCallback != nil {
			stage.OnAfterXhtml_blockquote_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_br_type:
		if stage.OnAfterXhtml_br_typeCreateCallback != nil {
			stage.OnAfterXhtml_br_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_caption_type:
		if stage.OnAfterXhtml_caption_typeCreateCallback != nil {
			stage.OnAfterXhtml_caption_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_cite_type:
		if stage.OnAfterXhtml_cite_typeCreateCallback != nil {
			stage.OnAfterXhtml_cite_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_code_type:
		if stage.OnAfterXhtml_code_typeCreateCallback != nil {
			stage.OnAfterXhtml_code_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_col_type:
		if stage.OnAfterXhtml_col_typeCreateCallback != nil {
			stage.OnAfterXhtml_col_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_colgroup_type:
		if stage.OnAfterXhtml_colgroup_typeCreateCallback != nil {
			stage.OnAfterXhtml_colgroup_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_dd_type:
		if stage.OnAfterXhtml_dd_typeCreateCallback != nil {
			stage.OnAfterXhtml_dd_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_dfn_type:
		if stage.OnAfterXhtml_dfn_typeCreateCallback != nil {
			stage.OnAfterXhtml_dfn_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_div_type:
		if stage.OnAfterXhtml_div_typeCreateCallback != nil {
			stage.OnAfterXhtml_div_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_dl_type:
		if stage.OnAfterXhtml_dl_typeCreateCallback != nil {
			stage.OnAfterXhtml_dl_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_dt_type:
		if stage.OnAfterXhtml_dt_typeCreateCallback != nil {
			stage.OnAfterXhtml_dt_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_edit_type:
		if stage.OnAfterXhtml_edit_typeCreateCallback != nil {
			stage.OnAfterXhtml_edit_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_em_type:
		if stage.OnAfterXhtml_em_typeCreateCallback != nil {
			stage.OnAfterXhtml_em_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_h1_type:
		if stage.OnAfterXhtml_h1_typeCreateCallback != nil {
			stage.OnAfterXhtml_h1_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_h2_type:
		if stage.OnAfterXhtml_h2_typeCreateCallback != nil {
			stage.OnAfterXhtml_h2_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_h3_type:
		if stage.OnAfterXhtml_h3_typeCreateCallback != nil {
			stage.OnAfterXhtml_h3_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_h4_type:
		if stage.OnAfterXhtml_h4_typeCreateCallback != nil {
			stage.OnAfterXhtml_h4_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_h5_type:
		if stage.OnAfterXhtml_h5_typeCreateCallback != nil {
			stage.OnAfterXhtml_h5_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_h6_type:
		if stage.OnAfterXhtml_h6_typeCreateCallback != nil {
			stage.OnAfterXhtml_h6_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_heading_type:
		if stage.OnAfterXhtml_heading_typeCreateCallback != nil {
			stage.OnAfterXhtml_heading_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_hr_type:
		if stage.OnAfterXhtml_hr_typeCreateCallback != nil {
			stage.OnAfterXhtml_hr_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_kbd_type:
		if stage.OnAfterXhtml_kbd_typeCreateCallback != nil {
			stage.OnAfterXhtml_kbd_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_li_type:
		if stage.OnAfterXhtml_li_typeCreateCallback != nil {
			stage.OnAfterXhtml_li_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_object_type:
		if stage.OnAfterXhtml_object_typeCreateCallback != nil {
			stage.OnAfterXhtml_object_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_ol_type:
		if stage.OnAfterXhtml_ol_typeCreateCallback != nil {
			stage.OnAfterXhtml_ol_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_p_type:
		if stage.OnAfterXhtml_p_typeCreateCallback != nil {
			stage.OnAfterXhtml_p_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_param_type:
		if stage.OnAfterXhtml_param_typeCreateCallback != nil {
			stage.OnAfterXhtml_param_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_pre_type:
		if stage.OnAfterXhtml_pre_typeCreateCallback != nil {
			stage.OnAfterXhtml_pre_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_q_type:
		if stage.OnAfterXhtml_q_typeCreateCallback != nil {
			stage.OnAfterXhtml_q_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_samp_type:
		if stage.OnAfterXhtml_samp_typeCreateCallback != nil {
			stage.OnAfterXhtml_samp_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_span_type:
		if stage.OnAfterXhtml_span_typeCreateCallback != nil {
			stage.OnAfterXhtml_span_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_strong_type:
		if stage.OnAfterXhtml_strong_typeCreateCallback != nil {
			stage.OnAfterXhtml_strong_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_table_type:
		if stage.OnAfterXhtml_table_typeCreateCallback != nil {
			stage.OnAfterXhtml_table_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_tbody_type:
		if stage.OnAfterXhtml_tbody_typeCreateCallback != nil {
			stage.OnAfterXhtml_tbody_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_td_type:
		if stage.OnAfterXhtml_td_typeCreateCallback != nil {
			stage.OnAfterXhtml_td_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_tfoot_type:
		if stage.OnAfterXhtml_tfoot_typeCreateCallback != nil {
			stage.OnAfterXhtml_tfoot_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_th_type:
		if stage.OnAfterXhtml_th_typeCreateCallback != nil {
			stage.OnAfterXhtml_th_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_thead_type:
		if stage.OnAfterXhtml_thead_typeCreateCallback != nil {
			stage.OnAfterXhtml_thead_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_tr_type:
		if stage.OnAfterXhtml_tr_typeCreateCallback != nil {
			stage.OnAfterXhtml_tr_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_ul_type:
		if stage.OnAfterXhtml_ul_typeCreateCallback != nil {
			stage.OnAfterXhtml_ul_typeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Xhtml_var_type:
		if stage.OnAfterXhtml_var_typeCreateCallback != nil {
			stage.OnAfterXhtml_var_typeCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		newTarget := any(new).(*ALTERNATIVE_ID)
		if stage.OnAfterALTERNATIVE_IDUpdateCallback != nil {
			stage.OnAfterALTERNATIVE_IDUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_BOOLEAN)
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_DATE)
		if stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_ENUMERATION)
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_INTEGER)
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_REAL)
		if stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_STRING)
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_XHTML)
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		newTarget := any(new).(*ATTRIBUTE_VALUE_BOOLEAN)
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_DATE:
		newTarget := any(new).(*ATTRIBUTE_VALUE_DATE)
		if stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		newTarget := any(new).(*ATTRIBUTE_VALUE_ENUMERATION)
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		newTarget := any(new).(*ATTRIBUTE_VALUE_INTEGER)
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_REAL:
		newTarget := any(new).(*ATTRIBUTE_VALUE_REAL)
		if stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_STRING:
		newTarget := any(new).(*ATTRIBUTE_VALUE_STRING)
		if stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_XHTML:
		newTarget := any(new).(*ATTRIBUTE_VALUE_XHTML)
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AnyType:
		newTarget := any(new).(*AnyType)
		if stage.OnAfterAnyTypeUpdateCallback != nil {
			stage.OnAfterAnyTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		newTarget := any(new).(*DATATYPE_DEFINITION_BOOLEAN)
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_DATE:
		newTarget := any(new).(*DATATYPE_DEFINITION_DATE)
		if stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		newTarget := any(new).(*DATATYPE_DEFINITION_ENUMERATION)
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_INTEGER:
		newTarget := any(new).(*DATATYPE_DEFINITION_INTEGER)
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_REAL:
		newTarget := any(new).(*DATATYPE_DEFINITION_REAL)
		if stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_STRING:
		newTarget := any(new).(*DATATYPE_DEFINITION_STRING)
		if stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_XHTML:
		newTarget := any(new).(*DATATYPE_DEFINITION_XHTML)
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EMBEDDED_VALUE:
		newTarget := any(new).(*EMBEDDED_VALUE)
		if stage.OnAfterEMBEDDED_VALUEUpdateCallback != nil {
			stage.OnAfterEMBEDDED_VALUEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ENUM_VALUE:
		newTarget := any(new).(*ENUM_VALUE)
		if stage.OnAfterENUM_VALUEUpdateCallback != nil {
			stage.OnAfterENUM_VALUEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RELATION_GROUP:
		newTarget := any(new).(*RELATION_GROUP)
		if stage.OnAfterRELATION_GROUPUpdateCallback != nil {
			stage.OnAfterRELATION_GROUPUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RELATION_GROUP_TYPE:
		newTarget := any(new).(*RELATION_GROUP_TYPE)
		if stage.OnAfterRELATION_GROUP_TYPEUpdateCallback != nil {
			stage.OnAfterRELATION_GROUP_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF:
		newTarget := any(new).(*REQ_IF)
		if stage.OnAfterREQ_IFUpdateCallback != nil {
			stage.OnAfterREQ_IFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF_CONTENT:
		newTarget := any(new).(*REQ_IF_CONTENT)
		if stage.OnAfterREQ_IF_CONTENTUpdateCallback != nil {
			stage.OnAfterREQ_IF_CONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF_HEADER:
		newTarget := any(new).(*REQ_IF_HEADER)
		if stage.OnAfterREQ_IF_HEADERUpdateCallback != nil {
			stage.OnAfterREQ_IF_HEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF_TOOL_EXTENSION:
		newTarget := any(new).(*REQ_IF_TOOL_EXTENSION)
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback != nil {
			stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECIFICATION:
		newTarget := any(new).(*SPECIFICATION)
		if stage.OnAfterSPECIFICATIONUpdateCallback != nil {
			stage.OnAfterSPECIFICATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECIFICATION_TYPE:
		newTarget := any(new).(*SPECIFICATION_TYPE)
		if stage.OnAfterSPECIFICATION_TYPEUpdateCallback != nil {
			stage.OnAfterSPECIFICATION_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_HIERARCHY:
		newTarget := any(new).(*SPEC_HIERARCHY)
		if stage.OnAfterSPEC_HIERARCHYUpdateCallback != nil {
			stage.OnAfterSPEC_HIERARCHYUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_OBJECT:
		newTarget := any(new).(*SPEC_OBJECT)
		if stage.OnAfterSPEC_OBJECTUpdateCallback != nil {
			stage.OnAfterSPEC_OBJECTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_OBJECT_TYPE:
		newTarget := any(new).(*SPEC_OBJECT_TYPE)
		if stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback != nil {
			stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_RELATION:
		newTarget := any(new).(*SPEC_RELATION)
		if stage.OnAfterSPEC_RELATIONUpdateCallback != nil {
			stage.OnAfterSPEC_RELATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_RELATION_TYPE:
		newTarget := any(new).(*SPEC_RELATION_TYPE)
		if stage.OnAfterSPEC_RELATION_TYPEUpdateCallback != nil {
			stage.OnAfterSPEC_RELATION_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *XHTML_CONTENT:
		newTarget := any(new).(*XHTML_CONTENT)
		if stage.OnAfterXHTML_CONTENTUpdateCallback != nil {
			stage.OnAfterXHTML_CONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_InlPres_type:
		newTarget := any(new).(*Xhtml_InlPres_type)
		if stage.OnAfterXhtml_InlPres_typeUpdateCallback != nil {
			stage.OnAfterXhtml_InlPres_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_a_type:
		newTarget := any(new).(*Xhtml_a_type)
		if stage.OnAfterXhtml_a_typeUpdateCallback != nil {
			stage.OnAfterXhtml_a_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_abbr_type:
		newTarget := any(new).(*Xhtml_abbr_type)
		if stage.OnAfterXhtml_abbr_typeUpdateCallback != nil {
			stage.OnAfterXhtml_abbr_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_acronym_type:
		newTarget := any(new).(*Xhtml_acronym_type)
		if stage.OnAfterXhtml_acronym_typeUpdateCallback != nil {
			stage.OnAfterXhtml_acronym_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_address_type:
		newTarget := any(new).(*Xhtml_address_type)
		if stage.OnAfterXhtml_address_typeUpdateCallback != nil {
			stage.OnAfterXhtml_address_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_blockquote_type:
		newTarget := any(new).(*Xhtml_blockquote_type)
		if stage.OnAfterXhtml_blockquote_typeUpdateCallback != nil {
			stage.OnAfterXhtml_blockquote_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_br_type:
		newTarget := any(new).(*Xhtml_br_type)
		if stage.OnAfterXhtml_br_typeUpdateCallback != nil {
			stage.OnAfterXhtml_br_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_caption_type:
		newTarget := any(new).(*Xhtml_caption_type)
		if stage.OnAfterXhtml_caption_typeUpdateCallback != nil {
			stage.OnAfterXhtml_caption_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_cite_type:
		newTarget := any(new).(*Xhtml_cite_type)
		if stage.OnAfterXhtml_cite_typeUpdateCallback != nil {
			stage.OnAfterXhtml_cite_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_code_type:
		newTarget := any(new).(*Xhtml_code_type)
		if stage.OnAfterXhtml_code_typeUpdateCallback != nil {
			stage.OnAfterXhtml_code_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_col_type:
		newTarget := any(new).(*Xhtml_col_type)
		if stage.OnAfterXhtml_col_typeUpdateCallback != nil {
			stage.OnAfterXhtml_col_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_colgroup_type:
		newTarget := any(new).(*Xhtml_colgroup_type)
		if stage.OnAfterXhtml_colgroup_typeUpdateCallback != nil {
			stage.OnAfterXhtml_colgroup_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_dd_type:
		newTarget := any(new).(*Xhtml_dd_type)
		if stage.OnAfterXhtml_dd_typeUpdateCallback != nil {
			stage.OnAfterXhtml_dd_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_dfn_type:
		newTarget := any(new).(*Xhtml_dfn_type)
		if stage.OnAfterXhtml_dfn_typeUpdateCallback != nil {
			stage.OnAfterXhtml_dfn_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_div_type:
		newTarget := any(new).(*Xhtml_div_type)
		if stage.OnAfterXhtml_div_typeUpdateCallback != nil {
			stage.OnAfterXhtml_div_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_dl_type:
		newTarget := any(new).(*Xhtml_dl_type)
		if stage.OnAfterXhtml_dl_typeUpdateCallback != nil {
			stage.OnAfterXhtml_dl_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_dt_type:
		newTarget := any(new).(*Xhtml_dt_type)
		if stage.OnAfterXhtml_dt_typeUpdateCallback != nil {
			stage.OnAfterXhtml_dt_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_edit_type:
		newTarget := any(new).(*Xhtml_edit_type)
		if stage.OnAfterXhtml_edit_typeUpdateCallback != nil {
			stage.OnAfterXhtml_edit_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_em_type:
		newTarget := any(new).(*Xhtml_em_type)
		if stage.OnAfterXhtml_em_typeUpdateCallback != nil {
			stage.OnAfterXhtml_em_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_h1_type:
		newTarget := any(new).(*Xhtml_h1_type)
		if stage.OnAfterXhtml_h1_typeUpdateCallback != nil {
			stage.OnAfterXhtml_h1_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_h2_type:
		newTarget := any(new).(*Xhtml_h2_type)
		if stage.OnAfterXhtml_h2_typeUpdateCallback != nil {
			stage.OnAfterXhtml_h2_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_h3_type:
		newTarget := any(new).(*Xhtml_h3_type)
		if stage.OnAfterXhtml_h3_typeUpdateCallback != nil {
			stage.OnAfterXhtml_h3_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_h4_type:
		newTarget := any(new).(*Xhtml_h4_type)
		if stage.OnAfterXhtml_h4_typeUpdateCallback != nil {
			stage.OnAfterXhtml_h4_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_h5_type:
		newTarget := any(new).(*Xhtml_h5_type)
		if stage.OnAfterXhtml_h5_typeUpdateCallback != nil {
			stage.OnAfterXhtml_h5_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_h6_type:
		newTarget := any(new).(*Xhtml_h6_type)
		if stage.OnAfterXhtml_h6_typeUpdateCallback != nil {
			stage.OnAfterXhtml_h6_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_heading_type:
		newTarget := any(new).(*Xhtml_heading_type)
		if stage.OnAfterXhtml_heading_typeUpdateCallback != nil {
			stage.OnAfterXhtml_heading_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_hr_type:
		newTarget := any(new).(*Xhtml_hr_type)
		if stage.OnAfterXhtml_hr_typeUpdateCallback != nil {
			stage.OnAfterXhtml_hr_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_kbd_type:
		newTarget := any(new).(*Xhtml_kbd_type)
		if stage.OnAfterXhtml_kbd_typeUpdateCallback != nil {
			stage.OnAfterXhtml_kbd_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_li_type:
		newTarget := any(new).(*Xhtml_li_type)
		if stage.OnAfterXhtml_li_typeUpdateCallback != nil {
			stage.OnAfterXhtml_li_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_object_type:
		newTarget := any(new).(*Xhtml_object_type)
		if stage.OnAfterXhtml_object_typeUpdateCallback != nil {
			stage.OnAfterXhtml_object_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_ol_type:
		newTarget := any(new).(*Xhtml_ol_type)
		if stage.OnAfterXhtml_ol_typeUpdateCallback != nil {
			stage.OnAfterXhtml_ol_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_p_type:
		newTarget := any(new).(*Xhtml_p_type)
		if stage.OnAfterXhtml_p_typeUpdateCallback != nil {
			stage.OnAfterXhtml_p_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_param_type:
		newTarget := any(new).(*Xhtml_param_type)
		if stage.OnAfterXhtml_param_typeUpdateCallback != nil {
			stage.OnAfterXhtml_param_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_pre_type:
		newTarget := any(new).(*Xhtml_pre_type)
		if stage.OnAfterXhtml_pre_typeUpdateCallback != nil {
			stage.OnAfterXhtml_pre_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_q_type:
		newTarget := any(new).(*Xhtml_q_type)
		if stage.OnAfterXhtml_q_typeUpdateCallback != nil {
			stage.OnAfterXhtml_q_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_samp_type:
		newTarget := any(new).(*Xhtml_samp_type)
		if stage.OnAfterXhtml_samp_typeUpdateCallback != nil {
			stage.OnAfterXhtml_samp_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_span_type:
		newTarget := any(new).(*Xhtml_span_type)
		if stage.OnAfterXhtml_span_typeUpdateCallback != nil {
			stage.OnAfterXhtml_span_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_strong_type:
		newTarget := any(new).(*Xhtml_strong_type)
		if stage.OnAfterXhtml_strong_typeUpdateCallback != nil {
			stage.OnAfterXhtml_strong_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_table_type:
		newTarget := any(new).(*Xhtml_table_type)
		if stage.OnAfterXhtml_table_typeUpdateCallback != nil {
			stage.OnAfterXhtml_table_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_tbody_type:
		newTarget := any(new).(*Xhtml_tbody_type)
		if stage.OnAfterXhtml_tbody_typeUpdateCallback != nil {
			stage.OnAfterXhtml_tbody_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_td_type:
		newTarget := any(new).(*Xhtml_td_type)
		if stage.OnAfterXhtml_td_typeUpdateCallback != nil {
			stage.OnAfterXhtml_td_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_tfoot_type:
		newTarget := any(new).(*Xhtml_tfoot_type)
		if stage.OnAfterXhtml_tfoot_typeUpdateCallback != nil {
			stage.OnAfterXhtml_tfoot_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_th_type:
		newTarget := any(new).(*Xhtml_th_type)
		if stage.OnAfterXhtml_th_typeUpdateCallback != nil {
			stage.OnAfterXhtml_th_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_thead_type:
		newTarget := any(new).(*Xhtml_thead_type)
		if stage.OnAfterXhtml_thead_typeUpdateCallback != nil {
			stage.OnAfterXhtml_thead_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_tr_type:
		newTarget := any(new).(*Xhtml_tr_type)
		if stage.OnAfterXhtml_tr_typeUpdateCallback != nil {
			stage.OnAfterXhtml_tr_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_ul_type:
		newTarget := any(new).(*Xhtml_ul_type)
		if stage.OnAfterXhtml_ul_typeUpdateCallback != nil {
			stage.OnAfterXhtml_ul_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Xhtml_var_type:
		newTarget := any(new).(*Xhtml_var_type)
		if stage.OnAfterXhtml_var_typeUpdateCallback != nil {
			stage.OnAfterXhtml_var_typeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *AnyType:
		if stage.OnAfterAnyTypeDeleteCallback != nil {
			staged := any(staged).(*AnyType)
			stage.OnAfterAnyTypeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *XHTML_CONTENT:
		if stage.OnAfterXHTML_CONTENTDeleteCallback != nil {
			staged := any(staged).(*XHTML_CONTENT)
			stage.OnAfterXHTML_CONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_InlPres_type:
		if stage.OnAfterXhtml_InlPres_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_InlPres_type)
			stage.OnAfterXhtml_InlPres_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_a_type:
		if stage.OnAfterXhtml_a_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_a_type)
			stage.OnAfterXhtml_a_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_abbr_type:
		if stage.OnAfterXhtml_abbr_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_abbr_type)
			stage.OnAfterXhtml_abbr_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_acronym_type:
		if stage.OnAfterXhtml_acronym_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_acronym_type)
			stage.OnAfterXhtml_acronym_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_address_type:
		if stage.OnAfterXhtml_address_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_address_type)
			stage.OnAfterXhtml_address_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_blockquote_type:
		if stage.OnAfterXhtml_blockquote_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_blockquote_type)
			stage.OnAfterXhtml_blockquote_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_br_type:
		if stage.OnAfterXhtml_br_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_br_type)
			stage.OnAfterXhtml_br_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_caption_type:
		if stage.OnAfterXhtml_caption_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_caption_type)
			stage.OnAfterXhtml_caption_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_cite_type:
		if stage.OnAfterXhtml_cite_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_cite_type)
			stage.OnAfterXhtml_cite_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_code_type:
		if stage.OnAfterXhtml_code_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_code_type)
			stage.OnAfterXhtml_code_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_col_type:
		if stage.OnAfterXhtml_col_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_col_type)
			stage.OnAfterXhtml_col_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_colgroup_type:
		if stage.OnAfterXhtml_colgroup_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_colgroup_type)
			stage.OnAfterXhtml_colgroup_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_dd_type:
		if stage.OnAfterXhtml_dd_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_dd_type)
			stage.OnAfterXhtml_dd_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_dfn_type:
		if stage.OnAfterXhtml_dfn_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_dfn_type)
			stage.OnAfterXhtml_dfn_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_div_type:
		if stage.OnAfterXhtml_div_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_div_type)
			stage.OnAfterXhtml_div_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_dl_type:
		if stage.OnAfterXhtml_dl_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_dl_type)
			stage.OnAfterXhtml_dl_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_dt_type:
		if stage.OnAfterXhtml_dt_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_dt_type)
			stage.OnAfterXhtml_dt_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_edit_type:
		if stage.OnAfterXhtml_edit_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_edit_type)
			stage.OnAfterXhtml_edit_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_em_type:
		if stage.OnAfterXhtml_em_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_em_type)
			stage.OnAfterXhtml_em_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_h1_type:
		if stage.OnAfterXhtml_h1_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_h1_type)
			stage.OnAfterXhtml_h1_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_h2_type:
		if stage.OnAfterXhtml_h2_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_h2_type)
			stage.OnAfterXhtml_h2_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_h3_type:
		if stage.OnAfterXhtml_h3_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_h3_type)
			stage.OnAfterXhtml_h3_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_h4_type:
		if stage.OnAfterXhtml_h4_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_h4_type)
			stage.OnAfterXhtml_h4_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_h5_type:
		if stage.OnAfterXhtml_h5_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_h5_type)
			stage.OnAfterXhtml_h5_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_h6_type:
		if stage.OnAfterXhtml_h6_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_h6_type)
			stage.OnAfterXhtml_h6_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_heading_type:
		if stage.OnAfterXhtml_heading_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_heading_type)
			stage.OnAfterXhtml_heading_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_hr_type:
		if stage.OnAfterXhtml_hr_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_hr_type)
			stage.OnAfterXhtml_hr_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_kbd_type:
		if stage.OnAfterXhtml_kbd_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_kbd_type)
			stage.OnAfterXhtml_kbd_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_li_type:
		if stage.OnAfterXhtml_li_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_li_type)
			stage.OnAfterXhtml_li_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_object_type:
		if stage.OnAfterXhtml_object_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_object_type)
			stage.OnAfterXhtml_object_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_ol_type:
		if stage.OnAfterXhtml_ol_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_ol_type)
			stage.OnAfterXhtml_ol_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_p_type:
		if stage.OnAfterXhtml_p_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_p_type)
			stage.OnAfterXhtml_p_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_param_type:
		if stage.OnAfterXhtml_param_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_param_type)
			stage.OnAfterXhtml_param_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_pre_type:
		if stage.OnAfterXhtml_pre_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_pre_type)
			stage.OnAfterXhtml_pre_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_q_type:
		if stage.OnAfterXhtml_q_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_q_type)
			stage.OnAfterXhtml_q_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_samp_type:
		if stage.OnAfterXhtml_samp_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_samp_type)
			stage.OnAfterXhtml_samp_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_span_type:
		if stage.OnAfterXhtml_span_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_span_type)
			stage.OnAfterXhtml_span_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_strong_type:
		if stage.OnAfterXhtml_strong_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_strong_type)
			stage.OnAfterXhtml_strong_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_table_type:
		if stage.OnAfterXhtml_table_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_table_type)
			stage.OnAfterXhtml_table_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_tbody_type:
		if stage.OnAfterXhtml_tbody_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_tbody_type)
			stage.OnAfterXhtml_tbody_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_td_type:
		if stage.OnAfterXhtml_td_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_td_type)
			stage.OnAfterXhtml_td_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_tfoot_type:
		if stage.OnAfterXhtml_tfoot_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_tfoot_type)
			stage.OnAfterXhtml_tfoot_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_th_type:
		if stage.OnAfterXhtml_th_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_th_type)
			stage.OnAfterXhtml_th_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_thead_type:
		if stage.OnAfterXhtml_thead_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_thead_type)
			stage.OnAfterXhtml_thead_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_tr_type:
		if stage.OnAfterXhtml_tr_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_tr_type)
			stage.OnAfterXhtml_tr_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_ul_type:
		if stage.OnAfterXhtml_ul_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_ul_type)
			stage.OnAfterXhtml_ul_typeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Xhtml_var_type:
		if stage.OnAfterXhtml_var_typeDeleteCallback != nil {
			staged := any(staged).(*Xhtml_var_type)
			stage.OnAfterXhtml_var_typeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *AnyType:
		if stage.OnAfterAnyTypeReadCallback != nil {
			stage.OnAfterAnyTypeReadCallback.OnAfterRead(stage, target)
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
	case *XHTML_CONTENT:
		if stage.OnAfterXHTML_CONTENTReadCallback != nil {
			stage.OnAfterXHTML_CONTENTReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_InlPres_type:
		if stage.OnAfterXhtml_InlPres_typeReadCallback != nil {
			stage.OnAfterXhtml_InlPres_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_a_type:
		if stage.OnAfterXhtml_a_typeReadCallback != nil {
			stage.OnAfterXhtml_a_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_abbr_type:
		if stage.OnAfterXhtml_abbr_typeReadCallback != nil {
			stage.OnAfterXhtml_abbr_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_acronym_type:
		if stage.OnAfterXhtml_acronym_typeReadCallback != nil {
			stage.OnAfterXhtml_acronym_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_address_type:
		if stage.OnAfterXhtml_address_typeReadCallback != nil {
			stage.OnAfterXhtml_address_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_blockquote_type:
		if stage.OnAfterXhtml_blockquote_typeReadCallback != nil {
			stage.OnAfterXhtml_blockquote_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_br_type:
		if stage.OnAfterXhtml_br_typeReadCallback != nil {
			stage.OnAfterXhtml_br_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_caption_type:
		if stage.OnAfterXhtml_caption_typeReadCallback != nil {
			stage.OnAfterXhtml_caption_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_cite_type:
		if stage.OnAfterXhtml_cite_typeReadCallback != nil {
			stage.OnAfterXhtml_cite_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_code_type:
		if stage.OnAfterXhtml_code_typeReadCallback != nil {
			stage.OnAfterXhtml_code_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_col_type:
		if stage.OnAfterXhtml_col_typeReadCallback != nil {
			stage.OnAfterXhtml_col_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_colgroup_type:
		if stage.OnAfterXhtml_colgroup_typeReadCallback != nil {
			stage.OnAfterXhtml_colgroup_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_dd_type:
		if stage.OnAfterXhtml_dd_typeReadCallback != nil {
			stage.OnAfterXhtml_dd_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_dfn_type:
		if stage.OnAfterXhtml_dfn_typeReadCallback != nil {
			stage.OnAfterXhtml_dfn_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_div_type:
		if stage.OnAfterXhtml_div_typeReadCallback != nil {
			stage.OnAfterXhtml_div_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_dl_type:
		if stage.OnAfterXhtml_dl_typeReadCallback != nil {
			stage.OnAfterXhtml_dl_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_dt_type:
		if stage.OnAfterXhtml_dt_typeReadCallback != nil {
			stage.OnAfterXhtml_dt_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_edit_type:
		if stage.OnAfterXhtml_edit_typeReadCallback != nil {
			stage.OnAfterXhtml_edit_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_em_type:
		if stage.OnAfterXhtml_em_typeReadCallback != nil {
			stage.OnAfterXhtml_em_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_h1_type:
		if stage.OnAfterXhtml_h1_typeReadCallback != nil {
			stage.OnAfterXhtml_h1_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_h2_type:
		if stage.OnAfterXhtml_h2_typeReadCallback != nil {
			stage.OnAfterXhtml_h2_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_h3_type:
		if stage.OnAfterXhtml_h3_typeReadCallback != nil {
			stage.OnAfterXhtml_h3_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_h4_type:
		if stage.OnAfterXhtml_h4_typeReadCallback != nil {
			stage.OnAfterXhtml_h4_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_h5_type:
		if stage.OnAfterXhtml_h5_typeReadCallback != nil {
			stage.OnAfterXhtml_h5_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_h6_type:
		if stage.OnAfterXhtml_h6_typeReadCallback != nil {
			stage.OnAfterXhtml_h6_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_heading_type:
		if stage.OnAfterXhtml_heading_typeReadCallback != nil {
			stage.OnAfterXhtml_heading_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_hr_type:
		if stage.OnAfterXhtml_hr_typeReadCallback != nil {
			stage.OnAfterXhtml_hr_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_kbd_type:
		if stage.OnAfterXhtml_kbd_typeReadCallback != nil {
			stage.OnAfterXhtml_kbd_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_li_type:
		if stage.OnAfterXhtml_li_typeReadCallback != nil {
			stage.OnAfterXhtml_li_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_object_type:
		if stage.OnAfterXhtml_object_typeReadCallback != nil {
			stage.OnAfterXhtml_object_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_ol_type:
		if stage.OnAfterXhtml_ol_typeReadCallback != nil {
			stage.OnAfterXhtml_ol_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_p_type:
		if stage.OnAfterXhtml_p_typeReadCallback != nil {
			stage.OnAfterXhtml_p_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_param_type:
		if stage.OnAfterXhtml_param_typeReadCallback != nil {
			stage.OnAfterXhtml_param_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_pre_type:
		if stage.OnAfterXhtml_pre_typeReadCallback != nil {
			stage.OnAfterXhtml_pre_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_q_type:
		if stage.OnAfterXhtml_q_typeReadCallback != nil {
			stage.OnAfterXhtml_q_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_samp_type:
		if stage.OnAfterXhtml_samp_typeReadCallback != nil {
			stage.OnAfterXhtml_samp_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_span_type:
		if stage.OnAfterXhtml_span_typeReadCallback != nil {
			stage.OnAfterXhtml_span_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_strong_type:
		if stage.OnAfterXhtml_strong_typeReadCallback != nil {
			stage.OnAfterXhtml_strong_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_table_type:
		if stage.OnAfterXhtml_table_typeReadCallback != nil {
			stage.OnAfterXhtml_table_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_tbody_type:
		if stage.OnAfterXhtml_tbody_typeReadCallback != nil {
			stage.OnAfterXhtml_tbody_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_td_type:
		if stage.OnAfterXhtml_td_typeReadCallback != nil {
			stage.OnAfterXhtml_td_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_tfoot_type:
		if stage.OnAfterXhtml_tfoot_typeReadCallback != nil {
			stage.OnAfterXhtml_tfoot_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_th_type:
		if stage.OnAfterXhtml_th_typeReadCallback != nil {
			stage.OnAfterXhtml_th_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_thead_type:
		if stage.OnAfterXhtml_thead_typeReadCallback != nil {
			stage.OnAfterXhtml_thead_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_tr_type:
		if stage.OnAfterXhtml_tr_typeReadCallback != nil {
			stage.OnAfterXhtml_tr_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_ul_type:
		if stage.OnAfterXhtml_ul_typeReadCallback != nil {
			stage.OnAfterXhtml_ul_typeReadCallback.OnAfterRead(stage, target)
		}
	case *Xhtml_var_type:
		if stage.OnAfterXhtml_var_typeReadCallback != nil {
			stage.OnAfterXhtml_var_typeReadCallback.OnAfterRead(stage, target)
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
	
	case *AnyType:
		stage.OnAfterAnyTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[AnyType])
	
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
	
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[XHTML_CONTENT])
	
	case *Xhtml_InlPres_type:
		stage.OnAfterXhtml_InlPres_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_InlPres_type])
	
	case *Xhtml_a_type:
		stage.OnAfterXhtml_a_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_a_type])
	
	case *Xhtml_abbr_type:
		stage.OnAfterXhtml_abbr_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_abbr_type])
	
	case *Xhtml_acronym_type:
		stage.OnAfterXhtml_acronym_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_acronym_type])
	
	case *Xhtml_address_type:
		stage.OnAfterXhtml_address_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_address_type])
	
	case *Xhtml_blockquote_type:
		stage.OnAfterXhtml_blockquote_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_blockquote_type])
	
	case *Xhtml_br_type:
		stage.OnAfterXhtml_br_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_br_type])
	
	case *Xhtml_caption_type:
		stage.OnAfterXhtml_caption_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_caption_type])
	
	case *Xhtml_cite_type:
		stage.OnAfterXhtml_cite_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_cite_type])
	
	case *Xhtml_code_type:
		stage.OnAfterXhtml_code_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_code_type])
	
	case *Xhtml_col_type:
		stage.OnAfterXhtml_col_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_col_type])
	
	case *Xhtml_colgroup_type:
		stage.OnAfterXhtml_colgroup_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_colgroup_type])
	
	case *Xhtml_dd_type:
		stage.OnAfterXhtml_dd_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_dd_type])
	
	case *Xhtml_dfn_type:
		stage.OnAfterXhtml_dfn_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_dfn_type])
	
	case *Xhtml_div_type:
		stage.OnAfterXhtml_div_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_div_type])
	
	case *Xhtml_dl_type:
		stage.OnAfterXhtml_dl_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_dl_type])
	
	case *Xhtml_dt_type:
		stage.OnAfterXhtml_dt_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_dt_type])
	
	case *Xhtml_edit_type:
		stage.OnAfterXhtml_edit_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_edit_type])
	
	case *Xhtml_em_type:
		stage.OnAfterXhtml_em_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_em_type])
	
	case *Xhtml_h1_type:
		stage.OnAfterXhtml_h1_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_h1_type])
	
	case *Xhtml_h2_type:
		stage.OnAfterXhtml_h2_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_h2_type])
	
	case *Xhtml_h3_type:
		stage.OnAfterXhtml_h3_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_h3_type])
	
	case *Xhtml_h4_type:
		stage.OnAfterXhtml_h4_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_h4_type])
	
	case *Xhtml_h5_type:
		stage.OnAfterXhtml_h5_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_h5_type])
	
	case *Xhtml_h6_type:
		stage.OnAfterXhtml_h6_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_h6_type])
	
	case *Xhtml_heading_type:
		stage.OnAfterXhtml_heading_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_heading_type])
	
	case *Xhtml_hr_type:
		stage.OnAfterXhtml_hr_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_hr_type])
	
	case *Xhtml_kbd_type:
		stage.OnAfterXhtml_kbd_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_kbd_type])
	
	case *Xhtml_li_type:
		stage.OnAfterXhtml_li_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_li_type])
	
	case *Xhtml_object_type:
		stage.OnAfterXhtml_object_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_object_type])
	
	case *Xhtml_ol_type:
		stage.OnAfterXhtml_ol_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_ol_type])
	
	case *Xhtml_p_type:
		stage.OnAfterXhtml_p_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_p_type])
	
	case *Xhtml_param_type:
		stage.OnAfterXhtml_param_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_param_type])
	
	case *Xhtml_pre_type:
		stage.OnAfterXhtml_pre_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_pre_type])
	
	case *Xhtml_q_type:
		stage.OnAfterXhtml_q_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_q_type])
	
	case *Xhtml_samp_type:
		stage.OnAfterXhtml_samp_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_samp_type])
	
	case *Xhtml_span_type:
		stage.OnAfterXhtml_span_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_span_type])
	
	case *Xhtml_strong_type:
		stage.OnAfterXhtml_strong_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_strong_type])
	
	case *Xhtml_table_type:
		stage.OnAfterXhtml_table_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_table_type])
	
	case *Xhtml_tbody_type:
		stage.OnAfterXhtml_tbody_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_tbody_type])
	
	case *Xhtml_td_type:
		stage.OnAfterXhtml_td_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_td_type])
	
	case *Xhtml_tfoot_type:
		stage.OnAfterXhtml_tfoot_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_tfoot_type])
	
	case *Xhtml_th_type:
		stage.OnAfterXhtml_th_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_th_type])
	
	case *Xhtml_thead_type:
		stage.OnAfterXhtml_thead_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_thead_type])
	
	case *Xhtml_tr_type:
		stage.OnAfterXhtml_tr_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_tr_type])
	
	case *Xhtml_ul_type:
		stage.OnAfterXhtml_ul_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_ul_type])
	
	case *Xhtml_var_type:
		stage.OnAfterXhtml_var_typeUpdateCallback = any(callback).(OnAfterUpdateInterface[Xhtml_var_type])
	
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
	
	case *AnyType:
		stage.OnAfterAnyTypeCreateCallback = any(callback).(OnAfterCreateInterface[AnyType])
	
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
	
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTCreateCallback = any(callback).(OnAfterCreateInterface[XHTML_CONTENT])
	
	case *Xhtml_InlPres_type:
		stage.OnAfterXhtml_InlPres_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_InlPres_type])
	
	case *Xhtml_a_type:
		stage.OnAfterXhtml_a_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_a_type])
	
	case *Xhtml_abbr_type:
		stage.OnAfterXhtml_abbr_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_abbr_type])
	
	case *Xhtml_acronym_type:
		stage.OnAfterXhtml_acronym_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_acronym_type])
	
	case *Xhtml_address_type:
		stage.OnAfterXhtml_address_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_address_type])
	
	case *Xhtml_blockquote_type:
		stage.OnAfterXhtml_blockquote_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_blockquote_type])
	
	case *Xhtml_br_type:
		stage.OnAfterXhtml_br_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_br_type])
	
	case *Xhtml_caption_type:
		stage.OnAfterXhtml_caption_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_caption_type])
	
	case *Xhtml_cite_type:
		stage.OnAfterXhtml_cite_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_cite_type])
	
	case *Xhtml_code_type:
		stage.OnAfterXhtml_code_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_code_type])
	
	case *Xhtml_col_type:
		stage.OnAfterXhtml_col_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_col_type])
	
	case *Xhtml_colgroup_type:
		stage.OnAfterXhtml_colgroup_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_colgroup_type])
	
	case *Xhtml_dd_type:
		stage.OnAfterXhtml_dd_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_dd_type])
	
	case *Xhtml_dfn_type:
		stage.OnAfterXhtml_dfn_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_dfn_type])
	
	case *Xhtml_div_type:
		stage.OnAfterXhtml_div_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_div_type])
	
	case *Xhtml_dl_type:
		stage.OnAfterXhtml_dl_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_dl_type])
	
	case *Xhtml_dt_type:
		stage.OnAfterXhtml_dt_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_dt_type])
	
	case *Xhtml_edit_type:
		stage.OnAfterXhtml_edit_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_edit_type])
	
	case *Xhtml_em_type:
		stage.OnAfterXhtml_em_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_em_type])
	
	case *Xhtml_h1_type:
		stage.OnAfterXhtml_h1_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_h1_type])
	
	case *Xhtml_h2_type:
		stage.OnAfterXhtml_h2_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_h2_type])
	
	case *Xhtml_h3_type:
		stage.OnAfterXhtml_h3_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_h3_type])
	
	case *Xhtml_h4_type:
		stage.OnAfterXhtml_h4_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_h4_type])
	
	case *Xhtml_h5_type:
		stage.OnAfterXhtml_h5_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_h5_type])
	
	case *Xhtml_h6_type:
		stage.OnAfterXhtml_h6_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_h6_type])
	
	case *Xhtml_heading_type:
		stage.OnAfterXhtml_heading_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_heading_type])
	
	case *Xhtml_hr_type:
		stage.OnAfterXhtml_hr_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_hr_type])
	
	case *Xhtml_kbd_type:
		stage.OnAfterXhtml_kbd_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_kbd_type])
	
	case *Xhtml_li_type:
		stage.OnAfterXhtml_li_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_li_type])
	
	case *Xhtml_object_type:
		stage.OnAfterXhtml_object_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_object_type])
	
	case *Xhtml_ol_type:
		stage.OnAfterXhtml_ol_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_ol_type])
	
	case *Xhtml_p_type:
		stage.OnAfterXhtml_p_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_p_type])
	
	case *Xhtml_param_type:
		stage.OnAfterXhtml_param_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_param_type])
	
	case *Xhtml_pre_type:
		stage.OnAfterXhtml_pre_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_pre_type])
	
	case *Xhtml_q_type:
		stage.OnAfterXhtml_q_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_q_type])
	
	case *Xhtml_samp_type:
		stage.OnAfterXhtml_samp_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_samp_type])
	
	case *Xhtml_span_type:
		stage.OnAfterXhtml_span_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_span_type])
	
	case *Xhtml_strong_type:
		stage.OnAfterXhtml_strong_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_strong_type])
	
	case *Xhtml_table_type:
		stage.OnAfterXhtml_table_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_table_type])
	
	case *Xhtml_tbody_type:
		stage.OnAfterXhtml_tbody_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_tbody_type])
	
	case *Xhtml_td_type:
		stage.OnAfterXhtml_td_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_td_type])
	
	case *Xhtml_tfoot_type:
		stage.OnAfterXhtml_tfoot_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_tfoot_type])
	
	case *Xhtml_th_type:
		stage.OnAfterXhtml_th_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_th_type])
	
	case *Xhtml_thead_type:
		stage.OnAfterXhtml_thead_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_thead_type])
	
	case *Xhtml_tr_type:
		stage.OnAfterXhtml_tr_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_tr_type])
	
	case *Xhtml_ul_type:
		stage.OnAfterXhtml_ul_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_ul_type])
	
	case *Xhtml_var_type:
		stage.OnAfterXhtml_var_typeCreateCallback = any(callback).(OnAfterCreateInterface[Xhtml_var_type])
	
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
	
	case *AnyType:
		stage.OnAfterAnyTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[AnyType])
	
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
	
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[XHTML_CONTENT])
	
	case *Xhtml_InlPres_type:
		stage.OnAfterXhtml_InlPres_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_InlPres_type])
	
	case *Xhtml_a_type:
		stage.OnAfterXhtml_a_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_a_type])
	
	case *Xhtml_abbr_type:
		stage.OnAfterXhtml_abbr_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_abbr_type])
	
	case *Xhtml_acronym_type:
		stage.OnAfterXhtml_acronym_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_acronym_type])
	
	case *Xhtml_address_type:
		stage.OnAfterXhtml_address_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_address_type])
	
	case *Xhtml_blockquote_type:
		stage.OnAfterXhtml_blockquote_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_blockquote_type])
	
	case *Xhtml_br_type:
		stage.OnAfterXhtml_br_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_br_type])
	
	case *Xhtml_caption_type:
		stage.OnAfterXhtml_caption_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_caption_type])
	
	case *Xhtml_cite_type:
		stage.OnAfterXhtml_cite_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_cite_type])
	
	case *Xhtml_code_type:
		stage.OnAfterXhtml_code_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_code_type])
	
	case *Xhtml_col_type:
		stage.OnAfterXhtml_col_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_col_type])
	
	case *Xhtml_colgroup_type:
		stage.OnAfterXhtml_colgroup_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_colgroup_type])
	
	case *Xhtml_dd_type:
		stage.OnAfterXhtml_dd_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_dd_type])
	
	case *Xhtml_dfn_type:
		stage.OnAfterXhtml_dfn_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_dfn_type])
	
	case *Xhtml_div_type:
		stage.OnAfterXhtml_div_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_div_type])
	
	case *Xhtml_dl_type:
		stage.OnAfterXhtml_dl_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_dl_type])
	
	case *Xhtml_dt_type:
		stage.OnAfterXhtml_dt_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_dt_type])
	
	case *Xhtml_edit_type:
		stage.OnAfterXhtml_edit_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_edit_type])
	
	case *Xhtml_em_type:
		stage.OnAfterXhtml_em_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_em_type])
	
	case *Xhtml_h1_type:
		stage.OnAfterXhtml_h1_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_h1_type])
	
	case *Xhtml_h2_type:
		stage.OnAfterXhtml_h2_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_h2_type])
	
	case *Xhtml_h3_type:
		stage.OnAfterXhtml_h3_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_h3_type])
	
	case *Xhtml_h4_type:
		stage.OnAfterXhtml_h4_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_h4_type])
	
	case *Xhtml_h5_type:
		stage.OnAfterXhtml_h5_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_h5_type])
	
	case *Xhtml_h6_type:
		stage.OnAfterXhtml_h6_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_h6_type])
	
	case *Xhtml_heading_type:
		stage.OnAfterXhtml_heading_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_heading_type])
	
	case *Xhtml_hr_type:
		stage.OnAfterXhtml_hr_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_hr_type])
	
	case *Xhtml_kbd_type:
		stage.OnAfterXhtml_kbd_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_kbd_type])
	
	case *Xhtml_li_type:
		stage.OnAfterXhtml_li_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_li_type])
	
	case *Xhtml_object_type:
		stage.OnAfterXhtml_object_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_object_type])
	
	case *Xhtml_ol_type:
		stage.OnAfterXhtml_ol_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_ol_type])
	
	case *Xhtml_p_type:
		stage.OnAfterXhtml_p_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_p_type])
	
	case *Xhtml_param_type:
		stage.OnAfterXhtml_param_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_param_type])
	
	case *Xhtml_pre_type:
		stage.OnAfterXhtml_pre_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_pre_type])
	
	case *Xhtml_q_type:
		stage.OnAfterXhtml_q_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_q_type])
	
	case *Xhtml_samp_type:
		stage.OnAfterXhtml_samp_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_samp_type])
	
	case *Xhtml_span_type:
		stage.OnAfterXhtml_span_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_span_type])
	
	case *Xhtml_strong_type:
		stage.OnAfterXhtml_strong_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_strong_type])
	
	case *Xhtml_table_type:
		stage.OnAfterXhtml_table_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_table_type])
	
	case *Xhtml_tbody_type:
		stage.OnAfterXhtml_tbody_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_tbody_type])
	
	case *Xhtml_td_type:
		stage.OnAfterXhtml_td_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_td_type])
	
	case *Xhtml_tfoot_type:
		stage.OnAfterXhtml_tfoot_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_tfoot_type])
	
	case *Xhtml_th_type:
		stage.OnAfterXhtml_th_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_th_type])
	
	case *Xhtml_thead_type:
		stage.OnAfterXhtml_thead_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_thead_type])
	
	case *Xhtml_tr_type:
		stage.OnAfterXhtml_tr_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_tr_type])
	
	case *Xhtml_ul_type:
		stage.OnAfterXhtml_ul_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_ul_type])
	
	case *Xhtml_var_type:
		stage.OnAfterXhtml_var_typeDeleteCallback = any(callback).(OnAfterDeleteInterface[Xhtml_var_type])
	
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
	
	case *AnyType:
		stage.OnAfterAnyTypeReadCallback = any(callback).(OnAfterReadInterface[AnyType])
	
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
	
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTReadCallback = any(callback).(OnAfterReadInterface[XHTML_CONTENT])
	
	case *Xhtml_InlPres_type:
		stage.OnAfterXhtml_InlPres_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_InlPres_type])
	
	case *Xhtml_a_type:
		stage.OnAfterXhtml_a_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_a_type])
	
	case *Xhtml_abbr_type:
		stage.OnAfterXhtml_abbr_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_abbr_type])
	
	case *Xhtml_acronym_type:
		stage.OnAfterXhtml_acronym_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_acronym_type])
	
	case *Xhtml_address_type:
		stage.OnAfterXhtml_address_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_address_type])
	
	case *Xhtml_blockquote_type:
		stage.OnAfterXhtml_blockquote_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_blockquote_type])
	
	case *Xhtml_br_type:
		stage.OnAfterXhtml_br_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_br_type])
	
	case *Xhtml_caption_type:
		stage.OnAfterXhtml_caption_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_caption_type])
	
	case *Xhtml_cite_type:
		stage.OnAfterXhtml_cite_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_cite_type])
	
	case *Xhtml_code_type:
		stage.OnAfterXhtml_code_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_code_type])
	
	case *Xhtml_col_type:
		stage.OnAfterXhtml_col_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_col_type])
	
	case *Xhtml_colgroup_type:
		stage.OnAfterXhtml_colgroup_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_colgroup_type])
	
	case *Xhtml_dd_type:
		stage.OnAfterXhtml_dd_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_dd_type])
	
	case *Xhtml_dfn_type:
		stage.OnAfterXhtml_dfn_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_dfn_type])
	
	case *Xhtml_div_type:
		stage.OnAfterXhtml_div_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_div_type])
	
	case *Xhtml_dl_type:
		stage.OnAfterXhtml_dl_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_dl_type])
	
	case *Xhtml_dt_type:
		stage.OnAfterXhtml_dt_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_dt_type])
	
	case *Xhtml_edit_type:
		stage.OnAfterXhtml_edit_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_edit_type])
	
	case *Xhtml_em_type:
		stage.OnAfterXhtml_em_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_em_type])
	
	case *Xhtml_h1_type:
		stage.OnAfterXhtml_h1_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_h1_type])
	
	case *Xhtml_h2_type:
		stage.OnAfterXhtml_h2_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_h2_type])
	
	case *Xhtml_h3_type:
		stage.OnAfterXhtml_h3_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_h3_type])
	
	case *Xhtml_h4_type:
		stage.OnAfterXhtml_h4_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_h4_type])
	
	case *Xhtml_h5_type:
		stage.OnAfterXhtml_h5_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_h5_type])
	
	case *Xhtml_h6_type:
		stage.OnAfterXhtml_h6_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_h6_type])
	
	case *Xhtml_heading_type:
		stage.OnAfterXhtml_heading_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_heading_type])
	
	case *Xhtml_hr_type:
		stage.OnAfterXhtml_hr_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_hr_type])
	
	case *Xhtml_kbd_type:
		stage.OnAfterXhtml_kbd_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_kbd_type])
	
	case *Xhtml_li_type:
		stage.OnAfterXhtml_li_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_li_type])
	
	case *Xhtml_object_type:
		stage.OnAfterXhtml_object_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_object_type])
	
	case *Xhtml_ol_type:
		stage.OnAfterXhtml_ol_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_ol_type])
	
	case *Xhtml_p_type:
		stage.OnAfterXhtml_p_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_p_type])
	
	case *Xhtml_param_type:
		stage.OnAfterXhtml_param_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_param_type])
	
	case *Xhtml_pre_type:
		stage.OnAfterXhtml_pre_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_pre_type])
	
	case *Xhtml_q_type:
		stage.OnAfterXhtml_q_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_q_type])
	
	case *Xhtml_samp_type:
		stage.OnAfterXhtml_samp_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_samp_type])
	
	case *Xhtml_span_type:
		stage.OnAfterXhtml_span_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_span_type])
	
	case *Xhtml_strong_type:
		stage.OnAfterXhtml_strong_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_strong_type])
	
	case *Xhtml_table_type:
		stage.OnAfterXhtml_table_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_table_type])
	
	case *Xhtml_tbody_type:
		stage.OnAfterXhtml_tbody_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_tbody_type])
	
	case *Xhtml_td_type:
		stage.OnAfterXhtml_td_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_td_type])
	
	case *Xhtml_tfoot_type:
		stage.OnAfterXhtml_tfoot_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_tfoot_type])
	
	case *Xhtml_th_type:
		stage.OnAfterXhtml_th_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_th_type])
	
	case *Xhtml_thead_type:
		stage.OnAfterXhtml_thead_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_thead_type])
	
	case *Xhtml_tr_type:
		stage.OnAfterXhtml_tr_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_tr_type])
	
	case *Xhtml_ul_type:
		stage.OnAfterXhtml_ul_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_ul_type])
	
	case *Xhtml_var_type:
		stage.OnAfterXhtml_var_typeReadCallback = any(callback).(OnAfterReadInterface[Xhtml_var_type])
	
	}
}
