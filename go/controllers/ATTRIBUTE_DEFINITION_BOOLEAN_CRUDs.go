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
var __ATTRIBUTE_DEFINITION_BOOLEAN__dummysDeclaration__ models.ATTRIBUTE_DEFINITION_BOOLEAN
var __ATTRIBUTE_DEFINITION_BOOLEAN_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_DEFINITION_BOOLEAN sync.Mutex

// An ATTRIBUTE_DEFINITION_BOOLEANID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_DEFINITION_BOOLEAN updateATTRIBUTE_DEFINITION_BOOLEAN deleteATTRIBUTE_DEFINITION_BOOLEAN
type ATTRIBUTE_DEFINITION_BOOLEANID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_DEFINITION_BOOLEANInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_DEFINITION_BOOLEAN updateATTRIBUTE_DEFINITION_BOOLEAN
type ATTRIBUTE_DEFINITION_BOOLEANInput struct {
	// The ATTRIBUTE_DEFINITION_BOOLEAN to submit or modify
	// in: body
	ATTRIBUTE_DEFINITION_BOOLEAN *orm.ATTRIBUTE_DEFINITION_BOOLEANAPI
}

// GetATTRIBUTE_DEFINITION_BOOLEANs
//
// swagger:route GET /attribute_definition_booleans attribute_definition_booleans getATTRIBUTE_DEFINITION_BOOLEANs
//
// # Get all attribute_definition_booleans
//
// Responses:
// default: genericError
//
//	200: attribute_definition_booleanDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_BOOLEANs(c *gin.Context) {

	// source slice
	var attribute_definition_booleanDBs []orm.ATTRIBUTE_DEFINITION_BOOLEANDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_BOOLEANs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.GetDB()

	_, err := db.Find(&attribute_definition_booleanDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_definition_booleanAPIs := make([]orm.ATTRIBUTE_DEFINITION_BOOLEANAPI, 0)

	// for each attribute_definition_boolean, update fields from the database nullable fields
	for idx := range attribute_definition_booleanDBs {
		attribute_definition_booleanDB := &attribute_definition_booleanDBs[idx]
		_ = attribute_definition_booleanDB
		var attribute_definition_booleanAPI orm.ATTRIBUTE_DEFINITION_BOOLEANAPI

		// insertion point for updating fields
		attribute_definition_booleanAPI.ID = attribute_definition_booleanDB.ID
		attribute_definition_booleanDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN_WOP(&attribute_definition_booleanAPI.ATTRIBUTE_DEFINITION_BOOLEAN_WOP)
		attribute_definition_booleanAPI.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding = attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding
		attribute_definition_booleanAPIs = append(attribute_definition_booleanAPIs, attribute_definition_booleanAPI)
	}

	c.JSON(http.StatusOK, attribute_definition_booleanAPIs)
}

// PostATTRIBUTE_DEFINITION_BOOLEAN
//
// swagger:route POST /attribute_definition_booleans attribute_definition_booleans postATTRIBUTE_DEFINITION_BOOLEAN
//
// Creates a attribute_definition_boolean
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_DEFINITION_BOOLEAN(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_BOOLEAN.Lock()
	defer mutexATTRIBUTE_DEFINITION_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_DEFINITION_BOOLEANs", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_BOOLEANAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_definition_boolean
	attribute_definition_booleanDB := orm.ATTRIBUTE_DEFINITION_BOOLEANDB{}
	attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding = input.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding
	attribute_definition_booleanDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEAN_WOP(&input.ATTRIBUTE_DEFINITION_BOOLEAN_WOP)

	_, err = db.Create(&attribute_definition_booleanDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseOneInstance(&attribute_definition_booleanDB)
	attribute_definition_boolean := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr[attribute_definition_booleanDB.ID]

	if attribute_definition_boolean != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_definition_boolean)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_definition_booleanDB)
}

