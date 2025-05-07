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
var __A_ATTRIBUTE_VALUE_XHTML__dummysDeclaration__ models.A_ATTRIBUTE_VALUE_XHTML
var __A_ATTRIBUTE_VALUE_XHTML_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_VALUE_XHTML sync.Mutex

// An A_ATTRIBUTE_VALUE_XHTMLID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_VALUE_XHTML updateA_ATTRIBUTE_VALUE_XHTML deleteA_ATTRIBUTE_VALUE_XHTML
type A_ATTRIBUTE_VALUE_XHTMLID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_VALUE_XHTMLInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_VALUE_XHTML updateA_ATTRIBUTE_VALUE_XHTML
type A_ATTRIBUTE_VALUE_XHTMLInput struct {
	// The A_ATTRIBUTE_VALUE_XHTML to submit or modify
	// in: body
	A_ATTRIBUTE_VALUE_XHTML *orm.A_ATTRIBUTE_VALUE_XHTMLAPI
}

// GetA_ATTRIBUTE_VALUE_XHTMLs
//
// swagger:route GET /a_attribute_value_xhtmls a_attribute_value_xhtmls getA_ATTRIBUTE_VALUE_XHTMLs
//
// # Get all a_attribute_value_xhtmls
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_xhtmlDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_XHTMLs(c *gin.Context) {

	// source slice
	var a_attribute_value_xhtmlDBs []orm.A_ATTRIBUTE_VALUE_XHTMLDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_XHTMLs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.GetDB()

	_, err := db.Find(&a_attribute_value_xhtmlDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_value_xhtmlAPIs := make([]orm.A_ATTRIBUTE_VALUE_XHTMLAPI, 0)

	// for each a_attribute_value_xhtml, update fields from the database nullable fields
	for idx := range a_attribute_value_xhtmlDBs {
		a_attribute_value_xhtmlDB := &a_attribute_value_xhtmlDBs[idx]
		_ = a_attribute_value_xhtmlDB
		var a_attribute_value_xhtmlAPI orm.A_ATTRIBUTE_VALUE_XHTMLAPI

		// insertion point for updating fields
		a_attribute_value_xhtmlAPI.ID = a_attribute_value_xhtmlDB.ID
		a_attribute_value_xhtmlDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_XHTML_WOP(&a_attribute_value_xhtmlAPI.A_ATTRIBUTE_VALUE_XHTML_WOP)
		a_attribute_value_xhtmlAPI.A_ATTRIBUTE_VALUE_XHTMLPointersEncoding = a_attribute_value_xhtmlDB.A_ATTRIBUTE_VALUE_XHTMLPointersEncoding
		a_attribute_value_xhtmlAPIs = append(a_attribute_value_xhtmlAPIs, a_attribute_value_xhtmlAPI)
	}

	c.JSON(http.StatusOK, a_attribute_value_xhtmlAPIs)
}

// PostA_ATTRIBUTE_VALUE_XHTML
//
// swagger:route POST /a_attribute_value_xhtmls a_attribute_value_xhtmls postA_ATTRIBUTE_VALUE_XHTML
//
// Creates a a_attribute_value_xhtml
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_VALUE_XHTML(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_XHTML.Lock()
	defer mutexA_ATTRIBUTE_VALUE_XHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_VALUE_XHTMLs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_XHTMLAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_value_xhtml
	a_attribute_value_xhtmlDB := orm.A_ATTRIBUTE_VALUE_XHTMLDB{}
	a_attribute_value_xhtmlDB.A_ATTRIBUTE_VALUE_XHTMLPointersEncoding = input.A_ATTRIBUTE_VALUE_XHTMLPointersEncoding
	a_attribute_value_xhtmlDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_XHTML_WOP(&input.A_ATTRIBUTE_VALUE_XHTML_WOP)

	_, err = db.Create(&a_attribute_value_xhtmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.CheckoutPhaseOneInstance(&a_attribute_value_xhtmlDB)
	a_attribute_value_xhtml := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.Map_A_ATTRIBUTE_VALUE_XHTMLDBID_A_ATTRIBUTE_VALUE_XHTMLPtr[a_attribute_value_xhtmlDB.ID]

	if a_attribute_value_xhtml != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_value_xhtml)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_value_xhtmlDB)
}

