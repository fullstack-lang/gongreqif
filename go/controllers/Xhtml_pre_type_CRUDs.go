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
var __Xhtml_pre_type__dummysDeclaration__ models.Xhtml_pre_type
var __Xhtml_pre_type_time__dummyDeclaration time.Duration

var mutexXhtml_pre_type sync.Mutex

// An Xhtml_pre_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_pre_type updateXhtml_pre_type deleteXhtml_pre_type
type Xhtml_pre_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_pre_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_pre_type updateXhtml_pre_type
type Xhtml_pre_typeInput struct {
	// The Xhtml_pre_type to submit or modify
	// in: body
	Xhtml_pre_type *orm.Xhtml_pre_typeAPI
}

// GetXhtml_pre_types
//
// swagger:route GET /xhtml_pre_types xhtml_pre_types getXhtml_pre_types
//
// # Get all xhtml_pre_types
//
// Responses:
// default: genericError
//
//	200: xhtml_pre_typeDBResponse
func (controller *Controller) GetXhtml_pre_types(c *gin.Context) {

	// source slice
	var xhtml_pre_typeDBs []orm.Xhtml_pre_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_pre_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_pre_type.GetDB()

	_, err := db.Find(&xhtml_pre_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_pre_typeAPIs := make([]orm.Xhtml_pre_typeAPI, 0)

	// for each xhtml_pre_type, update fields from the database nullable fields
	for idx := range xhtml_pre_typeDBs {
		xhtml_pre_typeDB := &xhtml_pre_typeDBs[idx]
		_ = xhtml_pre_typeDB
		var xhtml_pre_typeAPI orm.Xhtml_pre_typeAPI

		// insertion point for updating fields
		xhtml_pre_typeAPI.ID = xhtml_pre_typeDB.ID
		xhtml_pre_typeDB.CopyBasicFieldsToXhtml_pre_type_WOP(&xhtml_pre_typeAPI.Xhtml_pre_type_WOP)
		xhtml_pre_typeAPI.Xhtml_pre_typePointersEncoding = xhtml_pre_typeDB.Xhtml_pre_typePointersEncoding
		xhtml_pre_typeAPIs = append(xhtml_pre_typeAPIs, xhtml_pre_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_pre_typeAPIs)
}

// PostXhtml_pre_type
//
// swagger:route POST /xhtml_pre_types xhtml_pre_types postXhtml_pre_type
//
// Creates a xhtml_pre_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_pre_type(c *gin.Context) {

	mutexXhtml_pre_type.Lock()
	defer mutexXhtml_pre_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_pre_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_pre_type.GetDB()

	// Validate input
	var input orm.Xhtml_pre_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_pre_type
	xhtml_pre_typeDB := orm.Xhtml_pre_typeDB{}
	xhtml_pre_typeDB.Xhtml_pre_typePointersEncoding = input.Xhtml_pre_typePointersEncoding
	xhtml_pre_typeDB.CopyBasicFieldsFromXhtml_pre_type_WOP(&input.Xhtml_pre_type_WOP)

	_, err = db.Create(&xhtml_pre_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_pre_type.CheckoutPhaseOneInstance(&xhtml_pre_typeDB)
	xhtml_pre_type := backRepo.BackRepoXhtml_pre_type.Map_Xhtml_pre_typeDBID_Xhtml_pre_typePtr[xhtml_pre_typeDB.ID]

	if xhtml_pre_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_pre_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_pre_typeDB)
}

// GetXhtml_pre_type
//
// swagger:route GET /xhtml_pre_types/{ID} xhtml_pre_types getXhtml_pre_type
//
// Gets the details for a xhtml_pre_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_pre_typeDBResponse
func (controller *Controller) GetXhtml_pre_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_pre_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_pre_type.GetDB()

	// Get xhtml_pre_typeDB in DB
	var xhtml_pre_typeDB orm.Xhtml_pre_typeDB
	if _, err := db.First(&xhtml_pre_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_pre_typeAPI orm.Xhtml_pre_typeAPI
	xhtml_pre_typeAPI.ID = xhtml_pre_typeDB.ID
	xhtml_pre_typeAPI.Xhtml_pre_typePointersEncoding = xhtml_pre_typeDB.Xhtml_pre_typePointersEncoding
	xhtml_pre_typeDB.CopyBasicFieldsToXhtml_pre_type_WOP(&xhtml_pre_typeAPI.Xhtml_pre_type_WOP)

	c.JSON(http.StatusOK, xhtml_pre_typeAPI)
}

// UpdateXhtml_pre_type
//
// swagger:route PATCH /xhtml_pre_types/{ID} xhtml_pre_types updateXhtml_pre_type
//
// # Update a xhtml_pre_type
//
// Responses:
// default: genericError
//
//	200: xhtml_pre_typeDBResponse
func (controller *Controller) UpdateXhtml_pre_type(c *gin.Context) {

	mutexXhtml_pre_type.Lock()
	defer mutexXhtml_pre_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_pre_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_pre_type.GetDB()

	// Validate input
	var input orm.Xhtml_pre_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_pre_typeDB orm.Xhtml_pre_typeDB

	// fetch the xhtml_pre_type
	_, err := db.First(&xhtml_pre_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_pre_typeDB.CopyBasicFieldsFromXhtml_pre_type_WOP(&input.Xhtml_pre_type_WOP)
	xhtml_pre_typeDB.Xhtml_pre_typePointersEncoding = input.Xhtml_pre_typePointersEncoding

	db, _ = db.Model(&xhtml_pre_typeDB)
	_, err = db.Updates(&xhtml_pre_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_pre_typeNew := new(models.Xhtml_pre_type)
	xhtml_pre_typeDB.CopyBasicFieldsToXhtml_pre_type(xhtml_pre_typeNew)

	// redeem pointers
	xhtml_pre_typeDB.DecodePointers(backRepo, xhtml_pre_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_pre_typeOld := backRepo.BackRepoXhtml_pre_type.Map_Xhtml_pre_typeDBID_Xhtml_pre_typePtr[xhtml_pre_typeDB.ID]
	if xhtml_pre_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_pre_typeOld, xhtml_pre_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_pre_typeDB
	c.JSON(http.StatusOK, xhtml_pre_typeDB)
}

// DeleteXhtml_pre_type
//
// swagger:route DELETE /xhtml_pre_types/{ID} xhtml_pre_types deleteXhtml_pre_type
//
// # Delete a xhtml_pre_type
//
// default: genericError
//
//	200: xhtml_pre_typeDBResponse
func (controller *Controller) DeleteXhtml_pre_type(c *gin.Context) {

	mutexXhtml_pre_type.Lock()
	defer mutexXhtml_pre_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_pre_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_pre_type.GetDB()

	// Get model if exist
	var xhtml_pre_typeDB orm.Xhtml_pre_typeDB
	if _, err := db.First(&xhtml_pre_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_pre_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_pre_typeDeleted := new(models.Xhtml_pre_type)
	xhtml_pre_typeDB.CopyBasicFieldsToXhtml_pre_type(xhtml_pre_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_pre_typeStaged := backRepo.BackRepoXhtml_pre_type.Map_Xhtml_pre_typeDBID_Xhtml_pre_typePtr[xhtml_pre_typeDB.ID]
	if xhtml_pre_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_pre_typeStaged, xhtml_pre_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
