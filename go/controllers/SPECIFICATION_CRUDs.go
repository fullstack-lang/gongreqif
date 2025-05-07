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
var __SPECIFICATION__dummysDeclaration__ models.SPECIFICATION
var __SPECIFICATION_time__dummyDeclaration time.Duration

var mutexSPECIFICATION sync.Mutex

// An SPECIFICATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECIFICATION updateSPECIFICATION deleteSPECIFICATION
type SPECIFICATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECIFICATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECIFICATION updateSPECIFICATION
type SPECIFICATIONInput struct {
	// The SPECIFICATION to submit or modify
	// in: body
	SPECIFICATION *orm.SPECIFICATIONAPI
}

// GetSPECIFICATIONs
//
// swagger:route GET /specifications specifications getSPECIFICATIONs
//
// # Get all specifications
//
// Responses:
// default: genericError
//
//	200: specificationDBResponse
func (controller *Controller) GetSPECIFICATIONs(c *gin.Context) {

	// source slice
	var specificationDBs []orm.SPECIFICATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFICATIONs", "Name", stackPath)
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
	db := backRepo.BackRepoSPECIFICATION.GetDB()

	_, err := db.Find(&specificationDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specificationAPIs := make([]orm.SPECIFICATIONAPI, 0)

	// for each specification, update fields from the database nullable fields
	for idx := range specificationDBs {
		specificationDB := &specificationDBs[idx]
		_ = specificationDB
		var specificationAPI orm.SPECIFICATIONAPI

		// insertion point for updating fields
		specificationAPI.ID = specificationDB.ID
		specificationDB.CopyBasicFieldsToSPECIFICATION_WOP(&specificationAPI.SPECIFICATION_WOP)
		specificationAPI.SPECIFICATIONPointersEncoding = specificationDB.SPECIFICATIONPointersEncoding
		specificationAPIs = append(specificationAPIs, specificationAPI)
	}

	c.JSON(http.StatusOK, specificationAPIs)
}

// PostSPECIFICATION
//
// swagger:route POST /specifications specifications postSPECIFICATION
//
// Creates a specification
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECIFICATION(c *gin.Context) {

	mutexSPECIFICATION.Lock()
	defer mutexSPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECIFICATIONs", "Name", stackPath)
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
	db := backRepo.BackRepoSPECIFICATION.GetDB()

	// Validate input
	var input orm.SPECIFICATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specification
	specificationDB := orm.SPECIFICATIONDB{}
	specificationDB.SPECIFICATIONPointersEncoding = input.SPECIFICATIONPointersEncoding
	specificationDB.CopyBasicFieldsFromSPECIFICATION_WOP(&input.SPECIFICATION_WOP)

	_, err = db.Create(&specificationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECIFICATION.CheckoutPhaseOneInstance(&specificationDB)
	specification := backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr[specificationDB.ID]

	if specification != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specification)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specificationDB)
}

// GetSPECIFICATION
//
// swagger:route GET /specifications/{ID} specifications getSPECIFICATION
//
// Gets the details for a specification.
//
// Responses:
// default: genericError
//
//	200: specificationDBResponse
func (controller *Controller) GetSPECIFICATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFICATION", "Name", stackPath)
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
	db := backRepo.BackRepoSPECIFICATION.GetDB()

	// Get specificationDB in DB
	var specificationDB orm.SPECIFICATIONDB
	if _, err := db.First(&specificationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specificationAPI orm.SPECIFICATIONAPI
	specificationAPI.ID = specificationDB.ID
	specificationAPI.SPECIFICATIONPointersEncoding = specificationDB.SPECIFICATIONPointersEncoding
	specificationDB.CopyBasicFieldsToSPECIFICATION_WOP(&specificationAPI.SPECIFICATION_WOP)

	c.JSON(http.StatusOK, specificationAPI)
}

// UpdateSPECIFICATION
//
// swagger:route PATCH /specifications/{ID} specifications updateSPECIFICATION
//
// # Update a specification
//
// Responses:
// default: genericError
//
//	200: specificationDBResponse
func (controller *Controller) UpdateSPECIFICATION(c *gin.Context) {

	mutexSPECIFICATION.Lock()
	defer mutexSPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPECIFICATION", "Name", stackPath)
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
	db := backRepo.BackRepoSPECIFICATION.GetDB()

	// Validate input
	var input orm.SPECIFICATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specificationDB orm.SPECIFICATIONDB

	// fetch the specification
	_, err := db.First(&specificationDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specificationDB.CopyBasicFieldsFromSPECIFICATION_WOP(&input.SPECIFICATION_WOP)
	specificationDB.SPECIFICATIONPointersEncoding = input.SPECIFICATIONPointersEncoding

	db, _ = db.Model(&specificationDB)
	_, err = db.Updates(&specificationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specificationNew := new(models.SPECIFICATION)
	specificationDB.CopyBasicFieldsToSPECIFICATION(specificationNew)

	// redeem pointers
	specificationDB.DecodePointers(backRepo, specificationNew)

	// get stage instance from DB instance, and call callback function
	specificationOld := backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr[specificationDB.ID]
	if specificationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), specificationOld, specificationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specificationDB
	c.JSON(http.StatusOK, specificationDB)
}

// DeleteSPECIFICATION
//
// swagger:route DELETE /specifications/{ID} specifications deleteSPECIFICATION
//
// # Delete a specification
//
// default: genericError
//
//	200: specificationDBResponse
func (controller *Controller) DeleteSPECIFICATION(c *gin.Context) {

	mutexSPECIFICATION.Lock()
	defer mutexSPECIFICATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECIFICATION", "Name", stackPath)
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
	db := backRepo.BackRepoSPECIFICATION.GetDB()

	// Get model if exist
	var specificationDB orm.SPECIFICATIONDB
	if _, err := db.First(&specificationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&specificationDB)

	// get an instance (not staged) from DB instance, and call callback function
	specificationDeleted := new(models.SPECIFICATION)
	specificationDB.CopyBasicFieldsToSPECIFICATION(specificationDeleted)

	// get stage instance from DB instance, and call callback function
	specificationStaged := backRepo.BackRepoSPECIFICATION.Map_SPECIFICATIONDBID_SPECIFICATIONPtr[specificationDB.ID]
	if specificationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specificationStaged, specificationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
