package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"os"
)

var palette = []color.Color{color.White, color.Black}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		size    = 100
		nframes = 64
		delay   = 8
		// image canvas covers [-size..+size]
		// number of animation frames
		// delay between frames in 10ms units
	)
	//freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	shift := 0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		Rect(0+shift, 0, shift+15, 15, color.Black, img)
		shift++
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
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
	HLine(x1, y1, x2, col, img)
	HLine(x1, y2, x2, col, img)
	VLine(x1, y1, y2, col, img)
	VLine(x2, y1, y2, col, img)
}
