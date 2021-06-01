package converters

// StringsToBytes panic if not a number or more than 255 or less than 0
func StringsToBytes(strArray *[]string) *[]byte {
	results := make([]byte, len(*strArray))

	for i, v := range *strArray {
		vInt, err := StringToByte(v)

		if err != nil {
			panic(err)
		}

		results[i] = vInt
	}

	return &results
}
