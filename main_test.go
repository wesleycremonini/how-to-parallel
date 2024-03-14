package main

import "testing"

func Benchmark_normal_map(b *testing.B) {
	n := 10_000_000
	input := make([]Number, 0, n)
	for i := 0; i < n; i++ {
		input = append(input, Number{i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = normal_map(input, func(item Number, index int) Number {
			return Number{item.a * 2}
		})
	}
}

func Benchmark_map_concurrent_worker_pool(b *testing.B) {
	n := 10_000_000
	input := make([]Number, 0, n)
	for i := 0; i < n; i++ {
		input = append(input, Number{i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = map_concurrent_worker_pool(input, func(item Number, index int) Number {
			return Number{item.a * 2}
		})
	}
}

func Benchmark_map_concurrent_infinite_goroutines(b *testing.B) {
	n := 10_000_000
	input := make([]Number, 0, n)
	for i := 0; i < n; i++ {
		input = append(input, Number{i})
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = map_concurrent_infinite_goroutines(input, func(item Number, index int) Number {
			return Number{a: item.a * 2}
		})
	}
}
