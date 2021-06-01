package coreinterface

type ItemAtRemover interface {
	RemoveAt(index int) (isSuccess bool)
}
