package paintboard

import "image/color"

const RED = 1
const GREEN = 2
const YELLO = 3

type Pigment struct {
	Color color.Color
}
func NewPigment(r, g, b uint8) Pigment {
	return Pigment{color.RGBA{r,g,b,255}}
}
