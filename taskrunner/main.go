package main

import (
	"fmt"

	"github.com/vcollado/tasks"
)

func main() {
	var dna1, dna2 string = "GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"

	hammingDistance := tasks.Hamming(dna1, dna2)

	fmt.Println(dna1)
	fmt.Println(dna2)
	fmt.Println(hammingDistance)
}
