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
var __Xhtml_li_type__dummysDeclaration__ models.Xhtml_li_type
var __Xhtml_li_type_time__dummyDeclaration time.Duration

var mutexXhtml_li_type sync.Mutex

// An Xhtml_li_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_li_type updateXhtml_li_type deleteXhtml_li_type
type Xhtml_li_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_li_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_li_type updateXhtml_li_type
type Xhtml_li_typeInput struct {
	// The Xhtml_li_type to submit or modify
	// in: body
	Xhtml_li_type *orm.Xhtml_li_typeAPI
}

// GetXhtml_li_types
//
// swagger:route GET /xhtml_li_types xhtml_li_types getXhtml_li_types
//
// # Get all xhtml_li_types
//
// Responses:
// default: genericError
//
//	200: xhtml_li_typeDBResponse
func (controller *Controller) GetXhtml_li_types(c *gin.Context) {

	// source slice
	var xhtml_li_typeDBs []orm.Xhtml_li_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_li_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_li_type.GetDB()

	_, err := db.Find(&xhtml_li_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_li_typeAPIs := make([]orm.Xhtml_li_typeAPI, 0)

	// for each xhtml_li_type, update fields from the database nullable fields
	for idx := range xhtml_li_typeDBs {
		xhtml_li_typeDB := &xhtml_li_typeDBs[idx]
		_ = xhtml_li_typeDB
		var xhtml_li_typeAPI orm.Xhtml_li_typeAPI

		// insertion point for updating fields
		xhtml_li_typeAPI.ID = xhtml_li_typeDB.ID
		xhtml_li_typeDB.CopyBasicFieldsToXhtml_li_type_WOP(&xhtml_li_typeAPI.Xhtml_li_type_WOP)
		xhtml_li_typeAPI.Xhtml_li_typePointersEncoding = xhtml_li_typeDB.Xhtml_li_typePointersEncoding
		xhtml_li_typeAPIs = append(xhtml_li_typeAPIs, xhtml_li_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_li_typeAPIs)
}

// PostXhtml_li_type
//
// swagger:route POST /xhtml_li_types xhtml_li_types postXhtml_li_type
//
// Creates a xhtml_li_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_li_type(c *gin.Context) {

	mutexXhtml_li_type.Lock()
	defer mutexXhtml_li_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_li_types", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_li_type.GetDB()

	// Validate input
	var input orm.Xhtml_li_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_li_type
	xhtml_li_typeDB := orm.Xhtml_li_typeDB{}
	xhtml_li_typeDB.Xhtml_li_typePointersEncoding = input.Xhtml_li_typePointersEncoding
	xhtml_li_typeDB.CopyBasicFieldsFromXhtml_li_type_WOP(&input.Xhtml_li_type_WOP)

	_, err = db.Create(&xhtml_li_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_li_type.CheckoutPhaseOneInstance(&xhtml_li_typeDB)
	xhtml_li_type := backRepo.BackRepoXhtml_li_type.Map_Xhtml_li_typeDBID_Xhtml_li_typePtr[xhtml_li_typeDB.ID]

	if xhtml_li_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_li_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_li_typeDB)
}

// GetXhtml_li_type
//
// swagger:route GET /xhtml_li_types/{ID} xhtml_li_types getXhtml_li_type
//
// Gets the details for a xhtml_li_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_li_typeDBResponse
func (controller *Controller) GetXhtml_li_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_li_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_li_type.GetDB()

	// Get xhtml_li_typeDB in DB
	var xhtml_li_typeDB orm.Xhtml_li_typeDB
	if _, err := db.First(&xhtml_li_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_li_typeAPI orm.Xhtml_li_typeAPI
	xhtml_li_typeAPI.ID = xhtml_li_typeDB.ID
	xhtml_li_typeAPI.Xhtml_li_typePointersEncoding = xhtml_li_typeDB.Xhtml_li_typePointersEncoding
	xhtml_li_typeDB.CopyBasicFieldsToXhtml_li_type_WOP(&xhtml_li_typeAPI.Xhtml_li_type_WOP)

	c.JSON(http.StatusOK, xhtml_li_typeAPI)
}

// UpdateXhtml_li_type
//
// swagger:route PATCH /xhtml_li_types/{ID} xhtml_li_types updateXhtml_li_type
//
// # Update a xhtml_li_type
//
// Responses:
// default: genericError
//
//	200: xhtml_li_typeDBResponse
func (controller *Controller) UpdateXhtml_li_type(c *gin.Context) {

	mutexXhtml_li_type.Lock()
	defer mutexXhtml_li_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_li_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_li_type.GetDB()

	// Validate input
	var input orm.Xhtml_li_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_li_typeDB orm.Xhtml_li_typeDB

	// fetch the xhtml_li_type
	_, err := db.First(&xhtml_li_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_li_typeDB.CopyBasicFieldsFromXhtml_li_type_WOP(&input.Xhtml_li_type_WOP)
	xhtml_li_typeDB.Xhtml_li_typePointersEncoding = input.Xhtml_li_typePointersEncoding

	db, _ = db.Model(&xhtml_li_typeDB)
	_, err = db.Updates(&xhtml_li_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_li_typeNew := new(models.Xhtml_li_type)
	xhtml_li_typeDB.CopyBasicFieldsToXhtml_li_type(xhtml_li_typeNew)

	// redeem pointers
	xhtml_li_typeDB.DecodePointers(backRepo, xhtml_li_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_li_typeOld := backRepo.BackRepoXhtml_li_type.Map_Xhtml_li_typeDBID_Xhtml_li_typePtr[xhtml_li_typeDB.ID]
	if xhtml_li_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_li_typeOld, xhtml_li_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_li_typeDB
	c.JSON(http.StatusOK, xhtml_li_typeDB)
}

// DeleteXhtml_li_type
//
// swagger:route DELETE /xhtml_li_types/{ID} xhtml_li_types deleteXhtml_li_type
//
// # Delete a xhtml_li_type
//
// default: genericError
//
//	200: xhtml_li_typeDBResponse
func (controller *Controller) DeleteXhtml_li_type(c *gin.Context) {

	mutexXhtml_li_type.Lock()
	defer mutexXhtml_li_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_li_type", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoXhtml_li_type.GetDB()

	// Get model if exist
	var xhtml_li_typeDB orm.Xhtml_li_typeDB
	if _, err := db.First(&xhtml_li_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_li_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_li_typeDeleted := new(models.Xhtml_li_type)
	xhtml_li_typeDB.CopyBasicFieldsToXhtml_li_type(xhtml_li_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_li_typeStaged := backRepo.BackRepoXhtml_li_type.Map_Xhtml_li_typeDBID_Xhtml_li_typePtr[xhtml_li_typeDB.ID]
	if xhtml_li_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_li_typeStaged, xhtml_li_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
