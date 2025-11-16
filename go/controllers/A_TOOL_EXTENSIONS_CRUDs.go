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
var __A_TOOL_EXTENSIONS__dummysDeclaration__ models.A_TOOL_EXTENSIONS
var __A_TOOL_EXTENSIONS_time__dummyDeclaration time.Duration

var mutexA_TOOL_EXTENSIONS sync.Mutex

// An A_TOOL_EXTENSIONSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_TOOL_EXTENSIONS updateA_TOOL_EXTENSIONS deleteA_TOOL_EXTENSIONS
type A_TOOL_EXTENSIONSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_TOOL_EXTENSIONSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_TOOL_EXTENSIONS updateA_TOOL_EXTENSIONS
type A_TOOL_EXTENSIONSInput struct {
	// The A_TOOL_EXTENSIONS to submit or modify
	// in: body
	A_TOOL_EXTENSIONS *orm.A_TOOL_EXTENSIONSAPI
}

// GetA_TOOL_EXTENSIONSs
//
// swagger:route GET /a_tool_extensionss a_tool_extensionss getA_TOOL_EXTENSIONSs
//
// # Get all a_tool_extensionss
//
// Responses:
// default: genericError
//
//	200: a_tool_extensionsDBResponse
func (controller *Controller) GetA_TOOL_EXTENSIONSs(c *gin.Context) {

	// source slice
	var a_tool_extensionsDBs []orm.A_TOOL_EXTENSIONSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TOOL_EXTENSIONSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_TOOL_EXTENSIONS.GetDB()

	_, err := db.Find(&a_tool_extensionsDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_tool_extensionsAPIs := make([]orm.A_TOOL_EXTENSIONSAPI, 0)

	// for each a_tool_extensions, update fields from the database nullable fields
	for idx := range a_tool_extensionsDBs {
		a_tool_extensionsDB := &a_tool_extensionsDBs[idx]
		_ = a_tool_extensionsDB
		var a_tool_extensionsAPI orm.A_TOOL_EXTENSIONSAPI

		// insertion point for updating fields
		a_tool_extensionsAPI.ID = a_tool_extensionsDB.ID
		a_tool_extensionsDB.CopyBasicFieldsToA_TOOL_EXTENSIONS_WOP(&a_tool_extensionsAPI.A_TOOL_EXTENSIONS_WOP)
		a_tool_extensionsAPI.A_TOOL_EXTENSIONSPointersEncoding = a_tool_extensionsDB.A_TOOL_EXTENSIONSPointersEncoding
		a_tool_extensionsAPIs = append(a_tool_extensionsAPIs, a_tool_extensionsAPI)
	}

	c.JSON(http.StatusOK, a_tool_extensionsAPIs)
}

// PostA_TOOL_EXTENSIONS
//
// swagger:route POST /a_tool_extensionss a_tool_extensionss postA_TOOL_EXTENSIONS
//
// Creates a a_tool_extensions
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_TOOL_EXTENSIONS(c *gin.Context) {

	mutexA_TOOL_EXTENSIONS.Lock()
	defer mutexA_TOOL_EXTENSIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_TOOL_EXTENSIONSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_TOOL_EXTENSIONS.GetDB()

	// Validate input
	var input orm.A_TOOL_EXTENSIONSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_tool_extensions
	a_tool_extensionsDB := orm.A_TOOL_EXTENSIONSDB{}
	a_tool_extensionsDB.A_TOOL_EXTENSIONSPointersEncoding = input.A_TOOL_EXTENSIONSPointersEncoding
	a_tool_extensionsDB.CopyBasicFieldsFromA_TOOL_EXTENSIONS_WOP(&input.A_TOOL_EXTENSIONS_WOP)

	_, err = db.Create(&a_tool_extensionsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_TOOL_EXTENSIONS.CheckoutPhaseOneInstance(&a_tool_extensionsDB)
	a_tool_extensions := backRepo.BackRepoA_TOOL_EXTENSIONS.Map_A_TOOL_EXTENSIONSDBID_A_TOOL_EXTENSIONSPtr[a_tool_extensionsDB.ID]

	if a_tool_extensions != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_tool_extensions)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_tool_extensionsDB)
}

