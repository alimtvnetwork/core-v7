package iserror

func NotEmpty(err error) bool {
	return err != nil
}
