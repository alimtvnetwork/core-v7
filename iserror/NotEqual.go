package iserror

func NotEqual(left, right error) bool {
	return !Equal(left, right)
}
