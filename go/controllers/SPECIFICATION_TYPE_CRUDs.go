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
var __SPECIFICATION_TYPE__dummysDeclaration__ models.SPECIFICATION_TYPE
var __SPECIFICATION_TYPE_time__dummyDeclaration time.Duration

var mutexSPECIFICATION_TYPE sync.Mutex

// An SPECIFICATION_TYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPECIFICATION_TYPE updateSPECIFICATION_TYPE deleteSPECIFICATION_TYPE
type SPECIFICATION_TYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPECIFICATION_TYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPECIFICATION_TYPE updateSPECIFICATION_TYPE
type SPECIFICATION_TYPEInput struct {
	// The SPECIFICATION_TYPE to submit or modify
	// in: body
	SPECIFICATION_TYPE *orm.SPECIFICATION_TYPEAPI
}

// GetSPECIFICATION_TYPEs
//
// swagger:route GET /specification_types specification_types getSPECIFICATION_TYPEs
//
// # Get all specification_types
//
// Responses:
// default: genericError
//
//	200: specification_typeDBResponse
func (controller *Controller) GetSPECIFICATION_TYPEs(c *gin.Context) {

	// source slice
	var specification_typeDBs []orm.SPECIFICATION_TYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFICATION_TYPEs", "Name", stackPath)
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
	db := backRepo.BackRepoSPECIFICATION_TYPE.GetDB()

	_, err := db.Find(&specification_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	specification_typeAPIs := make([]orm.SPECIFICATION_TYPEAPI, 0)

	// for each specification_type, update fields from the database nullable fields
	for idx := range specification_typeDBs {
		specification_typeDB := &specification_typeDBs[idx]
		_ = specification_typeDB
		var specification_typeAPI orm.SPECIFICATION_TYPEAPI

		// insertion point for updating fields
		specification_typeAPI.ID = specification_typeDB.ID
		specification_typeDB.CopyBasicFieldsToSPECIFICATION_TYPE_WOP(&specification_typeAPI.SPECIFICATION_TYPE_WOP)
		specification_typeAPI.SPECIFICATION_TYPEPointersEncoding = specification_typeDB.SPECIFICATION_TYPEPointersEncoding
		specification_typeAPIs = append(specification_typeAPIs, specification_typeAPI)
	}

	c.JSON(http.StatusOK, specification_typeAPIs)
}

// PostSPECIFICATION_TYPE
//
// swagger:route POST /specification_types specification_types postSPECIFICATION_TYPE
//
// Creates a specification_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPECIFICATION_TYPE(c *gin.Context) {

	mutexSPECIFICATION_TYPE.Lock()
	defer mutexSPECIFICATION_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPECIFICATION_TYPEs", "Name", stackPath)
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
	db := backRepo.BackRepoSPECIFICATION_TYPE.GetDB()

	// Validate input
	var input orm.SPECIFICATION_TYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create specification_type
	specification_typeDB := orm.SPECIFICATION_TYPEDB{}
	specification_typeDB.SPECIFICATION_TYPEPointersEncoding = input.SPECIFICATION_TYPEPointersEncoding
	specification_typeDB.CopyBasicFieldsFromSPECIFICATION_TYPE_WOP(&input.SPECIFICATION_TYPE_WOP)

	_, err = db.Create(&specification_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPECIFICATION_TYPE.CheckoutPhaseOneInstance(&specification_typeDB)
	specification_type := backRepo.BackRepoSPECIFICATION_TYPE.Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEPtr[specification_typeDB.ID]

	if specification_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), specification_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, specification_typeDB)
}

// GetSPECIFICATION_TYPE
//
// swagger:route GET /specification_types/{ID} specification_types getSPECIFICATION_TYPE
//
// Gets the details for a specification_type.
//
// Responses:
// default: genericError
//
//	200: specification_typeDBResponse
func (controller *Controller) GetSPECIFICATION_TYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPECIFICATION_TYPE", "Name", stackPath)
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
	db := backRepo.BackRepoSPECIFICATION_TYPE.GetDB()

	// Get specification_typeDB in DB
	var specification_typeDB orm.SPECIFICATION_TYPEDB
	if _, err := db.First(&specification_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var specification_typeAPI orm.SPECIFICATION_TYPEAPI
	specification_typeAPI.ID = specification_typeDB.ID
	specification_typeAPI.SPECIFICATION_TYPEPointersEncoding = specification_typeDB.SPECIFICATION_TYPEPointersEncoding
	specification_typeDB.CopyBasicFieldsToSPECIFICATION_TYPE_WOP(&specification_typeAPI.SPECIFICATION_TYPE_WOP)

	c.JSON(http.StatusOK, specification_typeAPI)
}

// UpdateSPECIFICATION_TYPE
//
// swagger:route PATCH /specification_types/{ID} specification_types updateSPECIFICATION_TYPE
//
// # Update a specification_type
//
// Responses:
// default: genericError
//
//	200: specification_typeDBResponse
func (controller *Controller) UpdateSPECIFICATION_TYPE(c *gin.Context) {

	mutexSPECIFICATION_TYPE.Lock()
	defer mutexSPECIFICATION_TYPE.Unlock()

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
	db := backRepo.BackRepoSPECIFICATION_TYPE.GetDB()

	// Validate input
	var input orm.SPECIFICATION_TYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var specification_typeDB orm.SPECIFICATION_TYPEDB

	// fetch the specification_type
	_, err := db.First(&specification_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	specification_typeDB.CopyBasicFieldsFromSPECIFICATION_TYPE_WOP(&input.SPECIFICATION_TYPE_WOP)
	specification_typeDB.SPECIFICATION_TYPEPointersEncoding = input.SPECIFICATION_TYPEPointersEncoding

	db, _ = db.Model(&specification_typeDB)
	_, err = db.Updates(&specification_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	specification_typeNew := new(models.SPECIFICATION_TYPE)
	specification_typeDB.CopyBasicFieldsToSPECIFICATION_TYPE(specification_typeNew)

	// redeem pointers
	specification_typeDB.DecodePointers(backRepo, specification_typeNew)

	// get stage instance from DB instance, and call callback function
	specification_typeOld := backRepo.BackRepoSPECIFICATION_TYPE.Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEPtr[specification_typeDB.ID]
	if specification_typeOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), specification_typeOld, specification_typeNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), specification_typeOld, specification_typeNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the specification_typeDB
	c.JSON(http.StatusOK, specification_typeDB)
}

// DeleteSPECIFICATION_TYPE
//
// swagger:route DELETE /specification_types/{ID} specification_types deleteSPECIFICATION_TYPE
//
// # Delete a specification_type
//
// default: genericError
//
//	200: specification_typeDBResponse
func (controller *Controller) DeleteSPECIFICATION_TYPE(c *gin.Context) {

	mutexSPECIFICATION_TYPE.Lock()
	defer mutexSPECIFICATION_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPECIFICATION_TYPE", "Name", stackPath)
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
	db := backRepo.BackRepoSPECIFICATION_TYPE.GetDB()

	// Get model if exist
	var specification_typeDB orm.SPECIFICATION_TYPEDB
	if _, err := db.First(&specification_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&specification_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	specification_typeDeleted := new(models.SPECIFICATION_TYPE)
	specification_typeDB.CopyBasicFieldsToSPECIFICATION_TYPE(specification_typeDeleted)

	// get stage instance from DB instance, and call callback function
	specification_typeStaged := backRepo.BackRepoSPECIFICATION_TYPE.Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEPtr[specification_typeDB.ID]
	if specification_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), specification_typeStaged, specification_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
