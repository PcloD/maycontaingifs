package may

import (
	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
	"github.com/bit101/blg/geom"
	"github.com/bit101/blg/noise"
)

// May27 generates a gif
func May27() {
	width := 400.0
	height := 400.0

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		surface.ClearRGB(1, 1, 1)
		surface.SetLineWidth(0.25)
		res := 2.0
		scale := 0.012
		z := blmath.LerpSin(percent, -1, 1)

		for y := 0.0; y < height+res; y += res {
			var points []*geom.Point
			for x := 0.0; x < width+res; x += res {
				xx := x + noise.Perlin(x*scale, y*scale, z)*50.0
				yy := y + noise.Perlin(x*scale, y*scale, z+10.0)*50.0
				// fmt.Println(noise.Perlin(x*scale, y*scale, z))
				points = append(points, geom.NewPoint(xx, yy))
			}
			surface.StrokeMultiCurve(points)
		}

	})
}
