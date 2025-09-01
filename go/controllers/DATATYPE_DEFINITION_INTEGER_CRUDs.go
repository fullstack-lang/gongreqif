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
var __DATATYPE_DEFINITION_INTEGER__dummysDeclaration__ models.DATATYPE_DEFINITION_INTEGER
var __DATATYPE_DEFINITION_INTEGER_time__dummyDeclaration time.Duration

var mutexDATATYPE_DEFINITION_INTEGER sync.Mutex

// An DATATYPE_DEFINITION_INTEGERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPE_DEFINITION_INTEGER updateDATATYPE_DEFINITION_INTEGER deleteDATATYPE_DEFINITION_INTEGER
type DATATYPE_DEFINITION_INTEGERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPE_DEFINITION_INTEGERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPE_DEFINITION_INTEGER updateDATATYPE_DEFINITION_INTEGER
type DATATYPE_DEFINITION_INTEGERInput struct {
	// The DATATYPE_DEFINITION_INTEGER to submit or modify
	// in: body
	DATATYPE_DEFINITION_INTEGER *orm.DATATYPE_DEFINITION_INTEGERAPI
}

// GetDATATYPE_DEFINITION_INTEGERs
//
// swagger:route GET /datatype_definition_integers datatype_definition_integers getDATATYPE_DEFINITION_INTEGERs
//
// # Get all datatype_definition_integers
//
// Responses:
// default: genericError
//
//	200: datatype_definition_integerDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_INTEGERs(c *gin.Context) {

	// source slice
	var datatype_definition_integerDBs []orm.DATATYPE_DEFINITION_INTEGERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_INTEGERs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.GetDB()

	_, err := db.Find(&datatype_definition_integerDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatype_definition_integerAPIs := make([]orm.DATATYPE_DEFINITION_INTEGERAPI, 0)

	// for each datatype_definition_integer, update fields from the database nullable fields
	for idx := range datatype_definition_integerDBs {
		datatype_definition_integerDB := &datatype_definition_integerDBs[idx]
		_ = datatype_definition_integerDB
		var datatype_definition_integerAPI orm.DATATYPE_DEFINITION_INTEGERAPI

		// insertion point for updating fields
		datatype_definition_integerAPI.ID = datatype_definition_integerDB.ID
		datatype_definition_integerDB.CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER_WOP(&datatype_definition_integerAPI.DATATYPE_DEFINITION_INTEGER_WOP)
		datatype_definition_integerAPI.DATATYPE_DEFINITION_INTEGERPointersEncoding = datatype_definition_integerDB.DATATYPE_DEFINITION_INTEGERPointersEncoding
		datatype_definition_integerAPIs = append(datatype_definition_integerAPIs, datatype_definition_integerAPI)
	}

	c.JSON(http.StatusOK, datatype_definition_integerAPIs)
}

// PostDATATYPE_DEFINITION_INTEGER
//
// swagger:route POST /datatype_definition_integers datatype_definition_integers postDATATYPE_DEFINITION_INTEGER
//
// Creates a datatype_definition_integer
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPE_DEFINITION_INTEGER(c *gin.Context) {

	mutexDATATYPE_DEFINITION_INTEGER.Lock()
	defer mutexDATATYPE_DEFINITION_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPE_DEFINITION_INTEGERs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_INTEGERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatype_definition_integer
	datatype_definition_integerDB := orm.DATATYPE_DEFINITION_INTEGERDB{}
	datatype_definition_integerDB.DATATYPE_DEFINITION_INTEGERPointersEncoding = input.DATATYPE_DEFINITION_INTEGERPointersEncoding
	datatype_definition_integerDB.CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGER_WOP(&input.DATATYPE_DEFINITION_INTEGER_WOP)

	_, err = db.Create(&datatype_definition_integerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseOneInstance(&datatype_definition_integerDB)
	datatype_definition_integer := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr[datatype_definition_integerDB.ID]

	if datatype_definition_integer != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatype_definition_integer)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatype_definition_integerDB)
}

