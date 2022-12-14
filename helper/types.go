package helper

import "strconv"

func ConvFloat32(s string) float32 {
	v, e := strconv.ParseFloat(s, 32)
	if e != nil { panic(e) }
	return float32(v)
}

func ConvFloat64(s string) float64 {
	v, e := strconv.ParseFloat(s, 64)
	if e != nil { panic(e) }
	return v
}

func ConvInt(s string) int {
	v, e := strconv.Atoi(s)
	if e != nil { panic(e) }
	return v
}

