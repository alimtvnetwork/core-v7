package enumimpl

var (
	New                      = newCreator{}
	DefaultDiffCheckerImpl   = &differCheckerImpl{}
	LeftRightDiffCheckerImpl = &leftRightDiffCheckerImpl{}
)
