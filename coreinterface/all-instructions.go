package coreinterface

type FromToer interface {
	IsNull() bool
	IsFromEmpty() bool
	IsToEmpty() bool
	FromToNameSetter
	FromToNameGetter
}

type FromToNameSetter interface {
	SetToName(to string)
	SetFromName(form string)
}

type FromToNameGetter interface {
	ToName() string
	FromName() string
	String() string
}
