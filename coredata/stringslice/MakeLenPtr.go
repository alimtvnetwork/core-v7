package stringslice

func MakeLenPtr(length int) *[]string {
	slice := make([]string, length)

	return &slice
}
