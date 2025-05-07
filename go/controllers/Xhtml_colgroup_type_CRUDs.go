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
var __Xhtml_colgroup_type__dummysDeclaration__ models.Xhtml_colgroup_type
var __Xhtml_colgroup_type_time__dummyDeclaration time.Duration

var mutexXhtml_colgroup_type sync.Mutex

// An Xhtml_colgroup_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_colgroup_type updateXhtml_colgroup_type deleteXhtml_colgroup_type
type Xhtml_colgroup_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_colgroup_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_colgroup_type updateXhtml_colgroup_type
type Xhtml_colgroup_typeInput struct {
	// The Xhtml_colgroup_type to submit or modify
	// in: body
	Xhtml_colgroup_type *orm.Xhtml_colgroup_typeAPI
}

// GetXhtml_colgroup_types
//
// swagger:route GET /xhtml_colgroup_types xhtml_colgroup_types getXhtml_colgroup_types
//
// # Get all xhtml_colgroup_types
//
// Responses:
// default: genericError
//
//	200: xhtml_colgroup_typeDBResponse
func (controller *Controller) GetXhtml_colgroup_types(c *gin.Context) {

	// source slice
	var xhtml_colgroup_typeDBs []orm.Xhtml_colgroup_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_colgroup_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_colgroup_type.GetDB()

	_, err := db.Find(&xhtml_colgroup_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_colgroup_typeAPIs := make([]orm.Xhtml_colgroup_typeAPI, 0)

	// for each xhtml_colgroup_type, update fields from the database nullable fields
	for idx := range xhtml_colgroup_typeDBs {
		xhtml_colgroup_typeDB := &xhtml_colgroup_typeDBs[idx]
		_ = xhtml_colgroup_typeDB
		var xhtml_colgroup_typeAPI orm.Xhtml_colgroup_typeAPI

		// insertion point for updating fields
		xhtml_colgroup_typeAPI.ID = xhtml_colgroup_typeDB.ID
		xhtml_colgroup_typeDB.CopyBasicFieldsToXhtml_colgroup_type_WOP(&xhtml_colgroup_typeAPI.Xhtml_colgroup_type_WOP)
		xhtml_colgroup_typeAPI.Xhtml_colgroup_typePointersEncoding = xhtml_colgroup_typeDB.Xhtml_colgroup_typePointersEncoding
		xhtml_colgroup_typeAPIs = append(xhtml_colgroup_typeAPIs, xhtml_colgroup_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_colgroup_typeAPIs)
}

// PostXhtml_colgroup_type
//
// swagger:route POST /xhtml_colgroup_types xhtml_colgroup_types postXhtml_colgroup_type
//
// Creates a xhtml_colgroup_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_colgroup_type(c *gin.Context) {

	mutexXhtml_colgroup_type.Lock()
	defer mutexXhtml_colgroup_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_colgroup_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_colgroup_type.GetDB()

	// Validate input
	var input orm.Xhtml_colgroup_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_colgroup_type
	xhtml_colgroup_typeDB := orm.Xhtml_colgroup_typeDB{}
	xhtml_colgroup_typeDB.Xhtml_colgroup_typePointersEncoding = input.Xhtml_colgroup_typePointersEncoding
	xhtml_colgroup_typeDB.CopyBasicFieldsFromXhtml_colgroup_type_WOP(&input.Xhtml_colgroup_type_WOP)

	_, err = db.Create(&xhtml_colgroup_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_colgroup_type.CheckoutPhaseOneInstance(&xhtml_colgroup_typeDB)
	xhtml_colgroup_type := backRepo.BackRepoXhtml_colgroup_type.Map_Xhtml_colgroup_typeDBID_Xhtml_colgroup_typePtr[xhtml_colgroup_typeDB.ID]

	if xhtml_colgroup_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_colgroup_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_colgroup_typeDB)
}

// GetXhtml_colgroup_type
//
// swagger:route GET /xhtml_colgroup_types/{ID} xhtml_colgroup_types getXhtml_colgroup_type
//
// Gets the details for a xhtml_colgroup_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_colgroup_typeDBResponse
func (controller *Controller) GetXhtml_colgroup_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_colgroup_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_colgroup_type.GetDB()

	// Get xhtml_colgroup_typeDB in DB
	var xhtml_colgroup_typeDB orm.Xhtml_colgroup_typeDB
	if _, err := db.First(&xhtml_colgroup_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_colgroup_typeAPI orm.Xhtml_colgroup_typeAPI
	xhtml_colgroup_typeAPI.ID = xhtml_colgroup_typeDB.ID
	xhtml_colgroup_typeAPI.Xhtml_colgroup_typePointersEncoding = xhtml_colgroup_typeDB.Xhtml_colgroup_typePointersEncoding
	xhtml_colgroup_typeDB.CopyBasicFieldsToXhtml_colgroup_type_WOP(&xhtml_colgroup_typeAPI.Xhtml_colgroup_type_WOP)

	c.JSON(http.StatusOK, xhtml_colgroup_typeAPI)
}

// UpdateXhtml_colgroup_type
//
// swagger:route PATCH /xhtml_colgroup_types/{ID} xhtml_colgroup_types updateXhtml_colgroup_type
//
// # Update a xhtml_colgroup_type
//
// Responses:
// default: genericError
//
//	200: xhtml_colgroup_typeDBResponse
func (controller *Controller) UpdateXhtml_colgroup_type(c *gin.Context) {

	mutexXhtml_colgroup_type.Lock()
	defer mutexXhtml_colgroup_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_colgroup_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_colgroup_type.GetDB()

	// Validate input
	var input orm.Xhtml_colgroup_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_colgroup_typeDB orm.Xhtml_colgroup_typeDB

	// fetch the xhtml_colgroup_type
	_, err := db.First(&xhtml_colgroup_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_colgroup_typeDB.CopyBasicFieldsFromXhtml_colgroup_type_WOP(&input.Xhtml_colgroup_type_WOP)
	xhtml_colgroup_typeDB.Xhtml_colgroup_typePointersEncoding = input.Xhtml_colgroup_typePointersEncoding

	db, _ = db.Model(&xhtml_colgroup_typeDB)
	_, err = db.Updates(&xhtml_colgroup_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_colgroup_typeNew := new(models.Xhtml_colgroup_type)
	xhtml_colgroup_typeDB.CopyBasicFieldsToXhtml_colgroup_type(xhtml_colgroup_typeNew)

	// redeem pointers
	xhtml_colgroup_typeDB.DecodePointers(backRepo, xhtml_colgroup_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_colgroup_typeOld := backRepo.BackRepoXhtml_colgroup_type.Map_Xhtml_colgroup_typeDBID_Xhtml_colgroup_typePtr[xhtml_colgroup_typeDB.ID]
	if xhtml_colgroup_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_colgroup_typeOld, xhtml_colgroup_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_colgroup_typeDB
	c.JSON(http.StatusOK, xhtml_colgroup_typeDB)
}

// DeleteXhtml_colgroup_type
//
// swagger:route DELETE /xhtml_colgroup_types/{ID} xhtml_colgroup_types deleteXhtml_colgroup_type
//
// # Delete a xhtml_colgroup_type
//
// default: genericError
//
//	200: xhtml_colgroup_typeDBResponse
func (controller *Controller) DeleteXhtml_colgroup_type(c *gin.Context) {

	mutexXhtml_colgroup_type.Lock()
	defer mutexXhtml_colgroup_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_colgroup_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_colgroup_type.GetDB()

	// Get model if exist
	var xhtml_colgroup_typeDB orm.Xhtml_colgroup_typeDB
	if _, err := db.First(&xhtml_colgroup_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_colgroup_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_colgroup_typeDeleted := new(models.Xhtml_colgroup_type)
	xhtml_colgroup_typeDB.CopyBasicFieldsToXhtml_colgroup_type(xhtml_colgroup_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_colgroup_typeStaged := backRepo.BackRepoXhtml_colgroup_type.Map_Xhtml_colgroup_typeDBID_Xhtml_colgroup_typePtr[xhtml_colgroup_typeDB.ID]
	if xhtml_colgroup_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_colgroup_typeStaged, xhtml_colgroup_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
