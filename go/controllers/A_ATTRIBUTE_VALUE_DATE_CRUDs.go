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
var __A_ATTRIBUTE_VALUE_DATE__dummysDeclaration__ models.A_ATTRIBUTE_VALUE_DATE
var __A_ATTRIBUTE_VALUE_DATE_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_VALUE_DATE sync.Mutex

// An A_ATTRIBUTE_VALUE_DATEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_VALUE_DATE updateA_ATTRIBUTE_VALUE_DATE deleteA_ATTRIBUTE_VALUE_DATE
type A_ATTRIBUTE_VALUE_DATEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_VALUE_DATEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_VALUE_DATE updateA_ATTRIBUTE_VALUE_DATE
type A_ATTRIBUTE_VALUE_DATEInput struct {
	// The A_ATTRIBUTE_VALUE_DATE to submit or modify
	// in: body
	A_ATTRIBUTE_VALUE_DATE *orm.A_ATTRIBUTE_VALUE_DATEAPI
}

// GetA_ATTRIBUTE_VALUE_DATEs
//
// swagger:route GET /a_attribute_value_dates a_attribute_value_dates getA_ATTRIBUTE_VALUE_DATEs
//
// # Get all a_attribute_value_dates
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_dateDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_DATEs(c *gin.Context) {

	// source slice
	var a_attribute_value_dateDBs []orm.A_ATTRIBUTE_VALUE_DATEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_DATEs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.GetDB()

	_, err := db.Find(&a_attribute_value_dateDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_value_dateAPIs := make([]orm.A_ATTRIBUTE_VALUE_DATEAPI, 0)

	// for each a_attribute_value_date, update fields from the database nullable fields
	for idx := range a_attribute_value_dateDBs {
		a_attribute_value_dateDB := &a_attribute_value_dateDBs[idx]
		_ = a_attribute_value_dateDB
		var a_attribute_value_dateAPI orm.A_ATTRIBUTE_VALUE_DATEAPI

		// insertion point for updating fields
		a_attribute_value_dateAPI.ID = a_attribute_value_dateDB.ID
		a_attribute_value_dateDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_DATE_WOP(&a_attribute_value_dateAPI.A_ATTRIBUTE_VALUE_DATE_WOP)
		a_attribute_value_dateAPI.A_ATTRIBUTE_VALUE_DATEPointersEncoding = a_attribute_value_dateDB.A_ATTRIBUTE_VALUE_DATEPointersEncoding
		a_attribute_value_dateAPIs = append(a_attribute_value_dateAPIs, a_attribute_value_dateAPI)
	}

	c.JSON(http.StatusOK, a_attribute_value_dateAPIs)
}

// PostA_ATTRIBUTE_VALUE_DATE
//
// swagger:route POST /a_attribute_value_dates a_attribute_value_dates postA_ATTRIBUTE_VALUE_DATE
//
// Creates a a_attribute_value_date
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_VALUE_DATE(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_DATE.Lock()
	defer mutexA_ATTRIBUTE_VALUE_DATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_VALUE_DATEs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_DATEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_value_date
	a_attribute_value_dateDB := orm.A_ATTRIBUTE_VALUE_DATEDB{}
	a_attribute_value_dateDB.A_ATTRIBUTE_VALUE_DATEPointersEncoding = input.A_ATTRIBUTE_VALUE_DATEPointersEncoding
	a_attribute_value_dateDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_DATE_WOP(&input.A_ATTRIBUTE_VALUE_DATE_WOP)

	_, err = db.Create(&a_attribute_value_dateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.CheckoutPhaseOneInstance(&a_attribute_value_dateDB)
	a_attribute_value_date := backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.Map_A_ATTRIBUTE_VALUE_DATEDBID_A_ATTRIBUTE_VALUE_DATEPtr[a_attribute_value_dateDB.ID]

	if a_attribute_value_date != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_value_date)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_value_dateDB)
}

