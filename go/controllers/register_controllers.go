// generated code - do not edit
package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
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

		v1.GET("/v1/a_alternative_ids", GetController().GetA_ALTERNATIVE_IDs)
		v1.GET("/v1/a_alternative_ids/:id", GetController().GetA_ALTERNATIVE_ID)
		v1.POST("/v1/a_alternative_ids", GetController().PostA_ALTERNATIVE_ID)
		v1.PATCH("/v1/a_alternative_ids/:id", GetController().UpdateA_ALTERNATIVE_ID)
		v1.PUT("/v1/a_alternative_ids/:id", GetController().UpdateA_ALTERNATIVE_ID)
		v1.DELETE("/v1/a_alternative_ids/:id", GetController().DeleteA_ALTERNATIVE_ID)

		v1.GET("/v1/a_attribute_definition_boolean_refs", GetController().GetA_ATTRIBUTE_DEFINITION_BOOLEAN_REFs)
		v1.GET("/v1/a_attribute_definition_boolean_refs/:id", GetController().GetA_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		v1.POST("/v1/a_attribute_definition_boolean_refs", GetController().PostA_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		v1.PATCH("/v1/a_attribute_definition_boolean_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		v1.PUT("/v1/a_attribute_definition_boolean_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		v1.DELETE("/v1/a_attribute_definition_boolean_refs/:id", GetController().DeleteA_ATTRIBUTE_DEFINITION_BOOLEAN_REF)

		v1.GET("/v1/a_attribute_definition_date_refs", GetController().GetA_ATTRIBUTE_DEFINITION_DATE_REFs)
		v1.GET("/v1/a_attribute_definition_date_refs/:id", GetController().GetA_ATTRIBUTE_DEFINITION_DATE_REF)
		v1.POST("/v1/a_attribute_definition_date_refs", GetController().PostA_ATTRIBUTE_DEFINITION_DATE_REF)
		v1.PATCH("/v1/a_attribute_definition_date_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_DATE_REF)
		v1.PUT("/v1/a_attribute_definition_date_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_DATE_REF)
		v1.DELETE("/v1/a_attribute_definition_date_refs/:id", GetController().DeleteA_ATTRIBUTE_DEFINITION_DATE_REF)

		v1.GET("/v1/a_attribute_definition_enumeration_refs", GetController().GetA_ATTRIBUTE_DEFINITION_ENUMERATION_REFs)
		v1.GET("/v1/a_attribute_definition_enumeration_refs/:id", GetController().GetA_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		v1.POST("/v1/a_attribute_definition_enumeration_refs", GetController().PostA_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		v1.PATCH("/v1/a_attribute_definition_enumeration_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		v1.PUT("/v1/a_attribute_definition_enumeration_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		v1.DELETE("/v1/a_attribute_definition_enumeration_refs/:id", GetController().DeleteA_ATTRIBUTE_DEFINITION_ENUMERATION_REF)

		v1.GET("/v1/a_attribute_definition_integer_refs", GetController().GetA_ATTRIBUTE_DEFINITION_INTEGER_REFs)
		v1.GET("/v1/a_attribute_definition_integer_refs/:id", GetController().GetA_ATTRIBUTE_DEFINITION_INTEGER_REF)
		v1.POST("/v1/a_attribute_definition_integer_refs", GetController().PostA_ATTRIBUTE_DEFINITION_INTEGER_REF)
		v1.PATCH("/v1/a_attribute_definition_integer_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_INTEGER_REF)
		v1.PUT("/v1/a_attribute_definition_integer_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_INTEGER_REF)
		v1.DELETE("/v1/a_attribute_definition_integer_refs/:id", GetController().DeleteA_ATTRIBUTE_DEFINITION_INTEGER_REF)

		v1.GET("/v1/a_attribute_definition_real_refs", GetController().GetA_ATTRIBUTE_DEFINITION_REAL_REFs)
		v1.GET("/v1/a_attribute_definition_real_refs/:id", GetController().GetA_ATTRIBUTE_DEFINITION_REAL_REF)
		v1.POST("/v1/a_attribute_definition_real_refs", GetController().PostA_ATTRIBUTE_DEFINITION_REAL_REF)
		v1.PATCH("/v1/a_attribute_definition_real_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_REAL_REF)
		v1.PUT("/v1/a_attribute_definition_real_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_REAL_REF)
		v1.DELETE("/v1/a_attribute_definition_real_refs/:id", GetController().DeleteA_ATTRIBUTE_DEFINITION_REAL_REF)

		v1.GET("/v1/a_attribute_definition_string_refs", GetController().GetA_ATTRIBUTE_DEFINITION_STRING_REFs)
		v1.GET("/v1/a_attribute_definition_string_refs/:id", GetController().GetA_ATTRIBUTE_DEFINITION_STRING_REF)
		v1.POST("/v1/a_attribute_definition_string_refs", GetController().PostA_ATTRIBUTE_DEFINITION_STRING_REF)
		v1.PATCH("/v1/a_attribute_definition_string_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_STRING_REF)
		v1.PUT("/v1/a_attribute_definition_string_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_STRING_REF)
		v1.DELETE("/v1/a_attribute_definition_string_refs/:id", GetController().DeleteA_ATTRIBUTE_DEFINITION_STRING_REF)

		v1.GET("/v1/a_attribute_definition_xhtml_refs", GetController().GetA_ATTRIBUTE_DEFINITION_XHTML_REFs)
		v1.GET("/v1/a_attribute_definition_xhtml_refs/:id", GetController().GetA_ATTRIBUTE_DEFINITION_XHTML_REF)
		v1.POST("/v1/a_attribute_definition_xhtml_refs", GetController().PostA_ATTRIBUTE_DEFINITION_XHTML_REF)
		v1.PATCH("/v1/a_attribute_definition_xhtml_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_XHTML_REF)
		v1.PUT("/v1/a_attribute_definition_xhtml_refs/:id", GetController().UpdateA_ATTRIBUTE_DEFINITION_XHTML_REF)
		v1.DELETE("/v1/a_attribute_definition_xhtml_refs/:id", GetController().DeleteA_ATTRIBUTE_DEFINITION_XHTML_REF)

		v1.GET("/v1/a_attribute_value_booleans", GetController().GetA_ATTRIBUTE_VALUE_BOOLEANs)
		v1.GET("/v1/a_attribute_value_booleans/:id", GetController().GetA_ATTRIBUTE_VALUE_BOOLEAN)
		v1.POST("/v1/a_attribute_value_booleans", GetController().PostA_ATTRIBUTE_VALUE_BOOLEAN)
		v1.PATCH("/v1/a_attribute_value_booleans/:id", GetController().UpdateA_ATTRIBUTE_VALUE_BOOLEAN)
		v1.PUT("/v1/a_attribute_value_booleans/:id", GetController().UpdateA_ATTRIBUTE_VALUE_BOOLEAN)
		v1.DELETE("/v1/a_attribute_value_booleans/:id", GetController().DeleteA_ATTRIBUTE_VALUE_BOOLEAN)

		v1.GET("/v1/a_attribute_value_dates", GetController().GetA_ATTRIBUTE_VALUE_DATEs)
		v1.GET("/v1/a_attribute_value_dates/:id", GetController().GetA_ATTRIBUTE_VALUE_DATE)
		v1.POST("/v1/a_attribute_value_dates", GetController().PostA_ATTRIBUTE_VALUE_DATE)
		v1.PATCH("/v1/a_attribute_value_dates/:id", GetController().UpdateA_ATTRIBUTE_VALUE_DATE)
		v1.PUT("/v1/a_attribute_value_dates/:id", GetController().UpdateA_ATTRIBUTE_VALUE_DATE)
		v1.DELETE("/v1/a_attribute_value_dates/:id", GetController().DeleteA_ATTRIBUTE_VALUE_DATE)

		v1.GET("/v1/a_attribute_value_enumerations", GetController().GetA_ATTRIBUTE_VALUE_ENUMERATIONs)
		v1.GET("/v1/a_attribute_value_enumerations/:id", GetController().GetA_ATTRIBUTE_VALUE_ENUMERATION)
		v1.POST("/v1/a_attribute_value_enumerations", GetController().PostA_ATTRIBUTE_VALUE_ENUMERATION)
		v1.PATCH("/v1/a_attribute_value_enumerations/:id", GetController().UpdateA_ATTRIBUTE_VALUE_ENUMERATION)
		v1.PUT("/v1/a_attribute_value_enumerations/:id", GetController().UpdateA_ATTRIBUTE_VALUE_ENUMERATION)
		v1.DELETE("/v1/a_attribute_value_enumerations/:id", GetController().DeleteA_ATTRIBUTE_VALUE_ENUMERATION)

		v1.GET("/v1/a_attribute_value_integers", GetController().GetA_ATTRIBUTE_VALUE_INTEGERs)
		v1.GET("/v1/a_attribute_value_integers/:id", GetController().GetA_ATTRIBUTE_VALUE_INTEGER)
		v1.POST("/v1/a_attribute_value_integers", GetController().PostA_ATTRIBUTE_VALUE_INTEGER)
		v1.PATCH("/v1/a_attribute_value_integers/:id", GetController().UpdateA_ATTRIBUTE_VALUE_INTEGER)
		v1.PUT("/v1/a_attribute_value_integers/:id", GetController().UpdateA_ATTRIBUTE_VALUE_INTEGER)
		v1.DELETE("/v1/a_attribute_value_integers/:id", GetController().DeleteA_ATTRIBUTE_VALUE_INTEGER)

		v1.GET("/v1/a_attribute_value_reals", GetController().GetA_ATTRIBUTE_VALUE_REALs)
		v1.GET("/v1/a_attribute_value_reals/:id", GetController().GetA_ATTRIBUTE_VALUE_REAL)
		v1.POST("/v1/a_attribute_value_reals", GetController().PostA_ATTRIBUTE_VALUE_REAL)
		v1.PATCH("/v1/a_attribute_value_reals/:id", GetController().UpdateA_ATTRIBUTE_VALUE_REAL)
		v1.PUT("/v1/a_attribute_value_reals/:id", GetController().UpdateA_ATTRIBUTE_VALUE_REAL)
		v1.DELETE("/v1/a_attribute_value_reals/:id", GetController().DeleteA_ATTRIBUTE_VALUE_REAL)

		v1.GET("/v1/a_attribute_value_strings", GetController().GetA_ATTRIBUTE_VALUE_STRINGs)
		v1.GET("/v1/a_attribute_value_strings/:id", GetController().GetA_ATTRIBUTE_VALUE_STRING)
		v1.POST("/v1/a_attribute_value_strings", GetController().PostA_ATTRIBUTE_VALUE_STRING)
		v1.PATCH("/v1/a_attribute_value_strings/:id", GetController().UpdateA_ATTRIBUTE_VALUE_STRING)
		v1.PUT("/v1/a_attribute_value_strings/:id", GetController().UpdateA_ATTRIBUTE_VALUE_STRING)
		v1.DELETE("/v1/a_attribute_value_strings/:id", GetController().DeleteA_ATTRIBUTE_VALUE_STRING)

		v1.GET("/v1/a_attribute_value_xhtmls", GetController().GetA_ATTRIBUTE_VALUE_XHTMLs)
		v1.GET("/v1/a_attribute_value_xhtmls/:id", GetController().GetA_ATTRIBUTE_VALUE_XHTML)
		v1.POST("/v1/a_attribute_value_xhtmls", GetController().PostA_ATTRIBUTE_VALUE_XHTML)
		v1.PATCH("/v1/a_attribute_value_xhtmls/:id", GetController().UpdateA_ATTRIBUTE_VALUE_XHTML)
		v1.PUT("/v1/a_attribute_value_xhtmls/:id", GetController().UpdateA_ATTRIBUTE_VALUE_XHTML)
		v1.DELETE("/v1/a_attribute_value_xhtmls/:id", GetController().DeleteA_ATTRIBUTE_VALUE_XHTML)

		v1.GET("/v1/a_attribute_value_xhtml_1s", GetController().GetA_ATTRIBUTE_VALUE_XHTML_1s)
		v1.GET("/v1/a_attribute_value_xhtml_1s/:id", GetController().GetA_ATTRIBUTE_VALUE_XHTML_1)
		v1.POST("/v1/a_attribute_value_xhtml_1s", GetController().PostA_ATTRIBUTE_VALUE_XHTML_1)
		v1.PATCH("/v1/a_attribute_value_xhtml_1s/:id", GetController().UpdateA_ATTRIBUTE_VALUE_XHTML_1)
		v1.PUT("/v1/a_attribute_value_xhtml_1s/:id", GetController().UpdateA_ATTRIBUTE_VALUE_XHTML_1)
		v1.DELETE("/v1/a_attribute_value_xhtml_1s/:id", GetController().DeleteA_ATTRIBUTE_VALUE_XHTML_1)

		v1.GET("/v1/a_childrens", GetController().GetA_CHILDRENs)
		v1.GET("/v1/a_childrens/:id", GetController().GetA_CHILDREN)
		v1.POST("/v1/a_childrens", GetController().PostA_CHILDREN)
		v1.PATCH("/v1/a_childrens/:id", GetController().UpdateA_CHILDREN)
		v1.PUT("/v1/a_childrens/:id", GetController().UpdateA_CHILDREN)
		v1.DELETE("/v1/a_childrens/:id", GetController().DeleteA_CHILDREN)

		v1.GET("/v1/a_core_contents", GetController().GetA_CORE_CONTENTs)
		v1.GET("/v1/a_core_contents/:id", GetController().GetA_CORE_CONTENT)
		v1.POST("/v1/a_core_contents", GetController().PostA_CORE_CONTENT)
		v1.PATCH("/v1/a_core_contents/:id", GetController().UpdateA_CORE_CONTENT)
		v1.PUT("/v1/a_core_contents/:id", GetController().UpdateA_CORE_CONTENT)
		v1.DELETE("/v1/a_core_contents/:id", GetController().DeleteA_CORE_CONTENT)

		v1.GET("/v1/a_datatypess", GetController().GetA_DATATYPESs)
		v1.GET("/v1/a_datatypess/:id", GetController().GetA_DATATYPES)
		v1.POST("/v1/a_datatypess", GetController().PostA_DATATYPES)
		v1.PATCH("/v1/a_datatypess/:id", GetController().UpdateA_DATATYPES)
		v1.PUT("/v1/a_datatypess/:id", GetController().UpdateA_DATATYPES)
		v1.DELETE("/v1/a_datatypess/:id", GetController().DeleteA_DATATYPES)

		v1.GET("/v1/a_datatype_definition_boolean_refs", GetController().GetA_DATATYPE_DEFINITION_BOOLEAN_REFs)
		v1.GET("/v1/a_datatype_definition_boolean_refs/:id", GetController().GetA_DATATYPE_DEFINITION_BOOLEAN_REF)
		v1.POST("/v1/a_datatype_definition_boolean_refs", GetController().PostA_DATATYPE_DEFINITION_BOOLEAN_REF)
		v1.PATCH("/v1/a_datatype_definition_boolean_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_BOOLEAN_REF)
		v1.PUT("/v1/a_datatype_definition_boolean_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_BOOLEAN_REF)
		v1.DELETE("/v1/a_datatype_definition_boolean_refs/:id", GetController().DeleteA_DATATYPE_DEFINITION_BOOLEAN_REF)

		v1.GET("/v1/a_datatype_definition_date_refs", GetController().GetA_DATATYPE_DEFINITION_DATE_REFs)
		v1.GET("/v1/a_datatype_definition_date_refs/:id", GetController().GetA_DATATYPE_DEFINITION_DATE_REF)
		v1.POST("/v1/a_datatype_definition_date_refs", GetController().PostA_DATATYPE_DEFINITION_DATE_REF)
		v1.PATCH("/v1/a_datatype_definition_date_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_DATE_REF)
		v1.PUT("/v1/a_datatype_definition_date_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_DATE_REF)
		v1.DELETE("/v1/a_datatype_definition_date_refs/:id", GetController().DeleteA_DATATYPE_DEFINITION_DATE_REF)

		v1.GET("/v1/a_datatype_definition_enumeration_refs", GetController().GetA_DATATYPE_DEFINITION_ENUMERATION_REFs)
		v1.GET("/v1/a_datatype_definition_enumeration_refs/:id", GetController().GetA_DATATYPE_DEFINITION_ENUMERATION_REF)
		v1.POST("/v1/a_datatype_definition_enumeration_refs", GetController().PostA_DATATYPE_DEFINITION_ENUMERATION_REF)
		v1.PATCH("/v1/a_datatype_definition_enumeration_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_ENUMERATION_REF)
		v1.PUT("/v1/a_datatype_definition_enumeration_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_ENUMERATION_REF)
		v1.DELETE("/v1/a_datatype_definition_enumeration_refs/:id", GetController().DeleteA_DATATYPE_DEFINITION_ENUMERATION_REF)

		v1.GET("/v1/a_datatype_definition_integer_refs", GetController().GetA_DATATYPE_DEFINITION_INTEGER_REFs)
		v1.GET("/v1/a_datatype_definition_integer_refs/:id", GetController().GetA_DATATYPE_DEFINITION_INTEGER_REF)
		v1.POST("/v1/a_datatype_definition_integer_refs", GetController().PostA_DATATYPE_DEFINITION_INTEGER_REF)
		v1.PATCH("/v1/a_datatype_definition_integer_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_INTEGER_REF)
		v1.PUT("/v1/a_datatype_definition_integer_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_INTEGER_REF)
		v1.DELETE("/v1/a_datatype_definition_integer_refs/:id", GetController().DeleteA_DATATYPE_DEFINITION_INTEGER_REF)

		v1.GET("/v1/a_datatype_definition_real_refs", GetController().GetA_DATATYPE_DEFINITION_REAL_REFs)
		v1.GET("/v1/a_datatype_definition_real_refs/:id", GetController().GetA_DATATYPE_DEFINITION_REAL_REF)
		v1.POST("/v1/a_datatype_definition_real_refs", GetController().PostA_DATATYPE_DEFINITION_REAL_REF)
		v1.PATCH("/v1/a_datatype_definition_real_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_REAL_REF)
		v1.PUT("/v1/a_datatype_definition_real_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_REAL_REF)
		v1.DELETE("/v1/a_datatype_definition_real_refs/:id", GetController().DeleteA_DATATYPE_DEFINITION_REAL_REF)

		v1.GET("/v1/a_datatype_definition_string_refs", GetController().GetA_DATATYPE_DEFINITION_STRING_REFs)
		v1.GET("/v1/a_datatype_definition_string_refs/:id", GetController().GetA_DATATYPE_DEFINITION_STRING_REF)
		v1.POST("/v1/a_datatype_definition_string_refs", GetController().PostA_DATATYPE_DEFINITION_STRING_REF)
		v1.PATCH("/v1/a_datatype_definition_string_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_STRING_REF)
		v1.PUT("/v1/a_datatype_definition_string_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_STRING_REF)
		v1.DELETE("/v1/a_datatype_definition_string_refs/:id", GetController().DeleteA_DATATYPE_DEFINITION_STRING_REF)

		v1.GET("/v1/a_datatype_definition_xhtml_refs", GetController().GetA_DATATYPE_DEFINITION_XHTML_REFs)
		v1.GET("/v1/a_datatype_definition_xhtml_refs/:id", GetController().GetA_DATATYPE_DEFINITION_XHTML_REF)
		v1.POST("/v1/a_datatype_definition_xhtml_refs", GetController().PostA_DATATYPE_DEFINITION_XHTML_REF)
		v1.PATCH("/v1/a_datatype_definition_xhtml_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_XHTML_REF)
		v1.PUT("/v1/a_datatype_definition_xhtml_refs/:id", GetController().UpdateA_DATATYPE_DEFINITION_XHTML_REF)
		v1.DELETE("/v1/a_datatype_definition_xhtml_refs/:id", GetController().DeleteA_DATATYPE_DEFINITION_XHTML_REF)

		v1.GET("/v1/a_editable_attss", GetController().GetA_EDITABLE_ATTSs)
		v1.GET("/v1/a_editable_attss/:id", GetController().GetA_EDITABLE_ATTS)
		v1.POST("/v1/a_editable_attss", GetController().PostA_EDITABLE_ATTS)
		v1.PATCH("/v1/a_editable_attss/:id", GetController().UpdateA_EDITABLE_ATTS)
		v1.PUT("/v1/a_editable_attss/:id", GetController().UpdateA_EDITABLE_ATTS)
		v1.DELETE("/v1/a_editable_attss/:id", GetController().DeleteA_EDITABLE_ATTS)

		v1.GET("/v1/a_enum_value_refs", GetController().GetA_ENUM_VALUE_REFs)
		v1.GET("/v1/a_enum_value_refs/:id", GetController().GetA_ENUM_VALUE_REF)
		v1.POST("/v1/a_enum_value_refs", GetController().PostA_ENUM_VALUE_REF)
		v1.PATCH("/v1/a_enum_value_refs/:id", GetController().UpdateA_ENUM_VALUE_REF)
		v1.PUT("/v1/a_enum_value_refs/:id", GetController().UpdateA_ENUM_VALUE_REF)
		v1.DELETE("/v1/a_enum_value_refs/:id", GetController().DeleteA_ENUM_VALUE_REF)

		v1.GET("/v1/a_objects", GetController().GetA_OBJECTs)
		v1.GET("/v1/a_objects/:id", GetController().GetA_OBJECT)
		v1.POST("/v1/a_objects", GetController().PostA_OBJECT)
		v1.PATCH("/v1/a_objects/:id", GetController().UpdateA_OBJECT)
		v1.PUT("/v1/a_objects/:id", GetController().UpdateA_OBJECT)
		v1.DELETE("/v1/a_objects/:id", GetController().DeleteA_OBJECT)

		v1.GET("/v1/a_propertiess", GetController().GetA_PROPERTIESs)
		v1.GET("/v1/a_propertiess/:id", GetController().GetA_PROPERTIES)
		v1.POST("/v1/a_propertiess", GetController().PostA_PROPERTIES)
		v1.PATCH("/v1/a_propertiess/:id", GetController().UpdateA_PROPERTIES)
		v1.PUT("/v1/a_propertiess/:id", GetController().UpdateA_PROPERTIES)
		v1.DELETE("/v1/a_propertiess/:id", GetController().DeleteA_PROPERTIES)

		v1.GET("/v1/a_relation_group_type_refs", GetController().GetA_RELATION_GROUP_TYPE_REFs)
		v1.GET("/v1/a_relation_group_type_refs/:id", GetController().GetA_RELATION_GROUP_TYPE_REF)
		v1.POST("/v1/a_relation_group_type_refs", GetController().PostA_RELATION_GROUP_TYPE_REF)
		v1.PATCH("/v1/a_relation_group_type_refs/:id", GetController().UpdateA_RELATION_GROUP_TYPE_REF)
		v1.PUT("/v1/a_relation_group_type_refs/:id", GetController().UpdateA_RELATION_GROUP_TYPE_REF)
		v1.DELETE("/v1/a_relation_group_type_refs/:id", GetController().DeleteA_RELATION_GROUP_TYPE_REF)

		v1.GET("/v1/a_source_1s", GetController().GetA_SOURCE_1s)
		v1.GET("/v1/a_source_1s/:id", GetController().GetA_SOURCE_1)
		v1.POST("/v1/a_source_1s", GetController().PostA_SOURCE_1)
		v1.PATCH("/v1/a_source_1s/:id", GetController().UpdateA_SOURCE_1)
		v1.PUT("/v1/a_source_1s/:id", GetController().UpdateA_SOURCE_1)
		v1.DELETE("/v1/a_source_1s/:id", GetController().DeleteA_SOURCE_1)

		v1.GET("/v1/a_source_specification_1s", GetController().GetA_SOURCE_SPECIFICATION_1s)
		v1.GET("/v1/a_source_specification_1s/:id", GetController().GetA_SOURCE_SPECIFICATION_1)
		v1.POST("/v1/a_source_specification_1s", GetController().PostA_SOURCE_SPECIFICATION_1)
		v1.PATCH("/v1/a_source_specification_1s/:id", GetController().UpdateA_SOURCE_SPECIFICATION_1)
		v1.PUT("/v1/a_source_specification_1s/:id", GetController().UpdateA_SOURCE_SPECIFICATION_1)
		v1.DELETE("/v1/a_source_specification_1s/:id", GetController().DeleteA_SOURCE_SPECIFICATION_1)

		v1.GET("/v1/a_specificationss", GetController().GetA_SPECIFICATIONSs)
		v1.GET("/v1/a_specificationss/:id", GetController().GetA_SPECIFICATIONS)
		v1.POST("/v1/a_specificationss", GetController().PostA_SPECIFICATIONS)
		v1.PATCH("/v1/a_specificationss/:id", GetController().UpdateA_SPECIFICATIONS)
		v1.PUT("/v1/a_specificationss/:id", GetController().UpdateA_SPECIFICATIONS)
		v1.DELETE("/v1/a_specificationss/:id", GetController().DeleteA_SPECIFICATIONS)

		v1.GET("/v1/a_specification_type_refs", GetController().GetA_SPECIFICATION_TYPE_REFs)
		v1.GET("/v1/a_specification_type_refs/:id", GetController().GetA_SPECIFICATION_TYPE_REF)
		v1.POST("/v1/a_specification_type_refs", GetController().PostA_SPECIFICATION_TYPE_REF)
		v1.PATCH("/v1/a_specification_type_refs/:id", GetController().UpdateA_SPECIFICATION_TYPE_REF)
		v1.PUT("/v1/a_specification_type_refs/:id", GetController().UpdateA_SPECIFICATION_TYPE_REF)
		v1.DELETE("/v1/a_specification_type_refs/:id", GetController().DeleteA_SPECIFICATION_TYPE_REF)

		v1.GET("/v1/a_specified_valuess", GetController().GetA_SPECIFIED_VALUESs)
		v1.GET("/v1/a_specified_valuess/:id", GetController().GetA_SPECIFIED_VALUES)
		v1.POST("/v1/a_specified_valuess", GetController().PostA_SPECIFIED_VALUES)
		v1.PATCH("/v1/a_specified_valuess/:id", GetController().UpdateA_SPECIFIED_VALUES)
		v1.PUT("/v1/a_specified_valuess/:id", GetController().UpdateA_SPECIFIED_VALUES)
		v1.DELETE("/v1/a_specified_valuess/:id", GetController().DeleteA_SPECIFIED_VALUES)

		v1.GET("/v1/a_spec_attributess", GetController().GetA_SPEC_ATTRIBUTESs)
		v1.GET("/v1/a_spec_attributess/:id", GetController().GetA_SPEC_ATTRIBUTES)
		v1.POST("/v1/a_spec_attributess", GetController().PostA_SPEC_ATTRIBUTES)
		v1.PATCH("/v1/a_spec_attributess/:id", GetController().UpdateA_SPEC_ATTRIBUTES)
		v1.PUT("/v1/a_spec_attributess/:id", GetController().UpdateA_SPEC_ATTRIBUTES)
		v1.DELETE("/v1/a_spec_attributess/:id", GetController().DeleteA_SPEC_ATTRIBUTES)

		v1.GET("/v1/a_spec_objectss", GetController().GetA_SPEC_OBJECTSs)
		v1.GET("/v1/a_spec_objectss/:id", GetController().GetA_SPEC_OBJECTS)
		v1.POST("/v1/a_spec_objectss", GetController().PostA_SPEC_OBJECTS)
		v1.PATCH("/v1/a_spec_objectss/:id", GetController().UpdateA_SPEC_OBJECTS)
		v1.PUT("/v1/a_spec_objectss/:id", GetController().UpdateA_SPEC_OBJECTS)
		v1.DELETE("/v1/a_spec_objectss/:id", GetController().DeleteA_SPEC_OBJECTS)

		v1.GET("/v1/a_spec_object_type_refs", GetController().GetA_SPEC_OBJECT_TYPE_REFs)
		v1.GET("/v1/a_spec_object_type_refs/:id", GetController().GetA_SPEC_OBJECT_TYPE_REF)
		v1.POST("/v1/a_spec_object_type_refs", GetController().PostA_SPEC_OBJECT_TYPE_REF)
		v1.PATCH("/v1/a_spec_object_type_refs/:id", GetController().UpdateA_SPEC_OBJECT_TYPE_REF)
		v1.PUT("/v1/a_spec_object_type_refs/:id", GetController().UpdateA_SPEC_OBJECT_TYPE_REF)
		v1.DELETE("/v1/a_spec_object_type_refs/:id", GetController().DeleteA_SPEC_OBJECT_TYPE_REF)

		v1.GET("/v1/a_spec_relationss", GetController().GetA_SPEC_RELATIONSs)
		v1.GET("/v1/a_spec_relationss/:id", GetController().GetA_SPEC_RELATIONS)
		v1.POST("/v1/a_spec_relationss", GetController().PostA_SPEC_RELATIONS)
		v1.PATCH("/v1/a_spec_relationss/:id", GetController().UpdateA_SPEC_RELATIONS)
		v1.PUT("/v1/a_spec_relationss/:id", GetController().UpdateA_SPEC_RELATIONS)
		v1.DELETE("/v1/a_spec_relationss/:id", GetController().DeleteA_SPEC_RELATIONS)

		v1.GET("/v1/a_spec_relation_groupss", GetController().GetA_SPEC_RELATION_GROUPSs)
		v1.GET("/v1/a_spec_relation_groupss/:id", GetController().GetA_SPEC_RELATION_GROUPS)
		v1.POST("/v1/a_spec_relation_groupss", GetController().PostA_SPEC_RELATION_GROUPS)
		v1.PATCH("/v1/a_spec_relation_groupss/:id", GetController().UpdateA_SPEC_RELATION_GROUPS)
		v1.PUT("/v1/a_spec_relation_groupss/:id", GetController().UpdateA_SPEC_RELATION_GROUPS)
		v1.DELETE("/v1/a_spec_relation_groupss/:id", GetController().DeleteA_SPEC_RELATION_GROUPS)

		v1.GET("/v1/a_spec_relation_refs", GetController().GetA_SPEC_RELATION_REFs)
		v1.GET("/v1/a_spec_relation_refs/:id", GetController().GetA_SPEC_RELATION_REF)
		v1.POST("/v1/a_spec_relation_refs", GetController().PostA_SPEC_RELATION_REF)
		v1.PATCH("/v1/a_spec_relation_refs/:id", GetController().UpdateA_SPEC_RELATION_REF)
		v1.PUT("/v1/a_spec_relation_refs/:id", GetController().UpdateA_SPEC_RELATION_REF)
		v1.DELETE("/v1/a_spec_relation_refs/:id", GetController().DeleteA_SPEC_RELATION_REF)

		v1.GET("/v1/a_spec_relation_type_refs", GetController().GetA_SPEC_RELATION_TYPE_REFs)
		v1.GET("/v1/a_spec_relation_type_refs/:id", GetController().GetA_SPEC_RELATION_TYPE_REF)
		v1.POST("/v1/a_spec_relation_type_refs", GetController().PostA_SPEC_RELATION_TYPE_REF)
		v1.PATCH("/v1/a_spec_relation_type_refs/:id", GetController().UpdateA_SPEC_RELATION_TYPE_REF)
		v1.PUT("/v1/a_spec_relation_type_refs/:id", GetController().UpdateA_SPEC_RELATION_TYPE_REF)
		v1.DELETE("/v1/a_spec_relation_type_refs/:id", GetController().DeleteA_SPEC_RELATION_TYPE_REF)

		v1.GET("/v1/a_spec_typess", GetController().GetA_SPEC_TYPESs)
		v1.GET("/v1/a_spec_typess/:id", GetController().GetA_SPEC_TYPES)
		v1.POST("/v1/a_spec_typess", GetController().PostA_SPEC_TYPES)
		v1.PATCH("/v1/a_spec_typess/:id", GetController().UpdateA_SPEC_TYPES)
		v1.PUT("/v1/a_spec_typess/:id", GetController().UpdateA_SPEC_TYPES)
		v1.DELETE("/v1/a_spec_typess/:id", GetController().DeleteA_SPEC_TYPES)

		v1.GET("/v1/a_the_headers", GetController().GetA_THE_HEADERs)
		v1.GET("/v1/a_the_headers/:id", GetController().GetA_THE_HEADER)
		v1.POST("/v1/a_the_headers", GetController().PostA_THE_HEADER)
		v1.PATCH("/v1/a_the_headers/:id", GetController().UpdateA_THE_HEADER)
		v1.PUT("/v1/a_the_headers/:id", GetController().UpdateA_THE_HEADER)
		v1.DELETE("/v1/a_the_headers/:id", GetController().DeleteA_THE_HEADER)

		v1.GET("/v1/a_tool_extensionss", GetController().GetA_TOOL_EXTENSIONSs)
		v1.GET("/v1/a_tool_extensionss/:id", GetController().GetA_TOOL_EXTENSIONS)
		v1.POST("/v1/a_tool_extensionss", GetController().PostA_TOOL_EXTENSIONS)
		v1.PATCH("/v1/a_tool_extensionss/:id", GetController().UpdateA_TOOL_EXTENSIONS)
		v1.PUT("/v1/a_tool_extensionss/:id", GetController().UpdateA_TOOL_EXTENSIONS)
		v1.DELETE("/v1/a_tool_extensionss/:id", GetController().DeleteA_TOOL_EXTENSIONS)

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
			if origin == "" {
				log.Printf("CheckOrigin: Rejected - Origin header is empty. Request from: %s", r.RemoteAddr)
				return false // Or handle as per your security policy
			}

			u, err := url.Parse(origin)
			if err != nil {
				log.Printf("CheckOrigin: Rejected - Invalid Origin URL '%s'. Error: %v", origin, err)
				return false // Invalid URL
			}

			portStr := u.Port()

			if portStr == "" {
				// If no port is specified, it might be using default HTTP/HTTPS ports.
				// For this specific request, we'll assume a port must be present.
				log.Printf("CheckOrigin: Rejected - No port specified in Origin URL '%s'", origin)
				return false
			}

			port, err := strconv.Atoi(portStr)
			if err != nil {
				log.Printf("CheckOrigin: Rejected - Port '%s' in Origin URL '%s' is not a valid number. Error: %v", portStr, origin, err)
				return false // Port is not a valid number
			}

			// Check if the port is 4200 OR in the range 8000-9000
			allowed := port == 4200 || (port >= 8000 && port <= 9000)
			if !allowed {
				log.Printf("CheckOrigin: Rejected - Port %d from Origin '%s' is not in the allowed list (4200 or 8000-9000)", port, origin)
				return false
			}

			log.Printf("CheckOrigin: Accepted - Origin '%s' with port %d", origin, port)
			return true
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
