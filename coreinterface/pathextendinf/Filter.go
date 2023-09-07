package pathextendinf

type Filter interface {
	AllFilesFilter(isRecursive bool, filter FilterFunc) ([]string, error)
	AllDirsFilter(isRecursive bool, filter FilterFunc) ([]string, error)
	AllFilesOrDirsFilter(isRecursive bool, filter FilterFunc) ([]string, error)

	Filter(listType FileListerTyper, filter FilterFunc) ([]string, error)
	FilterAsExtender(listType FileListerTyper, filter FilterFunc) ([]string, error)
	// SkipFilter
	//
	//  Skips the matching criteria
	SkipFilter(listType FileListerTyper, filter FilterFunc) ([]string, error)
	// SkipFilterAsExtender
	//
	//  Skips the matching criteria
	SkipFilterAsExtender(listType FileListerTyper, filter FilterFunc) ([]string, error)
}
