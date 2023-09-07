package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/constants"
)

// Reference : https://ss64.com/bash/chmod.html
const (
	SingleRwxLengthString                           = "3"
	HyphenedRwxLength                               = constants.ArbitraryCapacity10
	HyphenedRwxLengthString                         = constants.N10String
	FullRwxLengthWithoutHyphenString                = constants.N9String
	FullRwxLengthWithoutHyphen                      = constants.ArbitraryCapacity9
	SingleRwxLength                                 = 3
	ReadValue                                       = 4
	WriteValue                                      = 2
	ExecuteValue                                    = 1
	ReadWriteValue                                  = ReadValue + WriteValue
	ReadExecuteValue                                = ReadValue + ExecuteValue
	WriteExecuteValue                               = WriteValue + ExecuteValue
	ReadWriteExecuteValue                           = ReadValue + WriteValue + ExecuteValue
	OwnerIndex                                      = 0
	GroupIndex                                      = 1
	OtherIndex                                      = 2
	ReadChar                            byte        = 'r'
	NopChar                             byte        = '-'
	WriteChar                           byte        = 'w'
	ExecuteChar                         byte        = 'x'
	AllWildcards                                    = "***"
	pathInvalidMessage                              = "Path invalid or permission issue. path : "
	chmodExpectationFailed                          = "chmod expectation failed"
	messageWithPathWrappedFormat                    = "%s Path: (\"%s\")"     // message, path
	fileDefaultChmod                    os.FileMode = 0644                    // cannot execute by everyone OwnerCanReadWriteGroupOtherCanReadOnly
	dirDefaultChmod                     os.FileMode = 0755                    // can execute by everyone OwnerCanDoAllExecuteGroupOtherCanReadExecute
	fileModeStringFriendlyDisplayFormat             = "{chmod : \"%s (%s)\"}" // chmod - 0777, -rwxrwxrwx
)
