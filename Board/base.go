package Board

import (
	"fmt"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

// HLine draws a horizontal line
func (b board) HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		b.Img.Set(x1, y, b.col)
	}
}

// VLine draws a veritcal line
func (b board) VLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		b.Img.Set(x, y1, b.col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func (b board) Rect(x1, y1, x2, y2 int) {
	b.HLine(x1, y1, x2)
	b.HLine(x1, y2, x2)
	b.VLine(x1, y1, y2)
	b.VLine(x2, y1, y2)

	for x := x1 + 1; x < x2; x++ {
		for y := y1 + 1; y < y2; y++ {
			b.BlackIt(x, y, 0)
		}
	}
}

func isItBlack(r uint32, g uint32, b uint32, a uint32) bool {
	if r != g || r != b {
		panic("oh oh")
	}
	switch int(r / 257) {
	case 0:
		return true
	case 255:
		return false
	default:
		panic("neither is_it_black or white")
	}
}

func (b board) SavePng() {
	f, _ := os.Create("board.png")

	err := png.Encode(f, b.Img)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func (b board) Done() bool {
	for i := range b.pixelDone {
		for j := range b.pixelDone[i] {
			if !b.pixelDone[i][j] {
				return false
			}
		}
	}
	return true
}

func (b board) PrintPixelInRect() {
	for i := range b.pixelDone {
		for j := range b.pixelDone[i] {
			fmt.Print(b.pixelDone[i][j], " ")
		}
		fmt.Println()
	}
}

func (b board) BlackIt(x int, y int, percentage int) {
	if rand.Intn(100) < percentage {
		b.Img.Set(x, y, color.Black)
	} else {
		b.Img.Set(x, y, color.White)
	}
}

func (b board) MarkPixelDone(x int, y int) {
	b.pixelDone[x][y] = true
}
