package strutilinternal

import (
	"strings"
)

type SliceToMapConverter []string

func (it SliceToMapConverter) SafeStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return it
}

func (it SliceToMapConverter) Strings() []string {
	return it
}

func (it SliceToMapConverter) Hashset() map[string]bool {
	length := it.Length()
	hashset := make(map[string]bool, length)

	for _, s := range it {
		hashset[s] = true
	}

	return hashset
}

func (it *SliceToMapConverter) Length() int {
	if it == nil || *it == nil {
		return 0
	}

	return len(*it)
}

func (it *SliceToMapConverter) IsEmpty() bool {
	return it.Length() == 0
}

func (it *SliceToMapConverter) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *SliceToMapConverter) LastIndex() int {
	return it.Length() - 1
}

func (it SliceToMapConverter) LineSplitMapOptions(
	isTrim bool,
	splitter string,
) map[string]string {
	if isTrim {
		return it.LineSplitMapTrim(splitter)
	}

	return it.LineSplitMap(splitter)
}

func (it SliceToMapConverter) LineProcessorMapOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key, val string),
) map[string]string {
	length := it.Length()
	if processorFunc == nil || length == 0 {
		return map[string]string{}
	}

	newMap := make(map[string]string, length+1)

	if isTrimBefore {
		for _, line := range it {
			trimmedLine := strings.TrimSpace(line)

			if trimmedLine == "" {
				continue
			}

			k, v := processorFunc(line)
			newMap[k] = v
		}

		return newMap
	}

	for _, line := range it {
		k, v := processorFunc(line)
		newMap[k] = v
	}

	return newMap
}

func (it SliceToMapConverter) LineProcessorMapStringIntegerTrim(
	processorFunc func(line string) (key string, val int),
) map[string]int {
	return it.LineProcessorMapStringIntegerOptions(
		true,
		processorFunc)
}

func (it SliceToMapConverter) LineProcessorMapStringIntegerOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key string, val int),
) map[string]int {
	length := it.Length()
	if processorFunc == nil || length == 0 {
		return map[string]int{}
	}

	newMap := make(map[string]int, length+1)

	if isTrimBefore {
		for _, line := range it {
			trimmedLine := strings.TrimSpace(line)

			if trimmedLine == "" {
				continue
			}

			k, v := processorFunc(line)
			newMap[k] = v
		}

		return newMap
	}

	for _, line := range it {
		k, v := processorFunc(line)
		newMap[k] = v
	}

	return newMap
}

func (it SliceToMapConverter) LineProcessorMapStringAnyTrim(
	processorFunc func(line string) (key string, val interface{}),
) map[string]interface{} {
	return it.LineProcessorMapStringAnyOptions(
		true,
		processorFunc)
}

func (it SliceToMapConverter) LineProcessorMapStringAnyOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key string, val interface{}),
) map[string]interface{} {
	length := it.Length()
	if processorFunc == nil || length == 0 {
		return map[string]interface{}{}
	}

	newMap := make(map[string]interface{}, length+1)

	if isTrimBefore {
		for _, line := range it {
			trimmedLine := strings.TrimSpace(line)

			if trimmedLine == "" {
				continue
			}

			k, v := processorFunc(line)
			newMap[k] = v
		}

		return newMap
	}

	for _, line := range it {
		k, v := processorFunc(line)
		newMap[k] = v
	}

	return newMap
}

func (it SliceToMapConverter) LineSplitMapTrim(
	splitter string,
) map[string]string {
	length := it.Length()
	if length == 0 {
		return map[string]string{}
	}

	newMap := make(map[string]string, length+1)

	for _, line := range it {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			continue
		}

		k, v := SplitLeftRightTrim(
			splitter,
			line)

		newMap[k] = v
	}

	return newMap
}

func (it SliceToMapConverter) LineSplitMap(
	splitter string,
) map[string]string {
	length := it.Length()
	if length == 0 {
		return map[string]string{}
	}

	newMap := make(map[string]string, length+1)

	for _, line := range it {
		k, v := SplitLeftRight(
			splitter,
			line)

		newMap[k] = v
	}

	return newMap
}
