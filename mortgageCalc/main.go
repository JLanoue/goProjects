package main

import (
	"fmt"
	"math"
)

func main() {
	var yrs float64
	var loan float64
	var rate float64
	yrs = 30
	loan = 300000
	rate = .0425
	var payment float64
	payment = loan * ((rate / 12) * (math.Pow((1+(rate/12)), (yrs*12)) / (math.Pow((1+(rate/12)), (yrs*12)) - 1)))
	fmt.Println(payment)
}

//func main (yrs int, loan int, rate float64) (payment float64) {
