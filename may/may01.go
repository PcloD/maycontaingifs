package may

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/geom"
	"github.com/bit101/blgo/random"
)

// May01 generates a gif
func May01() {
	width := 400.0
	height := 400.0
	var points []*geom.Point
	for i := 0; i < 100; i++ {
		points = append(points, geom.NewPoint(random.FloatRange(-width/2, width/2), random.FloatRange(-height/2, height/2)))
	}

	animation := anim.NewAnimation(width, height, 120)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		surface.ClearRGB(1, 1, 1)
		surface.Save()
		surface.Translate(width/2, height/2)
		surface.Rotate(math.Pi * 2.0 * percent)

		surface.SetSourceRGB(0, 0, 0)
		for _, p := range points {
			surface.FillCircle(p.X, p.Y, percent*100)
		}

		surface.SetSourceRGB(1, 1, 1)
		for _, p := range points {
			surface.FillCircle(p.X, p.Y, percent*70)
		}
		surface.Restore()
	})
}
