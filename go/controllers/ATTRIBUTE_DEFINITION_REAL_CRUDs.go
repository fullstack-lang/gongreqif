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
var __ATTRIBUTE_DEFINITION_REAL__dummysDeclaration__ models.ATTRIBUTE_DEFINITION_REAL
var __ATTRIBUTE_DEFINITION_REAL_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_DEFINITION_REAL sync.Mutex

// An ATTRIBUTE_DEFINITION_REALID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_DEFINITION_REAL updateATTRIBUTE_DEFINITION_REAL deleteATTRIBUTE_DEFINITION_REAL
type ATTRIBUTE_DEFINITION_REALID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_DEFINITION_REALInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_DEFINITION_REAL updateATTRIBUTE_DEFINITION_REAL
type ATTRIBUTE_DEFINITION_REALInput struct {
	// The ATTRIBUTE_DEFINITION_REAL to submit or modify
	// in: body
	ATTRIBUTE_DEFINITION_REAL *orm.ATTRIBUTE_DEFINITION_REALAPI
}

// GetATTRIBUTE_DEFINITION_REALs
//
// swagger:route GET /attribute_definition_reals attribute_definition_reals getATTRIBUTE_DEFINITION_REALs
//
// # Get all attribute_definition_reals
//
// Responses:
// default: genericError
//
//	200: attribute_definition_realDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_REALs(c *gin.Context) {

	// source slice
	var attribute_definition_realDBs []orm.ATTRIBUTE_DEFINITION_REALDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_REALs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.GetDB()

	_, err := db.Find(&attribute_definition_realDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_definition_realAPIs := make([]orm.ATTRIBUTE_DEFINITION_REALAPI, 0)

	// for each attribute_definition_real, update fields from the database nullable fields
	for idx := range attribute_definition_realDBs {
		attribute_definition_realDB := &attribute_definition_realDBs[idx]
		_ = attribute_definition_realDB
		var attribute_definition_realAPI orm.ATTRIBUTE_DEFINITION_REALAPI

		// insertion point for updating fields
		attribute_definition_realAPI.ID = attribute_definition_realDB.ID
		attribute_definition_realDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_REAL_WOP(&attribute_definition_realAPI.ATTRIBUTE_DEFINITION_REAL_WOP)
		attribute_definition_realAPI.ATTRIBUTE_DEFINITION_REALPointersEncoding = attribute_definition_realDB.ATTRIBUTE_DEFINITION_REALPointersEncoding
		attribute_definition_realAPIs = append(attribute_definition_realAPIs, attribute_definition_realAPI)
	}

	c.JSON(http.StatusOK, attribute_definition_realAPIs)
}

// PostATTRIBUTE_DEFINITION_REAL
//
// swagger:route POST /attribute_definition_reals attribute_definition_reals postATTRIBUTE_DEFINITION_REAL
//
// Creates a attribute_definition_real
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_DEFINITION_REAL(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_REAL.Lock()
	defer mutexATTRIBUTE_DEFINITION_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_DEFINITION_REALs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_REALAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_definition_real
	attribute_definition_realDB := orm.ATTRIBUTE_DEFINITION_REALDB{}
	attribute_definition_realDB.ATTRIBUTE_DEFINITION_REALPointersEncoding = input.ATTRIBUTE_DEFINITION_REALPointersEncoding
	attribute_definition_realDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_REAL_WOP(&input.ATTRIBUTE_DEFINITION_REAL_WOP)

	_, err = db.Create(&attribute_definition_realDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.CheckoutPhaseOneInstance(&attribute_definition_realDB)
	attribute_definition_real := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.Map_ATTRIBUTE_DEFINITION_REALDBID_ATTRIBUTE_DEFINITION_REALPtr[attribute_definition_realDB.ID]

	if attribute_definition_real != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_definition_real)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_definition_realDB)
}

