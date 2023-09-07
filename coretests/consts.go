package coretests

import "gitlab.com/auk-go/core/constants"

const (
	commonJoiner = constants.Space
	// notEqualComparisonMessageFormat
	//
	//  - Left FullPublicFieldsJson, f1String, f2Integer
	//  - Right FullPublicFieldsJson, f1String, f2Integer
	notEqualComparisonMessageFormat = "Compare :\n" +
		" Left  = %s f1String=%s f2Int=%d\n" +
		" Right = %s f1String=%s f2Int=%d\n"
)
