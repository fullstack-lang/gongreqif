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
var __Xhtml_h4_type__dummysDeclaration__ models.Xhtml_h4_type
var __Xhtml_h4_type_time__dummyDeclaration time.Duration

var mutexXhtml_h4_type sync.Mutex

// An Xhtml_h4_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_h4_type updateXhtml_h4_type deleteXhtml_h4_type
type Xhtml_h4_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_h4_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_h4_type updateXhtml_h4_type
type Xhtml_h4_typeInput struct {
	// The Xhtml_h4_type to submit or modify
	// in: body
	Xhtml_h4_type *orm.Xhtml_h4_typeAPI
}

// GetXhtml_h4_types
//
// swagger:route GET /xhtml_h4_types xhtml_h4_types getXhtml_h4_types
//
// # Get all xhtml_h4_types
//
// Responses:
// default: genericError
//
//	200: xhtml_h4_typeDBResponse
func (controller *Controller) GetXhtml_h4_types(c *gin.Context) {

	// source slice
	var xhtml_h4_typeDBs []orm.Xhtml_h4_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h4_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h4_type.GetDB()

	_, err := db.Find(&xhtml_h4_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_h4_typeAPIs := make([]orm.Xhtml_h4_typeAPI, 0)

	// for each xhtml_h4_type, update fields from the database nullable fields
	for idx := range xhtml_h4_typeDBs {
		xhtml_h4_typeDB := &xhtml_h4_typeDBs[idx]
		_ = xhtml_h4_typeDB
		var xhtml_h4_typeAPI orm.Xhtml_h4_typeAPI

		// insertion point for updating fields
		xhtml_h4_typeAPI.ID = xhtml_h4_typeDB.ID
		xhtml_h4_typeDB.CopyBasicFieldsToXhtml_h4_type_WOP(&xhtml_h4_typeAPI.Xhtml_h4_type_WOP)
		xhtml_h4_typeAPI.Xhtml_h4_typePointersEncoding = xhtml_h4_typeDB.Xhtml_h4_typePointersEncoding
		xhtml_h4_typeAPIs = append(xhtml_h4_typeAPIs, xhtml_h4_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_h4_typeAPIs)
}

// PostXhtml_h4_type
//
// swagger:route POST /xhtml_h4_types xhtml_h4_types postXhtml_h4_type
//
// Creates a xhtml_h4_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_h4_type(c *gin.Context) {

	mutexXhtml_h4_type.Lock()
	defer mutexXhtml_h4_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_h4_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h4_type.GetDB()

	// Validate input
	var input orm.Xhtml_h4_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_h4_type
	xhtml_h4_typeDB := orm.Xhtml_h4_typeDB{}
	xhtml_h4_typeDB.Xhtml_h4_typePointersEncoding = input.Xhtml_h4_typePointersEncoding
	xhtml_h4_typeDB.CopyBasicFieldsFromXhtml_h4_type_WOP(&input.Xhtml_h4_type_WOP)

	_, err = db.Create(&xhtml_h4_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_h4_type.CheckoutPhaseOneInstance(&xhtml_h4_typeDB)
	xhtml_h4_type := backRepo.BackRepoXhtml_h4_type.Map_Xhtml_h4_typeDBID_Xhtml_h4_typePtr[xhtml_h4_typeDB.ID]

	if xhtml_h4_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_h4_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_h4_typeDB)
}

// GetXhtml_h4_type
//
// swagger:route GET /xhtml_h4_types/{ID} xhtml_h4_types getXhtml_h4_type
//
// Gets the details for a xhtml_h4_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_h4_typeDBResponse
func (controller *Controller) GetXhtml_h4_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h4_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h4_type.GetDB()

	// Get xhtml_h4_typeDB in DB
	var xhtml_h4_typeDB orm.Xhtml_h4_typeDB
	if _, err := db.First(&xhtml_h4_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_h4_typeAPI orm.Xhtml_h4_typeAPI
	xhtml_h4_typeAPI.ID = xhtml_h4_typeDB.ID
	xhtml_h4_typeAPI.Xhtml_h4_typePointersEncoding = xhtml_h4_typeDB.Xhtml_h4_typePointersEncoding
	xhtml_h4_typeDB.CopyBasicFieldsToXhtml_h4_type_WOP(&xhtml_h4_typeAPI.Xhtml_h4_type_WOP)

	c.JSON(http.StatusOK, xhtml_h4_typeAPI)
}

// UpdateXhtml_h4_type
//
// swagger:route PATCH /xhtml_h4_types/{ID} xhtml_h4_types updateXhtml_h4_type
//
// # Update a xhtml_h4_type
//
// Responses:
// default: genericError
//
//	200: xhtml_h4_typeDBResponse
func (controller *Controller) UpdateXhtml_h4_type(c *gin.Context) {

	mutexXhtml_h4_type.Lock()
	defer mutexXhtml_h4_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_h4_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h4_type.GetDB()

	// Validate input
	var input orm.Xhtml_h4_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_h4_typeDB orm.Xhtml_h4_typeDB

	// fetch the xhtml_h4_type
	_, err := db.First(&xhtml_h4_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_h4_typeDB.CopyBasicFieldsFromXhtml_h4_type_WOP(&input.Xhtml_h4_type_WOP)
	xhtml_h4_typeDB.Xhtml_h4_typePointersEncoding = input.Xhtml_h4_typePointersEncoding

	db, _ = db.Model(&xhtml_h4_typeDB)
	_, err = db.Updates(&xhtml_h4_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h4_typeNew := new(models.Xhtml_h4_type)
	xhtml_h4_typeDB.CopyBasicFieldsToXhtml_h4_type(xhtml_h4_typeNew)

	// redeem pointers
	xhtml_h4_typeDB.DecodePointers(backRepo, xhtml_h4_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_h4_typeOld := backRepo.BackRepoXhtml_h4_type.Map_Xhtml_h4_typeDBID_Xhtml_h4_typePtr[xhtml_h4_typeDB.ID]
	if xhtml_h4_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_h4_typeOld, xhtml_h4_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_h4_typeDB
	c.JSON(http.StatusOK, xhtml_h4_typeDB)
}

// DeleteXhtml_h4_type
//
// swagger:route DELETE /xhtml_h4_types/{ID} xhtml_h4_types deleteXhtml_h4_type
//
// # Delete a xhtml_h4_type
//
// default: genericError
//
//	200: xhtml_h4_typeDBResponse
func (controller *Controller) DeleteXhtml_h4_type(c *gin.Context) {

	mutexXhtml_h4_type.Lock()
	defer mutexXhtml_h4_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_h4_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h4_type.GetDB()

	// Get model if exist
	var xhtml_h4_typeDB orm.Xhtml_h4_typeDB
	if _, err := db.First(&xhtml_h4_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_h4_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h4_typeDeleted := new(models.Xhtml_h4_type)
	xhtml_h4_typeDB.CopyBasicFieldsToXhtml_h4_type(xhtml_h4_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_h4_typeStaged := backRepo.BackRepoXhtml_h4_type.Map_Xhtml_h4_typeDBID_Xhtml_h4_typePtr[xhtml_h4_typeDB.ID]
	if xhtml_h4_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_h4_typeStaged, xhtml_h4_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
