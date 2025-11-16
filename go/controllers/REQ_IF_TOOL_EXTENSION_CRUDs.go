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
var __REQ_IF_TOOL_EXTENSION__dummysDeclaration__ models.REQ_IF_TOOL_EXTENSION
var __REQ_IF_TOOL_EXTENSION_time__dummyDeclaration time.Duration

var mutexREQ_IF_TOOL_EXTENSION sync.Mutex

// An REQ_IF_TOOL_EXTENSIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQ_IF_TOOL_EXTENSION updateREQ_IF_TOOL_EXTENSION deleteREQ_IF_TOOL_EXTENSION
type REQ_IF_TOOL_EXTENSIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQ_IF_TOOL_EXTENSIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQ_IF_TOOL_EXTENSION updateREQ_IF_TOOL_EXTENSION
type REQ_IF_TOOL_EXTENSIONInput struct {
	// The REQ_IF_TOOL_EXTENSION to submit or modify
	// in: body
	REQ_IF_TOOL_EXTENSION *orm.REQ_IF_TOOL_EXTENSIONAPI
}

// GetREQ_IF_TOOL_EXTENSIONs
//
// swagger:route GET /req_if_tool_extensions req_if_tool_extensions getREQ_IF_TOOL_EXTENSIONs
//
// # Get all req_if_tool_extensions
//
// Responses:
// default: genericError
//
//	200: req_if_tool_extensionDBResponse
func (controller *Controller) GetREQ_IF_TOOL_EXTENSIONs(c *gin.Context) {

	// source slice
	var req_if_tool_extensionDBs []orm.REQ_IF_TOOL_EXTENSIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQ_IF_TOOL_EXTENSIONs", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.GetDB()

	_, err := db.Find(&req_if_tool_extensionDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	req_if_tool_extensionAPIs := make([]orm.REQ_IF_TOOL_EXTENSIONAPI, 0)

	// for each req_if_tool_extension, update fields from the database nullable fields
	for idx := range req_if_tool_extensionDBs {
		req_if_tool_extensionDB := &req_if_tool_extensionDBs[idx]
		_ = req_if_tool_extensionDB
		var req_if_tool_extensionAPI orm.REQ_IF_TOOL_EXTENSIONAPI

		// insertion point for updating fields
		req_if_tool_extensionAPI.ID = req_if_tool_extensionDB.ID
		req_if_tool_extensionDB.CopyBasicFieldsToREQ_IF_TOOL_EXTENSION_WOP(&req_if_tool_extensionAPI.REQ_IF_TOOL_EXTENSION_WOP)
		req_if_tool_extensionAPI.REQ_IF_TOOL_EXTENSIONPointersEncoding = req_if_tool_extensionDB.REQ_IF_TOOL_EXTENSIONPointersEncoding
		req_if_tool_extensionAPIs = append(req_if_tool_extensionAPIs, req_if_tool_extensionAPI)
	}

	c.JSON(http.StatusOK, req_if_tool_extensionAPIs)
}

// PostREQ_IF_TOOL_EXTENSION
//
// swagger:route POST /req_if_tool_extensions req_if_tool_extensions postREQ_IF_TOOL_EXTENSION
//
// Creates a req_if_tool_extension
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQ_IF_TOOL_EXTENSION(c *gin.Context) {

	mutexREQ_IF_TOOL_EXTENSION.Lock()
	defer mutexREQ_IF_TOOL_EXTENSION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQ_IF_TOOL_EXTENSIONs", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.GetDB()

	// Validate input
	var input orm.REQ_IF_TOOL_EXTENSIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create req_if_tool_extension
	req_if_tool_extensionDB := orm.REQ_IF_TOOL_EXTENSIONDB{}
	req_if_tool_extensionDB.REQ_IF_TOOL_EXTENSIONPointersEncoding = input.REQ_IF_TOOL_EXTENSIONPointersEncoding
	req_if_tool_extensionDB.CopyBasicFieldsFromREQ_IF_TOOL_EXTENSION_WOP(&input.REQ_IF_TOOL_EXTENSION_WOP)

	_, err = db.Create(&req_if_tool_extensionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseOneInstance(&req_if_tool_extensionDB)
	req_if_tool_extension := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr[req_if_tool_extensionDB.ID]

	if req_if_tool_extension != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), req_if_tool_extension)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, req_if_tool_extensionDB)
}

