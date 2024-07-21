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
var __Xhtml_heading_type__dummysDeclaration__ models.Xhtml_heading_type
var __Xhtml_heading_type_time__dummyDeclaration time.Duration

var mutexXhtml_heading_type sync.Mutex

// An Xhtml_heading_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_heading_type updateXhtml_heading_type deleteXhtml_heading_type
type Xhtml_heading_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_heading_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_heading_type updateXhtml_heading_type
type Xhtml_heading_typeInput struct {
	// The Xhtml_heading_type to submit or modify
	// in: body
	Xhtml_heading_type *orm.Xhtml_heading_typeAPI
}

// GetXhtml_heading_types
//
// swagger:route GET /xhtml_heading_types xhtml_heading_types getXhtml_heading_types
//
// # Get all xhtml_heading_types
//
// Responses:
// default: genericError
//
//	200: xhtml_heading_typeDBResponse
func (controller *Controller) GetXhtml_heading_types(c *gin.Context) {

	// source slice
	var xhtml_heading_typeDBs []orm.Xhtml_heading_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_heading_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_heading_type.GetDB()

	query := db.Find(&xhtml_heading_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_heading_typeAPIs := make([]orm.Xhtml_heading_typeAPI, 0)

	// for each xhtml_heading_type, update fields from the database nullable fields
	for idx := range xhtml_heading_typeDBs {
		xhtml_heading_typeDB := &xhtml_heading_typeDBs[idx]
		_ = xhtml_heading_typeDB
		var xhtml_heading_typeAPI orm.Xhtml_heading_typeAPI

		// insertion point for updating fields
		xhtml_heading_typeAPI.ID = xhtml_heading_typeDB.ID
		xhtml_heading_typeDB.CopyBasicFieldsToXhtml_heading_type_WOP(&xhtml_heading_typeAPI.Xhtml_heading_type_WOP)
		xhtml_heading_typeAPI.Xhtml_heading_typePointersEncoding = xhtml_heading_typeDB.Xhtml_heading_typePointersEncoding
		xhtml_heading_typeAPIs = append(xhtml_heading_typeAPIs, xhtml_heading_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_heading_typeAPIs)
}

// PostXhtml_heading_type
//
// swagger:route POST /xhtml_heading_types xhtml_heading_types postXhtml_heading_type
//
// Creates a xhtml_heading_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_heading_type(c *gin.Context) {

	mutexXhtml_heading_type.Lock()
	defer mutexXhtml_heading_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_heading_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_heading_type.GetDB()

	// Validate input
	var input orm.Xhtml_heading_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_heading_type
	xhtml_heading_typeDB := orm.Xhtml_heading_typeDB{}
	xhtml_heading_typeDB.Xhtml_heading_typePointersEncoding = input.Xhtml_heading_typePointersEncoding
	xhtml_heading_typeDB.CopyBasicFieldsFromXhtml_heading_type_WOP(&input.Xhtml_heading_type_WOP)

	query := db.Create(&xhtml_heading_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_heading_type.CheckoutPhaseOneInstance(&xhtml_heading_typeDB)
	xhtml_heading_type := backRepo.BackRepoXhtml_heading_type.Map_Xhtml_heading_typeDBID_Xhtml_heading_typePtr[xhtml_heading_typeDB.ID]

	if xhtml_heading_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_heading_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_heading_typeDB)
}

// GetXhtml_heading_type
//
// swagger:route GET /xhtml_heading_types/{ID} xhtml_heading_types getXhtml_heading_type
//
// Gets the details for a xhtml_heading_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_heading_typeDBResponse
func (controller *Controller) GetXhtml_heading_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_heading_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_heading_type.GetDB()

	// Get xhtml_heading_typeDB in DB
	var xhtml_heading_typeDB orm.Xhtml_heading_typeDB
	if err := db.First(&xhtml_heading_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_heading_typeAPI orm.Xhtml_heading_typeAPI
	xhtml_heading_typeAPI.ID = xhtml_heading_typeDB.ID
	xhtml_heading_typeAPI.Xhtml_heading_typePointersEncoding = xhtml_heading_typeDB.Xhtml_heading_typePointersEncoding
	xhtml_heading_typeDB.CopyBasicFieldsToXhtml_heading_type_WOP(&xhtml_heading_typeAPI.Xhtml_heading_type_WOP)

	c.JSON(http.StatusOK, xhtml_heading_typeAPI)
}

// UpdateXhtml_heading_type
//
// swagger:route PATCH /xhtml_heading_types/{ID} xhtml_heading_types updateXhtml_heading_type
//
// # Update a xhtml_heading_type
//
// Responses:
// default: genericError
//
//	200: xhtml_heading_typeDBResponse
func (controller *Controller) UpdateXhtml_heading_type(c *gin.Context) {

	mutexXhtml_heading_type.Lock()
	defer mutexXhtml_heading_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_heading_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_heading_type.GetDB()

	// Validate input
	var input orm.Xhtml_heading_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_heading_typeDB orm.Xhtml_heading_typeDB

	// fetch the xhtml_heading_type
	query := db.First(&xhtml_heading_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_heading_typeDB.CopyBasicFieldsFromXhtml_heading_type_WOP(&input.Xhtml_heading_type_WOP)
	xhtml_heading_typeDB.Xhtml_heading_typePointersEncoding = input.Xhtml_heading_typePointersEncoding

	query = db.Model(&xhtml_heading_typeDB).Updates(xhtml_heading_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_heading_typeNew := new(models.Xhtml_heading_type)
	xhtml_heading_typeDB.CopyBasicFieldsToXhtml_heading_type(xhtml_heading_typeNew)

	// redeem pointers
	xhtml_heading_typeDB.DecodePointers(backRepo, xhtml_heading_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_heading_typeOld := backRepo.BackRepoXhtml_heading_type.Map_Xhtml_heading_typeDBID_Xhtml_heading_typePtr[xhtml_heading_typeDB.ID]
	if xhtml_heading_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_heading_typeOld, xhtml_heading_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_heading_typeDB
	c.JSON(http.StatusOK, xhtml_heading_typeDB)
}

// DeleteXhtml_heading_type
//
// swagger:route DELETE /xhtml_heading_types/{ID} xhtml_heading_types deleteXhtml_heading_type
//
// # Delete a xhtml_heading_type
//
// default: genericError
//
//	200: xhtml_heading_typeDBResponse
func (controller *Controller) DeleteXhtml_heading_type(c *gin.Context) {

	mutexXhtml_heading_type.Lock()
	defer mutexXhtml_heading_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_heading_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_heading_type.GetDB()

	// Get model if exist
	var xhtml_heading_typeDB orm.Xhtml_heading_typeDB
	if err := db.First(&xhtml_heading_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_heading_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_heading_typeDeleted := new(models.Xhtml_heading_type)
	xhtml_heading_typeDB.CopyBasicFieldsToXhtml_heading_type(xhtml_heading_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_heading_typeStaged := backRepo.BackRepoXhtml_heading_type.Map_Xhtml_heading_typeDBID_Xhtml_heading_typePtr[xhtml_heading_typeDB.ID]
	if xhtml_heading_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_heading_typeStaged, xhtml_heading_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
