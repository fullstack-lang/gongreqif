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
var __DATATYPE_DEFINITION_BOOLEAN__dummysDeclaration__ models.DATATYPE_DEFINITION_BOOLEAN
var __DATATYPE_DEFINITION_BOOLEAN_time__dummyDeclaration time.Duration

var mutexDATATYPE_DEFINITION_BOOLEAN sync.Mutex

// An DATATYPE_DEFINITION_BOOLEANID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPE_DEFINITION_BOOLEAN updateDATATYPE_DEFINITION_BOOLEAN deleteDATATYPE_DEFINITION_BOOLEAN
type DATATYPE_DEFINITION_BOOLEANID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPE_DEFINITION_BOOLEANInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPE_DEFINITION_BOOLEAN updateDATATYPE_DEFINITION_BOOLEAN
type DATATYPE_DEFINITION_BOOLEANInput struct {
	// The DATATYPE_DEFINITION_BOOLEAN to submit or modify
	// in: body
	DATATYPE_DEFINITION_BOOLEAN *orm.DATATYPE_DEFINITION_BOOLEANAPI
}

// GetDATATYPE_DEFINITION_BOOLEANs
//
// swagger:route GET /datatype_definition_booleans datatype_definition_booleans getDATATYPE_DEFINITION_BOOLEANs
//
// # Get all datatype_definition_booleans
//
// Responses:
// default: genericError
//
//	200: datatype_definition_booleanDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_BOOLEANs(c *gin.Context) {

	// source slice
	var datatype_definition_booleanDBs []orm.DATATYPE_DEFINITION_BOOLEANDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_BOOLEANs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.GetDB()

	_, err := db.Find(&datatype_definition_booleanDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatype_definition_booleanAPIs := make([]orm.DATATYPE_DEFINITION_BOOLEANAPI, 0)

	// for each datatype_definition_boolean, update fields from the database nullable fields
	for idx := range datatype_definition_booleanDBs {
		datatype_definition_booleanDB := &datatype_definition_booleanDBs[idx]
		_ = datatype_definition_booleanDB
		var datatype_definition_booleanAPI orm.DATATYPE_DEFINITION_BOOLEANAPI

		// insertion point for updating fields
		datatype_definition_booleanAPI.ID = datatype_definition_booleanDB.ID
		datatype_definition_booleanDB.CopyBasicFieldsToDATATYPE_DEFINITION_BOOLEAN_WOP(&datatype_definition_booleanAPI.DATATYPE_DEFINITION_BOOLEAN_WOP)
		datatype_definition_booleanAPI.DATATYPE_DEFINITION_BOOLEANPointersEncoding = datatype_definition_booleanDB.DATATYPE_DEFINITION_BOOLEANPointersEncoding
		datatype_definition_booleanAPIs = append(datatype_definition_booleanAPIs, datatype_definition_booleanAPI)
	}

	c.JSON(http.StatusOK, datatype_definition_booleanAPIs)
}

// PostDATATYPE_DEFINITION_BOOLEAN
//
// swagger:route POST /datatype_definition_booleans datatype_definition_booleans postDATATYPE_DEFINITION_BOOLEAN
//
// Creates a datatype_definition_boolean
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPE_DEFINITION_BOOLEAN(c *gin.Context) {

	mutexDATATYPE_DEFINITION_BOOLEAN.Lock()
	defer mutexDATATYPE_DEFINITION_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPE_DEFINITION_BOOLEANs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_BOOLEANAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatype_definition_boolean
	datatype_definition_booleanDB := orm.DATATYPE_DEFINITION_BOOLEANDB{}
	datatype_definition_booleanDB.DATATYPE_DEFINITION_BOOLEANPointersEncoding = input.DATATYPE_DEFINITION_BOOLEANPointersEncoding
	datatype_definition_booleanDB.CopyBasicFieldsFromDATATYPE_DEFINITION_BOOLEAN_WOP(&input.DATATYPE_DEFINITION_BOOLEAN_WOP)

	_, err = db.Create(&datatype_definition_booleanDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.CheckoutPhaseOneInstance(&datatype_definition_booleanDB)
	datatype_definition_boolean := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.Map_DATATYPE_DEFINITION_BOOLEANDBID_DATATYPE_DEFINITION_BOOLEANPtr[datatype_definition_booleanDB.ID]

	if datatype_definition_boolean != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatype_definition_boolean)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatype_definition_booleanDB)
}

