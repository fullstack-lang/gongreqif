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
var __ATTRIBUTE_VALUE_REAL__dummysDeclaration__ models.ATTRIBUTE_VALUE_REAL
var __ATTRIBUTE_VALUE_REAL_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_VALUE_REAL sync.Mutex

// An ATTRIBUTE_VALUE_REALID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_VALUE_REAL updateATTRIBUTE_VALUE_REAL deleteATTRIBUTE_VALUE_REAL
type ATTRIBUTE_VALUE_REALID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_VALUE_REALInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_VALUE_REAL updateATTRIBUTE_VALUE_REAL
type ATTRIBUTE_VALUE_REALInput struct {
	// The ATTRIBUTE_VALUE_REAL to submit or modify
	// in: body
	ATTRIBUTE_VALUE_REAL *orm.ATTRIBUTE_VALUE_REALAPI
}

// GetATTRIBUTE_VALUE_REALs
//
// swagger:route GET /attribute_value_reals attribute_value_reals getATTRIBUTE_VALUE_REALs
//
// # Get all attribute_value_reals
//
// Responses:
// default: genericError
//
//	200: attribute_value_realDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_REALs(c *gin.Context) {

	// source slice
	var attribute_value_realDBs []orm.ATTRIBUTE_VALUE_REALDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_REALs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_REAL.GetDB()

	_, err := db.Find(&attribute_value_realDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_value_realAPIs := make([]orm.ATTRIBUTE_VALUE_REALAPI, 0)

	// for each attribute_value_real, update fields from the database nullable fields
	for idx := range attribute_value_realDBs {
		attribute_value_realDB := &attribute_value_realDBs[idx]
		_ = attribute_value_realDB
		var attribute_value_realAPI orm.ATTRIBUTE_VALUE_REALAPI

		// insertion point for updating fields
		attribute_value_realAPI.ID = attribute_value_realDB.ID
		attribute_value_realDB.CopyBasicFieldsToATTRIBUTE_VALUE_REAL_WOP(&attribute_value_realAPI.ATTRIBUTE_VALUE_REAL_WOP)
		attribute_value_realAPI.ATTRIBUTE_VALUE_REALPointersEncoding = attribute_value_realDB.ATTRIBUTE_VALUE_REALPointersEncoding
		attribute_value_realAPIs = append(attribute_value_realAPIs, attribute_value_realAPI)
	}

	c.JSON(http.StatusOK, attribute_value_realAPIs)
}

// PostATTRIBUTE_VALUE_REAL
//
// swagger:route POST /attribute_value_reals attribute_value_reals postATTRIBUTE_VALUE_REAL
//
// Creates a attribute_value_real
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_VALUE_REAL(c *gin.Context) {

	mutexATTRIBUTE_VALUE_REAL.Lock()
	defer mutexATTRIBUTE_VALUE_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_VALUE_REALs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_REAL.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_REALAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_value_real
	attribute_value_realDB := orm.ATTRIBUTE_VALUE_REALDB{}
	attribute_value_realDB.ATTRIBUTE_VALUE_REALPointersEncoding = input.ATTRIBUTE_VALUE_REALPointersEncoding
	attribute_value_realDB.CopyBasicFieldsFromATTRIBUTE_VALUE_REAL_WOP(&input.ATTRIBUTE_VALUE_REAL_WOP)

	_, err = db.Create(&attribute_value_realDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.CheckoutPhaseOneInstance(&attribute_value_realDB)
	attribute_value_real := backRepo.BackRepoATTRIBUTE_VALUE_REAL.Map_ATTRIBUTE_VALUE_REALDBID_ATTRIBUTE_VALUE_REALPtr[attribute_value_realDB.ID]

	if attribute_value_real != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_value_real)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_value_realDB)
}

