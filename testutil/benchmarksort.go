package testutil

import "sort"

type sortFunc[T any] func(s []T) sort.Interface

func BaseSort[T any](s []T, sortFunc func(s []T) sort.Interface) []T {
	safe := make([]T, len(s))
	copy(safe, s)
	sort.Sort(sortFunc(safe))
	return safe
}

func BaseIntSlice(s []int) []int {
	return BaseSort(s, func(s []int) sort.Interface {
		return sort.IntSlice(s)
	})
}

func BaseFloat64Slice(s []float64) []float64 {
	return BaseSort(s, func(s []float64) sort.Interface {
		return sort.Float64Slice(s)
	})
}
