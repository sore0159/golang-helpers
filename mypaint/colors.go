package mypaint

import (
	//"fmt"
	"image/color"
)

var COLORS = [...]string{
	"white", "black", "charcoal", "grey", "auburn", "yellow", "golden",
	"pale", "amber", "tawny", "tan", "ochre", "olive", "khaki", "brown",
	"blue", "green", "purple", "red", "pink", "cherry", "strawberry",
}

func Pallate(name string) color.RGBA {
	switch name {
	case "shade":
		return color.RGBA{0, 0, 0, 40}
	case "whiteshade":
		return color.RGBA{0xff, 0xff, 0xff, 40}
	case "white":
		return color.RGBA{0xff, 0xff, 0xff, 255}
	case "black", "charcoal":
		return color.RGBA{0, 0, 0, 255}
	case "blackDarker":
		return color.RGBA{0x10, 0x10, 0x20, 255}
	case "lightgrey", "lightgray":
		return color.RGBA{200, 200, 200, 255}
	case "grey", "gray":
		return color.RGBA{100, 100, 100, 255}
	case "auburn", "yellow":
		return color.RGBA{0xf0, 0xa3, 0x0a, 255}
	case "golden":
		return color.RGBA{0xff, 0x99, 0x33, 255}
	case "pale":
		return color.RGBA{0xff, 0xcc, 0x66, 255}
	case "amber":
		return color.RGBA{0xbe, 0x72, 0x3c, 255}
	case "tawny":
		return color.RGBA{0xfa, 0x58, 0x00, 255}
	case "tan":
		return color.RGBA{0xca, 0xab, 0x67, 255}
	case "ochre":
		return color.RGBA{0xb5, 0x8a, 0x3f, 255}
	case "olive", "khaki":
		return color.RGBA{0x99, 0x33, 0x00, 255}
	case "brown":
		return color.RGBA{0x66, 0x33, 0x00, 255}
	case "blue":
		return color.RGBA{0x00, 0x33, 0xcc, 255}
	case "navy":
		return color.RGBA{0x00, 0x00, 0x66, 255}
	case "green":
		return color.RGBA{0x00, 0x66, 0x00, 255}
	case "purple":
		return color.RGBA{0x99, 0x33, 0x99, 255}
	case "red":
		return color.RGBA{0x99, 0, 0, 255}
	case "pink":
		return color.RGBA{0xff, 0x71, 0x71, 255}
	case "cherry", "strawberry":
		//return color.RGBA{0xff, 0x22, 0x00, 255}
		return color.RGBA{0xaf, 0x2f, 0x1f, 255}
	default:
		return color.RGBA{0xff, 0xcc, 0x66, 255}
		//panic("Unknown color name " + name + " !")
	}
}

func Lighter(c color.RGBA, s float64) color.RGBA {
	if s == 1 {
		return c
	}
	if (c == color.RGBA{0, 0, 0, 255}) {
		if s < 1 {
			return color.RGBA{0x05, 0x05, 0x20, 255}
		} else {
			return color.RGBA{0x10, 0x10, 0x30, 255}
		}
	}
	if s > 1 && (c == color.RGBA{0xff, 0xff, 0xff, 255}) {
		return color.RGBA{0xdf, 0xdf, 0xff, 255}
	}
	r1, g1, b1, a1 := c.RGBA()
	cl := make([]uint8, 4)
	for i, n := range []uint32{r1, g1, b1, a1} {
		if i == 3 {
			cl[i] = uint8(n)
			continue
		}
		x := float64(uint8(n)) * s
		if x > 255 {
			x = 255
		}
		if x == 0 && s > 1 {
			x = 255. * (s - 1)
		}
		cl[i] = uint8(x)
	}
	return color.RGBA{R: cl[0], G: cl[1], B: cl[2], A: cl[3]}

}

func LightC(name string, s float64) color.RGBA {
	return Lighter(Pallate(name), s)
}
