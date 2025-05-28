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
var __GeneratedImageMetamodel__dummysDeclaration__ models.GeneratedImageMetamodel
var __GeneratedImageMetamodel_time__dummyDeclaration time.Duration

var mutexGeneratedImageMetamodel sync.Mutex

// An GeneratedImageMetamodelID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGeneratedImageMetamodel updateGeneratedImageMetamodel deleteGeneratedImageMetamodel
type GeneratedImageMetamodelID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GeneratedImageMetamodelInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGeneratedImageMetamodel updateGeneratedImageMetamodel
type GeneratedImageMetamodelInput struct {
	// The GeneratedImageMetamodel to submit or modify
	// in: body
	GeneratedImageMetamodel *orm.GeneratedImageMetamodelAPI
}

// GetGeneratedImageMetamodels
//
// swagger:route GET /generatedimagemetamodels generatedimagemetamodels getGeneratedImageMetamodels
//
// # Get all generatedimagemetamodels
//
// Responses:
// default: genericError
//
//	200: generatedimagemetamodelDBResponse
func (controller *Controller) GetGeneratedImageMetamodels(c *gin.Context) {

	// source slice
	var generatedimagemetamodelDBs []orm.GeneratedImageMetamodelDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGeneratedImageMetamodels", "Name", stackPath)
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
	db := backRepo.BackRepoGeneratedImageMetamodel.GetDB()

	_, err := db.Find(&generatedimagemetamodelDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	generatedimagemetamodelAPIs := make([]orm.GeneratedImageMetamodelAPI, 0)

	// for each generatedimagemetamodel, update fields from the database nullable fields
	for idx := range generatedimagemetamodelDBs {
		generatedimagemetamodelDB := &generatedimagemetamodelDBs[idx]
		_ = generatedimagemetamodelDB
		var generatedimagemetamodelAPI orm.GeneratedImageMetamodelAPI

		// insertion point for updating fields
		generatedimagemetamodelAPI.ID = generatedimagemetamodelDB.ID
		generatedimagemetamodelDB.CopyBasicFieldsToGeneratedImageMetamodel_WOP(&generatedimagemetamodelAPI.GeneratedImageMetamodel_WOP)
		generatedimagemetamodelAPI.GeneratedImageMetamodelPointersEncoding = generatedimagemetamodelDB.GeneratedImageMetamodelPointersEncoding
		generatedimagemetamodelAPIs = append(generatedimagemetamodelAPIs, generatedimagemetamodelAPI)
	}

	c.JSON(http.StatusOK, generatedimagemetamodelAPIs)
}

// PostGeneratedImageMetamodel
//
// swagger:route POST /generatedimagemetamodels generatedimagemetamodels postGeneratedImageMetamodel
//
// Creates a generatedimagemetamodel
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostGeneratedImageMetamodel(c *gin.Context) {

	mutexGeneratedImageMetamodel.Lock()
	defer mutexGeneratedImageMetamodel.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostGeneratedImageMetamodels", "Name", stackPath)
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
	db := backRepo.BackRepoGeneratedImageMetamodel.GetDB()

	// Validate input
	var input orm.GeneratedImageMetamodelAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create generatedimagemetamodel
	generatedimagemetamodelDB := orm.GeneratedImageMetamodelDB{}
	generatedimagemetamodelDB.GeneratedImageMetamodelPointersEncoding = input.GeneratedImageMetamodelPointersEncoding
	generatedimagemetamodelDB.CopyBasicFieldsFromGeneratedImageMetamodel_WOP(&input.GeneratedImageMetamodel_WOP)

	_, err = db.Create(&generatedimagemetamodelDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoGeneratedImageMetamodel.CheckoutPhaseOneInstance(&generatedimagemetamodelDB)
	generatedimagemetamodel := backRepo.BackRepoGeneratedImageMetamodel.Map_GeneratedImageMetamodelDBID_GeneratedImageMetamodelPtr[generatedimagemetamodelDB.ID]

	if generatedimagemetamodel != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), generatedimagemetamodel)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, generatedimagemetamodelDB)
}

