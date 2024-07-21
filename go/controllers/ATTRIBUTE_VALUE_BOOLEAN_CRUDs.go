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
var __ATTRIBUTE_VALUE_BOOLEAN__dummysDeclaration__ models.ATTRIBUTE_VALUE_BOOLEAN
var __ATTRIBUTE_VALUE_BOOLEAN_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_VALUE_BOOLEAN sync.Mutex

// An ATTRIBUTE_VALUE_BOOLEANID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_VALUE_BOOLEAN updateATTRIBUTE_VALUE_BOOLEAN deleteATTRIBUTE_VALUE_BOOLEAN
type ATTRIBUTE_VALUE_BOOLEANID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_VALUE_BOOLEANInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_VALUE_BOOLEAN updateATTRIBUTE_VALUE_BOOLEAN
type ATTRIBUTE_VALUE_BOOLEANInput struct {
	// The ATTRIBUTE_VALUE_BOOLEAN to submit or modify
	// in: body
	ATTRIBUTE_VALUE_BOOLEAN *orm.ATTRIBUTE_VALUE_BOOLEANAPI
}

// GetATTRIBUTE_VALUE_BOOLEANs
//
// swagger:route GET /attribute_value_booleans attribute_value_booleans getATTRIBUTE_VALUE_BOOLEANs
//
// # Get all attribute_value_booleans
//
// Responses:
// default: genericError
//
//	200: attribute_value_booleanDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_BOOLEANs(c *gin.Context) {

	// source slice
	var attribute_value_booleanDBs []orm.ATTRIBUTE_VALUE_BOOLEANDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_BOOLEANs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.GetDB()

	query := db.Find(&attribute_value_booleanDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_value_booleanAPIs := make([]orm.ATTRIBUTE_VALUE_BOOLEANAPI, 0)

	// for each attribute_value_boolean, update fields from the database nullable fields
	for idx := range attribute_value_booleanDBs {
		attribute_value_booleanDB := &attribute_value_booleanDBs[idx]
		_ = attribute_value_booleanDB
		var attribute_value_booleanAPI orm.ATTRIBUTE_VALUE_BOOLEANAPI

		// insertion point for updating fields
		attribute_value_booleanAPI.ID = attribute_value_booleanDB.ID
		attribute_value_booleanDB.CopyBasicFieldsToATTRIBUTE_VALUE_BOOLEAN_WOP(&attribute_value_booleanAPI.ATTRIBUTE_VALUE_BOOLEAN_WOP)
		attribute_value_booleanAPI.ATTRIBUTE_VALUE_BOOLEANPointersEncoding = attribute_value_booleanDB.ATTRIBUTE_VALUE_BOOLEANPointersEncoding
		attribute_value_booleanAPIs = append(attribute_value_booleanAPIs, attribute_value_booleanAPI)
	}

	c.JSON(http.StatusOK, attribute_value_booleanAPIs)
}

// PostATTRIBUTE_VALUE_BOOLEAN
//
// swagger:route POST /attribute_value_booleans attribute_value_booleans postATTRIBUTE_VALUE_BOOLEAN
//
// Creates a attribute_value_boolean
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_VALUE_BOOLEAN(c *gin.Context) {

	mutexATTRIBUTE_VALUE_BOOLEAN.Lock()
	defer mutexATTRIBUTE_VALUE_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_VALUE_BOOLEANs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_BOOLEANAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_value_boolean
	attribute_value_booleanDB := orm.ATTRIBUTE_VALUE_BOOLEANDB{}
	attribute_value_booleanDB.ATTRIBUTE_VALUE_BOOLEANPointersEncoding = input.ATTRIBUTE_VALUE_BOOLEANPointersEncoding
	attribute_value_booleanDB.CopyBasicFieldsFromATTRIBUTE_VALUE_BOOLEAN_WOP(&input.ATTRIBUTE_VALUE_BOOLEAN_WOP)

	query := db.Create(&attribute_value_booleanDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.CheckoutPhaseOneInstance(&attribute_value_booleanDB)
	attribute_value_boolean := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANPtr[attribute_value_booleanDB.ID]

	if attribute_value_boolean != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_value_boolean)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_value_booleanDB)
}

