package chmodhelper

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
)

type RwxInstructionExecutors struct {
	items *[]*RwxInstructionExecutor
}

func NewRwxInstructionExecutors(capacity int) *RwxInstructionExecutors {
	slice := make([]*RwxInstructionExecutor, constants.Zero, capacity)

	return &RwxInstructionExecutors{
		items: &slice,
	}
}

// Add skips nil
func (it *RwxInstructionExecutors) Add(
	rwxInstructionExecutor *RwxInstructionExecutor,
) *RwxInstructionExecutors {
	if rwxInstructionExecutor == nil {
		return it
	}

	*it.items = append(*it.items, rwxInstructionExecutor)

	return it
}

// Adds skips nil
func (it *RwxInstructionExecutors) Adds(
	rwxInstructionExecutors ...*RwxInstructionExecutor,
) *RwxInstructionExecutors {
	if rwxInstructionExecutors == nil {
		return it
	}

	items := *it.items

	for _, executor := range rwxInstructionExecutors {
		if executor == nil {
			continue
		}

		items = append(items, executor)
	}

	*it.items = items

	return it
}

func (it *RwxInstructionExecutors) Length() int {
	if it.items == nil {
		return constants.Zero
	}

	return len(*it.items)
}

func (it *RwxInstructionExecutors) Count() int {
	return it.Length()
}

func (it *RwxInstructionExecutors) IsEmpty() bool {
	return it.Length() == 0
}

func (it *RwxInstructionExecutors) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *RwxInstructionExecutors) LastIndex() int {
	return it.Length() - 1
}

func (it *RwxInstructionExecutors) HasIndex(index int) bool {
	return it.LastIndex() >= index
}

func (it *RwxInstructionExecutors) VerifyRwxModifiers(
	isContinueOnErr,
	isRecursiveIgnore bool,
	locations []string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if isContinueOnErr {
		return it.verifyChmodErrorContinueOnErr(
			isRecursiveIgnore,
			locations)
	}

	for _, executor := range *it.items {
		err := executor.VerifyRwxModifiers(
			isRecursiveIgnore,
			locations)

		if err != nil {
			return err
		}
	}

	return nil
}

func (it *RwxInstructionExecutors) verifyChmodErrorContinueOnErr(
	isRecursiveIgnore bool,
	locations []string,
) error {
	var sliceErr []string

	for _, executor := range *it.items {
		err := executor.VerifyRwxModifiers(
			isRecursiveIgnore,
			locations)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error())
		}
	}

	return errcore.SliceToError(sliceErr)
}

func (it *RwxInstructionExecutors) Items() *[]*RwxInstructionExecutor {
	return it.items
}

func (it *RwxInstructionExecutors) ApplyOnPath(location string) error {
	if it.IsEmpty() {
		return nil
	}

	for _, executor := range *it.items {
		err := executor.ApplyOnPath(location)

		if err != nil {
			return err

		}
	}

	return nil
}

func (it *RwxInstructionExecutors) ApplyOnPaths(locations []string) error {
	if len(locations) == 0 {
		return nil
	}

	return it.ApplyOnPathsPtr(locations)
}

func (it *RwxInstructionExecutors) ApplyOnPathsPtr(locations []string) error {
	if it.IsEmpty() {
		return nil
	}

	for _, executor := range *it.items {
		err := executor.ApplyOnPathsPtr(&locations)

		if err != nil {
			return err
		}
	}

	return nil
}
