# Original `process` bugs

The starter function was supposed to square every `int` in a slice with a
worker pool and return results in **input order**. It never could.

```go
func process(items []int, workers int) []int {
    results := make([]int, 0)
    var wg sync.WaitGroup
    ch := make(chan int)
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            for it := range ch {
                results = append(results, it*it)
            }
            wg.Done()
        }()
    }
    for i := 0; i <= len(items); i++ {
        ch <- items[i]
    }
    wg.Wait()
    return results
}
```

## What it currently does

Workers are started and then blocked on an unbuffered `chan int`. The main
goroutine tries to send every element, then waits for the workers.

That sequence never completes successfully:

- If `workers < 1`, the first send has no receiver and deadlocks.
- Otherwise each in-range send is received, then `i == len(items)` indexes
  one past the end and panics. The channel is never closed, so surviving
  workers hang in `range ch`.

Even if the panic and deadlock were fixed, results would still be racy and
unordered.

## Bugs

1. **Off-by-one panic.** `for i := 0; i <= len(items); i++` reads `items[len(items)]`.
   An empty slice panics on the first iteration (`items[0]`).
2. **Channel never closed.** Workers `range` until close. `close(ch)` is missing,
   so they never finish and `wg.Wait()` cannot return.
3. **Wait / close ordering.** Even with a close, it must happen *before*
   `wg.Wait()`, after the last send. Waiting first deadlocks.
4. **`workers < 1` deadlock.** No consumers on an unbuffered channel; the first
   `ch <- items[i]` blocks forever.
5. **Data race on `results`.** Concurrent `append` mutates the slice header and
   backing array with no mutex. `go test -race` fails; lost updates or panics
   are possible.
6. **Order is not preserved.** Completion order is not input order, so append
   order cannot match `items` even under a mutex.
7. **No index on the job.** The channel carries only a value, so a worker has
   no slot to write into a preallocated result slice.
8. **`wg.Done()` is not deferred.** A panic in a worker leaks the WaitGroup
   (secondary to the bugs above).
