package main

import "fmt"

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

func expect(a int, b int) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func main() {
	a := []int{1, 5, 22, 33, 40}
	b := []int{1, 2, 3, 4, 5}

	aSearch := binarySearch(a, 22)
	bSearch := binarySearch(b, 22)
	cSearch := binarySearch(a, 1)
	dSearch := binarySearch(a, 33)

	fmt.Println("expect 2: ", expect(2, aSearch))
	fmt.Println("expect -1: ", expect(-1, bSearch))
	fmt.Println("expect 0: ", expect(0, cSearch))
	fmt.Println("expect 3: ", expect(3, dSearch))

}
