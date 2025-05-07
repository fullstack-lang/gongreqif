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
var __A_SPEC_TYPES__dummysDeclaration__ models.A_SPEC_TYPES
var __A_SPEC_TYPES_time__dummyDeclaration time.Duration

var mutexA_SPEC_TYPES sync.Mutex

// An A_SPEC_TYPESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SPEC_TYPES updateA_SPEC_TYPES deleteA_SPEC_TYPES
type A_SPEC_TYPESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SPEC_TYPESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SPEC_TYPES updateA_SPEC_TYPES
type A_SPEC_TYPESInput struct {
	// The A_SPEC_TYPES to submit or modify
	// in: body
	A_SPEC_TYPES *orm.A_SPEC_TYPESAPI
}

// GetA_SPEC_TYPESs
//
// swagger:route GET /a_spec_typess a_spec_typess getA_SPEC_TYPESs
//
// # Get all a_spec_typess
//
// Responses:
// default: genericError
//
//	200: a_spec_typesDBResponse
func (controller *Controller) GetA_SPEC_TYPESs(c *gin.Context) {

	// source slice
	var a_spec_typesDBs []orm.A_SPEC_TYPESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_TYPESs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_TYPES.GetDB()

	_, err := db.Find(&a_spec_typesDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_spec_typesAPIs := make([]orm.A_SPEC_TYPESAPI, 0)

	// for each a_spec_types, update fields from the database nullable fields
	for idx := range a_spec_typesDBs {
		a_spec_typesDB := &a_spec_typesDBs[idx]
		_ = a_spec_typesDB
		var a_spec_typesAPI orm.A_SPEC_TYPESAPI

		// insertion point for updating fields
		a_spec_typesAPI.ID = a_spec_typesDB.ID
		a_spec_typesDB.CopyBasicFieldsToA_SPEC_TYPES_WOP(&a_spec_typesAPI.A_SPEC_TYPES_WOP)
		a_spec_typesAPI.A_SPEC_TYPESPointersEncoding = a_spec_typesDB.A_SPEC_TYPESPointersEncoding
		a_spec_typesAPIs = append(a_spec_typesAPIs, a_spec_typesAPI)
	}

	c.JSON(http.StatusOK, a_spec_typesAPIs)
}

// PostA_SPEC_TYPES
//
// swagger:route POST /a_spec_typess a_spec_typess postA_SPEC_TYPES
//
// Creates a a_spec_types
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SPEC_TYPES(c *gin.Context) {

	mutexA_SPEC_TYPES.Lock()
	defer mutexA_SPEC_TYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SPEC_TYPESs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_TYPES.GetDB()

	// Validate input
	var input orm.A_SPEC_TYPESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_spec_types
	a_spec_typesDB := orm.A_SPEC_TYPESDB{}
	a_spec_typesDB.A_SPEC_TYPESPointersEncoding = input.A_SPEC_TYPESPointersEncoding
	a_spec_typesDB.CopyBasicFieldsFromA_SPEC_TYPES_WOP(&input.A_SPEC_TYPES_WOP)

	_, err = db.Create(&a_spec_typesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SPEC_TYPES.CheckoutPhaseOneInstance(&a_spec_typesDB)
	a_spec_types := backRepo.BackRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr[a_spec_typesDB.ID]

	if a_spec_types != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_spec_types)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_spec_typesDB)
}

// GetA_SPEC_TYPES
//
// swagger:route GET /a_spec_typess/{ID} a_spec_typess getA_SPEC_TYPES
//
// Gets the details for a a_spec_types.
//
// Responses:
// default: genericError
//
//	200: a_spec_typesDBResponse
func (controller *Controller) GetA_SPEC_TYPES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_TYPES", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_TYPES.GetDB()

	// Get a_spec_typesDB in DB
	var a_spec_typesDB orm.A_SPEC_TYPESDB
	if _, err := db.First(&a_spec_typesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_spec_typesAPI orm.A_SPEC_TYPESAPI
	a_spec_typesAPI.ID = a_spec_typesDB.ID
	a_spec_typesAPI.A_SPEC_TYPESPointersEncoding = a_spec_typesDB.A_SPEC_TYPESPointersEncoding
	a_spec_typesDB.CopyBasicFieldsToA_SPEC_TYPES_WOP(&a_spec_typesAPI.A_SPEC_TYPES_WOP)

	c.JSON(http.StatusOK, a_spec_typesAPI)
}

// UpdateA_SPEC_TYPES
//
// swagger:route PATCH /a_spec_typess/{ID} a_spec_typess updateA_SPEC_TYPES
//
// # Update a a_spec_types
//
// Responses:
// default: genericError
//
//	200: a_spec_typesDBResponse
func (controller *Controller) UpdateA_SPEC_TYPES(c *gin.Context) {

	mutexA_SPEC_TYPES.Lock()
	defer mutexA_SPEC_TYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_SPEC_TYPES", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_TYPES.GetDB()

	// Validate input
	var input orm.A_SPEC_TYPESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_spec_typesDB orm.A_SPEC_TYPESDB

	// fetch the a_spec_types
	_, err := db.First(&a_spec_typesDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_spec_typesDB.CopyBasicFieldsFromA_SPEC_TYPES_WOP(&input.A_SPEC_TYPES_WOP)
	a_spec_typesDB.A_SPEC_TYPESPointersEncoding = input.A_SPEC_TYPESPointersEncoding

	db, _ = db.Model(&a_spec_typesDB)
	_, err = db.Updates(&a_spec_typesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_typesNew := new(models.A_SPEC_TYPES)
	a_spec_typesDB.CopyBasicFieldsToA_SPEC_TYPES(a_spec_typesNew)

	// redeem pointers
	a_spec_typesDB.DecodePointers(backRepo, a_spec_typesNew)

	// get stage instance from DB instance, and call callback function
	a_spec_typesOld := backRepo.BackRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr[a_spec_typesDB.ID]
	if a_spec_typesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_spec_typesOld, a_spec_typesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_spec_typesDB
	c.JSON(http.StatusOK, a_spec_typesDB)
}

// DeleteA_SPEC_TYPES
//
// swagger:route DELETE /a_spec_typess/{ID} a_spec_typess deleteA_SPEC_TYPES
//
// # Delete a a_spec_types
//
// default: genericError
//
//	200: a_spec_typesDBResponse
func (controller *Controller) DeleteA_SPEC_TYPES(c *gin.Context) {

	mutexA_SPEC_TYPES.Lock()
	defer mutexA_SPEC_TYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SPEC_TYPES", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_TYPES.GetDB()

	// Get model if exist
	var a_spec_typesDB orm.A_SPEC_TYPESDB
	if _, err := db.First(&a_spec_typesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_spec_typesDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_typesDeleted := new(models.A_SPEC_TYPES)
	a_spec_typesDB.CopyBasicFieldsToA_SPEC_TYPES(a_spec_typesDeleted)

	// get stage instance from DB instance, and call callback function
	a_spec_typesStaged := backRepo.BackRepoA_SPEC_TYPES.Map_A_SPEC_TYPESDBID_A_SPEC_TYPESPtr[a_spec_typesDB.ID]
	if a_spec_typesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_spec_typesStaged, a_spec_typesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
