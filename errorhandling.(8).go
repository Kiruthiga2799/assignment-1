package main

import (
	"errors"
	"log"
)

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math:sqrt of negative number ")

	}
	return f * f, nil
}
