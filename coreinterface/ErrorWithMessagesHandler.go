package coreinterface

type ErrorWithMessagesHandler interface {
	HandleErrorWith(messages ...string)
}
