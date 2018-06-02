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

// May07 generates a gif
func May07() {
	width := 400.0
	height := 400.0
	numPoints := 800
	maxDist := 43.0
	var pointsA []*geom.Point
	var pointsB []*geom.Point
	for i := 0; i < numPoints; i++ {
		angle := random.FloatRange(0, math.Pi*2)
		radius := random.FloatRange(0, width)
		y := height/2 + math.Sin(angle)*radius
		pointsA = append(pointsA, geom.NewPoint(width/2+math.Cos(angle)*radius, y))
		angle = random.FloatRange(0, math.Pi*2)
		radius = random.FloatRange(0, width)
		pointsB = append(pointsB, geom.NewPoint(width/2+math.Cos(angle)*radius, y))
	}

	animation := anim.NewAnimation(width, height, 180)
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		var lpoints []*geom.Point
		for i := 0; i < len(pointsA); i++ {
			pa := pointsA[i]
			pb := pointsB[i]
			lp := geom.LerpPoint(blmath.CosRange(percent*math.Pi*2, 0, 1), pa, pb)
			lpoints = append(lpoints, &lp)
		}
		surface.ClearRGB(1, 1, 1)
		for i, p0 := range lpoints {
			for _, p1 := range lpoints[i:] {
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
