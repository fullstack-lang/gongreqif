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
var __A_ATTRIBUTE_VALUE_STRING__dummysDeclaration__ models.A_ATTRIBUTE_VALUE_STRING
var __A_ATTRIBUTE_VALUE_STRING_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_VALUE_STRING sync.Mutex

// An A_ATTRIBUTE_VALUE_STRINGID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_VALUE_STRING updateA_ATTRIBUTE_VALUE_STRING deleteA_ATTRIBUTE_VALUE_STRING
type A_ATTRIBUTE_VALUE_STRINGID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_VALUE_STRINGInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_VALUE_STRING updateA_ATTRIBUTE_VALUE_STRING
type A_ATTRIBUTE_VALUE_STRINGInput struct {
	// The A_ATTRIBUTE_VALUE_STRING to submit or modify
	// in: body
	A_ATTRIBUTE_VALUE_STRING *orm.A_ATTRIBUTE_VALUE_STRINGAPI
}

// GetA_ATTRIBUTE_VALUE_STRINGs
//
// swagger:route GET /a_attribute_value_strings a_attribute_value_strings getA_ATTRIBUTE_VALUE_STRINGs
//
// # Get all a_attribute_value_strings
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_stringDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_STRINGs(c *gin.Context) {

	// source slice
	var a_attribute_value_stringDBs []orm.A_ATTRIBUTE_VALUE_STRINGDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_STRINGs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.GetDB()

	_, err := db.Find(&a_attribute_value_stringDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_value_stringAPIs := make([]orm.A_ATTRIBUTE_VALUE_STRINGAPI, 0)

	// for each a_attribute_value_string, update fields from the database nullable fields
	for idx := range a_attribute_value_stringDBs {
		a_attribute_value_stringDB := &a_attribute_value_stringDBs[idx]
		_ = a_attribute_value_stringDB
		var a_attribute_value_stringAPI orm.A_ATTRIBUTE_VALUE_STRINGAPI

		// insertion point for updating fields
		a_attribute_value_stringAPI.ID = a_attribute_value_stringDB.ID
		a_attribute_value_stringDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_STRING_WOP(&a_attribute_value_stringAPI.A_ATTRIBUTE_VALUE_STRING_WOP)
		a_attribute_value_stringAPI.A_ATTRIBUTE_VALUE_STRINGPointersEncoding = a_attribute_value_stringDB.A_ATTRIBUTE_VALUE_STRINGPointersEncoding
		a_attribute_value_stringAPIs = append(a_attribute_value_stringAPIs, a_attribute_value_stringAPI)
	}

	c.JSON(http.StatusOK, a_attribute_value_stringAPIs)
}

// PostA_ATTRIBUTE_VALUE_STRING
//
// swagger:route POST /a_attribute_value_strings a_attribute_value_strings postA_ATTRIBUTE_VALUE_STRING
//
// Creates a a_attribute_value_string
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_VALUE_STRING(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_STRING.Lock()
	defer mutexA_ATTRIBUTE_VALUE_STRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_VALUE_STRINGs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_STRINGAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_value_string
	a_attribute_value_stringDB := orm.A_ATTRIBUTE_VALUE_STRINGDB{}
	a_attribute_value_stringDB.A_ATTRIBUTE_VALUE_STRINGPointersEncoding = input.A_ATTRIBUTE_VALUE_STRINGPointersEncoding
	a_attribute_value_stringDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_STRING_WOP(&input.A_ATTRIBUTE_VALUE_STRING_WOP)

	_, err = db.Create(&a_attribute_value_stringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.CheckoutPhaseOneInstance(&a_attribute_value_stringDB)
	a_attribute_value_string := backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.Map_A_ATTRIBUTE_VALUE_STRINGDBID_A_ATTRIBUTE_VALUE_STRINGPtr[a_attribute_value_stringDB.ID]

	if a_attribute_value_string != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_value_string)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_value_stringDB)
}

