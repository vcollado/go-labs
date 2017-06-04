package exercism

import (
	"strings"
)

// CalculateHammingDistance between two DNA strands.
func CalculateHammingDistance(dna1 string, dna2 string) string {

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
