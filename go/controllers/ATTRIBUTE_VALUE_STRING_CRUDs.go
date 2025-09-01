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
var __ATTRIBUTE_VALUE_STRING__dummysDeclaration__ models.ATTRIBUTE_VALUE_STRING
var __ATTRIBUTE_VALUE_STRING_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_VALUE_STRING sync.Mutex

// An ATTRIBUTE_VALUE_STRINGID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_VALUE_STRING updateATTRIBUTE_VALUE_STRING deleteATTRIBUTE_VALUE_STRING
type ATTRIBUTE_VALUE_STRINGID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_VALUE_STRINGInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_VALUE_STRING updateATTRIBUTE_VALUE_STRING
type ATTRIBUTE_VALUE_STRINGInput struct {
	// The ATTRIBUTE_VALUE_STRING to submit or modify
	// in: body
	ATTRIBUTE_VALUE_STRING *orm.ATTRIBUTE_VALUE_STRINGAPI
}

// GetATTRIBUTE_VALUE_STRINGs
//
// swagger:route GET /attribute_value_strings attribute_value_strings getATTRIBUTE_VALUE_STRINGs
//
// # Get all attribute_value_strings
//
// Responses:
// default: genericError
//
//	200: attribute_value_stringDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_STRINGs(c *gin.Context) {

	// source slice
	var attribute_value_stringDBs []orm.ATTRIBUTE_VALUE_STRINGDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_STRINGs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_STRING.GetDB()

	_, err := db.Find(&attribute_value_stringDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_value_stringAPIs := make([]orm.ATTRIBUTE_VALUE_STRINGAPI, 0)

	// for each attribute_value_string, update fields from the database nullable fields
	for idx := range attribute_value_stringDBs {
		attribute_value_stringDB := &attribute_value_stringDBs[idx]
		_ = attribute_value_stringDB
		var attribute_value_stringAPI orm.ATTRIBUTE_VALUE_STRINGAPI

		// insertion point for updating fields
		attribute_value_stringAPI.ID = attribute_value_stringDB.ID
		attribute_value_stringDB.CopyBasicFieldsToATTRIBUTE_VALUE_STRING_WOP(&attribute_value_stringAPI.ATTRIBUTE_VALUE_STRING_WOP)
		attribute_value_stringAPI.ATTRIBUTE_VALUE_STRINGPointersEncoding = attribute_value_stringDB.ATTRIBUTE_VALUE_STRINGPointersEncoding
		attribute_value_stringAPIs = append(attribute_value_stringAPIs, attribute_value_stringAPI)
	}

	c.JSON(http.StatusOK, attribute_value_stringAPIs)
}

// PostATTRIBUTE_VALUE_STRING
//
// swagger:route POST /attribute_value_strings attribute_value_strings postATTRIBUTE_VALUE_STRING
//
// Creates a attribute_value_string
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_VALUE_STRING(c *gin.Context) {

	mutexATTRIBUTE_VALUE_STRING.Lock()
	defer mutexATTRIBUTE_VALUE_STRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_VALUE_STRINGs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_STRING.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_STRINGAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_value_string
	attribute_value_stringDB := orm.ATTRIBUTE_VALUE_STRINGDB{}
	attribute_value_stringDB.ATTRIBUTE_VALUE_STRINGPointersEncoding = input.ATTRIBUTE_VALUE_STRINGPointersEncoding
	attribute_value_stringDB.CopyBasicFieldsFromATTRIBUTE_VALUE_STRING_WOP(&input.ATTRIBUTE_VALUE_STRING_WOP)

	_, err = db.Create(&attribute_value_stringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.CheckoutPhaseOneInstance(&attribute_value_stringDB)
	attribute_value_string := backRepo.BackRepoATTRIBUTE_VALUE_STRING.Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGPtr[attribute_value_stringDB.ID]

	if attribute_value_string != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_value_string)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_value_stringDB)
}

