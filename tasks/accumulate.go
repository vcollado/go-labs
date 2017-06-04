// Implement the accumulate operation, which, given a collection and an operation to perform on each element of the collection,
// returns a new collection containing the result of applying that operation to each element of the input collection.

// reference: https://stackoverflow.com/questions/3601796/

package tasks

// Accumulator accumulates two values
type Accumulator func(int) int

// Square two numbers
func Square(x int) int {
	return x * x
}

// accumulate a collection given an accumulator
func accumulate(coll []int, fn Accumulator) []int {
	result := make([]int, 0)
	for _, num := range coll {
		result = append(result, fn(num))
	}
	return result
}

// Accumulate given a collection applying an accumulator
func Accumulate(coll []int, fn Accumulator) []int {
	return accumulate(coll, fn)
}
