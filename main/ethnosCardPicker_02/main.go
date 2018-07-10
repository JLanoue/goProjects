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
	//var classesPicked = make([]string, 5)

	var S []string

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(hordes)) {
		S = append(S, hordes[i])
		//fmt.Print(s[:])
		//hordeSelection := strings.Fields(val)
		//fmt.Println(hordeSelection[0:6])
	}
	fmt.Println(S[:(x)])
}
