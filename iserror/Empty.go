package iserror

func Empty(err error) bool {
	return err == nil
}
