package codestack

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

func MethodNamePackageName(fullFuncName string) (fullMethodName, packageName, methodName string) {
	if fullFuncName == "" {
		return "", "", ""
	}

	hasComplexName := strings.HasPrefix(
		fullFuncName,
		gitlabDotCom) ||
		strings.ContainsRune(
			fullFuncName,
			constants.ForwardRune)

	if hasComplexName {
		forwardSlashFound := strings.LastIndexByte(
			fullFuncName,
			constants.ForwardChar)

		return MethodNamePackageName(fullFuncName[forwardSlashFound+1:])
	}

	splitsByDot := strings.Split(fullFuncName, constants.Dot)
	packageName, methodName = stringslice.FirstLastDefault(splitsByDot)

	return fullFuncName, packageName, methodName
}
