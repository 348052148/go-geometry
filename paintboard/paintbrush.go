package paintboard

import (
	"image/draw"
	"test/image/geometry"
)

//画笔
type PaintBrush struct {
	Palette draw.Image
}
//绘点
func (p PaintBrush)DrawPoint(x, y int, pigment Pigment)  {
	p.Palette.Set(x, y, pigment.Color)
}
//绘制
func (p PaintBrush)Draw(geometry geometry.Geometry, pigment Pigment)  {
	geometry.Draw(func(x, y int) {
		p.DrawPoint(x, y, pigment)
	})
}