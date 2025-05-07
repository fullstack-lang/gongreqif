// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type ALTERNATIVE_ID_WOP struct {
	// insertion point
	Name string
	IDENTIFIER string
}

func (from *ALTERNATIVE_ID) CopyBasicFields(to *ALTERNATIVE_ID) {
	// insertion point
	to.Name = from.Name
	to.IDENTIFIER = from.IDENTIFIER
}

type ATTRIBUTE_DEFINITION_BOOLEAN_WOP struct {
	// insertion point
	Name string
	DESC string
	IS_EDITABLE bool
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_BOOLEAN) CopyBasicFields(to *ATTRIBUTE_DEFINITION_BOOLEAN) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_DATE_WOP struct {
	// insertion point
	Name string
	DESC string
	IS_EDITABLE bool
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_DATE) CopyBasicFields(to *ATTRIBUTE_DEFINITION_DATE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_ENUMERATION_WOP struct {
	// insertion point
	Name string
	DESC string
	IS_EDITABLE bool
	LAST_CHANGE time.Time
	LONG_NAME string
	MULTI_VALUED bool
}

func (from *ATTRIBUTE_DEFINITION_ENUMERATION) CopyBasicFields(to *ATTRIBUTE_DEFINITION_ENUMERATION) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
	to.MULTI_VALUED = from.MULTI_VALUED
}

type ATTRIBUTE_DEFINITION_INTEGER_WOP struct {
	// insertion point
	Name string
	DESC string
	IS_EDITABLE bool
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_INTEGER) CopyBasicFields(to *ATTRIBUTE_DEFINITION_INTEGER) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_REAL_WOP struct {
	// insertion point
	Name string
	DESC string
	IS_EDITABLE bool
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_REAL) CopyBasicFields(to *ATTRIBUTE_DEFINITION_REAL) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_STRING_WOP struct {
	// insertion point
	Name string
	DESC string
	IS_EDITABLE bool
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_STRING) CopyBasicFields(to *ATTRIBUTE_DEFINITION_STRING) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_XHTML_WOP struct {
	// insertion point
	Name string
	DESC string
	IS_EDITABLE bool
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_XHTML) CopyBasicFields(to *ATTRIBUTE_DEFINITION_XHTML) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_VALUE_BOOLEAN_WOP struct {
	// insertion point
	Name string
	THE_VALUE bool
}

func (from *ATTRIBUTE_VALUE_BOOLEAN) CopyBasicFields(to *ATTRIBUTE_VALUE_BOOLEAN) {
	// insertion point
	to.Name = from.Name
	to.THE_VALUE = from.THE_VALUE
}

type ATTRIBUTE_VALUE_DATE_WOP struct {
	// insertion point
	Name string
	THE_VALUE time.Time
}

func (from *ATTRIBUTE_VALUE_DATE) CopyBasicFields(to *ATTRIBUTE_VALUE_DATE) {
	// insertion point
	to.Name = from.Name
	to.THE_VALUE = from.THE_VALUE
}

type ATTRIBUTE_VALUE_ENUMERATION_WOP struct {
	// insertion point
	Name string
}

func (from *ATTRIBUTE_VALUE_ENUMERATION) CopyBasicFields(to *ATTRIBUTE_VALUE_ENUMERATION) {
	// insertion point
	to.Name = from.Name
}

type ATTRIBUTE_VALUE_INTEGER_WOP struct {
	// insertion point
	Name string
}

func (from *ATTRIBUTE_VALUE_INTEGER) CopyBasicFields(to *ATTRIBUTE_VALUE_INTEGER) {
	// insertion point
	to.Name = from.Name
}

type ATTRIBUTE_VALUE_REAL_WOP struct {
	// insertion point
	Name string
	THE_VALUE float64
}

func (from *ATTRIBUTE_VALUE_REAL) CopyBasicFields(to *ATTRIBUTE_VALUE_REAL) {
	// insertion point
	to.Name = from.Name
	to.THE_VALUE = from.THE_VALUE
}

type ATTRIBUTE_VALUE_STRING_WOP struct {
	// insertion point
	Name string
	THE_VALUE string
}

func (from *ATTRIBUTE_VALUE_STRING) CopyBasicFields(to *ATTRIBUTE_VALUE_STRING) {
	// insertion point
	to.Name = from.Name
	to.THE_VALUE = from.THE_VALUE
}

type ATTRIBUTE_VALUE_XHTML_WOP struct {
	// insertion point
	Name string
	IS_SIMPLIFIED bool
}

func (from *ATTRIBUTE_VALUE_XHTML) CopyBasicFields(to *ATTRIBUTE_VALUE_XHTML) {
	// insertion point
	to.Name = from.Name
	to.IS_SIMPLIFIED = from.IS_SIMPLIFIED
}

