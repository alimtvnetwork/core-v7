package coredynamic

import (
	"reflect"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func TypeNamesStringUsingReflectType(
	isFullName bool,
	reflectTypes ...reflect.Type,
) string {
	return strings.Join(
		TypeNamesUsingReflectType(isFullName, reflectTypes...),
		constants.CommaSpace)
}
