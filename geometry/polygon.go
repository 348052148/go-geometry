package geometry

//平面图形
type Plane interface {
	GetVertexs()[]Point
	GetEdges() []Edge
	ContainerPoint(point Point) bool
}
type Edge LineSegment
//多边形
type Polygon struct {
	Vertexs []Point
}

//多边形
func NewPolygon(vertexs []Point) Polygon {
	return Polygon{
		Vertexs:vertexs,
	}
}

func (polygon Polygon)Draw(drawFunc func(x, y int)) {
	func(edges []Edge){
		for _, edge := range edges{
			LineSegment(edge).Draw(drawFunc)
		}
	}(polygon.GetEdges())
}
func (polygon Polygon)CenterPoint() Point {
	return Pt(0, 0)
}

//平移
func (polygon Polygon)Translation(distance float64, angle float64)  Geometry {
	polygon.Vertexs = func(vertexs []Point) []Point {
		var pvertexs []Point
		for _,v := range vertexs{
			pvertexs = append(pvertexs, v.Translation(distance, angle).(Point))
		}
		return pvertexs
	}(polygon.Vertexs)
	return polygon
}

func (polygon Polygon)Flip()  {

}
//旋转
func (polygon Polygon)Rotate(rpos Point,angle float64) Geometry {
	polygon.Vertexs = func(vertexs []Point) []Point {
		var pvertexs []Point
		for _,v := range vertexs{
			pvertexs = append(pvertexs, v.Rotate(rpos, angle).(Point))
		}
		return pvertexs
	}(polygon.Vertexs)
	return polygon
}

func solvePloygonEdge(vertexs []Point) []Edge {
	startVertex := vertexs[0]
	var edges []Edge
	for i := 1; i < len(vertexs); i++ {
		edges = append(edges, Edge(NewLineSegment(startVertex, vertexs[i])))
		startVertex = vertexs[i]
	}
	edges = append(edges, Edge(NewLineSegment(startVertex, vertexs[0])))
	return edges
}
//获取顶点
func (polygon Polygon)GetVertexs() []Point {
	return polygon.Vertexs
}
//获取边
func (polygon Polygon)GetEdges()[]Edge  {
	return solvePloygonEdge(polygon.Vertexs)
}
//点是否在多边形内
//存在照射点 垂直时判断失效
func (polygon Polygon)IsPolygobScope(point Point) bool {
	counter := 0
	for _, edge := range polygon.GetEdges() {
		//如果在顶点上
		if edge.EndPoint.Eq(point) {
			return true
		}
		//判断点是否在阴影范围
		if point.Y > Min(edge.StartPoint.Y, edge.EndPoint.Y) && point.Y < Max(edge.StartPoint.Y, edge.EndPoint.Y) {
			if edge.StartPoint.Y != edge.EndPoint.Y {
				//xinters = (point.Y - p1.Y) * (p2.X - p1.X) / (p2.Y - p1.Y) + p1.X;
				xinters := (point.Y - edge.StartPoint.Y) * (edge.EndPoint.X - edge.StartPoint.X) /
					(edge.EndPoint.Y - edge.StartPoint.Y) + edge.StartPoint.X
				//
				if edge.StartPoint.X == edge.EndPoint.X || point.X <= xinters {
					counter++
				}
			}
		}
	}
	return counter % 2 == 1
}

func Max(v1, v2 int) int  {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Min(v1, v2 int) int  {
	if v2 > v1 {
		return v1
	}
	return v2
}