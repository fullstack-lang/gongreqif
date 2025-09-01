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
var __A_ATTRIBUTE_VALUE_ENUMERATION__dummysDeclaration__ models.A_ATTRIBUTE_VALUE_ENUMERATION
var __A_ATTRIBUTE_VALUE_ENUMERATION_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_VALUE_ENUMERATION sync.Mutex

// An A_ATTRIBUTE_VALUE_ENUMERATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_VALUE_ENUMERATION updateA_ATTRIBUTE_VALUE_ENUMERATION deleteA_ATTRIBUTE_VALUE_ENUMERATION
type A_ATTRIBUTE_VALUE_ENUMERATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_VALUE_ENUMERATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_VALUE_ENUMERATION updateA_ATTRIBUTE_VALUE_ENUMERATION
type A_ATTRIBUTE_VALUE_ENUMERATIONInput struct {
	// The A_ATTRIBUTE_VALUE_ENUMERATION to submit or modify
	// in: body
	A_ATTRIBUTE_VALUE_ENUMERATION *orm.A_ATTRIBUTE_VALUE_ENUMERATIONAPI
}

// GetA_ATTRIBUTE_VALUE_ENUMERATIONs
//
// swagger:route GET /a_attribute_value_enumerations a_attribute_value_enumerations getA_ATTRIBUTE_VALUE_ENUMERATIONs
//
// # Get all a_attribute_value_enumerations
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_enumerationDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_ENUMERATIONs(c *gin.Context) {

	// source slice
	var a_attribute_value_enumerationDBs []orm.A_ATTRIBUTE_VALUE_ENUMERATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_ENUMERATIONs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.GetDB()

	_, err := db.Find(&a_attribute_value_enumerationDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_value_enumerationAPIs := make([]orm.A_ATTRIBUTE_VALUE_ENUMERATIONAPI, 0)

	// for each a_attribute_value_enumeration, update fields from the database nullable fields
	for idx := range a_attribute_value_enumerationDBs {
		a_attribute_value_enumerationDB := &a_attribute_value_enumerationDBs[idx]
		_ = a_attribute_value_enumerationDB
		var a_attribute_value_enumerationAPI orm.A_ATTRIBUTE_VALUE_ENUMERATIONAPI

		// insertion point for updating fields
		a_attribute_value_enumerationAPI.ID = a_attribute_value_enumerationDB.ID
		a_attribute_value_enumerationDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATION_WOP(&a_attribute_value_enumerationAPI.A_ATTRIBUTE_VALUE_ENUMERATION_WOP)
		a_attribute_value_enumerationAPI.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = a_attribute_value_enumerationDB.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
		a_attribute_value_enumerationAPIs = append(a_attribute_value_enumerationAPIs, a_attribute_value_enumerationAPI)
	}

	c.JSON(http.StatusOK, a_attribute_value_enumerationAPIs)
}

// PostA_ATTRIBUTE_VALUE_ENUMERATION
//
// swagger:route POST /a_attribute_value_enumerations a_attribute_value_enumerations postA_ATTRIBUTE_VALUE_ENUMERATION
//
// Creates a a_attribute_value_enumeration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_VALUE_ENUMERATION(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_ENUMERATION.Lock()
	defer mutexA_ATTRIBUTE_VALUE_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_VALUE_ENUMERATIONs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_ENUMERATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_value_enumeration
	a_attribute_value_enumerationDB := orm.A_ATTRIBUTE_VALUE_ENUMERATIONDB{}
	a_attribute_value_enumerationDB.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = input.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
	a_attribute_value_enumerationDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATION_WOP(&input.A_ATTRIBUTE_VALUE_ENUMERATION_WOP)

	_, err = db.Create(&a_attribute_value_enumerationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseOneInstance(&a_attribute_value_enumerationDB)
	a_attribute_value_enumeration := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr[a_attribute_value_enumerationDB.ID]

	if a_attribute_value_enumeration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_value_enumeration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_value_enumerationDB)
}

