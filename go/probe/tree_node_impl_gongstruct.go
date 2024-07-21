// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongreqif/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "ALTERNATIVE_ID" {
		fillUpTable[models.ALTERNATIVE_ID](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_BOOLEAN" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_DATE" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_ENUMERATION" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_INTEGER" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_REAL" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_STRING" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_XHTML" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_BOOLEAN" {
		fillUpTable[models.ATTRIBUTE_VALUE_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_DATE" {
		fillUpTable[models.ATTRIBUTE_VALUE_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_ENUMERATION" {
		fillUpTable[models.ATTRIBUTE_VALUE_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_INTEGER" {
		fillUpTable[models.ATTRIBUTE_VALUE_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_REAL" {
		fillUpTable[models.ATTRIBUTE_VALUE_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_STRING" {
		fillUpTable[models.ATTRIBUTE_VALUE_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_XHTML" {
		fillUpTable[models.ATTRIBUTE_VALUE_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AnyType" {
		fillUpTable[models.AnyType](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_BOOLEAN" {
		fillUpTable[models.DATATYPE_DEFINITION_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_DATE" {
		fillUpTable[models.DATATYPE_DEFINITION_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_ENUMERATION" {
		fillUpTable[models.DATATYPE_DEFINITION_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_INTEGER" {
		fillUpTable[models.DATATYPE_DEFINITION_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_REAL" {
		fillUpTable[models.DATATYPE_DEFINITION_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_STRING" {
		fillUpTable[models.DATATYPE_DEFINITION_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_XHTML" {
		fillUpTable[models.DATATYPE_DEFINITION_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "EMBEDDED_VALUE" {
		fillUpTable[models.EMBEDDED_VALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ENUM_VALUE" {
		fillUpTable[models.ENUM_VALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATION_GROUP" {
		fillUpTable[models.RELATION_GROUP](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATION_GROUP_TYPE" {
		fillUpTable[models.RELATION_GROUP_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF" {
		fillUpTable[models.REQ_IF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_CONTENT" {
		fillUpTable[models.REQ_IF_CONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_HEADER" {
		fillUpTable[models.REQ_IF_HEADER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_TOOL_EXTENSION" {
		fillUpTable[models.REQ_IF_TOOL_EXTENSION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATION" {
		fillUpTable[models.SPECIFICATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATION_TYPE" {
		fillUpTable[models.SPECIFICATION_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_HIERARCHY" {
		fillUpTable[models.SPEC_HIERARCHY](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_OBJECT" {
		fillUpTable[models.SPEC_OBJECT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_OBJECT_TYPE" {
		fillUpTable[models.SPEC_OBJECT_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_RELATION" {
		fillUpTable[models.SPEC_RELATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_RELATION_TYPE" {
		fillUpTable[models.SPEC_RELATION_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "XHTML_CONTENT" {
		fillUpTable[models.XHTML_CONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_InlPres_type" {
		fillUpTable[models.Xhtml_InlPres_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_a_type" {
		fillUpTable[models.Xhtml_a_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_abbr_type" {
		fillUpTable[models.Xhtml_abbr_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_acronym_type" {
		fillUpTable[models.Xhtml_acronym_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_address_type" {
		fillUpTable[models.Xhtml_address_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_blockquote_type" {
		fillUpTable[models.Xhtml_blockquote_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_br_type" {
		fillUpTable[models.Xhtml_br_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_caption_type" {
		fillUpTable[models.Xhtml_caption_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_cite_type" {
		fillUpTable[models.Xhtml_cite_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_code_type" {
		fillUpTable[models.Xhtml_code_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_col_type" {
		fillUpTable[models.Xhtml_col_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_colgroup_type" {
		fillUpTable[models.Xhtml_colgroup_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_dd_type" {
		fillUpTable[models.Xhtml_dd_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_dfn_type" {
		fillUpTable[models.Xhtml_dfn_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_div_type" {
		fillUpTable[models.Xhtml_div_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_dl_type" {
		fillUpTable[models.Xhtml_dl_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_dt_type" {
		fillUpTable[models.Xhtml_dt_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_edit_type" {
		fillUpTable[models.Xhtml_edit_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_em_type" {
		fillUpTable[models.Xhtml_em_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_h1_type" {
		fillUpTable[models.Xhtml_h1_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_h2_type" {
		fillUpTable[models.Xhtml_h2_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_h3_type" {
		fillUpTable[models.Xhtml_h3_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_h4_type" {
		fillUpTable[models.Xhtml_h4_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_h5_type" {
		fillUpTable[models.Xhtml_h5_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_h6_type" {
		fillUpTable[models.Xhtml_h6_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_heading_type" {
		fillUpTable[models.Xhtml_heading_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_hr_type" {
		fillUpTable[models.Xhtml_hr_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_kbd_type" {
		fillUpTable[models.Xhtml_kbd_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_li_type" {
		fillUpTable[models.Xhtml_li_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_object_type" {
		fillUpTable[models.Xhtml_object_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_ol_type" {
		fillUpTable[models.Xhtml_ol_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_p_type" {
		fillUpTable[models.Xhtml_p_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_param_type" {
		fillUpTable[models.Xhtml_param_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_pre_type" {
		fillUpTable[models.Xhtml_pre_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_q_type" {
		fillUpTable[models.Xhtml_q_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_samp_type" {
		fillUpTable[models.Xhtml_samp_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_span_type" {
		fillUpTable[models.Xhtml_span_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_strong_type" {
		fillUpTable[models.Xhtml_strong_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_table_type" {
		fillUpTable[models.Xhtml_table_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_tbody_type" {
		fillUpTable[models.Xhtml_tbody_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_td_type" {
		fillUpTable[models.Xhtml_td_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_tfoot_type" {
		fillUpTable[models.Xhtml_tfoot_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_th_type" {
		fillUpTable[models.Xhtml_th_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_thead_type" {
		fillUpTable[models.Xhtml_thead_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_tr_type" {
		fillUpTable[models.Xhtml_tr_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_ul_type" {
		fillUpTable[models.Xhtml_ul_type](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Xhtml_var_type" {
		fillUpTable[models.Xhtml_var_type](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
