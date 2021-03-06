package Board

import (
	"image"
	"image/color"
	"math/rand"
)

type board struct {
	Width     int
	Height    int
	Img       *image.RGBA
	pixelDone [][]bool
	col       color.Color
}

func New(width int, height int) *board {
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	pixelDone := make([][]bool, width)
	for i := range pixelDone {
		pixelDone[i] = make([]bool, height)
		for j := range pixelDone[i] {
			pixelDone[i][j] = false
			img.Set(i, j, color.Black)
		}
	}

	return &board{
		Width:     width,
		Height:    height,
		Img:       img,
		pixelDone: pixelDone,
		col:       color.Black,
	}
}

func (b board) SpawnRect() {
	minSize := 10
	maxSize := 80
	step := 30

	// stupid algo for sppawnfinder
	x, y := rand.Intn((b.Width)/step)*step, rand.Intn((b.Height)/step)*step

	step = 10

	xSize := rand.Intn((maxSize-minSize)/step)*step + minSize
	ySize := rand.Intn((maxSize-minSize)/step)*step + minSize

	b.Rect(x, y, x+xSize, y+ySize)

	for i := range b.pixelDone {
		for j := range b.pixelDone[i] {
			if x < i && i < x+xSize && y < j && j < y+ySize {
				b.pixelDone[i][j] = true
			}
			if j < y+ySize {
				break
			}
		}
		if i < x+xSize {
			break
		}
	}
}
