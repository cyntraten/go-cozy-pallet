package remap

import (
	"image/color"

	"github.com/cyntraten/go-cozy-pallet/internal/palette"
)

func ColorDistanceSq(c1 color.Color, c2 color.RGBA) uint32 {
	r, g, b, _ := c1.RGBA()
	r1 := r >> 8
	g1 := g >> 8
	b1 := b >> 8

	dr := int32(r1) - int32(c2.R)
	dg := int32(g1) - int32(c2.G)
	db := int32(b1) - int32(c2.B)

	sumDistance := dr*dr + dg*dg + db*db

	return uint32(sumDistance)
}

func FindClosest(c color.Color, p palette.Palette) color.RGBA {
	if len(p) == 0 {
		return color.RGBA{A: 255}
	}

	bestColor := p[0]

	minDist := ColorDistanceSq(c, bestColor)

	for _, candidate := range p[1:] {
		currentDist := ColorDistanceSq(c, candidate)

		if currentDist < minDist {
			minDist = currentDist
			bestColor = candidate
		}
	}

	return bestColor
}
