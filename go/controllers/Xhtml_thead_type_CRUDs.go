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
var __Xhtml_thead_type__dummysDeclaration__ models.Xhtml_thead_type
var __Xhtml_thead_type_time__dummyDeclaration time.Duration

var mutexXhtml_thead_type sync.Mutex

// An Xhtml_thead_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_thead_type updateXhtml_thead_type deleteXhtml_thead_type
type Xhtml_thead_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_thead_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_thead_type updateXhtml_thead_type
type Xhtml_thead_typeInput struct {
	// The Xhtml_thead_type to submit or modify
	// in: body
	Xhtml_thead_type *orm.Xhtml_thead_typeAPI
}

// GetXhtml_thead_types
//
// swagger:route GET /xhtml_thead_types xhtml_thead_types getXhtml_thead_types
//
// # Get all xhtml_thead_types
//
// Responses:
// default: genericError
//
//	200: xhtml_thead_typeDBResponse
func (controller *Controller) GetXhtml_thead_types(c *gin.Context) {

	// source slice
	var xhtml_thead_typeDBs []orm.Xhtml_thead_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_thead_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_thead_type.GetDB()

	_, err := db.Find(&xhtml_thead_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_thead_typeAPIs := make([]orm.Xhtml_thead_typeAPI, 0)

	// for each xhtml_thead_type, update fields from the database nullable fields
	for idx := range xhtml_thead_typeDBs {
		xhtml_thead_typeDB := &xhtml_thead_typeDBs[idx]
		_ = xhtml_thead_typeDB
		var xhtml_thead_typeAPI orm.Xhtml_thead_typeAPI

		// insertion point for updating fields
		xhtml_thead_typeAPI.ID = xhtml_thead_typeDB.ID
		xhtml_thead_typeDB.CopyBasicFieldsToXhtml_thead_type_WOP(&xhtml_thead_typeAPI.Xhtml_thead_type_WOP)
		xhtml_thead_typeAPI.Xhtml_thead_typePointersEncoding = xhtml_thead_typeDB.Xhtml_thead_typePointersEncoding
		xhtml_thead_typeAPIs = append(xhtml_thead_typeAPIs, xhtml_thead_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_thead_typeAPIs)
}

// PostXhtml_thead_type
//
// swagger:route POST /xhtml_thead_types xhtml_thead_types postXhtml_thead_type
//
// Creates a xhtml_thead_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_thead_type(c *gin.Context) {

	mutexXhtml_thead_type.Lock()
	defer mutexXhtml_thead_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_thead_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_thead_type.GetDB()

	// Validate input
	var input orm.Xhtml_thead_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_thead_type
	xhtml_thead_typeDB := orm.Xhtml_thead_typeDB{}
	xhtml_thead_typeDB.Xhtml_thead_typePointersEncoding = input.Xhtml_thead_typePointersEncoding
	xhtml_thead_typeDB.CopyBasicFieldsFromXhtml_thead_type_WOP(&input.Xhtml_thead_type_WOP)

	_, err = db.Create(&xhtml_thead_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_thead_type.CheckoutPhaseOneInstance(&xhtml_thead_typeDB)
	xhtml_thead_type := backRepo.BackRepoXhtml_thead_type.Map_Xhtml_thead_typeDBID_Xhtml_thead_typePtr[xhtml_thead_typeDB.ID]

	if xhtml_thead_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_thead_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_thead_typeDB)
}

// GetXhtml_thead_type
//
// swagger:route GET /xhtml_thead_types/{ID} xhtml_thead_types getXhtml_thead_type
//
// Gets the details for a xhtml_thead_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_thead_typeDBResponse
func (controller *Controller) GetXhtml_thead_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_thead_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_thead_type.GetDB()

	// Get xhtml_thead_typeDB in DB
	var xhtml_thead_typeDB orm.Xhtml_thead_typeDB
	if _, err := db.First(&xhtml_thead_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_thead_typeAPI orm.Xhtml_thead_typeAPI
	xhtml_thead_typeAPI.ID = xhtml_thead_typeDB.ID
	xhtml_thead_typeAPI.Xhtml_thead_typePointersEncoding = xhtml_thead_typeDB.Xhtml_thead_typePointersEncoding
	xhtml_thead_typeDB.CopyBasicFieldsToXhtml_thead_type_WOP(&xhtml_thead_typeAPI.Xhtml_thead_type_WOP)

	c.JSON(http.StatusOK, xhtml_thead_typeAPI)
}

// UpdateXhtml_thead_type
//
// swagger:route PATCH /xhtml_thead_types/{ID} xhtml_thead_types updateXhtml_thead_type
//
// # Update a xhtml_thead_type
//
// Responses:
// default: genericError
//
//	200: xhtml_thead_typeDBResponse
func (controller *Controller) UpdateXhtml_thead_type(c *gin.Context) {

	mutexXhtml_thead_type.Lock()
	defer mutexXhtml_thead_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_thead_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_thead_type.GetDB()

	// Validate input
	var input orm.Xhtml_thead_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_thead_typeDB orm.Xhtml_thead_typeDB

	// fetch the xhtml_thead_type
	_, err := db.First(&xhtml_thead_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_thead_typeDB.CopyBasicFieldsFromXhtml_thead_type_WOP(&input.Xhtml_thead_type_WOP)
	xhtml_thead_typeDB.Xhtml_thead_typePointersEncoding = input.Xhtml_thead_typePointersEncoding

	db, _ = db.Model(&xhtml_thead_typeDB)
	_, err = db.Updates(&xhtml_thead_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_thead_typeNew := new(models.Xhtml_thead_type)
	xhtml_thead_typeDB.CopyBasicFieldsToXhtml_thead_type(xhtml_thead_typeNew)

	// redeem pointers
	xhtml_thead_typeDB.DecodePointers(backRepo, xhtml_thead_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_thead_typeOld := backRepo.BackRepoXhtml_thead_type.Map_Xhtml_thead_typeDBID_Xhtml_thead_typePtr[xhtml_thead_typeDB.ID]
	if xhtml_thead_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_thead_typeOld, xhtml_thead_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_thead_typeDB
	c.JSON(http.StatusOK, xhtml_thead_typeDB)
}

// DeleteXhtml_thead_type
//
// swagger:route DELETE /xhtml_thead_types/{ID} xhtml_thead_types deleteXhtml_thead_type
//
// # Delete a xhtml_thead_type
//
// default: genericError
//
//	200: xhtml_thead_typeDBResponse
func (controller *Controller) DeleteXhtml_thead_type(c *gin.Context) {

	mutexXhtml_thead_type.Lock()
	defer mutexXhtml_thead_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_thead_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_thead_type.GetDB()

	// Get model if exist
	var xhtml_thead_typeDB orm.Xhtml_thead_typeDB
	if _, err := db.First(&xhtml_thead_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_thead_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_thead_typeDeleted := new(models.Xhtml_thead_type)
	xhtml_thead_typeDB.CopyBasicFieldsToXhtml_thead_type(xhtml_thead_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_thead_typeStaged := backRepo.BackRepoXhtml_thead_type.Map_Xhtml_thead_typeDBID_Xhtml_thead_typePtr[xhtml_thead_typeDB.ID]
	if xhtml_thead_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_thead_typeStaged, xhtml_thead_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
