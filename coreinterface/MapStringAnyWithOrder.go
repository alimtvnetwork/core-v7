package coreinterface

type MapStringAnyWithOrder interface {
	MapStringAnyGetter
	ValueTypeName() string
	AllKeysSorted() (sortedKeys []string)
	AllKeys() (keysUnsorted []string)
}
