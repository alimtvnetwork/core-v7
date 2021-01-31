package corejson

type ParseSelfInjector interface {
	JsonParseSelfInject(jsonResult *Result)
}
