// Table-driven tests and benchmarks for process. Bug list for the broken
// starter: docs/worker-pool-bugs.md in the project Agent Store (Context).
package workerpool

import "testing"

func sequential(items []int) []int {
	out := make([]int, len(items))
	for i, v := range items {
		out[i] = v * v
	}
	return out
}

func TestProcess(t *testing.T) {
	large := make([]int, 10000)
	for i := range large {
		large[i] = (i*17 + 3) % 1000
	}

	tests := []struct {
		name    string
		items   []int
		workers int
	}{
		{name: "empty input", items: []int{}, workers: 4},
		{name: "one item", items: []int{7}, workers: 3},
		{name: "workers greater than items", items: []int{3, 1, 4, 1, 5}, workers: 16},
		{name: "workers equals one", items: []int{9, 0, -2, 8, 6}, workers: 1},
		{name: "large input", items: large, workers: 8},
		{name: "non-positive workers", items: []int{2, -4, 6}, workers: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := process(tt.items, tt.workers)
			want := sequential(tt.items)
			if len(got) != len(want) {
				t.Fatalf("len(got)=%d len(want)=%d", len(got), len(want))
			}
			for i := range want {
				if got[i] != want[i] {
					t.Fatalf("index %d: got %d, want %d", i, got[i], want[i])
				}
			}
		})
	}
}

const benchItems = 10000

func benchmarkInput() []int {
	items := make([]int, benchItems)
	for i := range items {
		items[i] = i
	}
	return items
}

// benchSink keeps results live so the compiler cannot delete the work.
var benchSink int64

func BenchmarkSequential(b *testing.B) {
	items := benchmarkInput()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := sequential(items)
		benchSink += int64(out[len(out)-1])
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	items := benchmarkInput()
	workers := 8
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := process(items, workers)
		benchSink += int64(out[len(out)-1])
	}
}
