package geometry

import (
	"math"
)

type Circular struct {
	//圆心
	CenterOfCircular Point
	//半径
	Radius float64
	//Arc
	Rad float64
}

func NewCircular(centerOfCircular Point, radius float64, rad float64) Circular  {
	return Circular{
		CenterOfCircular:centerOfCircular,
		Radius:radius,
		Rad:rad,
	}
}


//旋转
func (circular Circular)Rotate(rpos Point,angle float64) Geometry {

	return circular
}
//翻转
func (circular Circular)Flip() {

}
//平移
func (circular Circular)Translation(distance float64,  angle float64)  Geometry {
	circular.CenterOfCircular = circular.CenterOfCircular.Translation(distance, angle).(Point)
	return circular
}

func (circular Circular)Draw(drawFunc func(x, y int))  {
	//x=x1+rcosθ
	//y=y1+rsinθ
	const SPEED  = 720 * 2
	var x, y float64
	for i:= 1.0 ; i <= SPEED * circular.Rad; i++ {
		y = float64(circular.CenterOfCircular.Y) + circular.Radius * math.Sin(2 * math.Pi / SPEED * i)
		x = float64(circular.CenterOfCircular.X) + circular.Radius * math.Cos(2 * math.Pi / SPEED * i)
		drawFunc(int(x), int(y))
	}
}

func (circular Circular)CenterPoint() Point  {
	return circular.CenterOfCircular
}
