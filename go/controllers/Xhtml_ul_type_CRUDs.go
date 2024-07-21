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
var __Xhtml_ul_type__dummysDeclaration__ models.Xhtml_ul_type
var __Xhtml_ul_type_time__dummyDeclaration time.Duration

var mutexXhtml_ul_type sync.Mutex

// An Xhtml_ul_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_ul_type updateXhtml_ul_type deleteXhtml_ul_type
type Xhtml_ul_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_ul_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_ul_type updateXhtml_ul_type
type Xhtml_ul_typeInput struct {
	// The Xhtml_ul_type to submit or modify
	// in: body
	Xhtml_ul_type *orm.Xhtml_ul_typeAPI
}

// GetXhtml_ul_types
//
// swagger:route GET /xhtml_ul_types xhtml_ul_types getXhtml_ul_types
//
// # Get all xhtml_ul_types
//
// Responses:
// default: genericError
//
//	200: xhtml_ul_typeDBResponse
func (controller *Controller) GetXhtml_ul_types(c *gin.Context) {

	// source slice
	var xhtml_ul_typeDBs []orm.Xhtml_ul_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_ul_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ul_type.GetDB()

	query := db.Find(&xhtml_ul_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_ul_typeAPIs := make([]orm.Xhtml_ul_typeAPI, 0)

	// for each xhtml_ul_type, update fields from the database nullable fields
	for idx := range xhtml_ul_typeDBs {
		xhtml_ul_typeDB := &xhtml_ul_typeDBs[idx]
		_ = xhtml_ul_typeDB
		var xhtml_ul_typeAPI orm.Xhtml_ul_typeAPI

		// insertion point for updating fields
		xhtml_ul_typeAPI.ID = xhtml_ul_typeDB.ID
		xhtml_ul_typeDB.CopyBasicFieldsToXhtml_ul_type_WOP(&xhtml_ul_typeAPI.Xhtml_ul_type_WOP)
		xhtml_ul_typeAPI.Xhtml_ul_typePointersEncoding = xhtml_ul_typeDB.Xhtml_ul_typePointersEncoding
		xhtml_ul_typeAPIs = append(xhtml_ul_typeAPIs, xhtml_ul_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_ul_typeAPIs)
}

// PostXhtml_ul_type
//
// swagger:route POST /xhtml_ul_types xhtml_ul_types postXhtml_ul_type
//
// Creates a xhtml_ul_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_ul_type(c *gin.Context) {

	mutexXhtml_ul_type.Lock()
	defer mutexXhtml_ul_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_ul_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ul_type.GetDB()

	// Validate input
	var input orm.Xhtml_ul_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_ul_type
	xhtml_ul_typeDB := orm.Xhtml_ul_typeDB{}
	xhtml_ul_typeDB.Xhtml_ul_typePointersEncoding = input.Xhtml_ul_typePointersEncoding
	xhtml_ul_typeDB.CopyBasicFieldsFromXhtml_ul_type_WOP(&input.Xhtml_ul_type_WOP)

	query := db.Create(&xhtml_ul_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_ul_type.CheckoutPhaseOneInstance(&xhtml_ul_typeDB)
	xhtml_ul_type := backRepo.BackRepoXhtml_ul_type.Map_Xhtml_ul_typeDBID_Xhtml_ul_typePtr[xhtml_ul_typeDB.ID]

	if xhtml_ul_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_ul_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_ul_typeDB)
}

// GetXhtml_ul_type
//
// swagger:route GET /xhtml_ul_types/{ID} xhtml_ul_types getXhtml_ul_type
//
// Gets the details for a xhtml_ul_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_ul_typeDBResponse
func (controller *Controller) GetXhtml_ul_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_ul_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ul_type.GetDB()

	// Get xhtml_ul_typeDB in DB
	var xhtml_ul_typeDB orm.Xhtml_ul_typeDB
	if err := db.First(&xhtml_ul_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_ul_typeAPI orm.Xhtml_ul_typeAPI
	xhtml_ul_typeAPI.ID = xhtml_ul_typeDB.ID
	xhtml_ul_typeAPI.Xhtml_ul_typePointersEncoding = xhtml_ul_typeDB.Xhtml_ul_typePointersEncoding
	xhtml_ul_typeDB.CopyBasicFieldsToXhtml_ul_type_WOP(&xhtml_ul_typeAPI.Xhtml_ul_type_WOP)

	c.JSON(http.StatusOK, xhtml_ul_typeAPI)
}

// UpdateXhtml_ul_type
//
// swagger:route PATCH /xhtml_ul_types/{ID} xhtml_ul_types updateXhtml_ul_type
//
// # Update a xhtml_ul_type
//
// Responses:
// default: genericError
//
//	200: xhtml_ul_typeDBResponse
func (controller *Controller) UpdateXhtml_ul_type(c *gin.Context) {

	mutexXhtml_ul_type.Lock()
	defer mutexXhtml_ul_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_ul_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ul_type.GetDB()

	// Validate input
	var input orm.Xhtml_ul_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_ul_typeDB orm.Xhtml_ul_typeDB

	// fetch the xhtml_ul_type
	query := db.First(&xhtml_ul_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_ul_typeDB.CopyBasicFieldsFromXhtml_ul_type_WOP(&input.Xhtml_ul_type_WOP)
	xhtml_ul_typeDB.Xhtml_ul_typePointersEncoding = input.Xhtml_ul_typePointersEncoding

	query = db.Model(&xhtml_ul_typeDB).Updates(xhtml_ul_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_ul_typeNew := new(models.Xhtml_ul_type)
	xhtml_ul_typeDB.CopyBasicFieldsToXhtml_ul_type(xhtml_ul_typeNew)

	// redeem pointers
	xhtml_ul_typeDB.DecodePointers(backRepo, xhtml_ul_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_ul_typeOld := backRepo.BackRepoXhtml_ul_type.Map_Xhtml_ul_typeDBID_Xhtml_ul_typePtr[xhtml_ul_typeDB.ID]
	if xhtml_ul_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_ul_typeOld, xhtml_ul_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_ul_typeDB
	c.JSON(http.StatusOK, xhtml_ul_typeDB)
}

// DeleteXhtml_ul_type
//
// swagger:route DELETE /xhtml_ul_types/{ID} xhtml_ul_types deleteXhtml_ul_type
//
// # Delete a xhtml_ul_type
//
// default: genericError
//
//	200: xhtml_ul_typeDBResponse
func (controller *Controller) DeleteXhtml_ul_type(c *gin.Context) {

	mutexXhtml_ul_type.Lock()
	defer mutexXhtml_ul_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_ul_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_ul_type.GetDB()

	// Get model if exist
	var xhtml_ul_typeDB orm.Xhtml_ul_typeDB
	if err := db.First(&xhtml_ul_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&xhtml_ul_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_ul_typeDeleted := new(models.Xhtml_ul_type)
	xhtml_ul_typeDB.CopyBasicFieldsToXhtml_ul_type(xhtml_ul_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_ul_typeStaged := backRepo.BackRepoXhtml_ul_type.Map_Xhtml_ul_typeDBID_Xhtml_ul_typePtr[xhtml_ul_typeDB.ID]
	if xhtml_ul_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_ul_typeStaged, xhtml_ul_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
