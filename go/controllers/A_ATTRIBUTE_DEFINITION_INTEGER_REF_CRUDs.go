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
var __A_ATTRIBUTE_DEFINITION_INTEGER_REF__dummysDeclaration__ models.A_ATTRIBUTE_DEFINITION_INTEGER_REF
var __A_ATTRIBUTE_DEFINITION_INTEGER_REF_time__dummyDeclaration time.Duration

var mutexA_ATTRIBUTE_DEFINITION_INTEGER_REF sync.Mutex

// An A_ATTRIBUTE_DEFINITION_INTEGER_REFID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_ATTRIBUTE_DEFINITION_INTEGER_REF updateA_ATTRIBUTE_DEFINITION_INTEGER_REF deleteA_ATTRIBUTE_DEFINITION_INTEGER_REF
type A_ATTRIBUTE_DEFINITION_INTEGER_REFID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_ATTRIBUTE_DEFINITION_INTEGER_REFInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_ATTRIBUTE_DEFINITION_INTEGER_REF updateA_ATTRIBUTE_DEFINITION_INTEGER_REF
type A_ATTRIBUTE_DEFINITION_INTEGER_REFInput struct {
	// The A_ATTRIBUTE_DEFINITION_INTEGER_REF to submit or modify
	// in: body
	A_ATTRIBUTE_DEFINITION_INTEGER_REF *orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFAPI
}

// GetA_ATTRIBUTE_DEFINITION_INTEGER_REFs
//
// swagger:route GET /a_attribute_definition_integer_refs a_attribute_definition_integer_refs getA_ATTRIBUTE_DEFINITION_INTEGER_REFs
//
// # Get all a_attribute_definition_integer_refs
//
// Responses:
// default: genericError
//
//	200: a_attribute_definition_integer_refDBResponse
func (controller *Controller) GetA_ATTRIBUTE_DEFINITION_INTEGER_REFs(c *gin.Context) {

	// source slice
	var a_attribute_definition_integer_refDBs []orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_DEFINITION_INTEGER_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.GetDB()

	_, err := db.Find(&a_attribute_definition_integer_refDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_attribute_definition_integer_refAPIs := make([]orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFAPI, 0)

	// for each a_attribute_definition_integer_ref, update fields from the database nullable fields
	for idx := range a_attribute_definition_integer_refDBs {
		a_attribute_definition_integer_refDB := &a_attribute_definition_integer_refDBs[idx]
		_ = a_attribute_definition_integer_refDB
		var a_attribute_definition_integer_refAPI orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFAPI

		// insertion point for updating fields
		a_attribute_definition_integer_refAPI.ID = a_attribute_definition_integer_refDB.ID
		a_attribute_definition_integer_refDB.CopyBasicFieldsToA_ATTRIBUTE_DEFINITION_INTEGER_REF_WOP(&a_attribute_definition_integer_refAPI.A_ATTRIBUTE_DEFINITION_INTEGER_REF_WOP)
		a_attribute_definition_integer_refAPI.A_ATTRIBUTE_DEFINITION_INTEGER_REFPointersEncoding = a_attribute_definition_integer_refDB.A_ATTRIBUTE_DEFINITION_INTEGER_REFPointersEncoding
		a_attribute_definition_integer_refAPIs = append(a_attribute_definition_integer_refAPIs, a_attribute_definition_integer_refAPI)
	}

	c.JSON(http.StatusOK, a_attribute_definition_integer_refAPIs)
}

// PostA_ATTRIBUTE_DEFINITION_INTEGER_REF
//
// swagger:route POST /a_attribute_definition_integer_refs a_attribute_definition_integer_refs postA_ATTRIBUTE_DEFINITION_INTEGER_REF
//
// Creates a a_attribute_definition_integer_ref
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_ATTRIBUTE_DEFINITION_INTEGER_REF(c *gin.Context) {

	mutexA_ATTRIBUTE_DEFINITION_INTEGER_REF.Lock()
	defer mutexA_ATTRIBUTE_DEFINITION_INTEGER_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_ATTRIBUTE_DEFINITION_INTEGER_REFs", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_attribute_definition_integer_ref
	a_attribute_definition_integer_refDB := orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFDB{}
	a_attribute_definition_integer_refDB.A_ATTRIBUTE_DEFINITION_INTEGER_REFPointersEncoding = input.A_ATTRIBUTE_DEFINITION_INTEGER_REFPointersEncoding
	a_attribute_definition_integer_refDB.CopyBasicFieldsFromA_ATTRIBUTE_DEFINITION_INTEGER_REF_WOP(&input.A_ATTRIBUTE_DEFINITION_INTEGER_REF_WOP)

	_, err = db.Create(&a_attribute_definition_integer_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.CheckoutPhaseOneInstance(&a_attribute_definition_integer_refDB)
	a_attribute_definition_integer_ref := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.Map_A_ATTRIBUTE_DEFINITION_INTEGER_REFDBID_A_ATTRIBUTE_DEFINITION_INTEGER_REFPtr[a_attribute_definition_integer_refDB.ID]

	if a_attribute_definition_integer_ref != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_attribute_definition_integer_ref)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_attribute_definition_integer_refDB)
}

