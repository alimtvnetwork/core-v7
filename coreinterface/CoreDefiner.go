package coreinterface

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/corejson"
)

type CoreDefiner interface {
	corejson.Jsoner
	AllSerializer

	fmt.Stringer
}
