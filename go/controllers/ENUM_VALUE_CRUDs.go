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
var __ENUM_VALUE__dummysDeclaration__ models.ENUM_VALUE
var __ENUM_VALUE_time__dummyDeclaration time.Duration

var mutexENUM_VALUE sync.Mutex

// An ENUM_VALUEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getENUM_VALUE updateENUM_VALUE deleteENUM_VALUE
type ENUM_VALUEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ENUM_VALUEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postENUM_VALUE updateENUM_VALUE
type ENUM_VALUEInput struct {
	// The ENUM_VALUE to submit or modify
	// in: body
	ENUM_VALUE *orm.ENUM_VALUEAPI
}

// GetENUM_VALUEs
//
// swagger:route GET /enum_values enum_values getENUM_VALUEs
//
// # Get all enum_values
//
// Responses:
// default: genericError
//
//	200: enum_valueDBResponse
func (controller *Controller) GetENUM_VALUEs(c *gin.Context) {

	// source slice
	var enum_valueDBs []orm.ENUM_VALUEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetENUM_VALUEs", "Name", stackPath)
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
	db := backRepo.BackRepoENUM_VALUE.GetDB()

	_, err := db.Find(&enum_valueDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	enum_valueAPIs := make([]orm.ENUM_VALUEAPI, 0)

	// for each enum_value, update fields from the database nullable fields
	for idx := range enum_valueDBs {
		enum_valueDB := &enum_valueDBs[idx]
		_ = enum_valueDB
		var enum_valueAPI orm.ENUM_VALUEAPI

		// insertion point for updating fields
		enum_valueAPI.ID = enum_valueDB.ID
		enum_valueDB.CopyBasicFieldsToENUM_VALUE_WOP(&enum_valueAPI.ENUM_VALUE_WOP)
		enum_valueAPI.ENUM_VALUEPointersEncoding = enum_valueDB.ENUM_VALUEPointersEncoding
		enum_valueAPIs = append(enum_valueAPIs, enum_valueAPI)
	}

	c.JSON(http.StatusOK, enum_valueAPIs)
}

// PostENUM_VALUE
//
// swagger:route POST /enum_values enum_values postENUM_VALUE
//
// Creates a enum_value
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostENUM_VALUE(c *gin.Context) {

	mutexENUM_VALUE.Lock()
	defer mutexENUM_VALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostENUM_VALUEs", "Name", stackPath)
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
	db := backRepo.BackRepoENUM_VALUE.GetDB()

	// Validate input
	var input orm.ENUM_VALUEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create enum_value
	enum_valueDB := orm.ENUM_VALUEDB{}
	enum_valueDB.ENUM_VALUEPointersEncoding = input.ENUM_VALUEPointersEncoding
	enum_valueDB.CopyBasicFieldsFromENUM_VALUE_WOP(&input.ENUM_VALUE_WOP)

	_, err = db.Create(&enum_valueDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoENUM_VALUE.CheckoutPhaseOneInstance(&enum_valueDB)
	enum_value := backRepo.BackRepoENUM_VALUE.Map_ENUM_VALUEDBID_ENUM_VALUEPtr[enum_valueDB.ID]

	if enum_value != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), enum_value)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, enum_valueDB)
}

// GetENUM_VALUE
//
// swagger:route GET /enum_values/{ID} enum_values getENUM_VALUE
//
// Gets the details for a enum_value.
//
// Responses:
// default: genericError
//
//	200: enum_valueDBResponse
func (controller *Controller) GetENUM_VALUE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetENUM_VALUE", "Name", stackPath)
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
	db := backRepo.BackRepoENUM_VALUE.GetDB()

	// Get enum_valueDB in DB
	var enum_valueDB orm.ENUM_VALUEDB
	if _, err := db.First(&enum_valueDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var enum_valueAPI orm.ENUM_VALUEAPI
	enum_valueAPI.ID = enum_valueDB.ID
	enum_valueAPI.ENUM_VALUEPointersEncoding = enum_valueDB.ENUM_VALUEPointersEncoding
	enum_valueDB.CopyBasicFieldsToENUM_VALUE_WOP(&enum_valueAPI.ENUM_VALUE_WOP)

	c.JSON(http.StatusOK, enum_valueAPI)
}

// UpdateENUM_VALUE
//
// swagger:route PATCH /enum_values/{ID} enum_values updateENUM_VALUE
//
// # Update a enum_value
//
// Responses:
// default: genericError
//
//	200: enum_valueDBResponse
func (controller *Controller) UpdateENUM_VALUE(c *gin.Context) {

	mutexENUM_VALUE.Lock()
	defer mutexENUM_VALUE.Unlock()

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
	db := backRepo.BackRepoENUM_VALUE.GetDB()

	// Validate input
	var input orm.ENUM_VALUEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var enum_valueDB orm.ENUM_VALUEDB

	// fetch the enum_value
	_, err := db.First(&enum_valueDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	enum_valueDB.CopyBasicFieldsFromENUM_VALUE_WOP(&input.ENUM_VALUE_WOP)
	enum_valueDB.ENUM_VALUEPointersEncoding = input.ENUM_VALUEPointersEncoding

	db, _ = db.Model(&enum_valueDB)
	_, err = db.Updates(&enum_valueDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	enum_valueNew := new(models.ENUM_VALUE)
	enum_valueDB.CopyBasicFieldsToENUM_VALUE(enum_valueNew)

	// redeem pointers
	enum_valueDB.DecodePointers(backRepo, enum_valueNew)

	// get stage instance from DB instance, and call callback function
	enum_valueOld := backRepo.BackRepoENUM_VALUE.Map_ENUM_VALUEDBID_ENUM_VALUEPtr[enum_valueDB.ID]
	if enum_valueOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), enum_valueOld, enum_valueNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the enum_valueDB
	c.JSON(http.StatusOK, enum_valueDB)
}

// DeleteENUM_VALUE
//
// swagger:route DELETE /enum_values/{ID} enum_values deleteENUM_VALUE
//
// # Delete a enum_value
//
// default: genericError
//
//	200: enum_valueDBResponse
func (controller *Controller) DeleteENUM_VALUE(c *gin.Context) {

	mutexENUM_VALUE.Lock()
	defer mutexENUM_VALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteENUM_VALUE", "Name", stackPath)
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
	db := backRepo.BackRepoENUM_VALUE.GetDB()

	// Get model if exist
	var enum_valueDB orm.ENUM_VALUEDB
	if _, err := db.First(&enum_valueDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&enum_valueDB)

	// get an instance (not staged) from DB instance, and call callback function
	enum_valueDeleted := new(models.ENUM_VALUE)
	enum_valueDB.CopyBasicFieldsToENUM_VALUE(enum_valueDeleted)

	// get stage instance from DB instance, and call callback function
	enum_valueStaged := backRepo.BackRepoENUM_VALUE.Map_ENUM_VALUEDBID_ENUM_VALUEPtr[enum_valueDB.ID]
	if enum_valueStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), enum_valueStaged, enum_valueDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
