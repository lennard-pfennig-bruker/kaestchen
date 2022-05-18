package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	width := 200
	height := 100

	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	x, y := 0, 0

	blackIt := func(percentage int) {
		if rand.Intn(100) < percentage {
			img.Set(x, y, color.Black)
		} else {
			img.Set(x, y, color.White)
		}
	}

	encode := func() {
		// Encode as PNG.
		f, _ := os.Create("board.png")
		err := png.Encode(f, img)
		if err != nil {
			log.Fatal(err)
		}
	}

	bruh := 0

	for x = 0; x < width; x++ {
		for y = 0; y < height; y++ {

			n1 := isItBlack(img.At(x, y-1).RGBA())   //   n2  n1
			n2 := isItBlack(img.At(x-1, y-1).RGBA()) //   n3  we
			n3 := isItBlack(img.At(x-1, y).RGBA())   //   n4
			n4 := isItBlack(img.At(x-1, y+1).RGBA())

			switch {
			case x == 0 && y == 0:
				blackIt(100)
			case x == width-1:
				blackIt(100)
			case y == height-1:
				blackIt(100)
			case x == 0:
				blackIt(100)
			case y == 0:
				blackIt(100)

			case n1 && n2 && !n3 && !n4:
				blackIt(10)
			case !n1 && n2 && n3 && n4:
				blackIt(10)

			case !n1 && !n2 && n3 && !n4:
				blackIt(100)
			case n1 && !n2 && !n3 && !n4:
				blackIt(100)
			case n1 && !n2 && n3:
				blackIt(100)
			case n1 && !n2 && !n3 && n4:
				blackIt(100)
			case !(n1 && n2 && n3 && n4):
				blackIt(0)
			default:
				bruh++
				blackIt(0)
			}

		}
	}

	encode()

	fmt.Println()
	fmt.Print("failed ")
	fmt.Print(bruh)
	fmt.Print("/")
	fmt.Println(width * height)
}
