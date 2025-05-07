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
var __Xhtml_caption_type__dummysDeclaration__ models.Xhtml_caption_type
var __Xhtml_caption_type_time__dummyDeclaration time.Duration

var mutexXhtml_caption_type sync.Mutex

// An Xhtml_caption_typeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getXhtml_caption_type updateXhtml_caption_type deleteXhtml_caption_type
type Xhtml_caption_typeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// Xhtml_caption_typeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postXhtml_caption_type updateXhtml_caption_type
type Xhtml_caption_typeInput struct {
	// The Xhtml_caption_type to submit or modify
	// in: body
	Xhtml_caption_type *orm.Xhtml_caption_typeAPI
}

// GetXhtml_caption_types
//
// swagger:route GET /xhtml_caption_types xhtml_caption_types getXhtml_caption_types
//
// # Get all xhtml_caption_types
//
// Responses:
// default: genericError
//
//	200: xhtml_caption_typeDBResponse
func (controller *Controller) GetXhtml_caption_types(c *gin.Context) {

	// source slice
	var xhtml_caption_typeDBs []orm.Xhtml_caption_typeDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_caption_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_caption_type.GetDB()

	_, err := db.Find(&xhtml_caption_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	xhtml_caption_typeAPIs := make([]orm.Xhtml_caption_typeAPI, 0)

	// for each xhtml_caption_type, update fields from the database nullable fields
	for idx := range xhtml_caption_typeDBs {
		xhtml_caption_typeDB := &xhtml_caption_typeDBs[idx]
		_ = xhtml_caption_typeDB
		var xhtml_caption_typeAPI orm.Xhtml_caption_typeAPI

		// insertion point for updating fields
		xhtml_caption_typeAPI.ID = xhtml_caption_typeDB.ID
		xhtml_caption_typeDB.CopyBasicFieldsToXhtml_caption_type_WOP(&xhtml_caption_typeAPI.Xhtml_caption_type_WOP)
		xhtml_caption_typeAPI.Xhtml_caption_typePointersEncoding = xhtml_caption_typeDB.Xhtml_caption_typePointersEncoding
		xhtml_caption_typeAPIs = append(xhtml_caption_typeAPIs, xhtml_caption_typeAPI)
	}

	c.JSON(http.StatusOK, xhtml_caption_typeAPIs)
}

// PostXhtml_caption_type
//
// swagger:route POST /xhtml_caption_types xhtml_caption_types postXhtml_caption_type
//
// Creates a xhtml_caption_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostXhtml_caption_type(c *gin.Context) {

	mutexXhtml_caption_type.Lock()
	defer mutexXhtml_caption_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostXhtml_caption_types", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_caption_type.GetDB()

	// Validate input
	var input orm.Xhtml_caption_typeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create xhtml_caption_type
	xhtml_caption_typeDB := orm.Xhtml_caption_typeDB{}
	xhtml_caption_typeDB.Xhtml_caption_typePointersEncoding = input.Xhtml_caption_typePointersEncoding
	xhtml_caption_typeDB.CopyBasicFieldsFromXhtml_caption_type_WOP(&input.Xhtml_caption_type_WOP)

	_, err = db.Create(&xhtml_caption_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoXhtml_caption_type.CheckoutPhaseOneInstance(&xhtml_caption_typeDB)
	xhtml_caption_type := backRepo.BackRepoXhtml_caption_type.Map_Xhtml_caption_typeDBID_Xhtml_caption_typePtr[xhtml_caption_typeDB.ID]

	if xhtml_caption_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), xhtml_caption_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, xhtml_caption_typeDB)
}

