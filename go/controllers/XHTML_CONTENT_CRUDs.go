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
var __XHTML_CONTENT__dummysDeclaration__ models.XHTML_CONTENT
var __XHTML_CONTENT_time__dummyDeclaration time.Duration

var mutexXHTML_CONTENT sync.Mutex

// An XHTML_CONTENTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXHTML_CONTENT updateXHTML_CONTENT deleteXHTML_CONTENT
type XHTML_CONTENTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// XHTML_CONTENTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXHTML_CONTENT updateXHTML_CONTENT
type XHTML_CONTENTInput struct {
	// The XHTML_CONTENT to submit or modify
	// in: body
	XHTML_CONTENT *orm.XHTML_CONTENTAPI
}

// GetXHTML_CONTENTs
//
// swagger:route GET /xhtml_contents xhtml_contents getXHTML_CONTENTs
//
// # Get all xhtml_contents
//
// Responses:
// default: genericError
//
//	200: xhtml_contentDBResponse
func (controller *Controller) GetXHTML_CONTENTs(c *gin.Context) {

	// source slice
	var xhtml_contentDBs []orm.XHTML_CONTENTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXHTML_CONTENTs", "Name", stackPath)
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
	db := backRepo.BackRepoXHTML_CONTENT.GetDB()

	_, err := db.Find(&xhtml_contentDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_contentAPIs := make([]orm.XHTML_CONTENTAPI, 0)

	// for each xhtml_content, update fields from the database nullable fields
	for idx := range xhtml_contentDBs {
		xhtml_contentDB := &xhtml_contentDBs[idx]
		_ = xhtml_contentDB
		var xhtml_contentAPI orm.XHTML_CONTENTAPI

		// insertion point for updating fields
		xhtml_contentAPI.ID = xhtml_contentDB.ID
		xhtml_contentDB.CopyBasicFieldsToXHTML_CONTENT_WOP(&xhtml_contentAPI.XHTML_CONTENT_WOP)
		xhtml_contentAPI.XHTML_CONTENTPointersEncoding = xhtml_contentDB.XHTML_CONTENTPointersEncoding
		xhtml_contentAPIs = append(xhtml_contentAPIs, xhtml_contentAPI)
	}

	c.JSON(http.StatusOK, xhtml_contentAPIs)
}

// PostXHTML_CONTENT
//
// swagger:route POST /xhtml_contents xhtml_contents postXHTML_CONTENT
//
// Creates a xhtml_content
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXHTML_CONTENT(c *gin.Context) {

	mutexXHTML_CONTENT.Lock()
	defer mutexXHTML_CONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXHTML_CONTENTs", "Name", stackPath)
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
	db := backRepo.BackRepoXHTML_CONTENT.GetDB()

	// Validate input
	var input orm.XHTML_CONTENTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_content
	xhtml_contentDB := orm.XHTML_CONTENTDB{}
	xhtml_contentDB.XHTML_CONTENTPointersEncoding = input.XHTML_CONTENTPointersEncoding
	xhtml_contentDB.CopyBasicFieldsFromXHTML_CONTENT_WOP(&input.XHTML_CONTENT_WOP)

	_, err = db.Create(&xhtml_contentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXHTML_CONTENT.CheckoutPhaseOneInstance(&xhtml_contentDB)
	xhtml_content := backRepo.BackRepoXHTML_CONTENT.Map_XHTML_CONTENTDBID_XHTML_CONTENTPtr[xhtml_contentDB.ID]

	if xhtml_content != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_content)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_contentDB)
}

// GetXHTML_CONTENT
//
// swagger:route GET /xhtml_contents/{ID} xhtml_contents getXHTML_CONTENT
//
// Gets the details for a xhtml_content.
//
// Responses:
// default: genericError
//
//	200: xhtml_contentDBResponse
func (controller *Controller) GetXHTML_CONTENT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXHTML_CONTENT", "Name", stackPath)
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
	db := backRepo.BackRepoXHTML_CONTENT.GetDB()

	// Get xhtml_contentDB in DB
	var xhtml_contentDB orm.XHTML_CONTENTDB
	if _, err := db.First(&xhtml_contentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_contentAPI orm.XHTML_CONTENTAPI
	xhtml_contentAPI.ID = xhtml_contentDB.ID
	xhtml_contentAPI.XHTML_CONTENTPointersEncoding = xhtml_contentDB.XHTML_CONTENTPointersEncoding
	xhtml_contentDB.CopyBasicFieldsToXHTML_CONTENT_WOP(&xhtml_contentAPI.XHTML_CONTENT_WOP)

	c.JSON(http.StatusOK, xhtml_contentAPI)
}

// UpdateXHTML_CONTENT
//
// swagger:route PATCH /xhtml_contents/{ID} xhtml_contents updateXHTML_CONTENT
//
// # Update a xhtml_content
//
// Responses:
// default: genericError
//
//	200: xhtml_contentDBResponse
func (controller *Controller) UpdateXHTML_CONTENT(c *gin.Context) {

	mutexXHTML_CONTENT.Lock()
	defer mutexXHTML_CONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXHTML_CONTENT", "Name", stackPath)
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
	db := backRepo.BackRepoXHTML_CONTENT.GetDB()

	// Validate input
	var input orm.XHTML_CONTENTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_contentDB orm.XHTML_CONTENTDB

	// fetch the xhtml_content
	_, err := db.First(&xhtml_contentDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_contentDB.CopyBasicFieldsFromXHTML_CONTENT_WOP(&input.XHTML_CONTENT_WOP)
	xhtml_contentDB.XHTML_CONTENTPointersEncoding = input.XHTML_CONTENTPointersEncoding

	db, _ = db.Model(&xhtml_contentDB)
	_, err = db.Updates(&xhtml_contentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_contentNew := new(models.XHTML_CONTENT)
	xhtml_contentDB.CopyBasicFieldsToXHTML_CONTENT(xhtml_contentNew)

	// redeem pointers
	xhtml_contentDB.DecodePointers(backRepo, xhtml_contentNew)

	// get stage instance from DB instance, and call callback function
	xhtml_contentOld := backRepo.BackRepoXHTML_CONTENT.Map_XHTML_CONTENTDBID_XHTML_CONTENTPtr[xhtml_contentDB.ID]
	if xhtml_contentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_contentOld, xhtml_contentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_contentDB
	c.JSON(http.StatusOK, xhtml_contentDB)
}

// DeleteXHTML_CONTENT
//
// swagger:route DELETE /xhtml_contents/{ID} xhtml_contents deleteXHTML_CONTENT
//
// # Delete a xhtml_content
//
// default: genericError
//
//	200: xhtml_contentDBResponse
func (controller *Controller) DeleteXHTML_CONTENT(c *gin.Context) {

	mutexXHTML_CONTENT.Lock()
	defer mutexXHTML_CONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXHTML_CONTENT", "Name", stackPath)
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
	db := backRepo.BackRepoXHTML_CONTENT.GetDB()

	// Get model if exist
	var xhtml_contentDB orm.XHTML_CONTENTDB
	if _, err := db.First(&xhtml_contentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_contentDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_contentDeleted := new(models.XHTML_CONTENT)
	xhtml_contentDB.CopyBasicFieldsToXHTML_CONTENT(xhtml_contentDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_contentStaged := backRepo.BackRepoXHTML_CONTENT.Map_XHTML_CONTENTDBID_XHTML_CONTENTPtr[xhtml_contentDB.ID]
	if xhtml_contentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_contentStaged, xhtml_contentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
