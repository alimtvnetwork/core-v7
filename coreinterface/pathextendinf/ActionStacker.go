package pathextendinf

type ActionStacker interface {
	Action(action ActionExecutorFunc) PathExtender
}
