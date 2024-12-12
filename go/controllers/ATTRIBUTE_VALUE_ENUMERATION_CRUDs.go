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
var __ATTRIBUTE_VALUE_ENUMERATION__dummysDeclaration__ models.ATTRIBUTE_VALUE_ENUMERATION
var __ATTRIBUTE_VALUE_ENUMERATION_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_VALUE_ENUMERATION sync.Mutex

// An ATTRIBUTE_VALUE_ENUMERATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_VALUE_ENUMERATION updateATTRIBUTE_VALUE_ENUMERATION deleteATTRIBUTE_VALUE_ENUMERATION
type ATTRIBUTE_VALUE_ENUMERATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_VALUE_ENUMERATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_VALUE_ENUMERATION updateATTRIBUTE_VALUE_ENUMERATION
type ATTRIBUTE_VALUE_ENUMERATIONInput struct {
	// The ATTRIBUTE_VALUE_ENUMERATION to submit or modify
	// in: body
	ATTRIBUTE_VALUE_ENUMERATION *orm.ATTRIBUTE_VALUE_ENUMERATIONAPI
}

// GetATTRIBUTE_VALUE_ENUMERATIONs
//
// swagger:route GET /attribute_value_enumerations attribute_value_enumerations getATTRIBUTE_VALUE_ENUMERATIONs
//
// # Get all attribute_value_enumerations
//
// Responses:
// default: genericError
//
//	200: attribute_value_enumerationDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_ENUMERATIONs(c *gin.Context) {

	// source slice
	var attribute_value_enumerationDBs []orm.ATTRIBUTE_VALUE_ENUMERATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_ENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.GetDB()

	_, err := db.Find(&attribute_value_enumerationDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_value_enumerationAPIs := make([]orm.ATTRIBUTE_VALUE_ENUMERATIONAPI, 0)

	// for each attribute_value_enumeration, update fields from the database nullable fields
	for idx := range attribute_value_enumerationDBs {
		attribute_value_enumerationDB := &attribute_value_enumerationDBs[idx]
		_ = attribute_value_enumerationDB
		var attribute_value_enumerationAPI orm.ATTRIBUTE_VALUE_ENUMERATIONAPI

		// insertion point for updating fields
		attribute_value_enumerationAPI.ID = attribute_value_enumerationDB.ID
		attribute_value_enumerationDB.CopyBasicFieldsToATTRIBUTE_VALUE_ENUMERATION_WOP(&attribute_value_enumerationAPI.ATTRIBUTE_VALUE_ENUMERATION_WOP)
		attribute_value_enumerationAPI.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = attribute_value_enumerationDB.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
		attribute_value_enumerationAPIs = append(attribute_value_enumerationAPIs, attribute_value_enumerationAPI)
	}

	c.JSON(http.StatusOK, attribute_value_enumerationAPIs)
}

// PostATTRIBUTE_VALUE_ENUMERATION
//
// swagger:route POST /attribute_value_enumerations attribute_value_enumerations postATTRIBUTE_VALUE_ENUMERATION
//
// Creates a attribute_value_enumeration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_VALUE_ENUMERATION(c *gin.Context) {

	mutexATTRIBUTE_VALUE_ENUMERATION.Lock()
	defer mutexATTRIBUTE_VALUE_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_VALUE_ENUMERATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_ENUMERATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_value_enumeration
	attribute_value_enumerationDB := orm.ATTRIBUTE_VALUE_ENUMERATIONDB{}
	attribute_value_enumerationDB.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = input.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
	attribute_value_enumerationDB.CopyBasicFieldsFromATTRIBUTE_VALUE_ENUMERATION_WOP(&input.ATTRIBUTE_VALUE_ENUMERATION_WOP)

	_, err = db.Create(&attribute_value_enumerationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseOneInstance(&attribute_value_enumerationDB)
	attribute_value_enumeration := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONPtr[attribute_value_enumerationDB.ID]

	if attribute_value_enumeration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_value_enumeration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_value_enumerationDB)
}