// GetA_ATTRIBUTE_VALUE_ENUMERATION
//
// swagger:route GET /a_attribute_value_enumerations/{ID} a_attribute_value_enumerations getA_ATTRIBUTE_VALUE_ENUMERATION
//
// Gets the details for a a_attribute_value_enumeration.
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_enumerationDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_ENUMERATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_ENUMERATION", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.GetDB()

	// Get a_attribute_value_enumerationDB in DB
	var a_attribute_value_enumerationDB orm.A_ATTRIBUTE_VALUE_ENUMERATIONDB
	if _, err := db.First(&a_attribute_value_enumerationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_value_enumerationAPI orm.A_ATTRIBUTE_VALUE_ENUMERATIONAPI
	a_attribute_value_enumerationAPI.ID = a_attribute_value_enumerationDB.ID
	a_attribute_value_enumerationAPI.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = a_attribute_value_enumerationDB.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
	a_attribute_value_enumerationDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATION_WOP(&a_attribute_value_enumerationAPI.A_ATTRIBUTE_VALUE_ENUMERATION_WOP)

	c.JSON(http.StatusOK, a_attribute_value_enumerationAPI)
}

// UpdateA_ATTRIBUTE_VALUE_ENUMERATION
//
// swagger:route PATCH /a_attribute_value_enumerations/{ID} a_attribute_value_enumerations updateA_ATTRIBUTE_VALUE_ENUMERATION
//
// # Update a a_attribute_value_enumeration
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_enumerationDBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_VALUE_ENUMERATION(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_ENUMERATION.Lock()
	defer mutexA_ATTRIBUTE_VALUE_ENUMERATION.Unlock()

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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_ENUMERATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_value_enumerationDB orm.A_ATTRIBUTE_VALUE_ENUMERATIONDB

	// fetch the a_attribute_value_enumeration
	_, err := db.First(&a_attribute_value_enumerationDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_value_enumerationDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATION_WOP(&input.A_ATTRIBUTE_VALUE_ENUMERATION_WOP)
	a_attribute_value_enumerationDB.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = input.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding

	db, _ = db.Model(&a_attribute_value_enumerationDB)
	_, err = db.Updates(&a_attribute_value_enumerationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_enumerationNew := new(models.A_ATTRIBUTE_VALUE_ENUMERATION)
	a_attribute_value_enumerationDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumerationNew)

	// redeem pointers
	a_attribute_value_enumerationDB.DecodePointers(backRepo, a_attribute_value_enumerationNew)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_enumerationOld := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr[a_attribute_value_enumerationDB.ID]
	if a_attribute_value_enumerationOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_enumerationOld, a_attribute_value_enumerationNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_enumerationOld, a_attribute_value_enumerationNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_value_enumerationDB
	c.JSON(http.StatusOK, a_attribute_value_enumerationDB)
}

// DeleteA_ATTRIBUTE_VALUE_ENUMERATION
//
// swagger:route DELETE /a_attribute_value_enumerations/{ID} a_attribute_value_enumerations deleteA_ATTRIBUTE_VALUE_ENUMERATION
//
// # Delete a a_attribute_value_enumeration
//
// default: genericError
//
//	200: a_attribute_value_enumerationDBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_VALUE_ENUMERATION(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_ENUMERATION.Lock()
	defer mutexA_ATTRIBUTE_VALUE_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_VALUE_ENUMERATION", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.GetDB()

	// Get model if exist
	var a_attribute_value_enumerationDB orm.A_ATTRIBUTE_VALUE_ENUMERATIONDB
	if _, err := db.First(&a_attribute_value_enumerationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_value_enumerationDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_enumerationDeleted := new(models.A_ATTRIBUTE_VALUE_ENUMERATION)
	a_attribute_value_enumerationDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumerationDeleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_enumerationStaged := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr[a_attribute_value_enumerationDB.ID]
	if a_attribute_value_enumerationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_value_enumerationStaged, a_attribute_value_enumerationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
