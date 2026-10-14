# Metrics

Both commands exited 0.

## `go test -race -count=1 -cover ./...`

```
ok  	workerpool	1.019s	coverage: 100.0% of statements
```

## `go test -bench . -benchmem -count=3 ./...`

```
goos: linux
goarch: amd64
pkg: workerpool
cpu: Intel(R) Xeon(R) Processor
BenchmarkSequential-4   	   84980	     13621 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSequential-4   	   84910	     13354 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSequential-4   	   89978	     13140 ns/op	   81920 B/op	       1 allocs/op
BenchmarkWorkerPool-4   	     505	   2169413 ns/op	   82477 B/op	      11 allocs/op
BenchmarkWorkerPool-4   	     568	   2208805 ns/op	   82442 B/op	      11 allocs/op
BenchmarkWorkerPool-4   	     590	   2042603 ns/op	   82426 B/op	      11 allocs/op
PASS
ok  	workerpool	8.134s
```

## Summary

Numbers below are copied from the command output above. Benchmark figures are the three recorded runs, not an average. The race command does not print a test count; a same-package `go test -race -count=1 -json` run recorded 6 subtests and the parent `TestProcess`, with 0 failures.

| Metric | Value |
| --- | --- |
| Number of tests | 6 subtests, plus parent TestProcess |
| Pass count | 7 |
| Fail count | 0 |
| Coverage | 100.0% of statements |
| Sequential ns/op | 13621, 13354, 13140 |
| Worker-pool ns/op | 2169413, 2208805, 2042603 |
| Sequential allocs/op | 1, 1, 1 |
| Worker-pool allocs/op | 11, 11, 11 |
