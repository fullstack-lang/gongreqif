// generated code - do not edit
package orm

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/fullstack-lang/gongreqif/go/db"
)

// Ensure DBLite implements DBInterface
var _ db.DBInterface = &DBLite{}

// DBLite is an in-memory database implementation of DBInterface
type DBLite struct {
	// Mutex to protect shared resources
	mu sync.RWMutex

	// insertion point definitions

	alternative_idDBs map[uint]*ALTERNATIVE_IDDB

	nextIDALTERNATIVE_IDDB uint

	attribute_definition_booleanDBs map[uint]*ATTRIBUTE_DEFINITION_BOOLEANDB

	nextIDATTRIBUTE_DEFINITION_BOOLEANDB uint

	attribute_definition_dateDBs map[uint]*ATTRIBUTE_DEFINITION_DATEDB

	nextIDATTRIBUTE_DEFINITION_DATEDB uint

	attribute_definition_enumerationDBs map[uint]*ATTRIBUTE_DEFINITION_ENUMERATIONDB

	nextIDATTRIBUTE_DEFINITION_ENUMERATIONDB uint

	attribute_definition_integerDBs map[uint]*ATTRIBUTE_DEFINITION_INTEGERDB

	nextIDATTRIBUTE_DEFINITION_INTEGERDB uint

	attribute_definition_realDBs map[uint]*ATTRIBUTE_DEFINITION_REALDB

	nextIDATTRIBUTE_DEFINITION_REALDB uint

	attribute_definition_stringDBs map[uint]*ATTRIBUTE_DEFINITION_STRINGDB

	nextIDATTRIBUTE_DEFINITION_STRINGDB uint

	attribute_definition_xhtmlDBs map[uint]*ATTRIBUTE_DEFINITION_XHTMLDB

	nextIDATTRIBUTE_DEFINITION_XHTMLDB uint

	attribute_value_booleanDBs map[uint]*ATTRIBUTE_VALUE_BOOLEANDB

	nextIDATTRIBUTE_VALUE_BOOLEANDB uint

	attribute_value_dateDBs map[uint]*ATTRIBUTE_VALUE_DATEDB

	nextIDATTRIBUTE_VALUE_DATEDB uint

	attribute_value_enumerationDBs map[uint]*ATTRIBUTE_VALUE_ENUMERATIONDB

	nextIDATTRIBUTE_VALUE_ENUMERATIONDB uint

	attribute_value_integerDBs map[uint]*ATTRIBUTE_VALUE_INTEGERDB

	nextIDATTRIBUTE_VALUE_INTEGERDB uint

	attribute_value_realDBs map[uint]*ATTRIBUTE_VALUE_REALDB

	nextIDATTRIBUTE_VALUE_REALDB uint

	attribute_value_stringDBs map[uint]*ATTRIBUTE_VALUE_STRINGDB

	nextIDATTRIBUTE_VALUE_STRINGDB uint

	attribute_value_xhtmlDBs map[uint]*ATTRIBUTE_VALUE_XHTMLDB

	nextIDATTRIBUTE_VALUE_XHTMLDB uint

	anytypeDBs map[uint]*AnyTypeDB

	nextIDAnyTypeDB uint

	datatype_definition_booleanDBs map[uint]*DATATYPE_DEFINITION_BOOLEANDB

	nextIDDATATYPE_DEFINITION_BOOLEANDB uint

	datatype_definition_dateDBs map[uint]*DATATYPE_DEFINITION_DATEDB

	nextIDDATATYPE_DEFINITION_DATEDB uint

	datatype_definition_enumerationDBs map[uint]*DATATYPE_DEFINITION_ENUMERATIONDB

	nextIDDATATYPE_DEFINITION_ENUMERATIONDB uint

	datatype_definition_integerDBs map[uint]*DATATYPE_DEFINITION_INTEGERDB

	nextIDDATATYPE_DEFINITION_INTEGERDB uint

	datatype_definition_realDBs map[uint]*DATATYPE_DEFINITION_REALDB

	nextIDDATATYPE_DEFINITION_REALDB uint

	datatype_definition_stringDBs map[uint]*DATATYPE_DEFINITION_STRINGDB

	nextIDDATATYPE_DEFINITION_STRINGDB uint

	datatype_definition_xhtmlDBs map[uint]*DATATYPE_DEFINITION_XHTMLDB

	nextIDDATATYPE_DEFINITION_XHTMLDB uint

	embedded_valueDBs map[uint]*EMBEDDED_VALUEDB

	nextIDEMBEDDED_VALUEDB uint

	enum_valueDBs map[uint]*ENUM_VALUEDB

	nextIDENUM_VALUEDB uint

	relation_groupDBs map[uint]*RELATION_GROUPDB

	nextIDRELATION_GROUPDB uint

	relation_group_typeDBs map[uint]*RELATION_GROUP_TYPEDB

	nextIDRELATION_GROUP_TYPEDB uint

	req_ifDBs map[uint]*REQ_IFDB

	nextIDREQ_IFDB uint

	req_if_contentDBs map[uint]*REQ_IF_CONTENTDB

	nextIDREQ_IF_CONTENTDB uint

	req_if_headerDBs map[uint]*REQ_IF_HEADERDB

	nextIDREQ_IF_HEADERDB uint

	req_if_tool_extensionDBs map[uint]*REQ_IF_TOOL_EXTENSIONDB

	nextIDREQ_IF_TOOL_EXTENSIONDB uint

	specificationDBs map[uint]*SPECIFICATIONDB

	nextIDSPECIFICATIONDB uint

	specification_typeDBs map[uint]*SPECIFICATION_TYPEDB

	nextIDSPECIFICATION_TYPEDB uint

	spec_hierarchyDBs map[uint]*SPEC_HIERARCHYDB

	nextIDSPEC_HIERARCHYDB uint

	spec_objectDBs map[uint]*SPEC_OBJECTDB

	nextIDSPEC_OBJECTDB uint

	spec_object_typeDBs map[uint]*SPEC_OBJECT_TYPEDB

	nextIDSPEC_OBJECT_TYPEDB uint

	spec_relationDBs map[uint]*SPEC_RELATIONDB

	nextIDSPEC_RELATIONDB uint

	spec_relation_typeDBs map[uint]*SPEC_RELATION_TYPEDB

	nextIDSPEC_RELATION_TYPEDB uint

	xhtml_contentDBs map[uint]*XHTML_CONTENTDB

	nextIDXHTML_CONTENTDB uint

	xhtml_inlpres_typeDBs map[uint]*Xhtml_InlPres_typeDB

	nextIDXhtml_InlPres_typeDB uint

	xhtml_a_typeDBs map[uint]*Xhtml_a_typeDB

	nextIDXhtml_a_typeDB uint

	xhtml_abbr_typeDBs map[uint]*Xhtml_abbr_typeDB

	nextIDXhtml_abbr_typeDB uint

	xhtml_acronym_typeDBs map[uint]*Xhtml_acronym_typeDB

	nextIDXhtml_acronym_typeDB uint

	xhtml_address_typeDBs map[uint]*Xhtml_address_typeDB

	nextIDXhtml_address_typeDB uint

	xhtml_blockquote_typeDBs map[uint]*Xhtml_blockquote_typeDB

	nextIDXhtml_blockquote_typeDB uint

	xhtml_br_typeDBs map[uint]*Xhtml_br_typeDB

	nextIDXhtml_br_typeDB uint

	xhtml_caption_typeDBs map[uint]*Xhtml_caption_typeDB

	nextIDXhtml_caption_typeDB uint

	xhtml_cite_typeDBs map[uint]*Xhtml_cite_typeDB

	nextIDXhtml_cite_typeDB uint

	xhtml_code_typeDBs map[uint]*Xhtml_code_typeDB

	nextIDXhtml_code_typeDB uint

	xhtml_col_typeDBs map[uint]*Xhtml_col_typeDB

	nextIDXhtml_col_typeDB uint

	xhtml_colgroup_typeDBs map[uint]*Xhtml_colgroup_typeDB

	nextIDXhtml_colgroup_typeDB uint

	xhtml_dd_typeDBs map[uint]*Xhtml_dd_typeDB

	nextIDXhtml_dd_typeDB uint

	xhtml_dfn_typeDBs map[uint]*Xhtml_dfn_typeDB

	nextIDXhtml_dfn_typeDB uint

	xhtml_div_typeDBs map[uint]*Xhtml_div_typeDB

	nextIDXhtml_div_typeDB uint

	xhtml_dl_typeDBs map[uint]*Xhtml_dl_typeDB

	nextIDXhtml_dl_typeDB uint

	xhtml_dt_typeDBs map[uint]*Xhtml_dt_typeDB

	nextIDXhtml_dt_typeDB uint

	xhtml_edit_typeDBs map[uint]*Xhtml_edit_typeDB

	nextIDXhtml_edit_typeDB uint

	xhtml_em_typeDBs map[uint]*Xhtml_em_typeDB

	nextIDXhtml_em_typeDB uint

	xhtml_h1_typeDBs map[uint]*Xhtml_h1_typeDB

	nextIDXhtml_h1_typeDB uint

	xhtml_h2_typeDBs map[uint]*Xhtml_h2_typeDB

	nextIDXhtml_h2_typeDB uint

	xhtml_h3_typeDBs map[uint]*Xhtml_h3_typeDB

	nextIDXhtml_h3_typeDB uint

	xhtml_h4_typeDBs map[uint]*Xhtml_h4_typeDB

	nextIDXhtml_h4_typeDB uint

	xhtml_h5_typeDBs map[uint]*Xhtml_h5_typeDB

	nextIDXhtml_h5_typeDB uint

	xhtml_h6_typeDBs map[uint]*Xhtml_h6_typeDB

	nextIDXhtml_h6_typeDB uint

	xhtml_heading_typeDBs map[uint]*Xhtml_heading_typeDB

	nextIDXhtml_heading_typeDB uint

	xhtml_hr_typeDBs map[uint]*Xhtml_hr_typeDB

	nextIDXhtml_hr_typeDB uint

	xhtml_kbd_typeDBs map[uint]*Xhtml_kbd_typeDB

	nextIDXhtml_kbd_typeDB uint

	xhtml_li_typeDBs map[uint]*Xhtml_li_typeDB

	nextIDXhtml_li_typeDB uint

	xhtml_object_typeDBs map[uint]*Xhtml_object_typeDB

	nextIDXhtml_object_typeDB uint

	xhtml_ol_typeDBs map[uint]*Xhtml_ol_typeDB

	nextIDXhtml_ol_typeDB uint

	xhtml_p_typeDBs map[uint]*Xhtml_p_typeDB

	nextIDXhtml_p_typeDB uint

	xhtml_param_typeDBs map[uint]*Xhtml_param_typeDB

	nextIDXhtml_param_typeDB uint

	xhtml_pre_typeDBs map[uint]*Xhtml_pre_typeDB

	nextIDXhtml_pre_typeDB uint

	xhtml_q_typeDBs map[uint]*Xhtml_q_typeDB

	nextIDXhtml_q_typeDB uint

	xhtml_samp_typeDBs map[uint]*Xhtml_samp_typeDB

	nextIDXhtml_samp_typeDB uint

	xhtml_span_typeDBs map[uint]*Xhtml_span_typeDB

	nextIDXhtml_span_typeDB uint

	xhtml_strong_typeDBs map[uint]*Xhtml_strong_typeDB

	nextIDXhtml_strong_typeDB uint

	xhtml_table_typeDBs map[uint]*Xhtml_table_typeDB

	nextIDXhtml_table_typeDB uint

	xhtml_tbody_typeDBs map[uint]*Xhtml_tbody_typeDB

	nextIDXhtml_tbody_typeDB uint

	xhtml_td_typeDBs map[uint]*Xhtml_td_typeDB

	nextIDXhtml_td_typeDB uint

	xhtml_tfoot_typeDBs map[uint]*Xhtml_tfoot_typeDB

	nextIDXhtml_tfoot_typeDB uint

	xhtml_th_typeDBs map[uint]*Xhtml_th_typeDB

	nextIDXhtml_th_typeDB uint

	xhtml_thead_typeDBs map[uint]*Xhtml_thead_typeDB

	nextIDXhtml_thead_typeDB uint

	xhtml_tr_typeDBs map[uint]*Xhtml_tr_typeDB

	nextIDXhtml_tr_typeDB uint

	xhtml_ul_typeDBs map[uint]*Xhtml_ul_typeDB

	nextIDXhtml_ul_typeDB uint

	xhtml_var_typeDBs map[uint]*Xhtml_var_typeDB

	nextIDXhtml_var_typeDB uint
}

