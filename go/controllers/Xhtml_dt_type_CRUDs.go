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
var __Xhtml_dt_type__dummysDeclaration__ models.Xhtml_dt_type
var __Xhtml_dt_type_time__dummyDeclaration time.Duration

var mutexXhtml_dt_type sync.Mutex

// An Xhtml_dt_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_dt_type updateXhtml_dt_type deleteXhtml_dt_type
type Xhtml_dt_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_dt_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_dt_type updateXhtml_dt_type
type Xhtml_dt_typeInput struct {
	// The Xhtml_dt_type to submit or modify
	// in: body
	Xhtml_dt_type *orm.Xhtml_dt_typeAPI
}

// GetXhtml_dt_types
//
// swagger:route GET /xhtml_dt_types xhtml_dt_types getXhtml_dt_types
//
// # Get all xhtml_dt_types
//
// Responses:
// default: genericError
//
//	200: xhtml_dt_typeDBResponse
func (controller *Controller) GetXhtml_dt_types(c *gin.Context) {

	// source slice
	var xhtml_dt_typeDBs []orm.Xhtml_dt_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_dt_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dt_type.GetDB()

	_, err := db.Find(&xhtml_dt_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_dt_typeAPIs := make([]orm.Xhtml_dt_typeAPI, 0)

	// for each xhtml_dt_type, update fields from the database nullable fields
	for idx := range xhtml_dt_typeDBs {
		xhtml_dt_typeDB := &xhtml_dt_typeDBs[idx]
		_ = xhtml_dt_typeDB
		var xhtml_dt_typeAPI orm.Xhtml_dt_typeAPI

		// insertion point for updating fields
		xhtml_dt_typeAPI.ID = xhtml_dt_typeDB.ID
		xhtml_dt_typeDB.CopyBasicFieldsToXhtml_dt_type_WOP(&xhtml_dt_typeAPI.Xhtml_dt_type_WOP)
		xhtml_dt_typeAPI.Xhtml_dt_typePointersEncoding = xhtml_dt_typeDB.Xhtml_dt_typePointersEncoding
		xhtml_dt_typeAPIs = append(xhtml_dt_typeAPIs, xhtml_dt_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_dt_typeAPIs)
}

// PostXhtml_dt_type
//
// swagger:route POST /xhtml_dt_types xhtml_dt_types postXhtml_dt_type
//
// Creates a xhtml_dt_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_dt_type(c *gin.Context) {

	mutexXhtml_dt_type.Lock()
	defer mutexXhtml_dt_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_dt_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dt_type.GetDB()

	// Validate input
	var input orm.Xhtml_dt_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_dt_type
	xhtml_dt_typeDB := orm.Xhtml_dt_typeDB{}
	xhtml_dt_typeDB.Xhtml_dt_typePointersEncoding = input.Xhtml_dt_typePointersEncoding
	xhtml_dt_typeDB.CopyBasicFieldsFromXhtml_dt_type_WOP(&input.Xhtml_dt_type_WOP)

	_, err = db.Create(&xhtml_dt_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_dt_type.CheckoutPhaseOneInstance(&xhtml_dt_typeDB)
	xhtml_dt_type := backRepo.BackRepoXhtml_dt_type.Map_Xhtml_dt_typeDBID_Xhtml_dt_typePtr[xhtml_dt_typeDB.ID]

	if xhtml_dt_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_dt_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_dt_typeDB)
}

// GetXhtml_dt_type
//
// swagger:route GET /xhtml_dt_types/{ID} xhtml_dt_types getXhtml_dt_type
//
// Gets the details for a xhtml_dt_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_dt_typeDBResponse
func (controller *Controller) GetXhtml_dt_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_dt_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dt_type.GetDB()

	// Get xhtml_dt_typeDB in DB
	var xhtml_dt_typeDB orm.Xhtml_dt_typeDB
	if _, err := db.First(&xhtml_dt_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_dt_typeAPI orm.Xhtml_dt_typeAPI
	xhtml_dt_typeAPI.ID = xhtml_dt_typeDB.ID
	xhtml_dt_typeAPI.Xhtml_dt_typePointersEncoding = xhtml_dt_typeDB.Xhtml_dt_typePointersEncoding
	xhtml_dt_typeDB.CopyBasicFieldsToXhtml_dt_type_WOP(&xhtml_dt_typeAPI.Xhtml_dt_type_WOP)

	c.JSON(http.StatusOK, xhtml_dt_typeAPI)
}

// UpdateXhtml_dt_type
//
// swagger:route PATCH /xhtml_dt_types/{ID} xhtml_dt_types updateXhtml_dt_type
//
// # Update a xhtml_dt_type
//
// Responses:
// default: genericError
//
//	200: xhtml_dt_typeDBResponse
func (controller *Controller) UpdateXhtml_dt_type(c *gin.Context) {

	mutexXhtml_dt_type.Lock()
	defer mutexXhtml_dt_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_dt_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dt_type.GetDB()

	// Validate input
	var input orm.Xhtml_dt_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_dt_typeDB orm.Xhtml_dt_typeDB

	// fetch the xhtml_dt_type
	_, err := db.First(&xhtml_dt_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_dt_typeDB.CopyBasicFieldsFromXhtml_dt_type_WOP(&input.Xhtml_dt_type_WOP)
	xhtml_dt_typeDB.Xhtml_dt_typePointersEncoding = input.Xhtml_dt_typePointersEncoding

	db, _ = db.Model(&xhtml_dt_typeDB)
	_, err = db.Updates(&xhtml_dt_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_dt_typeNew := new(models.Xhtml_dt_type)
	xhtml_dt_typeDB.CopyBasicFieldsToXhtml_dt_type(xhtml_dt_typeNew)

	// redeem pointers
	xhtml_dt_typeDB.DecodePointers(backRepo, xhtml_dt_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_dt_typeOld := backRepo.BackRepoXhtml_dt_type.Map_Xhtml_dt_typeDBID_Xhtml_dt_typePtr[xhtml_dt_typeDB.ID]
	if xhtml_dt_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_dt_typeOld, xhtml_dt_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_dt_typeDB
	c.JSON(http.StatusOK, xhtml_dt_typeDB)
}

// DeleteXhtml_dt_type
//
// swagger:route DELETE /xhtml_dt_types/{ID} xhtml_dt_types deleteXhtml_dt_type
//
// # Delete a xhtml_dt_type
//
// default: genericError
//
//	200: xhtml_dt_typeDBResponse
func (controller *Controller) DeleteXhtml_dt_type(c *gin.Context) {

	mutexXhtml_dt_type.Lock()
	defer mutexXhtml_dt_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_dt_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dt_type.GetDB()

	// Get model if exist
	var xhtml_dt_typeDB orm.Xhtml_dt_typeDB
	if _, err := db.First(&xhtml_dt_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_dt_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_dt_typeDeleted := new(models.Xhtml_dt_type)
	xhtml_dt_typeDB.CopyBasicFieldsToXhtml_dt_type(xhtml_dt_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_dt_typeStaged := backRepo.BackRepoXhtml_dt_type.Map_Xhtml_dt_typeDBID_Xhtml_dt_typePtr[xhtml_dt_typeDB.ID]
	if xhtml_dt_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_dt_typeStaged, xhtml_dt_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
