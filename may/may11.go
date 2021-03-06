package may

import (
	"fmt"
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
)

// May11 generates a gif
func May11() {
	width := 400.0
	height := 400.0

	type Point3D struct {
		x float64
		y float64
		z float64
	}

	radius := 100.0
	var points []*Point3D
	for y := -height * 1.5; y < height*1.5; y += 40.0 {
		angle := y / height * 3.0

		p0 := &Point3D{
			math.Cos(angle) * radius,
			y,
			math.Sin(angle) * radius,
		}
		points = append(points, p0)
		p1 := &Point3D{
			math.Cos(angle) * -radius,
			y,
			math.Sin(angle) * -radius,
		}
		points = append(points, p1)
	}

	fl := 500.0

	renderPoints := func(surface *blgo.Surface, p0 *Point3D, p1 *Point3D, angle float64) {
		xx := math.Cos(angle)*p0.x - math.Sin(angle)*p0.z
		zz := math.Cos(angle)*p0.z + math.Sin(angle)*p0.x
		scale := fl / (fl + zz + 200)
		x0 := scale * xx
		y0 := scale * p0.y
		r0 := scale * 10
		surface.FillCircle(x0, y0, r0)

		xx = math.Cos(angle)*p1.x - math.Sin(angle)*p1.z
		zz = math.Cos(angle)*p1.z + math.Sin(angle)*p1.x
		scale = fl / (fl + zz + 200)
		x1 := scale * xx
		y1 := scale * p0.y
		r1 := scale * 10
		surface.FillCircle(x1, y1, r1)
		surface.Line(x0, y0, x1, y1)
	}
	animation := anim.NewAnimation(width, height, 180)
	firstFrame := true
	animation.Render("frames", "frame", func(surface *blgo.Surface, percent float64) {
		fmt.Printf("\r%f", percent)
		n := blmath.SinRange(percent*math.Pi*2.0, 0.03, 1.0)
		if firstFrame {
			surface.ClearRGB(0, 0, 0)
			firstFrame = false
		} else {
			surface.SetSourceRGBA(0, 0, 0, n)
		}
		surface.Paint()
		surface.SetSourceRGB(n, n, 1)
		surface.Save()
		surface.Translate(width/2, height/2)
		angle := percent * math.Pi * 2

		for i := 0; i < len(points)-1; i += 2 {
			p0 := points[i]
			p1 := points[i+1]
			renderPoints(surface, p0, p1, angle)
		}
		surface.Restore()
	})

}
