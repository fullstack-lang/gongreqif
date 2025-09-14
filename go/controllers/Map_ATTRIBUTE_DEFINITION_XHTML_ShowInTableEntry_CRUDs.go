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
var __Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry__dummysDeclaration__ models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
var __Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry_time__dummyDeclaration time.Duration

var mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry sync.Mutex

// An Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry updateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry deleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
type Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry updateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
type Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryInput struct {
	// The Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry to submit or modify
	// in: body
	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry *orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryAPI
}

// GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntrys
//
// swagger:route GET /map_attribute_definition_xhtml_showintableentrys map_attribute_definition_xhtml_showintableentrys getMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntrys
//
// # Get all map_attribute_definition_xhtml_showintableentrys
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_xhtml_showintableentryDBResponse
func (controller *Controller) GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntrys(c *gin.Context) {

	// source slice
	var map_attribute_definition_xhtml_showintableentryDBs []orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.GetDB()

	_, err := db.Find(&map_attribute_definition_xhtml_showintableentryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	map_attribute_definition_xhtml_showintableentryAPIs := make([]orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryAPI, 0)

	// for each map_attribute_definition_xhtml_showintableentry, update fields from the database nullable fields
	for idx := range map_attribute_definition_xhtml_showintableentryDBs {
		map_attribute_definition_xhtml_showintableentryDB := &map_attribute_definition_xhtml_showintableentryDBs[idx]
		_ = map_attribute_definition_xhtml_showintableentryDB
		var map_attribute_definition_xhtml_showintableentryAPI orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryAPI

		// insertion point for updating fields
		map_attribute_definition_xhtml_showintableentryAPI.ID = map_attribute_definition_xhtml_showintableentryDB.ID
		map_attribute_definition_xhtml_showintableentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry_WOP(&map_attribute_definition_xhtml_showintableentryAPI.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry_WOP)
		map_attribute_definition_xhtml_showintableentryAPI.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPointersEncoding = map_attribute_definition_xhtml_showintableentryDB.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPointersEncoding
		map_attribute_definition_xhtml_showintableentryAPIs = append(map_attribute_definition_xhtml_showintableentryAPIs, map_attribute_definition_xhtml_showintableentryAPI)
	}

	c.JSON(http.StatusOK, map_attribute_definition_xhtml_showintableentryAPIs)
}

// PostMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
//
// swagger:route POST /map_attribute_definition_xhtml_showintableentrys map_attribute_definition_xhtml_showintableentrys postMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
//
// Creates a map_attribute_definition_xhtml_showintableentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.GetDB()

	// Validate input
	var input orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create map_attribute_definition_xhtml_showintableentry
	map_attribute_definition_xhtml_showintableentryDB := orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB{}
	map_attribute_definition_xhtml_showintableentryDB.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPointersEncoding = input.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPointersEncoding
	map_attribute_definition_xhtml_showintableentryDB.CopyBasicFieldsFromMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry_WOP(&input.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry_WOP)

	_, err = db.Create(&map_attribute_definition_xhtml_showintableentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.CheckoutPhaseOneInstance(&map_attribute_definition_xhtml_showintableentryDB)
	map_attribute_definition_xhtml_showintableentry := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDBID_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPtr[map_attribute_definition_xhtml_showintableentryDB.ID]

	if map_attribute_definition_xhtml_showintableentry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), map_attribute_definition_xhtml_showintableentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, map_attribute_definition_xhtml_showintableentryDB)
}

// GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
//
// swagger:route GET /map_attribute_definition_xhtml_showintableentrys/{ID} map_attribute_definition_xhtml_showintableentrys getMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
//
// Gets the details for a map_attribute_definition_xhtml_showintableentry.
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_xhtml_showintableentryDBResponse
func (controller *Controller) GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.GetDB()

	// Get map_attribute_definition_xhtml_showintableentryDB in DB
	var map_attribute_definition_xhtml_showintableentryDB orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB
	if _, err := db.First(&map_attribute_definition_xhtml_showintableentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var map_attribute_definition_xhtml_showintableentryAPI orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryAPI
	map_attribute_definition_xhtml_showintableentryAPI.ID = map_attribute_definition_xhtml_showintableentryDB.ID
	map_attribute_definition_xhtml_showintableentryAPI.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPointersEncoding = map_attribute_definition_xhtml_showintableentryDB.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPointersEncoding
	map_attribute_definition_xhtml_showintableentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry_WOP(&map_attribute_definition_xhtml_showintableentryAPI.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry_WOP)

	c.JSON(http.StatusOK, map_attribute_definition_xhtml_showintableentryAPI)
}

// UpdateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
//
// swagger:route PATCH /map_attribute_definition_xhtml_showintableentrys/{ID} map_attribute_definition_xhtml_showintableentrys updateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
//
// # Update a map_attribute_definition_xhtml_showintableentry
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_xhtml_showintableentryDBResponse
func (controller *Controller) UpdateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Unlock()

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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.GetDB()

	// Validate input
	var input orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var map_attribute_definition_xhtml_showintableentryDB orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB

	// fetch the map_attribute_definition_xhtml_showintableentry
	_, err := db.First(&map_attribute_definition_xhtml_showintableentryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	map_attribute_definition_xhtml_showintableentryDB.CopyBasicFieldsFromMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry_WOP(&input.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry_WOP)
	map_attribute_definition_xhtml_showintableentryDB.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPointersEncoding = input.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPointersEncoding

	db, _ = db.Model(&map_attribute_definition_xhtml_showintableentryDB)
	_, err = db.Updates(&map_attribute_definition_xhtml_showintableentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	map_attribute_definition_xhtml_showintableentryNew := new(models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry)
	map_attribute_definition_xhtml_showintableentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry(map_attribute_definition_xhtml_showintableentryNew)

	// redeem pointers
	map_attribute_definition_xhtml_showintableentryDB.DecodePointers(backRepo, map_attribute_definition_xhtml_showintableentryNew)

	// get stage instance from DB instance, and call callback function
	map_attribute_definition_xhtml_showintableentryOld := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDBID_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPtr[map_attribute_definition_xhtml_showintableentryDB.ID]
	if map_attribute_definition_xhtml_showintableentryOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_attribute_definition_xhtml_showintableentryOld, map_attribute_definition_xhtml_showintableentryNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_attribute_definition_xhtml_showintableentryOld, map_attribute_definition_xhtml_showintableentryNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the map_attribute_definition_xhtml_showintableentryDB
	c.JSON(http.StatusOK, map_attribute_definition_xhtml_showintableentryDB)
}

// DeleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
//
// swagger:route DELETE /map_attribute_definition_xhtml_showintableentrys/{ID} map_attribute_definition_xhtml_showintableentrys deleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry
//
// # Delete a map_attribute_definition_xhtml_showintableentry
//
// default: genericError
//
//	200: map_attribute_definition_xhtml_showintableentryDBResponse
func (controller *Controller) DeleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.GetDB()

	// Get model if exist
	var map_attribute_definition_xhtml_showintableentryDB orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDB
	if _, err := db.First(&map_attribute_definition_xhtml_showintableentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&map_attribute_definition_xhtml_showintableentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	map_attribute_definition_xhtml_showintableentryDeleted := new(models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry)
	map_attribute_definition_xhtml_showintableentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry(map_attribute_definition_xhtml_showintableentryDeleted)

	// get stage instance from DB instance, and call callback function
	map_attribute_definition_xhtml_showintableentryStaged := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntry.Map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryDBID_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTableEntryPtr[map_attribute_definition_xhtml_showintableentryDB.ID]
	if map_attribute_definition_xhtml_showintableentryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), map_attribute_definition_xhtml_showintableentryStaged, map_attribute_definition_xhtml_showintableentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
