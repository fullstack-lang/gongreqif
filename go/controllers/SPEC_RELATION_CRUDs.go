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
var __SPEC_RELATION__dummysDeclaration__ models.SPEC_RELATION
var __SPEC_RELATION_time__dummyDeclaration time.Duration

var mutexSPEC_RELATION sync.Mutex

// An SPEC_RELATIONID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPEC_RELATION updateSPEC_RELATION deleteSPEC_RELATION
type SPEC_RELATIONID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPEC_RELATIONInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPEC_RELATION updateSPEC_RELATION
type SPEC_RELATIONInput struct {
	// The SPEC_RELATION to submit or modify
	// in: body
	SPEC_RELATION *orm.SPEC_RELATIONAPI
}

// GetSPEC_RELATIONs
//
// swagger:route GET /spec_relations spec_relations getSPEC_RELATIONs
//
// # Get all spec_relations
//
// Responses:
// default: genericError
//
//	200: spec_relationDBResponse
func (controller *Controller) GetSPEC_RELATIONs(c *gin.Context) {

	// source slice
	var spec_relationDBs []orm.SPEC_RELATIONDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_RELATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_RELATION.GetDB()

	_, err := db.Find(&spec_relationDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spec_relationAPIs := make([]orm.SPEC_RELATIONAPI, 0)

	// for each spec_relation, update fields from the database nullable fields
	for idx := range spec_relationDBs {
		spec_relationDB := &spec_relationDBs[idx]
		_ = spec_relationDB
		var spec_relationAPI orm.SPEC_RELATIONAPI

		// insertion point for updating fields
		spec_relationAPI.ID = spec_relationDB.ID
		spec_relationDB.CopyBasicFieldsToSPEC_RELATION_WOP(&spec_relationAPI.SPEC_RELATION_WOP)
		spec_relationAPI.SPEC_RELATIONPointersEncoding = spec_relationDB.SPEC_RELATIONPointersEncoding
		spec_relationAPIs = append(spec_relationAPIs, spec_relationAPI)
	}

	c.JSON(http.StatusOK, spec_relationAPIs)
}

// PostSPEC_RELATION
//
// swagger:route POST /spec_relations spec_relations postSPEC_RELATION
//
// Creates a spec_relation
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPEC_RELATION(c *gin.Context) {

	mutexSPEC_RELATION.Lock()
	defer mutexSPEC_RELATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPEC_RELATIONs", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_RELATION.GetDB()

	// Validate input
	var input orm.SPEC_RELATIONAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spec_relation
	spec_relationDB := orm.SPEC_RELATIONDB{}
	spec_relationDB.SPEC_RELATIONPointersEncoding = input.SPEC_RELATIONPointersEncoding
	spec_relationDB.CopyBasicFieldsFromSPEC_RELATION_WOP(&input.SPEC_RELATION_WOP)

	_, err = db.Create(&spec_relationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPEC_RELATION.CheckoutPhaseOneInstance(&spec_relationDB)
	spec_relation := backRepo.BackRepoSPEC_RELATION.Map_SPEC_RELATIONDBID_SPEC_RELATIONPtr[spec_relationDB.ID]

	if spec_relation != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spec_relation)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spec_relationDB)
}

// GetSPEC_RELATION
//
// swagger:route GET /spec_relations/{ID} spec_relations getSPEC_RELATION
//
// Gets the details for a spec_relation.
//
// Responses:
// default: genericError
//
//	200: spec_relationDBResponse
func (controller *Controller) GetSPEC_RELATION(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_RELATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_RELATION.GetDB()

	// Get spec_relationDB in DB
	var spec_relationDB orm.SPEC_RELATIONDB
	if _, err := db.First(&spec_relationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spec_relationAPI orm.SPEC_RELATIONAPI
	spec_relationAPI.ID = spec_relationDB.ID
	spec_relationAPI.SPEC_RELATIONPointersEncoding = spec_relationDB.SPEC_RELATIONPointersEncoding
	spec_relationDB.CopyBasicFieldsToSPEC_RELATION_WOP(&spec_relationAPI.SPEC_RELATION_WOP)

	c.JSON(http.StatusOK, spec_relationAPI)
}

// UpdateSPEC_RELATION
//
// swagger:route PATCH /spec_relations/{ID} spec_relations updateSPEC_RELATION
//
// # Update a spec_relation
//
// Responses:
// default: genericError
//
//	200: spec_relationDBResponse
func (controller *Controller) UpdateSPEC_RELATION(c *gin.Context) {

	mutexSPEC_RELATION.Lock()
	defer mutexSPEC_RELATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateSPEC_RELATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_RELATION.GetDB()

	// Validate input
	var input orm.SPEC_RELATIONAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spec_relationDB orm.SPEC_RELATIONDB

	// fetch the spec_relation
	_, err := db.First(&spec_relationDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spec_relationDB.CopyBasicFieldsFromSPEC_RELATION_WOP(&input.SPEC_RELATION_WOP)
	spec_relationDB.SPEC_RELATIONPointersEncoding = input.SPEC_RELATIONPointersEncoding

	db, _ = db.Model(&spec_relationDB)
	_, err = db.Updates(&spec_relationDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spec_relationNew := new(models.SPEC_RELATION)
	spec_relationDB.CopyBasicFieldsToSPEC_RELATION(spec_relationNew)

	// redeem pointers
	spec_relationDB.DecodePointers(backRepo, spec_relationNew)

	// get stage instance from DB instance, and call callback function
	spec_relationOld := backRepo.BackRepoSPEC_RELATION.Map_SPEC_RELATIONDBID_SPEC_RELATIONPtr[spec_relationDB.ID]
	if spec_relationOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), spec_relationOld, spec_relationNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spec_relationDB
	c.JSON(http.StatusOK, spec_relationDB)
}

// DeleteSPEC_RELATION
//
// swagger:route DELETE /spec_relations/{ID} spec_relations deleteSPEC_RELATION
//
// # Delete a spec_relation
//
// default: genericError
//
//	200: spec_relationDBResponse
func (controller *Controller) DeleteSPEC_RELATION(c *gin.Context) {

	mutexSPEC_RELATION.Lock()
	defer mutexSPEC_RELATION.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPEC_RELATION", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoSPEC_RELATION.GetDB()

	// Get model if exist
	var spec_relationDB orm.SPEC_RELATIONDB
	if _, err := db.First(&spec_relationDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spec_relationDB)

	// get an instance (not staged) from DB instance, and call callback function
	spec_relationDeleted := new(models.SPEC_RELATION)
	spec_relationDB.CopyBasicFieldsToSPEC_RELATION(spec_relationDeleted)

	// get stage instance from DB instance, and call callback function
	spec_relationStaged := backRepo.BackRepoSPEC_RELATION.Map_SPEC_RELATIONDBID_SPEC_RELATIONPtr[spec_relationDB.ID]
	if spec_relationStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spec_relationStaged, spec_relationDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
