package main

import "testing"

func TestMax(t *testing.T) {
	// 测试用例 1: a > b
	result := Max(5, 3)
	if result != 5 {
		t.Errorf("Max(5, 3) should be 5, but got %d", result)
	}

	// 测试用例 2: a <= b
	result = Max(2, 7)
	if result != 7 {
		t.Errorf("Max(2, 7) should be 7, but got %d", result)
	}
}
