package baseactioninf

type Namer interface {
	Name() string
}

type Initializer interface {
	Initialize()
}

type Planner interface {
	Plan()
}

type IsApplier interface {
	IsApply() (isSuccess bool)
}
