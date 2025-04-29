package main

import (
	"fmt"
	"math"
)

// Shape описывает фигуры с методами площади и периметра
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Solid описывает 3D-фигуры с методами объема и площади поверхности
type Solid interface {
	Volume() float64
	SurfaceArea() float64
}

// Circle - структура описывающая круг через его радиус
type Circle struct {
	R float64
}

// Rectangle - структура, описывающая прямоугольник
type Rectangle struct {
	W, H float64
}

// Box - структура параллелепипеда, построенная на прямоугольнике
type Box struct {
	Rectangle
	Height float64
}

// Cylinder - структура цилиндра, построенная на круге
type Cylinder struct {
	Circle
	Height float64
}

func (P Box) Volume() float64 {
	return P.Rectangle.Area() * P.Height
}
func (P Box) SurfaceArea() float64 {
	return 2 * (P.Area() + P.H*P.Height + P.Height*P.W)
}
func (C Cylinder) Volume() float64 {
	return C.Circle.Area() * C.Height
}
func (C Cylinder) SurfaceArea() float64 {
	return 2*C.Circle.Area() + C.Circle.Perimeter()*C.Height
}

func (R Rectangle) Area() float64 {
	return R.W * R.H
}

func (C Circle) Area() float64 {
	return math.Pi * math.Pow(C.R, 2)
}

func (R Rectangle) Perimeter() float64 {
	return (R.W + R.H) * 2
}

func (C Circle) Perimeter() float64 {
	return 2 * math.Pi * C.R
}
func PrintShapeInfo(shapes []Shape) {
	var area, perim float64
	name := ""
	for _, shape := range shapes {
		switch option := shape.(type) {
		case Circle:
			area = option.Area()
			perim = option.Perimeter()
			name = "круг"
		case Rectangle:
			area = option.Area()
			perim = option.Perimeter()
			name = "прямоугольник"
		}
		fmt.Printf("Эта фигура - %s.\nЕе периметр: %.2f\n Площадь: %.2f\n------------\n", name, perim, area)
	}
}
func Print3DShapeInfo(shapes []Shape) {
	var area, volume float64
	name := ""
	for _, shape := range shapes {
		switch option := shape.(type) {
		case Box:
			area = option.SurfaceArea()
			volume = option.Volume()
			name = "параллелепипед"
		case Cylinder:
			area = option.SurfaceArea()
			volume = option.Volume()
			name = "цилиндр"
		}
		fmt.Printf("Эта фигура - %s.\nЕе объем: %.2f\n площадь поверхности: %.2f\n------------\n", name, volume, area)
	}

}
func Name(s Shape) {
	switch s.(type) {
	case Circle:
		fmt.Println("Это круг")
	case Rectangle:
		fmt.Println("Это прямоугольник")
	case Box:
		fmt.Println("Это параллелепипед")
	case Cylinder:
		fmt.Println("Это цилиндр")
	default:
		fmt.Println("Неизвестная фигура")
	}
}

func main() {
	circle := Circle{
		R: 3,
	}
	rectangle := Rectangle{
		W: 10,
		H: 15,
	}
	box := Box{
		Rectangle: Rectangle{
			W: 10,
			H: 15,
		},
		Height: 20,
	}
	cylinder := Cylinder{
		Circle: Circle{
			R: 3,
		},
		Height: 10,
	}
	shapes3D := []Shape{box, cylinder}
	Print3DShapeInfo(shapes3D)

	shapes2D := []Shape{circle, rectangle}
	PrintShapeInfo(shapes2D)
}
