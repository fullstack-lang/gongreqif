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
var __Xhtml_h3_type__dummysDeclaration__ models.Xhtml_h3_type
var __Xhtml_h3_type_time__dummyDeclaration time.Duration

var mutexXhtml_h3_type sync.Mutex

// An Xhtml_h3_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_h3_type updateXhtml_h3_type deleteXhtml_h3_type
type Xhtml_h3_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_h3_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_h3_type updateXhtml_h3_type
type Xhtml_h3_typeInput struct {
	// The Xhtml_h3_type to submit or modify
	// in: body
	Xhtml_h3_type *orm.Xhtml_h3_typeAPI
}

// GetXhtml_h3_types
//
// swagger:route GET /xhtml_h3_types xhtml_h3_types getXhtml_h3_types
//
// # Get all xhtml_h3_types
//
// Responses:
// default: genericError
//
//	200: xhtml_h3_typeDBResponse
func (controller *Controller) GetXhtml_h3_types(c *gin.Context) {

	// source slice
	var xhtml_h3_typeDBs []orm.Xhtml_h3_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h3_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h3_type.GetDB()

	query := db.Find(&xhtml_h3_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_h3_typeAPIs := make([]orm.Xhtml_h3_typeAPI, 0)

	// for each xhtml_h3_type, update fields from the database nullable fields
	for idx := range xhtml_h3_typeDBs {
		xhtml_h3_typeDB := &xhtml_h3_typeDBs[idx]
		_ = xhtml_h3_typeDB
		var xhtml_h3_typeAPI orm.Xhtml_h3_typeAPI

		// insertion point for updating fields
		xhtml_h3_typeAPI.ID = xhtml_h3_typeDB.ID
		xhtml_h3_typeDB.CopyBasicFieldsToXhtml_h3_type_WOP(&xhtml_h3_typeAPI.Xhtml_h3_type_WOP)
		xhtml_h3_typeAPI.Xhtml_h3_typePointersEncoding = xhtml_h3_typeDB.Xhtml_h3_typePointersEncoding
		xhtml_h3_typeAPIs = append(xhtml_h3_typeAPIs, xhtml_h3_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_h3_typeAPIs)
}

// PostXhtml_h3_type
//
// swagger:route POST /xhtml_h3_types xhtml_h3_types postXhtml_h3_type
//
// Creates a xhtml_h3_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_h3_type(c *gin.Context) {

	mutexXhtml_h3_type.Lock()
	defer mutexXhtml_h3_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_h3_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h3_type.GetDB()

	// Validate input
	var input orm.Xhtml_h3_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_h3_type
	xhtml_h3_typeDB := orm.Xhtml_h3_typeDB{}
	xhtml_h3_typeDB.Xhtml_h3_typePointersEncoding = input.Xhtml_h3_typePointersEncoding
	xhtml_h3_typeDB.CopyBasicFieldsFromXhtml_h3_type_WOP(&input.Xhtml_h3_type_WOP)

	query := db.Create(&xhtml_h3_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_h3_type.CheckoutPhaseOneInstance(&xhtml_h3_typeDB)
	xhtml_h3_type := backRepo.BackRepoXhtml_h3_type.Map_Xhtml_h3_typeDBID_Xhtml_h3_typePtr[xhtml_h3_typeDB.ID]

	if xhtml_h3_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_h3_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_h3_typeDB)
}

// GetXhtml_h3_type
//
// swagger:route GET /xhtml_h3_types/{ID} xhtml_h3_types getXhtml_h3_type
//
// Gets the details for a xhtml_h3_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_h3_typeDBResponse
func (controller *Controller) GetXhtml_h3_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h3_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h3_type.GetDB()

	// Get xhtml_h3_typeDB in DB
	var xhtml_h3_typeDB orm.Xhtml_h3_typeDB
	if err := db.First(&xhtml_h3_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_h3_typeAPI orm.Xhtml_h3_typeAPI
	xhtml_h3_typeAPI.ID = xhtml_h3_typeDB.ID
	xhtml_h3_typeAPI.Xhtml_h3_typePointersEncoding = xhtml_h3_typeDB.Xhtml_h3_typePointersEncoding
	xhtml_h3_typeDB.CopyBasicFieldsToXhtml_h3_type_WOP(&xhtml_h3_typeAPI.Xhtml_h3_type_WOP)

	c.JSON(http.StatusOK, xhtml_h3_typeAPI)
}

// UpdateXhtml_h3_type
//
// swagger:route PATCH /xhtml_h3_types/{ID} xhtml_h3_types updateXhtml_h3_type
//
// # Update a xhtml_h3_type
//
// Responses:
// default: genericError
//
//	200: xhtml_h3_typeDBResponse
func (controller *Controller) UpdateXhtml_h3_type(c *gin.Context) {

	mutexXhtml_h3_type.Lock()
	defer mutexXhtml_h3_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_h3_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h3_type.GetDB()

	// Validate input
	var input orm.Xhtml_h3_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_h3_typeDB orm.Xhtml_h3_typeDB

	// fetch the xhtml_h3_type
	query := db.First(&xhtml_h3_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_h3_typeDB.CopyBasicFieldsFromXhtml_h3_type_WOP(&input.Xhtml_h3_type_WOP)
	xhtml_h3_typeDB.Xhtml_h3_typePointersEncoding = input.Xhtml_h3_typePointersEncoding

	query = db.Model(&xhtml_h3_typeDB).Updates(xhtml_h3_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h3_typeNew := new(models.Xhtml_h3_type)
	xhtml_h3_typeDB.CopyBasicFieldsToXhtml_h3_type(xhtml_h3_typeNew)

	// redeem pointers
	xhtml_h3_typeDB.DecodePointers(backRepo, xhtml_h3_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_h3_typeOld := backRepo.BackRepoXhtml_h3_type.Map_Xhtml_h3_typeDBID_Xhtml_h3_typePtr[xhtml_h3_typeDB.ID]
	if xhtml_h3_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_h3_typeOld, xhtml_h3_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_h3_typeDB
	c.JSON(http.StatusOK, xhtml_h3_typeDB)
}

// DeleteXhtml_h3_type
//
// swagger:route DELETE /xhtml_h3_types/{ID} xhtml_h3_types deleteXhtml_h3_type
//
// # Delete a xhtml_h3_type
//
// default: genericError
//
//	200: xhtml_h3_typeDBResponse
func (controller *Controller) DeleteXhtml_h3_type(c *gin.Context) {

	mutexXhtml_h3_type.Lock()
	defer mutexXhtml_h3_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_h3_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_h3_type.GetDB()

	// Get model if exist
	var xhtml_h3_typeDB orm.Xhtml_h3_typeDB
	if err := db.First(&xhtml_h3_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_h3_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h3_typeDeleted := new(models.Xhtml_h3_type)
	xhtml_h3_typeDB.CopyBasicFieldsToXhtml_h3_type(xhtml_h3_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_h3_typeStaged := backRepo.BackRepoXhtml_h3_type.Map_Xhtml_h3_typeDBID_Xhtml_h3_typePtr[xhtml_h3_typeDB.ID]
	if xhtml_h3_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_h3_typeStaged, xhtml_h3_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
