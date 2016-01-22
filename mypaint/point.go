package mypaint

import (
	"math"
)

type P [2]float64

func (pt P) Add(dx, dy float64) P {
	return P{pt[0] + dx, pt[1] + dy}
}
func (pt P) AddP(pt2 P) P {
	return P{pt[0] + pt2[0], pt[1] + pt2[1]}
}
func (pt P) Mul(x float64) P {
	return P{pt[0] * x, pt[1] * x}
}
func (pt P) Sub(dx, dy float64) P {
	return pt.Add(-dx, -dy)
}
func (pt P) SubP(pt2 P) P {
	return pt.AddP(pt2.Mul(-1))
}
func (pt P) Mid(pt2 P) P {
	return P{.5 * (pt[0] + pt2[0]), .5 * (pt[1] + pt2[1])}
}
func (pt P) Partway(pt2 P, percent float64) P {
	r, theta := pt.PolarTo(pt2)
	return pt.GoPolar(r*percent, theta)
}

func (pt P) Floor() [2]int {
	return [2]int{int(math.Floor(pt[0])), int(math.Floor(pt[1]))}
}

func (pt P) Dist(p2 P) float64 {
	if pt == p2 {
		return 0
	}
	dx := pt[0] - p2[0]
	dy := pt[1] - p2[1]
	return math.Sqrt(dx*dx + dy*dy)
}

func (pt P) GoPolar(r, theta float64) P {
	if r <= 0 {
		return pt
	}
	theta = RadFix(theta)
	rads := theta * math.Pi
	dy, dx := math.Sincos(rads)
	return P{pt[0] + (r * dx), pt[1] + (r * dy)}
}

func (pt P) PolarTo(p2 P) (r float64, theta float64) {
	if pt == p2 {
		return 0, 0
	}
	r = pt.Dist(p2)
	cosT := (p2[0] - pt[0]) / r
	theta = math.Acos(cosT) / math.Pi
	if p2[1] < pt[1] {
		theta = 2 - theta
	}
	return
}

func (pt P) RotateAround(pt2 P, theta float64) P {
	r, theta2 := pt2.PolarTo(pt)
	return pt2.GoPolar(r, theta+theta2)
}
func (pt P) Rotate(theta float64) P {
	return pt.RotateAround(P{}, theta)
}

func RadFix(theta float64) float64 {
	for theta < 0 {
		theta += 2
	}
	for theta >= 2 {
		theta -= 2
	}
	return theta
}

func (p P) Bumpers(r, up float64) (P, P) {
	return p.GoPolar(r, up-.5), p.GoPolar(r, up+.5)
}
