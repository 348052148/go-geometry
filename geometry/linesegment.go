package geometry

type LineSegment struct {
	//起始点
	StartPoint Point
	//结束点
	EndPoint Point
}

func NewLineSegment(spoint Point, epoint Point) LineSegment {
	return LineSegment{
		StartPoint: spoint,
		EndPoint: epoint,
	}
}

//获取线段与线段之间的交点
func (lineSegment LineSegment)CrossPointAsLineSegment(geometry LineSegment) []Point  {
	k1, b1 := PrimaryFunc(lineSegment.StartPoint, lineSegment.EndPoint)
	k2, b2 := PrimaryFunc(geometry.StartPoint, geometry.EndPoint)
	//k1 * x + b1 = k2 * x + b2
	//x  = (b2 - b1) / (k1 - k2)
	x :=  int((b2 - b1) / (k1 - k2))
	fun := LinearFunc(lineSegment.StartPoint, lineSegment.EndPoint)
	return []Point{
		Pt(fun(x, 0)),
	}
}

//判断点是否在直线上
func (lineSegment LineSegment)IsCrossPointAsLineSegment(point Point) bool  {
	k1,_ := PrimaryFunc(point, lineSegment.StartPoint)
	k2,_ := PrimaryFunc(lineSegment.StartPoint, lineSegment.EndPoint)
	//如果斜率相同
	if k1 == k2 {
		//判断是否在线段范围
		if lineSegment.StartPoint.Y > lineSegment.EndPoint.Y {
			if lineSegment.EndPoint.Y <= point.Y && point.Y <= lineSegment.StartPoint.Y {
				return true
			}
		}else {
			if lineSegment.StartPoint.Y <= point.Y && point.Y <= lineSegment.EndPoint.Y {
				return true
			}
		}
	}
	return false
}

//平移
func (lineSegment LineSegment)Translation(distance float64, angle float64)  Geometry{
	lineSegment.StartPoint = lineSegment.StartPoint.Translation(distance, angle).(Point)
	lineSegment.EndPoint = lineSegment.EndPoint.Translation(distance, angle).(Point)
	return lineSegment
}

func (lineSegment LineSegment)Flip()  {
	
}
//旋转
func (lineSegment LineSegment)Rotate(rpos Point,angle float64) Geometry {
	return LineSegment{
		lineSegment.StartPoint.Rotate(rpos, angle).(Point),
		lineSegment.EndPoint.Rotate(rpos, angle).(Point),
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

func PrimaryFunc(pos1, pos2 Point) (float64,float64) {
	var k, b float64
	if (pos2.X - pos1.X) == 0 {
		k = 0
	}else {
		k = float64(pos2.Y - pos1.Y) / float64(pos2.X - pos1.X)
	}
	b = float64(pos1.Y) - k * float64( pos1.X)
	return k, b
}

func LinearFunc(pos1, pos2 Point) func(x, y int)(int, int) {
	// y = kx + b
	k, b := PrimaryFunc(pos1, pos2)
	return func(x, y int) (int, int) {
		return x, int(k * float64(x) + b)
	}
}

func InverseLinearFunc(pos1, pos2 Point) func(x, y int)(int, int) {
	// x = (y -b ) / k
	k, b := PrimaryFunc(pos1, pos2)
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

