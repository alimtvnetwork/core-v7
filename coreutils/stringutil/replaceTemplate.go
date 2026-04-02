package stringutil

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/alimtvnetwork/core/constants"
)

type replaceTemplate struct{}

func (it *replaceTemplate) CurlyOne(
	format string, // {key}-text...
	firstKey string, firstValue any,
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
				firstValue,
			),
		},
	)
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
		mapToReplace,
	)

}

func (it *replaceTemplate) CurlyTwo(
	format string, // {key}-text...
	firstKey string, firstValue any,
	secondKey string, secondValue any,
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
				firstValue,
			),
			secondKey: fmt.Sprintf(
				constants.SprintValueFormat,
				secondValue,
			),
		},
	)
}

func (it *replaceTemplate) DirectOne(
	format string, // key-text...
	firstKey string, firstValue any,
) string {
	if len(format) == 0 {
		return format
	}

	return strings.ReplaceAll(
		format,
		firstKey,
		fmt.Sprintf(
			constants.SprintValueFormat,
			firstValue,
		),
	)
}

func (it *replaceTemplate) DirectTwoItem(
	format string, // key-text...
	firstKey string, firstValue any,
	secondKey string, secondValue any,
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
				firstValue,
			),
			secondKey: fmt.Sprintf(
				constants.SprintValueFormat,
				secondValue,
			),
		},
	)
}

func (it *replaceTemplate) CurlyTwoItem(
	format string, // {key}-text...
	firstKey string, firstValue any,
	secondKey string, secondValue any,
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
				firstValue,
			),
			secondKey: fmt.Sprintf(
				constants.SprintValueFormat,
				secondValue,
			),
		},
	)
}

// DirectKeyUsingMap
//
//	Don't use wrap just use the Key and it will be replaced.
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
		mapToReplace,
	)
}

func (it *replaceTemplate) DirectKeyUsingKeyVal(
	format string, // key-text...
	keyValues ...KeyValReplacer,
) string {
	if len(keyValues) == 0 || len(format) == 0 {
		return format
	}

	for _, pair := range keyValues {
		format = strings.ReplaceAll(
			format,
			pair.Key,
			pair.Value,
		)
	}

	return format
}

func (it *replaceTemplate) DirectKeyUsingMapTrim(
	format string, // key-text...
	mapToReplace map[string]string,
) string {
	result := it.DirectKeyUsingMap(format, mapToReplace)

	return strings.TrimSpace(result)
}

// ReplaceWhiteSpaces
//
// Give "  some  nothing    \t" -> "somenothing"
func (it *replaceTemplate) ReplaceWhiteSpaces(
	textContainsWhitespaces string,
) string {
	trimmedText := strings.TrimSpace(textContainsWhitespaces)

	if len(trimmedText) == 0 {
		return trimmedText
	}

	var sb strings.Builder
	sb.Grow(len(trimmedText))

	for _, r := range trimmedText {
		if unicode.IsSpace(r) {
			continue
		}

		sb.WriteRune(r)
	}

	return sb.String()
}

// ReplaceWhiteSpacesToSingle
//
// Give "  some  nothing    \t" -> "some nothing"
func (it *replaceTemplate) ReplaceWhiteSpacesToSingle(
	textContainsWhitespaces string,
) string {
	trimmedText := strings.TrimSpace(textContainsWhitespaces)

	if len(trimmedText) == 0 {
		return trimmedText
	}

	var sb strings.Builder
	sb.Grow(len(trimmedText))
	hasSpaceAlready := false
	isSpace := false

	for _, r := range trimmedText {
		isSpace = unicode.IsSpace(r)
		if r == '\n' || r == '\f' || r == '\t' || r == '\r' || isSpace && hasSpaceAlready {
			continue
		}

		sb.WriteRune(r)

		hasSpaceAlready = isSpace
	}

	return sb.String()
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
		mapToReplace,
	)
}

// UsingMapOptions
//
//	Replaces format template using the map given.
//
//	format : {key}-text-something/{path}...
//
// Options:
//   - isConvKeysToCurlyBraceKeys : true
//     will convert map keys to {key} and then use
//     recursive replace to reduce the template format.
func (it *replaceTemplate) UsingMapOptions(
	isConvKeysToCurlyBraceKeys bool, // true: conv key to {key} before replace
	format string, // Template-format: {key}-text-something/{path}...
	mapToReplace map[string]string,
) string {
	if len(mapToReplace) == 0 || len(format) == 0 {
		return format
	}

	if isConvKeysToCurlyBraceKeys {
		for key, valueToReplace := range mapToReplace {
			keyCurly := fmt.Sprintf(
				constants.CurlyWrapFormat,
				key,
			)

			format = strings.ReplaceAll(
				format,
				keyCurly,
				valueToReplace,
			)
		}

		return format
	}

	for key, valueToReplace := range mapToReplace {
		format = strings.ReplaceAll(
			format,
			key,
			valueToReplace,
		)
	}

	return format
}

