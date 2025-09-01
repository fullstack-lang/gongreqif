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
var __A_ATTRIBUTE_VALUE_BOOLEAN__dummysDeclaration__ models.A_ATTRIBUTE_VALUE_BOOLEAN
var __A_ATTRIBUTE_VALUE_BOOLEAN_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_VALUE_BOOLEAN sync.Mutex

// An A_ATTRIBUTE_VALUE_BOOLEANID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_VALUE_BOOLEAN updateA_ATTRIBUTE_VALUE_BOOLEAN deleteA_ATTRIBUTE_VALUE_BOOLEAN
type A_ATTRIBUTE_VALUE_BOOLEANID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_VALUE_BOOLEANInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_VALUE_BOOLEAN updateA_ATTRIBUTE_VALUE_BOOLEAN
type A_ATTRIBUTE_VALUE_BOOLEANInput struct {
	// The A_ATTRIBUTE_VALUE_BOOLEAN to submit or modify
	// in: body
	A_ATTRIBUTE_VALUE_BOOLEAN *orm.A_ATTRIBUTE_VALUE_BOOLEANAPI
}

// GetA_ATTRIBUTE_VALUE_BOOLEANs
//
// swagger:route GET /a_attribute_value_booleans a_attribute_value_booleans getA_ATTRIBUTE_VALUE_BOOLEANs
//
// # Get all a_attribute_value_booleans
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_booleanDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_BOOLEANs(c *gin.Context) {

	// source slice
	var a_attribute_value_booleanDBs []orm.A_ATTRIBUTE_VALUE_BOOLEANDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_BOOLEANs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.GetDB()

	_, err := db.Find(&a_attribute_value_booleanDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_value_booleanAPIs := make([]orm.A_ATTRIBUTE_VALUE_BOOLEANAPI, 0)

	// for each a_attribute_value_boolean, update fields from the database nullable fields
	for idx := range a_attribute_value_booleanDBs {
		a_attribute_value_booleanDB := &a_attribute_value_booleanDBs[idx]
		_ = a_attribute_value_booleanDB
		var a_attribute_value_booleanAPI orm.A_ATTRIBUTE_VALUE_BOOLEANAPI

		// insertion point for updating fields
		a_attribute_value_booleanAPI.ID = a_attribute_value_booleanDB.ID
		a_attribute_value_booleanDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_BOOLEAN_WOP(&a_attribute_value_booleanAPI.A_ATTRIBUTE_VALUE_BOOLEAN_WOP)
		a_attribute_value_booleanAPI.A_ATTRIBUTE_VALUE_BOOLEANPointersEncoding = a_attribute_value_booleanDB.A_ATTRIBUTE_VALUE_BOOLEANPointersEncoding
		a_attribute_value_booleanAPIs = append(a_attribute_value_booleanAPIs, a_attribute_value_booleanAPI)
	}

	c.JSON(http.StatusOK, a_attribute_value_booleanAPIs)
}

// PostA_ATTRIBUTE_VALUE_BOOLEAN
//
// swagger:route POST /a_attribute_value_booleans a_attribute_value_booleans postA_ATTRIBUTE_VALUE_BOOLEAN
//
// Creates a a_attribute_value_boolean
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_VALUE_BOOLEAN(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_BOOLEAN.Lock()
	defer mutexA_ATTRIBUTE_VALUE_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_VALUE_BOOLEANs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_BOOLEANAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_value_boolean
	a_attribute_value_booleanDB := orm.A_ATTRIBUTE_VALUE_BOOLEANDB{}
	a_attribute_value_booleanDB.A_ATTRIBUTE_VALUE_BOOLEANPointersEncoding = input.A_ATTRIBUTE_VALUE_BOOLEANPointersEncoding
	a_attribute_value_booleanDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_BOOLEAN_WOP(&input.A_ATTRIBUTE_VALUE_BOOLEAN_WOP)

	_, err = db.Create(&a_attribute_value_booleanDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.CheckoutPhaseOneInstance(&a_attribute_value_booleanDB)
	a_attribute_value_boolean := backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.Map_A_ATTRIBUTE_VALUE_BOOLEANDBID_A_ATTRIBUTE_VALUE_BOOLEANPtr[a_attribute_value_booleanDB.ID]

	if a_attribute_value_boolean != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_value_boolean)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_value_booleanDB)
}

