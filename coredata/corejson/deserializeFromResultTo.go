package corejson

type deserializeFromResultTo struct{}

func (it deserializeFromResultTo) String(
	jsonResult *Result,
) (line string, err error) {
	err = Deserialize.Apply(jsonResult, &line)

	return line, err
}

func (it deserializeFromResultTo) Bool(
	jsonResult *Result,
) (isResult bool, err error) {
	err = Deserialize.Apply(jsonResult, &isResult)

	return isResult, err
}

func (it deserializeFromResultTo) Byte(
	jsonResult *Result,
) (byteVal byte, err error) {
	err = Deserialize.Apply(jsonResult, &byteVal)

	return byteVal, err
}

func (it deserializeFromResultTo) ByteMust(
	jsonResult *Result,
) byte {
	result, err := it.Byte(jsonResult)

	if err != nil {
		panic(err)
	}

	return result
}

func (it deserializeFromResultTo) BoolMust(
	jsonResult *Result,
) bool {
	result, err := it.Bool(jsonResult)

	if err != nil {
		panic(err)
	}

	return result
}

func (it deserializeFromResultTo) StringMust(
	jsonResult *Result,
) string {
	result, err := it.String(jsonResult)

	if err != nil {
		panic(err)
	}

	return result
}

func (it deserializeFromResultTo) StringsMust(
	jsonResult *Result,
) (lines []string) {
	err := jsonResult.Deserialize(&lines)

	if err != nil {
		panic(err)
	}

	return lines
}
