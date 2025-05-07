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
var __Xhtml_abbr_type__dummysDeclaration__ models.Xhtml_abbr_type
var __Xhtml_abbr_type_time__dummyDeclaration time.Duration

var mutexXhtml_abbr_type sync.Mutex

// An Xhtml_abbr_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_abbr_type updateXhtml_abbr_type deleteXhtml_abbr_type
type Xhtml_abbr_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_abbr_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_abbr_type updateXhtml_abbr_type
type Xhtml_abbr_typeInput struct {
	// The Xhtml_abbr_type to submit or modify
	// in: body
	Xhtml_abbr_type *orm.Xhtml_abbr_typeAPI
}

// GetXhtml_abbr_types
//
// swagger:route GET /xhtml_abbr_types xhtml_abbr_types getXhtml_abbr_types
//
// # Get all xhtml_abbr_types
//
// Responses:
// default: genericError
//
//	200: xhtml_abbr_typeDBResponse
func (controller *Controller) GetXhtml_abbr_types(c *gin.Context) {

	// source slice
	var xhtml_abbr_typeDBs []orm.Xhtml_abbr_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_abbr_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_abbr_type.GetDB()

	_, err := db.Find(&xhtml_abbr_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_abbr_typeAPIs := make([]orm.Xhtml_abbr_typeAPI, 0)

	// for each xhtml_abbr_type, update fields from the database nullable fields
	for idx := range xhtml_abbr_typeDBs {
		xhtml_abbr_typeDB := &xhtml_abbr_typeDBs[idx]
		_ = xhtml_abbr_typeDB
		var xhtml_abbr_typeAPI orm.Xhtml_abbr_typeAPI

		// insertion point for updating fields
		xhtml_abbr_typeAPI.ID = xhtml_abbr_typeDB.ID
		xhtml_abbr_typeDB.CopyBasicFieldsToXhtml_abbr_type_WOP(&xhtml_abbr_typeAPI.Xhtml_abbr_type_WOP)
		xhtml_abbr_typeAPI.Xhtml_abbr_typePointersEncoding = xhtml_abbr_typeDB.Xhtml_abbr_typePointersEncoding
		xhtml_abbr_typeAPIs = append(xhtml_abbr_typeAPIs, xhtml_abbr_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_abbr_typeAPIs)
}

// PostXhtml_abbr_type
//
// swagger:route POST /xhtml_abbr_types xhtml_abbr_types postXhtml_abbr_type
//
// Creates a xhtml_abbr_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_abbr_type(c *gin.Context) {

	mutexXhtml_abbr_type.Lock()
	defer mutexXhtml_abbr_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_abbr_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_abbr_type.GetDB()

	// Validate input
	var input orm.Xhtml_abbr_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_abbr_type
	xhtml_abbr_typeDB := orm.Xhtml_abbr_typeDB{}
	xhtml_abbr_typeDB.Xhtml_abbr_typePointersEncoding = input.Xhtml_abbr_typePointersEncoding
	xhtml_abbr_typeDB.CopyBasicFieldsFromXhtml_abbr_type_WOP(&input.Xhtml_abbr_type_WOP)

	_, err = db.Create(&xhtml_abbr_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_abbr_type.CheckoutPhaseOneInstance(&xhtml_abbr_typeDB)
	xhtml_abbr_type := backRepo.BackRepoXhtml_abbr_type.Map_Xhtml_abbr_typeDBID_Xhtml_abbr_typePtr[xhtml_abbr_typeDB.ID]

	if xhtml_abbr_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_abbr_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_abbr_typeDB)
}

// GetXhtml_abbr_type
//
// swagger:route GET /xhtml_abbr_types/{ID} xhtml_abbr_types getXhtml_abbr_type
//
// Gets the details for a xhtml_abbr_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_abbr_typeDBResponse
func (controller *Controller) GetXhtml_abbr_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_abbr_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_abbr_type.GetDB()

	// Get xhtml_abbr_typeDB in DB
	var xhtml_abbr_typeDB orm.Xhtml_abbr_typeDB
	if _, err := db.First(&xhtml_abbr_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_abbr_typeAPI orm.Xhtml_abbr_typeAPI
	xhtml_abbr_typeAPI.ID = xhtml_abbr_typeDB.ID
	xhtml_abbr_typeAPI.Xhtml_abbr_typePointersEncoding = xhtml_abbr_typeDB.Xhtml_abbr_typePointersEncoding
	xhtml_abbr_typeDB.CopyBasicFieldsToXhtml_abbr_type_WOP(&xhtml_abbr_typeAPI.Xhtml_abbr_type_WOP)

	c.JSON(http.StatusOK, xhtml_abbr_typeAPI)
}

// UpdateXhtml_abbr_type
//
// swagger:route PATCH /xhtml_abbr_types/{ID} xhtml_abbr_types updateXhtml_abbr_type
//
// # Update a xhtml_abbr_type
//
// Responses:
// default: genericError
//
//	200: xhtml_abbr_typeDBResponse
func (controller *Controller) UpdateXhtml_abbr_type(c *gin.Context) {

	mutexXhtml_abbr_type.Lock()
	defer mutexXhtml_abbr_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_abbr_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_abbr_type.GetDB()

	// Validate input
	var input orm.Xhtml_abbr_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_abbr_typeDB orm.Xhtml_abbr_typeDB

	// fetch the xhtml_abbr_type
	_, err := db.First(&xhtml_abbr_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_abbr_typeDB.CopyBasicFieldsFromXhtml_abbr_type_WOP(&input.Xhtml_abbr_type_WOP)
	xhtml_abbr_typeDB.Xhtml_abbr_typePointersEncoding = input.Xhtml_abbr_typePointersEncoding

	db, _ = db.Model(&xhtml_abbr_typeDB)
	_, err = db.Updates(&xhtml_abbr_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_abbr_typeNew := new(models.Xhtml_abbr_type)
	xhtml_abbr_typeDB.CopyBasicFieldsToXhtml_abbr_type(xhtml_abbr_typeNew)

	// redeem pointers
	xhtml_abbr_typeDB.DecodePointers(backRepo, xhtml_abbr_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_abbr_typeOld := backRepo.BackRepoXhtml_abbr_type.Map_Xhtml_abbr_typeDBID_Xhtml_abbr_typePtr[xhtml_abbr_typeDB.ID]
	if xhtml_abbr_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_abbr_typeOld, xhtml_abbr_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_abbr_typeDB
	c.JSON(http.StatusOK, xhtml_abbr_typeDB)
}

// DeleteXhtml_abbr_type
//
// swagger:route DELETE /xhtml_abbr_types/{ID} xhtml_abbr_types deleteXhtml_abbr_type
//
// # Delete a xhtml_abbr_type
//
// default: genericError
//
//	200: xhtml_abbr_typeDBResponse
func (controller *Controller) DeleteXhtml_abbr_type(c *gin.Context) {

	mutexXhtml_abbr_type.Lock()
	defer mutexXhtml_abbr_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_abbr_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_abbr_type.GetDB()

	// Get model if exist
	var xhtml_abbr_typeDB orm.Xhtml_abbr_typeDB
	if _, err := db.First(&xhtml_abbr_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_abbr_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_abbr_typeDeleted := new(models.Xhtml_abbr_type)
	xhtml_abbr_typeDB.CopyBasicFieldsToXhtml_abbr_type(xhtml_abbr_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_abbr_typeStaged := backRepo.BackRepoXhtml_abbr_type.Map_Xhtml_abbr_typeDBID_Xhtml_abbr_typePtr[xhtml_abbr_typeDB.ID]
	if xhtml_abbr_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_abbr_typeStaged, xhtml_abbr_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
