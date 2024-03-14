- NORMAL MAP BENCHMARK (ONE CORE) -> 24.67ms
- INFINITE GOROUTINES BENCHMARK (SLOWEST) -> 2400ms
- WORKER POOL (RESPECTING NUM OF CORES BENCHMARK) -> 14.64ms

That's because of ineffective memory sharing and excessive context switching from spawnning too many goroutines with only 4 CPU cores.
When using the max num of CPU cores, with efficient memory sharing (that will prevent others CPU core cache from being invalidated when accessed from another CPU core), and also making it so each goroutine only reads contiguos memory, preventing core context switch (which is very slow).
