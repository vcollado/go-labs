package main

import (
	"fmt"

	"github.com/vcollado/tasks"
)

func main() {
	var dna1, dna2 string = "GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"

	// hamming
	hammingDistance := tasks.Hamming(dna1, dna2)
	fmt.Println(dna1)
	fmt.Println(dna2)
	fmt.Println(hammingDistance)
	// output:
	// GAGCCTACTAACGGGAT
	// CATCGTAATGACGGCCT
	// ^ ^ ^  ^ ^    ^^

	// accumulator
	fmt.Println(tasks.Accumulate([]int{25, 5, 3}, tasks.Square))
	// output:
	// [625 25 9]

	// acronym
	fmt.Println(tasks.Acronym("I Am Your Gopher!"))

	// Bob
	p1 := "How are you?"
	fmt.Println(p1, tasks.Bob(p1))
	p2 := "Shut up!"
	fmt.Println(p2, tasks.Bob(p2))
	p3 := ""
	fmt.Println(p3, tasks.Bob(p3))
	p4 := "hi"
	fmt.Println(p4, tasks.Bob(p4))
}
