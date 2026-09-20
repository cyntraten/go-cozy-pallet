package palette

import "image/color"

type Palette []color.RGBA

type RGBFloat struct {
	R, G, B float64
}

var GruvboxDark = Palette{
	// Dark Backgrounds / Shades
	{R: 40, G: 40, B: 40, A: 255},    // bg0 (#282828)
	{R: 50, G: 48, B: 47, A: 255},    // bg1 (#32302f)
	{R: 60, G: 56, B: 54, A: 255},    // bg2 (#3c3836)
	{R: 80, G: 73, B: 69, A: 255},    // bg3 (#504945)
	{R: 146, G: 131, B: 116, A: 255}, // gray (#928374)

	// Light Foreground / Text
	{R: 235, G: 219, B: 178, A: 255}, // fg (#ebdbb2)

	// Normal Accent Colors
	{R: 204, G: 36, B: 29, A: 255},   // red (#cc241d)
	{R: 152, G: 151, B: 26, A: 255},  // green (#98971a)
	{R: 215, G: 153, B: 33, A: 255},  // yellow (#d79921)
	{R: 69, G: 133, B: 136, A: 255},  // blue (#458588)
	{R: 177, G: 98, B: 134, A: 255},  // purple (#b16286)
	{R: 104, G: 157, B: 106, A: 255}, // aqua (#689d6a)
	{R: 214, G: 93, B: 14, A: 255},   // orange (#d65d0e)

	// Bright Accent Colors
	{R: 251, G: 73, B: 52, A: 255},   // bright red (#fb4934)
	{R: 184, G: 187, B: 38, A: 255},  // bright green (#b8bb26)
	{R: 250, G: 189, B: 47, A: 255},  // bright yellow (#fabd2f)
	{R: 131, G: 165, B: 152, A: 255}, // bright blue (#83a598)
	{R: 211, G: 134, B: 155, A: 255}, // bright purple (#d3869b)
	{R: 142, G: 192, B: 124, A: 255}, // bright aqua (#8ec07c)
	{R: 254, G: 128, B: 25, A: 255},  // bright orange (#fe8019)
}
