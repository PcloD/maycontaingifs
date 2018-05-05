package may

import (
	"fmt"
	"math"

	"github.com/bit101/bitlibgo"
	"github.com/bit101/bitlibgo/anim"
	"github.com/bit101/bitlibgo/bitmath"
)

// May14 generates a gif
func May14() {
	width := 400.0
	height := 400.0
	filename := "out/may14.gif"
	numRows := 10.0
	numCols := 10.0
	rowHeight := height / numRows
	colWidth := width / numCols

	animation := anim.NewAnimation(filename)
	animation.SetSize(width, height)
	animation.Frames = 180
	animation.Render(func(surface *bitlibgo.BitSurface, percent float64) {
		fmt.Printf("\r%f", percent)
		// rowEven := true
		for r := 0.0; r < numRows; r++ {
			even := true
			surface.SetSourceRGB(0, 0, 0)
			surface.MoveTo(0, rowHeight*r)
			surface.LineTo(width, rowHeight*r)
			surface.Stroke()
			for c := -1.0; c < numCols+1.0; c++ {
				surface.Save()
				surface.Translate(math.Sin(r/numRows*math.Pi*4)*30*bitmath.SinRange(percent*math.Pi*2, -1, 1), rowHeight*r)
				if even {
					surface.SetSourceRGB(1, 1, 1)
				} else {
					surface.SetSourceRGB(0, 0, 0)
				}
				surface.FillRectangle(c*colWidth, 0, colWidth, rowHeight)
				even = !even
				surface.Restore()
			}
			// rowEven = !rowEven
		}

	})
}
