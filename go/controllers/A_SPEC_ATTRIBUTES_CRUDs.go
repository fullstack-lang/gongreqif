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
var __A_SPEC_ATTRIBUTES__dummysDeclaration__ models.A_SPEC_ATTRIBUTES
var __A_SPEC_ATTRIBUTES_time__dummyDeclaration time.Duration

var mutexA_SPEC_ATTRIBUTES sync.Mutex

// An A_SPEC_ATTRIBUTESID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SPEC_ATTRIBUTES updateA_SPEC_ATTRIBUTES deleteA_SPEC_ATTRIBUTES
type A_SPEC_ATTRIBUTESID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SPEC_ATTRIBUTESInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SPEC_ATTRIBUTES updateA_SPEC_ATTRIBUTES
type A_SPEC_ATTRIBUTESInput struct {
	// The A_SPEC_ATTRIBUTES to submit or modify
	// in: body
	A_SPEC_ATTRIBUTES *orm.A_SPEC_ATTRIBUTESAPI
}

// GetA_SPEC_ATTRIBUTESs
//
// swagger:route GET /a_spec_attributess a_spec_attributess getA_SPEC_ATTRIBUTESs
//
// # Get all a_spec_attributess
//
// Responses:
// default: genericError
//
//	200: a_spec_attributesDBResponse
func (controller *Controller) GetA_SPEC_ATTRIBUTESs(c *gin.Context) {

	// source slice
	var a_spec_attributesDBs []orm.A_SPEC_ATTRIBUTESDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_ATTRIBUTESs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_ATTRIBUTES.GetDB()

	_, err := db.Find(&a_spec_attributesDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_spec_attributesAPIs := make([]orm.A_SPEC_ATTRIBUTESAPI, 0)

	// for each a_spec_attributes, update fields from the database nullable fields
	for idx := range a_spec_attributesDBs {
		a_spec_attributesDB := &a_spec_attributesDBs[idx]
		_ = a_spec_attributesDB
		var a_spec_attributesAPI orm.A_SPEC_ATTRIBUTESAPI

		// insertion point for updating fields
		a_spec_attributesAPI.ID = a_spec_attributesDB.ID
		a_spec_attributesDB.CopyBasicFieldsToA_SPEC_ATTRIBUTES_WOP(&a_spec_attributesAPI.A_SPEC_ATTRIBUTES_WOP)
		a_spec_attributesAPI.A_SPEC_ATTRIBUTESPointersEncoding = a_spec_attributesDB.A_SPEC_ATTRIBUTESPointersEncoding
		a_spec_attributesAPIs = append(a_spec_attributesAPIs, a_spec_attributesAPI)
	}

	c.JSON(http.StatusOK, a_spec_attributesAPIs)
}

// PostA_SPEC_ATTRIBUTES
//
// swagger:route POST /a_spec_attributess a_spec_attributess postA_SPEC_ATTRIBUTES
//
// Creates a a_spec_attributes
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SPEC_ATTRIBUTES(c *gin.Context) {

	mutexA_SPEC_ATTRIBUTES.Lock()
	defer mutexA_SPEC_ATTRIBUTES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SPEC_ATTRIBUTESs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_ATTRIBUTES.GetDB()

	// Validate input
	var input orm.A_SPEC_ATTRIBUTESAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_spec_attributes
	a_spec_attributesDB := orm.A_SPEC_ATTRIBUTESDB{}
	a_spec_attributesDB.A_SPEC_ATTRIBUTESPointersEncoding = input.A_SPEC_ATTRIBUTESPointersEncoding
	a_spec_attributesDB.CopyBasicFieldsFromA_SPEC_ATTRIBUTES_WOP(&input.A_SPEC_ATTRIBUTES_WOP)

	_, err = db.Create(&a_spec_attributesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SPEC_ATTRIBUTES.CheckoutPhaseOneInstance(&a_spec_attributesDB)
	a_spec_attributes := backRepo.BackRepoA_SPEC_ATTRIBUTES.Map_A_SPEC_ATTRIBUTESDBID_A_SPEC_ATTRIBUTESPtr[a_spec_attributesDB.ID]

	if a_spec_attributes != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_spec_attributes)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_spec_attributesDB)
}

