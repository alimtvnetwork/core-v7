package coreinterface

import "gitlab.com/evatix-go/core/coredata/corestr"

type CollectionRangesGetter interface {
	RangesCollection() *corestr.Collection
}
