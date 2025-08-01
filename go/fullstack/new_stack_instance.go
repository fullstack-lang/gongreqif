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
	// This is a level 1 gong application, no need to import the angular code
	// therefore, the following line that is necessary in level 2 applications, is commented
	// _ "github.com/fullstack-lang/gongreqif/ng-github.com-fullstack-lang-gongreqif"
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
	stage *models.Stage,
	backRepo *orm.BackRepoStruct) {

	stage = models.NewStage(stackPath)

	if len(filenames) == 0 {
		filenames = append(filenames, ":memory:")
	}

	backRepo = orm.NewBackRepo(stage, filenames[0])

	controllers.GetController().AddBackRepo(backRepo, stackPath)

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
	models.SetOrchestratorOnAfterUpdate[models.A_ALTERNATIVE_ID](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_DATE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_REAL_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_STRING_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_DEFINITION_XHTML_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_BOOLEAN](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_DATE](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_ENUMERATION](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_INTEGER](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_REAL](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_STRING](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_XHTML](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ATTRIBUTE_VALUE_XHTML_1](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_CHILDREN](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_CORE_CONTENT](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_BOOLEAN_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_DATE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_ENUMERATION_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_INTEGER_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_REAL_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_STRING_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_DATATYPE_DEFINITION_XHTML_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_EDITABLE_ATTS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_ENUM_VALUE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_OBJECT](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_PROPERTIES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_RELATION_GROUP_TYPE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SOURCE_1](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SOURCE_SPECIFICATION_1](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPECIFICATIONS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPECIFICATION_TYPE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPECIFIED_VALUES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_ATTRIBUTES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_OBJECTS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_OBJECT_TYPE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATIONS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATION_GROUPS](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATION_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_RELATION_TYPE_REF](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_SPEC_TYPES](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_THE_HEADER](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_TOOL_EXTENSIONS](stage)
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
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSite](stage)
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSiteChapter](stage)
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSiteGeneratedImage](stage)
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSiteImage](stage)
	models.SetOrchestratorOnAfterUpdate[models.StaticWebSiteParagraph](stage)
	models.SetOrchestratorOnAfterUpdate[models.XHTML_CONTENT](stage)

	return
}
