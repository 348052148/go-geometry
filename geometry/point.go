package geometry

import (
	"math"
)

type Point struct {
	X, Y int
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

//
func (point Point)Add(distance int) Point  {
	point.X, point.Y = point.X + distance, point.Y +distance
	return point
}

func (point Point)Translation(distance float64, angle float64)  Geometry{
	point.X, point.Y = int(float64(point.X) + distance * math.Cos(2 * math.Pi / 360 * angle)), int(float64(point.Y) + distance * math.Sin(2 * math.Pi / 360 * angle))
	return point
}

func (point Point)Flip()  {
}

func (point Point)Rotate(rpos Point, angle float64) Geometry {
	return Pt(
		int(float64(point.X - rpos.X) * math.Cos(math.Pi/ 180 * angle) - float64(point.Y - rpos.Y) * math.Sin(math.Pi/ 180 * angle)) + rpos.X,
		int(float64(point.X - rpos.X) * math.Sin(math.Pi/ 180 * angle) + float64(point.Y - rpos.Y) * math.Cos(math.Pi/ 180 * angle)) + rpos.Y,
	)
}

func (point Point)Eq(point2 Point) bool {
	return point.X == point2.X && point.Y == point2.Y
}

func (point Point)Draw(drawFunc func(x, y int))  {
	drawFunc(point.X, point.Y)
}

func (point Point)CenterPoint() Point  {
	return point
}
