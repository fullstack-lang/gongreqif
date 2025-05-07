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
var __A_SPECIFIED_VALUES__dummysDeclaration__ models.A_SPECIFIED_VALUES
var __A_SPECIFIED_VALUES_time__dummyDeclaration time.Duration

var mutexA_SPECIFIED_VALUES sync.Mutex

// An A_SPECIFIED_VALUESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SPECIFIED_VALUES updateA_SPECIFIED_VALUES deleteA_SPECIFIED_VALUES
type A_SPECIFIED_VALUESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SPECIFIED_VALUESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SPECIFIED_VALUES updateA_SPECIFIED_VALUES
type A_SPECIFIED_VALUESInput struct {
	// The A_SPECIFIED_VALUES to submit or modify
	// in: body
	A_SPECIFIED_VALUES *orm.A_SPECIFIED_VALUESAPI
}

// GetA_SPECIFIED_VALUESs
//
// swagger:route GET /a_specified_valuess a_specified_valuess getA_SPECIFIED_VALUESs
//
// # Get all a_specified_valuess
//
// Responses:
// default: genericError
//
//	200: a_specified_valuesDBResponse
func (controller *Controller) GetA_SPECIFIED_VALUESs(c *gin.Context) {

	// source slice
	var a_specified_valuesDBs []orm.A_SPECIFIED_VALUESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPECIFIED_VALUESs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPECIFIED_VALUES.GetDB()

	_, err := db.Find(&a_specified_valuesDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_specified_valuesAPIs := make([]orm.A_SPECIFIED_VALUESAPI, 0)

	// for each a_specified_values, update fields from the database nullable fields
	for idx := range a_specified_valuesDBs {
		a_specified_valuesDB := &a_specified_valuesDBs[idx]
		_ = a_specified_valuesDB
		var a_specified_valuesAPI orm.A_SPECIFIED_VALUESAPI

		// insertion point for updating fields
		a_specified_valuesAPI.ID = a_specified_valuesDB.ID
		a_specified_valuesDB.CopyBasicFieldsToA_SPECIFIED_VALUES_WOP(&a_specified_valuesAPI.A_SPECIFIED_VALUES_WOP)
		a_specified_valuesAPI.A_SPECIFIED_VALUESPointersEncoding = a_specified_valuesDB.A_SPECIFIED_VALUESPointersEncoding
		a_specified_valuesAPIs = append(a_specified_valuesAPIs, a_specified_valuesAPI)
	}

	c.JSON(http.StatusOK, a_specified_valuesAPIs)
}

// PostA_SPECIFIED_VALUES
//
// swagger:route POST /a_specified_valuess a_specified_valuess postA_SPECIFIED_VALUES
//
// Creates a a_specified_values
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SPECIFIED_VALUES(c *gin.Context) {

	mutexA_SPECIFIED_VALUES.Lock()
	defer mutexA_SPECIFIED_VALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SPECIFIED_VALUESs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPECIFIED_VALUES.GetDB()

	// Validate input
	var input orm.A_SPECIFIED_VALUESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_specified_values
	a_specified_valuesDB := orm.A_SPECIFIED_VALUESDB{}
	a_specified_valuesDB.A_SPECIFIED_VALUESPointersEncoding = input.A_SPECIFIED_VALUESPointersEncoding
	a_specified_valuesDB.CopyBasicFieldsFromA_SPECIFIED_VALUES_WOP(&input.A_SPECIFIED_VALUES_WOP)

	_, err = db.Create(&a_specified_valuesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SPECIFIED_VALUES.CheckoutPhaseOneInstance(&a_specified_valuesDB)
	a_specified_values := backRepo.BackRepoA_SPECIFIED_VALUES.Map_A_SPECIFIED_VALUESDBID_A_SPECIFIED_VALUESPtr[a_specified_valuesDB.ID]

	if a_specified_values != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_specified_values)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_specified_valuesDB)
}

