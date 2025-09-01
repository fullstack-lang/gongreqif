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
var __A_SPECIFICATIONS__dummysDeclaration__ models.A_SPECIFICATIONS
var __A_SPECIFICATIONS_time__dummyDeclaration time.Duration

var mutexA_SPECIFICATIONS sync.Mutex

// An A_SPECIFICATIONSID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_SPECIFICATIONS updateA_SPECIFICATIONS deleteA_SPECIFICATIONS
type A_SPECIFICATIONSID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_SPECIFICATIONSInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_SPECIFICATIONS updateA_SPECIFICATIONS
type A_SPECIFICATIONSInput struct {
	// The A_SPECIFICATIONS to submit or modify
	// in: body
	A_SPECIFICATIONS *orm.A_SPECIFICATIONSAPI
}

// GetA_SPECIFICATIONSs
//
// swagger:route GET /a_specificationss a_specificationss getA_SPECIFICATIONSs
//
// # Get all a_specificationss
//
// Responses:
// default: genericError
//
//	200: a_specificationsDBResponse
func (controller *Controller) GetA_SPECIFICATIONSs(c *gin.Context) {

	// source slice
	var a_specificationsDBs []orm.A_SPECIFICATIONSDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPECIFICATIONSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPECIFICATIONS.GetDB()

	_, err := db.Find(&a_specificationsDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_specificationsAPIs := make([]orm.A_SPECIFICATIONSAPI, 0)

	// for each a_specifications, update fields from the database nullable fields
	for idx := range a_specificationsDBs {
		a_specificationsDB := &a_specificationsDBs[idx]
		_ = a_specificationsDB
		var a_specificationsAPI orm.A_SPECIFICATIONSAPI

		// insertion point for updating fields
		a_specificationsAPI.ID = a_specificationsDB.ID
		a_specificationsDB.CopyBasicFieldsToA_SPECIFICATIONS_WOP(&a_specificationsAPI.A_SPECIFICATIONS_WOP)
		a_specificationsAPI.A_SPECIFICATIONSPointersEncoding = a_specificationsDB.A_SPECIFICATIONSPointersEncoding
		a_specificationsAPIs = append(a_specificationsAPIs, a_specificationsAPI)
	}

	c.JSON(http.StatusOK, a_specificationsAPIs)
}

// PostA_SPECIFICATIONS
//
// swagger:route POST /a_specificationss a_specificationss postA_SPECIFICATIONS
//
// Creates a a_specifications
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_SPECIFICATIONS(c *gin.Context) {

	mutexA_SPECIFICATIONS.Lock()
	defer mutexA_SPECIFICATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_SPECIFICATIONSs", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPECIFICATIONS.GetDB()

	// Validate input
	var input orm.A_SPECIFICATIONSAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_specifications
	a_specificationsDB := orm.A_SPECIFICATIONSDB{}
	a_specificationsDB.A_SPECIFICATIONSPointersEncoding = input.A_SPECIFICATIONSPointersEncoding
	a_specificationsDB.CopyBasicFieldsFromA_SPECIFICATIONS_WOP(&input.A_SPECIFICATIONS_WOP)

	_, err = db.Create(&a_specificationsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_SPECIFICATIONS.CheckoutPhaseOneInstance(&a_specificationsDB)
	a_specifications := backRepo.BackRepoA_SPECIFICATIONS.Map_A_SPECIFICATIONSDBID_A_SPECIFICATIONSPtr[a_specificationsDB.ID]

	if a_specifications != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_specifications)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_specificationsDB)
}

