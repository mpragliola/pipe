package pipe_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mpragliola/pipe/pkg/pipe"
)

func TestScoop(t *testing.T) {
	got := pipe.Scoop(pipe.Of(1, 2, 3, 4, 5))
	want := []int{1, 2, 3, 4, 5}

	if !cmp.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
