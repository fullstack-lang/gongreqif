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
var __StaticWebSiteChapter__dummysDeclaration__ models.StaticWebSiteChapter
var __StaticWebSiteChapter_time__dummyDeclaration time.Duration

var mutexStaticWebSiteChapter sync.Mutex

// An StaticWebSiteChapterID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaticWebSiteChapter updateStaticWebSiteChapter deleteStaticWebSiteChapter
type StaticWebSiteChapterID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StaticWebSiteChapterInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaticWebSiteChapter updateStaticWebSiteChapter
type StaticWebSiteChapterInput struct {
	// The StaticWebSiteChapter to submit or modify
	// in: body
	StaticWebSiteChapter *orm.StaticWebSiteChapterAPI
}

// GetStaticWebSiteChapters
//
// swagger:route GET /staticwebsitechapters staticwebsitechapters getStaticWebSiteChapters
//
// # Get all staticwebsitechapters
//
// Responses:
// default: genericError
//
//	200: staticwebsitechapterDBResponse
func (controller *Controller) GetStaticWebSiteChapters(c *gin.Context) {

	// source slice
	var staticwebsitechapterDBs []orm.StaticWebSiteChapterDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSiteChapters", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteChapter.GetDB()

	_, err := db.Find(&staticwebsitechapterDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staticwebsitechapterAPIs := make([]orm.StaticWebSiteChapterAPI, 0)

	// for each staticwebsitechapter, update fields from the database nullable fields
	for idx := range staticwebsitechapterDBs {
		staticwebsitechapterDB := &staticwebsitechapterDBs[idx]
		_ = staticwebsitechapterDB
		var staticwebsitechapterAPI orm.StaticWebSiteChapterAPI

		// insertion point for updating fields
		staticwebsitechapterAPI.ID = staticwebsitechapterDB.ID
		staticwebsitechapterDB.CopyBasicFieldsToStaticWebSiteChapter_WOP(&staticwebsitechapterAPI.StaticWebSiteChapter_WOP)
		staticwebsitechapterAPI.StaticWebSiteChapterPointersEncoding = staticwebsitechapterDB.StaticWebSiteChapterPointersEncoding
		staticwebsitechapterAPIs = append(staticwebsitechapterAPIs, staticwebsitechapterAPI)
	}

	c.JSON(http.StatusOK, staticwebsitechapterAPIs)
}

// PostStaticWebSiteChapter
//
// swagger:route POST /staticwebsitechapters staticwebsitechapters postStaticWebSiteChapter
//
// Creates a staticwebsitechapter
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaticWebSiteChapter(c *gin.Context) {

	mutexStaticWebSiteChapter.Lock()
	defer mutexStaticWebSiteChapter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaticWebSiteChapters", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteChapter.GetDB()

	// Validate input
	var input orm.StaticWebSiteChapterAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staticwebsitechapter
	staticwebsitechapterDB := orm.StaticWebSiteChapterDB{}
	staticwebsitechapterDB.StaticWebSiteChapterPointersEncoding = input.StaticWebSiteChapterPointersEncoding
	staticwebsitechapterDB.CopyBasicFieldsFromStaticWebSiteChapter_WOP(&input.StaticWebSiteChapter_WOP)

	_, err = db.Create(&staticwebsitechapterDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaticWebSiteChapter.CheckoutPhaseOneInstance(&staticwebsitechapterDB)
	staticwebsitechapter := backRepo.BackRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr[staticwebsitechapterDB.ID]

	if staticwebsitechapter != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staticwebsitechapter)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staticwebsitechapterDB)
}

