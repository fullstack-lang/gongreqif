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
var __SPEC_RELATION_TYPE__dummysDeclaration__ models.SPEC_RELATION_TYPE
var __SPEC_RELATION_TYPE_time__dummyDeclaration time.Duration

var mutexSPEC_RELATION_TYPE sync.Mutex

// An SPEC_RELATION_TYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPEC_RELATION_TYPE updateSPEC_RELATION_TYPE deleteSPEC_RELATION_TYPE
type SPEC_RELATION_TYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPEC_RELATION_TYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPEC_RELATION_TYPE updateSPEC_RELATION_TYPE
type SPEC_RELATION_TYPEInput struct {
	// The SPEC_RELATION_TYPE to submit or modify
	// in: body
	SPEC_RELATION_TYPE *orm.SPEC_RELATION_TYPEAPI
}

// GetSPEC_RELATION_TYPEs
//
// swagger:route GET /spec_relation_types spec_relation_types getSPEC_RELATION_TYPEs
//
// # Get all spec_relation_types
//
// Responses:
// default: genericError
//
//	200: spec_relation_typeDBResponse
func (controller *Controller) GetSPEC_RELATION_TYPEs(c *gin.Context) {

	// source slice
	var spec_relation_typeDBs []orm.SPEC_RELATION_TYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_RELATION_TYPEs", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_RELATION_TYPE.GetDB()

	_, err := db.Find(&spec_relation_typeDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spec_relation_typeAPIs := make([]orm.SPEC_RELATION_TYPEAPI, 0)

	// for each spec_relation_type, update fields from the database nullable fields
	for idx := range spec_relation_typeDBs {
		spec_relation_typeDB := &spec_relation_typeDBs[idx]
		_ = spec_relation_typeDB
		var spec_relation_typeAPI orm.SPEC_RELATION_TYPEAPI

		// insertion point for updating fields
		spec_relation_typeAPI.ID = spec_relation_typeDB.ID
		spec_relation_typeDB.CopyBasicFieldsToSPEC_RELATION_TYPE_WOP(&spec_relation_typeAPI.SPEC_RELATION_TYPE_WOP)
		spec_relation_typeAPI.SPEC_RELATION_TYPEPointersEncoding = spec_relation_typeDB.SPEC_RELATION_TYPEPointersEncoding
		spec_relation_typeAPIs = append(spec_relation_typeAPIs, spec_relation_typeAPI)
	}

	c.JSON(http.StatusOK, spec_relation_typeAPIs)
}

// PostSPEC_RELATION_TYPE
//
// swagger:route POST /spec_relation_types spec_relation_types postSPEC_RELATION_TYPE
//
// Creates a spec_relation_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPEC_RELATION_TYPE(c *gin.Context) {

	mutexSPEC_RELATION_TYPE.Lock()
	defer mutexSPEC_RELATION_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPEC_RELATION_TYPEs", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_RELATION_TYPE.GetDB()

	// Validate input
	var input orm.SPEC_RELATION_TYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spec_relation_type
	spec_relation_typeDB := orm.SPEC_RELATION_TYPEDB{}
	spec_relation_typeDB.SPEC_RELATION_TYPEPointersEncoding = input.SPEC_RELATION_TYPEPointersEncoding
	spec_relation_typeDB.CopyBasicFieldsFromSPEC_RELATION_TYPE_WOP(&input.SPEC_RELATION_TYPE_WOP)

	_, err = db.Create(&spec_relation_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPEC_RELATION_TYPE.CheckoutPhaseOneInstance(&spec_relation_typeDB)
	spec_relation_type := backRepo.BackRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[spec_relation_typeDB.ID]

	if spec_relation_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spec_relation_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spec_relation_typeDB)
}

