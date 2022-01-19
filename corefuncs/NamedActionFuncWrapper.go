package corefuncs

type NamedActionFuncWrapper struct {
	Name   string
	Action NamedActionFunc
}

func (it NamedActionFuncWrapper) Exec() {
	it.Action(it.Name)
}

func (it NamedActionFuncWrapper) Next(next *NamedActionFuncWrapper) {
	it.Exec()
	next.Exec()
}

func (it NamedActionFuncWrapper) AsActionFunc() ActionFunc {
	return func() {
		it.Action(it.Name)
	}
}

func (it NamedActionFuncWrapper) AsActionReturnsErrorFunc() ActionReturnsErrorFunc {
	return func() error {
		it.Action(it.Name)

		return nil
	}
}
