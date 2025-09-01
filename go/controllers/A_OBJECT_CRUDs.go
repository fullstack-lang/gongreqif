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
var __A_OBJECT__dummysDeclaration__ models.A_OBJECT
var __A_OBJECT_time__dummyDeclaration time.Duration

var mutexA_OBJECT sync.Mutex

// An A_OBJECTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_OBJECT updateA_OBJECT deleteA_OBJECT
type A_OBJECTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_OBJECTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_OBJECT updateA_OBJECT
type A_OBJECTInput struct {
	// The A_OBJECT to submit or modify
	// in: body
	A_OBJECT *orm.A_OBJECTAPI
}

// GetA_OBJECTs
//
// swagger:route GET /a_objects a_objects getA_OBJECTs
//
// # Get all a_objects
//
// Responses:
// default: genericError
//
//	200: a_objectDBResponse
func (controller *Controller) GetA_OBJECTs(c *gin.Context) {

	// source slice
	var a_objectDBs []orm.A_OBJECTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_OBJECTs", "Name", stackPath)
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
	db := backRepo.BackRepoA_OBJECT.GetDB()

	_, err := db.Find(&a_objectDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_objectAPIs := make([]orm.A_OBJECTAPI, 0)

	// for each a_object, update fields from the database nullable fields
	for idx := range a_objectDBs {
		a_objectDB := &a_objectDBs[idx]
		_ = a_objectDB
		var a_objectAPI orm.A_OBJECTAPI

		// insertion point for updating fields
		a_objectAPI.ID = a_objectDB.ID
		a_objectDB.CopyBasicFieldsToA_OBJECT_WOP(&a_objectAPI.A_OBJECT_WOP)
		a_objectAPI.A_OBJECTPointersEncoding = a_objectDB.A_OBJECTPointersEncoding
		a_objectAPIs = append(a_objectAPIs, a_objectAPI)
	}

	c.JSON(http.StatusOK, a_objectAPIs)
}

// PostA_OBJECT
//
// swagger:route POST /a_objects a_objects postA_OBJECT
//
// Creates a a_object
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_OBJECT(c *gin.Context) {

	mutexA_OBJECT.Lock()
	defer mutexA_OBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_OBJECTs", "Name", stackPath)
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
	db := backRepo.BackRepoA_OBJECT.GetDB()

	// Validate input
	var input orm.A_OBJECTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_object
	a_objectDB := orm.A_OBJECTDB{}
	a_objectDB.A_OBJECTPointersEncoding = input.A_OBJECTPointersEncoding
	a_objectDB.CopyBasicFieldsFromA_OBJECT_WOP(&input.A_OBJECT_WOP)

	_, err = db.Create(&a_objectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_OBJECT.CheckoutPhaseOneInstance(&a_objectDB)
	a_object := backRepo.BackRepoA_OBJECT.Map_A_OBJECTDBID_A_OBJECTPtr[a_objectDB.ID]

	if a_object != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_object)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_objectDB)
}

// GetA_OBJECT
//
// swagger:route GET /a_objects/{ID} a_objects getA_OBJECT
//
// Gets the details for a a_object.
//
// Responses:
// default: genericError
//
//	200: a_objectDBResponse
func (controller *Controller) GetA_OBJECT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_OBJECT", "Name", stackPath)
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
	db := backRepo.BackRepoA_OBJECT.GetDB()

	// Get a_objectDB in DB
	var a_objectDB orm.A_OBJECTDB
	if _, err := db.First(&a_objectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_objectAPI orm.A_OBJECTAPI
	a_objectAPI.ID = a_objectDB.ID
	a_objectAPI.A_OBJECTPointersEncoding = a_objectDB.A_OBJECTPointersEncoding
	a_objectDB.CopyBasicFieldsToA_OBJECT_WOP(&a_objectAPI.A_OBJECT_WOP)

	c.JSON(http.StatusOK, a_objectAPI)
}

// UpdateA_OBJECT
//
// swagger:route PATCH /a_objects/{ID} a_objects updateA_OBJECT
//
// # Update a a_object
//
// Responses:
// default: genericError
//
//	200: a_objectDBResponse
func (controller *Controller) UpdateA_OBJECT(c *gin.Context) {

	mutexA_OBJECT.Lock()
	defer mutexA_OBJECT.Unlock()

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
	db := backRepo.BackRepoA_OBJECT.GetDB()

	// Validate input
	var input orm.A_OBJECTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_objectDB orm.A_OBJECTDB

	// fetch the a_object
	_, err := db.First(&a_objectDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_objectDB.CopyBasicFieldsFromA_OBJECT_WOP(&input.A_OBJECT_WOP)
	a_objectDB.A_OBJECTPointersEncoding = input.A_OBJECTPointersEncoding

	db, _ = db.Model(&a_objectDB)
	_, err = db.Updates(&a_objectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_objectNew := new(models.A_OBJECT)
	a_objectDB.CopyBasicFieldsToA_OBJECT(a_objectNew)

	// redeem pointers
	a_objectDB.DecodePointers(backRepo, a_objectNew)

	// get stage instance from DB instance, and call callback function
	a_objectOld := backRepo.BackRepoA_OBJECT.Map_A_OBJECTDBID_A_OBJECTPtr[a_objectDB.ID]
	if a_objectOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_objectOld, a_objectNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_objectOld, a_objectNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_objectDB
	c.JSON(http.StatusOK, a_objectDB)
}

// DeleteA_OBJECT
//
// swagger:route DELETE /a_objects/{ID} a_objects deleteA_OBJECT
//
// # Delete a a_object
//
// default: genericError
//
//	200: a_objectDBResponse
func (controller *Controller) DeleteA_OBJECT(c *gin.Context) {

	mutexA_OBJECT.Lock()
	defer mutexA_OBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_OBJECT", "Name", stackPath)
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
	db := backRepo.BackRepoA_OBJECT.GetDB()

	// Get model if exist
	var a_objectDB orm.A_OBJECTDB
	if _, err := db.First(&a_objectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_objectDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_objectDeleted := new(models.A_OBJECT)
	a_objectDB.CopyBasicFieldsToA_OBJECT(a_objectDeleted)

	// get stage instance from DB instance, and call callback function
	a_objectStaged := backRepo.BackRepoA_OBJECT.Map_A_OBJECTDBID_A_OBJECTPtr[a_objectDB.ID]
	if a_objectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_objectStaged, a_objectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
