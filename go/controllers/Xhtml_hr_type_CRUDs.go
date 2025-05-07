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
var __Xhtml_hr_type__dummysDeclaration__ models.Xhtml_hr_type
var __Xhtml_hr_type_time__dummyDeclaration time.Duration

var mutexXhtml_hr_type sync.Mutex

// An Xhtml_hr_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_hr_type updateXhtml_hr_type deleteXhtml_hr_type
type Xhtml_hr_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_hr_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_hr_type updateXhtml_hr_type
type Xhtml_hr_typeInput struct {
	// The Xhtml_hr_type to submit or modify
	// in: body
	Xhtml_hr_type *orm.Xhtml_hr_typeAPI
}

// GetXhtml_hr_types
//
// swagger:route GET /xhtml_hr_types xhtml_hr_types getXhtml_hr_types
//
// # Get all xhtml_hr_types
//
// Responses:
// default: genericError
//
//	200: xhtml_hr_typeDBResponse
func (controller *Controller) GetXhtml_hr_types(c *gin.Context) {

	// source slice
	var xhtml_hr_typeDBs []orm.Xhtml_hr_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_hr_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_hr_type.GetDB()

	_, err := db.Find(&xhtml_hr_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_hr_typeAPIs := make([]orm.Xhtml_hr_typeAPI, 0)

	// for each xhtml_hr_type, update fields from the database nullable fields
	for idx := range xhtml_hr_typeDBs {
		xhtml_hr_typeDB := &xhtml_hr_typeDBs[idx]
		_ = xhtml_hr_typeDB
		var xhtml_hr_typeAPI orm.Xhtml_hr_typeAPI

		// insertion point for updating fields
		xhtml_hr_typeAPI.ID = xhtml_hr_typeDB.ID
		xhtml_hr_typeDB.CopyBasicFieldsToXhtml_hr_type_WOP(&xhtml_hr_typeAPI.Xhtml_hr_type_WOP)
		xhtml_hr_typeAPI.Xhtml_hr_typePointersEncoding = xhtml_hr_typeDB.Xhtml_hr_typePointersEncoding
		xhtml_hr_typeAPIs = append(xhtml_hr_typeAPIs, xhtml_hr_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_hr_typeAPIs)
}

// PostXhtml_hr_type
//
// swagger:route POST /xhtml_hr_types xhtml_hr_types postXhtml_hr_type
//
// Creates a xhtml_hr_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_hr_type(c *gin.Context) {

	mutexXhtml_hr_type.Lock()
	defer mutexXhtml_hr_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_hr_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_hr_type.GetDB()

	// Validate input
	var input orm.Xhtml_hr_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_hr_type
	xhtml_hr_typeDB := orm.Xhtml_hr_typeDB{}
	xhtml_hr_typeDB.Xhtml_hr_typePointersEncoding = input.Xhtml_hr_typePointersEncoding
	xhtml_hr_typeDB.CopyBasicFieldsFromXhtml_hr_type_WOP(&input.Xhtml_hr_type_WOP)

	_, err = db.Create(&xhtml_hr_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_hr_type.CheckoutPhaseOneInstance(&xhtml_hr_typeDB)
	xhtml_hr_type := backRepo.BackRepoXhtml_hr_type.Map_Xhtml_hr_typeDBID_Xhtml_hr_typePtr[xhtml_hr_typeDB.ID]

	if xhtml_hr_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_hr_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_hr_typeDB)
}

// GetXhtml_hr_type
//
// swagger:route GET /xhtml_hr_types/{ID} xhtml_hr_types getXhtml_hr_type
//
// Gets the details for a xhtml_hr_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_hr_typeDBResponse
func (controller *Controller) GetXhtml_hr_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_hr_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_hr_type.GetDB()

	// Get xhtml_hr_typeDB in DB
	var xhtml_hr_typeDB orm.Xhtml_hr_typeDB
	if _, err := db.First(&xhtml_hr_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_hr_typeAPI orm.Xhtml_hr_typeAPI
	xhtml_hr_typeAPI.ID = xhtml_hr_typeDB.ID
	xhtml_hr_typeAPI.Xhtml_hr_typePointersEncoding = xhtml_hr_typeDB.Xhtml_hr_typePointersEncoding
	xhtml_hr_typeDB.CopyBasicFieldsToXhtml_hr_type_WOP(&xhtml_hr_typeAPI.Xhtml_hr_type_WOP)

	c.JSON(http.StatusOK, xhtml_hr_typeAPI)
}

// UpdateXhtml_hr_type
//
// swagger:route PATCH /xhtml_hr_types/{ID} xhtml_hr_types updateXhtml_hr_type
//
// # Update a xhtml_hr_type
//
// Responses:
// default: genericError
//
//	200: xhtml_hr_typeDBResponse
func (controller *Controller) UpdateXhtml_hr_type(c *gin.Context) {

	mutexXhtml_hr_type.Lock()
	defer mutexXhtml_hr_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_hr_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_hr_type.GetDB()

	// Validate input
	var input orm.Xhtml_hr_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_hr_typeDB orm.Xhtml_hr_typeDB

	// fetch the xhtml_hr_type
	_, err := db.First(&xhtml_hr_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_hr_typeDB.CopyBasicFieldsFromXhtml_hr_type_WOP(&input.Xhtml_hr_type_WOP)
	xhtml_hr_typeDB.Xhtml_hr_typePointersEncoding = input.Xhtml_hr_typePointersEncoding

	db, _ = db.Model(&xhtml_hr_typeDB)
	_, err = db.Updates(&xhtml_hr_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_hr_typeNew := new(models.Xhtml_hr_type)
	xhtml_hr_typeDB.CopyBasicFieldsToXhtml_hr_type(xhtml_hr_typeNew)

	// redeem pointers
	xhtml_hr_typeDB.DecodePointers(backRepo, xhtml_hr_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_hr_typeOld := backRepo.BackRepoXhtml_hr_type.Map_Xhtml_hr_typeDBID_Xhtml_hr_typePtr[xhtml_hr_typeDB.ID]
	if xhtml_hr_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_hr_typeOld, xhtml_hr_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_hr_typeDB
	c.JSON(http.StatusOK, xhtml_hr_typeDB)
}

// DeleteXhtml_hr_type
//
// swagger:route DELETE /xhtml_hr_types/{ID} xhtml_hr_types deleteXhtml_hr_type
//
// # Delete a xhtml_hr_type
//
// default: genericError
//
//	200: xhtml_hr_typeDBResponse
func (controller *Controller) DeleteXhtml_hr_type(c *gin.Context) {

	mutexXhtml_hr_type.Lock()
	defer mutexXhtml_hr_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_hr_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_hr_type.GetDB()

	// Get model if exist
	var xhtml_hr_typeDB orm.Xhtml_hr_typeDB
	if _, err := db.First(&xhtml_hr_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_hr_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_hr_typeDeleted := new(models.Xhtml_hr_type)
	xhtml_hr_typeDB.CopyBasicFieldsToXhtml_hr_type(xhtml_hr_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_hr_typeStaged := backRepo.BackRepoXhtml_hr_type.Map_Xhtml_hr_typeDBID_Xhtml_hr_typePtr[xhtml_hr_typeDB.ID]
	if xhtml_hr_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_hr_typeStaged, xhtml_hr_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