type AnyType_WOP struct {
	// insertion point
	Name string
	InnerXML string
}

func (from *AnyType) CopyBasicFields(to *AnyType) {
	// insertion point
	to.Name = from.Name
	to.InnerXML = from.InnerXML
}

type DATATYPE_DEFINITION_BOOLEAN_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_BOOLEAN) CopyBasicFields(to *DATATYPE_DEFINITION_BOOLEAN) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type DATATYPE_DEFINITION_DATE_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_DATE) CopyBasicFields(to *DATATYPE_DEFINITION_DATE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type DATATYPE_DEFINITION_ENUMERATION_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_ENUMERATION) CopyBasicFields(to *DATATYPE_DEFINITION_ENUMERATION) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type DATATYPE_DEFINITION_INTEGER_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_INTEGER) CopyBasicFields(to *DATATYPE_DEFINITION_INTEGER) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type DATATYPE_DEFINITION_REAL_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
	MAX float64
	MIN float64
}

func (from *DATATYPE_DEFINITION_REAL) CopyBasicFields(to *DATATYPE_DEFINITION_REAL) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
	to.MAX = from.MAX
	to.MIN = from.MIN
}

type DATATYPE_DEFINITION_STRING_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_STRING) CopyBasicFields(to *DATATYPE_DEFINITION_STRING) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type DATATYPE_DEFINITION_XHTML_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_XHTML) CopyBasicFields(to *DATATYPE_DEFINITION_XHTML) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type EMBEDDED_VALUE_WOP struct {
	// insertion point
	Name string
	OTHER_CONTENT string
}

func (from *EMBEDDED_VALUE) CopyBasicFields(to *EMBEDDED_VALUE) {
	// insertion point
	to.Name = from.Name
	to.OTHER_CONTENT = from.OTHER_CONTENT
}

type ENUM_VALUE_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *ENUM_VALUE) CopyBasicFields(to *ENUM_VALUE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type RELATION_GROUP_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *RELATION_GROUP) CopyBasicFields(to *RELATION_GROUP) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type RELATION_GROUP_TYPE_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *RELATION_GROUP_TYPE) CopyBasicFields(to *RELATION_GROUP_TYPE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type REQ_IF_WOP struct {
	// insertion point
	Name string
	Lang string
}

func (from *REQ_IF) CopyBasicFields(to *REQ_IF) {
	// insertion point
	to.Name = from.Name
	to.Lang = from.Lang
}

type REQ_IF_CONTENT_WOP struct {
	// insertion point
	Name string
}

func (from *REQ_IF_CONTENT) CopyBasicFields(to *REQ_IF_CONTENT) {
	// insertion point
	to.Name = from.Name
}

type REQ_IF_HEADER_WOP struct {
	// insertion point
	Name string
	COMMENT string
	CREATION_TIME time.Time
	REPOSITORY_ID string
	REQ_IF_TOOL_ID string
	REQ_IF_VERSION string
	SOURCE_TOOL_ID string
	TITLE string
}

func (from *REQ_IF_HEADER) CopyBasicFields(to *REQ_IF_HEADER) {
	// insertion point
	to.Name = from.Name
	to.COMMENT = from.COMMENT
	to.CREATION_TIME = from.CREATION_TIME
	to.REPOSITORY_ID = from.REPOSITORY_ID
	to.REQ_IF_TOOL_ID = from.REQ_IF_TOOL_ID
	to.REQ_IF_VERSION = from.REQ_IF_VERSION
	to.SOURCE_TOOL_ID = from.SOURCE_TOOL_ID
	to.TITLE = from.TITLE
}

type REQ_IF_TOOL_EXTENSION_WOP struct {
	// insertion point
	Name string
}

func (from *REQ_IF_TOOL_EXTENSION) CopyBasicFields(to *REQ_IF_TOOL_EXTENSION) {
	// insertion point
	to.Name = from.Name
}

type SPECIFICATION_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *SPECIFICATION) CopyBasicFields(to *SPECIFICATION) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPECIFICATION_TYPE_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *SPECIFICATION_TYPE) CopyBasicFields(to *SPECIFICATION_TYPE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_HIERARCHY_WOP struct {
	// insertion point
	Name string
	DESC string
	IS_EDITABLE bool
	IS_TABLE_INTERNAL bool
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *SPEC_HIERARCHY) CopyBasicFields(to *SPEC_HIERARCHY) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IS_EDITABLE = from.IS_EDITABLE
	to.IS_TABLE_INTERNAL = from.IS_TABLE_INTERNAL
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_OBJECT_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *SPEC_OBJECT) CopyBasicFields(to *SPEC_OBJECT) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_OBJECT_TYPE_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *SPEC_OBJECT_TYPE) CopyBasicFields(to *SPEC_OBJECT_TYPE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_RELATION_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *SPEC_RELATION) CopyBasicFields(to *SPEC_RELATION) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_RELATION_TYPE_WOP struct {
	// insertion point
	Name string
	DESC string
	LAST_CHANGE time.Time
	LONG_NAME string
}

