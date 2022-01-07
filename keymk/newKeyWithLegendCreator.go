package keymk

import "gitlab.com/evatix-go/core/constants"

type newKeyWithLegendCreator struct{}

// All
//
// Chain Sequence (Root-Package-Group-State-User-Item)
func (it *newKeyWithLegendCreator) All(
	option *Option,
	legendName LegendName,
	isAttachLegendNames bool,
	rootName,
	packageName,
	group,
	stateName string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          legendName,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		stateName:           stateName,
		isAttachLegendNames: isAttachLegendNames,
	}

	return keyWithLegend
}

// Create
//
// Chain Sequence (Root-Package-Group-State-User-Item)
func (it *newKeyWithLegendCreator) Create(
	option *Option,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          FullLegends,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: true,
	}

	return keyWithLegend
}

// NoLegend
//
// Chain Sequence (Root-Package-Group-State-User-Item)
func (it *newKeyWithLegendCreator) NoLegend(
	option *Option,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: false,
	}

	return keyWithLegend
}

// NoLegendPackage
//
// Chain Sequence (Root-Group-State-User-Item)
func (it *newKeyWithLegendCreator) NoLegendPackage(
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

func (it *newKeyWithLegendCreator) ShortLegend(
	option *Option,
	rootName,
	packageName,
	group string,
) *KeyWithLegend {
	keyWithLegend := &KeyWithLegend{
		option:              option,
		LegendName:          ShortLegends,
		rootName:            rootName,
		packageName:         packageName,
		groupName:           group,
		isAttachLegendNames: true,
	}

	return keyWithLegend
}
