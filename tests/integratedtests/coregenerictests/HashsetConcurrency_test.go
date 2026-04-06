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
// Test: Hashset AddLock — concurrent goroutines
// ==========================================

func Test_GenericHashset_AddLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashsetAddLockConcurrencyTestCase
	const goroutines = 500
	hs := coregeneric.NewHashset[int](goroutines)

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			hs.AddLock(idx)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashset AddSliceLock — concurrent batch adds
// ==========================================

func Test_GenericHashset_AddSliceLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashsetAddSliceLockConcurrencyTestCase
	const goroutines = 100
	const batchSize = 10
	hs := coregeneric.NewHashset[string](goroutines * batchSize)

	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			batch := make([]string, batchSize)
			for b := 0; b < batchSize; b++ {
				batch[b] = fmt.Sprintf("g%d-b%d", idx, b)
			}
			hs.AddSliceLock(batch)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashset ContainsLock — concurrent reads with writes
// ==========================================

func Test_GenericHashset_ContainsLock_ConcurrentReadsWrites(t *testing.T) {
	// Arrange
	tc := hashsetContainsLockConcurrencyTestCase
	const writers = 200
	const readers = 200
	hs := coregeneric.NewHashset[int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	for i := 0; i < writers; i++ {
		go func(idx int) {
			hs.AddLock(idx)
			wg.Done()
		}(i)
	}

	for i := 0; i < readers; i++ {
		go func(idx int) {
			_ = hs.ContainsLock(idx) // must not panic
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"finalLength": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashset RemoveLock — concurrent add and remove
// ==========================================

func Test_GenericHashset_RemoveLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashsetRemoveLockConcurrencyTestCase
	const items = 500
	hs := coregeneric.NewHashset[int](items)

	for i := 0; i < items; i++ {
		hs.Add(i)
	}

	wg := sync.WaitGroup{}
	wg.Add(items)

	for i := 0; i < items; i++ {
		go func(idx int) {
			hs.RemoveLock(idx)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashset LengthLock — concurrent reads during mutations
// ==========================================

func Test_GenericHashset_LengthLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashsetLengthLockConcurrencyTestCase
	const writers = 100
	const readers = 100
	hs := coregeneric.NewHashset[int](writers)

	wg := sync.WaitGroup{}
	wg.Add(writers + readers)

	var noNegativeLen atomic.Bool
	noNegativeLen.Store(true)

	for i := 0; i < writers; i++ {
		go func(idx int) {
			hs.AddLock(idx)
			wg.Done()
		}(i)
	}

	for i := 0; i < readers; i++ {
		go func() {
			length := hs.LengthLock()
			if length < 0 {
				noNegativeLen.Store(false)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// Act
	actual := args.Map{
		"finalLength":   hs.Length(),
		"noNegativeLen": noNegativeLen.Load(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Test: Hashset IsEmptyLock — concurrent check
// ==========================================

func Test_GenericHashset_IsEmptyLock_ConcurrentSafety(t *testing.T) {
	// Arrange
	tc := hashsetIsEmptyLockConcurrencyTestCase
	const goroutines = 100
	hs := coregeneric.NewHashset[int](goroutines)

	wg := sync.WaitGroup{}
	wg.Add(goroutines * 2)

	for i := 0; i < goroutines; i++ {
		go func(idx int) {
			hs.AddLock(idx)
			wg.Done()
		}(i)
		go func() {
			_ = hs.IsEmptyLock() // must not panic
			wg.Done()
		}()
	}

	wg.Wait()

	// Act
	actual := args.Map{"length": hs.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
