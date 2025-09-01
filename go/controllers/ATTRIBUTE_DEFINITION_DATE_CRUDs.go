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
var __ATTRIBUTE_DEFINITION_DATE__dummysDeclaration__ models.ATTRIBUTE_DEFINITION_DATE
var __ATTRIBUTE_DEFINITION_DATE_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_DEFINITION_DATE sync.Mutex

// An ATTRIBUTE_DEFINITION_DATEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_DEFINITION_DATE updateATTRIBUTE_DEFINITION_DATE deleteATTRIBUTE_DEFINITION_DATE
type ATTRIBUTE_DEFINITION_DATEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_DEFINITION_DATEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_DEFINITION_DATE updateATTRIBUTE_DEFINITION_DATE
type ATTRIBUTE_DEFINITION_DATEInput struct {
	// The ATTRIBUTE_DEFINITION_DATE to submit or modify
	// in: body
	ATTRIBUTE_DEFINITION_DATE *orm.ATTRIBUTE_DEFINITION_DATEAPI
}

// GetATTRIBUTE_DEFINITION_DATEs
//
// swagger:route GET /attribute_definition_dates attribute_definition_dates getATTRIBUTE_DEFINITION_DATEs
//
// # Get all attribute_definition_dates
//
// Responses:
// default: genericError
//
//	200: attribute_definition_dateDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_DATEs(c *gin.Context) {

	// source slice
	var attribute_definition_dateDBs []orm.ATTRIBUTE_DEFINITION_DATEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_DATEs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.GetDB()

	_, err := db.Find(&attribute_definition_dateDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_definition_dateAPIs := make([]orm.ATTRIBUTE_DEFINITION_DATEAPI, 0)

	// for each attribute_definition_date, update fields from the database nullable fields
	for idx := range attribute_definition_dateDBs {
		attribute_definition_dateDB := &attribute_definition_dateDBs[idx]
		_ = attribute_definition_dateDB
		var attribute_definition_dateAPI orm.ATTRIBUTE_DEFINITION_DATEAPI

		// insertion point for updating fields
		attribute_definition_dateAPI.ID = attribute_definition_dateDB.ID
		attribute_definition_dateDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_DATE_WOP(&attribute_definition_dateAPI.ATTRIBUTE_DEFINITION_DATE_WOP)
		attribute_definition_dateAPI.ATTRIBUTE_DEFINITION_DATEPointersEncoding = attribute_definition_dateDB.ATTRIBUTE_DEFINITION_DATEPointersEncoding
		attribute_definition_dateAPIs = append(attribute_definition_dateAPIs, attribute_definition_dateAPI)
	}

	c.JSON(http.StatusOK, attribute_definition_dateAPIs)
}

// PostATTRIBUTE_DEFINITION_DATE
//
// swagger:route POST /attribute_definition_dates attribute_definition_dates postATTRIBUTE_DEFINITION_DATE
//
// Creates a attribute_definition_date
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_DEFINITION_DATE(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_DATE.Lock()
	defer mutexATTRIBUTE_DEFINITION_DATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_DEFINITION_DATEs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_DATEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_definition_date
	attribute_definition_dateDB := orm.ATTRIBUTE_DEFINITION_DATEDB{}
	attribute_definition_dateDB.ATTRIBUTE_DEFINITION_DATEPointersEncoding = input.ATTRIBUTE_DEFINITION_DATEPointersEncoding
	attribute_definition_dateDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_DATE_WOP(&input.ATTRIBUTE_DEFINITION_DATE_WOP)

	_, err = db.Create(&attribute_definition_dateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.CheckoutPhaseOneInstance(&attribute_definition_dateDB)
	attribute_definition_date := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.Map_ATTRIBUTE_DEFINITION_DATEDBID_ATTRIBUTE_DEFINITION_DATEPtr[attribute_definition_dateDB.ID]

	if attribute_definition_date != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_definition_date)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_definition_dateDB)
}

