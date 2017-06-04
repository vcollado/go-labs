// directory structure for the tests: https://groups.google.com/d/msg/golang-nuts/3hPEtCNm3XA/lh5iaCLyAAAJ
// compare slices: // https://stackoverflow.com/a/15312182
package tests

import (
	"reflect"
	"testing"

	"github.com/vcollado/tasks"
)

func TestAccumulateWithSquare(t *testing.T) {

	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 4, 9}},
		{[]int{3, 2, 5}, []int{9, 4, 25}},
		{[]int{0, 25, 12}, []int{0, 621, 144}},
	}

	for _, c := range cases {
		got := tasks.Accumulate(c.in, tasks.Square)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Accumulate(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
