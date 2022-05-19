package main

import (
	"bensPlatte/Board"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func build() {
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
	counter := 100
	for {
		b.SpawnRect()
		counter -= 1
		if b.Done() || counter == 0 {
			fmt.Println("done")
			break
		}
	}

	b.SavePng("board.png")
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	build()
	fileBytes, err := ioutil.ReadFile("board.png")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write(fileBytes)
	if err != nil {
		return
	}
	return
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Serve on http://localhost:8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
