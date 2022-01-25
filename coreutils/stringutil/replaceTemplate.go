package stringutil

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

type replaceTemplate struct{}

func (it *replaceTemplate) CurlyOne(
	format string, // {key}-text...
	firstKey string, firstValue interface{},
) string {
	if len(format) == 0 {
		return format
	}

	return it.UsingMapOptions(
		true,
		format,
		map[string]string{
			firstKey: fmt.Sprintf(
				constants.SprintValueFormat,
				firstValue),
		})
}

func (it *replaceTemplate) Curly(
	format string, // {key}-text...
	mapToReplace map[string]string,
) string {
	if len(format) == 0 {
		return format
	}

	return it.UsingMapOptions(
		true,
		format,
		mapToReplace)

}

func (it *replaceTemplate) CurlyTwo(
	format string, // {key}-text...
	firstKey string, firstValue interface{},
	secondKey string, secondValue interface{},
) string {
	if len(format) == 0 {
		return format
	}

	return it.UsingMapOptions(
		true,
		format,
		map[string]string{
			firstKey: fmt.Sprintf(
				constants.SprintValueFormat,
				firstValue),
			secondKey: fmt.Sprintf(
				constants.SprintValueFormat,
				secondValue),
		})
}

func (it *replaceTemplate) DirectOne(
	format string, // key-text...
	firstKey string, firstValue interface{},
) string {
	if len(format) == 0 {
		return format
	}

	return strings.ReplaceAll(
		format,
		firstKey,
		fmt.Sprintf(
			constants.SprintValueFormat,
			firstValue))
}

func (it *replaceTemplate) DirectTwoItem(
	format string, // key-text...
	firstKey string, firstValue interface{},
	secondKey string, secondValue interface{},
) string {
	if len(format) == 0 {
		return format
	}

	return it.UsingMapOptions(
		false,
		format,
		map[string]string{
			firstKey: fmt.Sprintf(
				constants.SprintValueFormat,
				firstValue),
			secondKey: fmt.Sprintf(
				constants.SprintValueFormat,
				secondValue),
		})
}

func (it *replaceTemplate) CurlyTwoItem(
	format string, // {key}-text...
	firstKey string, firstValue interface{},
	secondKey string, secondValue interface{},
) string {
	if len(format) == 0 {
		return format
	}

	return it.UsingMapOptions(
		true,
		format,
		map[string]string{
			firstKey: fmt.Sprintf(
				constants.SprintValueFormat,
				firstValue),
			secondKey: fmt.Sprintf(
				constants.SprintValueFormat,
				secondValue),
		})
}

func (it *replaceTemplate) DirectKeyUsingMap(
	format string, // key-text...
	mapToReplace map[string]string,
) string {
	if len(mapToReplace) == 0 || len(format) == 0 {
		return format
	}

	return it.UsingMapOptions(
		false,
		format,
		mapToReplace)
}

func (it *replaceTemplate) CurlyKeyUsingMap(
	format string, // {key}-text...
	mapToReplace map[string]string,
) string {
	if len(mapToReplace) == 0 || len(format) == 0 {
		return format
	}

	return it.UsingMapOptions(
		true,
		format,
		mapToReplace)
}

func (it *replaceTemplate) UsingMapOptions(
	isConvKeysToCurlyBraceKeys bool, // conv key to {key} before replace
	format string, // {key}-text...
	mapToReplace map[string]string,
) string {
	if len(mapToReplace) == 0 || len(format) == 0 {
		return format
	}

	if isConvKeysToCurlyBraceKeys {
		for key, valueToReplace := range mapToReplace {
			keyCurly := fmt.Sprintf(
				constants.CurlyWrapFormat,
				key)

			format = strings.ReplaceAll(
				format,
				keyCurly,
				valueToReplace)
		}

		return format
	}

	for key, valueToReplace := range mapToReplace {
		format = strings.ReplaceAll(
			format,
			key,
			valueToReplace)
	}

	return format
}
