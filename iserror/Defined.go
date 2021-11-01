package iserror

func Defined(err error) bool {
	return err != nil
}
