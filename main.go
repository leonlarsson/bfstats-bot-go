package main

import (
	"flag"
	"fmt"
	"image/color"
	"math/rand/v2"
	"os"
	"sync"
	"time"

	"github.com/leonlarsson/go-image-gen/engine"
)

var iterations int

// init is called before main
func init() {

	// Parse command line arguments
	flag.IntVar(&iterations, "iterations", 1, "Number of iterations")
	flag.Parse()

	// Clean up old renders
	resetRendersFolder()
}

// generateImage generates an image and saves it to disk
func generateImage(width, height, identifier int, wg *sync.WaitGroup) {
	defer wg.Done()
	scene := engine.NewScene(width, height)
	scene.EachPixel(func(x, y int) color.RGBA {
		return color.RGBA{
			uint8(x * 255 / width),
			uint8(y * 255 / height),
			100,
			255,
		}
	})
	fileName := fmt.Sprintf("%d.png", identifier+1)
	scene.Save("./renders/" + fileName)
}

// resetRendersFolder removes the renders folder and creates a new one
func resetRendersFolder() {
	os.RemoveAll("./renders")
	os.MkdirAll("./renders", os.ModePerm)
}

func main() {
	startTime := time.Now()

	var wg sync.WaitGroup

	// Run one goroutine per iteration
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go generateImage(rand.IntN(500)+1, rand.IntN(500)+1, i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Generated %d images in %s\n", iterations, time.Since(startTime))
}
