package baseactioninf

type ActionReturnsErrorFuncBinder interface {
	Exec() (err error)
}
