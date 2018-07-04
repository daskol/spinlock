package spinlock

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	var cases = []int{1, 2, 4, 8, 16}
	var counter int
	var lock sync.Mutex

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
