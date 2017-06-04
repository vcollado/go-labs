package main

import (
	"fmt"

	"github.com/vcollado/exercism"
)

func main() {
	var dna1, dna2 string = "GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"

	hammingDistance := exercism.CalculateHammingDistance(dna1, dna2)

	fmt.Println(dna1)
	fmt.Println(dna2)
	fmt.Println(hammingDistance)
}
