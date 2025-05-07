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
var __RELATION_GROUP__dummysDeclaration__ models.RELATION_GROUP
var __RELATION_GROUP_time__dummyDeclaration time.Duration

var mutexRELATION_GROUP sync.Mutex

// An RELATION_GROUPID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getRELATION_GROUP updateRELATION_GROUP deleteRELATION_GROUP
type RELATION_GROUPID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// RELATION_GROUPInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postRELATION_GROUP updateRELATION_GROUP
type RELATION_GROUPInput struct {
	// The RELATION_GROUP to submit or modify
	// in: body
	RELATION_GROUP *orm.RELATION_GROUPAPI
}

// GetRELATION_GROUPs
//
// swagger:route GET /relation_groups relation_groups getRELATION_GROUPs
//
// # Get all relation_groups
//
// Responses:
// default: genericError
//
//	200: relation_groupDBResponse
func (controller *Controller) GetRELATION_GROUPs(c *gin.Context) {

	// source slice
	var relation_groupDBs []orm.RELATION_GROUPDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRELATION_GROUPs", "Name", stackPath)
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
	db := backRepo.BackRepoRELATION_GROUP.GetDB()

	_, err := db.Find(&relation_groupDBs)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	relation_groupAPIs := make([]orm.RELATION_GROUPAPI, 0)

	// for each relation_group, update fields from the database nullable fields
	for idx := range relation_groupDBs {
		relation_groupDB := &relation_groupDBs[idx]
		_ = relation_groupDB
		var relation_groupAPI orm.RELATION_GROUPAPI

		// insertion point for updating fields
		relation_groupAPI.ID = relation_groupDB.ID
		relation_groupDB.CopyBasicFieldsToRELATION_GROUP_WOP(&relation_groupAPI.RELATION_GROUP_WOP)
		relation_groupAPI.RELATION_GROUPPointersEncoding = relation_groupDB.RELATION_GROUPPointersEncoding
		relation_groupAPIs = append(relation_groupAPIs, relation_groupAPI)
	}

	c.JSON(http.StatusOK, relation_groupAPIs)
}

// PostRELATION_GROUP
//
// swagger:route POST /relation_groups relation_groups postRELATION_GROUP
//
// Creates a relation_group
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostRELATION_GROUP(c *gin.Context) {

	mutexRELATION_GROUP.Lock()
	defer mutexRELATION_GROUP.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostRELATION_GROUPs", "Name", stackPath)
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
	db := backRepo.BackRepoRELATION_GROUP.GetDB()

	// Validate input
	var input orm.RELATION_GROUPAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create relation_group
	relation_groupDB := orm.RELATION_GROUPDB{}
	relation_groupDB.RELATION_GROUPPointersEncoding = input.RELATION_GROUPPointersEncoding
	relation_groupDB.CopyBasicFieldsFromRELATION_GROUP_WOP(&input.RELATION_GROUP_WOP)

	_, err = db.Create(&relation_groupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoRELATION_GROUP.CheckoutPhaseOneInstance(&relation_groupDB)
	relation_group := backRepo.BackRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[relation_groupDB.ID]

	if relation_group != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), relation_group)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, relation_groupDB)
}

// GetRELATION_GROUP
//
// swagger:route GET /relation_groups/{ID} relation_groups getRELATION_GROUP
//
// Gets the details for a relation_group.
//
// Responses:
// default: genericError
//
//	200: relation_groupDBResponse
func (controller *Controller) GetRELATION_GROUP(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetRELATION_GROUP", "Name", stackPath)
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
	db := backRepo.BackRepoRELATION_GROUP.GetDB()

	// Get relation_groupDB in DB
	var relation_groupDB orm.RELATION_GROUPDB
	if _, err := db.First(&relation_groupDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var relation_groupAPI orm.RELATION_GROUPAPI
	relation_groupAPI.ID = relation_groupDB.ID
	relation_groupAPI.RELATION_GROUPPointersEncoding = relation_groupDB.RELATION_GROUPPointersEncoding
	relation_groupDB.CopyBasicFieldsToRELATION_GROUP_WOP(&relation_groupAPI.RELATION_GROUP_WOP)

	c.JSON(http.StatusOK, relation_groupAPI)
}

// UpdateRELATION_GROUP
//
// swagger:route PATCH /relation_groups/{ID} relation_groups updateRELATION_GROUP
//
// # Update a relation_group
//
// Responses:
// default: genericError
//
//	200: relation_groupDBResponse
func (controller *Controller) UpdateRELATION_GROUP(c *gin.Context) {

	mutexRELATION_GROUP.Lock()
	defer mutexRELATION_GROUP.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateRELATION_GROUP", "Name", stackPath)
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
	db := backRepo.BackRepoRELATION_GROUP.GetDB()

	// Validate input
	var input orm.RELATION_GROUPAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var relation_groupDB orm.RELATION_GROUPDB

	// fetch the relation_group
	_, err := db.First(&relation_groupDB, c.Param("id"))

	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	relation_groupDB.CopyBasicFieldsFromRELATION_GROUP_WOP(&input.RELATION_GROUP_WOP)
	relation_groupDB.RELATION_GROUPPointersEncoding = input.RELATION_GROUPPointersEncoding

	db, _ = db.Model(&relation_groupDB)
	_, err = db.Updates(&relation_groupDB)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	relation_groupNew := new(models.RELATION_GROUP)
	relation_groupDB.CopyBasicFieldsToRELATION_GROUP(relation_groupNew)

	// redeem pointers
	relation_groupDB.DecodePointers(backRepo, relation_groupNew)

	// get stage instance from DB instance, and call callback function
	relation_groupOld := backRepo.BackRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[relation_groupDB.ID]
	if relation_groupOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), relation_groupOld, relation_groupNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the relation_groupDB
	c.JSON(http.StatusOK, relation_groupDB)
}

// DeleteRELATION_GROUP
//
// swagger:route DELETE /relation_groups/{ID} relation_groups deleteRELATION_GROUP
//
// # Delete a relation_group
//
// default: genericError
//
//	200: relation_groupDBResponse
func (controller *Controller) DeleteRELATION_GROUP(c *gin.Context) {

	mutexRELATION_GROUP.Lock()
	defer mutexRELATION_GROUP.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["Name"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteRELATION_GROUP", "Name", stackPath)
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
	db := backRepo.BackRepoRELATION_GROUP.GetDB()

	// Get model if exist
	var relation_groupDB orm.RELATION_GROUPDB
	if _, err := db.First(&relation_groupDB, c.Param("id")); err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped()
	db.Delete(&relation_groupDB)

	// get an instance (not staged) from DB instance, and call callback function
	relation_groupDeleted := new(models.RELATION_GROUP)
	relation_groupDB.CopyBasicFieldsToRELATION_GROUP(relation_groupDeleted)

	// get stage instance from DB instance, and call callback function
	relation_groupStaged := backRepo.BackRepoRELATION_GROUP.Map_RELATION_GROUPDBID_RELATION_GROUPPtr[relation_groupDB.ID]
	if relation_groupStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), relation_groupStaged, relation_groupDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
