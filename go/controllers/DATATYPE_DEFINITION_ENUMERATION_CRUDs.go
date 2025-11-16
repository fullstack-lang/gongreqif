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
var __DATATYPE_DEFINITION_ENUMERATION__dummysDeclaration__ models.DATATYPE_DEFINITION_ENUMERATION
var __DATATYPE_DEFINITION_ENUMERATION_time__dummyDeclaration time.Duration

var mutexDATATYPE_DEFINITION_ENUMERATION sync.Mutex

// An DATATYPE_DEFINITION_ENUMERATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPE_DEFINITION_ENUMERATION updateDATATYPE_DEFINITION_ENUMERATION deleteDATATYPE_DEFINITION_ENUMERATION
type DATATYPE_DEFINITION_ENUMERATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPE_DEFINITION_ENUMERATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPE_DEFINITION_ENUMERATION updateDATATYPE_DEFINITION_ENUMERATION
type DATATYPE_DEFINITION_ENUMERATIONInput struct {
	// The DATATYPE_DEFINITION_ENUMERATION to submit or modify
	// in: body
	DATATYPE_DEFINITION_ENUMERATION *orm.DATATYPE_DEFINITION_ENUMERATIONAPI
}

// GetDATATYPE_DEFINITION_ENUMERATIONs
//
// swagger:route GET /datatype_definition_enumerations datatype_definition_enumerations getDATATYPE_DEFINITION_ENUMERATIONs
//
// # Get all datatype_definition_enumerations
//
// Responses:
// default: genericError
//
//	200: datatype_definition_enumerationDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_ENUMERATIONs(c *gin.Context) {

	// source slice
	var datatype_definition_enumerationDBs []orm.DATATYPE_DEFINITION_ENUMERATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_ENUMERATIONs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.GetDB()

	_, err := db.Find(&datatype_definition_enumerationDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatype_definition_enumerationAPIs := make([]orm.DATATYPE_DEFINITION_ENUMERATIONAPI, 0)

	// for each datatype_definition_enumeration, update fields from the database nullable fields
	for idx := range datatype_definition_enumerationDBs {
		datatype_definition_enumerationDB := &datatype_definition_enumerationDBs[idx]
		_ = datatype_definition_enumerationDB
		var datatype_definition_enumerationAPI orm.DATATYPE_DEFINITION_ENUMERATIONAPI

		// insertion point for updating fields
		datatype_definition_enumerationAPI.ID = datatype_definition_enumerationDB.ID
		datatype_definition_enumerationDB.CopyBasicFieldsToDATATYPE_DEFINITION_ENUMERATION_WOP(&datatype_definition_enumerationAPI.DATATYPE_DEFINITION_ENUMERATION_WOP)
		datatype_definition_enumerationAPI.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding = datatype_definition_enumerationDB.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding
		datatype_definition_enumerationAPIs = append(datatype_definition_enumerationAPIs, datatype_definition_enumerationAPI)
	}

	c.JSON(http.StatusOK, datatype_definition_enumerationAPIs)
}

// PostDATATYPE_DEFINITION_ENUMERATION
//
// swagger:route POST /datatype_definition_enumerations datatype_definition_enumerations postDATATYPE_DEFINITION_ENUMERATION
//
// Creates a datatype_definition_enumeration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPE_DEFINITION_ENUMERATION(c *gin.Context) {

	mutexDATATYPE_DEFINITION_ENUMERATION.Lock()
	defer mutexDATATYPE_DEFINITION_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPE_DEFINITION_ENUMERATIONs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_ENUMERATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatype_definition_enumeration
	datatype_definition_enumerationDB := orm.DATATYPE_DEFINITION_ENUMERATIONDB{}
	datatype_definition_enumerationDB.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding = input.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding
	datatype_definition_enumerationDB.CopyBasicFieldsFromDATATYPE_DEFINITION_ENUMERATION_WOP(&input.DATATYPE_DEFINITION_ENUMERATION_WOP)

	_, err = db.Create(&datatype_definition_enumerationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.CheckoutPhaseOneInstance(&datatype_definition_enumerationDB)
	datatype_definition_enumeration := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.Map_DATATYPE_DEFINITION_ENUMERATIONDBID_DATATYPE_DEFINITION_ENUMERATIONPtr[datatype_definition_enumerationDB.ID]

	if datatype_definition_enumeration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatype_definition_enumeration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatype_definition_enumerationDB)
}

