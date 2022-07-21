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

var palette = []color.Color{color.White, color.Black}

func main() {
	arr := generateRandomArray(10)
	lissajous(os.Stdout, arr)
}

func generateRandomArray(n int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(n)
}

func lissajous(out io.Writer, arr []int) {
	const (
		size  = 100
		delay = 20
		// image canvas covers [-size..+size]
		// number of animation frames
		// delay between frames in 10ms units
	)
	anim := gif.GIF{}
	width := 10
	left := width / 2
	deltaHeight := 2 * size / len(arr)
	drawArray(arr, width, deltaHeight, left, size, delay, &anim)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				drawArray(arr, width, deltaHeight, left, size, delay, &anim)
			}
		}
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func drawArray(arr []int, width int, deltaHeight int, left int, size int, delay int, anim *gif.GIF) {
	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	img := image.NewPaletted(rect, palette)
	for i := 0; i < len(arr); i++ {
		x1 := 0 + width*2*i
		x2 := x1 + width
		height := (arr[i] + 1) * deltaHeight
		Rect(left+x1, 0, left+x2, height, color.Black, img)
	}
	anim.Delay = append(anim.Delay, delay)
	anim.Image = append(anim.Image, img)
}

// HLine draws a horizontal line
func HLine(x1, y, x2 int, col color.Color, img *image.Paletted) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int, col color.Color, img *image.Paletted) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int, col color.Color, img *image.Paletted) {
	y1, y2 = img.Rect.Max.Y-y2, img.Rect.Max.Y-y1

	HLine(x1, y1, x2, col, img)
	HLine(x1, y2, x2, col, img)
	VLine(x1, y1, y2, col, img)
	VLine(x2, y1, y2, col, img)
}
