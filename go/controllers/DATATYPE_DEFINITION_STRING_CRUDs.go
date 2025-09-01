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
var __DATATYPE_DEFINITION_STRING__dummysDeclaration__ models.DATATYPE_DEFINITION_STRING
var __DATATYPE_DEFINITION_STRING_time__dummyDeclaration time.Duration

var mutexDATATYPE_DEFINITION_STRING sync.Mutex

// An DATATYPE_DEFINITION_STRINGID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPE_DEFINITION_STRING updateDATATYPE_DEFINITION_STRING deleteDATATYPE_DEFINITION_STRING
type DATATYPE_DEFINITION_STRINGID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPE_DEFINITION_STRINGInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPE_DEFINITION_STRING updateDATATYPE_DEFINITION_STRING
type DATATYPE_DEFINITION_STRINGInput struct {
	// The DATATYPE_DEFINITION_STRING to submit or modify
	// in: body
	DATATYPE_DEFINITION_STRING *orm.DATATYPE_DEFINITION_STRINGAPI
}

// GetDATATYPE_DEFINITION_STRINGs
//
// swagger:route GET /datatype_definition_strings datatype_definition_strings getDATATYPE_DEFINITION_STRINGs
//
// # Get all datatype_definition_strings
//
// Responses:
// default: genericError
//
//	200: datatype_definition_stringDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_STRINGs(c *gin.Context) {

	// source slice
	var datatype_definition_stringDBs []orm.DATATYPE_DEFINITION_STRINGDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_STRINGs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_STRING.GetDB()

	_, err := db.Find(&datatype_definition_stringDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatype_definition_stringAPIs := make([]orm.DATATYPE_DEFINITION_STRINGAPI, 0)

	// for each datatype_definition_string, update fields from the database nullable fields
	for idx := range datatype_definition_stringDBs {
		datatype_definition_stringDB := &datatype_definition_stringDBs[idx]
		_ = datatype_definition_stringDB
		var datatype_definition_stringAPI orm.DATATYPE_DEFINITION_STRINGAPI

		// insertion point for updating fields
		datatype_definition_stringAPI.ID = datatype_definition_stringDB.ID
		datatype_definition_stringDB.CopyBasicFieldsToDATATYPE_DEFINITION_STRING_WOP(&datatype_definition_stringAPI.DATATYPE_DEFINITION_STRING_WOP)
		datatype_definition_stringAPI.DATATYPE_DEFINITION_STRINGPointersEncoding = datatype_definition_stringDB.DATATYPE_DEFINITION_STRINGPointersEncoding
		datatype_definition_stringAPIs = append(datatype_definition_stringAPIs, datatype_definition_stringAPI)
	}

	c.JSON(http.StatusOK, datatype_definition_stringAPIs)
}

// PostDATATYPE_DEFINITION_STRING
//
// swagger:route POST /datatype_definition_strings datatype_definition_strings postDATATYPE_DEFINITION_STRING
//
// Creates a datatype_definition_string
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPE_DEFINITION_STRING(c *gin.Context) {

	mutexDATATYPE_DEFINITION_STRING.Lock()
	defer mutexDATATYPE_DEFINITION_STRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPE_DEFINITION_STRINGs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_STRING.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_STRINGAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatype_definition_string
	datatype_definition_stringDB := orm.DATATYPE_DEFINITION_STRINGDB{}
	datatype_definition_stringDB.DATATYPE_DEFINITION_STRINGPointersEncoding = input.DATATYPE_DEFINITION_STRINGPointersEncoding
	datatype_definition_stringDB.CopyBasicFieldsFromDATATYPE_DEFINITION_STRING_WOP(&input.DATATYPE_DEFINITION_STRING_WOP)

	_, err = db.Create(&datatype_definition_stringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseOneInstance(&datatype_definition_stringDB)
	datatype_definition_string := backRepo.BackRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr[datatype_definition_stringDB.ID]

	if datatype_definition_string != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatype_definition_string)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatype_definition_stringDB)
}

