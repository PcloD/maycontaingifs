package may

import (
	"math"
	"sort"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
)

// May30 generates a gif
func May30() {
	width := 400.0
	height := 400.0
	const count = 10000
	const size = 400.0
	const fl = 200.0

	type Params struct {
		a float64
		b float64
		c float64
		d float64
		e float64
		f float64
	}

	params := Params{
		2.3, -1.2, 2.0, 0.7, -0.6, 0.99,
	}

	plot := func(surface *blgo.Surface, p P3D) {
		scale := fl / (fl + p.z*size + size*2)
		// surface.FillRectangle(p.x*scale*size, p.y*scale*size, scale*1, scale*1)
		surface.SetLineWidth(0.2)
		surface.SetSourceRGB(1, 1, 1)
		surface.FillCircle(p.x*scale*size, p.y*scale*size, scale*20)
		surface.SetSourceRGB(0, 0, 0)
		surface.StrokeCircle(p.x*scale*size, p.y*scale*size, scale*20)
	}

	f := func(p *P3D) *P3D {
		x1 := p.z*math.Sin(params.a*p.x) + math.Cos(params.b*p.y)
		y1 := p.x*math.Sin(params.c*p.y) + math.Cos(params.d*p.z)
		z1 := p.y*math.Sin(params.e*p.z) + math.Cos(params.f*p.x)
		return &P3D{x1, y1, z1}
	}

	rotate := func(points P3DList, angle float64) P3DList {
		var ps P3DList
		for _, p := range points {
			x := p.x*math.Cos(angle) - p.z*math.Sin(angle)
			z := p.z*math.Cos(angle) + p.x*math.Sin(angle)
			ps = append(ps, &P3D{x, p.y, z})
		}
		return ps
	}

	p := &P3D{0.1, 0.1, 0.1}
	var points P3DList
	for i := 0; i < count; i++ {
		points = append(points, p)
		p = f(p)
	}

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		surface.Save()
		surface.Translate(width*0.5, height*0.2)
		surface.ClearRGB(0, 0, 0)
		surface.SetSourceRGB(1, 1, 1)

		ps := rotate(points, percent*math.Pi*2)
		sort.Sort(ps)
		for _, p := range ps {
			plot(surface, *p)
		}

		surface.Restore()

	})
}

// P3D is
type P3D struct {
	x float64
	y float64
	z float64
}

// P3DList is
type P3DList []*P3D

func (l P3DList) Len() int           { return len(l) }
func (l P3DList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l P3DList) Less(i, j int) bool { return l[i].z > l[i].z }
