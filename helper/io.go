package helper

import (
	"log"
	"os"
)

func ReadLines(fn string) string {
	bytes, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(bytes)
}
