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
var __Map_SPECIFICATION_Nodes_expandedEntry__dummysDeclaration__ models.Map_SPECIFICATION_Nodes_expandedEntry
var __Map_SPECIFICATION_Nodes_expandedEntry_time__dummyDeclaration time.Duration

var mutexMap_SPECIFICATION_Nodes_expandedEntry sync.Mutex

// An Map_SPECIFICATION_Nodes_expandedEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMap_SPECIFICATION_Nodes_expandedEntry updateMap_SPECIFICATION_Nodes_expandedEntry deleteMap_SPECIFICATION_Nodes_expandedEntry
type Map_SPECIFICATION_Nodes_expandedEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Map_SPECIFICATION_Nodes_expandedEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMap_SPECIFICATION_Nodes_expandedEntry updateMap_SPECIFICATION_Nodes_expandedEntry
type Map_SPECIFICATION_Nodes_expandedEntryInput struct {
	// The Map_SPECIFICATION_Nodes_expandedEntry to submit or modify
	// in: body
	Map_SPECIFICATION_Nodes_expandedEntry *orm.Map_SPECIFICATION_Nodes_expandedEntryAPI
}

// GetMap_SPECIFICATION_Nodes_expandedEntrys
//
// swagger:route GET /map_specification_nodes_expandedentrys map_specification_nodes_expandedentrys getMap_SPECIFICATION_Nodes_expandedEntrys
//
// # Get all map_specification_nodes_expandedentrys
//
// Responses:
// default: genericError
//
//	200: map_specification_nodes_expandedentryDBResponse
func (controller *Controller) GetMap_SPECIFICATION_Nodes_expandedEntrys(c *gin.Context) {

	// source slice
	var map_specification_nodes_expandedentryDBs []orm.Map_SPECIFICATION_Nodes_expandedEntryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_SPECIFICATION_Nodes_expandedEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.GetDB()

	_, err := db.Find(&map_specification_nodes_expandedentryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	map_specification_nodes_expandedentryAPIs := make([]orm.Map_SPECIFICATION_Nodes_expandedEntryAPI, 0)

	// for each map_specification_nodes_expandedentry, update fields from the database nullable fields
	for idx := range map_specification_nodes_expandedentryDBs {
		map_specification_nodes_expandedentryDB := &map_specification_nodes_expandedentryDBs[idx]
		_ = map_specification_nodes_expandedentryDB
		var map_specification_nodes_expandedentryAPI orm.Map_SPECIFICATION_Nodes_expandedEntryAPI

		// insertion point for updating fields
		map_specification_nodes_expandedentryAPI.ID = map_specification_nodes_expandedentryDB.ID
		map_specification_nodes_expandedentryDB.CopyBasicFieldsToMap_SPECIFICATION_Nodes_expandedEntry_WOP(&map_specification_nodes_expandedentryAPI.Map_SPECIFICATION_Nodes_expandedEntry_WOP)
		map_specification_nodes_expandedentryAPI.Map_SPECIFICATION_Nodes_expandedEntryPointersEncoding = map_specification_nodes_expandedentryDB.Map_SPECIFICATION_Nodes_expandedEntryPointersEncoding
		map_specification_nodes_expandedentryAPIs = append(map_specification_nodes_expandedentryAPIs, map_specification_nodes_expandedentryAPI)
	}

	c.JSON(http.StatusOK, map_specification_nodes_expandedentryAPIs)
}

// PostMap_SPECIFICATION_Nodes_expandedEntry
//
// swagger:route POST /map_specification_nodes_expandedentrys map_specification_nodes_expandedentrys postMap_SPECIFICATION_Nodes_expandedEntry
//
// Creates a map_specification_nodes_expandedentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMap_SPECIFICATION_Nodes_expandedEntry(c *gin.Context) {

	mutexMap_SPECIFICATION_Nodes_expandedEntry.Lock()
	defer mutexMap_SPECIFICATION_Nodes_expandedEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMap_SPECIFICATION_Nodes_expandedEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.GetDB()

	// Validate input
	var input orm.Map_SPECIFICATION_Nodes_expandedEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create map_specification_nodes_expandedentry
	map_specification_nodes_expandedentryDB := orm.Map_SPECIFICATION_Nodes_expandedEntryDB{}
	map_specification_nodes_expandedentryDB.Map_SPECIFICATION_Nodes_expandedEntryPointersEncoding = input.Map_SPECIFICATION_Nodes_expandedEntryPointersEncoding
	map_specification_nodes_expandedentryDB.CopyBasicFieldsFromMap_SPECIFICATION_Nodes_expandedEntry_WOP(&input.Map_SPECIFICATION_Nodes_expandedEntry_WOP)

	_, err = db.Create(&map_specification_nodes_expandedentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.CheckoutPhaseOneInstance(&map_specification_nodes_expandedentryDB)
	map_specification_nodes_expandedentry := backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.Map_Map_SPECIFICATION_Nodes_expandedEntryDBID_Map_SPECIFICATION_Nodes_expandedEntryPtr[map_specification_nodes_expandedentryDB.ID]

	if map_specification_nodes_expandedentry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), map_specification_nodes_expandedentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, map_specification_nodes_expandedentryDB)
}

