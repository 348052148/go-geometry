package geometry

//平行线
type ParallelLine struct {
	//线段
	Line LineSegment
	//距离
	Distance int
}

func NewParallelLine(line LineSegment, distance int) ParallelLine {
	return ParallelLine{
		Line:line,
		Distance:distance,
	}
}

//计算平行线中心点
func (parallelLine ParallelLine)CenterPoint() Point  {
	return Pt((parallelLine.Line.StartPoint.X+ parallelLine.Distance / 2) , (parallelLine.Line.StartPoint.Y + parallelLine.Line.EndPoint.Y) / 2)
}

func (parallelLine ParallelLine)Draw(drawFunc func(x, y int))  {
	//画Line01
	parallelLine.Line.Draw(drawFunc)
	//画Line02
	LineSegment{
		StartPoint:parallelLine.Line.StartPoint.Add(parallelLine.Distance),
		EndPoint:parallelLine.Line.EndPoint.Add(parallelLine.Distance),
	}.Draw(drawFunc)

}