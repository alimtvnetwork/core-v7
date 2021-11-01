package coreinterface

type ErrorRefMessageHandler interface {
	HandleErrorWithRefs(
		newMessage string,
		refVar,
		refVal interface{},
	)
}
