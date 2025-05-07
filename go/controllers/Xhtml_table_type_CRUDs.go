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
var __Xhtml_table_type__dummysDeclaration__ models.Xhtml_table_type
var __Xhtml_table_type_time__dummyDeclaration time.Duration

var mutexXhtml_table_type sync.Mutex

// An Xhtml_table_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_table_type updateXhtml_table_type deleteXhtml_table_type
type Xhtml_table_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_table_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_table_type updateXhtml_table_type
type Xhtml_table_typeInput struct {
	// The Xhtml_table_type to submit or modify
	// in: body
	Xhtml_table_type *orm.Xhtml_table_typeAPI
}

// GetXhtml_table_types
//
// swagger:route GET /xhtml_table_types xhtml_table_types getXhtml_table_types
//
// # Get all xhtml_table_types
//
// Responses:
// default: genericError
//
//	200: xhtml_table_typeDBResponse
func (controller *Controller) GetXhtml_table_types(c *gin.Context) {

	// source slice
	var xhtml_table_typeDBs []orm.Xhtml_table_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_table_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_table_type.GetDB()

	_, err := db.Find(&xhtml_table_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_table_typeAPIs := make([]orm.Xhtml_table_typeAPI, 0)

	// for each xhtml_table_type, update fields from the database nullable fields
	for idx := range xhtml_table_typeDBs {
		xhtml_table_typeDB := &xhtml_table_typeDBs[idx]
		_ = xhtml_table_typeDB
		var xhtml_table_typeAPI orm.Xhtml_table_typeAPI

		// insertion point for updating fields
		xhtml_table_typeAPI.ID = xhtml_table_typeDB.ID
		xhtml_table_typeDB.CopyBasicFieldsToXhtml_table_type_WOP(&xhtml_table_typeAPI.Xhtml_table_type_WOP)
		xhtml_table_typeAPI.Xhtml_table_typePointersEncoding = xhtml_table_typeDB.Xhtml_table_typePointersEncoding
		xhtml_table_typeAPIs = append(xhtml_table_typeAPIs, xhtml_table_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_table_typeAPIs)
}

// PostXhtml_table_type
//
// swagger:route POST /xhtml_table_types xhtml_table_types postXhtml_table_type
//
// Creates a xhtml_table_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_table_type(c *gin.Context) {

	mutexXhtml_table_type.Lock()
	defer mutexXhtml_table_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_table_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_table_type.GetDB()

	// Validate input
	var input orm.Xhtml_table_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_table_type
	xhtml_table_typeDB := orm.Xhtml_table_typeDB{}
	xhtml_table_typeDB.Xhtml_table_typePointersEncoding = input.Xhtml_table_typePointersEncoding
	xhtml_table_typeDB.CopyBasicFieldsFromXhtml_table_type_WOP(&input.Xhtml_table_type_WOP)

	_, err = db.Create(&xhtml_table_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_table_type.CheckoutPhaseOneInstance(&xhtml_table_typeDB)
	xhtml_table_type := backRepo.BackRepoXhtml_table_type.Map_Xhtml_table_typeDBID_Xhtml_table_typePtr[xhtml_table_typeDB.ID]

	if xhtml_table_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_table_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_table_typeDB)
}

// GetXhtml_table_type
//
// swagger:route GET /xhtml_table_types/{ID} xhtml_table_types getXhtml_table_type
//
// Gets the details for a xhtml_table_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_table_typeDBResponse
func (controller *Controller) GetXhtml_table_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_table_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_table_type.GetDB()

	// Get xhtml_table_typeDB in DB
	var xhtml_table_typeDB orm.Xhtml_table_typeDB
	if _, err := db.First(&xhtml_table_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_table_typeAPI orm.Xhtml_table_typeAPI
	xhtml_table_typeAPI.ID = xhtml_table_typeDB.ID
	xhtml_table_typeAPI.Xhtml_table_typePointersEncoding = xhtml_table_typeDB.Xhtml_table_typePointersEncoding
	xhtml_table_typeDB.CopyBasicFieldsToXhtml_table_type_WOP(&xhtml_table_typeAPI.Xhtml_table_type_WOP)

	c.JSON(http.StatusOK, xhtml_table_typeAPI)
}

// UpdateXhtml_table_type
//
// swagger:route PATCH /xhtml_table_types/{ID} xhtml_table_types updateXhtml_table_type
//
// # Update a xhtml_table_type
//
// Responses:
// default: genericError
//
//	200: xhtml_table_typeDBResponse
func (controller *Controller) UpdateXhtml_table_type(c *gin.Context) {

	mutexXhtml_table_type.Lock()
	defer mutexXhtml_table_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_table_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_table_type.GetDB()

	// Validate input
	var input orm.Xhtml_table_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_table_typeDB orm.Xhtml_table_typeDB

	// fetch the xhtml_table_type
	_, err := db.First(&xhtml_table_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_table_typeDB.CopyBasicFieldsFromXhtml_table_type_WOP(&input.Xhtml_table_type_WOP)
	xhtml_table_typeDB.Xhtml_table_typePointersEncoding = input.Xhtml_table_typePointersEncoding

	db, _ = db.Model(&xhtml_table_typeDB)
	_, err = db.Updates(&xhtml_table_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_table_typeNew := new(models.Xhtml_table_type)
	xhtml_table_typeDB.CopyBasicFieldsToXhtml_table_type(xhtml_table_typeNew)

	// redeem pointers
	xhtml_table_typeDB.DecodePointers(backRepo, xhtml_table_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_table_typeOld := backRepo.BackRepoXhtml_table_type.Map_Xhtml_table_typeDBID_Xhtml_table_typePtr[xhtml_table_typeDB.ID]
	if xhtml_table_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_table_typeOld, xhtml_table_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_table_typeDB
	c.JSON(http.StatusOK, xhtml_table_typeDB)
}

// DeleteXhtml_table_type
//
// swagger:route DELETE /xhtml_table_types/{ID} xhtml_table_types deleteXhtml_table_type
//
// # Delete a xhtml_table_type
//
// default: genericError
//
//	200: xhtml_table_typeDBResponse
func (controller *Controller) DeleteXhtml_table_type(c *gin.Context) {

	mutexXhtml_table_type.Lock()
	defer mutexXhtml_table_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_table_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_table_type.GetDB()

	// Get model if exist
	var xhtml_table_typeDB orm.Xhtml_table_typeDB
	if _, err := db.First(&xhtml_table_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_table_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_table_typeDeleted := new(models.Xhtml_table_type)
	xhtml_table_typeDB.CopyBasicFieldsToXhtml_table_type(xhtml_table_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_table_typeStaged := backRepo.BackRepoXhtml_table_type.Map_Xhtml_table_typeDBID_Xhtml_table_typePtr[xhtml_table_typeDB.ID]
	if xhtml_table_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_table_typeStaged, xhtml_table_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
