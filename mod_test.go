package clone_test

import (
	"testing"

	"github.com/nbarena/clone"
)

type foo struct {
	x int
}

func (f *foo) Clone() *foo {
	return &foo{f.x}
}

func TestCloneSlice(t *testing.T) {
	xs := []*foo{{1}}
	ys := clone.Slice(xs)
	xs[0].x = 500

	if ys[0].x != 1 {
		t.Errorf("slice was not cloned")
	}
}

func TestCloneMap(t *testing.T) {
	xs := map[string]*foo{"a": {1}}
	ys := clone.Map(xs)
	xs["a"].x = 500

	if ys["a"].x != 1 {
		t.Errorf("map was not cloned")
	}
}

func TestShallow(t *testing.T) {
	xp := clone.P(1)
	yp := clone.Shallow(xp)
	*xp = 500

	if *yp != 1 {
		t.Errorf("value pointer was not copied")
	}
}

func TestShallowNil(t *testing.T) {
	if clone.Shallow[int](nil) != nil {
		t.Errorf("nil value pointer was not copied")
	}
}

func TestCloneValuePointer(t *testing.T) {
	x := &foo{1}
	y := clone.ValuePointer(x)
	x.x = 500

	if y.x != 1 {
		t.Errorf("value pointer was not cloned")
	}
}

func TestCloneValuePointerNil(t *testing.T) {
	if clone.ValuePointer[foo](nil) != nil {
		t.Errorf("nil value pointer was not cloned")
	}
}
