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
var __Xhtml_kbd_type__dummysDeclaration__ models.Xhtml_kbd_type
var __Xhtml_kbd_type_time__dummyDeclaration time.Duration

var mutexXhtml_kbd_type sync.Mutex

// An Xhtml_kbd_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_kbd_type updateXhtml_kbd_type deleteXhtml_kbd_type
type Xhtml_kbd_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_kbd_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_kbd_type updateXhtml_kbd_type
type Xhtml_kbd_typeInput struct {
	// The Xhtml_kbd_type to submit or modify
	// in: body
	Xhtml_kbd_type *orm.Xhtml_kbd_typeAPI
}

// GetXhtml_kbd_types
//
// swagger:route GET /xhtml_kbd_types xhtml_kbd_types getXhtml_kbd_types
//
// # Get all xhtml_kbd_types
//
// Responses:
// default: genericError
//
//	200: xhtml_kbd_typeDBResponse
func (controller *Controller) GetXhtml_kbd_types(c *gin.Context) {

	// source slice
	var xhtml_kbd_typeDBs []orm.Xhtml_kbd_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_kbd_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_kbd_type.GetDB()

	query := db.Find(&xhtml_kbd_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_kbd_typeAPIs := make([]orm.Xhtml_kbd_typeAPI, 0)

	// for each xhtml_kbd_type, update fields from the database nullable fields
	for idx := range xhtml_kbd_typeDBs {
		xhtml_kbd_typeDB := &xhtml_kbd_typeDBs[idx]
		_ = xhtml_kbd_typeDB
		var xhtml_kbd_typeAPI orm.Xhtml_kbd_typeAPI

		// insertion point for updating fields
		xhtml_kbd_typeAPI.ID = xhtml_kbd_typeDB.ID
		xhtml_kbd_typeDB.CopyBasicFieldsToXhtml_kbd_type_WOP(&xhtml_kbd_typeAPI.Xhtml_kbd_type_WOP)
		xhtml_kbd_typeAPI.Xhtml_kbd_typePointersEncoding = xhtml_kbd_typeDB.Xhtml_kbd_typePointersEncoding
		xhtml_kbd_typeAPIs = append(xhtml_kbd_typeAPIs, xhtml_kbd_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_kbd_typeAPIs)
}

// PostXhtml_kbd_type
//
// swagger:route POST /xhtml_kbd_types xhtml_kbd_types postXhtml_kbd_type
//
// Creates a xhtml_kbd_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_kbd_type(c *gin.Context) {

	mutexXhtml_kbd_type.Lock()
	defer mutexXhtml_kbd_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_kbd_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_kbd_type.GetDB()

	// Validate input
	var input orm.Xhtml_kbd_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_kbd_type
	xhtml_kbd_typeDB := orm.Xhtml_kbd_typeDB{}
	xhtml_kbd_typeDB.Xhtml_kbd_typePointersEncoding = input.Xhtml_kbd_typePointersEncoding
	xhtml_kbd_typeDB.CopyBasicFieldsFromXhtml_kbd_type_WOP(&input.Xhtml_kbd_type_WOP)

	query := db.Create(&xhtml_kbd_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_kbd_type.CheckoutPhaseOneInstance(&xhtml_kbd_typeDB)
	xhtml_kbd_type := backRepo.BackRepoXhtml_kbd_type.Map_Xhtml_kbd_typeDBID_Xhtml_kbd_typePtr[xhtml_kbd_typeDB.ID]

	if xhtml_kbd_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_kbd_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_kbd_typeDB)
}

// GetXhtml_kbd_type
//
// swagger:route GET /xhtml_kbd_types/{ID} xhtml_kbd_types getXhtml_kbd_type
//
// Gets the details for a xhtml_kbd_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_kbd_typeDBResponse
func (controller *Controller) GetXhtml_kbd_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_kbd_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_kbd_type.GetDB()

	// Get xhtml_kbd_typeDB in DB
	var xhtml_kbd_typeDB orm.Xhtml_kbd_typeDB
	if err := db.First(&xhtml_kbd_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_kbd_typeAPI orm.Xhtml_kbd_typeAPI
	xhtml_kbd_typeAPI.ID = xhtml_kbd_typeDB.ID
	xhtml_kbd_typeAPI.Xhtml_kbd_typePointersEncoding = xhtml_kbd_typeDB.Xhtml_kbd_typePointersEncoding
	xhtml_kbd_typeDB.CopyBasicFieldsToXhtml_kbd_type_WOP(&xhtml_kbd_typeAPI.Xhtml_kbd_type_WOP)

	c.JSON(http.StatusOK, xhtml_kbd_typeAPI)
}

// UpdateXhtml_kbd_type
//
// swagger:route PATCH /xhtml_kbd_types/{ID} xhtml_kbd_types updateXhtml_kbd_type
//
// # Update a xhtml_kbd_type
//
// Responses:
// default: genericError
//
//	200: xhtml_kbd_typeDBResponse
func (controller *Controller) UpdateXhtml_kbd_type(c *gin.Context) {

	mutexXhtml_kbd_type.Lock()
	defer mutexXhtml_kbd_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_kbd_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_kbd_type.GetDB()

	// Validate input
	var input orm.Xhtml_kbd_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_kbd_typeDB orm.Xhtml_kbd_typeDB

	// fetch the xhtml_kbd_type
	query := db.First(&xhtml_kbd_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_kbd_typeDB.CopyBasicFieldsFromXhtml_kbd_type_WOP(&input.Xhtml_kbd_type_WOP)
	xhtml_kbd_typeDB.Xhtml_kbd_typePointersEncoding = input.Xhtml_kbd_typePointersEncoding

	query = db.Model(&xhtml_kbd_typeDB).Updates(xhtml_kbd_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_kbd_typeNew := new(models.Xhtml_kbd_type)
	xhtml_kbd_typeDB.CopyBasicFieldsToXhtml_kbd_type(xhtml_kbd_typeNew)

	// redeem pointers
	xhtml_kbd_typeDB.DecodePointers(backRepo, xhtml_kbd_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_kbd_typeOld := backRepo.BackRepoXhtml_kbd_type.Map_Xhtml_kbd_typeDBID_Xhtml_kbd_typePtr[xhtml_kbd_typeDB.ID]
	if xhtml_kbd_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_kbd_typeOld, xhtml_kbd_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_kbd_typeDB
	c.JSON(http.StatusOK, xhtml_kbd_typeDB)
}

// DeleteXhtml_kbd_type
//
// swagger:route DELETE /xhtml_kbd_types/{ID} xhtml_kbd_types deleteXhtml_kbd_type
//
// # Delete a xhtml_kbd_type
//
// default: genericError
//
//	200: xhtml_kbd_typeDBResponse
func (controller *Controller) DeleteXhtml_kbd_type(c *gin.Context) {

	mutexXhtml_kbd_type.Lock()
	defer mutexXhtml_kbd_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_kbd_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_kbd_type.GetDB()

	// Get model if exist
	var xhtml_kbd_typeDB orm.Xhtml_kbd_typeDB
	if err := db.First(&xhtml_kbd_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_kbd_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_kbd_typeDeleted := new(models.Xhtml_kbd_type)
	xhtml_kbd_typeDB.CopyBasicFieldsToXhtml_kbd_type(xhtml_kbd_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_kbd_typeStaged := backRepo.BackRepoXhtml_kbd_type.Map_Xhtml_kbd_typeDBID_Xhtml_kbd_typePtr[xhtml_kbd_typeDB.ID]
	if xhtml_kbd_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_kbd_typeStaged, xhtml_kbd_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
