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
var __Xhtml_dd_type__dummysDeclaration__ models.Xhtml_dd_type
var __Xhtml_dd_type_time__dummyDeclaration time.Duration

var mutexXhtml_dd_type sync.Mutex

// An Xhtml_dd_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_dd_type updateXhtml_dd_type deleteXhtml_dd_type
type Xhtml_dd_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_dd_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_dd_type updateXhtml_dd_type
type Xhtml_dd_typeInput struct {
	// The Xhtml_dd_type to submit or modify
	// in: body
	Xhtml_dd_type *orm.Xhtml_dd_typeAPI
}

// GetXhtml_dd_types
//
// swagger:route GET /xhtml_dd_types xhtml_dd_types getXhtml_dd_types
//
// # Get all xhtml_dd_types
//
// Responses:
// default: genericError
//
//	200: xhtml_dd_typeDBResponse
func (controller *Controller) GetXhtml_dd_types(c *gin.Context) {

	// source slice
	var xhtml_dd_typeDBs []orm.Xhtml_dd_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_dd_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dd_type.GetDB()

	_, err := db.Find(&xhtml_dd_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_dd_typeAPIs := make([]orm.Xhtml_dd_typeAPI, 0)

	// for each xhtml_dd_type, update fields from the database nullable fields
	for idx := range xhtml_dd_typeDBs {
		xhtml_dd_typeDB := &xhtml_dd_typeDBs[idx]
		_ = xhtml_dd_typeDB
		var xhtml_dd_typeAPI orm.Xhtml_dd_typeAPI

		// insertion point for updating fields
		xhtml_dd_typeAPI.ID = xhtml_dd_typeDB.ID
		xhtml_dd_typeDB.CopyBasicFieldsToXhtml_dd_type_WOP(&xhtml_dd_typeAPI.Xhtml_dd_type_WOP)
		xhtml_dd_typeAPI.Xhtml_dd_typePointersEncoding = xhtml_dd_typeDB.Xhtml_dd_typePointersEncoding
		xhtml_dd_typeAPIs = append(xhtml_dd_typeAPIs, xhtml_dd_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_dd_typeAPIs)
}

// PostXhtml_dd_type
//
// swagger:route POST /xhtml_dd_types xhtml_dd_types postXhtml_dd_type
//
// Creates a xhtml_dd_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_dd_type(c *gin.Context) {

	mutexXhtml_dd_type.Lock()
	defer mutexXhtml_dd_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_dd_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dd_type.GetDB()

	// Validate input
	var input orm.Xhtml_dd_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_dd_type
	xhtml_dd_typeDB := orm.Xhtml_dd_typeDB{}
	xhtml_dd_typeDB.Xhtml_dd_typePointersEncoding = input.Xhtml_dd_typePointersEncoding
	xhtml_dd_typeDB.CopyBasicFieldsFromXhtml_dd_type_WOP(&input.Xhtml_dd_type_WOP)

	_, err = db.Create(&xhtml_dd_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_dd_type.CheckoutPhaseOneInstance(&xhtml_dd_typeDB)
	xhtml_dd_type := backRepo.BackRepoXhtml_dd_type.Map_Xhtml_dd_typeDBID_Xhtml_dd_typePtr[xhtml_dd_typeDB.ID]

	if xhtml_dd_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_dd_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_dd_typeDB)
}

// GetXhtml_dd_type
//
// swagger:route GET /xhtml_dd_types/{ID} xhtml_dd_types getXhtml_dd_type
//
// Gets the details for a xhtml_dd_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_dd_typeDBResponse
func (controller *Controller) GetXhtml_dd_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_dd_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dd_type.GetDB()

	// Get xhtml_dd_typeDB in DB
	var xhtml_dd_typeDB orm.Xhtml_dd_typeDB
	if _, err := db.First(&xhtml_dd_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_dd_typeAPI orm.Xhtml_dd_typeAPI
	xhtml_dd_typeAPI.ID = xhtml_dd_typeDB.ID
	xhtml_dd_typeAPI.Xhtml_dd_typePointersEncoding = xhtml_dd_typeDB.Xhtml_dd_typePointersEncoding
	xhtml_dd_typeDB.CopyBasicFieldsToXhtml_dd_type_WOP(&xhtml_dd_typeAPI.Xhtml_dd_type_WOP)

	c.JSON(http.StatusOK, xhtml_dd_typeAPI)
}

// UpdateXhtml_dd_type
//
// swagger:route PATCH /xhtml_dd_types/{ID} xhtml_dd_types updateXhtml_dd_type
//
// # Update a xhtml_dd_type
//
// Responses:
// default: genericError
//
//	200: xhtml_dd_typeDBResponse
func (controller *Controller) UpdateXhtml_dd_type(c *gin.Context) {

	mutexXhtml_dd_type.Lock()
	defer mutexXhtml_dd_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_dd_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dd_type.GetDB()

	// Validate input
	var input orm.Xhtml_dd_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_dd_typeDB orm.Xhtml_dd_typeDB

	// fetch the xhtml_dd_type
	_, err := db.First(&xhtml_dd_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_dd_typeDB.CopyBasicFieldsFromXhtml_dd_type_WOP(&input.Xhtml_dd_type_WOP)
	xhtml_dd_typeDB.Xhtml_dd_typePointersEncoding = input.Xhtml_dd_typePointersEncoding

	db, _ = db.Model(&xhtml_dd_typeDB)
	_, err = db.Updates(&xhtml_dd_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_dd_typeNew := new(models.Xhtml_dd_type)
	xhtml_dd_typeDB.CopyBasicFieldsToXhtml_dd_type(xhtml_dd_typeNew)

	// redeem pointers
	xhtml_dd_typeDB.DecodePointers(backRepo, xhtml_dd_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_dd_typeOld := backRepo.BackRepoXhtml_dd_type.Map_Xhtml_dd_typeDBID_Xhtml_dd_typePtr[xhtml_dd_typeDB.ID]
	if xhtml_dd_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_dd_typeOld, xhtml_dd_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_dd_typeDB
	c.JSON(http.StatusOK, xhtml_dd_typeDB)
}

// DeleteXhtml_dd_type
//
// swagger:route DELETE /xhtml_dd_types/{ID} xhtml_dd_types deleteXhtml_dd_type
//
// # Delete a xhtml_dd_type
//
// default: genericError
//
//	200: xhtml_dd_typeDBResponse
func (controller *Controller) DeleteXhtml_dd_type(c *gin.Context) {

	mutexXhtml_dd_type.Lock()
	defer mutexXhtml_dd_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_dd_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dd_type.GetDB()

	// Get model if exist
	var xhtml_dd_typeDB orm.Xhtml_dd_typeDB
	if _, err := db.First(&xhtml_dd_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_dd_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_dd_typeDeleted := new(models.Xhtml_dd_type)
	xhtml_dd_typeDB.CopyBasicFieldsToXhtml_dd_type(xhtml_dd_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_dd_typeStaged := backRepo.BackRepoXhtml_dd_type.Map_Xhtml_dd_typeDBID_Xhtml_dd_typePtr[xhtml_dd_typeDB.ID]
	if xhtml_dd_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_dd_typeStaged, xhtml_dd_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
