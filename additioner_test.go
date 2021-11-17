package concurrency

import (
	"fmt"
	"testing"
)

func benchmarkAddWithMutex(concurrency int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddWithMutex(concurrency)
	}
}

func benchmarkAddWithChannels(concurrency int,b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddWithChannels(concurrency)
	}
}

func BenchmarkAdd(b *testing.B) {
	parallelism := []int{300, 500, 1000}
	for _, c := range parallelism {
		name := fmt.Sprintf("BenchmarkAddWithChannels%d", c)
		b.Run(name, func(b *testing.B) {
			benchmarkAddWithChannels(c, b)
		})

		name = fmt.Sprintf("BenchmarkAddWithMutex%d", c)
		b.Run(name, func(b *testing.B) {
			benchmarkAddWithMutex(c, b)
		})
	}
}
