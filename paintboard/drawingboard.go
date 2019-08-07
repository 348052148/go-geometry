package paintboard

import (
	"image"
	"image/draw"
)

type DrawBoard struct {
	//矩形
	Rectangle image.Rectangle
	//图形载体
	Palette draw.Image
}

//画板
func NewDrawBoard(width, height int) DrawBoard {
	return DrawBoard{
		image.Rect(0,0, width,height),
	image.NewRGBA(image.Rect(0,0, width,height)),
	}
}

//获取画笔
func (drawBoard DrawBoard)GetPaintBrush() PaintBrush  {
	return PaintBrush{drawBoard.Palette}
}