// GetA_ATTRIBUTE_VALUE_DATE
//
// swagger:route GET /a_attribute_value_dates/{ID} a_attribute_value_dates getA_ATTRIBUTE_VALUE_DATE
//
// Gets the details for a a_attribute_value_date.
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_dateDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_DATE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_DATE", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.GetDB()

	// Get a_attribute_value_dateDB in DB
	var a_attribute_value_dateDB orm.A_ATTRIBUTE_VALUE_DATEDB
	if _, err := db.First(&a_attribute_value_dateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_value_dateAPI orm.A_ATTRIBUTE_VALUE_DATEAPI
	a_attribute_value_dateAPI.ID = a_attribute_value_dateDB.ID
	a_attribute_value_dateAPI.A_ATTRIBUTE_VALUE_DATEPointersEncoding = a_attribute_value_dateDB.A_ATTRIBUTE_VALUE_DATEPointersEncoding
	a_attribute_value_dateDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_DATE_WOP(&a_attribute_value_dateAPI.A_ATTRIBUTE_VALUE_DATE_WOP)

	c.JSON(http.StatusOK, a_attribute_value_dateAPI)
}

// UpdateA_ATTRIBUTE_VALUE_DATE
//
// swagger:route PATCH /a_attribute_value_dates/{ID} a_attribute_value_dates updateA_ATTRIBUTE_VALUE_DATE
//
// # Update a a_attribute_value_date
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_dateDBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_VALUE_DATE(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_DATE.Lock()
	defer mutexA_ATTRIBUTE_VALUE_DATE.Unlock()

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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_DATEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_value_dateDB orm.A_ATTRIBUTE_VALUE_DATEDB

	// fetch the a_attribute_value_date
	_, err := db.First(&a_attribute_value_dateDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_value_dateDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_DATE_WOP(&input.A_ATTRIBUTE_VALUE_DATE_WOP)
	a_attribute_value_dateDB.A_ATTRIBUTE_VALUE_DATEPointersEncoding = input.A_ATTRIBUTE_VALUE_DATEPointersEncoding

	db, _ = db.Model(&a_attribute_value_dateDB)
	_, err = db.Updates(&a_attribute_value_dateDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_dateNew := new(models.A_ATTRIBUTE_VALUE_DATE)
	a_attribute_value_dateDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_DATE(a_attribute_value_dateNew)

	// redeem pointers
	a_attribute_value_dateDB.DecodePointers(backRepo, a_attribute_value_dateNew)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_dateOld := backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.Map_A_ATTRIBUTE_VALUE_DATEDBID_A_ATTRIBUTE_VALUE_DATEPtr[a_attribute_value_dateDB.ID]
	if a_attribute_value_dateOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_dateOld, a_attribute_value_dateNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_dateOld, a_attribute_value_dateNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_value_dateDB
	c.JSON(http.StatusOK, a_attribute_value_dateDB)
}

// DeleteA_ATTRIBUTE_VALUE_DATE
//
// swagger:route DELETE /a_attribute_value_dates/{ID} a_attribute_value_dates deleteA_ATTRIBUTE_VALUE_DATE
//
// # Delete a a_attribute_value_date
//
// default: genericError
//
//	200: a_attribute_value_dateDBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_VALUE_DATE(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_DATE.Lock()
	defer mutexA_ATTRIBUTE_VALUE_DATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_VALUE_DATE", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.GetDB()

	// Get model if exist
	var a_attribute_value_dateDB orm.A_ATTRIBUTE_VALUE_DATEDB
	if _, err := db.First(&a_attribute_value_dateDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_value_dateDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_dateDeleted := new(models.A_ATTRIBUTE_VALUE_DATE)
	a_attribute_value_dateDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_DATE(a_attribute_value_dateDeleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_dateStaged := backRepo.BackRepoA_ATTRIBUTE_VALUE_DATE.Map_A_ATTRIBUTE_VALUE_DATEDBID_A_ATTRIBUTE_VALUE_DATEPtr[a_attribute_value_dateDB.ID]
	if a_attribute_value_dateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_value_dateStaged, a_attribute_value_dateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