// GetDATATYPE_DEFINITION_INTEGER
//
// swagger:route GET /datatype_definition_integers/{ID} datatype_definition_integers getDATATYPE_DEFINITION_INTEGER
//
// Gets the details for a datatype_definition_integer.
//
// Responses:
// default: genericError
//
//	200: datatype_definition_integerDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_INTEGER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_INTEGER", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.GetDB()

	// Get datatype_definition_integerDB in DB
	var datatype_definition_integerDB orm.DATATYPE_DEFINITION_INTEGERDB
	if _, err := db.First(&datatype_definition_integerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatype_definition_integerAPI orm.DATATYPE_DEFINITION_INTEGERAPI
	datatype_definition_integerAPI.ID = datatype_definition_integerDB.ID
	datatype_definition_integerAPI.DATATYPE_DEFINITION_INTEGERPointersEncoding = datatype_definition_integerDB.DATATYPE_DEFINITION_INTEGERPointersEncoding
	datatype_definition_integerDB.CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER_WOP(&datatype_definition_integerAPI.DATATYPE_DEFINITION_INTEGER_WOP)

	c.JSON(http.StatusOK, datatype_definition_integerAPI)
}

// UpdateDATATYPE_DEFINITION_INTEGER
//
// swagger:route PATCH /datatype_definition_integers/{ID} datatype_definition_integers updateDATATYPE_DEFINITION_INTEGER
//
// # Update a datatype_definition_integer
//
// Responses:
// default: genericError
//
//	200: datatype_definition_integerDBResponse
func (controller *Controller) UpdateDATATYPE_DEFINITION_INTEGER(c *gin.Context) {

	mutexDATATYPE_DEFINITION_INTEGER.Lock()
	defer mutexDATATYPE_DEFINITION_INTEGER.Unlock()

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
	db := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_INTEGERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatype_definition_integerDB orm.DATATYPE_DEFINITION_INTEGERDB

	// fetch the datatype_definition_integer
	_, err := db.First(&datatype_definition_integerDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatype_definition_integerDB.CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGER_WOP(&input.DATATYPE_DEFINITION_INTEGER_WOP)
	datatype_definition_integerDB.DATATYPE_DEFINITION_INTEGERPointersEncoding = input.DATATYPE_DEFINITION_INTEGERPointersEncoding

	db, _ = db.Model(&datatype_definition_integerDB)
	_, err = db.Updates(&datatype_definition_integerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_integerNew := new(models.DATATYPE_DEFINITION_INTEGER)
	datatype_definition_integerDB.CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER(datatype_definition_integerNew)

	// redeem pointers
	datatype_definition_integerDB.DecodePointers(backRepo, datatype_definition_integerNew)

	// get stage instance from DB instance, and call callback function
	datatype_definition_integerOld := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr[datatype_definition_integerDB.ID]
	if datatype_definition_integerOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), datatype_definition_integerOld, datatype_definition_integerNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), datatype_definition_integerOld, datatype_definition_integerNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatype_definition_integerDB
	c.JSON(http.StatusOK, datatype_definition_integerDB)
}

// DeleteDATATYPE_DEFINITION_INTEGER
//
// swagger:route DELETE /datatype_definition_integers/{ID} datatype_definition_integers deleteDATATYPE_DEFINITION_INTEGER
//
// # Delete a datatype_definition_integer
//
// default: genericError
//
//	200: datatype_definition_integerDBResponse
func (controller *Controller) DeleteDATATYPE_DEFINITION_INTEGER(c *gin.Context) {

	mutexDATATYPE_DEFINITION_INTEGER.Lock()
	defer mutexDATATYPE_DEFINITION_INTEGER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPE_DEFINITION_INTEGER", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.GetDB()

	// Get model if exist
	var datatype_definition_integerDB orm.DATATYPE_DEFINITION_INTEGERDB
	if _, err := db.First(&datatype_definition_integerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&datatype_definition_integerDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_integerDeleted := new(models.DATATYPE_DEFINITION_INTEGER)
	datatype_definition_integerDB.CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER(datatype_definition_integerDeleted)

	// get stage instance from DB instance, and call callback function
	datatype_definition_integerStaged := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr[datatype_definition_integerDB.ID]
	if datatype_definition_integerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatype_definition_integerStaged, datatype_definition_integerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
