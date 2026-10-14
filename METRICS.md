# Metrics

Both commands exited 0.

## `go test -race -count=1 -cover ./...`

```
ok  	workerpool	1.365s	coverage: 100.0% of statements
```

## `go test -bench . -benchmem -count=3 ./...`

```
goos: darwin
goarch: arm64
pkg: workerpool
BenchmarkSequential-6   	  160322	      8862 ns/op	   81921 B/op	       1 allocs/op
BenchmarkSequential-6   	  122395	      9254 ns/op	   81922 B/op	       1 allocs/op
BenchmarkSequential-6   	  129022	      9232 ns/op	   81922 B/op	       1 allocs/op
BenchmarkWorkerPool-6   	     308	   3858821 ns/op	   82550 B/op	      11 allocs/op
BenchmarkWorkerPool-6   	     316	   4045337 ns/op	   82480 B/op	      11 allocs/op
BenchmarkWorkerPool-6   	     302	   3948706 ns/op	   82443 B/op	      11 allocs/op
PASS
ok  	workerpool	8.991s
```

## Summary

Numbers below are copied from the command output above. Benchmark figures are the three recorded runs, not an average. The race command does not print a test count; a same-package `go test -race -count=1 -json` run recorded 6 subtests and the parent `TestProcess`, with 0 failures.

| Metric | Value |
| --- | --- |
| Number of tests | 6 subtests, plus parent TestProcess |
| Pass count | 7 |
| Fail count | 0 |
| Coverage | 100.0% of statements |
| Sequential ns/op | 8862, 9254, 9232 |
| Worker-pool ns/op | 3858821, 4045337, 3948706 |
| Sequential allocs/op | 1, 1, 1 |
| Worker-pool allocs/op | 11, 11, 11 |
