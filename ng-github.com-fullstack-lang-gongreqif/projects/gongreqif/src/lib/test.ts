// generated code - do not edit

import { TestAPI } from './test-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Test {

	static GONGSTRUCT_NAME = "Test"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyTestToTestAPI(test: Test, testAPI: TestAPI) {

	testAPI.CreatedAt = test.CreatedAt
	testAPI.DeletedAt = test.DeletedAt
	testAPI.ID = test.ID

	// insertion point for basic fields copy operations
	testAPI.Name = test.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyTestAPIToTest update basic, pointers and slice of pointers fields of test
// from respectively the basic fields and encoded fields of pointers and slices of pointers of testAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyTestAPIToTest(testAPI: TestAPI, test: Test, frontRepo: FrontRepo) {

	test.CreatedAt = testAPI.CreatedAt
	test.DeletedAt = testAPI.DeletedAt
	test.ID = testAPI.ID

	// insertion point for basic fields copy operations
	test.Name = testAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
