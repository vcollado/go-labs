package test

import (
	"testing"

	"github.com/vcollado/tasks"
)

func TestAcronym(t *testing.T) {

	type testAcronymCase struct {
		in   string
		want string
	}

	case1 := testAcronymCase{"I am Your Gopher!", "IYG"}
	case2 := testAcronymCase{"This is Just A TesT", "TJATT"}
	cases := []testAcronymCase{case1, case2}

	for _, c := range cases {
		got := tasks.Acronym(c.in)

		if got != c.want {
			t.Errorf("Acronym(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
