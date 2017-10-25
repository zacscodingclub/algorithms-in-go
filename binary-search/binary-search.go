package main

import (
	"fmt"

	"github.com/zacscodingclub/algorithms-in-go/test-util"
)

// for integers
func binarySearch(values []int, search int) int {
	min := 0
	max := len(values) - 1

	for min <= max {
		mid := (min + max) / 2
		guess := values[mid]

		if guess == search {
			return mid
		} else if guess < search {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}

func main() {
	a := []int{1, 5, 22, 33, 40}
	b := []int{1, 2, 3, 4, 5}

	aSearch := binarySearch(a, 22)
	bSearch := binarySearch(b, 22)
	cSearch := binarySearch(a, 1)
	dSearch := binarySearch(a, 33)

	fmt.Println("expect 2: ", testutil.Expect(2, aSearch))
	fmt.Println("expect -1: ", testutil.Expect(-1, bSearch))
	fmt.Println("expect 0: ", testutil.Expect(0, cSearch))
	fmt.Println("expect 3: ", testutil.Expect(3, dSearch))

}
