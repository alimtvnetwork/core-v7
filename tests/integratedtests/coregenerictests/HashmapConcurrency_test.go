package coregenerictests

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coregeneric"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ==========================================
// Test: Hashmap SetLock — concurrent goroutines
// ==========================================

func Test_GenericHashmap_SetLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashmapSetLockConcurrencyTestCase
	const goroutines = 500
	hm := coregeneric.NewHashmap[int, string](goroutines)

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			hm.SetLock(idx, fmt.Sprintf("val-%d", idx))
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": hm.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashmap GetLock — concurrent reads with writes
// ==========================================

func Test_GenericHashmap_GetLock_ConcurrentReadsWrites(t *testing.T) {
	// Arrange
	tc := hashmapGetLockConcurrencyTestCase
	const writers = 200
	const readers = 200
	hm := coregeneric.NewHashmap[int, int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	for i := 0; i < writers; i++ {
		go func(idx int) {
			hm.SetLock(idx, idx*10)
			wg.Done()
		}(i)
	}

	for i := 0; i < readers; i++ {
		go func(idx int) {
			_, _ = hm.GetLock(idx) // must not panic
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"finalLength": hm.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashmap ContainsLock — concurrent reads
// ==========================================

func Test_GenericHashmap_ContainsLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashmapContainsLockConcurrencyTestCase
	const writers = 200
	const readers = 200
	hm := coregeneric.NewHashmap[string, int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	for i := 0; i < writers; i++ {
		go func(idx int) {
			hm.SetLock(fmt.Sprintf("key-%d", idx), idx)
			wg.Done()
		}(i)
	}

	for i := 0; i < readers; i++ {
		go func(idx int) {
			_ = hm.ContainsLock(fmt.Sprintf("key-%d", idx)) // must not panic
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"finalLength": hm.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashmap RemoveLock — concurrent removes
// ==========================================

func Test_GenericHashmap_RemoveLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashmapRemoveLockConcurrencyTestCase
	const items = 500
	hm := coregeneric.NewHashmap[int, string](items)

	for i := 0; i < items; i++ {
		hm.Set(i, fmt.Sprintf("val-%d", i))
	}

	wg := sync.WaitGroup{}
	wg.Add(items)

	for i := 0; i < items; i++ {
		go func(idx int) {
			hm.RemoveLock(idx)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": hm.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashmap LengthLock — concurrent reads during mutations
// ==========================================

func Test_GenericHashmap_LengthLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashmapLengthLockConcurrencyTestCase
	const writers = 100
	const readers = 100
	hm := coregeneric.NewHashmap[int, int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	var noNegativeLen atomic.Bool
	noNegativeLen.Store(true)

	for i := 0; i < writers; i++ {
		go func(idx int) {
			hm.SetLock(idx, idx)
			wg.Done()
		}(i)
	}

	for i := 0; i < readers; i++ {
		go func() {
			length := hm.LengthLock()
			if length < 0 {
				noNegativeLen.Store(false)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// Act
	actual := args.Map{
		"finalLength":   hm.Length(),
		"noNegativeLen": noNegativeLen.Load(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashmap IsEmptyLock — concurrent check
// ==========================================

func Test_GenericHashmap_IsEmptyLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashmapIsEmptyLockConcurrencyTestCase
	const goroutines = 100
	hm := coregeneric.NewHashmap[int, int](goroutines)

	wg := sync.WaitGroup{}
	wg.Add(goroutines * 2)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			hm.SetLock(idx, idx)
			wg.Done()
		}(i)
		go func() {
			_ = hm.IsEmptyLock() // must not panic
			wg.Done()
		}()
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": hm.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashmap mixed SetLock+GetLock+RemoveLock
// ==========================================

func Test_GenericHashmap_MixedOperations_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashmapMixedOpsConcurrencyTestCase
	const items = 300
	hm := coregeneric.NewHashmap[int, string](items)

	for i := 0; i < items/2; i++ {
		hm.Set(i, fmt.Sprintf("initial-%d", i))
	}

	wg := sync.WaitGroup{}
	wg.Add(items * 3)

	for i := items; i < items*2; i++ {
		go func(idx int) {
			hm.SetLock(idx, fmt.Sprintf("new-%d", idx))
			wg.Done()
		}(i)
	}

	for i := 0; i < items; i++ {
		go func(idx int) {
			_, _ = hm.GetLock(idx)
			wg.Done()
		}(i)
	}

	for i := 0; i < items; i++ {
		go func(idx int) {
			if idx < items/2 {
				hm.RemoveLock(idx)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"finalLength": hm.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
