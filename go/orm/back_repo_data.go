// generated code - do not edit
package orm

type BackRepoData struct {
	// insertion point for slices

	TestAPIs []*TestAPI
}

func CopyBackRepoToBackRepoData(backRepo *BackRepoStruct, backRepoData *BackRepoData) {
	// insertion point for slices copies
	for _, testDB := range backRepo.BackRepoTest.Map_TestDBID_TestDB {

		var testAPI TestAPI
		testAPI.ID = testDB.ID
		testAPI.TestPointersEncoding = testDB.TestPointersEncoding
		testDB.CopyBasicFieldsToTest_WOP(&testAPI.Test_WOP)

		backRepoData.TestAPIs = append(backRepoData.TestAPIs, &testAPI)
	}

}
