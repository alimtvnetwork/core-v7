package pathextendinf

type PathExtender interface {
	PathInfoer
	// ParentDirExtender
	//
	//  Refers to current full-path's parent dir.
	//  ParentDir or ParentDirExtender is different from RootDir.
	ParentDirExtender() PathExtender
	// RootExtender
	//
	//  Refers to start of the dir
	//  For example a repo start point
	RootExtender() PathExtender
	FullPathExtender() PathExtender
	RelativeExtender() PathExtender
	Joiner() Joiner

	IsEqual(right PathExtender) bool
	ActionStacker
	ActionExecutor
	Cloner() Cloner
}