// GetATTRIBUTE_VALUE_STRING
//
// swagger:route GET /attribute_value_strings/{ID} attribute_value_strings getATTRIBUTE_VALUE_STRING
//
// Gets the details for a attribute_value_string.
//
// Responses:
// default: genericError
//
//	200: attribute_value_stringDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_STRING(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_STRING", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_STRING.GetDB()

	// Get attribute_value_stringDB in DB
	var attribute_value_stringDB orm.ATTRIBUTE_VALUE_STRINGDB
	if _, err := db.First(&attribute_value_stringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_value_stringAPI orm.ATTRIBUTE_VALUE_STRINGAPI
	attribute_value_stringAPI.ID = attribute_value_stringDB.ID
	attribute_value_stringAPI.ATTRIBUTE_VALUE_STRINGPointersEncoding = attribute_value_stringDB.ATTRIBUTE_VALUE_STRINGPointersEncoding
	attribute_value_stringDB.CopyBasicFieldsToATTRIBUTE_VALUE_STRING_WOP(&attribute_value_stringAPI.ATTRIBUTE_VALUE_STRING_WOP)

	c.JSON(http.StatusOK, attribute_value_stringAPI)
}

// UpdateATTRIBUTE_VALUE_STRING
//
// swagger:route PATCH /attribute_value_strings/{ID} attribute_value_strings updateATTRIBUTE_VALUE_STRING
//
// # Update a attribute_value_string
//
// Responses:
// default: genericError
//
//	200: attribute_value_stringDBResponse
func (controller *Controller) UpdateATTRIBUTE_VALUE_STRING(c *gin.Context) {

	mutexATTRIBUTE_VALUE_STRING.Lock()
	defer mutexATTRIBUTE_VALUE_STRING.Unlock()

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
	db := backRepo.BackRepoATTRIBUTE_VALUE_STRING.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_STRINGAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_value_stringDB orm.ATTRIBUTE_VALUE_STRINGDB

	// fetch the attribute_value_string
	_, err := db.First(&attribute_value_stringDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_value_stringDB.CopyBasicFieldsFromATTRIBUTE_VALUE_STRING_WOP(&input.ATTRIBUTE_VALUE_STRING_WOP)
	attribute_value_stringDB.ATTRIBUTE_VALUE_STRINGPointersEncoding = input.ATTRIBUTE_VALUE_STRINGPointersEncoding

	db, _ = db.Model(&attribute_value_stringDB)
	_, err = db.Updates(&attribute_value_stringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_stringNew := new(models.ATTRIBUTE_VALUE_STRING)
	attribute_value_stringDB.CopyBasicFieldsToATTRIBUTE_VALUE_STRING(attribute_value_stringNew)

	// redeem pointers
	attribute_value_stringDB.DecodePointers(backRepo, attribute_value_stringNew)

	// get stage instance from DB instance, and call callback function
	attribute_value_stringOld := backRepo.BackRepoATTRIBUTE_VALUE_STRING.Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGPtr[attribute_value_stringDB.ID]
	if attribute_value_stringOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_value_stringOld, attribute_value_stringNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_value_stringOld, attribute_value_stringNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_value_stringDB
	c.JSON(http.StatusOK, attribute_value_stringDB)
}

// DeleteATTRIBUTE_VALUE_STRING
//
// swagger:route DELETE /attribute_value_strings/{ID} attribute_value_strings deleteATTRIBUTE_VALUE_STRING
//
// # Delete a attribute_value_string
//
// default: genericError
//
//	200: attribute_value_stringDBResponse
func (controller *Controller) DeleteATTRIBUTE_VALUE_STRING(c *gin.Context) {

	mutexATTRIBUTE_VALUE_STRING.Lock()
	defer mutexATTRIBUTE_VALUE_STRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_VALUE_STRING", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_STRING.GetDB()

	// Get model if exist
	var attribute_value_stringDB orm.ATTRIBUTE_VALUE_STRINGDB
	if _, err := db.First(&attribute_value_stringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_value_stringDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_stringDeleted := new(models.ATTRIBUTE_VALUE_STRING)
	attribute_value_stringDB.CopyBasicFieldsToATTRIBUTE_VALUE_STRING(attribute_value_stringDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_value_stringStaged := backRepo.BackRepoATTRIBUTE_VALUE_STRING.Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGPtr[attribute_value_stringDB.ID]
	if attribute_value_stringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_value_stringStaged, attribute_value_stringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
