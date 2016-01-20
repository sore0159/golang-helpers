package mydraw

func (path *Path) AngledEllipse(cn P, rx, ry, angle float64) {
	pt1, pt2 := cn.Bumpers(rx, angle)
	const scale = 1.22
	cv1, cv2 := pt1.GoPolar(scale*ry, angle), pt2.GoPolar(scale*ry, angle)
	path.MoveTo(pt1)
	path.CubicCurveTo(cv1, cv2, pt2)
	cv1, cv2 = pt1.GoPolar(scale*ry, angle+1), pt2.GoPolar(scale*ry, angle+1)
	path.CubicCurveTo(cv2, cv1, pt1)
	path.Close()
}

func (p *Path) Ellipse(cn P, rx, ry float64) {
	p.MoveTo(cn.Add(rx, 0))
	p.ArcTo(cn, rx, ry, 0, 2)
	p.Close()
}

func (p *Path) Circle(cn P, r float64) {
	p.MoveTo(cn.Add(r, 0))
	p.ArcTo(cn, r, r, 0, 2)
	p.Close()
}

func (p *Path) Rect(ul, lr P) {
	if ul[0] > lr[0] {
		ul[0], lr[0] = lr[0], ul[0]
	}
	if ul[1] > lr[1] {
		ul[1], lr[1] = lr[1], ul[1]
	}
	p.MoveTo(ul)
	p.LineTo(P{ul[0], lr[1]})
	p.LineTo(lr)
	p.LineTo(P{lr[0], ul[1]})
	p.Close()
}

func (p *Path) Polygon(path ...P) {
	if len(path) < 2 {
		return
	}
	p.MoveTo(path[len(path)-1])
	for _, pt := range path {
		p.LineTo(pt)
	}
	p.Close()
}

func (p *Path) Cylinder(pt1, pt2 P, w1, w2 float64) {
	if pt1 == pt2 {
		return
	}
	_, theta := pt1.PolarTo(pt2)
	phi := theta - .5
	pt1B := pt1.GoPolar(w1, phi+1)
	pt2A := pt2.GoPolar(w2, phi)
	p.MoveTo(pt1B)
	p.ArcTo(pt1, w1, w1, 1+phi, 1)
	p.LineTo(pt2A)
	p.ArcTo(pt2, w2, w2, phi, 1)
	p.LineTo(pt1B)
}

func (f *Filler) Blip(pt P) {
	f.ArcTo(pt, 4, 4, 0, 2)
	f.FillStroke()
}
