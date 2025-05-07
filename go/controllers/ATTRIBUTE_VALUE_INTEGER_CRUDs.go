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
var __ATTRIBUTE_VALUE_INTEGER__dummysDeclaration__ models.ATTRIBUTE_VALUE_INTEGER
var __ATTRIBUTE_VALUE_INTEGER_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_VALUE_INTEGER sync.Mutex

// An ATTRIBUTE_VALUE_INTEGERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_VALUE_INTEGER updateATTRIBUTE_VALUE_INTEGER deleteATTRIBUTE_VALUE_INTEGER
type ATTRIBUTE_VALUE_INTEGERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_VALUE_INTEGERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_VALUE_INTEGER updateATTRIBUTE_VALUE_INTEGER
type ATTRIBUTE_VALUE_INTEGERInput struct {
	// The ATTRIBUTE_VALUE_INTEGER to submit or modify
	// in: body
	ATTRIBUTE_VALUE_INTEGER *orm.ATTRIBUTE_VALUE_INTEGERAPI
}

// GetATTRIBUTE_VALUE_INTEGERs
//
// swagger:route GET /attribute_value_integers attribute_value_integers getATTRIBUTE_VALUE_INTEGERs
//
// # Get all attribute_value_integers
//
// Responses:
// default: genericError
//
//	200: attribute_value_integerDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_INTEGERs(c *gin.Context) {

	// source slice
	var attribute_value_integerDBs []orm.ATTRIBUTE_VALUE_INTEGERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_INTEGERs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.GetDB()

	_, err := db.Find(&attribute_value_integerDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_value_integerAPIs := make([]orm.ATTRIBUTE_VALUE_INTEGERAPI, 0)

	// for each attribute_value_integer, update fields from the database nullable fields
	for idx := range attribute_value_integerDBs {
		attribute_value_integerDB := &attribute_value_integerDBs[idx]
		_ = attribute_value_integerDB
		var attribute_value_integerAPI orm.ATTRIBUTE_VALUE_INTEGERAPI

		// insertion point for updating fields
		attribute_value_integerAPI.ID = attribute_value_integerDB.ID
		attribute_value_integerDB.CopyBasicFieldsToATTRIBUTE_VALUE_INTEGER_WOP(&attribute_value_integerAPI.ATTRIBUTE_VALUE_INTEGER_WOP)
		attribute_value_integerAPI.ATTRIBUTE_VALUE_INTEGERPointersEncoding = attribute_value_integerDB.ATTRIBUTE_VALUE_INTEGERPointersEncoding
		attribute_value_integerAPIs = append(attribute_value_integerAPIs, attribute_value_integerAPI)
	}

	c.JSON(http.StatusOK, attribute_value_integerAPIs)
}

// PostATTRIBUTE_VALUE_INTEGER
//
// swagger:route POST /attribute_value_integers attribute_value_integers postATTRIBUTE_VALUE_INTEGER
//
// Creates a attribute_value_integer
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_VALUE_INTEGER(c *gin.Context) {

	mutexATTRIBUTE_VALUE_INTEGER.Lock()
	defer mutexATTRIBUTE_VALUE_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_VALUE_INTEGERs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_INTEGERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_value_integer
	attribute_value_integerDB := orm.ATTRIBUTE_VALUE_INTEGERDB{}
	attribute_value_integerDB.ATTRIBUTE_VALUE_INTEGERPointersEncoding = input.ATTRIBUTE_VALUE_INTEGERPointersEncoding
	attribute_value_integerDB.CopyBasicFieldsFromATTRIBUTE_VALUE_INTEGER_WOP(&input.ATTRIBUTE_VALUE_INTEGER_WOP)

	_, err = db.Create(&attribute_value_integerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.CheckoutPhaseOneInstance(&attribute_value_integerDB)
	attribute_value_integer := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERPtr[attribute_value_integerDB.ID]

	if attribute_value_integer != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_value_integer)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_value_integerDB)
}

