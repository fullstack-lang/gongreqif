// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __MetaReference__dummysDeclaration__ models.MetaReference
var __MetaReference_time__dummyDeclaration time.Duration

var mutexMetaReference sync.Mutex

// An MetaReferenceID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getMetaReference updateMetaReference deleteMetaReference
type MetaReferenceID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// MetaReferenceInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postMetaReference updateMetaReference
type MetaReferenceInput struct {
	// The MetaReference to submit or modify
	// in: body
	MetaReference *orm.MetaReferenceAPI
}

// GetMetaReferences
//
// swagger:route GET /metareferences metareferences getMetaReferences
//
// # Get all metareferences
//
// Responses:
// default: genericError
//
//	200: metareferenceDBResponse
func (controller *Controller) GetMetaReferences(c *gin.Context) {

	// source slice
	var metareferenceDBs []orm.MetaReferenceDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetaReferences", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "GET Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMetaReference.GetDB()

	_, err := db.Find(&metareferenceDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	metareferenceAPIs := make([]orm.MetaReferenceAPI, 0)

	// for each metareference, update fields from the database nullable fields
	for idx := range metareferenceDBs {
		metareferenceDB := &metareferenceDBs[idx]
		_ = metareferenceDB
		var metareferenceAPI orm.MetaReferenceAPI

		// insertion point for updating fields
		metareferenceAPI.ID = metareferenceDB.ID
		metareferenceDB.CopyBasicFieldsToMetaReference_WOP(&metareferenceAPI.MetaReference_WOP)
		metareferenceAPI.MetaReferencePointersEncoding = metareferenceDB.MetaReferencePointersEncoding
		metareferenceAPIs = append(metareferenceAPIs, metareferenceAPI)
	}

	c.JSON(http.StatusOK, metareferenceAPIs)
}

// PostMetaReference
//
// swagger:route POST /metareferences metareferences postMetaReference
//
// Creates a metareference
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostMetaReference(c *gin.Context) {

	mutexMetaReference.Lock()
	defer mutexMetaReference.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostMetaReferences", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Post Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMetaReference.GetDB()

	// Validate input
	var input orm.MetaReferenceAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create metareference
	metareferenceDB := orm.MetaReferenceDB{}
	metareferenceDB.MetaReferencePointersEncoding = input.MetaReferencePointersEncoding
	metareferenceDB.CopyBasicFieldsFromMetaReference_WOP(&input.MetaReference_WOP)

	_, err = db.Create(&metareferenceDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoMetaReference.CheckoutPhaseOneInstance(&metareferenceDB)
	metareference := backRepo.BackRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr[metareferenceDB.ID]

	if metareference != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), metareference)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, metareferenceDB)
}

// GetMetaReference
//
// swagger:route GET /metareferences/{ID} metareferences getMetaReference
//
// Gets the details for a metareference.
//
// Responses:
// default: genericError
//
//	200: metareferenceDBResponse
func (controller *Controller) GetMetaReference(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetMetaReference", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMetaReference.GetDB()

	// Get metareferenceDB in DB
	var metareferenceDB orm.MetaReferenceDB
	if _, err := db.First(&metareferenceDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var metareferenceAPI orm.MetaReferenceAPI
	metareferenceAPI.ID = metareferenceDB.ID
	metareferenceAPI.MetaReferencePointersEncoding = metareferenceDB.MetaReferencePointersEncoding
	metareferenceDB.CopyBasicFieldsToMetaReference_WOP(&metareferenceAPI.MetaReference_WOP)

	c.JSON(http.StatusOK, metareferenceAPI)
}

// UpdateMetaReference
//
// swagger:route PATCH /metareferences/{ID} metareferences updateMetaReference
//
// # Update a metareference
//
// Responses:
// default: genericError
//
//	200: metareferenceDBResponse
func (controller *Controller) UpdateMetaReference(c *gin.Context) {

	mutexMetaReference.Lock()
	defer mutexMetaReference.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateMetaReference", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "PATCH Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMetaReference.GetDB()

	// Validate input
	var input orm.MetaReferenceAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var metareferenceDB orm.MetaReferenceDB

	// fetch the metareference
	_, err := db.First(&metareferenceDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	metareferenceDB.CopyBasicFieldsFromMetaReference_WOP(&input.MetaReference_WOP)
	metareferenceDB.MetaReferencePointersEncoding = input.MetaReferencePointersEncoding

	db, _ = db.Model(&metareferenceDB)
	_, err = db.Updates(&metareferenceDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	metareferenceNew := new(models.MetaReference)
	metareferenceDB.CopyBasicFieldsToMetaReference(metareferenceNew)

	// redeem pointers
	metareferenceDB.DecodePointers(backRepo, metareferenceNew)

	// get stage instance from DB instance, and call callback function
	metareferenceOld := backRepo.BackRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr[metareferenceDB.ID]
	if metareferenceOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), metareferenceOld, metareferenceNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the metareferenceDB
	c.JSON(http.StatusOK, metareferenceDB)
}

// DeleteMetaReference
//
// swagger:route DELETE /metareferences/{ID} metareferences deleteMetaReference
//
// # Delete a metareference
//
// default: genericError
//
//	200: metareferenceDBResponse
func (controller *Controller) DeleteMetaReference(c *gin.Context) {

	mutexMetaReference.Lock()
	defer mutexMetaReference.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteMetaReference", "Name", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		message := "DELETE Stack github.com/fullstack-lang/gong/go, Unkown stack: \"" + stackPath + "\"\n"

		message += "Availabe stack names are:\n"
		for k := range controller.Map_BackRepos {
			message += k + "\n"
		}

		log.Panic(message)
	}
	db := backRepo.BackRepoMetaReference.GetDB()

	// Get model if exist
	var metareferenceDB orm.MetaReferenceDB
	if _, err := db.First(&metareferenceDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&metareferenceDB)

	// get an instance (not staged) from DB instance, and call callback function
	metareferenceDeleted := new(models.MetaReference)
	metareferenceDB.CopyBasicFieldsToMetaReference(metareferenceDeleted)

	// get stage instance from DB instance, and call callback function
	metareferenceStaged := backRepo.BackRepoMetaReference.Map_MetaReferenceDBID_MetaReferencePtr[metareferenceDB.ID]
	if metareferenceStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), metareferenceStaged, metareferenceDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
