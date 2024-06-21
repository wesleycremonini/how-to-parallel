[Medium Article](https://medium.com/@wesleycremonini0/why-concurrent-code-can-be-slower-than-sequential-code-88d82d87ff09)

- NORMAL MAP BENCHMARK (ONE CORE) -> 24.67ms
- INFINITE GOROUTINES BENCHMARK (SLOWEST) -> 2400ms
- WORKER POOL (RESPECTING NUM OF CORES BENCHMARK) -> 14.64ms

This is due to ineffective memory sharing and excessive context switching from spawning too many goroutines with only 4 CPU cores. When using the maximum number of CPU cores, with efficient memory sharing (to prevent other CPU core caches from being invalidated when accessed from another CPU core), and also ensuring that each goroutine only reads contiguous memory, we can prevent core context switches, which are very slow.
