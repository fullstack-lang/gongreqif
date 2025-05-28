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
var __StaticWebSite__dummysDeclaration__ models.StaticWebSite
var __StaticWebSite_time__dummyDeclaration time.Duration

var mutexStaticWebSite sync.Mutex

// An StaticWebSiteID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaticWebSite updateStaticWebSite deleteStaticWebSite
type StaticWebSiteID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StaticWebSiteInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaticWebSite updateStaticWebSite
type StaticWebSiteInput struct {
	// The StaticWebSite to submit or modify
	// in: body
	StaticWebSite *orm.StaticWebSiteAPI
}

// GetStaticWebSites
//
// swagger:route GET /staticwebsites staticwebsites getStaticWebSites
//
// # Get all staticwebsites
//
// Responses:
// default: genericError
//
//	200: staticwebsiteDBResponse
func (controller *Controller) GetStaticWebSites(c *gin.Context) {

	// source slice
	var staticwebsiteDBs []orm.StaticWebSiteDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSites", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSite.GetDB()

	_, err := db.Find(&staticwebsiteDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staticwebsiteAPIs := make([]orm.StaticWebSiteAPI, 0)

	// for each staticwebsite, update fields from the database nullable fields
	for idx := range staticwebsiteDBs {
		staticwebsiteDB := &staticwebsiteDBs[idx]
		_ = staticwebsiteDB
		var staticwebsiteAPI orm.StaticWebSiteAPI

		// insertion point for updating fields
		staticwebsiteAPI.ID = staticwebsiteDB.ID
		staticwebsiteDB.CopyBasicFieldsToStaticWebSite_WOP(&staticwebsiteAPI.StaticWebSite_WOP)
		staticwebsiteAPI.StaticWebSitePointersEncoding = staticwebsiteDB.StaticWebSitePointersEncoding
		staticwebsiteAPIs = append(staticwebsiteAPIs, staticwebsiteAPI)
	}

	c.JSON(http.StatusOK, staticwebsiteAPIs)
}

// PostStaticWebSite
//
// swagger:route POST /staticwebsites staticwebsites postStaticWebSite
//
// Creates a staticwebsite
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaticWebSite(c *gin.Context) {

	mutexStaticWebSite.Lock()
	defer mutexStaticWebSite.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaticWebSites", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSite.GetDB()

	// Validate input
	var input orm.StaticWebSiteAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staticwebsite
	staticwebsiteDB := orm.StaticWebSiteDB{}
	staticwebsiteDB.StaticWebSitePointersEncoding = input.StaticWebSitePointersEncoding
	staticwebsiteDB.CopyBasicFieldsFromStaticWebSite_WOP(&input.StaticWebSite_WOP)

	_, err = db.Create(&staticwebsiteDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaticWebSite.CheckoutPhaseOneInstance(&staticwebsiteDB)
	staticwebsite := backRepo.BackRepoStaticWebSite.Map_StaticWebSiteDBID_StaticWebSitePtr[staticwebsiteDB.ID]

	if staticwebsite != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staticwebsite)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staticwebsiteDB)
}

// GetStaticWebSite
//
// swagger:route GET /staticwebsites/{ID} staticwebsites getStaticWebSite
//
// Gets the details for a staticwebsite.
//
// Responses:
// default: genericError
//
//	200: staticwebsiteDBResponse
func (controller *Controller) GetStaticWebSite(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSite", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSite.GetDB()

	// Get staticwebsiteDB in DB
	var staticwebsiteDB orm.StaticWebSiteDB
	if _, err := db.First(&staticwebsiteDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staticwebsiteAPI orm.StaticWebSiteAPI
	staticwebsiteAPI.ID = staticwebsiteDB.ID
	staticwebsiteAPI.StaticWebSitePointersEncoding = staticwebsiteDB.StaticWebSitePointersEncoding
	staticwebsiteDB.CopyBasicFieldsToStaticWebSite_WOP(&staticwebsiteAPI.StaticWebSite_WOP)

	c.JSON(http.StatusOK, staticwebsiteAPI)
}

// UpdateStaticWebSite
//
// swagger:route PATCH /staticwebsites/{ID} staticwebsites updateStaticWebSite
//
// # Update a staticwebsite
//
// Responses:
// default: genericError
//
//	200: staticwebsiteDBResponse
func (controller *Controller) UpdateStaticWebSite(c *gin.Context) {

	mutexStaticWebSite.Lock()
	defer mutexStaticWebSite.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStaticWebSite", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSite.GetDB()

	// Validate input
	var input orm.StaticWebSiteAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staticwebsiteDB orm.StaticWebSiteDB

	// fetch the staticwebsite
	_, err := db.First(&staticwebsiteDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staticwebsiteDB.CopyBasicFieldsFromStaticWebSite_WOP(&input.StaticWebSite_WOP)
	staticwebsiteDB.StaticWebSitePointersEncoding = input.StaticWebSitePointersEncoding

	db, _ = db.Model(&staticwebsiteDB)
	_, err = db.Updates(&staticwebsiteDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsiteNew := new(models.StaticWebSite)
	staticwebsiteDB.CopyBasicFieldsToStaticWebSite(staticwebsiteNew)

	// redeem pointers
	staticwebsiteDB.DecodePointers(backRepo, staticwebsiteNew)

	// get stage instance from DB instance, and call callback function
	staticwebsiteOld := backRepo.BackRepoStaticWebSite.Map_StaticWebSiteDBID_StaticWebSitePtr[staticwebsiteDB.ID]
	if staticwebsiteOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), staticwebsiteOld, staticwebsiteNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staticwebsiteDB
	c.JSON(http.StatusOK, staticwebsiteDB)
}

// DeleteStaticWebSite
//
// swagger:route DELETE /staticwebsites/{ID} staticwebsites deleteStaticWebSite
//
// # Delete a staticwebsite
//
// default: genericError
//
//	200: staticwebsiteDBResponse
func (controller *Controller) DeleteStaticWebSite(c *gin.Context) {

	mutexStaticWebSite.Lock()
	defer mutexStaticWebSite.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaticWebSite", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSite.GetDB()

	// Get model if exist
	var staticwebsiteDB orm.StaticWebSiteDB
	if _, err := db.First(&staticwebsiteDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&staticwebsiteDB)

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsiteDeleted := new(models.StaticWebSite)
	staticwebsiteDB.CopyBasicFieldsToStaticWebSite(staticwebsiteDeleted)

	// get stage instance from DB instance, and call callback function
	staticwebsiteStaged := backRepo.BackRepoStaticWebSite.Map_StaticWebSiteDBID_StaticWebSitePtr[staticwebsiteDB.ID]
	if staticwebsiteStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staticwebsiteStaged, staticwebsiteDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