// GetGeneratedImageMetamodel
//
// swagger:route GET /generatedimagemetamodels/{ID} generatedimagemetamodels getGeneratedImageMetamodel
//
// Gets the details for a generatedimagemetamodel.
//
// Responses:
// default: genericError
//
//	200: generatedimagemetamodelDBResponse
func (controller *Controller) GetGeneratedImageMetamodel(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetGeneratedImageMetamodel", "Name", stackPath)
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
	db := backRepo.BackRepoGeneratedImageMetamodel.GetDB()

	// Get generatedimagemetamodelDB in DB
	var generatedimagemetamodelDB orm.GeneratedImageMetamodelDB
	if _, err := db.First(&generatedimagemetamodelDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var generatedimagemetamodelAPI orm.GeneratedImageMetamodelAPI
	generatedimagemetamodelAPI.ID = generatedimagemetamodelDB.ID
	generatedimagemetamodelAPI.GeneratedImageMetamodelPointersEncoding = generatedimagemetamodelDB.GeneratedImageMetamodelPointersEncoding
	generatedimagemetamodelDB.CopyBasicFieldsToGeneratedImageMetamodel_WOP(&generatedimagemetamodelAPI.GeneratedImageMetamodel_WOP)

	c.JSON(http.StatusOK, generatedimagemetamodelAPI)
}

// UpdateGeneratedImageMetamodel
//
// swagger:route PATCH /generatedimagemetamodels/{ID} generatedimagemetamodels updateGeneratedImageMetamodel
//
// # Update a generatedimagemetamodel
//
// Responses:
// default: genericError
//
//	200: generatedimagemetamodelDBResponse
func (controller *Controller) UpdateGeneratedImageMetamodel(c *gin.Context) {

	mutexGeneratedImageMetamodel.Lock()
	defer mutexGeneratedImageMetamodel.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateGeneratedImageMetamodel", "Name", stackPath)
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
	db := backRepo.BackRepoGeneratedImageMetamodel.GetDB()

	// Validate input
	var input orm.GeneratedImageMetamodelAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var generatedimagemetamodelDB orm.GeneratedImageMetamodelDB

	// fetch the generatedimagemetamodel
	_, err := db.First(&generatedimagemetamodelDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	generatedimagemetamodelDB.CopyBasicFieldsFromGeneratedImageMetamodel_WOP(&input.GeneratedImageMetamodel_WOP)
	generatedimagemetamodelDB.GeneratedImageMetamodelPointersEncoding = input.GeneratedImageMetamodelPointersEncoding

	db, _ = db.Model(&generatedimagemetamodelDB)
	_, err = db.Updates(&generatedimagemetamodelDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	generatedimagemetamodelNew := new(models.GeneratedImageMetamodel)
	generatedimagemetamodelDB.CopyBasicFieldsToGeneratedImageMetamodel(generatedimagemetamodelNew)

	// redeem pointers
	generatedimagemetamodelDB.DecodePointers(backRepo, generatedimagemetamodelNew)

	// get stage instance from DB instance, and call callback function
	generatedimagemetamodelOld := backRepo.BackRepoGeneratedImageMetamodel.Map_GeneratedImageMetamodelDBID_GeneratedImageMetamodelPtr[generatedimagemetamodelDB.ID]
	if generatedimagemetamodelOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), generatedimagemetamodelOld, generatedimagemetamodelNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the generatedimagemetamodelDB
	c.JSON(http.StatusOK, generatedimagemetamodelDB)
}

// DeleteGeneratedImageMetamodel
//
// swagger:route DELETE /generatedimagemetamodels/{ID} generatedimagemetamodels deleteGeneratedImageMetamodel
//
// # Delete a generatedimagemetamodel
//
// default: genericError
//
//	200: generatedimagemetamodelDBResponse
func (controller *Controller) DeleteGeneratedImageMetamodel(c *gin.Context) {

	mutexGeneratedImageMetamodel.Lock()
	defer mutexGeneratedImageMetamodel.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteGeneratedImageMetamodel", "Name", stackPath)
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
	db := backRepo.BackRepoGeneratedImageMetamodel.GetDB()

	// Get model if exist
	var generatedimagemetamodelDB orm.GeneratedImageMetamodelDB
	if _, err := db.First(&generatedimagemetamodelDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&generatedimagemetamodelDB)

	// get an instance (not staged) from DB instance, and call callback function
	generatedimagemetamodelDeleted := new(models.GeneratedImageMetamodel)
	generatedimagemetamodelDB.CopyBasicFieldsToGeneratedImageMetamodel(generatedimagemetamodelDeleted)

	// get stage instance from DB instance, and call callback function
	generatedimagemetamodelStaged := backRepo.BackRepoGeneratedImageMetamodel.Map_GeneratedImageMetamodelDBID_GeneratedImageMetamodelPtr[generatedimagemetamodelDB.ID]
	if generatedimagemetamodelStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), generatedimagemetamodelStaged, generatedimagemetamodelDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
