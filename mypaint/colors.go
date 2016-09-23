package mypaint

import (
	//"fmt"
	"image/color"
	"math"
)

var COLORS = [...]string{
	"navy", "pink", "pale", "tan", "ochre", "amber",

	"golden", "auburn", "yellow", "tawny", "cherry",

	"strawberry", "red", "olive", "khaki", "brown",

	"black", "charcoal", "white", "grey", "emerald",

	"blue", "green", "slate", "purple", "plum",

	"neonred", "turquoise", "chartreuse", "dodger",

	"blonde", "skinblack",
}

func Pallate(name string) color.RGBA {
	switch name {
	case "chartreuse":
		return color.RGBA{0x7f, 0xff, 0x00, 255}
	case "emerald":
		return color.RGBA{0x00, 0xb9, 0x47, 255}
	case "green":
		return color.RGBA{0x11, 0x88, 0x11, 255}
	case "turquoise":
		return color.RGBA{0x00, 0xf5, 0xff, 255}
	case "dodger":
		return color.RGBA{0x1e, 0x90, 0xff, 255}
	case "slate":
		return color.RGBA{0x84, 0x70, 0xff, 255}
	case "plum":
		return color.RGBA{0x60, 0x11, 0x2d, 255}
	case "shade":
		return color.RGBA{0, 0, 0, 40}
	case "whiteshade":
		return color.RGBA{0xff, 0xff, 0xff, 40}
	case "white":
		return color.RGBA{0xff, 0xff, 0xff, 255}
	case "black", "charcoal":
		return color.RGBA{0, 0, 0, 255}
	case "skinblack":
		return color.RGBA{10, 10, 20, 255}
	case "blackDarker":
		return color.RGBA{0x10, 0x10, 0x20, 255}
	case "lightgrey", "lightgray":
		return color.RGBA{200, 200, 200, 255}
	case "background":
		return color.RGBA{255, 230, 230, 255}
	case "grey", "gray":
		return color.RGBA{100, 100, 100, 255}
	case "blonde":
		return color.RGBA{0xff, 0xff, 0x00, 255}
	case "yellow":
		return color.RGBA{0xf0, 0xa3, 0x0a, 255}
	case "golden":
		return color.RGBA{0xff, 0x99, 0x33, 255}
	case "amber":
		return color.RGBA{0xff, 0xcc, 0x66, 255}
	case "auburn":
		return color.RGBA{0xbe, 0x72, 0x3c, 255}
	case "tawny":
		return color.RGBA{0xda, 0x38, 0x00, 255}
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
	case "purple":
		return color.RGBA{0x99, 0x33, 0x99, 255}
	case "red":
		return color.RGBA{0x99, 0, 0, 255}
	case "pale":
		return color.RGBA{0xff, 0xaa, 0xaa, 255}
	case "palepink":
		panic("NO MORE PALEPINK")
	case "pink":
		return color.RGBA{0xff, 0x71, 0x71, 255}
	case "cherry", "strawberry":
		//return color.RGBA{0xff, 0x22, 0x00, 255}
		return color.RGBA{0xaf, 0x2f, 0x1f, 255}
	default:
		return color.RGBA{0xff, 0xff, 0xff, 255}
		//return color.RGBA{0xff, 0xcc, 0x66, 255}
		//panic("Unknown color name " + name + " !")
	}
}

func Lighter(c color.RGBA, s float64) color.RGBA {
	if s == 1 {
		return c
	}
	if s < 1 {
		return MixColors(c, color.RGBA{0, 0, 0, 255}, 1-s)
	} else {
		return MixColors(c, color.RGBA{255, 255, 255, 255}, s-1)
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

func Lighter2(c color.RGBA, s float64) color.RGBA {
	return MixColors(c, color.RGBA{255, 255, 255, 255}, s)
}

func MixColors(c1, c2 color.RGBA, proportion float64) color.RGBA {
	list1 := []uint8{c1.R, c1.G, c1.B, c1.A}
	list2 := []uint8{c2.R, c2.G, c2.B, c2.A}
	list3 := make([]uint8, 4)
	for i, x := range list1 {
		fl1, fl2 := float64(x), float64(list2[i])
		fl3 := (1-proportion)*fl1 + (proportion)*fl2
		list3[i] = uint8(math.Floor(fl3))
	}
	return color.RGBA{R: list3[0], G: list3[1], B: list3[2], A: list3[3]}
}

func LightC(name string, s float64) color.RGBA {
	return Lighter(Pallate(name), s)
}
