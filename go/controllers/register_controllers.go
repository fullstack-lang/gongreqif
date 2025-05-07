// generated code - do not edit
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongreqif/go/orm"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gongreqif/go")
	{ // insertion point for registrations
		v1.GET("/v1/alternative_ids", GetController().GetALTERNATIVE_IDs)
		v1.GET("/v1/alternative_ids/:id", GetController().GetALTERNATIVE_ID)
		v1.POST("/v1/alternative_ids", GetController().PostALTERNATIVE_ID)
		v1.PATCH("/v1/alternative_ids/:id", GetController().UpdateALTERNATIVE_ID)
		v1.PUT("/v1/alternative_ids/:id", GetController().UpdateALTERNATIVE_ID)
		v1.DELETE("/v1/alternative_ids/:id", GetController().DeleteALTERNATIVE_ID)

		v1.GET("/v1/attribute_definition_booleans", GetController().GetATTRIBUTE_DEFINITION_BOOLEANs)
		v1.GET("/v1/attribute_definition_booleans/:id", GetController().GetATTRIBUTE_DEFINITION_BOOLEAN)
		v1.POST("/v1/attribute_definition_booleans", GetController().PostATTRIBUTE_DEFINITION_BOOLEAN)
		v1.PATCH("/v1/attribute_definition_booleans/:id", GetController().UpdateATTRIBUTE_DEFINITION_BOOLEAN)
		v1.PUT("/v1/attribute_definition_booleans/:id", GetController().UpdateATTRIBUTE_DEFINITION_BOOLEAN)
		v1.DELETE("/v1/attribute_definition_booleans/:id", GetController().DeleteATTRIBUTE_DEFINITION_BOOLEAN)

		v1.GET("/v1/attribute_definition_dates", GetController().GetATTRIBUTE_DEFINITION_DATEs)
		v1.GET("/v1/attribute_definition_dates/:id", GetController().GetATTRIBUTE_DEFINITION_DATE)
		v1.POST("/v1/attribute_definition_dates", GetController().PostATTRIBUTE_DEFINITION_DATE)
		v1.PATCH("/v1/attribute_definition_dates/:id", GetController().UpdateATTRIBUTE_DEFINITION_DATE)
		v1.PUT("/v1/attribute_definition_dates/:id", GetController().UpdateATTRIBUTE_DEFINITION_DATE)
		v1.DELETE("/v1/attribute_definition_dates/:id", GetController().DeleteATTRIBUTE_DEFINITION_DATE)

		v1.GET("/v1/attribute_definition_enumerations", GetController().GetATTRIBUTE_DEFINITION_ENUMERATIONs)
		v1.GET("/v1/attribute_definition_enumerations/:id", GetController().GetATTRIBUTE_DEFINITION_ENUMERATION)
		v1.POST("/v1/attribute_definition_enumerations", GetController().PostATTRIBUTE_DEFINITION_ENUMERATION)
		v1.PATCH("/v1/attribute_definition_enumerations/:id", GetController().UpdateATTRIBUTE_DEFINITION_ENUMERATION)
		v1.PUT("/v1/attribute_definition_enumerations/:id", GetController().UpdateATTRIBUTE_DEFINITION_ENUMERATION)
		v1.DELETE("/v1/attribute_definition_enumerations/:id", GetController().DeleteATTRIBUTE_DEFINITION_ENUMERATION)

		v1.GET("/v1/attribute_definition_integers", GetController().GetATTRIBUTE_DEFINITION_INTEGERs)
		v1.GET("/v1/attribute_definition_integers/:id", GetController().GetATTRIBUTE_DEFINITION_INTEGER)
		v1.POST("/v1/attribute_definition_integers", GetController().PostATTRIBUTE_DEFINITION_INTEGER)
		v1.PATCH("/v1/attribute_definition_integers/:id", GetController().UpdateATTRIBUTE_DEFINITION_INTEGER)
		v1.PUT("/v1/attribute_definition_integers/:id", GetController().UpdateATTRIBUTE_DEFINITION_INTEGER)
		v1.DELETE("/v1/attribute_definition_integers/:id", GetController().DeleteATTRIBUTE_DEFINITION_INTEGER)

		v1.GET("/v1/attribute_definition_reals", GetController().GetATTRIBUTE_DEFINITION_REALs)
		v1.GET("/v1/attribute_definition_reals/:id", GetController().GetATTRIBUTE_DEFINITION_REAL)
		v1.POST("/v1/attribute_definition_reals", GetController().PostATTRIBUTE_DEFINITION_REAL)
		v1.PATCH("/v1/attribute_definition_reals/:id", GetController().UpdateATTRIBUTE_DEFINITION_REAL)
		v1.PUT("/v1/attribute_definition_reals/:id", GetController().UpdateATTRIBUTE_DEFINITION_REAL)
		v1.DELETE("/v1/attribute_definition_reals/:id", GetController().DeleteATTRIBUTE_DEFINITION_REAL)

		v1.GET("/v1/attribute_definition_strings", GetController().GetATTRIBUTE_DEFINITION_STRINGs)
		v1.GET("/v1/attribute_definition_strings/:id", GetController().GetATTRIBUTE_DEFINITION_STRING)
		v1.POST("/v1/attribute_definition_strings", GetController().PostATTRIBUTE_DEFINITION_STRING)
		v1.PATCH("/v1/attribute_definition_strings/:id", GetController().UpdateATTRIBUTE_DEFINITION_STRING)
		v1.PUT("/v1/attribute_definition_strings/:id", GetController().UpdateATTRIBUTE_DEFINITION_STRING)
		v1.DELETE("/v1/attribute_definition_strings/:id", GetController().DeleteATTRIBUTE_DEFINITION_STRING)

		v1.GET("/v1/attribute_definition_xhtmls", GetController().GetATTRIBUTE_DEFINITION_XHTMLs)
		v1.GET("/v1/attribute_definition_xhtmls/:id", GetController().GetATTRIBUTE_DEFINITION_XHTML)
		v1.POST("/v1/attribute_definition_xhtmls", GetController().PostATTRIBUTE_DEFINITION_XHTML)
		v1.PATCH("/v1/attribute_definition_xhtmls/:id", GetController().UpdateATTRIBUTE_DEFINITION_XHTML)
		v1.PUT("/v1/attribute_definition_xhtmls/:id", GetController().UpdateATTRIBUTE_DEFINITION_XHTML)
		v1.DELETE("/v1/attribute_definition_xhtmls/:id", GetController().DeleteATTRIBUTE_DEFINITION_XHTML)

		v1.GET("/v1/attribute_value_booleans", GetController().GetATTRIBUTE_VALUE_BOOLEANs)
		v1.GET("/v1/attribute_value_booleans/:id", GetController().GetATTRIBUTE_VALUE_BOOLEAN)
		v1.POST("/v1/attribute_value_booleans", GetController().PostATTRIBUTE_VALUE_BOOLEAN)
		v1.PATCH("/v1/attribute_value_booleans/:id", GetController().UpdateATTRIBUTE_VALUE_BOOLEAN)
		v1.PUT("/v1/attribute_value_booleans/:id", GetController().UpdateATTRIBUTE_VALUE_BOOLEAN)
		v1.DELETE("/v1/attribute_value_booleans/:id", GetController().DeleteATTRIBUTE_VALUE_BOOLEAN)

		v1.GET("/v1/attribute_value_dates", GetController().GetATTRIBUTE_VALUE_DATEs)
		v1.GET("/v1/attribute_value_dates/:id", GetController().GetATTRIBUTE_VALUE_DATE)
		v1.POST("/v1/attribute_value_dates", GetController().PostATTRIBUTE_VALUE_DATE)
		v1.PATCH("/v1/attribute_value_dates/:id", GetController().UpdateATTRIBUTE_VALUE_DATE)
		v1.PUT("/v1/attribute_value_dates/:id", GetController().UpdateATTRIBUTE_VALUE_DATE)
		v1.DELETE("/v1/attribute_value_dates/:id", GetController().DeleteATTRIBUTE_VALUE_DATE)

		v1.GET("/v1/attribute_value_enumerations", GetController().GetATTRIBUTE_VALUE_ENUMERATIONs)
		v1.GET("/v1/attribute_value_enumerations/:id", GetController().GetATTRIBUTE_VALUE_ENUMERATION)
		v1.POST("/v1/attribute_value_enumerations", GetController().PostATTRIBUTE_VALUE_ENUMERATION)
		v1.PATCH("/v1/attribute_value_enumerations/:id", GetController().UpdateATTRIBUTE_VALUE_ENUMERATION)
		v1.PUT("/v1/attribute_value_enumerations/:id", GetController().UpdateATTRIBUTE_VALUE_ENUMERATION)
		v1.DELETE("/v1/attribute_value_enumerations/:id", GetController().DeleteATTRIBUTE_VALUE_ENUMERATION)

		v1.GET("/v1/attribute_value_integers", GetController().GetATTRIBUTE_VALUE_INTEGERs)
		v1.GET("/v1/attribute_value_integers/:id", GetController().GetATTRIBUTE_VALUE_INTEGER)
		v1.POST("/v1/attribute_value_integers", GetController().PostATTRIBUTE_VALUE_INTEGER)
		v1.PATCH("/v1/attribute_value_integers/:id", GetController().UpdateATTRIBUTE_VALUE_INTEGER)
		v1.PUT("/v1/attribute_value_integers/:id", GetController().UpdateATTRIBUTE_VALUE_INTEGER)
		v1.DELETE("/v1/attribute_value_integers/:id", GetController().DeleteATTRIBUTE_VALUE_INTEGER)

		v1.GET("/v1/attribute_value_reals", GetController().GetATTRIBUTE_VALUE_REALs)
		v1.GET("/v1/attribute_value_reals/:id", GetController().GetATTRIBUTE_VALUE_REAL)
		v1.POST("/v1/attribute_value_reals", GetController().PostATTRIBUTE_VALUE_REAL)
		v1.PATCH("/v1/attribute_value_reals/:id", GetController().UpdateATTRIBUTE_VALUE_REAL)
		v1.PUT("/v1/attribute_value_reals/:id", GetController().UpdateATTRIBUTE_VALUE_REAL)
		v1.DELETE("/v1/attribute_value_reals/:id", GetController().DeleteATTRIBUTE_VALUE_REAL)

		v1.GET("/v1/attribute_value_strings", GetController().GetATTRIBUTE_VALUE_STRINGs)
		v1.GET("/v1/attribute_value_strings/:id", GetController().GetATTRIBUTE_VALUE_STRING)
		v1.POST("/v1/attribute_value_strings", GetController().PostATTRIBUTE_VALUE_STRING)
		v1.PATCH("/v1/attribute_value_strings/:id", GetController().UpdateATTRIBUTE_VALUE_STRING)
		v1.PUT("/v1/attribute_value_strings/:id", GetController().UpdateATTRIBUTE_VALUE_STRING)
		v1.DELETE("/v1/attribute_value_strings/:id", GetController().DeleteATTRIBUTE_VALUE_STRING)

		v1.GET("/v1/attribute_value_xhtmls", GetController().GetATTRIBUTE_VALUE_XHTMLs)
		v1.GET("/v1/attribute_value_xhtmls/:id", GetController().GetATTRIBUTE_VALUE_XHTML)
		v1.POST("/v1/attribute_value_xhtmls", GetController().PostATTRIBUTE_VALUE_XHTML)
		v1.PATCH("/v1/attribute_value_xhtmls/:id", GetController().UpdateATTRIBUTE_VALUE_XHTML)
		v1.PUT("/v1/attribute_value_xhtmls/:id", GetController().UpdateATTRIBUTE_VALUE_XHTML)
		v1.DELETE("/v1/attribute_value_xhtmls/:id", GetController().DeleteATTRIBUTE_VALUE_XHTML)

		v1.GET("/v1/anytypes", GetController().GetAnyTypes)
		v1.GET("/v1/anytypes/:id", GetController().GetAnyType)
		v1.POST("/v1/anytypes", GetController().PostAnyType)
		v1.PATCH("/v1/anytypes/:id", GetController().UpdateAnyType)
		v1.PUT("/v1/anytypes/:id", GetController().UpdateAnyType)
		v1.DELETE("/v1/anytypes/:id", GetController().DeleteAnyType)

		v1.GET("/v1/datatype_definition_booleans", GetController().GetDATATYPE_DEFINITION_BOOLEANs)
		v1.GET("/v1/datatype_definition_booleans/:id", GetController().GetDATATYPE_DEFINITION_BOOLEAN)
		v1.POST("/v1/datatype_definition_booleans", GetController().PostDATATYPE_DEFINITION_BOOLEAN)
		v1.PATCH("/v1/datatype_definition_booleans/:id", GetController().UpdateDATATYPE_DEFINITION_BOOLEAN)
		v1.PUT("/v1/datatype_definition_booleans/:id", GetController().UpdateDATATYPE_DEFINITION_BOOLEAN)
		v1.DELETE("/v1/datatype_definition_booleans/:id", GetController().DeleteDATATYPE_DEFINITION_BOOLEAN)

		v1.GET("/v1/datatype_definition_dates", GetController().GetDATATYPE_DEFINITION_DATEs)
		v1.GET("/v1/datatype_definition_dates/:id", GetController().GetDATATYPE_DEFINITION_DATE)
		v1.POST("/v1/datatype_definition_dates", GetController().PostDATATYPE_DEFINITION_DATE)
		v1.PATCH("/v1/datatype_definition_dates/:id", GetController().UpdateDATATYPE_DEFINITION_DATE)
		v1.PUT("/v1/datatype_definition_dates/:id", GetController().UpdateDATATYPE_DEFINITION_DATE)
		v1.DELETE("/v1/datatype_definition_dates/:id", GetController().DeleteDATATYPE_DEFINITION_DATE)

		v1.GET("/v1/datatype_definition_enumerations", GetController().GetDATATYPE_DEFINITION_ENUMERATIONs)
		v1.GET("/v1/datatype_definition_enumerations/:id", GetController().GetDATATYPE_DEFINITION_ENUMERATION)
		v1.POST("/v1/datatype_definition_enumerations", GetController().PostDATATYPE_DEFINITION_ENUMERATION)
		v1.PATCH("/v1/datatype_definition_enumerations/:id", GetController().UpdateDATATYPE_DEFINITION_ENUMERATION)
		v1.PUT("/v1/datatype_definition_enumerations/:id", GetController().UpdateDATATYPE_DEFINITION_ENUMERATION)
		v1.DELETE("/v1/datatype_definition_enumerations/:id", GetController().DeleteDATATYPE_DEFINITION_ENUMERATION)

		v1.GET("/v1/datatype_definition_integers", GetController().GetDATATYPE_DEFINITION_INTEGERs)
		v1.GET("/v1/datatype_definition_integers/:id", GetController().GetDATATYPE_DEFINITION_INTEGER)
		v1.POST("/v1/datatype_definition_integers", GetController().PostDATATYPE_DEFINITION_INTEGER)
		v1.PATCH("/v1/datatype_definition_integers/:id", GetController().UpdateDATATYPE_DEFINITION_INTEGER)
		v1.PUT("/v1/datatype_definition_integers/:id", GetController().UpdateDATATYPE_DEFINITION_INTEGER)
		v1.DELETE("/v1/datatype_definition_integers/:id", GetController().DeleteDATATYPE_DEFINITION_INTEGER)

		v1.GET("/v1/datatype_definition_reals", GetController().GetDATATYPE_DEFINITION_REALs)
		v1.GET("/v1/datatype_definition_reals/:id", GetController().GetDATATYPE_DEFINITION_REAL)
		v1.POST("/v1/datatype_definition_reals", GetController().PostDATATYPE_DEFINITION_REAL)
		v1.PATCH("/v1/datatype_definition_reals/:id", GetController().UpdateDATATYPE_DEFINITION_REAL)
		v1.PUT("/v1/datatype_definition_reals/:id", GetController().UpdateDATATYPE_DEFINITION_REAL)
		v1.DELETE("/v1/datatype_definition_reals/:id", GetController().DeleteDATATYPE_DEFINITION_REAL)

		v1.GET("/v1/datatype_definition_strings", GetController().GetDATATYPE_DEFINITION_STRINGs)
		v1.GET("/v1/datatype_definition_strings/:id", GetController().GetDATATYPE_DEFINITION_STRING)
		v1.POST("/v1/datatype_definition_strings", GetController().PostDATATYPE_DEFINITION_STRING)
		v1.PATCH("/v1/datatype_definition_strings/:id", GetController().UpdateDATATYPE_DEFINITION_STRING)
		v1.PUT("/v1/datatype_definition_strings/:id", GetController().UpdateDATATYPE_DEFINITION_STRING)
		v1.DELETE("/v1/datatype_definition_strings/:id", GetController().DeleteDATATYPE_DEFINITION_STRING)

		v1.GET("/v1/datatype_definition_xhtmls", GetController().GetDATATYPE_DEFINITION_XHTMLs)
		v1.GET("/v1/datatype_definition_xhtmls/:id", GetController().GetDATATYPE_DEFINITION_XHTML)
		v1.POST("/v1/datatype_definition_xhtmls", GetController().PostDATATYPE_DEFINITION_XHTML)
		v1.PATCH("/v1/datatype_definition_xhtmls/:id", GetController().UpdateDATATYPE_DEFINITION_XHTML)
		v1.PUT("/v1/datatype_definition_xhtmls/:id", GetController().UpdateDATATYPE_DEFINITION_XHTML)
		v1.DELETE("/v1/datatype_definition_xhtmls/:id", GetController().DeleteDATATYPE_DEFINITION_XHTML)

		v1.GET("/v1/embedded_values", GetController().GetEMBEDDED_VALUEs)
		v1.GET("/v1/embedded_values/:id", GetController().GetEMBEDDED_VALUE)
		v1.POST("/v1/embedded_values", GetController().PostEMBEDDED_VALUE)
		v1.PATCH("/v1/embedded_values/:id", GetController().UpdateEMBEDDED_VALUE)
		v1.PUT("/v1/embedded_values/:id", GetController().UpdateEMBEDDED_VALUE)
		v1.DELETE("/v1/embedded_values/:id", GetController().DeleteEMBEDDED_VALUE)

		v1.GET("/v1/enum_values", GetController().GetENUM_VALUEs)
		v1.GET("/v1/enum_values/:id", GetController().GetENUM_VALUE)
		v1.POST("/v1/enum_values", GetController().PostENUM_VALUE)
		v1.PATCH("/v1/enum_values/:id", GetController().UpdateENUM_VALUE)
		v1.PUT("/v1/enum_values/:id", GetController().UpdateENUM_VALUE)
		v1.DELETE("/v1/enum_values/:id", GetController().DeleteENUM_VALUE)

		v1.GET("/v1/relation_groups", GetController().GetRELATION_GROUPs)
		v1.GET("/v1/relation_groups/:id", GetController().GetRELATION_GROUP)
		v1.POST("/v1/relation_groups", GetController().PostRELATION_GROUP)
		v1.PATCH("/v1/relation_groups/:id", GetController().UpdateRELATION_GROUP)
		v1.PUT("/v1/relation_groups/:id", GetController().UpdateRELATION_GROUP)
		v1.DELETE("/v1/relation_groups/:id", GetController().DeleteRELATION_GROUP)

		v1.GET("/v1/relation_group_types", GetController().GetRELATION_GROUP_TYPEs)
		v1.GET("/v1/relation_group_types/:id", GetController().GetRELATION_GROUP_TYPE)
		v1.POST("/v1/relation_group_types", GetController().PostRELATION_GROUP_TYPE)
		v1.PATCH("/v1/relation_group_types/:id", GetController().UpdateRELATION_GROUP_TYPE)
		v1.PUT("/v1/relation_group_types/:id", GetController().UpdateRELATION_GROUP_TYPE)
		v1.DELETE("/v1/relation_group_types/:id", GetController().DeleteRELATION_GROUP_TYPE)

		v1.GET("/v1/req_ifs", GetController().GetREQ_IFs)
		v1.GET("/v1/req_ifs/:id", GetController().GetREQ_IF)
		v1.POST("/v1/req_ifs", GetController().PostREQ_IF)
		v1.PATCH("/v1/req_ifs/:id", GetController().UpdateREQ_IF)
		v1.PUT("/v1/req_ifs/:id", GetController().UpdateREQ_IF)
		v1.DELETE("/v1/req_ifs/:id", GetController().DeleteREQ_IF)

		v1.GET("/v1/req_if_contents", GetController().GetREQ_IF_CONTENTs)
		v1.GET("/v1/req_if_contents/:id", GetController().GetREQ_IF_CONTENT)
		v1.POST("/v1/req_if_contents", GetController().PostREQ_IF_CONTENT)
		v1.PATCH("/v1/req_if_contents/:id", GetController().UpdateREQ_IF_CONTENT)
		v1.PUT("/v1/req_if_contents/:id", GetController().UpdateREQ_IF_CONTENT)
		v1.DELETE("/v1/req_if_contents/:id", GetController().DeleteREQ_IF_CONTENT)

		v1.GET("/v1/req_if_headers", GetController().GetREQ_IF_HEADERs)
		v1.GET("/v1/req_if_headers/:id", GetController().GetREQ_IF_HEADER)
		v1.POST("/v1/req_if_headers", GetController().PostREQ_IF_HEADER)
		v1.PATCH("/v1/req_if_headers/:id", GetController().UpdateREQ_IF_HEADER)
		v1.PUT("/v1/req_if_headers/:id", GetController().UpdateREQ_IF_HEADER)
		v1.DELETE("/v1/req_if_headers/:id", GetController().DeleteREQ_IF_HEADER)

		v1.GET("/v1/req_if_tool_extensions", GetController().GetREQ_IF_TOOL_EXTENSIONs)
		v1.GET("/v1/req_if_tool_extensions/:id", GetController().GetREQ_IF_TOOL_EXTENSION)
		v1.POST("/v1/req_if_tool_extensions", GetController().PostREQ_IF_TOOL_EXTENSION)
		v1.PATCH("/v1/req_if_tool_extensions/:id", GetController().UpdateREQ_IF_TOOL_EXTENSION)
		v1.PUT("/v1/req_if_tool_extensions/:id", GetController().UpdateREQ_IF_TOOL_EXTENSION)
		v1.DELETE("/v1/req_if_tool_extensions/:id", GetController().DeleteREQ_IF_TOOL_EXTENSION)

		v1.GET("/v1/specifications", GetController().GetSPECIFICATIONs)
		v1.GET("/v1/specifications/:id", GetController().GetSPECIFICATION)
		v1.POST("/v1/specifications", GetController().PostSPECIFICATION)
		v1.PATCH("/v1/specifications/:id", GetController().UpdateSPECIFICATION)
		v1.PUT("/v1/specifications/:id", GetController().UpdateSPECIFICATION)
		v1.DELETE("/v1/specifications/:id", GetController().DeleteSPECIFICATION)

		v1.GET("/v1/specification_types", GetController().GetSPECIFICATION_TYPEs)
		v1.GET("/v1/specification_types/:id", GetController().GetSPECIFICATION_TYPE)
		v1.POST("/v1/specification_types", GetController().PostSPECIFICATION_TYPE)
		v1.PATCH("/v1/specification_types/:id", GetController().UpdateSPECIFICATION_TYPE)
		v1.PUT("/v1/specification_types/:id", GetController().UpdateSPECIFICATION_TYPE)
		v1.DELETE("/v1/specification_types/:id", GetController().DeleteSPECIFICATION_TYPE)

		v1.GET("/v1/spec_hierarchys", GetController().GetSPEC_HIERARCHYs)
		v1.GET("/v1/spec_hierarchys/:id", GetController().GetSPEC_HIERARCHY)
		v1.POST("/v1/spec_hierarchys", GetController().PostSPEC_HIERARCHY)
		v1.PATCH("/v1/spec_hierarchys/:id", GetController().UpdateSPEC_HIERARCHY)
		v1.PUT("/v1/spec_hierarchys/:id", GetController().UpdateSPEC_HIERARCHY)
		v1.DELETE("/v1/spec_hierarchys/:id", GetController().DeleteSPEC_HIERARCHY)

		v1.GET("/v1/spec_objects", GetController().GetSPEC_OBJECTs)
		v1.GET("/v1/spec_objects/:id", GetController().GetSPEC_OBJECT)
		v1.POST("/v1/spec_objects", GetController().PostSPEC_OBJECT)
		v1.PATCH("/v1/spec_objects/:id", GetController().UpdateSPEC_OBJECT)
		v1.PUT("/v1/spec_objects/:id", GetController().UpdateSPEC_OBJECT)
		v1.DELETE("/v1/spec_objects/:id", GetController().DeleteSPEC_OBJECT)

		v1.GET("/v1/spec_object_types", GetController().GetSPEC_OBJECT_TYPEs)
		v1.GET("/v1/spec_object_types/:id", GetController().GetSPEC_OBJECT_TYPE)
		v1.POST("/v1/spec_object_types", GetController().PostSPEC_OBJECT_TYPE)
		v1.PATCH("/v1/spec_object_types/:id", GetController().UpdateSPEC_OBJECT_TYPE)
		v1.PUT("/v1/spec_object_types/:id", GetController().UpdateSPEC_OBJECT_TYPE)
		v1.DELETE("/v1/spec_object_types/:id", GetController().DeleteSPEC_OBJECT_TYPE)

		v1.GET("/v1/spec_relations", GetController().GetSPEC_RELATIONs)
		v1.GET("/v1/spec_relations/:id", GetController().GetSPEC_RELATION)
		v1.POST("/v1/spec_relations", GetController().PostSPEC_RELATION)
		v1.PATCH("/v1/spec_relations/:id", GetController().UpdateSPEC_RELATION)
		v1.PUT("/v1/spec_relations/:id", GetController().UpdateSPEC_RELATION)
		v1.DELETE("/v1/spec_relations/:id", GetController().DeleteSPEC_RELATION)

		v1.GET("/v1/spec_relation_types", GetController().GetSPEC_RELATION_TYPEs)
		v1.GET("/v1/spec_relation_types/:id", GetController().GetSPEC_RELATION_TYPE)
		v1.POST("/v1/spec_relation_types", GetController().PostSPEC_RELATION_TYPE)
		v1.PATCH("/v1/spec_relation_types/:id", GetController().UpdateSPEC_RELATION_TYPE)
		v1.PUT("/v1/spec_relation_types/:id", GetController().UpdateSPEC_RELATION_TYPE)
		v1.DELETE("/v1/spec_relation_types/:id", GetController().DeleteSPEC_RELATION_TYPE)

		v1.GET("/v1/xhtml_contents", GetController().GetXHTML_CONTENTs)
		v1.GET("/v1/xhtml_contents/:id", GetController().GetXHTML_CONTENT)
		v1.POST("/v1/xhtml_contents", GetController().PostXHTML_CONTENT)
		v1.PATCH("/v1/xhtml_contents/:id", GetController().UpdateXHTML_CONTENT)
		v1.PUT("/v1/xhtml_contents/:id", GetController().UpdateXHTML_CONTENT)
		v1.DELETE("/v1/xhtml_contents/:id", GetController().DeleteXHTML_CONTENT)

		v1.GET("/v1/xhtml_inlpres_types", GetController().GetXhtml_InlPres_types)
		v1.GET("/v1/xhtml_inlpres_types/:id", GetController().GetXhtml_InlPres_type)
		v1.POST("/v1/xhtml_inlpres_types", GetController().PostXhtml_InlPres_type)
		v1.PATCH("/v1/xhtml_inlpres_types/:id", GetController().UpdateXhtml_InlPres_type)
		v1.PUT("/v1/xhtml_inlpres_types/:id", GetController().UpdateXhtml_InlPres_type)
		v1.DELETE("/v1/xhtml_inlpres_types/:id", GetController().DeleteXhtml_InlPres_type)

		v1.GET("/v1/xhtml_a_types", GetController().GetXhtml_a_types)
		v1.GET("/v1/xhtml_a_types/:id", GetController().GetXhtml_a_type)
		v1.POST("/v1/xhtml_a_types", GetController().PostXhtml_a_type)
		v1.PATCH("/v1/xhtml_a_types/:id", GetController().UpdateXhtml_a_type)
		v1.PUT("/v1/xhtml_a_types/:id", GetController().UpdateXhtml_a_type)
		v1.DELETE("/v1/xhtml_a_types/:id", GetController().DeleteXhtml_a_type)

		v1.GET("/v1/xhtml_abbr_types", GetController().GetXhtml_abbr_types)
		v1.GET("/v1/xhtml_abbr_types/:id", GetController().GetXhtml_abbr_type)
		v1.POST("/v1/xhtml_abbr_types", GetController().PostXhtml_abbr_type)
		v1.PATCH("/v1/xhtml_abbr_types/:id", GetController().UpdateXhtml_abbr_type)
		v1.PUT("/v1/xhtml_abbr_types/:id", GetController().UpdateXhtml_abbr_type)
		v1.DELETE("/v1/xhtml_abbr_types/:id", GetController().DeleteXhtml_abbr_type)

		v1.GET("/v1/xhtml_acronym_types", GetController().GetXhtml_acronym_types)
		v1.GET("/v1/xhtml_acronym_types/:id", GetController().GetXhtml_acronym_type)
		v1.POST("/v1/xhtml_acronym_types", GetController().PostXhtml_acronym_type)
		v1.PATCH("/v1/xhtml_acronym_types/:id", GetController().UpdateXhtml_acronym_type)
		v1.PUT("/v1/xhtml_acronym_types/:id", GetController().UpdateXhtml_acronym_type)
		v1.DELETE("/v1/xhtml_acronym_types/:id", GetController().DeleteXhtml_acronym_type)

		v1.GET("/v1/xhtml_address_types", GetController().GetXhtml_address_types)
		v1.GET("/v1/xhtml_address_types/:id", GetController().GetXhtml_address_type)
		v1.POST("/v1/xhtml_address_types", GetController().PostXhtml_address_type)
		v1.PATCH("/v1/xhtml_address_types/:id", GetController().UpdateXhtml_address_type)
		v1.PUT("/v1/xhtml_address_types/:id", GetController().UpdateXhtml_address_type)
		v1.DELETE("/v1/xhtml_address_types/:id", GetController().DeleteXhtml_address_type)

		v1.GET("/v1/xhtml_blockquote_types", GetController().GetXhtml_blockquote_types)
		v1.GET("/v1/xhtml_blockquote_types/:id", GetController().GetXhtml_blockquote_type)
		v1.POST("/v1/xhtml_blockquote_types", GetController().PostXhtml_blockquote_type)
		v1.PATCH("/v1/xhtml_blockquote_types/:id", GetController().UpdateXhtml_blockquote_type)
		v1.PUT("/v1/xhtml_blockquote_types/:id", GetController().UpdateXhtml_blockquote_type)
		v1.DELETE("/v1/xhtml_blockquote_types/:id", GetController().DeleteXhtml_blockquote_type)

		v1.GET("/v1/xhtml_br_types", GetController().GetXhtml_br_types)
		v1.GET("/v1/xhtml_br_types/:id", GetController().GetXhtml_br_type)
		v1.POST("/v1/xhtml_br_types", GetController().PostXhtml_br_type)
		v1.PATCH("/v1/xhtml_br_types/:id", GetController().UpdateXhtml_br_type)
		v1.PUT("/v1/xhtml_br_types/:id", GetController().UpdateXhtml_br_type)
		v1.DELETE("/v1/xhtml_br_types/:id", GetController().DeleteXhtml_br_type)

		v1.GET("/v1/xhtml_caption_types", GetController().GetXhtml_caption_types)
		v1.GET("/v1/xhtml_caption_types/:id", GetController().GetXhtml_caption_type)
		v1.POST("/v1/xhtml_caption_types", GetController().PostXhtml_caption_type)
		v1.PATCH("/v1/xhtml_caption_types/:id", GetController().UpdateXhtml_caption_type)
		v1.PUT("/v1/xhtml_caption_types/:id", GetController().UpdateXhtml_caption_type)
		v1.DELETE("/v1/xhtml_caption_types/:id", GetController().DeleteXhtml_caption_type)

		v1.GET("/v1/xhtml_cite_types", GetController().GetXhtml_cite_types)
		v1.GET("/v1/xhtml_cite_types/:id", GetController().GetXhtml_cite_type)
		v1.POST("/v1/xhtml_cite_types", GetController().PostXhtml_cite_type)
		v1.PATCH("/v1/xhtml_cite_types/:id", GetController().UpdateXhtml_cite_type)
		v1.PUT("/v1/xhtml_cite_types/:id", GetController().UpdateXhtml_cite_type)
		v1.DELETE("/v1/xhtml_cite_types/:id", GetController().DeleteXhtml_cite_type)

		v1.GET("/v1/xhtml_code_types", GetController().GetXhtml_code_types)
		v1.GET("/v1/xhtml_code_types/:id", GetController().GetXhtml_code_type)
		v1.POST("/v1/xhtml_code_types", GetController().PostXhtml_code_type)
		v1.PATCH("/v1/xhtml_code_types/:id", GetController().UpdateXhtml_code_type)
		v1.PUT("/v1/xhtml_code_types/:id", GetController().UpdateXhtml_code_type)
		v1.DELETE("/v1/xhtml_code_types/:id", GetController().DeleteXhtml_code_type)

		v1.GET("/v1/xhtml_col_types", GetController().GetXhtml_col_types)
		v1.GET("/v1/xhtml_col_types/:id", GetController().GetXhtml_col_type)
		v1.POST("/v1/xhtml_col_types", GetController().PostXhtml_col_type)
		v1.PATCH("/v1/xhtml_col_types/:id", GetController().UpdateXhtml_col_type)
		v1.PUT("/v1/xhtml_col_types/:id", GetController().UpdateXhtml_col_type)
		v1.DELETE("/v1/xhtml_col_types/:id", GetController().DeleteXhtml_col_type)

		v1.GET("/v1/xhtml_colgroup_types", GetController().GetXhtml_colgroup_types)
		v1.GET("/v1/xhtml_colgroup_types/:id", GetController().GetXhtml_colgroup_type)
		v1.POST("/v1/xhtml_colgroup_types", GetController().PostXhtml_colgroup_type)
		v1.PATCH("/v1/xhtml_colgroup_types/:id", GetController().UpdateXhtml_colgroup_type)
		v1.PUT("/v1/xhtml_colgroup_types/:id", GetController().UpdateXhtml_colgroup_type)
		v1.DELETE("/v1/xhtml_colgroup_types/:id", GetController().DeleteXhtml_colgroup_type)

		v1.GET("/v1/xhtml_dd_types", GetController().GetXhtml_dd_types)
		v1.GET("/v1/xhtml_dd_types/:id", GetController().GetXhtml_dd_type)
		v1.POST("/v1/xhtml_dd_types", GetController().PostXhtml_dd_type)
		v1.PATCH("/v1/xhtml_dd_types/:id", GetController().UpdateXhtml_dd_type)
		v1.PUT("/v1/xhtml_dd_types/:id", GetController().UpdateXhtml_dd_type)
		v1.DELETE("/v1/xhtml_dd_types/:id", GetController().DeleteXhtml_dd_type)

		v1.GET("/v1/xhtml_dfn_types", GetController().GetXhtml_dfn_types)
		v1.GET("/v1/xhtml_dfn_types/:id", GetController().GetXhtml_dfn_type)
		v1.POST("/v1/xhtml_dfn_types", GetController().PostXhtml_dfn_type)
		v1.PATCH("/v1/xhtml_dfn_types/:id", GetController().UpdateXhtml_dfn_type)
		v1.PUT("/v1/xhtml_dfn_types/:id", GetController().UpdateXhtml_dfn_type)
		v1.DELETE("/v1/xhtml_dfn_types/:id", GetController().DeleteXhtml_dfn_type)

		v1.GET("/v1/xhtml_div_types", GetController().GetXhtml_div_types)
		v1.GET("/v1/xhtml_div_types/:id", GetController().GetXhtml_div_type)
		v1.POST("/v1/xhtml_div_types", GetController().PostXhtml_div_type)
		v1.PATCH("/v1/xhtml_div_types/:id", GetController().UpdateXhtml_div_type)
		v1.PUT("/v1/xhtml_div_types/:id", GetController().UpdateXhtml_div_type)
		v1.DELETE("/v1/xhtml_div_types/:id", GetController().DeleteXhtml_div_type)

		v1.GET("/v1/xhtml_dl_types", GetController().GetXhtml_dl_types)
		v1.GET("/v1/xhtml_dl_types/:id", GetController().GetXhtml_dl_type)
		v1.POST("/v1/xhtml_dl_types", GetController().PostXhtml_dl_type)
		v1.PATCH("/v1/xhtml_dl_types/:id", GetController().UpdateXhtml_dl_type)
		v1.PUT("/v1/xhtml_dl_types/:id", GetController().UpdateXhtml_dl_type)
		v1.DELETE("/v1/xhtml_dl_types/:id", GetController().DeleteXhtml_dl_type)

		v1.GET("/v1/xhtml_dt_types", GetController().GetXhtml_dt_types)
		v1.GET("/v1/xhtml_dt_types/:id", GetController().GetXhtml_dt_type)
		v1.POST("/v1/xhtml_dt_types", GetController().PostXhtml_dt_type)
		v1.PATCH("/v1/xhtml_dt_types/:id", GetController().UpdateXhtml_dt_type)
		v1.PUT("/v1/xhtml_dt_types/:id", GetController().UpdateXhtml_dt_type)
		v1.DELETE("/v1/xhtml_dt_types/:id", GetController().DeleteXhtml_dt_type)

		v1.GET("/v1/xhtml_edit_types", GetController().GetXhtml_edit_types)
		v1.GET("/v1/xhtml_edit_types/:id", GetController().GetXhtml_edit_type)
		v1.POST("/v1/xhtml_edit_types", GetController().PostXhtml_edit_type)
		v1.PATCH("/v1/xhtml_edit_types/:id", GetController().UpdateXhtml_edit_type)
		v1.PUT("/v1/xhtml_edit_types/:id", GetController().UpdateXhtml_edit_type)
		v1.DELETE("/v1/xhtml_edit_types/:id", GetController().DeleteXhtml_edit_type)

		v1.GET("/v1/xhtml_em_types", GetController().GetXhtml_em_types)
		v1.GET("/v1/xhtml_em_types/:id", GetController().GetXhtml_em_type)
		v1.POST("/v1/xhtml_em_types", GetController().PostXhtml_em_type)
		v1.PATCH("/v1/xhtml_em_types/:id", GetController().UpdateXhtml_em_type)
		v1.PUT("/v1/xhtml_em_types/:id", GetController().UpdateXhtml_em_type)
		v1.DELETE("/v1/xhtml_em_types/:id", GetController().DeleteXhtml_em_type)

		v1.GET("/v1/xhtml_h1_types", GetController().GetXhtml_h1_types)
		v1.GET("/v1/xhtml_h1_types/:id", GetController().GetXhtml_h1_type)
		v1.POST("/v1/xhtml_h1_types", GetController().PostXhtml_h1_type)
		v1.PATCH("/v1/xhtml_h1_types/:id", GetController().UpdateXhtml_h1_type)
		v1.PUT("/v1/xhtml_h1_types/:id", GetController().UpdateXhtml_h1_type)
		v1.DELETE("/v1/xhtml_h1_types/:id", GetController().DeleteXhtml_h1_type)

		v1.GET("/v1/xhtml_h2_types", GetController().GetXhtml_h2_types)
		v1.GET("/v1/xhtml_h2_types/:id", GetController().GetXhtml_h2_type)
		v1.POST("/v1/xhtml_h2_types", GetController().PostXhtml_h2_type)
		v1.PATCH("/v1/xhtml_h2_types/:id", GetController().UpdateXhtml_h2_type)
		v1.PUT("/v1/xhtml_h2_types/:id", GetController().UpdateXhtml_h2_type)
		v1.DELETE("/v1/xhtml_h2_types/:id", GetController().DeleteXhtml_h2_type)

		v1.GET("/v1/xhtml_h3_types", GetController().GetXhtml_h3_types)
		v1.GET("/v1/xhtml_h3_types/:id", GetController().GetXhtml_h3_type)
		v1.POST("/v1/xhtml_h3_types", GetController().PostXhtml_h3_type)
		v1.PATCH("/v1/xhtml_h3_types/:id", GetController().UpdateXhtml_h3_type)
		v1.PUT("/v1/xhtml_h3_types/:id", GetController().UpdateXhtml_h3_type)
		v1.DELETE("/v1/xhtml_h3_types/:id", GetController().DeleteXhtml_h3_type)

		v1.GET("/v1/xhtml_h4_types", GetController().GetXhtml_h4_types)
		v1.GET("/v1/xhtml_h4_types/:id", GetController().GetXhtml_h4_type)
		v1.POST("/v1/xhtml_h4_types", GetController().PostXhtml_h4_type)
		v1.PATCH("/v1/xhtml_h4_types/:id", GetController().UpdateXhtml_h4_type)
		v1.PUT("/v1/xhtml_h4_types/:id", GetController().UpdateXhtml_h4_type)
		v1.DELETE("/v1/xhtml_h4_types/:id", GetController().DeleteXhtml_h4_type)

		v1.GET("/v1/xhtml_h5_types", GetController().GetXhtml_h5_types)
		v1.GET("/v1/xhtml_h5_types/:id", GetController().GetXhtml_h5_type)
		v1.POST("/v1/xhtml_h5_types", GetController().PostXhtml_h5_type)
		v1.PATCH("/v1/xhtml_h5_types/:id", GetController().UpdateXhtml_h5_type)
		v1.PUT("/v1/xhtml_h5_types/:id", GetController().UpdateXhtml_h5_type)
		v1.DELETE("/v1/xhtml_h5_types/:id", GetController().DeleteXhtml_h5_type)

		v1.GET("/v1/xhtml_h6_types", GetController().GetXhtml_h6_types)
		v1.GET("/v1/xhtml_h6_types/:id", GetController().GetXhtml_h6_type)
		v1.POST("/v1/xhtml_h6_types", GetController().PostXhtml_h6_type)
		v1.PATCH("/v1/xhtml_h6_types/:id", GetController().UpdateXhtml_h6_type)
		v1.PUT("/v1/xhtml_h6_types/:id", GetController().UpdateXhtml_h6_type)
		v1.DELETE("/v1/xhtml_h6_types/:id", GetController().DeleteXhtml_h6_type)

		v1.GET("/v1/xhtml_heading_types", GetController().GetXhtml_heading_types)
		v1.GET("/v1/xhtml_heading_types/:id", GetController().GetXhtml_heading_type)
		v1.POST("/v1/xhtml_heading_types", GetController().PostXhtml_heading_type)
		v1.PATCH("/v1/xhtml_heading_types/:id", GetController().UpdateXhtml_heading_type)
		v1.PUT("/v1/xhtml_heading_types/:id", GetController().UpdateXhtml_heading_type)
		v1.DELETE("/v1/xhtml_heading_types/:id", GetController().DeleteXhtml_heading_type)

		v1.GET("/v1/xhtml_hr_types", GetController().GetXhtml_hr_types)
		v1.GET("/v1/xhtml_hr_types/:id", GetController().GetXhtml_hr_type)
		v1.POST("/v1/xhtml_hr_types", GetController().PostXhtml_hr_type)
		v1.PATCH("/v1/xhtml_hr_types/:id", GetController().UpdateXhtml_hr_type)
		v1.PUT("/v1/xhtml_hr_types/:id", GetController().UpdateXhtml_hr_type)
		v1.DELETE("/v1/xhtml_hr_types/:id", GetController().DeleteXhtml_hr_type)

		v1.GET("/v1/xhtml_kbd_types", GetController().GetXhtml_kbd_types)
		v1.GET("/v1/xhtml_kbd_types/:id", GetController().GetXhtml_kbd_type)
		v1.POST("/v1/xhtml_kbd_types", GetController().PostXhtml_kbd_type)
		v1.PATCH("/v1/xhtml_kbd_types/:id", GetController().UpdateXhtml_kbd_type)
		v1.PUT("/v1/xhtml_kbd_types/:id", GetController().UpdateXhtml_kbd_type)
		v1.DELETE("/v1/xhtml_kbd_types/:id", GetController().DeleteXhtml_kbd_type)

		v1.GET("/v1/xhtml_li_types", GetController().GetXhtml_li_types)
		v1.GET("/v1/xhtml_li_types/:id", GetController().GetXhtml_li_type)
		v1.POST("/v1/xhtml_li_types", GetController().PostXhtml_li_type)
		v1.PATCH("/v1/xhtml_li_types/:id", GetController().UpdateXhtml_li_type)
		v1.PUT("/v1/xhtml_li_types/:id", GetController().UpdateXhtml_li_type)
		v1.DELETE("/v1/xhtml_li_types/:id", GetController().DeleteXhtml_li_type)

		v1.GET("/v1/xhtml_object_types", GetController().GetXhtml_object_types)
		v1.GET("/v1/xhtml_object_types/:id", GetController().GetXhtml_object_type)
		v1.POST("/v1/xhtml_object_types", GetController().PostXhtml_object_type)
		v1.PATCH("/v1/xhtml_object_types/:id", GetController().UpdateXhtml_object_type)
		v1.PUT("/v1/xhtml_object_types/:id", GetController().UpdateXhtml_object_type)
		v1.DELETE("/v1/xhtml_object_types/:id", GetController().DeleteXhtml_object_type)

		v1.GET("/v1/xhtml_ol_types", GetController().GetXhtml_ol_types)
		v1.GET("/v1/xhtml_ol_types/:id", GetController().GetXhtml_ol_type)
		v1.POST("/v1/xhtml_ol_types", GetController().PostXhtml_ol_type)
		v1.PATCH("/v1/xhtml_ol_types/:id", GetController().UpdateXhtml_ol_type)
		v1.PUT("/v1/xhtml_ol_types/:id", GetController().UpdateXhtml_ol_type)
		v1.DELETE("/v1/xhtml_ol_types/:id", GetController().DeleteXhtml_ol_type)

		v1.GET("/v1/xhtml_p_types", GetController().GetXhtml_p_types)
		v1.GET("/v1/xhtml_p_types/:id", GetController().GetXhtml_p_type)
		v1.POST("/v1/xhtml_p_types", GetController().PostXhtml_p_type)
		v1.PATCH("/v1/xhtml_p_types/:id", GetController().UpdateXhtml_p_type)
		v1.PUT("/v1/xhtml_p_types/:id", GetController().UpdateXhtml_p_type)
		v1.DELETE("/v1/xhtml_p_types/:id", GetController().DeleteXhtml_p_type)

		v1.GET("/v1/xhtml_param_types", GetController().GetXhtml_param_types)
		v1.GET("/v1/xhtml_param_types/:id", GetController().GetXhtml_param_type)
		v1.POST("/v1/xhtml_param_types", GetController().PostXhtml_param_type)
		v1.PATCH("/v1/xhtml_param_types/:id", GetController().UpdateXhtml_param_type)
		v1.PUT("/v1/xhtml_param_types/:id", GetController().UpdateXhtml_param_type)
		v1.DELETE("/v1/xhtml_param_types/:id", GetController().DeleteXhtml_param_type)

		v1.GET("/v1/xhtml_pre_types", GetController().GetXhtml_pre_types)
		v1.GET("/v1/xhtml_pre_types/:id", GetController().GetXhtml_pre_type)
		v1.POST("/v1/xhtml_pre_types", GetController().PostXhtml_pre_type)
		v1.PATCH("/v1/xhtml_pre_types/:id", GetController().UpdateXhtml_pre_type)
		v1.PUT("/v1/xhtml_pre_types/:id", GetController().UpdateXhtml_pre_type)
		v1.DELETE("/v1/xhtml_pre_types/:id", GetController().DeleteXhtml_pre_type)

		v1.GET("/v1/xhtml_q_types", GetController().GetXhtml_q_types)
		v1.GET("/v1/xhtml_q_types/:id", GetController().GetXhtml_q_type)
		v1.POST("/v1/xhtml_q_types", GetController().PostXhtml_q_type)
		v1.PATCH("/v1/xhtml_q_types/:id", GetController().UpdateXhtml_q_type)
		v1.PUT("/v1/xhtml_q_types/:id", GetController().UpdateXhtml_q_type)
		v1.DELETE("/v1/xhtml_q_types/:id", GetController().DeleteXhtml_q_type)

		v1.GET("/v1/xhtml_samp_types", GetController().GetXhtml_samp_types)
		v1.GET("/v1/xhtml_samp_types/:id", GetController().GetXhtml_samp_type)
		v1.POST("/v1/xhtml_samp_types", GetController().PostXhtml_samp_type)
		v1.PATCH("/v1/xhtml_samp_types/:id", GetController().UpdateXhtml_samp_type)
		v1.PUT("/v1/xhtml_samp_types/:id", GetController().UpdateXhtml_samp_type)
		v1.DELETE("/v1/xhtml_samp_types/:id", GetController().DeleteXhtml_samp_type)

		v1.GET("/v1/xhtml_span_types", GetController().GetXhtml_span_types)
		v1.GET("/v1/xhtml_span_types/:id", GetController().GetXhtml_span_type)
		v1.POST("/v1/xhtml_span_types", GetController().PostXhtml_span_type)
		v1.PATCH("/v1/xhtml_span_types/:id", GetController().UpdateXhtml_span_type)
		v1.PUT("/v1/xhtml_span_types/:id", GetController().UpdateXhtml_span_type)
		v1.DELETE("/v1/xhtml_span_types/:id", GetController().DeleteXhtml_span_type)

		v1.GET("/v1/xhtml_strong_types", GetController().GetXhtml_strong_types)
		v1.GET("/v1/xhtml_strong_types/:id", GetController().GetXhtml_strong_type)
		v1.POST("/v1/xhtml_strong_types", GetController().PostXhtml_strong_type)
		v1.PATCH("/v1/xhtml_strong_types/:id", GetController().UpdateXhtml_strong_type)
		v1.PUT("/v1/xhtml_strong_types/:id", GetController().UpdateXhtml_strong_type)
		v1.DELETE("/v1/xhtml_strong_types/:id", GetController().DeleteXhtml_strong_type)

		v1.GET("/v1/xhtml_table_types", GetController().GetXhtml_table_types)
		v1.GET("/v1/xhtml_table_types/:id", GetController().GetXhtml_table_type)
		v1.POST("/v1/xhtml_table_types", GetController().PostXhtml_table_type)
		v1.PATCH("/v1/xhtml_table_types/:id", GetController().UpdateXhtml_table_type)
		v1.PUT("/v1/xhtml_table_types/:id", GetController().UpdateXhtml_table_type)
		v1.DELETE("/v1/xhtml_table_types/:id", GetController().DeleteXhtml_table_type)

		v1.GET("/v1/xhtml_tbody_types", GetController().GetXhtml_tbody_types)
		v1.GET("/v1/xhtml_tbody_types/:id", GetController().GetXhtml_tbody_type)
		v1.POST("/v1/xhtml_tbody_types", GetController().PostXhtml_tbody_type)
		v1.PATCH("/v1/xhtml_tbody_types/:id", GetController().UpdateXhtml_tbody_type)
		v1.PUT("/v1/xhtml_tbody_types/:id", GetController().UpdateXhtml_tbody_type)
		v1.DELETE("/v1/xhtml_tbody_types/:id", GetController().DeleteXhtml_tbody_type)

		v1.GET("/v1/xhtml_td_types", GetController().GetXhtml_td_types)
		v1.GET("/v1/xhtml_td_types/:id", GetController().GetXhtml_td_type)
		v1.POST("/v1/xhtml_td_types", GetController().PostXhtml_td_type)
		v1.PATCH("/v1/xhtml_td_types/:id", GetController().UpdateXhtml_td_type)
		v1.PUT("/v1/xhtml_td_types/:id", GetController().UpdateXhtml_td_type)
		v1.DELETE("/v1/xhtml_td_types/:id", GetController().DeleteXhtml_td_type)

		v1.GET("/v1/xhtml_tfoot_types", GetController().GetXhtml_tfoot_types)
		v1.GET("/v1/xhtml_tfoot_types/:id", GetController().GetXhtml_tfoot_type)
		v1.POST("/v1/xhtml_tfoot_types", GetController().PostXhtml_tfoot_type)
		v1.PATCH("/v1/xhtml_tfoot_types/:id", GetController().UpdateXhtml_tfoot_type)
		v1.PUT("/v1/xhtml_tfoot_types/:id", GetController().UpdateXhtml_tfoot_type)
		v1.DELETE("/v1/xhtml_tfoot_types/:id", GetController().DeleteXhtml_tfoot_type)

		v1.GET("/v1/xhtml_th_types", GetController().GetXhtml_th_types)
		v1.GET("/v1/xhtml_th_types/:id", GetController().GetXhtml_th_type)
		v1.POST("/v1/xhtml_th_types", GetController().PostXhtml_th_type)
		v1.PATCH("/v1/xhtml_th_types/:id", GetController().UpdateXhtml_th_type)
		v1.PUT("/v1/xhtml_th_types/:id", GetController().UpdateXhtml_th_type)
		v1.DELETE("/v1/xhtml_th_types/:id", GetController().DeleteXhtml_th_type)

		v1.GET("/v1/xhtml_thead_types", GetController().GetXhtml_thead_types)
		v1.GET("/v1/xhtml_thead_types/:id", GetController().GetXhtml_thead_type)
		v1.POST("/v1/xhtml_thead_types", GetController().PostXhtml_thead_type)
		v1.PATCH("/v1/xhtml_thead_types/:id", GetController().UpdateXhtml_thead_type)
		v1.PUT("/v1/xhtml_thead_types/:id", GetController().UpdateXhtml_thead_type)
		v1.DELETE("/v1/xhtml_thead_types/:id", GetController().DeleteXhtml_thead_type)

		v1.GET("/v1/xhtml_tr_types", GetController().GetXhtml_tr_types)
		v1.GET("/v1/xhtml_tr_types/:id", GetController().GetXhtml_tr_type)
		v1.POST("/v1/xhtml_tr_types", GetController().PostXhtml_tr_type)
		v1.PATCH("/v1/xhtml_tr_types/:id", GetController().UpdateXhtml_tr_type)
		v1.PUT("/v1/xhtml_tr_types/:id", GetController().UpdateXhtml_tr_type)
		v1.DELETE("/v1/xhtml_tr_types/:id", GetController().DeleteXhtml_tr_type)

		v1.GET("/v1/xhtml_ul_types", GetController().GetXhtml_ul_types)
		v1.GET("/v1/xhtml_ul_types/:id", GetController().GetXhtml_ul_type)
		v1.POST("/v1/xhtml_ul_types", GetController().PostXhtml_ul_type)
		v1.PATCH("/v1/xhtml_ul_types/:id", GetController().UpdateXhtml_ul_type)
		v1.PUT("/v1/xhtml_ul_types/:id", GetController().UpdateXhtml_ul_type)
		v1.DELETE("/v1/xhtml_ul_types/:id", GetController().DeleteXhtml_ul_type)

		v1.GET("/v1/xhtml_var_types", GetController().GetXhtml_var_types)
		v1.GET("/v1/xhtml_var_types/:id", GetController().GetXhtml_var_type)
		v1.POST("/v1/xhtml_var_types", GetController().PostXhtml_var_type)
		v1.PATCH("/v1/xhtml_var_types/:id", GetController().UpdateXhtml_var_type)
		v1.PUT("/v1/xhtml_var_types/:id", GetController().UpdateXhtml_var_type)
		v1.DELETE("/v1/xhtml_var_types/:id", GetController().DeleteXhtml_var_type)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)

		v1.GET("/v1/ws/stage", GetController().onWebSocketRequestForBackRepoContent)

		v1.GET("/v1/stacks", GetController().stacks)
	}
}

