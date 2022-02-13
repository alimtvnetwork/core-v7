package internalpathextender

type BasePathExtender interface {
	Identifier
	PathInfoer
}

type PathRequestTyper interface {
	IsParentDirRequest() bool
	IsRootRequest() bool
	IsRelativeRequest() bool
}

type FileListerTyper interface {
	IsAllFilesOnly() bool
	IsAllFilesOnlyRecursive() bool
	IsAllFilesOrDirs() bool
	IsAllFilesOrDirsRecursive() bool
	IsAllDirsOnly() bool
	IsAllDirsOnlyRecursive() bool
}
