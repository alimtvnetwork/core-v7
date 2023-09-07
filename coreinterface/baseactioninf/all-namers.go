package baseactioninf

import "gitlab.com/auk-go/core/coreinterface/enuminf"

type CategoryTypeNamer interface {
	TypeName() enuminf.BasicEnumer
	Category() enuminf.BasicEnumer
}
