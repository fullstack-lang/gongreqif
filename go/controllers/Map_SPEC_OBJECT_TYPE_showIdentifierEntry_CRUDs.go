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
var __Map_SPEC_OBJECT_TYPE_showIdentifierEntry__dummysDeclaration__ models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry
var __Map_SPEC_OBJECT_TYPE_showIdentifierEntry_time__dummyDeclaration time.Duration

var mutexMap_SPEC_OBJECT_TYPE_showIdentifierEntry sync.Mutex

// An Map_SPEC_OBJECT_TYPE_showIdentifierEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMap_SPEC_OBJECT_TYPE_showIdentifierEntry updateMap_SPEC_OBJECT_TYPE_showIdentifierEntry deleteMap_SPEC_OBJECT_TYPE_showIdentifierEntry
type Map_SPEC_OBJECT_TYPE_showIdentifierEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Map_SPEC_OBJECT_TYPE_showIdentifierEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMap_SPEC_OBJECT_TYPE_showIdentifierEntry updateMap_SPEC_OBJECT_TYPE_showIdentifierEntry
type Map_SPEC_OBJECT_TYPE_showIdentifierEntryInput struct {
	// The Map_SPEC_OBJECT_TYPE_showIdentifierEntry to submit or modify
	// in: body
	Map_SPEC_OBJECT_TYPE_showIdentifierEntry *orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryAPI
}

// GetMap_SPEC_OBJECT_TYPE_showIdentifierEntrys
//
// swagger:route GET /map_spec_object_type_showidentifierentrys map_spec_object_type_showidentifierentrys getMap_SPEC_OBJECT_TYPE_showIdentifierEntrys
//
// # Get all map_spec_object_type_showidentifierentrys
//
// Responses:
// default: genericError
//
//	200: map_spec_object_type_showidentifierentryDBResponse
func (controller *Controller) GetMap_SPEC_OBJECT_TYPE_showIdentifierEntrys(c *gin.Context) {

	// source slice
	var map_spec_object_type_showidentifierentryDBs []orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_SPEC_OBJECT_TYPE_showIdentifierEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.GetDB()

	_, err := db.Find(&map_spec_object_type_showidentifierentryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	map_spec_object_type_showidentifierentryAPIs := make([]orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryAPI, 0)

	// for each map_spec_object_type_showidentifierentry, update fields from the database nullable fields
	for idx := range map_spec_object_type_showidentifierentryDBs {
		map_spec_object_type_showidentifierentryDB := &map_spec_object_type_showidentifierentryDBs[idx]
		_ = map_spec_object_type_showidentifierentryDB
		var map_spec_object_type_showidentifierentryAPI orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryAPI

		// insertion point for updating fields
		map_spec_object_type_showidentifierentryAPI.ID = map_spec_object_type_showidentifierentryDB.ID
		map_spec_object_type_showidentifierentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_showIdentifierEntry_WOP(&map_spec_object_type_showidentifierentryAPI.Map_SPEC_OBJECT_TYPE_showIdentifierEntry_WOP)
		map_spec_object_type_showidentifierentryAPI.Map_SPEC_OBJECT_TYPE_showIdentifierEntryPointersEncoding = map_spec_object_type_showidentifierentryDB.Map_SPEC_OBJECT_TYPE_showIdentifierEntryPointersEncoding
		map_spec_object_type_showidentifierentryAPIs = append(map_spec_object_type_showidentifierentryAPIs, map_spec_object_type_showidentifierentryAPI)
	}

	c.JSON(http.StatusOK, map_spec_object_type_showidentifierentryAPIs)
}

// PostMap_SPEC_OBJECT_TYPE_showIdentifierEntry
//
// swagger:route POST /map_spec_object_type_showidentifierentrys map_spec_object_type_showidentifierentrys postMap_SPEC_OBJECT_TYPE_showIdentifierEntry
//
// Creates a map_spec_object_type_showidentifierentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMap_SPEC_OBJECT_TYPE_showIdentifierEntry(c *gin.Context) {

	mutexMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Lock()
	defer mutexMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMap_SPEC_OBJECT_TYPE_showIdentifierEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.GetDB()

	// Validate input
	var input orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create map_spec_object_type_showidentifierentry
	map_spec_object_type_showidentifierentryDB := orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB{}
	map_spec_object_type_showidentifierentryDB.Map_SPEC_OBJECT_TYPE_showIdentifierEntryPointersEncoding = input.Map_SPEC_OBJECT_TYPE_showIdentifierEntryPointersEncoding
	map_spec_object_type_showidentifierentryDB.CopyBasicFieldsFromMap_SPEC_OBJECT_TYPE_showIdentifierEntry_WOP(&input.Map_SPEC_OBJECT_TYPE_showIdentifierEntry_WOP)

	_, err = db.Create(&map_spec_object_type_showidentifierentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.CheckoutPhaseOneInstance(&map_spec_object_type_showidentifierentryDB)
	map_spec_object_type_showidentifierentry := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Map_Map_SPEC_OBJECT_TYPE_showIdentifierEntryDBID_Map_SPEC_OBJECT_TYPE_showIdentifierEntryPtr[map_spec_object_type_showidentifierentryDB.ID]

	if map_spec_object_type_showidentifierentry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), map_spec_object_type_showidentifierentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, map_spec_object_type_showidentifierentryDB)
}

