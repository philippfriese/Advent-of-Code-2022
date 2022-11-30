package tests

import (
	"AoC2022/helper"
	"testing"
)

func TestSum(t *testing.T) {
	arr := make([]int,10)
	for i,_ := range arr { arr[i] = i}
	if helper.Sum(arr) != 45 {
		t.Fail()
	}

	arr2 := make([]float32,10)
	for i,_ := range arr2 { arr2[i] = float32(i)}
	if helper.Sum(arr2) != 45.0 {
		t.Fail()
	}

	if helper.Sum([]float32{1.5,2.5,.1}) != 4.1 {
		t.Fail()
	}
}
