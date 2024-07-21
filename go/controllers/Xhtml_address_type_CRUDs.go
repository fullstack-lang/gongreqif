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
var __Xhtml_address_type__dummysDeclaration__ models.Xhtml_address_type
var __Xhtml_address_type_time__dummyDeclaration time.Duration

var mutexXhtml_address_type sync.Mutex

// An Xhtml_address_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_address_type updateXhtml_address_type deleteXhtml_address_type
type Xhtml_address_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_address_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_address_type updateXhtml_address_type
type Xhtml_address_typeInput struct {
	// The Xhtml_address_type to submit or modify
	// in: body
	Xhtml_address_type *orm.Xhtml_address_typeAPI
}

// GetXhtml_address_types
//
// swagger:route GET /xhtml_address_types xhtml_address_types getXhtml_address_types
//
// # Get all xhtml_address_types
//
// Responses:
// default: genericError
//
//	200: xhtml_address_typeDBResponse
func (controller *Controller) GetXhtml_address_types(c *gin.Context) {

	// source slice
	var xhtml_address_typeDBs []orm.Xhtml_address_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_address_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_address_type.GetDB()

	query := db.Find(&xhtml_address_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_address_typeAPIs := make([]orm.Xhtml_address_typeAPI, 0)

	// for each xhtml_address_type, update fields from the database nullable fields
	for idx := range xhtml_address_typeDBs {
		xhtml_address_typeDB := &xhtml_address_typeDBs[idx]
		_ = xhtml_address_typeDB
		var xhtml_address_typeAPI orm.Xhtml_address_typeAPI

		// insertion point for updating fields
		xhtml_address_typeAPI.ID = xhtml_address_typeDB.ID
		xhtml_address_typeDB.CopyBasicFieldsToXhtml_address_type_WOP(&xhtml_address_typeAPI.Xhtml_address_type_WOP)
		xhtml_address_typeAPI.Xhtml_address_typePointersEncoding = xhtml_address_typeDB.Xhtml_address_typePointersEncoding
		xhtml_address_typeAPIs = append(xhtml_address_typeAPIs, xhtml_address_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_address_typeAPIs)
}

// PostXhtml_address_type
//
// swagger:route POST /xhtml_address_types xhtml_address_types postXhtml_address_type
//
// Creates a xhtml_address_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_address_type(c *gin.Context) {

	mutexXhtml_address_type.Lock()
	defer mutexXhtml_address_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_address_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_address_type.GetDB()

	// Validate input
	var input orm.Xhtml_address_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_address_type
	xhtml_address_typeDB := orm.Xhtml_address_typeDB{}
	xhtml_address_typeDB.Xhtml_address_typePointersEncoding = input.Xhtml_address_typePointersEncoding
	xhtml_address_typeDB.CopyBasicFieldsFromXhtml_address_type_WOP(&input.Xhtml_address_type_WOP)

	query := db.Create(&xhtml_address_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_address_type.CheckoutPhaseOneInstance(&xhtml_address_typeDB)
	xhtml_address_type := backRepo.BackRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr[xhtml_address_typeDB.ID]

	if xhtml_address_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_address_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_address_typeDB)
}

// GetXhtml_address_type
//
// swagger:route GET /xhtml_address_types/{ID} xhtml_address_types getXhtml_address_type
//
// Gets the details for a xhtml_address_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_address_typeDBResponse
func (controller *Controller) GetXhtml_address_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_address_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_address_type.GetDB()

	// Get xhtml_address_typeDB in DB
	var xhtml_address_typeDB orm.Xhtml_address_typeDB
	if err := db.First(&xhtml_address_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_address_typeAPI orm.Xhtml_address_typeAPI
	xhtml_address_typeAPI.ID = xhtml_address_typeDB.ID
	xhtml_address_typeAPI.Xhtml_address_typePointersEncoding = xhtml_address_typeDB.Xhtml_address_typePointersEncoding
	xhtml_address_typeDB.CopyBasicFieldsToXhtml_address_type_WOP(&xhtml_address_typeAPI.Xhtml_address_type_WOP)

	c.JSON(http.StatusOK, xhtml_address_typeAPI)
}

// UpdateXhtml_address_type
//
// swagger:route PATCH /xhtml_address_types/{ID} xhtml_address_types updateXhtml_address_type
//
// # Update a xhtml_address_type
//
// Responses:
// default: genericError
//
//	200: xhtml_address_typeDBResponse
func (controller *Controller) UpdateXhtml_address_type(c *gin.Context) {

	mutexXhtml_address_type.Lock()
	defer mutexXhtml_address_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_address_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_address_type.GetDB()

	// Validate input
	var input orm.Xhtml_address_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_address_typeDB orm.Xhtml_address_typeDB

	// fetch the xhtml_address_type
	query := db.First(&xhtml_address_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_address_typeDB.CopyBasicFieldsFromXhtml_address_type_WOP(&input.Xhtml_address_type_WOP)
	xhtml_address_typeDB.Xhtml_address_typePointersEncoding = input.Xhtml_address_typePointersEncoding

	query = db.Model(&xhtml_address_typeDB).Updates(xhtml_address_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_address_typeNew := new(models.Xhtml_address_type)
	xhtml_address_typeDB.CopyBasicFieldsToXhtml_address_type(xhtml_address_typeNew)

	// redeem pointers
	xhtml_address_typeDB.DecodePointers(backRepo, xhtml_address_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_address_typeOld := backRepo.BackRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr[xhtml_address_typeDB.ID]
	if xhtml_address_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_address_typeOld, xhtml_address_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_address_typeDB
	c.JSON(http.StatusOK, xhtml_address_typeDB)
}

// DeleteXhtml_address_type
//
// swagger:route DELETE /xhtml_address_types/{ID} xhtml_address_types deleteXhtml_address_type
//
// # Delete a xhtml_address_type
//
// default: genericError
//
//	200: xhtml_address_typeDBResponse
func (controller *Controller) DeleteXhtml_address_type(c *gin.Context) {

	mutexXhtml_address_type.Lock()
	defer mutexXhtml_address_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_address_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_address_type.GetDB()

	// Get model if exist
	var xhtml_address_typeDB orm.Xhtml_address_typeDB
	if err := db.First(&xhtml_address_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_address_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_address_typeDeleted := new(models.Xhtml_address_type)
	xhtml_address_typeDB.CopyBasicFieldsToXhtml_address_type(xhtml_address_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_address_typeStaged := backRepo.BackRepoXhtml_address_type.Map_Xhtml_address_typeDBID_Xhtml_address_typePtr[xhtml_address_typeDB.ID]
	if xhtml_address_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_address_typeStaged, xhtml_address_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
