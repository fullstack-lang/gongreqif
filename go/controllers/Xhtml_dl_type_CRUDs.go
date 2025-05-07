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
var __Xhtml_dl_type__dummysDeclaration__ models.Xhtml_dl_type
var __Xhtml_dl_type_time__dummyDeclaration time.Duration

var mutexXhtml_dl_type sync.Mutex

// An Xhtml_dl_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_dl_type updateXhtml_dl_type deleteXhtml_dl_type
type Xhtml_dl_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_dl_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_dl_type updateXhtml_dl_type
type Xhtml_dl_typeInput struct {
	// The Xhtml_dl_type to submit or modify
	// in: body
	Xhtml_dl_type *orm.Xhtml_dl_typeAPI
}

// GetXhtml_dl_types
//
// swagger:route GET /xhtml_dl_types xhtml_dl_types getXhtml_dl_types
//
// # Get all xhtml_dl_types
//
// Responses:
// default: genericError
//
//	200: xhtml_dl_typeDBResponse
func (controller *Controller) GetXhtml_dl_types(c *gin.Context) {

	// source slice
	var xhtml_dl_typeDBs []orm.Xhtml_dl_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_dl_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dl_type.GetDB()

	_, err := db.Find(&xhtml_dl_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_dl_typeAPIs := make([]orm.Xhtml_dl_typeAPI, 0)

	// for each xhtml_dl_type, update fields from the database nullable fields
	for idx := range xhtml_dl_typeDBs {
		xhtml_dl_typeDB := &xhtml_dl_typeDBs[idx]
		_ = xhtml_dl_typeDB
		var xhtml_dl_typeAPI orm.Xhtml_dl_typeAPI

		// insertion point for updating fields
		xhtml_dl_typeAPI.ID = xhtml_dl_typeDB.ID
		xhtml_dl_typeDB.CopyBasicFieldsToXhtml_dl_type_WOP(&xhtml_dl_typeAPI.Xhtml_dl_type_WOP)
		xhtml_dl_typeAPI.Xhtml_dl_typePointersEncoding = xhtml_dl_typeDB.Xhtml_dl_typePointersEncoding
		xhtml_dl_typeAPIs = append(xhtml_dl_typeAPIs, xhtml_dl_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_dl_typeAPIs)
}

// PostXhtml_dl_type
//
// swagger:route POST /xhtml_dl_types xhtml_dl_types postXhtml_dl_type
//
// Creates a xhtml_dl_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_dl_type(c *gin.Context) {

	mutexXhtml_dl_type.Lock()
	defer mutexXhtml_dl_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_dl_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dl_type.GetDB()

	// Validate input
	var input orm.Xhtml_dl_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_dl_type
	xhtml_dl_typeDB := orm.Xhtml_dl_typeDB{}
	xhtml_dl_typeDB.Xhtml_dl_typePointersEncoding = input.Xhtml_dl_typePointersEncoding
	xhtml_dl_typeDB.CopyBasicFieldsFromXhtml_dl_type_WOP(&input.Xhtml_dl_type_WOP)

	_, err = db.Create(&xhtml_dl_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_dl_type.CheckoutPhaseOneInstance(&xhtml_dl_typeDB)
	xhtml_dl_type := backRepo.BackRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr[xhtml_dl_typeDB.ID]

	if xhtml_dl_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_dl_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_dl_typeDB)
}

// GetXhtml_dl_type
//
// swagger:route GET /xhtml_dl_types/{ID} xhtml_dl_types getXhtml_dl_type
//
// Gets the details for a xhtml_dl_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_dl_typeDBResponse
func (controller *Controller) GetXhtml_dl_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_dl_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dl_type.GetDB()

	// Get xhtml_dl_typeDB in DB
	var xhtml_dl_typeDB orm.Xhtml_dl_typeDB
	if _, err := db.First(&xhtml_dl_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_dl_typeAPI orm.Xhtml_dl_typeAPI
	xhtml_dl_typeAPI.ID = xhtml_dl_typeDB.ID
	xhtml_dl_typeAPI.Xhtml_dl_typePointersEncoding = xhtml_dl_typeDB.Xhtml_dl_typePointersEncoding
	xhtml_dl_typeDB.CopyBasicFieldsToXhtml_dl_type_WOP(&xhtml_dl_typeAPI.Xhtml_dl_type_WOP)

	c.JSON(http.StatusOK, xhtml_dl_typeAPI)
}

// UpdateXhtml_dl_type
//
// swagger:route PATCH /xhtml_dl_types/{ID} xhtml_dl_types updateXhtml_dl_type
//
// # Update a xhtml_dl_type
//
// Responses:
// default: genericError
//
//	200: xhtml_dl_typeDBResponse
func (controller *Controller) UpdateXhtml_dl_type(c *gin.Context) {

	mutexXhtml_dl_type.Lock()
	defer mutexXhtml_dl_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_dl_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dl_type.GetDB()

	// Validate input
	var input orm.Xhtml_dl_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_dl_typeDB orm.Xhtml_dl_typeDB

	// fetch the xhtml_dl_type
	_, err := db.First(&xhtml_dl_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_dl_typeDB.CopyBasicFieldsFromXhtml_dl_type_WOP(&input.Xhtml_dl_type_WOP)
	xhtml_dl_typeDB.Xhtml_dl_typePointersEncoding = input.Xhtml_dl_typePointersEncoding

	db, _ = db.Model(&xhtml_dl_typeDB)
	_, err = db.Updates(&xhtml_dl_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_dl_typeNew := new(models.Xhtml_dl_type)
	xhtml_dl_typeDB.CopyBasicFieldsToXhtml_dl_type(xhtml_dl_typeNew)

	// redeem pointers
	xhtml_dl_typeDB.DecodePointers(backRepo, xhtml_dl_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_dl_typeOld := backRepo.BackRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr[xhtml_dl_typeDB.ID]
	if xhtml_dl_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_dl_typeOld, xhtml_dl_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_dl_typeDB
	c.JSON(http.StatusOK, xhtml_dl_typeDB)
}

// DeleteXhtml_dl_type
//
// swagger:route DELETE /xhtml_dl_types/{ID} xhtml_dl_types deleteXhtml_dl_type
//
// # Delete a xhtml_dl_type
//
// default: genericError
//
//	200: xhtml_dl_typeDBResponse
func (controller *Controller) DeleteXhtml_dl_type(c *gin.Context) {

	mutexXhtml_dl_type.Lock()
	defer mutexXhtml_dl_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_dl_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_dl_type.GetDB()

	// Get model if exist
	var xhtml_dl_typeDB orm.Xhtml_dl_typeDB
	if _, err := db.First(&xhtml_dl_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_dl_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_dl_typeDeleted := new(models.Xhtml_dl_type)
	xhtml_dl_typeDB.CopyBasicFieldsToXhtml_dl_type(xhtml_dl_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_dl_typeStaged := backRepo.BackRepoXhtml_dl_type.Map_Xhtml_dl_typeDBID_Xhtml_dl_typePtr[xhtml_dl_typeDB.ID]
	if xhtml_dl_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_dl_typeStaged, xhtml_dl_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
