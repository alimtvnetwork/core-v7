package coreinterface

type StandardTaskEntityDefiner interface {
	TaskEntityDefiner
	AttributesGetter
	IdentifierGetter
	IntegerIdentifier
	HasErrorChecker
	ValueReflectSetter
	SerializerDeserializer
}
