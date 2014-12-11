/*
Author:         Eric Hebert
Description:    Generates the Mandelbrot fractal to a file named mandelbrot.png
                    with the specified resolution and iterations.
Usage:          go run a6.go resolution iterations

"Mandelbrot Set." Rosetta Code. N.p., 10 Dec. 2014. Web. 10 Dec. 2014.
<http://rosettacode.org/wiki/Mandelbrot_set#Go>.
*/

package main

import (
    "fmt"
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
    "strconv"
)

func mandelbrot(x int, y int, resolution int, iterations int) bool {
    // Calculate the complex number to start at
    start := complex(float64(x) / float64(resolution) - 2, float64(y) / float64(resolution) - 1)
    
    // Iterate until z is 2 or greater
    i := 0
    for z := start; cmplx.Abs(z) < 2 && i < iterations; i++ {
        z = (z * z) + start
    }
    return (iterations == i)
}

func generateFractal(resolution int, iterations int) {
    // Calculate resolution
    width := resolution * 3
    height := resolution * 2

    // Setup bounds and image
    bounds := image.Rect(0, 0, width, height)
    pixels := image.NewNRGBA(bounds)

    // Generate each pixel of the fractal
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            if mandelbrot(x, y, resolution, iterations) {
                pixels.Set(x, y, color.White)
            } else {
                pixels.Set(x, y, color.Black)
            }
        }
    }

    // Create a PNG file
    f, err := os.Create("mandelbrot.png")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    // Store the pixels to the image file
    if err = png.Encode(f, pixels); err != nil {
        fmt.Println(err)
    }
    if err = f.Close(); err != nil {
        fmt.Println(err)
    }
    return
}

func main() {
    // Check command line arguments
    if len(os.Args) != 3 {
        fmt.Printf("Usage: %s resolution iterations\n", os.Args[0])
        return
    }

    // Read command line arguments
    resolution, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    iterations, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println(err)
        return
    }

    // Generate the fractal
    generateFractal(resolution, iterations)
}