func (it *replaceTemplate) UsingNamerMapOptions(
	isConvKeysToCurlyBraceKeys bool, // true: conv key to {key} before replace
	format string, // Template-format: {key}-text-something/{path}...
	mapToReplace map[namer]string,
) string {
	if len(mapToReplace) == 0 || len(format) == 0 {
		return format
	}

	if isConvKeysToCurlyBraceKeys {
		for key, valueToReplace := range mapToReplace {
			keyCurly := fmt.Sprintf(
				constants.CurlyWrapFormat,
				key.Name(),
			)

			format = strings.ReplaceAll(
				format,
				keyCurly,
				valueToReplace,
			)
		}

		return format
	}

	for keyNamer, valueToReplace := range mapToReplace {
		format = strings.ReplaceAll(
			format,
			keyNamer.Name(),
			valueToReplace,
		)
	}

	return format
}

func (it *replaceTemplate) UsingStringerMapOptions(
	isConvKeysToCurlyBraceKeys bool, // true: conv key to {key} before replace
	format string, // Template-format: {key}-text-something/{path}...
	mapToReplace map[fmt.Stringer]string,
) string {
	if len(mapToReplace) == 0 || len(format) == 0 {
		return format
	}

	if isConvKeysToCurlyBraceKeys {
		for key, valueToReplace := range mapToReplace {
			keyCurly := fmt.Sprintf(
				constants.CurlyWrapFormat,
				key,
			)

			format = strings.ReplaceAll(
				format,
				keyCurly,
				valueToReplace,
			)
		}

		return format
	}

	for keyStringer, valueToReplace := range mapToReplace {
		format = strings.ReplaceAll(
			format,
			keyStringer.String(),
			valueToReplace,
		)
	}

	return format
}

// UsingWrappedTemplate
//
//	Replaces all constants.WrappedTemplate {wrapped} with replaceText
//
//	Format:
//	    - "some format {wrapped} text here."
func (it *replaceTemplate) UsingWrappedTemplate(
	format,
	replacingText string,
) string {
	if len(format) == 0 {
		return format
	}

	return strings.ReplaceAll(
		format,
		constants.WrappedTemplate,
		replacingText,
	)
}

// UsingBracketsWrappedTemplate
//
//	Replaces all constants.BracketsWrappedTemplate {brackets-wrapped} with replaceText
//
//	Format:
//	    - "some format {brackets-wrapped} text here."
func (it *replaceTemplate) UsingBracketsWrappedTemplate(
	format,
	replacingText string,
) string {
	if len(format) == 0 {
		return format
	}

	return strings.ReplaceAll(
		format,
		constants.BracketsWrappedTemplate,
		replacingText,
	)
}

// UsingQuotesWrappedTemplate
//
//	Replaces all constants.QuotesWrappedTemplate "{quotes-wrapped}" with replaceText
//
//	Format:
//	    - "some format {quotes-wrapped} text here."
func (it *replaceTemplate) UsingQuotesWrappedTemplate(
	format,
	replacingText string,
) string {
	if len(format) == 0 {
		return format
	}

	return strings.ReplaceAll(
		format,
		constants.QuotesWrappedTemplate,
		replacingText,
	)
}

// UsingValueTemplate
//
//	Replaces all constants.ValueTemplate "{value}" with replaceText
//
//	Format:
//	    - "some format {value} text here."
func (it *replaceTemplate) UsingValueTemplate(
	format,
	replacingText string,
) string {
	if len(format) == 0 {
		return format
	}

	return strings.ReplaceAll(
		format,
		constants.ValueTemplate,
		replacingText,
	)
}

// UsingValueWithFieldsTemplate
//
//	Replaces all constants.ValueWithFieldsTemplate "{value-fields}" with replaceText
//
//	Format:
//	    - "some format {value-fields} text here."
func (it *replaceTemplate) UsingValueWithFieldsTemplate(
	format,
	replacingText string,
) string {
	if len(format) == 0 {
		return format
	}

	return strings.ReplaceAll(
		format,
		constants.ValueWithFieldsTemplate,
		replacingText,
	)
}
