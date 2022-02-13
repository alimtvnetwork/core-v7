package pathextendinf

type AttachModifier interface {
	// ParentDir
	//
	//  Actions related to ParentDir
	ParentDir() Modifier
	// RootDir
	//
	//  Actions related to RootDir
	RootDir() Modifier
	// FilePath
	//
	//  Actions related to FilePath
	FilePath() Modifier
}
