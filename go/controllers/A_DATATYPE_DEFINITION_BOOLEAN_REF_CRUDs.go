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
var __A_DATATYPE_DEFINITION_BOOLEAN_REF__dummysDeclaration__ models.A_DATATYPE_DEFINITION_BOOLEAN_REF
var __A_DATATYPE_DEFINITION_BOOLEAN_REF_time__dummyDeclaration time.Duration

var mutexA_DATATYPE_DEFINITION_BOOLEAN_REF sync.Mutex

// An A_DATATYPE_DEFINITION_BOOLEAN_REFID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_DATATYPE_DEFINITION_BOOLEAN_REF updateA_DATATYPE_DEFINITION_BOOLEAN_REF deleteA_DATATYPE_DEFINITION_BOOLEAN_REF
type A_DATATYPE_DEFINITION_BOOLEAN_REFID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_DATATYPE_DEFINITION_BOOLEAN_REFInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_DATATYPE_DEFINITION_BOOLEAN_REF updateA_DATATYPE_DEFINITION_BOOLEAN_REF
type A_DATATYPE_DEFINITION_BOOLEAN_REFInput struct {
	// The A_DATATYPE_DEFINITION_BOOLEAN_REF to submit or modify
	// in: body
	A_DATATYPE_DEFINITION_BOOLEAN_REF *orm.A_DATATYPE_DEFINITION_BOOLEAN_REFAPI
}

// GetA_DATATYPE_DEFINITION_BOOLEAN_REFs
//
// swagger:route GET /a_datatype_definition_boolean_refs a_datatype_definition_boolean_refs getA_DATATYPE_DEFINITION_BOOLEAN_REFs
//
// # Get all a_datatype_definition_boolean_refs
//
// Responses:
// default: genericError
//
//	200: a_datatype_definition_boolean_refDBResponse
func (controller *Controller) GetA_DATATYPE_DEFINITION_BOOLEAN_REFs(c *gin.Context) {

	// source slice
	var a_datatype_definition_boolean_refDBs []orm.A_DATATYPE_DEFINITION_BOOLEAN_REFDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DATATYPE_DEFINITION_BOOLEAN_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.GetDB()

	_, err := db.Find(&a_datatype_definition_boolean_refDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_datatype_definition_boolean_refAPIs := make([]orm.A_DATATYPE_DEFINITION_BOOLEAN_REFAPI, 0)

	// for each a_datatype_definition_boolean_ref, update fields from the database nullable fields
	for idx := range a_datatype_definition_boolean_refDBs {
		a_datatype_definition_boolean_refDB := &a_datatype_definition_boolean_refDBs[idx]
		_ = a_datatype_definition_boolean_refDB
		var a_datatype_definition_boolean_refAPI orm.A_DATATYPE_DEFINITION_BOOLEAN_REFAPI

		// insertion point for updating fields
		a_datatype_definition_boolean_refAPI.ID = a_datatype_definition_boolean_refDB.ID
		a_datatype_definition_boolean_refDB.CopyBasicFieldsToA_DATATYPE_DEFINITION_BOOLEAN_REF_WOP(&a_datatype_definition_boolean_refAPI.A_DATATYPE_DEFINITION_BOOLEAN_REF_WOP)
		a_datatype_definition_boolean_refAPI.A_DATATYPE_DEFINITION_BOOLEAN_REFPointersEncoding = a_datatype_definition_boolean_refDB.A_DATATYPE_DEFINITION_BOOLEAN_REFPointersEncoding
		a_datatype_definition_boolean_refAPIs = append(a_datatype_definition_boolean_refAPIs, a_datatype_definition_boolean_refAPI)
	}

	c.JSON(http.StatusOK, a_datatype_definition_boolean_refAPIs)
}

// PostA_DATATYPE_DEFINITION_BOOLEAN_REF
//
// swagger:route POST /a_datatype_definition_boolean_refs a_datatype_definition_boolean_refs postA_DATATYPE_DEFINITION_BOOLEAN_REF
//
// Creates a a_datatype_definition_boolean_ref
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_DATATYPE_DEFINITION_BOOLEAN_REF(c *gin.Context) {

	mutexA_DATATYPE_DEFINITION_BOOLEAN_REF.Lock()
	defer mutexA_DATATYPE_DEFINITION_BOOLEAN_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_DATATYPE_DEFINITION_BOOLEAN_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.GetDB()

	// Validate input
	var input orm.A_DATATYPE_DEFINITION_BOOLEAN_REFAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_datatype_definition_boolean_ref
	a_datatype_definition_boolean_refDB := orm.A_DATATYPE_DEFINITION_BOOLEAN_REFDB{}
	a_datatype_definition_boolean_refDB.A_DATATYPE_DEFINITION_BOOLEAN_REFPointersEncoding = input.A_DATATYPE_DEFINITION_BOOLEAN_REFPointersEncoding
	a_datatype_definition_boolean_refDB.CopyBasicFieldsFromA_DATATYPE_DEFINITION_BOOLEAN_REF_WOP(&input.A_DATATYPE_DEFINITION_BOOLEAN_REF_WOP)

	_, err = db.Create(&a_datatype_definition_boolean_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.CheckoutPhaseOneInstance(&a_datatype_definition_boolean_refDB)
	a_datatype_definition_boolean_ref := backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.Map_A_DATATYPE_DEFINITION_BOOLEAN_REFDBID_A_DATATYPE_DEFINITION_BOOLEAN_REFPtr[a_datatype_definition_boolean_refDB.ID]

	if a_datatype_definition_boolean_ref != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_datatype_definition_boolean_ref)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_datatype_definition_boolean_refDB)
}

