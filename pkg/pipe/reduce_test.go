package pipe_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mpragliola/pipe/pkg/pipe"
)

func TestReduce(t *testing.T) {
	sum := func(a, b int) int {
		return a + b
	}

	want := []int{1, 3, 6, 10}

	got := pipe.Scoop(pipe.Reduce(sum, 0, pipe.Of(1, 2, 3, 4)))

	if !cmp.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
