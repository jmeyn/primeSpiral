package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type coord struct {
	X, Y, val int
}

func main() {
	const SIZE, LEN = 121, 11

	img := image.NewRGBA(image.Rect(0, 0, LEN, LEN))

	for i, lim := 0, 1; i < SIZE-1; lim++ {

	}

	f, err := os.Create("draw.png")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	png.Encode(f, img)
}
