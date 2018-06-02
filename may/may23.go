package may

import (
	"fmt"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/geom"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/random"
)

// May23 generates a gif
func May23() {
	width := 400.0
	height := 400.0

	var points []*geom.Point

	random.Seed(0)
	count := 1000.0
	for i := 0.0; i < count; i++ {
		p := geom.NewPoint(
			random.FloatRange(0, width),
			random.FloatRange(0, height),
		)

		points = append(points, p)
	}
	scale := 0.007

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		r := 10.0
		for _, p := range points {
			z := blmath.LerpSin(percent, 0, 500)
			rr := noise.Perlin(p.X*scale, p.Y*scale, z*scale)*r + r
			surface.Save()
			surface.Translate(p.X, p.Y)
			surface.SetSourceRGB(0, 0, 0)
			surface.FillCircle(0, 0, rr)

			surface.SetSourceRGB(1, 1, 1)
			surface.FillCircle(0, 0, rr*0.95)

			surface.SetSourceRGB(0, 0, 0)
			surface.FillCircle(rr*0.1, rr*0.1, rr*0.85)

			surface.Restore()
		}

	})
}
