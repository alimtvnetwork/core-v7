package jsoninternal

import (
	"bytes"
	"encoding/json"
)

type bytesToPrettyConvert struct{}

func (it bytesToPrettyConvert) Safe(
	prefix string, b []byte,
) string {
	output, _ := it.Prefix(prefix, b)

	return output
}

func (it bytesToPrettyConvert) SafeDefault(
	b []byte,
) string {
	output, _ := it.Prefix("", b)

	return output
}

func (it bytesToPrettyConvert) Prefix(
	prefix string,
	b []byte,
) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, b, prefix, indent)

	return prettyJSON.String(), err
}

func (it bytesToPrettyConvert) Indent(
	prefix,
	curIndent string,
	b []byte,
) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, b, prefix, curIndent)

	return prettyJSON.String(), err
}

func (it bytesToPrettyConvert) PrefixMust(prefix string, b []byte) string {
	output, err := it.Prefix(prefix, b)

	if err != nil {
		panic(err)
	}

	return output
}

func (it bytesToPrettyConvert) DefaultMust(b []byte) string {
	output, err := it.Prefix("", b)

	if err != nil {
		panic(err)
	}

	return output
}
