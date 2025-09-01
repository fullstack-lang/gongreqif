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
var __A_SPEC_RELATIONS__dummysDeclaration__ models.A_SPEC_RELATIONS
var __A_SPEC_RELATIONS_time__dummyDeclaration time.Duration

var mutexA_SPEC_RELATIONS sync.Mutex

// An A_SPEC_RELATIONSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SPEC_RELATIONS updateA_SPEC_RELATIONS deleteA_SPEC_RELATIONS
type A_SPEC_RELATIONSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SPEC_RELATIONSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SPEC_RELATIONS updateA_SPEC_RELATIONS
type A_SPEC_RELATIONSInput struct {
	// The A_SPEC_RELATIONS to submit or modify
	// in: body
	A_SPEC_RELATIONS *orm.A_SPEC_RELATIONSAPI
}

// GetA_SPEC_RELATIONSs
//
// swagger:route GET /a_spec_relationss a_spec_relationss getA_SPEC_RELATIONSs
//
// # Get all a_spec_relationss
//
// Responses:
// default: genericError
//
//	200: a_spec_relationsDBResponse
func (controller *Controller) GetA_SPEC_RELATIONSs(c *gin.Context) {

	// source slice
	var a_spec_relationsDBs []orm.A_SPEC_RELATIONSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_RELATIONSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_RELATIONS.GetDB()

	_, err := db.Find(&a_spec_relationsDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_spec_relationsAPIs := make([]orm.A_SPEC_RELATIONSAPI, 0)

	// for each a_spec_relations, update fields from the database nullable fields
	for idx := range a_spec_relationsDBs {
		a_spec_relationsDB := &a_spec_relationsDBs[idx]
		_ = a_spec_relationsDB
		var a_spec_relationsAPI orm.A_SPEC_RELATIONSAPI

		// insertion point for updating fields
		a_spec_relationsAPI.ID = a_spec_relationsDB.ID
		a_spec_relationsDB.CopyBasicFieldsToA_SPEC_RELATIONS_WOP(&a_spec_relationsAPI.A_SPEC_RELATIONS_WOP)
		a_spec_relationsAPI.A_SPEC_RELATIONSPointersEncoding = a_spec_relationsDB.A_SPEC_RELATIONSPointersEncoding
		a_spec_relationsAPIs = append(a_spec_relationsAPIs, a_spec_relationsAPI)
	}

	c.JSON(http.StatusOK, a_spec_relationsAPIs)
}

// PostA_SPEC_RELATIONS
//
// swagger:route POST /a_spec_relationss a_spec_relationss postA_SPEC_RELATIONS
//
// Creates a a_spec_relations
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SPEC_RELATIONS(c *gin.Context) {

	mutexA_SPEC_RELATIONS.Lock()
	defer mutexA_SPEC_RELATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SPEC_RELATIONSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_RELATIONS.GetDB()

	// Validate input
	var input orm.A_SPEC_RELATIONSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_spec_relations
	a_spec_relationsDB := orm.A_SPEC_RELATIONSDB{}
	a_spec_relationsDB.A_SPEC_RELATIONSPointersEncoding = input.A_SPEC_RELATIONSPointersEncoding
	a_spec_relationsDB.CopyBasicFieldsFromA_SPEC_RELATIONS_WOP(&input.A_SPEC_RELATIONS_WOP)

	_, err = db.Create(&a_spec_relationsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SPEC_RELATIONS.CheckoutPhaseOneInstance(&a_spec_relationsDB)
	a_spec_relations := backRepo.BackRepoA_SPEC_RELATIONS.Map_A_SPEC_RELATIONSDBID_A_SPEC_RELATIONSPtr[a_spec_relationsDB.ID]

	if a_spec_relations != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_spec_relations)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_spec_relationsDB)
}

