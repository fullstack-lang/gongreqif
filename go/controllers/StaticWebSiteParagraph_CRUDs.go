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
var __StaticWebSiteParagraph__dummysDeclaration__ models.StaticWebSiteParagraph
var __StaticWebSiteParagraph_time__dummyDeclaration time.Duration

var mutexStaticWebSiteParagraph sync.Mutex

// An StaticWebSiteParagraphID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStaticWebSiteParagraph updateStaticWebSiteParagraph deleteStaticWebSiteParagraph
type StaticWebSiteParagraphID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StaticWebSiteParagraphInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStaticWebSiteParagraph updateStaticWebSiteParagraph
type StaticWebSiteParagraphInput struct {
	// The StaticWebSiteParagraph to submit or modify
	// in: body
	StaticWebSiteParagraph *orm.StaticWebSiteParagraphAPI
}

// GetStaticWebSiteParagraphs
//
// swagger:route GET /staticwebsiteparagraphs staticwebsiteparagraphs getStaticWebSiteParagraphs
//
// # Get all staticwebsiteparagraphs
//
// Responses:
// default: genericError
//
//	200: staticwebsiteparagraphDBResponse
func (controller *Controller) GetStaticWebSiteParagraphs(c *gin.Context) {

	// source slice
	var staticwebsiteparagraphDBs []orm.StaticWebSiteParagraphDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSiteParagraphs", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteParagraph.GetDB()

	_, err := db.Find(&staticwebsiteparagraphDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	staticwebsiteparagraphAPIs := make([]orm.StaticWebSiteParagraphAPI, 0)

	// for each staticwebsiteparagraph, update fields from the database nullable fields
	for idx := range staticwebsiteparagraphDBs {
		staticwebsiteparagraphDB := &staticwebsiteparagraphDBs[idx]
		_ = staticwebsiteparagraphDB
		var staticwebsiteparagraphAPI orm.StaticWebSiteParagraphAPI

		// insertion point for updating fields
		staticwebsiteparagraphAPI.ID = staticwebsiteparagraphDB.ID
		staticwebsiteparagraphDB.CopyBasicFieldsToStaticWebSiteParagraph_WOP(&staticwebsiteparagraphAPI.StaticWebSiteParagraph_WOP)
		staticwebsiteparagraphAPI.StaticWebSiteParagraphPointersEncoding = staticwebsiteparagraphDB.StaticWebSiteParagraphPointersEncoding
		staticwebsiteparagraphAPIs = append(staticwebsiteparagraphAPIs, staticwebsiteparagraphAPI)
	}

	c.JSON(http.StatusOK, staticwebsiteparagraphAPIs)
}

// PostStaticWebSiteParagraph
//
// swagger:route POST /staticwebsiteparagraphs staticwebsiteparagraphs postStaticWebSiteParagraph
//
// Creates a staticwebsiteparagraph
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostStaticWebSiteParagraph(c *gin.Context) {

	mutexStaticWebSiteParagraph.Lock()
	defer mutexStaticWebSiteParagraph.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostStaticWebSiteParagraphs", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteParagraph.GetDB()

	// Validate input
	var input orm.StaticWebSiteParagraphAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create staticwebsiteparagraph
	staticwebsiteparagraphDB := orm.StaticWebSiteParagraphDB{}
	staticwebsiteparagraphDB.StaticWebSiteParagraphPointersEncoding = input.StaticWebSiteParagraphPointersEncoding
	staticwebsiteparagraphDB.CopyBasicFieldsFromStaticWebSiteParagraph_WOP(&input.StaticWebSiteParagraph_WOP)

	_, err = db.Create(&staticwebsiteparagraphDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoStaticWebSiteParagraph.CheckoutPhaseOneInstance(&staticwebsiteparagraphDB)
	staticwebsiteparagraph := backRepo.BackRepoStaticWebSiteParagraph.Map_StaticWebSiteParagraphDBID_StaticWebSiteParagraphPtr[staticwebsiteparagraphDB.ID]

	if staticwebsiteparagraph != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), staticwebsiteparagraph)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, staticwebsiteparagraphDB)
}

