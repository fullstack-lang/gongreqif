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
var __EMBEDDED_VALUE__dummysDeclaration__ models.EMBEDDED_VALUE
var __EMBEDDED_VALUE_time__dummyDeclaration time.Duration

var mutexEMBEDDED_VALUE sync.Mutex

// An EMBEDDED_VALUEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getEMBEDDED_VALUE updateEMBEDDED_VALUE deleteEMBEDDED_VALUE
type EMBEDDED_VALUEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// EMBEDDED_VALUEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postEMBEDDED_VALUE updateEMBEDDED_VALUE
type EMBEDDED_VALUEInput struct {
	// The EMBEDDED_VALUE to submit or modify
	// in: body
	EMBEDDED_VALUE *orm.EMBEDDED_VALUEAPI
}

// GetEMBEDDED_VALUEs
//
// swagger:route GET /embedded_values embedded_values getEMBEDDED_VALUEs
//
// # Get all embedded_values
//
// Responses:
// default: genericError
//
//	200: embedded_valueDBResponse
func (controller *Controller) GetEMBEDDED_VALUEs(c *gin.Context) {

	// source slice
	var embedded_valueDBs []orm.EMBEDDED_VALUEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEMBEDDED_VALUEs", "Name", stackPath)
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
	db := backRepo.BackRepoEMBEDDED_VALUE.GetDB()

	_, err := db.Find(&embedded_valueDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	embedded_valueAPIs := make([]orm.EMBEDDED_VALUEAPI, 0)

	// for each embedded_value, update fields from the database nullable fields
	for idx := range embedded_valueDBs {
		embedded_valueDB := &embedded_valueDBs[idx]
		_ = embedded_valueDB
		var embedded_valueAPI orm.EMBEDDED_VALUEAPI

		// insertion point for updating fields
		embedded_valueAPI.ID = embedded_valueDB.ID
		embedded_valueDB.CopyBasicFieldsToEMBEDDED_VALUE_WOP(&embedded_valueAPI.EMBEDDED_VALUE_WOP)
		embedded_valueAPI.EMBEDDED_VALUEPointersEncoding = embedded_valueDB.EMBEDDED_VALUEPointersEncoding
		embedded_valueAPIs = append(embedded_valueAPIs, embedded_valueAPI)
	}

	c.JSON(http.StatusOK, embedded_valueAPIs)
}

// PostEMBEDDED_VALUE
//
// swagger:route POST /embedded_values embedded_values postEMBEDDED_VALUE
//
// Creates a embedded_value
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostEMBEDDED_VALUE(c *gin.Context) {

	mutexEMBEDDED_VALUE.Lock()
	defer mutexEMBEDDED_VALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostEMBEDDED_VALUEs", "Name", stackPath)
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
	db := backRepo.BackRepoEMBEDDED_VALUE.GetDB()

	// Validate input
	var input orm.EMBEDDED_VALUEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create embedded_value
	embedded_valueDB := orm.EMBEDDED_VALUEDB{}
	embedded_valueDB.EMBEDDED_VALUEPointersEncoding = input.EMBEDDED_VALUEPointersEncoding
	embedded_valueDB.CopyBasicFieldsFromEMBEDDED_VALUE_WOP(&input.EMBEDDED_VALUE_WOP)

	_, err = db.Create(&embedded_valueDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoEMBEDDED_VALUE.CheckoutPhaseOneInstance(&embedded_valueDB)
	embedded_value := backRepo.BackRepoEMBEDDED_VALUE.Map_EMBEDDED_VALUEDBID_EMBEDDED_VALUEPtr[embedded_valueDB.ID]

	if embedded_value != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), embedded_value)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, embedded_valueDB)
}

