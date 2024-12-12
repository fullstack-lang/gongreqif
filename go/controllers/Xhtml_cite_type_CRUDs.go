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
var __Xhtml_cite_type__dummysDeclaration__ models.Xhtml_cite_type
var __Xhtml_cite_type_time__dummyDeclaration time.Duration

var mutexXhtml_cite_type sync.Mutex

// An Xhtml_cite_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_cite_type updateXhtml_cite_type deleteXhtml_cite_type
type Xhtml_cite_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_cite_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_cite_type updateXhtml_cite_type
type Xhtml_cite_typeInput struct {
	// The Xhtml_cite_type to submit or modify
	// in: body
	Xhtml_cite_type *orm.Xhtml_cite_typeAPI
}

// GetXhtml_cite_types
//
// swagger:route GET /xhtml_cite_types xhtml_cite_types getXhtml_cite_types
//
// # Get all xhtml_cite_types
//
// Responses:
// default: genericError
//
//	200: xhtml_cite_typeDBResponse
func (controller *Controller) GetXhtml_cite_types(c *gin.Context) {

	// source slice
	var xhtml_cite_typeDBs []orm.Xhtml_cite_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_cite_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_cite_type.GetDB()

	_, err := db.Find(&xhtml_cite_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_cite_typeAPIs := make([]orm.Xhtml_cite_typeAPI, 0)

	// for each xhtml_cite_type, update fields from the database nullable fields
	for idx := range xhtml_cite_typeDBs {
		xhtml_cite_typeDB := &xhtml_cite_typeDBs[idx]
		_ = xhtml_cite_typeDB
		var xhtml_cite_typeAPI orm.Xhtml_cite_typeAPI

		// insertion point for updating fields
		xhtml_cite_typeAPI.ID = xhtml_cite_typeDB.ID
		xhtml_cite_typeDB.CopyBasicFieldsToXhtml_cite_type_WOP(&xhtml_cite_typeAPI.Xhtml_cite_type_WOP)
		xhtml_cite_typeAPI.Xhtml_cite_typePointersEncoding = xhtml_cite_typeDB.Xhtml_cite_typePointersEncoding
		xhtml_cite_typeAPIs = append(xhtml_cite_typeAPIs, xhtml_cite_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_cite_typeAPIs)
}

// PostXhtml_cite_type
//
// swagger:route POST /xhtml_cite_types xhtml_cite_types postXhtml_cite_type
//
// Creates a xhtml_cite_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_cite_type(c *gin.Context) {

	mutexXhtml_cite_type.Lock()
	defer mutexXhtml_cite_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_cite_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_cite_type.GetDB()

	// Validate input
	var input orm.Xhtml_cite_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_cite_type
	xhtml_cite_typeDB := orm.Xhtml_cite_typeDB{}
	xhtml_cite_typeDB.Xhtml_cite_typePointersEncoding = input.Xhtml_cite_typePointersEncoding
	xhtml_cite_typeDB.CopyBasicFieldsFromXhtml_cite_type_WOP(&input.Xhtml_cite_type_WOP)

	_, err = db.Create(&xhtml_cite_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_cite_type.CheckoutPhaseOneInstance(&xhtml_cite_typeDB)
	xhtml_cite_type := backRepo.BackRepoXhtml_cite_type.Map_Xhtml_cite_typeDBID_Xhtml_cite_typePtr[xhtml_cite_typeDB.ID]

	if xhtml_cite_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_cite_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_cite_typeDB)
}

// GetXhtml_cite_type
//
// swagger:route GET /xhtml_cite_types/{ID} xhtml_cite_types getXhtml_cite_type
//
// Gets the details for a xhtml_cite_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_cite_typeDBResponse
func (controller *Controller) GetXhtml_cite_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_cite_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_cite_type.GetDB()

	// Get xhtml_cite_typeDB in DB
	var xhtml_cite_typeDB orm.Xhtml_cite_typeDB
	if _, err := db.First(&xhtml_cite_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_cite_typeAPI orm.Xhtml_cite_typeAPI
	xhtml_cite_typeAPI.ID = xhtml_cite_typeDB.ID
	xhtml_cite_typeAPI.Xhtml_cite_typePointersEncoding = xhtml_cite_typeDB.Xhtml_cite_typePointersEncoding
	xhtml_cite_typeDB.CopyBasicFieldsToXhtml_cite_type_WOP(&xhtml_cite_typeAPI.Xhtml_cite_type_WOP)

	c.JSON(http.StatusOK, xhtml_cite_typeAPI)
}

// UpdateXhtml_cite_type
//
// swagger:route PATCH /xhtml_cite_types/{ID} xhtml_cite_types updateXhtml_cite_type
//
// # Update a xhtml_cite_type
//
// Responses:
// default: genericError
//
//	200: xhtml_cite_typeDBResponse
func (controller *Controller) UpdateXhtml_cite_type(c *gin.Context) {

	mutexXhtml_cite_type.Lock()
	defer mutexXhtml_cite_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_cite_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_cite_type.GetDB()

	// Validate input
	var input orm.Xhtml_cite_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_cite_typeDB orm.Xhtml_cite_typeDB

	// fetch the xhtml_cite_type
	_, err := db.First(&xhtml_cite_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_cite_typeDB.CopyBasicFieldsFromXhtml_cite_type_WOP(&input.Xhtml_cite_type_WOP)
	xhtml_cite_typeDB.Xhtml_cite_typePointersEncoding = input.Xhtml_cite_typePointersEncoding

	db, _ = db.Model(&xhtml_cite_typeDB)
	_, err = db.Updates(&xhtml_cite_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_cite_typeNew := new(models.Xhtml_cite_type)
	xhtml_cite_typeDB.CopyBasicFieldsToXhtml_cite_type(xhtml_cite_typeNew)

	// redeem pointers
	xhtml_cite_typeDB.DecodePointers(backRepo, xhtml_cite_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_cite_typeOld := backRepo.BackRepoXhtml_cite_type.Map_Xhtml_cite_typeDBID_Xhtml_cite_typePtr[xhtml_cite_typeDB.ID]
	if xhtml_cite_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_cite_typeOld, xhtml_cite_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_cite_typeDB
	c.JSON(http.StatusOK, xhtml_cite_typeDB)
}

// DeleteXhtml_cite_type
//
// swagger:route DELETE /xhtml_cite_types/{ID} xhtml_cite_types deleteXhtml_cite_type
//
// # Delete a xhtml_cite_type
//
// default: genericError
//
//	200: xhtml_cite_typeDBResponse
func (controller *Controller) DeleteXhtml_cite_type(c *gin.Context) {

	mutexXhtml_cite_type.Lock()
	defer mutexXhtml_cite_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_cite_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_cite_type.GetDB()

	// Get model if exist
	var xhtml_cite_typeDB orm.Xhtml_cite_typeDB
	if _, err := db.First(&xhtml_cite_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_cite_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_cite_typeDeleted := new(models.Xhtml_cite_type)
	xhtml_cite_typeDB.CopyBasicFieldsToXhtml_cite_type(xhtml_cite_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_cite_typeStaged := backRepo.BackRepoXhtml_cite_type.Map_Xhtml_cite_typeDBID_Xhtml_cite_typePtr[xhtml_cite_typeDB.ID]
	if xhtml_cite_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_cite_typeStaged, xhtml_cite_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
