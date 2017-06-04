package tests

import (
	"testing"

	"github.com/vcollado/tasks"
)

func TestHamming(t *testing.T) {

	// define the input data for the hamming function test
	type testHammingCase struct {
		in   []string
		want string
	}

	case1 := testHammingCase{[]string{"GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"}, "^ ^ ^  ^ ^    ^^ "}
	case2 := testHammingCase{[]string{"GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCC"}, "^ ^ ^  ^ ^    ^^^"}
	cases := []testHammingCase{case1, case2}

	for _, c := range cases {
		got := tasks.Hamming(c.in[0], c.in[1])

		if got != c.want {
			t.Errorf("Accumulate(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
