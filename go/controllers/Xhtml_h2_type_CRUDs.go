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
var __Xhtml_h2_type__dummysDeclaration__ models.Xhtml_h2_type
var __Xhtml_h2_type_time__dummyDeclaration time.Duration

var mutexXhtml_h2_type sync.Mutex

// An Xhtml_h2_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_h2_type updateXhtml_h2_type deleteXhtml_h2_type
type Xhtml_h2_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_h2_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_h2_type updateXhtml_h2_type
type Xhtml_h2_typeInput struct {
	// The Xhtml_h2_type to submit or modify
	// in: body
	Xhtml_h2_type *orm.Xhtml_h2_typeAPI
}

// GetXhtml_h2_types
//
// swagger:route GET /xhtml_h2_types xhtml_h2_types getXhtml_h2_types
//
// # Get all xhtml_h2_types
//
// Responses:
// default: genericError
//
//	200: xhtml_h2_typeDBResponse
func (controller *Controller) GetXhtml_h2_types(c *gin.Context) {

	// source slice
	var xhtml_h2_typeDBs []orm.Xhtml_h2_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h2_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h2_type.GetDB()

	_, err := db.Find(&xhtml_h2_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_h2_typeAPIs := make([]orm.Xhtml_h2_typeAPI, 0)

	// for each xhtml_h2_type, update fields from the database nullable fields
	for idx := range xhtml_h2_typeDBs {
		xhtml_h2_typeDB := &xhtml_h2_typeDBs[idx]
		_ = xhtml_h2_typeDB
		var xhtml_h2_typeAPI orm.Xhtml_h2_typeAPI

		// insertion point for updating fields
		xhtml_h2_typeAPI.ID = xhtml_h2_typeDB.ID
		xhtml_h2_typeDB.CopyBasicFieldsToXhtml_h2_type_WOP(&xhtml_h2_typeAPI.Xhtml_h2_type_WOP)
		xhtml_h2_typeAPI.Xhtml_h2_typePointersEncoding = xhtml_h2_typeDB.Xhtml_h2_typePointersEncoding
		xhtml_h2_typeAPIs = append(xhtml_h2_typeAPIs, xhtml_h2_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_h2_typeAPIs)
}

// PostXhtml_h2_type
//
// swagger:route POST /xhtml_h2_types xhtml_h2_types postXhtml_h2_type
//
// Creates a xhtml_h2_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_h2_type(c *gin.Context) {

	mutexXhtml_h2_type.Lock()
	defer mutexXhtml_h2_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_h2_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h2_type.GetDB()

	// Validate input
	var input orm.Xhtml_h2_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_h2_type
	xhtml_h2_typeDB := orm.Xhtml_h2_typeDB{}
	xhtml_h2_typeDB.Xhtml_h2_typePointersEncoding = input.Xhtml_h2_typePointersEncoding
	xhtml_h2_typeDB.CopyBasicFieldsFromXhtml_h2_type_WOP(&input.Xhtml_h2_type_WOP)

	_, err = db.Create(&xhtml_h2_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_h2_type.CheckoutPhaseOneInstance(&xhtml_h2_typeDB)
	xhtml_h2_type := backRepo.BackRepoXhtml_h2_type.Map_Xhtml_h2_typeDBID_Xhtml_h2_typePtr[xhtml_h2_typeDB.ID]

	if xhtml_h2_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_h2_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_h2_typeDB)
}

// GetXhtml_h2_type
//
// swagger:route GET /xhtml_h2_types/{ID} xhtml_h2_types getXhtml_h2_type
//
// Gets the details for a xhtml_h2_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_h2_typeDBResponse
func (controller *Controller) GetXhtml_h2_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h2_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h2_type.GetDB()

	// Get xhtml_h2_typeDB in DB
	var xhtml_h2_typeDB orm.Xhtml_h2_typeDB
	if _, err := db.First(&xhtml_h2_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_h2_typeAPI orm.Xhtml_h2_typeAPI
	xhtml_h2_typeAPI.ID = xhtml_h2_typeDB.ID
	xhtml_h2_typeAPI.Xhtml_h2_typePointersEncoding = xhtml_h2_typeDB.Xhtml_h2_typePointersEncoding
	xhtml_h2_typeDB.CopyBasicFieldsToXhtml_h2_type_WOP(&xhtml_h2_typeAPI.Xhtml_h2_type_WOP)

	c.JSON(http.StatusOK, xhtml_h2_typeAPI)
}

// UpdateXhtml_h2_type
//
// swagger:route PATCH /xhtml_h2_types/{ID} xhtml_h2_types updateXhtml_h2_type
//
// # Update a xhtml_h2_type
//
// Responses:
// default: genericError
//
//	200: xhtml_h2_typeDBResponse
func (controller *Controller) UpdateXhtml_h2_type(c *gin.Context) {

	mutexXhtml_h2_type.Lock()
	defer mutexXhtml_h2_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_h2_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h2_type.GetDB()

	// Validate input
	var input orm.Xhtml_h2_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_h2_typeDB orm.Xhtml_h2_typeDB

	// fetch the xhtml_h2_type
	_, err := db.First(&xhtml_h2_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_h2_typeDB.CopyBasicFieldsFromXhtml_h2_type_WOP(&input.Xhtml_h2_type_WOP)
	xhtml_h2_typeDB.Xhtml_h2_typePointersEncoding = input.Xhtml_h2_typePointersEncoding

	db, _ = db.Model(&xhtml_h2_typeDB)
	_, err = db.Updates(&xhtml_h2_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h2_typeNew := new(models.Xhtml_h2_type)
	xhtml_h2_typeDB.CopyBasicFieldsToXhtml_h2_type(xhtml_h2_typeNew)

	// redeem pointers
	xhtml_h2_typeDB.DecodePointers(backRepo, xhtml_h2_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_h2_typeOld := backRepo.BackRepoXhtml_h2_type.Map_Xhtml_h2_typeDBID_Xhtml_h2_typePtr[xhtml_h2_typeDB.ID]
	if xhtml_h2_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_h2_typeOld, xhtml_h2_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_h2_typeDB
	c.JSON(http.StatusOK, xhtml_h2_typeDB)
}

// DeleteXhtml_h2_type
//
// swagger:route DELETE /xhtml_h2_types/{ID} xhtml_h2_types deleteXhtml_h2_type
//
// # Delete a xhtml_h2_type
//
// default: genericError
//
//	200: xhtml_h2_typeDBResponse
func (controller *Controller) DeleteXhtml_h2_type(c *gin.Context) {

	mutexXhtml_h2_type.Lock()
	defer mutexXhtml_h2_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_h2_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h2_type.GetDB()

	// Get model if exist
	var xhtml_h2_typeDB orm.Xhtml_h2_typeDB
	if _, err := db.First(&xhtml_h2_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_h2_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h2_typeDeleted := new(models.Xhtml_h2_type)
	xhtml_h2_typeDB.CopyBasicFieldsToXhtml_h2_type(xhtml_h2_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_h2_typeStaged := backRepo.BackRepoXhtml_h2_type.Map_Xhtml_h2_typeDBID_Xhtml_h2_typePtr[xhtml_h2_typeDB.ID]
	if xhtml_h2_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_h2_typeStaged, xhtml_h2_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
