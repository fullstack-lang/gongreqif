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
var __Xhtml_th_type__dummysDeclaration__ models.Xhtml_th_type
var __Xhtml_th_type_time__dummyDeclaration time.Duration

var mutexXhtml_th_type sync.Mutex

// An Xhtml_th_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_th_type updateXhtml_th_type deleteXhtml_th_type
type Xhtml_th_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_th_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_th_type updateXhtml_th_type
type Xhtml_th_typeInput struct {
	// The Xhtml_th_type to submit or modify
	// in: body
	Xhtml_th_type *orm.Xhtml_th_typeAPI
}

// GetXhtml_th_types
//
// swagger:route GET /xhtml_th_types xhtml_th_types getXhtml_th_types
//
// # Get all xhtml_th_types
//
// Responses:
// default: genericError
//
//	200: xhtml_th_typeDBResponse
func (controller *Controller) GetXhtml_th_types(c *gin.Context) {

	// source slice
	var xhtml_th_typeDBs []orm.Xhtml_th_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_th_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_th_type.GetDB()

	_, err := db.Find(&xhtml_th_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_th_typeAPIs := make([]orm.Xhtml_th_typeAPI, 0)

	// for each xhtml_th_type, update fields from the database nullable fields
	for idx := range xhtml_th_typeDBs {
		xhtml_th_typeDB := &xhtml_th_typeDBs[idx]
		_ = xhtml_th_typeDB
		var xhtml_th_typeAPI orm.Xhtml_th_typeAPI

		// insertion point for updating fields
		xhtml_th_typeAPI.ID = xhtml_th_typeDB.ID
		xhtml_th_typeDB.CopyBasicFieldsToXhtml_th_type_WOP(&xhtml_th_typeAPI.Xhtml_th_type_WOP)
		xhtml_th_typeAPI.Xhtml_th_typePointersEncoding = xhtml_th_typeDB.Xhtml_th_typePointersEncoding
		xhtml_th_typeAPIs = append(xhtml_th_typeAPIs, xhtml_th_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_th_typeAPIs)
}

// PostXhtml_th_type
//
// swagger:route POST /xhtml_th_types xhtml_th_types postXhtml_th_type
//
// Creates a xhtml_th_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_th_type(c *gin.Context) {

	mutexXhtml_th_type.Lock()
	defer mutexXhtml_th_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_th_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_th_type.GetDB()

	// Validate input
	var input orm.Xhtml_th_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_th_type
	xhtml_th_typeDB := orm.Xhtml_th_typeDB{}
	xhtml_th_typeDB.Xhtml_th_typePointersEncoding = input.Xhtml_th_typePointersEncoding
	xhtml_th_typeDB.CopyBasicFieldsFromXhtml_th_type_WOP(&input.Xhtml_th_type_WOP)

	_, err = db.Create(&xhtml_th_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_th_type.CheckoutPhaseOneInstance(&xhtml_th_typeDB)
	xhtml_th_type := backRepo.BackRepoXhtml_th_type.Map_Xhtml_th_typeDBID_Xhtml_th_typePtr[xhtml_th_typeDB.ID]

	if xhtml_th_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_th_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_th_typeDB)
}

// GetXhtml_th_type
//
// swagger:route GET /xhtml_th_types/{ID} xhtml_th_types getXhtml_th_type
//
// Gets the details for a xhtml_th_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_th_typeDBResponse
func (controller *Controller) GetXhtml_th_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_th_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_th_type.GetDB()

	// Get xhtml_th_typeDB in DB
	var xhtml_th_typeDB orm.Xhtml_th_typeDB
	if _, err := db.First(&xhtml_th_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_th_typeAPI orm.Xhtml_th_typeAPI
	xhtml_th_typeAPI.ID = xhtml_th_typeDB.ID
	xhtml_th_typeAPI.Xhtml_th_typePointersEncoding = xhtml_th_typeDB.Xhtml_th_typePointersEncoding
	xhtml_th_typeDB.CopyBasicFieldsToXhtml_th_type_WOP(&xhtml_th_typeAPI.Xhtml_th_type_WOP)

	c.JSON(http.StatusOK, xhtml_th_typeAPI)
}

// UpdateXhtml_th_type
//
// swagger:route PATCH /xhtml_th_types/{ID} xhtml_th_types updateXhtml_th_type
//
// # Update a xhtml_th_type
//
// Responses:
// default: genericError
//
//	200: xhtml_th_typeDBResponse
func (controller *Controller) UpdateXhtml_th_type(c *gin.Context) {

	mutexXhtml_th_type.Lock()
	defer mutexXhtml_th_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_th_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_th_type.GetDB()

	// Validate input
	var input orm.Xhtml_th_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_th_typeDB orm.Xhtml_th_typeDB

	// fetch the xhtml_th_type
	_, err := db.First(&xhtml_th_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_th_typeDB.CopyBasicFieldsFromXhtml_th_type_WOP(&input.Xhtml_th_type_WOP)
	xhtml_th_typeDB.Xhtml_th_typePointersEncoding = input.Xhtml_th_typePointersEncoding

	db, _ = db.Model(&xhtml_th_typeDB)
	_, err = db.Updates(&xhtml_th_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_th_typeNew := new(models.Xhtml_th_type)
	xhtml_th_typeDB.CopyBasicFieldsToXhtml_th_type(xhtml_th_typeNew)

	// redeem pointers
	xhtml_th_typeDB.DecodePointers(backRepo, xhtml_th_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_th_typeOld := backRepo.BackRepoXhtml_th_type.Map_Xhtml_th_typeDBID_Xhtml_th_typePtr[xhtml_th_typeDB.ID]
	if xhtml_th_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_th_typeOld, xhtml_th_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_th_typeDB
	c.JSON(http.StatusOK, xhtml_th_typeDB)
}

// DeleteXhtml_th_type
//
// swagger:route DELETE /xhtml_th_types/{ID} xhtml_th_types deleteXhtml_th_type
//
// # Delete a xhtml_th_type
//
// default: genericError
//
//	200: xhtml_th_typeDBResponse
func (controller *Controller) DeleteXhtml_th_type(c *gin.Context) {

	mutexXhtml_th_type.Lock()
	defer mutexXhtml_th_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_th_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_th_type.GetDB()

	// Get model if exist
	var xhtml_th_typeDB orm.Xhtml_th_typeDB
	if _, err := db.First(&xhtml_th_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_th_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_th_typeDeleted := new(models.Xhtml_th_type)
	xhtml_th_typeDB.CopyBasicFieldsToXhtml_th_type(xhtml_th_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_th_typeStaged := backRepo.BackRepoXhtml_th_type.Map_Xhtml_th_typeDBID_Xhtml_th_typePtr[xhtml_th_typeDB.ID]
	if xhtml_th_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_th_typeStaged, xhtml_th_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
