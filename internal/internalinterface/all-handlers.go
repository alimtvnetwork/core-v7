package internalinterface

type ErrorHandler interface {
	// HandleError
	//
	// Only call panic
	// if it has any error
	HandleError()
}

type ErrorMessageHandler interface {
	HandleErrorWithRefs(
		newMessage string,
		refVar,
		refVal interface{},
	)
	// HandleErrorWithMsg
	//
	//  Only call panic on has currentError
	HandleErrorWithMsg(newMessage string)
}
