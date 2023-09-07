package pathextendinf

type DirFileLister interface {
	// Files
	//
	//  Returns all files from top level, no nesting.
	Files() (string, error)
	// FilesRecursively
	//
	//  Returns all files from top level to nesting (Recursively).
	FilesRecursively() (string, error)
	// FilesOrDirs
	//
	//  Returns all files or dirs from top level, no nesting.
	FilesOrDirs() (string, error)
	FilesOrDirsRecursively() (string, error)
	Dirs() (string, error)
	DirsRecursively() (string, error)
	Filter
	AttachModifier
}
