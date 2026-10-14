# Benchmark Cursor's Models

I ran one identical prompt in Cursor against three different models and compared correctness, speed, and how each one handled problems along the way.

| Model | Provider in Cursor |
|---|---|
| **Grok 4.6** | Cursor Models pool |
| **Grok 4.7** | Cursor Models pool |
| **Composer 2.5** | Cursor's own agentic coding model |

---

## The Task

I gave each model a short, buggy Go function that is supposed to square every integer in a slice using a worker pool, returning results in the **same order as the input**. I planted four bugs on purpose so there is a known answer key:

1. **Data race** on the shared `results` slice (concurrent `append`)
2. **Off-by-one** in the send loop (`i <= len(items)` panics)
3. **Deadlock**: the channel is never closed, so `wg.Wait()` never returns
4. **Nondeterministic order**: results come back in completion order, not input order

### The prompt (identical for every model)

1. Explain what the code currently does and list every bug.
2. Rewrite it so it is correct, race-free, and preserves input order.
3. Write table-driven tests (empty input, 1 item, workers > items, workers = 1, large input) that pass with `go test -race`.
4. Add a benchmark comparing the worker pool against a plain sequential loop.
5. Run `go test -race -count=1 -cover ./...` and `go test -bench . -benchmem -count=3 ./...`, paste the raw output into `METRICS.md`, and summarize it in a table. Do not estimate or guess any numbers.

## Method

- Same prompt, word for word, in a fresh chat per model (Agent mode).
- No follow-up prompts.
- Time measured by stopwatch from sending the prompt to the end of the output.
- Metrics (tests, coverage, benchmarks) measured by running `measure.sh` in each project folder.

---

## Results

### Side by side

| | **Grok 4.6** | **Grok 4.7** | **Composer 2.5** |
|---|---|---|---|
| **Time (prompt to end)** | 3:11 | 7:19 (run 1), 0:60 (run 2) | **1:08** |
| **Seeded bugs found** | 4 / 4 | 4 / 4 | 4 / 4 |
| **Total issues listed** | **8** | 7 | 7 |
| **Fix approach** | Preallocated slice, indexed jobs | Same | Same |
| **Tests** | 6 subtests + parent (7 pass / 0 fail) | 6 subtests + parent (7 / 0) | 6 subtests + parent (7 / 0) |
| **Coverage** | 100% | 100% | 100% |
| **`go test -race`** | Pass | Pass | Pass |
| **Sequential ns/op (avg of 3)** | ~15,400 | ~9,100 | ~12,400 |
| **Worker pool ns/op (avg of 3)** | ~3,070,000 | ~3,950,000 | ~1,990,000 |
| **Pool vs sequential** | ~200x slower | ~430x slower | ~160x slower |
| **Allocs/op (seq / pool)** | 1 / 11 | 1 / 11 | 1 / 11 |
| **Machine** | Linux amd64, 4 cores | macOS arm64, 6 cores | Linux amd64, 4 cores |

> **Note on benchmarks:** absolute ns/op values are not comparable across models because Grok 4.7 ran on a Mac and the other two on a Linux VM. The meaningful result is the ratio: in every run the worker pool was far slower than the sequential loop, because one multiply per element costs less than a channel send. All three models reported this correctly.

> **Note on Grok 4.7's time:** the 7:19 first run included Go toolchain setup. The 0:60 second run is the fairer number.


---

## Pros and Cons

### Grok 4.6
**Pros**
- Most thorough bug list (8 issues), including the non-deferred `wg.Done()` and the missing job index
- Hit a missing `bc` command, installed it, and re-ran the pipeline
- Documented the bugs in a separate file and opened a PR

**Cons**
- Went beyond the brief (PR and extra docs weren't requested)
- Its two summary tables disagree on test count (7 vs 1), because one counted only `func Test` lines
- Slower than Composer 2.5

### Grok 4.7
**Pros**
- Most careful about metrics: noticed the race command doesn't print a test count and obtained one another way (`-json` run)
- Diagnosed a macOS-specific `paste` failure and explained the fix
- Included a clear summary table and noted its numbers were copied, not averaged

**Cons**
- Slowest first run (7:19), partly toolchain setup
- Wordiest output of the three

### Composer 2.5
**Pros**
- Fastest by a wide margin (1:08)
- Concise and stayed in scope, with the fix shown inline
- Same correct design and 100% coverage as the others

**Cons**
- Stopped at the `bc: command not found` error and reported exit 127 without fixing it
- Shortest explanation of the bugs, with no discussion of secondary issues like `wg.Done()`

---

## Which Model Is Good at What?

| Strength | Best pick | Evidence |
|---|---|---|
| Speed | **Composer 2.5** | 1:08 vs 3:11 and 7:19 |
| Bug-finding thoroughness | **Grok 4.6** | 8 issues listed |
| Metrics rigor | **Grok 4.7** | Verified test count independently |
| Recovering from tooling errors | **Grok 4.6 / 4.7** | Fixed or diagnosed `bc` / `paste` |
| Staying in scope | **Composer 2.5** | No unrequested extras |
| Core correctness | **Tie** | All passed `-race` with 100% coverage |

## Takeaways

- **All three solved the core task correctly.** Same design, same test cases, same coverage, all four seeded bugs found. For a task this size, correctness did not separate them.
- **The differences are in process:** Composer 2.5 is the fastest and most concise, while the Grok models are slower but more thorough and better at working around environment problems.
- **Limitation:** one run per model on one task. A harder prompt (subtler concurrency bugs, a multi-file task) would likely separate them more.

## Reproduce

```bash
# in each model's project folder
./measure.sh
```

Check [cursor.com/docs](https://cursor.com/docs) for the latest model list, since names and availability change often.
