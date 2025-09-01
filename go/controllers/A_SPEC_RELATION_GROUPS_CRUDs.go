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
var __A_SPEC_RELATION_GROUPS__dummysDeclaration__ models.A_SPEC_RELATION_GROUPS
var __A_SPEC_RELATION_GROUPS_time__dummyDeclaration time.Duration

var mutexA_SPEC_RELATION_GROUPS sync.Mutex

// An A_SPEC_RELATION_GROUPSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SPEC_RELATION_GROUPS updateA_SPEC_RELATION_GROUPS deleteA_SPEC_RELATION_GROUPS
type A_SPEC_RELATION_GROUPSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SPEC_RELATION_GROUPSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SPEC_RELATION_GROUPS updateA_SPEC_RELATION_GROUPS
type A_SPEC_RELATION_GROUPSInput struct {
	// The A_SPEC_RELATION_GROUPS to submit or modify
	// in: body
	A_SPEC_RELATION_GROUPS *orm.A_SPEC_RELATION_GROUPSAPI
}

// GetA_SPEC_RELATION_GROUPSs
//
// swagger:route GET /a_spec_relation_groupss a_spec_relation_groupss getA_SPEC_RELATION_GROUPSs
//
// # Get all a_spec_relation_groupss
//
// Responses:
// default: genericError
//
//	200: a_spec_relation_groupsDBResponse
func (controller *Controller) GetA_SPEC_RELATION_GROUPSs(c *gin.Context) {

	// source slice
	var a_spec_relation_groupsDBs []orm.A_SPEC_RELATION_GROUPSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_RELATION_GROUPSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_RELATION_GROUPS.GetDB()

	_, err := db.Find(&a_spec_relation_groupsDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_spec_relation_groupsAPIs := make([]orm.A_SPEC_RELATION_GROUPSAPI, 0)

	// for each a_spec_relation_groups, update fields from the database nullable fields
	for idx := range a_spec_relation_groupsDBs {
		a_spec_relation_groupsDB := &a_spec_relation_groupsDBs[idx]
		_ = a_spec_relation_groupsDB
		var a_spec_relation_groupsAPI orm.A_SPEC_RELATION_GROUPSAPI

		// insertion point for updating fields
		a_spec_relation_groupsAPI.ID = a_spec_relation_groupsDB.ID
		a_spec_relation_groupsDB.CopyBasicFieldsToA_SPEC_RELATION_GROUPS_WOP(&a_spec_relation_groupsAPI.A_SPEC_RELATION_GROUPS_WOP)
		a_spec_relation_groupsAPI.A_SPEC_RELATION_GROUPSPointersEncoding = a_spec_relation_groupsDB.A_SPEC_RELATION_GROUPSPointersEncoding
		a_spec_relation_groupsAPIs = append(a_spec_relation_groupsAPIs, a_spec_relation_groupsAPI)
	}

	c.JSON(http.StatusOK, a_spec_relation_groupsAPIs)
}

// PostA_SPEC_RELATION_GROUPS
//
// swagger:route POST /a_spec_relation_groupss a_spec_relation_groupss postA_SPEC_RELATION_GROUPS
//
// Creates a a_spec_relation_groups
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SPEC_RELATION_GROUPS(c *gin.Context) {

	mutexA_SPEC_RELATION_GROUPS.Lock()
	defer mutexA_SPEC_RELATION_GROUPS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SPEC_RELATION_GROUPSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_RELATION_GROUPS.GetDB()

	// Validate input
	var input orm.A_SPEC_RELATION_GROUPSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_spec_relation_groups
	a_spec_relation_groupsDB := orm.A_SPEC_RELATION_GROUPSDB{}
	a_spec_relation_groupsDB.A_SPEC_RELATION_GROUPSPointersEncoding = input.A_SPEC_RELATION_GROUPSPointersEncoding
	a_spec_relation_groupsDB.CopyBasicFieldsFromA_SPEC_RELATION_GROUPS_WOP(&input.A_SPEC_RELATION_GROUPS_WOP)

	_, err = db.Create(&a_spec_relation_groupsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SPEC_RELATION_GROUPS.CheckoutPhaseOneInstance(&a_spec_relation_groupsDB)
	a_spec_relation_groups := backRepo.BackRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr[a_spec_relation_groupsDB.ID]

	if a_spec_relation_groups != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_spec_relation_groups)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_spec_relation_groupsDB)
}

