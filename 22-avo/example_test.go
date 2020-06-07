package example

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x, y uint64
	x, y = 2, 3
	result := Add(x, y)
	if result != x+y {
		t.Fatalf("expect %v+%v=%v, but get %v", x, y, x+y, result)
	}
}
