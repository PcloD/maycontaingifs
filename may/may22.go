package may

import (
	"fmt"
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
	"github.com/bit101/blg/geom"
	"github.com/bit101/blg/noise"
	"github.com/bit101/blg/random"
)

// May22 generates a gif
func May22() {
	width := 400.0
	height := 400.0

	type Particle struct {
		x float64
		y float64
	}

	scale := 0.01

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		var points []*geom.Point

		random.Seed(0)
		for i := 0; i < 500; i++ {
			angle := random.FloatRange(0, math.Pi*2.0)
			// r := random.FloatRange(0, 100)
			r := 100.0
			x := width/2.0 + math.Cos(angle)*r
			y := height*0.7 + math.Sin(angle)*r
			p := geom.NewPoint(x, y)
			points = append(points, p)
		}
		for i := 0.0; i < blmath.LerpSin(percent, 1, 200); i++ {
			for _, p := range points {
				surface.FillRectangle(p.X, p.Y, 0.5, 0.5)
				a := blmath.Map(noise.Perlin(p.X*scale, p.Y*scale, 0), -0.25, 0.75, -math.Pi, math.Pi)
				vx := math.Cos(a) * 1
				vy := math.Sin(a) * 1
				p.X += vx
				p.Y += vy
			}
		}

	})
}
