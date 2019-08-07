package geometry


type Rectangle struct {
	//宽
	Width int
	//高
	Height int
	//笛卡尔坐标位置
	X, Y int

	angle float64
	rpos Point
}

func NewRectangle(x, y int, width, height int) Rectangle  {
	return Rectangle{
		Width:width,
		Height:height,
		X:x,
		Y:y,
		angle:0,
	}
}

//平移
func (rectangle Rectangle)Translation(distance float64)  Geometry {
	rectangle.X , rectangle.Y = int(float64(rectangle.X) + distance) , int(float64(rectangle.Y) + distance)
	return rectangle
}

func (rectangle Rectangle)Flip()  {

}
//旋转
func (rectangle Rectangle)Rotate(rpos Point,angle float64) Geometry {
	return Rectangle{
		X:rectangle.X,
		Y:rectangle.Y,
		Width:rectangle.Width,
		Height:rectangle.Height,
		rpos:rpos,
		angle:angle,
	}
}

func (rectangle Rectangle)SeekSides() []LineSegment {
	if rectangle.angle == 0 {
		return []LineSegment{
			LineSegment{Pt(rectangle.X, rectangle.Y), Pt(rectangle.X+rectangle.Width, rectangle.Y)},
			LineSegment{Pt(rectangle.X, rectangle.Y+rectangle.Height), Pt(rectangle.X+rectangle.Width, rectangle.Y+rectangle.Height)},
			LineSegment{Pt(rectangle.X, rectangle.Y), Pt(rectangle.X, rectangle.Y+rectangle.Height)},
			LineSegment{Pt(rectangle.X+rectangle.Width, rectangle.Y+rectangle.Height), Pt(rectangle.X+rectangle.Width, rectangle.Y)},
		}
	}else {
		//旋转各个线段
		return []LineSegment{
			LineSegment{Pt(rectangle.X, rectangle.Y), Pt(rectangle.X+rectangle.Width, rectangle.Y)}.Rotate(rectangle.rpos,rectangle.angle).(LineSegment),
			LineSegment{Pt(rectangle.X, rectangle.Y+rectangle.Height), Pt(rectangle.X+rectangle.Width, rectangle.Y+rectangle.Height)}.Rotate(rectangle.rpos,rectangle.angle).(LineSegment),
			LineSegment{Pt(rectangle.X, rectangle.Y), Pt(rectangle.X, rectangle.Y+rectangle.Height)}.Rotate(rectangle.rpos,rectangle.angle).(LineSegment),
			LineSegment{Pt(rectangle.X+rectangle.Width, rectangle.Y+rectangle.Height), Pt(rectangle.X+rectangle.Width, rectangle.Y)}.Rotate(rectangle.rpos,rectangle.angle).(LineSegment),
		}
	}
}
//绘制矩形
func(rectangle Rectangle)Draw(drawFunc func(x, y int))   {
	for _, lineSegment := range rectangle.SeekSides() {
		lineSegment.Draw(drawFunc)
	}
}

func (rectangle Rectangle)CenterPoint() Point  {
	return Pt(rectangle.X+ rectangle.Width/2, rectangle.Y+rectangle.Height/2)
}