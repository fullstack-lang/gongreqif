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
var __Xhtml_h5_type__dummysDeclaration__ models.Xhtml_h5_type
var __Xhtml_h5_type_time__dummyDeclaration time.Duration

var mutexXhtml_h5_type sync.Mutex

// An Xhtml_h5_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_h5_type updateXhtml_h5_type deleteXhtml_h5_type
type Xhtml_h5_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_h5_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_h5_type updateXhtml_h5_type
type Xhtml_h5_typeInput struct {
	// The Xhtml_h5_type to submit or modify
	// in: body
	Xhtml_h5_type *orm.Xhtml_h5_typeAPI
}

// GetXhtml_h5_types
//
// swagger:route GET /xhtml_h5_types xhtml_h5_types getXhtml_h5_types
//
// # Get all xhtml_h5_types
//
// Responses:
// default: genericError
//
//	200: xhtml_h5_typeDBResponse
func (controller *Controller) GetXhtml_h5_types(c *gin.Context) {

	// source slice
	var xhtml_h5_typeDBs []orm.Xhtml_h5_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h5_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h5_type.GetDB()

	_, err := db.Find(&xhtml_h5_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_h5_typeAPIs := make([]orm.Xhtml_h5_typeAPI, 0)

	// for each xhtml_h5_type, update fields from the database nullable fields
	for idx := range xhtml_h5_typeDBs {
		xhtml_h5_typeDB := &xhtml_h5_typeDBs[idx]
		_ = xhtml_h5_typeDB
		var xhtml_h5_typeAPI orm.Xhtml_h5_typeAPI

		// insertion point for updating fields
		xhtml_h5_typeAPI.ID = xhtml_h5_typeDB.ID
		xhtml_h5_typeDB.CopyBasicFieldsToXhtml_h5_type_WOP(&xhtml_h5_typeAPI.Xhtml_h5_type_WOP)
		xhtml_h5_typeAPI.Xhtml_h5_typePointersEncoding = xhtml_h5_typeDB.Xhtml_h5_typePointersEncoding
		xhtml_h5_typeAPIs = append(xhtml_h5_typeAPIs, xhtml_h5_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_h5_typeAPIs)
}

// PostXhtml_h5_type
//
// swagger:route POST /xhtml_h5_types xhtml_h5_types postXhtml_h5_type
//
// Creates a xhtml_h5_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_h5_type(c *gin.Context) {

	mutexXhtml_h5_type.Lock()
	defer mutexXhtml_h5_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_h5_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h5_type.GetDB()

	// Validate input
	var input orm.Xhtml_h5_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_h5_type
	xhtml_h5_typeDB := orm.Xhtml_h5_typeDB{}
	xhtml_h5_typeDB.Xhtml_h5_typePointersEncoding = input.Xhtml_h5_typePointersEncoding
	xhtml_h5_typeDB.CopyBasicFieldsFromXhtml_h5_type_WOP(&input.Xhtml_h5_type_WOP)

	_, err = db.Create(&xhtml_h5_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_h5_type.CheckoutPhaseOneInstance(&xhtml_h5_typeDB)
	xhtml_h5_type := backRepo.BackRepoXhtml_h5_type.Map_Xhtml_h5_typeDBID_Xhtml_h5_typePtr[xhtml_h5_typeDB.ID]

	if xhtml_h5_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_h5_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_h5_typeDB)
}

// GetXhtml_h5_type
//
// swagger:route GET /xhtml_h5_types/{ID} xhtml_h5_types getXhtml_h5_type
//
// Gets the details for a xhtml_h5_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_h5_typeDBResponse
func (controller *Controller) GetXhtml_h5_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_h5_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h5_type.GetDB()

	// Get xhtml_h5_typeDB in DB
	var xhtml_h5_typeDB orm.Xhtml_h5_typeDB
	if _, err := db.First(&xhtml_h5_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_h5_typeAPI orm.Xhtml_h5_typeAPI
	xhtml_h5_typeAPI.ID = xhtml_h5_typeDB.ID
	xhtml_h5_typeAPI.Xhtml_h5_typePointersEncoding = xhtml_h5_typeDB.Xhtml_h5_typePointersEncoding
	xhtml_h5_typeDB.CopyBasicFieldsToXhtml_h5_type_WOP(&xhtml_h5_typeAPI.Xhtml_h5_type_WOP)

	c.JSON(http.StatusOK, xhtml_h5_typeAPI)
}

// UpdateXhtml_h5_type
//
// swagger:route PATCH /xhtml_h5_types/{ID} xhtml_h5_types updateXhtml_h5_type
//
// # Update a xhtml_h5_type
//
// Responses:
// default: genericError
//
//	200: xhtml_h5_typeDBResponse
func (controller *Controller) UpdateXhtml_h5_type(c *gin.Context) {

	mutexXhtml_h5_type.Lock()
	defer mutexXhtml_h5_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_h5_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h5_type.GetDB()

	// Validate input
	var input orm.Xhtml_h5_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_h5_typeDB orm.Xhtml_h5_typeDB

	// fetch the xhtml_h5_type
	_, err := db.First(&xhtml_h5_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_h5_typeDB.CopyBasicFieldsFromXhtml_h5_type_WOP(&input.Xhtml_h5_type_WOP)
	xhtml_h5_typeDB.Xhtml_h5_typePointersEncoding = input.Xhtml_h5_typePointersEncoding

	db, _ = db.Model(&xhtml_h5_typeDB)
	_, err = db.Updates(&xhtml_h5_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h5_typeNew := new(models.Xhtml_h5_type)
	xhtml_h5_typeDB.CopyBasicFieldsToXhtml_h5_type(xhtml_h5_typeNew)

	// redeem pointers
	xhtml_h5_typeDB.DecodePointers(backRepo, xhtml_h5_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_h5_typeOld := backRepo.BackRepoXhtml_h5_type.Map_Xhtml_h5_typeDBID_Xhtml_h5_typePtr[xhtml_h5_typeDB.ID]
	if xhtml_h5_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_h5_typeOld, xhtml_h5_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_h5_typeDB
	c.JSON(http.StatusOK, xhtml_h5_typeDB)
}

// DeleteXhtml_h5_type
//
// swagger:route DELETE /xhtml_h5_types/{ID} xhtml_h5_types deleteXhtml_h5_type
//
// # Delete a xhtml_h5_type
//
// default: genericError
//
//	200: xhtml_h5_typeDBResponse
func (controller *Controller) DeleteXhtml_h5_type(c *gin.Context) {

	mutexXhtml_h5_type.Lock()
	defer mutexXhtml_h5_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_h5_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_h5_type.GetDB()

	// Get model if exist
	var xhtml_h5_typeDB orm.Xhtml_h5_typeDB
	if _, err := db.First(&xhtml_h5_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_h5_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_h5_typeDeleted := new(models.Xhtml_h5_type)
	xhtml_h5_typeDB.CopyBasicFieldsToXhtml_h5_type(xhtml_h5_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_h5_typeStaged := backRepo.BackRepoXhtml_h5_type.Map_Xhtml_h5_typeDBID_Xhtml_h5_typePtr[xhtml_h5_typeDB.ID]
	if xhtml_h5_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_h5_typeStaged, xhtml_h5_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
