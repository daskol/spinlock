// +build amd64

package spinlock

//go:nosplit

// Unlock unlocks spinlock s.
func (s *SpinlockInAsm) Unlock() {
	Unlock(&s.uint32)
}

//go:noescape
//go:nosplit

// Lock is a manual implementation of locking with CAS idiom in assembly.
func Lock(flag *uint32)

//go:noescape
//go:nosplit

// Unlock is a manual implementation of unlocking with CAS idiom in assembly.
func Unlock(flag *uint32)

// SpinlockInAsm wraps native Lock() and Unlock(). This is a bit slower but
// implements Locker interface.
type SpinlockInAsm struct {
	uint32
}

//go:nosplit

// Lock locks spinlock s. If the lock is already in use, the calling goroutine
// blocks until the spinlock is avaliable.
func (s *SpinlockInAsm) Lock() {
	Lock(&s.uint32)
}

// Spinlock is an alias for spinlock implementation in assembly.
type Spinlock = SpinlockInAsm
