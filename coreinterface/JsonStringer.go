package coreinterface

type JsonCombineStringer interface {
	StringJsoner
	// JustMustStringer panic if any error
	JustMustStringer
}
