package main

import (
	"fmt"
	"math"
	"image"
	"image/png"
	"os"
	"image/color"
)

type coord struct {
	X, Y, val int
}

func isPrime(guess uint32) bool {
	if guess % 2 == 0 {
		if guess == 2 {
			return true
		}
		return false
	}

	for i := uint32(3); i < uint32(math.Sqrt(float64(guess))) + 1; i += 2 {
		if guess % i == 0 {
			return false
		}
	}

	if guess == 1 {
		return false
	}

	return true
}

func main() {
	//primes := []int{2}
	const SIZE, LEN = 1234321, 1111
	grid := make([]coord, SIZE)
	var table [LEN][LEN]int

	x, y := 0, 0
	direction := true
	img := image.NewRGBA(image.Rect(0, 0, LEN, LEN))

	grid[0] = coord{x, y, 1}

	for i, lim := 0, 1; i < SIZE-1; lim++ {
		if direction == true {
			for k := 0; k < lim && i < SIZE-1; k++ {
				x++
				i++
				grid[i] = coord{x, y, i + 1}
			}
			for k := 0; k < lim && i < SIZE-1; k++ {
				y++
				i++
				grid[i] = coord{x, y, i + 1}
			}
		} else {
			for k := 0; k < lim && i < SIZE-1; k++ {
				x--
				i++
				grid[i] = coord{x, y, i + 1}
			}
			for k := 0; k < lim && i < SIZE-1; k++ {
				y--
				i++
				grid[i] = coord{x, y, i + 1}
			}
		}
		direction = !direction
	}

	offset := int(LEN/2)
	for i := 0; i < len(grid); i++ {
		table[offset - grid[i].Y][offset + grid[i].X] = grid[i].val
		if isPrime(uint32(grid[i].val)) == true {
			if grid[i].val == 2 {
				img.Set(offset + grid[i].X, offset - grid[i].Y, color.Gray16{0})
			} else {
				img.Set(offset + grid[i].X, offset - grid[i].Y, color.RGBA{0, 255, 0, 255})
			}
		} else {
			img.Set(offset + grid[i].X, offset - grid[i].Y, color.Gray16{0xffff})
		}
	}

	f, err := os.Create("draw.png")

	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer f.Close()
	png.Encode(f, img)
}
