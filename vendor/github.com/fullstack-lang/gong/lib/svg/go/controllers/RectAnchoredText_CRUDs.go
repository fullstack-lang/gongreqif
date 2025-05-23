// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __RectAnchoredText__dummysDeclaration__ models.RectAnchoredText
var __RectAnchoredText_time__dummyDeclaration time.Duration

var mutexRectAnchoredText sync.Mutex

// An RectAnchoredTextID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRectAnchoredText updateRectAnchoredText deleteRectAnchoredText
type RectAnchoredTextID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RectAnchoredTextInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRectAnchoredText updateRectAnchoredText
type RectAnchoredTextInput struct {
	// The RectAnchoredText to submit or modify
	// in: body
	RectAnchoredText *orm.RectAnchoredTextAPI
}

// GetRectAnchoredTexts
//
// swagger:route GET /rectanchoredtexts rectanchoredtexts getRectAnchoredTexts
//
// # Get all rectanchoredtexts
//
// Responses:
// default: genericError
//
//	200: rectanchoredtextDBResponse
func (controller *Controller) GetRectAnchoredTexts(c *gin.Context) {

	// source slice
	var rectanchoredtextDBs []orm.RectAnchoredTextDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredTexts", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoRectAnchoredText.GetDB()

	_, err := db.Find(&rectanchoredtextDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	rectanchoredtextAPIs := make([]orm.RectAnchoredTextAPI, 0)

	// for each rectanchoredtext, update fields from the database nullable fields
	for idx := range rectanchoredtextDBs {
		rectanchoredtextDB := &rectanchoredtextDBs[idx]
		_ = rectanchoredtextDB
		var rectanchoredtextAPI orm.RectAnchoredTextAPI

		// insertion point for updating fields
		rectanchoredtextAPI.ID = rectanchoredtextDB.ID
		rectanchoredtextDB.CopyBasicFieldsToRectAnchoredText_WOP(&rectanchoredtextAPI.RectAnchoredText_WOP)
		rectanchoredtextAPI.RectAnchoredTextPointersEncoding = rectanchoredtextDB.RectAnchoredTextPointersEncoding
		rectanchoredtextAPIs = append(rectanchoredtextAPIs, rectanchoredtextAPI)
	}

	c.JSON(http.StatusOK, rectanchoredtextAPIs)
}

// PostRectAnchoredText
//
// swagger:route POST /rectanchoredtexts rectanchoredtexts postRectAnchoredText
//
// Creates a rectanchoredtext
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRectAnchoredText(c *gin.Context) {

	mutexRectAnchoredText.Lock()
	defer mutexRectAnchoredText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRectAnchoredTexts", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoRectAnchoredText.GetDB()

	// Validate input
	var input orm.RectAnchoredTextAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create rectanchoredtext
	rectanchoredtextDB := orm.RectAnchoredTextDB{}
	rectanchoredtextDB.RectAnchoredTextPointersEncoding = input.RectAnchoredTextPointersEncoding
	rectanchoredtextDB.CopyBasicFieldsFromRectAnchoredText_WOP(&input.RectAnchoredText_WOP)

	_, err = db.Create(&rectanchoredtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRectAnchoredText.CheckoutPhaseOneInstance(&rectanchoredtextDB)
	rectanchoredtext := backRepo.BackRepoRectAnchoredText.Map_RectAnchoredTextDBID_RectAnchoredTextPtr[rectanchoredtextDB.ID]

	if rectanchoredtext != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), rectanchoredtext)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, rectanchoredtextDB)
}

// GetRectAnchoredText
//
// swagger:route GET /rectanchoredtexts/{ID} rectanchoredtexts getRectAnchoredText
//
// Gets the details for a rectanchoredtext.
//
// Responses:
// default: genericError
//
//	200: rectanchoredtextDBResponse
func (controller *Controller) GetRectAnchoredText(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRectAnchoredText", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoRectAnchoredText.GetDB()

	// Get rectanchoredtextDB in DB
	var rectanchoredtextDB orm.RectAnchoredTextDB
	if _, err := db.First(&rectanchoredtextDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var rectanchoredtextAPI orm.RectAnchoredTextAPI
	rectanchoredtextAPI.ID = rectanchoredtextDB.ID
	rectanchoredtextAPI.RectAnchoredTextPointersEncoding = rectanchoredtextDB.RectAnchoredTextPointersEncoding
	rectanchoredtextDB.CopyBasicFieldsToRectAnchoredText_WOP(&rectanchoredtextAPI.RectAnchoredText_WOP)

	c.JSON(http.StatusOK, rectanchoredtextAPI)
}

// UpdateRectAnchoredText
//
// swagger:route PATCH /rectanchoredtexts/{ID} rectanchoredtexts updateRectAnchoredText
//
// # Update a rectanchoredtext
//
// Responses:
// default: genericError
//
//	200: rectanchoredtextDBResponse
func (controller *Controller) UpdateRectAnchoredText(c *gin.Context) {

	mutexRectAnchoredText.Lock()
	defer mutexRectAnchoredText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRectAnchoredText", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoRectAnchoredText.GetDB()

	// Validate input
	var input orm.RectAnchoredTextAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var rectanchoredtextDB orm.RectAnchoredTextDB

	// fetch the rectanchoredtext
	_, err := db.First(&rectanchoredtextDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	rectanchoredtextDB.CopyBasicFieldsFromRectAnchoredText_WOP(&input.RectAnchoredText_WOP)
	rectanchoredtextDB.RectAnchoredTextPointersEncoding = input.RectAnchoredTextPointersEncoding

	db, _ = db.Model(&rectanchoredtextDB)
	_, err = db.Updates(&rectanchoredtextDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredtextNew := new(models.RectAnchoredText)
	rectanchoredtextDB.CopyBasicFieldsToRectAnchoredText(rectanchoredtextNew)

	// redeem pointers
	rectanchoredtextDB.DecodePointers(backRepo, rectanchoredtextNew)

	// get stage instance from DB instance, and call callback function
	rectanchoredtextOld := backRepo.BackRepoRectAnchoredText.Map_RectAnchoredTextDBID_RectAnchoredTextPtr[rectanchoredtextDB.ID]
	if rectanchoredtextOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), rectanchoredtextOld, rectanchoredtextNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the rectanchoredtextDB
	c.JSON(http.StatusOK, rectanchoredtextDB)
}

// DeleteRectAnchoredText
//
// swagger:route DELETE /rectanchoredtexts/{ID} rectanchoredtexts deleteRectAnchoredText
//
// # Delete a rectanchoredtext
//
// default: genericError
//
//	200: rectanchoredtextDBResponse
func (controller *Controller) DeleteRectAnchoredText(c *gin.Context) {

	mutexRectAnchoredText.Lock()
	defer mutexRectAnchoredText.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRectAnchoredText", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/lib/svg/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoRectAnchoredText.GetDB()

	// Get model if exist
	var rectanchoredtextDB orm.RectAnchoredTextDB
	if _, err := db.First(&rectanchoredtextDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&rectanchoredtextDB)

	// get an instance (not staged) from DB instance, and call callback function
	rectanchoredtextDeleted := new(models.RectAnchoredText)
	rectanchoredtextDB.CopyBasicFieldsToRectAnchoredText(rectanchoredtextDeleted)

	// get stage instance from DB instance, and call callback function
	rectanchoredtextStaged := backRepo.BackRepoRectAnchoredText.Map_RectAnchoredTextDBID_RectAnchoredTextPtr[rectanchoredtextDB.ID]
	if rectanchoredtextStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), rectanchoredtextStaged, rectanchoredtextDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
