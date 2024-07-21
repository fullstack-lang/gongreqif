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
var __Xhtml_em_type__dummysDeclaration__ models.Xhtml_em_type
var __Xhtml_em_type_time__dummyDeclaration time.Duration

var mutexXhtml_em_type sync.Mutex

// An Xhtml_em_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_em_type updateXhtml_em_type deleteXhtml_em_type
type Xhtml_em_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_em_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_em_type updateXhtml_em_type
type Xhtml_em_typeInput struct {
	// The Xhtml_em_type to submit or modify
	// in: body
	Xhtml_em_type *orm.Xhtml_em_typeAPI
}

// GetXhtml_em_types
//
// swagger:route GET /xhtml_em_types xhtml_em_types getXhtml_em_types
//
// # Get all xhtml_em_types
//
// Responses:
// default: genericError
//
//	200: xhtml_em_typeDBResponse
func (controller *Controller) GetXhtml_em_types(c *gin.Context) {

	// source slice
	var xhtml_em_typeDBs []orm.Xhtml_em_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_em_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_em_type.GetDB()

	query := db.Find(&xhtml_em_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_em_typeAPIs := make([]orm.Xhtml_em_typeAPI, 0)

	// for each xhtml_em_type, update fields from the database nullable fields
	for idx := range xhtml_em_typeDBs {
		xhtml_em_typeDB := &xhtml_em_typeDBs[idx]
		_ = xhtml_em_typeDB
		var xhtml_em_typeAPI orm.Xhtml_em_typeAPI

		// insertion point for updating fields
		xhtml_em_typeAPI.ID = xhtml_em_typeDB.ID
		xhtml_em_typeDB.CopyBasicFieldsToXhtml_em_type_WOP(&xhtml_em_typeAPI.Xhtml_em_type_WOP)
		xhtml_em_typeAPI.Xhtml_em_typePointersEncoding = xhtml_em_typeDB.Xhtml_em_typePointersEncoding
		xhtml_em_typeAPIs = append(xhtml_em_typeAPIs, xhtml_em_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_em_typeAPIs)
}

// PostXhtml_em_type
//
// swagger:route POST /xhtml_em_types xhtml_em_types postXhtml_em_type
//
// Creates a xhtml_em_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_em_type(c *gin.Context) {

	mutexXhtml_em_type.Lock()
	defer mutexXhtml_em_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_em_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_em_type.GetDB()

	// Validate input
	var input orm.Xhtml_em_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_em_type
	xhtml_em_typeDB := orm.Xhtml_em_typeDB{}
	xhtml_em_typeDB.Xhtml_em_typePointersEncoding = input.Xhtml_em_typePointersEncoding
	xhtml_em_typeDB.CopyBasicFieldsFromXhtml_em_type_WOP(&input.Xhtml_em_type_WOP)

	query := db.Create(&xhtml_em_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_em_type.CheckoutPhaseOneInstance(&xhtml_em_typeDB)
	xhtml_em_type := backRepo.BackRepoXhtml_em_type.Map_Xhtml_em_typeDBID_Xhtml_em_typePtr[xhtml_em_typeDB.ID]

	if xhtml_em_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_em_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_em_typeDB)
}

// GetXhtml_em_type
//
// swagger:route GET /xhtml_em_types/{ID} xhtml_em_types getXhtml_em_type
//
// Gets the details for a xhtml_em_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_em_typeDBResponse
func (controller *Controller) GetXhtml_em_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_em_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_em_type.GetDB()

	// Get xhtml_em_typeDB in DB
	var xhtml_em_typeDB orm.Xhtml_em_typeDB
	if err := db.First(&xhtml_em_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_em_typeAPI orm.Xhtml_em_typeAPI
	xhtml_em_typeAPI.ID = xhtml_em_typeDB.ID
	xhtml_em_typeAPI.Xhtml_em_typePointersEncoding = xhtml_em_typeDB.Xhtml_em_typePointersEncoding
	xhtml_em_typeDB.CopyBasicFieldsToXhtml_em_type_WOP(&xhtml_em_typeAPI.Xhtml_em_type_WOP)

	c.JSON(http.StatusOK, xhtml_em_typeAPI)
}

// UpdateXhtml_em_type
//
// swagger:route PATCH /xhtml_em_types/{ID} xhtml_em_types updateXhtml_em_type
//
// # Update a xhtml_em_type
//
// Responses:
// default: genericError
//
//	200: xhtml_em_typeDBResponse
func (controller *Controller) UpdateXhtml_em_type(c *gin.Context) {

	mutexXhtml_em_type.Lock()
	defer mutexXhtml_em_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_em_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_em_type.GetDB()

	// Validate input
	var input orm.Xhtml_em_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_em_typeDB orm.Xhtml_em_typeDB

	// fetch the xhtml_em_type
	query := db.First(&xhtml_em_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_em_typeDB.CopyBasicFieldsFromXhtml_em_type_WOP(&input.Xhtml_em_type_WOP)
	xhtml_em_typeDB.Xhtml_em_typePointersEncoding = input.Xhtml_em_typePointersEncoding

	query = db.Model(&xhtml_em_typeDB).Updates(xhtml_em_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_em_typeNew := new(models.Xhtml_em_type)
	xhtml_em_typeDB.CopyBasicFieldsToXhtml_em_type(xhtml_em_typeNew)

	// redeem pointers
	xhtml_em_typeDB.DecodePointers(backRepo, xhtml_em_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_em_typeOld := backRepo.BackRepoXhtml_em_type.Map_Xhtml_em_typeDBID_Xhtml_em_typePtr[xhtml_em_typeDB.ID]
	if xhtml_em_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_em_typeOld, xhtml_em_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_em_typeDB
	c.JSON(http.StatusOK, xhtml_em_typeDB)
}

// DeleteXhtml_em_type
//
// swagger:route DELETE /xhtml_em_types/{ID} xhtml_em_types deleteXhtml_em_type
//
// # Delete a xhtml_em_type
//
// default: genericError
//
//	200: xhtml_em_typeDBResponse
func (controller *Controller) DeleteXhtml_em_type(c *gin.Context) {

	mutexXhtml_em_type.Lock()
	defer mutexXhtml_em_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_em_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_em_type.GetDB()

	// Get model if exist
	var xhtml_em_typeDB orm.Xhtml_em_typeDB
	if err := db.First(&xhtml_em_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_em_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_em_typeDeleted := new(models.Xhtml_em_type)
	xhtml_em_typeDB.CopyBasicFieldsToXhtml_em_type(xhtml_em_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_em_typeStaged := backRepo.BackRepoXhtml_em_type.Map_Xhtml_em_typeDBID_Xhtml_em_typePtr[xhtml_em_typeDB.ID]
	if xhtml_em_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_em_typeStaged, xhtml_em_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
