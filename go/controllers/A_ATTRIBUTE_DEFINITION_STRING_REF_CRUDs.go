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
var __A_ATTRIBUTE_DEFINITION_STRING_REF__dummysDeclaration__ models.A_ATTRIBUTE_DEFINITION_STRING_REF
var __A_ATTRIBUTE_DEFINITION_STRING_REF_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_DEFINITION_STRING_REF sync.Mutex

// An A_ATTRIBUTE_DEFINITION_STRING_REFID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_DEFINITION_STRING_REF updateA_ATTRIBUTE_DEFINITION_STRING_REF deleteA_ATTRIBUTE_DEFINITION_STRING_REF
type A_ATTRIBUTE_DEFINITION_STRING_REFID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_DEFINITION_STRING_REFInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_DEFINITION_STRING_REF updateA_ATTRIBUTE_DEFINITION_STRING_REF
type A_ATTRIBUTE_DEFINITION_STRING_REFInput struct {
	// The A_ATTRIBUTE_DEFINITION_STRING_REF to submit or modify
	// in: body
	A_ATTRIBUTE_DEFINITION_STRING_REF *orm.A_ATTRIBUTE_DEFINITION_STRING_REFAPI
}

// GetA_ATTRIBUTE_DEFINITION_STRING_REFs
//
// swagger:route GET /a_attribute_definition_string_refs a_attribute_definition_string_refs getA_ATTRIBUTE_DEFINITION_STRING_REFs
//
// # Get all a_attribute_definition_string_refs
//
// Responses:
// default: genericError
//
//	200: a_attribute_definition_string_refDBResponse
func (controller *Controller) GetA_ATTRIBUTE_DEFINITION_STRING_REFs(c *gin.Context) {

	// source slice
	var a_attribute_definition_string_refDBs []orm.A_ATTRIBUTE_DEFINITION_STRING_REFDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_DEFINITION_STRING_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.GetDB()

	_, err := db.Find(&a_attribute_definition_string_refDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_definition_string_refAPIs := make([]orm.A_ATTRIBUTE_DEFINITION_STRING_REFAPI, 0)

	// for each a_attribute_definition_string_ref, update fields from the database nullable fields
	for idx := range a_attribute_definition_string_refDBs {
		a_attribute_definition_string_refDB := &a_attribute_definition_string_refDBs[idx]
		_ = a_attribute_definition_string_refDB
		var a_attribute_definition_string_refAPI orm.A_ATTRIBUTE_DEFINITION_STRING_REFAPI

		// insertion point for updating fields
		a_attribute_definition_string_refAPI.ID = a_attribute_definition_string_refDB.ID
		a_attribute_definition_string_refDB.CopyBasicFieldsToA_ATTRIBUTE_DEFINITION_STRING_REF_WOP(&a_attribute_definition_string_refAPI.A_ATTRIBUTE_DEFINITION_STRING_REF_WOP)
		a_attribute_definition_string_refAPI.A_ATTRIBUTE_DEFINITION_STRING_REFPointersEncoding = a_attribute_definition_string_refDB.A_ATTRIBUTE_DEFINITION_STRING_REFPointersEncoding
		a_attribute_definition_string_refAPIs = append(a_attribute_definition_string_refAPIs, a_attribute_definition_string_refAPI)
	}

	c.JSON(http.StatusOK, a_attribute_definition_string_refAPIs)
}

// PostA_ATTRIBUTE_DEFINITION_STRING_REF
//
// swagger:route POST /a_attribute_definition_string_refs a_attribute_definition_string_refs postA_ATTRIBUTE_DEFINITION_STRING_REF
//
// Creates a a_attribute_definition_string_ref
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_DEFINITION_STRING_REF(c *gin.Context) {

	mutexA_ATTRIBUTE_DEFINITION_STRING_REF.Lock()
	defer mutexA_ATTRIBUTE_DEFINITION_STRING_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_DEFINITION_STRING_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_DEFINITION_STRING_REFAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_definition_string_ref
	a_attribute_definition_string_refDB := orm.A_ATTRIBUTE_DEFINITION_STRING_REFDB{}
	a_attribute_definition_string_refDB.A_ATTRIBUTE_DEFINITION_STRING_REFPointersEncoding = input.A_ATTRIBUTE_DEFINITION_STRING_REFPointersEncoding
	a_attribute_definition_string_refDB.CopyBasicFieldsFromA_ATTRIBUTE_DEFINITION_STRING_REF_WOP(&input.A_ATTRIBUTE_DEFINITION_STRING_REF_WOP)

	_, err = db.Create(&a_attribute_definition_string_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.CheckoutPhaseOneInstance(&a_attribute_definition_string_refDB)
	a_attribute_definition_string_ref := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.Map_A_ATTRIBUTE_DEFINITION_STRING_REFDBID_A_ATTRIBUTE_DEFINITION_STRING_REFPtr[a_attribute_definition_string_refDB.ID]

	if a_attribute_definition_string_ref != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_definition_string_ref)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_definition_string_refDB)
}