// GetA_TOOL_EXTENSIONS
//
// swagger:route GET /a_tool_extensionss/{ID} a_tool_extensionss getA_TOOL_EXTENSIONS
//
// Gets the details for a a_tool_extensions.
//
// Responses:
// default: genericError
//
//	200: a_tool_extensionsDBResponse
func (controller *Controller) GetA_TOOL_EXTENSIONS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_TOOL_EXTENSIONS", "Name", stackPath)
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
	db := backRepo.BackRepoA_TOOL_EXTENSIONS.GetDB()

	// Get a_tool_extensionsDB in DB
	var a_tool_extensionsDB orm.A_TOOL_EXTENSIONSDB
	if _, err := db.First(&a_tool_extensionsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_tool_extensionsAPI orm.A_TOOL_EXTENSIONSAPI
	a_tool_extensionsAPI.ID = a_tool_extensionsDB.ID
	a_tool_extensionsAPI.A_TOOL_EXTENSIONSPointersEncoding = a_tool_extensionsDB.A_TOOL_EXTENSIONSPointersEncoding
	a_tool_extensionsDB.CopyBasicFieldsToA_TOOL_EXTENSIONS_WOP(&a_tool_extensionsAPI.A_TOOL_EXTENSIONS_WOP)

	c.JSON(http.StatusOK, a_tool_extensionsAPI)
}

// UpdateA_TOOL_EXTENSIONS
//
// swagger:route PATCH /a_tool_extensionss/{ID} a_tool_extensionss updateA_TOOL_EXTENSIONS
//
// # Update a a_tool_extensions
//
// Responses:
// default: genericError
//
//	200: a_tool_extensionsDBResponse
func (controller *Controller) UpdateA_TOOL_EXTENSIONS(c *gin.Context) {

	mutexA_TOOL_EXTENSIONS.Lock()
	defer mutexA_TOOL_EXTENSIONS.Unlock()

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
	db := backRepo.BackRepoA_TOOL_EXTENSIONS.GetDB()

	// Validate input
	var input orm.A_TOOL_EXTENSIONSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_tool_extensionsDB orm.A_TOOL_EXTENSIONSDB

	// fetch the a_tool_extensions
	_, err := db.First(&a_tool_extensionsDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_tool_extensionsDB.CopyBasicFieldsFromA_TOOL_EXTENSIONS_WOP(&input.A_TOOL_EXTENSIONS_WOP)
	a_tool_extensionsDB.A_TOOL_EXTENSIONSPointersEncoding = input.A_TOOL_EXTENSIONSPointersEncoding

	db, _ = db.Model(&a_tool_extensionsDB)
	_, err = db.Updates(&a_tool_extensionsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_tool_extensionsNew := new(models.A_TOOL_EXTENSIONS)
	a_tool_extensionsDB.CopyBasicFieldsToA_TOOL_EXTENSIONS(a_tool_extensionsNew)

	// redeem pointers
	a_tool_extensionsDB.DecodePointers(backRepo, a_tool_extensionsNew)

	// get stage instance from DB instance, and call callback function
	a_tool_extensionsOld := backRepo.BackRepoA_TOOL_EXTENSIONS.Map_A_TOOL_EXTENSIONSDBID_A_TOOL_EXTENSIONSPtr[a_tool_extensionsDB.ID]
	if a_tool_extensionsOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_tool_extensionsOld, a_tool_extensionsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_tool_extensionsDB
	c.JSON(http.StatusOK, a_tool_extensionsDB)
}

// DeleteA_TOOL_EXTENSIONS
//
// swagger:route DELETE /a_tool_extensionss/{ID} a_tool_extensionss deleteA_TOOL_EXTENSIONS
//
// # Delete a a_tool_extensions
//
// default: genericError
//
//	200: a_tool_extensionsDBResponse
func (controller *Controller) DeleteA_TOOL_EXTENSIONS(c *gin.Context) {

	mutexA_TOOL_EXTENSIONS.Lock()
	defer mutexA_TOOL_EXTENSIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_TOOL_EXTENSIONS", "Name", stackPath)
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
	db := backRepo.BackRepoA_TOOL_EXTENSIONS.GetDB()

	// Get model if exist
	var a_tool_extensionsDB orm.A_TOOL_EXTENSIONSDB
	if _, err := db.First(&a_tool_extensionsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_tool_extensionsDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_tool_extensionsDeleted := new(models.A_TOOL_EXTENSIONS)
	a_tool_extensionsDB.CopyBasicFieldsToA_TOOL_EXTENSIONS(a_tool_extensionsDeleted)

	// get stage instance from DB instance, and call callback function
	a_tool_extensionsStaged := backRepo.BackRepoA_TOOL_EXTENSIONS.Map_A_TOOL_EXTENSIONSDBID_A_TOOL_EXTENSIONSPtr[a_tool_extensionsDB.ID]
	if a_tool_extensionsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_tool_extensionsStaged, a_tool_extensionsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
