package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var x int
	fmt.Println("How many classes do you need?")

	_, err := fmt.Scanf("%d", &x)

	if err != nil {
		fmt.Println(err)
	}

	hordes := []string{"Centaurs", "Dwarves", "Elves", "Giants", "Halflings", "Merfolk", "Minotaurs",
		"Orcs", "Skeletons", "Trolls", "Wingfolk", "Wizard"}

	var s []string

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(hordes)) {
		s = append(s, hordes[i])

	}
	fmt.Println(s[:(x)])
}
