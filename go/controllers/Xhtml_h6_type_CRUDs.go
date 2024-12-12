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
var __Xhtml_h6_type__dummysDeclaration__ models.Xhtml_h6_type
var __Xhtml_h6_type_time__dummyDeclaration time.Duration

var mutexXhtml_h6_type sync.Mutex

// An Xhtml_h6_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_h6_type updateXhtml_h6_type deleteXhtml_h6_type
type Xhtml_h6_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_h6_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_h6_type updateXhtml_h6_type
type Xhtml_h6_typeInput struct {
	// The Xhtml_h6_type to submit or modify
	// in: body
	Xhtml_h6_type *orm.Xhtml_h6_typeAPI
}

// GetXhtml_h6_types
//
// swagger:route GET /xhtml_h6_types xhtml_h6_types getXhtml_h6_types
//
// # Get all xhtml_h6_types
//
// Responses:
// default: genericError
//
//	200: xhtml_h6_typeDBResponse
func (controller *Controller) GetXhtml_h6_types(c *gin.Context) {

	// source slice
	var xhtml_h6_typeDBs []orm.Xhtml_h6_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h6_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h6_type.GetDB()

	_, err := db.Find(&xhtml_h6_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_h6_typeAPIs := make([]orm.Xhtml_h6_typeAPI, 0)

	// for each xhtml_h6_type, update fields from the database nullable fields
	for idx := range xhtml_h6_typeDBs {
		xhtml_h6_typeDB := &xhtml_h6_typeDBs[idx]
		_ = xhtml_h6_typeDB
		var xhtml_h6_typeAPI orm.Xhtml_h6_typeAPI

		// insertion point for updating fields
		xhtml_h6_typeAPI.ID = xhtml_h6_typeDB.ID
		xhtml_h6_typeDB.CopyBasicFieldsToXhtml_h6_type_WOP(&xhtml_h6_typeAPI.Xhtml_h6_type_WOP)
		xhtml_h6_typeAPI.Xhtml_h6_typePointersEncoding = xhtml_h6_typeDB.Xhtml_h6_typePointersEncoding
		xhtml_h6_typeAPIs = append(xhtml_h6_typeAPIs, xhtml_h6_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_h6_typeAPIs)
}

// PostXhtml_h6_type
//
// swagger:route POST /xhtml_h6_types xhtml_h6_types postXhtml_h6_type
//
// Creates a xhtml_h6_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_h6_type(c *gin.Context) {

	mutexXhtml_h6_type.Lock()
	defer mutexXhtml_h6_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_h6_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h6_type.GetDB()

	// Validate input
	var input orm.Xhtml_h6_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_h6_type
	xhtml_h6_typeDB := orm.Xhtml_h6_typeDB{}
	xhtml_h6_typeDB.Xhtml_h6_typePointersEncoding = input.Xhtml_h6_typePointersEncoding
	xhtml_h6_typeDB.CopyBasicFieldsFromXhtml_h6_type_WOP(&input.Xhtml_h6_type_WOP)

	_, err = db.Create(&xhtml_h6_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_h6_type.CheckoutPhaseOneInstance(&xhtml_h6_typeDB)
	xhtml_h6_type := backRepo.BackRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr[xhtml_h6_typeDB.ID]

	if xhtml_h6_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_h6_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_h6_typeDB)
}

// GetXhtml_h6_type
//
// swagger:route GET /xhtml_h6_types/{ID} xhtml_h6_types getXhtml_h6_type
//
// Gets the details for a xhtml_h6_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_h6_typeDBResponse
func (controller *Controller) GetXhtml_h6_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h6_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h6_type.GetDB()

	// Get xhtml_h6_typeDB in DB
	var xhtml_h6_typeDB orm.Xhtml_h6_typeDB
	if _, err := db.First(&xhtml_h6_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_h6_typeAPI orm.Xhtml_h6_typeAPI
	xhtml_h6_typeAPI.ID = xhtml_h6_typeDB.ID
	xhtml_h6_typeAPI.Xhtml_h6_typePointersEncoding = xhtml_h6_typeDB.Xhtml_h6_typePointersEncoding
	xhtml_h6_typeDB.CopyBasicFieldsToXhtml_h6_type_WOP(&xhtml_h6_typeAPI.Xhtml_h6_type_WOP)

	c.JSON(http.StatusOK, xhtml_h6_typeAPI)
}

// UpdateXhtml_h6_type
//
// swagger:route PATCH /xhtml_h6_types/{ID} xhtml_h6_types updateXhtml_h6_type
//
// # Update a xhtml_h6_type
//
// Responses:
// default: genericError
//
//	200: xhtml_h6_typeDBResponse
func (controller *Controller) UpdateXhtml_h6_type(c *gin.Context) {

	mutexXhtml_h6_type.Lock()
	defer mutexXhtml_h6_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_h6_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h6_type.GetDB()

	// Validate input
	var input orm.Xhtml_h6_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_h6_typeDB orm.Xhtml_h6_typeDB

	// fetch the xhtml_h6_type
	_, err := db.First(&xhtml_h6_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_h6_typeDB.CopyBasicFieldsFromXhtml_h6_type_WOP(&input.Xhtml_h6_type_WOP)
	xhtml_h6_typeDB.Xhtml_h6_typePointersEncoding = input.Xhtml_h6_typePointersEncoding

	db, _ = db.Model(&xhtml_h6_typeDB)
	_, err = db.Updates(&xhtml_h6_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h6_typeNew := new(models.Xhtml_h6_type)
	xhtml_h6_typeDB.CopyBasicFieldsToXhtml_h6_type(xhtml_h6_typeNew)

	// redeem pointers
	xhtml_h6_typeDB.DecodePointers(backRepo, xhtml_h6_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_h6_typeOld := backRepo.BackRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr[xhtml_h6_typeDB.ID]
	if xhtml_h6_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_h6_typeOld, xhtml_h6_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_h6_typeDB
	c.JSON(http.StatusOK, xhtml_h6_typeDB)
}

// DeleteXhtml_h6_type
//
// swagger:route DELETE /xhtml_h6_types/{ID} xhtml_h6_types deleteXhtml_h6_type
//
// # Delete a xhtml_h6_type
//
// default: genericError
//
//	200: xhtml_h6_typeDBResponse
func (controller *Controller) DeleteXhtml_h6_type(c *gin.Context) {

	mutexXhtml_h6_type.Lock()
	defer mutexXhtml_h6_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_h6_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h6_type.GetDB()

	// Get model if exist
	var xhtml_h6_typeDB orm.Xhtml_h6_typeDB
	if _, err := db.First(&xhtml_h6_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_h6_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h6_typeDeleted := new(models.Xhtml_h6_type)
	xhtml_h6_typeDB.CopyBasicFieldsToXhtml_h6_type(xhtml_h6_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_h6_typeStaged := backRepo.BackRepoXhtml_h6_type.Map_Xhtml_h6_typeDBID_Xhtml_h6_typePtr[xhtml_h6_typeDB.ID]
	if xhtml_h6_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_h6_typeStaged, xhtml_h6_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
