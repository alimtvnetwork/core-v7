package stringslice

import "sync"

func ProcessAsync(
	processor func(index int, item interface{}) string,
	items ...interface{},
) []string {
	if len(items) == 0 {
		return []string{}
	}

	list := make([]string, len(items))

	wg := sync.WaitGroup{}

	singleProcessorFunc := func(index int, item interface{}) {
		list[index] = processor(index, item)

		wg.Done()
	}

	wg.Add(len(items))
	for i, item := range items {
		go singleProcessorFunc(i, item)
	}

	wg.Wait()

	return list
}
