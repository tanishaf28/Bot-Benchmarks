# Metrics

Both commands exited 0.

## `go test -race -count=1 -cover ./...`

```
ok  	workerpool	1.024s	coverage: 100.0% of statements
```

## `go test -bench . -benchmem -count=3 ./...`

```
goos: linux
goarch: amd64
pkg: workerpool
cpu: Intel(R) Xeon(R) Processor
BenchmarkSequential-4   	   73150	     14705 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSequential-4   	   80496	     15140 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSequential-4   	   76377	     16379 ns/op	   81920 B/op	       1 allocs/op
BenchmarkWorkerPool-4   	     392	   3081720 ns/op	   82488 B/op	      11 allocs/op
BenchmarkWorkerPool-4   	     372	   3081965 ns/op	   82468 B/op	      11 allocs/op
BenchmarkWorkerPool-4   	     382	   3056120 ns/op	   82438 B/op	      11 allocs/op
PASS
ok  	workerpool	8.523s
```

## Summary

Numbers below are copied from the command output above. Benchmark figures are the three recorded runs, not an average. The race command does not print a test count; a same-package `go test -race -count=1 -json ./...` run recorded these pass events: `TestProcess` plus 6 subtests, with 0 failures.

| Metric | Value |
| --- | --- |
| Number of tests | 6 subtests, plus parent TestProcess |
| Pass count | 7 |
| Fail count | 0 |
| Coverage | 100.0% of statements |
| Sequential ns/op | 14705, 15140, 16379 |
| Worker-pool ns/op | 3081720, 3081965, 3056120 |
| Sequential allocs/op | 1, 1, 1 |
| Worker-pool allocs/op | 11, 11, 11 |
