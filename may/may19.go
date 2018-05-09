package may

import (
	"fmt"
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
	"github.com/bit101/blg/random"
)

// May19 generates a gif
func May19() {
	width := 400.0
	height := 400.0

	type Point3D struct {
		x float64
		y float64
		z float64
	}

	var points []*Point3D
	for i := 0; i < 100; i++ {
		angle := random.FloatRange(0, math.Pi)
		radius := random.FloatRange(-250, 250)
		p := &Point3D{
			math.Cos(angle) * radius,
			random.FloatRange(-200, 200),
			math.Sin(angle) * radius,
		}
		points = append(points, p)
	}

	fl := 300.0

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(0, 0, 0)
		surface.SetSourceRGB(1, 1, 1)
		surface.SetLineWidth(0.9)
		surface.Save()
		surface.Translate(width/2, height/2)
		angle := blmath.LerpSin(percent, -1.0, 1.0)

		for _, p := range points {
			xx := math.Cos(angle)*p.x - math.Sin(angle)*p.z
			zz := math.Cos(angle)*p.z + math.Sin(angle)*p.x
			scale := fl / (fl + zz + 200)
			x := scale * xx
			y := scale * p.y

			surface.LineTo(x, y)
		}
		surface.Stroke()
		surface.Restore()
	})
}
