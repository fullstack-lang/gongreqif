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
var __Map_SPEC_OBJECT_TYPE_showNameEntry__dummysDeclaration__ models.Map_SPEC_OBJECT_TYPE_showNameEntry
var __Map_SPEC_OBJECT_TYPE_showNameEntry_time__dummyDeclaration time.Duration

var mutexMap_SPEC_OBJECT_TYPE_showNameEntry sync.Mutex

// An Map_SPEC_OBJECT_TYPE_showNameEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMap_SPEC_OBJECT_TYPE_showNameEntry updateMap_SPEC_OBJECT_TYPE_showNameEntry deleteMap_SPEC_OBJECT_TYPE_showNameEntry
type Map_SPEC_OBJECT_TYPE_showNameEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Map_SPEC_OBJECT_TYPE_showNameEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMap_SPEC_OBJECT_TYPE_showNameEntry updateMap_SPEC_OBJECT_TYPE_showNameEntry
type Map_SPEC_OBJECT_TYPE_showNameEntryInput struct {
	// The Map_SPEC_OBJECT_TYPE_showNameEntry to submit or modify
	// in: body
	Map_SPEC_OBJECT_TYPE_showNameEntry *orm.Map_SPEC_OBJECT_TYPE_showNameEntryAPI
}

// GetMap_SPEC_OBJECT_TYPE_showNameEntrys
//
// swagger:route GET /map_spec_object_type_shownameentrys map_spec_object_type_shownameentrys getMap_SPEC_OBJECT_TYPE_showNameEntrys
//
// # Get all map_spec_object_type_shownameentrys
//
// Responses:
// default: genericError
//
//	200: map_spec_object_type_shownameentryDBResponse
func (controller *Controller) GetMap_SPEC_OBJECT_TYPE_showNameEntrys(c *gin.Context) {

	// source slice
	var map_spec_object_type_shownameentryDBs []orm.Map_SPEC_OBJECT_TYPE_showNameEntryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_SPEC_OBJECT_TYPE_showNameEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.GetDB()

	_, err := db.Find(&map_spec_object_type_shownameentryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	map_spec_object_type_shownameentryAPIs := make([]orm.Map_SPEC_OBJECT_TYPE_showNameEntryAPI, 0)

	// for each map_spec_object_type_shownameentry, update fields from the database nullable fields
	for idx := range map_spec_object_type_shownameentryDBs {
		map_spec_object_type_shownameentryDB := &map_spec_object_type_shownameentryDBs[idx]
		_ = map_spec_object_type_shownameentryDB
		var map_spec_object_type_shownameentryAPI orm.Map_SPEC_OBJECT_TYPE_showNameEntryAPI

		// insertion point for updating fields
		map_spec_object_type_shownameentryAPI.ID = map_spec_object_type_shownameentryDB.ID
		map_spec_object_type_shownameentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_showNameEntry_WOP(&map_spec_object_type_shownameentryAPI.Map_SPEC_OBJECT_TYPE_showNameEntry_WOP)
		map_spec_object_type_shownameentryAPI.Map_SPEC_OBJECT_TYPE_showNameEntryPointersEncoding = map_spec_object_type_shownameentryDB.Map_SPEC_OBJECT_TYPE_showNameEntryPointersEncoding
		map_spec_object_type_shownameentryAPIs = append(map_spec_object_type_shownameentryAPIs, map_spec_object_type_shownameentryAPI)
	}

	c.JSON(http.StatusOK, map_spec_object_type_shownameentryAPIs)
}

// PostMap_SPEC_OBJECT_TYPE_showNameEntry
//
// swagger:route POST /map_spec_object_type_shownameentrys map_spec_object_type_shownameentrys postMap_SPEC_OBJECT_TYPE_showNameEntry
//
// Creates a map_spec_object_type_shownameentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMap_SPEC_OBJECT_TYPE_showNameEntry(c *gin.Context) {

	mutexMap_SPEC_OBJECT_TYPE_showNameEntry.Lock()
	defer mutexMap_SPEC_OBJECT_TYPE_showNameEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMap_SPEC_OBJECT_TYPE_showNameEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.GetDB()

	// Validate input
	var input orm.Map_SPEC_OBJECT_TYPE_showNameEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create map_spec_object_type_shownameentry
	map_spec_object_type_shownameentryDB := orm.Map_SPEC_OBJECT_TYPE_showNameEntryDB{}
	map_spec_object_type_shownameentryDB.Map_SPEC_OBJECT_TYPE_showNameEntryPointersEncoding = input.Map_SPEC_OBJECT_TYPE_showNameEntryPointersEncoding
	map_spec_object_type_shownameentryDB.CopyBasicFieldsFromMap_SPEC_OBJECT_TYPE_showNameEntry_WOP(&input.Map_SPEC_OBJECT_TYPE_showNameEntry_WOP)

	_, err = db.Create(&map_spec_object_type_shownameentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.CheckoutPhaseOneInstance(&map_spec_object_type_shownameentryDB)
	map_spec_object_type_shownameentry := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.Map_Map_SPEC_OBJECT_TYPE_showNameEntryDBID_Map_SPEC_OBJECT_TYPE_showNameEntryPtr[map_spec_object_type_shownameentryDB.ID]

	if map_spec_object_type_shownameentry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), map_spec_object_type_shownameentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, map_spec_object_type_shownameentryDB)
}

