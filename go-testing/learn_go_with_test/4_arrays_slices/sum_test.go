package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := [5]int{1, 2, 3, 4, 5}

		result := Sum(nums)

		expected := 15

		if result != expected {
			t.Errorf("result %d want %d given, %v", result, expected, nums)

		}
	})
}

func TestSum2(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{10, 20, 30, 40, 50}

		result := Sum2(nums)

		expected := 150

		if result != expected {
			t.Errorf("result %d want %d given, %v", result, expected, nums)
		}
	})
}

func TestSumAll(t *testing.T) {

	result := SumAll([]int{1, 2}, []int{9, 8})
	expected := []int{3, 17}

	// if !reflect.DeepEqual(result, expected) {
	// 	t.Errorf("result %v expected %v", result, expected)
	// }

	// After 1.21
	if !slices.Equal(result, expected) {
		t.Errorf("result %v expected %v", result, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		result := SumAllTails([]int{1, 2}, []int{9, 8})
		expected := []int{2, 8}

		if !slices.Equal(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{99, 100})
		expected := []int{0, 100}

		if !slices.Equal(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})
}
