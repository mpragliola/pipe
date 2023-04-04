package pipe_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mpragliola/pipe/pkg/pipe"
)

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	oddFilter := func(i int) bool { return i%2 == 0 }

	got := pipe.Scoop(pipe.Filter(oddFilter, pipe.Of(input...)))
	want := []int{2, 4, 6, 8, 10}

	if !cmp.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
