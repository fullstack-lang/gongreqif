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
var __DATATYPE_DEFINITION_XHTML__dummysDeclaration__ models.DATATYPE_DEFINITION_XHTML
var __DATATYPE_DEFINITION_XHTML_time__dummyDeclaration time.Duration

var mutexDATATYPE_DEFINITION_XHTML sync.Mutex

// An DATATYPE_DEFINITION_XHTMLID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDATATYPE_DEFINITION_XHTML updateDATATYPE_DEFINITION_XHTML deleteDATATYPE_DEFINITION_XHTML
type DATATYPE_DEFINITION_XHTMLID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DATATYPE_DEFINITION_XHTMLInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDATATYPE_DEFINITION_XHTML updateDATATYPE_DEFINITION_XHTML
type DATATYPE_DEFINITION_XHTMLInput struct {
	// The DATATYPE_DEFINITION_XHTML to submit or modify
	// in: body
	DATATYPE_DEFINITION_XHTML *orm.DATATYPE_DEFINITION_XHTMLAPI
}

// GetDATATYPE_DEFINITION_XHTMLs
//
// swagger:route GET /datatype_definition_xhtmls datatype_definition_xhtmls getDATATYPE_DEFINITION_XHTMLs
//
// # Get all datatype_definition_xhtmls
//
// Responses:
// default: genericError
//
//	200: datatype_definition_xhtmlDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_XHTMLs(c *gin.Context) {

	// source slice
	var datatype_definition_xhtmlDBs []orm.DATATYPE_DEFINITION_XHTMLDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_XHTMLs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.GetDB()

	_, err := db.Find(&datatype_definition_xhtmlDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	datatype_definition_xhtmlAPIs := make([]orm.DATATYPE_DEFINITION_XHTMLAPI, 0)

	// for each datatype_definition_xhtml, update fields from the database nullable fields
	for idx := range datatype_definition_xhtmlDBs {
		datatype_definition_xhtmlDB := &datatype_definition_xhtmlDBs[idx]
		_ = datatype_definition_xhtmlDB
		var datatype_definition_xhtmlAPI orm.DATATYPE_DEFINITION_XHTMLAPI

		// insertion point for updating fields
		datatype_definition_xhtmlAPI.ID = datatype_definition_xhtmlDB.ID
		datatype_definition_xhtmlDB.CopyBasicFieldsToDATATYPE_DEFINITION_XHTML_WOP(&datatype_definition_xhtmlAPI.DATATYPE_DEFINITION_XHTML_WOP)
		datatype_definition_xhtmlAPI.DATATYPE_DEFINITION_XHTMLPointersEncoding = datatype_definition_xhtmlDB.DATATYPE_DEFINITION_XHTMLPointersEncoding
		datatype_definition_xhtmlAPIs = append(datatype_definition_xhtmlAPIs, datatype_definition_xhtmlAPI)
	}

	c.JSON(http.StatusOK, datatype_definition_xhtmlAPIs)
}

// PostDATATYPE_DEFINITION_XHTML
//
// swagger:route POST /datatype_definition_xhtmls datatype_definition_xhtmls postDATATYPE_DEFINITION_XHTML
//
// Creates a datatype_definition_xhtml
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostDATATYPE_DEFINITION_XHTML(c *gin.Context) {

	mutexDATATYPE_DEFINITION_XHTML.Lock()
	defer mutexDATATYPE_DEFINITION_XHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostDATATYPE_DEFINITION_XHTMLs", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_XHTMLAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create datatype_definition_xhtml
	datatype_definition_xhtmlDB := orm.DATATYPE_DEFINITION_XHTMLDB{}
	datatype_definition_xhtmlDB.DATATYPE_DEFINITION_XHTMLPointersEncoding = input.DATATYPE_DEFINITION_XHTMLPointersEncoding
	datatype_definition_xhtmlDB.CopyBasicFieldsFromDATATYPE_DEFINITION_XHTML_WOP(&input.DATATYPE_DEFINITION_XHTML_WOP)

	_, err = db.Create(&datatype_definition_xhtmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.CheckoutPhaseOneInstance(&datatype_definition_xhtmlDB)
	datatype_definition_xhtml := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.Map_DATATYPE_DEFINITION_XHTMLDBID_DATATYPE_DEFINITION_XHTMLPtr[datatype_definition_xhtmlDB.ID]

	if datatype_definition_xhtml != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), datatype_definition_xhtml)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, datatype_definition_xhtmlDB)
}

