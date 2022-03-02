package coreinterface

type KeyValueStringDefiner interface {
	KeyStringNameGetter
	VariableNameStringGetter
	ExplicitValueStringGetter
	IsVariableNameEqual(name string) bool
	IsValueEqual(valueString string) bool
	IsEqualKeyValueStringDefiner(right KeyValueStringDefiner) bool

	CoreDefiner
	StringCompiler
}

type KeyAnyValueDefiner interface {
	KeyStringNameGetter
	VariableNameStringGetter
	ValueAnyGetter
	ExplicitValueStringGetter
	IsVariableNameEqual(name string) bool
	IsAnyValueEqual(right interface{}) bool
	IsEqualKeyAnyValueDefiner(right KeyAnyValueDefiner) bool

	CoreDefiner
	StringCompiler
}

type KeyStringValuesCollectionDefiner interface {
	KeyValueStringDefiners() []KeyValueStringDefiner

	AllKeysStringer
	AllKeysSortedStringer

	HashmapGetter
	KeysHashsetGetter
	StringsGetter
	HasKeyChecker

	ValueOfKey(key string) (val string)
	ValueOfKeys(keys ...string) (values []string)

	IsEqualKeyStringValuesCollectionDefiner(
		right KeyStringValuesCollectionDefiner,
	) bool
	CoreDefiner
	StringCompiler
}

type KeyAnyValuesCollectionDefiner interface {
	KeyValueStringDefiners() []KeyAnyValueDefiner

	AllKeysStringer
	AllKeysSortedStringer

	HashmapGetter
	KeysHashsetGetter
	StringsGetter
	MapStringAnyGetter
	HasKeyChecker
	ValueOfKey(key string) (valInf interface{})
	ValueOfKeys(keys ...string) (valuesAnyItems []interface{})

	IsEqualKeyAnyValuesCollectionDefiner(
		right KeyAnyValuesCollectionDefiner,
	) bool

	CoreDefiner
	StringCompiler
}
