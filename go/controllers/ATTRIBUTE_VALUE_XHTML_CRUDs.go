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
var __ATTRIBUTE_VALUE_XHTML__dummysDeclaration__ models.ATTRIBUTE_VALUE_XHTML
var __ATTRIBUTE_VALUE_XHTML_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_VALUE_XHTML sync.Mutex

// An ATTRIBUTE_VALUE_XHTMLID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_VALUE_XHTML updateATTRIBUTE_VALUE_XHTML deleteATTRIBUTE_VALUE_XHTML
type ATTRIBUTE_VALUE_XHTMLID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_VALUE_XHTMLInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_VALUE_XHTML updateATTRIBUTE_VALUE_XHTML
type ATTRIBUTE_VALUE_XHTMLInput struct {
	// The ATTRIBUTE_VALUE_XHTML to submit or modify
	// in: body
	ATTRIBUTE_VALUE_XHTML *orm.ATTRIBUTE_VALUE_XHTMLAPI
}

// GetATTRIBUTE_VALUE_XHTMLs
//
// swagger:route GET /attribute_value_xhtmls attribute_value_xhtmls getATTRIBUTE_VALUE_XHTMLs
//
// # Get all attribute_value_xhtmls
//
// Responses:
// default: genericError
//
//	200: attribute_value_xhtmlDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_XHTMLs(c *gin.Context) {

	// source slice
	var attribute_value_xhtmlDBs []orm.ATTRIBUTE_VALUE_XHTMLDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_XHTMLs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.GetDB()

	_, err := db.Find(&attribute_value_xhtmlDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_value_xhtmlAPIs := make([]orm.ATTRIBUTE_VALUE_XHTMLAPI, 0)

	// for each attribute_value_xhtml, update fields from the database nullable fields
	for idx := range attribute_value_xhtmlDBs {
		attribute_value_xhtmlDB := &attribute_value_xhtmlDBs[idx]
		_ = attribute_value_xhtmlDB
		var attribute_value_xhtmlAPI orm.ATTRIBUTE_VALUE_XHTMLAPI

		// insertion point for updating fields
		attribute_value_xhtmlAPI.ID = attribute_value_xhtmlDB.ID
		attribute_value_xhtmlDB.CopyBasicFieldsToATTRIBUTE_VALUE_XHTML_WOP(&attribute_value_xhtmlAPI.ATTRIBUTE_VALUE_XHTML_WOP)
		attribute_value_xhtmlAPI.ATTRIBUTE_VALUE_XHTMLPointersEncoding = attribute_value_xhtmlDB.ATTRIBUTE_VALUE_XHTMLPointersEncoding
		attribute_value_xhtmlAPIs = append(attribute_value_xhtmlAPIs, attribute_value_xhtmlAPI)
	}

	c.JSON(http.StatusOK, attribute_value_xhtmlAPIs)
}

// PostATTRIBUTE_VALUE_XHTML
//
// swagger:route POST /attribute_value_xhtmls attribute_value_xhtmls postATTRIBUTE_VALUE_XHTML
//
// Creates a attribute_value_xhtml
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_VALUE_XHTML(c *gin.Context) {

	mutexATTRIBUTE_VALUE_XHTML.Lock()
	defer mutexATTRIBUTE_VALUE_XHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_VALUE_XHTMLs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_XHTMLAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_value_xhtml
	attribute_value_xhtmlDB := orm.ATTRIBUTE_VALUE_XHTMLDB{}
	attribute_value_xhtmlDB.ATTRIBUTE_VALUE_XHTMLPointersEncoding = input.ATTRIBUTE_VALUE_XHTMLPointersEncoding
	attribute_value_xhtmlDB.CopyBasicFieldsFromATTRIBUTE_VALUE_XHTML_WOP(&input.ATTRIBUTE_VALUE_XHTML_WOP)

	_, err = db.Create(&attribute_value_xhtmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseOneInstance(&attribute_value_xhtmlDB)
	attribute_value_xhtml := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr[attribute_value_xhtmlDB.ID]

	if attribute_value_xhtml != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_value_xhtml)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_value_xhtmlDB)
}