// GetA_ATTRIBUTE_DEFINITION_INTEGER_REF
//
// swagger:route GET /a_attribute_definition_integer_refs/{ID} a_attribute_definition_integer_refs getA_ATTRIBUTE_DEFINITION_INTEGER_REF
//
// Gets the details for a a_attribute_definition_integer_ref.
//
// Responses:
// default: genericError
//
//	200: a_attribute_definition_integer_refDBResponse
func (controller *Controller) GetA_ATTRIBUTE_DEFINITION_INTEGER_REF(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_ATTRIBUTE_DEFINITION_INTEGER_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.GetDB()

	// Get a_attribute_definition_integer_refDB in DB
	var a_attribute_definition_integer_refDB orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFDB
	if _, err := db.First(&a_attribute_definition_integer_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_attribute_definition_integer_refAPI orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFAPI
	a_attribute_definition_integer_refAPI.ID = a_attribute_definition_integer_refDB.ID
	a_attribute_definition_integer_refAPI.A_ATTRIBUTE_DEFINITION_INTEGER_REFPointersEncoding = a_attribute_definition_integer_refDB.A_ATTRIBUTE_DEFINITION_INTEGER_REFPointersEncoding
	a_attribute_definition_integer_refDB.CopyBasicFieldsToA_ATTRIBUTE_DEFINITION_INTEGER_REF_WOP(&a_attribute_definition_integer_refAPI.A_ATTRIBUTE_DEFINITION_INTEGER_REF_WOP)

	c.JSON(http.StatusOK, a_attribute_definition_integer_refAPI)
}

// UpdateA_ATTRIBUTE_DEFINITION_INTEGER_REF
//
// swagger:route PATCH /a_attribute_definition_integer_refs/{ID} a_attribute_definition_integer_refs updateA_ATTRIBUTE_DEFINITION_INTEGER_REF
//
// # Update a a_attribute_definition_integer_ref
//
// Responses:
// default: genericError
//
//	200: a_attribute_definition_integer_refDBResponse
func (controller *Controller) UpdateA_ATTRIBUTE_DEFINITION_INTEGER_REF(c *gin.Context) {

	mutexA_ATTRIBUTE_DEFINITION_INTEGER_REF.Lock()
	defer mutexA_ATTRIBUTE_DEFINITION_INTEGER_REF.Unlock()

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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.GetDB()

	// Validate input
	var input orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_attribute_definition_integer_refDB orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFDB

	// fetch the a_attribute_definition_integer_ref
	_, err := db.First(&a_attribute_definition_integer_refDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_attribute_definition_integer_refDB.CopyBasicFieldsFromA_ATTRIBUTE_DEFINITION_INTEGER_REF_WOP(&input.A_ATTRIBUTE_DEFINITION_INTEGER_REF_WOP)
	a_attribute_definition_integer_refDB.A_ATTRIBUTE_DEFINITION_INTEGER_REFPointersEncoding = input.A_ATTRIBUTE_DEFINITION_INTEGER_REFPointersEncoding

	db, _ = db.Model(&a_attribute_definition_integer_refDB)
	_, err = db.Updates(&a_attribute_definition_integer_refDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_definition_integer_refNew := new(models.A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	a_attribute_definition_integer_refDB.CopyBasicFieldsToA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_refNew)

	// redeem pointers
	a_attribute_definition_integer_refDB.DecodePointers(backRepo, a_attribute_definition_integer_refNew)

	// get stage instance from DB instance, and call callback function
	a_attribute_definition_integer_refOld := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.Map_A_ATTRIBUTE_DEFINITION_INTEGER_REFDBID_A_ATTRIBUTE_DEFINITION_INTEGER_REFPtr[a_attribute_definition_integer_refDB.ID]
	if a_attribute_definition_integer_refOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_attribute_definition_integer_refOld, a_attribute_definition_integer_refNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_attribute_definition_integer_refDB
	c.JSON(http.StatusOK, a_attribute_definition_integer_refDB)
}

// DeleteA_ATTRIBUTE_DEFINITION_INTEGER_REF
//
// swagger:route DELETE /a_attribute_definition_integer_refs/{ID} a_attribute_definition_integer_refs deleteA_ATTRIBUTE_DEFINITION_INTEGER_REF
//
// # Delete a a_attribute_definition_integer_ref
//
// default: genericError
//
//	200: a_attribute_definition_integer_refDBResponse
func (controller *Controller) DeleteA_ATTRIBUTE_DEFINITION_INTEGER_REF(c *gin.Context) {

	mutexA_ATTRIBUTE_DEFINITION_INTEGER_REF.Lock()
	defer mutexA_ATTRIBUTE_DEFINITION_INTEGER_REF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_ATTRIBUTE_DEFINITION_INTEGER_REF", "Name", stackPath)
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
	db := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.GetDB()

	// Get model if exist
	var a_attribute_definition_integer_refDB orm.A_ATTRIBUTE_DEFINITION_INTEGER_REFDB
	if _, err := db.First(&a_attribute_definition_integer_refDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_attribute_definition_integer_refDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_attribute_definition_integer_refDeleted := new(models.A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	a_attribute_definition_integer_refDB.CopyBasicFieldsToA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_refDeleted)

	// get stage instance from DB instance, and call callback function
	a_attribute_definition_integer_refStaged := backRepo.BackRepoA_ATTRIBUTE_DEFINITION_INTEGER_REF.Map_A_ATTRIBUTE_DEFINITION_INTEGER_REFDBID_A_ATTRIBUTE_DEFINITION_INTEGER_REFPtr[a_attribute_definition_integer_refDB.ID]
	if a_attribute_definition_integer_refStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_attribute_definition_integer_refStaged, a_attribute_definition_integer_refDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
