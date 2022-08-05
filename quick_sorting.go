package main

import (
	"os"
)

func main() {
	generateRandomArray2(arrLen)
	draw(os.Stdout)
}

func quickSort(arr []int) {
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, first, last int) {
	if first < last {
		pivot := partition(arr, first, last)
		quickSortHelper(arr, first, pivot-1)
		quickSortHelper(arr, pivot+1, last)
	}
}

func partition(arr []int, first, last int) int {
	pivotvalue := arr[first]

	leftmark := first + 1
	rightmark := last
	done := false

	for !done {
		for leftmark <= rightmark && arr[leftmark] <= pivotvalue {
			leftmark = leftmark + 1
		}
		for arr[rightmark] >= pivotvalue && rightmark >= leftmark {
			rightmark = rightmark - 1
		}
		if rightmark < leftmark {
			done = true
		} else {
			arr[leftmark], arr[rightmark] = arr[rightmark], arr[leftmark]
		}
	}
	arr[first], arr[rightmark] = arr[rightmark], arr[first]

	return rightmark
}
