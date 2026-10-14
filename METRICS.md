# Metrics

Recorded by running the `measure.sh` commands from the project root.
`go test -race` and `go test -bench` exited 0. `bc` was not installed on
the first pass; that error is pasted below. After `apt-get install bc`,
the test-count pipeline was re-run and printed `1`.

## `measure.sh` output (unedited)

```
== tests ==
ok  	workerpool	1.029s	coverage: 100.0% of statements
== bench ==
goos: linux
goarch: amd64
pkg: workerpool
cpu: Intel(R) Xeon(R) Processor
BenchmarkSequential-4   	   74118	     15820 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSequential-4   	   79615	     15632 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSequential-4   	   79022	     16069 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSequential-4   	   77985	     14785 ns/op	   81920 B/op	       1 allocs/op
BenchmarkSequential-4   	   85144	     16463 ns/op	   81920 B/op	       1 allocs/op
BenchmarkWorkerPool-4   	     386	   3223679 ns/op	   82454 B/op	      11 allocs/op
BenchmarkWorkerPool-4   	     374	   3202914 ns/op	   82452 B/op	      11 allocs/op
BenchmarkWorkerPool-4   	     386	   3090290 ns/op	   82448 B/op	      11 allocs/op
BenchmarkWorkerPool-4   	     388	   3173697 ns/op	   82445 B/op	      11 allocs/op
BenchmarkWorkerPool-4   	     579	   1964955 ns/op	   82444 B/op	      11 allocs/op
PASS
ok  	workerpool	14.526s
== size ==
 123 total
== test count ==
--: line 4: bc: command not found
```

## `== test count ==` after installing `bc`

```
1
```

`tests.txt` and `bench.txt` contain the tee'd `go test` output from this run.

## Summary

Numbers below are copied from the output above. Benchmark figures are the
five recorded runs, not an average. Test count is the `grep -rhc "func Test"`
pipeline (one `func TestProcess`). Fail count is 0 because `go test -race`
printed `ok`.

| Metric | Value |
| --- | --- |
| Number of tests | 1 (`func Test` count) |
| Pass count | 1 (`ok` from `go test -race`) |
| Fail count | 0 |
| Coverage | 100.0% of statements |
| Go lines (`wc -l` total) | 123 |
| Sequential ns/op | 15820, 15632, 16069, 14785, 16463 |
| Worker-pool ns/op | 3223679, 3202914, 3090290, 3173697, 1964955 |
| Sequential allocs/op | 1, 1, 1, 1, 1 |
| Worker-pool allocs/op | 11, 11, 11, 11, 11 |
