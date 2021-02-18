package core

func EmptyAnysPtr() *[]interface{} {
	return &([]interface{}{})
}

func EmptyFloat32Ptr() *[]float32 {
	return &([]float32{})
}

func EmptyFloat64Ptr() *[]float64 {
	return &([]float64{})
}

func EmptyBoolsPtr() *[]bool {
	return &([]bool{})
}

func EmptyIntsPtr() *[]int {
	return &([]int{})
}

func EmptyBytePtr() *[]byte {
	return &([]byte{})
}

func EmptyStringsMapPtr() *map[string]string {
	return &(map[string]string{})
}

func EmptyStringToIntMapPtr() *map[string]int {
	return &(map[string]int{})
}

func EmptyStringsPtr() *[]string {
	return &([]string{})
}

func EmptyPointerStringsPtr() *[]*string {
	return &([]*string{})
}

func StringsPtrByLength(length int) *[]string {
	list := make([]string, length)

	return &(list)
}

func StringsPtrByCapacity(length, cap int) *[]string {
	list := make([]string, length, cap)

	return &(list)
}

func PointerStringsPtrByCapacity(length, cap int) *[]*string {
	list := make([]*string, length, cap)

	return &(list)
}
