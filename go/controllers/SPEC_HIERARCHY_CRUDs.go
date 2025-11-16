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
var __SPEC_HIERARCHY__dummysDeclaration__ models.SPEC_HIERARCHY
var __SPEC_HIERARCHY_time__dummyDeclaration time.Duration

var mutexSPEC_HIERARCHY sync.Mutex

// An SPEC_HIERARCHYID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getSPEC_HIERARCHY updateSPEC_HIERARCHY deleteSPEC_HIERARCHY
type SPEC_HIERARCHYID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// SPEC_HIERARCHYInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postSPEC_HIERARCHY updateSPEC_HIERARCHY
type SPEC_HIERARCHYInput struct {
	// The SPEC_HIERARCHY to submit or modify
	// in: body
	SPEC_HIERARCHY *orm.SPEC_HIERARCHYAPI
}

// GetSPEC_HIERARCHYs
//
// swagger:route GET /spec_hierarchys spec_hierarchys getSPEC_HIERARCHYs
//
// # Get all spec_hierarchys
//
// Responses:
// default: genericError
//
//	200: spec_hierarchyDBResponse
func (controller *Controller) GetSPEC_HIERARCHYs(c *gin.Context) {

	// source slice
	var spec_hierarchyDBs []orm.SPEC_HIERARCHYDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_HIERARCHYs", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_HIERARCHY.GetDB()

	_, err := db.Find(&spec_hierarchyDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	spec_hierarchyAPIs := make([]orm.SPEC_HIERARCHYAPI, 0)

	// for each spec_hierarchy, update fields from the database nullable fields
	for idx := range spec_hierarchyDBs {
		spec_hierarchyDB := &spec_hierarchyDBs[idx]
		_ = spec_hierarchyDB
		var spec_hierarchyAPI orm.SPEC_HIERARCHYAPI

		// insertion point for updating fields
		spec_hierarchyAPI.ID = spec_hierarchyDB.ID
		spec_hierarchyDB.CopyBasicFieldsToSPEC_HIERARCHY_WOP(&spec_hierarchyAPI.SPEC_HIERARCHY_WOP)
		spec_hierarchyAPI.SPEC_HIERARCHYPointersEncoding = spec_hierarchyDB.SPEC_HIERARCHYPointersEncoding
		spec_hierarchyAPIs = append(spec_hierarchyAPIs, spec_hierarchyAPI)
	}

	c.JSON(http.StatusOK, spec_hierarchyAPIs)
}

// PostSPEC_HIERARCHY
//
// swagger:route POST /spec_hierarchys spec_hierarchys postSPEC_HIERARCHY
//
// Creates a spec_hierarchy
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostSPEC_HIERARCHY(c *gin.Context) {

	mutexSPEC_HIERARCHY.Lock()
	defer mutexSPEC_HIERARCHY.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostSPEC_HIERARCHYs", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_HIERARCHY.GetDB()

	// Validate input
	var input orm.SPEC_HIERARCHYAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create spec_hierarchy
	spec_hierarchyDB := orm.SPEC_HIERARCHYDB{}
	spec_hierarchyDB.SPEC_HIERARCHYPointersEncoding = input.SPEC_HIERARCHYPointersEncoding
	spec_hierarchyDB.CopyBasicFieldsFromSPEC_HIERARCHY_WOP(&input.SPEC_HIERARCHY_WOP)

	_, err = db.Create(&spec_hierarchyDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoSPEC_HIERARCHY.CheckoutPhaseOneInstance(&spec_hierarchyDB)
	spec_hierarchy := backRepo.BackRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr[spec_hierarchyDB.ID]

	if spec_hierarchy != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), spec_hierarchy)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, spec_hierarchyDB)
}

