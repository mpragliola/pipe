package pipe_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mpragliola/pipe/pkg/pipe"
)

func TestSpread(t *testing.T) {
	data := []string{"1,2,3", "4,5", "6,7,8", "9", "10"}
	split := func(s string) []string { return strings.Split(s, ",") }

	got := pipe.Scoop(pipe.Spread(pipe.Pipe(split, pipe.Of(data...))))
	want := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	if !cmp.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
