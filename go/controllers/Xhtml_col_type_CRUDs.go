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
var __Xhtml_col_type__dummysDeclaration__ models.Xhtml_col_type
var __Xhtml_col_type_time__dummyDeclaration time.Duration

var mutexXhtml_col_type sync.Mutex

// An Xhtml_col_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_col_type updateXhtml_col_type deleteXhtml_col_type
type Xhtml_col_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_col_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_col_type updateXhtml_col_type
type Xhtml_col_typeInput struct {
	// The Xhtml_col_type to submit or modify
	// in: body
	Xhtml_col_type *orm.Xhtml_col_typeAPI
}

// GetXhtml_col_types
//
// swagger:route GET /xhtml_col_types xhtml_col_types getXhtml_col_types
//
// # Get all xhtml_col_types
//
// Responses:
// default: genericError
//
//	200: xhtml_col_typeDBResponse
func (controller *Controller) GetXhtml_col_types(c *gin.Context) {

	// source slice
	var xhtml_col_typeDBs []orm.Xhtml_col_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_col_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_col_type.GetDB()

	query := db.Find(&xhtml_col_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_col_typeAPIs := make([]orm.Xhtml_col_typeAPI, 0)

	// for each xhtml_col_type, update fields from the database nullable fields
	for idx := range xhtml_col_typeDBs {
		xhtml_col_typeDB := &xhtml_col_typeDBs[idx]
		_ = xhtml_col_typeDB
		var xhtml_col_typeAPI orm.Xhtml_col_typeAPI

		// insertion point for updating fields
		xhtml_col_typeAPI.ID = xhtml_col_typeDB.ID
		xhtml_col_typeDB.CopyBasicFieldsToXhtml_col_type_WOP(&xhtml_col_typeAPI.Xhtml_col_type_WOP)
		xhtml_col_typeAPI.Xhtml_col_typePointersEncoding = xhtml_col_typeDB.Xhtml_col_typePointersEncoding
		xhtml_col_typeAPIs = append(xhtml_col_typeAPIs, xhtml_col_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_col_typeAPIs)
}

// PostXhtml_col_type
//
// swagger:route POST /xhtml_col_types xhtml_col_types postXhtml_col_type
//
// Creates a xhtml_col_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_col_type(c *gin.Context) {

	mutexXhtml_col_type.Lock()
	defer mutexXhtml_col_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_col_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_col_type.GetDB()

	// Validate input
	var input orm.Xhtml_col_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_col_type
	xhtml_col_typeDB := orm.Xhtml_col_typeDB{}
	xhtml_col_typeDB.Xhtml_col_typePointersEncoding = input.Xhtml_col_typePointersEncoding
	xhtml_col_typeDB.CopyBasicFieldsFromXhtml_col_type_WOP(&input.Xhtml_col_type_WOP)

	query := db.Create(&xhtml_col_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_col_type.CheckoutPhaseOneInstance(&xhtml_col_typeDB)
	xhtml_col_type := backRepo.BackRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr[xhtml_col_typeDB.ID]

	if xhtml_col_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_col_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_col_typeDB)
}

// GetXhtml_col_type
//
// swagger:route GET /xhtml_col_types/{ID} xhtml_col_types getXhtml_col_type
//
// Gets the details for a xhtml_col_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_col_typeDBResponse
func (controller *Controller) GetXhtml_col_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_col_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_col_type.GetDB()

	// Get xhtml_col_typeDB in DB
	var xhtml_col_typeDB orm.Xhtml_col_typeDB
	if err := db.First(&xhtml_col_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_col_typeAPI orm.Xhtml_col_typeAPI
	xhtml_col_typeAPI.ID = xhtml_col_typeDB.ID
	xhtml_col_typeAPI.Xhtml_col_typePointersEncoding = xhtml_col_typeDB.Xhtml_col_typePointersEncoding
	xhtml_col_typeDB.CopyBasicFieldsToXhtml_col_type_WOP(&xhtml_col_typeAPI.Xhtml_col_type_WOP)

	c.JSON(http.StatusOK, xhtml_col_typeAPI)
}

// UpdateXhtml_col_type
//
// swagger:route PATCH /xhtml_col_types/{ID} xhtml_col_types updateXhtml_col_type
//
// # Update a xhtml_col_type
//
// Responses:
// default: genericError
//
//	200: xhtml_col_typeDBResponse
func (controller *Controller) UpdateXhtml_col_type(c *gin.Context) {

	mutexXhtml_col_type.Lock()
	defer mutexXhtml_col_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_col_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_col_type.GetDB()

	// Validate input
	var input orm.Xhtml_col_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_col_typeDB orm.Xhtml_col_typeDB

	// fetch the xhtml_col_type
	query := db.First(&xhtml_col_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_col_typeDB.CopyBasicFieldsFromXhtml_col_type_WOP(&input.Xhtml_col_type_WOP)
	xhtml_col_typeDB.Xhtml_col_typePointersEncoding = input.Xhtml_col_typePointersEncoding

	query = db.Model(&xhtml_col_typeDB).Updates(xhtml_col_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_col_typeNew := new(models.Xhtml_col_type)
	xhtml_col_typeDB.CopyBasicFieldsToXhtml_col_type(xhtml_col_typeNew)

	// redeem pointers
	xhtml_col_typeDB.DecodePointers(backRepo, xhtml_col_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_col_typeOld := backRepo.BackRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr[xhtml_col_typeDB.ID]
	if xhtml_col_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_col_typeOld, xhtml_col_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_col_typeDB
	c.JSON(http.StatusOK, xhtml_col_typeDB)
}

// DeleteXhtml_col_type
//
// swagger:route DELETE /xhtml_col_types/{ID} xhtml_col_types deleteXhtml_col_type
//
// # Delete a xhtml_col_type
//
// default: genericError
//
//	200: xhtml_col_typeDBResponse
func (controller *Controller) DeleteXhtml_col_type(c *gin.Context) {

	mutexXhtml_col_type.Lock()
	defer mutexXhtml_col_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_col_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_col_type.GetDB()

	// Get model if exist
	var xhtml_col_typeDB orm.Xhtml_col_typeDB
	if err := db.First(&xhtml_col_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_col_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_col_typeDeleted := new(models.Xhtml_col_type)
	xhtml_col_typeDB.CopyBasicFieldsToXhtml_col_type(xhtml_col_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_col_typeStaged := backRepo.BackRepoXhtml_col_type.Map_Xhtml_col_typeDBID_Xhtml_col_typePtr[xhtml_col_typeDB.ID]
	if xhtml_col_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_col_typeStaged, xhtml_col_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