// GetA_SPEC_ATTRIBUTES
//
// swagger:route GET /a_spec_attributess/{ID} a_spec_attributess getA_SPEC_ATTRIBUTES
//
// Gets the details for a a_spec_attributes.
//
// Responses:
// default: genericError
//
//	200: a_spec_attributesDBResponse
func (controller *Controller) GetA_SPEC_ATTRIBUTES(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_ATTRIBUTES", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_ATTRIBUTES.GetDB()

	// Get a_spec_attributesDB in DB
	var a_spec_attributesDB orm.A_SPEC_ATTRIBUTESDB
	if _, err := db.First(&a_spec_attributesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_spec_attributesAPI orm.A_SPEC_ATTRIBUTESAPI
	a_spec_attributesAPI.ID = a_spec_attributesDB.ID
	a_spec_attributesAPI.A_SPEC_ATTRIBUTESPointersEncoding = a_spec_attributesDB.A_SPEC_ATTRIBUTESPointersEncoding
	a_spec_attributesDB.CopyBasicFieldsToA_SPEC_ATTRIBUTES_WOP(&a_spec_attributesAPI.A_SPEC_ATTRIBUTES_WOP)

	c.JSON(http.StatusOK, a_spec_attributesAPI)
}

// UpdateA_SPEC_ATTRIBUTES
//
// swagger:route PATCH /a_spec_attributess/{ID} a_spec_attributess updateA_SPEC_ATTRIBUTES
//
// # Update a a_spec_attributes
//
// Responses:
// default: genericError
//
//	200: a_spec_attributesDBResponse
func (controller *Controller) UpdateA_SPEC_ATTRIBUTES(c *gin.Context) {

	mutexA_SPEC_ATTRIBUTES.Lock()
	defer mutexA_SPEC_ATTRIBUTES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_SPEC_ATTRIBUTES", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_ATTRIBUTES.GetDB()

	// Validate input
	var input orm.A_SPEC_ATTRIBUTESAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_spec_attributesDB orm.A_SPEC_ATTRIBUTESDB

	// fetch the a_spec_attributes
	_, err := db.First(&a_spec_attributesDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_spec_attributesDB.CopyBasicFieldsFromA_SPEC_ATTRIBUTES_WOP(&input.A_SPEC_ATTRIBUTES_WOP)
	a_spec_attributesDB.A_SPEC_ATTRIBUTESPointersEncoding = input.A_SPEC_ATTRIBUTESPointersEncoding

	db, _ = db.Model(&a_spec_attributesDB)
	_, err = db.Updates(&a_spec_attributesDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_attributesNew := new(models.A_SPEC_ATTRIBUTES)
	a_spec_attributesDB.CopyBasicFieldsToA_SPEC_ATTRIBUTES(a_spec_attributesNew)

	// redeem pointers
	a_spec_attributesDB.DecodePointers(backRepo, a_spec_attributesNew)

	// get stage instance from DB instance, and call callback function
	a_spec_attributesOld := backRepo.BackRepoA_SPEC_ATTRIBUTES.Map_A_SPEC_ATTRIBUTESDBID_A_SPEC_ATTRIBUTESPtr[a_spec_attributesDB.ID]
	if a_spec_attributesOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_spec_attributesOld, a_spec_attributesNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_spec_attributesDB
	c.JSON(http.StatusOK, a_spec_attributesDB)
}

// DeleteA_SPEC_ATTRIBUTES
//
// swagger:route DELETE /a_spec_attributess/{ID} a_spec_attributess deleteA_SPEC_ATTRIBUTES
//
// # Delete a a_spec_attributes
//
// default: genericError
//
//	200: a_spec_attributesDBResponse
func (controller *Controller) DeleteA_SPEC_ATTRIBUTES(c *gin.Context) {

	mutexA_SPEC_ATTRIBUTES.Lock()
	defer mutexA_SPEC_ATTRIBUTES.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SPEC_ATTRIBUTES", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_ATTRIBUTES.GetDB()

	// Get model if exist
	var a_spec_attributesDB orm.A_SPEC_ATTRIBUTESDB
	if _, err := db.First(&a_spec_attributesDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_spec_attributesDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_attributesDeleted := new(models.A_SPEC_ATTRIBUTES)
	a_spec_attributesDB.CopyBasicFieldsToA_SPEC_ATTRIBUTES(a_spec_attributesDeleted)

	// get stage instance from DB instance, and call callback function
	a_spec_attributesStaged := backRepo.BackRepoA_SPEC_ATTRIBUTES.Map_A_SPEC_ATTRIBUTESDBID_A_SPEC_ATTRIBUTESPtr[a_spec_attributesDB.ID]
	if a_spec_attributesStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_spec_attributesStaged, a_spec_attributesDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