// GetATTRIBUTE_DEFINITION_REAL
//
// swagger:route GET /attribute_definition_reals/{ID} attribute_definition_reals getATTRIBUTE_DEFINITION_REAL
//
// Gets the details for a attribute_definition_real.
//
// Responses:
// default: genericError
//
//	200: attribute_definition_realDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_REAL(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_REAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.GetDB()

	// Get attribute_definition_realDB in DB
	var attribute_definition_realDB orm.ATTRIBUTE_DEFINITION_REALDB
	if _, err := db.First(&attribute_definition_realDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_definition_realAPI orm.ATTRIBUTE_DEFINITION_REALAPI
	attribute_definition_realAPI.ID = attribute_definition_realDB.ID
	attribute_definition_realAPI.ATTRIBUTE_DEFINITION_REALPointersEncoding = attribute_definition_realDB.ATTRIBUTE_DEFINITION_REALPointersEncoding
	attribute_definition_realDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_REAL_WOP(&attribute_definition_realAPI.ATTRIBUTE_DEFINITION_REAL_WOP)

	c.JSON(http.StatusOK, attribute_definition_realAPI)
}

// UpdateATTRIBUTE_DEFINITION_REAL
//
// swagger:route PATCH /attribute_definition_reals/{ID} attribute_definition_reals updateATTRIBUTE_DEFINITION_REAL
//
// # Update a attribute_definition_real
//
// Responses:
// default: genericError
//
//	200: attribute_definition_realDBResponse
func (controller *Controller) UpdateATTRIBUTE_DEFINITION_REAL(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_REAL.Lock()
	defer mutexATTRIBUTE_DEFINITION_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateATTRIBUTE_DEFINITION_REAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_REALAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_definition_realDB orm.ATTRIBUTE_DEFINITION_REALDB

	// fetch the attribute_definition_real
	_, err := db.First(&attribute_definition_realDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_definition_realDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_REAL_WOP(&input.ATTRIBUTE_DEFINITION_REAL_WOP)
	attribute_definition_realDB.ATTRIBUTE_DEFINITION_REALPointersEncoding = input.ATTRIBUTE_DEFINITION_REALPointersEncoding

	db, _ = db.Model(&attribute_definition_realDB)
	_, err = db.Updates(&attribute_definition_realDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_realNew := new(models.ATTRIBUTE_DEFINITION_REAL)
	attribute_definition_realDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_REAL(attribute_definition_realNew)

	// redeem pointers
	attribute_definition_realDB.DecodePointers(backRepo, attribute_definition_realNew)

	// get stage instance from DB instance, and call callback function
	attribute_definition_realOld := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.Map_ATTRIBUTE_DEFINITION_REALDBID_ATTRIBUTE_DEFINITION_REALPtr[attribute_definition_realDB.ID]
	if attribute_definition_realOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), attribute_definition_realOld, attribute_definition_realNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_definition_realDB
	c.JSON(http.StatusOK, attribute_definition_realDB)
}

// DeleteATTRIBUTE_DEFINITION_REAL
//
// swagger:route DELETE /attribute_definition_reals/{ID} attribute_definition_reals deleteATTRIBUTE_DEFINITION_REAL
//
// # Delete a attribute_definition_real
//
// default: genericError
//
//	200: attribute_definition_realDBResponse
func (controller *Controller) DeleteATTRIBUTE_DEFINITION_REAL(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_REAL.Lock()
	defer mutexATTRIBUTE_DEFINITION_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_DEFINITION_REAL", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.GetDB()

	// Get model if exist
	var attribute_definition_realDB orm.ATTRIBUTE_DEFINITION_REALDB
	if _, err := db.First(&attribute_definition_realDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_definition_realDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_realDeleted := new(models.ATTRIBUTE_DEFINITION_REAL)
	attribute_definition_realDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_REAL(attribute_definition_realDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_definition_realStaged := backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.Map_ATTRIBUTE_DEFINITION_REALDBID_ATTRIBUTE_DEFINITION_REALPtr[attribute_definition_realDB.ID]
	if attribute_definition_realStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_definition_realStaged, attribute_definition_realDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
