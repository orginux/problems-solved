package main

import "testing"

func TestDubicates(t *testing.T) {
	t.Run("Odd", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 3}
		want := 3
		got := removeDuplicates(nums)

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Even", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 3, 3}
		want := 3
		got := removeDuplicates(nums)

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("One", func(t *testing.T) {
		nums := []int{0}
		want := 1
		got := removeDuplicates(nums)

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Zero", func(t *testing.T) {
		nums := []int{}
		want := 0
		got := removeDuplicates(nums)

		if want != got {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
