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
var __Xhtml_q_type__dummysDeclaration__ models.Xhtml_q_type
var __Xhtml_q_type_time__dummyDeclaration time.Duration

var mutexXhtml_q_type sync.Mutex

// An Xhtml_q_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_q_type updateXhtml_q_type deleteXhtml_q_type
type Xhtml_q_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_q_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_q_type updateXhtml_q_type
type Xhtml_q_typeInput struct {
	// The Xhtml_q_type to submit or modify
	// in: body
	Xhtml_q_type *orm.Xhtml_q_typeAPI
}

// GetXhtml_q_types
//
// swagger:route GET /xhtml_q_types xhtml_q_types getXhtml_q_types
//
// # Get all xhtml_q_types
//
// Responses:
// default: genericError
//
//	200: xhtml_q_typeDBResponse
func (controller *Controller) GetXhtml_q_types(c *gin.Context) {

	// source slice
	var xhtml_q_typeDBs []orm.Xhtml_q_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_q_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_q_type.GetDB()

	query := db.Find(&xhtml_q_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_q_typeAPIs := make([]orm.Xhtml_q_typeAPI, 0)

	// for each xhtml_q_type, update fields from the database nullable fields
	for idx := range xhtml_q_typeDBs {
		xhtml_q_typeDB := &xhtml_q_typeDBs[idx]
		_ = xhtml_q_typeDB
		var xhtml_q_typeAPI orm.Xhtml_q_typeAPI

		// insertion point for updating fields
		xhtml_q_typeAPI.ID = xhtml_q_typeDB.ID
		xhtml_q_typeDB.CopyBasicFieldsToXhtml_q_type_WOP(&xhtml_q_typeAPI.Xhtml_q_type_WOP)
		xhtml_q_typeAPI.Xhtml_q_typePointersEncoding = xhtml_q_typeDB.Xhtml_q_typePointersEncoding
		xhtml_q_typeAPIs = append(xhtml_q_typeAPIs, xhtml_q_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_q_typeAPIs)
}

// PostXhtml_q_type
//
// swagger:route POST /xhtml_q_types xhtml_q_types postXhtml_q_type
//
// Creates a xhtml_q_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_q_type(c *gin.Context) {

	mutexXhtml_q_type.Lock()
	defer mutexXhtml_q_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_q_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_q_type.GetDB()

	// Validate input
	var input orm.Xhtml_q_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_q_type
	xhtml_q_typeDB := orm.Xhtml_q_typeDB{}
	xhtml_q_typeDB.Xhtml_q_typePointersEncoding = input.Xhtml_q_typePointersEncoding
	xhtml_q_typeDB.CopyBasicFieldsFromXhtml_q_type_WOP(&input.Xhtml_q_type_WOP)

	query := db.Create(&xhtml_q_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_q_type.CheckoutPhaseOneInstance(&xhtml_q_typeDB)
	xhtml_q_type := backRepo.BackRepoXhtml_q_type.Map_Xhtml_q_typeDBID_Xhtml_q_typePtr[xhtml_q_typeDB.ID]

	if xhtml_q_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_q_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_q_typeDB)
}

// GetXhtml_q_type
//
// swagger:route GET /xhtml_q_types/{ID} xhtml_q_types getXhtml_q_type
//
// Gets the details for a xhtml_q_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_q_typeDBResponse
func (controller *Controller) GetXhtml_q_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_q_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_q_type.GetDB()

	// Get xhtml_q_typeDB in DB
	var xhtml_q_typeDB orm.Xhtml_q_typeDB
	if err := db.First(&xhtml_q_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_q_typeAPI orm.Xhtml_q_typeAPI
	xhtml_q_typeAPI.ID = xhtml_q_typeDB.ID
	xhtml_q_typeAPI.Xhtml_q_typePointersEncoding = xhtml_q_typeDB.Xhtml_q_typePointersEncoding
	xhtml_q_typeDB.CopyBasicFieldsToXhtml_q_type_WOP(&xhtml_q_typeAPI.Xhtml_q_type_WOP)

	c.JSON(http.StatusOK, xhtml_q_typeAPI)
}

// UpdateXhtml_q_type
//
// swagger:route PATCH /xhtml_q_types/{ID} xhtml_q_types updateXhtml_q_type
//
// # Update a xhtml_q_type
//
// Responses:
// default: genericError
//
//	200: xhtml_q_typeDBResponse
func (controller *Controller) UpdateXhtml_q_type(c *gin.Context) {

	mutexXhtml_q_type.Lock()
	defer mutexXhtml_q_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_q_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_q_type.GetDB()

	// Validate input
	var input orm.Xhtml_q_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_q_typeDB orm.Xhtml_q_typeDB

	// fetch the xhtml_q_type
	query := db.First(&xhtml_q_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_q_typeDB.CopyBasicFieldsFromXhtml_q_type_WOP(&input.Xhtml_q_type_WOP)
	xhtml_q_typeDB.Xhtml_q_typePointersEncoding = input.Xhtml_q_typePointersEncoding

	query = db.Model(&xhtml_q_typeDB).Updates(xhtml_q_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_q_typeNew := new(models.Xhtml_q_type)
	xhtml_q_typeDB.CopyBasicFieldsToXhtml_q_type(xhtml_q_typeNew)

	// redeem pointers
	xhtml_q_typeDB.DecodePointers(backRepo, xhtml_q_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_q_typeOld := backRepo.BackRepoXhtml_q_type.Map_Xhtml_q_typeDBID_Xhtml_q_typePtr[xhtml_q_typeDB.ID]
	if xhtml_q_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_q_typeOld, xhtml_q_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_q_typeDB
	c.JSON(http.StatusOK, xhtml_q_typeDB)
}

// DeleteXhtml_q_type
//
// swagger:route DELETE /xhtml_q_types/{ID} xhtml_q_types deleteXhtml_q_type
//
// # Delete a xhtml_q_type
//
// default: genericError
//
//	200: xhtml_q_typeDBResponse
func (controller *Controller) DeleteXhtml_q_type(c *gin.Context) {

	mutexXhtml_q_type.Lock()
	defer mutexXhtml_q_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_q_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_q_type.GetDB()

	// Get model if exist
	var xhtml_q_typeDB orm.Xhtml_q_typeDB
	if err := db.First(&xhtml_q_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_q_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_q_typeDeleted := new(models.Xhtml_q_type)
	xhtml_q_typeDB.CopyBasicFieldsToXhtml_q_type(xhtml_q_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_q_typeStaged := backRepo.BackRepoXhtml_q_type.Map_Xhtml_q_typeDBID_Xhtml_q_typePtr[xhtml_q_typeDB.ID]
	if xhtml_q_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_q_typeStaged, xhtml_q_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
