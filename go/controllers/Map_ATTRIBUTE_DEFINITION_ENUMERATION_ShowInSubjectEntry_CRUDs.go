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
var __Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry__dummysDeclaration__ models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
var __Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry_time__dummyDeclaration time.Duration

var mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry sync.Mutex

// An Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry updateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry deleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
type Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry updateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
type Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryInput struct {
	// The Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry to submit or modify
	// in: body
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry *orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryAPI
}

// GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntrys
//
// swagger:route GET /map_attribute_definition_enumeration_showinsubjectentrys map_attribute_definition_enumeration_showinsubjectentrys getMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntrys
//
// # Get all map_attribute_definition_enumeration_showinsubjectentrys
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_enumeration_showinsubjectentryDBResponse
func (controller *Controller) GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntrys(c *gin.Context) {

	// source slice
	var map_attribute_definition_enumeration_showinsubjectentryDBs []orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.GetDB()

	_, err := db.Find(&map_attribute_definition_enumeration_showinsubjectentryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	map_attribute_definition_enumeration_showinsubjectentryAPIs := make([]orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryAPI, 0)

	// for each map_attribute_definition_enumeration_showinsubjectentry, update fields from the database nullable fields
	for idx := range map_attribute_definition_enumeration_showinsubjectentryDBs {
		map_attribute_definition_enumeration_showinsubjectentryDB := &map_attribute_definition_enumeration_showinsubjectentryDBs[idx]
		_ = map_attribute_definition_enumeration_showinsubjectentryDB
		var map_attribute_definition_enumeration_showinsubjectentryAPI orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryAPI

		// insertion point for updating fields
		map_attribute_definition_enumeration_showinsubjectentryAPI.ID = map_attribute_definition_enumeration_showinsubjectentryDB.ID
		map_attribute_definition_enumeration_showinsubjectentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry_WOP(&map_attribute_definition_enumeration_showinsubjectentryAPI.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry_WOP)
		map_attribute_definition_enumeration_showinsubjectentryAPI.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPointersEncoding = map_attribute_definition_enumeration_showinsubjectentryDB.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPointersEncoding
		map_attribute_definition_enumeration_showinsubjectentryAPIs = append(map_attribute_definition_enumeration_showinsubjectentryAPIs, map_attribute_definition_enumeration_showinsubjectentryAPI)
	}

	c.JSON(http.StatusOK, map_attribute_definition_enumeration_showinsubjectentryAPIs)
}

// PostMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
//
// swagger:route POST /map_attribute_definition_enumeration_showinsubjectentrys map_attribute_definition_enumeration_showinsubjectentrys postMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
//
// Creates a map_attribute_definition_enumeration_showinsubjectentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.GetDB()

	// Validate input
	var input orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create map_attribute_definition_enumeration_showinsubjectentry
	map_attribute_definition_enumeration_showinsubjectentryDB := orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB{}
	map_attribute_definition_enumeration_showinsubjectentryDB.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPointersEncoding = input.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPointersEncoding
	map_attribute_definition_enumeration_showinsubjectentryDB.CopyBasicFieldsFromMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry_WOP(&input.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry_WOP)

	_, err = db.Create(&map_attribute_definition_enumeration_showinsubjectentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.CheckoutPhaseOneInstance(&map_attribute_definition_enumeration_showinsubjectentryDB)
	map_attribute_definition_enumeration_showinsubjectentry := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDBID_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPtr[map_attribute_definition_enumeration_showinsubjectentryDB.ID]

	if map_attribute_definition_enumeration_showinsubjectentry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), map_attribute_definition_enumeration_showinsubjectentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, map_attribute_definition_enumeration_showinsubjectentryDB)
}

// GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
//
// swagger:route GET /map_attribute_definition_enumeration_showinsubjectentrys/{ID} map_attribute_definition_enumeration_showinsubjectentrys getMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
//
// Gets the details for a map_attribute_definition_enumeration_showinsubjectentry.
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_enumeration_showinsubjectentryDBResponse
func (controller *Controller) GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.GetDB()

	// Get map_attribute_definition_enumeration_showinsubjectentryDB in DB
	var map_attribute_definition_enumeration_showinsubjectentryDB orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB
	if _, err := db.First(&map_attribute_definition_enumeration_showinsubjectentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var map_attribute_definition_enumeration_showinsubjectentryAPI orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryAPI
	map_attribute_definition_enumeration_showinsubjectentryAPI.ID = map_attribute_definition_enumeration_showinsubjectentryDB.ID
	map_attribute_definition_enumeration_showinsubjectentryAPI.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPointersEncoding = map_attribute_definition_enumeration_showinsubjectentryDB.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPointersEncoding
	map_attribute_definition_enumeration_showinsubjectentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry_WOP(&map_attribute_definition_enumeration_showinsubjectentryAPI.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry_WOP)

	c.JSON(http.StatusOK, map_attribute_definition_enumeration_showinsubjectentryAPI)
}

// UpdateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
//
// swagger:route PATCH /map_attribute_definition_enumeration_showinsubjectentrys/{ID} map_attribute_definition_enumeration_showinsubjectentrys updateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
//
// # Update a map_attribute_definition_enumeration_showinsubjectentry
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_enumeration_showinsubjectentryDBResponse
func (controller *Controller) UpdateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Unlock()

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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.GetDB()

	// Validate input
	var input orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var map_attribute_definition_enumeration_showinsubjectentryDB orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB

	// fetch the map_attribute_definition_enumeration_showinsubjectentry
	_, err := db.First(&map_attribute_definition_enumeration_showinsubjectentryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	map_attribute_definition_enumeration_showinsubjectentryDB.CopyBasicFieldsFromMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry_WOP(&input.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry_WOP)
	map_attribute_definition_enumeration_showinsubjectentryDB.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPointersEncoding = input.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPointersEncoding

	db, _ = db.Model(&map_attribute_definition_enumeration_showinsubjectentryDB)
	_, err = db.Updates(&map_attribute_definition_enumeration_showinsubjectentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	map_attribute_definition_enumeration_showinsubjectentryNew := new(models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry)
	map_attribute_definition_enumeration_showinsubjectentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry(map_attribute_definition_enumeration_showinsubjectentryNew)

	// redeem pointers
	map_attribute_definition_enumeration_showinsubjectentryDB.DecodePointers(backRepo, map_attribute_definition_enumeration_showinsubjectentryNew)

	// get stage instance from DB instance, and call callback function
	map_attribute_definition_enumeration_showinsubjectentryOld := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDBID_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPtr[map_attribute_definition_enumeration_showinsubjectentryDB.ID]
	if map_attribute_definition_enumeration_showinsubjectentryOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_attribute_definition_enumeration_showinsubjectentryOld, map_attribute_definition_enumeration_showinsubjectentryNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), map_attribute_definition_enumeration_showinsubjectentryOld, map_attribute_definition_enumeration_showinsubjectentryNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the map_attribute_definition_enumeration_showinsubjectentryDB
	c.JSON(http.StatusOK, map_attribute_definition_enumeration_showinsubjectentryDB)
}

// DeleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
//
// swagger:route DELETE /map_attribute_definition_enumeration_showinsubjectentrys/{ID} map_attribute_definition_enumeration_showinsubjectentrys deleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry
//
// # Delete a map_attribute_definition_enumeration_showinsubjectentry
//
// default: genericError
//
//	200: map_attribute_definition_enumeration_showinsubjectentryDBResponse
func (controller *Controller) DeleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.GetDB()

	// Get model if exist
	var map_attribute_definition_enumeration_showinsubjectentryDB orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDB
	if _, err := db.First(&map_attribute_definition_enumeration_showinsubjectentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&map_attribute_definition_enumeration_showinsubjectentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	map_attribute_definition_enumeration_showinsubjectentryDeleted := new(models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry)
	map_attribute_definition_enumeration_showinsubjectentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry(map_attribute_definition_enumeration_showinsubjectentryDeleted)

	// get stage instance from DB instance, and call callback function
	map_attribute_definition_enumeration_showinsubjectentryStaged := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntry.Map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryDBID_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInSubjectEntryPtr[map_attribute_definition_enumeration_showinsubjectentryDB.ID]
	if map_attribute_definition_enumeration_showinsubjectentryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), map_attribute_definition_enumeration_showinsubjectentryStaged, map_attribute_definition_enumeration_showinsubjectentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