// GetATTRIBUTE_VALUE_REAL
//
// swagger:route GET /attribute_value_reals/{ID} attribute_value_reals getATTRIBUTE_VALUE_REAL
//
// Gets the details for a attribute_value_real.
//
// Responses:
// default: genericError
//
//	200: attribute_value_realDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_REAL(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_REAL", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_REAL.GetDB()

	// Get attribute_value_realDB in DB
	var attribute_value_realDB orm.ATTRIBUTE_VALUE_REALDB
	if _, err := db.First(&attribute_value_realDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_value_realAPI orm.ATTRIBUTE_VALUE_REALAPI
	attribute_value_realAPI.ID = attribute_value_realDB.ID
	attribute_value_realAPI.ATTRIBUTE_VALUE_REALPointersEncoding = attribute_value_realDB.ATTRIBUTE_VALUE_REALPointersEncoding
	attribute_value_realDB.CopyBasicFieldsToATTRIBUTE_VALUE_REAL_WOP(&attribute_value_realAPI.ATTRIBUTE_VALUE_REAL_WOP)

	c.JSON(http.StatusOK, attribute_value_realAPI)
}

// UpdateATTRIBUTE_VALUE_REAL
//
// swagger:route PATCH /attribute_value_reals/{ID} attribute_value_reals updateATTRIBUTE_VALUE_REAL
//
// # Update a attribute_value_real
//
// Responses:
// default: genericError
//
//	200: attribute_value_realDBResponse
func (controller *Controller) UpdateATTRIBUTE_VALUE_REAL(c *gin.Context) {

	mutexATTRIBUTE_VALUE_REAL.Lock()
	defer mutexATTRIBUTE_VALUE_REAL.Unlock()

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
	db := backRepo.BackRepoATTRIBUTE_VALUE_REAL.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_REALAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_value_realDB orm.ATTRIBUTE_VALUE_REALDB

	// fetch the attribute_value_real
	_, err := db.First(&attribute_value_realDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_value_realDB.CopyBasicFieldsFromATTRIBUTE_VALUE_REAL_WOP(&input.ATTRIBUTE_VALUE_REAL_WOP)
	attribute_value_realDB.ATTRIBUTE_VALUE_REALPointersEncoding = input.ATTRIBUTE_VALUE_REALPointersEncoding

	db, _ = db.Model(&attribute_value_realDB)
	_, err = db.Updates(&attribute_value_realDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_realNew := new(models.ATTRIBUTE_VALUE_REAL)
	attribute_value_realDB.CopyBasicFieldsToATTRIBUTE_VALUE_REAL(attribute_value_realNew)

	// redeem pointers
	attribute_value_realDB.DecodePointers(backRepo, attribute_value_realNew)

	// get stage instance from DB instance, and call callback function
	attribute_value_realOld := backRepo.BackRepoATTRIBUTE_VALUE_REAL.Map_ATTRIBUTE_VALUE_REALDBID_ATTRIBUTE_VALUE_REALPtr[attribute_value_realDB.ID]
	if attribute_value_realOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_value_realOld, attribute_value_realNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_value_realDB
	c.JSON(http.StatusOK, attribute_value_realDB)
}

// DeleteATTRIBUTE_VALUE_REAL
//
// swagger:route DELETE /attribute_value_reals/{ID} attribute_value_reals deleteATTRIBUTE_VALUE_REAL
//
// # Delete a attribute_value_real
//
// default: genericError
//
//	200: attribute_value_realDBResponse
func (controller *Controller) DeleteATTRIBUTE_VALUE_REAL(c *gin.Context) {

	mutexATTRIBUTE_VALUE_REAL.Lock()
	defer mutexATTRIBUTE_VALUE_REAL.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_VALUE_REAL", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_REAL.GetDB()

	// Get model if exist
	var attribute_value_realDB orm.ATTRIBUTE_VALUE_REALDB
	if _, err := db.First(&attribute_value_realDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_value_realDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_realDeleted := new(models.ATTRIBUTE_VALUE_REAL)
	attribute_value_realDB.CopyBasicFieldsToATTRIBUTE_VALUE_REAL(attribute_value_realDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_value_realStaged := backRepo.BackRepoATTRIBUTE_VALUE_REAL.Map_ATTRIBUTE_VALUE_REALDBID_ATTRIBUTE_VALUE_REALPtr[attribute_value_realDB.ID]
	if attribute_value_realStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_value_realStaged, attribute_value_realDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