// GetA_ATTRIBUTE_VALUE_XHTML
//
// swagger:route GET /a_attribute_value_xhtmls/{ID} a_attribute_value_xhtmls getA_ATTRIBUTE_VALUE_XHTML
//
// Gets the details for a a_attribute_value_xhtml.
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_xhtmlDBResponse
func (controller *Controller) GetA_ATTRIBUTE_VALUE_XHTML(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_VALUE_XHTML", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.GetDB()

	// Get a_attribute_value_xhtmlDB in DB
	var a_attribute_value_xhtmlDB orm.A_ATTRIBUTE_VALUE_XHTMLDB
	if _, err := db.First(&a_attribute_value_xhtmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_value_xhtmlAPI orm.A_ATTRIBUTE_VALUE_XHTMLAPI
	a_attribute_value_xhtmlAPI.ID = a_attribute_value_xhtmlDB.ID
	a_attribute_value_xhtmlAPI.A_ATTRIBUTE_VALUE_XHTMLPointersEncoding = a_attribute_value_xhtmlDB.A_ATTRIBUTE_VALUE_XHTMLPointersEncoding
	a_attribute_value_xhtmlDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_XHTML_WOP(&a_attribute_value_xhtmlAPI.A_ATTRIBUTE_VALUE_XHTML_WOP)

	c.JSON(http.StatusOK, a_attribute_value_xhtmlAPI)
}

// UpdateA_ATTRIBUTE_VALUE_XHTML
//
// swagger:route PATCH /a_attribute_value_xhtmls/{ID} a_attribute_value_xhtmls updateA_ATTRIBUTE_VALUE_XHTML
//
// # Update a a_attribute_value_xhtml
//
// Responses:
// default: genericError
//
//	200: a_attribute_value_xhtmlDBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_VALUE_XHTML(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_XHTML.Lock()
	defer mutexA_ATTRIBUTE_VALUE_XHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_ATTRIBUTE_VALUE_XHTML", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_VALUE_XHTMLAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_value_xhtmlDB orm.A_ATTRIBUTE_VALUE_XHTMLDB

	// fetch the a_attribute_value_xhtml
	_, err := db.First(&a_attribute_value_xhtmlDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_value_xhtmlDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_XHTML_WOP(&input.A_ATTRIBUTE_VALUE_XHTML_WOP)
	a_attribute_value_xhtmlDB.A_ATTRIBUTE_VALUE_XHTMLPointersEncoding = input.A_ATTRIBUTE_VALUE_XHTMLPointersEncoding

	db, _ = db.Model(&a_attribute_value_xhtmlDB)
	_, err = db.Updates(&a_attribute_value_xhtmlDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_xhtmlNew := new(models.A_ATTRIBUTE_VALUE_XHTML)
	a_attribute_value_xhtmlDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtmlNew)

	// redeem pointers
	a_attribute_value_xhtmlDB.DecodePointers(backRepo, a_attribute_value_xhtmlNew)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_xhtmlOld := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.Map_A_ATTRIBUTE_VALUE_XHTMLDBID_A_ATTRIBUTE_VALUE_XHTMLPtr[a_attribute_value_xhtmlDB.ID]
	if a_attribute_value_xhtmlOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_attribute_value_xhtmlOld, a_attribute_value_xhtmlNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_value_xhtmlDB
	c.JSON(http.StatusOK, a_attribute_value_xhtmlDB)
}

// DeleteA_ATTRIBUTE_VALUE_XHTML
//
// swagger:route DELETE /a_attribute_value_xhtmls/{ID} a_attribute_value_xhtmls deleteA_ATTRIBUTE_VALUE_XHTML
//
// # Delete a a_attribute_value_xhtml
//
// default: genericError
//
//	200: a_attribute_value_xhtmlDBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_VALUE_XHTML(c *gin.Context) {

	mutexA_ATTRIBUTE_VALUE_XHTML.Lock()
	defer mutexA_ATTRIBUTE_VALUE_XHTML.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_VALUE_XHTML", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.GetDB()

	// Get model if exist
	var a_attribute_value_xhtmlDB orm.A_ATTRIBUTE_VALUE_XHTMLDB
	if _, err := db.First(&a_attribute_value_xhtmlDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_value_xhtmlDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_value_xhtmlDeleted := new(models.A_ATTRIBUTE_VALUE_XHTML)
	a_attribute_value_xhtmlDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtmlDeleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_value_xhtmlStaged := backRepo.BackRepoA_ATTRIBUTE_VALUE_XHTML.Map_A_ATTRIBUTE_VALUE_XHTMLDBID_A_ATTRIBUTE_VALUE_XHTMLPtr[a_attribute_value_xhtmlDB.ID]
	if a_attribute_value_xhtmlStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_value_xhtmlStaged, a_attribute_value_xhtmlDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
