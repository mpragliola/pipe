package pipe_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mpragliola/pipe/pkg/pipe"
)

func TestPipe(t *testing.T) {
	double := func(i int) int { return i * 2 }
	add10 := func(i int) int { return i + 10 }

	got := pipe.Scoop(pipe.Pipe(add10, pipe.Pipe(double, pipe.Of(1, 2, 3))))
	want := []int{12, 14, 16}

	if cmp.Equal(got, want) == false {
		t.Errorf("got %v, want %v", got, want)
	}
}