// GetA_ATTRIBUTE_DEFINITION_STRING_REF
//
// swagger:route GET /a_attribute_definition_string_refs/{ID} a_attribute_definition_string_refs getA_ATTRIBUTE_DEFINITION_STRING_REF
//
// Gets the details for a a_attribute_definition_string_ref.
//
// Responses:
// default: genericError
//
//	200: a_attribute_definition_string_refDBResponse
func (controller *Controller) GetA_ATTRIBUTE_DEFINITION_STRING_REF(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_DEFINITION_STRING_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.GetDB()

	// Get a_attribute_definition_string_refDB in DB
	var a_attribute_definition_string_refDB orm.A_ATTRIBUTE_DEFINITION_STRING_REFDB
	if _, err := db.First(&a_attribute_definition_string_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_definition_string_refAPI orm.A_ATTRIBUTE_DEFINITION_STRING_REFAPI
	a_attribute_definition_string_refAPI.ID = a_attribute_definition_string_refDB.ID
	a_attribute_definition_string_refAPI.A_ATTRIBUTE_DEFINITION_STRING_REFPointersEncoding = a_attribute_definition_string_refDB.A_ATTRIBUTE_DEFINITION_STRING_REFPointersEncoding
	a_attribute_definition_string_refDB.CopyBasicFieldsToA_ATTRIBUTE_DEFINITION_STRING_REF_WOP(&a_attribute_definition_string_refAPI.A_ATTRIBUTE_DEFINITION_STRING_REF_WOP)

	c.JSON(http.StatusOK, a_attribute_definition_string_refAPI)
}

// UpdateA_ATTRIBUTE_DEFINITION_STRING_REF
//
// swagger:route PATCH /a_attribute_definition_string_refs/{ID} a_attribute_definition_string_refs updateA_ATTRIBUTE_DEFINITION_STRING_REF
//
// # Update a a_attribute_definition_string_ref
//
// Responses:
// default: genericError
//
//	200: a_attribute_definition_string_refDBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_DEFINITION_STRING_REF(c *gin.Context) {

	mutexA_ATTRIBUTE_DEFINITION_STRING_REF.Lock()
	defer mutexA_ATTRIBUTE_DEFINITION_STRING_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_ATTRIBUTE_DEFINITION_STRING_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_DEFINITION_STRING_REFAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_definition_string_refDB orm.A_ATTRIBUTE_DEFINITION_STRING_REFDB

	// fetch the a_attribute_definition_string_ref
	_, err := db.First(&a_attribute_definition_string_refDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_definition_string_refDB.CopyBasicFieldsFromA_ATTRIBUTE_DEFINITION_STRING_REF_WOP(&input.A_ATTRIBUTE_DEFINITION_STRING_REF_WOP)
	a_attribute_definition_string_refDB.A_ATTRIBUTE_DEFINITION_STRING_REFPointersEncoding = input.A_ATTRIBUTE_DEFINITION_STRING_REFPointersEncoding

	db, _ = db.Model(&a_attribute_definition_string_refDB)
	_, err = db.Updates(&a_attribute_definition_string_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_definition_string_refNew := new(models.A_ATTRIBUTE_DEFINITION_STRING_REF)
	a_attribute_definition_string_refDB.CopyBasicFieldsToA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_refNew)

	// redeem pointers
	a_attribute_definition_string_refDB.DecodePointers(backRepo, a_attribute_definition_string_refNew)

	// get stage instance from DB instance, and call callback function
	a_attribute_definition_string_refOld := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.Map_A_ATTRIBUTE_DEFINITION_STRING_REFDBID_A_ATTRIBUTE_DEFINITION_STRING_REFPtr[a_attribute_definition_string_refDB.ID]
	if a_attribute_definition_string_refOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_attribute_definition_string_refOld, a_attribute_definition_string_refNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_definition_string_refDB
	c.JSON(http.StatusOK, a_attribute_definition_string_refDB)
}

// DeleteA_ATTRIBUTE_DEFINITION_STRING_REF
//
// swagger:route DELETE /a_attribute_definition_string_refs/{ID} a_attribute_definition_string_refs deleteA_ATTRIBUTE_DEFINITION_STRING_REF
//
// # Delete a a_attribute_definition_string_ref
//
// default: genericError
//
//	200: a_attribute_definition_string_refDBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_DEFINITION_STRING_REF(c *gin.Context) {

	mutexA_ATTRIBUTE_DEFINITION_STRING_REF.Lock()
	defer mutexA_ATTRIBUTE_DEFINITION_STRING_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_DEFINITION_STRING_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.GetDB()

	// Get model if exist
	var a_attribute_definition_string_refDB orm.A_ATTRIBUTE_DEFINITION_STRING_REFDB
	if _, err := db.First(&a_attribute_definition_string_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_definition_string_refDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_definition_string_refDeleted := new(models.A_ATTRIBUTE_DEFINITION_STRING_REF)
	a_attribute_definition_string_refDB.CopyBasicFieldsToA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_refDeleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_definition_string_refStaged := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_STRING_REF.Map_A_ATTRIBUTE_DEFINITION_STRING_REFDBID_A_ATTRIBUTE_DEFINITION_STRING_REFPtr[a_attribute_definition_string_refDB.ID]
	if a_attribute_definition_string_refStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_definition_string_refStaged, a_attribute_definition_string_refDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
