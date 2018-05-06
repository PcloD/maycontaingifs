package may

import (
	"fmt"
	"math"

	"github.com/bit101/bitlibgo"
	"github.com/bit101/bitlibgo/anim"
)

// May15 generates a gif
func May15() {
	width := 400.0
	height := 400.0
	filename := "out/may15.gif"

	var circ func(*bitlibgo.BitSurface, float64, float64, float64, float64, int)
	circ = func(surface *bitlibgo.BitSurface, rotation, x, y, r float64, d int) {
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

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 180
	animation.Render(func(surface *bitlibgo.BitSurface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		surface.Save()
		surface.Translate(width/2, height/2)

		circ(surface, percent*math.Pi*2.0/3.0, 0, 0, 60, 7)

		surface.Restore()
	})

}
