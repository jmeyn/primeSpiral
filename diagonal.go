package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func isPrime(guess int) bool {
	if guess%2 == 0 {
		if guess == 2 {
			return true
		}

		return false
	}

	for i := 3; i < int(math.Sqrt(float64(guess))+1); i += 2 {
		if i == 5 {
			continue
		}
		if guess%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	const SIZE, LEN = 1000000, 1000
	/*red := 255
	green := 0
	blue := 0*/
	n := 2
	img := image.NewRGBA(image.Rect(0, 0, LEN, LEN))

	for i := 0; i < LEN-1; i++ {
		for j, k := 0, i; k >= 0; {
			if isPrime(n) == true {
				img.Set(j, k, color.RGBA{0, 0, 0, 255})
			} else {
				img.Set(j, k, color.RGBA{255, 0, 255, 255})
			}

			/*
				if red == 255 && blue == 0 && green == 0 {
					blue = 0
					red = 0
					green += 255
				} else if red == 0 && blue == 0 && green == 255 {
					blue += 255
					red = 0
					green = 0
				} else {
					blue = 0
					red += 255
					green = 0
				}*/

			j++
			k--
			n++
		}
	}
	for i := 0; i < LEN; i++ {
		for j, k := i, LEN-1; k >= i; {
			if isPrime(n) == true {
				img.Set(j, k, color.RGBA{0, 0, 0, 255})
			} else {
				img.Set(j, k, color.RGBA{255, 0, 255, 255})
			}

			/*if red == 255 && blue == 0 && green == 0 {
				blue = 0
				red = 0
				green += 255
			} else if red == 0 && blue == 0 && green == 255 {
				blue += 255
				red = 0
				green = 0
			} else {
				blue = 0
				red += 255
				green = 0
			}*/

			j++
			k--
			n++
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
