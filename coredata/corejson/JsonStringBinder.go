package corejson

type JsonStringBinder interface {
	JsonStringer
	PrettyJsonStringer
	AsJsonStringBinder() JsonStringBinder
}
