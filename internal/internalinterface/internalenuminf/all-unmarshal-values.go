package internalenuminf

type UnmarshallToValueByter interface {
	UnmarshallToValue(isMappedToFirstIfEmpty bool, jsonUnmarshallingValue []byte) (byte, error)
}

type UnmarshallToValueInter interface {
	UnmarshallToValue(isMappedToFirstIfEmpty bool, jsonUnmarshallingValue []byte) (int, error)
}

type UnmarshallToValueInt8er interface {
	UnmarshallToValue(isMappedToFirstIfEmpty bool, jsonUnmarshallingValue []byte) (int8, error)
}

type UnmarshallToValueInt16er interface {
	UnmarshallToValue(isMappedToFirstIfEmpty bool, jsonUnmarshallingValue []byte) (int16, error)
}

type UnmarshallEnumToValueByter interface {
	UnmarshallEnumToValue(jsonUnmarshallingValue []byte) (byte, error)
}
