package math

import (
	"testing"
)

func TestMax(t *testing.T) {
	want := 42
	if got := Max(10, 42); got != want {
		t.Errorf("Max(10, 42) = %v, want %v", got, want)
	}
}

func TestMin(t *testing.T) {
	want := 10
	if got := Min(10, 42); got != want {
		t.Errorf("Min(10, 42) = %v, want %v", got, want)
	}
}