// GetMap_SPEC_OBJECT_TYPE_showNameEntry
//
// swagger:route GET /map_spec_object_type_shownameentrys/{ID} map_spec_object_type_shownameentrys getMap_SPEC_OBJECT_TYPE_showNameEntry
//
// Gets the details for a map_spec_object_type_shownameentry.
//
// Responses:
// default: genericError
//
//	200: map_spec_object_type_shownameentryDBResponse
func (controller *Controller) GetMap_SPEC_OBJECT_TYPE_showNameEntry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_SPEC_OBJECT_TYPE_showNameEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.GetDB()

	// Get map_spec_object_type_shownameentryDB in DB
	var map_spec_object_type_shownameentryDB orm.Map_SPEC_OBJECT_TYPE_showNameEntryDB
	if _, err := db.First(&map_spec_object_type_shownameentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var map_spec_object_type_shownameentryAPI orm.Map_SPEC_OBJECT_TYPE_showNameEntryAPI
	map_spec_object_type_shownameentryAPI.ID = map_spec_object_type_shownameentryDB.ID
	map_spec_object_type_shownameentryAPI.Map_SPEC_OBJECT_TYPE_showNameEntryPointersEncoding = map_spec_object_type_shownameentryDB.Map_SPEC_OBJECT_TYPE_showNameEntryPointersEncoding
	map_spec_object_type_shownameentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_showNameEntry_WOP(&map_spec_object_type_shownameentryAPI.Map_SPEC_OBJECT_TYPE_showNameEntry_WOP)

	c.JSON(http.StatusOK, map_spec_object_type_shownameentryAPI)
}

// UpdateMap_SPEC_OBJECT_TYPE_showNameEntry
//
// swagger:route PATCH /map_spec_object_type_shownameentrys/{ID} map_spec_object_type_shownameentrys updateMap_SPEC_OBJECT_TYPE_showNameEntry
//
// # Update a map_spec_object_type_shownameentry
//
// Responses:
// default: genericError
//
//	200: map_spec_object_type_shownameentryDBResponse
func (controller *Controller) UpdateMap_SPEC_OBJECT_TYPE_showNameEntry(c *gin.Context) {

	mutexMap_SPEC_OBJECT_TYPE_showNameEntry.Lock()
	defer mutexMap_SPEC_OBJECT_TYPE_showNameEntry.Unlock()

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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.GetDB()

	// Validate input
	var input orm.Map_SPEC_OBJECT_TYPE_showNameEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var map_spec_object_type_shownameentryDB orm.Map_SPEC_OBJECT_TYPE_showNameEntryDB

	// fetch the map_spec_object_type_shownameentry
	_, err := db.First(&map_spec_object_type_shownameentryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	map_spec_object_type_shownameentryDB.CopyBasicFieldsFromMap_SPEC_OBJECT_TYPE_showNameEntry_WOP(&input.Map_SPEC_OBJECT_TYPE_showNameEntry_WOP)
	map_spec_object_type_shownameentryDB.Map_SPEC_OBJECT_TYPE_showNameEntryPointersEncoding = input.Map_SPEC_OBJECT_TYPE_showNameEntryPointersEncoding

	db, _ = db.Model(&map_spec_object_type_shownameentryDB)
	_, err = db.Updates(&map_spec_object_type_shownameentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	map_spec_object_type_shownameentryNew := new(models.Map_SPEC_OBJECT_TYPE_showNameEntry)
	map_spec_object_type_shownameentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_showNameEntry(map_spec_object_type_shownameentryNew)

	// redeem pointers
	map_spec_object_type_shownameentryDB.DecodePointers(backRepo, map_spec_object_type_shownameentryNew)

	// get stage instance from DB instance, and call callback function
	map_spec_object_type_shownameentryOld := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.Map_Map_SPEC_OBJECT_TYPE_showNameEntryDBID_Map_SPEC_OBJECT_TYPE_showNameEntryPtr[map_spec_object_type_shownameentryDB.ID]
	if map_spec_object_type_shownameentryOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_spec_object_type_shownameentryOld, map_spec_object_type_shownameentryNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_spec_object_type_shownameentryOld, map_spec_object_type_shownameentryNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the map_spec_object_type_shownameentryDB
	c.JSON(http.StatusOK, map_spec_object_type_shownameentryDB)
}

// DeleteMap_SPEC_OBJECT_TYPE_showNameEntry
//
// swagger:route DELETE /map_spec_object_type_shownameentrys/{ID} map_spec_object_type_shownameentrys deleteMap_SPEC_OBJECT_TYPE_showNameEntry
//
// # Delete a map_spec_object_type_shownameentry
//
// default: genericError
//
//	200: map_spec_object_type_shownameentryDBResponse
func (controller *Controller) DeleteMap_SPEC_OBJECT_TYPE_showNameEntry(c *gin.Context) {

	mutexMap_SPEC_OBJECT_TYPE_showNameEntry.Lock()
	defer mutexMap_SPEC_OBJECT_TYPE_showNameEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMap_SPEC_OBJECT_TYPE_showNameEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.GetDB()

	// Get model if exist
	var map_spec_object_type_shownameentryDB orm.Map_SPEC_OBJECT_TYPE_showNameEntryDB
	if _, err := db.First(&map_spec_object_type_shownameentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&map_spec_object_type_shownameentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	map_spec_object_type_shownameentryDeleted := new(models.Map_SPEC_OBJECT_TYPE_showNameEntry)
	map_spec_object_type_shownameentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_showNameEntry(map_spec_object_type_shownameentryDeleted)

	// get stage instance from DB instance, and call callback function
	map_spec_object_type_shownameentryStaged := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_showNameEntry.Map_Map_SPEC_OBJECT_TYPE_showNameEntryDBID_Map_SPEC_OBJECT_TYPE_showNameEntryPtr[map_spec_object_type_shownameentryDB.ID]
	if map_spec_object_type_shownameentryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), map_spec_object_type_shownameentryStaged, map_spec_object_type_shownameentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
