package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calc())
}

func calc(yrs, loan, rate float64) float64 {
	return loan * ((rate / 12) * (math.Pow((1+(rate/12)), (yrs*12)) / (math.Pow((1+(rate/12)), (yrs*12)) - 1)))
}

//func main (yrs int, loan int, rate float64) (payment float64) {