// GetA_SPEC_RELATION_GROUPS
//
// swagger:route GET /a_spec_relation_groupss/{ID} a_spec_relation_groupss getA_SPEC_RELATION_GROUPS
//
// Gets the details for a a_spec_relation_groups.
//
// Responses:
// default: genericError
//
//	200: a_spec_relation_groupsDBResponse
func (controller *Controller) GetA_SPEC_RELATION_GROUPS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_RELATION_GROUPS", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_RELATION_GROUPS.GetDB()

	// Get a_spec_relation_groupsDB in DB
	var a_spec_relation_groupsDB orm.A_SPEC_RELATION_GROUPSDB
	if _, err := db.First(&a_spec_relation_groupsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_spec_relation_groupsAPI orm.A_SPEC_RELATION_GROUPSAPI
	a_spec_relation_groupsAPI.ID = a_spec_relation_groupsDB.ID
	a_spec_relation_groupsAPI.A_SPEC_RELATION_GROUPSPointersEncoding = a_spec_relation_groupsDB.A_SPEC_RELATION_GROUPSPointersEncoding
	a_spec_relation_groupsDB.CopyBasicFieldsToA_SPEC_RELATION_GROUPS_WOP(&a_spec_relation_groupsAPI.A_SPEC_RELATION_GROUPS_WOP)

	c.JSON(http.StatusOK, a_spec_relation_groupsAPI)
}

// UpdateA_SPEC_RELATION_GROUPS
//
// swagger:route PATCH /a_spec_relation_groupss/{ID} a_spec_relation_groupss updateA_SPEC_RELATION_GROUPS
//
// # Update a a_spec_relation_groups
//
// Responses:
// default: genericError
//
//	200: a_spec_relation_groupsDBResponse
func (controller *Controller) UpdateA_SPEC_RELATION_GROUPS(c *gin.Context) {

	mutexA_SPEC_RELATION_GROUPS.Lock()
	defer mutexA_SPEC_RELATION_GROUPS.Unlock()

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
	db := backRepo.BackRepoA_SPEC_RELATION_GROUPS.GetDB()

	// Validate input
	var input orm.A_SPEC_RELATION_GROUPSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_spec_relation_groupsDB orm.A_SPEC_RELATION_GROUPSDB

	// fetch the a_spec_relation_groups
	_, err := db.First(&a_spec_relation_groupsDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_spec_relation_groupsDB.CopyBasicFieldsFromA_SPEC_RELATION_GROUPS_WOP(&input.A_SPEC_RELATION_GROUPS_WOP)
	a_spec_relation_groupsDB.A_SPEC_RELATION_GROUPSPointersEncoding = input.A_SPEC_RELATION_GROUPSPointersEncoding

	db, _ = db.Model(&a_spec_relation_groupsDB)
	_, err = db.Updates(&a_spec_relation_groupsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_relation_groupsNew := new(models.A_SPEC_RELATION_GROUPS)
	a_spec_relation_groupsDB.CopyBasicFieldsToA_SPEC_RELATION_GROUPS(a_spec_relation_groupsNew)

	// redeem pointers
	a_spec_relation_groupsDB.DecodePointers(backRepo, a_spec_relation_groupsNew)

	// get stage instance from DB instance, and call callback function
	a_spec_relation_groupsOld := backRepo.BackRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr[a_spec_relation_groupsDB.ID]
	if a_spec_relation_groupsOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_spec_relation_groupsOld, a_spec_relation_groupsNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_spec_relation_groupsOld, a_spec_relation_groupsNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_spec_relation_groupsDB
	c.JSON(http.StatusOK, a_spec_relation_groupsDB)
}

// DeleteA_SPEC_RELATION_GROUPS
//
// swagger:route DELETE /a_spec_relation_groupss/{ID} a_spec_relation_groupss deleteA_SPEC_RELATION_GROUPS
//
// # Delete a a_spec_relation_groups
//
// default: genericError
//
//	200: a_spec_relation_groupsDBResponse
func (controller *Controller) DeleteA_SPEC_RELATION_GROUPS(c *gin.Context) {

	mutexA_SPEC_RELATION_GROUPS.Lock()
	defer mutexA_SPEC_RELATION_GROUPS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SPEC_RELATION_GROUPS", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_RELATION_GROUPS.GetDB()

	// Get model if exist
	var a_spec_relation_groupsDB orm.A_SPEC_RELATION_GROUPSDB
	if _, err := db.First(&a_spec_relation_groupsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_spec_relation_groupsDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_relation_groupsDeleted := new(models.A_SPEC_RELATION_GROUPS)
	a_spec_relation_groupsDB.CopyBasicFieldsToA_SPEC_RELATION_GROUPS(a_spec_relation_groupsDeleted)

	// get stage instance from DB instance, and call callback function
	a_spec_relation_groupsStaged := backRepo.BackRepoA_SPEC_RELATION_GROUPS.Map_A_SPEC_RELATION_GROUPSDBID_A_SPEC_RELATION_GROUPSPtr[a_spec_relation_groupsDB.ID]
	if a_spec_relation_groupsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_spec_relation_groupsStaged, a_spec_relation_groupsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
