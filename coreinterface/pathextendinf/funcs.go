package pathextendinf

type (
	FilterFunc         func(filePath string) (isTake bool, err error)
	ActionExecutorFunc func(extender PathExtender) error
)
