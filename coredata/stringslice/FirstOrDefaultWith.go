package stringslice

func FirstOrDefaultWith(slice []string, defaultValue string) (result string, isSuccess bool) {
	if len(slice) == 0 {
		return defaultValue, false
	}

	return slice[0], true
}
