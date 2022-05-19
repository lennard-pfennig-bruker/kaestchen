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
	maxSize := 100
	x, y := 0, 0
	for {
		x, y = rand.Intn(b.Width-minSize), rand.Intn(b.Height-minSize)
		if !b.pixelDone[x][y] {
			break
		}
	}

	xSize := rand.Intn(maxSize - minSize)
	ySize := rand.Intn(maxSize - minSize)

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
