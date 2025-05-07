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
var __A_THE_HEADER__dummysDeclaration__ models.A_THE_HEADER
var __A_THE_HEADER_time__dummyDeclaration time.Duration

var mutexA_THE_HEADER sync.Mutex

// An A_THE_HEADERID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_THE_HEADER updateA_THE_HEADER deleteA_THE_HEADER
type A_THE_HEADERID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_THE_HEADERInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_THE_HEADER updateA_THE_HEADER
type A_THE_HEADERInput struct {
	// The A_THE_HEADER to submit or modify
	// in: body
	A_THE_HEADER *orm.A_THE_HEADERAPI
}

// GetA_THE_HEADERs
//
// swagger:route GET /a_the_headers a_the_headers getA_THE_HEADERs
//
// # Get all a_the_headers
//
// Responses:
// default: genericError
//
//	200: a_the_headerDBResponse
func (controller *Controller) GetA_THE_HEADERs(c *gin.Context) {

	// source slice
	var a_the_headerDBs []orm.A_THE_HEADERDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_THE_HEADERs", "Name", stackPath)
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
	db := backRepo.BackRepoA_THE_HEADER.GetDB()

	_, err := db.Find(&a_the_headerDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_the_headerAPIs := make([]orm.A_THE_HEADERAPI, 0)

	// for each a_the_header, update fields from the database nullable fields
	for idx := range a_the_headerDBs {
		a_the_headerDB := &a_the_headerDBs[idx]
		_ = a_the_headerDB
		var a_the_headerAPI orm.A_THE_HEADERAPI

		// insertion point for updating fields
		a_the_headerAPI.ID = a_the_headerDB.ID
		a_the_headerDB.CopyBasicFieldsToA_THE_HEADER_WOP(&a_the_headerAPI.A_THE_HEADER_WOP)
		a_the_headerAPI.A_THE_HEADERPointersEncoding = a_the_headerDB.A_THE_HEADERPointersEncoding
		a_the_headerAPIs = append(a_the_headerAPIs, a_the_headerAPI)
	}

	c.JSON(http.StatusOK, a_the_headerAPIs)
}

// PostA_THE_HEADER
//
// swagger:route POST /a_the_headers a_the_headers postA_THE_HEADER
//
// Creates a a_the_header
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_THE_HEADER(c *gin.Context) {

	mutexA_THE_HEADER.Lock()
	defer mutexA_THE_HEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_THE_HEADERs", "Name", stackPath)
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
	db := backRepo.BackRepoA_THE_HEADER.GetDB()

	// Validate input
	var input orm.A_THE_HEADERAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_the_header
	a_the_headerDB := orm.A_THE_HEADERDB{}
	a_the_headerDB.A_THE_HEADERPointersEncoding = input.A_THE_HEADERPointersEncoding
	a_the_headerDB.CopyBasicFieldsFromA_THE_HEADER_WOP(&input.A_THE_HEADER_WOP)

	_, err = db.Create(&a_the_headerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_THE_HEADER.CheckoutPhaseOneInstance(&a_the_headerDB)
	a_the_header := backRepo.BackRepoA_THE_HEADER.Map_A_THE_HEADERDBID_A_THE_HEADERPtr[a_the_headerDB.ID]

	if a_the_header != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_the_header)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_the_headerDB)
}

// GetA_THE_HEADER
//
// swagger:route GET /a_the_headers/{ID} a_the_headers getA_THE_HEADER
//
// Gets the details for a a_the_header.
//
// Responses:
// default: genericError
//
//	200: a_the_headerDBResponse
func (controller *Controller) GetA_THE_HEADER(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_THE_HEADER", "Name", stackPath)
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
	db := backRepo.BackRepoA_THE_HEADER.GetDB()

	// Get a_the_headerDB in DB
	var a_the_headerDB orm.A_THE_HEADERDB
	if _, err := db.First(&a_the_headerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_the_headerAPI orm.A_THE_HEADERAPI
	a_the_headerAPI.ID = a_the_headerDB.ID
	a_the_headerAPI.A_THE_HEADERPointersEncoding = a_the_headerDB.A_THE_HEADERPointersEncoding
	a_the_headerDB.CopyBasicFieldsToA_THE_HEADER_WOP(&a_the_headerAPI.A_THE_HEADER_WOP)

	c.JSON(http.StatusOK, a_the_headerAPI)
}

// UpdateA_THE_HEADER
//
// swagger:route PATCH /a_the_headers/{ID} a_the_headers updateA_THE_HEADER
//
// # Update a a_the_header
//
// Responses:
// default: genericError
//
//	200: a_the_headerDBResponse
func (controller *Controller) UpdateA_THE_HEADER(c *gin.Context) {

	mutexA_THE_HEADER.Lock()
	defer mutexA_THE_HEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateA_THE_HEADER", "Name", stackPath)
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
	db := backRepo.BackRepoA_THE_HEADER.GetDB()

	// Validate input
	var input orm.A_THE_HEADERAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_the_headerDB orm.A_THE_HEADERDB

	// fetch the a_the_header
	_, err := db.First(&a_the_headerDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_the_headerDB.CopyBasicFieldsFromA_THE_HEADER_WOP(&input.A_THE_HEADER_WOP)
	a_the_headerDB.A_THE_HEADERPointersEncoding = input.A_THE_HEADERPointersEncoding

	db, _ = db.Model(&a_the_headerDB)
	_, err = db.Updates(&a_the_headerDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_the_headerNew := new(models.A_THE_HEADER)
	a_the_headerDB.CopyBasicFieldsToA_THE_HEADER(a_the_headerNew)

	// redeem pointers
	a_the_headerDB.DecodePointers(backRepo, a_the_headerNew)

	// get stage instance from DB instance, and call callback function
	a_the_headerOld := backRepo.BackRepoA_THE_HEADER.Map_A_THE_HEADERDBID_A_THE_HEADERPtr[a_the_headerDB.ID]
	if a_the_headerOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), a_the_headerOld, a_the_headerNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_the_headerDB
	c.JSON(http.StatusOK, a_the_headerDB)
}

// DeleteA_THE_HEADER
//
// swagger:route DELETE /a_the_headers/{ID} a_the_headers deleteA_THE_HEADER
//
// # Delete a a_the_header
//
// default: genericError
//
//	200: a_the_headerDBResponse
func (controller *Controller) DeleteA_THE_HEADER(c *gin.Context) {

	mutexA_THE_HEADER.Lock()
	defer mutexA_THE_HEADER.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_THE_HEADER", "Name", stackPath)
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
	db := backRepo.BackRepoA_THE_HEADER.GetDB()

	// Get model if exist
	var a_the_headerDB orm.A_THE_HEADERDB
	if _, err := db.First(&a_the_headerDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_the_headerDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_the_headerDeleted := new(models.A_THE_HEADER)
	a_the_headerDB.CopyBasicFieldsToA_THE_HEADER(a_the_headerDeleted)

	// get stage instance from DB instance, and call callback function
	a_the_headerStaged := backRepo.BackRepoA_THE_HEADER.Map_A_THE_HEADERDBID_A_THE_HEADERPtr[a_the_headerDB.ID]
	if a_the_headerStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_the_headerStaged, a_the_headerDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
