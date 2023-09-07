package internalpathextender

type FileListerTyper interface {
	IsAllFilesOnly() bool
	IsAllFilesOnlyRecursive() bool
	IsAllFilesOrDirs() bool
	IsAllFilesOrDirsRecursive() bool
	IsAllDirsOnly() bool
	IsAllDirsOnlyRecursive() bool
}
