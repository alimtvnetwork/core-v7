package corejson

type Jsoner interface {
	Json() *Result
	JsonModelAny() interface{}
}
