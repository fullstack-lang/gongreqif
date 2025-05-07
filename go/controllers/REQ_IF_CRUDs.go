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
var __REQ_IF__dummysDeclaration__ models.REQ_IF
var __REQ_IF_time__dummyDeclaration time.Duration

var mutexREQ_IF sync.Mutex

// An REQ_IFID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getREQ_IF updateREQ_IF deleteREQ_IF
type REQ_IFID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// REQ_IFInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postREQ_IF updateREQ_IF
type REQ_IFInput struct {
	// The REQ_IF to submit or modify
	// in: body
	REQ_IF *orm.REQ_IFAPI
}

// GetREQ_IFs
//
// swagger:route GET /req_ifs req_ifs getREQ_IFs
//
// # Get all req_ifs
//
// Responses:
// default: genericError
//
//	200: req_ifDBResponse
func (controller *Controller) GetREQ_IFs(c *gin.Context) {

	// source slice
	var req_ifDBs []orm.REQ_IFDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQ_IFs", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF.GetDB()

	_, err := db.Find(&req_ifDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	req_ifAPIs := make([]orm.REQ_IFAPI, 0)

	// for each req_if, update fields from the database nullable fields
	for idx := range req_ifDBs {
		req_ifDB := &req_ifDBs[idx]
		_ = req_ifDB
		var req_ifAPI orm.REQ_IFAPI

		// insertion point for updating fields
		req_ifAPI.ID = req_ifDB.ID
		req_ifDB.CopyBasicFieldsToREQ_IF_WOP(&req_ifAPI.REQ_IF_WOP)
		req_ifAPI.REQ_IFPointersEncoding = req_ifDB.REQ_IFPointersEncoding
		req_ifAPIs = append(req_ifAPIs, req_ifAPI)
	}

	c.JSON(http.StatusOK, req_ifAPIs)
}

// PostREQ_IF
//
// swagger:route POST /req_ifs req_ifs postREQ_IF
//
// Creates a req_if
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostREQ_IF(c *gin.Context) {

	mutexREQ_IF.Lock()
	defer mutexREQ_IF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostREQ_IFs", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF.GetDB()

	// Validate input
	var input orm.REQ_IFAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create req_if
	req_ifDB := orm.REQ_IFDB{}
	req_ifDB.REQ_IFPointersEncoding = input.REQ_IFPointersEncoding
	req_ifDB.CopyBasicFieldsFromREQ_IF_WOP(&input.REQ_IF_WOP)

	_, err = db.Create(&req_ifDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoREQ_IF.CheckoutPhaseOneInstance(&req_ifDB)
	req_if := backRepo.BackRepoREQ_IF.Map_REQ_IFDBID_REQ_IFPtr[req_ifDB.ID]

	if req_if != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), req_if)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, req_ifDB)
}

// GetREQ_IF
//
// swagger:route GET /req_ifs/{ID} req_ifs getREQ_IF
//
// Gets the details for a req_if.
//
// Responses:
// default: genericError
//
//	200: req_ifDBResponse
func (controller *Controller) GetREQ_IF(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetREQ_IF", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF.GetDB()

	// Get req_ifDB in DB
	var req_ifDB orm.REQ_IFDB
	if _, err := db.First(&req_ifDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var req_ifAPI orm.REQ_IFAPI
	req_ifAPI.ID = req_ifDB.ID
	req_ifAPI.REQ_IFPointersEncoding = req_ifDB.REQ_IFPointersEncoding
	req_ifDB.CopyBasicFieldsToREQ_IF_WOP(&req_ifAPI.REQ_IF_WOP)

	c.JSON(http.StatusOK, req_ifAPI)
}

// UpdateREQ_IF
//
// swagger:route PATCH /req_ifs/{ID} req_ifs updateREQ_IF
//
// # Update a req_if
//
// Responses:
// default: genericError
//
//	200: req_ifDBResponse
func (controller *Controller) UpdateREQ_IF(c *gin.Context) {

	mutexREQ_IF.Lock()
	defer mutexREQ_IF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateREQ_IF", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF.GetDB()

	// Validate input
	var input orm.REQ_IFAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var req_ifDB orm.REQ_IFDB

	// fetch the req_if
	_, err := db.First(&req_ifDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	req_ifDB.CopyBasicFieldsFromREQ_IF_WOP(&input.REQ_IF_WOP)
	req_ifDB.REQ_IFPointersEncoding = input.REQ_IFPointersEncoding

	db, _ = db.Model(&req_ifDB)
	_, err = db.Updates(&req_ifDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	req_ifNew := new(models.REQ_IF)
	req_ifDB.CopyBasicFieldsToREQ_IF(req_ifNew)

	// redeem pointers
	req_ifDB.DecodePointers(backRepo, req_ifNew)

	// get stage instance from DB instance, and call callback function
	req_ifOld := backRepo.BackRepoREQ_IF.Map_REQ_IFDBID_REQ_IFPtr[req_ifDB.ID]
	if req_ifOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), req_ifOld, req_ifNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the req_ifDB
	c.JSON(http.StatusOK, req_ifDB)
}

// DeleteREQ_IF
//
// swagger:route DELETE /req_ifs/{ID} req_ifs deleteREQ_IF
//
// # Delete a req_if
//
// default: genericError
//
//	200: req_ifDBResponse
func (controller *Controller) DeleteREQ_IF(c *gin.Context) {

	mutexREQ_IF.Lock()
	defer mutexREQ_IF.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteREQ_IF", "Name", stackPath)
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
	db := backRepo.BackRepoREQ_IF.GetDB()

	// Get model if exist
	var req_ifDB orm.REQ_IFDB
	if _, err := db.First(&req_ifDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&req_ifDB)

	// get an instance (not staged) from DB instance, and call callback function
	req_ifDeleted := new(models.REQ_IF)
	req_ifDB.CopyBasicFieldsToREQ_IF(req_ifDeleted)

	// get stage instance from DB instance, and call callback function
	req_ifStaged := backRepo.BackRepoREQ_IF.Map_REQ_IFDBID_REQ_IFPtr[req_ifDB.ID]
	if req_ifStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), req_ifStaged, req_ifDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
