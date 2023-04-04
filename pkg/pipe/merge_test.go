package pipe_test

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mpragliola/pipe/pkg/pipe"
)

func TestMerge(t *testing.T) {
	c1 := pipe.Of(1, 2, 3)
	c2 := pipe.Of(4, 5, 6)

	got := pipe.Scoop(pipe.Merge(c1, c2))
	want := []int{1, 2, 3, 4, 5, 6}

	// order might be sequential but arbitrary between the merged pipes
	sort.Ints(got)

	if !cmp.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
