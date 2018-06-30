package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	vals := []string{"Centaurs", "Dwarves", "Elves", "Giants", "Halflings", "Merfolk", "Minotaurs",
		"Orcs", "Skeletons", "Trolls", "Wingfolk", "Wizard"}
	//var classesPicked = make([]string, 5)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(vals)) {
		val := vals[i]
		fmt.Println(val)
	}
}

/*
func main() {
	rand.Seed(time.Now().UTC.UnixNano())
}
*/
