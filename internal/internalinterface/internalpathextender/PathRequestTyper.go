package internalpathextender

type PathRequestTyper interface {
	IsParentDirRequest() bool
	IsRootRequest() bool
	IsRelativeRequest() bool
}
