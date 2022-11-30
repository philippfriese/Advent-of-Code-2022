package tests

import (
	"AoC2022/helper"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	fn, _ := filepath.Abs("./data/string_test")

	iarr := make([]int, 10)
	for i, _ := range iarr { iarr[i] = i }
	if !reflect.DeepEqual(helper.Split[float32](fn, "\n", helper.ConvFloat32),
						  helper.Collect(iarr, func(i int) float32 { return float32(i) })) {
		t.Fail()
	}

	if !reflect.DeepEqual(helper.Split[int](fn, "\n", helper.ConvInt),
		helper.Collect(iarr, func(i int) int { return int(i) })) {
		t.Fail()
	}
}
