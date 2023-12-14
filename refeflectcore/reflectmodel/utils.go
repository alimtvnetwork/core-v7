package reflectmodel

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"gitlab.com/auk-go/core/internal/convertinteranl"
)

type rvUtils struct{}

func (it rvUtils) ArgsToReflectValues(args []interface{}) []reflect.Value {
	if len(args) == 0 {
		return []reflect.Value{}
	}

	list := make(
		[]reflect.Value,
		len(args),
	)

	for i, arg := range args {
		list[i] = reflect.ValueOf(arg)
	}

	return list
}

func (it rvUtils) ReflectValuesToInterfaces(
	reflectValues []reflect.Value,
) []interface{} {
	if len(reflectValues) == 0 {
		return []interface{}{}
	}

	list := make(
		[]interface{},
		len(reflectValues),
	)

	for i, rv := range reflectValues {
		list[i] = it.ReflectValueToAnyValue(rv)
	}

	return list
}

func (it rvUtils) ReflectValueToAnyValue(rv reflect.Value) interface{} {
	if it.IsNull(rv) {
		return nil
	}

	k := rv.Kind()

	switch k {
	case reflect.Ptr, reflect.Interface:
		return rv.Elem().Interface()
	default:
		return rv.Interface()
	}
}

func (it rvUtils) IsNull(item interface{}) bool {
	if item == nil {
		return true
	}

	rv := reflect.ValueOf(item)

	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

func (it rvUtils) InterfacesToTypesNamesWithValues(items []interface{}) []string {
	if len(items) == 0 {
		return []string{}
	}

	var output []string

	for i, item := range items {
		toType := reflect.TypeOf(item)
		compiledString := fmt.Sprintf(
			"%d. %s [value: %s]",
			i,
			toType.Name(),
			convertinteranl.AnyTo.SmartString(item),
		)

		output = append(output, compiledString)

	}

	return output
}

func (it rvUtils) InterfacesToTypes(items []interface{}) []reflect.Type {
	if len(items) == 0 {
		return []reflect.Type{}
	}

	var output []reflect.Type

	for _, item := range items {
		toType := reflect.TypeOf(item)
		output = append(output, toType)
	}

	return output
}

func (it rvUtils) errorMessageForTypeVerification(i int, errFirst error) string {
	return fmt.Sprintf(
		"- Index {%d}, %s args : %s",
		i,
		it.IndexToPosition(i),
		errFirst.Error(),
	)
}

func (it rvUtils) VerifyReflectTypes(
	rootName string,
	expectedArgs,
	givenArgs []reflect.Type,
) (isOkay bool, err error) {
	leftLen := len(expectedArgs)
	rightLen := len(givenArgs)

	if leftLen != rightLen {
		errMsg := fmt.Sprintf(
			"Expected Length (%d) != (%d) Given Length",
			leftLen,
			rightLen,
		)

		return false, errors.New(errMsg)
	}

	var errLines []string

	for i := 0; i < leftLen; i++ {
		expected := expectedArgs[i]
		given := givenArgs[i]

		isCurrTypeOkay, errFirst := it.IsReflectTypeMatch(expected, given)

		if isCurrTypeOkay {
			continue
		}

		if errFirst != nil {
			errLines = append(
				errLines,
				it.errorMessageForTypeVerification(i, errFirst),
			)
		}
	}

	if len(errLines) == 0 {
		return true, nil
	}

	finalErrMessages := it.PrependWithSpaces(
		4,
		errLines,
		0,
		fmt.Sprintf("%s =>", rootName),
	)

	finalMsg := strings.Join(finalErrMessages, "\n")

	return false, errors.New(finalMsg)
}

func (it rvUtils) IsReflectTypeMatch(expectedType, givenType reflect.Type) (isOkay bool, err error) {
	if expectedType == givenType {
		return true, nil
	}

	errMsg := fmt.Sprintf(
		"Expected Type (%s) != (%s) Given Type",
		expectedType.Name(),
		givenType.Name(),
	)

	return false, errors.New(errMsg)
}

func (it rvUtils) IsReflectTypeMatchAny(expected, given interface{}) (isOkay bool, err error) {
	ex := reflect.TypeOf(expected)
	gi := reflect.TypeOf(given)

	return it.IsReflectTypeMatch(ex, gi)
}

// IndexToPosition
//
// Index to position format
//
//   - Index 0 => Position => 1st
//   - Index 1 => Position => 2nd
//   - Index 3 => Position => 3rd
//   - Index Rest => Position => %dth
func (it rvUtils) IndexToPosition(
	index int,
) string {
	position := index + 1

	switch position {
	case 1:
		return "1st"
	case 2:
		return "2nd"
	case 3:
		return "3rd"
	default:
		return fmt.Sprintf(
			"%dth",
			position,
		)
	}
}

func (it rvUtils) PrependWithSpaces(
	spaceCountLines int,
	existingLines []string,
	prependingLinesSpaceCount int,
	prependingLines ...string,
) []string {
	var newSlice []string

	if prependingLinesSpaceCount > 0 {
		prependingLines = it.
			WithSpaces(
				prependingLinesSpaceCount,
				prependingLines...,
			)
	}

	newSlice = append(
		newSlice,
		prependingLines...,
	)

	if spaceCountLines > 0 {
		existingLines = it.
			WithSpaces(
				spaceCountLines,
				existingLines...,
			)
	}

	newSlice = append(newSlice, existingLines...)

	return newSlice
}

func (it rvUtils) WithSpaces(spaceCount int, lines ...string) []string {
	if len(lines) == 0 {
		return []string{}
	}

	newLines := make([]string, len(lines))
	prefix := strings.Repeat(
		" ",
		spaceCount,
	)

	for i, line := range lines {
		newLines[i] = fmt.Sprintf(
			"%s%s",
			prefix,
			line,
		)
	}

	return newLines
}
