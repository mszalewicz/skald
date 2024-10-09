package gui

import "math"

func CalculateResolution(xIn, yIn int, scale float64) (x int, y int) {
	if scale < 0 || scale > 1 {
		x, y = xIn, yIn
		return
	}

	x, y = int(math.Ceil((float64(xIn) * scale))), int(math.Ceil((float64(yIn) * scale)))
	return
}
