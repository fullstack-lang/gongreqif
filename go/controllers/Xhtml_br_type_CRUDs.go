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
var __Xhtml_br_type__dummysDeclaration__ models.Xhtml_br_type
var __Xhtml_br_type_time__dummyDeclaration time.Duration

var mutexXhtml_br_type sync.Mutex

// An Xhtml_br_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_br_type updateXhtml_br_type deleteXhtml_br_type
type Xhtml_br_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_br_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_br_type updateXhtml_br_type
type Xhtml_br_typeInput struct {
	// The Xhtml_br_type to submit or modify
	// in: body
	Xhtml_br_type *orm.Xhtml_br_typeAPI
}

// GetXhtml_br_types
//
// swagger:route GET /xhtml_br_types xhtml_br_types getXhtml_br_types
//
// # Get all xhtml_br_types
//
// Responses:
// default: genericError
//
//	200: xhtml_br_typeDBResponse
func (controller *Controller) GetXhtml_br_types(c *gin.Context) {

	// source slice
	var xhtml_br_typeDBs []orm.Xhtml_br_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_br_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_br_type.GetDB()

	_, err := db.Find(&xhtml_br_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_br_typeAPIs := make([]orm.Xhtml_br_typeAPI, 0)

	// for each xhtml_br_type, update fields from the database nullable fields
	for idx := range xhtml_br_typeDBs {
		xhtml_br_typeDB := &xhtml_br_typeDBs[idx]
		_ = xhtml_br_typeDB
		var xhtml_br_typeAPI orm.Xhtml_br_typeAPI

		// insertion point for updating fields
		xhtml_br_typeAPI.ID = xhtml_br_typeDB.ID
		xhtml_br_typeDB.CopyBasicFieldsToXhtml_br_type_WOP(&xhtml_br_typeAPI.Xhtml_br_type_WOP)
		xhtml_br_typeAPI.Xhtml_br_typePointersEncoding = xhtml_br_typeDB.Xhtml_br_typePointersEncoding
		xhtml_br_typeAPIs = append(xhtml_br_typeAPIs, xhtml_br_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_br_typeAPIs)
}

// PostXhtml_br_type
//
// swagger:route POST /xhtml_br_types xhtml_br_types postXhtml_br_type
//
// Creates a xhtml_br_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_br_type(c *gin.Context) {

	mutexXhtml_br_type.Lock()
	defer mutexXhtml_br_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_br_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_br_type.GetDB()

	// Validate input
	var input orm.Xhtml_br_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_br_type
	xhtml_br_typeDB := orm.Xhtml_br_typeDB{}
	xhtml_br_typeDB.Xhtml_br_typePointersEncoding = input.Xhtml_br_typePointersEncoding
	xhtml_br_typeDB.CopyBasicFieldsFromXhtml_br_type_WOP(&input.Xhtml_br_type_WOP)

	_, err = db.Create(&xhtml_br_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_br_type.CheckoutPhaseOneInstance(&xhtml_br_typeDB)
	xhtml_br_type := backRepo.BackRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr[xhtml_br_typeDB.ID]

	if xhtml_br_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_br_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_br_typeDB)
}

// GetXhtml_br_type
//
// swagger:route GET /xhtml_br_types/{ID} xhtml_br_types getXhtml_br_type
//
// Gets the details for a xhtml_br_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_br_typeDBResponse
func (controller *Controller) GetXhtml_br_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_br_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_br_type.GetDB()

	// Get xhtml_br_typeDB in DB
	var xhtml_br_typeDB orm.Xhtml_br_typeDB
	if _, err := db.First(&xhtml_br_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_br_typeAPI orm.Xhtml_br_typeAPI
	xhtml_br_typeAPI.ID = xhtml_br_typeDB.ID
	xhtml_br_typeAPI.Xhtml_br_typePointersEncoding = xhtml_br_typeDB.Xhtml_br_typePointersEncoding
	xhtml_br_typeDB.CopyBasicFieldsToXhtml_br_type_WOP(&xhtml_br_typeAPI.Xhtml_br_type_WOP)

	c.JSON(http.StatusOK, xhtml_br_typeAPI)
}

// UpdateXhtml_br_type
//
// swagger:route PATCH /xhtml_br_types/{ID} xhtml_br_types updateXhtml_br_type
//
// # Update a xhtml_br_type
//
// Responses:
// default: genericError
//
//	200: xhtml_br_typeDBResponse
func (controller *Controller) UpdateXhtml_br_type(c *gin.Context) {

	mutexXhtml_br_type.Lock()
	defer mutexXhtml_br_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_br_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_br_type.GetDB()

	// Validate input
	var input orm.Xhtml_br_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_br_typeDB orm.Xhtml_br_typeDB

	// fetch the xhtml_br_type
	_, err := db.First(&xhtml_br_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_br_typeDB.CopyBasicFieldsFromXhtml_br_type_WOP(&input.Xhtml_br_type_WOP)
	xhtml_br_typeDB.Xhtml_br_typePointersEncoding = input.Xhtml_br_typePointersEncoding

	db, _ = db.Model(&xhtml_br_typeDB)
	_, err = db.Updates(&xhtml_br_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_br_typeNew := new(models.Xhtml_br_type)
	xhtml_br_typeDB.CopyBasicFieldsToXhtml_br_type(xhtml_br_typeNew)

	// redeem pointers
	xhtml_br_typeDB.DecodePointers(backRepo, xhtml_br_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_br_typeOld := backRepo.BackRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr[xhtml_br_typeDB.ID]
	if xhtml_br_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_br_typeOld, xhtml_br_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_br_typeDB
	c.JSON(http.StatusOK, xhtml_br_typeDB)
}

// DeleteXhtml_br_type
//
// swagger:route DELETE /xhtml_br_types/{ID} xhtml_br_types deleteXhtml_br_type
//
// # Delete a xhtml_br_type
//
// default: genericError
//
//	200: xhtml_br_typeDBResponse
func (controller *Controller) DeleteXhtml_br_type(c *gin.Context) {

	mutexXhtml_br_type.Lock()
	defer mutexXhtml_br_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_br_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_br_type.GetDB()

	// Get model if exist
	var xhtml_br_typeDB orm.Xhtml_br_typeDB
	if _, err := db.First(&xhtml_br_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_br_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_br_typeDeleted := new(models.Xhtml_br_type)
	xhtml_br_typeDB.CopyBasicFieldsToXhtml_br_type(xhtml_br_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_br_typeStaged := backRepo.BackRepoXhtml_br_type.Map_Xhtml_br_typeDBID_Xhtml_br_typePtr[xhtml_br_typeDB.ID]
	if xhtml_br_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_br_typeStaged, xhtml_br_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
