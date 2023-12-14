package convertinteranl

import "go/format"

type codeFormatter struct{}

func (it codeFormatter) GolangRaw(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return []byte{}, nil
	}

	return format.Source(b)
}

func (it codeFormatter) Golang(code string) (string, error) {
	if len(code) == 0 {
		return "", nil
	}

	formattedCode, err := format.Source([]byte(code))

	return string(formattedCode), err
}
