package chmodhelper

import "gitlab.com/evatix-go/core/constants"

// Reference : https://ss64.com/bash/chmod.html
const (
	SingleRwxLengthString                 = "3"
	HyphenedRwxLength                     = constants.ArbitraryCapacity10
	HyphenedRwxLengthString               = constants.N10String
	FullRwxLengthWithoutHyphenString      = constants.N9String
	FullRwxLengthWithoutHyphen            = constants.ArbitraryCapacity9
	SingleRwxLength                       = 3
	ReadValue                             = 4
	WriteValue                            = 2
	ExecuteValue                          = 1
	ReadWriteValue                        = ReadValue + WriteValue
	ReadExecuteValue                      = ReadValue + ExecuteValue
	WriteExecuteValue                     = WriteValue + ExecuteValue
	ReadWriteExecuteValue                 = ReadValue + WriteValue + ExecuteValue
	chmod                                 = "chmod"
	OwnerIndex                            = 0
	GroupIndex                            = 1
	OtherIndex                            = 2
	ReadChar                         byte = 'r'
	NopChar                          byte = '-'
	WriteChar                        byte = 'w'
	ExecuteChar                      byte = 'x'
	AllWildcards                          = "***"
)
