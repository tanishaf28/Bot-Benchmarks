package workerpool

import "sync"

// process squares every item using a worker pool and returns the squares
// in the same order as items. A worker count below 1 is treated as 1.
func process(items []int, workers int) []int {
	n := len(items)
	results := make([]int, n)
	if n == 0 {
		return results
	}
	if workers < 1 {
		workers = 1
	}
	if workers > n {
		workers = n
	}

	// Index travels with the value so each job writes exactly one slot.
	// Distinct indices do not race; WaitGroup.Wait happens-before return.
	type job struct {
		index int
		value int
	}
	jobs := make(chan job)

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := range jobs {
				results[j.index] = j.value * j.value
			}
		}()
	}

	for i, v := range items {
		jobs <- job{index: i, value: v}
	}
	close(jobs)
	wg.Wait()
	return results
}
