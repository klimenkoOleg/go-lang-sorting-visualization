package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	size  = 500
	delay = 8
	// image canvas covers [-size..+size]
	// delay between frames in 10ms units
)

var palette2 = []color.Color{color.White, color.Black, color.RGBA{0xff, 0x00, 0x00, 0xff}}
var anim = gif.GIF{}
var width = 10
var left = width / 2
var arrLen = 50
var totalHeight = 2 * size
var deltaHeight = totalHeight / arrLen

var globalArr []int
var sortedGlobalArr []int

func main() {
	generateRandomArray2(arrLen)
	draw(os.Stdout)
}

func generateRandomArray2(n int) {
	rand.Seed(time.Now().Unix())
	globalArr = rand.Perm(n)
	sortedGlobalArr = make([]int, n)
}

func mergesort(arr []int, lIdx, rIdx int) {
	n := len(arr)
	if n == 1 {
		return
	}
	mid := n / 2
	l1 := append([]int(nil), arr[:mid]...)
	l2 := append([]int(nil), arr[mid:]...)
	mergesort(l1, lIdx, lIdx+mid)
	mergesort(l2, lIdx+mid, rIdx)
	merge(arr, l1, l2, lIdx, rIdx)
}

func merge(arr []int, l1, l2 []int, lIdx, rIdx int) {
	i := 0
	j := 0
	k := 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			arr[k] = l1[i]
			drawArray2(arr, lIdx, rIdx, lIdx+i, lIdx+k)
			i++
		} else {
			arr[k] = l2[j]
			drawArray2(arr, lIdx, rIdx, lIdx+j, lIdx+k)
			j++
		}
		k++
	}
	for i < len(l1) {
		arr[k] = l1[i]
		drawArray2(arr, lIdx, rIdx, lIdx+i, lIdx+k)
		k++
		i++
	}
	for j < len(l2) {
		arr[k] = l2[j]
		drawArray2(arr, lIdx, rIdx, lIdx+j, lIdx+k)
		k++
		j++
	}
}

func draw(out io.Writer) {
	drawArray2(globalArr, 0, len(globalArr))
	mergesort(globalArr, 0, len(globalArr))
	drawArray2(globalArr, 0, len(globalArr))
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func drawArray2(arr []int, lIdx, rIdx int, points ...int) {
	if points != nil {
		sortedGlobalArr[points[1]] = arr[points[1]-lIdx]
	}
	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	img := image.NewPaletted(rect, palette2)
	//for i := 0; i < len(arr); i++ {
	for i := 0; i < len(globalArr); i++ {
		x1 := 0 + width*2*i
		x2 := x1 + width
		var height int
		//if i < lIdx || i >= rIdx {
		if i >= rIdx {
			height = (globalArr[i] + 1) * deltaHeight
			rect1(left+x1, 0, left+x2, height, color.Black, img, false)
		} else {
			//height = (arr[i-lIdx] + 1) * deltaHeight
			height = (sortedGlobalArr[i] + 1) * deltaHeight
			rect1(left+x1, 0, left+x2, height, color.Black, img, true)
		}
	}
	for _, point := range points {
		x1 := 0 + width*2*point
		x2 := x1 + width
		rect1(left+x1, 0, left+x2, 10, color.RGBA{0xff, 0, 0, 0xff}, img, true)
	}
	anim.Delay = append(anim.Delay, delay)
	anim.Image = append(anim.Image, img)
}

// HLine draws a horizontal line
func hLine(x1, y, x2 int, col color.Color, img *image.Paletted) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a vertical line
func vLine(x, y1, y2 int, col color.Color, img *image.Paletted) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func rect1(x1, y1, x2, y2 int, col color.Color, img *image.Paletted, filled bool) {
	y1, y2 = img.Rect.Max.Y-y2, img.Rect.Max.Y-y1
	hLine(x1, y1, x2, col, img)
	hLine(x1, y2, x2, col, img)
	vLine(x1, y1, y2, col, img)
	if filled {
		for i := x1; i < x2; i++ {
			vLine(i, y1, y2, col, img)
		}
	}
	vLine(x2, y1, y2, col, img)
}