// GetA_SPECIFIED_VALUES
//
// swagger:route GET /a_specified_valuess/{ID} a_specified_valuess getA_SPECIFIED_VALUES
//
// Gets the details for a a_specified_values.
//
// Responses:
// default: genericError
//
//	200: a_specified_valuesDBResponse
func (controller *Controller) GetA_SPECIFIED_VALUES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPECIFIED_VALUES", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPECIFIED_VALUES.GetDB()

	// Get a_specified_valuesDB in DB
	var a_specified_valuesDB orm.A_SPECIFIED_VALUESDB
	if _, err := db.First(&a_specified_valuesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_specified_valuesAPI orm.A_SPECIFIED_VALUESAPI
	a_specified_valuesAPI.ID = a_specified_valuesDB.ID
	a_specified_valuesAPI.A_SPECIFIED_VALUESPointersEncoding = a_specified_valuesDB.A_SPECIFIED_VALUESPointersEncoding
	a_specified_valuesDB.CopyBasicFieldsToA_SPECIFIED_VALUES_WOP(&a_specified_valuesAPI.A_SPECIFIED_VALUES_WOP)

	c.JSON(http.StatusOK, a_specified_valuesAPI)
}

// UpdateA_SPECIFIED_VALUES
//
// swagger:route PATCH /a_specified_valuess/{ID} a_specified_valuess updateA_SPECIFIED_VALUES
//
// # Update a a_specified_values
//
// Responses:
// default: genericError
//
//	200: a_specified_valuesDBResponse
func (controller *Controller) UpdateA_SPECIFIED_VALUES(c *gin.Context) {

	mutexA_SPECIFIED_VALUES.Lock()
	defer mutexA_SPECIFIED_VALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_SPECIFIED_VALUES", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPECIFIED_VALUES.GetDB()

	// Validate input
	var input orm.A_SPECIFIED_VALUESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_specified_valuesDB orm.A_SPECIFIED_VALUESDB

	// fetch the a_specified_values
	_, err := db.First(&a_specified_valuesDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_specified_valuesDB.CopyBasicFieldsFromA_SPECIFIED_VALUES_WOP(&input.A_SPECIFIED_VALUES_WOP)
	a_specified_valuesDB.A_SPECIFIED_VALUESPointersEncoding = input.A_SPECIFIED_VALUESPointersEncoding

	db, _ = db.Model(&a_specified_valuesDB)
	_, err = db.Updates(&a_specified_valuesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_specified_valuesNew := new(models.A_SPECIFIED_VALUES)
	a_specified_valuesDB.CopyBasicFieldsToA_SPECIFIED_VALUES(a_specified_valuesNew)

	// redeem pointers
	a_specified_valuesDB.DecodePointers(backRepo, a_specified_valuesNew)

	// get stage instance from DB instance, and call callback function
	a_specified_valuesOld := backRepo.BackRepoA_SPECIFIED_VALUES.Map_A_SPECIFIED_VALUESDBID_A_SPECIFIED_VALUESPtr[a_specified_valuesDB.ID]
	if a_specified_valuesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_specified_valuesOld, a_specified_valuesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_specified_valuesDB
	c.JSON(http.StatusOK, a_specified_valuesDB)
}

// DeleteA_SPECIFIED_VALUES
//
// swagger:route DELETE /a_specified_valuess/{ID} a_specified_valuess deleteA_SPECIFIED_VALUES
//
// # Delete a a_specified_values
//
// default: genericError
//
//	200: a_specified_valuesDBResponse
func (controller *Controller) DeleteA_SPECIFIED_VALUES(c *gin.Context) {

	mutexA_SPECIFIED_VALUES.Lock()
	defer mutexA_SPECIFIED_VALUES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SPECIFIED_VALUES", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPECIFIED_VALUES.GetDB()

	// Get model if exist
	var a_specified_valuesDB orm.A_SPECIFIED_VALUESDB
	if _, err := db.First(&a_specified_valuesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_specified_valuesDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_specified_valuesDeleted := new(models.A_SPECIFIED_VALUES)
	a_specified_valuesDB.CopyBasicFieldsToA_SPECIFIED_VALUES(a_specified_valuesDeleted)

	// get stage instance from DB instance, and call callback function
	a_specified_valuesStaged := backRepo.BackRepoA_SPECIFIED_VALUES.Map_A_SPECIFIED_VALUESDBID_A_SPECIFIED_VALUESPtr[a_specified_valuesDB.ID]
	if a_specified_valuesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_specified_valuesStaged, a_specified_valuesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