// NewDBLite creates a new instance of DBLite
func NewDBLite() *DBLite {
	return &DBLite{
		// insertion point maps init

		alternative_idDBs: make(map[uint]*ALTERNATIVE_IDDB),

		attribute_definition_booleanDBs: make(map[uint]*ATTRIBUTE_DEFINITION_BOOLEANDB),

		attribute_definition_dateDBs: make(map[uint]*ATTRIBUTE_DEFINITION_DATEDB),

		attribute_definition_enumerationDBs: make(map[uint]*ATTRIBUTE_DEFINITION_ENUMERATIONDB),

		attribute_definition_integerDBs: make(map[uint]*ATTRIBUTE_DEFINITION_INTEGERDB),

		attribute_definition_realDBs: make(map[uint]*ATTRIBUTE_DEFINITION_REALDB),

		attribute_definition_stringDBs: make(map[uint]*ATTRIBUTE_DEFINITION_STRINGDB),

		attribute_definition_xhtmlDBs: make(map[uint]*ATTRIBUTE_DEFINITION_XHTMLDB),

		attribute_value_booleanDBs: make(map[uint]*ATTRIBUTE_VALUE_BOOLEANDB),

		attribute_value_dateDBs: make(map[uint]*ATTRIBUTE_VALUE_DATEDB),

		attribute_value_enumerationDBs: make(map[uint]*ATTRIBUTE_VALUE_ENUMERATIONDB),

		attribute_value_integerDBs: make(map[uint]*ATTRIBUTE_VALUE_INTEGERDB),

		attribute_value_realDBs: make(map[uint]*ATTRIBUTE_VALUE_REALDB),

		attribute_value_stringDBs: make(map[uint]*ATTRIBUTE_VALUE_STRINGDB),

		attribute_value_xhtmlDBs: make(map[uint]*ATTRIBUTE_VALUE_XHTMLDB),

		anytypeDBs: make(map[uint]*AnyTypeDB),

		datatype_definition_booleanDBs: make(map[uint]*DATATYPE_DEFINITION_BOOLEANDB),

		datatype_definition_dateDBs: make(map[uint]*DATATYPE_DEFINITION_DATEDB),

		datatype_definition_enumerationDBs: make(map[uint]*DATATYPE_DEFINITION_ENUMERATIONDB),

		datatype_definition_integerDBs: make(map[uint]*DATATYPE_DEFINITION_INTEGERDB),

		datatype_definition_realDBs: make(map[uint]*DATATYPE_DEFINITION_REALDB),

		datatype_definition_stringDBs: make(map[uint]*DATATYPE_DEFINITION_STRINGDB),

		datatype_definition_xhtmlDBs: make(map[uint]*DATATYPE_DEFINITION_XHTMLDB),

		embedded_valueDBs: make(map[uint]*EMBEDDED_VALUEDB),

		enum_valueDBs: make(map[uint]*ENUM_VALUEDB),

		relation_groupDBs: make(map[uint]*RELATION_GROUPDB),

		relation_group_typeDBs: make(map[uint]*RELATION_GROUP_TYPEDB),

		req_ifDBs: make(map[uint]*REQ_IFDB),

		req_if_contentDBs: make(map[uint]*REQ_IF_CONTENTDB),

		req_if_headerDBs: make(map[uint]*REQ_IF_HEADERDB),

		req_if_tool_extensionDBs: make(map[uint]*REQ_IF_TOOL_EXTENSIONDB),

		specificationDBs: make(map[uint]*SPECIFICATIONDB),

		specification_typeDBs: make(map[uint]*SPECIFICATION_TYPEDB),

		spec_hierarchyDBs: make(map[uint]*SPEC_HIERARCHYDB),

		spec_objectDBs: make(map[uint]*SPEC_OBJECTDB),

		spec_object_typeDBs: make(map[uint]*SPEC_OBJECT_TYPEDB),

		spec_relationDBs: make(map[uint]*SPEC_RELATIONDB),

		spec_relation_typeDBs: make(map[uint]*SPEC_RELATION_TYPEDB),

		xhtml_contentDBs: make(map[uint]*XHTML_CONTENTDB),

		xhtml_inlpres_typeDBs: make(map[uint]*Xhtml_InlPres_typeDB),

		xhtml_a_typeDBs: make(map[uint]*Xhtml_a_typeDB),

		xhtml_abbr_typeDBs: make(map[uint]*Xhtml_abbr_typeDB),

		xhtml_acronym_typeDBs: make(map[uint]*Xhtml_acronym_typeDB),

		xhtml_address_typeDBs: make(map[uint]*Xhtml_address_typeDB),

		xhtml_blockquote_typeDBs: make(map[uint]*Xhtml_blockquote_typeDB),

		xhtml_br_typeDBs: make(map[uint]*Xhtml_br_typeDB),

		xhtml_caption_typeDBs: make(map[uint]*Xhtml_caption_typeDB),

		xhtml_cite_typeDBs: make(map[uint]*Xhtml_cite_typeDB),

		xhtml_code_typeDBs: make(map[uint]*Xhtml_code_typeDB),

		xhtml_col_typeDBs: make(map[uint]*Xhtml_col_typeDB),

		xhtml_colgroup_typeDBs: make(map[uint]*Xhtml_colgroup_typeDB),

		xhtml_dd_typeDBs: make(map[uint]*Xhtml_dd_typeDB),

		xhtml_dfn_typeDBs: make(map[uint]*Xhtml_dfn_typeDB),

		xhtml_div_typeDBs: make(map[uint]*Xhtml_div_typeDB),

		xhtml_dl_typeDBs: make(map[uint]*Xhtml_dl_typeDB),

		xhtml_dt_typeDBs: make(map[uint]*Xhtml_dt_typeDB),

		xhtml_edit_typeDBs: make(map[uint]*Xhtml_edit_typeDB),

		xhtml_em_typeDBs: make(map[uint]*Xhtml_em_typeDB),

		xhtml_h1_typeDBs: make(map[uint]*Xhtml_h1_typeDB),

		xhtml_h2_typeDBs: make(map[uint]*Xhtml_h2_typeDB),

		xhtml_h3_typeDBs: make(map[uint]*Xhtml_h3_typeDB),

		xhtml_h4_typeDBs: make(map[uint]*Xhtml_h4_typeDB),

		xhtml_h5_typeDBs: make(map[uint]*Xhtml_h5_typeDB),

		xhtml_h6_typeDBs: make(map[uint]*Xhtml_h6_typeDB),

		xhtml_heading_typeDBs: make(map[uint]*Xhtml_heading_typeDB),

		xhtml_hr_typeDBs: make(map[uint]*Xhtml_hr_typeDB),

		xhtml_kbd_typeDBs: make(map[uint]*Xhtml_kbd_typeDB),

		xhtml_li_typeDBs: make(map[uint]*Xhtml_li_typeDB),

		xhtml_object_typeDBs: make(map[uint]*Xhtml_object_typeDB),

		xhtml_ol_typeDBs: make(map[uint]*Xhtml_ol_typeDB),

		xhtml_p_typeDBs: make(map[uint]*Xhtml_p_typeDB),

		xhtml_param_typeDBs: make(map[uint]*Xhtml_param_typeDB),

		xhtml_pre_typeDBs: make(map[uint]*Xhtml_pre_typeDB),

		xhtml_q_typeDBs: make(map[uint]*Xhtml_q_typeDB),

		xhtml_samp_typeDBs: make(map[uint]*Xhtml_samp_typeDB),

		xhtml_span_typeDBs: make(map[uint]*Xhtml_span_typeDB),

		xhtml_strong_typeDBs: make(map[uint]*Xhtml_strong_typeDB),

		xhtml_table_typeDBs: make(map[uint]*Xhtml_table_typeDB),

		xhtml_tbody_typeDBs: make(map[uint]*Xhtml_tbody_typeDB),

		xhtml_td_typeDBs: make(map[uint]*Xhtml_td_typeDB),

		xhtml_tfoot_typeDBs: make(map[uint]*Xhtml_tfoot_typeDB),

		xhtml_th_typeDBs: make(map[uint]*Xhtml_th_typeDB),

		xhtml_thead_typeDBs: make(map[uint]*Xhtml_thead_typeDB),

		xhtml_tr_typeDBs: make(map[uint]*Xhtml_tr_typeDB),

		xhtml_ul_typeDBs: make(map[uint]*Xhtml_ul_typeDB),

		xhtml_var_typeDBs: make(map[uint]*Xhtml_var_typeDB),
	}
}

