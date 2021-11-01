package corejson

type Jsoner interface {
	Json() Result
	JsonPtr() *Result
}
