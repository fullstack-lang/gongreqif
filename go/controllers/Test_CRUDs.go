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
var __Test__dummysDeclaration__ models.Test
var __Test_time__dummyDeclaration time.Duration

var mutexTest sync.Mutex

// An TestID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getTest updateTest deleteTest
type TestID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// TestInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postTest updateTest
type TestInput struct {
	// The Test to submit or modify
	// in: body
	Test *orm.TestAPI
}

// GetTests
//
// swagger:route GET /tests tests getTests
//
// # Get all tests
//
// Responses:
// default: genericError
//
//	200: testDBResponse
func (controller *Controller) GetTests(c *gin.Context) {

	// source slice
	var testDBs []orm.TestDB

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTests", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTest.GetDB()

	query := db.Find(&testDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	testAPIs := make([]orm.TestAPI, 0)

	// for each test, update fields from the database nullable fields
	for idx := range testDBs {
		testDB := &testDBs[idx]
		_ = testDB
		var testAPI orm.TestAPI

		// insertion point for updating fields
		testAPI.ID = testDB.ID
		testDB.CopyBasicFieldsToTest_WOP(&testAPI.Test_WOP)
		testAPI.TestPointersEncoding = testDB.TestPointersEncoding
		testAPIs = append(testAPIs, testAPI)
	}

	c.JSON(http.StatusOK, testAPIs)
}

// PostTest
//
// swagger:route POST /tests tests postTest
//
// Creates a test
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostTest(c *gin.Context) {

	mutexTest.Lock()
	defer mutexTest.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostTests", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTest.GetDB()

	// Validate input
	var input orm.TestAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create test
	testDB := orm.TestDB{}
	testDB.TestPointersEncoding = input.TestPointersEncoding
	testDB.CopyBasicFieldsFromTest_WOP(&input.Test_WOP)

	query := db.Create(&testDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoTest.CheckoutPhaseOneInstance(&testDB)
	test := backRepo.BackRepoTest.Map_TestDBID_TestPtr[testDB.ID]

	if test != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), test)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, testDB)
}

// GetTest
//
// swagger:route GET /tests/{ID} tests getTest
//
// Gets the details for a test.
//
// Responses:
// default: genericError
//
//	200: testDBResponse
func (controller *Controller) GetTest(c *gin.Context) {

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetTest", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTest.GetDB()

	// Get testDB in DB
	var testDB orm.TestDB
	if err := db.First(&testDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var testAPI orm.TestAPI
	testAPI.ID = testDB.ID
	testAPI.TestPointersEncoding = testDB.TestPointersEncoding
	testDB.CopyBasicFieldsToTest_WOP(&testAPI.Test_WOP)

	c.JSON(http.StatusOK, testAPI)
}

// UpdateTest
//
// swagger:route PATCH /tests/{ID} tests updateTest
//
// # Update a test
//
// Responses:
// default: genericError
//
//	200: testDBResponse
func (controller *Controller) UpdateTest(c *gin.Context) {

	mutexTest.Lock()
	defer mutexTest.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateTest", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTest.GetDB()

	// Validate input
	var input orm.TestAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var testDB orm.TestDB

	// fetch the test
	query := db.First(&testDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	testDB.CopyBasicFieldsFromTest_WOP(&input.Test_WOP)
	testDB.TestPointersEncoding = input.TestPointersEncoding

	query = db.Model(&testDB).Updates(testDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	testNew := new(models.Test)
	testDB.CopyBasicFieldsToTest(testNew)

	// redeem pointers
	testDB.DecodePointers(backRepo, testNew)

	// get stage instance from DB instance, and call callback function
	testOld := backRepo.BackRepoTest.Map_TestDBID_TestPtr[testDB.ID]
	if testOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), testOld, testNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the testDB
	c.JSON(http.StatusOK, testDB)
}

// DeleteTest
//
// swagger:route DELETE /tests/{ID} tests deleteTest
//
// # Delete a test
//
// default: genericError
//
//	200: testDBResponse
func (controller *Controller) DeleteTest(c *gin.Context) {

	mutexTest.Lock()
	defer mutexTest.Unlock()

	_values := c.Request.URL.Query()
	stackPath := ""
	if len(_values) == 1 {
		value := _values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteTest", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongreqif/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoTest.GetDB()

	// Get model if exist
	var testDB orm.TestDB
	if err := db.First(&testDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&testDB)

	// get an instance (not staged) from DB instance, and call callback function
	testDeleted := new(models.Test)
	testDB.CopyBasicFieldsToTest(testDeleted)

	// get stage instance from DB instance, and call callback function
	testStaged := backRepo.BackRepoTest.Map_TestDBID_TestPtr[testDB.ID]
	if testStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), testStaged, testDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