// GetATTRIBUTE_VALUE_BOOLEAN
//
// swagger:route GET /attribute_value_booleans/{ID} attribute_value_booleans getATTRIBUTE_VALUE_BOOLEAN
//
// Gets the details for a attribute_value_boolean.
//
// Responses:
// default: genericError
//
//	200: attribute_value_booleanDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_BOOLEAN(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_BOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.GetDB()

	// Get attribute_value_booleanDB in DB
	var attribute_value_booleanDB orm.ATTRIBUTE_VALUE_BOOLEANDB
	if err := db.First(&attribute_value_booleanDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_value_booleanAPI orm.ATTRIBUTE_VALUE_BOOLEANAPI
	attribute_value_booleanAPI.ID = attribute_value_booleanDB.ID
	attribute_value_booleanAPI.ATTRIBUTE_VALUE_BOOLEANPointersEncoding = attribute_value_booleanDB.ATTRIBUTE_VALUE_BOOLEANPointersEncoding
	attribute_value_booleanDB.CopyBasicFieldsToATTRIBUTE_VALUE_BOOLEAN_WOP(&attribute_value_booleanAPI.ATTRIBUTE_VALUE_BOOLEAN_WOP)

	c.JSON(http.StatusOK, attribute_value_booleanAPI)
}

// UpdateATTRIBUTE_VALUE_BOOLEAN
//
// swagger:route PATCH /attribute_value_booleans/{ID} attribute_value_booleans updateATTRIBUTE_VALUE_BOOLEAN
//
// # Update a attribute_value_boolean
//
// Responses:
// default: genericError
//
//	200: attribute_value_booleanDBResponse
func (controller *Controller) UpdateATTRIBUTE_VALUE_BOOLEAN(c *gin.Context) {

	mutexATTRIBUTE_VALUE_BOOLEAN.Lock()
	defer mutexATTRIBUTE_VALUE_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTE_VALUE_BOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_BOOLEANAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_value_booleanDB orm.ATTRIBUTE_VALUE_BOOLEANDB

	// fetch the attribute_value_boolean
	query := db.First(&attribute_value_booleanDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_value_booleanDB.CopyBasicFieldsFromATTRIBUTE_VALUE_BOOLEAN_WOP(&input.ATTRIBUTE_VALUE_BOOLEAN_WOP)
	attribute_value_booleanDB.ATTRIBUTE_VALUE_BOOLEANPointersEncoding = input.ATTRIBUTE_VALUE_BOOLEANPointersEncoding

	query = db.Model(&attribute_value_booleanDB).Updates(attribute_value_booleanDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_booleanNew := new(models.ATTRIBUTE_VALUE_BOOLEAN)
	attribute_value_booleanDB.CopyBasicFieldsToATTRIBUTE_VALUE_BOOLEAN(attribute_value_booleanNew)

	// redeem pointers
	attribute_value_booleanDB.DecodePointers(backRepo, attribute_value_booleanNew)

	// get stage instance from DB instance, and call callback function
	attribute_value_booleanOld := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANPtr[attribute_value_booleanDB.ID]
	if attribute_value_booleanOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attribute_value_booleanOld, attribute_value_booleanNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_value_booleanDB
	c.JSON(http.StatusOK, attribute_value_booleanDB)
}

// DeleteATTRIBUTE_VALUE_BOOLEAN
//
// swagger:route DELETE /attribute_value_booleans/{ID} attribute_value_booleans deleteATTRIBUTE_VALUE_BOOLEAN
//
// # Delete a attribute_value_boolean
//
// default: genericError
//
//	200: attribute_value_booleanDBResponse
func (controller *Controller) DeleteATTRIBUTE_VALUE_BOOLEAN(c *gin.Context) {

	mutexATTRIBUTE_VALUE_BOOLEAN.Lock()
	defer mutexATTRIBUTE_VALUE_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_VALUE_BOOLEAN", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.GetDB()

	// Get model if exist
	var attribute_value_booleanDB orm.ATTRIBUTE_VALUE_BOOLEANDB
	if err := db.First(&attribute_value_booleanDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attribute_value_booleanDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_booleanDeleted := new(models.ATTRIBUTE_VALUE_BOOLEAN)
	attribute_value_booleanDB.CopyBasicFieldsToATTRIBUTE_VALUE_BOOLEAN(attribute_value_booleanDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_value_booleanStaged := backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANPtr[attribute_value_booleanDB.ID]
	if attribute_value_booleanStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_value_booleanStaged, attribute_value_booleanDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
