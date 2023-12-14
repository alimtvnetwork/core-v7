package reflectinternal

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type reflectUtils struct{}

func (it reflectUtils) MaxLimit(currentLength, maxCapacity int) int {
	if maxCapacity <= -1 {
		return currentLength
	}

	if currentLength >= maxCapacity {
		return maxCapacity
	}

	return currentLength
}

func (it reflectUtils) AppendArgs(appendingItem interface{}, args []interface{}) []interface{} {
	if len(args) == 0 {
		return []interface{}{appendingItem}
	}

	list := make(
		[]interface{},
		len(args)+1,
	)

	list[0] = appendingItem

	for i, arg := range args {
		list[i+1] = reflect.ValueOf(arg)
	}

	return list
}

func (it reflectUtils) VerifyReflectTypesAny(left, right []interface{}) (isOkay bool, err error) {
	leftLen := len(left)
	rightLen := len(right)

	if leftLen != rightLen {
		errMsg := fmt.Sprintf(
			"Left Len(%d) != Right Len (%d)",
			leftLen,
			rightLen,
		)

		return false, errors.New(errMsg)
	}

	var errSlice []string

	for i := 0; i < leftLen; i++ {
		l := left[i]
		r := right[i]

		isCurrTypeOkay, errFirst := it.IsReflectTypeMatchAny(l, r)

		if isCurrTypeOkay {
			continue
		}

		if errFirst != nil {
			errSlice = append(
				errSlice,
				it.errorMessageForTypeVerification(i, errFirst),
			)
		}
	}

	if len(errSlice) == 0 {
		return true, nil
	}

	return false, errors.New(strings.Join(errSlice, "\n"))
}

func (it reflectUtils) errorMessageForTypeVerification(i int, errFirst error) string {
	return fmt.Sprintf(
		"- Index {%d}, %s args : %s",
		i,
		indexToPositionFunc(i),
		errFirst.Error(),
	)
}

func (it reflectUtils) VerifyReflectTypes(
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

	finalErrMessage := prependWithSpacesFunc(
		4,
		errLines,
		0,
		fmt.Sprintf("%s =>", rootName),
	)

	return false, errors.New(finalErrMessage)
}

func (it reflectUtils) PkgNameOnly(
	i interface{},
) string {
	fullName := GetFunc.FullName(i)
	_, pkgName, _ := GetFunc.All(fullName)

	return GetFunc.fixFinalFuncName(pkgName)
}

func (it reflectUtils) FullNameToPkgName(
	fullName string,
) string {
	_, pkgName, _ := GetFunc.All(fullName)

	return GetFunc.fixFinalFuncName(pkgName)
}

func (it reflectUtils) IsReflectTypeMatch(expectedType, givenType reflect.Type) (isOkay bool, err error) {
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

func (it reflectUtils) IsReflectTypeMatchAny(expected, given interface{}) (isOkay bool, err error) {
	ex := reflect.TypeOf(expected)
	gi := reflect.TypeOf(given)

	return it.IsReflectTypeMatch(ex, gi)
}
