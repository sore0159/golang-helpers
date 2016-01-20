package mydraw

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/draw"
)

type Drawer struct {
	Img draw.Image
	*Filler
}

func BlankDrawer(w, h int, cName string) *Drawer {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	if cName != "" {
		c := Pallate(cName)
		draw.Draw(img, img.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	}
	return MakeDrawer(img)
}

func MakeDrawer(img draw.Image) *Drawer {
	f := NewFiller(img)
	return &Drawer{
		Img:    img,
		Filler: f,
	}
}

func (p *Drawer) SaveToPng(name string) {
	if name == "" {
		name = "DEFAULT"
	}
	draw2dimg.SaveToPngFile(name+".png", p.Img)
}
