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
var __RenderingConfiguration__dummysDeclaration__ models.RenderingConfiguration
var __RenderingConfiguration_time__dummyDeclaration time.Duration

var mutexRenderingConfiguration sync.Mutex

// An RenderingConfigurationID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRenderingConfiguration updateRenderingConfiguration deleteRenderingConfiguration
type RenderingConfigurationID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RenderingConfigurationInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRenderingConfiguration updateRenderingConfiguration
type RenderingConfigurationInput struct {
	// The RenderingConfiguration to submit or modify
	// in: body
	RenderingConfiguration *orm.RenderingConfigurationAPI
}

// GetRenderingConfigurations
//
// swagger:route GET /renderingconfigurations renderingconfigurations getRenderingConfigurations
//
// # Get all renderingconfigurations
//
// Responses:
// default: genericError
//
//	200: renderingconfigurationDBResponse
func (controller *Controller) GetRenderingConfigurations(c *gin.Context) {

	// source slice
	var renderingconfigurationDBs []orm.RenderingConfigurationDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRenderingConfigurations", "Name", stackPath)
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
	db := backRepo.BackRepoRenderingConfiguration.GetDB()

	_, err := db.Find(&renderingconfigurationDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	renderingconfigurationAPIs := make([]orm.RenderingConfigurationAPI, 0)

	// for each renderingconfiguration, update fields from the database nullable fields
	for idx := range renderingconfigurationDBs {
		renderingconfigurationDB := &renderingconfigurationDBs[idx]
		_ = renderingconfigurationDB
		var renderingconfigurationAPI orm.RenderingConfigurationAPI

		// insertion point for updating fields
		renderingconfigurationAPI.ID = renderingconfigurationDB.ID
		renderingconfigurationDB.CopyBasicFieldsToRenderingConfiguration_WOP(&renderingconfigurationAPI.RenderingConfiguration_WOP)
		renderingconfigurationAPI.RenderingConfigurationPointersEncoding = renderingconfigurationDB.RenderingConfigurationPointersEncoding
		renderingconfigurationAPIs = append(renderingconfigurationAPIs, renderingconfigurationAPI)
	}

	c.JSON(http.StatusOK, renderingconfigurationAPIs)
}

// PostRenderingConfiguration
//
// swagger:route POST /renderingconfigurations renderingconfigurations postRenderingConfiguration
//
// Creates a renderingconfiguration
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRenderingConfiguration(c *gin.Context) {

	mutexRenderingConfiguration.Lock()
	defer mutexRenderingConfiguration.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRenderingConfigurations", "Name", stackPath)
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
	db := backRepo.BackRepoRenderingConfiguration.GetDB()

	// Validate input
	var input orm.RenderingConfigurationAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create renderingconfiguration
	renderingconfigurationDB := orm.RenderingConfigurationDB{}
	renderingconfigurationDB.RenderingConfigurationPointersEncoding = input.RenderingConfigurationPointersEncoding
	renderingconfigurationDB.CopyBasicFieldsFromRenderingConfiguration_WOP(&input.RenderingConfiguration_WOP)

	_, err = db.Create(&renderingconfigurationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRenderingConfiguration.CheckoutPhaseOneInstance(&renderingconfigurationDB)
	renderingconfiguration := backRepo.BackRepoRenderingConfiguration.Map_RenderingConfigurationDBID_RenderingConfigurationPtr[renderingconfigurationDB.ID]

	if renderingconfiguration != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), renderingconfiguration)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, renderingconfigurationDB)
}

