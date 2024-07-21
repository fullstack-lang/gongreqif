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
var __SPEC_OBJECT_TYPE__dummysDeclaration__ models.SPEC_OBJECT_TYPE
var __SPEC_OBJECT_TYPE_time__dummyDeclaration time.Duration

var mutexSPEC_OBJECT_TYPE sync.Mutex

// An SPEC_OBJECT_TYPEID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPEC_OBJECT_TYPE updateSPEC_OBJECT_TYPE deleteSPEC_OBJECT_TYPE
type SPEC_OBJECT_TYPEID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPEC_OBJECT_TYPEInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPEC_OBJECT_TYPE updateSPEC_OBJECT_TYPE
type SPEC_OBJECT_TYPEInput struct {
	// The SPEC_OBJECT_TYPE to submit or modify
	// in: body
	SPEC_OBJECT_TYPE *orm.SPEC_OBJECT_TYPEAPI
}

// GetSPEC_OBJECT_TYPEs
//
// swagger:route GET /spec_object_types spec_object_types getSPEC_OBJECT_TYPEs
//
// # Get all spec_object_types
//
// Responses:
// default: genericError
//
//	200: spec_object_typeDBResponse
func (controller *Controller) GetSPEC_OBJECT_TYPEs(c *gin.Context) {

	// source slice
	var spec_object_typeDBs []orm.SPEC_OBJECT_TYPEDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_OBJECT_TYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_OBJECT_TYPE.GetDB()

	query := db.Find(&spec_object_typeDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spec_object_typeAPIs := make([]orm.SPEC_OBJECT_TYPEAPI, 0)

	// for each spec_object_type, update fields from the database nullable fields
	for idx := range spec_object_typeDBs {
		spec_object_typeDB := &spec_object_typeDBs[idx]
		_ = spec_object_typeDB
		var spec_object_typeAPI orm.SPEC_OBJECT_TYPEAPI

		// insertion point for updating fields
		spec_object_typeAPI.ID = spec_object_typeDB.ID
		spec_object_typeDB.CopyBasicFieldsToSPEC_OBJECT_TYPE_WOP(&spec_object_typeAPI.SPEC_OBJECT_TYPE_WOP)
		spec_object_typeAPI.SPEC_OBJECT_TYPEPointersEncoding = spec_object_typeDB.SPEC_OBJECT_TYPEPointersEncoding
		spec_object_typeAPIs = append(spec_object_typeAPIs, spec_object_typeAPI)
	}

	c.JSON(http.StatusOK, spec_object_typeAPIs)
}

// PostSPEC_OBJECT_TYPE
//
// swagger:route POST /spec_object_types spec_object_types postSPEC_OBJECT_TYPE
//
// Creates a spec_object_type
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPEC_OBJECT_TYPE(c *gin.Context) {

	mutexSPEC_OBJECT_TYPE.Lock()
	defer mutexSPEC_OBJECT_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPEC_OBJECT_TYPEs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_OBJECT_TYPE.GetDB()

	// Validate input
	var input orm.SPEC_OBJECT_TYPEAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spec_object_type
	spec_object_typeDB := orm.SPEC_OBJECT_TYPEDB{}
	spec_object_typeDB.SPEC_OBJECT_TYPEPointersEncoding = input.SPEC_OBJECT_TYPEPointersEncoding
	spec_object_typeDB.CopyBasicFieldsFromSPEC_OBJECT_TYPE_WOP(&input.SPEC_OBJECT_TYPE_WOP)

	query := db.Create(&spec_object_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPEC_OBJECT_TYPE.CheckoutPhaseOneInstance(&spec_object_typeDB)
	spec_object_type := backRepo.BackRepoSPEC_OBJECT_TYPE.Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEPtr[spec_object_typeDB.ID]

	if spec_object_type != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spec_object_type)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spec_object_typeDB)
}

