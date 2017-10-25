package main

import (
	"fmt"
	"math/rand"

	testutil "github.com/zacscodingclub/algorithms-in-go/test-util"
)

func quickSort(values []int) []int {
	length := len(values)
	if length < 2 {
		return values
	}
	less, greater := 0, length-1
	pivot := rand.Int() % length

	values[pivot], values[greater] = values[greater], values[pivot]

	for i := range values {
		if values[i] < values[greater] {
			values[less], values[i] = values[i], values[less]
			less++
		}
	}

	values[less], values[greater] = values[greater], values[less]

	quickSort(values[:less])
	quickSort(values[less+1:])

	return values
}

func main() {
	a := []int{12, 2, 15, 9, 8, 200, 1, 11, 20, 23, 3, 4, 13, 18, 17, 21, 29, 33, 55, 70}
	sorted := quickSort(a)

	testCase := []int{1, 2, 3, 4, 8, 9, 11, 12, 13, 15, 17, 18, 20, 21, 23, 29, 33, 55, 70, 200}
	fmt.Println("expect true: ", testutil.ExpectSlice(sorted, testCase))
}
