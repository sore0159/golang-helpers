package mydraw

import (
	"github.com/llgcode/draw2d"
	"math"
)

type Path struct {
	*draw2d.Path
}

func NewPath() *Path {
	return &Path{
		&draw2d.Path{
			//
			Components: []draw2d.PathCmp{},
			Points:     []float64{},
		},
	}
}

func (p *Path) AddPaths(paths ...*Path) {
	for i, p2 := range paths {
		p.Path.Points = append(p.Path.Points, p2.Path.Points...)
		p.Path.Components = append(p.Path.Components, p2.Path.Components...)
		if i == len(paths)-1 {
			x, y := p2.Path.LastPoint()
			p.Path.MoveTo(x, y)
		}
	}
}

func (p *Path) Copy() *Path {
	return &Path{p.Path.Copy()}
}
func (p Path) MoveTo(pt P) {
	p.Path.MoveTo(pt[0], pt[1])
}
func (p Path) LineTo(pt P) {
	p.Path.LineTo(pt[0], pt[1])
}
func (p Path) QuadCurveTo(cv, pt P) {
	p.Path.QuadCurveTo(cv[0], cv[1], pt[0], pt[1])
}
func (p Path) CubicCurveTo(cv1, cv2, pt P) {
	p.Path.CubicCurveTo(cv1[0], cv1[1], cv2[0], cv2[1], pt[0], pt[1])
}
func (p Path) Arc(cn P, r, startA, widthA float64) {
	p.Path.ArcTo(cn[0], cn[1], r, r, startA*math.Pi, widthA*math.Pi)
}

func (p Path) ArcTo(cn P, rx, ry, startA, widthA float64) {
	p.Path.ArcTo(cn[0], cn[1], rx, ry, startA*math.Pi, widthA*math.Pi)
}
func (p Path) LastP() P {
	x, y := p.LastPoint()
	return P{x, y}
}

func ConvertPaths(mypaths []*Path) (paths []*draw2d.Path) {
	paths = make([]*draw2d.Path, len(mypaths))
	for i, p := range mypaths {
		paths[i] = p.Path
	}
	return
}
