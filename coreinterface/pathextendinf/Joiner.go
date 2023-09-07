package pathextendinf

type Joiner interface {
	// ParentDir
	//
	//  Join from ParentDir
	ParentDir(relativePaths ...string) string
	ParentDirExtender(relativePaths ...string) PathExtender

	// RootDir
	//
	//  Join from Root Dir
	RootDir(relativePaths ...string) string
	JoinRootDirExtender(relativePaths ...string) PathExtender

	// FullPath
	//
	//  Join from JoinFullPath
	FullPath(relativePaths ...string) string
	FullPathExtender(relativePaths ...string) PathExtender
}