// GetDATATYPE_DEFINITION_XHTML
//
// swagger:route GET /datatype_definition_xhtmls/{ID} datatype_definition_xhtmls getDATATYPE_DEFINITION_XHTML
//
// Gets the details for a datatype_definition_xhtml.
//
// Responses:
// default: genericError
//
//	200: datatype_definition_xhtmlDBResponse
func (controller *Controller) GetDATATYPE_DEFINITION_XHTML(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetDATATYPE_DEFINITION_XHTML", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.GetDB()

	// Get datatype_definition_xhtmlDB in DB
	var datatype_definition_xhtmlDB orm.DATATYPE_DEFINITION_XHTMLDB
	if _, err := db.First(&datatype_definition_xhtmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var datatype_definition_xhtmlAPI orm.DATATYPE_DEFINITION_XHTMLAPI
	datatype_definition_xhtmlAPI.ID = datatype_definition_xhtmlDB.ID
	datatype_definition_xhtmlAPI.DATATYPE_DEFINITION_XHTMLPointersEncoding = datatype_definition_xhtmlDB.DATATYPE_DEFINITION_XHTMLPointersEncoding
	datatype_definition_xhtmlDB.CopyBasicFieldsToDATATYPE_DEFINITION_XHTML_WOP(&datatype_definition_xhtmlAPI.DATATYPE_DEFINITION_XHTML_WOP)

	c.JSON(http.StatusOK, datatype_definition_xhtmlAPI)
}

// UpdateDATATYPE_DEFINITION_XHTML
//
// swagger:route PATCH /datatype_definition_xhtmls/{ID} datatype_definition_xhtmls updateDATATYPE_DEFINITION_XHTML
//
// # Update a datatype_definition_xhtml
//
// Responses:
// default: genericError
//
//	200: datatype_definition_xhtmlDBResponse
func (controller *Controller) UpdateDATATYPE_DEFINITION_XHTML(c *gin.Context) {

	mutexDATATYPE_DEFINITION_XHTML.Lock()
	defer mutexDATATYPE_DEFINITION_XHTML.Unlock()

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
	db := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.GetDB()

	// Validate input
	var input orm.DATATYPE_DEFINITION_XHTMLAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var datatype_definition_xhtmlDB orm.DATATYPE_DEFINITION_XHTMLDB

	// fetch the datatype_definition_xhtml
	_, err := db.First(&datatype_definition_xhtmlDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	datatype_definition_xhtmlDB.CopyBasicFieldsFromDATATYPE_DEFINITION_XHTML_WOP(&input.DATATYPE_DEFINITION_XHTML_WOP)
	datatype_definition_xhtmlDB.DATATYPE_DEFINITION_XHTMLPointersEncoding = input.DATATYPE_DEFINITION_XHTMLPointersEncoding

	db, _ = db.Model(&datatype_definition_xhtmlDB)
	_, err = db.Updates(&datatype_definition_xhtmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_xhtmlNew := new(models.DATATYPE_DEFINITION_XHTML)
	datatype_definition_xhtmlDB.CopyBasicFieldsToDATATYPE_DEFINITION_XHTML(datatype_definition_xhtmlNew)

	// redeem pointers
	datatype_definition_xhtmlDB.DecodePointers(backRepo, datatype_definition_xhtmlNew)

	// get stage instance from DB instance, and call callback function
	datatype_definition_xhtmlOld := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.Map_DATATYPE_DEFINITION_XHTMLDBID_DATATYPE_DEFINITION_XHTMLPtr[datatype_definition_xhtmlDB.ID]
	if datatype_definition_xhtmlOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), datatype_definition_xhtmlOld, datatype_definition_xhtmlNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the datatype_definition_xhtmlDB
	c.JSON(http.StatusOK, datatype_definition_xhtmlDB)
}

// DeleteDATATYPE_DEFINITION_XHTML
//
// swagger:route DELETE /datatype_definition_xhtmls/{ID} datatype_definition_xhtmls deleteDATATYPE_DEFINITION_XHTML
//
// # Delete a datatype_definition_xhtml
//
// default: genericError
//
//	200: datatype_definition_xhtmlDBResponse
func (controller *Controller) DeleteDATATYPE_DEFINITION_XHTML(c *gin.Context) {

	mutexDATATYPE_DEFINITION_XHTML.Lock()
	defer mutexDATATYPE_DEFINITION_XHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteDATATYPE_DEFINITION_XHTML", "Name", stackPath)
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
	db := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.GetDB()

	// Get model if exist
	var datatype_definition_xhtmlDB orm.DATATYPE_DEFINITION_XHTMLDB
	if _, err := db.First(&datatype_definition_xhtmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&datatype_definition_xhtmlDB)

	// get an instance (not staged) from DB instance, and call callback function
	datatype_definition_xhtmlDeleted := new(models.DATATYPE_DEFINITION_XHTML)
	datatype_definition_xhtmlDB.CopyBasicFieldsToDATATYPE_DEFINITION_XHTML(datatype_definition_xhtmlDeleted)

	// get stage instance from DB instance, and call callback function
	datatype_definition_xhtmlStaged := backRepo.BackRepoDATATYPE_DEFINITION_XHTML.Map_DATATYPE_DEFINITION_XHTMLDBID_DATATYPE_DEFINITION_XHTMLPtr[datatype_definition_xhtmlDB.ID]
	if datatype_definition_xhtmlStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), datatype_definition_xhtmlStaged, datatype_definition_xhtmlDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
