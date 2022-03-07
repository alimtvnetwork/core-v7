package baseactioninf

import "gitlab.com/evatix-go/core/coreinterface/enuminf"

type CategoryTypeNamer interface {
	TypeName() enuminf.BasicEnumer
	Category() enuminf.BasicEnumer
}