// GetA_SPEC_RELATIONS
//
// swagger:route GET /a_spec_relationss/{ID} a_spec_relationss getA_SPEC_RELATIONS
//
// Gets the details for a a_spec_relations.
//
// Responses:
// default: genericError
//
//	200: a_spec_relationsDBResponse
func (controller *Controller) GetA_SPEC_RELATIONS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_RELATIONS", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_RELATIONS.GetDB()

	// Get a_spec_relationsDB in DB
	var a_spec_relationsDB orm.A_SPEC_RELATIONSDB
	if _, err := db.First(&a_spec_relationsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_spec_relationsAPI orm.A_SPEC_RELATIONSAPI
	a_spec_relationsAPI.ID = a_spec_relationsDB.ID
	a_spec_relationsAPI.A_SPEC_RELATIONSPointersEncoding = a_spec_relationsDB.A_SPEC_RELATIONSPointersEncoding
	a_spec_relationsDB.CopyBasicFieldsToA_SPEC_RELATIONS_WOP(&a_spec_relationsAPI.A_SPEC_RELATIONS_WOP)

	c.JSON(http.StatusOK, a_spec_relationsAPI)
}

// UpdateA_SPEC_RELATIONS
//
// swagger:route PATCH /a_spec_relationss/{ID} a_spec_relationss updateA_SPEC_RELATIONS
//
// # Update a a_spec_relations
//
// Responses:
// default: genericError
//
//	200: a_spec_relationsDBResponse
func (controller *Controller) UpdateA_SPEC_RELATIONS(c *gin.Context) {

	mutexA_SPEC_RELATIONS.Lock()
	defer mutexA_SPEC_RELATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	hasMouseEvent := false
	shiftKey := false
	_ = shiftKey
	if len(_values) >= 1 {
		_nameValues := _values["Name"]
		if len(_nameValues) == 1 {
			stackPath = _nameValues[0]
		}
	}

	if len(_values) >= 2 {
		hasMouseEvent = true
		_shiftKeyValues := _values["shiftKey"]
		if len(_shiftKeyValues) == 1 {
			shiftKey = _shiftKeyValues[0] == "true"
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
	db := backRepo.BackRepoA_SPEC_RELATIONS.GetDB()

	// Validate input
	var input orm.A_SPEC_RELATIONSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_spec_relationsDB orm.A_SPEC_RELATIONSDB

	// fetch the a_spec_relations
	_, err := db.First(&a_spec_relationsDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_spec_relationsDB.CopyBasicFieldsFromA_SPEC_RELATIONS_WOP(&input.A_SPEC_RELATIONS_WOP)
	a_spec_relationsDB.A_SPEC_RELATIONSPointersEncoding = input.A_SPEC_RELATIONSPointersEncoding

	db, _ = db.Model(&a_spec_relationsDB)
	_, err = db.Updates(&a_spec_relationsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_relationsNew := new(models.A_SPEC_RELATIONS)
	a_spec_relationsDB.CopyBasicFieldsToA_SPEC_RELATIONS(a_spec_relationsNew)

	// redeem pointers
	a_spec_relationsDB.DecodePointers(backRepo, a_spec_relationsNew)

	// get stage instance from DB instance, and call callback function
	a_spec_relationsOld := backRepo.BackRepoA_SPEC_RELATIONS.Map_A_SPEC_RELATIONSDBID_A_SPEC_RELATIONSPtr[a_spec_relationsDB.ID]
	if a_spec_relationsOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_spec_relationsOld, a_spec_relationsNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_spec_relationsOld, a_spec_relationsNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_spec_relationsDB
	c.JSON(http.StatusOK, a_spec_relationsDB)
}

// DeleteA_SPEC_RELATIONS
//
// swagger:route DELETE /a_spec_relationss/{ID} a_spec_relationss deleteA_SPEC_RELATIONS
//
// # Delete a a_spec_relations
//
// default: genericError
//
//	200: a_spec_relationsDBResponse
func (controller *Controller) DeleteA_SPEC_RELATIONS(c *gin.Context) {

	mutexA_SPEC_RELATIONS.Lock()
	defer mutexA_SPEC_RELATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SPEC_RELATIONS", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_RELATIONS.GetDB()

	// Get model if exist
	var a_spec_relationsDB orm.A_SPEC_RELATIONSDB
	if _, err := db.First(&a_spec_relationsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_spec_relationsDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_relationsDeleted := new(models.A_SPEC_RELATIONS)
	a_spec_relationsDB.CopyBasicFieldsToA_SPEC_RELATIONS(a_spec_relationsDeleted)

	// get stage instance from DB instance, and call callback function
	a_spec_relationsStaged := backRepo.BackRepoA_SPEC_RELATIONS.Map_A_SPEC_RELATIONSDBID_A_SPEC_RELATIONSPtr[a_spec_relationsDB.ID]
	if a_spec_relationsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_spec_relationsStaged, a_spec_relationsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
