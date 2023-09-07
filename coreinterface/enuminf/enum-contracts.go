package enuminf

type BasicContractsEnumer interface {
	BasicEnumer
	TypeNameWithRangeNamesCsvGetter
}

type BasicByteContractsEnumer interface {
	BasicContractsEnumer
	IsValueByteEqualer
	IsAnyValueByteEqualer
	BasicByteEnumer
}

type BasicByteEnumContractsBinder interface {
	BasicByteContractsEnumer
	AsBasicByteEnumContractsBinder() BasicByteEnumContractsBinder
}

type BasicByteEnumContractsDelegateBinder interface {
	AsBasicByteEnumContractsDelegateBinder() BasicByteEnumContractsDelegateBinder
}

type BasicInt8ContractsEnumer interface {
	BasicContractsEnumer
	BasicInt8Enumer
	Int8ToEnumStringer
	IsValueInteger8Equaler
	IsAnyValueInteger8Equaler
}

type BasicInt8EnumContractsBinder interface {
	BasicInt8ContractsEnumer
	AsBasicInt8EnumContractsBinder() BasicInt8EnumContractsBinder
}

type BasicInt16ContractsEnumer interface {
	BasicContractsEnumer
	BasicInt16Enumer
	Int16ToEnumStringer
	IsValueInteger16Equaler
	IsAnyValueInteger16Equaler
}

type BasicInt16EnumContractsBinder interface {
	BasicInt16ContractsEnumer
	AsBasicIn16EnumContractsBinder() BasicInt16ContractsEnumer
}

type BasicInt32ContractsEnumer interface {
	BasicContractsEnumer
	BasicInt32Enumer
	Int32ToEnumStringer
	IsValueInteger32Equaler
	IsAnyValueInteger32Equaler
}

type BasicInt32EnumContractsBinder interface {
	BasicInt32ContractsEnumer
	AsBasicInt32EnumContractsBinder() BasicInt32ContractsEnumer
}

type BasicEnumContractsBinder interface {
	BasicContractsEnumer
	AsBasicEnumContractsBinder() BasicEnumContractsBinder
}

type StandardEnumerContractsBinder interface {
	StandardEnumer
	AsStandardEnumerContractsBinder() StandardEnumerContractsBinder
}