// GetSPEC_RELATION_TYPE
//
// swagger:route GET /spec_relation_types/{ID} spec_relation_types getSPEC_RELATION_TYPE
//
// Gets the details for a spec_relation_type.
//
// Responses:
// default: genericError
//
//	200: spec_relation_typeDBResponse
func (controller *Controller) GetSPEC_RELATION_TYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_RELATION_TYPE", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_RELATION_TYPE.GetDB()

	// Get spec_relation_typeDB in DB
	var spec_relation_typeDB orm.SPEC_RELATION_TYPEDB
	if _, err := db.First(&spec_relation_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spec_relation_typeAPI orm.SPEC_RELATION_TYPEAPI
	spec_relation_typeAPI.ID = spec_relation_typeDB.ID
	spec_relation_typeAPI.SPEC_RELATION_TYPEPointersEncoding = spec_relation_typeDB.SPEC_RELATION_TYPEPointersEncoding
	spec_relation_typeDB.CopyBasicFieldsToSPEC_RELATION_TYPE_WOP(&spec_relation_typeAPI.SPEC_RELATION_TYPE_WOP)

	c.JSON(http.StatusOK, spec_relation_typeAPI)
}

// UpdateSPEC_RELATION_TYPE
//
// swagger:route PATCH /spec_relation_types/{ID} spec_relation_types updateSPEC_RELATION_TYPE
//
// # Update a spec_relation_type
//
// Responses:
// default: genericError
//
//	200: spec_relation_typeDBResponse
func (controller *Controller) UpdateSPEC_RELATION_TYPE(c *gin.Context) {

	mutexSPEC_RELATION_TYPE.Lock()
	defer mutexSPEC_RELATION_TYPE.Unlock()

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
	db := backRepo.BackRepoSPEC_RELATION_TYPE.GetDB()

	// Validate input
	var input orm.SPEC_RELATION_TYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spec_relation_typeDB orm.SPEC_RELATION_TYPEDB

	// fetch the spec_relation_type
	_, err := db.First(&spec_relation_typeDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spec_relation_typeDB.CopyBasicFieldsFromSPEC_RELATION_TYPE_WOP(&input.SPEC_RELATION_TYPE_WOP)
	spec_relation_typeDB.SPEC_RELATION_TYPEPointersEncoding = input.SPEC_RELATION_TYPEPointersEncoding

	db, _ = db.Model(&spec_relation_typeDB)
	_, err = db.Updates(&spec_relation_typeDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spec_relation_typeNew := new(models.SPEC_RELATION_TYPE)
	spec_relation_typeDB.CopyBasicFieldsToSPEC_RELATION_TYPE(spec_relation_typeNew)

	// redeem pointers
	spec_relation_typeDB.DecodePointers(backRepo, spec_relation_typeNew)

	// get stage instance from DB instance, and call callback function
	spec_relation_typeOld := backRepo.BackRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[spec_relation_typeDB.ID]
	if spec_relation_typeOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), spec_relation_typeOld, spec_relation_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spec_relation_typeDB
	c.JSON(http.StatusOK, spec_relation_typeDB)
}

// DeleteSPEC_RELATION_TYPE
//
// swagger:route DELETE /spec_relation_types/{ID} spec_relation_types deleteSPEC_RELATION_TYPE
//
// # Delete a spec_relation_type
//
// default: genericError
//
//	200: spec_relation_typeDBResponse
func (controller *Controller) DeleteSPEC_RELATION_TYPE(c *gin.Context) {

	mutexSPEC_RELATION_TYPE.Lock()
	defer mutexSPEC_RELATION_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPEC_RELATION_TYPE", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_RELATION_TYPE.GetDB()

	// Get model if exist
	var spec_relation_typeDB orm.SPEC_RELATION_TYPEDB
	if _, err := db.First(&spec_relation_typeDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spec_relation_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	spec_relation_typeDeleted := new(models.SPEC_RELATION_TYPE)
	spec_relation_typeDB.CopyBasicFieldsToSPEC_RELATION_TYPE(spec_relation_typeDeleted)

	// get stage instance from DB instance, and call callback function
	spec_relation_typeStaged := backRepo.BackRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[spec_relation_typeDB.ID]
	if spec_relation_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spec_relation_typeStaged, spec_relation_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
