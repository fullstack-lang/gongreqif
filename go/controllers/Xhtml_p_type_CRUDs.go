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
var __Xhtml_p_type__dummysDeclaration__ models.Xhtml_p_type
var __Xhtml_p_type_time__dummyDeclaration time.Duration

var mutexXhtml_p_type sync.Mutex

// An Xhtml_p_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_p_type updateXhtml_p_type deleteXhtml_p_type
type Xhtml_p_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_p_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_p_type updateXhtml_p_type
type Xhtml_p_typeInput struct {
	// The Xhtml_p_type to submit or modify
	// in: body
	Xhtml_p_type *orm.Xhtml_p_typeAPI
}

// GetXhtml_p_types
//
// swagger:route GET /xhtml_p_types xhtml_p_types getXhtml_p_types
//
// # Get all xhtml_p_types
//
// Responses:
// default: genericError
//
//	200: xhtml_p_typeDBResponse
func (controller *Controller) GetXhtml_p_types(c *gin.Context) {

	// source slice
	var xhtml_p_typeDBs []orm.Xhtml_p_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_p_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_p_type.GetDB()

	_, err := db.Find(&xhtml_p_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_p_typeAPIs := make([]orm.Xhtml_p_typeAPI, 0)

	// for each xhtml_p_type, update fields from the database nullable fields
	for idx := range xhtml_p_typeDBs {
		xhtml_p_typeDB := &xhtml_p_typeDBs[idx]
		_ = xhtml_p_typeDB
		var xhtml_p_typeAPI orm.Xhtml_p_typeAPI

		// insertion point for updating fields
		xhtml_p_typeAPI.ID = xhtml_p_typeDB.ID
		xhtml_p_typeDB.CopyBasicFieldsToXhtml_p_type_WOP(&xhtml_p_typeAPI.Xhtml_p_type_WOP)
		xhtml_p_typeAPI.Xhtml_p_typePointersEncoding = xhtml_p_typeDB.Xhtml_p_typePointersEncoding
		xhtml_p_typeAPIs = append(xhtml_p_typeAPIs, xhtml_p_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_p_typeAPIs)
}

// PostXhtml_p_type
//
// swagger:route POST /xhtml_p_types xhtml_p_types postXhtml_p_type
//
// Creates a xhtml_p_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_p_type(c *gin.Context) {

	mutexXhtml_p_type.Lock()
	defer mutexXhtml_p_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_p_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_p_type.GetDB()

	// Validate input
	var input orm.Xhtml_p_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_p_type
	xhtml_p_typeDB := orm.Xhtml_p_typeDB{}
	xhtml_p_typeDB.Xhtml_p_typePointersEncoding = input.Xhtml_p_typePointersEncoding
	xhtml_p_typeDB.CopyBasicFieldsFromXhtml_p_type_WOP(&input.Xhtml_p_type_WOP)

	_, err = db.Create(&xhtml_p_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_p_type.CheckoutPhaseOneInstance(&xhtml_p_typeDB)
	xhtml_p_type := backRepo.BackRepoXhtml_p_type.Map_Xhtml_p_typeDBID_Xhtml_p_typePtr[xhtml_p_typeDB.ID]

	if xhtml_p_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_p_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_p_typeDB)
}

// GetXhtml_p_type
//
// swagger:route GET /xhtml_p_types/{ID} xhtml_p_types getXhtml_p_type
//
// Gets the details for a xhtml_p_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_p_typeDBResponse
func (controller *Controller) GetXhtml_p_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_p_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_p_type.GetDB()

	// Get xhtml_p_typeDB in DB
	var xhtml_p_typeDB orm.Xhtml_p_typeDB
	if _, err := db.First(&xhtml_p_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_p_typeAPI orm.Xhtml_p_typeAPI
	xhtml_p_typeAPI.ID = xhtml_p_typeDB.ID
	xhtml_p_typeAPI.Xhtml_p_typePointersEncoding = xhtml_p_typeDB.Xhtml_p_typePointersEncoding
	xhtml_p_typeDB.CopyBasicFieldsToXhtml_p_type_WOP(&xhtml_p_typeAPI.Xhtml_p_type_WOP)

	c.JSON(http.StatusOK, xhtml_p_typeAPI)
}

// UpdateXhtml_p_type
//
// swagger:route PATCH /xhtml_p_types/{ID} xhtml_p_types updateXhtml_p_type
//
// # Update a xhtml_p_type
//
// Responses:
// default: genericError
//
//	200: xhtml_p_typeDBResponse
func (controller *Controller) UpdateXhtml_p_type(c *gin.Context) {

	mutexXhtml_p_type.Lock()
	defer mutexXhtml_p_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_p_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_p_type.GetDB()

	// Validate input
	var input orm.Xhtml_p_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_p_typeDB orm.Xhtml_p_typeDB

	// fetch the xhtml_p_type
	_, err := db.First(&xhtml_p_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_p_typeDB.CopyBasicFieldsFromXhtml_p_type_WOP(&input.Xhtml_p_type_WOP)
	xhtml_p_typeDB.Xhtml_p_typePointersEncoding = input.Xhtml_p_typePointersEncoding

	db, _ = db.Model(&xhtml_p_typeDB)
	_, err = db.Updates(&xhtml_p_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_p_typeNew := new(models.Xhtml_p_type)
	xhtml_p_typeDB.CopyBasicFieldsToXhtml_p_type(xhtml_p_typeNew)

	// redeem pointers
	xhtml_p_typeDB.DecodePointers(backRepo, xhtml_p_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_p_typeOld := backRepo.BackRepoXhtml_p_type.Map_Xhtml_p_typeDBID_Xhtml_p_typePtr[xhtml_p_typeDB.ID]
	if xhtml_p_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_p_typeOld, xhtml_p_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_p_typeDB
	c.JSON(http.StatusOK, xhtml_p_typeDB)
}

// DeleteXhtml_p_type
//
// swagger:route DELETE /xhtml_p_types/{ID} xhtml_p_types deleteXhtml_p_type
//
// # Delete a xhtml_p_type
//
// default: genericError
//
//	200: xhtml_p_typeDBResponse
func (controller *Controller) DeleteXhtml_p_type(c *gin.Context) {

	mutexXhtml_p_type.Lock()
	defer mutexXhtml_p_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_p_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_p_type.GetDB()

	// Get model if exist
	var xhtml_p_typeDB orm.Xhtml_p_typeDB
	if _, err := db.First(&xhtml_p_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_p_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_p_typeDeleted := new(models.Xhtml_p_type)
	xhtml_p_typeDB.CopyBasicFieldsToXhtml_p_type(xhtml_p_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_p_typeStaged := backRepo.BackRepoXhtml_p_type.Map_Xhtml_p_typeDBID_Xhtml_p_typePtr[xhtml_p_typeDB.ID]
	if xhtml_p_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_p_typeStaged, xhtml_p_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
