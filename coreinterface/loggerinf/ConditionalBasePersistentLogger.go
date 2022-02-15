package loggerinf

type ConditionalBasePersistentLogger interface {
	On(isCondition bool) BasePersistentLogger
	OnErr(err error) BasePersistentLogger
	OnString(message string) BasePersistentLogger
	OnBytes(rawBytes []byte) BasePersistentLogger
}
