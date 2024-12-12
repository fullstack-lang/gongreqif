// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongreqif/go/models"
	"github.com/fullstack-lang/gongreqif/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Xhtml_InlPres_type__dummysDeclaration__ models.Xhtml_InlPres_type
var __Xhtml_InlPres_type_time__dummyDeclaration time.Duration

var mutexXhtml_InlPres_type sync.Mutex

// An Xhtml_InlPres_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_InlPres_type updateXhtml_InlPres_type deleteXhtml_InlPres_type
type Xhtml_InlPres_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_InlPres_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_InlPres_type updateXhtml_InlPres_type
type Xhtml_InlPres_typeInput struct {
	// The Xhtml_InlPres_type to submit or modify
	// in: body
	Xhtml_InlPres_type *orm.Xhtml_InlPres_typeAPI
}

// GetXhtml_InlPres_types
//
// swagger:route GET /xhtml_inlpres_types xhtml_inlpres_types getXhtml_InlPres_types
//
// # Get all xhtml_inlpres_types
//
// Responses:
// default: genericError
//
//	200: xhtml_inlpres_typeDBResponse
func (controller *Controller) GetXhtml_InlPres_types(c *gin.Context) {

	// source slice
	var xhtml_inlpres_typeDBs []orm.Xhtml_InlPres_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_InlPres_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_InlPres_type.GetDB()

	_, err := db.Find(&xhtml_inlpres_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_inlpres_typeAPIs := make([]orm.Xhtml_InlPres_typeAPI, 0)

	// for each xhtml_inlpres_type, update fields from the database nullable fields
	for idx := range xhtml_inlpres_typeDBs {
		xhtml_inlpres_typeDB := &xhtml_inlpres_typeDBs[idx]
		_ = xhtml_inlpres_typeDB
		var xhtml_inlpres_typeAPI orm.Xhtml_InlPres_typeAPI

		// insertion point for updating fields
		xhtml_inlpres_typeAPI.ID = xhtml_inlpres_typeDB.ID
		xhtml_inlpres_typeDB.CopyBasicFieldsToXhtml_InlPres_type_WOP(&xhtml_inlpres_typeAPI.Xhtml_InlPres_type_WOP)
		xhtml_inlpres_typeAPI.Xhtml_InlPres_typePointersEncoding = xhtml_inlpres_typeDB.Xhtml_InlPres_typePointersEncoding
		xhtml_inlpres_typeAPIs = append(xhtml_inlpres_typeAPIs, xhtml_inlpres_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_inlpres_typeAPIs)
}

// PostXhtml_InlPres_type
//
// swagger:route POST /xhtml_inlpres_types xhtml_inlpres_types postXhtml_InlPres_type
//
// Creates a xhtml_inlpres_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_InlPres_type(c *gin.Context) {

	mutexXhtml_InlPres_type.Lock()
	defer mutexXhtml_InlPres_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_InlPres_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_InlPres_type.GetDB()

	// Validate input
	var input orm.Xhtml_InlPres_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_inlpres_type
	xhtml_inlpres_typeDB := orm.Xhtml_InlPres_typeDB{}
	xhtml_inlpres_typeDB.Xhtml_InlPres_typePointersEncoding = input.Xhtml_InlPres_typePointersEncoding
	xhtml_inlpres_typeDB.CopyBasicFieldsFromXhtml_InlPres_type_WOP(&input.Xhtml_InlPres_type_WOP)

	_, err = db.Create(&xhtml_inlpres_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_InlPres_type.CheckoutPhaseOneInstance(&xhtml_inlpres_typeDB)
	xhtml_inlpres_type := backRepo.BackRepoXhtml_InlPres_type.Map_Xhtml_InlPres_typeDBID_Xhtml_InlPres_typePtr[xhtml_inlpres_typeDB.ID]

	if xhtml_inlpres_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_inlpres_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_inlpres_typeDB)
}

// GetXhtml_InlPres_type
//
// swagger:route GET /xhtml_inlpres_types/{ID} xhtml_inlpres_types getXhtml_InlPres_type
//
// Gets the details for a xhtml_inlpres_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_inlpres_typeDBResponse
func (controller *Controller) GetXhtml_InlPres_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_InlPres_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_InlPres_type.GetDB()

	// Get xhtml_inlpres_typeDB in DB
	var xhtml_inlpres_typeDB orm.Xhtml_InlPres_typeDB
	if _, err := db.First(&xhtml_inlpres_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_inlpres_typeAPI orm.Xhtml_InlPres_typeAPI
	xhtml_inlpres_typeAPI.ID = xhtml_inlpres_typeDB.ID
	xhtml_inlpres_typeAPI.Xhtml_InlPres_typePointersEncoding = xhtml_inlpres_typeDB.Xhtml_InlPres_typePointersEncoding
	xhtml_inlpres_typeDB.CopyBasicFieldsToXhtml_InlPres_type_WOP(&xhtml_inlpres_typeAPI.Xhtml_InlPres_type_WOP)

	c.JSON(http.StatusOK, xhtml_inlpres_typeAPI)
}

// UpdateXhtml_InlPres_type
//
// swagger:route PATCH /xhtml_inlpres_types/{ID} xhtml_inlpres_types updateXhtml_InlPres_type
//
// # Update a xhtml_inlpres_type
//
// Responses:
// default: genericError
//
//	200: xhtml_inlpres_typeDBResponse
func (controller *Controller) UpdateXhtml_InlPres_type(c *gin.Context) {

	mutexXhtml_InlPres_type.Lock()
	defer mutexXhtml_InlPres_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_InlPres_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_InlPres_type.GetDB()

	// Validate input
	var input orm.Xhtml_InlPres_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_inlpres_typeDB orm.Xhtml_InlPres_typeDB

	// fetch the xhtml_inlpres_type
	_, err := db.First(&xhtml_inlpres_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_inlpres_typeDB.CopyBasicFieldsFromXhtml_InlPres_type_WOP(&input.Xhtml_InlPres_type_WOP)
	xhtml_inlpres_typeDB.Xhtml_InlPres_typePointersEncoding = input.Xhtml_InlPres_typePointersEncoding

	db, _ = db.Model(&xhtml_inlpres_typeDB)
	_, err = db.Updates(&xhtml_inlpres_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_inlpres_typeNew := new(models.Xhtml_InlPres_type)
	xhtml_inlpres_typeDB.CopyBasicFieldsToXhtml_InlPres_type(xhtml_inlpres_typeNew)

	// redeem pointers
	xhtml_inlpres_typeDB.DecodePointers(backRepo, xhtml_inlpres_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_inlpres_typeOld := backRepo.BackRepoXhtml_InlPres_type.Map_Xhtml_InlPres_typeDBID_Xhtml_InlPres_typePtr[xhtml_inlpres_typeDB.ID]
	if xhtml_inlpres_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_inlpres_typeOld, xhtml_inlpres_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_inlpres_typeDB
	c.JSON(http.StatusOK, xhtml_inlpres_typeDB)
}

// DeleteXhtml_InlPres_type
//
// swagger:route DELETE /xhtml_inlpres_types/{ID} xhtml_inlpres_types deleteXhtml_InlPres_type
//
// # Delete a xhtml_inlpres_type
//
// default: genericError
//
//	200: xhtml_inlpres_typeDBResponse
func (controller *Controller) DeleteXhtml_InlPres_type(c *gin.Context) {

	mutexXhtml_InlPres_type.Lock()
	defer mutexXhtml_InlPres_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_InlPres_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_InlPres_type.GetDB()

	// Get model if exist
	var xhtml_inlpres_typeDB orm.Xhtml_InlPres_typeDB
	if _, err := db.First(&xhtml_inlpres_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_inlpres_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_inlpres_typeDeleted := new(models.Xhtml_InlPres_type)
	xhtml_inlpres_typeDB.CopyBasicFieldsToXhtml_InlPres_type(xhtml_inlpres_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_inlpres_typeStaged := backRepo.BackRepoXhtml_InlPres_type.Map_Xhtml_InlPres_typeDBID_Xhtml_InlPres_typePtr[xhtml_inlpres_typeDB.ID]
	if xhtml_inlpres_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_inlpres_typeStaged, xhtml_inlpres_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
