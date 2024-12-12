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
var __Xhtml_param_type__dummysDeclaration__ models.Xhtml_param_type
var __Xhtml_param_type_time__dummyDeclaration time.Duration

var mutexXhtml_param_type sync.Mutex

// An Xhtml_param_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_param_type updateXhtml_param_type deleteXhtml_param_type
type Xhtml_param_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_param_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_param_type updateXhtml_param_type
type Xhtml_param_typeInput struct {
	// The Xhtml_param_type to submit or modify
	// in: body
	Xhtml_param_type *orm.Xhtml_param_typeAPI
}

// GetXhtml_param_types
//
// swagger:route GET /xhtml_param_types xhtml_param_types getXhtml_param_types
//
// # Get all xhtml_param_types
//
// Responses:
// default: genericError
//
//	200: xhtml_param_typeDBResponse
func (controller *Controller) GetXhtml_param_types(c *gin.Context) {

	// source slice
	var xhtml_param_typeDBs []orm.Xhtml_param_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_param_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_param_type.GetDB()

	_, err := db.Find(&xhtml_param_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_param_typeAPIs := make([]orm.Xhtml_param_typeAPI, 0)

	// for each xhtml_param_type, update fields from the database nullable fields
	for idx := range xhtml_param_typeDBs {
		xhtml_param_typeDB := &xhtml_param_typeDBs[idx]
		_ = xhtml_param_typeDB
		var xhtml_param_typeAPI orm.Xhtml_param_typeAPI

		// insertion point for updating fields
		xhtml_param_typeAPI.ID = xhtml_param_typeDB.ID
		xhtml_param_typeDB.CopyBasicFieldsToXhtml_param_type_WOP(&xhtml_param_typeAPI.Xhtml_param_type_WOP)
		xhtml_param_typeAPI.Xhtml_param_typePointersEncoding = xhtml_param_typeDB.Xhtml_param_typePointersEncoding
		xhtml_param_typeAPIs = append(xhtml_param_typeAPIs, xhtml_param_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_param_typeAPIs)
}

// PostXhtml_param_type
//
// swagger:route POST /xhtml_param_types xhtml_param_types postXhtml_param_type
//
// Creates a xhtml_param_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_param_type(c *gin.Context) {

	mutexXhtml_param_type.Lock()
	defer mutexXhtml_param_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_param_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_param_type.GetDB()

	// Validate input
	var input orm.Xhtml_param_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_param_type
	xhtml_param_typeDB := orm.Xhtml_param_typeDB{}
	xhtml_param_typeDB.Xhtml_param_typePointersEncoding = input.Xhtml_param_typePointersEncoding
	xhtml_param_typeDB.CopyBasicFieldsFromXhtml_param_type_WOP(&input.Xhtml_param_type_WOP)

	_, err = db.Create(&xhtml_param_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_param_type.CheckoutPhaseOneInstance(&xhtml_param_typeDB)
	xhtml_param_type := backRepo.BackRepoXhtml_param_type.Map_Xhtml_param_typeDBID_Xhtml_param_typePtr[xhtml_param_typeDB.ID]

	if xhtml_param_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_param_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_param_typeDB)
}

// GetXhtml_param_type
//
// swagger:route GET /xhtml_param_types/{ID} xhtml_param_types getXhtml_param_type
//
// Gets the details for a xhtml_param_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_param_typeDBResponse
func (controller *Controller) GetXhtml_param_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_param_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_param_type.GetDB()

	// Get xhtml_param_typeDB in DB
	var xhtml_param_typeDB orm.Xhtml_param_typeDB
	if _, err := db.First(&xhtml_param_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_param_typeAPI orm.Xhtml_param_typeAPI
	xhtml_param_typeAPI.ID = xhtml_param_typeDB.ID
	xhtml_param_typeAPI.Xhtml_param_typePointersEncoding = xhtml_param_typeDB.Xhtml_param_typePointersEncoding
	xhtml_param_typeDB.CopyBasicFieldsToXhtml_param_type_WOP(&xhtml_param_typeAPI.Xhtml_param_type_WOP)

	c.JSON(http.StatusOK, xhtml_param_typeAPI)
}

// UpdateXhtml_param_type
//
// swagger:route PATCH /xhtml_param_types/{ID} xhtml_param_types updateXhtml_param_type
//
// # Update a xhtml_param_type
//
// Responses:
// default: genericError
//
//	200: xhtml_param_typeDBResponse
func (controller *Controller) UpdateXhtml_param_type(c *gin.Context) {

	mutexXhtml_param_type.Lock()
	defer mutexXhtml_param_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_param_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_param_type.GetDB()

	// Validate input
	var input orm.Xhtml_param_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_param_typeDB orm.Xhtml_param_typeDB

	// fetch the xhtml_param_type
	_, err := db.First(&xhtml_param_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_param_typeDB.CopyBasicFieldsFromXhtml_param_type_WOP(&input.Xhtml_param_type_WOP)
	xhtml_param_typeDB.Xhtml_param_typePointersEncoding = input.Xhtml_param_typePointersEncoding

	db, _ = db.Model(&xhtml_param_typeDB)
	_, err = db.Updates(&xhtml_param_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_param_typeNew := new(models.Xhtml_param_type)
	xhtml_param_typeDB.CopyBasicFieldsToXhtml_param_type(xhtml_param_typeNew)

	// redeem pointers
	xhtml_param_typeDB.DecodePointers(backRepo, xhtml_param_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_param_typeOld := backRepo.BackRepoXhtml_param_type.Map_Xhtml_param_typeDBID_Xhtml_param_typePtr[xhtml_param_typeDB.ID]
	if xhtml_param_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_param_typeOld, xhtml_param_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_param_typeDB
	c.JSON(http.StatusOK, xhtml_param_typeDB)
}

// DeleteXhtml_param_type
//
// swagger:route DELETE /xhtml_param_types/{ID} xhtml_param_types deleteXhtml_param_type
//
// # Delete a xhtml_param_type
//
// default: genericError
//
//	200: xhtml_param_typeDBResponse
func (controller *Controller) DeleteXhtml_param_type(c *gin.Context) {

	mutexXhtml_param_type.Lock()
	defer mutexXhtml_param_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_param_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_param_type.GetDB()

	// Get model if exist
	var xhtml_param_typeDB orm.Xhtml_param_typeDB
	if _, err := db.First(&xhtml_param_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_param_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_param_typeDeleted := new(models.Xhtml_param_type)
	xhtml_param_typeDB.CopyBasicFieldsToXhtml_param_type(xhtml_param_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_param_typeStaged := backRepo.BackRepoXhtml_param_type.Map_Xhtml_param_typeDBID_Xhtml_param_typePtr[xhtml_param_typeDB.ID]
	if xhtml_param_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_param_typeStaged, xhtml_param_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
