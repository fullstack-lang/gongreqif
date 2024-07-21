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
var __Xhtml_acronym_type__dummysDeclaration__ models.Xhtml_acronym_type
var __Xhtml_acronym_type_time__dummyDeclaration time.Duration

var mutexXhtml_acronym_type sync.Mutex

// An Xhtml_acronym_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_acronym_type updateXhtml_acronym_type deleteXhtml_acronym_type
type Xhtml_acronym_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_acronym_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_acronym_type updateXhtml_acronym_type
type Xhtml_acronym_typeInput struct {
	// The Xhtml_acronym_type to submit or modify
	// in: body
	Xhtml_acronym_type *orm.Xhtml_acronym_typeAPI
}

// GetXhtml_acronym_types
//
// swagger:route GET /xhtml_acronym_types xhtml_acronym_types getXhtml_acronym_types
//
// # Get all xhtml_acronym_types
//
// Responses:
// default: genericError
//
//	200: xhtml_acronym_typeDBResponse
func (controller *Controller) GetXhtml_acronym_types(c *gin.Context) {

	// source slice
	var xhtml_acronym_typeDBs []orm.Xhtml_acronym_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_acronym_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_acronym_type.GetDB()

	query := db.Find(&xhtml_acronym_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_acronym_typeAPIs := make([]orm.Xhtml_acronym_typeAPI, 0)

	// for each xhtml_acronym_type, update fields from the database nullable fields
	for idx := range xhtml_acronym_typeDBs {
		xhtml_acronym_typeDB := &xhtml_acronym_typeDBs[idx]
		_ = xhtml_acronym_typeDB
		var xhtml_acronym_typeAPI orm.Xhtml_acronym_typeAPI

		// insertion point for updating fields
		xhtml_acronym_typeAPI.ID = xhtml_acronym_typeDB.ID
		xhtml_acronym_typeDB.CopyBasicFieldsToXhtml_acronym_type_WOP(&xhtml_acronym_typeAPI.Xhtml_acronym_type_WOP)
		xhtml_acronym_typeAPI.Xhtml_acronym_typePointersEncoding = xhtml_acronym_typeDB.Xhtml_acronym_typePointersEncoding
		xhtml_acronym_typeAPIs = append(xhtml_acronym_typeAPIs, xhtml_acronym_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_acronym_typeAPIs)
}

// PostXhtml_acronym_type
//
// swagger:route POST /xhtml_acronym_types xhtml_acronym_types postXhtml_acronym_type
//
// Creates a xhtml_acronym_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_acronym_type(c *gin.Context) {

	mutexXhtml_acronym_type.Lock()
	defer mutexXhtml_acronym_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_acronym_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_acronym_type.GetDB()

	// Validate input
	var input orm.Xhtml_acronym_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_acronym_type
	xhtml_acronym_typeDB := orm.Xhtml_acronym_typeDB{}
	xhtml_acronym_typeDB.Xhtml_acronym_typePointersEncoding = input.Xhtml_acronym_typePointersEncoding
	xhtml_acronym_typeDB.CopyBasicFieldsFromXhtml_acronym_type_WOP(&input.Xhtml_acronym_type_WOP)

	query := db.Create(&xhtml_acronym_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_acronym_type.CheckoutPhaseOneInstance(&xhtml_acronym_typeDB)
	xhtml_acronym_type := backRepo.BackRepoXhtml_acronym_type.Map_Xhtml_acronym_typeDBID_Xhtml_acronym_typePtr[xhtml_acronym_typeDB.ID]

	if xhtml_acronym_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_acronym_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_acronym_typeDB)
}

// GetXhtml_acronym_type
//
// swagger:route GET /xhtml_acronym_types/{ID} xhtml_acronym_types getXhtml_acronym_type
//
// Gets the details for a xhtml_acronym_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_acronym_typeDBResponse
func (controller *Controller) GetXhtml_acronym_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_acronym_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_acronym_type.GetDB()

	// Get xhtml_acronym_typeDB in DB
	var xhtml_acronym_typeDB orm.Xhtml_acronym_typeDB
	if err := db.First(&xhtml_acronym_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_acronym_typeAPI orm.Xhtml_acronym_typeAPI
	xhtml_acronym_typeAPI.ID = xhtml_acronym_typeDB.ID
	xhtml_acronym_typeAPI.Xhtml_acronym_typePointersEncoding = xhtml_acronym_typeDB.Xhtml_acronym_typePointersEncoding
	xhtml_acronym_typeDB.CopyBasicFieldsToXhtml_acronym_type_WOP(&xhtml_acronym_typeAPI.Xhtml_acronym_type_WOP)

	c.JSON(http.StatusOK, xhtml_acronym_typeAPI)
}

// UpdateXhtml_acronym_type
//
// swagger:route PATCH /xhtml_acronym_types/{ID} xhtml_acronym_types updateXhtml_acronym_type
//
// # Update a xhtml_acronym_type
//
// Responses:
// default: genericError
//
//	200: xhtml_acronym_typeDBResponse
func (controller *Controller) UpdateXhtml_acronym_type(c *gin.Context) {

	mutexXhtml_acronym_type.Lock()
	defer mutexXhtml_acronym_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_acronym_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_acronym_type.GetDB()

	// Validate input
	var input orm.Xhtml_acronym_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_acronym_typeDB orm.Xhtml_acronym_typeDB

	// fetch the xhtml_acronym_type
	query := db.First(&xhtml_acronym_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_acronym_typeDB.CopyBasicFieldsFromXhtml_acronym_type_WOP(&input.Xhtml_acronym_type_WOP)
	xhtml_acronym_typeDB.Xhtml_acronym_typePointersEncoding = input.Xhtml_acronym_typePointersEncoding

	query = db.Model(&xhtml_acronym_typeDB).Updates(xhtml_acronym_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_acronym_typeNew := new(models.Xhtml_acronym_type)
	xhtml_acronym_typeDB.CopyBasicFieldsToXhtml_acronym_type(xhtml_acronym_typeNew)

	// redeem pointers
	xhtml_acronym_typeDB.DecodePointers(backRepo, xhtml_acronym_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_acronym_typeOld := backRepo.BackRepoXhtml_acronym_type.Map_Xhtml_acronym_typeDBID_Xhtml_acronym_typePtr[xhtml_acronym_typeDB.ID]
	if xhtml_acronym_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_acronym_typeOld, xhtml_acronym_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_acronym_typeDB
	c.JSON(http.StatusOK, xhtml_acronym_typeDB)
}

// DeleteXhtml_acronym_type
//
// swagger:route DELETE /xhtml_acronym_types/{ID} xhtml_acronym_types deleteXhtml_acronym_type
//
// # Delete a xhtml_acronym_type
//
// default: genericError
//
//	200: xhtml_acronym_typeDBResponse
func (controller *Controller) DeleteXhtml_acronym_type(c *gin.Context) {

	mutexXhtml_acronym_type.Lock()
	defer mutexXhtml_acronym_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_acronym_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_acronym_type.GetDB()

	// Get model if exist
	var xhtml_acronym_typeDB orm.Xhtml_acronym_typeDB
	if err := db.First(&xhtml_acronym_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_acronym_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_acronym_typeDeleted := new(models.Xhtml_acronym_type)
	xhtml_acronym_typeDB.CopyBasicFieldsToXhtml_acronym_type(xhtml_acronym_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_acronym_typeStaged := backRepo.BackRepoXhtml_acronym_type.Map_Xhtml_acronym_typeDBID_Xhtml_acronym_typePtr[xhtml_acronym_typeDB.ID]
	if xhtml_acronym_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_acronym_typeStaged, xhtml_acronym_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
