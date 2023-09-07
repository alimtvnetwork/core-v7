package converters

import (
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

type StringsToMapConverter []string

func (it StringsToMapConverter) SafeStrings() []string {
	if it.IsEmpty() {
		return []string{}
	}

	return it
}

func (it StringsToMapConverter) Strings() []string {
	return it
}

func (it *StringsToMapConverter) Length() int {
	if it == nil || *it == nil {
		return 0
	}

	return len(*it)
}

func (it *StringsToMapConverter) IsEmpty() bool {
	return it.Length() == 0
}

func (it *StringsToMapConverter) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *StringsToMapConverter) LastIndex() int {
	return it.Length() - 1
}

func (it StringsToMapConverter) LineSplitMapOptions(
	isTrim bool,
	splitter string,
) map[string]string {
	if isTrim {
		return it.LineSplitMapTrim(splitter)
	}

	return it.LineSplitMap(splitter)
}

func (it StringsToMapConverter) LineProcessorMapOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key, val string),
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineProcessorMapOptions(
			isTrimBefore,
			processorFunc)
}

func (it StringsToMapConverter) LineProcessorMapStringIntegerTrim(
	processorFunc func(line string) (key string, val int),
) map[string]int {
	return it.LineProcessorMapStringIntegerOptions(
		true,
		processorFunc)
}

func (it StringsToMapConverter) LineProcessorMapStringIntegerOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key string, val int),
) map[string]int {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineProcessorMapStringIntegerOptions(
			isTrimBefore,
			processorFunc)
}

func (it StringsToMapConverter) LineProcessorMapStringAnyTrim(
	processorFunc func(line string) (key string, val interface{}),
) map[string]interface{} {
	return it.LineProcessorMapStringAnyOptions(
		true,
		processorFunc)
}

func (it StringsToMapConverter) LineProcessorMapStringAnyOptions(
	isTrimBefore bool,
	processorFunc func(line string) (key string, val interface{}),
) map[string]interface{} {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineProcessorMapStringAnyOptions(
			isTrimBefore,
			processorFunc)
}

func (it StringsToMapConverter) LineSplitMapTrim(
	splitter string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineSplitMapTrim(
			splitter,
		)
}

func (it StringsToMapConverter) LineSplitMap(
	splitter string,
) map[string]string {
	return strutilinternal.
		SliceToMapConverter(it.Strings()).
		LineSplitMap(
			splitter,
		)
}
