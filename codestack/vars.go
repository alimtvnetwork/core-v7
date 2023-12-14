package codestack

import (
	"gitlab.com/auk-go/core/internal/reflectinternal"
)

var (
	NameOf   = currentNameOf{}
	New      = newCreator{}
	StacksTo = stacksTo{}
	File     = fileGetter{}
	Dir      = dirGetter{}

	getFuncEverything = reflectinternal.GetFunc.All
)