// GetREQ_IF_TOOL_EXTENSION
//
// swagger:route GET /req_if_tool_extensions/{ID} req_if_tool_extensions getREQ_IF_TOOL_EXTENSION
//
// Gets the details for a req_if_tool_extension.
//
// Responses:
// default: genericError
//
//	200: req_if_tool_extensionDBResponse
func (controller *Controller) GetREQ_IF_TOOL_EXTENSION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQ_IF_TOOL_EXTENSION", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.GetDB()

	// Get req_if_tool_extensionDB in DB
	var req_if_tool_extensionDB orm.REQ_IF_TOOL_EXTENSIONDB
	if _, err := db.First(&req_if_tool_extensionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var req_if_tool_extensionAPI orm.REQ_IF_TOOL_EXTENSIONAPI
	req_if_tool_extensionAPI.ID = req_if_tool_extensionDB.ID
	req_if_tool_extensionAPI.REQ_IF_TOOL_EXTENSIONPointersEncoding = req_if_tool_extensionDB.REQ_IF_TOOL_EXTENSIONPointersEncoding
	req_if_tool_extensionDB.CopyBasicFieldsToREQ_IF_TOOL_EXTENSION_WOP(&req_if_tool_extensionAPI.REQ_IF_TOOL_EXTENSION_WOP)

	c.JSON(http.StatusOK, req_if_tool_extensionAPI)
}

// UpdateREQ_IF_TOOL_EXTENSION
//
// swagger:route PATCH /req_if_tool_extensions/{ID} req_if_tool_extensions updateREQ_IF_TOOL_EXTENSION
//
// # Update a req_if_tool_extension
//
// Responses:
// default: genericError
//
//	200: req_if_tool_extensionDBResponse
func (controller *Controller) UpdateREQ_IF_TOOL_EXTENSION(c *gin.Context) {

	mutexREQ_IF_TOOL_EXTENSION.Lock()
	defer mutexREQ_IF_TOOL_EXTENSION.Unlock()

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
	db := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.GetDB()

	// Validate input
	var input orm.REQ_IF_TOOL_EXTENSIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var req_if_tool_extensionDB orm.REQ_IF_TOOL_EXTENSIONDB

	// fetch the req_if_tool_extension
	_, err := db.First(&req_if_tool_extensionDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	req_if_tool_extensionDB.CopyBasicFieldsFromREQ_IF_TOOL_EXTENSION_WOP(&input.REQ_IF_TOOL_EXTENSION_WOP)
	req_if_tool_extensionDB.REQ_IF_TOOL_EXTENSIONPointersEncoding = input.REQ_IF_TOOL_EXTENSIONPointersEncoding

	db, _ = db.Model(&req_if_tool_extensionDB)
	_, err = db.Updates(&req_if_tool_extensionDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	req_if_tool_extensionNew := new(models.REQ_IF_TOOL_EXTENSION)
	req_if_tool_extensionDB.CopyBasicFieldsToREQ_IF_TOOL_EXTENSION(req_if_tool_extensionNew)

	// redeem pointers
	req_if_tool_extensionDB.DecodePointers(backRepo, req_if_tool_extensionNew)

	// get stage instance from DB instance, and call callback function
	req_if_tool_extensionOld := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr[req_if_tool_extensionDB.ID]
	if req_if_tool_extensionOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), req_if_tool_extensionOld, req_if_tool_extensionNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the req_if_tool_extensionDB
	c.JSON(http.StatusOK, req_if_tool_extensionDB)
}

// DeleteREQ_IF_TOOL_EXTENSION
//
// swagger:route DELETE /req_if_tool_extensions/{ID} req_if_tool_extensions deleteREQ_IF_TOOL_EXTENSION
//
// # Delete a req_if_tool_extension
//
// default: genericError
//
//	200: req_if_tool_extensionDBResponse
func (controller *Controller) DeleteREQ_IF_TOOL_EXTENSION(c *gin.Context) {

	mutexREQ_IF_TOOL_EXTENSION.Lock()
	defer mutexREQ_IF_TOOL_EXTENSION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQ_IF_TOOL_EXTENSION", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.GetDB()

	// Get model if exist
	var req_if_tool_extensionDB orm.REQ_IF_TOOL_EXTENSIONDB
	if _, err := db.First(&req_if_tool_extensionDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&req_if_tool_extensionDB)

	// get an instance (not staged) from DB instance, and call callback function
	req_if_tool_extensionDeleted := new(models.REQ_IF_TOOL_EXTENSION)
	req_if_tool_extensionDB.CopyBasicFieldsToREQ_IF_TOOL_EXTENSION(req_if_tool_extensionDeleted)

	// get stage instance from DB instance, and call callback function
	req_if_tool_extensionStaged := backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr[req_if_tool_extensionDB.ID]
	if req_if_tool_extensionStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), req_if_tool_extensionStaged, req_if_tool_extensionDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