// GetATTRIBUTE_VALUE_ENUMERATION
//
// swagger:route GET /attribute_value_enumerations/{ID} attribute_value_enumerations getATTRIBUTE_VALUE_ENUMERATION
//
// Gets the details for a attribute_value_enumeration.
//
// Responses:
// default: genericError
//
//	200: attribute_value_enumerationDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_ENUMERATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_ENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.GetDB()

	// Get attribute_value_enumerationDB in DB
	var attribute_value_enumerationDB orm.ATTRIBUTE_VALUE_ENUMERATIONDB
	if _, err := db.First(&attribute_value_enumerationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_value_enumerationAPI orm.ATTRIBUTE_VALUE_ENUMERATIONAPI
	attribute_value_enumerationAPI.ID = attribute_value_enumerationDB.ID
	attribute_value_enumerationAPI.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = attribute_value_enumerationDB.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
	attribute_value_enumerationDB.CopyBasicFieldsToATTRIBUTE_VALUE_ENUMERATION_WOP(&attribute_value_enumerationAPI.ATTRIBUTE_VALUE_ENUMERATION_WOP)

	c.JSON(http.StatusOK, attribute_value_enumerationAPI)
}

// UpdateATTRIBUTE_VALUE_ENUMERATION
//
// swagger:route PATCH /attribute_value_enumerations/{ID} attribute_value_enumerations updateATTRIBUTE_VALUE_ENUMERATION
//
// # Update a attribute_value_enumeration
//
// Responses:
// default: genericError
//
//	200: attribute_value_enumerationDBResponse
func (controller *Controller) UpdateATTRIBUTE_VALUE_ENUMERATION(c *gin.Context) {

	mutexATTRIBUTE_VALUE_ENUMERATION.Lock()
	defer mutexATTRIBUTE_VALUE_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTE_VALUE_ENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_ENUMERATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_value_enumerationDB orm.ATTRIBUTE_VALUE_ENUMERATIONDB

	// fetch the attribute_value_enumeration
	_, err := db.First(&attribute_value_enumerationDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_value_enumerationDB.CopyBasicFieldsFromATTRIBUTE_VALUE_ENUMERATION_WOP(&input.ATTRIBUTE_VALUE_ENUMERATION_WOP)
	attribute_value_enumerationDB.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding = input.ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding

	db, _ = db.Model(&attribute_value_enumerationDB)
	_, err = db.Updates(&attribute_value_enumerationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_enumerationNew := new(models.ATTRIBUTE_VALUE_ENUMERATION)
	attribute_value_enumerationDB.CopyBasicFieldsToATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumerationNew)

	// redeem pointers
	attribute_value_enumerationDB.DecodePointers(backRepo, attribute_value_enumerationNew)

	// get stage instance from DB instance, and call callback function
	attribute_value_enumerationOld := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONPtr[attribute_value_enumerationDB.ID]
	if attribute_value_enumerationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attribute_value_enumerationOld, attribute_value_enumerationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_value_enumerationDB
	c.JSON(http.StatusOK, attribute_value_enumerationDB)
}

// DeleteATTRIBUTE_VALUE_ENUMERATION
//
// swagger:route DELETE /attribute_value_enumerations/{ID} attribute_value_enumerations deleteATTRIBUTE_VALUE_ENUMERATION
//
// # Delete a attribute_value_enumeration
//
// default: genericError
//
//	200: attribute_value_enumerationDBResponse
func (controller *Controller) DeleteATTRIBUTE_VALUE_ENUMERATION(c *gin.Context) {

	mutexATTRIBUTE_VALUE_ENUMERATION.Lock()
	defer mutexATTRIBUTE_VALUE_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_VALUE_ENUMERATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.GetDB()

	// Get model if exist
	var attribute_value_enumerationDB orm.ATTRIBUTE_VALUE_ENUMERATIONDB
	if _, err := db.First(&attribute_value_enumerationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_value_enumerationDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_enumerationDeleted := new(models.ATTRIBUTE_VALUE_ENUMERATION)
	attribute_value_enumerationDB.CopyBasicFieldsToATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumerationDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_value_enumerationStaged := backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONPtr[attribute_value_enumerationDB.ID]
	if attribute_value_enumerationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_value_enumerationStaged, attribute_value_enumerationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
