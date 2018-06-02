package may

import (
	"fmt"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
)

// May15 generates a gif
func May15() {
	width := 400.0
	height := 400.0

	var circ func(*blgo.Surface, float64, float64, float64, float64, int)
	circ = func(surface *blgo.Surface, rotation, x, y, r float64, d int) {
		surface.Save()
		surface.Translate(x, y)
		surface.Rotate(rotation)
		surface.FillCircle(0, 0, r)
		if d > 0 {
			radius := r * 1.6
			for i := 0.0; i < 3.0; i += 1.0 {
				angle := math.Pi * 2.0 / 3.0 * i
				circ(surface, rotation, math.Cos(angle)*radius, math.Sin(angle)*radius, r*0.6, d-1)
			}
		}
		surface.Restore()
	}

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		surface.Save()
		surface.Translate(width/2, height/2)

		circ(surface, percent*math.Pi*2.0/3.0, 0, 0, 60, 7)

		surface.Restore()
	})

}
