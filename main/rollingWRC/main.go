package main

import (
	"fmt"
)

func main() {
	wrc := []int{30, -3, 145, 318, 281, 181, 120, 212, 266, 720, 51, -100, 145, 115, 913, 120, 332, 30, -3, -27, 442, -100, 427, 330, -100, 1086, 332, 231, -27, 181, 302, -100, 387, -43, 344, 215, 172, -100, 181, 336, 505, 231, 145, 382, 115, 30, 349, 94, -1, 255}

	abs := []int{4, 4, 5, 5, 5, 2, 6, 4, 5, 5, 5, 5, 5, 2, 5, 5, 5, 4, 4, 5, 5, 5, 4, 3, 2, 4, 5, 4, 5, 2, 4, 5, 5, 6, 5, 5, 5, 5, 4, 5, 5, 4, 5, 4, 4, 4, 4, 5, 5, 5}

	games := 15 //Set to how many games rolling average should be

	low := 0
	high := games

	//fmt.Println(wrcSelection)
	//fmt.Println(absSelection)

	var sum int

	for j := 0; j < 30; j++ {

		var wrcSelection []int = wrc[low:high]
		var absSelection []int = abs[low:high]

		for i := 0; i < games; i++ {
			x := absSelection[i] * wrcSelection[i]
			sum += x
		}
		//fmt.Println(sum)

		absSum := 0
		for _, num := range absSelection {
			absSum += num
		}
		//fmt.Println(absSum)

		wrcavg := sum / absSum
		fmt.Println(wrcavg)

		low++
		high++
		sum = 0

	}
}
