package sleepsort

import (
	"sync"
	"time"
)

func Sort[T any](items []T, duration func(val T) time.Duration) []T {
	resultChan := make(chan T, len(items))
	wg := sync.WaitGroup{}

	var result []T
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(items); i++ {
			result = append(result, <-resultChan)
		}
		close(resultChan)
	}()

	for _, v := range items {
		wg.Add(1)
		go func(v T) {
			defer wg.Done()
			time.Sleep(duration(v))
			resultChan <- v
		}(v)
	}

	wg.Wait()
	return result
}

func IntSlice(s []int) []int {
	return Sort(s, func(v int) time.Duration {
		return time.Duration(v) * time.Millisecond
	})
}

func Float64Slice(s []float64) []float64 {
	return Sort(s, func(v float64) time.Duration {
		return time.Duration(v) * time.Millisecond
	})
}
