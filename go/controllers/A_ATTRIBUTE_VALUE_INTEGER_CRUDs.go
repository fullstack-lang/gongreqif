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
var __A_ATTRIBUTE_VALUE_INTEGER__dummysDeclaration__ models.A_ATTRIBUTE_VALUE_INTEGER
var __A_ATTRIBUTE_VALUE_INTEGER_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_VALUE_INTEGER sync.Mutex

// An A_ATTRIBUTE_VALUE_INTEGERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_VALUE_INTEGER updateA_ATTRIBUTE_VALUE_INTEGER deleteA_ATTRIBUTE_VALUE_INTEGER
type A_ATTRIBUTE_VALUE_INTEGERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_VALUE_INTEGERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_VALUE_INTEGER updateA_ATTRIBUTE_VALUE_INTEGER
type A_ATTRIBUTE_VALUE_INTEGERInput struct {
	// The A_ATTRIBUTE_VALUE_INTEGER to submit or modify
	// in: body
	A_ATTRIBUTE_VALUE_INTEGER *orm.A_ATTRIBUTE_VALUE_INTEGERAPI
}

// GetA_ATTRIBUTE_VALUE_INTEGERs
//
// swagger:route GET /a_attribute_value_integers a_attribute_value_integers getA_ATTRIBUTE_VALUE_INTEGERs
//
// # Get all a_attribute_value_integers
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_integerDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_INTEGERs(c *gin.Context) {

	// source slice
	var a_attribute_value_integerDBs []orm.A_ATTRIBUTE_VALUE_INTEGERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_INTEGERs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.GetDB()

	_, err := db.Find(&a_attribute_value_integerDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_value_integerAPIs := make([]orm.A_ATTRIBUTE_VALUE_INTEGERAPI, 0)

	// for each a_attribute_value_integer, update fields from the database nullable fields
	for idx := range a_attribute_value_integerDBs {
		a_attribute_value_integerDB := &a_attribute_value_integerDBs[idx]
		_ = a_attribute_value_integerDB
		var a_attribute_value_integerAPI orm.A_ATTRIBUTE_VALUE_INTEGERAPI

		// insertion point for updating fields
		a_attribute_value_integerAPI.ID = a_attribute_value_integerDB.ID
		a_attribute_value_integerDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_INTEGER_WOP(&a_attribute_value_integerAPI.A_ATTRIBUTE_VALUE_INTEGER_WOP)
		a_attribute_value_integerAPI.A_ATTRIBUTE_VALUE_INTEGERPointersEncoding = a_attribute_value_integerDB.A_ATTRIBUTE_VALUE_INTEGERPointersEncoding
		a_attribute_value_integerAPIs = append(a_attribute_value_integerAPIs, a_attribute_value_integerAPI)
	}

	c.JSON(http.StatusOK, a_attribute_value_integerAPIs)
}

// PostA_ATTRIBUTE_VALUE_INTEGER
//
// swagger:route POST /a_attribute_value_integers a_attribute_value_integers postA_ATTRIBUTE_VALUE_INTEGER
//
// Creates a a_attribute_value_integer
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_VALUE_INTEGER(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_INTEGER.Lock()
	defer mutexA_ATTRIBUTE_VALUE_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_VALUE_INTEGERs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_INTEGERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_value_integer
	a_attribute_value_integerDB := orm.A_ATTRIBUTE_VALUE_INTEGERDB{}
	a_attribute_value_integerDB.A_ATTRIBUTE_VALUE_INTEGERPointersEncoding = input.A_ATTRIBUTE_VALUE_INTEGERPointersEncoding
	a_attribute_value_integerDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_INTEGER_WOP(&input.A_ATTRIBUTE_VALUE_INTEGER_WOP)

	_, err = db.Create(&a_attribute_value_integerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.CheckoutPhaseOneInstance(&a_attribute_value_integerDB)
	a_attribute_value_integer := backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.Map_A_ATTRIBUTE_VALUE_INTEGERDBID_A_ATTRIBUTE_VALUE_INTEGERPtr[a_attribute_value_integerDB.ID]

	if a_attribute_value_integer != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_value_integer)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_value_integerDB)
}

