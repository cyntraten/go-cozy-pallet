package remap

import (
	"image"

	"github.com/cyntraten/go-cozy-pallet/internal/palette"
)

func ApplyPalette(src image.Image, p palette.Palette) *image.RGBA {

	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			origColor := src.At(x, y)
			newColor := FindClosest(origColor, p)
			dst.Set(x, y, newColor)
		}

	}

	return dst

}