// GetMap_SPEC_OBJECT_TYPE_showIdentifierEntry
//
// swagger:route GET /map_spec_object_type_showidentifierentrys/{ID} map_spec_object_type_showidentifierentrys getMap_SPEC_OBJECT_TYPE_showIdentifierEntry
//
// Gets the details for a map_spec_object_type_showidentifierentry.
//
// Responses:
// default: genericError
//
//	200: map_spec_object_type_showidentifierentryDBResponse
func (controller *Controller) GetMap_SPEC_OBJECT_TYPE_showIdentifierEntry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_SPEC_OBJECT_TYPE_showIdentifierEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.GetDB()

	// Get map_spec_object_type_showidentifierentryDB in DB
	var map_spec_object_type_showidentifierentryDB orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB
	if _, err := db.First(&map_spec_object_type_showidentifierentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var map_spec_object_type_showidentifierentryAPI orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryAPI
	map_spec_object_type_showidentifierentryAPI.ID = map_spec_object_type_showidentifierentryDB.ID
	map_spec_object_type_showidentifierentryAPI.Map_SPEC_OBJECT_TYPE_showIdentifierEntryPointersEncoding = map_spec_object_type_showidentifierentryDB.Map_SPEC_OBJECT_TYPE_showIdentifierEntryPointersEncoding
	map_spec_object_type_showidentifierentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_showIdentifierEntry_WOP(&map_spec_object_type_showidentifierentryAPI.Map_SPEC_OBJECT_TYPE_showIdentifierEntry_WOP)

	c.JSON(http.StatusOK, map_spec_object_type_showidentifierentryAPI)
}

// UpdateMap_SPEC_OBJECT_TYPE_showIdentifierEntry
//
// swagger:route PATCH /map_spec_object_type_showidentifierentrys/{ID} map_spec_object_type_showidentifierentrys updateMap_SPEC_OBJECT_TYPE_showIdentifierEntry
//
// # Update a map_spec_object_type_showidentifierentry
//
// Responses:
// default: genericError
//
//	200: map_spec_object_type_showidentifierentryDBResponse
func (controller *Controller) UpdateMap_SPEC_OBJECT_TYPE_showIdentifierEntry(c *gin.Context) {

	mutexMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Lock()
	defer mutexMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Unlock()

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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.GetDB()

	// Validate input
	var input orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var map_spec_object_type_showidentifierentryDB orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB

	// fetch the map_spec_object_type_showidentifierentry
	_, err := db.First(&map_spec_object_type_showidentifierentryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	map_spec_object_type_showidentifierentryDB.CopyBasicFieldsFromMap_SPEC_OBJECT_TYPE_showIdentifierEntry_WOP(&input.Map_SPEC_OBJECT_TYPE_showIdentifierEntry_WOP)
	map_spec_object_type_showidentifierentryDB.Map_SPEC_OBJECT_TYPE_showIdentifierEntryPointersEncoding = input.Map_SPEC_OBJECT_TYPE_showIdentifierEntryPointersEncoding

	db, _ = db.Model(&map_spec_object_type_showidentifierentryDB)
	_, err = db.Updates(&map_spec_object_type_showidentifierentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	map_spec_object_type_showidentifierentryNew := new(models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry)
	map_spec_object_type_showidentifierentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_showIdentifierEntry(map_spec_object_type_showidentifierentryNew)

	// redeem pointers
	map_spec_object_type_showidentifierentryDB.DecodePointers(backRepo, map_spec_object_type_showidentifierentryNew)

	// get stage instance from DB instance, and call callback function
	map_spec_object_type_showidentifierentryOld := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Map_Map_SPEC_OBJECT_TYPE_showIdentifierEntryDBID_Map_SPEC_OBJECT_TYPE_showIdentifierEntryPtr[map_spec_object_type_showidentifierentryDB.ID]
	if map_spec_object_type_showidentifierentryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), map_spec_object_type_showidentifierentryOld, map_spec_object_type_showidentifierentryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the map_spec_object_type_showidentifierentryDB
	c.JSON(http.StatusOK, map_spec_object_type_showidentifierentryDB)
}

// DeleteMap_SPEC_OBJECT_TYPE_showIdentifierEntry
//
// swagger:route DELETE /map_spec_object_type_showidentifierentrys/{ID} map_spec_object_type_showidentifierentrys deleteMap_SPEC_OBJECT_TYPE_showIdentifierEntry
//
// # Delete a map_spec_object_type_showidentifierentry
//
// default: genericError
//
//	200: map_spec_object_type_showidentifierentryDBResponse
func (controller *Controller) DeleteMap_SPEC_OBJECT_TYPE_showIdentifierEntry(c *gin.Context) {

	mutexMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Lock()
	defer mutexMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMap_SPEC_OBJECT_TYPE_showIdentifierEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.GetDB()

	// Get model if exist
	var map_spec_object_type_showidentifierentryDB orm.Map_SPEC_OBJECT_TYPE_showIdentifierEntryDB
	if _, err := db.First(&map_spec_object_type_showidentifierentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&map_spec_object_type_showidentifierentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	map_spec_object_type_showidentifierentryDeleted := new(models.Map_SPEC_OBJECT_TYPE_showIdentifierEntry)
	map_spec_object_type_showidentifierentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_showIdentifierEntry(map_spec_object_type_showidentifierentryDeleted)

	// get stage instance from DB instance, and call callback function
	map_spec_object_type_showidentifierentryStaged := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showIdentifierEntry.Map_Map_SPEC_OBJECT_TYPE_showIdentifierEntryDBID_Map_SPEC_OBJECT_TYPE_showIdentifierEntryPtr[map_spec_object_type_showidentifierentryDB.ID]
	if map_spec_object_type_showidentifierentryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), map_spec_object_type_showidentifierentryStaged, map_spec_object_type_showidentifierentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
