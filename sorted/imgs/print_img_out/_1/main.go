package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	// Open the image file
	reader, err := os.Open("gopher.jpg")
	if err != nil {
		fmt.Printf("Error opening image: %v\n", err)
		return
	}
	defer reader.Close()

	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		return
	}

	// Convert to ASCII with resizing
	targetWidth := 80 // Desired ASCII art width
	printASCII(img, targetWidth)
}

func printASCII(img image.Image, targetWidth int) {
	bounds := img.Bounds()
	origWidth := bounds.Max.X - bounds.Min.X
	origHeight := bounds.Max.Y - bounds.Min.Y

	// Calculate new height maintaining aspect ratio
	// Multiply by 0.5 to account for terminal characters being taller than wide
	ratio := float64(origHeight) / float64(origWidth) * 0.5
	targetHeight := int(float64(targetWidth) * ratio)

	// Print ASCII art with calculated sampling
	for y := 0; y < targetHeight; y++ {
		for x := 0; x < targetWidth; x++ {
			// Map target coordinates back to original image
			srcX := x * origWidth / targetWidth
			srcY := y * origHeight / targetHeight

			// Get color at the mapped coordinate
			r, g, b, _ := img.At(srcX, srcY).RGBA()
			// Convert to grayscale
			gray := (r + g + b) / 3 >> 8
			// Print corresponding ASCII character
			fmt.Print(getASCIIChar(uint8(gray)))
		}
		fmt.Println()
	}
}

func getASCIIChar(gray uint8) string {
	// Reduced character set for better readability
	chars := []string{"@", "#", "8", "&", "o", ":", "*", ".", " "}
	index := int(gray) * len(chars) / 256
	if index >= len(chars) {
		index = len(chars) - 1
	}
	return chars[index]
}
