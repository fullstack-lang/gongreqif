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
var __Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry__dummysDeclaration__ models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry
var __Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry_time__dummyDeclaration time.Duration

var mutexMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry sync.Mutex

// An Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry updateMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry deleteMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
type Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry updateMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
type Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryInput struct {
	// The Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry to submit or modify
	// in: body
	Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry *orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryAPI
}

// GetMap_SPEC_OBJECT_TYPE_isNodeExpandedEntrys
//
// swagger:route GET /map_spec_object_type_isnodeexpandedentrys map_spec_object_type_isnodeexpandedentrys getMap_SPEC_OBJECT_TYPE_isNodeExpandedEntrys
//
// # Get all map_spec_object_type_isnodeexpandedentrys
//
// Responses:
// default: genericError
//
//	200: map_spec_object_type_isnodeexpandedentryDBResponse
func (controller *Controller) GetMap_SPEC_OBJECT_TYPE_isNodeExpandedEntrys(c *gin.Context) {

	// source slice
	var map_spec_object_type_isnodeexpandedentryDBs []orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_SPEC_OBJECT_TYPE_isNodeExpandedEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.GetDB()

	_, err := db.Find(&map_spec_object_type_isnodeexpandedentryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	map_spec_object_type_isnodeexpandedentryAPIs := make([]orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryAPI, 0)

	// for each map_spec_object_type_isnodeexpandedentry, update fields from the database nullable fields
	for idx := range map_spec_object_type_isnodeexpandedentryDBs {
		map_spec_object_type_isnodeexpandedentryDB := &map_spec_object_type_isnodeexpandedentryDBs[idx]
		_ = map_spec_object_type_isnodeexpandedentryDB
		var map_spec_object_type_isnodeexpandedentryAPI orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryAPI

		// insertion point for updating fields
		map_spec_object_type_isnodeexpandedentryAPI.ID = map_spec_object_type_isnodeexpandedentryDB.ID
		map_spec_object_type_isnodeexpandedentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry_WOP(&map_spec_object_type_isnodeexpandedentryAPI.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry_WOP)
		map_spec_object_type_isnodeexpandedentryAPI.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPointersEncoding = map_spec_object_type_isnodeexpandedentryDB.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPointersEncoding
		map_spec_object_type_isnodeexpandedentryAPIs = append(map_spec_object_type_isnodeexpandedentryAPIs, map_spec_object_type_isnodeexpandedentryAPI)
	}

	c.JSON(http.StatusOK, map_spec_object_type_isnodeexpandedentryAPIs)
}

// PostMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
//
// swagger:route POST /map_spec_object_type_isnodeexpandedentrys map_spec_object_type_isnodeexpandedentrys postMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
//
// Creates a map_spec_object_type_isnodeexpandedentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry(c *gin.Context) {

	mutexMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Lock()
	defer mutexMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMap_SPEC_OBJECT_TYPE_isNodeExpandedEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.GetDB()

	// Validate input
	var input orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create map_spec_object_type_isnodeexpandedentry
	map_spec_object_type_isnodeexpandedentryDB := orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB{}
	map_spec_object_type_isnodeexpandedentryDB.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPointersEncoding = input.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPointersEncoding
	map_spec_object_type_isnodeexpandedentryDB.CopyBasicFieldsFromMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry_WOP(&input.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry_WOP)

	_, err = db.Create(&map_spec_object_type_isnodeexpandedentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.CheckoutPhaseOneInstance(&map_spec_object_type_isnodeexpandedentryDB)
	map_spec_object_type_isnodeexpandedentry := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Map_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDBID_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPtr[map_spec_object_type_isnodeexpandedentryDB.ID]

	if map_spec_object_type_isnodeexpandedentry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), map_spec_object_type_isnodeexpandedentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, map_spec_object_type_isnodeexpandedentryDB)
}

// GetMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
//
// swagger:route GET /map_spec_object_type_isnodeexpandedentrys/{ID} map_spec_object_type_isnodeexpandedentrys getMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
//
// Gets the details for a map_spec_object_type_isnodeexpandedentry.
//
// Responses:
// default: genericError
//
//	200: map_spec_object_type_isnodeexpandedentryDBResponse
func (controller *Controller) GetMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.GetDB()

	// Get map_spec_object_type_isnodeexpandedentryDB in DB
	var map_spec_object_type_isnodeexpandedentryDB orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB
	if _, err := db.First(&map_spec_object_type_isnodeexpandedentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var map_spec_object_type_isnodeexpandedentryAPI orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryAPI
	map_spec_object_type_isnodeexpandedentryAPI.ID = map_spec_object_type_isnodeexpandedentryDB.ID
	map_spec_object_type_isnodeexpandedentryAPI.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPointersEncoding = map_spec_object_type_isnodeexpandedentryDB.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPointersEncoding
	map_spec_object_type_isnodeexpandedentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry_WOP(&map_spec_object_type_isnodeexpandedentryAPI.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry_WOP)

	c.JSON(http.StatusOK, map_spec_object_type_isnodeexpandedentryAPI)
}

// UpdateMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
//
// swagger:route PATCH /map_spec_object_type_isnodeexpandedentrys/{ID} map_spec_object_type_isnodeexpandedentrys updateMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
//
// # Update a map_spec_object_type_isnodeexpandedentry
//
// Responses:
// default: genericError
//
//	200: map_spec_object_type_isnodeexpandedentryDBResponse
func (controller *Controller) UpdateMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry(c *gin.Context) {

	mutexMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Lock()
	defer mutexMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Unlock()

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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.GetDB()

	// Validate input
	var input orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var map_spec_object_type_isnodeexpandedentryDB orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB

	// fetch the map_spec_object_type_isnodeexpandedentry
	_, err := db.First(&map_spec_object_type_isnodeexpandedentryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	map_spec_object_type_isnodeexpandedentryDB.CopyBasicFieldsFromMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry_WOP(&input.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry_WOP)
	map_spec_object_type_isnodeexpandedentryDB.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPointersEncoding = input.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPointersEncoding

	db, _ = db.Model(&map_spec_object_type_isnodeexpandedentryDB)
	_, err = db.Updates(&map_spec_object_type_isnodeexpandedentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	map_spec_object_type_isnodeexpandedentryNew := new(models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry)
	map_spec_object_type_isnodeexpandedentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry(map_spec_object_type_isnodeexpandedentryNew)

	// redeem pointers
	map_spec_object_type_isnodeexpandedentryDB.DecodePointers(backRepo, map_spec_object_type_isnodeexpandedentryNew)

	// get stage instance from DB instance, and call callback function
	map_spec_object_type_isnodeexpandedentryOld := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Map_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDBID_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPtr[map_spec_object_type_isnodeexpandedentryDB.ID]
	if map_spec_object_type_isnodeexpandedentryOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_spec_object_type_isnodeexpandedentryOld, map_spec_object_type_isnodeexpandedentryNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_spec_object_type_isnodeexpandedentryOld, map_spec_object_type_isnodeexpandedentryNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the map_spec_object_type_isnodeexpandedentryDB
	c.JSON(http.StatusOK, map_spec_object_type_isnodeexpandedentryDB)
}

// DeleteMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
//
// swagger:route DELETE /map_spec_object_type_isnodeexpandedentrys/{ID} map_spec_object_type_isnodeexpandedentrys deleteMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry
//
// # Delete a map_spec_object_type_isnodeexpandedentry
//
// default: genericError
//
//	200: map_spec_object_type_isnodeexpandedentryDBResponse
func (controller *Controller) DeleteMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry(c *gin.Context) {

	mutexMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Lock()
	defer mutexMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.GetDB()

	// Get model if exist
	var map_spec_object_type_isnodeexpandedentryDB orm.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDB
	if _, err := db.First(&map_spec_object_type_isnodeexpandedentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&map_spec_object_type_isnodeexpandedentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	map_spec_object_type_isnodeexpandedentryDeleted := new(models.Map_SPEC_OBJECT_TYPE_isNodeExpandedEntry)
	map_spec_object_type_isnodeexpandedentryDB.CopyBasicFieldsToMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry(map_spec_object_type_isnodeexpandedentryDeleted)

	// get stage instance from DB instance, and call callback function
	map_spec_object_type_isnodeexpandedentryStaged := backRepo.BackRepoMap_SPEC_OBJECT_TYPE_isNodeExpandedEntry.Map_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryDBID_Map_SPEC_OBJECT_TYPE_isNodeExpandedEntryPtr[map_spec_object_type_isnodeexpandedentryDB.ID]
	if map_spec_object_type_isnodeexpandedentryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), map_spec_object_type_isnodeexpandedentryStaged, map_spec_object_type_isnodeexpandedentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
