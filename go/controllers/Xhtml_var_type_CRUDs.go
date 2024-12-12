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
var __Xhtml_var_type__dummysDeclaration__ models.Xhtml_var_type
var __Xhtml_var_type_time__dummyDeclaration time.Duration

var mutexXhtml_var_type sync.Mutex

// An Xhtml_var_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_var_type updateXhtml_var_type deleteXhtml_var_type
type Xhtml_var_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_var_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_var_type updateXhtml_var_type
type Xhtml_var_typeInput struct {
	// The Xhtml_var_type to submit or modify
	// in: body
	Xhtml_var_type *orm.Xhtml_var_typeAPI
}

// GetXhtml_var_types
//
// swagger:route GET /xhtml_var_types xhtml_var_types getXhtml_var_types
//
// # Get all xhtml_var_types
//
// Responses:
// default: genericError
//
//	200: xhtml_var_typeDBResponse
func (controller *Controller) GetXhtml_var_types(c *gin.Context) {

	// source slice
	var xhtml_var_typeDBs []orm.Xhtml_var_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_var_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_var_type.GetDB()

	_, err := db.Find(&xhtml_var_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_var_typeAPIs := make([]orm.Xhtml_var_typeAPI, 0)

	// for each xhtml_var_type, update fields from the database nullable fields
	for idx := range xhtml_var_typeDBs {
		xhtml_var_typeDB := &xhtml_var_typeDBs[idx]
		_ = xhtml_var_typeDB
		var xhtml_var_typeAPI orm.Xhtml_var_typeAPI

		// insertion point for updating fields
		xhtml_var_typeAPI.ID = xhtml_var_typeDB.ID
		xhtml_var_typeDB.CopyBasicFieldsToXhtml_var_type_WOP(&xhtml_var_typeAPI.Xhtml_var_type_WOP)
		xhtml_var_typeAPI.Xhtml_var_typePointersEncoding = xhtml_var_typeDB.Xhtml_var_typePointersEncoding
		xhtml_var_typeAPIs = append(xhtml_var_typeAPIs, xhtml_var_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_var_typeAPIs)
}

// PostXhtml_var_type
//
// swagger:route POST /xhtml_var_types xhtml_var_types postXhtml_var_type
//
// Creates a xhtml_var_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_var_type(c *gin.Context) {

	mutexXhtml_var_type.Lock()
	defer mutexXhtml_var_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_var_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_var_type.GetDB()

	// Validate input
	var input orm.Xhtml_var_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_var_type
	xhtml_var_typeDB := orm.Xhtml_var_typeDB{}
	xhtml_var_typeDB.Xhtml_var_typePointersEncoding = input.Xhtml_var_typePointersEncoding
	xhtml_var_typeDB.CopyBasicFieldsFromXhtml_var_type_WOP(&input.Xhtml_var_type_WOP)

	_, err = db.Create(&xhtml_var_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_var_type.CheckoutPhaseOneInstance(&xhtml_var_typeDB)
	xhtml_var_type := backRepo.BackRepoXhtml_var_type.Map_Xhtml_var_typeDBID_Xhtml_var_typePtr[xhtml_var_typeDB.ID]

	if xhtml_var_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_var_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_var_typeDB)
}

// GetXhtml_var_type
//
// swagger:route GET /xhtml_var_types/{ID} xhtml_var_types getXhtml_var_type
//
// Gets the details for a xhtml_var_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_var_typeDBResponse
func (controller *Controller) GetXhtml_var_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_var_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_var_type.GetDB()

	// Get xhtml_var_typeDB in DB
	var xhtml_var_typeDB orm.Xhtml_var_typeDB
	if _, err := db.First(&xhtml_var_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_var_typeAPI orm.Xhtml_var_typeAPI
	xhtml_var_typeAPI.ID = xhtml_var_typeDB.ID
	xhtml_var_typeAPI.Xhtml_var_typePointersEncoding = xhtml_var_typeDB.Xhtml_var_typePointersEncoding
	xhtml_var_typeDB.CopyBasicFieldsToXhtml_var_type_WOP(&xhtml_var_typeAPI.Xhtml_var_type_WOP)

	c.JSON(http.StatusOK, xhtml_var_typeAPI)
}

// UpdateXhtml_var_type
//
// swagger:route PATCH /xhtml_var_types/{ID} xhtml_var_types updateXhtml_var_type
//
// # Update a xhtml_var_type
//
// Responses:
// default: genericError
//
//	200: xhtml_var_typeDBResponse
func (controller *Controller) UpdateXhtml_var_type(c *gin.Context) {

	mutexXhtml_var_type.Lock()
	defer mutexXhtml_var_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_var_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_var_type.GetDB()

	// Validate input
	var input orm.Xhtml_var_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_var_typeDB orm.Xhtml_var_typeDB

	// fetch the xhtml_var_type
	_, err := db.First(&xhtml_var_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_var_typeDB.CopyBasicFieldsFromXhtml_var_type_WOP(&input.Xhtml_var_type_WOP)
	xhtml_var_typeDB.Xhtml_var_typePointersEncoding = input.Xhtml_var_typePointersEncoding

	db, _ = db.Model(&xhtml_var_typeDB)
	_, err = db.Updates(&xhtml_var_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_var_typeNew := new(models.Xhtml_var_type)
	xhtml_var_typeDB.CopyBasicFieldsToXhtml_var_type(xhtml_var_typeNew)

	// redeem pointers
	xhtml_var_typeDB.DecodePointers(backRepo, xhtml_var_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_var_typeOld := backRepo.BackRepoXhtml_var_type.Map_Xhtml_var_typeDBID_Xhtml_var_typePtr[xhtml_var_typeDB.ID]
	if xhtml_var_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_var_typeOld, xhtml_var_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_var_typeDB
	c.JSON(http.StatusOK, xhtml_var_typeDB)
}

// DeleteXhtml_var_type
//
// swagger:route DELETE /xhtml_var_types/{ID} xhtml_var_types deleteXhtml_var_type
//
// # Delete a xhtml_var_type
//
// default: genericError
//
//	200: xhtml_var_typeDBResponse
func (controller *Controller) DeleteXhtml_var_type(c *gin.Context) {

	mutexXhtml_var_type.Lock()
	defer mutexXhtml_var_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_var_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_var_type.GetDB()

	// Get model if exist
	var xhtml_var_typeDB orm.Xhtml_var_typeDB
	if _, err := db.First(&xhtml_var_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_var_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_var_typeDeleted := new(models.Xhtml_var_type)
	xhtml_var_typeDB.CopyBasicFieldsToXhtml_var_type(xhtml_var_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_var_typeStaged := backRepo.BackRepoXhtml_var_type.Map_Xhtml_var_typeDBID_Xhtml_var_typePtr[xhtml_var_typeDB.ID]
	if xhtml_var_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_var_typeStaged, xhtml_var_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
