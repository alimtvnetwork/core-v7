package coreinterface

type NameValuer interface {
	// NameValue
	// should be a combined string output using name[value]
	NameValue() string
}