// GetXhtml_caption_type
//
// swagger:route GET /xhtml_caption_types/{ID} xhtml_caption_types getXhtml_caption_type
//
// Gets the details for a xhtml_caption_type.
//
// Responses:
// default: genericError
//
//	200: xhtml_caption_typeDBResponse
func (controller *Controller) GetXhtml_caption_type(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetXhtml_caption_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_caption_type.GetDB()

	// Get xhtml_caption_typeDB in DB
	var xhtml_caption_typeDB orm.Xhtml_caption_typeDB
	if _, err := db.First(&xhtml_caption_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var xhtml_caption_typeAPI orm.Xhtml_caption_typeAPI
	xhtml_caption_typeAPI.ID = xhtml_caption_typeDB.ID
	xhtml_caption_typeAPI.Xhtml_caption_typePointersEncoding = xhtml_caption_typeDB.Xhtml_caption_typePointersEncoding
	xhtml_caption_typeDB.CopyBasicFieldsToXhtml_caption_type_WOP(&xhtml_caption_typeAPI.Xhtml_caption_type_WOP)

	c.JSON(http.StatusOK, xhtml_caption_typeAPI)
}

// UpdateXhtml_caption_type
//
// swagger:route PATCH /xhtml_caption_types/{ID} xhtml_caption_types updateXhtml_caption_type
//
// # Update a xhtml_caption_type
//
// Responses:
// default: genericError
//
//	200: xhtml_caption_typeDBResponse
func (controller *Controller) UpdateXhtml_caption_type(c *gin.Context) {

	mutexXhtml_caption_type.Lock()
	defer mutexXhtml_caption_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateXhtml_caption_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_caption_type.GetDB()

	// Validate input
	var input orm.Xhtml_caption_typeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var xhtml_caption_typeDB orm.Xhtml_caption_typeDB

	// fetch the xhtml_caption_type
	_, err := db.First(&xhtml_caption_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	xhtml_caption_typeDB.CopyBasicFieldsFromXhtml_caption_type_WOP(&input.Xhtml_caption_type_WOP)
	xhtml_caption_typeDB.Xhtml_caption_typePointersEncoding = input.Xhtml_caption_typePointersEncoding

	db, _ = db.Model(&xhtml_caption_typeDB)
	_, err = db.Updates(&xhtml_caption_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_caption_typeNew := new(models.Xhtml_caption_type)
	xhtml_caption_typeDB.CopyBasicFieldsToXhtml_caption_type(xhtml_caption_typeNew)

	// redeem pointers
	xhtml_caption_typeDB.DecodePointers(backRepo, xhtml_caption_typeNew)

	// get stage instance from DB instance, and call callback function
	xhtml_caption_typeOld := backRepo.BackRepoXhtml_caption_type.Map_Xhtml_caption_typeDBID_Xhtml_caption_typePtr[xhtml_caption_typeDB.ID]
	if xhtml_caption_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), xhtml_caption_typeOld, xhtml_caption_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the xhtml_caption_typeDB
	c.JSON(http.StatusOK, xhtml_caption_typeDB)
}

// DeleteXhtml_caption_type
//
// swagger:route DELETE /xhtml_caption_types/{ID} xhtml_caption_types deleteXhtml_caption_type
//
// # Delete a xhtml_caption_type
//
// default: genericError
//
//	200: xhtml_caption_typeDBResponse
func (controller *Controller) DeleteXhtml_caption_type(c *gin.Context) {

	mutexXhtml_caption_type.Lock()
	defer mutexXhtml_caption_type.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteXhtml_caption_type", "Name", stackPath)
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
	db := backRepo.BackRepoXhtml_caption_type.GetDB()

	// Get model if exist
	var xhtml_caption_typeDB orm.Xhtml_caption_typeDB
	if _, err := db.First(&xhtml_caption_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&xhtml_caption_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	xhtml_caption_typeDeleted := new(models.Xhtml_caption_type)
	xhtml_caption_typeDB.CopyBasicFieldsToXhtml_caption_type(xhtml_caption_typeDeleted)

	// get stage instance from DB instance, and call callback function
	xhtml_caption_typeStaged := backRepo.BackRepoXhtml_caption_type.Map_Xhtml_caption_typeDBID_Xhtml_caption_typePtr[xhtml_caption_typeDB.ID]
	if xhtml_caption_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), xhtml_caption_typeStaged, xhtml_caption_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
