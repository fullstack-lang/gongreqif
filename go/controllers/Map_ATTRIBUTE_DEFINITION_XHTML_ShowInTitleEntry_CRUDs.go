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
var __Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry__dummysDeclaration__ models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
var __Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry_time__dummyDeclaration time.Duration

var mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry sync.Mutex

// An Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry updateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry deleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
type Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry updateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
type Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryInput struct {
	// The Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry to submit or modify
	// in: body
	Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry *orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryAPI
}

// GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntrys
//
// swagger:route GET /map_attribute_definition_xhtml_showintitleentrys map_attribute_definition_xhtml_showintitleentrys getMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntrys
//
// # Get all map_attribute_definition_xhtml_showintitleentrys
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_xhtml_showintitleentryDBResponse
func (controller *Controller) GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntrys(c *gin.Context) {

	// source slice
	var map_attribute_definition_xhtml_showintitleentryDBs []orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.GetDB()

	_, err := db.Find(&map_attribute_definition_xhtml_showintitleentryDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	map_attribute_definition_xhtml_showintitleentryAPIs := make([]orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryAPI, 0)

	// for each map_attribute_definition_xhtml_showintitleentry, update fields from the database nullable fields
	for idx := range map_attribute_definition_xhtml_showintitleentryDBs {
		map_attribute_definition_xhtml_showintitleentryDB := &map_attribute_definition_xhtml_showintitleentryDBs[idx]
		_ = map_attribute_definition_xhtml_showintitleentryDB
		var map_attribute_definition_xhtml_showintitleentryAPI orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryAPI

		// insertion point for updating fields
		map_attribute_definition_xhtml_showintitleentryAPI.ID = map_attribute_definition_xhtml_showintitleentryDB.ID
		map_attribute_definition_xhtml_showintitleentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry_WOP(&map_attribute_definition_xhtml_showintitleentryAPI.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry_WOP)
		map_attribute_definition_xhtml_showintitleentryAPI.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPointersEncoding = map_attribute_definition_xhtml_showintitleentryDB.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPointersEncoding
		map_attribute_definition_xhtml_showintitleentryAPIs = append(map_attribute_definition_xhtml_showintitleentryAPIs, map_attribute_definition_xhtml_showintitleentryAPI)
	}

	c.JSON(http.StatusOK, map_attribute_definition_xhtml_showintitleentryAPIs)
}

// PostMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
//
// swagger:route POST /map_attribute_definition_xhtml_showintitleentrys map_attribute_definition_xhtml_showintitleentrys postMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
//
// Creates a map_attribute_definition_xhtml_showintitleentry
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntrys", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.GetDB()

	// Validate input
	var input orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create map_attribute_definition_xhtml_showintitleentry
	map_attribute_definition_xhtml_showintitleentryDB := orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB{}
	map_attribute_definition_xhtml_showintitleentryDB.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPointersEncoding = input.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPointersEncoding
	map_attribute_definition_xhtml_showintitleentryDB.CopyBasicFieldsFromMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry_WOP(&input.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry_WOP)

	_, err = db.Create(&map_attribute_definition_xhtml_showintitleentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.CheckoutPhaseOneInstance(&map_attribute_definition_xhtml_showintitleentryDB)
	map_attribute_definition_xhtml_showintitleentry := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDBID_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPtr[map_attribute_definition_xhtml_showintitleentryDB.ID]

	if map_attribute_definition_xhtml_showintitleentry != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), map_attribute_definition_xhtml_showintitleentry)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, map_attribute_definition_xhtml_showintitleentryDB)
}

// GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
//
// swagger:route GET /map_attribute_definition_xhtml_showintitleentrys/{ID} map_attribute_definition_xhtml_showintitleentrys getMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
//
// Gets the details for a map_attribute_definition_xhtml_showintitleentry.
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_xhtml_showintitleentryDBResponse
func (controller *Controller) GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.GetDB()

	// Get map_attribute_definition_xhtml_showintitleentryDB in DB
	var map_attribute_definition_xhtml_showintitleentryDB orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB
	if _, err := db.First(&map_attribute_definition_xhtml_showintitleentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var map_attribute_definition_xhtml_showintitleentryAPI orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryAPI
	map_attribute_definition_xhtml_showintitleentryAPI.ID = map_attribute_definition_xhtml_showintitleentryDB.ID
	map_attribute_definition_xhtml_showintitleentryAPI.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPointersEncoding = map_attribute_definition_xhtml_showintitleentryDB.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPointersEncoding
	map_attribute_definition_xhtml_showintitleentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry_WOP(&map_attribute_definition_xhtml_showintitleentryAPI.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry_WOP)

	c.JSON(http.StatusOK, map_attribute_definition_xhtml_showintitleentryAPI)
}

// UpdateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
//
// swagger:route PATCH /map_attribute_definition_xhtml_showintitleentrys/{ID} map_attribute_definition_xhtml_showintitleentrys updateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
//
// # Update a map_attribute_definition_xhtml_showintitleentry
//
// Responses:
// default: genericError
//
//	200: map_attribute_definition_xhtml_showintitleentryDBResponse
func (controller *Controller) UpdateMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Unlock()

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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.GetDB()

	// Validate input
	var input orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var map_attribute_definition_xhtml_showintitleentryDB orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB

	// fetch the map_attribute_definition_xhtml_showintitleentry
	_, err := db.First(&map_attribute_definition_xhtml_showintitleentryDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	map_attribute_definition_xhtml_showintitleentryDB.CopyBasicFieldsFromMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry_WOP(&input.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry_WOP)
	map_attribute_definition_xhtml_showintitleentryDB.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPointersEncoding = input.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPointersEncoding

	db, _ = db.Model(&map_attribute_definition_xhtml_showintitleentryDB)
	_, err = db.Updates(&map_attribute_definition_xhtml_showintitleentryDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	map_attribute_definition_xhtml_showintitleentryNew := new(models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry)
	map_attribute_definition_xhtml_showintitleentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry(map_attribute_definition_xhtml_showintitleentryNew)

	// redeem pointers
	map_attribute_definition_xhtml_showintitleentryDB.DecodePointers(backRepo, map_attribute_definition_xhtml_showintitleentryNew)

	// get stage instance from DB instance, and call callback function
	map_attribute_definition_xhtml_showintitleentryOld := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDBID_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPtr[map_attribute_definition_xhtml_showintitleentryDB.ID]
	if map_attribute_definition_xhtml_showintitleentryOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), map_attribute_definition_xhtml_showintitleentryOld, map_attribute_definition_xhtml_showintitleentryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the map_attribute_definition_xhtml_showintitleentryDB
	c.JSON(http.StatusOK, map_attribute_definition_xhtml_showintitleentryDB)
}

// DeleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
//
// swagger:route DELETE /map_attribute_definition_xhtml_showintitleentrys/{ID} map_attribute_definition_xhtml_showintitleentrys deleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry
//
// # Delete a map_attribute_definition_xhtml_showintitleentry
//
// default: genericError
//
//	200: map_attribute_definition_xhtml_showintitleentryDBResponse
func (controller *Controller) DeleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry(c *gin.Context) {

	mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Lock()
	defer mutexMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry", "Name", stackPath)
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
	db := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.GetDB()

	// Get model if exist
	var map_attribute_definition_xhtml_showintitleentryDB orm.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDB
	if _, err := db.First(&map_attribute_definition_xhtml_showintitleentryDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&map_attribute_definition_xhtml_showintitleentryDB)

	// get an instance (not staged) from DB instance, and call callback function
	map_attribute_definition_xhtml_showintitleentryDeleted := new(models.Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry)
	map_attribute_definition_xhtml_showintitleentryDB.CopyBasicFieldsToMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry(map_attribute_definition_xhtml_showintitleentryDeleted)

	// get stage instance from DB instance, and call callback function
	map_attribute_definition_xhtml_showintitleentryStaged := backRepo.BackRepoMap_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntry.Map_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryDBID_Map_ATTRIBUTE_DEFINITION_XHTML_ShowInTitleEntryPtr[map_attribute_definition_xhtml_showintitleentryDB.ID]
	if map_attribute_definition_xhtml_showintitleentryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), map_attribute_definition_xhtml_showintitleentryStaged, map_attribute_definition_xhtml_showintitleentryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
