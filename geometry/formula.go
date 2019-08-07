package geometry


type FormulaFunc func(float64) (int, int)

type Formula struct {
	ranges [2]float64
	formulaFuncList []FormulaFunc
	originPoint Point
}

func NewFormula(ranges [2]float64) Formula  {
	return Formula{
		ranges:ranges,
		originPoint:Pt(0,0),
	}
}

func (f Formula)Draw(drawFunc func(x, y int)) {
	for _, ff := range f.formulaFuncList {
		for i := f.ranges[0]; i < f.ranges[1]; i+=0.1 {
			x, y := ff(i)
			drawFunc(f.originPoint.X + x, f.originPoint.Y + y)
		}
	}
}

func Translation()  {

}

func (f Formula)AddFormula(ff FormulaFunc) Formula  {
	f.formulaFuncList = append(f.formulaFuncList, ff)
	return f
}

//定义原点
func (f Formula)Define(originPoint Point) Formula {
	f.originPoint = originPoint
	return f
}

func (f Formula)CenterPoint() Point {
	return Pt(0, 0)
}
