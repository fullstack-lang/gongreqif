// do not modify, generated file
package fullstack

import (
	"github.com/fullstack-lang/gongreqif/go/controllers"
	"github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/orm"

	"github.com/gin-gonic/gin"

	// this will import the angular front end source code directory (versionned with git) in the vendor directory
	// this path will be included in the "tsconfig.json" front end compilation paths
	// to include this stack front end code
	_ "github.com/fullstack-lang/gongreqif/ng-github.com-fullstack-lang-gongreqif/projects"
)

// NewStackInstance creates a new stack instance from the Stack Model
// and returns the backRepo of the stack instance (you can get the stage from backRepo.GetStage()
//
// - the stackPath is the unique identifier of the stack
// - the optional parameter filenames is for the name of the database filename
// if filenames is omitted, the database is persisted in memory
func NewStackInstance(
	r *gin.Engine,
	stackPath string,
	// filesnames is an optional parameter for the name of the database
	filenames ...string) (
	stage *models.StageStruct,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	if stackPath != "" {
		controllers.GetController().AddBackRepo(backRepo, stackPath)
	}

	controllers.Register(r)

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.ALTERNATIVE_ID](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_DEFINITION_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.ATTRIBUTE_VALUE_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.AnyType](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.DATATYPE_DEFINITION_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.EMBEDDED_VALUE](stage)
	models.SetOrchestratorOnAfterUpdate[models.ENUM_VALUE](stage)
	models.SetOrchestratorOnAfterUpdate[models.RELATION_GROUP](stage)
	models.SetOrchestratorOnAfterUpdate[models.RELATION_GROUP_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF_CONTENT](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF_HEADER](stage)
	models.SetOrchestratorOnAfterUpdate[models.REQ_IF_TOOL_EXTENSION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPECIFICATION_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_HIERARCHY](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_OBJECT_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_RELATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.SPEC_RELATION_TYPE](stage)
	models.SetOrchestratorOnAfterUpdate[models.XHTML_CONTENT](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_InlPres_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_a_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_abbr_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_acronym_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_address_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_blockquote_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_br_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_caption_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_cite_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_code_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_col_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_colgroup_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_dd_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_dfn_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_div_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_dl_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_dt_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_edit_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_em_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_h1_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_h2_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_h3_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_h4_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_h5_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_h6_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_heading_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_hr_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_kbd_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_li_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_object_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_ol_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_p_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_param_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_pre_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_q_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_samp_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_span_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_strong_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_table_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_tbody_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_td_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_tfoot_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_th_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_thead_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_tr_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_ul_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Xhtml_var_type](stage)

	return
}
