package convertinteranl

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/internal/jsoninternal"
)

type anyTo struct{}

// ValueString
//
// If nil then returns ""
// Or, returns %v of the value given.
func (it anyTo) ValueString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem,
	)
}

func (it anyTo) FullPropertyString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintFullPropertyNameValueFormat,
		anyItem,
	)
}

func (it anyTo) TypeName(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	return fmt.Sprintf(
		constants.SprintTypeFormat,
		anyItem,
	)
}

// SmartString
//
//   - If nil return ""
//   - If string return just returns
//   - Or, else return %v of value
func (it anyTo) SmartString(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	toStr, isSuccess := anyItem.(string)

	if isSuccess {
		return toStr
	}

	toNamer, isNamer := anyItem.(Namer)

	if isNamer {
		return toNamer.Name()
	}

	toStringer, isStringer := anyItem.(fmt.Stringer)

	if isStringer {
		return toStringer.String()
	}

	return fmt.Sprintf(
		constants.SprintValueFormat,
		anyItem,
	)
}

func (it anyTo) SmartJson(anyItem interface{}) string {
	if anyItem == nil {
		return ""
	}

	switch v := anyItem.(type) {
	case []string:
		return strings.Join(
			v,
			constants.NewLineUnix,
		)
	case string:
		return v
	case int, int32, byte, int64, float64, float32, bool:
		return it.SmartString(v)
	case error:
		if v == nil {
			return ""
		}

		return v.Error()
	default:
		toPrettyJson := jsoninternal.Pretty.
			AnyTo.
			SafeString(anyItem)

		return toPrettyJson
	}
}

func (it anyTo) SmartPrettyJsonLines(anyItem interface{}) []string {
	if anyItem == nil {
		return []string{}
	}

	switch v := anyItem.(type) {
	case []string:
		return v
	case string:
		return strings.Split(
			v,
			constants.NewLineUnix,
		)

	default:
		return it.PrettyJsonLines(anyItem)
	}
}

func (it anyTo) PrettyJsonLines(anyItem interface{}) []string {
	if anyItem == nil {
		return []string{}
	}

	toPrettyJson := jsoninternal.Pretty.
		AnyTo.
		PrettyStringDefault(anyItem)

	return strings.Split(
		toPrettyJson,
		constants.NewLineUnix,
	)
}

// Strings
//
//	This function will display complex objects to simpler form
//	for the integration testing validation and expectations.
//
// # Steps:
//  01. string to []string
//  02. []string to as is.
//  03. []interface{} to []string
//  04. map[string]interface{} (fmt - "%s : SmartJson(%s)") to []string
//  05. map[interface{}]interface{} (fmt - SmartJson("%s) : SmartJson(%s)") to []string
//  06. map[string]string (fmt - %s : %s)") to []string
//  07. map[string]int (fmt - %s : %d)") to []string
//  08. map[int]string (fmt - %d : %s)") to []string
//  09. int to []string
//  10. byte to []string
//  11. bool to []string
//  12. any to PrettyJSON
func (it anyTo) Strings(
	any interface{},
) []string {
	switch v := any.(type) {
	case string:
		if v == "" {
			return []string{}
		}

		return strings.Split(
			v,
			constants.NewLineUnix,
		)
	case error:
		if v == nil {
			return []string{}
		}

		return strings.Split(
			v.Error(),
			constants.NewLineUnix,
		)
	case []string:
		return v
	case []interface{}:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))

		for i, line := range v {
			lines[i] = it.SmartJson(line)
		}

		return lines

	case map[string]interface{}:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %s",
				key,
				it.SmartJson(value),
			)

			index++
		}

		sort.Strings(lines)

		return lines
	case map[interface{}]interface{}:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %s",
				it.SmartJson(key),
				it.SmartJson(value),
			)

			index++
		}

		sort.Strings(lines)

		return lines
	case map[string]string:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %s",
				key,
				value,
			)

			index++
		}

		sort.Strings(lines)

		return lines
	case map[string]int:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %d",
				key,
				value,
			)

			index++
		}

		sort.Strings(lines)

		return lines

	case map[int]string:
		if len(v) == 0 {
			return []string{}
		}

		lines := make([]string, len(v))
		index := 0

		for key, value := range v {
			lines[index] = fmt.Sprintf(
				"%s : %s",
				it.ValueString(key),
				value,
			)

			index++
		}

		sort.Strings(lines)

		return lines
	case int:
		return []string{
			strconv.Itoa(v),
		}
	case byte:
		return []string{
			strconv.Itoa(int(v)),
		}
	case bool:
		return []string{
			strconv.FormatBool(v),
		}
	default:
		return it.PrettyJsonLines(any)
	}
}

// String
//
//	This function will display complex objects to simpler form
//	for the integration testing validation and expectations.
//
// # Steps:
//  01. string to []string
//  02. []string to as is.
//  03. []interface{} to []string
//  04. map[string]interface{} (fmt - "%s : SmartJson(%s)") to []string
//  05. map[interface{}]interface{} (fmt - SmartJson("%s) : SmartJson(%s)") to []string
//  06. map[string]string (fmt - %s : %s)") to []string
//  07. map[string]int (fmt - %s : %d)") to []string
//  08. map[int]string (fmt - %d : %s)") to []string
//  09. int to []string
//  10. byte to []string
//  11. bool to []string
//  12. any to PrettyJSON
func (it anyTo) String(
	any interface{},
) string {
	v := it.Strings(any)

	return strings.Join(v, constants.NewLineUnix)
}
