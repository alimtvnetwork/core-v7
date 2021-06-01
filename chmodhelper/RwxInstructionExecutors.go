package chmodhelper

import "gitlab.com/evatix-go/core/constants"

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

func (receiver *RwxInstructionExecutors) Items() *[]*RwxInstructionExecutor {
	return receiver.items
}

func (receiver *RwxInstructionExecutors) ApplyOnPath(location string) error {
	if receiver.IsEmpty() {
		return nil
	}

	for _, instruction := range *receiver.items {
		err := instruction.ApplyOnPath(location)

		if err != nil {
			return err

		}
	}

	return nil
}

func (receiver *RwxInstructionExecutors) ApplyOnPaths(locations *[]string) error {
	if receiver.IsEmpty() {
		return nil
	}

	for _, instruction := range *receiver.items {
		err := instruction.ApplyOnPaths(locations)

		if err != nil {
			return err
		}
	}

	return nil
}