// GetATTRIBUTE_DEFINITION_DATE
//
// swagger:route GET /attribute_definition_dates/{ID} attribute_definition_dates getATTRIBUTE_DEFINITION_DATE
//
// Gets the details for a attribute_definition_date.
//
// Responses:
// default: genericError
//
//	200: attribute_definition_dateDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_DATE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_DATE", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.GetDB()

	// Get attribute_definition_dateDB in DB
	var attribute_definition_dateDB orm.ATTRIBUTE_DEFINITION_DATEDB
	if _, err := db.First(&attribute_definition_dateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_definition_dateAPI orm.ATTRIBUTE_DEFINITION_DATEAPI
	attribute_definition_dateAPI.ID = attribute_definition_dateDB.ID
	attribute_definition_dateAPI.ATTRIBUTE_DEFINITION_DATEPointersEncoding = attribute_definition_dateDB.ATTRIBUTE_DEFINITION_DATEPointersEncoding
	attribute_definition_dateDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_DATE_WOP(&attribute_definition_dateAPI.ATTRIBUTE_DEFINITION_DATE_WOP)

	c.JSON(http.StatusOK, attribute_definition_dateAPI)
}

// UpdateATTRIBUTE_DEFINITION_DATE
//
// swagger:route PATCH /attribute_definition_dates/{ID} attribute_definition_dates updateATTRIBUTE_DEFINITION_DATE
//
// # Update a attribute_definition_date
//
// Responses:
// default: genericError
//
//	200: attribute_definition_dateDBResponse
func (controller *Controller) UpdateATTRIBUTE_DEFINITION_DATE(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_DATE.Lock()
	defer mutexATTRIBUTE_DEFINITION_DATE.Unlock()

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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_DATEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_definition_dateDB orm.ATTRIBUTE_DEFINITION_DATEDB

	// fetch the attribute_definition_date
	_, err := db.First(&attribute_definition_dateDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_definition_dateDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_DATE_WOP(&input.ATTRIBUTE_DEFINITION_DATE_WOP)
	attribute_definition_dateDB.ATTRIBUTE_DEFINITION_DATEPointersEncoding = input.ATTRIBUTE_DEFINITION_DATEPointersEncoding

	db, _ = db.Model(&attribute_definition_dateDB)
	_, err = db.Updates(&attribute_definition_dateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_dateNew := new(models.ATTRIBUTE_DEFINITION_DATE)
	attribute_definition_dateDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_DATE(attribute_definition_dateNew)

	// redeem pointers
	attribute_definition_dateDB.DecodePointers(backRepo, attribute_definition_dateNew)

	// get stage instance from DB instance, and call callback function
	attribute_definition_dateOld := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.Map_ATTRIBUTE_DEFINITION_DATEDBID_ATTRIBUTE_DEFINITION_DATEPtr[attribute_definition_dateDB.ID]
	if attribute_definition_dateOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_definition_dateOld, attribute_definition_dateNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_definition_dateOld, attribute_definition_dateNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_definition_dateDB
	c.JSON(http.StatusOK, attribute_definition_dateDB)
}

// DeleteATTRIBUTE_DEFINITION_DATE
//
// swagger:route DELETE /attribute_definition_dates/{ID} attribute_definition_dates deleteATTRIBUTE_DEFINITION_DATE
//
// # Delete a attribute_definition_date
//
// default: genericError
//
//	200: attribute_definition_dateDBResponse
func (controller *Controller) DeleteATTRIBUTE_DEFINITION_DATE(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_DATE.Lock()
	defer mutexATTRIBUTE_DEFINITION_DATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_DEFINITION_DATE", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.GetDB()

	// Get model if exist
	var attribute_definition_dateDB orm.ATTRIBUTE_DEFINITION_DATEDB
	if _, err := db.First(&attribute_definition_dateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_definition_dateDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_dateDeleted := new(models.ATTRIBUTE_DEFINITION_DATE)
	attribute_definition_dateDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_DATE(attribute_definition_dateDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_definition_dateStaged := backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.Map_ATTRIBUTE_DEFINITION_DATEDBID_ATTRIBUTE_DEFINITION_DATEPtr[attribute_definition_dateDB.ID]
	if attribute_definition_dateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_definition_dateStaged, attribute_definition_dateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
