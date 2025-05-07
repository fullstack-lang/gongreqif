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
var __Xhtml_span_type__dummysDeclaration__ models.Xhtml_span_type
var __Xhtml_span_type_time__dummyDeclaration time.Duration

var mutexXhtml_span_type sync.Mutex

// An Xhtml_span_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_span_type updateXhtml_span_type deleteXhtml_span_type
type Xhtml_span_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_span_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_span_type updateXhtml_span_type
type Xhtml_span_typeInput struct {
	// The Xhtml_span_type to submit or modify
	// in: body
	Xhtml_span_type *orm.Xhtml_span_typeAPI
}

// GetXhtml_span_types
//
// swagger:route GET /xhtml_span_types xhtml_span_types getXhtml_span_types
//
// # Get all xhtml_span_types
//
// Responses:
// default: genericError
//
//	200: xhtml_span_typeDBResponse
func (controller *Controller) GetXhtml_span_types(c *gin.Context) {

	// source slice
	var xhtml_span_typeDBs []orm.Xhtml_span_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_span_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_span_type.GetDB()

	_, err := db.Find(&xhtml_span_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_span_typeAPIs := make([]orm.Xhtml_span_typeAPI, 0)

	// for each xhtml_span_type, update fields from the database nullable fields
	for idx := range xhtml_span_typeDBs {
		xhtml_span_typeDB := &xhtml_span_typeDBs[idx]
		_ = xhtml_span_typeDB
		var xhtml_span_typeAPI orm.Xhtml_span_typeAPI

		// insertion point for updating fields
		xhtml_span_typeAPI.ID = xhtml_span_typeDB.ID
		xhtml_span_typeDB.CopyBasicFieldsToXhtml_span_type_WOP(&xhtml_span_typeAPI.Xhtml_span_type_WOP)
		xhtml_span_typeAPI.Xhtml_span_typePointersEncoding = xhtml_span_typeDB.Xhtml_span_typePointersEncoding
		xhtml_span_typeAPIs = append(xhtml_span_typeAPIs, xhtml_span_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_span_typeAPIs)
}

// PostXhtml_span_type
//
// swagger:route POST /xhtml_span_types xhtml_span_types postXhtml_span_type
//
// Creates a xhtml_span_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_span_type(c *gin.Context) {

	mutexXhtml_span_type.Lock()
	defer mutexXhtml_span_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_span_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_span_type.GetDB()

	// Validate input
	var input orm.Xhtml_span_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_span_type
	xhtml_span_typeDB := orm.Xhtml_span_typeDB{}
	xhtml_span_typeDB.Xhtml_span_typePointersEncoding = input.Xhtml_span_typePointersEncoding
	xhtml_span_typeDB.CopyBasicFieldsFromXhtml_span_type_WOP(&input.Xhtml_span_type_WOP)

	_, err = db.Create(&xhtml_span_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_span_type.CheckoutPhaseOneInstance(&xhtml_span_typeDB)
	xhtml_span_type := backRepo.BackRepoXhtml_span_type.Map_Xhtml_span_typeDBID_Xhtml_span_typePtr[xhtml_span_typeDB.ID]

	if xhtml_span_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_span_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_span_typeDB)
}

// GetXhtml_span_type
//
// swagger:route GET /xhtml_span_types/{ID} xhtml_span_types getXhtml_span_type
//
// Gets the details for a xhtml_span_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_span_typeDBResponse
func (controller *Controller) GetXhtml_span_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_span_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_span_type.GetDB()

	// Get xhtml_span_typeDB in DB
	var xhtml_span_typeDB orm.Xhtml_span_typeDB
	if _, err := db.First(&xhtml_span_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_span_typeAPI orm.Xhtml_span_typeAPI
	xhtml_span_typeAPI.ID = xhtml_span_typeDB.ID
	xhtml_span_typeAPI.Xhtml_span_typePointersEncoding = xhtml_span_typeDB.Xhtml_span_typePointersEncoding
	xhtml_span_typeDB.CopyBasicFieldsToXhtml_span_type_WOP(&xhtml_span_typeAPI.Xhtml_span_type_WOP)

	c.JSON(http.StatusOK, xhtml_span_typeAPI)
}

// UpdateXhtml_span_type
//
// swagger:route PATCH /xhtml_span_types/{ID} xhtml_span_types updateXhtml_span_type
//
// # Update a xhtml_span_type
//
// Responses:
// default: genericError
//
//	200: xhtml_span_typeDBResponse
func (controller *Controller) UpdateXhtml_span_type(c *gin.Context) {

	mutexXhtml_span_type.Lock()
	defer mutexXhtml_span_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_span_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_span_type.GetDB()

	// Validate input
	var input orm.Xhtml_span_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_span_typeDB orm.Xhtml_span_typeDB

	// fetch the xhtml_span_type
	_, err := db.First(&xhtml_span_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_span_typeDB.CopyBasicFieldsFromXhtml_span_type_WOP(&input.Xhtml_span_type_WOP)
	xhtml_span_typeDB.Xhtml_span_typePointersEncoding = input.Xhtml_span_typePointersEncoding

	db, _ = db.Model(&xhtml_span_typeDB)
	_, err = db.Updates(&xhtml_span_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_span_typeNew := new(models.Xhtml_span_type)
	xhtml_span_typeDB.CopyBasicFieldsToXhtml_span_type(xhtml_span_typeNew)

	// redeem pointers
	xhtml_span_typeDB.DecodePointers(backRepo, xhtml_span_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_span_typeOld := backRepo.BackRepoXhtml_span_type.Map_Xhtml_span_typeDBID_Xhtml_span_typePtr[xhtml_span_typeDB.ID]
	if xhtml_span_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_span_typeOld, xhtml_span_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_span_typeDB
	c.JSON(http.StatusOK, xhtml_span_typeDB)
}

// DeleteXhtml_span_type
//
// swagger:route DELETE /xhtml_span_types/{ID} xhtml_span_types deleteXhtml_span_type
//
// # Delete a xhtml_span_type
//
// default: genericError
//
//	200: xhtml_span_typeDBResponse
func (controller *Controller) DeleteXhtml_span_type(c *gin.Context) {

	mutexXhtml_span_type.Lock()
	defer mutexXhtml_span_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_span_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_span_type.GetDB()

	// Get model if exist
	var xhtml_span_typeDB orm.Xhtml_span_typeDB
	if _, err := db.First(&xhtml_span_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_span_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_span_typeDeleted := new(models.Xhtml_span_type)
	xhtml_span_typeDB.CopyBasicFieldsToXhtml_span_type(xhtml_span_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_span_typeStaged := backRepo.BackRepoXhtml_span_type.Map_Xhtml_span_typeDBID_Xhtml_span_typePtr[xhtml_span_typeDB.ID]
	if xhtml_span_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_span_typeStaged, xhtml_span_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
