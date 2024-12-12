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
var __Xhtml_object_type__dummysDeclaration__ models.Xhtml_object_type
var __Xhtml_object_type_time__dummyDeclaration time.Duration

var mutexXhtml_object_type sync.Mutex

// An Xhtml_object_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_object_type updateXhtml_object_type deleteXhtml_object_type
type Xhtml_object_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_object_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_object_type updateXhtml_object_type
type Xhtml_object_typeInput struct {
	// The Xhtml_object_type to submit or modify
	// in: body
	Xhtml_object_type *orm.Xhtml_object_typeAPI
}

// GetXhtml_object_types
//
// swagger:route GET /xhtml_object_types xhtml_object_types getXhtml_object_types
//
// # Get all xhtml_object_types
//
// Responses:
// default: genericError
//
//	200: xhtml_object_typeDBResponse
func (controller *Controller) GetXhtml_object_types(c *gin.Context) {

	// source slice
	var xhtml_object_typeDBs []orm.Xhtml_object_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_object_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_object_type.GetDB()

	_, err := db.Find(&xhtml_object_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_object_typeAPIs := make([]orm.Xhtml_object_typeAPI, 0)

	// for each xhtml_object_type, update fields from the database nullable fields
	for idx := range xhtml_object_typeDBs {
		xhtml_object_typeDB := &xhtml_object_typeDBs[idx]
		_ = xhtml_object_typeDB
		var xhtml_object_typeAPI orm.Xhtml_object_typeAPI

		// insertion point for updating fields
		xhtml_object_typeAPI.ID = xhtml_object_typeDB.ID
		xhtml_object_typeDB.CopyBasicFieldsToXhtml_object_type_WOP(&xhtml_object_typeAPI.Xhtml_object_type_WOP)
		xhtml_object_typeAPI.Xhtml_object_typePointersEncoding = xhtml_object_typeDB.Xhtml_object_typePointersEncoding
		xhtml_object_typeAPIs = append(xhtml_object_typeAPIs, xhtml_object_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_object_typeAPIs)
}

// PostXhtml_object_type
//
// swagger:route POST /xhtml_object_types xhtml_object_types postXhtml_object_type
//
// Creates a xhtml_object_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_object_type(c *gin.Context) {

	mutexXhtml_object_type.Lock()
	defer mutexXhtml_object_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_object_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_object_type.GetDB()

	// Validate input
	var input orm.Xhtml_object_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_object_type
	xhtml_object_typeDB := orm.Xhtml_object_typeDB{}
	xhtml_object_typeDB.Xhtml_object_typePointersEncoding = input.Xhtml_object_typePointersEncoding
	xhtml_object_typeDB.CopyBasicFieldsFromXhtml_object_type_WOP(&input.Xhtml_object_type_WOP)

	_, err = db.Create(&xhtml_object_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_object_type.CheckoutPhaseOneInstance(&xhtml_object_typeDB)
	xhtml_object_type := backRepo.BackRepoXhtml_object_type.Map_Xhtml_object_typeDBID_Xhtml_object_typePtr[xhtml_object_typeDB.ID]

	if xhtml_object_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_object_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_object_typeDB)
}

// GetXhtml_object_type
//
// swagger:route GET /xhtml_object_types/{ID} xhtml_object_types getXhtml_object_type
//
// Gets the details for a xhtml_object_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_object_typeDBResponse
func (controller *Controller) GetXhtml_object_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_object_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_object_type.GetDB()

	// Get xhtml_object_typeDB in DB
	var xhtml_object_typeDB orm.Xhtml_object_typeDB
	if _, err := db.First(&xhtml_object_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_object_typeAPI orm.Xhtml_object_typeAPI
	xhtml_object_typeAPI.ID = xhtml_object_typeDB.ID
	xhtml_object_typeAPI.Xhtml_object_typePointersEncoding = xhtml_object_typeDB.Xhtml_object_typePointersEncoding
	xhtml_object_typeDB.CopyBasicFieldsToXhtml_object_type_WOP(&xhtml_object_typeAPI.Xhtml_object_type_WOP)

	c.JSON(http.StatusOK, xhtml_object_typeAPI)
}

// UpdateXhtml_object_type
//
// swagger:route PATCH /xhtml_object_types/{ID} xhtml_object_types updateXhtml_object_type
//
// # Update a xhtml_object_type
//
// Responses:
// default: genericError
//
//	200: xhtml_object_typeDBResponse
func (controller *Controller) UpdateXhtml_object_type(c *gin.Context) {

	mutexXhtml_object_type.Lock()
	defer mutexXhtml_object_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_object_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_object_type.GetDB()

	// Validate input
	var input orm.Xhtml_object_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_object_typeDB orm.Xhtml_object_typeDB

	// fetch the xhtml_object_type
	_, err := db.First(&xhtml_object_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_object_typeDB.CopyBasicFieldsFromXhtml_object_type_WOP(&input.Xhtml_object_type_WOP)
	xhtml_object_typeDB.Xhtml_object_typePointersEncoding = input.Xhtml_object_typePointersEncoding

	db, _ = db.Model(&xhtml_object_typeDB)
	_, err = db.Updates(&xhtml_object_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_object_typeNew := new(models.Xhtml_object_type)
	xhtml_object_typeDB.CopyBasicFieldsToXhtml_object_type(xhtml_object_typeNew)

	// redeem pointers
	xhtml_object_typeDB.DecodePointers(backRepo, xhtml_object_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_object_typeOld := backRepo.BackRepoXhtml_object_type.Map_Xhtml_object_typeDBID_Xhtml_object_typePtr[xhtml_object_typeDB.ID]
	if xhtml_object_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_object_typeOld, xhtml_object_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_object_typeDB
	c.JSON(http.StatusOK, xhtml_object_typeDB)
}

// DeleteXhtml_object_type
//
// swagger:route DELETE /xhtml_object_types/{ID} xhtml_object_types deleteXhtml_object_type
//
// # Delete a xhtml_object_type
//
// default: genericError
//
//	200: xhtml_object_typeDBResponse
func (controller *Controller) DeleteXhtml_object_type(c *gin.Context) {

	mutexXhtml_object_type.Lock()
	defer mutexXhtml_object_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_object_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_object_type.GetDB()

	// Get model if exist
	var xhtml_object_typeDB orm.Xhtml_object_typeDB
	if _, err := db.First(&xhtml_object_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_object_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_object_typeDeleted := new(models.Xhtml_object_type)
	xhtml_object_typeDB.CopyBasicFieldsToXhtml_object_type(xhtml_object_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_object_typeStaged := backRepo.BackRepoXhtml_object_type.Map_Xhtml_object_typeDBID_Xhtml_object_typePtr[xhtml_object_typeDB.ID]
	if xhtml_object_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_object_typeStaged, xhtml_object_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
