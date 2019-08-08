package main

import (
	"image/png"
	"os"
	"test/image/paintboard"
	"test/image/geometry"
	"fmt"
)

func main() {
	file, err := os.Create("rect.png")
	if err != nil {
		panic(err)
	}
	//创建一个画板
	borad := paintboard.NewDrawBoard(600,600)
	//从画板中获取画笔（这里貌似不符合现实世界逻辑
	paintBrush := borad.GetPaintBrush()
	//创建一个矩形的结构
	rectangle := geometry.NewRectangle(200,200, 100,200)
	//使用画笔 用 红色的颜料 将顺时针旋转90度的矩形给画出来
	paintBrush.Draw(rectangle.Rotate(rectangle.CenterPoint(),90), paintboard.NewPigment(255,0,0))
	//创建一个线段结构
	linear := geometry.NewLineSegment(geometry.Pt(200,200),geometry.Pt(300,400))
	//使用画笔 用 绿色颜料 将顺时针旋转45度线段画出来
	paintBrush.Draw(linear, paintboard.NewPigment(0,255,0))
	//把线段的两个端点作为圆心用蓝色颜料画出来
	paintBrush.Draw(geometry.NewCircular(linear.StartPoint, 10, 1),paintboard.NewPigment(0,0,255))
	paintBrush.Draw(geometry.NewCircular(linear.EndPoint, 10, 1),paintboard.NewPigment(0,0,255))
	//公式 二次函数
	f := geometry.NewFormula([2]float64{-100 , 100})
	paintBrush.Draw(f.AddFormula(func(x float64) (int, int) {
		y := 0.01 * x * x + 10
		return int(x), int(y)
	}).Define(geometry.Pt(200,200)), paintboard.NewPigment(255,0,0))
	//反比例函数
	paintBrush.Draw(f.AddFormula(func(x float64) (int, int) {
		y := 200 / x + 10
		return int(x), int(y)
	}).Define(geometry.Pt(200,200)), paintboard.NewPigment(255,0,0))
	//复合图形
	paintBrush.Draw(geometry.ComplexGraph{}.AddGeometry(
		geometry.NewCircular(geometry.Pt(400,400),20,1),
		).AddGeometry(geometry.NewLineSegment(geometry.Pt(300,300),geometry.Pt(500,500))).AddGeometry(
			geometry.NewRectangle(350,375, 100,50),
		).Rotate(geometry.Pt(400,400), 45),
		paintboard.NewPigment(125,0,125),
		)
	//两线段交点
	line01 := geometry.NewLineSegment(geometry.Pt(100,100),geometry.Pt(300,300))
	line02 := geometry.NewLineSegment(geometry.Pt(300,100),geometry.Pt(100,300))
	paintBrush.Draw(line01, paintboard.NewPigment(200,145,0))
	paintBrush.Draw(line02, paintboard.NewPigment(0,145,200))
	for _, pos  := range line01.CrossPointAsLineSegment(line02) {
		paintBrush.Draw(geometry.NewCircular(pos, 5, 1), paintboard.NewPigment(100,200,100))
	}
	//画 多边形 以及判断点是否在多边形内
	py := geometry.NewPolygon([]geometry.Point{
		//geometry.Pt(100,100),
		geometry.Pt(100,200),
		geometry.Pt(210,200),
		geometry.Pt(200,100),
	})
	p :=  geometry.Pt(285,130)
	paintBrush.Draw(py, paintboard.NewPigment(255,255,0))
	paintBrush.Draw(geometry.NewCircular(p, 5, 1), paintboard.NewPigment(0,255,255))
	fmt.Println(py.IsPolygobScope(p))
	//矩形
	rect := geometry.NewRectangle(100,100, 200,200)
	paintBrush.Draw(rect.Rotate(rect.CenterPoint(), 0), paintboard.NewPigment(164, 20, 240))
	fmt.Println(rect.Rotate(rect.CenterPoint(), 0).(geometry.Polygon).IsPolygobScope(p))
	//矩形平移
	paintBrush.Draw(rect.Translation(20, 90), paintboard.NewPigment(134, 252, 10   ))
	//将画好的图形保持至png文件
	png.Encode(file, paintBrush.Palette)
}