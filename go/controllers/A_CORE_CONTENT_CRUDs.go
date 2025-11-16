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
var __A_CORE_CONTENT__dummysDeclaration__ models.A_CORE_CONTENT
var __A_CORE_CONTENT_time__dummyDeclaration time.Duration

var mutexA_CORE_CONTENT sync.Mutex

// An A_CORE_CONTENTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_CORE_CONTENT updateA_CORE_CONTENT deleteA_CORE_CONTENT
type A_CORE_CONTENTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_CORE_CONTENTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_CORE_CONTENT updateA_CORE_CONTENT
type A_CORE_CONTENTInput struct {
	// The A_CORE_CONTENT to submit or modify
	// in: body
	A_CORE_CONTENT *orm.A_CORE_CONTENTAPI
}

// GetA_CORE_CONTENTs
//
// swagger:route GET /a_core_contents a_core_contents getA_CORE_CONTENTs
//
// # Get all a_core_contents
//
// Responses:
// default: genericError
//
//	200: a_core_contentDBResponse
func (controller *Controller) GetA_CORE_CONTENTs(c *gin.Context) {

	// source slice
	var a_core_contentDBs []orm.A_CORE_CONTENTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_CORE_CONTENTs", "Name", stackPath)
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
	db := backRepo.BackRepoA_CORE_CONTENT.GetDB()

	_, err := db.Find(&a_core_contentDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_core_contentAPIs := make([]orm.A_CORE_CONTENTAPI, 0)

	// for each a_core_content, update fields from the database nullable fields
	for idx := range a_core_contentDBs {
		a_core_contentDB := &a_core_contentDBs[idx]
		_ = a_core_contentDB
		var a_core_contentAPI orm.A_CORE_CONTENTAPI

		// insertion point for updating fields
		a_core_contentAPI.ID = a_core_contentDB.ID
		a_core_contentDB.CopyBasicFieldsToA_CORE_CONTENT_WOP(&a_core_contentAPI.A_CORE_CONTENT_WOP)
		a_core_contentAPI.A_CORE_CONTENTPointersEncoding = a_core_contentDB.A_CORE_CONTENTPointersEncoding
		a_core_contentAPIs = append(a_core_contentAPIs, a_core_contentAPI)
	}

	c.JSON(http.StatusOK, a_core_contentAPIs)
}

// PostA_CORE_CONTENT
//
// swagger:route POST /a_core_contents a_core_contents postA_CORE_CONTENT
//
// Creates a a_core_content
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_CORE_CONTENT(c *gin.Context) {

	mutexA_CORE_CONTENT.Lock()
	defer mutexA_CORE_CONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_CORE_CONTENTs", "Name", stackPath)
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
	db := backRepo.BackRepoA_CORE_CONTENT.GetDB()

	// Validate input
	var input orm.A_CORE_CONTENTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_core_content
	a_core_contentDB := orm.A_CORE_CONTENTDB{}
	a_core_contentDB.A_CORE_CONTENTPointersEncoding = input.A_CORE_CONTENTPointersEncoding
	a_core_contentDB.CopyBasicFieldsFromA_CORE_CONTENT_WOP(&input.A_CORE_CONTENT_WOP)

	_, err = db.Create(&a_core_contentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_CORE_CONTENT.CheckoutPhaseOneInstance(&a_core_contentDB)
	a_core_content := backRepo.BackRepoA_CORE_CONTENT.Map_A_CORE_CONTENTDBID_A_CORE_CONTENTPtr[a_core_contentDB.ID]

	if a_core_content != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_core_content)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_core_contentDB)
}

// GetA_CORE_CONTENT
//
// swagger:route GET /a_core_contents/{ID} a_core_contents getA_CORE_CONTENT
//
// Gets the details for a a_core_content.
//
// Responses:
// default: genericError
//
//	200: a_core_contentDBResponse
func (controller *Controller) GetA_CORE_CONTENT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_CORE_CONTENT", "Name", stackPath)
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
	db := backRepo.BackRepoA_CORE_CONTENT.GetDB()

	// Get a_core_contentDB in DB
	var a_core_contentDB orm.A_CORE_CONTENTDB
	if _, err := db.First(&a_core_contentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_core_contentAPI orm.A_CORE_CONTENTAPI
	a_core_contentAPI.ID = a_core_contentDB.ID
	a_core_contentAPI.A_CORE_CONTENTPointersEncoding = a_core_contentDB.A_CORE_CONTENTPointersEncoding
	a_core_contentDB.CopyBasicFieldsToA_CORE_CONTENT_WOP(&a_core_contentAPI.A_CORE_CONTENT_WOP)

	c.JSON(http.StatusOK, a_core_contentAPI)
}

// UpdateA_CORE_CONTENT
//
// swagger:route PATCH /a_core_contents/{ID} a_core_contents updateA_CORE_CONTENT
//
// # Update a a_core_content
//
// Responses:
// default: genericError
//
//	200: a_core_contentDBResponse
func (controller *Controller) UpdateA_CORE_CONTENT(c *gin.Context) {

	mutexA_CORE_CONTENT.Lock()
	defer mutexA_CORE_CONTENT.Unlock()

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
	db := backRepo.BackRepoA_CORE_CONTENT.GetDB()

	// Validate input
	var input orm.A_CORE_CONTENTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_core_contentDB orm.A_CORE_CONTENTDB

	// fetch the a_core_content
	_, err := db.First(&a_core_contentDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_core_contentDB.CopyBasicFieldsFromA_CORE_CONTENT_WOP(&input.A_CORE_CONTENT_WOP)
	a_core_contentDB.A_CORE_CONTENTPointersEncoding = input.A_CORE_CONTENTPointersEncoding

	db, _ = db.Model(&a_core_contentDB)
	_, err = db.Updates(&a_core_contentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_core_contentNew := new(models.A_CORE_CONTENT)
	a_core_contentDB.CopyBasicFieldsToA_CORE_CONTENT(a_core_contentNew)

	// redeem pointers
	a_core_contentDB.DecodePointers(backRepo, a_core_contentNew)

	// get stage instance from DB instance, and call callback function
	a_core_contentOld := backRepo.BackRepoA_CORE_CONTENT.Map_A_CORE_CONTENTDBID_A_CORE_CONTENTPtr[a_core_contentDB.ID]
	if a_core_contentOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_core_contentOld, a_core_contentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_core_contentDB
	c.JSON(http.StatusOK, a_core_contentDB)
}

// DeleteA_CORE_CONTENT
//
// swagger:route DELETE /a_core_contents/{ID} a_core_contents deleteA_CORE_CONTENT
//
// # Delete a a_core_content
//
// default: genericError
//
//	200: a_core_contentDBResponse
func (controller *Controller) DeleteA_CORE_CONTENT(c *gin.Context) {

	mutexA_CORE_CONTENT.Lock()
	defer mutexA_CORE_CONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_CORE_CONTENT", "Name", stackPath)
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
	db := backRepo.BackRepoA_CORE_CONTENT.GetDB()

	// Get model if exist
	var a_core_contentDB orm.A_CORE_CONTENTDB
	if _, err := db.First(&a_core_contentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_core_contentDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_core_contentDeleted := new(models.A_CORE_CONTENT)
	a_core_contentDB.CopyBasicFieldsToA_CORE_CONTENT(a_core_contentDeleted)

	// get stage instance from DB instance, and call callback function
	a_core_contentStaged := backRepo.BackRepoA_CORE_CONTENT.Map_A_CORE_CONTENTDBID_A_CORE_CONTENTPtr[a_core_contentDB.ID]
	if a_core_contentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_core_contentStaged, a_core_contentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
