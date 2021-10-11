package chmodhelper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"
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
func (receiver *RwxInstructionExecutors) Add(
	rwxInstructionExecutor *RwxInstructionExecutor,
) *RwxInstructionExecutors {
	if rwxInstructionExecutor == nil {
		return receiver
	}

	*receiver.items = append(*receiver.items, rwxInstructionExecutor)

	return receiver
}

// Adds skips nil
func (receiver *RwxInstructionExecutors) Adds(
	rwxInstructionExecutors ...*RwxInstructionExecutor,
) *RwxInstructionExecutors {
	if rwxInstructionExecutors == nil {
		return receiver
	}

	items := *receiver.items

	for _, executor := range rwxInstructionExecutors {
		items = append(items, executor)
	}

	*receiver.items = items

	return receiver
}

func (receiver *RwxInstructionExecutors) Length() int {
	if receiver.items == nil {
		return constants.Zero
	}

	return len(*receiver.items)
}

func (receiver *RwxInstructionExecutors) Count() int {
	return receiver.Length()
}

func (receiver *RwxInstructionExecutors) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *RwxInstructionExecutors) HasAnyItem() bool {
	return receiver.Length() > 0
}

func (receiver *RwxInstructionExecutors) LastIndex() int {
	return receiver.Length() - 1
}

func (receiver *RwxInstructionExecutors) HasIndex(index int) bool {
	return receiver.LastIndex() >= index
}

func (receiver *RwxInstructionExecutors) VerifyRwxModifiers(
	isContinueOnErr,
	isRecursiveIgnore bool,
	locations []string,
) error {
	if len(locations) == 0 {
		return nil
	}

	if isContinueOnErr {
		return receiver.verifyChmodErrorContinueOnErr(
			isRecursiveIgnore,
			locations)
	}

	for _, executor := range *receiver.items {
		err := executor.VerifyRwxModifiers(
			isRecursiveIgnore,
			locations)

		if err != nil {
			return err
		}
	}

	return nil
}

func (receiver *RwxInstructionExecutors) verifyChmodErrorContinueOnErr(
	isRecursiveIgnore bool,
	locations []string,
) error {
	var sliceErr []string

	for _, executor := range *receiver.items {
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

func (receiver *RwxInstructionExecutors) Items() *[]*RwxInstructionExecutor {
	return receiver.items
}

func (receiver *RwxInstructionExecutors) ApplyOnPath(location string) error {
	if receiver.IsEmpty() {
		return nil
	}

	for _, executor := range *receiver.items {
		err := executor.ApplyOnPath(location)

		if err != nil {
			return err

		}
	}

	return nil
}

func (receiver *RwxInstructionExecutors) ApplyOnPaths(locations []string) error {
	if len(locations) == 0 {
		return nil
	}

	return receiver.ApplyOnPathsPtr(&locations)
}

func (receiver *RwxInstructionExecutors) ApplyOnPathsPtr(locations *[]string) error {
	if receiver.IsEmpty() {
		return nil
	}

	for _, executor := range *receiver.items {
		err := executor.ApplyOnPathsPtr(locations)

		if err != nil {
			return err
		}
	}

	return nil
}
