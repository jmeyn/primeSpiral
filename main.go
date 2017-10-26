package main

import (
	"fmt"
	"math"
)

type coord struct {
	X, Y, val int
}

func isPrime(guess uint32) bool {
	if guess%2 == 0 {
		return false
	}

	for i := uint32(3); i < uint32(math.Sqrt(float64(guess)))+1; i += 2 {
		if guess%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	//primes := []int{2}
	grid := make([]coord, 400)
	var table [20][20]int

	max := 400
	x, y := 0, 0
	direction := true

	grid[0] = coord{x, y, 2}

	for i, lim := 0, 1; i < max-1; lim++ {
		if direction == true {
			for k := 0; k < lim && i < max-1; k++ {
				x++
				i++
				grid[i] = coord{x, y, i + 2}
			}
			for k := 0; k < lim && i < max-1; k++ {
				y++
				i++
				grid[i] = coord{x, y, i + 2}
			}
		} else {
			for k := 0; k < lim && i < max-1; k++ {
				x--
				i++
				grid[i] = coord{x, y, i + 2}
			}
			for k := 0; k < lim && i < max-1; k++ {
				y--
				i++
				grid[i] = coord{x, y, i + 2}
			}
		}
		direction = !direction
	}

	for i := 0; i < len(grid); i++ {
		table[9 + grid[i].X][9 + grid[i].Y] = grid[i].val
	}

	for i := 0; i < 401; i++ {
		for j := 0; j < 20; j++ {
			fmt.Printf("%4d ", table[i][j])
		}
		fmt.Println()
	}
}
