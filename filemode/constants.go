package filemode

// Reference : https://ss64.com/bash/chmod.html
const (
	SupportedLengthString = "3"
	SupportedLength       = 3
	ReadValue             = 4
	WriteValue            = 2
	ExecuteValue          = 1

	ReadWriteValue        = ReadValue + WriteValue
	ReadExecuteValue      = ReadValue + ExecuteValue
	WriteExecuteValue     = WriteValue + ExecuteValue
	ReadWriteExecuteValue = ReadValue + WriteValue + ExecuteValue

	OwnerIndex = 0
	GroupIndex = 1
	OtherIndex = 2

	ReadChar    byte = 'r'
	WriteChar   byte = 'w'
	ExecuteChar byte = 'x'
)
