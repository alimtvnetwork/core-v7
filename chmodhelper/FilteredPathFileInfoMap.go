package chmodhelper

import (
	"os"
	"strings"

	"gitlab.com/auk-go/core/constants"
)

type FilteredPathFileInfoMap struct {
	lazyValidLocations                   []string
	lazyValidLocationFileInfoRwxWrappers []LocationFileInfoRwxWrapper
	FilesToInfoMap                       map[string]os.FileInfo
	PathsRemainsUnProcessed              []string // TODO
	MissingOrOtherPathIssues             []string
	Error                                error
}

func (it *FilteredPathFileInfoMap) LazyValidLocationFileInfoRwxWrappers() []LocationFileInfoRwxWrapper {
	if it.lazyValidLocationFileInfoRwxWrappers != nil {
		return it.lazyValidLocationFileInfoRwxWrappers
	}

	it.lazyValidLocationFileInfoRwxWrappers =
		it.ValidLocationFileInfoRwxWrappers()

	return it.lazyValidLocationFileInfoRwxWrappers
}

func (it *FilteredPathFileInfoMap) LazyValidLocations() []string {
	if it.lazyValidLocations != nil {
		return it.lazyValidLocations
	}

	it.lazyValidLocations = it.ValidLocations()

	return it.lazyValidLocations
}

func (it *FilteredPathFileInfoMap) ValidLocations() []string {
	length := len(it.FilesToInfoMap)
	slice := make([]string, length)

	if length == 0 {
		return slice
	}

	index := 0
	for s := range it.FilesToInfoMap {
		slice[index] = s
		index++
	}

	return slice
}

func (it *FilteredPathFileInfoMap) ValidFileInfos() []os.FileInfo {
	length := len(it.FilesToInfoMap)
	slice := make([]os.FileInfo, length)

	if length == 0 {
		return slice
	}

	index := 0
	for _, fileInfo := range it.FilesToInfoMap {
		slice[index] = fileInfo
		index++
	}

	return slice
}

func (it *FilteredPathFileInfoMap) ValidLocationFileInfoRwxWrappers() []LocationFileInfoRwxWrapper {
	length := len(it.FilesToInfoMap)
	slice := make([]LocationFileInfoRwxWrapper, length)

	if length == 0 {
		return slice
	}

	index := 0
	for location, fileInfo := range it.FilesToInfoMap {
		slice[index] = LocationFileInfoRwxWrapper{
			Location:   location,
			FileInfo:   fileInfo,
			RwxWrapper: New.RwxWrapper.UsingFileModePtr(fileInfo.Mode()),
		}

		index++
	}

	return slice
}

func InvalidFilteredPathFileInfoMap() *FilteredPathFileInfoMap {
	return &FilteredPathFileInfoMap{
		FilesToInfoMap:           map[string]os.FileInfo{},
		MissingOrOtherPathIssues: []string{},
		Error:                    nil,
	}
}

func (it *FilteredPathFileInfoMap) HasAnyValidFileInfo() bool {
	return len(it.FilesToInfoMap) > 0
}

func (it *FilteredPathFileInfoMap) IsEmptyValidFileInfos() bool {
	return len(it.FilesToInfoMap) == 0
}

func (it *FilteredPathFileInfoMap) LengthOfIssues() int {
	return len(it.MissingOrOtherPathIssues)
}

func (it *FilteredPathFileInfoMap) IsEmptyIssues() bool {
	return len(it.MissingOrOtherPathIssues) == 0
}

func (it *FilteredPathFileInfoMap) HasAnyIssues() bool {
	return it.Error != nil || it.HasAnyMissingPaths()
}

func (it *FilteredPathFileInfoMap) HasError() bool {
	return it.Error != nil
}

func (it *FilteredPathFileInfoMap) HasAnyMissingPaths() bool {
	return len(it.MissingOrOtherPathIssues) > 0
}

func (it *FilteredPathFileInfoMap) MissingPathsToString() string {
	return strings.Join(it.MissingOrOtherPathIssues, constants.CommaSpace)
}
