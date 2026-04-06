package stringslice

func MakePtr(length, capacity int) []string {
	return make([]string, length, capacity)
}
