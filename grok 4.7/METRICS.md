# Metrics

## `go test -race -count=1 -cover ./...`

```
ok  	workerpool	1.318s	coverage: 100.0% of statements
```

## `go test -bench . -benchmem -count=3 ./...`

```
goos: darwin
goarch: arm64
pkg: workerpool
BenchmarkSequential-6   	  176931	      7621 ns/op	   81923 B/op	       1 allocs/op
BenchmarkSequential-6   	  157638	      7526 ns/op	   81924 B/op	       1 allocs/op
BenchmarkSequential-6   	  160228	      7861 ns/op	   81924 B/op	       1 allocs/op
BenchmarkWorkerPool-6   	     384	   3435570 ns/op	   82522 B/op	      11 allocs/op
BenchmarkWorkerPool-6   	     376	   3207198 ns/op	   82476 B/op	      11 allocs/op
BenchmarkWorkerPool-6   	     367	   3271905 ns/op	   82452 B/op	      11 allocs/op
PASS
ok  	workerpool	9.877s
```

## Summary

Both commands exited 0. Benchmark figures are the three recorded runs, not an average. The race command does not print a test count; `go test -race -count=1 -json ./...` recorded 6 subtests and the parent `TestProcess`, all pass, 0 fail.

| Metric | Value |
| --- | --- |
| Number of tests | 6 subtests, plus parent TestProcess |
| Pass count | 7 |
| Fail count | 0 |
| Coverage | 100.0% of statements |
| Sequential ns/op | 7621, 7526, 7861 |
| Worker-pool ns/op | 3435570, 3207198, 3271905 |
| Sequential allocs/op | 1, 1, 1 |
| Worker-pool allocs/op | 11, 11, 11 |
