package spinlock

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestSpinlockInAsmSimple(t *testing.T) {
	lock := SpinlockInAsm{}
	lock.Lock()
	lock.Unlock()
}

func TestSpinlockInAsmCorrectness(t *testing.T) {
	var lock SpinlockInAsm
	var group sync.WaitGroup
	var begin, end time.Time

	go func() {
		group.Add(1)
		lock.Lock()
		time.Sleep(1500 * time.Millisecond)
		lock.Unlock()
		group.Done()
	}()

	time.Sleep(300 * time.Millisecond)

	go func() {
		group.Add(1)
		begin = time.Now()
		lock.Lock()
		end = time.Now()
		lock.Unlock()
		group.Done()
	}()

	group.Wait()

	var dur = end.Sub(begin)

	if dur < 1100*time.Millisecond || dur > 1300*time.Millisecond {
		t.Errorf("SpinlockInAsm.Lock() is broken: %s", dur)

	}
}

func TestSpinlockInAsmConcurrency(t *testing.T) {
	var counter uint
	var group sync.WaitGroup
	var lock SpinlockInAsm

	for i := 0; i != 100; i++ {
		go func() {
			group.Add(1)
			defer group.Done()
			for i := 0; i != 100000; i++ {
				lock.Lock()
				counter++
				lock.Unlock()
			}
		}()
	}

	group.Wait()

	if counter != 100*100000 {
		t.Errorf("Wrong counter value: %d", counter)
	}
}

func TestSpinlockInGoCorrectness(t *testing.T) {
	var lock SpinlockInGo
	var group sync.WaitGroup
	var begin, end time.Time

	go func() {
		group.Add(1)
		lock.Lock()
		time.Sleep(1500 * time.Millisecond)
		lock.Unlock()
		group.Done()
	}()

	time.Sleep(300 * time.Millisecond)

	go func() {
		group.Add(1)
		begin = time.Now()
		lock.Lock()
		end = time.Now()
		lock.Unlock()
		group.Done()
	}()

	group.Wait()

	var dur = end.Sub(begin)

	if dur < 1100*time.Millisecond || dur > 1300*time.Millisecond {
		t.Errorf("SpinlockInGo.Lock() is broken: %s", dur)
	}
}

func TestSpinlockInGoConcurrency(t *testing.T) {
	var counter uint
	var group sync.WaitGroup
	var lock SpinlockInGo

	for i := 0; i != 100; i++ {
		go func() {
			group.Add(1)
			defer group.Done()
			for i := 0; i != 100000; i++ {
				lock.Lock()
				counter++
				lock.Unlock()
			}
		}()
	}

	group.Wait()

	if counter != 100*100000 {
		t.Errorf("Wrong counter value: %d", counter)
	}
}

func BenchmarkSpinlockInAsm(b *testing.B) {
	var cases = []int{1, 2, 4, 8, 16}
	var counter int
	var lock SpinlockInAsm

	for _, concurrency := range cases {
		b.Run(strconv.Itoa(concurrency), func(b *testing.B) {
			b.SetParallelism(concurrency)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					lock.Lock()
					counter++
					lock.Unlock()
				}
			})
		})
	}
}

func BenchmarkSpinlockInGo(b *testing.B) {
	var cases = []int{1, 2, 4, 8, 16}
	var counter int
	var lock SpinlockInGo

	for _, concurrency := range cases {
		b.Run(strconv.Itoa(concurrency), func(b *testing.B) {
			b.SetParallelism(concurrency)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					lock.Lock()
					counter++
					lock.Unlock()
				}
			})
		})
	}
}

func BenchmarkSpinlockThin(b *testing.B) {
	var cases = []int{1, 2, 4, 8, 16}
	var counter int
	var lock uint32

	for _, concurrency := range cases {
		b.Run(strconv.Itoa(concurrency), func(b *testing.B) {
			b.SetParallelism(concurrency)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					Lock(&lock)
					counter++
					Unlock(&lock)
				}
			})
		})
	}
}
