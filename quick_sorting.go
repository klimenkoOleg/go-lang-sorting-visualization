package main

import (
	"go-lang-sorting-visualization/internal/utils"
	"go-lang-sorting-visualization/internal/utils/drawing"
	"image/gif"
)

var globalArr2 []int
var sortedGlobalArr2 []int

var arrLen2 = 50

func main() {
	anim := gif.GIF{}
	globalArr2 = utils.GenerateRandomArray(arrLen2)
	sortedGlobalArr2 = utils.GenerateRandomArray(arrLen2)

	quickSort(&anim, globalArr2)
	drawing.SaveToOutput(&anim)
}

func quickSort(anim *gif.GIF, arr []int) {
	quickSortHelper(anim, arr, 0, len(arr)-1)
}

func quickSortHelper(anim *gif.GIF, arr []int, first, last int) {
	if first < last {
		pivot := partition(anim, arr, first, last)
		quickSortHelper(anim, arr, first, pivot-1)
		quickSortHelper(anim, arr, pivot+1, last)
	}
}

func partition(anim *gif.GIF, arr []int, first, last int) int {
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
			drawing.DrawArray(anim, arr, arr, 0, len(arr), leftmark, rightmark)
		}
	}
	arr[first], arr[rightmark] = arr[rightmark], arr[first]
	drawing.DrawArray(anim, arr, arr, 0, len(arr), first, rightmark)
	return rightmark
}
