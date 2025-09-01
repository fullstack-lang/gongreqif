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
var __SPEC_OBJECT__dummysDeclaration__ models.SPEC_OBJECT
var __SPEC_OBJECT_time__dummyDeclaration time.Duration

var mutexSPEC_OBJECT sync.Mutex

// An SPEC_OBJECTID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPEC_OBJECT updateSPEC_OBJECT deleteSPEC_OBJECT
type SPEC_OBJECTID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPEC_OBJECTInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPEC_OBJECT updateSPEC_OBJECT
type SPEC_OBJECTInput struct {
	// The SPEC_OBJECT to submit or modify
	// in: body
	SPEC_OBJECT *orm.SPEC_OBJECTAPI
}

// GetSPEC_OBJECTs
//
// swagger:route GET /spec_objects spec_objects getSPEC_OBJECTs
//
// # Get all spec_objects
//
// Responses:
// default: genericError
//
//	200: spec_objectDBResponse
func (controller *Controller) GetSPEC_OBJECTs(c *gin.Context) {

	// source slice
	var spec_objectDBs []orm.SPEC_OBJECTDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_OBJECTs", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_OBJECT.GetDB()

	_, err := db.Find(&spec_objectDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spec_objectAPIs := make([]orm.SPEC_OBJECTAPI, 0)

	// for each spec_object, update fields from the database nullable fields
	for idx := range spec_objectDBs {
		spec_objectDB := &spec_objectDBs[idx]
		_ = spec_objectDB
		var spec_objectAPI orm.SPEC_OBJECTAPI

		// insertion point for updating fields
		spec_objectAPI.ID = spec_objectDB.ID
		spec_objectDB.CopyBasicFieldsToSPEC_OBJECT_WOP(&spec_objectAPI.SPEC_OBJECT_WOP)
		spec_objectAPI.SPEC_OBJECTPointersEncoding = spec_objectDB.SPEC_OBJECTPointersEncoding
		spec_objectAPIs = append(spec_objectAPIs, spec_objectAPI)
	}

	c.JSON(http.StatusOK, spec_objectAPIs)
}

// PostSPEC_OBJECT
//
// swagger:route POST /spec_objects spec_objects postSPEC_OBJECT
//
// Creates a spec_object
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPEC_OBJECT(c *gin.Context) {

	mutexSPEC_OBJECT.Lock()
	defer mutexSPEC_OBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPEC_OBJECTs", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_OBJECT.GetDB()

	// Validate input
	var input orm.SPEC_OBJECTAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spec_object
	spec_objectDB := orm.SPEC_OBJECTDB{}
	spec_objectDB.SPEC_OBJECTPointersEncoding = input.SPEC_OBJECTPointersEncoding
	spec_objectDB.CopyBasicFieldsFromSPEC_OBJECT_WOP(&input.SPEC_OBJECT_WOP)

	_, err = db.Create(&spec_objectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPEC_OBJECT.CheckoutPhaseOneInstance(&spec_objectDB)
	spec_object := backRepo.BackRepoSPEC_OBJECT.Map_SPEC_OBJECTDBID_SPEC_OBJECTPtr[spec_objectDB.ID]

	if spec_object != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spec_object)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spec_objectDB)
}

// GetSPEC_OBJECT
//
// swagger:route GET /spec_objects/{ID} spec_objects getSPEC_OBJECT
//
// Gets the details for a spec_object.
//
// Responses:
// default: genericError
//
//	200: spec_objectDBResponse
func (controller *Controller) GetSPEC_OBJECT(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_OBJECT", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_OBJECT.GetDB()

	// Get spec_objectDB in DB
	var spec_objectDB orm.SPEC_OBJECTDB
	if _, err := db.First(&spec_objectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spec_objectAPI orm.SPEC_OBJECTAPI
	spec_objectAPI.ID = spec_objectDB.ID
	spec_objectAPI.SPEC_OBJECTPointersEncoding = spec_objectDB.SPEC_OBJECTPointersEncoding
	spec_objectDB.CopyBasicFieldsToSPEC_OBJECT_WOP(&spec_objectAPI.SPEC_OBJECT_WOP)

	c.JSON(http.StatusOK, spec_objectAPI)
}

// UpdateSPEC_OBJECT
//
// swagger:route PATCH /spec_objects/{ID} spec_objects updateSPEC_OBJECT
//
// # Update a spec_object
//
// Responses:
// default: genericError
//
//	200: spec_objectDBResponse
func (controller *Controller) UpdateSPEC_OBJECT(c *gin.Context) {

	mutexSPEC_OBJECT.Lock()
	defer mutexSPEC_OBJECT.Unlock()

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
	db := backRepo.BackRepoSPEC_OBJECT.GetDB()

	// Validate input
	var input orm.SPEC_OBJECTAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spec_objectDB orm.SPEC_OBJECTDB

	// fetch the spec_object
	_, err := db.First(&spec_objectDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spec_objectDB.CopyBasicFieldsFromSPEC_OBJECT_WOP(&input.SPEC_OBJECT_WOP)
	spec_objectDB.SPEC_OBJECTPointersEncoding = input.SPEC_OBJECTPointersEncoding

	db, _ = db.Model(&spec_objectDB)
	_, err = db.Updates(&spec_objectDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spec_objectNew := new(models.SPEC_OBJECT)
	spec_objectDB.CopyBasicFieldsToSPEC_OBJECT(spec_objectNew)

	// redeem pointers
	spec_objectDB.DecodePointers(backRepo, spec_objectNew)

	// get stage instance from DB instance, and call callback function
	spec_objectOld := backRepo.BackRepoSPEC_OBJECT.Map_SPEC_OBJECTDBID_SPEC_OBJECTPtr[spec_objectDB.ID]
	if spec_objectOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), spec_objectOld, spec_objectNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), spec_objectOld, spec_objectNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spec_objectDB
	c.JSON(http.StatusOK, spec_objectDB)
}

// DeleteSPEC_OBJECT
//
// swagger:route DELETE /spec_objects/{ID} spec_objects deleteSPEC_OBJECT
//
// # Delete a spec_object
//
// default: genericError
//
//	200: spec_objectDBResponse
func (controller *Controller) DeleteSPEC_OBJECT(c *gin.Context) {

	mutexSPEC_OBJECT.Lock()
	defer mutexSPEC_OBJECT.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPEC_OBJECT", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_OBJECT.GetDB()

	// Get model if exist
	var spec_objectDB orm.SPEC_OBJECTDB
	if _, err := db.First(&spec_objectDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spec_objectDB)

	// get an instance (not staged) from DB instance, and call callback function
	spec_objectDeleted := new(models.SPEC_OBJECT)
	spec_objectDB.CopyBasicFieldsToSPEC_OBJECT(spec_objectDeleted)

	// get stage instance from DB instance, and call callback function
	spec_objectStaged := backRepo.BackRepoSPEC_OBJECT.Map_SPEC_OBJECTDBID_SPEC_OBJECTPtr[spec_objectDB.ID]
	if spec_objectStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spec_objectStaged, spec_objectDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
