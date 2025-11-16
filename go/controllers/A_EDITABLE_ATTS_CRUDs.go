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
var __A_EDITABLE_ATTS__dummysDeclaration__ models.A_EDITABLE_ATTS
var __A_EDITABLE_ATTS_time__dummyDeclaration time.Duration

var mutexA_EDITABLE_ATTS sync.Mutex

// An A_EDITABLE_ATTSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_EDITABLE_ATTS updateA_EDITABLE_ATTS deleteA_EDITABLE_ATTS
type A_EDITABLE_ATTSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_EDITABLE_ATTSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_EDITABLE_ATTS updateA_EDITABLE_ATTS
type A_EDITABLE_ATTSInput struct {
	// The A_EDITABLE_ATTS to submit or modify
	// in: body
	A_EDITABLE_ATTS *orm.A_EDITABLE_ATTSAPI
}

// GetA_EDITABLE_ATTSs
//
// swagger:route GET /a_editable_attss a_editable_attss getA_EDITABLE_ATTSs
//
// # Get all a_editable_attss
//
// Responses:
// default: genericError
//
//	200: a_editable_attsDBResponse
func (controller *Controller) GetA_EDITABLE_ATTSs(c *gin.Context) {

	// source slice
	var a_editable_attsDBs []orm.A_EDITABLE_ATTSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_EDITABLE_ATTSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_EDITABLE_ATTS.GetDB()

	_, err := db.Find(&a_editable_attsDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_editable_attsAPIs := make([]orm.A_EDITABLE_ATTSAPI, 0)

	// for each a_editable_atts, update fields from the database nullable fields
	for idx := range a_editable_attsDBs {
		a_editable_attsDB := &a_editable_attsDBs[idx]
		_ = a_editable_attsDB
		var a_editable_attsAPI orm.A_EDITABLE_ATTSAPI

		// insertion point for updating fields
		a_editable_attsAPI.ID = a_editable_attsDB.ID
		a_editable_attsDB.CopyBasicFieldsToA_EDITABLE_ATTS_WOP(&a_editable_attsAPI.A_EDITABLE_ATTS_WOP)
		a_editable_attsAPI.A_EDITABLE_ATTSPointersEncoding = a_editable_attsDB.A_EDITABLE_ATTSPointersEncoding
		a_editable_attsAPIs = append(a_editable_attsAPIs, a_editable_attsAPI)
	}

	c.JSON(http.StatusOK, a_editable_attsAPIs)
}

// PostA_EDITABLE_ATTS
//
// swagger:route POST /a_editable_attss a_editable_attss postA_EDITABLE_ATTS
//
// Creates a a_editable_atts
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_EDITABLE_ATTS(c *gin.Context) {

	mutexA_EDITABLE_ATTS.Lock()
	defer mutexA_EDITABLE_ATTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_EDITABLE_ATTSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_EDITABLE_ATTS.GetDB()

	// Validate input
	var input orm.A_EDITABLE_ATTSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_editable_atts
	a_editable_attsDB := orm.A_EDITABLE_ATTSDB{}
	a_editable_attsDB.A_EDITABLE_ATTSPointersEncoding = input.A_EDITABLE_ATTSPointersEncoding
	a_editable_attsDB.CopyBasicFieldsFromA_EDITABLE_ATTS_WOP(&input.A_EDITABLE_ATTS_WOP)

	_, err = db.Create(&a_editable_attsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_EDITABLE_ATTS.CheckoutPhaseOneInstance(&a_editable_attsDB)
	a_editable_atts := backRepo.BackRepoA_EDITABLE_ATTS.Map_A_EDITABLE_ATTSDBID_A_EDITABLE_ATTSPtr[a_editable_attsDB.ID]

	if a_editable_atts != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_editable_atts)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_editable_attsDB)
}

