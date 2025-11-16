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
var __ATTRIBUTE_DEFINITION_XHTML__dummysDeclaration__ models.ATTRIBUTE_DEFINITION_XHTML
var __ATTRIBUTE_DEFINITION_XHTML_time__dummyDeclaration time.Duration

var mutexATTRIBUTE_DEFINITION_XHTML sync.Mutex

// An ATTRIBUTE_DEFINITION_XHTMLID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getATTRIBUTE_DEFINITION_XHTML updateATTRIBUTE_DEFINITION_XHTML deleteATTRIBUTE_DEFINITION_XHTML
type ATTRIBUTE_DEFINITION_XHTMLID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ATTRIBUTE_DEFINITION_XHTMLInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postATTRIBUTE_DEFINITION_XHTML updateATTRIBUTE_DEFINITION_XHTML
type ATTRIBUTE_DEFINITION_XHTMLInput struct {
	// The ATTRIBUTE_DEFINITION_XHTML to submit or modify
	// in: body
	ATTRIBUTE_DEFINITION_XHTML *orm.ATTRIBUTE_DEFINITION_XHTMLAPI
}

// GetATTRIBUTE_DEFINITION_XHTMLs
//
// swagger:route GET /attribute_definition_xhtmls attribute_definition_xhtmls getATTRIBUTE_DEFINITION_XHTMLs
//
// # Get all attribute_definition_xhtmls
//
// Responses:
// default: genericError
//
//	200: attribute_definition_xhtmlDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_XHTMLs(c *gin.Context) {

	// source slice
	var attribute_definition_xhtmlDBs []orm.ATTRIBUTE_DEFINITION_XHTMLDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_XHTMLs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.GetDB()

	_, err := db.Find(&attribute_definition_xhtmlDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	attribute_definition_xhtmlAPIs := make([]orm.ATTRIBUTE_DEFINITION_XHTMLAPI, 0)

	// for each attribute_definition_xhtml, update fields from the database nullable fields
	for idx := range attribute_definition_xhtmlDBs {
		attribute_definition_xhtmlDB := &attribute_definition_xhtmlDBs[idx]
		_ = attribute_definition_xhtmlDB
		var attribute_definition_xhtmlAPI orm.ATTRIBUTE_DEFINITION_XHTMLAPI

		// insertion point for updating fields
		attribute_definition_xhtmlAPI.ID = attribute_definition_xhtmlDB.ID
		attribute_definition_xhtmlDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_XHTML_WOP(&attribute_definition_xhtmlAPI.ATTRIBUTE_DEFINITION_XHTML_WOP)
		attribute_definition_xhtmlAPI.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding = attribute_definition_xhtmlDB.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding
		attribute_definition_xhtmlAPIs = append(attribute_definition_xhtmlAPIs, attribute_definition_xhtmlAPI)
	}

	c.JSON(http.StatusOK, attribute_definition_xhtmlAPIs)
}

// PostATTRIBUTE_DEFINITION_XHTML
//
// swagger:route POST /attribute_definition_xhtmls attribute_definition_xhtmls postATTRIBUTE_DEFINITION_XHTML
//
// Creates a attribute_definition_xhtml
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostATTRIBUTE_DEFINITION_XHTML(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_XHTML.Lock()
	defer mutexATTRIBUTE_DEFINITION_XHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostATTRIBUTE_DEFINITION_XHTMLs", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_XHTMLAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create attribute_definition_xhtml
	attribute_definition_xhtmlDB := orm.ATTRIBUTE_DEFINITION_XHTMLDB{}
	attribute_definition_xhtmlDB.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding = input.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding
	attribute_definition_xhtmlDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_XHTML_WOP(&input.ATTRIBUTE_DEFINITION_XHTML_WOP)

	_, err = db.Create(&attribute_definition_xhtmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.CheckoutPhaseOneInstance(&attribute_definition_xhtmlDB)
	attribute_definition_xhtml := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.Map_ATTRIBUTE_DEFINITION_XHTMLDBID_ATTRIBUTE_DEFINITION_XHTMLPtr[attribute_definition_xhtmlDB.ID]

	if attribute_definition_xhtml != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), attribute_definition_xhtml)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, attribute_definition_xhtmlDB)
}

