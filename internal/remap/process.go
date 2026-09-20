package remap

import (
	"image"
	"image/color"

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

func clamp(val float64) float64 {
	if val < 0.0 {
		return 0.0
	}

	if val > 255.0 {
		return 255.0
	}

	return val
}

func ApplyDithering(src image.Image, p palette.Palette) *image.RGBA {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	buff := ImageToRGBFloat(src)

	h := len(buff)
	w := len(buff[0])

	for y := range h {
		for x := range w {
			current := buff[y][x]

			clampedChannelR := clamp(current.R)
			clampedChannelG := clamp(current.G)
			clampedChannelB := clamp(current.B)

			clampedColor := color.RGBA{
				R: uint8(clampedChannelR),
				G: uint8(clampedChannelG),
				B: uint8(clampedChannelB),
				A: 255,
			}

			closest := FindClosest(clampedColor, p)

			dst.Set(bounds.Min.X+x, bounds.Min.Y+y, closest)

			//Dithering strength
			strength := 1.25

			errR := current.R - float64(closest.R)*strength
			errG := current.G - float64(closest.G)*strength
			errB := current.B - float64(closest.B)*strength

			if x+1 < w {
				buff[y][x+1].R += errR * 7.0 / 16.0
				buff[y][x+1].G += errG * 7.0 / 16.0
				buff[y][x+1].B += errB * 7.0 / 16.0
			}

			if x-1 >= 0 && y+1 < h {
				buff[y+1][x-1].R += errR * 3.0 / 16.0
				buff[y+1][x-1].G += errG * 3.0 / 16.0
				buff[y+1][x-1].B += errB * 3.0 / 16.0
			}

			if y+1 < h {
				buff[y+1][x].R += errR * 5.0 / 16.0
				buff[y+1][x].G += errG * 5.0 / 16.0
				buff[y+1][x].B += errB * 5.0 / 16.0
			}

			if y+1 < h && x+1 < w {
				buff[y+1][x+1].R += errR * 1.0 / 16.0
				buff[y+1][x+1].G += errG * 1.0 / 16.0
				buff[y+1][x+1].B += errB * 1.0 / 16.0
			}

		}
	}

	return dst

}
