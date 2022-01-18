package coretests

func AnyToDraftType(anyItem interface{}) *DraftType {
	switch expectedAs := anyItem.(type) {
	case DraftType:
		return &expectedAs
	case *DraftType:
		return expectedAs
	default:
		return nil
	}
}
