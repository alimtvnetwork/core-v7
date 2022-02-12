package loggerinf

type ConditionalPersistentLogger interface {
	On(isCondition bool) PersistentLogger
	OnErr(err error) PersistentLogger
	OnString(message string) PersistentLogger
	OnBytes(rawBytes []byte) PersistentLogger
}
