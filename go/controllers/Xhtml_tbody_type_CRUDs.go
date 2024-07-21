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
var __Xhtml_tbody_type__dummysDeclaration__ models.Xhtml_tbody_type
var __Xhtml_tbody_type_time__dummyDeclaration time.Duration

var mutexXhtml_tbody_type sync.Mutex

// An Xhtml_tbody_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_tbody_type updateXhtml_tbody_type deleteXhtml_tbody_type
type Xhtml_tbody_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_tbody_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_tbody_type updateXhtml_tbody_type
type Xhtml_tbody_typeInput struct {
	// The Xhtml_tbody_type to submit or modify
	// in: body
	Xhtml_tbody_type *orm.Xhtml_tbody_typeAPI
}

// GetXhtml_tbody_types
//
// swagger:route GET /xhtml_tbody_types xhtml_tbody_types getXhtml_tbody_types
//
// # Get all xhtml_tbody_types
//
// Responses:
// default: genericError
//
//	200: xhtml_tbody_typeDBResponse
func (controller *Controller) GetXhtml_tbody_types(c *gin.Context) {

	// source slice
	var xhtml_tbody_typeDBs []orm.Xhtml_tbody_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_tbody_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_tbody_type.GetDB()

	query := db.Find(&xhtml_tbody_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_tbody_typeAPIs := make([]orm.Xhtml_tbody_typeAPI, 0)

	// for each xhtml_tbody_type, update fields from the database nullable fields
	for idx := range xhtml_tbody_typeDBs {
		xhtml_tbody_typeDB := &xhtml_tbody_typeDBs[idx]
		_ = xhtml_tbody_typeDB
		var xhtml_tbody_typeAPI orm.Xhtml_tbody_typeAPI

		// insertion point for updating fields
		xhtml_tbody_typeAPI.ID = xhtml_tbody_typeDB.ID
		xhtml_tbody_typeDB.CopyBasicFieldsToXhtml_tbody_type_WOP(&xhtml_tbody_typeAPI.Xhtml_tbody_type_WOP)
		xhtml_tbody_typeAPI.Xhtml_tbody_typePointersEncoding = xhtml_tbody_typeDB.Xhtml_tbody_typePointersEncoding
		xhtml_tbody_typeAPIs = append(xhtml_tbody_typeAPIs, xhtml_tbody_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_tbody_typeAPIs)
}

// PostXhtml_tbody_type
//
// swagger:route POST /xhtml_tbody_types xhtml_tbody_types postXhtml_tbody_type
//
// Creates a xhtml_tbody_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_tbody_type(c *gin.Context) {

	mutexXhtml_tbody_type.Lock()
	defer mutexXhtml_tbody_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_tbody_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_tbody_type.GetDB()

	// Validate input
	var input orm.Xhtml_tbody_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_tbody_type
	xhtml_tbody_typeDB := orm.Xhtml_tbody_typeDB{}
	xhtml_tbody_typeDB.Xhtml_tbody_typePointersEncoding = input.Xhtml_tbody_typePointersEncoding
	xhtml_tbody_typeDB.CopyBasicFieldsFromXhtml_tbody_type_WOP(&input.Xhtml_tbody_type_WOP)

	query := db.Create(&xhtml_tbody_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_tbody_type.CheckoutPhaseOneInstance(&xhtml_tbody_typeDB)
	xhtml_tbody_type := backRepo.BackRepoXhtml_tbody_type.Map_Xhtml_tbody_typeDBID_Xhtml_tbody_typePtr[xhtml_tbody_typeDB.ID]

	if xhtml_tbody_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_tbody_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_tbody_typeDB)
}

// GetXhtml_tbody_type
//
// swagger:route GET /xhtml_tbody_types/{ID} xhtml_tbody_types getXhtml_tbody_type
//
// Gets the details for a xhtml_tbody_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_tbody_typeDBResponse
func (controller *Controller) GetXhtml_tbody_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_tbody_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_tbody_type.GetDB()

	// Get xhtml_tbody_typeDB in DB
	var xhtml_tbody_typeDB orm.Xhtml_tbody_typeDB
	if err := db.First(&xhtml_tbody_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_tbody_typeAPI orm.Xhtml_tbody_typeAPI
	xhtml_tbody_typeAPI.ID = xhtml_tbody_typeDB.ID
	xhtml_tbody_typeAPI.Xhtml_tbody_typePointersEncoding = xhtml_tbody_typeDB.Xhtml_tbody_typePointersEncoding
	xhtml_tbody_typeDB.CopyBasicFieldsToXhtml_tbody_type_WOP(&xhtml_tbody_typeAPI.Xhtml_tbody_type_WOP)

	c.JSON(http.StatusOK, xhtml_tbody_typeAPI)
}

// UpdateXhtml_tbody_type
//
// swagger:route PATCH /xhtml_tbody_types/{ID} xhtml_tbody_types updateXhtml_tbody_type
//
// # Update a xhtml_tbody_type
//
// Responses:
// default: genericError
//
//	200: xhtml_tbody_typeDBResponse
func (controller *Controller) UpdateXhtml_tbody_type(c *gin.Context) {

	mutexXhtml_tbody_type.Lock()
	defer mutexXhtml_tbody_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_tbody_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_tbody_type.GetDB()

	// Validate input
	var input orm.Xhtml_tbody_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_tbody_typeDB orm.Xhtml_tbody_typeDB

	// fetch the xhtml_tbody_type
	query := db.First(&xhtml_tbody_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_tbody_typeDB.CopyBasicFieldsFromXhtml_tbody_type_WOP(&input.Xhtml_tbody_type_WOP)
	xhtml_tbody_typeDB.Xhtml_tbody_typePointersEncoding = input.Xhtml_tbody_typePointersEncoding

	query = db.Model(&xhtml_tbody_typeDB).Updates(xhtml_tbody_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_tbody_typeNew := new(models.Xhtml_tbody_type)
	xhtml_tbody_typeDB.CopyBasicFieldsToXhtml_tbody_type(xhtml_tbody_typeNew)

	// redeem pointers
	xhtml_tbody_typeDB.DecodePointers(backRepo, xhtml_tbody_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_tbody_typeOld := backRepo.BackRepoXhtml_tbody_type.Map_Xhtml_tbody_typeDBID_Xhtml_tbody_typePtr[xhtml_tbody_typeDB.ID]
	if xhtml_tbody_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_tbody_typeOld, xhtml_tbody_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_tbody_typeDB
	c.JSON(http.StatusOK, xhtml_tbody_typeDB)
}

// DeleteXhtml_tbody_type
//
// swagger:route DELETE /xhtml_tbody_types/{ID} xhtml_tbody_types deleteXhtml_tbody_type
//
// # Delete a xhtml_tbody_type
//
// default: genericError
//
//	200: xhtml_tbody_typeDBResponse
func (controller *Controller) DeleteXhtml_tbody_type(c *gin.Context) {

	mutexXhtml_tbody_type.Lock()
	defer mutexXhtml_tbody_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_tbody_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_tbody_type.GetDB()

	// Get model if exist
	var xhtml_tbody_typeDB orm.Xhtml_tbody_typeDB
	if err := db.First(&xhtml_tbody_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_tbody_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_tbody_typeDeleted := new(models.Xhtml_tbody_type)
	xhtml_tbody_typeDB.CopyBasicFieldsToXhtml_tbody_type(xhtml_tbody_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_tbody_typeStaged := backRepo.BackRepoXhtml_tbody_type.Map_Xhtml_tbody_typeDBID_Xhtml_tbody_typePtr[xhtml_tbody_typeDB.ID]
	if xhtml_tbody_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_tbody_typeStaged, xhtml_tbody_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