// GetSPEC_OBJECT_TYPE
//
// swagger:route GET /spec_object_types/{ID} spec_object_types getSPEC_OBJECT_TYPE
//
// Gets the details for a spec_object_type.
//
// Responses:
// default: genericError
//
//	200: spec_object_typeDBResponse
func (controller *Controller) GetSPEC_OBJECT_TYPE(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_OBJECT_TYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_OBJECT_TYPE.GetDB()

	// Get spec_object_typeDB in DB
	var spec_object_typeDB orm.SPEC_OBJECT_TYPEDB
	if err := db.First(&spec_object_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spec_object_typeAPI orm.SPEC_OBJECT_TYPEAPI
	spec_object_typeAPI.ID = spec_object_typeDB.ID
	spec_object_typeAPI.SPEC_OBJECT_TYPEPointersEncoding = spec_object_typeDB.SPEC_OBJECT_TYPEPointersEncoding
	spec_object_typeDB.CopyBasicFieldsToSPEC_OBJECT_TYPE_WOP(&spec_object_typeAPI.SPEC_OBJECT_TYPE_WOP)

	c.JSON(http.StatusOK, spec_object_typeAPI)
}

// UpdateSPEC_OBJECT_TYPE
//
// swagger:route PATCH /spec_object_types/{ID} spec_object_types updateSPEC_OBJECT_TYPE
//
// # Update a spec_object_type
//
// Responses:
// default: genericError
//
//	200: spec_object_typeDBResponse
func (controller *Controller) UpdateSPEC_OBJECT_TYPE(c *gin.Context) {

	mutexSPEC_OBJECT_TYPE.Lock()
	defer mutexSPEC_OBJECT_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPEC_OBJECT_TYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_OBJECT_TYPE.GetDB()

	// Validate input
	var input orm.SPEC_OBJECT_TYPEAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spec_object_typeDB orm.SPEC_OBJECT_TYPEDB

	// fetch the spec_object_type
	query := db.First(&spec_object_typeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spec_object_typeDB.CopyBasicFieldsFromSPEC_OBJECT_TYPE_WOP(&input.SPEC_OBJECT_TYPE_WOP)
	spec_object_typeDB.SPEC_OBJECT_TYPEPointersEncoding = input.SPEC_OBJECT_TYPEPointersEncoding

	query = db.Model(&spec_object_typeDB).Updates(spec_object_typeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spec_object_typeNew := new(models.SPEC_OBJECT_TYPE)
	spec_object_typeDB.CopyBasicFieldsToSPEC_OBJECT_TYPE(spec_object_typeNew)

	// redeem pointers
	spec_object_typeDB.DecodePointers(backRepo, spec_object_typeNew)

	// get stage instance from DB instance, and call callback function
	spec_object_typeOld := backRepo.BackRepoSPEC_OBJECT_TYPE.Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEPtr[spec_object_typeDB.ID]
	if spec_object_typeOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spec_object_typeOld, spec_object_typeNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spec_object_typeDB
	c.JSON(http.StatusOK, spec_object_typeDB)
}

// DeleteSPEC_OBJECT_TYPE
//
// swagger:route DELETE /spec_object_types/{ID} spec_object_types deleteSPEC_OBJECT_TYPE
//
// # Delete a spec_object_type
//
// default: genericError
//
//	200: spec_object_typeDBResponse
func (controller *Controller) DeleteSPEC_OBJECT_TYPE(c *gin.Context) {

	mutexSPEC_OBJECT_TYPE.Lock()
	defer mutexSPEC_OBJECT_TYPE.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPEC_OBJECT_TYPE", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_OBJECT_TYPE.GetDB()

	// Get model if exist
	var spec_object_typeDB orm.SPEC_OBJECT_TYPEDB
	if err := db.First(&spec_object_typeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&spec_object_typeDB)

	// get an instance (not staged) from DB instance, and call callback function
	spec_object_typeDeleted := new(models.SPEC_OBJECT_TYPE)
	spec_object_typeDB.CopyBasicFieldsToSPEC_OBJECT_TYPE(spec_object_typeDeleted)

	// get stage instance from DB instance, and call callback function
	spec_object_typeStaged := backRepo.BackRepoSPEC_OBJECT_TYPE.Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEPtr[spec_object_typeDB.ID]
	if spec_object_typeStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spec_object_typeStaged, spec_object_typeDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
