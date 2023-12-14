package jsoninternal

import (
	"bytes"
	"encoding/json"
)

type stringToPrettyConvert struct{}

func (it stringToPrettyConvert) Safe(
	prefix,
	src string,
) string {
	output, _ := it.Prefix(prefix, src)

	return output
}

func (it stringToPrettyConvert) SafeDefault(
	src string,
) string {
	output, _ := it.Prefix("", src)

	return output
}

func (it stringToPrettyConvert) Prefix(
	prefix string,
	src string,
) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(src), prefix, indent)

	return prettyJSON.String(), err
}

func (it stringToPrettyConvert) Indent(
	prefix string,
	curIndent,
	src string,
) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(
		&prettyJSON,
		[]byte(src),
		prefix,
		curIndent,
	)

	return prettyJSON.String(), err
}

func (it stringToPrettyConvert) PrefixMust(prefix string, src string) string {
	output, err := it.Prefix(prefix, src)

	if err != nil {
		panic(err)
	}

	return output
}

func (it stringToPrettyConvert) DefaultMust(src string) string {
	output, err := it.Prefix("", src)

	if err != nil {
		panic(err)
	}

	return output
}
