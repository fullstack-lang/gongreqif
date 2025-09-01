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
var __DATATYPE_DEFINITION_DATE__dummysDeclaration__ models.DATATYPE_DEFINITION_DATE
var __DATATYPE_DEFINITION_DATE_time__dummyDeclaration time.Duration

var mutexDATATYPE_DEFINITION_DATE sync.Mutex

// An DATATYPE_DEFINITION_DATEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPE_DEFINITION_DATE updateDATATYPE_DEFINITION_DATE deleteDATATYPE_DEFINITION_DATE
type DATATYPE_DEFINITION_DATEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPE_DEFINITION_DATEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPE_DEFINITION_DATE updateDATATYPE_DEFINITION_DATE
type DATATYPE_DEFINITION_DATEInput struct {
	// The DATATYPE_DEFINITION_DATE to submit or modify
	// in: body
	DATATYPE_DEFINITION_DATE *orm.DATATYPE_DEFINITION_DATEAPI
}

// GetDATATYPE_DEFINITION_DATEs
//
// swagger:route GET /datatype_definition_dates datatype_definition_dates getDATATYPE_DEFINITION_DATEs
//
// # Get all datatype_definition_dates
//
// Responses:
// default: genericError
//
//	200: datatype_definition_dateDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_DATEs(c *gin.Context) {

	// source slice
	var datatype_definition_dateDBs []orm.DATATYPE_DEFINITION_DATEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_DATEs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_DATE.GetDB()

	_, err := db.Find(&datatype_definition_dateDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatype_definition_dateAPIs := make([]orm.DATATYPE_DEFINITION_DATEAPI, 0)

	// for each datatype_definition_date, update fields from the database nullable fields
	for idx := range datatype_definition_dateDBs {
		datatype_definition_dateDB := &datatype_definition_dateDBs[idx]
		_ = datatype_definition_dateDB
		var datatype_definition_dateAPI orm.DATATYPE_DEFINITION_DATEAPI

		// insertion point for updating fields
		datatype_definition_dateAPI.ID = datatype_definition_dateDB.ID
		datatype_definition_dateDB.CopyBasicFieldsToDATATYPE_DEFINITION_DATE_WOP(&datatype_definition_dateAPI.DATATYPE_DEFINITION_DATE_WOP)
		datatype_definition_dateAPI.DATATYPE_DEFINITION_DATEPointersEncoding = datatype_definition_dateDB.DATATYPE_DEFINITION_DATEPointersEncoding
		datatype_definition_dateAPIs = append(datatype_definition_dateAPIs, datatype_definition_dateAPI)
	}

	c.JSON(http.StatusOK, datatype_definition_dateAPIs)
}

// PostDATATYPE_DEFINITION_DATE
//
// swagger:route POST /datatype_definition_dates datatype_definition_dates postDATATYPE_DEFINITION_DATE
//
// Creates a datatype_definition_date
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPE_DEFINITION_DATE(c *gin.Context) {

	mutexDATATYPE_DEFINITION_DATE.Lock()
	defer mutexDATATYPE_DEFINITION_DATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPE_DEFINITION_DATEs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_DATE.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_DATEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatype_definition_date
	datatype_definition_dateDB := orm.DATATYPE_DEFINITION_DATEDB{}
	datatype_definition_dateDB.DATATYPE_DEFINITION_DATEPointersEncoding = input.DATATYPE_DEFINITION_DATEPointersEncoding
	datatype_definition_dateDB.CopyBasicFieldsFromDATATYPE_DEFINITION_DATE_WOP(&input.DATATYPE_DEFINITION_DATE_WOP)

	_, err = db.Create(&datatype_definition_dateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.CheckoutPhaseOneInstance(&datatype_definition_dateDB)
	datatype_definition_date := backRepo.BackRepoDATATYPE_DEFINITION_DATE.Map_DATATYPE_DEFINITION_DATEDBID_DATATYPE_DEFINITION_DATEPtr[datatype_definition_dateDB.ID]

	if datatype_definition_date != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatype_definition_date)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatype_definition_dateDB)
}

