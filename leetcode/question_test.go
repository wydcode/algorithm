package leetcode

import (
	"fmt"
	"testing"
)

func TestQuestion_pivotIndex(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	idx1 := pivotIndex(nums)
	if idx1 != 3 {
		fmt.Println("idx1: ", idx1)
		t.Fail()
	}

	nums1 := []int{1, 2, 3}
	idx2 := pivotIndex(nums1)
	if idx2 != -1 {
		fmt.Println("idx2: ", idx2)
		t.Fail()
	}
}

func TestQuestion_pivotIndex2(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	idx1 := pivotIndex2(nums)
	if idx1 != 3 {
		fmt.Println("idx1: ", idx1)
		t.Fail()
	}

	nums1 := []int{1, 2, 3}
	idx2 := pivotIndex2(nums1)
	if idx2 != -1 {
		fmt.Println("idx2: ", idx2)
		t.Fail()
	}
}

func TestQuestion_canCompleteCircuit(t *testing.T) {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	idx := canCompleteCircuit(gas, cost)
	if idx != 3 {
		fmt.Println("idx: ", idx)
		t.Fail()
	}
}
