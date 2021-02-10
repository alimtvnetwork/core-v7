package intunique

func GetMap(intSlice *[]int) *map[int]bool {
	if intSlice == nil {
		return nil
	}

	length := len(*intSlice)
	dict := make(map[int]bool, length)

	if length == 0 {
		return &dict
	}

	for _, v := range *intSlice {
		dict[v] = true
	}

	return &dict
}