func (from *SPEC_RELATION_TYPE) CopyBasicFields(to *SPEC_RELATION_TYPE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type XHTML_CONTENT_WOP struct {
	// insertion point
	Name string
}

func (from *XHTML_CONTENT) CopyBasicFields(to *XHTML_CONTENT) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_InlPres_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_InlPres_type) CopyBasicFields(to *Xhtml_InlPres_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_a_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_a_type) CopyBasicFields(to *Xhtml_a_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_abbr_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_abbr_type) CopyBasicFields(to *Xhtml_abbr_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_acronym_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_acronym_type) CopyBasicFields(to *Xhtml_acronym_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_address_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_address_type) CopyBasicFields(to *Xhtml_address_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_blockquote_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_blockquote_type) CopyBasicFields(to *Xhtml_blockquote_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_br_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_br_type) CopyBasicFields(to *Xhtml_br_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_caption_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_caption_type) CopyBasicFields(to *Xhtml_caption_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_cite_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_cite_type) CopyBasicFields(to *Xhtml_cite_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_code_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_code_type) CopyBasicFields(to *Xhtml_code_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_col_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_col_type) CopyBasicFields(to *Xhtml_col_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_colgroup_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_colgroup_type) CopyBasicFields(to *Xhtml_colgroup_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_dd_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_dd_type) CopyBasicFields(to *Xhtml_dd_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_dfn_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_dfn_type) CopyBasicFields(to *Xhtml_dfn_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_div_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_div_type) CopyBasicFields(to *Xhtml_div_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_dl_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_dl_type) CopyBasicFields(to *Xhtml_dl_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_dt_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_dt_type) CopyBasicFields(to *Xhtml_dt_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_edit_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_edit_type) CopyBasicFields(to *Xhtml_edit_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_em_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_em_type) CopyBasicFields(to *Xhtml_em_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_h1_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_h1_type) CopyBasicFields(to *Xhtml_h1_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_h2_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_h2_type) CopyBasicFields(to *Xhtml_h2_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_h3_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_h3_type) CopyBasicFields(to *Xhtml_h3_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_h4_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_h4_type) CopyBasicFields(to *Xhtml_h4_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_h5_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_h5_type) CopyBasicFields(to *Xhtml_h5_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_h6_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_h6_type) CopyBasicFields(to *Xhtml_h6_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_heading_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_heading_type) CopyBasicFields(to *Xhtml_heading_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_hr_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_hr_type) CopyBasicFields(to *Xhtml_hr_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_kbd_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_kbd_type) CopyBasicFields(to *Xhtml_kbd_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_li_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_li_type) CopyBasicFields(to *Xhtml_li_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_object_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_object_type) CopyBasicFields(to *Xhtml_object_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_ol_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_ol_type) CopyBasicFields(to *Xhtml_ol_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_p_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_p_type) CopyBasicFields(to *Xhtml_p_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_param_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_param_type) CopyBasicFields(to *Xhtml_param_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_pre_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_pre_type) CopyBasicFields(to *Xhtml_pre_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_q_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_q_type) CopyBasicFields(to *Xhtml_q_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_samp_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_samp_type) CopyBasicFields(to *Xhtml_samp_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_span_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_span_type) CopyBasicFields(to *Xhtml_span_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_strong_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_strong_type) CopyBasicFields(to *Xhtml_strong_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_table_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_table_type) CopyBasicFields(to *Xhtml_table_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_tbody_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_tbody_type) CopyBasicFields(to *Xhtml_tbody_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_td_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_td_type) CopyBasicFields(to *Xhtml_td_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_tfoot_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_tfoot_type) CopyBasicFields(to *Xhtml_tfoot_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_th_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_th_type) CopyBasicFields(to *Xhtml_th_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_thead_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_thead_type) CopyBasicFields(to *Xhtml_thead_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_tr_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_tr_type) CopyBasicFields(to *Xhtml_tr_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_ul_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_ul_type) CopyBasicFields(to *Xhtml_ul_type) {
	// insertion point
	to.Name = from.Name
}

type Xhtml_var_type_WOP struct {
	// insertion point
	Name string
}

func (from *Xhtml_var_type) CopyBasicFields(to *Xhtml_var_type) {
	// insertion point
	to.Name = from.Name
}

