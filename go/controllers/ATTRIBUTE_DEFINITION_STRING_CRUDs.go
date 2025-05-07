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
var __ATTRIBUTE_DEFINITION_STRING__dummysDeclaration__ models.ATTRIBUTE_DEFINITION_STRING
var __ATTRIBUTE_DEFINITION_STRING_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_DEFINITION_STRING sync.Mutex

// An ATTRIBUTE_DEFINITION_STRINGID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_DEFINITION_STRING updateATTRIBUTE_DEFINITION_STRING deleteATTRIBUTE_DEFINITION_STRING
type ATTRIBUTE_DEFINITION_STRINGID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_DEFINITION_STRINGInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_DEFINITION_STRING updateATTRIBUTE_DEFINITION_STRING
type ATTRIBUTE_DEFINITION_STRINGInput struct {
	// The ATTRIBUTE_DEFINITION_STRING to submit or modify
	// in: body
	ATTRIBUTE_DEFINITION_STRING *orm.ATTRIBUTE_DEFINITION_STRINGAPI
}

// GetATTRIBUTE_DEFINITION_STRINGs
//
// swagger:route GET /attribute_definition_strings attribute_definition_strings getATTRIBUTE_DEFINITION_STRINGs
//
// # Get all attribute_definition_strings
//
// Responses:
// default: genericError
//
//	200: attribute_definition_stringDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_STRINGs(c *gin.Context) {

	// source slice
	var attribute_definition_stringDBs []orm.ATTRIBUTE_DEFINITION_STRINGDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_STRINGs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.GetDB()

	_, err := db.Find(&attribute_definition_stringDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_definition_stringAPIs := make([]orm.ATTRIBUTE_DEFINITION_STRINGAPI, 0)

	// for each attribute_definition_string, update fields from the database nullable fields
	for idx := range attribute_definition_stringDBs {
		attribute_definition_stringDB := &attribute_definition_stringDBs[idx]
		_ = attribute_definition_stringDB
		var attribute_definition_stringAPI orm.ATTRIBUTE_DEFINITION_STRINGAPI

		// insertion point for updating fields
		attribute_definition_stringAPI.ID = attribute_definition_stringDB.ID
		attribute_definition_stringDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING_WOP(&attribute_definition_stringAPI.ATTRIBUTE_DEFINITION_STRING_WOP)
		attribute_definition_stringAPI.ATTRIBUTE_DEFINITION_STRINGPointersEncoding = attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding
		attribute_definition_stringAPIs = append(attribute_definition_stringAPIs, attribute_definition_stringAPI)
	}

	c.JSON(http.StatusOK, attribute_definition_stringAPIs)
}

// PostATTRIBUTE_DEFINITION_STRING
//
// swagger:route POST /attribute_definition_strings attribute_definition_strings postATTRIBUTE_DEFINITION_STRING
//
// Creates a attribute_definition_string
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_DEFINITION_STRING(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_STRING.Lock()
	defer mutexATTRIBUTE_DEFINITION_STRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_DEFINITION_STRINGs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_STRINGAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_definition_string
	attribute_definition_stringDB := orm.ATTRIBUTE_DEFINITION_STRINGDB{}
	attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding = input.ATTRIBUTE_DEFINITION_STRINGPointersEncoding
	attribute_definition_stringDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRING_WOP(&input.ATTRIBUTE_DEFINITION_STRING_WOP)

	_, err = db.Create(&attribute_definition_stringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseOneInstance(&attribute_definition_stringDB)
	attribute_definition_string := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr[attribute_definition_stringDB.ID]

	if attribute_definition_string != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_definition_string)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_definition_stringDB)
}

// GetATTRIBUTE_DEFINITION_STRING
//
// swagger:route GET /attribute_definition_strings/{ID} attribute_definition_strings getATTRIBUTE_DEFINITION_STRING
//
// Gets the details for a attribute_definition_string.
//
// Responses:
// default: genericError
//
//	200: attribute_definition_stringDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_STRING(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_STRING", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.GetDB()

	// Get attribute_definition_stringDB in DB
	var attribute_definition_stringDB orm.ATTRIBUTE_DEFINITION_STRINGDB
	if _, err := db.First(&attribute_definition_stringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_definition_stringAPI orm.ATTRIBUTE_DEFINITION_STRINGAPI
	attribute_definition_stringAPI.ID = attribute_definition_stringDB.ID
	attribute_definition_stringAPI.ATTRIBUTE_DEFINITION_STRINGPointersEncoding = attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding
	attribute_definition_stringDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING_WOP(&attribute_definition_stringAPI.ATTRIBUTE_DEFINITION_STRING_WOP)

	c.JSON(http.StatusOK, attribute_definition_stringAPI)
}

// UpdateATTRIBUTE_DEFINITION_STRING
//
// swagger:route PATCH /attribute_definition_strings/{ID} attribute_definition_strings updateATTRIBUTE_DEFINITION_STRING
//
// # Update a attribute_definition_string
//
// Responses:
// default: genericError
//
//	200: attribute_definition_stringDBResponse
func (controller *Controller) UpdateATTRIBUTE_DEFINITION_STRING(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_STRING.Lock()
	defer mutexATTRIBUTE_DEFINITION_STRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTE_DEFINITION_STRING", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_STRINGAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_definition_stringDB orm.ATTRIBUTE_DEFINITION_STRINGDB

	// fetch the attribute_definition_string
	_, err := db.First(&attribute_definition_stringDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_definition_stringDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_STRING_WOP(&input.ATTRIBUTE_DEFINITION_STRING_WOP)
	attribute_definition_stringDB.ATTRIBUTE_DEFINITION_STRINGPointersEncoding = input.ATTRIBUTE_DEFINITION_STRINGPointersEncoding

	db, _ = db.Model(&attribute_definition_stringDB)
	_, err = db.Updates(&attribute_definition_stringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_stringNew := new(models.ATTRIBUTE_DEFINITION_STRING)
	attribute_definition_stringDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING(attribute_definition_stringNew)

	// redeem pointers
	attribute_definition_stringDB.DecodePointers(backRepo, attribute_definition_stringNew)

	// get stage instance from DB instance, and call callback function
	attribute_definition_stringOld := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr[attribute_definition_stringDB.ID]
	if attribute_definition_stringOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attribute_definition_stringOld, attribute_definition_stringNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_definition_stringDB
	c.JSON(http.StatusOK, attribute_definition_stringDB)
}

// DeleteATTRIBUTE_DEFINITION_STRING
//
// swagger:route DELETE /attribute_definition_strings/{ID} attribute_definition_strings deleteATTRIBUTE_DEFINITION_STRING
//
// # Delete a attribute_definition_string
//
// default: genericError
//
//	200: attribute_definition_stringDBResponse
func (controller *Controller) DeleteATTRIBUTE_DEFINITION_STRING(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_STRING.Lock()
	defer mutexATTRIBUTE_DEFINITION_STRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_DEFINITION_STRING", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.GetDB()

	// Get model if exist
	var attribute_definition_stringDB orm.ATTRIBUTE_DEFINITION_STRINGDB
	if _, err := db.First(&attribute_definition_stringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_definition_stringDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_stringDeleted := new(models.ATTRIBUTE_DEFINITION_STRING)
	attribute_definition_stringDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_STRING(attribute_definition_stringDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_definition_stringStaged := backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr[attribute_definition_stringDB.ID]
	if attribute_definition_stringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_definition_stringStaged, attribute_definition_stringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
