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
var __REQ_IF_CONTENT__dummysDeclaration__ models.REQ_IF_CONTENT
var __REQ_IF_CONTENT_time__dummyDeclaration time.Duration

var mutexREQ_IF_CONTENT sync.Mutex

// An REQ_IF_CONTENTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQ_IF_CONTENT updateREQ_IF_CONTENT deleteREQ_IF_CONTENT
type REQ_IF_CONTENTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQ_IF_CONTENTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQ_IF_CONTENT updateREQ_IF_CONTENT
type REQ_IF_CONTENTInput struct {
	// The REQ_IF_CONTENT to submit or modify
	// in: body
	REQ_IF_CONTENT *orm.REQ_IF_CONTENTAPI
}

// GetREQ_IF_CONTENTs
//
// swagger:route GET /req_if_contents req_if_contents getREQ_IF_CONTENTs
//
// # Get all req_if_contents
//
// Responses:
// default: genericError
//
//	200: req_if_contentDBResponse
func (controller *Controller) GetREQ_IF_CONTENTs(c *gin.Context) {

	// source slice
	var req_if_contentDBs []orm.REQ_IF_CONTENTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQ_IF_CONTENTs", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF_CONTENT.GetDB()

	_, err := db.Find(&req_if_contentDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	req_if_contentAPIs := make([]orm.REQ_IF_CONTENTAPI, 0)

	// for each req_if_content, update fields from the database nullable fields
	for idx := range req_if_contentDBs {
		req_if_contentDB := &req_if_contentDBs[idx]
		_ = req_if_contentDB
		var req_if_contentAPI orm.REQ_IF_CONTENTAPI

		// insertion point for updating fields
		req_if_contentAPI.ID = req_if_contentDB.ID
		req_if_contentDB.CopyBasicFieldsToREQ_IF_CONTENT_WOP(&req_if_contentAPI.REQ_IF_CONTENT_WOP)
		req_if_contentAPI.REQ_IF_CONTENTPointersEncoding = req_if_contentDB.REQ_IF_CONTENTPointersEncoding
		req_if_contentAPIs = append(req_if_contentAPIs, req_if_contentAPI)
	}

	c.JSON(http.StatusOK, req_if_contentAPIs)
}

// PostREQ_IF_CONTENT
//
// swagger:route POST /req_if_contents req_if_contents postREQ_IF_CONTENT
//
// Creates a req_if_content
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQ_IF_CONTENT(c *gin.Context) {

	mutexREQ_IF_CONTENT.Lock()
	defer mutexREQ_IF_CONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQ_IF_CONTENTs", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF_CONTENT.GetDB()

	// Validate input
	var input orm.REQ_IF_CONTENTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create req_if_content
	req_if_contentDB := orm.REQ_IF_CONTENTDB{}
	req_if_contentDB.REQ_IF_CONTENTPointersEncoding = input.REQ_IF_CONTENTPointersEncoding
	req_if_contentDB.CopyBasicFieldsFromREQ_IF_CONTENT_WOP(&input.REQ_IF_CONTENT_WOP)

	_, err = db.Create(&req_if_contentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQ_IF_CONTENT.CheckoutPhaseOneInstance(&req_if_contentDB)
	req_if_content := backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[req_if_contentDB.ID]

	if req_if_content != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), req_if_content)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, req_if_contentDB)
}

// GetREQ_IF_CONTENT
//
// swagger:route GET /req_if_contents/{ID} req_if_contents getREQ_IF_CONTENT
//
// Gets the details for a req_if_content.
//
// Responses:
// default: genericError
//
//	200: req_if_contentDBResponse
func (controller *Controller) GetREQ_IF_CONTENT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQ_IF_CONTENT", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF_CONTENT.GetDB()

	// Get req_if_contentDB in DB
	var req_if_contentDB orm.REQ_IF_CONTENTDB
	if _, err := db.First(&req_if_contentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var req_if_contentAPI orm.REQ_IF_CONTENTAPI
	req_if_contentAPI.ID = req_if_contentDB.ID
	req_if_contentAPI.REQ_IF_CONTENTPointersEncoding = req_if_contentDB.REQ_IF_CONTENTPointersEncoding
	req_if_contentDB.CopyBasicFieldsToREQ_IF_CONTENT_WOP(&req_if_contentAPI.REQ_IF_CONTENT_WOP)

	c.JSON(http.StatusOK, req_if_contentAPI)
}

// UpdateREQ_IF_CONTENT
//
// swagger:route PATCH /req_if_contents/{ID} req_if_contents updateREQ_IF_CONTENT
//
// # Update a req_if_content
//
// Responses:
// default: genericError
//
//	200: req_if_contentDBResponse
func (controller *Controller) UpdateREQ_IF_CONTENT(c *gin.Context) {

	mutexREQ_IF_CONTENT.Lock()
	defer mutexREQ_IF_CONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateREQ_IF_CONTENT", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF_CONTENT.GetDB()

	// Validate input
	var input orm.REQ_IF_CONTENTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var req_if_contentDB orm.REQ_IF_CONTENTDB

	// fetch the req_if_content
	_, err := db.First(&req_if_contentDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	req_if_contentDB.CopyBasicFieldsFromREQ_IF_CONTENT_WOP(&input.REQ_IF_CONTENT_WOP)
	req_if_contentDB.REQ_IF_CONTENTPointersEncoding = input.REQ_IF_CONTENTPointersEncoding

	db, _ = db.Model(&req_if_contentDB)
	_, err = db.Updates(&req_if_contentDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	req_if_contentNew := new(models.REQ_IF_CONTENT)
	req_if_contentDB.CopyBasicFieldsToREQ_IF_CONTENT(req_if_contentNew)

	// redeem pointers
	req_if_contentDB.DecodePointers(backRepo, req_if_contentNew)

	// get stage instance from DB instance, and call callback function
	req_if_contentOld := backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[req_if_contentDB.ID]
	if req_if_contentOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), req_if_contentOld, req_if_contentNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the req_if_contentDB
	c.JSON(http.StatusOK, req_if_contentDB)
}

// DeleteREQ_IF_CONTENT
//
// swagger:route DELETE /req_if_contents/{ID} req_if_contents deleteREQ_IF_CONTENT
//
// # Delete a req_if_content
//
// default: genericError
//
//	200: req_if_contentDBResponse
func (controller *Controller) DeleteREQ_IF_CONTENT(c *gin.Context) {

	mutexREQ_IF_CONTENT.Lock()
	defer mutexREQ_IF_CONTENT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQ_IF_CONTENT", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF_CONTENT.GetDB()

	// Get model if exist
	var req_if_contentDB orm.REQ_IF_CONTENTDB
	if _, err := db.First(&req_if_contentDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&req_if_contentDB)

	// get an instance (not staged) from DB instance, and call callback function
	req_if_contentDeleted := new(models.REQ_IF_CONTENT)
	req_if_contentDB.CopyBasicFieldsToREQ_IF_CONTENT(req_if_contentDeleted)

	// get stage instance from DB instance, and call callback function
	req_if_contentStaged := backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[req_if_contentDB.ID]
	if req_if_contentStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), req_if_contentStaged, req_if_contentDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
