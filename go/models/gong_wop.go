// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Test_WOP struct {
	// insertion point
	Name string
}

func (from *Test) CopyBasicFields(to *Test) {
	// insertion point
	to.Name = from.Name
}
