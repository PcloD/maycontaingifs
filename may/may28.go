package may

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
)

// May28 generates a gif
func May28() {
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
		0.0, 1.2, 3.0, 0.7, -0.6, 0.99,
	}

	plot := func(surface *blgo.Surface, x, y, z float64) {
		scale := fl / (fl + z)
		surface.FillRectangle(x*scale*size, y*scale*size, scale*1, scale*1)
	}

	f := func(x, y, z float64) (float64, float64, float64) {
		x1 := z*math.Sin(params.a*x) + math.Cos(params.b*y)
		y1 := x*math.Sin(params.c*y) + math.Cos(params.d*z)
		z1 := y*math.Sin(params.e*z) + math.Cos(params.f*x)
		return x1, y1, z1
	}

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		surface.Save()
		surface.Translate(width*0.27, height*0.2)
		surface.ClearRGB(0, 0, 0)
		surface.SetSourceRGB(1, 1, 1)
		x := 0.1
		y := 0.1
		z := 0.1

		params.a = blmath.LerpSin(percent, 0.5, 1.5)
		for i := 0; i < count; i++ {
			plot(surface, x, y, z)
			x, y, z = f(x, y, z)
		}
		surface.Restore()

	})
}
