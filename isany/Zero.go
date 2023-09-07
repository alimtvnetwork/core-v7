package isany

import "gitlab.com/auk-go/core/internal/reflectinternal"

// Zero
//
//  returns true if the current value is null
//  or reflect value is zero
//
// Reference:
//  - Stackoverflow Example : https://stackoverflow.com/a/23555352
func Zero(anyItem interface{}) bool {
	return reflectinternal.IsZero(anyItem)
}
