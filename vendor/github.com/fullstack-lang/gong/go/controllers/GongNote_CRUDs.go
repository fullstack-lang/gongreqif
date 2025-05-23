// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __GongNote__dummysDeclaration__ models.GongNote
var __GongNote_time__dummyDeclaration time.Duration

var mutexGongNote sync.Mutex

// An GongNoteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGongNote updateGongNote deleteGongNote
type GongNoteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GongNoteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGongNote updateGongNote
type GongNoteInput struct {
	// The GongNote to submit or modify
	// in: body
	GongNote *orm.GongNoteAPI
}

// GetGongNotes
//
// swagger:route GET /gongnotes gongnotes getGongNotes
//
// # Get all gongnotes
//
// Responses:
// default: genericError
//
//	200: gongnoteDBResponse
func (controller *Controller) GetGongNotes(c *gin.Context) {

	// source slice
	var gongnoteDBs []orm.GongNoteDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongNotes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongNote.GetDB()

	_, err := db.Find(&gongnoteDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	gongnoteAPIs := make([]orm.GongNoteAPI, 0)

	// for each gongnote, update fields from the database nullable fields
	for idx := range gongnoteDBs {
		gongnoteDB := &gongnoteDBs[idx]
		_ = gongnoteDB
		var gongnoteAPI orm.GongNoteAPI

		// insertion point for updating fields
		gongnoteAPI.ID = gongnoteDB.ID
		gongnoteDB.CopyBasicFieldsToGongNote_WOP(&gongnoteAPI.GongNote_WOP)
		gongnoteAPI.GongNotePointersEncoding = gongnoteDB.GongNotePointersEncoding
		gongnoteAPIs = append(gongnoteAPIs, gongnoteAPI)
	}

	c.JSON(http.StatusOK, gongnoteAPIs)
}

// PostGongNote
//
// swagger:route POST /gongnotes gongnotes postGongNote
//
// Creates a gongnote
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGongNote(c *gin.Context) {

	mutexGongNote.Lock()
	defer mutexGongNote.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGongNotes", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongNote.GetDB()

	// Validate input
	var input orm.GongNoteAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gongnote
	gongnoteDB := orm.GongNoteDB{}
	gongnoteDB.GongNotePointersEncoding = input.GongNotePointersEncoding
	gongnoteDB.CopyBasicFieldsFromGongNote_WOP(&input.GongNote_WOP)

	_, err = db.Create(&gongnoteDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGongNote.CheckoutPhaseOneInstance(&gongnoteDB)
	gongnote := backRepo.BackRepoGongNote.Map_GongNoteDBID_GongNotePtr[gongnoteDB.ID]

	if gongnote != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), gongnote)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gongnoteDB)
}

// GetGongNote
//
// swagger:route GET /gongnotes/{ID} gongnotes getGongNote
//
// Gets the details for a gongnote.
//
// Responses:
// default: genericError
//
//	200: gongnoteDBResponse
func (controller *Controller) GetGongNote(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGongNote", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongNote.GetDB()

	// Get gongnoteDB in DB
	var gongnoteDB orm.GongNoteDB
	if _, err := db.First(&gongnoteDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var gongnoteAPI orm.GongNoteAPI
	gongnoteAPI.ID = gongnoteDB.ID
	gongnoteAPI.GongNotePointersEncoding = gongnoteDB.GongNotePointersEncoding
	gongnoteDB.CopyBasicFieldsToGongNote_WOP(&gongnoteAPI.GongNote_WOP)

	c.JSON(http.StatusOK, gongnoteAPI)
}

// UpdateGongNote
//
// swagger:route PATCH /gongnotes/{ID} gongnotes updateGongNote
//
// # Update a gongnote
//
// Responses:
// default: genericError
//
//	200: gongnoteDBResponse
func (controller *Controller) UpdateGongNote(c *gin.Context) {

	mutexGongNote.Lock()
	defer mutexGongNote.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGongNote", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongNote.GetDB()

	// Validate input
	var input orm.GongNoteAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var gongnoteDB orm.GongNoteDB

	// fetch the gongnote
	_, err := db.First(&gongnoteDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	gongnoteDB.CopyBasicFieldsFromGongNote_WOP(&input.GongNote_WOP)
	gongnoteDB.GongNotePointersEncoding = input.GongNotePointersEncoding

	db, _ = db.Model(&gongnoteDB)
	_, err = db.Updates(&gongnoteDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	gongnoteNew := new(models.GongNote)
	gongnoteDB.CopyBasicFieldsToGongNote(gongnoteNew)

	// redeem pointers
	gongnoteDB.DecodePointers(backRepo, gongnoteNew)

	// get stage instance from DB instance, and call callback function
	gongnoteOld := backRepo.BackRepoGongNote.Map_GongNoteDBID_GongNotePtr[gongnoteDB.ID]
	if gongnoteOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), gongnoteOld, gongnoteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the gongnoteDB
	c.JSON(http.StatusOK, gongnoteDB)
}

// DeleteGongNote
//
// swagger:route DELETE /gongnotes/{ID} gongnotes deleteGongNote
//
// # Delete a gongnote
//
// default: genericError
//
//	200: gongnoteDBResponse
func (controller *Controller) DeleteGongNote(c *gin.Context) {

	mutexGongNote.Lock()
	defer mutexGongNote.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGongNote", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoGongNote.GetDB()

	// Get model if exist
	var gongnoteDB orm.GongNoteDB
	if _, err := db.First(&gongnoteDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&gongnoteDB)

	// get an instance (not staged) from DB instance, and call callback function
	gongnoteDeleted := new(models.GongNote)
	gongnoteDB.CopyBasicFieldsToGongNote(gongnoteDeleted)

	// get stage instance from DB instance, and call callback function
	gongnoteStaged := backRepo.BackRepoGongNote.Map_GongNoteDBID_GongNotePtr[gongnoteDB.ID]
	if gongnoteStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), gongnoteStaged, gongnoteDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
