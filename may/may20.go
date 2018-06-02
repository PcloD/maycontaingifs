package may

import (
	"fmt"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/geom"
	"github.com/bit101/blgo/random"
)

// May20 generates a gif
func May20() {
	width := 400.0
	height := 400.0
	var path []*geom.Point
	random.Seed(3)
	for i := 0.0; i < 60.0; i += 1.0 {
		angle := math.Pi * 2.0 / 60.0 * i
		radius := random.FloatRange(50, 80)
		path = append(path, geom.NewPoint(math.Cos(angle)*radius, math.Sin(angle)*radius))
	}

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		surface.SetLineWidth(0.5)
		for x := -100.0; x < width+100.0; x += 5.0 {
			surface.Save()
			surface.Translate(x, height/2)
			surface.Rotate(x * 0.003)
			s := blmath.LerpSin(x/(width+200.0*2.0), 0.5, 2.0)
			surface.Scale(s, s)
			surface.Rotate(blmath.LerpSin(percent, -1, 1))
			surface.MultiLoop(path)
			surface.SetSourceRGB(1, 1, 1)
			surface.FillPreserve()
			surface.SetSourceRGB(0, 0, 0)
			surface.Stroke()
			surface.Restore()
		}

	})
}
