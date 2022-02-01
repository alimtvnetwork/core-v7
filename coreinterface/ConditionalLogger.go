package coreinterface

type ConditionalLogger interface {
	On(isCondition bool) StandardLogger
	OnErr(err error) StandardLogger
	OnString(message string) StandardLogger
	OnBytes(rawBytes []byte) StandardLogger
}