// GetA_SPECIFICATIONS
//
// swagger:route GET /a_specificationss/{ID} a_specificationss getA_SPECIFICATIONS
//
// Gets the details for a a_specifications.
//
// Responses:
// default: genericError
//
//	200: a_specificationsDBResponse
func (controller *Controller) GetA_SPECIFICATIONS(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_SPECIFICATIONS", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPECIFICATIONS.GetDB()

	// Get a_specificationsDB in DB
	var a_specificationsDB orm.A_SPECIFICATIONSDB
	if _, err := db.First(&a_specificationsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_specificationsAPI orm.A_SPECIFICATIONSAPI
	a_specificationsAPI.ID = a_specificationsDB.ID
	a_specificationsAPI.A_SPECIFICATIONSPointersEncoding = a_specificationsDB.A_SPECIFICATIONSPointersEncoding
	a_specificationsDB.CopyBasicFieldsToA_SPECIFICATIONS_WOP(&a_specificationsAPI.A_SPECIFICATIONS_WOP)

	c.JSON(http.StatusOK, a_specificationsAPI)
}

// UpdateA_SPECIFICATIONS
//
// swagger:route PATCH /a_specificationss/{ID} a_specificationss updateA_SPECIFICATIONS
//
// # Update a a_specifications
//
// Responses:
// default: genericError
//
//	200: a_specificationsDBResponse
func (controller *Controller) UpdateA_SPECIFICATIONS(c *gin.Context) {

	mutexA_SPECIFICATIONS.Lock()
	defer mutexA_SPECIFICATIONS.Unlock()

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
	db := backRepo.BackRepoA_SPECIFICATIONS.GetDB()

	// Validate input
	var input orm.A_SPECIFICATIONSAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_specificationsDB orm.A_SPECIFICATIONSDB

	// fetch the a_specifications
	_, err := db.First(&a_specificationsDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_specificationsDB.CopyBasicFieldsFromA_SPECIFICATIONS_WOP(&input.A_SPECIFICATIONS_WOP)
	a_specificationsDB.A_SPECIFICATIONSPointersEncoding = input.A_SPECIFICATIONSPointersEncoding

	db, _ = db.Model(&a_specificationsDB)
	_, err = db.Updates(&a_specificationsDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_specificationsNew := new(models.A_SPECIFICATIONS)
	a_specificationsDB.CopyBasicFieldsToA_SPECIFICATIONS(a_specificationsNew)

	// redeem pointers
	a_specificationsDB.DecodePointers(backRepo, a_specificationsNew)

	// get stage instance from DB instance, and call callback function
	a_specificationsOld := backRepo.BackRepoA_SPECIFICATIONS.Map_A_SPECIFICATIONSDBID_A_SPECIFICATIONSPtr[a_specificationsDB.ID]
	if a_specificationsOld != nil {
		if !hasMouseEvent {
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_specificationsOld, a_specificationsNew, nil)
		} else {
			mouseEvent := &models.Gong__MouseEvent{
				ShiftKey: shiftKey,
			}
			models.OnAfterUpdateFromFront(backRepo.GetStage(), a_specificationsOld, a_specificationsNew, mouseEvent)

		}
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_specificationsDB
	c.JSON(http.StatusOK, a_specificationsDB)
}

// DeleteA_SPECIFICATIONS
//
// swagger:route DELETE /a_specificationss/{ID} a_specificationss deleteA_SPECIFICATIONS
//
// # Delete a a_specifications
//
// default: genericError
//
//	200: a_specificationsDBResponse
func (controller *Controller) DeleteA_SPECIFICATIONS(c *gin.Context) {

	mutexA_SPECIFICATIONS.Lock()
	defer mutexA_SPECIFICATIONS.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_SPECIFICATIONS", "Name", stackPath)
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
	db := backRepo.BackRepoA_SPECIFICATIONS.GetDB()

	// Get model if exist
	var a_specificationsDB orm.A_SPECIFICATIONSDB
	if _, err := db.First(&a_specificationsDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_specificationsDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_specificationsDeleted := new(models.A_SPECIFICATIONS)
	a_specificationsDB.CopyBasicFieldsToA_SPECIFICATIONS(a_specificationsDeleted)

	// get stage instance from DB instance, and call callback function
	a_specificationsStaged := backRepo.BackRepoA_SPECIFICATIONS.Map_A_SPECIFICATIONSDBID_A_SPECIFICATIONSPtr[a_specificationsDB.ID]
	if a_specificationsStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_specificationsStaged, a_specificationsDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
