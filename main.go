package main

import (
    "fmt"
    //"image"
    //"image/color"
    //"image/png"
    //"math"
    //"os"
)

//brute forces to find if number is prime
func isPrime(guess uint) bool {
    var i uint

    if(guess % 2 == 0) {
        return false
    }

    for i = 3; i < guess; i += 2 {
        if(guess % i == 0) {
            return false
        }
    }
    return true
}

func main() {
    primes := [50000]uint{2}
    j := 0
    var i uint
    //img := image.NewRGBA(image.Rect(0, 0, 1000, 1000))

    for i = 3; i < uint(200000); i += 2 {
        if(isPrime(i) == true) {
            primes[j] = i
            j++
        }
    }
    
    for _, i := range primes {
        if(i > 0) {
            fmt.Println(i)
        }
    }
    /* f, err := os.OpenFile("rgb.png", os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }

    defer f.Close()
    png.Encode(f, img)*/
}