// GetATTRIBUTE_VALUE_XHTML
//
// swagger:route GET /attribute_value_xhtmls/{ID} attribute_value_xhtmls getATTRIBUTE_VALUE_XHTML
//
// Gets the details for a attribute_value_xhtml.
//
// Responses:
// default: genericError
//
//	200: attribute_value_xhtmlDBResponse
func (controller *Controller) GetATTRIBUTE_VALUE_XHTML(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_VALUE_XHTML", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.GetDB()

	// Get attribute_value_xhtmlDB in DB
	var attribute_value_xhtmlDB orm.ATTRIBUTE_VALUE_XHTMLDB
	if _, err := db.First(&attribute_value_xhtmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_value_xhtmlAPI orm.ATTRIBUTE_VALUE_XHTMLAPI
	attribute_value_xhtmlAPI.ID = attribute_value_xhtmlDB.ID
	attribute_value_xhtmlAPI.ATTRIBUTE_VALUE_XHTMLPointersEncoding = attribute_value_xhtmlDB.ATTRIBUTE_VALUE_XHTMLPointersEncoding
	attribute_value_xhtmlDB.CopyBasicFieldsToATTRIBUTE_VALUE_XHTML_WOP(&attribute_value_xhtmlAPI.ATTRIBUTE_VALUE_XHTML_WOP)

	c.JSON(http.StatusOK, attribute_value_xhtmlAPI)
}

// UpdateATTRIBUTE_VALUE_XHTML
//
// swagger:route PATCH /attribute_value_xhtmls/{ID} attribute_value_xhtmls updateATTRIBUTE_VALUE_XHTML
//
// # Update a attribute_value_xhtml
//
// Responses:
// default: genericError
//
//	200: attribute_value_xhtmlDBResponse
func (controller *Controller) UpdateATTRIBUTE_VALUE_XHTML(c *gin.Context) {

	mutexATTRIBUTE_VALUE_XHTML.Lock()
	defer mutexATTRIBUTE_VALUE_XHTML.Unlock()

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
	db := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_VALUE_XHTMLAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_value_xhtmlDB orm.ATTRIBUTE_VALUE_XHTMLDB

	// fetch the attribute_value_xhtml
	_, err := db.First(&attribute_value_xhtmlDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_value_xhtmlDB.CopyBasicFieldsFromATTRIBUTE_VALUE_XHTML_WOP(&input.ATTRIBUTE_VALUE_XHTML_WOP)
	attribute_value_xhtmlDB.ATTRIBUTE_VALUE_XHTMLPointersEncoding = input.ATTRIBUTE_VALUE_XHTMLPointersEncoding

	db, _ = db.Model(&attribute_value_xhtmlDB)
	_, err = db.Updates(&attribute_value_xhtmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_xhtmlNew := new(models.ATTRIBUTE_VALUE_XHTML)
	attribute_value_xhtmlDB.CopyBasicFieldsToATTRIBUTE_VALUE_XHTML(attribute_value_xhtmlNew)

	// redeem pointers
	attribute_value_xhtmlDB.DecodePointers(backRepo, attribute_value_xhtmlNew)

	// get stage instance from DB instance, and call callback function
	attribute_value_xhtmlOld := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr[attribute_value_xhtmlDB.ID]
	if attribute_value_xhtmlOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_value_xhtmlOld, attribute_value_xhtmlNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_value_xhtmlOld, attribute_value_xhtmlNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_value_xhtmlDB
	c.JSON(http.StatusOK, attribute_value_xhtmlDB)
}

// DeleteATTRIBUTE_VALUE_XHTML
//
// swagger:route DELETE /attribute_value_xhtmls/{ID} attribute_value_xhtmls deleteATTRIBUTE_VALUE_XHTML
//
// # Delete a attribute_value_xhtml
//
// default: genericError
//
//	200: attribute_value_xhtmlDBResponse
func (controller *Controller) DeleteATTRIBUTE_VALUE_XHTML(c *gin.Context) {

	mutexATTRIBUTE_VALUE_XHTML.Lock()
	defer mutexATTRIBUTE_VALUE_XHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_VALUE_XHTML", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.GetDB()

	// Get model if exist
	var attribute_value_xhtmlDB orm.ATTRIBUTE_VALUE_XHTMLDB
	if _, err := db.First(&attribute_value_xhtmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_value_xhtmlDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_value_xhtmlDeleted := new(models.ATTRIBUTE_VALUE_XHTML)
	attribute_value_xhtmlDB.CopyBasicFieldsToATTRIBUTE_VALUE_XHTML(attribute_value_xhtmlDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_value_xhtmlStaged := backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr[attribute_value_xhtmlDB.ID]
	if attribute_value_xhtmlStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_value_xhtmlStaged, attribute_value_xhtmlDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