// GetStaticWebSiteParagraph
//
// swagger:route GET /staticwebsiteparagraphs/{ID} staticwebsiteparagraphs getStaticWebSiteParagraph
//
// Gets the details for a staticwebsiteparagraph.
//
// Responses:
// default: genericError
//
//	200: staticwebsiteparagraphDBResponse
func (controller *Controller) GetStaticWebSiteParagraph(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetStaticWebSiteParagraph", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteParagraph.GetDB()

	// Get staticwebsiteparagraphDB in DB
	var staticwebsiteparagraphDB orm.StaticWebSiteParagraphDB
	if _, err := db.First(&staticwebsiteparagraphDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var staticwebsiteparagraphAPI orm.StaticWebSiteParagraphAPI
	staticwebsiteparagraphAPI.ID = staticwebsiteparagraphDB.ID
	staticwebsiteparagraphAPI.StaticWebSiteParagraphPointersEncoding = staticwebsiteparagraphDB.StaticWebSiteParagraphPointersEncoding
	staticwebsiteparagraphDB.CopyBasicFieldsToStaticWebSiteParagraph_WOP(&staticwebsiteparagraphAPI.StaticWebSiteParagraph_WOP)

	c.JSON(http.StatusOK, staticwebsiteparagraphAPI)
}

// UpdateStaticWebSiteParagraph
//
// swagger:route PATCH /staticwebsiteparagraphs/{ID} staticwebsiteparagraphs updateStaticWebSiteParagraph
//
// # Update a staticwebsiteparagraph
//
// Responses:
// default: genericError
//
//	200: staticwebsiteparagraphDBResponse
func (controller *Controller) UpdateStaticWebSiteParagraph(c *gin.Context) {

	mutexStaticWebSiteParagraph.Lock()
	defer mutexStaticWebSiteParagraph.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateStaticWebSiteParagraph", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteParagraph.GetDB()

	// Validate input
	var input orm.StaticWebSiteParagraphAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var staticwebsiteparagraphDB orm.StaticWebSiteParagraphDB

	// fetch the staticwebsiteparagraph
	_, err := db.First(&staticwebsiteparagraphDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	staticwebsiteparagraphDB.CopyBasicFieldsFromStaticWebSiteParagraph_WOP(&input.StaticWebSiteParagraph_WOP)
	staticwebsiteparagraphDB.StaticWebSiteParagraphPointersEncoding = input.StaticWebSiteParagraphPointersEncoding

	db, _ = db.Model(&staticwebsiteparagraphDB)
	_, err = db.Updates(&staticwebsiteparagraphDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsiteparagraphNew := new(models.StaticWebSiteParagraph)
	staticwebsiteparagraphDB.CopyBasicFieldsToStaticWebSiteParagraph(staticwebsiteparagraphNew)

	// redeem pointers
	staticwebsiteparagraphDB.DecodePointers(backRepo, staticwebsiteparagraphNew)

	// get stage instance from DB instance, and call callback function
	staticwebsiteparagraphOld := backRepo.BackRepoStaticWebSiteParagraph.Map_StaticWebSiteParagraphDBID_StaticWebSiteParagraphPtr[staticwebsiteparagraphDB.ID]
	if staticwebsiteparagraphOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), staticwebsiteparagraphOld, staticwebsiteparagraphNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the staticwebsiteparagraphDB
	c.JSON(http.StatusOK, staticwebsiteparagraphDB)
}

// DeleteStaticWebSiteParagraph
//
// swagger:route DELETE /staticwebsiteparagraphs/{ID} staticwebsiteparagraphs deleteStaticWebSiteParagraph
//
// # Delete a staticwebsiteparagraph
//
// default: genericError
//
//	200: staticwebsiteparagraphDBResponse
func (controller *Controller) DeleteStaticWebSiteParagraph(c *gin.Context) {

	mutexStaticWebSiteParagraph.Lock()
	defer mutexStaticWebSiteParagraph.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteStaticWebSiteParagraph", "Name", stackPath)
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
	db := backRepo.BackRepoStaticWebSiteParagraph.GetDB()

	// Get model if exist
	var staticwebsiteparagraphDB orm.StaticWebSiteParagraphDB
	if _, err := db.First(&staticwebsiteparagraphDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&staticwebsiteparagraphDB)

	// get an instance (not staged) from DB instance, and call callback function
	staticwebsiteparagraphDeleted := new(models.StaticWebSiteParagraph)
	staticwebsiteparagraphDB.CopyBasicFieldsToStaticWebSiteParagraph(staticwebsiteparagraphDeleted)

	// get stage instance from DB instance, and call callback function
	staticwebsiteparagraphStaged := backRepo.BackRepoStaticWebSiteParagraph.Map_StaticWebSiteParagraphDBID_StaticWebSiteParagraphPtr[staticwebsiteparagraphDB.ID]
	if staticwebsiteparagraphStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), staticwebsiteparagraphStaged, staticwebsiteparagraphDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
