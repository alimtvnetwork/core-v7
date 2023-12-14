package jsoninternal

import "encoding/json"

type anyToConvert struct{}

func (it anyToConvert) SafeString(anyItem interface{}) string {
	s, _ := it.String(anyItem)

	return s
}

// String
//
// It is not pretty JSON
func (it anyToConvert) String(
	anyItem interface{},
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return string(b), err
}

func (it anyToConvert) PrettyString(
	prefix string,
	anyItem interface{},
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return Pretty.Bytes.Prefix(prefix, b)
}

func (it anyToConvert) PrettyStringIndent(
	prefix,
	curIndent string,
	anyItem interface{},
) (string, error) {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return "", err
	}

	return Pretty.Bytes.Indent(prefix, curIndent, b)
}

func (it anyToConvert) SafePrettyString(
	prefix string,
	anyItem interface{},
) string {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return ""
	}

	return Pretty.Bytes.Safe(prefix, b)
}

func (it anyToConvert) PrettyStringDefault(
	anyItem interface{},
) string {
	b, err := json.Marshal(anyItem)

	if len(b) == 0 || err != nil {
		return ""
	}

	return Pretty.Bytes.Safe("", b)
}

func (it anyToConvert) PrettyStringDefaultMust(
	anyItem interface{},
) string {
	b, err := json.Marshal(anyItem)

	if err != nil {
		panic(err)
	}

	if len(b) == 0 || err != nil {
		return ""
	}

	return Pretty.Bytes.DefaultMust(b)
}
