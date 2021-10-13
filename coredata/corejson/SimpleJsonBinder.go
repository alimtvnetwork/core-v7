package corejson

type SimpleJsonBinder interface {
	Jsoner
	AsSimpleJsonBinder() SimpleJsonBinder
}
