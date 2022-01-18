package coretests

func AnyToBytesPtr(anyItem interface{}) *[]byte {
	switch expectedAs := anyItem.(type) {
	case []byte:
		if expectedAs == nil {
			return nil
		}

		return &expectedAs
	case *[]byte:
		if expectedAs == nil {
			return nil
		}

		return expectedAs
	default:
		return nil
	}
}
