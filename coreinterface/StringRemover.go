package coreinterface

type StringRemover interface {
	Remove(key string) (isSuccess bool)
}
