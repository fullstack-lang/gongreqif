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
var __Xhtml_edit_type__dummysDeclaration__ models.Xhtml_edit_type
var __Xhtml_edit_type_time__dummyDeclaration time.Duration

var mutexXhtml_edit_type sync.Mutex

// An Xhtml_edit_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_edit_type updateXhtml_edit_type deleteXhtml_edit_type
type Xhtml_edit_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_edit_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_edit_type updateXhtml_edit_type
type Xhtml_edit_typeInput struct {
	// The Xhtml_edit_type to submit or modify
	// in: body
	Xhtml_edit_type *orm.Xhtml_edit_typeAPI
}

// GetXhtml_edit_types
//
// swagger:route GET /xhtml_edit_types xhtml_edit_types getXhtml_edit_types
//
// # Get all xhtml_edit_types
//
// Responses:
// default: genericError
//
//	200: xhtml_edit_typeDBResponse
func (controller *Controller) GetXhtml_edit_types(c *gin.Context) {

	// source slice
	var xhtml_edit_typeDBs []orm.Xhtml_edit_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_edit_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_edit_type.GetDB()

	query := db.Find(&xhtml_edit_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_edit_typeAPIs := make([]orm.Xhtml_edit_typeAPI, 0)

	// for each xhtml_edit_type, update fields from the database nullable fields
	for idx := range xhtml_edit_typeDBs {
		xhtml_edit_typeDB := &xhtml_edit_typeDBs[idx]
		_ = xhtml_edit_typeDB
		var xhtml_edit_typeAPI orm.Xhtml_edit_typeAPI

		// insertion point for updating fields
		xhtml_edit_typeAPI.ID = xhtml_edit_typeDB.ID
		xhtml_edit_typeDB.CopyBasicFieldsToXhtml_edit_type_WOP(&xhtml_edit_typeAPI.Xhtml_edit_type_WOP)
		xhtml_edit_typeAPI.Xhtml_edit_typePointersEncoding = xhtml_edit_typeDB.Xhtml_edit_typePointersEncoding
		xhtml_edit_typeAPIs = append(xhtml_edit_typeAPIs, xhtml_edit_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_edit_typeAPIs)
}

// PostXhtml_edit_type
//
// swagger:route POST /xhtml_edit_types xhtml_edit_types postXhtml_edit_type
//
// Creates a xhtml_edit_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_edit_type(c *gin.Context) {

	mutexXhtml_edit_type.Lock()
	defer mutexXhtml_edit_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_edit_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_edit_type.GetDB()

	// Validate input
	var input orm.Xhtml_edit_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_edit_type
	xhtml_edit_typeDB := orm.Xhtml_edit_typeDB{}
	xhtml_edit_typeDB.Xhtml_edit_typePointersEncoding = input.Xhtml_edit_typePointersEncoding
	xhtml_edit_typeDB.CopyBasicFieldsFromXhtml_edit_type_WOP(&input.Xhtml_edit_type_WOP)

	query := db.Create(&xhtml_edit_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_edit_type.CheckoutPhaseOneInstance(&xhtml_edit_typeDB)
	xhtml_edit_type := backRepo.BackRepoXhtml_edit_type.Map_Xhtml_edit_typeDBID_Xhtml_edit_typePtr[xhtml_edit_typeDB.ID]

	if xhtml_edit_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_edit_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_edit_typeDB)
}

// GetXhtml_edit_type
//
// swagger:route GET /xhtml_edit_types/{ID} xhtml_edit_types getXhtml_edit_type
//
// Gets the details for a xhtml_edit_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_edit_typeDBResponse
func (controller *Controller) GetXhtml_edit_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_edit_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_edit_type.GetDB()

	// Get xhtml_edit_typeDB in DB
	var xhtml_edit_typeDB orm.Xhtml_edit_typeDB
	if err := db.First(&xhtml_edit_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_edit_typeAPI orm.Xhtml_edit_typeAPI
	xhtml_edit_typeAPI.ID = xhtml_edit_typeDB.ID
	xhtml_edit_typeAPI.Xhtml_edit_typePointersEncoding = xhtml_edit_typeDB.Xhtml_edit_typePointersEncoding
	xhtml_edit_typeDB.CopyBasicFieldsToXhtml_edit_type_WOP(&xhtml_edit_typeAPI.Xhtml_edit_type_WOP)

	c.JSON(http.StatusOK, xhtml_edit_typeAPI)
}

// UpdateXhtml_edit_type
//
// swagger:route PATCH /xhtml_edit_types/{ID} xhtml_edit_types updateXhtml_edit_type
//
// # Update a xhtml_edit_type
//
// Responses:
// default: genericError
//
//	200: xhtml_edit_typeDBResponse
func (controller *Controller) UpdateXhtml_edit_type(c *gin.Context) {

	mutexXhtml_edit_type.Lock()
	defer mutexXhtml_edit_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_edit_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_edit_type.GetDB()

	// Validate input
	var input orm.Xhtml_edit_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_edit_typeDB orm.Xhtml_edit_typeDB

	// fetch the xhtml_edit_type
	query := db.First(&xhtml_edit_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_edit_typeDB.CopyBasicFieldsFromXhtml_edit_type_WOP(&input.Xhtml_edit_type_WOP)
	xhtml_edit_typeDB.Xhtml_edit_typePointersEncoding = input.Xhtml_edit_typePointersEncoding

	query = db.Model(&xhtml_edit_typeDB).Updates(xhtml_edit_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_edit_typeNew := new(models.Xhtml_edit_type)
	xhtml_edit_typeDB.CopyBasicFieldsToXhtml_edit_type(xhtml_edit_typeNew)

	// redeem pointers
	xhtml_edit_typeDB.DecodePointers(backRepo, xhtml_edit_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_edit_typeOld := backRepo.BackRepoXhtml_edit_type.Map_Xhtml_edit_typeDBID_Xhtml_edit_typePtr[xhtml_edit_typeDB.ID]
	if xhtml_edit_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_edit_typeOld, xhtml_edit_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_edit_typeDB
	c.JSON(http.StatusOK, xhtml_edit_typeDB)
}

// DeleteXhtml_edit_type
//
// swagger:route DELETE /xhtml_edit_types/{ID} xhtml_edit_types deleteXhtml_edit_type
//
// # Delete a xhtml_edit_type
//
// default: genericError
//
//	200: xhtml_edit_typeDBResponse
func (controller *Controller) DeleteXhtml_edit_type(c *gin.Context) {

	mutexXhtml_edit_type.Lock()
	defer mutexXhtml_edit_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_edit_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_edit_type.GetDB()

	// Get model if exist
	var xhtml_edit_typeDB orm.Xhtml_edit_typeDB
	if err := db.First(&xhtml_edit_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_edit_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_edit_typeDeleted := new(models.Xhtml_edit_type)
	xhtml_edit_typeDB.CopyBasicFieldsToXhtml_edit_type(xhtml_edit_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_edit_typeStaged := backRepo.BackRepoXhtml_edit_type.Map_Xhtml_edit_typeDBID_Xhtml_edit_typePtr[xhtml_edit_typeDB.ID]
	if xhtml_edit_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_edit_typeStaged, xhtml_edit_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
