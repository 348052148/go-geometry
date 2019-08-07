package main

import (
	"image/png"
	"os"
	"test/image/paintboard"
	"test/image/geometry"
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
	linear := geometry.NewLineSegment([2]int{200,200}, [2]int{300,400})
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
		).AddGeometry(geometry.NewLineSegment([2]int{300,300},[2]int{500,500})).AddGeometry(
			geometry.NewRectangle(350,375, 100,50),
		).Rotate(geometry.Pt(400,400), 45),
		paintboard.NewPigment(125,0,125),
		)
	//将画好的图形保持至png文件
	png.Encode(file, paintBrush.Palette)
}