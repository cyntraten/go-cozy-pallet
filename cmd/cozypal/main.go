package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"

	"github.com/cyntraten/go-cozy-pallet/internal/palette"
	"github.com/cyntraten/go-cozy-pallet/internal/remap"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("Doesn not have args, you have: %d, need: 3\nUse cozypal <input image> <output image>\n", len(args))
		os.Exit(1)
	}

	inputPath := args[1]
	outputPath := args[2]
	fmt.Printf("Program name %v\n", args[0])
	fmt.Printf("Image src: %v\n", inputPath)
	fmt.Printf("Image output: %v\n", outputPath)

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("Open image error: %s\n", err)
	}

	defer file.Close()

	img, formatName, err := image.Decode(file)

	if err != nil {
		log.Fatalf("Decode image error: %s\n", err)
	}

	imageInfo := img.Bounds()

	fmt.Printf("Image format is: %v, Width(Dx): %v, Height(Dy): %v \n", formatName, imageInfo.Dx(), imageInfo.Dy())

	outputImg := remap.ApplyPalette(img, palette.GruvboxDark)
	outFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Create image error: %s\n", err)
	}

	defer outFile.Close()

	err = png.Encode(outFile, outputImg)

	if err != nil {
		log.Fatalf("Encode image error: %s\n", err)
	}

}
