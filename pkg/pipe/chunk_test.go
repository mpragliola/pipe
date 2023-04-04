package pipe_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mpragliola/pipe/pkg/pipe"
)

func TestChunk(t *testing.T) {
	type test struct {
		in  []int
		out [][]int
	}

	tests := []test{
		{in: []int{}, out: [][]int{}},
		{in: []int{1}, out: [][]int{{1}}},
		{in: []int{1, 2}, out: [][]int{{1, 2}}},
		{in: []int{1, 2, 3}, out: [][]int{{1, 2, 3}}},
		{in: []int{1, 2, 3, 4}, out: [][]int{{1, 2, 3}, {4}}},
		{in: []int{1, 2, 3, 4, 5}, out: [][]int{{1, 2, 3}, {4, 5}}},
		{in: []int{1, 2, 3, 4, 5, 6}, out: [][]int{{1, 2, 3}, {4, 5, 6}}},
		{in: []int{1, 2, 3, 4, 5, 6, 7}, out: [][]int{{1, 2, 3}, {4, 5, 6}, {7}}},
	}

	for _, test := range tests {
		got := pipe.Scoop(pipe.Chunk(3, pipe.Of(test.in...)))

		if !cmp.Equal(got, test.out) {
			t.Errorf("Chunk(3, %v) = %v, want %v", test.in, got, test.out)
		}
	}
}
