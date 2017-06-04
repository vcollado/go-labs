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

	// Nil
	p1 := "Ey Nil, bon dia"
	fmt.Println(p1, tasks.Nil(p1))
	p2 := "Saps que ens faran fora del pis si no paguem les factures?"
	fmt.Println(p2, tasks.Nil(p2))
	p3 := "Perdoni, voster és el llogater d'aquesta finca no?!"
	fmt.Println(p3, tasks.Nil(p3))
	p4 := ""
	fmt.Println(p4, tasks.Nil(p4))
}
