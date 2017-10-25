package main

import (
	"fmt"

	testutil "github.com/zacscodingclub/algorithms-in-go/test-util"
)

func findSmallest(values []int) int {
	smallest := values[0]
	smallestIndex := 0

	for i := 1; i < len(values); i++ {
		if values[i] < smallest {
			smallest = values[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func selectionSort(values []int) []int {
	collector := make([]int, len(values))
	totalLength := len(values)

	for i := 0; i < totalLength; i++ {
		smallest := findSmallest(values)
		collector[i] = values[smallest]

		values = append(values[:smallest], values[smallest+1:]...)
	}

	return collector
}

func main() {
	a := []int{12, 2, 15, 9, 8, 200, 1, 11, 20, 23, 3, 4, 13, 18, 17, 21, 29, 33, 55, 70}

	sorted := selectionSort(a)

	testCase := []int{1, 2, 3, 4, 8, 9, 11, 12, 13, 15, 17, 18, 20, 21, 23, 29, 33, 55, 70, 200}
	fmt.Println("expect true: ", testutil.ExpectSlice(sorted, testCase))
}
