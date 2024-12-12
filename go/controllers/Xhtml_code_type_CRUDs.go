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
var __Xhtml_code_type__dummysDeclaration__ models.Xhtml_code_type
var __Xhtml_code_type_time__dummyDeclaration time.Duration

var mutexXhtml_code_type sync.Mutex

// An Xhtml_code_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_code_type updateXhtml_code_type deleteXhtml_code_type
type Xhtml_code_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_code_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_code_type updateXhtml_code_type
type Xhtml_code_typeInput struct {
	// The Xhtml_code_type to submit or modify
	// in: body
	Xhtml_code_type *orm.Xhtml_code_typeAPI
}

// GetXhtml_code_types
//
// swagger:route GET /xhtml_code_types xhtml_code_types getXhtml_code_types
//
// # Get all xhtml_code_types
//
// Responses:
// default: genericError
//
//	200: xhtml_code_typeDBResponse
func (controller *Controller) GetXhtml_code_types(c *gin.Context) {

	// source slice
	var xhtml_code_typeDBs []orm.Xhtml_code_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_code_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_code_type.GetDB()

	_, err := db.Find(&xhtml_code_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_code_typeAPIs := make([]orm.Xhtml_code_typeAPI, 0)

	// for each xhtml_code_type, update fields from the database nullable fields
	for idx := range xhtml_code_typeDBs {
		xhtml_code_typeDB := &xhtml_code_typeDBs[idx]
		_ = xhtml_code_typeDB
		var xhtml_code_typeAPI orm.Xhtml_code_typeAPI

		// insertion point for updating fields
		xhtml_code_typeAPI.ID = xhtml_code_typeDB.ID
		xhtml_code_typeDB.CopyBasicFieldsToXhtml_code_type_WOP(&xhtml_code_typeAPI.Xhtml_code_type_WOP)
		xhtml_code_typeAPI.Xhtml_code_typePointersEncoding = xhtml_code_typeDB.Xhtml_code_typePointersEncoding
		xhtml_code_typeAPIs = append(xhtml_code_typeAPIs, xhtml_code_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_code_typeAPIs)
}

// PostXhtml_code_type
//
// swagger:route POST /xhtml_code_types xhtml_code_types postXhtml_code_type
//
// Creates a xhtml_code_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_code_type(c *gin.Context) {

	mutexXhtml_code_type.Lock()
	defer mutexXhtml_code_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_code_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_code_type.GetDB()

	// Validate input
	var input orm.Xhtml_code_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_code_type
	xhtml_code_typeDB := orm.Xhtml_code_typeDB{}
	xhtml_code_typeDB.Xhtml_code_typePointersEncoding = input.Xhtml_code_typePointersEncoding
	xhtml_code_typeDB.CopyBasicFieldsFromXhtml_code_type_WOP(&input.Xhtml_code_type_WOP)

	_, err = db.Create(&xhtml_code_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_code_type.CheckoutPhaseOneInstance(&xhtml_code_typeDB)
	xhtml_code_type := backRepo.BackRepoXhtml_code_type.Map_Xhtml_code_typeDBID_Xhtml_code_typePtr[xhtml_code_typeDB.ID]

	if xhtml_code_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_code_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_code_typeDB)
}

// GetXhtml_code_type
//
// swagger:route GET /xhtml_code_types/{ID} xhtml_code_types getXhtml_code_type
//
// Gets the details for a xhtml_code_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_code_typeDBResponse
func (controller *Controller) GetXhtml_code_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_code_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_code_type.GetDB()

	// Get xhtml_code_typeDB in DB
	var xhtml_code_typeDB orm.Xhtml_code_typeDB
	if _, err := db.First(&xhtml_code_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_code_typeAPI orm.Xhtml_code_typeAPI
	xhtml_code_typeAPI.ID = xhtml_code_typeDB.ID
	xhtml_code_typeAPI.Xhtml_code_typePointersEncoding = xhtml_code_typeDB.Xhtml_code_typePointersEncoding
	xhtml_code_typeDB.CopyBasicFieldsToXhtml_code_type_WOP(&xhtml_code_typeAPI.Xhtml_code_type_WOP)

	c.JSON(http.StatusOK, xhtml_code_typeAPI)
}

// UpdateXhtml_code_type
//
// swagger:route PATCH /xhtml_code_types/{ID} xhtml_code_types updateXhtml_code_type
//
// # Update a xhtml_code_type
//
// Responses:
// default: genericError
//
//	200: xhtml_code_typeDBResponse
func (controller *Controller) UpdateXhtml_code_type(c *gin.Context) {

	mutexXhtml_code_type.Lock()
	defer mutexXhtml_code_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_code_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_code_type.GetDB()

	// Validate input
	var input orm.Xhtml_code_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_code_typeDB orm.Xhtml_code_typeDB

	// fetch the xhtml_code_type
	_, err := db.First(&xhtml_code_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_code_typeDB.CopyBasicFieldsFromXhtml_code_type_WOP(&input.Xhtml_code_type_WOP)
	xhtml_code_typeDB.Xhtml_code_typePointersEncoding = input.Xhtml_code_typePointersEncoding

	db, _ = db.Model(&xhtml_code_typeDB)
	_, err = db.Updates(&xhtml_code_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_code_typeNew := new(models.Xhtml_code_type)
	xhtml_code_typeDB.CopyBasicFieldsToXhtml_code_type(xhtml_code_typeNew)

	// redeem pointers
	xhtml_code_typeDB.DecodePointers(backRepo, xhtml_code_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_code_typeOld := backRepo.BackRepoXhtml_code_type.Map_Xhtml_code_typeDBID_Xhtml_code_typePtr[xhtml_code_typeDB.ID]
	if xhtml_code_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_code_typeOld, xhtml_code_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_code_typeDB
	c.JSON(http.StatusOK, xhtml_code_typeDB)
}

// DeleteXhtml_code_type
//
// swagger:route DELETE /xhtml_code_types/{ID} xhtml_code_types deleteXhtml_code_type
//
// # Delete a xhtml_code_type
//
// default: genericError
//
//	200: xhtml_code_typeDBResponse
func (controller *Controller) DeleteXhtml_code_type(c *gin.Context) {

	mutexXhtml_code_type.Lock()
	defer mutexXhtml_code_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_code_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_code_type.GetDB()

	// Get model if exist
	var xhtml_code_typeDB orm.Xhtml_code_typeDB
	if _, err := db.First(&xhtml_code_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_code_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_code_typeDeleted := new(models.Xhtml_code_type)
	xhtml_code_typeDB.CopyBasicFieldsToXhtml_code_type(xhtml_code_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_code_typeStaged := backRepo.BackRepoXhtml_code_type.Map_Xhtml_code_typeDBID_Xhtml_code_typePtr[xhtml_code_typeDB.ID]
	if xhtml_code_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_code_typeStaged, xhtml_code_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