// GetRenderingConfiguration
//
// swagger:route GET /renderingconfigurations/{ID} renderingconfigurations getRenderingConfiguration
//
// Gets the details for a renderingconfiguration.
//
// Responses:
// default: genericError
//
//	200: renderingconfigurationDBResponse
func (controller *Controller) GetRenderingConfiguration(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRenderingConfiguration", "Name", stackPath)
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
	db := backRepo.BackRepoRenderingConfiguration.GetDB()

	// Get renderingconfigurationDB in DB
	var renderingconfigurationDB orm.RenderingConfigurationDB
	if _, err := db.First(&renderingconfigurationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var renderingconfigurationAPI orm.RenderingConfigurationAPI
	renderingconfigurationAPI.ID = renderingconfigurationDB.ID
	renderingconfigurationAPI.RenderingConfigurationPointersEncoding = renderingconfigurationDB.RenderingConfigurationPointersEncoding
	renderingconfigurationDB.CopyBasicFieldsToRenderingConfiguration_WOP(&renderingconfigurationAPI.RenderingConfiguration_WOP)

	c.JSON(http.StatusOK, renderingconfigurationAPI)
}

// UpdateRenderingConfiguration
//
// swagger:route PATCH /renderingconfigurations/{ID} renderingconfigurations updateRenderingConfiguration
//
// # Update a renderingconfiguration
//
// Responses:
// default: genericError
//
//	200: renderingconfigurationDBResponse
func (controller *Controller) UpdateRenderingConfiguration(c *gin.Context) {

	mutexRenderingConfiguration.Lock()
	defer mutexRenderingConfiguration.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
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
	db := backRepo.BackRepoRenderingConfiguration.GetDB()

	// Validate input
	var input orm.RenderingConfigurationAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var renderingconfigurationDB orm.RenderingConfigurationDB

	// fetch the renderingconfiguration
	_, err := db.First(&renderingconfigurationDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	renderingconfigurationDB.CopyBasicFieldsFromRenderingConfiguration_WOP(&input.RenderingConfiguration_WOP)
	renderingconfigurationDB.RenderingConfigurationPointersEncoding = input.RenderingConfigurationPointersEncoding

	db, _ = db.Model(&renderingconfigurationDB)
	_, err = db.Updates(&renderingconfigurationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	renderingconfigurationNew := new(models.RenderingConfiguration)
	renderingconfigurationDB.CopyBasicFieldsToRenderingConfiguration(renderingconfigurationNew)

	// redeem pointers
	renderingconfigurationDB.DecodePointers(backRepo, renderingconfigurationNew)

	// get stage instance from DB instance, and call callback function
	renderingconfigurationOld := backRepo.BackRepoRenderingConfiguration.Map_RenderingConfigurationDBID_RenderingConfigurationPtr[renderingconfigurationDB.ID]
	if renderingconfigurationOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), renderingconfigurationOld, renderingconfigurationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the renderingconfigurationDB
	c.JSON(http.StatusOK, renderingconfigurationDB)
}

// DeleteRenderingConfiguration
//
// swagger:route DELETE /renderingconfigurations/{ID} renderingconfigurations deleteRenderingConfiguration
//
// # Delete a renderingconfiguration
//
// default: genericError
//
//	200: renderingconfigurationDBResponse
func (controller *Controller) DeleteRenderingConfiguration(c *gin.Context) {

	mutexRenderingConfiguration.Lock()
	defer mutexRenderingConfiguration.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRenderingConfiguration", "Name", stackPath)
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
	db := backRepo.BackRepoRenderingConfiguration.GetDB()

	// Get model if exist
	var renderingconfigurationDB orm.RenderingConfigurationDB
	if _, err := db.First(&renderingconfigurationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&renderingconfigurationDB)

	// get an instance (not staged) from DB instance, and call callback function
	renderingconfigurationDeleted := new(models.RenderingConfiguration)
	renderingconfigurationDB.CopyBasicFieldsToRenderingConfiguration(renderingconfigurationDeleted)

	// get stage instance from DB instance, and call callback function
	renderingconfigurationStaged := backRepo.BackRepoRenderingConfiguration.Map_RenderingConfigurationDBID_RenderingConfigurationPtr[renderingconfigurationDB.ID]
	if renderingconfigurationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), renderingconfigurationStaged, renderingconfigurationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
