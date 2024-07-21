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
var __Xhtml_dfn_type__dummysDeclaration__ models.Xhtml_dfn_type
var __Xhtml_dfn_type_time__dummyDeclaration time.Duration

var mutexXhtml_dfn_type sync.Mutex

// An Xhtml_dfn_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_dfn_type updateXhtml_dfn_type deleteXhtml_dfn_type
type Xhtml_dfn_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_dfn_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_dfn_type updateXhtml_dfn_type
type Xhtml_dfn_typeInput struct {
	// The Xhtml_dfn_type to submit or modify
	// in: body
	Xhtml_dfn_type *orm.Xhtml_dfn_typeAPI
}

// GetXhtml_dfn_types
//
// swagger:route GET /xhtml_dfn_types xhtml_dfn_types getXhtml_dfn_types
//
// # Get all xhtml_dfn_types
//
// Responses:
// default: genericError
//
//	200: xhtml_dfn_typeDBResponse
func (controller *Controller) GetXhtml_dfn_types(c *gin.Context) {

	// source slice
	var xhtml_dfn_typeDBs []orm.Xhtml_dfn_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_dfn_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dfn_type.GetDB()

	query := db.Find(&xhtml_dfn_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_dfn_typeAPIs := make([]orm.Xhtml_dfn_typeAPI, 0)

	// for each xhtml_dfn_type, update fields from the database nullable fields
	for idx := range xhtml_dfn_typeDBs {
		xhtml_dfn_typeDB := &xhtml_dfn_typeDBs[idx]
		_ = xhtml_dfn_typeDB
		var xhtml_dfn_typeAPI orm.Xhtml_dfn_typeAPI

		// insertion point for updating fields
		xhtml_dfn_typeAPI.ID = xhtml_dfn_typeDB.ID
		xhtml_dfn_typeDB.CopyBasicFieldsToXhtml_dfn_type_WOP(&xhtml_dfn_typeAPI.Xhtml_dfn_type_WOP)
		xhtml_dfn_typeAPI.Xhtml_dfn_typePointersEncoding = xhtml_dfn_typeDB.Xhtml_dfn_typePointersEncoding
		xhtml_dfn_typeAPIs = append(xhtml_dfn_typeAPIs, xhtml_dfn_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_dfn_typeAPIs)
}

// PostXhtml_dfn_type
//
// swagger:route POST /xhtml_dfn_types xhtml_dfn_types postXhtml_dfn_type
//
// Creates a xhtml_dfn_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_dfn_type(c *gin.Context) {

	mutexXhtml_dfn_type.Lock()
	defer mutexXhtml_dfn_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_dfn_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dfn_type.GetDB()

	// Validate input
	var input orm.Xhtml_dfn_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_dfn_type
	xhtml_dfn_typeDB := orm.Xhtml_dfn_typeDB{}
	xhtml_dfn_typeDB.Xhtml_dfn_typePointersEncoding = input.Xhtml_dfn_typePointersEncoding
	xhtml_dfn_typeDB.CopyBasicFieldsFromXhtml_dfn_type_WOP(&input.Xhtml_dfn_type_WOP)

	query := db.Create(&xhtml_dfn_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_dfn_type.CheckoutPhaseOneInstance(&xhtml_dfn_typeDB)
	xhtml_dfn_type := backRepo.BackRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr[xhtml_dfn_typeDB.ID]

	if xhtml_dfn_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_dfn_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_dfn_typeDB)
}

// GetXhtml_dfn_type
//
// swagger:route GET /xhtml_dfn_types/{ID} xhtml_dfn_types getXhtml_dfn_type
//
// Gets the details for a xhtml_dfn_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_dfn_typeDBResponse
func (controller *Controller) GetXhtml_dfn_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_dfn_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dfn_type.GetDB()

	// Get xhtml_dfn_typeDB in DB
	var xhtml_dfn_typeDB orm.Xhtml_dfn_typeDB
	if err := db.First(&xhtml_dfn_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_dfn_typeAPI orm.Xhtml_dfn_typeAPI
	xhtml_dfn_typeAPI.ID = xhtml_dfn_typeDB.ID
	xhtml_dfn_typeAPI.Xhtml_dfn_typePointersEncoding = xhtml_dfn_typeDB.Xhtml_dfn_typePointersEncoding
	xhtml_dfn_typeDB.CopyBasicFieldsToXhtml_dfn_type_WOP(&xhtml_dfn_typeAPI.Xhtml_dfn_type_WOP)

	c.JSON(http.StatusOK, xhtml_dfn_typeAPI)
}

// UpdateXhtml_dfn_type
//
// swagger:route PATCH /xhtml_dfn_types/{ID} xhtml_dfn_types updateXhtml_dfn_type
//
// # Update a xhtml_dfn_type
//
// Responses:
// default: genericError
//
//	200: xhtml_dfn_typeDBResponse
func (controller *Controller) UpdateXhtml_dfn_type(c *gin.Context) {

	mutexXhtml_dfn_type.Lock()
	defer mutexXhtml_dfn_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_dfn_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dfn_type.GetDB()

	// Validate input
	var input orm.Xhtml_dfn_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_dfn_typeDB orm.Xhtml_dfn_typeDB

	// fetch the xhtml_dfn_type
	query := db.First(&xhtml_dfn_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_dfn_typeDB.CopyBasicFieldsFromXhtml_dfn_type_WOP(&input.Xhtml_dfn_type_WOP)
	xhtml_dfn_typeDB.Xhtml_dfn_typePointersEncoding = input.Xhtml_dfn_typePointersEncoding

	query = db.Model(&xhtml_dfn_typeDB).Updates(xhtml_dfn_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_dfn_typeNew := new(models.Xhtml_dfn_type)
	xhtml_dfn_typeDB.CopyBasicFieldsToXhtml_dfn_type(xhtml_dfn_typeNew)

	// redeem pointers
	xhtml_dfn_typeDB.DecodePointers(backRepo, xhtml_dfn_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_dfn_typeOld := backRepo.BackRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr[xhtml_dfn_typeDB.ID]
	if xhtml_dfn_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_dfn_typeOld, xhtml_dfn_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_dfn_typeDB
	c.JSON(http.StatusOK, xhtml_dfn_typeDB)
}

// DeleteXhtml_dfn_type
//
// swagger:route DELETE /xhtml_dfn_types/{ID} xhtml_dfn_types deleteXhtml_dfn_type
//
// # Delete a xhtml_dfn_type
//
// default: genericError
//
//	200: xhtml_dfn_typeDBResponse
func (controller *Controller) DeleteXhtml_dfn_type(c *gin.Context) {

	mutexXhtml_dfn_type.Lock()
	defer mutexXhtml_dfn_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_dfn_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_dfn_type.GetDB()

	// Get model if exist
	var xhtml_dfn_typeDB orm.Xhtml_dfn_typeDB
	if err := db.First(&xhtml_dfn_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_dfn_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_dfn_typeDeleted := new(models.Xhtml_dfn_type)
	xhtml_dfn_typeDB.CopyBasicFieldsToXhtml_dfn_type(xhtml_dfn_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_dfn_typeStaged := backRepo.BackRepoXhtml_dfn_type.Map_Xhtml_dfn_typeDBID_Xhtml_dfn_typePtr[xhtml_dfn_typeDB.ID]
	if xhtml_dfn_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_dfn_typeStaged, xhtml_dfn_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
