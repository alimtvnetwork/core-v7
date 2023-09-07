package loggerinf

type BasePersistentLogger interface {
	Attr() AttrPersistentLogger
	NewGeneralWriter
	AllLogWriter
	StandardLogger

	PersistentIdGetter
	ParentPersistentIdGetter
}