func (controller *Controller) stacks(c *gin.Context) {

	var res []string

	for k := range controller.Map_BackRepos {
		res = append(res, k)
	}

	c.JSON(http.StatusOK, res)
}

// onWebSocketRequestForBackRepoContent is a function that is started each time
// a web socket request is received
//
// 1. upgrade the incomming web connection to a web socket
// 1. it subscribe to the backend commit number broadcaster
// 1. it stays live and pool for incomming backend commit number broadcast and forward
// them on the web socket connection
func (controller *Controller) onWebSocketRequestForBackRepoContent(c *gin.Context) {

	// log.Println("Stack github.com/fullstack-lang/gongreqif/go, onWebSocketRequestForBackRepoContent")

	// Upgrader specifies parameters for upgrading an HTTP connection to a
	// WebSocket connection.
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:8080" || origin == "http://localhost:4200"
		},
	}

	wsConnection, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wsConnection.Close()

	// Create a context that is canceled when the connection is closed
	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "Name", stackPath)
		}
	}

	index := controller.listenerIndex
	controller.listenerIndex++
	log.Printf(
		"%s github.com/fullstack-lang/gongreqif/go: Con: '%s', index %d",
		time.Now().Format("2006-01-02 15:04:05.000000"),
		stackPath, index,
	)

	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "WebSocket Connect, Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	updateCommitBackRepoNbChannel := backRepo.SubscribeToCommitNb(ctx)

	// Start a goroutine to read from the WebSocket to detect disconnection
	go func() {
		for {
			// ReadMessage is used to detect client disconnection
			_, _, err := wsConnection.ReadMessage()
			if err != nil {
				log.Println("github.com/fullstack-lang/gongreqif/go", stackPath, "WS client disconnected:", err)
				cancel() // Cancel the context
				return
			}
		}
	}()

	backRepoData := new(orm.BackRepoData)
	orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
	backRepoData.GONG__Index = index
	
	refresh := 0
	err = wsConnection.WriteJSON(backRepoData)
	if err != nil {
		log.Println("github.com/fullstack-lang/gongreqif/go:\n",
			"client no longer receiver web socket message, assuming it is no longer alive, closing websocket handler")
		fmt.Println(err)
		return
	} else {
	log.Printf(
		"%s github.com/fullstack-lang/gongreqif/go: %03d: '%s', index %d",
		time.Now().Format("2006-01-02 15:04:05.000000"),
		refresh,
		stackPath, index,
	)
	}
	for {
		select {
		case <-ctx.Done():
			// Context canceled, exit the loop
			return
		default:
			for nbCommitBackRepo := range updateCommitBackRepoNbChannel {
				_ = nbCommitBackRepo
				refresh += 1

				backRepoData := new(orm.BackRepoData)
				orm.CopyBackRepoToBackRepoData(backRepo, backRepoData)
				backRepoData.GONG__Index = index

				// Set write deadline to prevent blocking indefinitely
				wsConnection.SetWriteDeadline(time.Now().Add(10 * time.Second))

				// Send backRepo data
				err = wsConnection.WriteJSON(backRepoData)
				if err != nil {
					log.Println("github.com/fullstack-lang/gongreqif/go:\n", stackPath,
						"client no longer receiver web socket message,closing websocket handler")
					fmt.Println(err)
					cancel() // Cancel the context
					return
				} else {
					log.Printf(
						"%s github.com/fullstack-lang/gongreqif/go: %03d: '%s', index %d",
						time.Now().Format("2006-01-02 15:04:05.000000"),
						refresh,
						stackPath, index,
					)
				}
			}
		}
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
