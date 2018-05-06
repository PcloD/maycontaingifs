package may

import (
	"fmt"
	"math"

	"github.com/bit101/blg"
	"github.com/bit101/blg/anim"
	"github.com/bit101/blg/blmath"
	"github.com/bit101/blg/geom"
	"github.com/bit101/blg/random"
)

// May03 generates a gif
func May03() {
	width := 400.0
	height := 400.0
	numPoints := 500
	var points []*geom.Point
	for i := 0; i < numPoints; i++ {
		points = append(points, geom.NewPoint(random.FloatRange(0, width), random.FloatRange(0, height)))
	}

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blg.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		x := blmath.Lerp(percent, -400, width+400)
		for i, p0 := range points {
			dx := math.Abs(x - p0.X)
			maxDist := (1.0 - dx/width) * 100.0
			for _, p1 := range points[i:] {
				dist := p0.Distance(p1)
				if dist < maxDist {
					surface.SetLineWidth((1 - dist/maxDist) * 0.5)
					surface.MoveTo(p0.X, p0.Y)
					surface.LineTo(p1.X, p1.Y)
					surface.Stroke()
				}
			}
		}
	})
}
