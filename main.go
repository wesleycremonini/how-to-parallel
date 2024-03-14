package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Number struct {
	a int
}

func main() {
	const n = 10_000_000
	arr := make([]Number, 0, n)
	for i := range n {
		arr = append(arr, Number{i})
	}

	returnInt := func(t Number, index int) int {
		return t.a
	}

	func() {
		defer timer("normal map")()
		normal_map(arr, returnInt)
	}()

	func() {
		defer timer("infinite goroutines map")()
		map_concurrent_infinite_goroutines(arr, returnInt)
	}()

	func() {
		defer timer("worker pool map")()
		map_concurrent_worker_pool(arr, returnInt)
	}()
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		println()
		fmt.Printf("%s took %v\n", name, time.Since(start))
		println()
	}
}

func normal_map[T1, T2 any](arr []T1, f func(item T1, index int) T2) []T2 {
	arrT2 := make([]T2, len(arr))

	for i, t := range arr {
		t2 := f(t, i)
		arrT2[i] = t2
	}

	return arrT2
}

func map_concurrent_infinite_goroutines[T1, T2 any](arr []T1, f func(item T1, index int) T2) []T2 {
	var wg sync.WaitGroup
	wg.Add(len(arr))

	arrT2 := make([]T2, len(arr))

	for i, t := range arr {
		go func() {
			t2 := f(t, i)
			arrT2[i] = t2

			wg.Done()
		}()
	}

	wg.Wait()
	return arrT2
}

func map_concurrent_worker_pool[T1, T2 any](arr []T1, f func(item T1, index int) T2) []T2 {
	arrT2 := make([]T2, len(arr))

	numworkers := max_parallelism()

	var wg sync.WaitGroup
	wg.Add(numworkers)

	worker := func(startIndex, endIndex int) {
		for i := startIndex; i < endIndex; i++ {
			t2 := f(arr[i], i)
			arrT2[i] = t2
		}

		wg.Done()
	}

	chunkSize := len(arr) / numworkers
	for i := 0; i < numworkers; i++ {
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if i == numworkers-1 {
			endIndex = len(arr)
		}

		go worker(startIndex, endIndex)
	}

	wg.Wait()

	return arrT2
}

func max_parallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}