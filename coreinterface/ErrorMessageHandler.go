package coreinterface

type ErrorMessageHandler interface {
	HandleErrorWithMsg(newMessage string)
	HandleErrorWithRefs(
		newMessage string,
		refVar,
		refVal interface{},
	)
}
