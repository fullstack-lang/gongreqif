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
var __A_SPEC_OBJECTS__dummysDeclaration__ models.A_SPEC_OBJECTS
var __A_SPEC_OBJECTS_time__dummyDeclaration time.Duration

var mutexA_SPEC_OBJECTS sync.Mutex

// An A_SPEC_OBJECTSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SPEC_OBJECTS updateA_SPEC_OBJECTS deleteA_SPEC_OBJECTS
type A_SPEC_OBJECTSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SPEC_OBJECTSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SPEC_OBJECTS updateA_SPEC_OBJECTS
type A_SPEC_OBJECTSInput struct {
	// The A_SPEC_OBJECTS to submit or modify
	// in: body
	A_SPEC_OBJECTS *orm.A_SPEC_OBJECTSAPI
}

// GetA_SPEC_OBJECTSs
//
// swagger:route GET /a_spec_objectss a_spec_objectss getA_SPEC_OBJECTSs
//
// # Get all a_spec_objectss
//
// Responses:
// default: genericError
//
//	200: a_spec_objectsDBResponse
func (controller *Controller) GetA_SPEC_OBJECTSs(c *gin.Context) {

	// source slice
	var a_spec_objectsDBs []orm.A_SPEC_OBJECTSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_OBJECTSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_OBJECTS.GetDB()

	_, err := db.Find(&a_spec_objectsDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_spec_objectsAPIs := make([]orm.A_SPEC_OBJECTSAPI, 0)

	// for each a_spec_objects, update fields from the database nullable fields
	for idx := range a_spec_objectsDBs {
		a_spec_objectsDB := &a_spec_objectsDBs[idx]
		_ = a_spec_objectsDB
		var a_spec_objectsAPI orm.A_SPEC_OBJECTSAPI

		// insertion point for updating fields
		a_spec_objectsAPI.ID = a_spec_objectsDB.ID
		a_spec_objectsDB.CopyBasicFieldsToA_SPEC_OBJECTS_WOP(&a_spec_objectsAPI.A_SPEC_OBJECTS_WOP)
		a_spec_objectsAPI.A_SPEC_OBJECTSPointersEncoding = a_spec_objectsDB.A_SPEC_OBJECTSPointersEncoding
		a_spec_objectsAPIs = append(a_spec_objectsAPIs, a_spec_objectsAPI)
	}

	c.JSON(http.StatusOK, a_spec_objectsAPIs)
}

// PostA_SPEC_OBJECTS
//
// swagger:route POST /a_spec_objectss a_spec_objectss postA_SPEC_OBJECTS
//
// Creates a a_spec_objects
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SPEC_OBJECTS(c *gin.Context) {

	mutexA_SPEC_OBJECTS.Lock()
	defer mutexA_SPEC_OBJECTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SPEC_OBJECTSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_OBJECTS.GetDB()

	// Validate input
	var input orm.A_SPEC_OBJECTSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_spec_objects
	a_spec_objectsDB := orm.A_SPEC_OBJECTSDB{}
	a_spec_objectsDB.A_SPEC_OBJECTSPointersEncoding = input.A_SPEC_OBJECTSPointersEncoding
	a_spec_objectsDB.CopyBasicFieldsFromA_SPEC_OBJECTS_WOP(&input.A_SPEC_OBJECTS_WOP)

	_, err = db.Create(&a_spec_objectsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SPEC_OBJECTS.CheckoutPhaseOneInstance(&a_spec_objectsDB)
	a_spec_objects := backRepo.BackRepoA_SPEC_OBJECTS.Map_A_SPEC_OBJECTSDBID_A_SPEC_OBJECTSPtr[a_spec_objectsDB.ID]

	if a_spec_objects != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_spec_objects)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_spec_objectsDB)
}

