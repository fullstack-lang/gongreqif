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
var __ATTRIBUTE_VALUE_DATE__dummysDeclaration__ models.ATTRIBUTE_VALUE_DATE
var __ATTRIBUTE_VALUE_DATE_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_VALUE_DATE sync.Mutex

// An ATTRIBUTE_VALUE_DATEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_VALUE_DATE updateATTRIBUTE_VALUE_DATE deleteATTRIBUTE_VALUE_DATE
type ATTRIBUTE_VALUE_DATEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_VALUE_DATEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_VALUE_DATE updateATTRIBUTE_VALUE_DATE
type ATTRIBUTE_VALUE_DATEInput struct {
	// The ATTRIBUTE_VALUE_DATE to submit or modify
	// in: body
	ATTRIBUTE_VALUE_DATE *orm.ATTRIBUTE_VALUE_DATEAPI
}

// GetATTRIBUTE_VALUE_DATEs
//
// swagger:route GET /attribute_value_dates attribute_value_dates getATTRIBUTE_VALUE_DATEs
//
// # Get all attribute_value_dates
//
// Responses:
// default: genericError
//
//	200: attribute_value_dateDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_DATEs(c *gin.Context) {

	// source slice
	var attribute_value_dateDBs []orm.ATTRIBUTE_VALUE_DATEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_DATEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_DATE.GetDB()

	query := db.Find(&attribute_value_dateDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_value_dateAPIs := make([]orm.ATTRIBUTE_VALUE_DATEAPI, 0)

	// for each attribute_value_date, update fields from the database nullable fields
	for idx := range attribute_value_dateDBs {
		attribute_value_dateDB := &attribute_value_dateDBs[idx]
		_ = attribute_value_dateDB
		var attribute_value_dateAPI orm.ATTRIBUTE_VALUE_DATEAPI

		// insertion point for updating fields
		attribute_value_dateAPI.ID = attribute_value_dateDB.ID
		attribute_value_dateDB.CopyBasicFieldsToATTRIBUTE_VALUE_DATE_WOP(&attribute_value_dateAPI.ATTRIBUTE_VALUE_DATE_WOP)
		attribute_value_dateAPI.ATTRIBUTE_VALUE_DATEPointersEncoding = attribute_value_dateDB.ATTRIBUTE_VALUE_DATEPointersEncoding
		attribute_value_dateAPIs = append(attribute_value_dateAPIs, attribute_value_dateAPI)
	}

	c.JSON(http.StatusOK, attribute_value_dateAPIs)
}

// PostATTRIBUTE_VALUE_DATE
//
// swagger:route POST /attribute_value_dates attribute_value_dates postATTRIBUTE_VALUE_DATE
//
// Creates a attribute_value_date
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_VALUE_DATE(c *gin.Context) {

	mutexATTRIBUTE_VALUE_DATE.Lock()
	defer mutexATTRIBUTE_VALUE_DATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_VALUE_DATEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_DATE.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_DATEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_value_date
	attribute_value_dateDB := orm.ATTRIBUTE_VALUE_DATEDB{}
	attribute_value_dateDB.ATTRIBUTE_VALUE_DATEPointersEncoding = input.ATTRIBUTE_VALUE_DATEPointersEncoding
	attribute_value_dateDB.CopyBasicFieldsFromATTRIBUTE_VALUE_DATE_WOP(&input.ATTRIBUTE_VALUE_DATE_WOP)

	query := db.Create(&attribute_value_dateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.CheckoutPhaseOneInstance(&attribute_value_dateDB)
	attribute_value_date := backRepo.BackRepoATTRIBUTE_VALUE_DATE.Map_ATTRIBUTE_VALUE_DATEDBID_ATTRIBUTE_VALUE_DATEPtr[attribute_value_dateDB.ID]

	if attribute_value_date != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_value_date)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_value_dateDB)
}

