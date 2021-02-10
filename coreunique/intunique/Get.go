package intunique

func Get(intSlice *[]int) *[]int {
	if intSlice == nil {
		return intSlice
	}

	length := len(*intSlice)
	if length == 0 {
		return intSlice
	}

	dict := make(map[int]bool, length)

	for _, v := range *intSlice {
		dict[v] = true
	}

	newSlice := make([]int, 0, len(dict))

	for k := range dict {
		newSlice = append(newSlice, k)
	}

	return &newSlice
}
