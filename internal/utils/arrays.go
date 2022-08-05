package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(n)
	return arr
}

func CreateSlice(size int) []int {
	arr := make([]int, size)
	return arr
}
