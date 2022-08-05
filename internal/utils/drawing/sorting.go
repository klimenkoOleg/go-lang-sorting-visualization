package drawing

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

const (
	size  = 500
	delay = 8
	// image canvas covers [-size..+size]
	// delay between frames in 10ms units
)

var palette2 = []color.Color{color.White, color.Black, color.RGBA{0xff, 0x00, 0x00, 0xff}}
var width = 10
var left = width / 2
var arrLen = 50
var totalHeight = 2 * size
var deltaHeight = totalHeight / arrLen

func SaveToOutput(anim *gif.GIF) {
	gif.EncodeAll(os.Stdout, anim)
}

func DrawArray(anim *gif.GIF, globalArr []int, sortedGlobalArr []int, lIdx, rIdx int, points ...int) {
	if points != nil {
		sortedGlobalArr[points[1]] = globalArr[points[1]-lIdx]
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
			Rect(left+x1, 0, left+x2, height, color.Black, img, false)
		} else {
			//height = (arr[i-lIdx] + 1) * deltaHeight
			height = (sortedGlobalArr[i] + 1) * deltaHeight
			Rect(left+x1, 0, left+x2, height, color.Black, img, true)
		}
	}
	for _, point := range points {
		x1 := 0 + width*2*point
		x2 := x1 + width
		Rect(left+x1, 0, left+x2, 10, color.RGBA{0xff, 0, 0, 0xff}, img, true)
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

// VLine draws a vertical line
func VLine(x, y1, y2 int, col color.Color, img *image.Paletted) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int, col color.Color, img *image.Paletted, filled bool) {
	y1, y2 = img.Rect.Max.Y-y2, img.Rect.Max.Y-y1
	HLine(x1, y1, x2, col, img)
	HLine(x1, y2, x2, col, img)
	VLine(x1, y1, y2, col, img)
	if filled {
		for i := x1; i < x2; i++ {
			VLine(i, y1, y2, col, img)
		}
	}
	VLine(x2, y1, y2, col, img)
}