// GetA_ATTRIBUTE_VALUE_STRING
//
// swagger:route GET /a_attribute_value_strings/{ID} a_attribute_value_strings getA_ATTRIBUTE_VALUE_STRING
//
// Gets the details for a a_attribute_value_string.
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_stringDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_STRING(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_STRING", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.GetDB()

	// Get a_attribute_value_stringDB in DB
	var a_attribute_value_stringDB orm.A_ATTRIBUTE_VALUE_STRINGDB
	if _, err := db.First(&a_attribute_value_stringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_value_stringAPI orm.A_ATTRIBUTE_VALUE_STRINGAPI
	a_attribute_value_stringAPI.ID = a_attribute_value_stringDB.ID
	a_attribute_value_stringAPI.A_ATTRIBUTE_VALUE_STRINGPointersEncoding = a_attribute_value_stringDB.A_ATTRIBUTE_VALUE_STRINGPointersEncoding
	a_attribute_value_stringDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_STRING_WOP(&a_attribute_value_stringAPI.A_ATTRIBUTE_VALUE_STRING_WOP)

	c.JSON(http.StatusOK, a_attribute_value_stringAPI)
}

// UpdateA_ATTRIBUTE_VALUE_STRING
//
// swagger:route PATCH /a_attribute_value_strings/{ID} a_attribute_value_strings updateA_ATTRIBUTE_VALUE_STRING
//
// # Update a a_attribute_value_string
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_stringDBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_VALUE_STRING(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_STRING.Lock()
	defer mutexA_ATTRIBUTE_VALUE_STRING.Unlock()

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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_STRINGAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_value_stringDB orm.A_ATTRIBUTE_VALUE_STRINGDB

	// fetch the a_attribute_value_string
	_, err := db.First(&a_attribute_value_stringDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_value_stringDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_STRING_WOP(&input.A_ATTRIBUTE_VALUE_STRING_WOP)
	a_attribute_value_stringDB.A_ATTRIBUTE_VALUE_STRINGPointersEncoding = input.A_ATTRIBUTE_VALUE_STRINGPointersEncoding

	db, _ = db.Model(&a_attribute_value_stringDB)
	_, err = db.Updates(&a_attribute_value_stringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_stringNew := new(models.A_ATTRIBUTE_VALUE_STRING)
	a_attribute_value_stringDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_STRING(a_attribute_value_stringNew)

	// redeem pointers
	a_attribute_value_stringDB.DecodePointers(backRepo, a_attribute_value_stringNew)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_stringOld := backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.Map_A_ATTRIBUTE_VALUE_STRINGDBID_A_ATTRIBUTE_VALUE_STRINGPtr[a_attribute_value_stringDB.ID]
	if a_attribute_value_stringOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_stringOld, a_attribute_value_stringNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_stringOld, a_attribute_value_stringNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_value_stringDB
	c.JSON(http.StatusOK, a_attribute_value_stringDB)
}

// DeleteA_ATTRIBUTE_VALUE_STRING
//
// swagger:route DELETE /a_attribute_value_strings/{ID} a_attribute_value_strings deleteA_ATTRIBUTE_VALUE_STRING
//
// # Delete a a_attribute_value_string
//
// default: genericError
//
//	200: a_attribute_value_stringDBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_VALUE_STRING(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_STRING.Lock()
	defer mutexA_ATTRIBUTE_VALUE_STRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_VALUE_STRING", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.GetDB()

	// Get model if exist
	var a_attribute_value_stringDB orm.A_ATTRIBUTE_VALUE_STRINGDB
	if _, err := db.First(&a_attribute_value_stringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_value_stringDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_stringDeleted := new(models.A_ATTRIBUTE_VALUE_STRING)
	a_attribute_value_stringDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_STRING(a_attribute_value_stringDeleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_stringStaged := backRepo.BackRepoA_ATTRIBUTE_VALUE_STRING.Map_A_ATTRIBUTE_VALUE_STRINGDBID_A_ATTRIBUTE_VALUE_STRINGPtr[a_attribute_value_stringDB.ID]
	if a_attribute_value_stringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_value_stringStaged, a_attribute_value_stringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