// GetMap_SPECIFICATION_Nodes_expandedEntry
//
// swagger:route GET /map_specification_nodes_expandedentrys/{ID} map_specification_nodes_expandedentrys getMap_SPECIFICATION_Nodes_expandedEntry
//
// Gets the details for a map_specification_nodes_expandedentry.
//
// Responses:
// default: genericError
//
//	200: map_specification_nodes_expandedentryDBResponse
func (controller *Controller) GetMap_SPECIFICATION_Nodes_expandedEntry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_SPECIFICATION_Nodes_expandedEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.GetDB()

	// Get map_specification_nodes_expandedentryDB in DB
	var map_specification_nodes_expandedentryDB orm.Map_SPECIFICATION_Nodes_expandedEntryDB
	if _, err := db.First(&map_specification_nodes_expandedentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var map_specification_nodes_expandedentryAPI orm.Map_SPECIFICATION_Nodes_expandedEntryAPI
	map_specification_nodes_expandedentryAPI.ID = map_specification_nodes_expandedentryDB.ID
	map_specification_nodes_expandedentryAPI.Map_SPECIFICATION_Nodes_expandedEntryPointersEncoding = map_specification_nodes_expandedentryDB.Map_SPECIFICATION_Nodes_expandedEntryPointersEncoding
	map_specification_nodes_expandedentryDB.CopyBasicFieldsToMap_SPECIFICATION_Nodes_expandedEntry_WOP(&map_specification_nodes_expandedentryAPI.Map_SPECIFICATION_Nodes_expandedEntry_WOP)

	c.JSON(http.StatusOK, map_specification_nodes_expandedentryAPI)
}

// UpdateMap_SPECIFICATION_Nodes_expandedEntry
//
// swagger:route PATCH /map_specification_nodes_expandedentrys/{ID} map_specification_nodes_expandedentrys updateMap_SPECIFICATION_Nodes_expandedEntry
//
// # Update a map_specification_nodes_expandedentry
//
// Responses:
// default: genericError
//
//	200: map_specification_nodes_expandedentryDBResponse
func (controller *Controller) UpdateMap_SPECIFICATION_Nodes_expandedEntry(c *gin.Context) {

	mutexMap_SPECIFICATION_Nodes_expandedEntry.Lock()
	defer mutexMap_SPECIFICATION_Nodes_expandedEntry.Unlock()

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
	db := backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.GetDB()

	// Validate input
	var input orm.Map_SPECIFICATION_Nodes_expandedEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var map_specification_nodes_expandedentryDB orm.Map_SPECIFICATION_Nodes_expandedEntryDB

	// fetch the map_specification_nodes_expandedentry
	_, err := db.First(&map_specification_nodes_expandedentryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	map_specification_nodes_expandedentryDB.CopyBasicFieldsFromMap_SPECIFICATION_Nodes_expandedEntry_WOP(&input.Map_SPECIFICATION_Nodes_expandedEntry_WOP)
	map_specification_nodes_expandedentryDB.Map_SPECIFICATION_Nodes_expandedEntryPointersEncoding = input.Map_SPECIFICATION_Nodes_expandedEntryPointersEncoding

	db, _ = db.Model(&map_specification_nodes_expandedentryDB)
	_, err = db.Updates(&map_specification_nodes_expandedentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	map_specification_nodes_expandedentryNew := new(models.Map_SPECIFICATION_Nodes_expandedEntry)
	map_specification_nodes_expandedentryDB.CopyBasicFieldsToMap_SPECIFICATION_Nodes_expandedEntry(map_specification_nodes_expandedentryNew)

	// redeem pointers
	map_specification_nodes_expandedentryDB.DecodePointers(backRepo, map_specification_nodes_expandedentryNew)

	// get stage instance from DB instance, and call callback function
	map_specification_nodes_expandedentryOld := backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.Map_Map_SPECIFICATION_Nodes_expandedEntryDBID_Map_SPECIFICATION_Nodes_expandedEntryPtr[map_specification_nodes_expandedentryDB.ID]
	if map_specification_nodes_expandedentryOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_specification_nodes_expandedentryOld, map_specification_nodes_expandedentryNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_specification_nodes_expandedentryOld, map_specification_nodes_expandedentryNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the map_specification_nodes_expandedentryDB
	c.JSON(http.StatusOK, map_specification_nodes_expandedentryDB)
}

// DeleteMap_SPECIFICATION_Nodes_expandedEntry
//
// swagger:route DELETE /map_specification_nodes_expandedentrys/{ID} map_specification_nodes_expandedentrys deleteMap_SPECIFICATION_Nodes_expandedEntry
//
// # Delete a map_specification_nodes_expandedentry
//
// default: genericError
//
//	200: map_specification_nodes_expandedentryDBResponse
func (controller *Controller) DeleteMap_SPECIFICATION_Nodes_expandedEntry(c *gin.Context) {

	mutexMap_SPECIFICATION_Nodes_expandedEntry.Lock()
	defer mutexMap_SPECIFICATION_Nodes_expandedEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMap_SPECIFICATION_Nodes_expandedEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.GetDB()

	// Get model if exist
	var map_specification_nodes_expandedentryDB orm.Map_SPECIFICATION_Nodes_expandedEntryDB
	if _, err := db.First(&map_specification_nodes_expandedentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&map_specification_nodes_expandedentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	map_specification_nodes_expandedentryDeleted := new(models.Map_SPECIFICATION_Nodes_expandedEntry)
	map_specification_nodes_expandedentryDB.CopyBasicFieldsToMap_SPECIFICATION_Nodes_expandedEntry(map_specification_nodes_expandedentryDeleted)

	// get stage instance from DB instance, and call callback function
	map_specification_nodes_expandedentryStaged := backRepo.BackRepoMap_SPECIFICATION_Nodes_expandedEntry.Map_Map_SPECIFICATION_Nodes_expandedEntryDBID_Map_SPECIFICATION_Nodes_expandedEntryPtr[map_specification_nodes_expandedentryDB.ID]
	if map_specification_nodes_expandedentryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), map_specification_nodes_expandedentryStaged, map_specification_nodes_expandedentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
