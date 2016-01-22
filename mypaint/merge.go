package mypaint

import (
	"image"
	"image/draw"
)

func (p *Drawer) PasteOn(img draw.Image) {
	bounds := img.Bounds()
	draw.Draw(p.Img, bounds, img, bounds.Min, draw.Over)
}

func (p *Drawer) PasteAt(img draw.Image, pt P) {
	flr := pt.Floor()
	bounds := img.Bounds()
	dp := image.Point{flr[0], flr[1]}
	rec := bounds.Sub(bounds.Min).Add(dp)
	draw.Draw(p.Img, rec, img, bounds.Min, draw.Over)
}

func (p *Drawer) NoAddMerge(img draw.Image) {
	dest := p.Img
	srcBounds := img.Bounds()
	destB := dest.Bounds()
	destRect := srcBounds.Sub(srcBounds.Min).Add(destB.Min)
	source := img
	sourcePt := srcBounds.Min
	maskPt := destRect.Min
	mask := p.Img
	draw.DrawMask(dest, destRect, source, sourcePt, mask, maskPt, draw.Over)
}

func (p *Drawer) ImgCopy() draw.Image {
	bounds := p.Img.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, p.Img, bounds.Min, draw.Src)
	return img
}

func (p *Drawer) MaskAt(pt P, sourceD, maskD *Drawer) {
	dest := p.Img
	source := sourceD.Img
	mask := maskD.Img

	rect := source.Bounds()
	rect = rect.Sub(rect.Min)

	floor := pt.Floor()
	destPt := image.Point{floor[0], floor[1]}

	destRect := rect.Add(destPt)

	sourcePt := source.Bounds().Min
	maskPt := mask.Bounds().Min
	//
	draw.DrawMask(dest, destRect, source, sourcePt, mask, maskPt, draw.Over)
}
