package geometry

import (
	"math"
)


type LineSegment struct {
	//起始点
	StartPoint Point
	//结束点
	EndPoint Point
}

func NewLineSegment(spoint [2]int, epoint [2]int) LineSegment {
	return LineSegment{
		StartPoint:Pt(spoint[0], spoint[1]),
		EndPoint:Pt(epoint[0], epoint[1]),
	}
}

//平移
func (lineSegment LineSegment)Translation(distance float64)  Geometry{
	lineSegment.StartPoint.X, lineSegment.StartPoint.Y = int(float64(lineSegment.StartPoint.X) + distance), int(float64(lineSegment.StartPoint.Y) + distance)
	lineSegment.EndPoint.X, lineSegment.EndPoint.Y = int(float64(lineSegment.EndPoint.X) + distance), int(float64(lineSegment.EndPoint.Y) + distance)
	return lineSegment
}

func (lineSegment LineSegment)Flip()  {
	
}
//旋转
func (lineSegment LineSegment)Rotate(rpos Point,angle float64) Geometry {
	return LineSegment{
		Pt(
			int(float64(lineSegment.StartPoint.X - rpos.X) * math.Cos(math.Pi/ 180 * angle) - float64(lineSegment.StartPoint.Y - rpos.Y) * math.Sin(math.Pi/ 180 * angle)) + rpos.X,
			int(float64(lineSegment.StartPoint.X - rpos.X) * math.Sin(math.Pi/ 180 * angle) + float64(lineSegment.StartPoint.Y - rpos.Y) * math.Cos(math.Pi/ 180 * angle)) + rpos.Y,
		),
		Pt(
			int(float64(lineSegment.EndPoint.X - rpos.X) * math.Cos(math.Pi/ 180 * angle) - float64(lineSegment.EndPoint.Y - rpos.Y) * math.Sin(math.Pi/ 180 * angle)) + rpos.X,
			int(float64(lineSegment.EndPoint.X - rpos.X) * math.Sin(math.Pi/ 180 * angle) + float64(lineSegment.EndPoint.Y - rpos.Y) * math.Cos(math.Pi/ 180 * angle)) + rpos.Y,
		),
	}
}

func (lineSegment LineSegment)CenterPoint() Point  {
	return Pt((lineSegment.StartPoint.X+lineSegment.EndPoint.X) / 2 , (lineSegment.StartPoint.Y + lineSegment.EndPoint.Y) / 2)
}

type InearFormula func(x, y int) (int, int)
func (lineSegment LineSegment)Draw(drawFunc func(x, y int))  {
	p1, p2 := CombioPoint(lineSegment.StartPoint, lineSegment.EndPoint)
	var fn InearFormula
	if (p2.X - p1.X < p2.Y - p1.Y) {
		fn = InverseLinearFunc(lineSegment.StartPoint, lineSegment.EndPoint)
	}else {
		//根据2点算出 公式 y = kx + b 中 未知数 k 和 b
		fn = LinearFunc(lineSegment.StartPoint, lineSegment.EndPoint)
	}
	//描点
	for y := p1.Y; y <= p2.Y; y++ {
		for x := p1.X; x <= p2.X ; x++  {
			drawFunc(fn(x,y))
		}
	}
}

func LinearFunc(pos1, pos2 Point) func(x, y int)(int, int) {
	// y = kx + b
	var k, b float64
	if (pos2.X - pos1.X) == 0 {
		k = 0
	}else {
		k = float64(pos2.Y - pos1.Y) / float64(pos2.X - pos1.X)
	}
	b = float64(pos1.Y) - k * float64( pos1.X)
	return func(x, y int) (int, int) {
		return x, int(k * float64(x) + b)
	}
}

func InverseLinearFunc(pos1, pos2 Point) func(x, y int)(int, int) {
	// x = (y -b ) / k
	var k, b float64
	if (pos2.X - pos1.X) == 0 {
		k = 0
	}else {
		k = float64(pos2.Y - pos1.Y) / float64(pos2.X - pos1.X)
	}
	b = float64(pos1.Y) - k * float64( pos1.X)
	return func(x, y int) (int,int) {
		if k == 0 {
			return x, y
		}
		return int((float64(y) - b) / k), y
	}
}

//重组点
func CombioPoint( pos1, pos2 Point) (Point, Point) {
	cpos1, cpos2 :=  Point{},Point{}
	if pos1.X > pos2.X {
		cpos2.X, cpos1.X = pos1.X, pos2.X
	}else {
		cpos1.X, cpos2.X = pos1.X, pos2.X
	}
	if pos1.Y > pos2.Y {
		cpos2.Y, cpos1.Y = pos1.Y, pos2.Y
	}else {
		cpos2.Y, cpos1.Y = pos2.Y, pos1.Y
	}
	return cpos1,cpos2
}

