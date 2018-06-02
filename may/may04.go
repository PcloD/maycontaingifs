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

// May04 generates a gif
func May04() {
	width := 380.0
	height := 380.0
	center := geom.NewPoint(width/2, height/2)
	numPoints := 400
	var points []*geom.Point
	for i := 0; i < numPoints; i++ {
		points = append(points, geom.NewPoint(random.FloatRange(0, width), random.FloatRange(0, height)))
	}

	animation := anim.NewAnimation(width, height, 120)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		surface.ClearRGB(1, 1, 1)
		for i, p0 := range points {
			dActual := p0.Distance(center)
			dIdeal := blmath.SinRange(percent*math.Pi*2, 0, width)
			dx := math.Abs(dActual - dIdeal)
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
