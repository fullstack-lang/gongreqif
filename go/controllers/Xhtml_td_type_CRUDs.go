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
var __Xhtml_td_type__dummysDeclaration__ models.Xhtml_td_type
var __Xhtml_td_type_time__dummyDeclaration time.Duration

var mutexXhtml_td_type sync.Mutex

// An Xhtml_td_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_td_type updateXhtml_td_type deleteXhtml_td_type
type Xhtml_td_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_td_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_td_type updateXhtml_td_type
type Xhtml_td_typeInput struct {
	// The Xhtml_td_type to submit or modify
	// in: body
	Xhtml_td_type *orm.Xhtml_td_typeAPI
}

// GetXhtml_td_types
//
// swagger:route GET /xhtml_td_types xhtml_td_types getXhtml_td_types
//
// # Get all xhtml_td_types
//
// Responses:
// default: genericError
//
//	200: xhtml_td_typeDBResponse
func (controller *Controller) GetXhtml_td_types(c *gin.Context) {

	// source slice
	var xhtml_td_typeDBs []orm.Xhtml_td_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_td_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_td_type.GetDB()

	_, err := db.Find(&xhtml_td_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_td_typeAPIs := make([]orm.Xhtml_td_typeAPI, 0)

	// for each xhtml_td_type, update fields from the database nullable fields
	for idx := range xhtml_td_typeDBs {
		xhtml_td_typeDB := &xhtml_td_typeDBs[idx]
		_ = xhtml_td_typeDB
		var xhtml_td_typeAPI orm.Xhtml_td_typeAPI

		// insertion point for updating fields
		xhtml_td_typeAPI.ID = xhtml_td_typeDB.ID
		xhtml_td_typeDB.CopyBasicFieldsToXhtml_td_type_WOP(&xhtml_td_typeAPI.Xhtml_td_type_WOP)
		xhtml_td_typeAPI.Xhtml_td_typePointersEncoding = xhtml_td_typeDB.Xhtml_td_typePointersEncoding
		xhtml_td_typeAPIs = append(xhtml_td_typeAPIs, xhtml_td_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_td_typeAPIs)
}

// PostXhtml_td_type
//
// swagger:route POST /xhtml_td_types xhtml_td_types postXhtml_td_type
//
// Creates a xhtml_td_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_td_type(c *gin.Context) {

	mutexXhtml_td_type.Lock()
	defer mutexXhtml_td_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_td_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_td_type.GetDB()

	// Validate input
	var input orm.Xhtml_td_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_td_type
	xhtml_td_typeDB := orm.Xhtml_td_typeDB{}
	xhtml_td_typeDB.Xhtml_td_typePointersEncoding = input.Xhtml_td_typePointersEncoding
	xhtml_td_typeDB.CopyBasicFieldsFromXhtml_td_type_WOP(&input.Xhtml_td_type_WOP)

	_, err = db.Create(&xhtml_td_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_td_type.CheckoutPhaseOneInstance(&xhtml_td_typeDB)
	xhtml_td_type := backRepo.BackRepoXhtml_td_type.Map_Xhtml_td_typeDBID_Xhtml_td_typePtr[xhtml_td_typeDB.ID]

	if xhtml_td_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_td_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_td_typeDB)
}

// GetXhtml_td_type
//
// swagger:route GET /xhtml_td_types/{ID} xhtml_td_types getXhtml_td_type
//
// Gets the details for a xhtml_td_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_td_typeDBResponse
func (controller *Controller) GetXhtml_td_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_td_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_td_type.GetDB()

	// Get xhtml_td_typeDB in DB
	var xhtml_td_typeDB orm.Xhtml_td_typeDB
	if _, err := db.First(&xhtml_td_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_td_typeAPI orm.Xhtml_td_typeAPI
	xhtml_td_typeAPI.ID = xhtml_td_typeDB.ID
	xhtml_td_typeAPI.Xhtml_td_typePointersEncoding = xhtml_td_typeDB.Xhtml_td_typePointersEncoding
	xhtml_td_typeDB.CopyBasicFieldsToXhtml_td_type_WOP(&xhtml_td_typeAPI.Xhtml_td_type_WOP)

	c.JSON(http.StatusOK, xhtml_td_typeAPI)
}

// UpdateXhtml_td_type
//
// swagger:route PATCH /xhtml_td_types/{ID} xhtml_td_types updateXhtml_td_type
//
// # Update a xhtml_td_type
//
// Responses:
// default: genericError
//
//	200: xhtml_td_typeDBResponse
func (controller *Controller) UpdateXhtml_td_type(c *gin.Context) {

	mutexXhtml_td_type.Lock()
	defer mutexXhtml_td_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_td_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_td_type.GetDB()

	// Validate input
	var input orm.Xhtml_td_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_td_typeDB orm.Xhtml_td_typeDB

	// fetch the xhtml_td_type
	_, err := db.First(&xhtml_td_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_td_typeDB.CopyBasicFieldsFromXhtml_td_type_WOP(&input.Xhtml_td_type_WOP)
	xhtml_td_typeDB.Xhtml_td_typePointersEncoding = input.Xhtml_td_typePointersEncoding

	db, _ = db.Model(&xhtml_td_typeDB)
	_, err = db.Updates(&xhtml_td_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_td_typeNew := new(models.Xhtml_td_type)
	xhtml_td_typeDB.CopyBasicFieldsToXhtml_td_type(xhtml_td_typeNew)

	// redeem pointers
	xhtml_td_typeDB.DecodePointers(backRepo, xhtml_td_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_td_typeOld := backRepo.BackRepoXhtml_td_type.Map_Xhtml_td_typeDBID_Xhtml_td_typePtr[xhtml_td_typeDB.ID]
	if xhtml_td_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_td_typeOld, xhtml_td_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_td_typeDB
	c.JSON(http.StatusOK, xhtml_td_typeDB)
}

// DeleteXhtml_td_type
//
// swagger:route DELETE /xhtml_td_types/{ID} xhtml_td_types deleteXhtml_td_type
//
// # Delete a xhtml_td_type
//
// default: genericError
//
//	200: xhtml_td_typeDBResponse
func (controller *Controller) DeleteXhtml_td_type(c *gin.Context) {

	mutexXhtml_td_type.Lock()
	defer mutexXhtml_td_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_td_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_td_type.GetDB()

	// Get model if exist
	var xhtml_td_typeDB orm.Xhtml_td_typeDB
	if _, err := db.First(&xhtml_td_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_td_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_td_typeDeleted := new(models.Xhtml_td_type)
	xhtml_td_typeDB.CopyBasicFieldsToXhtml_td_type(xhtml_td_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_td_typeStaged := backRepo.BackRepoXhtml_td_type.Map_Xhtml_td_typeDBID_Xhtml_td_typePtr[xhtml_td_typeDB.ID]
	if xhtml_td_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_td_typeStaged, xhtml_td_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