// Create inserts a new record into the database
func (db *DBLite) Create(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point create
	case *ALTERNATIVE_IDDB:
		db.nextIDALTERNATIVE_IDDB++
		v.ID = db.nextIDALTERNATIVE_IDDB
		db.alternative_idDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		db.nextIDATTRIBUTE_DEFINITION_BOOLEANDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_BOOLEANDB
		db.attribute_definition_booleanDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_DATEDB:
		db.nextIDATTRIBUTE_DEFINITION_DATEDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_DATEDB
		db.attribute_definition_dateDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		db.nextIDATTRIBUTE_DEFINITION_ENUMERATIONDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_ENUMERATIONDB
		db.attribute_definition_enumerationDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		db.nextIDATTRIBUTE_DEFINITION_INTEGERDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_INTEGERDB
		db.attribute_definition_integerDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_REALDB:
		db.nextIDATTRIBUTE_DEFINITION_REALDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_REALDB
		db.attribute_definition_realDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		db.nextIDATTRIBUTE_DEFINITION_STRINGDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_STRINGDB
		db.attribute_definition_stringDBs[v.ID] = v
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		db.nextIDATTRIBUTE_DEFINITION_XHTMLDB++
		v.ID = db.nextIDATTRIBUTE_DEFINITION_XHTMLDB
		db.attribute_definition_xhtmlDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		db.nextIDATTRIBUTE_VALUE_BOOLEANDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_BOOLEANDB
		db.attribute_value_booleanDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_DATEDB:
		db.nextIDATTRIBUTE_VALUE_DATEDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_DATEDB
		db.attribute_value_dateDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		db.nextIDATTRIBUTE_VALUE_ENUMERATIONDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_ENUMERATIONDB
		db.attribute_value_enumerationDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_INTEGERDB:
		db.nextIDATTRIBUTE_VALUE_INTEGERDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_INTEGERDB
		db.attribute_value_integerDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_REALDB:
		db.nextIDATTRIBUTE_VALUE_REALDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_REALDB
		db.attribute_value_realDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_STRINGDB:
		db.nextIDATTRIBUTE_VALUE_STRINGDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_STRINGDB
		db.attribute_value_stringDBs[v.ID] = v
	case *ATTRIBUTE_VALUE_XHTMLDB:
		db.nextIDATTRIBUTE_VALUE_XHTMLDB++
		v.ID = db.nextIDATTRIBUTE_VALUE_XHTMLDB
		db.attribute_value_xhtmlDBs[v.ID] = v
	case *AnyTypeDB:
		db.nextIDAnyTypeDB++
		v.ID = db.nextIDAnyTypeDB
		db.anytypeDBs[v.ID] = v
	case *DATATYPE_DEFINITION_BOOLEANDB:
		db.nextIDDATATYPE_DEFINITION_BOOLEANDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_BOOLEANDB
		db.datatype_definition_booleanDBs[v.ID] = v
	case *DATATYPE_DEFINITION_DATEDB:
		db.nextIDDATATYPE_DEFINITION_DATEDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_DATEDB
		db.datatype_definition_dateDBs[v.ID] = v
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		db.nextIDDATATYPE_DEFINITION_ENUMERATIONDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_ENUMERATIONDB
		db.datatype_definition_enumerationDBs[v.ID] = v
	case *DATATYPE_DEFINITION_INTEGERDB:
		db.nextIDDATATYPE_DEFINITION_INTEGERDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_INTEGERDB
		db.datatype_definition_integerDBs[v.ID] = v
	case *DATATYPE_DEFINITION_REALDB:
		db.nextIDDATATYPE_DEFINITION_REALDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_REALDB
		db.datatype_definition_realDBs[v.ID] = v
	case *DATATYPE_DEFINITION_STRINGDB:
		db.nextIDDATATYPE_DEFINITION_STRINGDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_STRINGDB
		db.datatype_definition_stringDBs[v.ID] = v
	case *DATATYPE_DEFINITION_XHTMLDB:
		db.nextIDDATATYPE_DEFINITION_XHTMLDB++
		v.ID = db.nextIDDATATYPE_DEFINITION_XHTMLDB
		db.datatype_definition_xhtmlDBs[v.ID] = v
	case *EMBEDDED_VALUEDB:
		db.nextIDEMBEDDED_VALUEDB++
		v.ID = db.nextIDEMBEDDED_VALUEDB
		db.embedded_valueDBs[v.ID] = v
	case *ENUM_VALUEDB:
		db.nextIDENUM_VALUEDB++
		v.ID = db.nextIDENUM_VALUEDB
		db.enum_valueDBs[v.ID] = v
	case *RELATION_GROUPDB:
		db.nextIDRELATION_GROUPDB++
		v.ID = db.nextIDRELATION_GROUPDB
		db.relation_groupDBs[v.ID] = v
	case *RELATION_GROUP_TYPEDB:
		db.nextIDRELATION_GROUP_TYPEDB++
		v.ID = db.nextIDRELATION_GROUP_TYPEDB
		db.relation_group_typeDBs[v.ID] = v
	case *REQ_IFDB:
		db.nextIDREQ_IFDB++
		v.ID = db.nextIDREQ_IFDB
		db.req_ifDBs[v.ID] = v
	case *REQ_IF_CONTENTDB:
		db.nextIDREQ_IF_CONTENTDB++
		v.ID = db.nextIDREQ_IF_CONTENTDB
		db.req_if_contentDBs[v.ID] = v
	case *REQ_IF_HEADERDB:
		db.nextIDREQ_IF_HEADERDB++
		v.ID = db.nextIDREQ_IF_HEADERDB
		db.req_if_headerDBs[v.ID] = v
	case *REQ_IF_TOOL_EXTENSIONDB:
		db.nextIDREQ_IF_TOOL_EXTENSIONDB++
		v.ID = db.nextIDREQ_IF_TOOL_EXTENSIONDB
		db.req_if_tool_extensionDBs[v.ID] = v
	case *SPECIFICATIONDB:
		db.nextIDSPECIFICATIONDB++
		v.ID = db.nextIDSPECIFICATIONDB
		db.specificationDBs[v.ID] = v
	case *SPECIFICATION_TYPEDB:
		db.nextIDSPECIFICATION_TYPEDB++
		v.ID = db.nextIDSPECIFICATION_TYPEDB
		db.specification_typeDBs[v.ID] = v
	case *SPEC_HIERARCHYDB:
		db.nextIDSPEC_HIERARCHYDB++
		v.ID = db.nextIDSPEC_HIERARCHYDB
		db.spec_hierarchyDBs[v.ID] = v
	case *SPEC_OBJECTDB:
		db.nextIDSPEC_OBJECTDB++
		v.ID = db.nextIDSPEC_OBJECTDB
		db.spec_objectDBs[v.ID] = v
	case *SPEC_OBJECT_TYPEDB:
		db.nextIDSPEC_OBJECT_TYPEDB++
		v.ID = db.nextIDSPEC_OBJECT_TYPEDB
		db.spec_object_typeDBs[v.ID] = v
	case *SPEC_RELATIONDB:
		db.nextIDSPEC_RELATIONDB++
		v.ID = db.nextIDSPEC_RELATIONDB
		db.spec_relationDBs[v.ID] = v
	case *SPEC_RELATION_TYPEDB:
		db.nextIDSPEC_RELATION_TYPEDB++
		v.ID = db.nextIDSPEC_RELATION_TYPEDB
		db.spec_relation_typeDBs[v.ID] = v
	case *XHTML_CONTENTDB:
		db.nextIDXHTML_CONTENTDB++
		v.ID = db.nextIDXHTML_CONTENTDB
		db.xhtml_contentDBs[v.ID] = v
	case *Xhtml_InlPres_typeDB:
		db.nextIDXhtml_InlPres_typeDB++
		v.ID = db.nextIDXhtml_InlPres_typeDB
		db.xhtml_inlpres_typeDBs[v.ID] = v
	case *Xhtml_a_typeDB:
		db.nextIDXhtml_a_typeDB++
		v.ID = db.nextIDXhtml_a_typeDB
		db.xhtml_a_typeDBs[v.ID] = v
	case *Xhtml_abbr_typeDB:
		db.nextIDXhtml_abbr_typeDB++
		v.ID = db.nextIDXhtml_abbr_typeDB
		db.xhtml_abbr_typeDBs[v.ID] = v
	case *Xhtml_acronym_typeDB:
		db.nextIDXhtml_acronym_typeDB++
		v.ID = db.nextIDXhtml_acronym_typeDB
		db.xhtml_acronym_typeDBs[v.ID] = v
	case *Xhtml_address_typeDB:
		db.nextIDXhtml_address_typeDB++
		v.ID = db.nextIDXhtml_address_typeDB
		db.xhtml_address_typeDBs[v.ID] = v
	case *Xhtml_blockquote_typeDB:
		db.nextIDXhtml_blockquote_typeDB++
		v.ID = db.nextIDXhtml_blockquote_typeDB
		db.xhtml_blockquote_typeDBs[v.ID] = v
	case *Xhtml_br_typeDB:
		db.nextIDXhtml_br_typeDB++
		v.ID = db.nextIDXhtml_br_typeDB
		db.xhtml_br_typeDBs[v.ID] = v
	case *Xhtml_caption_typeDB:
		db.nextIDXhtml_caption_typeDB++
		v.ID = db.nextIDXhtml_caption_typeDB
		db.xhtml_caption_typeDBs[v.ID] = v
	case *Xhtml_cite_typeDB:
		db.nextIDXhtml_cite_typeDB++
		v.ID = db.nextIDXhtml_cite_typeDB
		db.xhtml_cite_typeDBs[v.ID] = v
	case *Xhtml_code_typeDB:
		db.nextIDXhtml_code_typeDB++
		v.ID = db.nextIDXhtml_code_typeDB
		db.xhtml_code_typeDBs[v.ID] = v
	case *Xhtml_col_typeDB:
		db.nextIDXhtml_col_typeDB++
		v.ID = db.nextIDXhtml_col_typeDB
		db.xhtml_col_typeDBs[v.ID] = v
	case *Xhtml_colgroup_typeDB:
		db.nextIDXhtml_colgroup_typeDB++
		v.ID = db.nextIDXhtml_colgroup_typeDB
		db.xhtml_colgroup_typeDBs[v.ID] = v
	case *Xhtml_dd_typeDB:
		db.nextIDXhtml_dd_typeDB++
		v.ID = db.nextIDXhtml_dd_typeDB
		db.xhtml_dd_typeDBs[v.ID] = v
	case *Xhtml_dfn_typeDB:
		db.nextIDXhtml_dfn_typeDB++
		v.ID = db.nextIDXhtml_dfn_typeDB
		db.xhtml_dfn_typeDBs[v.ID] = v
	case *Xhtml_div_typeDB:
		db.nextIDXhtml_div_typeDB++
		v.ID = db.nextIDXhtml_div_typeDB
		db.xhtml_div_typeDBs[v.ID] = v
	case *Xhtml_dl_typeDB:
		db.nextIDXhtml_dl_typeDB++
		v.ID = db.nextIDXhtml_dl_typeDB
		db.xhtml_dl_typeDBs[v.ID] = v
	case *Xhtml_dt_typeDB:
		db.nextIDXhtml_dt_typeDB++
		v.ID = db.nextIDXhtml_dt_typeDB
		db.xhtml_dt_typeDBs[v.ID] = v
	case *Xhtml_edit_typeDB:
		db.nextIDXhtml_edit_typeDB++
		v.ID = db.nextIDXhtml_edit_typeDB
		db.xhtml_edit_typeDBs[v.ID] = v
	case *Xhtml_em_typeDB:
		db.nextIDXhtml_em_typeDB++
		v.ID = db.nextIDXhtml_em_typeDB
		db.xhtml_em_typeDBs[v.ID] = v
	case *Xhtml_h1_typeDB:
		db.nextIDXhtml_h1_typeDB++
		v.ID = db.nextIDXhtml_h1_typeDB
		db.xhtml_h1_typeDBs[v.ID] = v
	case *Xhtml_h2_typeDB:
		db.nextIDXhtml_h2_typeDB++
		v.ID = db.nextIDXhtml_h2_typeDB
		db.xhtml_h2_typeDBs[v.ID] = v
	case *Xhtml_h3_typeDB:
		db.nextIDXhtml_h3_typeDB++
		v.ID = db.nextIDXhtml_h3_typeDB
		db.xhtml_h3_typeDBs[v.ID] = v
	case *Xhtml_h4_typeDB:
		db.nextIDXhtml_h4_typeDB++
		v.ID = db.nextIDXhtml_h4_typeDB
		db.xhtml_h4_typeDBs[v.ID] = v
	case *Xhtml_h5_typeDB:
		db.nextIDXhtml_h5_typeDB++
		v.ID = db.nextIDXhtml_h5_typeDB
		db.xhtml_h5_typeDBs[v.ID] = v
	case *Xhtml_h6_typeDB:
		db.nextIDXhtml_h6_typeDB++
		v.ID = db.nextIDXhtml_h6_typeDB
		db.xhtml_h6_typeDBs[v.ID] = v
	case *Xhtml_heading_typeDB:
		db.nextIDXhtml_heading_typeDB++
		v.ID = db.nextIDXhtml_heading_typeDB
		db.xhtml_heading_typeDBs[v.ID] = v
	case *Xhtml_hr_typeDB:
		db.nextIDXhtml_hr_typeDB++
		v.ID = db.nextIDXhtml_hr_typeDB
		db.xhtml_hr_typeDBs[v.ID] = v
	case *Xhtml_kbd_typeDB:
		db.nextIDXhtml_kbd_typeDB++
		v.ID = db.nextIDXhtml_kbd_typeDB
		db.xhtml_kbd_typeDBs[v.ID] = v
	case *Xhtml_li_typeDB:
		db.nextIDXhtml_li_typeDB++
		v.ID = db.nextIDXhtml_li_typeDB
		db.xhtml_li_typeDBs[v.ID] = v
	case *Xhtml_object_typeDB:
		db.nextIDXhtml_object_typeDB++
		v.ID = db.nextIDXhtml_object_typeDB
		db.xhtml_object_typeDBs[v.ID] = v
	case *Xhtml_ol_typeDB:
		db.nextIDXhtml_ol_typeDB++
		v.ID = db.nextIDXhtml_ol_typeDB
		db.xhtml_ol_typeDBs[v.ID] = v
	case *Xhtml_p_typeDB:
		db.nextIDXhtml_p_typeDB++
		v.ID = db.nextIDXhtml_p_typeDB
		db.xhtml_p_typeDBs[v.ID] = v
	case *Xhtml_param_typeDB:
		db.nextIDXhtml_param_typeDB++
		v.ID = db.nextIDXhtml_param_typeDB
		db.xhtml_param_typeDBs[v.ID] = v
	case *Xhtml_pre_typeDB:
		db.nextIDXhtml_pre_typeDB++
		v.ID = db.nextIDXhtml_pre_typeDB
		db.xhtml_pre_typeDBs[v.ID] = v
	case *Xhtml_q_typeDB:
		db.nextIDXhtml_q_typeDB++
		v.ID = db.nextIDXhtml_q_typeDB
		db.xhtml_q_typeDBs[v.ID] = v
	case *Xhtml_samp_typeDB:
		db.nextIDXhtml_samp_typeDB++
		v.ID = db.nextIDXhtml_samp_typeDB
		db.xhtml_samp_typeDBs[v.ID] = v
	case *Xhtml_span_typeDB:
		db.nextIDXhtml_span_typeDB++
		v.ID = db.nextIDXhtml_span_typeDB
		db.xhtml_span_typeDBs[v.ID] = v
	case *Xhtml_strong_typeDB:
		db.nextIDXhtml_strong_typeDB++
		v.ID = db.nextIDXhtml_strong_typeDB
		db.xhtml_strong_typeDBs[v.ID] = v
	case *Xhtml_table_typeDB:
		db.nextIDXhtml_table_typeDB++
		v.ID = db.nextIDXhtml_table_typeDB
		db.xhtml_table_typeDBs[v.ID] = v
	case *Xhtml_tbody_typeDB:
		db.nextIDXhtml_tbody_typeDB++
		v.ID = db.nextIDXhtml_tbody_typeDB
		db.xhtml_tbody_typeDBs[v.ID] = v
	case *Xhtml_td_typeDB:
		db.nextIDXhtml_td_typeDB++
		v.ID = db.nextIDXhtml_td_typeDB
		db.xhtml_td_typeDBs[v.ID] = v
	case *Xhtml_tfoot_typeDB:
		db.nextIDXhtml_tfoot_typeDB++
		v.ID = db.nextIDXhtml_tfoot_typeDB
		db.xhtml_tfoot_typeDBs[v.ID] = v
	case *Xhtml_th_typeDB:
		db.nextIDXhtml_th_typeDB++
		v.ID = db.nextIDXhtml_th_typeDB
		db.xhtml_th_typeDBs[v.ID] = v
	case *Xhtml_thead_typeDB:
		db.nextIDXhtml_thead_typeDB++
		v.ID = db.nextIDXhtml_thead_typeDB
		db.xhtml_thead_typeDBs[v.ID] = v
	case *Xhtml_tr_typeDB:
		db.nextIDXhtml_tr_typeDB++
		v.ID = db.nextIDXhtml_tr_typeDB
		db.xhtml_tr_typeDBs[v.ID] = v
	case *Xhtml_ul_typeDB:
		db.nextIDXhtml_ul_typeDB++
		v.ID = db.nextIDXhtml_ul_typeDB
		db.xhtml_ul_typeDBs[v.ID] = v
	case *Xhtml_var_typeDB:
		db.nextIDXhtml_var_typeDB++
		v.ID = db.nextIDXhtml_var_typeDB
		db.xhtml_var_typeDBs[v.ID] = v
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, unsupported type in Create")
	}
	return db, nil
}

// Unscoped sets the unscoped flag for soft-deletes (not used in this implementation)
func (db *DBLite) Unscoped() (db.DBInterface, error) {
	return db, nil
}

