package coreinterface

type SafeStringRemover interface {
	SafeRemove(key string) (isSuccess bool)
}
