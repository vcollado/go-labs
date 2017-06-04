// Calculate the Hamming difference between two DNA strands.
// GAGCCTACTAACGGGAT
// CATCGTAATGACGGCCT
// ^ ^ ^  ^ ^    ^^

package tasks

import (
	"strings"
)

// Hamming between two DNA strands.
func Hamming(dna1 string, dna2 string) string {

	sliceDna1 := strings.Split(dna1, "")
	var sliceDna2 = strings.Split(dna2, "")
	var sliceDna1Len = len(sliceDna1)

	hammingDistance := make([]string, 0)

	for i := 0; i < sliceDna1Len; i++ {
		if sliceDna1[i] != sliceDna2[i] {
			hammingDistance = append(hammingDistance, "^")
		} else {
			hammingDistance = append(hammingDistance, " ")
		}
	}

	return strings.Join(hammingDistance, "")
}
