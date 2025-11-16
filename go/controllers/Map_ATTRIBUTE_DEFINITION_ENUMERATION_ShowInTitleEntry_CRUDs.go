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
var __Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry__dummysDeclaration__ models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
var __Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry_time__dummyDeclaration time.Duration

var mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry sync.Mutex

// An Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry updateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry deleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
type Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry updateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
type Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryInput struct {
	// The Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry to submit or modify
	// in: body
	Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry *orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryAPI
}

// GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntrys
//
// swagger:route GET /map_attribute_definition_enumeration_showintitleentrys map_attribute_definition_enumeration_showintitleentrys getMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntrys
//
// # Get all map_attribute_definition_enumeration_showintitleentrys
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_enumeration_showintitleentryDBResponse
func (controller *Controller) GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntrys(c *gin.Context) {

	// source slice
	var map_attribute_definition_enumeration_showintitleentryDBs []orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.GetDB()

	_, err := db.Find(&map_attribute_definition_enumeration_showintitleentryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	map_attribute_definition_enumeration_showintitleentryAPIs := make([]orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryAPI, 0)

	// for each map_attribute_definition_enumeration_showintitleentry, update fields from the database nullable fields
	for idx := range map_attribute_definition_enumeration_showintitleentryDBs {
		map_attribute_definition_enumeration_showintitleentryDB := &map_attribute_definition_enumeration_showintitleentryDBs[idx]
		_ = map_attribute_definition_enumeration_showintitleentryDB
		var map_attribute_definition_enumeration_showintitleentryAPI orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryAPI

		// insertion point for updating fields
		map_attribute_definition_enumeration_showintitleentryAPI.ID = map_attribute_definition_enumeration_showintitleentryDB.ID
		map_attribute_definition_enumeration_showintitleentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry_WOP(&map_attribute_definition_enumeration_showintitleentryAPI.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry_WOP)
		map_attribute_definition_enumeration_showintitleentryAPI.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPointersEncoding = map_attribute_definition_enumeration_showintitleentryDB.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPointersEncoding
		map_attribute_definition_enumeration_showintitleentryAPIs = append(map_attribute_definition_enumeration_showintitleentryAPIs, map_attribute_definition_enumeration_showintitleentryAPI)
	}

	c.JSON(http.StatusOK, map_attribute_definition_enumeration_showintitleentryAPIs)
}

// PostMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
//
// swagger:route POST /map_attribute_definition_enumeration_showintitleentrys map_attribute_definition_enumeration_showintitleentrys postMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
//
// Creates a map_attribute_definition_enumeration_showintitleentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.GetDB()

	// Validate input
	var input orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create map_attribute_definition_enumeration_showintitleentry
	map_attribute_definition_enumeration_showintitleentryDB := orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB{}
	map_attribute_definition_enumeration_showintitleentryDB.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPointersEncoding = input.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPointersEncoding
	map_attribute_definition_enumeration_showintitleentryDB.CopyBasicFieldsFromMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry_WOP(&input.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry_WOP)

	_, err = db.Create(&map_attribute_definition_enumeration_showintitleentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.CheckoutPhaseOneInstance(&map_attribute_definition_enumeration_showintitleentryDB)
	map_attribute_definition_enumeration_showintitleentry := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDBID_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPtr[map_attribute_definition_enumeration_showintitleentryDB.ID]

	if map_attribute_definition_enumeration_showintitleentry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), map_attribute_definition_enumeration_showintitleentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, map_attribute_definition_enumeration_showintitleentryDB)
}

// GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
//
// swagger:route GET /map_attribute_definition_enumeration_showintitleentrys/{ID} map_attribute_definition_enumeration_showintitleentrys getMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
//
// Gets the details for a map_attribute_definition_enumeration_showintitleentry.
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_enumeration_showintitleentryDBResponse
func (controller *Controller) GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.GetDB()

	// Get map_attribute_definition_enumeration_showintitleentryDB in DB
	var map_attribute_definition_enumeration_showintitleentryDB orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB
	if _, err := db.First(&map_attribute_definition_enumeration_showintitleentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var map_attribute_definition_enumeration_showintitleentryAPI orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryAPI
	map_attribute_definition_enumeration_showintitleentryAPI.ID = map_attribute_definition_enumeration_showintitleentryDB.ID
	map_attribute_definition_enumeration_showintitleentryAPI.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPointersEncoding = map_attribute_definition_enumeration_showintitleentryDB.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPointersEncoding
	map_attribute_definition_enumeration_showintitleentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry_WOP(&map_attribute_definition_enumeration_showintitleentryAPI.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry_WOP)

	c.JSON(http.StatusOK, map_attribute_definition_enumeration_showintitleentryAPI)
}

// UpdateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
//
// swagger:route PATCH /map_attribute_definition_enumeration_showintitleentrys/{ID} map_attribute_definition_enumeration_showintitleentrys updateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
//
// # Update a map_attribute_definition_enumeration_showintitleentry
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_enumeration_showintitleentryDBResponse
func (controller *Controller) UpdateMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Unlock()

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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.GetDB()

	// Validate input
	var input orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var map_attribute_definition_enumeration_showintitleentryDB orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB

	// fetch the map_attribute_definition_enumeration_showintitleentry
	_, err := db.First(&map_attribute_definition_enumeration_showintitleentryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	map_attribute_definition_enumeration_showintitleentryDB.CopyBasicFieldsFromMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry_WOP(&input.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry_WOP)
	map_attribute_definition_enumeration_showintitleentryDB.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPointersEncoding = input.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPointersEncoding

	db, _ = db.Model(&map_attribute_definition_enumeration_showintitleentryDB)
	_, err = db.Updates(&map_attribute_definition_enumeration_showintitleentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	map_attribute_definition_enumeration_showintitleentryNew := new(models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry)
	map_attribute_definition_enumeration_showintitleentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry(map_attribute_definition_enumeration_showintitleentryNew)

	// redeem pointers
	map_attribute_definition_enumeration_showintitleentryDB.DecodePointers(backRepo, map_attribute_definition_enumeration_showintitleentryNew)

	// get stage instance from DB instance, and call callback function
	map_attribute_definition_enumeration_showintitleentryOld := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDBID_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPtr[map_attribute_definition_enumeration_showintitleentryDB.ID]
	if map_attribute_definition_enumeration_showintitleentryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), map_attribute_definition_enumeration_showintitleentryOld, map_attribute_definition_enumeration_showintitleentryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the map_attribute_definition_enumeration_showintitleentryDB
	c.JSON(http.StatusOK, map_attribute_definition_enumeration_showintitleentryDB)
}

// DeleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
//
// swagger:route DELETE /map_attribute_definition_enumeration_showintitleentrys/{ID} map_attribute_definition_enumeration_showintitleentrys deleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry
//
// # Delete a map_attribute_definition_enumeration_showintitleentry
//
// default: genericError
//
//	200: map_attribute_definition_enumeration_showintitleentryDBResponse
func (controller *Controller) DeleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.GetDB()

	// Get model if exist
	var map_attribute_definition_enumeration_showintitleentryDB orm.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDB
	if _, err := db.First(&map_attribute_definition_enumeration_showintitleentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&map_attribute_definition_enumeration_showintitleentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	map_attribute_definition_enumeration_showintitleentryDeleted := new(models.Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry)
	map_attribute_definition_enumeration_showintitleentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry(map_attribute_definition_enumeration_showintitleentryDeleted)

	// get stage instance from DB instance, and call callback function
	map_attribute_definition_enumeration_showintitleentryStaged := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntry.Map_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryDBID_Map_ATTRIBUTE_DEFINITION_ENUMERATION_ShowInTitleEntryPtr[map_attribute_definition_enumeration_showintitleentryDB.ID]
	if map_attribute_definition_enumeration_showintitleentryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), map_attribute_definition_enumeration_showintitleentryStaged, map_attribute_definition_enumeration_showintitleentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
