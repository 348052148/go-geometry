package geometry

//几何
type Geometry interface {
	Draw(drawFunc func(x, y int))
	CenterPoint() Point
}

//变换
type Transformation interface {
	//平移
	Translation(distance float64)  Geometry
	//旋转
	Rotate(rpos Point,angle float64) Geometry
	//翻转
	Flip()
}
//复合图形
type ComplexGraph struct {
	geometryList []Geometry
}

func (complex ComplexGraph)AddGeometry(geometry Geometry) ComplexGraph {
	complex.geometryList = append(complex.geometryList, geometry)
	return complex
}

func (complex ComplexGraph)Translation(distance float64)  Geometry {
	complex.geometryList = func(geomes []Geometry) []Geometry {
		var combiGeomes []Geometry
		for _, geome := range geomes  {
			combiGeomes = append(combiGeomes, geome.(Transformation).Translation(distance))
		}
		return combiGeomes
	}(complex.geometryList)
	return complex
}

//旋转
func (complex ComplexGraph)Rotate(rpos Point,angle float64) Geometry {
	complex.geometryList = func(geomes []Geometry) []Geometry {
		var combiGeomes []Geometry
		for _, geome := range geomes  {
			combiGeomes = append(combiGeomes, geome.(Transformation).Rotate(rpos, angle))
		}
		return combiGeomes
	}(complex.geometryList)
	return complex
}
//翻转
func Flip() {

}

func (complex ComplexGraph)Draw(drawFunc func(x, y int)) {
	for _,geometry := range complex.geometryList {
		geometry.Draw(drawFunc)
	}
}
//计算所有复合图形最小x和最大y中间点
func (complex ComplexGraph)CenterPoint() Point {
	return Pt(0,0)
}