// GetATTRIBUTE_VALUE_DATE
//
// swagger:route GET /attribute_value_dates/{ID} attribute_value_dates getATTRIBUTE_VALUE_DATE
//
// Gets the details for a attribute_value_date.
//
// Responses:
// default: genericError
//
//	200: attribute_value_dateDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_DATE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_DATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_DATE.GetDB()

	// Get attribute_value_dateDB in DB
	var attribute_value_dateDB orm.ATTRIBUTE_VALUE_DATEDB
	if err := db.First(&attribute_value_dateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_value_dateAPI orm.ATTRIBUTE_VALUE_DATEAPI
	attribute_value_dateAPI.ID = attribute_value_dateDB.ID
	attribute_value_dateAPI.ATTRIBUTE_VALUE_DATEPointersEncoding = attribute_value_dateDB.ATTRIBUTE_VALUE_DATEPointersEncoding
	attribute_value_dateDB.CopyBasicFieldsToATTRIBUTE_VALUE_DATE_WOP(&attribute_value_dateAPI.ATTRIBUTE_VALUE_DATE_WOP)

	c.JSON(http.StatusOK, attribute_value_dateAPI)
}

// UpdateATTRIBUTE_VALUE_DATE
//
// swagger:route PATCH /attribute_value_dates/{ID} attribute_value_dates updateATTRIBUTE_VALUE_DATE
//
// # Update a attribute_value_date
//
// Responses:
// default: genericError
//
//	200: attribute_value_dateDBResponse
func (controller *Controller) UpdateATTRIBUTE_VALUE_DATE(c *gin.Context) {

	mutexATTRIBUTE_VALUE_DATE.Lock()
	defer mutexATTRIBUTE_VALUE_DATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTE_VALUE_DATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_DATE.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_DATEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_value_dateDB orm.ATTRIBUTE_VALUE_DATEDB

	// fetch the attribute_value_date
	query := db.First(&attribute_value_dateDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_value_dateDB.CopyBasicFieldsFromATTRIBUTE_VALUE_DATE_WOP(&input.ATTRIBUTE_VALUE_DATE_WOP)
	attribute_value_dateDB.ATTRIBUTE_VALUE_DATEPointersEncoding = input.ATTRIBUTE_VALUE_DATEPointersEncoding

	query = db.Model(&attribute_value_dateDB).Updates(attribute_value_dateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_dateNew := new(models.ATTRIBUTE_VALUE_DATE)
	attribute_value_dateDB.CopyBasicFieldsToATTRIBUTE_VALUE_DATE(attribute_value_dateNew)

	// redeem pointers
	attribute_value_dateDB.DecodePointers(backRepo, attribute_value_dateNew)

	// get stage instance from DB instance, and call callback function
	attribute_value_dateOld := backRepo.BackRepoATTRIBUTE_VALUE_DATE.Map_ATTRIBUTE_VALUE_DATEDBID_ATTRIBUTE_VALUE_DATEPtr[attribute_value_dateDB.ID]
	if attribute_value_dateOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attribute_value_dateOld, attribute_value_dateNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_value_dateDB
	c.JSON(http.StatusOK, attribute_value_dateDB)
}

// DeleteATTRIBUTE_VALUE_DATE
//
// swagger:route DELETE /attribute_value_dates/{ID} attribute_value_dates deleteATTRIBUTE_VALUE_DATE
//
// # Delete a attribute_value_date
//
// default: genericError
//
//	200: attribute_value_dateDBResponse
func (controller *Controller) DeleteATTRIBUTE_VALUE_DATE(c *gin.Context) {

	mutexATTRIBUTE_VALUE_DATE.Lock()
	defer mutexATTRIBUTE_VALUE_DATE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_VALUE_DATE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_DATE.GetDB()

	// Get model if exist
	var attribute_value_dateDB orm.ATTRIBUTE_VALUE_DATEDB
	if err := db.First(&attribute_value_dateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&attribute_value_dateDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_dateDeleted := new(models.ATTRIBUTE_VALUE_DATE)
	attribute_value_dateDB.CopyBasicFieldsToATTRIBUTE_VALUE_DATE(attribute_value_dateDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_value_dateStaged := backRepo.BackRepoATTRIBUTE_VALUE_DATE.Map_ATTRIBUTE_VALUE_DATEDBID_ATTRIBUTE_VALUE_DATEPtr[attribute_value_dateDB.ID]
	if attribute_value_dateStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_value_dateStaged, attribute_value_dateDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