// GetEMBEDDED_VALUE
//
// swagger:route GET /embedded_values/{ID} embedded_values getEMBEDDED_VALUE
//
// Gets the details for a embedded_value.
//
// Responses:
// default: genericError
//
//	200: embedded_valueDBResponse
func (controller *Controller) GetEMBEDDED_VALUE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetEMBEDDED_VALUE", "Name", stackPath)
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
	db := backRepo.BackRepoEMBEDDED_VALUE.GetDB()

	// Get embedded_valueDB in DB
	var embedded_valueDB orm.EMBEDDED_VALUEDB
	if _, err := db.First(&embedded_valueDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var embedded_valueAPI orm.EMBEDDED_VALUEAPI
	embedded_valueAPI.ID = embedded_valueDB.ID
	embedded_valueAPI.EMBEDDED_VALUEPointersEncoding = embedded_valueDB.EMBEDDED_VALUEPointersEncoding
	embedded_valueDB.CopyBasicFieldsToEMBEDDED_VALUE_WOP(&embedded_valueAPI.EMBEDDED_VALUE_WOP)

	c.JSON(http.StatusOK, embedded_valueAPI)
}

// UpdateEMBEDDED_VALUE
//
// swagger:route PATCH /embedded_values/{ID} embedded_values updateEMBEDDED_VALUE
//
// # Update a embedded_value
//
// Responses:
// default: genericError
//
//	200: embedded_valueDBResponse
func (controller *Controller) UpdateEMBEDDED_VALUE(c *gin.Context) {

	mutexEMBEDDED_VALUE.Lock()
	defer mutexEMBEDDED_VALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateEMBEDDED_VALUE", "Name", stackPath)
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
	db := backRepo.BackRepoEMBEDDED_VALUE.GetDB()

	// Validate input
	var input orm.EMBEDDED_VALUEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var embedded_valueDB orm.EMBEDDED_VALUEDB

	// fetch the embedded_value
	_, err := db.First(&embedded_valueDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	embedded_valueDB.CopyBasicFieldsFromEMBEDDED_VALUE_WOP(&input.EMBEDDED_VALUE_WOP)
	embedded_valueDB.EMBEDDED_VALUEPointersEncoding = input.EMBEDDED_VALUEPointersEncoding

	db, _ = db.Model(&embedded_valueDB)
	_, err = db.Updates(&embedded_valueDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	embedded_valueNew := new(models.EMBEDDED_VALUE)
	embedded_valueDB.CopyBasicFieldsToEMBEDDED_VALUE(embedded_valueNew)

	// redeem pointers
	embedded_valueDB.DecodePointers(backRepo, embedded_valueNew)

	// get stage instance from DB instance, and call callback function
	embedded_valueOld := backRepo.BackRepoEMBEDDED_VALUE.Map_EMBEDDED_VALUEDBID_EMBEDDED_VALUEPtr[embedded_valueDB.ID]
	if embedded_valueOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), embedded_valueOld, embedded_valueNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the embedded_valueDB
	c.JSON(http.StatusOK, embedded_valueDB)
}

// DeleteEMBEDDED_VALUE
//
// swagger:route DELETE /embedded_values/{ID} embedded_values deleteEMBEDDED_VALUE
//
// # Delete a embedded_value
//
// default: genericError
//
//	200: embedded_valueDBResponse
func (controller *Controller) DeleteEMBEDDED_VALUE(c *gin.Context) {

	mutexEMBEDDED_VALUE.Lock()
	defer mutexEMBEDDED_VALUE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteEMBEDDED_VALUE", "Name", stackPath)
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
	db := backRepo.BackRepoEMBEDDED_VALUE.GetDB()

	// Get model if exist
	var embedded_valueDB orm.EMBEDDED_VALUEDB
	if _, err := db.First(&embedded_valueDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&embedded_valueDB)

	// get an instance (not staged) from DB instance, and call callback function
	embedded_valueDeleted := new(models.EMBEDDED_VALUE)
	embedded_valueDB.CopyBasicFieldsToEMBEDDED_VALUE(embedded_valueDeleted)

	// get stage instance from DB instance, and call callback function
	embedded_valueStaged := backRepo.BackRepoEMBEDDED_VALUE.Map_EMBEDDED_VALUEDBID_EMBEDDED_VALUEPtr[embedded_valueDB.ID]
	if embedded_valueStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), embedded_valueStaged, embedded_valueDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
