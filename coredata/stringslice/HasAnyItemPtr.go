package stringslice

func HasAnyItemPtr(slice []string) bool {
	return !IsEmptyPtr(slice)
}