// GetATTRIBUTE_DEFINITION_BOOLEAN
//
// swagger:route GET /attribute_definition_booleans/{ID} attribute_definition_booleans getATTRIBUTE_DEFINITION_BOOLEAN
//
// Gets the details for a attribute_definition_boolean.
//
// Responses:
// default: genericError
//
//	200: attribute_definition_booleanDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_BOOLEAN(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_BOOLEAN", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.GetDB()

	// Get attribute_definition_booleanDB in DB
	var attribute_definition_booleanDB orm.ATTRIBUTE_DEFINITION_BOOLEANDB
	if _, err := db.First(&attribute_definition_booleanDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_definition_booleanAPI orm.ATTRIBUTE_DEFINITION_BOOLEANAPI
	attribute_definition_booleanAPI.ID = attribute_definition_booleanDB.ID
	attribute_definition_booleanAPI.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding = attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding
	attribute_definition_booleanDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN_WOP(&attribute_definition_booleanAPI.ATTRIBUTE_DEFINITION_BOOLEAN_WOP)

	c.JSON(http.StatusOK, attribute_definition_booleanAPI)
}

// UpdateATTRIBUTE_DEFINITION_BOOLEAN
//
// swagger:route PATCH /attribute_definition_booleans/{ID} attribute_definition_booleans updateATTRIBUTE_DEFINITION_BOOLEAN
//
// # Update a attribute_definition_boolean
//
// Responses:
// default: genericError
//
//	200: attribute_definition_booleanDBResponse
func (controller *Controller) UpdateATTRIBUTE_DEFINITION_BOOLEAN(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_BOOLEAN.Lock()
	defer mutexATTRIBUTE_DEFINITION_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTE_DEFINITION_BOOLEAN", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_BOOLEANAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_definition_booleanDB orm.ATTRIBUTE_DEFINITION_BOOLEANDB

	// fetch the attribute_definition_boolean
	_, err := db.First(&attribute_definition_booleanDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_definition_booleanDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_BOOLEAN_WOP(&input.ATTRIBUTE_DEFINITION_BOOLEAN_WOP)
	attribute_definition_booleanDB.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding = input.ATTRIBUTE_DEFINITION_BOOLEANPointersEncoding

	db, _ = db.Model(&attribute_definition_booleanDB)
	_, err = db.Updates(&attribute_definition_booleanDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_booleanNew := new(models.ATTRIBUTE_DEFINITION_BOOLEAN)
	attribute_definition_booleanDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_booleanNew)

	// redeem pointers
	attribute_definition_booleanDB.DecodePointers(backRepo, attribute_definition_booleanNew)

	// get stage instance from DB instance, and call callback function
	attribute_definition_booleanOld := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr[attribute_definition_booleanDB.ID]
	if attribute_definition_booleanOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attribute_definition_booleanOld, attribute_definition_booleanNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_definition_booleanDB
	c.JSON(http.StatusOK, attribute_definition_booleanDB)
}

// DeleteATTRIBUTE_DEFINITION_BOOLEAN
//
// swagger:route DELETE /attribute_definition_booleans/{ID} attribute_definition_booleans deleteATTRIBUTE_DEFINITION_BOOLEAN
//
// # Delete a attribute_definition_boolean
//
// default: genericError
//
//	200: attribute_definition_booleanDBResponse
func (controller *Controller) DeleteATTRIBUTE_DEFINITION_BOOLEAN(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_BOOLEAN.Lock()
	defer mutexATTRIBUTE_DEFINITION_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_DEFINITION_BOOLEAN", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gongreqif/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.GetDB()

	// Get model if exist
	var attribute_definition_booleanDB orm.ATTRIBUTE_DEFINITION_BOOLEANDB
	if _, err := db.First(&attribute_definition_booleanDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_definition_booleanDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_booleanDeleted := new(models.ATTRIBUTE_DEFINITION_BOOLEAN)
	attribute_definition_booleanDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_booleanDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_definition_booleanStaged := backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr[attribute_definition_booleanDB.ID]
	if attribute_definition_booleanStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_definition_booleanStaged, attribute_definition_booleanDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
