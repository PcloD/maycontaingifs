package may

import (
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/geom"
	"github.com/bit101/blg/random"
)

// May01 generates a gif
func May01() {
	width := 400.0
	height := 400.0
	filename := "out/may01.gif"
	var points []*geom.Point
	for i := 0; i < 100; i++ {
		points = append(points, geom.NewPoint(random.FloatRange(-width/2, width/2), random.FloatRange(-height/2, height/2)))
	}

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 120
	animation.Render(func(surface *blg.Surface, percent float64) {
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
