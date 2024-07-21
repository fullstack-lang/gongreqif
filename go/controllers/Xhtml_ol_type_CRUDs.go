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
var __Xhtml_ol_type__dummysDeclaration__ models.Xhtml_ol_type
var __Xhtml_ol_type_time__dummyDeclaration time.Duration

var mutexXhtml_ol_type sync.Mutex

// An Xhtml_ol_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_ol_type updateXhtml_ol_type deleteXhtml_ol_type
type Xhtml_ol_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_ol_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_ol_type updateXhtml_ol_type
type Xhtml_ol_typeInput struct {
	// The Xhtml_ol_type to submit or modify
	// in: body
	Xhtml_ol_type *orm.Xhtml_ol_typeAPI
}

// GetXhtml_ol_types
//
// swagger:route GET /xhtml_ol_types xhtml_ol_types getXhtml_ol_types
//
// # Get all xhtml_ol_types
//
// Responses:
// default: genericError
//
//	200: xhtml_ol_typeDBResponse
func (controller *Controller) GetXhtml_ol_types(c *gin.Context) {

	// source slice
	var xhtml_ol_typeDBs []orm.Xhtml_ol_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_ol_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ol_type.GetDB()

	query := db.Find(&xhtml_ol_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_ol_typeAPIs := make([]orm.Xhtml_ol_typeAPI, 0)

	// for each xhtml_ol_type, update fields from the database nullable fields
	for idx := range xhtml_ol_typeDBs {
		xhtml_ol_typeDB := &xhtml_ol_typeDBs[idx]
		_ = xhtml_ol_typeDB
		var xhtml_ol_typeAPI orm.Xhtml_ol_typeAPI

		// insertion point for updating fields
		xhtml_ol_typeAPI.ID = xhtml_ol_typeDB.ID
		xhtml_ol_typeDB.CopyBasicFieldsToXhtml_ol_type_WOP(&xhtml_ol_typeAPI.Xhtml_ol_type_WOP)
		xhtml_ol_typeAPI.Xhtml_ol_typePointersEncoding = xhtml_ol_typeDB.Xhtml_ol_typePointersEncoding
		xhtml_ol_typeAPIs = append(xhtml_ol_typeAPIs, xhtml_ol_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_ol_typeAPIs)
}

// PostXhtml_ol_type
//
// swagger:route POST /xhtml_ol_types xhtml_ol_types postXhtml_ol_type
//
// Creates a xhtml_ol_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_ol_type(c *gin.Context) {

	mutexXhtml_ol_type.Lock()
	defer mutexXhtml_ol_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_ol_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ol_type.GetDB()

	// Validate input
	var input orm.Xhtml_ol_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_ol_type
	xhtml_ol_typeDB := orm.Xhtml_ol_typeDB{}
	xhtml_ol_typeDB.Xhtml_ol_typePointersEncoding = input.Xhtml_ol_typePointersEncoding
	xhtml_ol_typeDB.CopyBasicFieldsFromXhtml_ol_type_WOP(&input.Xhtml_ol_type_WOP)

	query := db.Create(&xhtml_ol_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_ol_type.CheckoutPhaseOneInstance(&xhtml_ol_typeDB)
	xhtml_ol_type := backRepo.BackRepoXhtml_ol_type.Map_Xhtml_ol_typeDBID_Xhtml_ol_typePtr[xhtml_ol_typeDB.ID]

	if xhtml_ol_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_ol_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_ol_typeDB)
}

// GetXhtml_ol_type
//
// swagger:route GET /xhtml_ol_types/{ID} xhtml_ol_types getXhtml_ol_type
//
// Gets the details for a xhtml_ol_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_ol_typeDBResponse
func (controller *Controller) GetXhtml_ol_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_ol_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ol_type.GetDB()

	// Get xhtml_ol_typeDB in DB
	var xhtml_ol_typeDB orm.Xhtml_ol_typeDB
	if err := db.First(&xhtml_ol_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_ol_typeAPI orm.Xhtml_ol_typeAPI
	xhtml_ol_typeAPI.ID = xhtml_ol_typeDB.ID
	xhtml_ol_typeAPI.Xhtml_ol_typePointersEncoding = xhtml_ol_typeDB.Xhtml_ol_typePointersEncoding
	xhtml_ol_typeDB.CopyBasicFieldsToXhtml_ol_type_WOP(&xhtml_ol_typeAPI.Xhtml_ol_type_WOP)

	c.JSON(http.StatusOK, xhtml_ol_typeAPI)
}

// UpdateXhtml_ol_type
//
// swagger:route PATCH /xhtml_ol_types/{ID} xhtml_ol_types updateXhtml_ol_type
//
// # Update a xhtml_ol_type
//
// Responses:
// default: genericError
//
//	200: xhtml_ol_typeDBResponse
func (controller *Controller) UpdateXhtml_ol_type(c *gin.Context) {

	mutexXhtml_ol_type.Lock()
	defer mutexXhtml_ol_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_ol_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ol_type.GetDB()

	// Validate input
	var input orm.Xhtml_ol_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_ol_typeDB orm.Xhtml_ol_typeDB

	// fetch the xhtml_ol_type
	query := db.First(&xhtml_ol_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_ol_typeDB.CopyBasicFieldsFromXhtml_ol_type_WOP(&input.Xhtml_ol_type_WOP)
	xhtml_ol_typeDB.Xhtml_ol_typePointersEncoding = input.Xhtml_ol_typePointersEncoding

	query = db.Model(&xhtml_ol_typeDB).Updates(xhtml_ol_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_ol_typeNew := new(models.Xhtml_ol_type)
	xhtml_ol_typeDB.CopyBasicFieldsToXhtml_ol_type(xhtml_ol_typeNew)

	// redeem pointers
	xhtml_ol_typeDB.DecodePointers(backRepo, xhtml_ol_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_ol_typeOld := backRepo.BackRepoXhtml_ol_type.Map_Xhtml_ol_typeDBID_Xhtml_ol_typePtr[xhtml_ol_typeDB.ID]
	if xhtml_ol_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_ol_typeOld, xhtml_ol_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_ol_typeDB
	c.JSON(http.StatusOK, xhtml_ol_typeDB)
}

// DeleteXhtml_ol_type
//
// swagger:route DELETE /xhtml_ol_types/{ID} xhtml_ol_types deleteXhtml_ol_type
//
// # Delete a xhtml_ol_type
//
// default: genericError
//
//	200: xhtml_ol_typeDBResponse
func (controller *Controller) DeleteXhtml_ol_type(c *gin.Context) {

	mutexXhtml_ol_type.Lock()
	defer mutexXhtml_ol_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_ol_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ol_type.GetDB()

	// Get model if exist
	var xhtml_ol_typeDB orm.Xhtml_ol_typeDB
	if err := db.First(&xhtml_ol_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_ol_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_ol_typeDeleted := new(models.Xhtml_ol_type)
	xhtml_ol_typeDB.CopyBasicFieldsToXhtml_ol_type(xhtml_ol_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_ol_typeStaged := backRepo.BackRepoXhtml_ol_type.Map_Xhtml_ol_typeDBID_Xhtml_ol_typePtr[xhtml_ol_typeDB.ID]
	if xhtml_ol_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_ol_typeStaged, xhtml_ol_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
