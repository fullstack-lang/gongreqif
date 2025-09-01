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
var __ATTRIBUTE_DEFINITION_INTEGER__dummysDeclaration__ models.ATTRIBUTE_DEFINITION_INTEGER
var __ATTRIBUTE_DEFINITION_INTEGER_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_DEFINITION_INTEGER sync.Mutex

// An ATTRIBUTE_DEFINITION_INTEGERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_DEFINITION_INTEGER updateATTRIBUTE_DEFINITION_INTEGER deleteATTRIBUTE_DEFINITION_INTEGER
type ATTRIBUTE_DEFINITION_INTEGERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_DEFINITION_INTEGERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_DEFINITION_INTEGER updateATTRIBUTE_DEFINITION_INTEGER
type ATTRIBUTE_DEFINITION_INTEGERInput struct {
	// The ATTRIBUTE_DEFINITION_INTEGER to submit or modify
	// in: body
	ATTRIBUTE_DEFINITION_INTEGER *orm.ATTRIBUTE_DEFINITION_INTEGERAPI
}

// GetATTRIBUTE_DEFINITION_INTEGERs
//
// swagger:route GET /attribute_definition_integers attribute_definition_integers getATTRIBUTE_DEFINITION_INTEGERs
//
// # Get all attribute_definition_integers
//
// Responses:
// default: genericError
//
//	200: attribute_definition_integerDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_INTEGERs(c *gin.Context) {

	// source slice
	var attribute_definition_integerDBs []orm.ATTRIBUTE_DEFINITION_INTEGERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_INTEGERs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.GetDB()

	_, err := db.Find(&attribute_definition_integerDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_definition_integerAPIs := make([]orm.ATTRIBUTE_DEFINITION_INTEGERAPI, 0)

	// for each attribute_definition_integer, update fields from the database nullable fields
	for idx := range attribute_definition_integerDBs {
		attribute_definition_integerDB := &attribute_definition_integerDBs[idx]
		_ = attribute_definition_integerDB
		var attribute_definition_integerAPI orm.ATTRIBUTE_DEFINITION_INTEGERAPI

		// insertion point for updating fields
		attribute_definition_integerAPI.ID = attribute_definition_integerDB.ID
		attribute_definition_integerDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER_WOP(&attribute_definition_integerAPI.ATTRIBUTE_DEFINITION_INTEGER_WOP)
		attribute_definition_integerAPI.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding = attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding
		attribute_definition_integerAPIs = append(attribute_definition_integerAPIs, attribute_definition_integerAPI)
	}

	c.JSON(http.StatusOK, attribute_definition_integerAPIs)
}

// PostATTRIBUTE_DEFINITION_INTEGER
//
// swagger:route POST /attribute_definition_integers attribute_definition_integers postATTRIBUTE_DEFINITION_INTEGER
//
// Creates a attribute_definition_integer
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_DEFINITION_INTEGER(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_INTEGER.Lock()
	defer mutexATTRIBUTE_DEFINITION_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_DEFINITION_INTEGERs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_INTEGERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_definition_integer
	attribute_definition_integerDB := orm.ATTRIBUTE_DEFINITION_INTEGERDB{}
	attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding = input.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding
	attribute_definition_integerDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGER_WOP(&input.ATTRIBUTE_DEFINITION_INTEGER_WOP)

	_, err = db.Create(&attribute_definition_integerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseOneInstance(&attribute_definition_integerDB)
	attribute_definition_integer := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr[attribute_definition_integerDB.ID]

	if attribute_definition_integer != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_definition_integer)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_definition_integerDB)
}

// GetATTRIBUTE_DEFINITION_INTEGER
//
// swagger:route GET /attribute_definition_integers/{ID} attribute_definition_integers getATTRIBUTE_DEFINITION_INTEGER
//
// Gets the details for a attribute_definition_integer.
//
// Responses:
// default: genericError
//
//	200: attribute_definition_integerDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_INTEGER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_INTEGER", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.GetDB()

	// Get attribute_definition_integerDB in DB
	var attribute_definition_integerDB orm.ATTRIBUTE_DEFINITION_INTEGERDB
	if _, err := db.First(&attribute_definition_integerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_definition_integerAPI orm.ATTRIBUTE_DEFINITION_INTEGERAPI
	attribute_definition_integerAPI.ID = attribute_definition_integerDB.ID
	attribute_definition_integerAPI.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding = attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding
	attribute_definition_integerDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER_WOP(&attribute_definition_integerAPI.ATTRIBUTE_DEFINITION_INTEGER_WOP)

	c.JSON(http.StatusOK, attribute_definition_integerAPI)
}

// UpdateATTRIBUTE_DEFINITION_INTEGER
//
// swagger:route PATCH /attribute_definition_integers/{ID} attribute_definition_integers updateATTRIBUTE_DEFINITION_INTEGER
//
// # Update a attribute_definition_integer
//
// Responses:
// default: genericError
//
//	200: attribute_definition_integerDBResponse
func (controller *Controller) UpdateATTRIBUTE_DEFINITION_INTEGER(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_INTEGER.Lock()
	defer mutexATTRIBUTE_DEFINITION_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	hasMouseEvent := false
	shiftKey := false
	_ = shiftKey
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	if len(_values) >= 2 {
		hasMouseEvent = true
		_shiftKeyValues := _values["shiftKey"]
		if len(_shiftKeyValues) == 1 {
			shiftKey = _shiftKeyValues[0] == "true"
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_INTEGERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_definition_integerDB orm.ATTRIBUTE_DEFINITION_INTEGERDB

	// fetch the attribute_definition_integer
	_, err := db.First(&attribute_definition_integerDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_definition_integerDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_INTEGER_WOP(&input.ATTRIBUTE_DEFINITION_INTEGER_WOP)
	attribute_definition_integerDB.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding = input.ATTRIBUTE_DEFINITION_INTEGERPointersEncoding

	db, _ = db.Model(&attribute_definition_integerDB)
	_, err = db.Updates(&attribute_definition_integerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_integerNew := new(models.ATTRIBUTE_DEFINITION_INTEGER)
	attribute_definition_integerDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integerNew)

	// redeem pointers
	attribute_definition_integerDB.DecodePointers(backRepo, attribute_definition_integerNew)

	// get stage instance from DB instance, and call callback function
	attribute_definition_integerOld := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr[attribute_definition_integerDB.ID]
	if attribute_definition_integerOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_definition_integerOld, attribute_definition_integerNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_definition_integerOld, attribute_definition_integerNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_definition_integerDB
	c.JSON(http.StatusOK, attribute_definition_integerDB)
}

// DeleteATTRIBUTE_DEFINITION_INTEGER
//
// swagger:route DELETE /attribute_definition_integers/{ID} attribute_definition_integers deleteATTRIBUTE_DEFINITION_INTEGER
//
// # Delete a attribute_definition_integer
//
// default: genericError
//
//	200: attribute_definition_integerDBResponse
func (controller *Controller) DeleteATTRIBUTE_DEFINITION_INTEGER(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_INTEGER.Lock()
	defer mutexATTRIBUTE_DEFINITION_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_DEFINITION_INTEGER", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.GetDB()

	// Get model if exist
	var attribute_definition_integerDB orm.ATTRIBUTE_DEFINITION_INTEGERDB
	if _, err := db.First(&attribute_definition_integerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_definition_integerDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_integerDeleted := new(models.ATTRIBUTE_DEFINITION_INTEGER)
	attribute_definition_integerDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integerDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_definition_integerStaged := backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr[attribute_definition_integerDB.ID]
	if attribute_definition_integerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_definition_integerStaged, attribute_definition_integerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
