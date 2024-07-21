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
var __Xhtml_samp_type__dummysDeclaration__ models.Xhtml_samp_type
var __Xhtml_samp_type_time__dummyDeclaration time.Duration

var mutexXhtml_samp_type sync.Mutex

// An Xhtml_samp_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_samp_type updateXhtml_samp_type deleteXhtml_samp_type
type Xhtml_samp_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_samp_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_samp_type updateXhtml_samp_type
type Xhtml_samp_typeInput struct {
	// The Xhtml_samp_type to submit or modify
	// in: body
	Xhtml_samp_type *orm.Xhtml_samp_typeAPI
}

// GetXhtml_samp_types
//
// swagger:route GET /xhtml_samp_types xhtml_samp_types getXhtml_samp_types
//
// # Get all xhtml_samp_types
//
// Responses:
// default: genericError
//
//	200: xhtml_samp_typeDBResponse
func (controller *Controller) GetXhtml_samp_types(c *gin.Context) {

	// source slice
	var xhtml_samp_typeDBs []orm.Xhtml_samp_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_samp_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_samp_type.GetDB()

	query := db.Find(&xhtml_samp_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_samp_typeAPIs := make([]orm.Xhtml_samp_typeAPI, 0)

	// for each xhtml_samp_type, update fields from the database nullable fields
	for idx := range xhtml_samp_typeDBs {
		xhtml_samp_typeDB := &xhtml_samp_typeDBs[idx]
		_ = xhtml_samp_typeDB
		var xhtml_samp_typeAPI orm.Xhtml_samp_typeAPI

		// insertion point for updating fields
		xhtml_samp_typeAPI.ID = xhtml_samp_typeDB.ID
		xhtml_samp_typeDB.CopyBasicFieldsToXhtml_samp_type_WOP(&xhtml_samp_typeAPI.Xhtml_samp_type_WOP)
		xhtml_samp_typeAPI.Xhtml_samp_typePointersEncoding = xhtml_samp_typeDB.Xhtml_samp_typePointersEncoding
		xhtml_samp_typeAPIs = append(xhtml_samp_typeAPIs, xhtml_samp_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_samp_typeAPIs)
}

// PostXhtml_samp_type
//
// swagger:route POST /xhtml_samp_types xhtml_samp_types postXhtml_samp_type
//
// Creates a xhtml_samp_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_samp_type(c *gin.Context) {

	mutexXhtml_samp_type.Lock()
	defer mutexXhtml_samp_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_samp_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_samp_type.GetDB()

	// Validate input
	var input orm.Xhtml_samp_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_samp_type
	xhtml_samp_typeDB := orm.Xhtml_samp_typeDB{}
	xhtml_samp_typeDB.Xhtml_samp_typePointersEncoding = input.Xhtml_samp_typePointersEncoding
	xhtml_samp_typeDB.CopyBasicFieldsFromXhtml_samp_type_WOP(&input.Xhtml_samp_type_WOP)

	query := db.Create(&xhtml_samp_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_samp_type.CheckoutPhaseOneInstance(&xhtml_samp_typeDB)
	xhtml_samp_type := backRepo.BackRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr[xhtml_samp_typeDB.ID]

	if xhtml_samp_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_samp_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_samp_typeDB)
}

// GetXhtml_samp_type
//
// swagger:route GET /xhtml_samp_types/{ID} xhtml_samp_types getXhtml_samp_type
//
// Gets the details for a xhtml_samp_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_samp_typeDBResponse
func (controller *Controller) GetXhtml_samp_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_samp_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_samp_type.GetDB()

	// Get xhtml_samp_typeDB in DB
	var xhtml_samp_typeDB orm.Xhtml_samp_typeDB
	if err := db.First(&xhtml_samp_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_samp_typeAPI orm.Xhtml_samp_typeAPI
	xhtml_samp_typeAPI.ID = xhtml_samp_typeDB.ID
	xhtml_samp_typeAPI.Xhtml_samp_typePointersEncoding = xhtml_samp_typeDB.Xhtml_samp_typePointersEncoding
	xhtml_samp_typeDB.CopyBasicFieldsToXhtml_samp_type_WOP(&xhtml_samp_typeAPI.Xhtml_samp_type_WOP)

	c.JSON(http.StatusOK, xhtml_samp_typeAPI)
}

// UpdateXhtml_samp_type
//
// swagger:route PATCH /xhtml_samp_types/{ID} xhtml_samp_types updateXhtml_samp_type
//
// # Update a xhtml_samp_type
//
// Responses:
// default: genericError
//
//	200: xhtml_samp_typeDBResponse
func (controller *Controller) UpdateXhtml_samp_type(c *gin.Context) {

	mutexXhtml_samp_type.Lock()
	defer mutexXhtml_samp_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_samp_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_samp_type.GetDB()

	// Validate input
	var input orm.Xhtml_samp_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_samp_typeDB orm.Xhtml_samp_typeDB

	// fetch the xhtml_samp_type
	query := db.First(&xhtml_samp_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_samp_typeDB.CopyBasicFieldsFromXhtml_samp_type_WOP(&input.Xhtml_samp_type_WOP)
	xhtml_samp_typeDB.Xhtml_samp_typePointersEncoding = input.Xhtml_samp_typePointersEncoding

	query = db.Model(&xhtml_samp_typeDB).Updates(xhtml_samp_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_samp_typeNew := new(models.Xhtml_samp_type)
	xhtml_samp_typeDB.CopyBasicFieldsToXhtml_samp_type(xhtml_samp_typeNew)

	// redeem pointers
	xhtml_samp_typeDB.DecodePointers(backRepo, xhtml_samp_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_samp_typeOld := backRepo.BackRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr[xhtml_samp_typeDB.ID]
	if xhtml_samp_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_samp_typeOld, xhtml_samp_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_samp_typeDB
	c.JSON(http.StatusOK, xhtml_samp_typeDB)
}

// DeleteXhtml_samp_type
//
// swagger:route DELETE /xhtml_samp_types/{ID} xhtml_samp_types deleteXhtml_samp_type
//
// # Delete a xhtml_samp_type
//
// default: genericError
//
//	200: xhtml_samp_typeDBResponse
func (controller *Controller) DeleteXhtml_samp_type(c *gin.Context) {

	mutexXhtml_samp_type.Lock()
	defer mutexXhtml_samp_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_samp_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_samp_type.GetDB()

	// Get model if exist
	var xhtml_samp_typeDB orm.Xhtml_samp_typeDB
	if err := db.First(&xhtml_samp_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_samp_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_samp_typeDeleted := new(models.Xhtml_samp_type)
	xhtml_samp_typeDB.CopyBasicFieldsToXhtml_samp_type(xhtml_samp_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_samp_typeStaged := backRepo.BackRepoXhtml_samp_type.Map_Xhtml_samp_typeDBID_Xhtml_samp_typePtr[xhtml_samp_typeDB.ID]
	if xhtml_samp_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_samp_typeStaged, xhtml_samp_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
