package calc

import "testing"

func TestSumBasic(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Fatalf("want 6, got %d", got)
	}
}

func TestSumEmpty(t *testing.T) {
	if got := Sum(nil); got != 0 {
		t.Fatalf("want 0, got %d", got)
	}
}

func TestSumMixed(t *testing.T) {
	if got := Sum([]int{10, -3, 2}); got != 9 {
		t.Fatalf("want 9, got %d", got)
	}
}