// GetStaticWebSiteChapter
//
// swagger:route GET /staticwebsitechapters/{ID} staticwebsitechapters getStaticWebSiteChapter
//
// Gets the details for a staticwebsitechapter.
//
// Responses:
// default: genericError
//
//	200: staticwebsitechapterDBResponse
func (controller *Controller) GetStaticWebSiteChapter(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSiteChapter", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteChapter.GetDB()

	// Get staticwebsitechapterDB in DB
	var staticwebsitechapterDB orm.StaticWebSiteChapterDB
	if _, err := db.First(&staticwebsitechapterDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staticwebsitechapterAPI orm.StaticWebSiteChapterAPI
	staticwebsitechapterAPI.ID = staticwebsitechapterDB.ID
	staticwebsitechapterAPI.StaticWebSiteChapterPointersEncoding = staticwebsitechapterDB.StaticWebSiteChapterPointersEncoding
	staticwebsitechapterDB.CopyBasicFieldsToStaticWebSiteChapter_WOP(&staticwebsitechapterAPI.StaticWebSiteChapter_WOP)

	c.JSON(http.StatusOK, staticwebsitechapterAPI)
}

// UpdateStaticWebSiteChapter
//
// swagger:route PATCH /staticwebsitechapters/{ID} staticwebsitechapters updateStaticWebSiteChapter
//
// # Update a staticwebsitechapter
//
// Responses:
// default: genericError
//
//	200: staticwebsitechapterDBResponse
func (controller *Controller) UpdateStaticWebSiteChapter(c *gin.Context) {

	mutexStaticWebSiteChapter.Lock()
	defer mutexStaticWebSiteChapter.Unlock()

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
	db := backRepo.BackRepoStaticWebSiteChapter.GetDB()

	// Validate input
	var input orm.StaticWebSiteChapterAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staticwebsitechapterDB orm.StaticWebSiteChapterDB

	// fetch the staticwebsitechapter
	_, err := db.First(&staticwebsitechapterDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staticwebsitechapterDB.CopyBasicFieldsFromStaticWebSiteChapter_WOP(&input.StaticWebSiteChapter_WOP)
	staticwebsitechapterDB.StaticWebSiteChapterPointersEncoding = input.StaticWebSiteChapterPointersEncoding

	db, _ = db.Model(&staticwebsitechapterDB)
	_, err = db.Updates(&staticwebsitechapterDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsitechapterNew := new(models.StaticWebSiteChapter)
	staticwebsitechapterDB.CopyBasicFieldsToStaticWebSiteChapter(staticwebsitechapterNew)

	// redeem pointers
	staticwebsitechapterDB.DecodePointers(backRepo, staticwebsitechapterNew)

	// get stage instance from DB instance, and call callback function
	staticwebsitechapterOld := backRepo.BackRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr[staticwebsitechapterDB.ID]
	if staticwebsitechapterOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), staticwebsitechapterOld, staticwebsitechapterNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), staticwebsitechapterOld, staticwebsitechapterNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staticwebsitechapterDB
	c.JSON(http.StatusOK, staticwebsitechapterDB)
}

// DeleteStaticWebSiteChapter
//
// swagger:route DELETE /staticwebsitechapters/{ID} staticwebsitechapters deleteStaticWebSiteChapter
//
// # Delete a staticwebsitechapter
//
// default: genericError
//
//	200: staticwebsitechapterDBResponse
func (controller *Controller) DeleteStaticWebSiteChapter(c *gin.Context) {

	mutexStaticWebSiteChapter.Lock()
	defer mutexStaticWebSiteChapter.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaticWebSiteChapter", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteChapter.GetDB()

	// Get model if exist
	var staticwebsitechapterDB orm.StaticWebSiteChapterDB
	if _, err := db.First(&staticwebsitechapterDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&staticwebsitechapterDB)

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsitechapterDeleted := new(models.StaticWebSiteChapter)
	staticwebsitechapterDB.CopyBasicFieldsToStaticWebSiteChapter(staticwebsitechapterDeleted)

	// get stage instance from DB instance, and call callback function
	staticwebsitechapterStaged := backRepo.BackRepoStaticWebSiteChapter.Map_StaticWebSiteChapterDBID_StaticWebSiteChapterPtr[staticwebsitechapterDB.ID]
	if staticwebsitechapterStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staticwebsitechapterStaged, staticwebsitechapterDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
