package internalpathextender

type IsPathChecker interface {
	// IsFile
	//
	//  Will be set during creation.
	IsFile() bool
	// IsDir
	//
	//  Will be set during creation.
	IsDir() bool
	IsExist() bool
}
