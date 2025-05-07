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
var __Xhtml_div_type__dummysDeclaration__ models.Xhtml_div_type
var __Xhtml_div_type_time__dummyDeclaration time.Duration

var mutexXhtml_div_type sync.Mutex

// An Xhtml_div_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_div_type updateXhtml_div_type deleteXhtml_div_type
type Xhtml_div_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_div_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_div_type updateXhtml_div_type
type Xhtml_div_typeInput struct {
	// The Xhtml_div_type to submit or modify
	// in: body
	Xhtml_div_type *orm.Xhtml_div_typeAPI
}

// GetXhtml_div_types
//
// swagger:route GET /xhtml_div_types xhtml_div_types getXhtml_div_types
//
// # Get all xhtml_div_types
//
// Responses:
// default: genericError
//
//	200: xhtml_div_typeDBResponse
func (controller *Controller) GetXhtml_div_types(c *gin.Context) {

	// source slice
	var xhtml_div_typeDBs []orm.Xhtml_div_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_div_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_div_type.GetDB()

	_, err := db.Find(&xhtml_div_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_div_typeAPIs := make([]orm.Xhtml_div_typeAPI, 0)

	// for each xhtml_div_type, update fields from the database nullable fields
	for idx := range xhtml_div_typeDBs {
		xhtml_div_typeDB := &xhtml_div_typeDBs[idx]
		_ = xhtml_div_typeDB
		var xhtml_div_typeAPI orm.Xhtml_div_typeAPI

		// insertion point for updating fields
		xhtml_div_typeAPI.ID = xhtml_div_typeDB.ID
		xhtml_div_typeDB.CopyBasicFieldsToXhtml_div_type_WOP(&xhtml_div_typeAPI.Xhtml_div_type_WOP)
		xhtml_div_typeAPI.Xhtml_div_typePointersEncoding = xhtml_div_typeDB.Xhtml_div_typePointersEncoding
		xhtml_div_typeAPIs = append(xhtml_div_typeAPIs, xhtml_div_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_div_typeAPIs)
}

// PostXhtml_div_type
//
// swagger:route POST /xhtml_div_types xhtml_div_types postXhtml_div_type
//
// Creates a xhtml_div_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_div_type(c *gin.Context) {

	mutexXhtml_div_type.Lock()
	defer mutexXhtml_div_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_div_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_div_type.GetDB()

	// Validate input
	var input orm.Xhtml_div_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_div_type
	xhtml_div_typeDB := orm.Xhtml_div_typeDB{}
	xhtml_div_typeDB.Xhtml_div_typePointersEncoding = input.Xhtml_div_typePointersEncoding
	xhtml_div_typeDB.CopyBasicFieldsFromXhtml_div_type_WOP(&input.Xhtml_div_type_WOP)

	_, err = db.Create(&xhtml_div_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_div_type.CheckoutPhaseOneInstance(&xhtml_div_typeDB)
	xhtml_div_type := backRepo.BackRepoXhtml_div_type.Map_Xhtml_div_typeDBID_Xhtml_div_typePtr[xhtml_div_typeDB.ID]

	if xhtml_div_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_div_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_div_typeDB)
}

// GetXhtml_div_type
//
// swagger:route GET /xhtml_div_types/{ID} xhtml_div_types getXhtml_div_type
//
// Gets the details for a xhtml_div_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_div_typeDBResponse
func (controller *Controller) GetXhtml_div_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_div_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_div_type.GetDB()

	// Get xhtml_div_typeDB in DB
	var xhtml_div_typeDB orm.Xhtml_div_typeDB
	if _, err := db.First(&xhtml_div_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_div_typeAPI orm.Xhtml_div_typeAPI
	xhtml_div_typeAPI.ID = xhtml_div_typeDB.ID
	xhtml_div_typeAPI.Xhtml_div_typePointersEncoding = xhtml_div_typeDB.Xhtml_div_typePointersEncoding
	xhtml_div_typeDB.CopyBasicFieldsToXhtml_div_type_WOP(&xhtml_div_typeAPI.Xhtml_div_type_WOP)

	c.JSON(http.StatusOK, xhtml_div_typeAPI)
}

// UpdateXhtml_div_type
//
// swagger:route PATCH /xhtml_div_types/{ID} xhtml_div_types updateXhtml_div_type
//
// # Update a xhtml_div_type
//
// Responses:
// default: genericError
//
//	200: xhtml_div_typeDBResponse
func (controller *Controller) UpdateXhtml_div_type(c *gin.Context) {

	mutexXhtml_div_type.Lock()
	defer mutexXhtml_div_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_div_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_div_type.GetDB()

	// Validate input
	var input orm.Xhtml_div_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_div_typeDB orm.Xhtml_div_typeDB

	// fetch the xhtml_div_type
	_, err := db.First(&xhtml_div_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_div_typeDB.CopyBasicFieldsFromXhtml_div_type_WOP(&input.Xhtml_div_type_WOP)
	xhtml_div_typeDB.Xhtml_div_typePointersEncoding = input.Xhtml_div_typePointersEncoding

	db, _ = db.Model(&xhtml_div_typeDB)
	_, err = db.Updates(&xhtml_div_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_div_typeNew := new(models.Xhtml_div_type)
	xhtml_div_typeDB.CopyBasicFieldsToXhtml_div_type(xhtml_div_typeNew)

	// redeem pointers
	xhtml_div_typeDB.DecodePointers(backRepo, xhtml_div_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_div_typeOld := backRepo.BackRepoXhtml_div_type.Map_Xhtml_div_typeDBID_Xhtml_div_typePtr[xhtml_div_typeDB.ID]
	if xhtml_div_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_div_typeOld, xhtml_div_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_div_typeDB
	c.JSON(http.StatusOK, xhtml_div_typeDB)
}

// DeleteXhtml_div_type
//
// swagger:route DELETE /xhtml_div_types/{ID} xhtml_div_types deleteXhtml_div_type
//
// # Delete a xhtml_div_type
//
// default: genericError
//
//	200: xhtml_div_typeDBResponse
func (controller *Controller) DeleteXhtml_div_type(c *gin.Context) {

	mutexXhtml_div_type.Lock()
	defer mutexXhtml_div_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_div_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_div_type.GetDB()

	// Get model if exist
	var xhtml_div_typeDB orm.Xhtml_div_typeDB
	if _, err := db.First(&xhtml_div_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_div_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_div_typeDeleted := new(models.Xhtml_div_type)
	xhtml_div_typeDB.CopyBasicFieldsToXhtml_div_type(xhtml_div_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_div_typeStaged := backRepo.BackRepoXhtml_div_type.Map_Xhtml_div_typeDBID_Xhtml_div_typePtr[xhtml_div_typeDB.ID]
	if xhtml_div_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_div_typeStaged, xhtml_div_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
