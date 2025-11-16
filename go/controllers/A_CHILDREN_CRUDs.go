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
var __A_CHILDREN__dummysDeclaration__ models.A_CHILDREN
var __A_CHILDREN_time__dummyDeclaration time.Duration

var mutexA_CHILDREN sync.Mutex

// An A_CHILDRENID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getA_CHILDREN updateA_CHILDREN deleteA_CHILDREN
type A_CHILDRENID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// A_CHILDRENInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postA_CHILDREN updateA_CHILDREN
type A_CHILDRENInput struct {
	// The A_CHILDREN to submit or modify
	// in: body
	A_CHILDREN *orm.A_CHILDRENAPI
}

// GetA_CHILDRENs
//
// swagger:route GET /a_childrens a_childrens getA_CHILDRENs
//
// # Get all a_childrens
//
// Responses:
// default: genericError
//
//	200: a_childrenDBResponse
func (controller *Controller) GetA_CHILDRENs(c *gin.Context) {

	// source slice
	var a_childrenDBs []orm.A_CHILDRENDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_CHILDRENs", "Name", stackPath)
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
	db := backRepo.BackRepoA_CHILDREN.GetDB()

	_, err := db.Find(&a_childrenDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	a_childrenAPIs := make([]orm.A_CHILDRENAPI, 0)

	// for each a_children, update fields from the database nullable fields
	for idx := range a_childrenDBs {
		a_childrenDB := &a_childrenDBs[idx]
		_ = a_childrenDB
		var a_childrenAPI orm.A_CHILDRENAPI

		// insertion point for updating fields
		a_childrenAPI.ID = a_childrenDB.ID
		a_childrenDB.CopyBasicFieldsToA_CHILDREN_WOP(&a_childrenAPI.A_CHILDREN_WOP)
		a_childrenAPI.A_CHILDRENPointersEncoding = a_childrenDB.A_CHILDRENPointersEncoding
		a_childrenAPIs = append(a_childrenAPIs, a_childrenAPI)
	}

	c.JSON(http.StatusOK, a_childrenAPIs)
}

// PostA_CHILDREN
//
// swagger:route POST /a_childrens a_childrens postA_CHILDREN
//
// Creates a a_children
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostA_CHILDREN(c *gin.Context) {

	mutexA_CHILDREN.Lock()
	defer mutexA_CHILDREN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostA_CHILDRENs", "Name", stackPath)
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
	db := backRepo.BackRepoA_CHILDREN.GetDB()

	// Validate input
	var input orm.A_CHILDRENAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create a_children
	a_childrenDB := orm.A_CHILDRENDB{}
	a_childrenDB.A_CHILDRENPointersEncoding = input.A_CHILDRENPointersEncoding
	a_childrenDB.CopyBasicFieldsFromA_CHILDREN_WOP(&input.A_CHILDREN_WOP)

	_, err = db.Create(&a_childrenDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoA_CHILDREN.CheckoutPhaseOneInstance(&a_childrenDB)
	a_children := backRepo.BackRepoA_CHILDREN.Map_A_CHILDRENDBID_A_CHILDRENPtr[a_childrenDB.ID]

	if a_children != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), a_children)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, a_childrenDB)
}

// GetA_CHILDREN
//
// swagger:route GET /a_childrens/{ID} a_childrens getA_CHILDREN
//
// Gets the details for a a_children.
//
// Responses:
// default: genericError
//
//	200: a_childrenDBResponse
func (controller *Controller) GetA_CHILDREN(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetA_CHILDREN", "Name", stackPath)
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
	db := backRepo.BackRepoA_CHILDREN.GetDB()

	// Get a_childrenDB in DB
	var a_childrenDB orm.A_CHILDRENDB
	if _, err := db.First(&a_childrenDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var a_childrenAPI orm.A_CHILDRENAPI
	a_childrenAPI.ID = a_childrenDB.ID
	a_childrenAPI.A_CHILDRENPointersEncoding = a_childrenDB.A_CHILDRENPointersEncoding
	a_childrenDB.CopyBasicFieldsToA_CHILDREN_WOP(&a_childrenAPI.A_CHILDREN_WOP)

	c.JSON(http.StatusOK, a_childrenAPI)
}

// UpdateA_CHILDREN
//
// swagger:route PATCH /a_childrens/{ID} a_childrens updateA_CHILDREN
//
// # Update a a_children
//
// Responses:
// default: genericError
//
//	200: a_childrenDBResponse
func (controller *Controller) UpdateA_CHILDREN(c *gin.Context) {

	mutexA_CHILDREN.Lock()
	defer mutexA_CHILDREN.Unlock()

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
	db := backRepo.BackRepoA_CHILDREN.GetDB()

	// Validate input
	var input orm.A_CHILDRENAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var a_childrenDB orm.A_CHILDRENDB

	// fetch the a_children
	_, err := db.First(&a_childrenDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	a_childrenDB.CopyBasicFieldsFromA_CHILDREN_WOP(&input.A_CHILDREN_WOP)
	a_childrenDB.A_CHILDRENPointersEncoding = input.A_CHILDRENPointersEncoding

	db, _ = db.Model(&a_childrenDB)
	_, err = db.Updates(&a_childrenDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	a_childrenNew := new(models.A_CHILDREN)
	a_childrenDB.CopyBasicFieldsToA_CHILDREN(a_childrenNew)

	// redeem pointers
	a_childrenDB.DecodePointers(backRepo, a_childrenNew)

	// get stage instance from DB instance, and call callback function
	a_childrenOld := backRepo.BackRepoA_CHILDREN.Map_A_CHILDRENDBID_A_CHILDRENPtr[a_childrenDB.ID]
	if a_childrenOld != nil {
		models.OnAfterUpdateFromFront(backRepo.GetStage(), a_childrenOld, a_childrenNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the a_childrenDB
	c.JSON(http.StatusOK, a_childrenDB)
}

// DeleteA_CHILDREN
//
// swagger:route DELETE /a_childrens/{ID} a_childrens deleteA_CHILDREN
//
// # Delete a a_children
//
// default: genericError
//
//	200: a_childrenDBResponse
func (controller *Controller) DeleteA_CHILDREN(c *gin.Context) {

	mutexA_CHILDREN.Lock()
	defer mutexA_CHILDREN.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteA_CHILDREN", "Name", stackPath)
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
	db := backRepo.BackRepoA_CHILDREN.GetDB()

	// Get model if exist
	var a_childrenDB orm.A_CHILDRENDB
	if _, err := db.First(&a_childrenDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&a_childrenDB)

	// get an instance (not staged) from DB instance, and call callback function
	a_childrenDeleted := new(models.A_CHILDREN)
	a_childrenDB.CopyBasicFieldsToA_CHILDREN(a_childrenDeleted)

	// get stage instance from DB instance, and call callback function
	a_childrenStaged := backRepo.BackRepoA_CHILDREN.Map_A_CHILDRENDBID_A_CHILDRENPtr[a_childrenDB.ID]
	if a_childrenStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), a_childrenStaged, a_childrenDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
