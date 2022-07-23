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
	size  = 100
	delay = 20
	// image canvas covers [-size..+size]
	// number of animation frames
	// delay between frames in 10ms units
)

var palette2 = []color.Color{color.White, color.Black}
var anim = gif.GIF{}
var width = 10
var left = width / 2
var arrLen = 10
var deltaHeight = 2 * size / arrLen

func main() {
	arr := generateRandomArray2(arrLen)
	draw(os.Stdout, arr)
}

func generateRandomArray2(n int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(n)
}

func mergesort(arr []int) {
	n := len(arr)
	if n == 1 {
		return
	}
	l1 := append([]int(nil), arr[:n/2]...)
	l2 := append([]int(nil), arr[n/2:]...)
	mergesort(l1)
	mergesort(l2)
	merge(arr, l1, l2)
}

func merge(arr []int, l1, l2 []int) {
	//c := make([]int, len(l1)+len(l2))
	i := 0
	j := 0
	k := 0
	for i < len(l1) && j < len(l2) {
		if l1[i] < l2[j] {
			arr[k] = l1[i]
			i++
		} else {
			arr[k] = l2[j]
			j++
		}
		k++
	}
	for i < len(l1) {
		arr[k] = l1[i]
		k++
		i++
	}
	for j < len(l2) {
		arr[k] = l2[j]
		k++
		j++
	}
	drawArray2(arr)
}

func draw(out io.Writer, arr []int) {
	drawArray2(arr)
	mergesort(arr)
	drawArray2(arr)
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func drawArray2(arr []int) {
	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	img := image.NewPaletted(rect, palette2)
	for i := 0; i < len(arr); i++ {
		x1 := 0 + width*2*i
		x2 := x1 + width
		height := (arr[i] + 1) * deltaHeight
		rect1(left+x1, 0, left+x2, height, color.Black, img)
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

// VLine draws a veritcal line
func vLine(x, y1, y2 int, col color.Color, img *image.Paletted) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func rect1(x1, y1, x2, y2 int, col color.Color, img *image.Paletted) {
	y1, y2 = img.Rect.Max.Y-y2, img.Rect.Max.Y-y1

	hLine(x1, y1, x2, col, img)
	hLine(x1, y2, x2, col, img)
	vLine(x1, y1, y2, col, img)
	vLine(x2, y1, y2, col, img)
}
