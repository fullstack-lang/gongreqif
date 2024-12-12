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
var __Xhtml_h1_type__dummysDeclaration__ models.Xhtml_h1_type
var __Xhtml_h1_type_time__dummyDeclaration time.Duration

var mutexXhtml_h1_type sync.Mutex

// An Xhtml_h1_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_h1_type updateXhtml_h1_type deleteXhtml_h1_type
type Xhtml_h1_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_h1_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_h1_type updateXhtml_h1_type
type Xhtml_h1_typeInput struct {
	// The Xhtml_h1_type to submit or modify
	// in: body
	Xhtml_h1_type *orm.Xhtml_h1_typeAPI
}

// GetXhtml_h1_types
//
// swagger:route GET /xhtml_h1_types xhtml_h1_types getXhtml_h1_types
//
// # Get all xhtml_h1_types
//
// Responses:
// default: genericError
//
//	200: xhtml_h1_typeDBResponse
func (controller *Controller) GetXhtml_h1_types(c *gin.Context) {

	// source slice
	var xhtml_h1_typeDBs []orm.Xhtml_h1_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h1_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h1_type.GetDB()

	_, err := db.Find(&xhtml_h1_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_h1_typeAPIs := make([]orm.Xhtml_h1_typeAPI, 0)

	// for each xhtml_h1_type, update fields from the database nullable fields
	for idx := range xhtml_h1_typeDBs {
		xhtml_h1_typeDB := &xhtml_h1_typeDBs[idx]
		_ = xhtml_h1_typeDB
		var xhtml_h1_typeAPI orm.Xhtml_h1_typeAPI

		// insertion point for updating fields
		xhtml_h1_typeAPI.ID = xhtml_h1_typeDB.ID
		xhtml_h1_typeDB.CopyBasicFieldsToXhtml_h1_type_WOP(&xhtml_h1_typeAPI.Xhtml_h1_type_WOP)
		xhtml_h1_typeAPI.Xhtml_h1_typePointersEncoding = xhtml_h1_typeDB.Xhtml_h1_typePointersEncoding
		xhtml_h1_typeAPIs = append(xhtml_h1_typeAPIs, xhtml_h1_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_h1_typeAPIs)
}

// PostXhtml_h1_type
//
// swagger:route POST /xhtml_h1_types xhtml_h1_types postXhtml_h1_type
//
// Creates a xhtml_h1_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_h1_type(c *gin.Context) {

	mutexXhtml_h1_type.Lock()
	defer mutexXhtml_h1_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_h1_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h1_type.GetDB()

	// Validate input
	var input orm.Xhtml_h1_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_h1_type
	xhtml_h1_typeDB := orm.Xhtml_h1_typeDB{}
	xhtml_h1_typeDB.Xhtml_h1_typePointersEncoding = input.Xhtml_h1_typePointersEncoding
	xhtml_h1_typeDB.CopyBasicFieldsFromXhtml_h1_type_WOP(&input.Xhtml_h1_type_WOP)

	_, err = db.Create(&xhtml_h1_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_h1_type.CheckoutPhaseOneInstance(&xhtml_h1_typeDB)
	xhtml_h1_type := backRepo.BackRepoXhtml_h1_type.Map_Xhtml_h1_typeDBID_Xhtml_h1_typePtr[xhtml_h1_typeDB.ID]

	if xhtml_h1_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_h1_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_h1_typeDB)
}

// GetXhtml_h1_type
//
// swagger:route GET /xhtml_h1_types/{ID} xhtml_h1_types getXhtml_h1_type
//
// Gets the details for a xhtml_h1_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_h1_typeDBResponse
func (controller *Controller) GetXhtml_h1_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h1_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h1_type.GetDB()

	// Get xhtml_h1_typeDB in DB
	var xhtml_h1_typeDB orm.Xhtml_h1_typeDB
	if _, err := db.First(&xhtml_h1_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_h1_typeAPI orm.Xhtml_h1_typeAPI
	xhtml_h1_typeAPI.ID = xhtml_h1_typeDB.ID
	xhtml_h1_typeAPI.Xhtml_h1_typePointersEncoding = xhtml_h1_typeDB.Xhtml_h1_typePointersEncoding
	xhtml_h1_typeDB.CopyBasicFieldsToXhtml_h1_type_WOP(&xhtml_h1_typeAPI.Xhtml_h1_type_WOP)

	c.JSON(http.StatusOK, xhtml_h1_typeAPI)
}

// UpdateXhtml_h1_type
//
// swagger:route PATCH /xhtml_h1_types/{ID} xhtml_h1_types updateXhtml_h1_type
//
// # Update a xhtml_h1_type
//
// Responses:
// default: genericError
//
//	200: xhtml_h1_typeDBResponse
func (controller *Controller) UpdateXhtml_h1_type(c *gin.Context) {

	mutexXhtml_h1_type.Lock()
	defer mutexXhtml_h1_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_h1_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h1_type.GetDB()

	// Validate input
	var input orm.Xhtml_h1_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_h1_typeDB orm.Xhtml_h1_typeDB

	// fetch the xhtml_h1_type
	_, err := db.First(&xhtml_h1_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_h1_typeDB.CopyBasicFieldsFromXhtml_h1_type_WOP(&input.Xhtml_h1_type_WOP)
	xhtml_h1_typeDB.Xhtml_h1_typePointersEncoding = input.Xhtml_h1_typePointersEncoding

	db, _ = db.Model(&xhtml_h1_typeDB)
	_, err = db.Updates(&xhtml_h1_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h1_typeNew := new(models.Xhtml_h1_type)
	xhtml_h1_typeDB.CopyBasicFieldsToXhtml_h1_type(xhtml_h1_typeNew)

	// redeem pointers
	xhtml_h1_typeDB.DecodePointers(backRepo, xhtml_h1_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_h1_typeOld := backRepo.BackRepoXhtml_h1_type.Map_Xhtml_h1_typeDBID_Xhtml_h1_typePtr[xhtml_h1_typeDB.ID]
	if xhtml_h1_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_h1_typeOld, xhtml_h1_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_h1_typeDB
	c.JSON(http.StatusOK, xhtml_h1_typeDB)
}

// DeleteXhtml_h1_type
//
// swagger:route DELETE /xhtml_h1_types/{ID} xhtml_h1_types deleteXhtml_h1_type
//
// # Delete a xhtml_h1_type
//
// default: genericError
//
//	200: xhtml_h1_typeDBResponse
func (controller *Controller) DeleteXhtml_h1_type(c *gin.Context) {

	mutexXhtml_h1_type.Lock()
	defer mutexXhtml_h1_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_h1_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h1_type.GetDB()

	// Get model if exist
	var xhtml_h1_typeDB orm.Xhtml_h1_typeDB
	if _, err := db.First(&xhtml_h1_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_h1_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h1_typeDeleted := new(models.Xhtml_h1_type)
	xhtml_h1_typeDB.CopyBasicFieldsToXhtml_h1_type(xhtml_h1_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_h1_typeStaged := backRepo.BackRepoXhtml_h1_type.Map_Xhtml_h1_typeDBID_Xhtml_h1_typePtr[xhtml_h1_typeDB.ID]
	if xhtml_h1_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_h1_typeStaged, xhtml_h1_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
