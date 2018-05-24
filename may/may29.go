package may

import (
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
)

// May29 generates a gif
func May29() {
	width := 400.0
	height := 400.0
	const count = 100000
	const size = 200.0
	const fl = 300.0

	type Params struct {
		a float64
		b float64
		c float64
		d float64
		e float64
		f float64
	}

	params := Params{
		2.3, 1.2, 3.0, 0.7, -0.6, 0.99,
	}

	type P3D struct {
		x float64
		y float64
		z float64
	}

	plot := func(surface *blg.Surface, p P3D) {
		scale := fl / (fl + p.z*size + size)
		surface.FillRectangle(p.x*scale*size, p.y*scale*size, scale*1, scale*1)
	}

	f := func(p P3D) P3D {
		x1 := p.z*math.Sin(params.a*p.x) + math.Cos(params.b*p.y)
		y1 := p.x*math.Sin(params.c*p.y) + math.Cos(params.d*p.z)
		z1 := p.y*math.Sin(params.e*p.z) + math.Cos(params.f*p.x)
		return P3D{x1, y1, z1}
	}

	rotate := func(points []P3D, angle float64) []P3D {
		var ps []P3D
		for _, p := range points {
			x := p.x*math.Cos(angle) - p.z*math.Sin(angle)
			z := p.z*math.Cos(angle) + p.x*math.Sin(angle)
			ps = append(ps, P3D{x, p.y, z})
		}
		return ps
	}

	p := P3D{0.1, 0.1, 0.1}
	var points []P3D
	for i := 0; i < count; i++ {
		points = append(points, p)
		p = f(p)
	}

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		surface.Save()
		surface.Translate(width*0.5, height*0.1)
		surface.ClearRGB(0, 0, 0)
		surface.SetSourceRGB(1, 1, 1)

		ps := rotate(points, percent*math.Pi*2)
		for _, p := range ps {
			plot(surface, p)
		}

		surface.Restore()

	})
}