// GetA_EDITABLE_ATTS
//
// swagger:route GET /a_editable_attss/{ID} a_editable_attss getA_EDITABLE_ATTS
//
// Gets the details for a a_editable_atts.
//
// Responses:
// default: genericError
//
//	200: a_editable_attsDBResponse
func (controller *Controller) GetA_EDITABLE_ATTS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_EDITABLE_ATTS", "Name", stackPath)
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
	db := backRepo.BackRepoA_EDITABLE_ATTS.GetDB()

	// Get a_editable_attsDB in DB
	var a_editable_attsDB orm.A_EDITABLE_ATTSDB
	if _, err := db.First(&a_editable_attsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_editable_attsAPI orm.A_EDITABLE_ATTSAPI
	a_editable_attsAPI.ID = a_editable_attsDB.ID
	a_editable_attsAPI.A_EDITABLE_ATTSPointersEncoding = a_editable_attsDB.A_EDITABLE_ATTSPointersEncoding
	a_editable_attsDB.CopyBasicFieldsToA_EDITABLE_ATTS_WOP(&a_editable_attsAPI.A_EDITABLE_ATTS_WOP)

	c.JSON(http.StatusOK, a_editable_attsAPI)
}

// UpdateA_EDITABLE_ATTS
//
// swagger:route PATCH /a_editable_attss/{ID} a_editable_attss updateA_EDITABLE_ATTS
//
// # Update a a_editable_atts
//
// Responses:
// default: genericError
//
//	200: a_editable_attsDBResponse
func (controller *Controller) UpdateA_EDITABLE_ATTS(c *gin.Context) {

	mutexA_EDITABLE_ATTS.Lock()
	defer mutexA_EDITABLE_ATTS.Unlock()

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
	db := backRepo.BackRepoA_EDITABLE_ATTS.GetDB()

	// Validate input
	var input orm.A_EDITABLE_ATTSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_editable_attsDB orm.A_EDITABLE_ATTSDB

	// fetch the a_editable_atts
	_, err := db.First(&a_editable_attsDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_editable_attsDB.CopyBasicFieldsFromA_EDITABLE_ATTS_WOP(&input.A_EDITABLE_ATTS_WOP)
	a_editable_attsDB.A_EDITABLE_ATTSPointersEncoding = input.A_EDITABLE_ATTSPointersEncoding

	db, _ = db.Model(&a_editable_attsDB)
	_, err = db.Updates(&a_editable_attsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_editable_attsNew := new(models.A_EDITABLE_ATTS)
	a_editable_attsDB.CopyBasicFieldsToA_EDITABLE_ATTS(a_editable_attsNew)

	// redeem pointers
	a_editable_attsDB.DecodePointers(backRepo, a_editable_attsNew)

	// get stage instance from DB instance, and call callback function
	a_editable_attsOld := backRepo.BackRepoA_EDITABLE_ATTS.Map_A_EDITABLE_ATTSDBID_A_EDITABLE_ATTSPtr[a_editable_attsDB.ID]
	if a_editable_attsOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_editable_attsOld, a_editable_attsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_editable_attsDB
	c.JSON(http.StatusOK, a_editable_attsDB)
}

// DeleteA_EDITABLE_ATTS
//
// swagger:route DELETE /a_editable_attss/{ID} a_editable_attss deleteA_EDITABLE_ATTS
//
// # Delete a a_editable_atts
//
// default: genericError
//
//	200: a_editable_attsDBResponse
func (controller *Controller) DeleteA_EDITABLE_ATTS(c *gin.Context) {

	mutexA_EDITABLE_ATTS.Lock()
	defer mutexA_EDITABLE_ATTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_EDITABLE_ATTS", "Name", stackPath)
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
	db := backRepo.BackRepoA_EDITABLE_ATTS.GetDB()

	// Get model if exist
	var a_editable_attsDB orm.A_EDITABLE_ATTSDB
	if _, err := db.First(&a_editable_attsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_editable_attsDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_editable_attsDeleted := new(models.A_EDITABLE_ATTS)
	a_editable_attsDB.CopyBasicFieldsToA_EDITABLE_ATTS(a_editable_attsDeleted)

	// get stage instance from DB instance, and call callback function
	a_editable_attsStaged := backRepo.BackRepoA_EDITABLE_ATTS.Map_A_EDITABLE_ATTSDBID_A_EDITABLE_ATTSPtr[a_editable_attsDB.ID]
	if a_editable_attsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_editable_attsStaged, a_editable_attsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
