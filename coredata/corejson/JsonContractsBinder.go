package corejson

type JsonContractsBinder interface {
	Jsoner
	JsonMarshaller
	JsonParseSelfInjector
	AsJsonContractsBinder() JsonContractsBinder
}