// GetSPEC_HIERARCHY
//
// swagger:route GET /spec_hierarchys/{ID} spec_hierarchys getSPEC_HIERARCHY
//
// Gets the details for a spec_hierarchy.
//
// Responses:
// default: genericError
//
//	200: spec_hierarchyDBResponse
func (controller *Controller) GetSPEC_HIERARCHY(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetSPEC_HIERARCHY", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_HIERARCHY.GetDB()

	// Get spec_hierarchyDB in DB
	var spec_hierarchyDB orm.SPEC_HIERARCHYDB
	if _, err := db.First(&spec_hierarchyDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var spec_hierarchyAPI orm.SPEC_HIERARCHYAPI
	spec_hierarchyAPI.ID = spec_hierarchyDB.ID
	spec_hierarchyAPI.SPEC_HIERARCHYPointersEncoding = spec_hierarchyDB.SPEC_HIERARCHYPointersEncoding
	spec_hierarchyDB.CopyBasicFieldsToSPEC_HIERARCHY_WOP(&spec_hierarchyAPI.SPEC_HIERARCHY_WOP)

	c.JSON(http.StatusOK, spec_hierarchyAPI)
}

// UpdateSPEC_HIERARCHY
//
// swagger:route PATCH /spec_hierarchys/{ID} spec_hierarchys updateSPEC_HIERARCHY
//
// # Update a spec_hierarchy
//
// Responses:
// default: genericError
//
//	200: spec_hierarchyDBResponse
func (controller *Controller) UpdateSPEC_HIERARCHY(c *gin.Context) {

	mutexSPEC_HIERARCHY.Lock()
	defer mutexSPEC_HIERARCHY.Unlock()

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
	db := backRepo.BackRepoSPEC_HIERARCHY.GetDB()

	// Validate input
	var input orm.SPEC_HIERARCHYAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var spec_hierarchyDB orm.SPEC_HIERARCHYDB

	// fetch the spec_hierarchy
	_, err := db.First(&spec_hierarchyDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	spec_hierarchyDB.CopyBasicFieldsFromSPEC_HIERARCHY_WOP(&input.SPEC_HIERARCHY_WOP)
	spec_hierarchyDB.SPEC_HIERARCHYPointersEncoding = input.SPEC_HIERARCHYPointersEncoding

	db, _ = db.Model(&spec_hierarchyDB)
	_, err = db.Updates(&spec_hierarchyDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	spec_hierarchyNew := new(models.SPEC_HIERARCHY)
	spec_hierarchyDB.CopyBasicFieldsToSPEC_HIERARCHY(spec_hierarchyNew)

	// redeem pointers
	spec_hierarchyDB.DecodePointers(backRepo, spec_hierarchyNew)

	// get stage instance from DB instance, and call callback function
	spec_hierarchyOld := backRepo.BackRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr[spec_hierarchyDB.ID]
	if spec_hierarchyOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), spec_hierarchyOld, spec_hierarchyNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the spec_hierarchyDB
	c.JSON(http.StatusOK, spec_hierarchyDB)
}

// DeleteSPEC_HIERARCHY
//
// swagger:route DELETE /spec_hierarchys/{ID} spec_hierarchys deleteSPEC_HIERARCHY
//
// # Delete a spec_hierarchy
//
// default: genericError
//
//	200: spec_hierarchyDBResponse
func (controller *Controller) DeleteSPEC_HIERARCHY(c *gin.Context) {

	mutexSPEC_HIERARCHY.Lock()
	defer mutexSPEC_HIERARCHY.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteSPEC_HIERARCHY", "Name", stackPath)
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
	db := backRepo.BackRepoSPEC_HIERARCHY.GetDB()

	// Get model if exist
	var spec_hierarchyDB orm.SPEC_HIERARCHYDB
	if _, err := db.First(&spec_hierarchyDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&spec_hierarchyDB)

	// get an instance (not staged) from DB instance, and call callback function
	spec_hierarchyDeleted := new(models.SPEC_HIERARCHY)
	spec_hierarchyDB.CopyBasicFieldsToSPEC_HIERARCHY(spec_hierarchyDeleted)

	// get stage instance from DB instance, and call callback function
	spec_hierarchyStaged := backRepo.BackRepoSPEC_HIERARCHY.Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr[spec_hierarchyDB.ID]
	if spec_hierarchyStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), spec_hierarchyStaged, spec_hierarchyDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
