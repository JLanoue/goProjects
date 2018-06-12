// Calculate value to nullify credit card fee
// This is for a non-profit fundraiser where people want to pay more to cover cc fees
package main

import (
	"fmt"
)

func main() {
	var total float64
	var cost float64
	var fee float64
	var ccfee float64
	var takehome float64
	cost = 100
	fee = .025
	total = cost / (1 - fee)
	ccfee = total * fee
	takehome = total - (total * fee)
	fmt.Println(total)
	fmt.Println(ccfee)
	fmt.Println(takehome)
}
