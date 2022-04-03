package sleepsort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/connormckelvey/sleep-sort/testutil"
	"github.com/stretchr/testify/assert"
)

func TestIntSlice(t *testing.T) {
	tests := []struct {
		input []int
	}{
		{[]int{2, 3, 5, 3, 8}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			expected := testutil.BaseIntSlice(test.input)
			actual := IntSlice(test.input)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestFloat64Slice(t *testing.T) {
	tests := []struct {
		input []float64
	}{
		{[]float64{0, 2.1, 3.2, 5.3, 3.4, 8.5}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			expected := testutil.BaseFloat64Slice(test.input)
			actual := Float64Slice(test.input)
			assert.ElementsMatch(t, expected, actual)
		})
	}
}

func BenchmarkBaseIntSlice(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		testutil.BaseIntSlice(rand.Perm(100))
	}
}

func BenchmarkSleepSortIntSlice(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		IntSlice(rand.Perm(100))
	}
}
