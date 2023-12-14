package jsoninternal

type prettyConverter struct {
	Bytes  bytesToPrettyConvert
	String stringToPrettyConvert
	AnyTo  anyToConvert // it doesn't convert all to Pretty JSON
}