// GetDATATYPE_DEFINITION_DATE
//
// swagger:route GET /datatype_definition_dates/{ID} datatype_definition_dates getDATATYPE_DEFINITION_DATE
//
// Gets the details for a datatype_definition_date.
//
// Responses:
// default: genericError
//
//	200: datatype_definition_dateDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_DATE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_DATE", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_DATE.GetDB()

	// Get datatype_definition_dateDB in DB
	var datatype_definition_dateDB orm.DATATYPE_DEFINITION_DATEDB
	if _, err := db.First(&datatype_definition_dateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatype_definition_dateAPI orm.DATATYPE_DEFINITION_DATEAPI
	datatype_definition_dateAPI.ID = datatype_definition_dateDB.ID
	datatype_definition_dateAPI.DATATYPE_DEFINITION_DATEPointersEncoding = datatype_definition_dateDB.DATATYPE_DEFINITION_DATEPointersEncoding
	datatype_definition_dateDB.CopyBasicFieldsToDATATYPE_DEFINITION_DATE_WOP(&datatype_definition_dateAPI.DATATYPE_DEFINITION_DATE_WOP)

	c.JSON(http.StatusOK, datatype_definition_dateAPI)
}

// UpdateDATATYPE_DEFINITION_DATE
//
// swagger:route PATCH /datatype_definition_dates/{ID} datatype_definition_dates updateDATATYPE_DEFINITION_DATE
//
// # Update a datatype_definition_date
//
// Responses:
// default: genericError
//
//	200: datatype_definition_dateDBResponse
func (controller *Controller) UpdateDATATYPE_DEFINITION_DATE(c *gin.Context) {

	mutexDATATYPE_DEFINITION_DATE.Lock()
	defer mutexDATATYPE_DEFINITION_DATE.Unlock()

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
	db := backRepo.BackRepoDATATYPE_DEFINITION_DATE.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_DATEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatype_definition_dateDB orm.DATATYPE_DEFINITION_DATEDB

	// fetch the datatype_definition_date
	_, err := db.First(&datatype_definition_dateDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatype_definition_dateDB.CopyBasicFieldsFromDATATYPE_DEFINITION_DATE_WOP(&input.DATATYPE_DEFINITION_DATE_WOP)
	datatype_definition_dateDB.DATATYPE_DEFINITION_DATEPointersEncoding = input.DATATYPE_DEFINITION_DATEPointersEncoding

	db, _ = db.Model(&datatype_definition_dateDB)
	_, err = db.Updates(&datatype_definition_dateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_dateNew := new(models.DATATYPE_DEFINITION_DATE)
	datatype_definition_dateDB.CopyBasicFieldsToDATATYPE_DEFINITION_DATE(datatype_definition_dateNew)

	// redeem pointers
	datatype_definition_dateDB.DecodePointers(backRepo, datatype_definition_dateNew)

	// get stage instance from DB instance, and call callback function
	datatype_definition_dateOld := backRepo.BackRepoDATATYPE_DEFINITION_DATE.Map_DATATYPE_DEFINITION_DATEDBID_DATATYPE_DEFINITION_DATEPtr[datatype_definition_dateDB.ID]
	if datatype_definition_dateOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), datatype_definition_dateOld, datatype_definition_dateNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), datatype_definition_dateOld, datatype_definition_dateNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatype_definition_dateDB
	c.JSON(http.StatusOK, datatype_definition_dateDB)
}

// DeleteDATATYPE_DEFINITION_DATE
//
// swagger:route DELETE /datatype_definition_dates/{ID} datatype_definition_dates deleteDATATYPE_DEFINITION_DATE
//
// # Delete a datatype_definition_date
//
// default: genericError
//
//	200: datatype_definition_dateDBResponse
func (controller *Controller) DeleteDATATYPE_DEFINITION_DATE(c *gin.Context) {

	mutexDATATYPE_DEFINITION_DATE.Lock()
	defer mutexDATATYPE_DEFINITION_DATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPE_DEFINITION_DATE", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_DATE.GetDB()

	// Get model if exist
	var datatype_definition_dateDB orm.DATATYPE_DEFINITION_DATEDB
	if _, err := db.First(&datatype_definition_dateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&datatype_definition_dateDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_dateDeleted := new(models.DATATYPE_DEFINITION_DATE)
	datatype_definition_dateDB.CopyBasicFieldsToDATATYPE_DEFINITION_DATE(datatype_definition_dateDeleted)

	// get stage instance from DB instance, and call callback function
	datatype_definition_dateStaged := backRepo.BackRepoDATATYPE_DEFINITION_DATE.Map_DATATYPE_DEFINITION_DATEDBID_DATATYPE_DEFINITION_DATEPtr[datatype_definition_dateDB.ID]
	if datatype_definition_dateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatype_definition_dateStaged, datatype_definition_dateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
