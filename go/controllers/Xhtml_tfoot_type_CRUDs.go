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
var __Xhtml_tfoot_type__dummysDeclaration__ models.Xhtml_tfoot_type
var __Xhtml_tfoot_type_time__dummyDeclaration time.Duration

var mutexXhtml_tfoot_type sync.Mutex

// An Xhtml_tfoot_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_tfoot_type updateXhtml_tfoot_type deleteXhtml_tfoot_type
type Xhtml_tfoot_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_tfoot_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_tfoot_type updateXhtml_tfoot_type
type Xhtml_tfoot_typeInput struct {
	// The Xhtml_tfoot_type to submit or modify
	// in: body
	Xhtml_tfoot_type *orm.Xhtml_tfoot_typeAPI
}

// GetXhtml_tfoot_types
//
// swagger:route GET /xhtml_tfoot_types xhtml_tfoot_types getXhtml_tfoot_types
//
// # Get all xhtml_tfoot_types
//
// Responses:
// default: genericError
//
//	200: xhtml_tfoot_typeDBResponse
func (controller *Controller) GetXhtml_tfoot_types(c *gin.Context) {

	// source slice
	var xhtml_tfoot_typeDBs []orm.Xhtml_tfoot_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_tfoot_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tfoot_type.GetDB()

	_, err := db.Find(&xhtml_tfoot_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_tfoot_typeAPIs := make([]orm.Xhtml_tfoot_typeAPI, 0)

	// for each xhtml_tfoot_type, update fields from the database nullable fields
	for idx := range xhtml_tfoot_typeDBs {
		xhtml_tfoot_typeDB := &xhtml_tfoot_typeDBs[idx]
		_ = xhtml_tfoot_typeDB
		var xhtml_tfoot_typeAPI orm.Xhtml_tfoot_typeAPI

		// insertion point for updating fields
		xhtml_tfoot_typeAPI.ID = xhtml_tfoot_typeDB.ID
		xhtml_tfoot_typeDB.CopyBasicFieldsToXhtml_tfoot_type_WOP(&xhtml_tfoot_typeAPI.Xhtml_tfoot_type_WOP)
		xhtml_tfoot_typeAPI.Xhtml_tfoot_typePointersEncoding = xhtml_tfoot_typeDB.Xhtml_tfoot_typePointersEncoding
		xhtml_tfoot_typeAPIs = append(xhtml_tfoot_typeAPIs, xhtml_tfoot_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_tfoot_typeAPIs)
}

// PostXhtml_tfoot_type
//
// swagger:route POST /xhtml_tfoot_types xhtml_tfoot_types postXhtml_tfoot_type
//
// Creates a xhtml_tfoot_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_tfoot_type(c *gin.Context) {

	mutexXhtml_tfoot_type.Lock()
	defer mutexXhtml_tfoot_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_tfoot_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tfoot_type.GetDB()

	// Validate input
	var input orm.Xhtml_tfoot_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_tfoot_type
	xhtml_tfoot_typeDB := orm.Xhtml_tfoot_typeDB{}
	xhtml_tfoot_typeDB.Xhtml_tfoot_typePointersEncoding = input.Xhtml_tfoot_typePointersEncoding
	xhtml_tfoot_typeDB.CopyBasicFieldsFromXhtml_tfoot_type_WOP(&input.Xhtml_tfoot_type_WOP)

	_, err = db.Create(&xhtml_tfoot_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_tfoot_type.CheckoutPhaseOneInstance(&xhtml_tfoot_typeDB)
	xhtml_tfoot_type := backRepo.BackRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr[xhtml_tfoot_typeDB.ID]

	if xhtml_tfoot_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_tfoot_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_tfoot_typeDB)
}

// GetXhtml_tfoot_type
//
// swagger:route GET /xhtml_tfoot_types/{ID} xhtml_tfoot_types getXhtml_tfoot_type
//
// Gets the details for a xhtml_tfoot_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_tfoot_typeDBResponse
func (controller *Controller) GetXhtml_tfoot_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_tfoot_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tfoot_type.GetDB()

	// Get xhtml_tfoot_typeDB in DB
	var xhtml_tfoot_typeDB orm.Xhtml_tfoot_typeDB
	if _, err := db.First(&xhtml_tfoot_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_tfoot_typeAPI orm.Xhtml_tfoot_typeAPI
	xhtml_tfoot_typeAPI.ID = xhtml_tfoot_typeDB.ID
	xhtml_tfoot_typeAPI.Xhtml_tfoot_typePointersEncoding = xhtml_tfoot_typeDB.Xhtml_tfoot_typePointersEncoding
	xhtml_tfoot_typeDB.CopyBasicFieldsToXhtml_tfoot_type_WOP(&xhtml_tfoot_typeAPI.Xhtml_tfoot_type_WOP)

	c.JSON(http.StatusOK, xhtml_tfoot_typeAPI)
}

// UpdateXhtml_tfoot_type
//
// swagger:route PATCH /xhtml_tfoot_types/{ID} xhtml_tfoot_types updateXhtml_tfoot_type
//
// # Update a xhtml_tfoot_type
//
// Responses:
// default: genericError
//
//	200: xhtml_tfoot_typeDBResponse
func (controller *Controller) UpdateXhtml_tfoot_type(c *gin.Context) {

	mutexXhtml_tfoot_type.Lock()
	defer mutexXhtml_tfoot_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_tfoot_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tfoot_type.GetDB()

	// Validate input
	var input orm.Xhtml_tfoot_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_tfoot_typeDB orm.Xhtml_tfoot_typeDB

	// fetch the xhtml_tfoot_type
	_, err := db.First(&xhtml_tfoot_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_tfoot_typeDB.CopyBasicFieldsFromXhtml_tfoot_type_WOP(&input.Xhtml_tfoot_type_WOP)
	xhtml_tfoot_typeDB.Xhtml_tfoot_typePointersEncoding = input.Xhtml_tfoot_typePointersEncoding

	db, _ = db.Model(&xhtml_tfoot_typeDB)
	_, err = db.Updates(&xhtml_tfoot_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_tfoot_typeNew := new(models.Xhtml_tfoot_type)
	xhtml_tfoot_typeDB.CopyBasicFieldsToXhtml_tfoot_type(xhtml_tfoot_typeNew)

	// redeem pointers
	xhtml_tfoot_typeDB.DecodePointers(backRepo, xhtml_tfoot_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_tfoot_typeOld := backRepo.BackRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr[xhtml_tfoot_typeDB.ID]
	if xhtml_tfoot_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_tfoot_typeOld, xhtml_tfoot_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_tfoot_typeDB
	c.JSON(http.StatusOK, xhtml_tfoot_typeDB)
}

// DeleteXhtml_tfoot_type
//
// swagger:route DELETE /xhtml_tfoot_types/{ID} xhtml_tfoot_types deleteXhtml_tfoot_type
//
// # Delete a xhtml_tfoot_type
//
// default: genericError
//
//	200: xhtml_tfoot_typeDBResponse
func (controller *Controller) DeleteXhtml_tfoot_type(c *gin.Context) {

	mutexXhtml_tfoot_type.Lock()
	defer mutexXhtml_tfoot_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_tfoot_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tfoot_type.GetDB()

	// Get model if exist
	var xhtml_tfoot_typeDB orm.Xhtml_tfoot_typeDB
	if _, err := db.First(&xhtml_tfoot_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_tfoot_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_tfoot_typeDeleted := new(models.Xhtml_tfoot_type)
	xhtml_tfoot_typeDB.CopyBasicFieldsToXhtml_tfoot_type(xhtml_tfoot_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_tfoot_typeStaged := backRepo.BackRepoXhtml_tfoot_type.Map_Xhtml_tfoot_typeDBID_Xhtml_tfoot_typePtr[xhtml_tfoot_typeDB.ID]
	if xhtml_tfoot_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_tfoot_typeStaged, xhtml_tfoot_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
