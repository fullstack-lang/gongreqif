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
var __A_DATATYPES__dummysDeclaration__ models.A_DATATYPES
var __A_DATATYPES_time__dummyDeclaration time.Duration

var mutexA_DATATYPES sync.Mutex

// An A_DATATYPESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DATATYPES updateA_DATATYPES deleteA_DATATYPES
type A_DATATYPESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DATATYPESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DATATYPES updateA_DATATYPES
type A_DATATYPESInput struct {
	// The A_DATATYPES to submit or modify
	// in: body
	A_DATATYPES *orm.A_DATATYPESAPI
}

// GetA_DATATYPESs
//
// swagger:route GET /a_datatypess a_datatypess getA_DATATYPESs
//
// # Get all a_datatypess
//
// Responses:
// default: genericError
//
//	200: a_datatypesDBResponse
func (controller *Controller) GetA_DATATYPESs(c *gin.Context) {

	// source slice
	var a_datatypesDBs []orm.A_DATATYPESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DATATYPESs", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPES.GetDB()

	_, err := db.Find(&a_datatypesDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_datatypesAPIs := make([]orm.A_DATATYPESAPI, 0)

	// for each a_datatypes, update fields from the database nullable fields
	for idx := range a_datatypesDBs {
		a_datatypesDB := &a_datatypesDBs[idx]
		_ = a_datatypesDB
		var a_datatypesAPI orm.A_DATATYPESAPI

		// insertion point for updating fields
		a_datatypesAPI.ID = a_datatypesDB.ID
		a_datatypesDB.CopyBasicFieldsToA_DATATYPES_WOP(&a_datatypesAPI.A_DATATYPES_WOP)
		a_datatypesAPI.A_DATATYPESPointersEncoding = a_datatypesDB.A_DATATYPESPointersEncoding
		a_datatypesAPIs = append(a_datatypesAPIs, a_datatypesAPI)
	}

	c.JSON(http.StatusOK, a_datatypesAPIs)
}

// PostA_DATATYPES
//
// swagger:route POST /a_datatypess a_datatypess postA_DATATYPES
//
// Creates a a_datatypes
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DATATYPES(c *gin.Context) {

	mutexA_DATATYPES.Lock()
	defer mutexA_DATATYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DATATYPESs", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPES.GetDB()

	// Validate input
	var input orm.A_DATATYPESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_datatypes
	a_datatypesDB := orm.A_DATATYPESDB{}
	a_datatypesDB.A_DATATYPESPointersEncoding = input.A_DATATYPESPointersEncoding
	a_datatypesDB.CopyBasicFieldsFromA_DATATYPES_WOP(&input.A_DATATYPES_WOP)

	_, err = db.Create(&a_datatypesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DATATYPES.CheckoutPhaseOneInstance(&a_datatypesDB)
	a_datatypes := backRepo.BackRepoA_DATATYPES.Map_A_DATATYPESDBID_A_DATATYPESPtr[a_datatypesDB.ID]

	if a_datatypes != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_datatypes)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_datatypesDB)
}

// GetA_DATATYPES
//
// swagger:route GET /a_datatypess/{ID} a_datatypess getA_DATATYPES
//
// Gets the details for a a_datatypes.
//
// Responses:
// default: genericError
//
//	200: a_datatypesDBResponse
func (controller *Controller) GetA_DATATYPES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DATATYPES", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPES.GetDB()

	// Get a_datatypesDB in DB
	var a_datatypesDB orm.A_DATATYPESDB
	if _, err := db.First(&a_datatypesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_datatypesAPI orm.A_DATATYPESAPI
	a_datatypesAPI.ID = a_datatypesDB.ID
	a_datatypesAPI.A_DATATYPESPointersEncoding = a_datatypesDB.A_DATATYPESPointersEncoding
	a_datatypesDB.CopyBasicFieldsToA_DATATYPES_WOP(&a_datatypesAPI.A_DATATYPES_WOP)

	c.JSON(http.StatusOK, a_datatypesAPI)
}

// UpdateA_DATATYPES
//
// swagger:route PATCH /a_datatypess/{ID} a_datatypess updateA_DATATYPES
//
// # Update a a_datatypes
//
// Responses:
// default: genericError
//
//	200: a_datatypesDBResponse
func (controller *Controller) UpdateA_DATATYPES(c *gin.Context) {

	mutexA_DATATYPES.Lock()
	defer mutexA_DATATYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DATATYPES", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPES.GetDB()

	// Validate input
	var input orm.A_DATATYPESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_datatypesDB orm.A_DATATYPESDB

	// fetch the a_datatypes
	_, err := db.First(&a_datatypesDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_datatypesDB.CopyBasicFieldsFromA_DATATYPES_WOP(&input.A_DATATYPES_WOP)
	a_datatypesDB.A_DATATYPESPointersEncoding = input.A_DATATYPESPointersEncoding

	db, _ = db.Model(&a_datatypesDB)
	_, err = db.Updates(&a_datatypesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_datatypesNew := new(models.A_DATATYPES)
	a_datatypesDB.CopyBasicFieldsToA_DATATYPES(a_datatypesNew)

	// redeem pointers
	a_datatypesDB.DecodePointers(backRepo, a_datatypesNew)

	// get stage instance from DB instance, and call callback function
	a_datatypesOld := backRepo.BackRepoA_DATATYPES.Map_A_DATATYPESDBID_A_DATATYPESPtr[a_datatypesDB.ID]
	if a_datatypesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_datatypesOld, a_datatypesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_datatypesDB
	c.JSON(http.StatusOK, a_datatypesDB)
}

// DeleteA_DATATYPES
//
// swagger:route DELETE /a_datatypess/{ID} a_datatypess deleteA_DATATYPES
//
// # Delete a a_datatypes
//
// default: genericError
//
//	200: a_datatypesDBResponse
func (controller *Controller) DeleteA_DATATYPES(c *gin.Context) {

	mutexA_DATATYPES.Lock()
	defer mutexA_DATATYPES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DATATYPES", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPES.GetDB()

	// Get model if exist
	var a_datatypesDB orm.A_DATATYPESDB
	if _, err := db.First(&a_datatypesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_datatypesDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_datatypesDeleted := new(models.A_DATATYPES)
	a_datatypesDB.CopyBasicFieldsToA_DATATYPES(a_datatypesDeleted)

	// get stage instance from DB instance, and call callback function
	a_datatypesStaged := backRepo.BackRepoA_DATATYPES.Map_A_DATATYPESDBID_A_DATATYPESPtr[a_datatypesDB.ID]
	if a_datatypesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_datatypesStaged, a_datatypesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
