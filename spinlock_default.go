package spinlock

import (
	"runtime"
	"sync/atomic"
)

// SpinlockInGo is a spinlock implementation in pure Go.
type SpinlockInGo struct {
	uint32
}

func (s *SpinlockInGo) get() *uint32 {
	return &s.uint32
}

// Lock locks spinlock s.
func (s *SpinlockInGo) Lock() {
	for !atomic.CompareAndSwapUint32(s.get(), 0, 1) {
		runtime.Gosched()
	}
}

// Unlock unlocks spinlock s.
func (s *SpinlockInGo) Unlock() {
	atomic.StoreUint32(s.get(), 0)
}
