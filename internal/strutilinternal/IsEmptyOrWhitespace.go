package strutilinternal

import "strings"

func IsEmptyOrWhitespace(str string) bool {
	return str == "" || str == " " || str == "\n" || strings.TrimSpace(str) == ""
}