// GetA_ATTRIBUTE_VALUE_INTEGER
//
// swagger:route GET /a_attribute_value_integers/{ID} a_attribute_value_integers getA_ATTRIBUTE_VALUE_INTEGER
//
// Gets the details for a a_attribute_value_integer.
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_integerDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_INTEGER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_INTEGER", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.GetDB()

	// Get a_attribute_value_integerDB in DB
	var a_attribute_value_integerDB orm.A_ATTRIBUTE_VALUE_INTEGERDB
	if _, err := db.First(&a_attribute_value_integerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_value_integerAPI orm.A_ATTRIBUTE_VALUE_INTEGERAPI
	a_attribute_value_integerAPI.ID = a_attribute_value_integerDB.ID
	a_attribute_value_integerAPI.A_ATTRIBUTE_VALUE_INTEGERPointersEncoding = a_attribute_value_integerDB.A_ATTRIBUTE_VALUE_INTEGERPointersEncoding
	a_attribute_value_integerDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_INTEGER_WOP(&a_attribute_value_integerAPI.A_ATTRIBUTE_VALUE_INTEGER_WOP)

	c.JSON(http.StatusOK, a_attribute_value_integerAPI)
}

// UpdateA_ATTRIBUTE_VALUE_INTEGER
//
// swagger:route PATCH /a_attribute_value_integers/{ID} a_attribute_value_integers updateA_ATTRIBUTE_VALUE_INTEGER
//
// # Update a a_attribute_value_integer
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_integerDBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_VALUE_INTEGER(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_INTEGER.Lock()
	defer mutexA_ATTRIBUTE_VALUE_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_INTEGERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_value_integerDB orm.A_ATTRIBUTE_VALUE_INTEGERDB

	// fetch the a_attribute_value_integer
	_, err := db.First(&a_attribute_value_integerDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_value_integerDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_INTEGER_WOP(&input.A_ATTRIBUTE_VALUE_INTEGER_WOP)
	a_attribute_value_integerDB.A_ATTRIBUTE_VALUE_INTEGERPointersEncoding = input.A_ATTRIBUTE_VALUE_INTEGERPointersEncoding

	db, _ = db.Model(&a_attribute_value_integerDB)
	_, err = db.Updates(&a_attribute_value_integerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_integerNew := new(models.A_ATTRIBUTE_VALUE_INTEGER)
	a_attribute_value_integerDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integerNew)

	// redeem pointers
	a_attribute_value_integerDB.DecodePointers(backRepo, a_attribute_value_integerNew)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_integerOld := backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.Map_A_ATTRIBUTE_VALUE_INTEGERDBID_A_ATTRIBUTE_VALUE_INTEGERPtr[a_attribute_value_integerDB.ID]
	if a_attribute_value_integerOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_integerOld, a_attribute_value_integerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_value_integerDB
	c.JSON(http.StatusOK, a_attribute_value_integerDB)
}

// DeleteA_ATTRIBUTE_VALUE_INTEGER
//
// swagger:route DELETE /a_attribute_value_integers/{ID} a_attribute_value_integers deleteA_ATTRIBUTE_VALUE_INTEGER
//
// # Delete a a_attribute_value_integer
//
// default: genericError
//
//	200: a_attribute_value_integerDBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_VALUE_INTEGER(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_INTEGER.Lock()
	defer mutexA_ATTRIBUTE_VALUE_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_VALUE_INTEGER", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.GetDB()

	// Get model if exist
	var a_attribute_value_integerDB orm.A_ATTRIBUTE_VALUE_INTEGERDB
	if _, err := db.First(&a_attribute_value_integerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_value_integerDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_integerDeleted := new(models.A_ATTRIBUTE_VALUE_INTEGER)
	a_attribute_value_integerDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integerDeleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_integerStaged := backRepo.BackRepoA_ATTRIBUTE_VALUE_INTEGER.Map_A_ATTRIBUTE_VALUE_INTEGERDBID_A_ATTRIBUTE_VALUE_INTEGERPtr[a_attribute_value_integerDB.ID]
	if a_attribute_value_integerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_value_integerStaged, a_attribute_value_integerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