// GetDATATYPE_DEFINITION_BOOLEAN
//
// swagger:route GET /datatype_definition_booleans/{ID} datatype_definition_booleans getDATATYPE_DEFINITION_BOOLEAN
//
// Gets the details for a datatype_definition_boolean.
//
// Responses:
// default: genericError
//
//	200: datatype_definition_booleanDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_BOOLEAN(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_BOOLEAN", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.GetDB()

	// Get datatype_definition_booleanDB in DB
	var datatype_definition_booleanDB orm.DATATYPE_DEFINITION_BOOLEANDB
	if _, err := db.First(&datatype_definition_booleanDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatype_definition_booleanAPI orm.DATATYPE_DEFINITION_BOOLEANAPI
	datatype_definition_booleanAPI.ID = datatype_definition_booleanDB.ID
	datatype_definition_booleanAPI.DATATYPE_DEFINITION_BOOLEANPointersEncoding = datatype_definition_booleanDB.DATATYPE_DEFINITION_BOOLEANPointersEncoding
	datatype_definition_booleanDB.CopyBasicFieldsToDATATYPE_DEFINITION_BOOLEAN_WOP(&datatype_definition_booleanAPI.DATATYPE_DEFINITION_BOOLEAN_WOP)

	c.JSON(http.StatusOK, datatype_definition_booleanAPI)
}

// UpdateDATATYPE_DEFINITION_BOOLEAN
//
// swagger:route PATCH /datatype_definition_booleans/{ID} datatype_definition_booleans updateDATATYPE_DEFINITION_BOOLEAN
//
// # Update a datatype_definition_boolean
//
// Responses:
// default: genericError
//
//	200: datatype_definition_booleanDBResponse
func (controller *Controller) UpdateDATATYPE_DEFINITION_BOOLEAN(c *gin.Context) {

	mutexDATATYPE_DEFINITION_BOOLEAN.Lock()
	defer mutexDATATYPE_DEFINITION_BOOLEAN.Unlock()

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
	db := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_BOOLEANAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatype_definition_booleanDB orm.DATATYPE_DEFINITION_BOOLEANDB

	// fetch the datatype_definition_boolean
	_, err := db.First(&datatype_definition_booleanDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatype_definition_booleanDB.CopyBasicFieldsFromDATATYPE_DEFINITION_BOOLEAN_WOP(&input.DATATYPE_DEFINITION_BOOLEAN_WOP)
	datatype_definition_booleanDB.DATATYPE_DEFINITION_BOOLEANPointersEncoding = input.DATATYPE_DEFINITION_BOOLEANPointersEncoding

	db, _ = db.Model(&datatype_definition_booleanDB)
	_, err = db.Updates(&datatype_definition_booleanDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_booleanNew := new(models.DATATYPE_DEFINITION_BOOLEAN)
	datatype_definition_booleanDB.CopyBasicFieldsToDATATYPE_DEFINITION_BOOLEAN(datatype_definition_booleanNew)

	// redeem pointers
	datatype_definition_booleanDB.DecodePointers(backRepo, datatype_definition_booleanNew)

	// get stage instance from DB instance, and call callback function
	datatype_definition_booleanOld := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.Map_DATATYPE_DEFINITION_BOOLEANDBID_DATATYPE_DEFINITION_BOOLEANPtr[datatype_definition_booleanDB.ID]
	if datatype_definition_booleanOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), datatype_definition_booleanOld, datatype_definition_booleanNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatype_definition_booleanDB
	c.JSON(http.StatusOK, datatype_definition_booleanDB)
}

// DeleteDATATYPE_DEFINITION_BOOLEAN
//
// swagger:route DELETE /datatype_definition_booleans/{ID} datatype_definition_booleans deleteDATATYPE_DEFINITION_BOOLEAN
//
// # Delete a datatype_definition_boolean
//
// default: genericError
//
//	200: datatype_definition_booleanDBResponse
func (controller *Controller) DeleteDATATYPE_DEFINITION_BOOLEAN(c *gin.Context) {

	mutexDATATYPE_DEFINITION_BOOLEAN.Lock()
	defer mutexDATATYPE_DEFINITION_BOOLEAN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPE_DEFINITION_BOOLEAN", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.GetDB()

	// Get model if exist
	var datatype_definition_booleanDB orm.DATATYPE_DEFINITION_BOOLEANDB
	if _, err := db.First(&datatype_definition_booleanDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&datatype_definition_booleanDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_booleanDeleted := new(models.DATATYPE_DEFINITION_BOOLEAN)
	datatype_definition_booleanDB.CopyBasicFieldsToDATATYPE_DEFINITION_BOOLEAN(datatype_definition_booleanDeleted)

	// get stage instance from DB instance, and call callback function
	datatype_definition_booleanStaged := backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.Map_DATATYPE_DEFINITION_BOOLEANDBID_DATATYPE_DEFINITION_BOOLEANPtr[datatype_definition_booleanDB.ID]
	if datatype_definition_booleanStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatype_definition_booleanStaged, datatype_definition_booleanDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
