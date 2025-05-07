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
var __Xhtml_a_type__dummysDeclaration__ models.Xhtml_a_type
var __Xhtml_a_type_time__dummyDeclaration time.Duration

var mutexXhtml_a_type sync.Mutex

// An Xhtml_a_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_a_type updateXhtml_a_type deleteXhtml_a_type
type Xhtml_a_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_a_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_a_type updateXhtml_a_type
type Xhtml_a_typeInput struct {
	// The Xhtml_a_type to submit or modify
	// in: body
	Xhtml_a_type *orm.Xhtml_a_typeAPI
}

// GetXhtml_a_types
//
// swagger:route GET /xhtml_a_types xhtml_a_types getXhtml_a_types
//
// # Get all xhtml_a_types
//
// Responses:
// default: genericError
//
//	200: xhtml_a_typeDBResponse
func (controller *Controller) GetXhtml_a_types(c *gin.Context) {

	// source slice
	var xhtml_a_typeDBs []orm.Xhtml_a_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_a_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_a_type.GetDB()

	_, err := db.Find(&xhtml_a_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_a_typeAPIs := make([]orm.Xhtml_a_typeAPI, 0)

	// for each xhtml_a_type, update fields from the database nullable fields
	for idx := range xhtml_a_typeDBs {
		xhtml_a_typeDB := &xhtml_a_typeDBs[idx]
		_ = xhtml_a_typeDB
		var xhtml_a_typeAPI orm.Xhtml_a_typeAPI

		// insertion point for updating fields
		xhtml_a_typeAPI.ID = xhtml_a_typeDB.ID
		xhtml_a_typeDB.CopyBasicFieldsToXhtml_a_type_WOP(&xhtml_a_typeAPI.Xhtml_a_type_WOP)
		xhtml_a_typeAPI.Xhtml_a_typePointersEncoding = xhtml_a_typeDB.Xhtml_a_typePointersEncoding
		xhtml_a_typeAPIs = append(xhtml_a_typeAPIs, xhtml_a_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_a_typeAPIs)
}

// PostXhtml_a_type
//
// swagger:route POST /xhtml_a_types xhtml_a_types postXhtml_a_type
//
// Creates a xhtml_a_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_a_type(c *gin.Context) {

	mutexXhtml_a_type.Lock()
	defer mutexXhtml_a_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_a_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_a_type.GetDB()

	// Validate input
	var input orm.Xhtml_a_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_a_type
	xhtml_a_typeDB := orm.Xhtml_a_typeDB{}
	xhtml_a_typeDB.Xhtml_a_typePointersEncoding = input.Xhtml_a_typePointersEncoding
	xhtml_a_typeDB.CopyBasicFieldsFromXhtml_a_type_WOP(&input.Xhtml_a_type_WOP)

	_, err = db.Create(&xhtml_a_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_a_type.CheckoutPhaseOneInstance(&xhtml_a_typeDB)
	xhtml_a_type := backRepo.BackRepoXhtml_a_type.Map_Xhtml_a_typeDBID_Xhtml_a_typePtr[xhtml_a_typeDB.ID]

	if xhtml_a_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_a_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_a_typeDB)
}

// GetXhtml_a_type
//
// swagger:route GET /xhtml_a_types/{ID} xhtml_a_types getXhtml_a_type
//
// Gets the details for a xhtml_a_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_a_typeDBResponse
func (controller *Controller) GetXhtml_a_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_a_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_a_type.GetDB()

	// Get xhtml_a_typeDB in DB
	var xhtml_a_typeDB orm.Xhtml_a_typeDB
	if _, err := db.First(&xhtml_a_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_a_typeAPI orm.Xhtml_a_typeAPI
	xhtml_a_typeAPI.ID = xhtml_a_typeDB.ID
	xhtml_a_typeAPI.Xhtml_a_typePointersEncoding = xhtml_a_typeDB.Xhtml_a_typePointersEncoding
	xhtml_a_typeDB.CopyBasicFieldsToXhtml_a_type_WOP(&xhtml_a_typeAPI.Xhtml_a_type_WOP)

	c.JSON(http.StatusOK, xhtml_a_typeAPI)
}

// UpdateXhtml_a_type
//
// swagger:route PATCH /xhtml_a_types/{ID} xhtml_a_types updateXhtml_a_type
//
// # Update a xhtml_a_type
//
// Responses:
// default: genericError
//
//	200: xhtml_a_typeDBResponse
func (controller *Controller) UpdateXhtml_a_type(c *gin.Context) {

	mutexXhtml_a_type.Lock()
	defer mutexXhtml_a_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_a_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_a_type.GetDB()

	// Validate input
	var input orm.Xhtml_a_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_a_typeDB orm.Xhtml_a_typeDB

	// fetch the xhtml_a_type
	_, err := db.First(&xhtml_a_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_a_typeDB.CopyBasicFieldsFromXhtml_a_type_WOP(&input.Xhtml_a_type_WOP)
	xhtml_a_typeDB.Xhtml_a_typePointersEncoding = input.Xhtml_a_typePointersEncoding

	db, _ = db.Model(&xhtml_a_typeDB)
	_, err = db.Updates(&xhtml_a_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_a_typeNew := new(models.Xhtml_a_type)
	xhtml_a_typeDB.CopyBasicFieldsToXhtml_a_type(xhtml_a_typeNew)

	// redeem pointers
	xhtml_a_typeDB.DecodePointers(backRepo, xhtml_a_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_a_typeOld := backRepo.BackRepoXhtml_a_type.Map_Xhtml_a_typeDBID_Xhtml_a_typePtr[xhtml_a_typeDB.ID]
	if xhtml_a_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_a_typeOld, xhtml_a_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_a_typeDB
	c.JSON(http.StatusOK, xhtml_a_typeDB)
}

// DeleteXhtml_a_type
//
// swagger:route DELETE /xhtml_a_types/{ID} xhtml_a_types deleteXhtml_a_type
//
// # Delete a xhtml_a_type
//
// default: genericError
//
//	200: xhtml_a_typeDBResponse
func (controller *Controller) DeleteXhtml_a_type(c *gin.Context) {

	mutexXhtml_a_type.Lock()
	defer mutexXhtml_a_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_a_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_a_type.GetDB()

	// Get model if exist
	var xhtml_a_typeDB orm.Xhtml_a_typeDB
	if _, err := db.First(&xhtml_a_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_a_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_a_typeDeleted := new(models.Xhtml_a_type)
	xhtml_a_typeDB.CopyBasicFieldsToXhtml_a_type(xhtml_a_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_a_typeStaged := backRepo.BackRepoXhtml_a_type.Map_Xhtml_a_typeDBID_Xhtml_a_typePtr[xhtml_a_typeDB.ID]
	if xhtml_a_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_a_typeStaged, xhtml_a_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