// GetA_DATATYPE_DEFINITION_BOOLEAN_REF
//
// swagger:route GET /a_datatype_definition_boolean_refs/{ID} a_datatype_definition_boolean_refs getA_DATATYPE_DEFINITION_BOOLEAN_REF
//
// Gets the details for a a_datatype_definition_boolean_ref.
//
// Responses:
// default: genericError
//
//	200: a_datatype_definition_boolean_refDBResponse
func (controller *Controller) GetA_DATATYPE_DEFINITION_BOOLEAN_REF(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_DATATYPE_DEFINITION_BOOLEAN_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.GetDB()

	// Get a_datatype_definition_boolean_refDB in DB
	var a_datatype_definition_boolean_refDB orm.A_DATATYPE_DEFINITION_BOOLEAN_REFDB
	if _, err := db.First(&a_datatype_definition_boolean_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_datatype_definition_boolean_refAPI orm.A_DATATYPE_DEFINITION_BOOLEAN_REFAPI
	a_datatype_definition_boolean_refAPI.ID = a_datatype_definition_boolean_refDB.ID
	a_datatype_definition_boolean_refAPI.A_DATATYPE_DEFINITION_BOOLEAN_REFPointersEncoding = a_datatype_definition_boolean_refDB.A_DATATYPE_DEFINITION_BOOLEAN_REFPointersEncoding
	a_datatype_definition_boolean_refDB.CopyBasicFieldsToA_DATATYPE_DEFINITION_BOOLEAN_REF_WOP(&a_datatype_definition_boolean_refAPI.A_DATATYPE_DEFINITION_BOOLEAN_REF_WOP)

	c.JSON(http.StatusOK, a_datatype_definition_boolean_refAPI)
}

// UpdateA_DATATYPE_DEFINITION_BOOLEAN_REF
//
// swagger:route PATCH /a_datatype_definition_boolean_refs/{ID} a_datatype_definition_boolean_refs updateA_DATATYPE_DEFINITION_BOOLEAN_REF
//
// # Update a a_datatype_definition_boolean_ref
//
// Responses:
// default: genericError
//
//	200: a_datatype_definition_boolean_refDBResponse
func (controller *Controller) UpdateA_DATATYPE_DEFINITION_BOOLEAN_REF(c *gin.Context) {

	mutexA_DATATYPE_DEFINITION_BOOLEAN_REF.Lock()
	defer mutexA_DATATYPE_DEFINITION_BOOLEAN_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_DATATYPE_DEFINITION_BOOLEAN_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.GetDB()

	// Validate input
	var input orm.A_DATATYPE_DEFINITION_BOOLEAN_REFAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_datatype_definition_boolean_refDB orm.A_DATATYPE_DEFINITION_BOOLEAN_REFDB

	// fetch the a_datatype_definition_boolean_ref
	_, err := db.First(&a_datatype_definition_boolean_refDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_datatype_definition_boolean_refDB.CopyBasicFieldsFromA_DATATYPE_DEFINITION_BOOLEAN_REF_WOP(&input.A_DATATYPE_DEFINITION_BOOLEAN_REF_WOP)
	a_datatype_definition_boolean_refDB.A_DATATYPE_DEFINITION_BOOLEAN_REFPointersEncoding = input.A_DATATYPE_DEFINITION_BOOLEAN_REFPointersEncoding

	db, _ = db.Model(&a_datatype_definition_boolean_refDB)
	_, err = db.Updates(&a_datatype_definition_boolean_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_datatype_definition_boolean_refNew := new(models.A_DATATYPE_DEFINITION_BOOLEAN_REF)
	a_datatype_definition_boolean_refDB.CopyBasicFieldsToA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_refNew)

	// redeem pointers
	a_datatype_definition_boolean_refDB.DecodePointers(backRepo, a_datatype_definition_boolean_refNew)

	// get stage instance from DB instance, and call callback function
	a_datatype_definition_boolean_refOld := backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.Map_A_DATATYPE_DEFINITION_BOOLEAN_REFDBID_A_DATATYPE_DEFINITION_BOOLEAN_REFPtr[a_datatype_definition_boolean_refDB.ID]
	if a_datatype_definition_boolean_refOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_datatype_definition_boolean_refOld, a_datatype_definition_boolean_refNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_datatype_definition_boolean_refDB
	c.JSON(http.StatusOK, a_datatype_definition_boolean_refDB)
}

// DeleteA_DATATYPE_DEFINITION_BOOLEAN_REF
//
// swagger:route DELETE /a_datatype_definition_boolean_refs/{ID} a_datatype_definition_boolean_refs deleteA_DATATYPE_DEFINITION_BOOLEAN_REF
//
// # Delete a a_datatype_definition_boolean_ref
//
// default: genericError
//
//	200: a_datatype_definition_boolean_refDBResponse
func (controller *Controller) DeleteA_DATATYPE_DEFINITION_BOOLEAN_REF(c *gin.Context) {

	mutexA_DATATYPE_DEFINITION_BOOLEAN_REF.Lock()
	defer mutexA_DATATYPE_DEFINITION_BOOLEAN_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_DATATYPE_DEFINITION_BOOLEAN_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.GetDB()

	// Get model if exist
	var a_datatype_definition_boolean_refDB orm.A_DATATYPE_DEFINITION_BOOLEAN_REFDB
	if _, err := db.First(&a_datatype_definition_boolean_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_datatype_definition_boolean_refDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_datatype_definition_boolean_refDeleted := new(models.A_DATATYPE_DEFINITION_BOOLEAN_REF)
	a_datatype_definition_boolean_refDB.CopyBasicFieldsToA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_refDeleted)

	// get stage instance from DB instance, and call callback function
	a_datatype_definition_boolean_refStaged := backRepo.BackRepoA_DATATYPE_DEFINITION_BOOLEAN_REF.Map_A_DATATYPE_DEFINITION_BOOLEAN_REFDBID_A_DATATYPE_DEFINITION_BOOLEAN_REFPtr[a_datatype_definition_boolean_refDB.ID]
	if a_datatype_definition_boolean_refStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_datatype_definition_boolean_refStaged, a_datatype_definition_boolean_refDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
