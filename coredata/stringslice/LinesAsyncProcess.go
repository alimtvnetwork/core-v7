package stringslice

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
)

func LinesAsyncProcess(
	lines []string,
	lineProcessor func(index int, lineIn string) (lineOut string),
) []string {
	if len(lines) == 0 {
		return []string{}
	}

	slice := Make(constants.Zero, len(lines))
	wg := sync.WaitGroup{}

	wg.Add(len(lines))

	asyncProcessor := func(index int, lineIn string) {
		slice[index] = lineProcessor(index, lineIn)

		wg.Done()
	}

	for i, lineIn := range lines {
		go asyncProcessor(i, lineIn)
	}

	wg.Wait()

	return slice
}