// Model is a placeholder in this implementation
func (db *DBLite) Model(instanceDB any) (db.DBInterface, error) {
	// Not implemented as types are handled directly
	return db, nil
}

// Delete removes a record from the database
func (db *DBLite) Delete(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ALTERNATIVE_IDDB:
		delete(db.alternative_idDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		delete(db.attribute_definition_booleanDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_DATEDB:
		delete(db.attribute_definition_dateDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		delete(db.attribute_definition_enumerationDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		delete(db.attribute_definition_integerDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_REALDB:
		delete(db.attribute_definition_realDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		delete(db.attribute_definition_stringDBs, v.ID)
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		delete(db.attribute_definition_xhtmlDBs, v.ID)
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		delete(db.attribute_value_booleanDBs, v.ID)
	case *ATTRIBUTE_VALUE_DATEDB:
		delete(db.attribute_value_dateDBs, v.ID)
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		delete(db.attribute_value_enumerationDBs, v.ID)
	case *ATTRIBUTE_VALUE_INTEGERDB:
		delete(db.attribute_value_integerDBs, v.ID)
	case *ATTRIBUTE_VALUE_REALDB:
		delete(db.attribute_value_realDBs, v.ID)
	case *ATTRIBUTE_VALUE_STRINGDB:
		delete(db.attribute_value_stringDBs, v.ID)
	case *ATTRIBUTE_VALUE_XHTMLDB:
		delete(db.attribute_value_xhtmlDBs, v.ID)
	case *AnyTypeDB:
		delete(db.anytypeDBs, v.ID)
	case *DATATYPE_DEFINITION_BOOLEANDB:
		delete(db.datatype_definition_booleanDBs, v.ID)
	case *DATATYPE_DEFINITION_DATEDB:
		delete(db.datatype_definition_dateDBs, v.ID)
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		delete(db.datatype_definition_enumerationDBs, v.ID)
	case *DATATYPE_DEFINITION_INTEGERDB:
		delete(db.datatype_definition_integerDBs, v.ID)
	case *DATATYPE_DEFINITION_REALDB:
		delete(db.datatype_definition_realDBs, v.ID)
	case *DATATYPE_DEFINITION_STRINGDB:
		delete(db.datatype_definition_stringDBs, v.ID)
	case *DATATYPE_DEFINITION_XHTMLDB:
		delete(db.datatype_definition_xhtmlDBs, v.ID)
	case *EMBEDDED_VALUEDB:
		delete(db.embedded_valueDBs, v.ID)
	case *ENUM_VALUEDB:
		delete(db.enum_valueDBs, v.ID)
	case *RELATION_GROUPDB:
		delete(db.relation_groupDBs, v.ID)
	case *RELATION_GROUP_TYPEDB:
		delete(db.relation_group_typeDBs, v.ID)
	case *REQ_IFDB:
		delete(db.req_ifDBs, v.ID)
	case *REQ_IF_CONTENTDB:
		delete(db.req_if_contentDBs, v.ID)
	case *REQ_IF_HEADERDB:
		delete(db.req_if_headerDBs, v.ID)
	case *REQ_IF_TOOL_EXTENSIONDB:
		delete(db.req_if_tool_extensionDBs, v.ID)
	case *SPECIFICATIONDB:
		delete(db.specificationDBs, v.ID)
	case *SPECIFICATION_TYPEDB:
		delete(db.specification_typeDBs, v.ID)
	case *SPEC_HIERARCHYDB:
		delete(db.spec_hierarchyDBs, v.ID)
	case *SPEC_OBJECTDB:
		delete(db.spec_objectDBs, v.ID)
	case *SPEC_OBJECT_TYPEDB:
		delete(db.spec_object_typeDBs, v.ID)
	case *SPEC_RELATIONDB:
		delete(db.spec_relationDBs, v.ID)
	case *SPEC_RELATION_TYPEDB:
		delete(db.spec_relation_typeDBs, v.ID)
	case *XHTML_CONTENTDB:
		delete(db.xhtml_contentDBs, v.ID)
	case *Xhtml_InlPres_typeDB:
		delete(db.xhtml_inlpres_typeDBs, v.ID)
	case *Xhtml_a_typeDB:
		delete(db.xhtml_a_typeDBs, v.ID)
	case *Xhtml_abbr_typeDB:
		delete(db.xhtml_abbr_typeDBs, v.ID)
	case *Xhtml_acronym_typeDB:
		delete(db.xhtml_acronym_typeDBs, v.ID)
	case *Xhtml_address_typeDB:
		delete(db.xhtml_address_typeDBs, v.ID)
	case *Xhtml_blockquote_typeDB:
		delete(db.xhtml_blockquote_typeDBs, v.ID)
	case *Xhtml_br_typeDB:
		delete(db.xhtml_br_typeDBs, v.ID)
	case *Xhtml_caption_typeDB:
		delete(db.xhtml_caption_typeDBs, v.ID)
	case *Xhtml_cite_typeDB:
		delete(db.xhtml_cite_typeDBs, v.ID)
	case *Xhtml_code_typeDB:
		delete(db.xhtml_code_typeDBs, v.ID)
	case *Xhtml_col_typeDB:
		delete(db.xhtml_col_typeDBs, v.ID)
	case *Xhtml_colgroup_typeDB:
		delete(db.xhtml_colgroup_typeDBs, v.ID)
	case *Xhtml_dd_typeDB:
		delete(db.xhtml_dd_typeDBs, v.ID)
	case *Xhtml_dfn_typeDB:
		delete(db.xhtml_dfn_typeDBs, v.ID)
	case *Xhtml_div_typeDB:
		delete(db.xhtml_div_typeDBs, v.ID)
	case *Xhtml_dl_typeDB:
		delete(db.xhtml_dl_typeDBs, v.ID)
	case *Xhtml_dt_typeDB:
		delete(db.xhtml_dt_typeDBs, v.ID)
	case *Xhtml_edit_typeDB:
		delete(db.xhtml_edit_typeDBs, v.ID)
	case *Xhtml_em_typeDB:
		delete(db.xhtml_em_typeDBs, v.ID)
	case *Xhtml_h1_typeDB:
		delete(db.xhtml_h1_typeDBs, v.ID)
	case *Xhtml_h2_typeDB:
		delete(db.xhtml_h2_typeDBs, v.ID)
	case *Xhtml_h3_typeDB:
		delete(db.xhtml_h3_typeDBs, v.ID)
	case *Xhtml_h4_typeDB:
		delete(db.xhtml_h4_typeDBs, v.ID)
	case *Xhtml_h5_typeDB:
		delete(db.xhtml_h5_typeDBs, v.ID)
	case *Xhtml_h6_typeDB:
		delete(db.xhtml_h6_typeDBs, v.ID)
	case *Xhtml_heading_typeDB:
		delete(db.xhtml_heading_typeDBs, v.ID)
	case *Xhtml_hr_typeDB:
		delete(db.xhtml_hr_typeDBs, v.ID)
	case *Xhtml_kbd_typeDB:
		delete(db.xhtml_kbd_typeDBs, v.ID)
	case *Xhtml_li_typeDB:
		delete(db.xhtml_li_typeDBs, v.ID)
	case *Xhtml_object_typeDB:
		delete(db.xhtml_object_typeDBs, v.ID)
	case *Xhtml_ol_typeDB:
		delete(db.xhtml_ol_typeDBs, v.ID)
	case *Xhtml_p_typeDB:
		delete(db.xhtml_p_typeDBs, v.ID)
	case *Xhtml_param_typeDB:
		delete(db.xhtml_param_typeDBs, v.ID)
	case *Xhtml_pre_typeDB:
		delete(db.xhtml_pre_typeDBs, v.ID)
	case *Xhtml_q_typeDB:
		delete(db.xhtml_q_typeDBs, v.ID)
	case *Xhtml_samp_typeDB:
		delete(db.xhtml_samp_typeDBs, v.ID)
	case *Xhtml_span_typeDB:
		delete(db.xhtml_span_typeDBs, v.ID)
	case *Xhtml_strong_typeDB:
		delete(db.xhtml_strong_typeDBs, v.ID)
	case *Xhtml_table_typeDB:
		delete(db.xhtml_table_typeDBs, v.ID)
	case *Xhtml_tbody_typeDB:
		delete(db.xhtml_tbody_typeDBs, v.ID)
	case *Xhtml_td_typeDB:
		delete(db.xhtml_td_typeDBs, v.ID)
	case *Xhtml_tfoot_typeDB:
		delete(db.xhtml_tfoot_typeDBs, v.ID)
	case *Xhtml_th_typeDB:
		delete(db.xhtml_th_typeDBs, v.ID)
	case *Xhtml_thead_typeDB:
		delete(db.xhtml_thead_typeDBs, v.ID)
	case *Xhtml_tr_typeDB:
		delete(db.xhtml_tr_typeDBs, v.ID)
	case *Xhtml_ul_typeDB:
		delete(db.xhtml_ul_typeDBs, v.ID)
	case *Xhtml_var_typeDB:
		delete(db.xhtml_var_typeDBs, v.ID)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, unsupported type in Delete")
	}
	return db, nil
}

// Save updates or inserts a record into the database
func (db *DBLite) Save(instanceDB any) (db.DBInterface, error) {

	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ALTERNATIVE_IDDB:
		db.alternative_idDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		db.attribute_definition_booleanDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_DATEDB:
		db.attribute_definition_dateDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		db.attribute_definition_enumerationDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		db.attribute_definition_integerDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_REALDB:
		db.attribute_definition_realDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		db.attribute_definition_stringDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		db.attribute_definition_xhtmlDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		db.attribute_value_booleanDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_DATEDB:
		db.attribute_value_dateDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		db.attribute_value_enumerationDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_INTEGERDB:
		db.attribute_value_integerDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_REALDB:
		db.attribute_value_realDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_STRINGDB:
		db.attribute_value_stringDBs[v.ID] = v
		return db, nil
	case *ATTRIBUTE_VALUE_XHTMLDB:
		db.attribute_value_xhtmlDBs[v.ID] = v
		return db, nil
	case *AnyTypeDB:
		db.anytypeDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_BOOLEANDB:
		db.datatype_definition_booleanDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_DATEDB:
		db.datatype_definition_dateDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		db.datatype_definition_enumerationDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_INTEGERDB:
		db.datatype_definition_integerDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_REALDB:
		db.datatype_definition_realDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_STRINGDB:
		db.datatype_definition_stringDBs[v.ID] = v
		return db, nil
	case *DATATYPE_DEFINITION_XHTMLDB:
		db.datatype_definition_xhtmlDBs[v.ID] = v
		return db, nil
	case *EMBEDDED_VALUEDB:
		db.embedded_valueDBs[v.ID] = v
		return db, nil
	case *ENUM_VALUEDB:
		db.enum_valueDBs[v.ID] = v
		return db, nil
	case *RELATION_GROUPDB:
		db.relation_groupDBs[v.ID] = v
		return db, nil
	case *RELATION_GROUP_TYPEDB:
		db.relation_group_typeDBs[v.ID] = v
		return db, nil
	case *REQ_IFDB:
		db.req_ifDBs[v.ID] = v
		return db, nil
	case *REQ_IF_CONTENTDB:
		db.req_if_contentDBs[v.ID] = v
		return db, nil
	case *REQ_IF_HEADERDB:
		db.req_if_headerDBs[v.ID] = v
		return db, nil
	case *REQ_IF_TOOL_EXTENSIONDB:
		db.req_if_tool_extensionDBs[v.ID] = v
		return db, nil
	case *SPECIFICATIONDB:
		db.specificationDBs[v.ID] = v
		return db, nil
	case *SPECIFICATION_TYPEDB:
		db.specification_typeDBs[v.ID] = v
		return db, nil
	case *SPEC_HIERARCHYDB:
		db.spec_hierarchyDBs[v.ID] = v
		return db, nil
	case *SPEC_OBJECTDB:
		db.spec_objectDBs[v.ID] = v
		return db, nil
	case *SPEC_OBJECT_TYPEDB:
		db.spec_object_typeDBs[v.ID] = v
		return db, nil
	case *SPEC_RELATIONDB:
		db.spec_relationDBs[v.ID] = v
		return db, nil
	case *SPEC_RELATION_TYPEDB:
		db.spec_relation_typeDBs[v.ID] = v
		return db, nil
	case *XHTML_CONTENTDB:
		db.xhtml_contentDBs[v.ID] = v
		return db, nil
	case *Xhtml_InlPres_typeDB:
		db.xhtml_inlpres_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_a_typeDB:
		db.xhtml_a_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_abbr_typeDB:
		db.xhtml_abbr_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_acronym_typeDB:
		db.xhtml_acronym_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_address_typeDB:
		db.xhtml_address_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_blockquote_typeDB:
		db.xhtml_blockquote_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_br_typeDB:
		db.xhtml_br_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_caption_typeDB:
		db.xhtml_caption_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_cite_typeDB:
		db.xhtml_cite_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_code_typeDB:
		db.xhtml_code_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_col_typeDB:
		db.xhtml_col_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_colgroup_typeDB:
		db.xhtml_colgroup_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_dd_typeDB:
		db.xhtml_dd_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_dfn_typeDB:
		db.xhtml_dfn_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_div_typeDB:
		db.xhtml_div_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_dl_typeDB:
		db.xhtml_dl_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_dt_typeDB:
		db.xhtml_dt_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_edit_typeDB:
		db.xhtml_edit_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_em_typeDB:
		db.xhtml_em_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_h1_typeDB:
		db.xhtml_h1_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_h2_typeDB:
		db.xhtml_h2_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_h3_typeDB:
		db.xhtml_h3_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_h4_typeDB:
		db.xhtml_h4_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_h5_typeDB:
		db.xhtml_h5_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_h6_typeDB:
		db.xhtml_h6_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_heading_typeDB:
		db.xhtml_heading_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_hr_typeDB:
		db.xhtml_hr_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_kbd_typeDB:
		db.xhtml_kbd_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_li_typeDB:
		db.xhtml_li_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_object_typeDB:
		db.xhtml_object_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_ol_typeDB:
		db.xhtml_ol_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_p_typeDB:
		db.xhtml_p_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_param_typeDB:
		db.xhtml_param_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_pre_typeDB:
		db.xhtml_pre_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_q_typeDB:
		db.xhtml_q_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_samp_typeDB:
		db.xhtml_samp_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_span_typeDB:
		db.xhtml_span_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_strong_typeDB:
		db.xhtml_strong_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_table_typeDB:
		db.xhtml_table_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_tbody_typeDB:
		db.xhtml_tbody_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_td_typeDB:
		db.xhtml_td_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_tfoot_typeDB:
		db.xhtml_tfoot_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_th_typeDB:
		db.xhtml_th_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_thead_typeDB:
		db.xhtml_thead_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_tr_typeDB:
		db.xhtml_tr_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_ul_typeDB:
		db.xhtml_ul_typeDBs[v.ID] = v
		return db, nil
	case *Xhtml_var_typeDB:
		db.xhtml_var_typeDBs[v.ID] = v
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, Save: unsupported type")
	}
}

// Updates modifies an existing record in the database
func (db *DBLite) Updates(instanceDB any) (db.DBInterface, error) {
	if instanceDB == nil {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, instanceDB cannot be nil")
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	switch v := instanceDB.(type) {
	// insertion point delete
	case *ALTERNATIVE_IDDB:
		if existing, ok := db.alternative_idDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ALTERNATIVE_ID github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		if existing, ok := db.attribute_definition_booleanDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_BOOLEAN github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_DATEDB:
		if existing, ok := db.attribute_definition_dateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_DATE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		if existing, ok := db.attribute_definition_enumerationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_ENUMERATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		if existing, ok := db.attribute_definition_integerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_INTEGER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_REALDB:
		if existing, ok := db.attribute_definition_realDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_REAL github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		if existing, ok := db.attribute_definition_stringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_STRING github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		if existing, ok := db.attribute_definition_xhtmlDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_DEFINITION_XHTML github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		if existing, ok := db.attribute_value_booleanDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_BOOLEAN github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_DATEDB:
		if existing, ok := db.attribute_value_dateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_DATE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		if existing, ok := db.attribute_value_enumerationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_ENUMERATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_INTEGERDB:
		if existing, ok := db.attribute_value_integerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_INTEGER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_REALDB:
		if existing, ok := db.attribute_value_realDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_REAL github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_STRINGDB:
		if existing, ok := db.attribute_value_stringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_STRING github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ATTRIBUTE_VALUE_XHTMLDB:
		if existing, ok := db.attribute_value_xhtmlDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ATTRIBUTE_VALUE_XHTML github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *AnyTypeDB:
		if existing, ok := db.anytypeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db AnyType github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_BOOLEANDB:
		if existing, ok := db.datatype_definition_booleanDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_BOOLEAN github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_DATEDB:
		if existing, ok := db.datatype_definition_dateDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_DATE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		if existing, ok := db.datatype_definition_enumerationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_ENUMERATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_INTEGERDB:
		if existing, ok := db.datatype_definition_integerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_INTEGER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_REALDB:
		if existing, ok := db.datatype_definition_realDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_REAL github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_STRINGDB:
		if existing, ok := db.datatype_definition_stringDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_STRING github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *DATATYPE_DEFINITION_XHTMLDB:
		if existing, ok := db.datatype_definition_xhtmlDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db DATATYPE_DEFINITION_XHTML github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *EMBEDDED_VALUEDB:
		if existing, ok := db.embedded_valueDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db EMBEDDED_VALUE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *ENUM_VALUEDB:
		if existing, ok := db.enum_valueDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db ENUM_VALUE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *RELATION_GROUPDB:
		if existing, ok := db.relation_groupDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RELATION_GROUP github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *RELATION_GROUP_TYPEDB:
		if existing, ok := db.relation_group_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db RELATION_GROUP_TYPE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *REQ_IFDB:
		if existing, ok := db.req_ifDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db REQ_IF github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *REQ_IF_CONTENTDB:
		if existing, ok := db.req_if_contentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db REQ_IF_CONTENT github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *REQ_IF_HEADERDB:
		if existing, ok := db.req_if_headerDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db REQ_IF_HEADER github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *REQ_IF_TOOL_EXTENSIONDB:
		if existing, ok := db.req_if_tool_extensionDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db REQ_IF_TOOL_EXTENSION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPECIFICATIONDB:
		if existing, ok := db.specificationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPECIFICATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPECIFICATION_TYPEDB:
		if existing, ok := db.specification_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPECIFICATION_TYPE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_HIERARCHYDB:
		if existing, ok := db.spec_hierarchyDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_HIERARCHY github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_OBJECTDB:
		if existing, ok := db.spec_objectDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_OBJECT github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_OBJECT_TYPEDB:
		if existing, ok := db.spec_object_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_OBJECT_TYPE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_RELATIONDB:
		if existing, ok := db.spec_relationDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_RELATION github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *SPEC_RELATION_TYPEDB:
		if existing, ok := db.spec_relation_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db SPEC_RELATION_TYPE github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *XHTML_CONTENTDB:
		if existing, ok := db.xhtml_contentDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db XHTML_CONTENT github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_InlPres_typeDB:
		if existing, ok := db.xhtml_inlpres_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_InlPres_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_a_typeDB:
		if existing, ok := db.xhtml_a_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_a_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_abbr_typeDB:
		if existing, ok := db.xhtml_abbr_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_abbr_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_acronym_typeDB:
		if existing, ok := db.xhtml_acronym_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_acronym_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_address_typeDB:
		if existing, ok := db.xhtml_address_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_address_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_blockquote_typeDB:
		if existing, ok := db.xhtml_blockquote_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_blockquote_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_br_typeDB:
		if existing, ok := db.xhtml_br_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_br_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_caption_typeDB:
		if existing, ok := db.xhtml_caption_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_caption_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_cite_typeDB:
		if existing, ok := db.xhtml_cite_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_cite_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_code_typeDB:
		if existing, ok := db.xhtml_code_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_code_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_col_typeDB:
		if existing, ok := db.xhtml_col_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_col_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_colgroup_typeDB:
		if existing, ok := db.xhtml_colgroup_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_colgroup_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_dd_typeDB:
		if existing, ok := db.xhtml_dd_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_dd_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_dfn_typeDB:
		if existing, ok := db.xhtml_dfn_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_dfn_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_div_typeDB:
		if existing, ok := db.xhtml_div_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_div_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_dl_typeDB:
		if existing, ok := db.xhtml_dl_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_dl_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_dt_typeDB:
		if existing, ok := db.xhtml_dt_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_dt_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_edit_typeDB:
		if existing, ok := db.xhtml_edit_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_edit_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_em_typeDB:
		if existing, ok := db.xhtml_em_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_em_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_h1_typeDB:
		if existing, ok := db.xhtml_h1_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_h1_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_h2_typeDB:
		if existing, ok := db.xhtml_h2_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_h2_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_h3_typeDB:
		if existing, ok := db.xhtml_h3_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_h3_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_h4_typeDB:
		if existing, ok := db.xhtml_h4_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_h4_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_h5_typeDB:
		if existing, ok := db.xhtml_h5_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_h5_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_h6_typeDB:
		if existing, ok := db.xhtml_h6_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_h6_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_heading_typeDB:
		if existing, ok := db.xhtml_heading_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_heading_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_hr_typeDB:
		if existing, ok := db.xhtml_hr_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_hr_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_kbd_typeDB:
		if existing, ok := db.xhtml_kbd_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_kbd_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_li_typeDB:
		if existing, ok := db.xhtml_li_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_li_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_object_typeDB:
		if existing, ok := db.xhtml_object_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_object_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_ol_typeDB:
		if existing, ok := db.xhtml_ol_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_ol_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_p_typeDB:
		if existing, ok := db.xhtml_p_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_p_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_param_typeDB:
		if existing, ok := db.xhtml_param_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_param_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_pre_typeDB:
		if existing, ok := db.xhtml_pre_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_pre_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_q_typeDB:
		if existing, ok := db.xhtml_q_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_q_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_samp_typeDB:
		if existing, ok := db.xhtml_samp_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_samp_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_span_typeDB:
		if existing, ok := db.xhtml_span_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_span_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_strong_typeDB:
		if existing, ok := db.xhtml_strong_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_strong_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_table_typeDB:
		if existing, ok := db.xhtml_table_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_table_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_tbody_typeDB:
		if existing, ok := db.xhtml_tbody_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_tbody_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_td_typeDB:
		if existing, ok := db.xhtml_td_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_td_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_tfoot_typeDB:
		if existing, ok := db.xhtml_tfoot_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_tfoot_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_th_typeDB:
		if existing, ok := db.xhtml_th_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_th_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_thead_typeDB:
		if existing, ok := db.xhtml_thead_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_thead_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_tr_typeDB:
		if existing, ok := db.xhtml_tr_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_tr_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_ul_typeDB:
		if existing, ok := db.xhtml_ul_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_ul_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	case *Xhtml_var_typeDB:
		if existing, ok := db.xhtml_var_typeDBs[v.ID]; ok {
			*existing = *v
		} else {
			return nil, errors.New("db Xhtml_var_type github.com/fullstack-lang/gongreqif/go, record not found")
		}
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, unsupported type in Updates")
	}
	return db, nil
}

// Find retrieves all records of a type from the database
func (db *DBLite) Find(instanceDBs any) (db.DBInterface, error) {

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch ptr := instanceDBs.(type) {
	// insertion point find
	case *[]ALTERNATIVE_IDDB:
		*ptr = make([]ALTERNATIVE_IDDB, 0, len(db.alternative_idDBs))
		for _, v := range db.alternative_idDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_BOOLEANDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_BOOLEANDB, 0, len(db.attribute_definition_booleanDBs))
		for _, v := range db.attribute_definition_booleanDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_DATEDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_DATEDB, 0, len(db.attribute_definition_dateDBs))
		for _, v := range db.attribute_definition_dateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_ENUMERATIONDB, 0, len(db.attribute_definition_enumerationDBs))
		for _, v := range db.attribute_definition_enumerationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_INTEGERDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_INTEGERDB, 0, len(db.attribute_definition_integerDBs))
		for _, v := range db.attribute_definition_integerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_REALDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_REALDB, 0, len(db.attribute_definition_realDBs))
		for _, v := range db.attribute_definition_realDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_STRINGDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_STRINGDB, 0, len(db.attribute_definition_stringDBs))
		for _, v := range db.attribute_definition_stringDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_DEFINITION_XHTMLDB:
		*ptr = make([]ATTRIBUTE_DEFINITION_XHTMLDB, 0, len(db.attribute_definition_xhtmlDBs))
		for _, v := range db.attribute_definition_xhtmlDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_BOOLEANDB:
		*ptr = make([]ATTRIBUTE_VALUE_BOOLEANDB, 0, len(db.attribute_value_booleanDBs))
		for _, v := range db.attribute_value_booleanDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_DATEDB:
		*ptr = make([]ATTRIBUTE_VALUE_DATEDB, 0, len(db.attribute_value_dateDBs))
		for _, v := range db.attribute_value_dateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_ENUMERATIONDB:
		*ptr = make([]ATTRIBUTE_VALUE_ENUMERATIONDB, 0, len(db.attribute_value_enumerationDBs))
		for _, v := range db.attribute_value_enumerationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_INTEGERDB:
		*ptr = make([]ATTRIBUTE_VALUE_INTEGERDB, 0, len(db.attribute_value_integerDBs))
		for _, v := range db.attribute_value_integerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_REALDB:
		*ptr = make([]ATTRIBUTE_VALUE_REALDB, 0, len(db.attribute_value_realDBs))
		for _, v := range db.attribute_value_realDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_STRINGDB:
		*ptr = make([]ATTRIBUTE_VALUE_STRINGDB, 0, len(db.attribute_value_stringDBs))
		for _, v := range db.attribute_value_stringDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ATTRIBUTE_VALUE_XHTMLDB:
		*ptr = make([]ATTRIBUTE_VALUE_XHTMLDB, 0, len(db.attribute_value_xhtmlDBs))
		for _, v := range db.attribute_value_xhtmlDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]AnyTypeDB:
		*ptr = make([]AnyTypeDB, 0, len(db.anytypeDBs))
		for _, v := range db.anytypeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_BOOLEANDB:
		*ptr = make([]DATATYPE_DEFINITION_BOOLEANDB, 0, len(db.datatype_definition_booleanDBs))
		for _, v := range db.datatype_definition_booleanDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_DATEDB:
		*ptr = make([]DATATYPE_DEFINITION_DATEDB, 0, len(db.datatype_definition_dateDBs))
		for _, v := range db.datatype_definition_dateDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_ENUMERATIONDB:
		*ptr = make([]DATATYPE_DEFINITION_ENUMERATIONDB, 0, len(db.datatype_definition_enumerationDBs))
		for _, v := range db.datatype_definition_enumerationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_INTEGERDB:
		*ptr = make([]DATATYPE_DEFINITION_INTEGERDB, 0, len(db.datatype_definition_integerDBs))
		for _, v := range db.datatype_definition_integerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_REALDB:
		*ptr = make([]DATATYPE_DEFINITION_REALDB, 0, len(db.datatype_definition_realDBs))
		for _, v := range db.datatype_definition_realDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_STRINGDB:
		*ptr = make([]DATATYPE_DEFINITION_STRINGDB, 0, len(db.datatype_definition_stringDBs))
		for _, v := range db.datatype_definition_stringDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]DATATYPE_DEFINITION_XHTMLDB:
		*ptr = make([]DATATYPE_DEFINITION_XHTMLDB, 0, len(db.datatype_definition_xhtmlDBs))
		for _, v := range db.datatype_definition_xhtmlDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]EMBEDDED_VALUEDB:
		*ptr = make([]EMBEDDED_VALUEDB, 0, len(db.embedded_valueDBs))
		for _, v := range db.embedded_valueDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]ENUM_VALUEDB:
		*ptr = make([]ENUM_VALUEDB, 0, len(db.enum_valueDBs))
		for _, v := range db.enum_valueDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RELATION_GROUPDB:
		*ptr = make([]RELATION_GROUPDB, 0, len(db.relation_groupDBs))
		for _, v := range db.relation_groupDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]RELATION_GROUP_TYPEDB:
		*ptr = make([]RELATION_GROUP_TYPEDB, 0, len(db.relation_group_typeDBs))
		for _, v := range db.relation_group_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]REQ_IFDB:
		*ptr = make([]REQ_IFDB, 0, len(db.req_ifDBs))
		for _, v := range db.req_ifDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]REQ_IF_CONTENTDB:
		*ptr = make([]REQ_IF_CONTENTDB, 0, len(db.req_if_contentDBs))
		for _, v := range db.req_if_contentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]REQ_IF_HEADERDB:
		*ptr = make([]REQ_IF_HEADERDB, 0, len(db.req_if_headerDBs))
		for _, v := range db.req_if_headerDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]REQ_IF_TOOL_EXTENSIONDB:
		*ptr = make([]REQ_IF_TOOL_EXTENSIONDB, 0, len(db.req_if_tool_extensionDBs))
		for _, v := range db.req_if_tool_extensionDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPECIFICATIONDB:
		*ptr = make([]SPECIFICATIONDB, 0, len(db.specificationDBs))
		for _, v := range db.specificationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPECIFICATION_TYPEDB:
		*ptr = make([]SPECIFICATION_TYPEDB, 0, len(db.specification_typeDBs))
		for _, v := range db.specification_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_HIERARCHYDB:
		*ptr = make([]SPEC_HIERARCHYDB, 0, len(db.spec_hierarchyDBs))
		for _, v := range db.spec_hierarchyDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_OBJECTDB:
		*ptr = make([]SPEC_OBJECTDB, 0, len(db.spec_objectDBs))
		for _, v := range db.spec_objectDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_OBJECT_TYPEDB:
		*ptr = make([]SPEC_OBJECT_TYPEDB, 0, len(db.spec_object_typeDBs))
		for _, v := range db.spec_object_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_RELATIONDB:
		*ptr = make([]SPEC_RELATIONDB, 0, len(db.spec_relationDBs))
		for _, v := range db.spec_relationDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]SPEC_RELATION_TYPEDB:
		*ptr = make([]SPEC_RELATION_TYPEDB, 0, len(db.spec_relation_typeDBs))
		for _, v := range db.spec_relation_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]XHTML_CONTENTDB:
		*ptr = make([]XHTML_CONTENTDB, 0, len(db.xhtml_contentDBs))
		for _, v := range db.xhtml_contentDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_InlPres_typeDB:
		*ptr = make([]Xhtml_InlPres_typeDB, 0, len(db.xhtml_inlpres_typeDBs))
		for _, v := range db.xhtml_inlpres_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_a_typeDB:
		*ptr = make([]Xhtml_a_typeDB, 0, len(db.xhtml_a_typeDBs))
		for _, v := range db.xhtml_a_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_abbr_typeDB:
		*ptr = make([]Xhtml_abbr_typeDB, 0, len(db.xhtml_abbr_typeDBs))
		for _, v := range db.xhtml_abbr_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_acronym_typeDB:
		*ptr = make([]Xhtml_acronym_typeDB, 0, len(db.xhtml_acronym_typeDBs))
		for _, v := range db.xhtml_acronym_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_address_typeDB:
		*ptr = make([]Xhtml_address_typeDB, 0, len(db.xhtml_address_typeDBs))
		for _, v := range db.xhtml_address_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_blockquote_typeDB:
		*ptr = make([]Xhtml_blockquote_typeDB, 0, len(db.xhtml_blockquote_typeDBs))
		for _, v := range db.xhtml_blockquote_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_br_typeDB:
		*ptr = make([]Xhtml_br_typeDB, 0, len(db.xhtml_br_typeDBs))
		for _, v := range db.xhtml_br_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_caption_typeDB:
		*ptr = make([]Xhtml_caption_typeDB, 0, len(db.xhtml_caption_typeDBs))
		for _, v := range db.xhtml_caption_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_cite_typeDB:
		*ptr = make([]Xhtml_cite_typeDB, 0, len(db.xhtml_cite_typeDBs))
		for _, v := range db.xhtml_cite_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_code_typeDB:
		*ptr = make([]Xhtml_code_typeDB, 0, len(db.xhtml_code_typeDBs))
		for _, v := range db.xhtml_code_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_col_typeDB:
		*ptr = make([]Xhtml_col_typeDB, 0, len(db.xhtml_col_typeDBs))
		for _, v := range db.xhtml_col_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_colgroup_typeDB:
		*ptr = make([]Xhtml_colgroup_typeDB, 0, len(db.xhtml_colgroup_typeDBs))
		for _, v := range db.xhtml_colgroup_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_dd_typeDB:
		*ptr = make([]Xhtml_dd_typeDB, 0, len(db.xhtml_dd_typeDBs))
		for _, v := range db.xhtml_dd_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_dfn_typeDB:
		*ptr = make([]Xhtml_dfn_typeDB, 0, len(db.xhtml_dfn_typeDBs))
		for _, v := range db.xhtml_dfn_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_div_typeDB:
		*ptr = make([]Xhtml_div_typeDB, 0, len(db.xhtml_div_typeDBs))
		for _, v := range db.xhtml_div_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_dl_typeDB:
		*ptr = make([]Xhtml_dl_typeDB, 0, len(db.xhtml_dl_typeDBs))
		for _, v := range db.xhtml_dl_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_dt_typeDB:
		*ptr = make([]Xhtml_dt_typeDB, 0, len(db.xhtml_dt_typeDBs))
		for _, v := range db.xhtml_dt_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_edit_typeDB:
		*ptr = make([]Xhtml_edit_typeDB, 0, len(db.xhtml_edit_typeDBs))
		for _, v := range db.xhtml_edit_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_em_typeDB:
		*ptr = make([]Xhtml_em_typeDB, 0, len(db.xhtml_em_typeDBs))
		for _, v := range db.xhtml_em_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_h1_typeDB:
		*ptr = make([]Xhtml_h1_typeDB, 0, len(db.xhtml_h1_typeDBs))
		for _, v := range db.xhtml_h1_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_h2_typeDB:
		*ptr = make([]Xhtml_h2_typeDB, 0, len(db.xhtml_h2_typeDBs))
		for _, v := range db.xhtml_h2_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_h3_typeDB:
		*ptr = make([]Xhtml_h3_typeDB, 0, len(db.xhtml_h3_typeDBs))
		for _, v := range db.xhtml_h3_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_h4_typeDB:
		*ptr = make([]Xhtml_h4_typeDB, 0, len(db.xhtml_h4_typeDBs))
		for _, v := range db.xhtml_h4_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_h5_typeDB:
		*ptr = make([]Xhtml_h5_typeDB, 0, len(db.xhtml_h5_typeDBs))
		for _, v := range db.xhtml_h5_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_h6_typeDB:
		*ptr = make([]Xhtml_h6_typeDB, 0, len(db.xhtml_h6_typeDBs))
		for _, v := range db.xhtml_h6_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_heading_typeDB:
		*ptr = make([]Xhtml_heading_typeDB, 0, len(db.xhtml_heading_typeDBs))
		for _, v := range db.xhtml_heading_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_hr_typeDB:
		*ptr = make([]Xhtml_hr_typeDB, 0, len(db.xhtml_hr_typeDBs))
		for _, v := range db.xhtml_hr_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_kbd_typeDB:
		*ptr = make([]Xhtml_kbd_typeDB, 0, len(db.xhtml_kbd_typeDBs))
		for _, v := range db.xhtml_kbd_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_li_typeDB:
		*ptr = make([]Xhtml_li_typeDB, 0, len(db.xhtml_li_typeDBs))
		for _, v := range db.xhtml_li_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_object_typeDB:
		*ptr = make([]Xhtml_object_typeDB, 0, len(db.xhtml_object_typeDBs))
		for _, v := range db.xhtml_object_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_ol_typeDB:
		*ptr = make([]Xhtml_ol_typeDB, 0, len(db.xhtml_ol_typeDBs))
		for _, v := range db.xhtml_ol_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_p_typeDB:
		*ptr = make([]Xhtml_p_typeDB, 0, len(db.xhtml_p_typeDBs))
		for _, v := range db.xhtml_p_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_param_typeDB:
		*ptr = make([]Xhtml_param_typeDB, 0, len(db.xhtml_param_typeDBs))
		for _, v := range db.xhtml_param_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_pre_typeDB:
		*ptr = make([]Xhtml_pre_typeDB, 0, len(db.xhtml_pre_typeDBs))
		for _, v := range db.xhtml_pre_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_q_typeDB:
		*ptr = make([]Xhtml_q_typeDB, 0, len(db.xhtml_q_typeDBs))
		for _, v := range db.xhtml_q_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_samp_typeDB:
		*ptr = make([]Xhtml_samp_typeDB, 0, len(db.xhtml_samp_typeDBs))
		for _, v := range db.xhtml_samp_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_span_typeDB:
		*ptr = make([]Xhtml_span_typeDB, 0, len(db.xhtml_span_typeDBs))
		for _, v := range db.xhtml_span_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_strong_typeDB:
		*ptr = make([]Xhtml_strong_typeDB, 0, len(db.xhtml_strong_typeDBs))
		for _, v := range db.xhtml_strong_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_table_typeDB:
		*ptr = make([]Xhtml_table_typeDB, 0, len(db.xhtml_table_typeDBs))
		for _, v := range db.xhtml_table_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_tbody_typeDB:
		*ptr = make([]Xhtml_tbody_typeDB, 0, len(db.xhtml_tbody_typeDBs))
		for _, v := range db.xhtml_tbody_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_td_typeDB:
		*ptr = make([]Xhtml_td_typeDB, 0, len(db.xhtml_td_typeDBs))
		for _, v := range db.xhtml_td_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_tfoot_typeDB:
		*ptr = make([]Xhtml_tfoot_typeDB, 0, len(db.xhtml_tfoot_typeDBs))
		for _, v := range db.xhtml_tfoot_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_th_typeDB:
		*ptr = make([]Xhtml_th_typeDB, 0, len(db.xhtml_th_typeDBs))
		for _, v := range db.xhtml_th_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_thead_typeDB:
		*ptr = make([]Xhtml_thead_typeDB, 0, len(db.xhtml_thead_typeDBs))
		for _, v := range db.xhtml_thead_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_tr_typeDB:
		*ptr = make([]Xhtml_tr_typeDB, 0, len(db.xhtml_tr_typeDBs))
		for _, v := range db.xhtml_tr_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_ul_typeDB:
		*ptr = make([]Xhtml_ul_typeDB, 0, len(db.xhtml_ul_typeDBs))
		for _, v := range db.xhtml_ul_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	case *[]Xhtml_var_typeDB:
		*ptr = make([]Xhtml_var_typeDB, 0, len(db.xhtml_var_typeDBs))
		for _, v := range db.xhtml_var_typeDBs {
			*ptr = append(*ptr, *v)
		}
		return db, nil
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, Find: unsupported type")
	}
}