// GetATTRIBUTE_VALUE_INTEGER
//
// swagger:route GET /attribute_value_integers/{ID} attribute_value_integers getATTRIBUTE_VALUE_INTEGER
//
// Gets the details for a attribute_value_integer.
//
// Responses:
// default: genericError
//
//	200: attribute_value_integerDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_INTEGER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_INTEGER", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.GetDB()

	// Get attribute_value_integerDB in DB
	var attribute_value_integerDB orm.ATTRIBUTE_VALUE_INTEGERDB
	if _, err := db.First(&attribute_value_integerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_value_integerAPI orm.ATTRIBUTE_VALUE_INTEGERAPI
	attribute_value_integerAPI.ID = attribute_value_integerDB.ID
	attribute_value_integerAPI.ATTRIBUTE_VALUE_INTEGERPointersEncoding = attribute_value_integerDB.ATTRIBUTE_VALUE_INTEGERPointersEncoding
	attribute_value_integerDB.CopyBasicFieldsToATTRIBUTE_VALUE_INTEGER_WOP(&attribute_value_integerAPI.ATTRIBUTE_VALUE_INTEGER_WOP)

	c.JSON(http.StatusOK, attribute_value_integerAPI)
}

// UpdateATTRIBUTE_VALUE_INTEGER
//
// swagger:route PATCH /attribute_value_integers/{ID} attribute_value_integers updateATTRIBUTE_VALUE_INTEGER
//
// # Update a attribute_value_integer
//
// Responses:
// default: genericError
//
//	200: attribute_value_integerDBResponse
func (controller *Controller) UpdateATTRIBUTE_VALUE_INTEGER(c *gin.Context) {

	mutexATTRIBUTE_VALUE_INTEGER.Lock()
	defer mutexATTRIBUTE_VALUE_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTE_VALUE_INTEGER", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_INTEGERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_value_integerDB orm.ATTRIBUTE_VALUE_INTEGERDB

	// fetch the attribute_value_integer
	_, err := db.First(&attribute_value_integerDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_value_integerDB.CopyBasicFieldsFromATTRIBUTE_VALUE_INTEGER_WOP(&input.ATTRIBUTE_VALUE_INTEGER_WOP)
	attribute_value_integerDB.ATTRIBUTE_VALUE_INTEGERPointersEncoding = input.ATTRIBUTE_VALUE_INTEGERPointersEncoding

	db, _ = db.Model(&attribute_value_integerDB)
	_, err = db.Updates(&attribute_value_integerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_integerNew := new(models.ATTRIBUTE_VALUE_INTEGER)
	attribute_value_integerDB.CopyBasicFieldsToATTRIBUTE_VALUE_INTEGER(attribute_value_integerNew)

	// redeem pointers
	attribute_value_integerDB.DecodePointers(backRepo, attribute_value_integerNew)

	// get stage instance from DB instance, and call callback function
	attribute_value_integerOld := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERPtr[attribute_value_integerDB.ID]
	if attribute_value_integerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attribute_value_integerOld, attribute_value_integerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_value_integerDB
	c.JSON(http.StatusOK, attribute_value_integerDB)
}

// DeleteATTRIBUTE_VALUE_INTEGER
//
// swagger:route DELETE /attribute_value_integers/{ID} attribute_value_integers deleteATTRIBUTE_VALUE_INTEGER
//
// # Delete a attribute_value_integer
//
// default: genericError
//
//	200: attribute_value_integerDBResponse
func (controller *Controller) DeleteATTRIBUTE_VALUE_INTEGER(c *gin.Context) {

	mutexATTRIBUTE_VALUE_INTEGER.Lock()
	defer mutexATTRIBUTE_VALUE_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_VALUE_INTEGER", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.GetDB()

	// Get model if exist
	var attribute_value_integerDB orm.ATTRIBUTE_VALUE_INTEGERDB
	if _, err := db.First(&attribute_value_integerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_value_integerDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_integerDeleted := new(models.ATTRIBUTE_VALUE_INTEGER)
	attribute_value_integerDB.CopyBasicFieldsToATTRIBUTE_VALUE_INTEGER(attribute_value_integerDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_value_integerStaged := backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERPtr[attribute_value_integerDB.ID]
	if attribute_value_integerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_value_integerStaged, attribute_value_integerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
