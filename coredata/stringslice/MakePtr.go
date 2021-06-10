package stringslice

func MakePtr(length, capacity int) *[]string {
	slice := make([]string, length, capacity)

	return &slice
}
