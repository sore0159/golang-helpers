package mypaint

import (
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	"image/draw"
)

type Filler struct {
	GC             draw2d.GraphicContext
	CurFillColor   color.RGBA
	CurStrokeColor color.RGBA
	*Path
}

func NewFiller(img draw.Image) *Filler {
	return &Filler{
		GC:   draw2dimg.NewGraphicContext(img),
		Path: NewPath(),
	}
}

func (f *Filler) SetLineWidth(x float64) {
	f.GC.SetLineWidth(x)
}

func (f *Filler) SetFillRuleWinding(x bool) {
	if x {
		f.GC.SetFillRule(draw2d.FillRuleWinding)
	} else {
		f.GC.SetFillRule(draw2d.FillRuleEvenOdd)
	}
}

func (f *Filler) Clear() {
	f.Path = NewPath()
}
func (f *Filler) TakePath(p *Path) {
	f.Path = p
}
func (f *Filler) Scale(sx, sy float64) {
	f.GC.Scale(sx, sy)
}
func (f *Filler) Translate(sx, sy float64) {
	f.GC.Translate(sx, sy)
}

func (f *Filler) Fill(mypaths ...*Path) {
	if len(mypaths) == 0 {
		f.GC.Fill(f.Path.Path)
		f.Clear()
	} else {
		f.GC.Fill(ConvertPaths(mypaths)...)
	}
}

func (f *Filler) Stroke(mypaths ...*Path) {
	if len(mypaths) == 0 {
		f.GC.Stroke(f.Path.Path)
		f.Clear()
	} else {
		f.GC.Stroke(ConvertPaths(mypaths)...)
	}
}

func (f *Filler) FillStroke(mypaths ...*Path) {
	if len(mypaths) == 0 {
		f.GC.FillStroke(f.Path.Path)
		f.Clear()
	} else {
		f.GC.FillStroke(ConvertPaths(mypaths)...)
	}
}

func (p *Filler) SetFillColor(R, G, B, A uint8) {
	c := color.RGBA{R, G, B, A}
	p.CurFillColor = c
	p.GC.SetFillColor(c)
}

func (p *Filler) SetStrokeColor(R, G, B, A uint8) {
	c := color.RGBA{R, G, B, A}
	p.CurStrokeColor = c
	p.GC.SetStrokeColor(c)
}

func (p *Filler) SetFillC(name string) {
	c := Pallate(name)
	p.CurFillColor = c
	p.GC.SetFillColor(c)
}
func (p *Filler) SetStrokeC(name string) {
	c := Pallate(name)
	p.CurStrokeColor = c
	p.GC.SetStrokeColor(c)
}
func (p *Filler) SetFillCL(name string, s float64) {
	c := Pallate(name)
	p.CurFillColor = c
	p.GC.SetFillColor(Lighter(c, s))
}
func (p *Filler) SetStrokeCL(name string, s float64) {
	c := Pallate(name)
	p.CurStrokeColor = c
	p.GC.SetStrokeColor(Lighter(c, s))
}

func (p *Filler) SetFillLight(s float64) {
	c := Lighter(p.CurFillColor, s)
	p.GC.SetFillColor(c)

}
func (p *Filler) SetStrokeLight(s float64) {
	c := Lighter(p.CurStrokeColor, s)
	p.GC.SetStrokeColor(c)
}

func (f *Filler) Pencil() {
	f.SetLineWidth(1)
	f.SetStrokeC("black")
}
func (f *Filler) IsFillWhite() bool {
	return f.CurFillColor == color.RGBA{0xff, 0xff, 0xff, 40}
}
func (f *Filler) IsFillBlack() bool {
	return f.CurFillColor == color.RGBA{0, 0, 0, 255}
}