// GetATTRIBUTE_DEFINITION_XHTML
//
// swagger:route GET /attribute_definition_xhtmls/{ID} attribute_definition_xhtmls getATTRIBUTE_DEFINITION_XHTML
//
// Gets the details for a attribute_definition_xhtml.
//
// Responses:
// default: genericError
//
//	200: attribute_definition_xhtmlDBResponse
func (controller *Controller) GetATTRIBUTE_DEFINITION_XHTML(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetATTRIBUTE_DEFINITION_XHTML", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.GetDB()

	// Get attribute_definition_xhtmlDB in DB
	var attribute_definition_xhtmlDB orm.ATTRIBUTE_DEFINITION_XHTMLDB
	if _, err := db.First(&attribute_definition_xhtmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var attribute_definition_xhtmlAPI orm.ATTRIBUTE_DEFINITION_XHTMLAPI
	attribute_definition_xhtmlAPI.ID = attribute_definition_xhtmlDB.ID
	attribute_definition_xhtmlAPI.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding = attribute_definition_xhtmlDB.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding
	attribute_definition_xhtmlDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_XHTML_WOP(&attribute_definition_xhtmlAPI.ATTRIBUTE_DEFINITION_XHTML_WOP)

	c.JSON(http.StatusOK, attribute_definition_xhtmlAPI)
}

// UpdateATTRIBUTE_DEFINITION_XHTML
//
// swagger:route PATCH /attribute_definition_xhtmls/{ID} attribute_definition_xhtmls updateATTRIBUTE_DEFINITION_XHTML
//
// # Update a attribute_definition_xhtml
//
// Responses:
// default: genericError
//
//	200: attribute_definition_xhtmlDBResponse
func (controller *Controller) UpdateATTRIBUTE_DEFINITION_XHTML(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_XHTML.Lock()
	defer mutexATTRIBUTE_DEFINITION_XHTML.Unlock()

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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.GetDB()

	// Validate input
	var input orm.ATTRIBUTE_DEFINITION_XHTMLAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var attribute_definition_xhtmlDB orm.ATTRIBUTE_DEFINITION_XHTMLDB

	// fetch the attribute_definition_xhtml
	_, err := db.First(&attribute_definition_xhtmlDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	attribute_definition_xhtmlDB.CopyBasicFieldsFromATTRIBUTE_DEFINITION_XHTML_WOP(&input.ATTRIBUTE_DEFINITION_XHTML_WOP)
	attribute_definition_xhtmlDB.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding = input.ATTRIBUTE_DEFINITION_XHTMLPointersEncoding

	db, _ = db.Model(&attribute_definition_xhtmlDB)
	_, err = db.Updates(&attribute_definition_xhtmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_xhtmlNew := new(models.ATTRIBUTE_DEFINITION_XHTML)
	attribute_definition_xhtmlDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtmlNew)

	// redeem pointers
	attribute_definition_xhtmlDB.DecodePointers(backRepo, attribute_definition_xhtmlNew)

	// get stage instance from DB instance, and call callback function
	attribute_definition_xhtmlOld := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.Map_ATTRIBUTE_DEFINITION_XHTMLDBID_ATTRIBUTE_DEFINITION_XHTMLPtr[attribute_definition_xhtmlDB.ID]
	if attribute_definition_xhtmlOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), attribute_definition_xhtmlOld, attribute_definition_xhtmlNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the attribute_definition_xhtmlDB
	c.JSON(http.StatusOK, attribute_definition_xhtmlDB)
}

// DeleteATTRIBUTE_DEFINITION_XHTML
//
// swagger:route DELETE /attribute_definition_xhtmls/{ID} attribute_definition_xhtmls deleteATTRIBUTE_DEFINITION_XHTML
//
// # Delete a attribute_definition_xhtml
//
// default: genericError
//
//	200: attribute_definition_xhtmlDBResponse
func (controller *Controller) DeleteATTRIBUTE_DEFINITION_XHTML(c *gin.Context) {

	mutexATTRIBUTE_DEFINITION_XHTML.Lock()
	defer mutexATTRIBUTE_DEFINITION_XHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteATTRIBUTE_DEFINITION_XHTML", "Name", stackPath)
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
	db := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.GetDB()

	// Get model if exist
	var attribute_definition_xhtmlDB orm.ATTRIBUTE_DEFINITION_XHTMLDB
	if _, err := db.First(&attribute_definition_xhtmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&attribute_definition_xhtmlDB)

	// get an instance (not staged) from DB instance, and call callback function
	attribute_definition_xhtmlDeleted := new(models.ATTRIBUTE_DEFINITION_XHTML)
	attribute_definition_xhtmlDB.CopyBasicFieldsToATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtmlDeleted)

	// get stage instance from DB instance, and call callback function
	attribute_definition_xhtmlStaged := backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.Map_ATTRIBUTE_DEFINITION_XHTMLDBID_ATTRIBUTE_DEFINITION_XHTMLPtr[attribute_definition_xhtmlDB.ID]
	if attribute_definition_xhtmlStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), attribute_definition_xhtmlStaged, attribute_definition_xhtmlDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