// GetDATATYPE_DEFINITION_STRING
//
// swagger:route GET /datatype_definition_strings/{ID} datatype_definition_strings getDATATYPE_DEFINITION_STRING
//
// Gets the details for a datatype_definition_string.
//
// Responses:
// default: genericError
//
//	200: datatype_definition_stringDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_STRING(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_STRING", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_STRING.GetDB()

	// Get datatype_definition_stringDB in DB
	var datatype_definition_stringDB orm.DATATYPE_DEFINITION_STRINGDB
	if _, err := db.First(&datatype_definition_stringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatype_definition_stringAPI orm.DATATYPE_DEFINITION_STRINGAPI
	datatype_definition_stringAPI.ID = datatype_definition_stringDB.ID
	datatype_definition_stringAPI.DATATYPE_DEFINITION_STRINGPointersEncoding = datatype_definition_stringDB.DATATYPE_DEFINITION_STRINGPointersEncoding
	datatype_definition_stringDB.CopyBasicFieldsToDATATYPE_DEFINITION_STRING_WOP(&datatype_definition_stringAPI.DATATYPE_DEFINITION_STRING_WOP)

	c.JSON(http.StatusOK, datatype_definition_stringAPI)
}

// UpdateDATATYPE_DEFINITION_STRING
//
// swagger:route PATCH /datatype_definition_strings/{ID} datatype_definition_strings updateDATATYPE_DEFINITION_STRING
//
// # Update a datatype_definition_string
//
// Responses:
// default: genericError
//
//	200: datatype_definition_stringDBResponse
func (controller *Controller) UpdateDATATYPE_DEFINITION_STRING(c *gin.Context) {

	mutexDATATYPE_DEFINITION_STRING.Lock()
	defer mutexDATATYPE_DEFINITION_STRING.Unlock()

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
	db := backRepo.BackRepoDATATYPE_DEFINITION_STRING.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_STRINGAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatype_definition_stringDB orm.DATATYPE_DEFINITION_STRINGDB

	// fetch the datatype_definition_string
	_, err := db.First(&datatype_definition_stringDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatype_definition_stringDB.CopyBasicFieldsFromDATATYPE_DEFINITION_STRING_WOP(&input.DATATYPE_DEFINITION_STRING_WOP)
	datatype_definition_stringDB.DATATYPE_DEFINITION_STRINGPointersEncoding = input.DATATYPE_DEFINITION_STRINGPointersEncoding

	db, _ = db.Model(&datatype_definition_stringDB)
	_, err = db.Updates(&datatype_definition_stringDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_stringNew := new(models.DATATYPE_DEFINITION_STRING)
	datatype_definition_stringDB.CopyBasicFieldsToDATATYPE_DEFINITION_STRING(datatype_definition_stringNew)

	// redeem pointers
	datatype_definition_stringDB.DecodePointers(backRepo, datatype_definition_stringNew)

	// get stage instance from DB instance, and call callback function
	datatype_definition_stringOld := backRepo.BackRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr[datatype_definition_stringDB.ID]
	if datatype_definition_stringOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), datatype_definition_stringOld, datatype_definition_stringNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), datatype_definition_stringOld, datatype_definition_stringNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatype_definition_stringDB
	c.JSON(http.StatusOK, datatype_definition_stringDB)
}

// DeleteDATATYPE_DEFINITION_STRING
//
// swagger:route DELETE /datatype_definition_strings/{ID} datatype_definition_strings deleteDATATYPE_DEFINITION_STRING
//
// # Delete a datatype_definition_string
//
// default: genericError
//
//	200: datatype_definition_stringDBResponse
func (controller *Controller) DeleteDATATYPE_DEFINITION_STRING(c *gin.Context) {

	mutexDATATYPE_DEFINITION_STRING.Lock()
	defer mutexDATATYPE_DEFINITION_STRING.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPE_DEFINITION_STRING", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_STRING.GetDB()

	// Get model if exist
	var datatype_definition_stringDB orm.DATATYPE_DEFINITION_STRINGDB
	if _, err := db.First(&datatype_definition_stringDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&datatype_definition_stringDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_stringDeleted := new(models.DATATYPE_DEFINITION_STRING)
	datatype_definition_stringDB.CopyBasicFieldsToDATATYPE_DEFINITION_STRING(datatype_definition_stringDeleted)

	// get stage instance from DB instance, and call callback function
	datatype_definition_stringStaged := backRepo.BackRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr[datatype_definition_stringDB.ID]
	if datatype_definition_stringStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatype_definition_stringStaged, datatype_definition_stringDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
