package stringslice

func Make(length, capacity int) []string {
	return make([]string, length, capacity)
}