// GetDATATYPE_DEFINITION_ENUMERATION
//
// swagger:route GET /datatype_definition_enumerations/{ID} datatype_definition_enumerations getDATATYPE_DEFINITION_ENUMERATION
//
// Gets the details for a datatype_definition_enumeration.
//
// Responses:
// default: genericError
//
//	200: datatype_definition_enumerationDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_ENUMERATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_ENUMERATION", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.GetDB()

	// Get datatype_definition_enumerationDB in DB
	var datatype_definition_enumerationDB orm.DATATYPE_DEFINITION_ENUMERATIONDB
	if _, err := db.First(&datatype_definition_enumerationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatype_definition_enumerationAPI orm.DATATYPE_DEFINITION_ENUMERATIONAPI
	datatype_definition_enumerationAPI.ID = datatype_definition_enumerationDB.ID
	datatype_definition_enumerationAPI.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding = datatype_definition_enumerationDB.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding
	datatype_definition_enumerationDB.CopyBasicFieldsToDATATYPE_DEFINITION_ENUMERATION_WOP(&datatype_definition_enumerationAPI.DATATYPE_DEFINITION_ENUMERATION_WOP)

	c.JSON(http.StatusOK, datatype_definition_enumerationAPI)
}

// UpdateDATATYPE_DEFINITION_ENUMERATION
//
// swagger:route PATCH /datatype_definition_enumerations/{ID} datatype_definition_enumerations updateDATATYPE_DEFINITION_ENUMERATION
//
// # Update a datatype_definition_enumeration
//
// Responses:
// default: genericError
//
//	200: datatype_definition_enumerationDBResponse
func (controller *Controller) UpdateDATATYPE_DEFINITION_ENUMERATION(c *gin.Context) {

	mutexDATATYPE_DEFINITION_ENUMERATION.Lock()
	defer mutexDATATYPE_DEFINITION_ENUMERATION.Unlock()

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
	db := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_ENUMERATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatype_definition_enumerationDB orm.DATATYPE_DEFINITION_ENUMERATIONDB

	// fetch the datatype_definition_enumeration
	_, err := db.First(&datatype_definition_enumerationDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatype_definition_enumerationDB.CopyBasicFieldsFromDATATYPE_DEFINITION_ENUMERATION_WOP(&input.DATATYPE_DEFINITION_ENUMERATION_WOP)
	datatype_definition_enumerationDB.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding = input.DATATYPE_DEFINITION_ENUMERATIONPointersEncoding

	db, _ = db.Model(&datatype_definition_enumerationDB)
	_, err = db.Updates(&datatype_definition_enumerationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_enumerationNew := new(models.DATATYPE_DEFINITION_ENUMERATION)
	datatype_definition_enumerationDB.CopyBasicFieldsToDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumerationNew)

	// redeem pointers
	datatype_definition_enumerationDB.DecodePointers(backRepo, datatype_definition_enumerationNew)

	// get stage instance from DB instance, and call callback function
	datatype_definition_enumerationOld := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.Map_DATATYPE_DEFINITION_ENUMERATIONDBID_DATATYPE_DEFINITION_ENUMERATIONPtr[datatype_definition_enumerationDB.ID]
	if datatype_definition_enumerationOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), datatype_definition_enumerationOld, datatype_definition_enumerationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatype_definition_enumerationDB
	c.JSON(http.StatusOK, datatype_definition_enumerationDB)
}

// DeleteDATATYPE_DEFINITION_ENUMERATION
//
// swagger:route DELETE /datatype_definition_enumerations/{ID} datatype_definition_enumerations deleteDATATYPE_DEFINITION_ENUMERATION
//
// # Delete a datatype_definition_enumeration
//
// default: genericError
//
//	200: datatype_definition_enumerationDBResponse
func (controller *Controller) DeleteDATATYPE_DEFINITION_ENUMERATION(c *gin.Context) {

	mutexDATATYPE_DEFINITION_ENUMERATION.Lock()
	defer mutexDATATYPE_DEFINITION_ENUMERATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPE_DEFINITION_ENUMERATION", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.GetDB()

	// Get model if exist
	var datatype_definition_enumerationDB orm.DATATYPE_DEFINITION_ENUMERATIONDB
	if _, err := db.First(&datatype_definition_enumerationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&datatype_definition_enumerationDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_enumerationDeleted := new(models.DATATYPE_DEFINITION_ENUMERATION)
	datatype_definition_enumerationDB.CopyBasicFieldsToDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumerationDeleted)

	// get stage instance from DB instance, and call callback function
	datatype_definition_enumerationStaged := backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.Map_DATATYPE_DEFINITION_ENUMERATIONDBID_DATATYPE_DEFINITION_ENUMERATIONPtr[datatype_definition_enumerationDB.ID]
	if datatype_definition_enumerationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatype_definition_enumerationStaged, datatype_definition_enumerationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
