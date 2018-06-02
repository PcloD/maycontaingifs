package may

import (
	"fmt"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
)

// May18 generates a gif
func May18() {
	width := 400.0
	height := 400.0

	var circ func(*blgo.Surface, float64, float64, float64, float64, int, float64)
	circ = func(surface *blgo.Surface, rotation, x, y, r float64, d int, w float64) {
		surface.Save()
		surface.Translate(x, y)
		surface.Rotate(rotation)
		if d > 0 {
			radius := r * 1.6
			surface.SetLineWidth(w)
			for i := 0.0; i < 3.0; i += 1.0 {
				angle := math.Pi * 2.0 / 3.0 * i
				surface.Line(0, 0, math.Cos(angle)*radius, math.Sin(angle)*radius)
				circ(surface, rotation, math.Cos(angle)*radius, math.Sin(angle)*radius, r*0.6, d-1, w*0.6)
			}
		}
		surface.Restore()
	}

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(0, 0, 0)
		surface.Save()
		surface.Translate(width/2, height/2)

		surface.SetSourceRGB(1, 1, 1)
		circ(surface, percent*math.Pi*2.0/3.0, 0, 0, 60, 6, 14)
		surface.SetSourceRGB(0, 0, 0)
		circ(surface, percent*math.Pi*2.0/3.0, 0, 0, 60, 6, 10)

		surface.Restore()
	})

}
