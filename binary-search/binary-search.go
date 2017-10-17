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

func main() {
	a := []int{1, 5, 22, 33, 40}
	b := []int{1, 2, 3, 4, 5}
	aSearch := binarySearch(a, 22)

	bSearch := binarySearch(b, 22)
	fmt.Println("expect truthy: ", aSearch)
	fmt.Println("expect falsey: ", bSearch)
}
