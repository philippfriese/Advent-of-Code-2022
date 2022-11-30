package helper

import "strconv"

func ConvFloat32(s string) float32 {
	v, _ := strconv.ParseFloat(s, 32)
	return float32(v)
}

func ConvFloat64(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func ConvInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

