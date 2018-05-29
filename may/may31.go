package may

import (
	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
	"github.com/bit101/gometro/metro"
)

// May31 generates a gif
func May31() {
	width := 400.0
	height := 400.0
	box := metro.NewBox(100, 100, 100)
	box.Position(200, 340, 0)

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		surface.ClearRGB(0, 0, 0)
		w := blmath.LerpSin(percent, 20, 200)
		d := 220 - w
		h := blmath.LerpSin(percent*2.0-0.25, 20, 200)
		box.Size(w, d, h)
		box.Render(surface)

	})
}