// GetA_SPEC_OBJECTS
//
// swagger:route GET /a_spec_objectss/{ID} a_spec_objectss getA_SPEC_OBJECTS
//
// Gets the details for a a_spec_objects.
//
// Responses:
// default: genericError
//
//	200: a_spec_objectsDBResponse
func (controller *Controller) GetA_SPEC_OBJECTS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPEC_OBJECTS", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_OBJECTS.GetDB()

	// Get a_spec_objectsDB in DB
	var a_spec_objectsDB orm.A_SPEC_OBJECTSDB
	if _, err := db.First(&a_spec_objectsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_spec_objectsAPI orm.A_SPEC_OBJECTSAPI
	a_spec_objectsAPI.ID = a_spec_objectsDB.ID
	a_spec_objectsAPI.A_SPEC_OBJECTSPointersEncoding = a_spec_objectsDB.A_SPEC_OBJECTSPointersEncoding
	a_spec_objectsDB.CopyBasicFieldsToA_SPEC_OBJECTS_WOP(&a_spec_objectsAPI.A_SPEC_OBJECTS_WOP)

	c.JSON(http.StatusOK, a_spec_objectsAPI)
}

// UpdateA_SPEC_OBJECTS
//
// swagger:route PATCH /a_spec_objectss/{ID} a_spec_objectss updateA_SPEC_OBJECTS
//
// # Update a a_spec_objects
//
// Responses:
// default: genericError
//
//	200: a_spec_objectsDBResponse
func (controller *Controller) UpdateA_SPEC_OBJECTS(c *gin.Context) {

	mutexA_SPEC_OBJECTS.Lock()
	defer mutexA_SPEC_OBJECTS.Unlock()

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
	db := backRepo.BackRepoA_SPEC_OBJECTS.GetDB()

	// Validate input
	var input orm.A_SPEC_OBJECTSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_spec_objectsDB orm.A_SPEC_OBJECTSDB

	// fetch the a_spec_objects
	_, err := db.First(&a_spec_objectsDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_spec_objectsDB.CopyBasicFieldsFromA_SPEC_OBJECTS_WOP(&input.A_SPEC_OBJECTS_WOP)
	a_spec_objectsDB.A_SPEC_OBJECTSPointersEncoding = input.A_SPEC_OBJECTSPointersEncoding

	db, _ = db.Model(&a_spec_objectsDB)
	_, err = db.Updates(&a_spec_objectsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_objectsNew := new(models.A_SPEC_OBJECTS)
	a_spec_objectsDB.CopyBasicFieldsToA_SPEC_OBJECTS(a_spec_objectsNew)

	// redeem pointers
	a_spec_objectsDB.DecodePointers(backRepo, a_spec_objectsNew)

	// get stage instance from DB instance, and call callback function
	a_spec_objectsOld := backRepo.BackRepoA_SPEC_OBJECTS.Map_A_SPEC_OBJECTSDBID_A_SPEC_OBJECTSPtr[a_spec_objectsDB.ID]
	if a_spec_objectsOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_spec_objectsOld, a_spec_objectsNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_spec_objectsDB
	c.JSON(http.StatusOK, a_spec_objectsDB)
}

// DeleteA_SPEC_OBJECTS
//
// swagger:route DELETE /a_spec_objectss/{ID} a_spec_objectss deleteA_SPEC_OBJECTS
//
// # Delete a a_spec_objects
//
// default: genericError
//
//	200: a_spec_objectsDBResponse
func (controller *Controller) DeleteA_SPEC_OBJECTS(c *gin.Context) {

	mutexA_SPEC_OBJECTS.Lock()
	defer mutexA_SPEC_OBJECTS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SPEC_OBJECTS", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPEC_OBJECTS.GetDB()

	// Get model if exist
	var a_spec_objectsDB orm.A_SPEC_OBJECTSDB
	if _, err := db.First(&a_spec_objectsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_spec_objectsDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_spec_objectsDeleted := new(models.A_SPEC_OBJECTS)
	a_spec_objectsDB.CopyBasicFieldsToA_SPEC_OBJECTS(a_spec_objectsDeleted)

	// get stage instance from DB instance, and call callback function
	a_spec_objectsStaged := backRepo.BackRepoA_SPEC_OBJECTS.Map_A_SPEC_OBJECTSDBID_A_SPEC_OBJECTSPtr[a_spec_objectsDB.ID]
	if a_spec_objectsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_spec_objectsStaged, a_spec_objectsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
