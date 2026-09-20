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

func ImageToRGBFloat(src image.Image) [][]palette.RGBFloat {
	bounds := src.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()

	buffer := make([][]palette.RGBFloat, h)

	for y := range h {
		buffer[y] = make([]palette.RGBFloat, w)
		for x := range w {
			srcX := bounds.Min.X + x
			srcY := bounds.Min.Y + y
			pixelColor := src.At(srcX, srcY)

			r, g, b, _ := pixelColor.RGBA()

			buffer[y][x] = palette.RGBFloat{
				R: float64(r >> 8),
				G: float64(g >> 8),
				B: float64(b >> 8),
			}
		}
	}

	return buffer

}
