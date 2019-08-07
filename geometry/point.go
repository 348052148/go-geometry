package geometry

import (
	"math"
)

type Point struct {
	X, Y int
	angle float64
	distance float64
}

func Pt(x, y int) Point {
	return Point{
		X:x,
		Y:y,
	}
}

func NewPoint(x, y int) Point  {
	return Point{
		X:x,
		Y:y,
	}
}

//旋转
func (point Point)Rotate(angle float64)  {
	point.angle = angle
}

func (point Point)Add(distance int) Point  {
	point.X, point.Y = point.X + distance, point.Y +distance
	return point
}

func (point Point)Translation(distance float64)  Geometry{
	point.X, point.Y = int(float64(point.X) + distance), int(float64(point.Y) + distance)
	return point
}

//前进
func (point Point)Forward( d float64) Point {
	point.distance = d
	return point
}

func (point Point)Draw(drawFunc func(x, y int))  {
	var x, y float64
	y = float64(point.Y) + point.distance * math.Sin(2 * math.Pi / 360 * point.angle)
	x = float64(point.X) + point.distance * math.Cos(2 * math.Pi / 360 * point.angle)
	LineSegment{
		StartPoint:Pt(point.X, point.Y),
		EndPoint:Pt(int(x), int(y)),
	}.Draw(drawFunc)
}

func (point Point)CenterPoint() Point  {
	return point
}