// First retrieves the first record of a type from the database
func (db *DBLite) First(instanceDB any, conds ...any) (db.DBInterface, error) {
	if len(conds) != 1 {
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, Do not process when conds is not a single parameter")
	}

	var i uint64
	var err error

	switch cond := conds[0].(type) {
	case string:
		i, err = strconv.ParseUint(cond, 10, 32) // Base 10, 32-bit unsigned int
		if err != nil {
			return nil, errors.New("github.com/fullstack-lang/gongreqif/go, conds[0] is not a string number")
		}
	case uint64:
		i = cond
	case uint:
		i = uint64(cond)
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, conds[0] is not a string or uint64")
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	switch instanceDB.(type) {
	// insertion point first
	case *ALTERNATIVE_IDDB:
		tmp, ok := db.alternative_idDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ALTERNATIVE_ID Unkown entry %d", i))
		}

		alternative_idDB, _ := instanceDB.(*ALTERNATIVE_IDDB)
		*alternative_idDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_BOOLEANDB:
		tmp, ok := db.attribute_definition_booleanDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_BOOLEAN Unkown entry %d", i))
		}

		attribute_definition_booleanDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_BOOLEANDB)
		*attribute_definition_booleanDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_DATEDB:
		tmp, ok := db.attribute_definition_dateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_DATE Unkown entry %d", i))
		}

		attribute_definition_dateDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_DATEDB)
		*attribute_definition_dateDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_ENUMERATIONDB:
		tmp, ok := db.attribute_definition_enumerationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_ENUMERATION Unkown entry %d", i))
		}

		attribute_definition_enumerationDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_ENUMERATIONDB)
		*attribute_definition_enumerationDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_INTEGERDB:
		tmp, ok := db.attribute_definition_integerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_INTEGER Unkown entry %d", i))
		}

		attribute_definition_integerDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_INTEGERDB)
		*attribute_definition_integerDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_REALDB:
		tmp, ok := db.attribute_definition_realDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_REAL Unkown entry %d", i))
		}

		attribute_definition_realDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_REALDB)
		*attribute_definition_realDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_STRINGDB:
		tmp, ok := db.attribute_definition_stringDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_STRING Unkown entry %d", i))
		}

		attribute_definition_stringDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_STRINGDB)
		*attribute_definition_stringDB = *tmp
		
	case *ATTRIBUTE_DEFINITION_XHTMLDB:
		tmp, ok := db.attribute_definition_xhtmlDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_DEFINITION_XHTML Unkown entry %d", i))
		}

		attribute_definition_xhtmlDB, _ := instanceDB.(*ATTRIBUTE_DEFINITION_XHTMLDB)
		*attribute_definition_xhtmlDB = *tmp
		
	case *ATTRIBUTE_VALUE_BOOLEANDB:
		tmp, ok := db.attribute_value_booleanDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_BOOLEAN Unkown entry %d", i))
		}

		attribute_value_booleanDB, _ := instanceDB.(*ATTRIBUTE_VALUE_BOOLEANDB)
		*attribute_value_booleanDB = *tmp
		
	case *ATTRIBUTE_VALUE_DATEDB:
		tmp, ok := db.attribute_value_dateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_DATE Unkown entry %d", i))
		}

		attribute_value_dateDB, _ := instanceDB.(*ATTRIBUTE_VALUE_DATEDB)
		*attribute_value_dateDB = *tmp
		
	case *ATTRIBUTE_VALUE_ENUMERATIONDB:
		tmp, ok := db.attribute_value_enumerationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_ENUMERATION Unkown entry %d", i))
		}

		attribute_value_enumerationDB, _ := instanceDB.(*ATTRIBUTE_VALUE_ENUMERATIONDB)
		*attribute_value_enumerationDB = *tmp
		
	case *ATTRIBUTE_VALUE_INTEGERDB:
		tmp, ok := db.attribute_value_integerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_INTEGER Unkown entry %d", i))
		}

		attribute_value_integerDB, _ := instanceDB.(*ATTRIBUTE_VALUE_INTEGERDB)
		*attribute_value_integerDB = *tmp
		
	case *ATTRIBUTE_VALUE_REALDB:
		tmp, ok := db.attribute_value_realDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_REAL Unkown entry %d", i))
		}

		attribute_value_realDB, _ := instanceDB.(*ATTRIBUTE_VALUE_REALDB)
		*attribute_value_realDB = *tmp
		
	case *ATTRIBUTE_VALUE_STRINGDB:
		tmp, ok := db.attribute_value_stringDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_STRING Unkown entry %d", i))
		}

		attribute_value_stringDB, _ := instanceDB.(*ATTRIBUTE_VALUE_STRINGDB)
		*attribute_value_stringDB = *tmp
		
	case *ATTRIBUTE_VALUE_XHTMLDB:
		tmp, ok := db.attribute_value_xhtmlDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ATTRIBUTE_VALUE_XHTML Unkown entry %d", i))
		}

		attribute_value_xhtmlDB, _ := instanceDB.(*ATTRIBUTE_VALUE_XHTMLDB)
		*attribute_value_xhtmlDB = *tmp
		
	case *AnyTypeDB:
		tmp, ok := db.anytypeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First AnyType Unkown entry %d", i))
		}

		anytypeDB, _ := instanceDB.(*AnyTypeDB)
		*anytypeDB = *tmp
		
	case *DATATYPE_DEFINITION_BOOLEANDB:
		tmp, ok := db.datatype_definition_booleanDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_BOOLEAN Unkown entry %d", i))
		}

		datatype_definition_booleanDB, _ := instanceDB.(*DATATYPE_DEFINITION_BOOLEANDB)
		*datatype_definition_booleanDB = *tmp
		
	case *DATATYPE_DEFINITION_DATEDB:
		tmp, ok := db.datatype_definition_dateDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_DATE Unkown entry %d", i))
		}

		datatype_definition_dateDB, _ := instanceDB.(*DATATYPE_DEFINITION_DATEDB)
		*datatype_definition_dateDB = *tmp
		
	case *DATATYPE_DEFINITION_ENUMERATIONDB:
		tmp, ok := db.datatype_definition_enumerationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_ENUMERATION Unkown entry %d", i))
		}

		datatype_definition_enumerationDB, _ := instanceDB.(*DATATYPE_DEFINITION_ENUMERATIONDB)
		*datatype_definition_enumerationDB = *tmp
		
	case *DATATYPE_DEFINITION_INTEGERDB:
		tmp, ok := db.datatype_definition_integerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_INTEGER Unkown entry %d", i))
		}

		datatype_definition_integerDB, _ := instanceDB.(*DATATYPE_DEFINITION_INTEGERDB)
		*datatype_definition_integerDB = *tmp
		
	case *DATATYPE_DEFINITION_REALDB:
		tmp, ok := db.datatype_definition_realDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_REAL Unkown entry %d", i))
		}

		datatype_definition_realDB, _ := instanceDB.(*DATATYPE_DEFINITION_REALDB)
		*datatype_definition_realDB = *tmp
		
	case *DATATYPE_DEFINITION_STRINGDB:
		tmp, ok := db.datatype_definition_stringDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_STRING Unkown entry %d", i))
		}

		datatype_definition_stringDB, _ := instanceDB.(*DATATYPE_DEFINITION_STRINGDB)
		*datatype_definition_stringDB = *tmp
		
	case *DATATYPE_DEFINITION_XHTMLDB:
		tmp, ok := db.datatype_definition_xhtmlDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First DATATYPE_DEFINITION_XHTML Unkown entry %d", i))
		}

		datatype_definition_xhtmlDB, _ := instanceDB.(*DATATYPE_DEFINITION_XHTMLDB)
		*datatype_definition_xhtmlDB = *tmp
		
	case *EMBEDDED_VALUEDB:
		tmp, ok := db.embedded_valueDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First EMBEDDED_VALUE Unkown entry %d", i))
		}

		embedded_valueDB, _ := instanceDB.(*EMBEDDED_VALUEDB)
		*embedded_valueDB = *tmp
		
	case *ENUM_VALUEDB:
		tmp, ok := db.enum_valueDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First ENUM_VALUE Unkown entry %d", i))
		}

		enum_valueDB, _ := instanceDB.(*ENUM_VALUEDB)
		*enum_valueDB = *tmp
		
	case *RELATION_GROUPDB:
		tmp, ok := db.relation_groupDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RELATION_GROUP Unkown entry %d", i))
		}

		relation_groupDB, _ := instanceDB.(*RELATION_GROUPDB)
		*relation_groupDB = *tmp
		
	case *RELATION_GROUP_TYPEDB:
		tmp, ok := db.relation_group_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First RELATION_GROUP_TYPE Unkown entry %d", i))
		}

		relation_group_typeDB, _ := instanceDB.(*RELATION_GROUP_TYPEDB)
		*relation_group_typeDB = *tmp
		
	case *REQ_IFDB:
		tmp, ok := db.req_ifDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First REQ_IF Unkown entry %d", i))
		}

		req_ifDB, _ := instanceDB.(*REQ_IFDB)
		*req_ifDB = *tmp
		
	case *REQ_IF_CONTENTDB:
		tmp, ok := db.req_if_contentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First REQ_IF_CONTENT Unkown entry %d", i))
		}

		req_if_contentDB, _ := instanceDB.(*REQ_IF_CONTENTDB)
		*req_if_contentDB = *tmp
		
	case *REQ_IF_HEADERDB:
		tmp, ok := db.req_if_headerDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First REQ_IF_HEADER Unkown entry %d", i))
		}

		req_if_headerDB, _ := instanceDB.(*REQ_IF_HEADERDB)
		*req_if_headerDB = *tmp
		
	case *REQ_IF_TOOL_EXTENSIONDB:
		tmp, ok := db.req_if_tool_extensionDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First REQ_IF_TOOL_EXTENSION Unkown entry %d", i))
		}

		req_if_tool_extensionDB, _ := instanceDB.(*REQ_IF_TOOL_EXTENSIONDB)
		*req_if_tool_extensionDB = *tmp
		
	case *SPECIFICATIONDB:
		tmp, ok := db.specificationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPECIFICATION Unkown entry %d", i))
		}

		specificationDB, _ := instanceDB.(*SPECIFICATIONDB)
		*specificationDB = *tmp
		
	case *SPECIFICATION_TYPEDB:
		tmp, ok := db.specification_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPECIFICATION_TYPE Unkown entry %d", i))
		}

		specification_typeDB, _ := instanceDB.(*SPECIFICATION_TYPEDB)
		*specification_typeDB = *tmp
		
	case *SPEC_HIERARCHYDB:
		tmp, ok := db.spec_hierarchyDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_HIERARCHY Unkown entry %d", i))
		}

		spec_hierarchyDB, _ := instanceDB.(*SPEC_HIERARCHYDB)
		*spec_hierarchyDB = *tmp
		
	case *SPEC_OBJECTDB:
		tmp, ok := db.spec_objectDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_OBJECT Unkown entry %d", i))
		}

		spec_objectDB, _ := instanceDB.(*SPEC_OBJECTDB)
		*spec_objectDB = *tmp
		
	case *SPEC_OBJECT_TYPEDB:
		tmp, ok := db.spec_object_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_OBJECT_TYPE Unkown entry %d", i))
		}

		spec_object_typeDB, _ := instanceDB.(*SPEC_OBJECT_TYPEDB)
		*spec_object_typeDB = *tmp
		
	case *SPEC_RELATIONDB:
		tmp, ok := db.spec_relationDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_RELATION Unkown entry %d", i))
		}

		spec_relationDB, _ := instanceDB.(*SPEC_RELATIONDB)
		*spec_relationDB = *tmp
		
	case *SPEC_RELATION_TYPEDB:
		tmp, ok := db.spec_relation_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First SPEC_RELATION_TYPE Unkown entry %d", i))
		}

		spec_relation_typeDB, _ := instanceDB.(*SPEC_RELATION_TYPEDB)
		*spec_relation_typeDB = *tmp
		
	case *XHTML_CONTENTDB:
		tmp, ok := db.xhtml_contentDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First XHTML_CONTENT Unkown entry %d", i))
		}

		xhtml_contentDB, _ := instanceDB.(*XHTML_CONTENTDB)
		*xhtml_contentDB = *tmp
		
	case *Xhtml_InlPres_typeDB:
		tmp, ok := db.xhtml_inlpres_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_InlPres_type Unkown entry %d", i))
		}

		xhtml_inlpres_typeDB, _ := instanceDB.(*Xhtml_InlPres_typeDB)
		*xhtml_inlpres_typeDB = *tmp
		
	case *Xhtml_a_typeDB:
		tmp, ok := db.xhtml_a_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_a_type Unkown entry %d", i))
		}

		xhtml_a_typeDB, _ := instanceDB.(*Xhtml_a_typeDB)
		*xhtml_a_typeDB = *tmp
		
	case *Xhtml_abbr_typeDB:
		tmp, ok := db.xhtml_abbr_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_abbr_type Unkown entry %d", i))
		}

		xhtml_abbr_typeDB, _ := instanceDB.(*Xhtml_abbr_typeDB)
		*xhtml_abbr_typeDB = *tmp
		
	case *Xhtml_acronym_typeDB:
		tmp, ok := db.xhtml_acronym_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_acronym_type Unkown entry %d", i))
		}

		xhtml_acronym_typeDB, _ := instanceDB.(*Xhtml_acronym_typeDB)
		*xhtml_acronym_typeDB = *tmp
		
	case *Xhtml_address_typeDB:
		tmp, ok := db.xhtml_address_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_address_type Unkown entry %d", i))
		}

		xhtml_address_typeDB, _ := instanceDB.(*Xhtml_address_typeDB)
		*xhtml_address_typeDB = *tmp
		
	case *Xhtml_blockquote_typeDB:
		tmp, ok := db.xhtml_blockquote_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_blockquote_type Unkown entry %d", i))
		}

		xhtml_blockquote_typeDB, _ := instanceDB.(*Xhtml_blockquote_typeDB)
		*xhtml_blockquote_typeDB = *tmp
		
	case *Xhtml_br_typeDB:
		tmp, ok := db.xhtml_br_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_br_type Unkown entry %d", i))
		}

		xhtml_br_typeDB, _ := instanceDB.(*Xhtml_br_typeDB)
		*xhtml_br_typeDB = *tmp
		
	case *Xhtml_caption_typeDB:
		tmp, ok := db.xhtml_caption_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_caption_type Unkown entry %d", i))
		}

		xhtml_caption_typeDB, _ := instanceDB.(*Xhtml_caption_typeDB)
		*xhtml_caption_typeDB = *tmp
		
	case *Xhtml_cite_typeDB:
		tmp, ok := db.xhtml_cite_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_cite_type Unkown entry %d", i))
		}

		xhtml_cite_typeDB, _ := instanceDB.(*Xhtml_cite_typeDB)
		*xhtml_cite_typeDB = *tmp
		
	case *Xhtml_code_typeDB:
		tmp, ok := db.xhtml_code_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_code_type Unkown entry %d", i))
		}

		xhtml_code_typeDB, _ := instanceDB.(*Xhtml_code_typeDB)
		*xhtml_code_typeDB = *tmp
		
	case *Xhtml_col_typeDB:
		tmp, ok := db.xhtml_col_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_col_type Unkown entry %d", i))
		}

		xhtml_col_typeDB, _ := instanceDB.(*Xhtml_col_typeDB)
		*xhtml_col_typeDB = *tmp
		
	case *Xhtml_colgroup_typeDB:
		tmp, ok := db.xhtml_colgroup_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_colgroup_type Unkown entry %d", i))
		}

		xhtml_colgroup_typeDB, _ := instanceDB.(*Xhtml_colgroup_typeDB)
		*xhtml_colgroup_typeDB = *tmp
		
	case *Xhtml_dd_typeDB:
		tmp, ok := db.xhtml_dd_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_dd_type Unkown entry %d", i))
		}

		xhtml_dd_typeDB, _ := instanceDB.(*Xhtml_dd_typeDB)
		*xhtml_dd_typeDB = *tmp
		
	case *Xhtml_dfn_typeDB:
		tmp, ok := db.xhtml_dfn_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_dfn_type Unkown entry %d", i))
		}

		xhtml_dfn_typeDB, _ := instanceDB.(*Xhtml_dfn_typeDB)
		*xhtml_dfn_typeDB = *tmp
		
	case *Xhtml_div_typeDB:
		tmp, ok := db.xhtml_div_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_div_type Unkown entry %d", i))
		}

		xhtml_div_typeDB, _ := instanceDB.(*Xhtml_div_typeDB)
		*xhtml_div_typeDB = *tmp
		
	case *Xhtml_dl_typeDB:
		tmp, ok := db.xhtml_dl_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_dl_type Unkown entry %d", i))
		}

		xhtml_dl_typeDB, _ := instanceDB.(*Xhtml_dl_typeDB)
		*xhtml_dl_typeDB = *tmp
		
	case *Xhtml_dt_typeDB:
		tmp, ok := db.xhtml_dt_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_dt_type Unkown entry %d", i))
		}

		xhtml_dt_typeDB, _ := instanceDB.(*Xhtml_dt_typeDB)
		*xhtml_dt_typeDB = *tmp
		
	case *Xhtml_edit_typeDB:
		tmp, ok := db.xhtml_edit_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_edit_type Unkown entry %d", i))
		}

		xhtml_edit_typeDB, _ := instanceDB.(*Xhtml_edit_typeDB)
		*xhtml_edit_typeDB = *tmp
		
	case *Xhtml_em_typeDB:
		tmp, ok := db.xhtml_em_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_em_type Unkown entry %d", i))
		}

		xhtml_em_typeDB, _ := instanceDB.(*Xhtml_em_typeDB)
		*xhtml_em_typeDB = *tmp
		
	case *Xhtml_h1_typeDB:
		tmp, ok := db.xhtml_h1_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_h1_type Unkown entry %d", i))
		}

		xhtml_h1_typeDB, _ := instanceDB.(*Xhtml_h1_typeDB)
		*xhtml_h1_typeDB = *tmp
		
	case *Xhtml_h2_typeDB:
		tmp, ok := db.xhtml_h2_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_h2_type Unkown entry %d", i))
		}

		xhtml_h2_typeDB, _ := instanceDB.(*Xhtml_h2_typeDB)
		*xhtml_h2_typeDB = *tmp
		
	case *Xhtml_h3_typeDB:
		tmp, ok := db.xhtml_h3_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_h3_type Unkown entry %d", i))
		}

		xhtml_h3_typeDB, _ := instanceDB.(*Xhtml_h3_typeDB)
		*xhtml_h3_typeDB = *tmp
		
	case *Xhtml_h4_typeDB:
		tmp, ok := db.xhtml_h4_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_h4_type Unkown entry %d", i))
		}

		xhtml_h4_typeDB, _ := instanceDB.(*Xhtml_h4_typeDB)
		*xhtml_h4_typeDB = *tmp
		
	case *Xhtml_h5_typeDB:
		tmp, ok := db.xhtml_h5_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_h5_type Unkown entry %d", i))
		}

		xhtml_h5_typeDB, _ := instanceDB.(*Xhtml_h5_typeDB)
		*xhtml_h5_typeDB = *tmp
		
	case *Xhtml_h6_typeDB:
		tmp, ok := db.xhtml_h6_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_h6_type Unkown entry %d", i))
		}

		xhtml_h6_typeDB, _ := instanceDB.(*Xhtml_h6_typeDB)
		*xhtml_h6_typeDB = *tmp
		
	case *Xhtml_heading_typeDB:
		tmp, ok := db.xhtml_heading_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_heading_type Unkown entry %d", i))
		}

		xhtml_heading_typeDB, _ := instanceDB.(*Xhtml_heading_typeDB)
		*xhtml_heading_typeDB = *tmp
		
	case *Xhtml_hr_typeDB:
		tmp, ok := db.xhtml_hr_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_hr_type Unkown entry %d", i))
		}

		xhtml_hr_typeDB, _ := instanceDB.(*Xhtml_hr_typeDB)
		*xhtml_hr_typeDB = *tmp
		
	case *Xhtml_kbd_typeDB:
		tmp, ok := db.xhtml_kbd_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_kbd_type Unkown entry %d", i))
		}

		xhtml_kbd_typeDB, _ := instanceDB.(*Xhtml_kbd_typeDB)
		*xhtml_kbd_typeDB = *tmp
		
	case *Xhtml_li_typeDB:
		tmp, ok := db.xhtml_li_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_li_type Unkown entry %d", i))
		}

		xhtml_li_typeDB, _ := instanceDB.(*Xhtml_li_typeDB)
		*xhtml_li_typeDB = *tmp
		
	case *Xhtml_object_typeDB:
		tmp, ok := db.xhtml_object_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_object_type Unkown entry %d", i))
		}

		xhtml_object_typeDB, _ := instanceDB.(*Xhtml_object_typeDB)
		*xhtml_object_typeDB = *tmp
		
	case *Xhtml_ol_typeDB:
		tmp, ok := db.xhtml_ol_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_ol_type Unkown entry %d", i))
		}

		xhtml_ol_typeDB, _ := instanceDB.(*Xhtml_ol_typeDB)
		*xhtml_ol_typeDB = *tmp
		
	case *Xhtml_p_typeDB:
		tmp, ok := db.xhtml_p_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_p_type Unkown entry %d", i))
		}

		xhtml_p_typeDB, _ := instanceDB.(*Xhtml_p_typeDB)
		*xhtml_p_typeDB = *tmp
		
	case *Xhtml_param_typeDB:
		tmp, ok := db.xhtml_param_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_param_type Unkown entry %d", i))
		}

		xhtml_param_typeDB, _ := instanceDB.(*Xhtml_param_typeDB)
		*xhtml_param_typeDB = *tmp
		
	case *Xhtml_pre_typeDB:
		tmp, ok := db.xhtml_pre_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_pre_type Unkown entry %d", i))
		}

		xhtml_pre_typeDB, _ := instanceDB.(*Xhtml_pre_typeDB)
		*xhtml_pre_typeDB = *tmp
		
	case *Xhtml_q_typeDB:
		tmp, ok := db.xhtml_q_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_q_type Unkown entry %d", i))
		}

		xhtml_q_typeDB, _ := instanceDB.(*Xhtml_q_typeDB)
		*xhtml_q_typeDB = *tmp
		
	case *Xhtml_samp_typeDB:
		tmp, ok := db.xhtml_samp_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_samp_type Unkown entry %d", i))
		}

		xhtml_samp_typeDB, _ := instanceDB.(*Xhtml_samp_typeDB)
		*xhtml_samp_typeDB = *tmp
		
	case *Xhtml_span_typeDB:
		tmp, ok := db.xhtml_span_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_span_type Unkown entry %d", i))
		}

		xhtml_span_typeDB, _ := instanceDB.(*Xhtml_span_typeDB)
		*xhtml_span_typeDB = *tmp
		
	case *Xhtml_strong_typeDB:
		tmp, ok := db.xhtml_strong_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_strong_type Unkown entry %d", i))
		}

		xhtml_strong_typeDB, _ := instanceDB.(*Xhtml_strong_typeDB)
		*xhtml_strong_typeDB = *tmp
		
	case *Xhtml_table_typeDB:
		tmp, ok := db.xhtml_table_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_table_type Unkown entry %d", i))
		}

		xhtml_table_typeDB, _ := instanceDB.(*Xhtml_table_typeDB)
		*xhtml_table_typeDB = *tmp
		
	case *Xhtml_tbody_typeDB:
		tmp, ok := db.xhtml_tbody_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_tbody_type Unkown entry %d", i))
		}

		xhtml_tbody_typeDB, _ := instanceDB.(*Xhtml_tbody_typeDB)
		*xhtml_tbody_typeDB = *tmp
		
	case *Xhtml_td_typeDB:
		tmp, ok := db.xhtml_td_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_td_type Unkown entry %d", i))
		}

		xhtml_td_typeDB, _ := instanceDB.(*Xhtml_td_typeDB)
		*xhtml_td_typeDB = *tmp
		
	case *Xhtml_tfoot_typeDB:
		tmp, ok := db.xhtml_tfoot_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_tfoot_type Unkown entry %d", i))
		}

		xhtml_tfoot_typeDB, _ := instanceDB.(*Xhtml_tfoot_typeDB)
		*xhtml_tfoot_typeDB = *tmp
		
	case *Xhtml_th_typeDB:
		tmp, ok := db.xhtml_th_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_th_type Unkown entry %d", i))
		}

		xhtml_th_typeDB, _ := instanceDB.(*Xhtml_th_typeDB)
		*xhtml_th_typeDB = *tmp
		
	case *Xhtml_thead_typeDB:
		tmp, ok := db.xhtml_thead_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_thead_type Unkown entry %d", i))
		}

		xhtml_thead_typeDB, _ := instanceDB.(*Xhtml_thead_typeDB)
		*xhtml_thead_typeDB = *tmp
		
	case *Xhtml_tr_typeDB:
		tmp, ok := db.xhtml_tr_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_tr_type Unkown entry %d", i))
		}

		xhtml_tr_typeDB, _ := instanceDB.(*Xhtml_tr_typeDB)
		*xhtml_tr_typeDB = *tmp
		
	case *Xhtml_ul_typeDB:
		tmp, ok := db.xhtml_ul_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_ul_type Unkown entry %d", i))
		}

		xhtml_ul_typeDB, _ := instanceDB.(*Xhtml_ul_typeDB)
		*xhtml_ul_typeDB = *tmp
		
	case *Xhtml_var_typeDB:
		tmp, ok := db.xhtml_var_typeDBs[uint(i)]

		if !ok {
			return nil, errors.New(fmt.Sprintf("db.First Xhtml_var_type Unkown entry %d", i))
		}

		xhtml_var_typeDB, _ := instanceDB.(*Xhtml_var_typeDB)
		*xhtml_var_typeDB = *tmp
		
	default:
		return nil, errors.New("github.com/fullstack-lang/gongreqif/go, Unkown type")
	}
	
	return db, nil
}