// GetA_ATTRIBUTE_VALUE_BOOLEAN
//
// swagger:route GET /a_attribute_value_booleans/{ID} a_attribute_value_booleans getA_ATTRIBUTE_VALUE_BOOLEAN
//
// Gets the details for a a_attribute_value_boolean.
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_booleanDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_BOOLEAN(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_BOOLEAN", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.GetDB()

	// Get a_attribute_value_booleanDB in DB
	var a_attribute_value_booleanDB orm.A_ATTRIBUTE_VALUE_BOOLEANDB
	if _, err := db.First(&a_attribute_value_booleanDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_value_booleanAPI orm.A_ATTRIBUTE_VALUE_BOOLEANAPI
	a_attribute_value_booleanAPI.ID = a_attribute_value_booleanDB.ID
	a_attribute_value_booleanAPI.A_ATTRIBUTE_VALUE_BOOLEANPointersEncoding = a_attribute_value_booleanDB.A_ATTRIBUTE_VALUE_BOOLEANPointersEncoding
	a_attribute_value_booleanDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_BOOLEAN_WOP(&a_attribute_value_booleanAPI.A_ATTRIBUTE_VALUE_BOOLEAN_WOP)

	c.JSON(http.StatusOK, a_attribute_value_booleanAPI)
}

// UpdateA_ATTRIBUTE_VALUE_BOOLEAN
//
// swagger:route PATCH /a_attribute_value_booleans/{ID} a_attribute_value_booleans updateA_ATTRIBUTE_VALUE_BOOLEAN
//
// # Update a a_attribute_value_boolean
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_booleanDBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_VALUE_BOOLEAN(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_BOOLEAN.Lock()
	defer mutexA_ATTRIBUTE_VALUE_BOOLEAN.Unlock()

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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_BOOLEANAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_value_booleanDB orm.A_ATTRIBUTE_VALUE_BOOLEANDB

	// fetch the a_attribute_value_boolean
	_, err := db.First(&a_attribute_value_booleanDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_value_booleanDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_BOOLEAN_WOP(&input.A_ATTRIBUTE_VALUE_BOOLEAN_WOP)
	a_attribute_value_booleanDB.A_ATTRIBUTE_VALUE_BOOLEANPointersEncoding = input.A_ATTRIBUTE_VALUE_BOOLEANPointersEncoding

	db, _ = db.Model(&a_attribute_value_booleanDB)
	_, err = db.Updates(&a_attribute_value_booleanDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_booleanNew := new(models.A_ATTRIBUTE_VALUE_BOOLEAN)
	a_attribute_value_booleanDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_booleanNew)

	// redeem pointers
	a_attribute_value_booleanDB.DecodePointers(backRepo, a_attribute_value_booleanNew)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_booleanOld := backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.Map_A_ATTRIBUTE_VALUE_BOOLEANDBID_A_ATTRIBUTE_VALUE_BOOLEANPtr[a_attribute_value_booleanDB.ID]
	if a_attribute_value_booleanOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_booleanOld, a_attribute_value_booleanNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_booleanOld, a_attribute_value_booleanNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_value_booleanDB
	c.JSON(http.StatusOK, a_attribute_value_booleanDB)
}

// DeleteA_ATTRIBUTE_VALUE_BOOLEAN
//
// swagger:route DELETE /a_attribute_value_booleans/{ID} a_attribute_value_booleans deleteA_ATTRIBUTE_VALUE_BOOLEAN
//
// # Delete a a_attribute_value_boolean
//
// default: genericError
//
//	200: a_attribute_value_booleanDBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_VALUE_BOOLEAN(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_BOOLEAN.Lock()
	defer mutexA_ATTRIBUTE_VALUE_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_VALUE_BOOLEAN", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.GetDB()

	// Get model if exist
	var a_attribute_value_booleanDB orm.A_ATTRIBUTE_VALUE_BOOLEANDB
	if _, err := db.First(&a_attribute_value_booleanDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_value_booleanDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_booleanDeleted := new(models.A_ATTRIBUTE_VALUE_BOOLEAN)
	a_attribute_value_booleanDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_booleanDeleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_booleanStaged := backRepo.BackRepoA_ATTRIBUTE_VALUE_BOOLEAN.Map_A_ATTRIBUTE_VALUE_BOOLEANDBID_A_ATTRIBUTE_VALUE_BOOLEANPtr[a_attribute_value_booleanDB.ID]
	if a_attribute_value_booleanStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_value_booleanStaged, a_attribute_value_booleanDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
