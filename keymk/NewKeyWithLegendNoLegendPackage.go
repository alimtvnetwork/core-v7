package keymk

import "gitlab.com/evatix-go/core/constants"

// NewKeyWithLegendNoLegendPackage
//
// Chain Sequence (Root-Group-User-Item)
func NewKeyWithLegendNoLegendPackage(
	isAttachLegend bool,
	option *Option,
	rootName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		rootName:            rootName,
		packageName:         constants.EmptyString,
		groupName:           group,
		isAttachLegendNames: isAttachLegend,
	}

	return keyWithLegend
}
