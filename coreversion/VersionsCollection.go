package coreversion

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface"
)

type VersionsCollection struct {
	Versions []Version
}

func (it *VersionsCollection) Add(
	version string,
) *VersionsCollection {
	it.Versions = append(
		it.Versions,
		New.Create(version))

	return it
}

func (it *VersionsCollection) AddSkipInvalid(
	version string,
) *VersionsCollection {
	v := New.Create(version)

	if v.IsEmptyOrInvalid() {
		return it
	}

	it.Versions = append(
		it.Versions,
		v)

	return it
}

func (it *VersionsCollection) AddVersionsRaw(
	versions ...string,
) *VersionsCollection {
	for _, v := range versions {
		it.Versions = append(
			it.Versions,
			New.Create(v))
	}

	return it
}

func (it *VersionsCollection) AddVersions(
	versions ...Version,
) *VersionsCollection {
	for _, v := range versions {
		it.Versions = append(
			it.Versions, v)
	}

	return it
}

func (it *VersionsCollection) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Versions)
}

func (it *VersionsCollection) Count() int {
	return it.Length()
}

func (it *VersionsCollection) IsEmpty() bool {
	return it.Length() == 0
}

func (it *VersionsCollection) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *VersionsCollection) LastIndex() int {
	return it.Length() - 1
}

func (it *VersionsCollection) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *VersionsCollection) VersionCompactStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, version := range it.Versions {
		slice[i] = version.VersionCompact
	}

	return slice
}

func (it *VersionsCollection) VersionsStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	slice := make([]string, it.Length())

	for i, version := range it.Versions {
		slice[i] = version.VersionDisplay()
	}

	return slice
}

func (it *VersionsCollection) IndexOf(
	versionString string,
) int {
	lookupVersion := New.Create(versionString)

	for i, version := range it.Versions {
		if version.VersionCompact == lookupVersion.VersionCompact {
			return i
		}
	}

	return constants.InvalidValue
}

func (it *VersionsCollection) IsContainsVersion(
	versionString string,
) bool {
	return it.IndexOf(versionString) > constants.InvalidValue
}

func (it *VersionsCollection) IsEqual(
	another *VersionsCollection,
) bool {
	if it == nil && another == nil {
		return true
	}

	if it == nil || another == nil {
		return false
	}

	if it.Length() != another.Length() {
		return false
	}

	for i, version := range it.Versions {
		anotherV := another.Versions[i]

		if version.IsVersionCompareNotEqual(anotherV.VersionCompact) {
			return false
		}
	}

	return true
}

func (it *VersionsCollection) String() string {
	return strings.Join(it.VersionsStrings(), constants.NewLineUnix)
}

func (it *VersionsCollection) Json() corejson.Result {
	return corejson.New(it)
}

func (it *VersionsCollection) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *VersionsCollection) JsonParseSelfInject(jsonResult *corejson.Result) error {
	return jsonResult.Deserialize(it)
}

func (it *VersionsCollection) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *VersionsCollection) AsBasicSliceContractsBinder() coreinterface.BasicSlicerContractsBinder {
	return it
}
