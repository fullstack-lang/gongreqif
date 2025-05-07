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
var __Xhtml_blockquote_type__dummysDeclaration__ models.Xhtml_blockquote_type
var __Xhtml_blockquote_type_time__dummyDeclaration time.Duration

var mutexXhtml_blockquote_type sync.Mutex

// An Xhtml_blockquote_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_blockquote_type updateXhtml_blockquote_type deleteXhtml_blockquote_type
type Xhtml_blockquote_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_blockquote_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_blockquote_type updateXhtml_blockquote_type
type Xhtml_blockquote_typeInput struct {
	// The Xhtml_blockquote_type to submit or modify
	// in: body
	Xhtml_blockquote_type *orm.Xhtml_blockquote_typeAPI
}

// GetXhtml_blockquote_types
//
// swagger:route GET /xhtml_blockquote_types xhtml_blockquote_types getXhtml_blockquote_types
//
// # Get all xhtml_blockquote_types
//
// Responses:
// default: genericError
//
//	200: xhtml_blockquote_typeDBResponse
func (controller *Controller) GetXhtml_blockquote_types(c *gin.Context) {

	// source slice
	var xhtml_blockquote_typeDBs []orm.Xhtml_blockquote_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_blockquote_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_blockquote_type.GetDB()

	_, err := db.Find(&xhtml_blockquote_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_blockquote_typeAPIs := make([]orm.Xhtml_blockquote_typeAPI, 0)

	// for each xhtml_blockquote_type, update fields from the database nullable fields
	for idx := range xhtml_blockquote_typeDBs {
		xhtml_blockquote_typeDB := &xhtml_blockquote_typeDBs[idx]
		_ = xhtml_blockquote_typeDB
		var xhtml_blockquote_typeAPI orm.Xhtml_blockquote_typeAPI

		// insertion point for updating fields
		xhtml_blockquote_typeAPI.ID = xhtml_blockquote_typeDB.ID
		xhtml_blockquote_typeDB.CopyBasicFieldsToXhtml_blockquote_type_WOP(&xhtml_blockquote_typeAPI.Xhtml_blockquote_type_WOP)
		xhtml_blockquote_typeAPI.Xhtml_blockquote_typePointersEncoding = xhtml_blockquote_typeDB.Xhtml_blockquote_typePointersEncoding
		xhtml_blockquote_typeAPIs = append(xhtml_blockquote_typeAPIs, xhtml_blockquote_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_blockquote_typeAPIs)
}

// PostXhtml_blockquote_type
//
// swagger:route POST /xhtml_blockquote_types xhtml_blockquote_types postXhtml_blockquote_type
//
// Creates a xhtml_blockquote_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_blockquote_type(c *gin.Context) {

	mutexXhtml_blockquote_type.Lock()
	defer mutexXhtml_blockquote_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_blockquote_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_blockquote_type.GetDB()

	// Validate input
	var input orm.Xhtml_blockquote_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_blockquote_type
	xhtml_blockquote_typeDB := orm.Xhtml_blockquote_typeDB{}
	xhtml_blockquote_typeDB.Xhtml_blockquote_typePointersEncoding = input.Xhtml_blockquote_typePointersEncoding
	xhtml_blockquote_typeDB.CopyBasicFieldsFromXhtml_blockquote_type_WOP(&input.Xhtml_blockquote_type_WOP)

	_, err = db.Create(&xhtml_blockquote_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_blockquote_type.CheckoutPhaseOneInstance(&xhtml_blockquote_typeDB)
	xhtml_blockquote_type := backRepo.BackRepoXhtml_blockquote_type.Map_Xhtml_blockquote_typeDBID_Xhtml_blockquote_typePtr[xhtml_blockquote_typeDB.ID]

	if xhtml_blockquote_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_blockquote_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_blockquote_typeDB)
}

// GetXhtml_blockquote_type
//
// swagger:route GET /xhtml_blockquote_types/{ID} xhtml_blockquote_types getXhtml_blockquote_type
//
// Gets the details for a xhtml_blockquote_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_blockquote_typeDBResponse
func (controller *Controller) GetXhtml_blockquote_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_blockquote_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_blockquote_type.GetDB()

	// Get xhtml_blockquote_typeDB in DB
	var xhtml_blockquote_typeDB orm.Xhtml_blockquote_typeDB
	if _, err := db.First(&xhtml_blockquote_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_blockquote_typeAPI orm.Xhtml_blockquote_typeAPI
	xhtml_blockquote_typeAPI.ID = xhtml_blockquote_typeDB.ID
	xhtml_blockquote_typeAPI.Xhtml_blockquote_typePointersEncoding = xhtml_blockquote_typeDB.Xhtml_blockquote_typePointersEncoding
	xhtml_blockquote_typeDB.CopyBasicFieldsToXhtml_blockquote_type_WOP(&xhtml_blockquote_typeAPI.Xhtml_blockquote_type_WOP)

	c.JSON(http.StatusOK, xhtml_blockquote_typeAPI)
}

// UpdateXhtml_blockquote_type
//
// swagger:route PATCH /xhtml_blockquote_types/{ID} xhtml_blockquote_types updateXhtml_blockquote_type
//
// # Update a xhtml_blockquote_type
//
// Responses:
// default: genericError
//
//	200: xhtml_blockquote_typeDBResponse
func (controller *Controller) UpdateXhtml_blockquote_type(c *gin.Context) {

	mutexXhtml_blockquote_type.Lock()
	defer mutexXhtml_blockquote_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_blockquote_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_blockquote_type.GetDB()

	// Validate input
	var input orm.Xhtml_blockquote_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_blockquote_typeDB orm.Xhtml_blockquote_typeDB

	// fetch the xhtml_blockquote_type
	_, err := db.First(&xhtml_blockquote_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_blockquote_typeDB.CopyBasicFieldsFromXhtml_blockquote_type_WOP(&input.Xhtml_blockquote_type_WOP)
	xhtml_blockquote_typeDB.Xhtml_blockquote_typePointersEncoding = input.Xhtml_blockquote_typePointersEncoding

	db, _ = db.Model(&xhtml_blockquote_typeDB)
	_, err = db.Updates(&xhtml_blockquote_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_blockquote_typeNew := new(models.Xhtml_blockquote_type)
	xhtml_blockquote_typeDB.CopyBasicFieldsToXhtml_blockquote_type(xhtml_blockquote_typeNew)

	// redeem pointers
	xhtml_blockquote_typeDB.DecodePointers(backRepo, xhtml_blockquote_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_blockquote_typeOld := backRepo.BackRepoXhtml_blockquote_type.Map_Xhtml_blockquote_typeDBID_Xhtml_blockquote_typePtr[xhtml_blockquote_typeDB.ID]
	if xhtml_blockquote_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_blockquote_typeOld, xhtml_blockquote_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_blockquote_typeDB
	c.JSON(http.StatusOK, xhtml_blockquote_typeDB)
}

// DeleteXhtml_blockquote_type
//
// swagger:route DELETE /xhtml_blockquote_types/{ID} xhtml_blockquote_types deleteXhtml_blockquote_type
//
// # Delete a xhtml_blockquote_type
//
// default: genericError
//
//	200: xhtml_blockquote_typeDBResponse
func (controller *Controller) DeleteXhtml_blockquote_type(c *gin.Context) {

	mutexXhtml_blockquote_type.Lock()
	defer mutexXhtml_blockquote_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_blockquote_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_blockquote_type.GetDB()

	// Get model if exist
	var xhtml_blockquote_typeDB orm.Xhtml_blockquote_typeDB
	if _, err := db.First(&xhtml_blockquote_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_blockquote_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_blockquote_typeDeleted := new(models.Xhtml_blockquote_type)
	xhtml_blockquote_typeDB.CopyBasicFieldsToXhtml_blockquote_type(xhtml_blockquote_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_blockquote_typeStaged := backRepo.BackRepoXhtml_blockquote_type.Map_Xhtml_blockquote_typeDBID_Xhtml_blockquote_typePtr[xhtml_blockquote_typeDB.ID]
	if xhtml_blockquote_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_blockquote_typeStaged, xhtml_blockquote_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
