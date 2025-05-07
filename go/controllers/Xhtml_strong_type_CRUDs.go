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
var __Xhtml_strong_type__dummysDeclaration__ models.Xhtml_strong_type
var __Xhtml_strong_type_time__dummyDeclaration time.Duration

var mutexXhtml_strong_type sync.Mutex

// An Xhtml_strong_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_strong_type updateXhtml_strong_type deleteXhtml_strong_type
type Xhtml_strong_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_strong_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_strong_type updateXhtml_strong_type
type Xhtml_strong_typeInput struct {
	// The Xhtml_strong_type to submit or modify
	// in: body
	Xhtml_strong_type *orm.Xhtml_strong_typeAPI
}

// GetXhtml_strong_types
//
// swagger:route GET /xhtml_strong_types xhtml_strong_types getXhtml_strong_types
//
// # Get all xhtml_strong_types
//
// Responses:
// default: genericError
//
//	200: xhtml_strong_typeDBResponse
func (controller *Controller) GetXhtml_strong_types(c *gin.Context) {

	// source slice
	var xhtml_strong_typeDBs []orm.Xhtml_strong_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_strong_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_strong_type.GetDB()

	_, err := db.Find(&xhtml_strong_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_strong_typeAPIs := make([]orm.Xhtml_strong_typeAPI, 0)

	// for each xhtml_strong_type, update fields from the database nullable fields
	for idx := range xhtml_strong_typeDBs {
		xhtml_strong_typeDB := &xhtml_strong_typeDBs[idx]
		_ = xhtml_strong_typeDB
		var xhtml_strong_typeAPI orm.Xhtml_strong_typeAPI

		// insertion point for updating fields
		xhtml_strong_typeAPI.ID = xhtml_strong_typeDB.ID
		xhtml_strong_typeDB.CopyBasicFieldsToXhtml_strong_type_WOP(&xhtml_strong_typeAPI.Xhtml_strong_type_WOP)
		xhtml_strong_typeAPI.Xhtml_strong_typePointersEncoding = xhtml_strong_typeDB.Xhtml_strong_typePointersEncoding
		xhtml_strong_typeAPIs = append(xhtml_strong_typeAPIs, xhtml_strong_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_strong_typeAPIs)
}

// PostXhtml_strong_type
//
// swagger:route POST /xhtml_strong_types xhtml_strong_types postXhtml_strong_type
//
// Creates a xhtml_strong_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_strong_type(c *gin.Context) {

	mutexXhtml_strong_type.Lock()
	defer mutexXhtml_strong_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_strong_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_strong_type.GetDB()

	// Validate input
	var input orm.Xhtml_strong_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_strong_type
	xhtml_strong_typeDB := orm.Xhtml_strong_typeDB{}
	xhtml_strong_typeDB.Xhtml_strong_typePointersEncoding = input.Xhtml_strong_typePointersEncoding
	xhtml_strong_typeDB.CopyBasicFieldsFromXhtml_strong_type_WOP(&input.Xhtml_strong_type_WOP)

	_, err = db.Create(&xhtml_strong_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_strong_type.CheckoutPhaseOneInstance(&xhtml_strong_typeDB)
	xhtml_strong_type := backRepo.BackRepoXhtml_strong_type.Map_Xhtml_strong_typeDBID_Xhtml_strong_typePtr[xhtml_strong_typeDB.ID]

	if xhtml_strong_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_strong_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_strong_typeDB)
}

// GetXhtml_strong_type
//
// swagger:route GET /xhtml_strong_types/{ID} xhtml_strong_types getXhtml_strong_type
//
// Gets the details for a xhtml_strong_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_strong_typeDBResponse
func (controller *Controller) GetXhtml_strong_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_strong_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_strong_type.GetDB()

	// Get xhtml_strong_typeDB in DB
	var xhtml_strong_typeDB orm.Xhtml_strong_typeDB
	if _, err := db.First(&xhtml_strong_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_strong_typeAPI orm.Xhtml_strong_typeAPI
	xhtml_strong_typeAPI.ID = xhtml_strong_typeDB.ID
	xhtml_strong_typeAPI.Xhtml_strong_typePointersEncoding = xhtml_strong_typeDB.Xhtml_strong_typePointersEncoding
	xhtml_strong_typeDB.CopyBasicFieldsToXhtml_strong_type_WOP(&xhtml_strong_typeAPI.Xhtml_strong_type_WOP)

	c.JSON(http.StatusOK, xhtml_strong_typeAPI)
}

// UpdateXhtml_strong_type
//
// swagger:route PATCH /xhtml_strong_types/{ID} xhtml_strong_types updateXhtml_strong_type
//
// # Update a xhtml_strong_type
//
// Responses:
// default: genericError
//
//	200: xhtml_strong_typeDBResponse
func (controller *Controller) UpdateXhtml_strong_type(c *gin.Context) {

	mutexXhtml_strong_type.Lock()
	defer mutexXhtml_strong_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_strong_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_strong_type.GetDB()

	// Validate input
	var input orm.Xhtml_strong_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_strong_typeDB orm.Xhtml_strong_typeDB

	// fetch the xhtml_strong_type
	_, err := db.First(&xhtml_strong_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_strong_typeDB.CopyBasicFieldsFromXhtml_strong_type_WOP(&input.Xhtml_strong_type_WOP)
	xhtml_strong_typeDB.Xhtml_strong_typePointersEncoding = input.Xhtml_strong_typePointersEncoding

	db, _ = db.Model(&xhtml_strong_typeDB)
	_, err = db.Updates(&xhtml_strong_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_strong_typeNew := new(models.Xhtml_strong_type)
	xhtml_strong_typeDB.CopyBasicFieldsToXhtml_strong_type(xhtml_strong_typeNew)

	// redeem pointers
	xhtml_strong_typeDB.DecodePointers(backRepo, xhtml_strong_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_strong_typeOld := backRepo.BackRepoXhtml_strong_type.Map_Xhtml_strong_typeDBID_Xhtml_strong_typePtr[xhtml_strong_typeDB.ID]
	if xhtml_strong_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_strong_typeOld, xhtml_strong_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_strong_typeDB
	c.JSON(http.StatusOK, xhtml_strong_typeDB)
}

// DeleteXhtml_strong_type
//
// swagger:route DELETE /xhtml_strong_types/{ID} xhtml_strong_types deleteXhtml_strong_type
//
// # Delete a xhtml_strong_type
//
// default: genericError
//
//	200: xhtml_strong_typeDBResponse
func (controller *Controller) DeleteXhtml_strong_type(c *gin.Context) {

	mutexXhtml_strong_type.Lock()
	defer mutexXhtml_strong_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_strong_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_strong_type.GetDB()

	// Get model if exist
	var xhtml_strong_typeDB orm.Xhtml_strong_typeDB
	if _, err := db.First(&xhtml_strong_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_strong_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_strong_typeDeleted := new(models.Xhtml_strong_type)
	xhtml_strong_typeDB.CopyBasicFieldsToXhtml_strong_type(xhtml_strong_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_strong_typeStaged := backRepo.BackRepoXhtml_strong_type.Map_Xhtml_strong_typeDBID_Xhtml_strong_typePtr[xhtml_strong_typeDB.ID]
	if xhtml_strong_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_strong_typeStaged, xhtml_strong_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
