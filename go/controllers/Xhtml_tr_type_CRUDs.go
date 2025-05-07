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
var __Xhtml_tr_type__dummysDeclaration__ models.Xhtml_tr_type
var __Xhtml_tr_type_time__dummyDeclaration time.Duration

var mutexXhtml_tr_type sync.Mutex

// An Xhtml_tr_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_tr_type updateXhtml_tr_type deleteXhtml_tr_type
type Xhtml_tr_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_tr_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_tr_type updateXhtml_tr_type
type Xhtml_tr_typeInput struct {
	// The Xhtml_tr_type to submit or modify
	// in: body
	Xhtml_tr_type *orm.Xhtml_tr_typeAPI
}

// GetXhtml_tr_types
//
// swagger:route GET /xhtml_tr_types xhtml_tr_types getXhtml_tr_types
//
// # Get all xhtml_tr_types
//
// Responses:
// default: genericError
//
//	200: xhtml_tr_typeDBResponse
func (controller *Controller) GetXhtml_tr_types(c *gin.Context) {

	// source slice
	var xhtml_tr_typeDBs []orm.Xhtml_tr_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_tr_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tr_type.GetDB()

	_, err := db.Find(&xhtml_tr_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_tr_typeAPIs := make([]orm.Xhtml_tr_typeAPI, 0)

	// for each xhtml_tr_type, update fields from the database nullable fields
	for idx := range xhtml_tr_typeDBs {
		xhtml_tr_typeDB := &xhtml_tr_typeDBs[idx]
		_ = xhtml_tr_typeDB
		var xhtml_tr_typeAPI orm.Xhtml_tr_typeAPI

		// insertion point for updating fields
		xhtml_tr_typeAPI.ID = xhtml_tr_typeDB.ID
		xhtml_tr_typeDB.CopyBasicFieldsToXhtml_tr_type_WOP(&xhtml_tr_typeAPI.Xhtml_tr_type_WOP)
		xhtml_tr_typeAPI.Xhtml_tr_typePointersEncoding = xhtml_tr_typeDB.Xhtml_tr_typePointersEncoding
		xhtml_tr_typeAPIs = append(xhtml_tr_typeAPIs, xhtml_tr_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_tr_typeAPIs)
}

// PostXhtml_tr_type
//
// swagger:route POST /xhtml_tr_types xhtml_tr_types postXhtml_tr_type
//
// Creates a xhtml_tr_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_tr_type(c *gin.Context) {

	mutexXhtml_tr_type.Lock()
	defer mutexXhtml_tr_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_tr_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tr_type.GetDB()

	// Validate input
	var input orm.Xhtml_tr_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_tr_type
	xhtml_tr_typeDB := orm.Xhtml_tr_typeDB{}
	xhtml_tr_typeDB.Xhtml_tr_typePointersEncoding = input.Xhtml_tr_typePointersEncoding
	xhtml_tr_typeDB.CopyBasicFieldsFromXhtml_tr_type_WOP(&input.Xhtml_tr_type_WOP)

	_, err = db.Create(&xhtml_tr_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_tr_type.CheckoutPhaseOneInstance(&xhtml_tr_typeDB)
	xhtml_tr_type := backRepo.BackRepoXhtml_tr_type.Map_Xhtml_tr_typeDBID_Xhtml_tr_typePtr[xhtml_tr_typeDB.ID]

	if xhtml_tr_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_tr_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_tr_typeDB)
}

// GetXhtml_tr_type
//
// swagger:route GET /xhtml_tr_types/{ID} xhtml_tr_types getXhtml_tr_type
//
// Gets the details for a xhtml_tr_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_tr_typeDBResponse
func (controller *Controller) GetXhtml_tr_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_tr_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tr_type.GetDB()

	// Get xhtml_tr_typeDB in DB
	var xhtml_tr_typeDB orm.Xhtml_tr_typeDB
	if _, err := db.First(&xhtml_tr_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_tr_typeAPI orm.Xhtml_tr_typeAPI
	xhtml_tr_typeAPI.ID = xhtml_tr_typeDB.ID
	xhtml_tr_typeAPI.Xhtml_tr_typePointersEncoding = xhtml_tr_typeDB.Xhtml_tr_typePointersEncoding
	xhtml_tr_typeDB.CopyBasicFieldsToXhtml_tr_type_WOP(&xhtml_tr_typeAPI.Xhtml_tr_type_WOP)

	c.JSON(http.StatusOK, xhtml_tr_typeAPI)
}

// UpdateXhtml_tr_type
//
// swagger:route PATCH /xhtml_tr_types/{ID} xhtml_tr_types updateXhtml_tr_type
//
// # Update a xhtml_tr_type
//
// Responses:
// default: genericError
//
//	200: xhtml_tr_typeDBResponse
func (controller *Controller) UpdateXhtml_tr_type(c *gin.Context) {

	mutexXhtml_tr_type.Lock()
	defer mutexXhtml_tr_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_tr_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tr_type.GetDB()

	// Validate input
	var input orm.Xhtml_tr_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_tr_typeDB orm.Xhtml_tr_typeDB

	// fetch the xhtml_tr_type
	_, err := db.First(&xhtml_tr_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_tr_typeDB.CopyBasicFieldsFromXhtml_tr_type_WOP(&input.Xhtml_tr_type_WOP)
	xhtml_tr_typeDB.Xhtml_tr_typePointersEncoding = input.Xhtml_tr_typePointersEncoding

	db, _ = db.Model(&xhtml_tr_typeDB)
	_, err = db.Updates(&xhtml_tr_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_tr_typeNew := new(models.Xhtml_tr_type)
	xhtml_tr_typeDB.CopyBasicFieldsToXhtml_tr_type(xhtml_tr_typeNew)

	// redeem pointers
	xhtml_tr_typeDB.DecodePointers(backRepo, xhtml_tr_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_tr_typeOld := backRepo.BackRepoXhtml_tr_type.Map_Xhtml_tr_typeDBID_Xhtml_tr_typePtr[xhtml_tr_typeDB.ID]
	if xhtml_tr_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_tr_typeOld, xhtml_tr_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_tr_typeDB
	c.JSON(http.StatusOK, xhtml_tr_typeDB)
}

// DeleteXhtml_tr_type
//
// swagger:route DELETE /xhtml_tr_types/{ID} xhtml_tr_types deleteXhtml_tr_type
//
// # Delete a xhtml_tr_type
//
// default: genericError
//
//	200: xhtml_tr_typeDBResponse
func (controller *Controller) DeleteXhtml_tr_type(c *gin.Context) {

	mutexXhtml_tr_type.Lock()
	defer mutexXhtml_tr_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_tr_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_tr_type.GetDB()

	// Get model if exist
	var xhtml_tr_typeDB orm.Xhtml_tr_typeDB
	if _, err := db.First(&xhtml_tr_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_tr_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_tr_typeDeleted := new(models.Xhtml_tr_type)
	xhtml_tr_typeDB.CopyBasicFieldsToXhtml_tr_type(xhtml_tr_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_tr_typeStaged := backRepo.BackRepoXhtml_tr_type.Map_Xhtml_tr_typeDBID_Xhtml_tr_typePtr[xhtml_tr_typeDB.ID]
	if xhtml_tr_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_tr_typeStaged, xhtml_tr_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
