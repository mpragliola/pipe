package pipe_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mpragliola/pipe/pkg/pipe"
)

func TestOfType(t *testing.T) {
	intOf := pipe.Of(1, 2, 3, 4, 5)

	if fmt.Sprintf("%T", intOf) != "<-chan int" {
		t.Error("Not a <-chan int type")
	}

	stringOf := pipe.Of("1", "2", "3", "4", "5")

	if fmt.Sprintf("%T", stringOf) != "<-chan string" {
		t.Error("Not a <-chan string type")
	}
}

func TestOfChan(t *testing.T) {
	intOf := pipe.Of(1, 2, 3, 4, 5)

	got := pipe.Scoop(intOf)
	want := []int{1, 2, 3, 4, 5}

	if !cmp.Equal(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestOfFunc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var i, j = 0, 1
	p := pipe.OfFunc(
		ctx,
		func() int {
			i, j = j, j+i

			if i > 10 {
				cancel()
			}

			return i
		},
	)

	got := pipe.Scoop(p)
	want := []int{1, 1, 2, 3, 5, 8, 13}

	if !cmp.Equal(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
