package coreinterface

type JsonCombineStringer interface {
	JsonStringer
	// MustJsonStringer panic if any error
	MustJsonStringer
}
