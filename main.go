package main

import (
	"bensPlatte/Board"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	b := Board.New(400, 200)

	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			switch {
			case y == 0 || x == 0:
				b.BlackIt(x, y, 100)
			default:
				b.BlackIt(x, y, 0)
			}
		}
	}
	fmt.Println("running random rect spawning...")
	counter := 50
	for {
		b.SpawnRect()
		counter -= 1
		if b.Done() || counter == 0 {
			fmt.Println("done")
			break
		}
	}

	b.SavePng()

}
