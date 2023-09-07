package mutexbykey

import (
	"sync"

	"gitlab.com/auk-go/core/constants"
)

type mutexMap struct {
	items map[string]*sync.Mutex
}

var globalMutex = sync.Mutex{}

var items = make(
	map[string]*sync.Mutex,
	constants.ArbitraryCapacity10)

var internalMap = mutexMap{
	items: items,
}

func Get(key string) *sync.Mutex {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	mutex, has := internalMap.items[key]

	if has {
		return mutex
	}

	// not there
	newMutex := &sync.Mutex{}
	internalMap.items[key] = newMutex

	return newMutex
}

func Delete(key string) bool {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	_, has := internalMap.items[key]

	if has {
		delete(internalMap.items, key)
	}

	return has
}
