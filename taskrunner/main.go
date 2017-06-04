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
}
